package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___letf2(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32 {
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	v8 = int32(1)
	v12 = l1 & int64(9223372036854775807)
	v13 = int64(9223090561878065152)
	if v12 == v13 {
		v17 = base.B2i32(l0 != int64(0))
	} else {
		v17 = base.B2i32(base.Ui64(v13) < base.Ui64(v12))
	}
	if v17 != 0 {
		v66 = v8
		return v66
	} else {
		v21 = l3 & int64(9223372036854775807)
		v22 = int64(9223090561878065152)
		if v21 == v22 {
			v26 = base.B2i32(l2 != int64(0))
		} else {
			v26 = base.B2i32(base.Ui64(v22) < base.Ui64(v21))
		}
		if v26 != 0 {
			v66 = v8
			return v66
		} else {
			if base.B2i32(l2|l0|(v21|v12) == int64(0)) == int32(0) {
				if l3&l1 < int64(0) {
					if l1 == l3 {
						v56 = base.B2i32(base.Ui64(l2) < base.Ui64(l0))
					} else {
						v56 = base.B2i32(l3 < l1)
					}
					if v56 == int32(0) {
						v66 = base.B2i32(l0^l2|(l1^l3) != int64(0))
						return v66
					} else {
						return int32(-1)
					}
				} else {
					if l1 == l3 {
						v42 = base.B2i32(base.Ui64(l0) < base.Ui64(l2))
					} else {
						v42 = base.B2i32(l1 < l3)
					}
					if v42 == int32(0) {
						return base.B2i32(l0^l2|(l1^l3) != int64(0))
					} else {
						return int32(-1)
					}
				}
			} else {
				return int32(0)
			}
		}
	}
}
func F___lock(m *base.Module, l0 int32) {
	return
}
func F___lseek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v13 = m.Wasi_snapshot_preview1.Fd_seek(m, l0, l1, l2&int32(255), v7+int32(8))
	mBase = m.M
	if v13 != 0 {
		v15 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v13
	} else {
	}
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	m.G0 = v7 + int32(16)
	if v13 != 0 {
		v24 = int64(-1)
	} else {
		v24 = v19
	}
	return v24
}
func F_lazyFreeEvalScripts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_freeEvalScripts(m, v5, v8, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeEvalScripts[0]))
		v15 = v6 + v7
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeEvalScripts[0])) = v14 - v15
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeEvalScripts[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeEvalScripts[1])) = v20 + v15
		return
	}
}
func F_lazyFreeFunctionsCtx(m *base.Module, l0 int32) {
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v9 = v7 + v8
	F_functionsLibCtxFree(m, v5, int32(0), v4)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeFunctionsCtx[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeFunctionsCtx[0])) = v15 - v9
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_lazyFreeFunctionsCtx[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyFreeFunctionsCtx[1])) = v9 + v20
		return
	}
}
func F_lazyfreeFreeObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_decrRefCount(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeObject[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeObject[0])) = v7 + int32(-1)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeObject[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_lazyfreeFreeObject[1])) = v13 + int32(1)
		return
	}
}
func F_lazyfreeGetFreeEffort(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v54 int64
	_ = v54
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(304)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = v12 & int32(255)
	switch v14 + int32(-34) {
	case 0:
		v23 = F_objectGetVal(m, l1)
		mBase = m.M
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		v96 = v24 + v25
		m.G0 = v10 + int32(304)
		return v96
	case 1:
		switch v12&int32(15) + int32(-5) {
		case 0:
			v91 = F_moduleGetFreeEffort(m, l0, l1, l2)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				if v91 != 0 {
					v94 = v91
				} else {
					v94 = int32(-1)
				}
				v96 = v94
				m.G0 = v10 + int32(304)
				return v96
			}
		case 1:
			v39 = F_objectGetVal(m, l1)
			mBase = m.M
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			if v42 == int32(0) {
				v89 = v41
				v96 = base.I32_wrap_i64(v89)
				m.G0 = v10 + int32(304)
				return v96
			} else {
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
				if v45 == int64(0) {
					v89 = v41
					v96 = base.I32_wrap_i64(v89)
					m.G0 = v10 + int32(304)
					return v96
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(128)
					v54 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = v54
					*(*int64)(unsafe.Add(mBase, uint32(v10)+296)) = v54
					*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = int64(137438953472)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v10 + int32(168)
					v67 = int32(0)
					v69 = F_raxSeek(m, v10, int32(_a_F_lazyfreeGetFreeEffort_0), v67, v67)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = F_raxNext(m, v10)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							if v73 == int32(0) {
								F__serverAssert(m, int32(_a_F_lazyfreeGetFreeEffort_1), int32(_a_F_lazyfreeGetFreeEffort_2), int32(168))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
								v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
								F_raxStop(m, v10)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = v79*(v81+int64(1)) + v41
									v96 = base.I32_wrap_i64(v89)
									m.G0 = v10 + int32(304)
									return v96
								}
							}
						}
					}
				}
			}
		default:
			v96 = int32(1)
			m.G0 = v10 + int32(304)
			return v96
		}
	case 2:
		v30 = F_objectGetVal(m, l1)
		mBase = m.M
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
		v96 = v31 + v32
		m.G0 = v10 + int32(304)
		return v96
	default:
		if v14 == int32(115) {
			v27 = F_objectGetVal(m, l1)
			mBase = m.M
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			v96 = v29
			m.G0 = v10 + int32(304)
			return v96
		} else {
			if v14 != int32(145) {
				switch v12&int32(15) + int32(-5) {
				case 0:
					v91 = F_moduleGetFreeEffort(m, l0, l1, l2)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						if v91 != 0 {
							v94 = v91
						} else {
							v94 = int32(-1)
						}
						v96 = v94
						m.G0 = v10 + int32(304)
						return v96
					}
				case 1:
					v39 = F_objectGetVal(m, l1)
					mBase = m.M
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					if v42 == int32(0) {
						v89 = v41
						v96 = base.I32_wrap_i64(v89)
						m.G0 = v10 + int32(304)
						return v96
					} else {
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
						if v45 == int64(0) {
							v89 = v41
							v96 = base.I32_wrap_i64(v89)
							m.G0 = v10 + int32(304)
							return v96
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(128)
							v54 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+12)) = v54
							*(*int64)(unsafe.Add(mBase, uint32(v10)+296)) = v54
							*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = int64(137438953472)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v10 + int32(24)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v10 + int32(168)
							v67 = int32(0)
							v69 = F_raxSeek(m, v10, int32(_a_F_lazyfreeGetFreeEffort_0), v67, v67)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								v73 = F_raxNext(m, v10)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 == int32(0) {
										F__serverAssert(m, int32(_a_F_lazyfreeGetFreeEffort_1), int32(_a_F_lazyfreeGetFreeEffort_2), int32(168))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
										v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
										F_raxStop(m, v10)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											v89 = v79*(v81+int64(1)) + v41
											v96 = base.I32_wrap_i64(v89)
											m.G0 = v10 + int32(304)
											return v96
										}
									}
								}
							}
						}
					}
				default:
					v96 = int32(1)
					m.G0 = v10 + int32(304)
					return v96
				}
			} else {
				v21 = F_objectGetVal(m, l1)
				mBase = m.M
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				v96 = v22
				m.G0 = v10 + int32(304)
				return v96
			}
		}
	}
}
func F_lazyfreeGetPendingObjectsCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_lazyfreeGetPendingObjectsCount[0]))
	return v2
}
func F_libraryUnlink(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = F_dictGetIterator(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L6
L3:
	;
	F__serverAssert(m, int32(_a_F_libraryUnlink_0), int32(_a_F_libraryUnlink_1), int32(335))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	F__serverAssert(m, int32(_a_F_libraryUnlink_2), int32(_a_F_libraryUnlink_1), int32(323))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L74
	}
L5:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L40
	}
L6:
	;
	v21 = v7 + int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v117 == int32(0) {
		goto L5
	} else {
		goto L34
	}
L9:
	;
	v28 = v21
	v29 = v25
	goto L12
L10:
	;
	v25 = int32(1)
	goto L9
L11:
	;
	v25 = int32(0)
	goto L9
L12:
	;
	switch v29 {
	case 0:
		goto L17
	default:
		goto L16
	}
L14:
	;
	v29 = int32(0)
	goto L12
L15:
	;
	goto L8
L16:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v109
	if v109 == int32(0) {
		goto L14
	} else {
		goto L33
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v33 != int32(-1) {
		v72 = v33
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = int32(1)
	v74 = v72 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v74
	v76 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v80+int32(26)))))
	if v84 == int32(255) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v37 != 0 {
		v72 = int32(-1)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v39 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v66 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+16)))
	v47 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+27)))
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+12)))
	v50 = int64(*(*int8)(unsafe.Add(mBase, uint32(v38)+26)))
	v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+4)))
	v52 = F_wangHash64(m, v51)
	mBase = m.M
	v54 = F_wangHash64(m, v50+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v49+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v48+v56)
	mBase = m.M
	v60 = F_wangHash64(m, v47+v58)
	mBase = m.M
	v62 = F_wangHash64(m, v46+v60)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v65 = v64
	goto L21
L23:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)))
	v44 = v42 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+24)) = uint16(v44)
	v65 = v38
	goto L21
L24:
	;
	v72 = v66 + int32(-1)
	goto L18
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v72 = v69
	goto L18
L26:
	;
	v99 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v79+v97<<(uint(v99)%32)+int32(4))))
	v28 = v104 + v98<<(uint(v99)%32)
	v29 = int32(1)
	goto L12
L27:
	;
	v88 = v76
	goto L29
L28:
	;
	v88 = v73 << (uint(v84) % 32)
	goto L29
L29:
	;
	if v74 < v88 {
		v97 = v80
		v98 = v74
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if v80 != 0 {
		v117 = v76
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	if v90 == int32(-1) {
		v117 = v76
		goto L15
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v97 = int32(1)
	v98 = int32(0)
	goto L26
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v113
	v117 = v109
	goto L15
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	v127 = F_objectGetVal(m, v126)
	mBase = m.M
	v128 = F_dictDelete(m, v124, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v128 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123+int32(-8))))
	goto L38
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v140 = F_scriptingEngineCallGetFunctionMemoryOverhead(m, v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v142 - (v140 + (v132&int32(2147483647) + int32(8)))
	goto L6
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v150 = F_dictUnlink(m, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = int32(0)
	goto L42
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_dictFreeUnlinkedEntry(m, v155, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-8))))
	goto L44
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v172 = v165 + int32(-1)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	v175 = v173 & int32(7)
	switch v175 {
	case 0:
		goto L51
	case 1:
		v181 = int32(4)
		goto L46
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		goto L47
	}
L45:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v213 = v206 + int32(-1)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v216 = v214 & int32(7)
	switch v216 {
	case 0:
		goto L64
	case 1:
		v222 = int32(4)
		goto L59
	case 2:
		goto L63
	case 3:
		goto L62
	case 4:
		goto L61
	default:
		goto L60
	}
L46:
	;
	switch v175 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		v201 = int32(0)
		goto L52
	}
L47:
	;
	v181 = int32(1)
	goto L46
L48:
	;
	v181 = int32(18)
	goto L46
L49:
	;
	v181 = int32(10)
	goto L46
L50:
	;
	v181 = int32(6)
	goto L46
L51:
	;
	v176 = F_zmalloc_usable_size(m, v172)
	mBase = m.M
	v205 = v176
	goto L45
L52:
	;
	v205 = v181 + v201
	goto L45
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v165+int32(-9))))
	v201 = v200
	goto L52
L54:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v165+int32(-5))))
	v205 = v181 + v196
	goto L45
L55:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165+int32(-3)))))
	v205 = v181 + v192
	goto L45
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+int32(-2)))))
	v205 = v181 + v188
	goto L45
L57:
	;
	v205 = v181 + int32(base.Ui32(v173)>>(uint(int32(3))%32))
	goto L45
L58:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v247 - (v246 + (v205 + (v160&int32(2147483647) + int32(8))))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	goto L71
L59:
	;
	switch v216 {
	case 0:
		goto L70
	case 1:
		goto L69
	case 2:
		goto L68
	case 3:
		goto L67
	case 4:
		goto L66
	default:
		v242 = int32(0)
		goto L65
	}
L60:
	;
	v222 = int32(1)
	goto L59
L61:
	;
	v222 = int32(18)
	goto L59
L62:
	;
	v222 = int32(10)
	goto L59
L63:
	;
	v222 = int32(6)
	goto L59
L64:
	;
	v217 = F_zmalloc_usable_size(m, v213)
	mBase = m.M
	v246 = v217
	goto L58
L65:
	;
	v246 = v222 + v242
	goto L58
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(-9))))
	v242 = v241
	goto L65
L67:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(-5))))
	v246 = v222 + v237
	goto L58
L68:
	;
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206+int32(-3)))))
	v246 = v222 + v233
	goto L58
L69:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+int32(-2)))))
	v246 = v222 + v229
	goto L58
L70:
	;
	v246 = v222 + int32(base.Ui32(v214)>>(uint(int32(3))%32))
	goto L58
L71:
	;
	v255 = F_dictFetchValue(m, v252, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v255 == int32(0) {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v259 + int32(-1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+16))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v263 - (v265 + v266)
	return
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_listen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v3 = int32(0)
	v7 = m.Env.X__syscall_listen(m, l0, l1, v3, v3, v3, v3)
	mBase = m.M
	if base.Ui32(v7) < base.Ui32(int32(-4095)) {
		v15 = v7
	} else {
		v10 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0) - v7
		v15 = int32(-1)
	}
	return v15
}
func F_listenerByType(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v7 = F_connectionByType(m, l0)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = l0*int32(88) + int32(_a_F_listenerByType_0)
		} else {
			v11 = int32(0)
		}
		return v11
	}
}
func F_llex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
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
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
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
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v363 int64
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v475 int64
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v739 int64
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
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
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	goto L3
L1:
	;
	m.G0 = v11 + int32(176)
	return v1101
L2:
	;
	F_read_numeral(m, l0, l1)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L19
	} else {
		goto L265
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = v24
	goto L6
L4:
	;
	F_save(m, l0, int32(46))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L19
	} else {
		goto L242
	}
L5:
	;
	goto L4
L6:
	;
	switch v28 + int32(1) {
	case 0:
		v1101 = int32(287)
		goto L1
	default:
		goto L9
	case 10, 12, 13, 33:
		goto L8
	case 11, 14:
		goto L18
	case 35, 40:
		goto L11
	case 46:
		goto L17
	case 47:
		goto L5
	case 61:
		goto L14
	case 62:
		goto L15
	case 63:
		goto L13
	case 92:
		goto L16
	case 127:
		goto L12
	}
L8:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	*(*int32)(unsafe.Add(mBase, uint32(v1001))) = v1002 + int32(-1)
	if v1002 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L9:
	;
	if base.Ui32(v28+int32(-48)) <= base.Ui32(int32(9)) {
		goto L2
	} else {
		goto L212
	}
L10:
	;
	F_read_long_string(m, l0, int32(0), v79)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L19
	} else {
		goto L211
	}
L11:
	;
	F_save(m, l0, v28)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L19
	} else {
		goto L87
	}
L12:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v249 + int32(-1)
	if v249 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L13:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v211 + int32(-1)
	if v211 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L14:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v173 + int32(-1)
	if v173 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L15:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v135 + int32(-1)
	if v135 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	v120 = F_skip_sep(m, l0)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L43
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v41 + int32(-1)
	if v41 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	F_inclinenumber(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L3
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v54
	v57 = int32(45)
	if v54 != v57 {
		v1101 = v57
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v52 = F_luaZ_fill(m, v40)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v47 + int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v54 = v51
	goto L21
L24:
	;
	v54 = v52
	goto L21
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61 + int32(-1)
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v74
	if v74 != int32(91) {
		v87 = v74
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v72 = F_luaZ_fill(m, v60)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v67 + int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v74 = v71
	goto L26
L29:
	;
	v74 = v72
	goto L26
L30:
	;
	v91 = v87
	goto L34
L31:
	;
	v79 = F_skip_sep(m, l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = int32(0)
	if base.Ui32(int32(1)) < base.Ui32(v79) {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = v86
	goto L30
L34:
	;
	v97 = v91 + int32(1)
	if base.Ui32(int32(14)) < base.Ui32(v97) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v105 + int32(-1)
	if v105 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if int32(1)<<(uint(v97)%32)&int32(18433) != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v117 = F_luaZ_fill(m, v104)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L19
	} else {
		goto L41
	}
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v111 + int32(1)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v115
	v91 = v115
	goto L34
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v117
	v91 = v117
	goto L34
L42:
	;
	v127 = int32(91)
	if v120 != 0 {
		v1101 = v127
		goto L1
	} else {
		goto L46
	}
L43:
	;
	if base.Ui32(v120) < base.Ui32(int32(2)) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_read_long_string(m, l0, l1, v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	v1101 = int32(286)
	goto L1
L46:
	;
	v128 = m.G3
	F_luaX_lexerror(m, l0, v128+int32(_a_F_llex_0), int32(286))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	v1101 = v127
	goto L1
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v148
	v151 = int32(61)
	if v148 != v151 {
		v1101 = v151
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v146 = F_luaZ_fill(m, v134)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L51
	}
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v141 + int32(1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v148 = v145
	goto L48
L51:
	;
	v148 = v146
	goto L48
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v155 + int32(-1)
	if v155 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v168
	v1101 = int32(280)
	goto L1
L54:
	;
	v166 = F_luaZ_fill(m, v154)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L56
	}
L55:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v161 + int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v168 = v165
	goto L53
L56:
	;
	v168 = v166
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v186
	if v186 == int32(61) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v184 = F_luaZ_fill(m, v172)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L19
	} else {
		goto L60
	}
L59:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v179 + int32(1)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v186 = v183
	goto L57
L60:
	;
	v186 = v184
	goto L57
L61:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v193 + int32(-1)
	if v193 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v1101 = int32(60)
	goto L1
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v206
	v1101 = int32(282)
	goto L1
L64:
	;
	v204 = F_luaZ_fill(m, v192)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L19
	} else {
		goto L66
	}
L65:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v199 + int32(1)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	v206 = v203
	goto L63
L66:
	;
	v206 = v204
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v224
	if v224 == int32(61) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v222 = F_luaZ_fill(m, v210)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L70
	}
L69:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v217 + int32(1)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v224 = v221
	goto L67
L70:
	;
	v224 = v222
	goto L67
L71:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v231 + int32(-1)
	if v231 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v1101 = int32(62)
	goto L1
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v244
	v1101 = int32(281)
	goto L1
L74:
	;
	v242 = F_luaZ_fill(m, v230)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L19
	} else {
		goto L76
	}
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+4)) = v237 + int32(1)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v244 = v241
	goto L73
L76:
	;
	v244 = v242
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v262
	if v262 == int32(61) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v260 = F_luaZ_fill(m, v248)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L19
	} else {
		goto L80
	}
L79:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = v255 + int32(1)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v262 = v259
	goto L77
L80:
	;
	v262 = v260
	goto L77
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v269 + int32(-1)
	if v269 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v1101 = int32(126)
	goto L1
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v282
	v1101 = int32(283)
	goto L1
L84:
	;
	v280 = F_luaZ_fill(m, v268)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L19
	} else {
		goto L86
	}
L85:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v275 + int32(1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v282 = v279
	goto L83
L86:
	;
	v282 = v280
	goto L83
L87:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v289 + int32(-1)
	if v289 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v302
	if v302 == v28 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v300 = F_luaZ_fill(m, v288)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L19
	} else {
		goto L91
	}
L90:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v295 + int32(1)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v302 = v299
	goto L88
L91:
	;
	v302 = v300
	goto L88
L92:
	;
	F_save(m, l0, v28)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L19
	} else {
		goto L200
	}
L93:
	;
	v310 = v302
	goto L94
L94:
	;
	switch v310 + int32(1) {
	case 0:
		goto L101
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13:
		goto L98
	case 11, 14:
		goto L100
	default:
		goto L99
	}
L95:
	;
	goto L92
L96:
	;
	if v837 != v28 {
		v310 = v837
		goto L94
	} else {
		goto L199
	}
L97:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v567 + int32(-1)
	if v567 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L98:
	;
	F_save(m, l0, v310)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L19
	} else {
		goto L134
	}
L99:
	;
	if v310 == int32(92) {
		goto L97
	} else {
		goto L133
	}
L100:
	;
	v429 = v11 + int32(96)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v432 = v430 + int32(16)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	switch v436 + int32(-61) {
	case 0:
		goto L120
	default:
		goto L118
	case 3:
		goto L119
	}
L101:
	;
	v317 = v11 + int32(96)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v320 = v318 + int32(16)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	switch v324 + int32(-61) {
	case 0:
		goto L105
	default:
		goto L103
	case 3:
		goto L104
	}
L102:
	;
	v398 = m.G3
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v398 + int32(_a_F_llex_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(96)
	v412 = F_luaO_pushfstring(m, v399, v398+int32(_a_F_llex_2), v11+int32(16))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L19
	} else {
		goto L114
	}
L103:
	;
	v351 = m.G3
	v354 = F_strcspn(m, v320, v351+int32(_a_F_llex_3))
	mBase = m.M
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_llex[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(104)))) = uint16(v361)
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_llex[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v317))) = v363
	v366 = int32(63)
	if base.Ui32(v354) < base.Ui32(v366) {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v336 = v318 + int32(17)
	v337 = F_strlen(m, v336)
	mBase = m.M
	v338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v338)
	v341 = int32(72)
	if base.Ui32(v337) <= base.Ui32(v341) {
		v349 = v336
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v329 = F_strncpy(m, v317, v318+int32(17), int32(80))
	mBase = m.M
	v333 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329+int32(79)))) = uint8(v333)
	goto L102
L106:
	;
	v350 = F_strcat(m, v317, v349)
	mBase = m.M
	goto L102
L107:
	;
	v343 = F_strlen(m, v317)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v317+v343))) = int32(3026478)
	v349 = v336 + (v337 - v341)
	goto L106
L108:
	;
	v380 = F_strlen(m, v317)
	mBase = m.M
	v381 = v317 + v380
	v382 = m.G3
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+uint32(_c_F_llex[2]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v381))) = uint16(v385)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+uint32(_c_F_llex[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v381+int32(2)))) = uint8(v391)
	goto L102
L109:
	;
	v378 = F_strcat(m, v317, v320)
	mBase = m.M
	goto L108
L110:
	;
	v368 = v354
	goto L112
L111:
	;
	v368 = v366
	goto L112
L112:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v368))))
	if v370 == int32(0) {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v373 = F_strncat(m, v317, v320, v368)
	mBase = m.M
	v374 = F_strlen(m, v373)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v373+v374))) = int32(3026478)
	goto L108
L114:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v398 + int32(_a_F_llex_4)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v412
	v421 = F_luaO_pushfstring(m, v414, v398+int32(_a_F_llex_5), v11)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_luaD_throw(m, v423, int32(3))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v427
	goto L96
L117:
	;
	v510 = m.G3
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v510 + int32(_a_F_llex_1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(96)
	v524 = F_luaO_pushfstring(m, v511, v510+int32(_a_F_llex_2), v11+int32(48))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L19
	} else {
		goto L129
	}
L118:
	;
	v463 = m.G3
	v466 = F_strcspn(m, v432, v463+int32(_a_F_llex_3))
	mBase = m.M
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v463)+uint32(_c_F_llex[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(104)))) = uint16(v473)
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v463)+uint32(_c_F_llex[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v429))) = v475
	v478 = int32(63)
	if base.Ui32(v466) < base.Ui32(v478) {
		goto L125
	} else {
		goto L126
	}
L119:
	;
	v448 = v430 + int32(17)
	v449 = F_strlen(m, v448)
	mBase = m.M
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v429))) = uint8(v450)
	v453 = int32(72)
	if base.Ui32(v449) <= base.Ui32(v453) {
		v461 = v448
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v441 = F_strncpy(m, v429, v430+int32(17), int32(80))
	mBase = m.M
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v441+int32(79)))) = uint8(v445)
	goto L117
L121:
	;
	v462 = F_strcat(m, v429, v461)
	mBase = m.M
	goto L117
L122:
	;
	v455 = F_strlen(m, v429)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v429+v455))) = int32(3026478)
	v461 = v448 + (v449 - v453)
	goto L121
L123:
	;
	v492 = F_strlen(m, v429)
	mBase = m.M
	v493 = v429 + v492
	v494 = m.G3
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v494)+uint32(_c_F_llex[2]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v493))) = uint16(v497)
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+uint32(_c_F_llex[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v493+int32(2)))) = uint8(v503)
	goto L117
L124:
	;
	v490 = F_strcat(m, v429, v432)
	mBase = m.M
	goto L123
L125:
	;
	v480 = v466
	goto L127
L126:
	;
	v480 = v478
	goto L127
L127:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432+v480))))
	if v482 == int32(0) {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v485 = F_strncat(m, v429, v432, v480)
	mBase = m.M
	v486 = F_strlen(m, v485)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v485+v486))) = int32(3026478)
	goto L123
L129:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_save(m, l0, int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L19
	} else {
		goto L130
	}
L130:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v524
	v538 = F_luaO_pushfstring(m, v526, v510+int32(_a_F_llex_5), v11+int32(32))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_luaD_throw(m, v540, int32(3))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v544
	goto L96
L133:
	;
	goto L98
L134:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v550 + int32(-1)
	if v550 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v563
	v837 = v563
	goto L96
L136:
	;
	v561 = F_luaZ_fill(m, v549)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L19
	} else {
		goto L138
	}
L137:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = v556 + int32(1)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v563 = v560
	goto L135
L138:
	;
	v563 = v561
	goto L135
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v580
	switch v580 + int32(-97) {
	case 0:
		v816 = int32(7)
		goto L143
	case 1:
		goto L144
	case 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 18, 20:
		goto L145
	case 5:
		goto L152
	case 13:
		goto L151
	case 17:
		goto L150
	case 19:
		goto L149
	case 21:
		goto L148
	default:
		goto L153
	}
L140:
	;
	v578 = F_luaZ_fill(m, v566)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L19
	} else {
		goto L142
	}
L141:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v566)+4)) = v573 + int32(1)
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	v580 = v577
	goto L139
L142:
	;
	v580 = v578
	goto L139
L143:
	;
	F_save(m, l0, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L19
	} else {
		goto L195
	}
L144:
	;
	v816 = int32(8)
	goto L143
L145:
	;
	v600 = v580 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v600) {
		goto L158
	} else {
		goto L159
	}
L146:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v598
	goto L96
L147:
	;
	F_save(m, l0, int32(10))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L19
	} else {
		goto L154
	}
L148:
	;
	v816 = int32(11)
	goto L143
L149:
	;
	v816 = int32(9)
	goto L143
L150:
	;
	v816 = int32(13)
	goto L143
L151:
	;
	v816 = int32(10)
	goto L143
L152:
	;
	v816 = int32(12)
	goto L143
L153:
	;
	switch v580 + int32(1) {
	case 0:
		goto L146
	default:
		goto L145
	case 11, 14:
		goto L147
	}
L154:
	;
	F_inclinenumber(m, l0)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L19
	} else {
		goto L155
	}
L155:
	;
	goto L146
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v637
	if base.Ui32(int32(9)) < base.Ui32(v637+int32(-48)) {
		v809 = v600
		goto L166
	} else {
		goto L167
	}
L157:
	;
	v635 = F_luaZ_fill(m, v603)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L19
	} else {
		goto L165
	}
L158:
	;
	F_save(m, l0, v580)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L19
	} else {
		goto L161
	}
L159:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	*(*int32)(unsafe.Add(mBase, uint32(v603))) = v604 + int32(-1)
	if v604 == int32(0) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v610 + int32(1)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	v637 = v614
	goto L156
L161:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v618 + int32(-1)
	if v618 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v631 = F_luaZ_fill(m, v617)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L19
	} else {
		goto L164
	}
L163:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v617)+4)) = v624 + int32(1)
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v628
	v837 = v628
	goto L96
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v631
	v837 = v631
	goto L96
L165:
	;
	v637 = v635
	goto L156
L166:
	;
	F_save(m, l0, v809)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L19
	} else {
		goto L194
	}
L167:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v645 + int32(-1)
	if v645 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v662 = int32(-48)
	v663 = v600*int32(10) + v637 + v662
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v661
	if base.Ui32(int32(9)) < base.Ui32(v661+v662) {
		v809 = v663
		goto L166
	} else {
		goto L172
	}
L169:
	;
	v659 = F_luaZ_fill(m, v644)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L19
	} else {
		goto L171
	}
L170:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v644)+4)) = v654 + int32(1)
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	v661 = v658
	goto L168
L171:
	;
	v661 = v659
	goto L168
L172:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v670 + int32(-1)
	v678 = v663*int32(10) + v661 + int32(-48)
	if v670 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v688
	if base.Ui32(v678) < base.Ui32(int32(256)) {
		v809 = v678
		goto L166
	} else {
		goto L177
	}
L174:
	;
	v686 = F_luaZ_fill(m, v669)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L19
	} else {
		goto L176
	}
L175:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+4)) = v681 + int32(1)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v688 = v685
	goto L173
L176:
	;
	v688 = v686
	goto L173
L177:
	;
	v693 = v11 + int32(96)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v696 = v694 + int32(16)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	switch v700 + int32(-61) {
	case 0:
		goto L181
	default:
		goto L179
	case 3:
		goto L180
	}
L178:
	;
	v774 = m.G3
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v774 + int32(_a_F_llex_6)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(96)
	v788 = F_luaO_pushfstring(m, v775, v774+int32(_a_F_llex_2), v11+int32(80))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L19
	} else {
		goto L190
	}
L179:
	;
	v727 = m.G3
	v730 = F_strcspn(m, v696, v727+int32(_a_F_llex_3))
	mBase = m.M
	v737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+uint32(_c_F_llex[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(104)))) = uint16(v737)
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v727)+uint32(_c_F_llex[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v693))) = v739
	v742 = int32(63)
	if base.Ui32(v730) < base.Ui32(v742) {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	v712 = v694 + int32(17)
	v713 = F_strlen(m, v712)
	mBase = m.M
	v714 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v693))) = uint8(v714)
	v717 = int32(72)
	if base.Ui32(v713) <= base.Ui32(v717) {
		v725 = v712
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v705 = F_strncpy(m, v693, v694+int32(17), int32(80))
	mBase = m.M
	v709 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v705+int32(79)))) = uint8(v709)
	goto L178
L182:
	;
	v726 = F_strcat(m, v693, v725)
	mBase = m.M
	goto L178
L183:
	;
	v719 = F_strlen(m, v693)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v693+v719))) = int32(3026478)
	v725 = v712 + (v713 - v717)
	goto L182
L184:
	;
	v756 = F_strlen(m, v693)
	mBase = m.M
	v757 = v693 + v756
	v758 = m.G3
	v761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_llex[2]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v757))) = uint16(v761)
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758)+uint32(_c_F_llex[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v757+int32(2)))) = uint8(v767)
	goto L178
L185:
	;
	v754 = F_strcat(m, v693, v696)
	mBase = m.M
	goto L184
L186:
	;
	v744 = v730
	goto L188
L187:
	;
	v744 = v742
	goto L188
L188:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v744))))
	if v746 == int32(0) {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v749 = F_strncat(m, v693, v696, v744)
	mBase = m.M
	v750 = F_strlen(m, v749)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v749+v750))) = int32(3026478)
	goto L184
L190:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_save(m, l0, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L19
	} else {
		goto L191
	}
L191:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v795
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v788
	v802 = F_luaO_pushfstring(m, v790, v774+int32(_a_F_llex_5), v11+int32(64))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L19
	} else {
		goto L192
	}
L192:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_luaD_throw(m, v804, int32(3))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L19
	} else {
		goto L193
	}
L193:
	;
	v809 = v678
	goto L166
L194:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v837 = v814
	goto L96
L195:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	*(*int32)(unsafe.Add(mBase, uint32(v819))) = v820 + int32(-1)
	if v820 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v833 = F_luaZ_fill(m, v819)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L19
	} else {
		goto L198
	}
L197:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v819)+4)) = v826 + int32(1)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v830
	v837 = v830
	goto L96
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v833
	v837 = v833
	goto L96
L199:
	;
	goto L95
L200:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	*(*int32)(unsafe.Add(mBase, uint32(v852))) = v853 + int32(-1)
	if v853 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v866
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	v877 = F_luaS_newlstr(m, v869, v871+int32(1), v874+int32(-2))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L19
	} else {
		goto L205
	}
L202:
	;
	v864 = F_luaZ_fill(m, v852)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L19
	} else {
		goto L204
	}
L203:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v852)+4)) = v859 + int32(1)
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	v866 = v863
	goto L201
L204:
	;
	v866 = v864
	goto L201
L205:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)+4))
	v881 = F_luaH_setstr(m, v869, v880, v877)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L19
	} else {
		goto L207
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v877
	v1101 = int32(286)
	goto L1
L207:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v881)+8))
	if v883 != 0 {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v884 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v881)+8)) = v884
	*(*int32)(unsafe.Add(mBase, uint32(v881))) = v884
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v869)+16))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)+68))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v888)+64))
	if base.Ui32(v889) < base.Ui32(v890) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	F_luaC_step(m, v869)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L19
	} else {
		goto L210
	}
L210:
	;
	goto L206
L211:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v900)+4)) = int32(0)
	goto L3
L212:
	;
	if v28 == int32(95) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	*(*int32)(unsafe.Add(mBase, uint32(v984))) = v985 + int32(-1)
	if v985 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L214:
	;
	v918 = v28
	goto L217
L215:
	;
	if base.Ui32(int32(25)) < base.Ui32(v28|int32(32)+int32(-97)) {
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	F_save(m, l0, v918)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L19
	} else {
		goto L219
	}
L218:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v956)+4))
	v959 = F_luaS_newlstr(m, v955, v957, v958)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L19
	} else {
		goto L227
	}
L219:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	*(*int32)(unsafe.Add(mBase, uint32(v925))) = v926 + int32(-1)
	if v926 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v939
	goto L224
L221:
	;
	v937 = F_luaZ_fill(m, v925)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L19
	} else {
		goto L223
	}
L222:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v925)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v925)+4)) = v932 + int32(1)
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932))))
	v939 = v936
	goto L220
L223:
	;
	v939 = v937
	goto L220
L224:
	;
	if v939 == int32(95) {
		v918 = v939
		goto L217
	} else {
		goto L225
	}
L225:
	;
	if base.B2i32(base.Ui32(v939+int32(-48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v939|int32(32)+int32(-97)) < base.Ui32(int32(26))) != 0 {
		v918 = v939
		goto L217
	} else {
		goto L226
	}
L226:
	;
	goto L218
L227:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v961)+4))
	v963 = F_luaH_setstr(m, v955, v962, v959)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L19
	} else {
		goto L229
	}
L228:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+6)))
	if v977 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L229:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v963)+8))
	if v965 != 0 {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v966 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v963)+8)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v963))) = v966
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v955)+16))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)+68))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v970)+64))
	if base.Ui32(v971) < base.Ui32(v972) {
		goto L228
	} else {
		goto L231
	}
L231:
	;
	F_luaC_step(m, v955)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L19
	} else {
		goto L232
	}
L232:
	;
	goto L228
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v959
	v1101 = int32(285)
	goto L1
L234:
	;
	v1101 = v977 | int32(256)
	goto L1
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v998
	v1101 = v28
	goto L1
L236:
	;
	v996 = F_luaZ_fill(m, v984)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L19
	} else {
		goto L238
	}
L237:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v984)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v984)+4)) = v991 + int32(1)
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991))))
	v998 = v995
	goto L235
L238:
	;
	v998 = v996
	goto L235
L239:
	;
	v1014 = F_luaZ_fill(m, v1001)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L19
	} else {
		goto L241
	}
L240:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1001)+4)) = v1008 + int32(1)
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1012
	v28 = v1012
	goto L6
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1014
	v28 = v1014
	goto L6
L242:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1021 + int32(-1)
	if v1021 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1034
	v1038 = v1034 & int32(255)
	if v1038 == int32(46) {
		goto L248
	} else {
		goto L249
	}
L244:
	;
	v1032 = F_luaZ_fill(m, v1020)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L19
	} else {
		goto L246
	}
L245:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1020)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1020)+4)) = v1027 + int32(1)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027))))
	v1034 = v1031
	goto L243
L246:
	;
	v1034 = v1032
	goto L243
L247:
	;
	if base.Ui32(v1034+int32(-48)) <= base.Ui32(int32(9)) {
		goto L2
	} else {
		goto L264
	}
L248:
	;
	F_save(m, l0, v1034)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L19
	} else {
		goto L251
	}
L249:
	;
	if v1038 != 0 {
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)))
	*(*int32)(unsafe.Add(mBase, uint32(v1043))) = v1044 + int32(-1)
	if v1044 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1057
	v1061 = v1057 & int32(255)
	if v1061 == int32(46) {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	v1055 = F_luaZ_fill(m, v1043)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L19
	} else {
		goto L255
	}
L254:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+4)) = v1050 + int32(1)
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
	v1057 = v1054
	goto L252
L255:
	;
	v1057 = v1055
	goto L252
L256:
	;
	F_save(m, l0, v1057)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L19
	} else {
		goto L259
	}
L257:
	;
	if v1061 == int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1101 = int32(278)
	goto L1
L259:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	*(*int32)(unsafe.Add(mBase, uint32(v1069))) = v1070 + int32(-1)
	if v1070 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1083
	v1101 = int32(279)
	goto L1
L261:
	;
	v1081 = F_luaZ_fill(m, v1069)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L19
	} else {
		goto L263
	}
L262:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1069)+4)) = v1076 + int32(1)
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076))))
	v1083 = v1080
	goto L260
L263:
	;
	v1083 = v1081
	goto L260
L264:
	;
	v1101 = int32(46)
	goto L1
L265:
	;
	v1101 = int32(284)
	goto L1
}
func F_llrint(m *base.Module, l0 float64) int64 {
	var v2 float64
	_ = v2
	var v8 int64
	_ = v8
	v2 = base.F64_nearest(l0)
	if base.F64_lt(base.F64_abs(v2), float64(9.223372036854776e+18)) == int32(0) {
		return int64(-9223372036854775807 - 1)
	} else {
		v8 = base.I64_trunc_f64_s(v2)
		return v8
	}
}
func F_lmoveHandlePush(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	if l2 != 0 {
		v22 = l2
		v26 = int32(0)
		F_listTypeTryConversionRaw(m, v22, int32(1), v9+int32(8), v26, v26, v26, v26)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			F_listTypePush(m, v32, l3, l4)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				F_signalModifiedKey(m, l0, v35, l1)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if l4 != 0 {
						v41 = int32(_a_F_lmoveHandlePush_0)
					} else {
						v41 = int32(_a_F_lmoveHandlePush_1)
					}
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
					F_notifyKeyspaceEvent(m, int32(16), v41, l1, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						F_addReplyBulk(m, l0, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v13 = F_createListListpackObject(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v13
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			F_dbAdd(m, v16, l1, v9+int32(12))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v22 = v21
				v26 = int32(0)
				F_listTypeTryConversionRaw(m, v22, int32(1), v9+int32(8), v26, v26, v26, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					F_listTypePush(m, v32, l3, l4)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						F_signalModifiedKey(m, l0, v35, l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if l4 != 0 {
								v41 = int32(_a_F_lmoveHandlePush_0)
							} else {
								v41 = int32(_a_F_lmoveHandlePush_1)
							}
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
							F_notifyKeyspaceEvent(m, int32(16), v41, l1, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								F_addReplyBulk(m, l0, v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
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
func F_lmpopGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = int32(1)
	v9 = F_genericGetKeys(m, int32(0), v6, int32(2), v6, l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_loadSingleAppendOnlyFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v415 int64
	_ = v415
	var v421 int32
	_ = v421
	var v433 int64
	_ = v433
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v439 int32
	_ = v439
	var v446 int64
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int64
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int64
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v569 int32
	_ = v569
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v603 int64
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int64
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int64
	_ = v696
	var v697 int32
	_ = v697
	var v698 int64
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int64
	_ = v707
	var v720 int64
	_ = v720
	var v721 int64
	_ = v721
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v742 int64
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v761 int64
	_ = v761
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v814 int64
	_ = v814
	var v817 int64
	_ = v817
	var v826 int32
	_ = v826
	var v835 int32
	_ = v835
	var v845 int64
	_ = v845
	var v849 int64
	_ = v849
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v915 int64
	_ = v915
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v963 int32
	_ = v963
	var v985 int32
	_ = v985
	var v994 int32
	_ = v994
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1025 int64
	_ = v1025
	var v1036 int64
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int64
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1059 int32
	_ = v1059
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	v21 = m.G0
	v23 = v21 - int32(1344)
	m.G0 = v23
	v25 = int32(_a_F_loadSingleAppendOnlyFile_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[0]))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[1]))
	v29 = F_makePath(m, v28, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v1124 + int32(1344)
	return v1129
L2:
	;
	F_sdsfree(m, v1104)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L4
	} else {
		goto L275
	}
L3:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if int32(-1) < v78 {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	return int32(0)
L5:
	;
	v34 = F_fopen(m, v29, int32(_a_F_loadSingleAppendOnlyFile_1))
	mBase = m.M
	if v34 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(9116376)
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[2]))
	v41 = F___fstatat(m, int32(-100), v29, v23+int32(1248), int32(0))
	mBase = m.M
	goto L10
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v60 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v47 = int32(3)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if v47 < v49 {
		v1102 = v23
		v1104 = v29
		v1107 = v47
		goto L2
	} else {
		goto L13
	}
L10:
	;
	if v41 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[2]))
	if v44 == int32(44) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v52 = F___strerror_l(m, v36, v36)
	mBase = m.M
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_2), v23)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v1102 = v23
	v1104 = v29
	v1107 = v47
	goto L2
L16:
	;
	F_sdsfree(m, v29)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v63 = int32(44)
	v64 = F___strerror_l(m, v63, v63)
	mBase = m.M
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_3), v23+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v1124 = v23
	v1129 = int32(1)
	goto L1
L21:
	;
	v116 = int32(_a_F_loadSingleAppendOnlyFile_0)
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[0])) = v117
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[4]))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[5]))
	v123 = int64(0)
	v125 = F_createClient(m, v117)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L36
	}
L22:
	;
	if int32(-1) < v95 {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	if int32(-1) < v87 {
		v95 = v87
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v82 = F___lockfile(m, v34)
	mBase = m.M
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	if v82 == int32(0) {
		v87 = v83
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v87 = v81
	goto L23
L26:
	;
	F___unlockfile(m, v34)
	mBase = m.M
	v87 = v83
	goto L23
L27:
	;
	goto L22
L28:
	;
	v91 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(8)
	v95 = int32(-1)
	goto L27
L29:
	;
	if v105 == int32(-1) {
		goto L21
	} else {
		goto L32
	}
L30:
	;
	v104 = F___fstatat(m, v95, int32(_a_F_loadSingleAppendOnlyFile_4), v23+int32(1248), int32(4096))
	mBase = m.M
	v105 = v104
	goto L29
L31:
	;
	v101 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v105 = v101
	goto L29
L32:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v23)+1272))
	if v108 != int64(0) {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v111 = F_fclose(m, v34)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_sdsfree(m, v29)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v1124 = v23
	v1129 = int32(2)
	goto L1
L36:
	;
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v125)+200)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v125)+328)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v125))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+204)) = int32(268439552)
	*(*int64)(unsafe.Add(mBase, uint32(v125+int32(208)))) = v127
	F_initClientReplicationData(m, v125)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v125)+104))
	v142 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142
	v144 = int32(_a_F_loadSingleAppendOnlyFile_0)
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[4])) = v125
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[5])) = v125
	v149 = v125 + int32(200)
	v154 = F_fread(m, v23+int32(1242), int32(1), v142, v34)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L51
	}
L38:
	;
	v1095 = int32(_a_F_loadSingleAppendOnlyFile_0)
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[5])) = v122
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[4])) = v120
	v1099 = F_fclose(m, v34)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L4
	} else {
		goto L274
	}
L39:
	;
	v1073 = F_freeClient(m, v125)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L273
	}
L40:
	;
	v1036 = F___ftello(m, v34)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L4
	} else {
		goto L269
	}
L41:
	;
	v1015 = int32(4)
	if v125 != 0 {
		v1059 = v1015
		goto L39
	} else {
		goto L268
	}
L42:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v985 {
		goto L41
	} else {
		goto L266
	}
L43:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[6]))
	if v857 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L44:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v826 {
		goto L232
	} else {
		goto L233
	}
L45:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if int32(-1) < v770 {
		goto L223
	} else {
		goto L224
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+24)) = v744
	F_freeClientArgv(m, v125)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L219
	}
L47:
	;
	v742 = int64(0)
	v757 = v123
	v760 = v742
	v761 = v742
	goto L45
L48:
	;
	v421 = v23 + int32(208) | int32(1)
	v433 = v415
	v436 = int64(0)
	v437 = v415
	v439 = int32(0)
	goto L123
L49:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[7]))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v311 == int32(0) {
		v334 = v310
		v335 = v311
		goto L90
	} else {
		goto L91
	}
L50:
	;
	v298 = int64(0)
	v300 = int32(0)
	v302 = F_fseek(m, v34, v300, v300)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L87
	}
L51:
	;
	if v154 != int32(6) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v159 = v23 + int32(1242)
	v160 = int32(_a_F_loadSingleAppendOnlyFile_5)
	v161 = int32(6)
	goto L57
L53:
	;
	if v225 == int32(0) {
		goto L49
	} else {
		goto L69
	}
L54:
	;
	v225 = int32(0)
	goto L53
L55:
	;
	v197 = v192
	v198 = v193
	v199 = v194
	goto L65
L56:
	;
	if v182 == int32(0) {
		goto L54
	} else {
		goto L63
	}
L57:
	;
	if (v160|v159)&int32(3) != 0 {
		v192 = v159
		v193 = v160
		v194 = v161
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v169 = v159
	v170 = v160
	v171 = v161
	goto L59
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v174 != v175 {
		v192 = v169
		v193 = v170
		v194 = v171
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L56
L61:
	;
	v177 = int32(4)
	v178 = v170 + v177
	v180 = v169 + v177
	v182 = v171 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v182) {
		v169 = v180
		v170 = v178
		v171 = v182
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v192 = v180
	v193 = v178
	v194 = v182
	goto L55
L64:
	;
	v225 = v202 - v203
	goto L53
L65:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v202 != v203 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v205 = int32(1)
	v210 = v199 + int32(-1)
	if v210 == int32(0) {
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v197 = v197 + v205
	v198 = v198 + v205
	v199 = v210
	goto L65
L69:
	;
	v229 = v23 + int32(1242)
	v230 = int32(_a_F_loadSingleAppendOnlyFile_6)
	v231 = int32(6)
	goto L74
L70:
	;
	if v295 == int32(0) {
		goto L49
	} else {
		goto L86
	}
L71:
	;
	v295 = int32(0)
	goto L70
L72:
	;
	v267 = v262
	v268 = v263
	v269 = v264
	goto L82
L73:
	;
	if v252 == int32(0) {
		goto L71
	} else {
		goto L80
	}
L74:
	;
	if (v230|v229)&int32(3) != 0 {
		v262 = v229
		v263 = v230
		v264 = v231
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v239 = v229
	v240 = v230
	v241 = v231
	goto L76
L76:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	if v244 != v245 {
		v262 = v239
		v263 = v240
		v264 = v241
		goto L72
	} else {
		goto L78
	}
L77:
	;
	goto L73
L78:
	;
	v247 = int32(4)
	v248 = v240 + v247
	v250 = v239 + v247
	v252 = v241 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v252) {
		v239 = v250
		v240 = v248
		v241 = v252
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v262 = v250
	v263 = v248
	v264 = v252
	goto L72
L81:
	;
	v295 = v272 - v273
	goto L70
L82:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v272 != v273 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v275 = int32(1)
	v280 = v269 + int32(-1)
	if v280 == int32(0) {
		goto L71
	} else {
		goto L85
	}
L85:
	;
	v267 = v267 + v275
	v268 = v268 + v275
	v269 = v280
	goto L82
L86:
	;
	goto L50
L87:
	;
	if v302 != int32(-1) {
		v415 = v123
		goto L48
	} else {
		goto L88
	}
L88:
	;
	v757 = v123
	v760 = v298
	v761 = v298
	goto L45
L89:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(2) < v340 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	v338 = v335 - v334&int32(255)
	goto L89
L91:
	;
	if v311 != v310&int32(255) {
		v334 = v310
		v335 = v311
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v317 = l0
	v318 = v307
	goto L93
L93:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if v322 == int32(0) {
		v334 = v321
		v335 = v322
		goto L90
	} else {
		goto L95
	}
L94:
	;
	v334 = v321
	v335 = v322
	goto L90
L95:
	;
	v325 = int32(1)
	if v322 == v321&int32(255) {
		v317 = v317 + v325
		v318 = v318 + v325
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v350 = int32(0)
	v352 = F_fseek(m, v34, v350, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L103
	}
L98:
	;
	if v338 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v346 = int32(_a_F_loadSingleAppendOnlyFile_7)
	goto L101
L100:
	;
	v346 = int32(_a_F_loadSingleAppendOnlyFile_8)
	goto L101
L101:
	;
	F__serverLog(m, int32(2), v346, int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	goto L97
L103:
	;
	if v352 == int32(-1) {
		goto L47
	} else {
		goto L104
	}
L104:
	;
	v360 = F___memcpy(m, v23+int32(208), int32(_a_F_loadSingleAppendOnlyFile_9), int32(80))
	mBase = m.M
	v361 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v360)+56)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v360)+48)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v360+int32(64)))) = v361
	v370 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360+int32(72)))) = uint8(v370)
	goto L105
L105:
	;
	v376 = F_rdbLoadRio(m, v23+int32(208), int32(1), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	v394 = F___ftello(m, v34)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L115
	}
L107:
	;
	if v376 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v381 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v1059 = int32(4)
	goto L39
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = l0
	if v338 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v388 = int32(_a_F_loadSingleAppendOnlyFile_10)
	goto L113
L112:
	;
	v388 = int32(_a_F_loadSingleAppendOnlyFile_11)
	goto L113
L113:
	;
	F__serverLog(m, int32(3), v388, v23+int32(192))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	goto L109
L115:
	;
	v396 = int32(_a_F_loadSingleAppendOnlyFile_0)
	*(*int64)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[8])) = v394
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9]))
	v400 = F_zmalloc_used_memory(m)
	mBase = m.M
	if base.Ui32(v400) <= base.Ui32(v399) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v338 != 0 {
		v415 = v394
		goto L48
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v403 = F_zmalloc_used_memory(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9])) = v403
	goto L117
L119:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(2) < v406 {
		v415 = v394
		goto L48
	} else {
		goto L120
	}
L120:
	;
	F__serverLog(m, int32(2), int32(_a_F_loadSingleAppendOnlyFile_12), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v415 = v394
	goto L48
L122:
	;
	F__serverAssert(m, int32(_a_F_loadSingleAppendOnlyFile_13), int32(_a_F_loadSingleAppendOnlyFile_14), int32(1698))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L218
	}
L123:
	;
	if v439&int32(1023) != 0 {
		v466 = v433
		goto L125
	} else {
		goto L126
	}
L124:
	;
	F__serverAssert(m, int32(_a_F_loadSingleAppendOnlyFile_15), int32(_a_F_loadSingleAppendOnlyFile_14), int32(1695))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L217
	}
L125:
	;
	v471 = F_fgets(m, v23+int32(208), int32(1024), v34)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L134
	}
L126:
	;
	v446 = F___ftello(m, v34)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v449 = int32(_a_F_loadSingleAppendOnlyFile_0)
	v451 = *(*int64)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[8])) = v451 + (v446 - v433)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9]))
	v456 = F_zmalloc_used_memory(m)
	mBase = m.M
	if base.Ui32(v456) <= base.Ui32(v455) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	F_processEventsWhileBlocked(m)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v459 = F_zmalloc_used_memory(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9])) = v459
	goto L129
L131:
	;
	F_processModuleLoadingProgressEvent(m, int32(1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v466 = v446
	goto L125
L133:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+208)))
	switch v496 + int32(-35) {
	case 0:
		v720 = v436
		v721 = v437
		goto L144
	default:
		goto L42
	case 7:
		goto L145
	}
L134:
	;
	if v471 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if int32(-1) < v475 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	if int32(base.Ui32(v484)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v757 = v466
		v760 = v436
		v761 = v437
		goto L45
	} else {
		goto L141
	}
L137:
	;
	goto L136
L138:
	;
	v479 = F___lockfile(m, v34)
	mBase = m.M
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v479 == int32(0) {
		v484 = v480
		goto L137
	} else {
		goto L140
	}
L139:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v484 = v478
	goto L137
L140:
	;
	F___unlockfile(m, v34)
	mBase = m.M
	v484 = v480
	goto L137
L141:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v493&int32(8) != 0 {
		v814 = v466
		v817 = v436
		goto L44
	} else {
		goto L142
	}
L142:
	;
	v1022 = int32(0)
	v1025 = v466
	goto L40
L143:
	;
	goto L124
L144:
	;
	v433 = v466
	v436 = v720
	v437 = v721
	v439 = v439 + int32(1)
	goto L123
L145:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+209)))
	if v499 == int32(0) {
		v757 = v466
		v760 = v436
		v761 = v437
		goto L45
	} else {
		goto L146
	}
L146:
	;
	v505 = v421
	goto L148
L147:
	;
	if base.Ui32(v550+int32(-1073741824)) < base.Ui32(int32(-1073741823)) {
		goto L42
	} else {
		goto L162
	}
L148:
	;
	v510 = v505 + int32(1)
	v511 = int32(*(*int8)(unsafe.Add(mBase, uint32(v505))))
	v512 = F___isspace_1(m, v511)
	mBase = m.M
	if v512 != 0 {
		v505 = v510
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v513 = int32(1)
	switch v511&int32(255) + int32(-43) {
	case 0:
		v519 = v513
		goto L152
	default:
		v521 = v505
		v522 = v511
		v523 = v513
		goto L151
	case 2:
		goto L153
	}
L150:
	;
	goto L149
L151:
	;
	v526 = v522 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v526) {
		v544 = int32(0)
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v520 = int32(*(*int8)(unsafe.Add(mBase, uint32(v510))))
	v521 = v510
	v522 = v520
	v523 = v519
	goto L151
L153:
	;
	v519 = int32(0)
	goto L152
L154:
	;
	if v523 != 0 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v530 = int32(0)
	v531 = v521
	v532 = v526
	goto L156
L156:
	;
	v534 = int32(10)
	v536 = v530*v534 - v532
	v537 = int32(*(*int8)(unsafe.Add(mBase, uint32(v531)+1)))
	v541 = v537 + int32(-48)
	if base.Ui32(v541) < base.Ui32(v534) {
		v530 = v536
		v531 = v531 + int32(1)
		v532 = v541
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v544 = v536
	goto L154
L158:
	;
	goto L157
L159:
	;
	v550 = int32(0) - v544
	goto L161
L160:
	;
	v550 = v544
	goto L161
L161:
	;
	goto L147
L162:
	;
	v557 = F_valkey_malloc(m, v550<<(uint(int32(2))%32))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+28)) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v125)+24)) = v550
	v569 = int32(0)
	goto L164
L164:
	;
	v586 = F_fgets(m, v23+int32(208), int32(1024), v34)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L168
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = int32(0)
	v636 = F_lookupCommand(m, v557, v550)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L184
	}
L166:
	;
	v603 = F_strtox_2(m, v421, int32(0), int32(10), int64(2147483648))
	mBase = m.M
	v604 = base.I32_wrap_i64(v603)
	goto L173
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+24)) = v569
	F_freeClientArgv(m, v125)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L171
	}
L168:
	;
	if v586 == int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+208)))
	if v590&int32(255) == int32(36) {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	goto L167
L171:
	;
	if v586 == int32(0) {
		v757 = v466
		v760 = v436
		v761 = v437
		goto L45
	} else {
		goto L172
	}
L172:
	;
	goto L42
L173:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[10]))
	v607 = F_sdsnewlen(m, v606, v604)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	if v604 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v620 = F_createObject(m, int32(0), v607)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L180
	}
L176:
	;
	v612 = F_fread(m, v607, v604, int32(1), v34)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	if v612 != 0 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	F_sdsfree(m, v607)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v744 = v569
	goto L46
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v557+v569<<(uint(int32(2))%32)))) = v620
	v623 = int32(1)
	v624 = v569 + v623
	v629 = F_fread(m, v23+int32(208), int32(2), v623, v34)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	if v629 == int32(0) {
		v744 = v624
		goto L46
	} else {
		goto L182
	}
L182:
	;
	if v624 != v550 {
		v569 = v624
		goto L164
	} else {
		goto L183
	}
L183:
	;
	goto L165
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+68)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(v125)+72)) = v636
	if v636 != 0 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v636)+48))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v670&int32(8) == int32(0) {
		goto L199
	} else {
		goto L200
	}
L186:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v651 {
		goto L193
	} else {
		goto L194
	}
L187:
	;
	v648 = F_commandCheckArity(m, v636, v550, v23+int32(204))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L191
	}
L188:
	;
	v642 = F_commandCheckExistence(m, v125, v23+int32(204))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	if v642 == int32(0) {
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L185
L191:
	;
	if v648 != 0 {
		goto L185
	} else {
		goto L192
	}
L192:
	;
	goto L186
L193:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
	F_sdsfree(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L4
	} else {
		goto L196
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = l0
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v655
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_16), v23+int32(176))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	F_freeClientArgv(m, v125)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	v1059 = int32(4)
	goto L39
L198:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v125)+180))
	if v684 != 0 {
		goto L143
	} else {
		goto L204
	}
L199:
	;
	m.T0[v669].(func(*base.Module, int32))(m, v125)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L4
	} else {
		goto L203
	}
L200:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v125)+68))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)+48))
	if v676 == int32(17) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v636)+56))
	F_queueMultiCommand(m, v125, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	goto L198
L204:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v125)+132))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+20))
	if v686 != 0 {
		goto L143
	} else {
		goto L205
	}
L205:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v687&int32(16) != 0 {
		goto L122
	} else {
		goto L206
	}
L206:
	;
	F_freeClientArgv(m, v125)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[6]))
	if v693 == int32(0) {
		v698 = v437
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[11]))
	if v700 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v696 = F___ftello(m, v34)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	v698 = v696
	goto L208
L211:
	;
	if v669 == int32(18) {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	F_debugDelay(m, v700)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	v707 = v437
	goto L216
L215:
	;
	v707 = v436
	goto L216
L216:
	;
	v720 = v707
	v721 = v698
	goto L144
L217:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	v757 = v466
	v760 = v436
	v761 = v437
	goto L45
L220:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v800&int32(8) == int32(0) {
		v845 = v757
		v849 = v761
		goto L43
	} else {
		goto L231
	}
L221:
	;
	if int32(base.Ui32(v779)>>(uint(int32(4))%32))&int32(1) != 0 {
		goto L220
	} else {
		goto L226
	}
L222:
	;
	goto L221
L223:
	;
	v774 = F___lockfile(m, v34)
	mBase = m.M
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v774 == int32(0) {
		v779 = v775
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v779 = v773
	goto L222
L225:
	;
	F___unlockfile(m, v34)
	mBase = m.M
	v779 = v775
	goto L222
L226:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v786 {
		goto L41
	} else {
		goto L227
	}
L227:
	;
	goto L228
L228:
	;
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[2]))
	v791 = F___strerror_l(m, v790, v790)
	mBase = m.M
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_17), v23+int32(32))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	goto L41
L231:
	;
	v814 = v757
	v817 = v760
	goto L44
L232:
	;
	v845 = v814
	v849 = v817
	goto L43
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_18), v23+int32(144))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	v952 = int32(4)
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v954 {
		v1059 = v952
		goto L39
	} else {
		goto L264
	}
L236:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v861 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	if v849 == int64(-1) {
		goto L245
	} else {
		goto L246
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_19), v23+int32(128))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v872 {
		goto L237
	} else {
		goto L240
	}
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_20), v23+int32(112))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	goto L237
L242:
	;
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_21), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L263
	}
L243:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[12]))
	if v910 == int32(-1) {
		goto L253
	} else {
		goto L254
	}
L244:
	;
	goto L250
L245:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if v894 <= int32(3) {
		goto L242
	} else {
		goto L249
	}
L246:
	;
	v885 = F_truncate(m, v29, v849)
	mBase = m.M
	if v885 != int32(-1) {
		goto L243
	} else {
		goto L247
	}
L247:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if v889 <= int32(3) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v1059 = int32(4)
	goto L39
L249:
	;
	v1059 = int32(4)
	goto L39
L250:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[2]))
	v900 = F___strerror_l(m, v899, v899)
	mBase = m.M
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_22), v23+int32(64))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	goto L235
L253:
	;
	v934 = int32(5)
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if int32(3) < v936 {
		v1022 = v934
		v1025 = v845
		goto L40
	} else {
		goto L261
	}
L254:
	;
	v915 = F___lseek(m, v910, int64(0), int32(2))
	mBase = m.M
	if v915 != int64(-1) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[3]))
	if v919 <= int32(3) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	goto L258
L257:
	;
	v1059 = int32(4)
	goto L39
L258:
	;
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[2]))
	v925 = F___strerror_l(m, v924, v924)
	mBase = m.M
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v925
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_23), v23+int32(96))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	goto L235
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_24), v23+int32(80))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	v1022 = v934
	v1025 = v845
	goto L40
L263:
	;
	goto L235
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_25), v23+int32(48))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	v1059 = v952
	goto L39
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = l0
	F__serverLog(m, int32(3), int32(_a_F_loadSingleAppendOnlyFile_26), v23+int32(160))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	goto L41
L268:
	;
	v1076 = v23
	v1078 = v29
	v1081 = v1015
	goto L38
L269:
	;
	v1039 = int32(_a_F_loadSingleAppendOnlyFile_0)
	v1041 = *(*int64)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[8])) = v1041 + (v1036 - v1025)
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9]))
	v1046 = F_zmalloc_used_memory(m)
	mBase = m.M
	if base.Ui32(v1046) <= base.Ui32(v1045) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[0])) = v26
	v1059 = v1022
	goto L39
L271:
	;
	goto L270
L272:
	;
	v1049 = F_zmalloc_used_memory(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_loadSingleAppendOnlyFile[9])) = v1049
	goto L271
L273:
	;
	v1076 = v23
	v1078 = v29
	v1081 = v1059
	goto L38
L274:
	;
	v1102 = v1076
	v1104 = v1078
	v1107 = v1081
	goto L2
L275:
	;
	v1124 = v1102
	v1129 = v1107
	goto L1
}
func F_locking_getc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v5 = l0 + int32(76)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v7 != 0 {
		v9 = v7
	} else {
		v9 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
	if v7 == int32(0) {
	} else {
	}
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 == v15 {
		v21 = F___uflow(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = v21
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
			if v27&int32(1073741824) == int32(0) {
			} else {
				v35 = F_emscripten_futex_wake(m, v5, int32(1))
				mBase = m.M
			}
			return v25
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 + int32(1)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
		v25 = v20
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
		if v27&int32(1073741824) == int32(0) {
		} else {
			v35 = F_emscripten_futex_wake(m, v5, int32(1))
			mBase = m.M
		}
		return v25
	}
}
func F_locking_putc_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	v1 = l0
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0]))
	if v7 != 0 {
		v9 = v7
	} else {
		v9 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0])) = v9
	if v7 == int32(0) {
	} else {
	}
	v16 = v1 & int32(255)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[1]))
	if v16 == v18 {
		v32 = F___overflow(m, int32(_a_F_locking_putc_2_0), v16)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v36 = v32
			v38 = int32(0)
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0])) = v38
			if v40&int32(1073741824) == int32(0) {
			} else {
				v50 = F_emscripten_futex_wake(m, int32(_a_F_locking_putc_2_1), int32(1))
				mBase = m.M
			}
			return v36
		}
	} else {
		v20 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[2]))
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[3]))
		if v21 == v23 {
			v32 = F___overflow(m, int32(_a_F_locking_putc_2_0), v16)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = v32
				v38 = int32(0)
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0]))
				*(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0])) = v38
				if v40&int32(1073741824) == int32(0) {
				} else {
					v50 = F_emscripten_futex_wake(m, int32(_a_F_locking_putc_2_1), int32(1))
					mBase = m.M
				}
				return v36
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[2])) = v21 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v1)
			v36 = v16
			v38 = int32(0)
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_locking_putc_2[0])) = v38
			if v40&int32(1073741824) == int32(0) {
			} else {
				v50 = F_emscripten_futex_wake(m, int32(_a_F_locking_putc_2_1), int32(1))
				mBase = m.M
			}
			return v36
		}
	}
}
func F_log(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v96 float64
	_ = v96
	var v104 int32
	_ = v104
	var v108 float64
	_ = v108
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v132 int32
	_ = v132
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v146 float64
	_ = v146
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v157 float64
	_ = v157
	var v160 float64
	_ = v160
	var v164 float64
	_ = v164
	var v167 float64
	_ = v167
	var v172 float64
	_ = v172
	var v175 float64
	_ = v175
	var v183 float64
	_ = v183
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(48)) % 64)))
	v15 = base.I64_reinterpret_f64(l0)
	if base.Ui64(int64(854320534781951)) < base.Ui64(v15+int64(-4606619468846596096)) {
		if base.Ui32(int32(-32737)) < base.Ui32(v14+int32(-32752)) {
			v116 = v15
			v118 = v116 + int64(-4604367669032910848)
			v122 = base.F64_convert_i32_s(base.I32_wrap_i64(v118 >> (uint(int64(52)) % 64)))
			v123 = int32(0)
			v124 = *(*float64)(unsafe.Add(mBase, _c_F_log[0]))
			v132 = base.I32_wrap_i64(int64(base.Ui64(v118)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
			v135 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[1])))
			v136 = base.F64_add(base.F64_mul(v122, v124), v135)
			v139 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[2])))
			v146 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[3])))
			v150 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[4])))
			v152 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v116-v118&int64(-4503599627370496)), v146), v150))
			v153 = base.F64_add(v136, v152)
			v154 = base.F64_mul(v152, v152)
			v157 = *(*float64)(unsafe.Add(mBase, _c_F_log[5]))
			v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[6]))
			v164 = *(*float64)(unsafe.Add(mBase, _c_F_log[7]))
			v167 = *(*float64)(unsafe.Add(mBase, _c_F_log[8]))
			v172 = *(*float64)(unsafe.Add(mBase, _c_F_log[9]))
			v175 = *(*float64)(unsafe.Add(mBase, _c_F_log[10]))
			v183 = base.F64_add(v153, base.F64_add(base.F64_mul(base.F64_mul(v152, v154), base.F64_add(base.F64_mul(v154, base.F64_add(base.F64_mul(v152, v157), v160)), base.F64_add(base.F64_mul(v152, v164), v167))), base.F64_add(base.F64_mul(v154, v172), base.F64_add(base.F64_mul(v122, v175), base.F64_add(v152, base.F64_sub(v136, v153))))))
			return v183
		} else {
			if base.F64_ne(l0, float64(0)) != 0 {
				if v15 == int64(9218868437227405312) {
					v183 = l0
					return v183
				} else {
					if base.Ui32(int32(32767)) < base.Ui32(v14) {
						v108 = base.F64_sub(l0, l0)
						return base.F64_div(v108, v108)
					} else {
						v104 = int32(32752)
						if v14&v104 != v104 {
							v116 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15))) + int64(-234187180623265792)
							v118 = v116 + int64(-4604367669032910848)
							v122 = base.F64_convert_i32_s(base.I32_wrap_i64(v118 >> (uint(int64(52)) % 64)))
							v123 = int32(0)
							v124 = *(*float64)(unsafe.Add(mBase, _c_F_log[0]))
							v132 = base.I32_wrap_i64(int64(base.Ui64(v118)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
							v135 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[1])))
							v136 = base.F64_add(base.F64_mul(v122, v124), v135)
							v139 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[2])))
							v146 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[3])))
							v150 = *(*float64)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_log[4])))
							v152 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v116-v118&int64(-4503599627370496)), v146), v150))
							v153 = base.F64_add(v136, v152)
							v154 = base.F64_mul(v152, v152)
							v157 = *(*float64)(unsafe.Add(mBase, _c_F_log[5]))
							v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[6]))
							v164 = *(*float64)(unsafe.Add(mBase, _c_F_log[7]))
							v167 = *(*float64)(unsafe.Add(mBase, _c_F_log[8]))
							v172 = *(*float64)(unsafe.Add(mBase, _c_F_log[9]))
							v175 = *(*float64)(unsafe.Add(mBase, _c_F_log[10]))
							v183 = base.F64_add(v153, base.F64_add(base.F64_mul(base.F64_mul(v152, v154), base.F64_add(base.F64_mul(v154, base.F64_add(base.F64_mul(v152, v157), v160)), base.F64_add(base.F64_mul(v152, v164), v167))), base.F64_add(base.F64_mul(v154, v172), base.F64_add(base.F64_mul(v122, v175), base.F64_add(v152, base.F64_sub(v136, v153))))))
							return v183
						} else {
							v108 = base.F64_sub(l0, l0)
							return base.F64_div(v108, v108)
						}
					}
				}
			} else {
				v96 = F_fp_barrier_3(m, float64(-1))
				mBase = m.M
				return base.F64_div(v96, float64(0))
			}
		}
	} else {
		if v15 != int64(4607182418800017408) {
			v25 = base.F64_add(l0, float64(-1))
			v27 = base.F64_mul(v25, float64(1.34217728e+08))
			v29 = base.F64_sub(base.F64_add(v25, v27), v27)
			v31 = int32(0)
			v32 = *(*float64)(unsafe.Add(mBase, _c_F_log[11]))
			v33 = base.F64_mul(base.F64_mul(v29, v29), v32)
			v34 = base.F64_add(v25, v33)
			v35 = base.F64_mul(v25, v25)
			v36 = base.F64_mul(v25, v35)
			v38 = *(*float64)(unsafe.Add(mBase, _c_F_log[12]))
			v41 = *(*float64)(unsafe.Add(mBase, _c_F_log[13]))
			v44 = *(*float64)(unsafe.Add(mBase, _c_F_log[14]))
			v47 = *(*float64)(unsafe.Add(mBase, _c_F_log[15]))
			v53 = *(*float64)(unsafe.Add(mBase, _c_F_log[16]))
			v56 = *(*float64)(unsafe.Add(mBase, _c_F_log[17]))
			v59 = *(*float64)(unsafe.Add(mBase, _c_F_log[18]))
			v65 = *(*float64)(unsafe.Add(mBase, _c_F_log[19]))
			v68 = *(*float64)(unsafe.Add(mBase, _c_F_log[20]))
			v71 = *(*float64)(unsafe.Add(mBase, _c_F_log[21]))
			return base.F64_add(v34, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, v38), base.F64_add(base.F64_mul(v35, v41), base.F64_add(base.F64_mul(v25, v44), v47)))), base.F64_add(base.F64_mul(v35, v53), base.F64_add(base.F64_mul(v25, v56), v59)))), base.F64_add(base.F64_mul(v35, v65), base.F64_add(base.F64_mul(v25, v68), v71)))), base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(v25, v29), v32), base.F64_add(v25, v29)), base.F64_add(v33, base.F64_sub(v25, v34)))))
		} else {
			return float64(0)
		}
	}
}
func F_logCurrentClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v137 int32
	_ = v137
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
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
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
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(160)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = F_sdsempty(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = l1
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_0), v9+int32(144))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[1]))
	v28 = F_catClientInfoString(m, v24, l0, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v31 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_sdsfree(m, v28)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v28
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_1), v9+int32(128))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v55 < int32(1) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v47
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_2), v9+int32(112))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v60 = F_getArgvReprString(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v63 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_sdsfree(m, v60)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = int32(0)
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_3), v9+int32(96))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = F_objectGetVal(m, v78)
	mBase = m.M
	v80 = int32(_a_F_logCurrentClient_4)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	if v346 < int32(2) {
		goto L1
	} else {
		goto L96
	}
L23:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v346 = v342
	goto L22
L24:
	;
	if v115-v117 == int32(0) {
		goto L23
	} else {
		goto L36
	}
L25:
	;
	v115 = F_tolower(m, v111)
	mBase = m.M
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v117 = F_tolower(m, v116)
	mBase = m.M
	goto L24
L26:
	;
	v85 = v79
	v86 = v80
	v87 = v83
	goto L29
L27:
	;
	v111 = int32(0)
	v112 = v80
	goto L25
L28:
	;
	v111 = v108 & int32(255)
	v112 = v107
	goto L25
L29:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v89 == int32(0) {
		v107 = v86
		v108 = v87
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v107 = v101
	v108 = int32(0)
	goto L28
L31:
	;
	v93 = v87 & int32(255)
	if v93 == v89 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = int32(1)
	v101 = v86 + v100
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v102 != 0 {
		v85 = v85 + v100
		v86 = v101
		v87 = v102
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v95 = F_tolower(m, v93)
	mBase = m.M
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v97 = F_tolower(m, v96)
	mBase = m.M
	if v95 == v97 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v107 = v86
	v108 = v99
	goto L28
L35:
	;
	goto L30
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v123 = F_objectGetVal(m, v122)
	mBase = m.M
	v124 = int32(_a_F_logCurrentClient_5)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v159-v161 == int32(0) {
		goto L23
	} else {
		goto L49
	}
L38:
	;
	v159 = F_tolower(m, v155)
	mBase = m.M
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v161 = F_tolower(m, v160)
	mBase = m.M
	goto L37
L39:
	;
	v129 = v123
	v130 = v124
	v131 = v127
	goto L42
L40:
	;
	v155 = int32(0)
	v156 = v124
	goto L38
L41:
	;
	v155 = v152 & int32(255)
	v156 = v151
	goto L38
L42:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v133 == int32(0) {
		v151 = v130
		v152 = v131
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v151 = v145
	v152 = int32(0)
	goto L41
L44:
	;
	v137 = v131 & int32(255)
	if v137 == v133 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v144 = int32(1)
	v145 = v130 + v144
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	if v146 != 0 {
		v129 = v129 + v144
		v130 = v145
		v131 = v146
		goto L42
	} else {
		goto L48
	}
L46:
	;
	v139 = F_tolower(m, v137)
	mBase = m.M
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v141 = F_tolower(m, v140)
	mBase = m.M
	if v139 == v141 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v151 = v130
	v152 = v143
	goto L41
L48:
	;
	goto L43
L49:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v165 < int32(2) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v170 = int32(1)
	goto L51
L51:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[1]))
	if v176 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v333 = v170 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v333 < v334 {
		v170 = v333
		goto L51
	} else {
		goto L95
	}
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v220 = v170 << (uint(int32(2)) % 32)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v220)))
	v223 = F_getArgvReprString(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L64
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v180 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v170<<(uint(int32(2))%32))))
	v189 = F_objectGetVal(m, v188)
	mBase = m.M
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+int32(-1)))))
	switch v192 & int32(7) {
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
		v209 = int32(0)
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v170
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_6), v9+int32(80))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L63
	}
L58:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v189+int32(-17))))
	v209 = v208
	goto L57
L59:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v189+int32(-9))))
	v209 = v205
	goto L57
L60:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189+int32(-5)))))
	v209 = v202
	goto L57
L61:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+int32(-3)))))
	v209 = v199
	goto L57
L62:
	;
	v209 = int32(base.Ui32(v192) >> (uint(int32(3)) % 32))
	goto L57
L63:
	;
	goto L53
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v226 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_sdsfree(m, v223)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v170
	F__serverLog(m, int32(1027), int32(_a_F_logCurrentClient_3), v9+int32(64))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v239+v220)))
	v242 = F_objectGetVal(m, v241)
	mBase = m.M
	v243 = int32(_a_F_logCurrentClient_4)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v246 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v278-v280 == int32(0) {
		goto L23
	} else {
		goto L81
	}
L70:
	;
	v278 = F_tolower(m, v274)
	mBase = m.M
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v280 = F_tolower(m, v279)
	mBase = m.M
	goto L69
L71:
	;
	v248 = v242
	v249 = v243
	v250 = v246
	goto L74
L72:
	;
	v274 = int32(0)
	v275 = v243
	goto L70
L73:
	;
	v274 = v271 & int32(255)
	v275 = v270
	goto L70
L74:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == int32(0) {
		v270 = v249
		v271 = v250
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v270 = v264
	v271 = int32(0)
	goto L73
L76:
	;
	v256 = v250 & int32(255)
	if v256 == v252 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v263 = int32(1)
	v264 = v249 + v263
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v265 != 0 {
		v248 = v248 + v263
		v249 = v264
		v250 = v265
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v258 = F_tolower(m, v256)
	mBase = m.M
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v260 = F_tolower(m, v259)
	mBase = m.M
	if v258 == v260 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v270 = v249
	v271 = v262
	goto L73
L80:
	;
	goto L75
L81:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v284+v220)))
	v287 = F_objectGetVal(m, v286)
	mBase = m.M
	v288 = int32(_a_F_logCurrentClient_5)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v291 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v323-v325 == int32(0) {
		goto L23
	} else {
		goto L94
	}
L83:
	;
	v323 = F_tolower(m, v319)
	mBase = m.M
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	v325 = F_tolower(m, v324)
	mBase = m.M
	goto L82
L84:
	;
	v293 = v287
	v294 = v288
	v295 = v291
	goto L87
L85:
	;
	v319 = int32(0)
	v320 = v288
	goto L83
L86:
	;
	v319 = v316 & int32(255)
	v320 = v315
	goto L83
L87:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == int32(0) {
		v315 = v294
		v316 = v295
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v315 = v309
	v316 = int32(0)
	goto L86
L89:
	;
	v301 = v295 & int32(255)
	if v301 == v297 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v308 = int32(1)
	v309 = v294 + v308
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	if v310 != 0 {
		v293 = v293 + v308
		v294 = v309
		v295 = v310
		goto L87
	} else {
		goto L93
	}
L91:
	;
	v303 = F_tolower(m, v301)
	mBase = m.M
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v305 = F_tolower(m, v304)
	mBase = m.M
	if v303 == v305 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v315 = v294
	v316 = v307
	goto L86
L93:
	;
	goto L88
L94:
	;
	goto L53
L95:
	;
	v346 = v334
	goto L22
L96:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v353 = F_getDecodedObject(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v356 = F_objectGetVal(m, v353)
	mBase = m.M
	v357 = F_dbFind(m, v355, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L99
	}
L98:
	;
	F_decrRefCount(m, v353)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L114
	}
L99:
	;
	if v357 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v361 = int32(_a_F_logCurrentClient_7)
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[1]))
	if v364 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v385 {
		goto L98
	} else {
		goto L108
	}
L102:
	;
	if int32(3) < v362 {
		goto L98
	} else {
		goto L106
	}
L103:
	;
	if int32(3) < v362 {
		goto L98
	} else {
		goto L104
	}
L104:
	;
	F__serverLog(m, int32(3), int32(_a_F_logCurrentClient_8), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	v376 = F_objectGetVal(m, v353)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v376
	F__serverLog(m, int32(3), int32(_a_F_logCurrentClient_9), v9+int32(48))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	goto L101
L108:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v388 & int32(15)
	F__serverLog(m, int32(3), int32(_a_F_logCurrentClient_10), v9+int32(32))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v399 {
		goto L98
	} else {
		goto L110
	}
L110:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(base.Ui32(v402)>>(uint(int32(4))%32)) & int32(15)
	F__serverLog(m, int32(3), int32(_a_F_logCurrentClient_11), v9+int32(16))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_logCurrentClient[0]))
	if int32(3) < v415 {
		goto L98
	} else {
		goto L112
	}
L112:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v419 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(base.Ui32(v418) >> (uint(v419) % 32))
	F__serverLog(m, v419, int32(_a_F_logCurrentClient_12), v9)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	goto L98
L114:
	;
	goto L1
}
func F_logServerInfo(m *base.Module) {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_serverLogRaw(m, int32(1027), int32(_a_F_logServerInfo_0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
		v20 = F_createStringObject_1(m, int32(_a_F_logServerInfo_1), int32(3))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
			v31 = F_genInfoSectionDict(m, v8+int32(4), int32(1), int32(0), v8+int32(12), v8+int32(8))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v35 = F_genValkeyInfoString(m, v31, v33, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_logServerInfo[0]))
					if v38 == int32(0) {
						v66 = v35
						F_serverLogRaw(m, int32(1027), v66)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							F_serverLogRaw(m, int32(1027), int32(_a_F_logServerInfo_2))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_logServerInfo[1]))
								v80 = F_getAllClientsInfoString(m, int32(-1), v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_serverLogRaw(m, int32(1027), v80)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_sdsfree(m, v66)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_sdsfree(m, v80)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												F_releaseInfoSectionDict(m, v31)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
													F_decrRefCount(m, v90)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
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
						v41 = F_sdsempty(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = F_genClusterInfoString(m, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v45 = int32(0)
								v48 = F_clusterGenNodesDescription(m, v45, v45, v45)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v52 = F_sdscatprintf(m, v35, int32(_a_F_logServerInfo_3), int32(0))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v54 = F_sdscatsds(m, v52, v43)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											v58 = F_sdscatprintf(m, v54, int32(_a_F_logServerInfo_4), int32(0))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												v60 = F_sdscatsds(m, v58, v48)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													F_sdsfree(m, v43)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														F_sdsfree(m, v48)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return
														} else {
															v66 = v60
															F_serverLogRaw(m, int32(1027), v66)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return
															} else {
																F_serverLogRaw(m, int32(1027), int32(_a_F_logServerInfo_2))
																mBase = m.M
																v75 = m.ExcPending
																if v75 != 0 {
																	return
																} else {
																	v79 = *(*int32)(unsafe.Add(mBase, _c_F_logServerInfo[1]))
																	v80 = F_getAllClientsInfoString(m, int32(-1), v79)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		F_serverLogRaw(m, int32(1027), v80)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return
																		} else {
																			F_sdsfree(m, v66)
																			mBase = m.M
																			v85 = m.ExcPending
																			if v85 != 0 {
																				return
																			} else {
																				F_sdsfree(m, v80)
																				mBase = m.M
																				v87 = m.ExcPending
																				if v87 != 0 {
																					return
																				} else {
																					F_releaseInfoSectionDict(m, v31)
																					mBase = m.M
																					v89 = m.ExcPending
																					if v89 != 0 {
																						return
																					} else {
																						v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
																						F_decrRefCount(m, v90)
																						mBase = m.M
																						v92 = m.ExcPending
																						if v92 != 0 {
																							return
																						} else {
																							m.G0 = v8 + int32(16)
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
				}
			}
		}
	}
}
func F_log_iter_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v144 int64
	_ = v144
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 float64
	_ = v183
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v203 int64
	_ = v203
	var v221 int32
	_ = v221
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = int64(0)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v12 <= v11 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v221
L2:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if v48 <= v53 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
	if v15 < v17 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v48 = v14
	goto L2
L5:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v26 = v15 + int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v30 = v26 >> (uint(v29) % 32)
	if v21 < v30 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	v33 = v22
	goto L9
L8:
	;
	v33 = v21
	goto L9
L9:
	;
	v36 = int32(1)
	if v36 < v30 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v30
	goto L12
L11:
	;
	v39 = v36
	goto L12
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if base.I64_extend_i32_s((v22+int32(-1))&v26+v33)<<(uint(base.I64_extend_i32_u(v39+v40+int32(-1)))%64) <= v46 {
		v221 = v21
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v48 = v46
	goto L2
L14:
	;
	v221 = int32(1)
	goto L1
L15:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v179
	v183 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	if base.F64_lt(base.F64_abs(v183), float64(9.223372036854776e+18)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L16:
	;
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = v73 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	if v78 <= v75 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L15
L19:
	;
	if base.B2i32(v75 < v78) == int32(0) {
		goto L14
	} else {
		goto L36
	}
L20:
	;
	goto L19
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+64))
	if v80 == int32(0) {
		v93 = v75
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v77)+96))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v94+v93<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v98 + v100
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v77)+32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v77)+28))
	v109 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v111 = v75 >> (uint(v110) % 32)
	if v109 < v111 {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v83 = int32(0)
	v86 = v75 - v80
	if v86 < v78 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v88 = v83
	goto L26
L25:
	;
	v88 = v83 - v78
	goto L26
L26:
	;
	if v86 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = v78
	goto L29
L28:
	;
	v91 = v88
	goto L29
L29:
	;
	v93 = v91 + v86
	goto L22
L30:
	;
	v114 = v105
	goto L32
L31:
	;
	v114 = v109
	goto L32
L32:
	;
	v117 = int32(1)
	if v117 < v111 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v120 = v111
	goto L35
L34:
	;
	v120 = v117
	goto L35
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v126 = base.I64_extend_i32_s((v105+int32(-1))&v75+v114) << (uint(base.I64_extend_i32_u(v120+v121+int32(-1))) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v126
	v133 = int32(63) - (v110 + base.I32_wrap_i64(base.I64_clz(v126|v104)))
	v134 = base.I64_extend_i32_u(v133)
	v135 = v126 >> (uint(v134) % 64)
	v137 = base.I64_extend32_s(v135) << (uint(v134) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v137
	v139 = int64(1)
	v144 = v139 << (uint(base.I64_extend_i32_u(v133+base.B2i32(v103 <= base.I32_wrap_i64(v135)))) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v144>>(uint(v139)%64) + v137
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v137 + v144 + int64(-1)
	goto L20
L36:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v164 + v165
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	if v168 < v169 {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	goto L18
L38:
	;
	v192 = v180 * v191
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v195)+32))
	v203 = base.I64_extend_i32_u(int32(63) - (v196 + base.I32_wrap_i64(base.I64_clz(v192|v197))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = base.I64_extend32_s(v192>>(uint(v203)%64)) << (uint(v203) % 64)
	goto L14
L39:
	;
	v191 = int64(-9223372036854775807 - 1)
	goto L38
L40:
	;
	v189 = base.I64_trunc_f64_s(v183)
	v191 = v189
	goto L38
}
func F_lolwut5Command(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(66)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(12)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v28 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(32)
	return
L2:
	;
	v58 = int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v59 < v58 {
		v65 = v58
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = F_getLongFromObjectOrReply(m, l0, v32, v20+int32(24), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v38 < int32(3) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = F_getLongFromObjectOrReply(m, l0, v42, v20+int32(20), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 < int32(4) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v56 = F_getLongFromObjectOrReply(m, l0, v52, v20+int32(16), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L2
L13:
	;
	v69 = int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v70 < v69 {
		v76 = v69
		goto L18
	} else {
		goto L19
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v65
	v68 = v65
	goto L13
L15:
	;
	if base.Ui32(v59) < base.Ui32(int32(1001)) {
		v68 = v59
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v65 = int32(1000)
	goto L14
L17:
	;
	v80 = int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v81 < v80 {
		v87 = v80
		goto L22
	} else {
		goto L23
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v76
	v78 = v76
	goto L17
L19:
	;
	if base.Ui32(v70) < base.Ui32(int32(201)) {
		v78 = v70
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v76 = int32(200)
	goto L18
L21:
	;
	v91 = F_lwDrawSchotter(m, v68, v78, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v87
	v89 = v87
	goto L21
L23:
	;
	if base.Ui32(v81) < base.Ui32(int32(201)) {
		v89 = v81
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v87 = int32(200)
	goto L22
L25:
	;
	v93 = F_sdsempty(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v95 < int32(1) {
		v378 = v93
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut5Command[0]))
	if v392 != 0 {
		goto L110
	} else {
		goto L111
	}
L28:
	;
	v101 = v95
	v104 = int32(0)
	v105 = v93
	goto L29
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v116 < int32(1) {
		v344 = v101
		v348 = v105
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v378 = v368
	goto L27
L31:
	;
	if v104 == v344+int32(-1) {
		v367 = v344
		v368 = v348
		goto L106
	} else {
		goto L107
	}
L32:
	;
	v120 = v104 | int32(3)
	v122 = v104 | int32(2)
	v124 = v104 | int32(1)
	v128 = int32(0)
	v132 = v105
	goto L33
L33:
	;
	v143 = int32(0)
	if v128 < v143 {
		v160 = v143
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v344 = v341
	v348 = v335
	goto L31
L35:
	;
	v162 = int32(0)
	if v128 < v162 {
		v179 = v162
		goto L42
	} else {
		goto L43
	}
L36:
	;
	goto L35
L37:
	;
	v148 = int32(0)
	if v104 < v148 {
		v160 = v148
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v151 <= v128 {
		v160 = v148
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v153 <= v104 {
		v160 = v148
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v155+v128+v151*v104))))
	v160 = v159
	goto L36
L41:
	;
	v181 = int32(0)
	if v128 < v181 {
		v198 = v181
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v167 = int32(0)
	if v124 < v167 {
		v179 = v167
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v170 <= v128 {
		v179 = v167
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v172 <= v124 {
		v179 = v167
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v178 = int32(*(*int8)(unsafe.Add(mBase, uint32(v174+v128+v170*v124))))
	v179 = v178
	goto L42
L47:
	;
	v201 = v128 | int32(1)
	v202 = int32(0)
	if v201 < v202 {
		v219 = v202
		goto L54
	} else {
		goto L55
	}
L48:
	;
	goto L47
L49:
	;
	v186 = int32(0)
	if v122 < v186 {
		v198 = v186
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v189 <= v128 {
		v198 = v186
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v191 <= v122 {
		v198 = v186
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v193+v128+v189*v122))))
	v198 = v197
	goto L48
L53:
	;
	v221 = int32(0)
	if v201 < v221 {
		v238 = v221
		goto L60
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v207 = int32(0)
	if v104 < v207 {
		v219 = v207
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v210 <= v201 {
		v219 = v207
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v212 <= v104 {
		v219 = v207
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v214+v201+v210*v104))))
	v219 = v218
	goto L54
L59:
	;
	v240 = int32(0)
	if v201 < v240 {
		v257 = v240
		goto L66
	} else {
		goto L67
	}
L60:
	;
	goto L59
L61:
	;
	v226 = int32(0)
	if v124 < v226 {
		v238 = v226
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v229 <= v201 {
		v238 = v226
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v231 <= v124 {
		v238 = v226
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v237 = int32(*(*int8)(unsafe.Add(mBase, uint32(v233+v201+v229*v124))))
	v238 = v237
	goto L60
L65:
	;
	v259 = int32(0)
	if v128 < v259 {
		v276 = v259
		goto L72
	} else {
		goto L73
	}
L66:
	;
	goto L65
L67:
	;
	v245 = int32(0)
	if v122 < v245 {
		v257 = v245
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v248 <= v201 {
		v257 = v245
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v250 <= v122 {
		v257 = v245
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v252+v201+v248*v122))))
	v257 = v256
	goto L66
L71:
	;
	v278 = int32(0)
	if v201 < v278 {
		v295 = v278
		goto L78
	} else {
		goto L79
	}
L72:
	;
	goto L71
L73:
	;
	v264 = int32(0)
	if v120 < v264 {
		v276 = v264
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v267 <= v128 {
		v276 = v264
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v269 <= v120 {
		v276 = v264
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v275 = int32(*(*int8)(unsafe.Add(mBase, uint32(v271+v128+v267*v120))))
	v276 = v275
	goto L72
L77:
	;
	v297 = int32(226)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+29)) = uint8(v297)
	v300 = base.B2i32(v160 != int32(0))
	if v179 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	goto L77
L79:
	;
	v283 = int32(0)
	if v120 < v283 {
		v295 = v283
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v286 <= v201 {
		v295 = v283
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v288 <= v120 {
		v295 = v283
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v294 = int32(*(*int8)(unsafe.Add(mBase, uint32(v290+v201+v286*v120))))
	v295 = v294
	goto L78
L83:
	;
	v303 = v300 | int32(2)
	goto L85
L84:
	;
	v303 = v300
	goto L85
L85:
	;
	if v198 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v306 = v303 | int32(4)
	goto L88
L87:
	;
	v306 = v303
	goto L88
L88:
	;
	if v219 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v309 = v306 | int32(8)
	goto L91
L90:
	;
	v309 = v306
	goto L91
L91:
	;
	if v238 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v312 = v309 | int32(16)
	goto L94
L93:
	;
	v312 = v309
	goto L94
L94:
	;
	if v257 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v315 = v312 | int32(32)
	goto L97
L96:
	;
	v315 = v312
	goto L97
L97:
	;
	if v276 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v318 = v315 | int32(64)
	goto L100
L99:
	;
	v318 = v315
	goto L100
L100:
	;
	if v295 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v321 = v318 | int32(128)
	goto L103
L102:
	;
	v321 = v318
	goto L103
L103:
	;
	v325 = v321&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+31)) = uint8(v325)
	v330 = int32(base.Ui32(v321)>>(uint(int32(6))%32)) ^ int32(160)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+30)) = uint8(v330)
	v335 = F_sdscatlen(m, v132, v20+int32(29), int32(3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v338 = v128 + int32(2)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v338 < v339 {
		v128 = v338
		v132 = v335
		goto L33
	} else {
		goto L105
	}
L105:
	;
	goto L34
L106:
	;
	v370 = v104 + int32(4)
	if v370 < v367 {
		v101 = v367
		v104 = v370
		v105 = v368
		goto L29
	} else {
		goto L109
	}
L107:
	;
	v364 = F_sdscatlen(m, v348, int32(_a_F_lolwut5Command_0), int32(1))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v367 = v366
	v368 = v364
	goto L106
L109:
	;
	goto L30
L110:
	;
	v393 = int32(_a_F_lolwut5Command_1)
	goto L112
L111:
	;
	v393 = int32(_a_F_lolwut5Command_2)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v393
	v397 = F_sdscatprintf(m, v378, int32(_a_F_lolwut5Command_3), v20)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L119
	}
L113:
	;
	F_addReplyVerbatim(m, l0, v408, v429, int32(_a_F_lolwut5Command_4))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L125
	}
L114:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v408+int32(-17))))
	v429 = v428
	goto L113
L115:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v408+int32(-9))))
	v429 = v425
	goto L113
L116:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408+int32(-5)))))
	v429 = v422
	goto L113
L117:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+int32(-3)))))
	v429 = v419
	goto L113
L118:
	;
	v429 = int32(base.Ui32(v412) >> (uint(int32(3)) % 32))
	goto L113
L119:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut5Command[0]))
	if v402 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v403 = int32(_a_F_lolwut5Command_5)
	goto L122
L121:
	;
	v403 = int32(_a_F_lolwut5Command_6)
	goto L122
L122:
	;
	v404 = F_sdscat(m, v397, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v408 = F_sdscatlen(m, v404, int32(_a_F_lolwut5Command_0), int32(1))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+int32(-1)))))
	switch v412 & int32(7) {
	case 0:
		goto L118
	case 1:
		goto L117
	case 2:
		goto L116
	case 3:
		goto L115
	case 4:
		goto L114
	default:
		v429 = int32(0)
		goto L113
	}
L125:
	;
	F_sdsfree(m, v408)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_lwFreeCanvas(m, v91)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	goto L1
}
func F_lolwut9Command(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 float32
	_ = v101
	var v102 float32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v117 float32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v135 float32
	_ = v135
	var v141 float32
	_ = v141
	var v142 float32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v214 int32
	_ = v214
	var v220 float32
	_ = v220
	var v225 float32
	_ = v225
	var v231 float32
	_ = v231
	var v232 float32
	_ = v232
	var v235 float32
	_ = v235
	var v236 float32
	_ = v236
	var v244 float32
	_ = v244
	var v248 float32
	_ = v248
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	v17 = m.G0
	v18 = int32(80)
	v19 = v17 - v18
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = int32(40)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v25 + int32(-2) {
	case 0, 2:
		goto L3
	default:
		goto L4
	}
L1:
	;
	m.G0 = v19 + int32(80)
	return
L2:
	;
	if v25 < int32(2) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a_F_lolwut9Command_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if v25 < int32(6) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	v52 = int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v53 < v52 {
		v59 = v52
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = F_getLongFromObjectOrReply(m, l0, v36, v19+int32(76), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v42 < int32(3) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v50 = F_getLongFromObjectOrReply(m, l0, v46, v19+int32(72), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v63 < v62 {
		v69 = v62
		goto L20
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v59
	goto L15
L17:
	;
	if base.Ui32(v53) < base.Ui32(int32(161)) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v59 = int32(160)
	goto L16
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v72 != int32(5) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v69
	goto L19
L21:
	;
	if base.Ui32(v63) < base.Ui32(int32(81)) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v69 = int32(80)
	goto L20
L23:
	;
	v144 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v151 = F_sdsnewlen(m, v144, (v146+int32(1))*v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L32
	}
L24:
	;
	v104 = int32(0)
	v106 = *(*int64)(unsafe.Add(mBase, _c_F_lolwut9Command[0]))
	v110 = v106*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_lolwut9Command[0])) = v110
	goto L30
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v80 = F_getLongDoubleFromObjectOrReply(m, l0, v76, v19+int32(56), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v80 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v87 = F_getLongDoubleFromObjectOrReply(m, l0, v83, v19+int32(40), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v19+int32(64))))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v19)+40))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v19+int32(48))))
	v101 = F___trunctfsf2(m, v95, v100)
	mBase = m.M
	v102 = F___trunctfsf2(m, v94, v93)
	mBase = m.M
	v141 = v101
	v142 = v102
	goto L23
L30:
	;
	v117 = base.F32_mul(base.F32_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v110)>>(uint(int64(33))%64)))), float32(4.656613e-10))
	v122 = int32(0)
	v124 = *(*int64)(unsafe.Add(mBase, _c_F_lolwut9Command[0]))
	v128 = v124*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _c_F_lolwut9Command[0])) = v128
	goto L31
L31:
	;
	v135 = base.F32_mul(base.F32_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v128)>>(uint(int64(33))%64)))), float32(4.656613e-10))
	v141 = base.F32_add(base.F32_add(v135, v135), float32(-1))
	v142 = base.F32_add(base.F32_add(v117, v117), float32(-1))
	goto L23
L32:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v153 < int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = base.F64_promote_f32(v141)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = base.F64_promote_f32(v142)
	v327 = F_sdscatprintf(m, v151, int32(_a_F_lolwut9Command_1), v19+int32(16))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L54
	}
L34:
	;
	v165 = v144
	goto L35
L35:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v172 < int32(1) {
		v288 = v172
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L33
L37:
	;
	v293 = int32(1)
	v298 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+(v288+v293)*v165+v288))) = uint8(v298)
	v301 = v165 + v293
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v301 < v302 {
		v165 = v301
		goto L35
	} else {
		goto L53
	}
L38:
	;
	v184 = int32(0)
	v192 = v172
	goto L39
L39:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v214 = int32(0)
	v220 = base.F32_add(base.F32_div(base.F32_mul(base.F32_add(base.F32_convert_i32_u(v184), float32(0.5)), float32(4)), base.F32_convert_i32_s(v192)), float32(-2))
	v225 = base.F32_sub(float32(2), base.F32_div(base.F32_mul(base.F32_add(base.F32_convert_i32_u(v165), float32(0.5)), float32(4)), base.F32_convert_i32_s(v198)))
	goto L42
L40:
	;
	v288 = v275
	goto L37
L41:
	;
	v265 = int32(10)
	if base.Ui32(v257) < base.Ui32(v265) {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v231 = base.F32_add(v142, base.F32_sub(base.F32_mul(v220, v220), base.F32_mul(v225, v225)))
	v232 = base.F32_mul(v231, v231)
	v235 = base.F32_add(base.F32_mul(base.F32_add(v220, v220), v225), v141)
	v236 = base.F32_mul(v235, v235)
	if base.F32_gt(base.F32_add(v232, v236), float32(4)) != 0 {
		v257 = v214
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v257 = v214 | int32(1)
	goto L41
L44:
	;
	if v214 != int32(10) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v244 = base.F32_add(v142, base.F32_sub(v232, v236))
	v248 = base.F32_add(base.F32_mul(base.F32_add(v231, v231), v235), v141)
	if base.F32_gt(base.F32_add(base.F32_mul(v244, v244), base.F32_mul(v248, v248)), float32(4)) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v257 = int32(11)
	goto L41
L47:
	;
	goto L43
L48:
	;
	v214 = v214 + int32(2)
	v220 = v244
	v225 = v248
	goto L42
L49:
	;
	v268 = v257
	goto L51
L50:
	;
	v268 = v265
	goto L51
L51:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_c_F_lolwut9Command[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v151+(v192+int32(1))*v165+v184))) = uint8(v271)
	v274 = v184 + int32(1)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v274 < v275 {
		v184 = v274
		v192 = v275
		goto L39
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	goto L36
L54:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut9Command[2]))
	if v332 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v333 = int32(_a_F_lolwut9Command_2)
	goto L57
L56:
	;
	v333 = int32(_a_F_lolwut9Command_3)
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v333
	v337 = F_sdscatprintf(m, v327, int32(_a_F_lolwut9Command_4), v19)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L64
	}
L58:
	;
	F_addReplyVerbatim(m, l0, v348, v369, int32(_a_F_lolwut9Command_5))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L70
	}
L59:
	;
	v369 = int32(base.Ui32(v352) >> (uint(int32(3)) % 32))
	goto L58
L60:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348+int32(-3)))))
	v369 = v366
	goto L58
L61:
	;
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348+int32(-5)))))
	v369 = v363
	goto L58
L62:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348+int32(-9))))
	v369 = v360
	goto L58
L63:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348+int32(-17))))
	v369 = v357
	goto L58
L64:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_lolwut9Command[2]))
	if v342 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v343 = int32(_a_F_lolwut9Command_6)
	goto L67
L66:
	;
	v343 = int32(_a_F_lolwut9Command_7)
	goto L67
L67:
	;
	v344 = F_sdscat(m, v337, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v348 = F_sdscatlen(m, v344, int32(_a_F_lolwut9Command_8), int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348+int32(-1)))))
	switch v352 & int32(7) {
	case 0:
		goto L59
	case 1:
		goto L60
	case 2:
		goto L61
	case 3:
		goto L62
	case 4:
		goto L63
	default:
		v369 = int32(0)
		goto L58
	}
L70:
	;
	F_sdsfree(m, v348)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L1
}
func F_lolwutUnstableCommand(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_sdsempty(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_lolwutUnstableCommand[0]))
		if v15 != 0 {
			v16 = int32(_a_F_lolwutUnstableCommand_0)
		} else {
			v16 = int32(_a_F_lolwutUnstableCommand_1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
		v20 = F_sdscatprintf(m, v10, int32(_a_F_lolwutUnstableCommand_2), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_lolwutUnstableCommand[0]))
			if v25 != 0 {
				v26 = int32(_a_F_lolwutUnstableCommand_3)
			} else {
				v26 = int32(_a_F_lolwutUnstableCommand_4)
			}
			v27 = F_sdscat(m, v20, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v31 = F_sdscatlen(m, v27, int32(_a_F_lolwutUnstableCommand_5), int32(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-1)))))
					switch v35 & int32(7) {
					case 0:
						v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
					case 1:
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-3)))))
						v52 = v42
					case 2:
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+int32(-5)))))
						v52 = v45
					case 3:
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-9))))
						v52 = v48
					case 4:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v31+int32(-17))))
						v52 = v51
					default:
						v52 = int32(0)
					}
					F_addReplyVerbatim(m, l0, v31, v52, int32(_a_F_lolwutUnstableCommand_6))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_sdsfree(m, v31)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_lrangeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v13 = F_getLongFromObjectOrReply(m, l0, v9, v6+int32(12), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
			v20 = F_getLongFromObjectOrReply(m, l0, v16, v6+int32(8), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				if v20 != 0 {
					m.G0 = v6 + int32(16)
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_lrangeCommand[0]))
					v26 = F_lookupKeyReadOrReply(m, l0, v23, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 == int32(0) {
							m.G0 = v6 + int32(16)
							return
						} else {
							v31 = F_checkType(m, l0, v26, int32(1))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								if v31 != 0 {
									m.G0 = v6 + int32(16)
									return
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
									v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									F_addListRangeReply(m, l0, v26, v33, v34, int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
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
func F_lru_import(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_lru_import[0]))
	return (v3 - l0) & int32(16777215)
}
func F_lrulfu_getIdleness(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	v3 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lrulfu_getIdleness[0])))
	if v8 != int32(1) {
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_getIdleness[1]))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = (v40 - l0) & int32(16777215)
		return l0
	} else {
		v11 = int32(0)
		v12 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_lrulfu_getIdleness[2])))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_lrulfu_getIdleness[3]))
		if v14 == v11 {
			v25 = v3
		} else {
			v20 = int32(65535)
			v22 = base.I32_div_s((v12-int32(base.Ui32(l0)>>(uint(int32(8))%32)))&v20, v14)
			v25 = v22 & v20
		}
		v28 = l0 & int32(255)
		v29 = v28 - v25
		if base.Ui32(v28) < base.Ui32(v29) {
			v31 = int32(0)
		} else {
			v31 = v29
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31 ^ int32(255)
		return v31 | v12<<(uint(int32(8))%32)
	}
}
func F_lrulfu_isUsingLFU(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lrulfu_isUsingLFU[0])))
	return v2
}
func F_lsetCommand(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_lsetCommand[0]))
	v14 = F_lookupKeyWriteOrReply(m, l0, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			m.G0 = v8 + int32(16)
			return
		} else {
			v19 = F_checkType(m, l0, v14, int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					m.G0 = v8 + int32(16)
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
					v27 = F_getLongFromObjectOrReply(m, l0, v23, v8+int32(12), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != 0 {
							m.G0 = v8 + int32(16)
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v31 = int32(3)
							v33 = int32(0)
							F_listTypeTryConversionRaw(m, v14, int32(1), v30, v31, v31, v33, v33)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v38 = F_listTypeReplaceAtIndex(m, v14, v37, v22)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									if v38 == int32(0) {
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_lsetCommand[1]))
										F_addReplyErrorObject(m, l0, v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									} else {
										v43 = int32(0)
										F_listTypeTryConversionRaw(m, v14, int32(2), v43, v43, v43, v43, v43)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
											F_signalModifiedKey(m, l0, v50, v52)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
												F_notifyKeyspaceEvent(m, int32(16), int32(_a_F_lsetCommand_0), v58, v60)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													v63 = int32(_a_F_lsetCommand_1)
													v65 = *(*int64)(unsafe.Add(mBase, _c_F_lsetCommand[2]))
													*(*int64)(unsafe.Add(mBase, _c_F_lsetCommand[2])) = v65 + int64(1)
													v70 = *(*int32)(unsafe.Add(mBase, _c_F_lsetCommand[3]))
													F_addReply(m, l0, v70)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
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
func F_ltrimCommand(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = F_getLongFromObjectOrReply(m, l0, v14, v11+int32(12), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if v18 != 0 {
			m.G0 = v11 + int32(16)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v25 = F_getLongFromObjectOrReply(m, l0, v21, v11+int32(8), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					m.G0 = v11 + int32(16)
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
					v31 = F_lookupKeyWriteOrReply(m, l0, v28, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						if v31 == int32(0) {
							m.G0 = v11 + int32(16)
							return
						} else {
							v36 = F_checkType(m, l0, v31, int32(1))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								if v36 != 0 {
									m.G0 = v11 + int32(16)
									return
								} else {
									v38 = F_listTypeLength(m, v31)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										if int32(-1) < v40 {
											v45 = v40
										} else {
											v43 = v40 + v38
											*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v43
											v45 = v43
										}
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										if int32(-1) < v46 {
											v51 = v46
										} else {
											v49 = v46 + v38
											*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v49
											v51 = v49
										}
										if int32(-1) < v45 {
											v57 = v45
										} else {
											v54 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v54
											v57 = v54
										}
										v58 = int32(0)
										if v57 <= v51 {
											if v57 < v38 {
												if base.Ui32(v51) < base.Ui32(v38) {
													v65 = v51
												} else {
													v63 = v38 + int32(-1)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v63
													v65 = v63
												}
												v69 = v57
												v71 = v38 + (v65 ^ int32(-1))
											} else {
												v69 = v38
												v71 = v58
											}
										} else {
											v69 = v38
											v71 = v58
										}
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
										switch int32(base.Ui32(v72)>>(uint(int32(4))%32))&int32(15) + int32(-9) {
										case 0:
											v99 = F_objectGetVal(m, v31)
											mBase = m.M
											v101 = F_quicklistDelRange(m, v99, int32(0), v69)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = F_objectGetVal(m, v31)
												mBase = m.M
												v106 = F_quicklistDelRange(m, v103, int32(0)-v71, v71)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return
												} else {
													v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
													v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
													F_notifyKeyspaceEvent(m, int32(16), int32(_a_F_ltrimCommand_0), v111, v113)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														v116 = F_listTypeLength(m, v31)
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return
														} else {
															if v116 != 0 {
																v132 = int32(0)
																F_listTypeTryConversionRaw(m, v31, int32(2), v132, v132, v132, v132, v132)
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
																	return
																} else {
																	v139 = v69 + v71
																	if v139 == int32(0) {
																		v148 = int32(_a_F_ltrimCommand_1)
																		v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																		*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																		v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																		F_addReply(m, l0, v154)
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return
																		} else {
																			m.G0 = v11 + int32(16)
																			return
																		}
																	} else {
																		v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
																		F_signalModifiedKey(m, l0, v143, v145)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return
																		} else {
																			v148 = int32(_a_F_ltrimCommand_1)
																			v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																			*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																			v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																			F_addReply(m, l0, v154)
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return
																			} else {
																				m.G0 = v11 + int32(16)
																				return
																			}
																		}
																	}
																}
															} else {
																v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																v121 = F_dbDelete(m, v118, v120)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return
																} else {
																	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
																	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+28))
																	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_ltrimCommand_2), v126, v128)
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
																		return
																	} else {
																		v139 = v69 + v71
																		if v139 == int32(0) {
																			v148 = int32(_a_F_ltrimCommand_1)
																			v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																			*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																			v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																			F_addReply(m, l0, v154)
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return
																			} else {
																				m.G0 = v11 + int32(16)
																				return
																			}
																		} else {
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
																			F_signalModifiedKey(m, l0, v143, v145)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return
																			} else {
																				v148 = int32(_a_F_ltrimCommand_1)
																				v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																				*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																				v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																				F_addReply(m, l0, v154)
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					m.G0 = v11 + int32(16)
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
											F__serverPanic_1(m, int32(_a_F_ltrimCommand_3), int32(904), int32(_a_F_ltrimCommand_4), int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										case 2:
											v79 = F_objectGetVal(m, v31)
											mBase = m.M
											v81 = F_lpDeleteRange(m, v79, int32(0), v69)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												F_objectSetVal(m, v31, v81)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													v85 = F_objectGetVal(m, v31)
													mBase = m.M
													v88 = F_lpDeleteRange(m, v85, int32(0)-v71, v71)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														F_objectSetVal(m, v31, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return
														} else {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
															F_notifyKeyspaceEvent(m, int32(16), int32(_a_F_ltrimCommand_0), v111, v113)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																v116 = F_listTypeLength(m, v31)
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return
																} else {
																	if v116 != 0 {
																		v132 = int32(0)
																		F_listTypeTryConversionRaw(m, v31, int32(2), v132, v132, v132, v132, v132)
																		mBase = m.M
																		v138 = m.ExcPending
																		if v138 != 0 {
																			return
																		} else {
																			v139 = v69 + v71
																			if v139 == int32(0) {
																				v148 = int32(_a_F_ltrimCommand_1)
																				v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																				*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																				v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																				F_addReply(m, l0, v154)
																				mBase = m.M
																				v156 = m.ExcPending
																				if v156 != 0 {
																					return
																				} else {
																					m.G0 = v11 + int32(16)
																					return
																				}
																			} else {
																				v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
																				F_signalModifiedKey(m, l0, v143, v145)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return
																				} else {
																					v148 = int32(_a_F_ltrimCommand_1)
																					v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																					*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																					v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																					F_addReply(m, l0, v154)
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						m.G0 = v11 + int32(16)
																						return
																					}
																				}
																			}
																		}
																	} else {
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																		v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
																		v121 = F_dbDelete(m, v118, v120)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
																			v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+28))
																			F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_ltrimCommand_2), v126, v128)
																			mBase = m.M
																			v130 = m.ExcPending
																			if v130 != 0 {
																				return
																			} else {
																				v139 = v69 + v71
																				if v139 == int32(0) {
																					v148 = int32(_a_F_ltrimCommand_1)
																					v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																					*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																					v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																					F_addReply(m, l0, v154)
																					mBase = m.M
																					v156 = m.ExcPending
																					if v156 != 0 {
																						return
																					} else {
																						m.G0 = v11 + int32(16)
																						return
																					}
																				} else {
																					v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																					v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
																					F_signalModifiedKey(m, l0, v143, v145)
																					mBase = m.M
																					v147 = m.ExcPending
																					if v147 != 0 {
																						return
																					} else {
																						v148 = int32(_a_F_ltrimCommand_1)
																						v150 = *(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1]))
																						*(*int64)(unsafe.Add(mBase, _c_F_ltrimCommand[1])) = v150 + base.I64_extend_i32_s(v139)
																						v154 = *(*int32)(unsafe.Add(mBase, _c_F_ltrimCommand[0]))
																						F_addReply(m, l0, v154)
																						mBase = m.M
																						v156 = m.ExcPending
																						if v156 != 0 {
																							return
																						} else {
																							m.G0 = v11 + int32(16)
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
				}
			}
		}
	}
}
func F_luaG_typeerror(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = m.G399
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v18<<(uint(int32(2))%32))))
	if base.Ui32(v16) <= base.Ui32(v17) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	v68 = m.G3
	F_luaG_runerror(m, l0, v68+int32(_a_F_luaG_typeerror_0), v11)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L12
	}
L3:
	;
	v31 = v17
	goto L5
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = F_getobjname(m, l0, v15, (l1-v37)>>(uint(int32(4))%32), v11+int32(44))
	mBase = m.M
	if v43 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L5:
	;
	if l1 == v31 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = v31 + int32(16)
	if base.Ui32(v16) <= base.Ui32(v35) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v31 = v35
	goto L5
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v43
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v23
	v51 = m.G3
	F_luaG_runerror(m, l0, v51+int32(_a_F_luaG_typeerror_1), v11+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	goto L1
L12:
	;
	goto L1
}
func F_luaM_growaux_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v10 = base.I32_div_s(l4, int32(2))
	if v8 < v10 {
		v19 = v8 << (uint(int32(1)) % 32)
		v20 = int32(4)
		if v20 < v19 {
			v23 = v19
		} else {
			v23 = v20
		}
		v24 = v23
		v28 = base.I32_div_u_s(int32(-3), l3)
		if base.Ui32(v28) < base.Ui32(v24+int32(1)) {
			v49 = m.G3
			F_luaG_runerror(m, l0, v49+int32(_a_F_luaM_growaux__0), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
				return int32(0)
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v33 = v32 * l3
			v34 = v24 * l3
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
			v36 = m.T0[v35].(func(*base.Module, int32, int32, int32, int32) int32)(m, v31, l1, v33, v34)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
					return v36
				} else {
					if v36 != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
						return v36
					} else {
						F_luaD_throw(m, l0, int32(4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
							return v36
						}
					}
				}
			}
		}
	} else {
		if v8 < l4 {
			v24 = l4
			v28 = base.I32_div_u_s(int32(-3), l3)
			if base.Ui32(v28) < base.Ui32(v24+int32(1)) {
				v49 = m.G3
				F_luaG_runerror(m, l0, v49+int32(_a_F_luaM_growaux__0), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
					return int32(0)
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v33 = v32 * l3
				v34 = v24 * l3
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				v36 = m.T0[v35].(func(*base.Module, int32, int32, int32, int32) int32)(m, v31, l1, v33, v34)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v34 == int32(0) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
						return v36
					} else {
						if v36 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
							return v36
						} else {
							F_luaD_throw(m, l0, int32(4))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
								*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
								return v36
							}
						}
					}
				}
			}
		} else {
			F_luaG_runerror(m, l0, l5, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v24 = l4
				v28 = base.I32_div_u_s(int32(-3), l3)
				if base.Ui32(v28) < base.Ui32(v24+int32(1)) {
					v49 = m.G3
					F_luaG_runerror(m, l0, v49+int32(_a_F_luaM_growaux__0), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
						return int32(0)
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v33 = v32 * l3
					v34 = v24 * l3
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
					v36 = m.T0[v35].(func(*base.Module, int32, int32, int32, int32) int32)(m, v31, l1, v33, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
							return v36
						} else {
							if v36 != 0 {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
								*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
								return v36
							} else {
								F_luaD_throw(m, l0, int32(4))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
									*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v34 - v33 + v44
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
									return v36
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_luaM_realloc_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32, int32) int32)(m, v7, l1, l2, l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l3 == int32(0) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = l3 - l2 + v19
			return v9
		} else {
			if v9 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = l3 - l2 + v19
				return v9
			} else {
				F_luaD_throw(m, l0, int32(4))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = l3 - l2 + v19
					return v9
				}
			}
		}
	}
}
func F_luaS_newlstr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	if base.Ui32(l2) <= base.Ui32(int32(3)) {
		v46 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v52 = l1 + l2&int32(-4)
	v53 = int32(0)
	switch l2 & int32(3) {
	default:
		v77 = v46
		goto L6
	case 1:
		v64 = v53
		goto L7
	case 2:
		v59 = v53
		goto L8
	case 3:
		goto L9
	}
L2:
	;
	v16 = l2
	v18 = int32(0)
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+v18<<(uint(int32(2))%32))))
	v39 = base.I32_rotl((int32(base.Ui32(v23*int32(-862048943))>>(uint(int32(17))%32))|v23*int32(380141568))*int32(461845907)^v16, int32(13))*int32(5) + int32(-430675100)
	v41 = v18 + int32(1)
	if v41 != int32(base.Ui32(l2)>>(uint(int32(2))%32)) {
		v16 = v39
		v18 = v41
		goto L3
	} else {
		goto L5
	}
L4:
	;
	v46 = v39
	goto L1
L5:
	;
	goto L4
L6:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = v77 ^ l2
	v82 = int32(16)
	v86 = (int32(base.Ui32(v81)>>(uint(v82)%32)) ^ v81) * int32(-2048144789)
	v91 = (int32(base.Ui32(v86)>>(uint(int32(13))%32)) ^ v86) * int32(-1028477387)
	v94 = int32(base.Ui32(v91)>>(uint(v82)%32)) ^ v91
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v80+v94&(v95+int32(-1))<<(uint(int32(2))%32))))
	if v102 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v66 = v64 ^ v65
	v77 = (int32(base.Ui32(v66*int32(-862048943))>>(uint(int32(17))%32))|v66*int32(380141568))*int32(461845907) ^ v46
	goto L6
L8:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v64 = v60<<(uint(int32(8))%32) | v59
	goto L7
L9:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+2)))
	v59 = v56 << (uint(int32(16)) % 32)
	goto L8
L10:
	;
	return v266
L11:
	;
	if base.Ui32(int32(17)) < base.Ui32(l2+int32(19)) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	v110 = v102
	goto L13
L13:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v112 != l2 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L11
L15:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v193 != 0 {
		v110 = v193
		goto L13
	} else {
		goto L35
	}
L16:
	;
	v115 = v110 + int32(16)
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		v139 = l1
		v140 = v115
		v141 = l2
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v179 != 0 {
		goto L15
	} else {
		goto L33
	}
L18:
	;
	v179 = int32(0)
	goto L17
L19:
	;
	v151 = v146
	v152 = v147
	v153 = v148
	goto L29
L20:
	;
	if v141 == int32(0) {
		goto L18
	} else {
		goto L27
	}
L21:
	;
	if (v115|l1)&int32(3) != 0 {
		v146 = l1
		v147 = v115
		v148 = l2
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v123 = l1
	v124 = v115
	v125 = l2
	goto L23
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v128 != v129 {
		v146 = v123
		v147 = v124
		v148 = v125
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v139 = v134
	v140 = v132
	v141 = v136
	goto L20
L25:
	;
	v131 = int32(4)
	v132 = v124 + v131
	v134 = v123 + v131
	v136 = v125 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v136) {
		v123 = v134
		v124 = v132
		v125 = v136
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v146 = v139
	v147 = v140
	v148 = v141
	goto L19
L28:
	;
	v179 = v156 - v157
	goto L17
L29:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v156 != v157 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v159 = int32(1)
	v164 = v153 + int32(-1)
	if v164 == int32(0) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v151 = v151 + v159
	v152 = v152 + v159
	v153 = v164
	goto L29
L33:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+5)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)))
	if v180&(v181^int32(-1))&int32(3) == int32(0) {
		v266 = v110
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v190 = v180 ^ int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+5)) = uint8(v190)
	return v110
L35:
	;
	goto L14
L36:
	;
	v209 = int32(0)
	v213 = F_luaM_realloc_(m, l0, v209, v209, l2+int32(17))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L38
	} else {
		goto L40
	}
L37:
	;
	v205 = F_luaM_toobig(m, l0)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v213)+12)) = l2
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+20)))
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)) = uint8(v219)
	v221 = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)) = uint8(v221)
	v224 = v218 & int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)) = uint8(v224)
	v227 = v213 + int32(16)
	if l2 == v219 {
		v231 = v227
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231+l2))) = uint8(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	v242 = (v237 + int32(-1)) & v94 << (uint(int32(2)) % 32)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v236+v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v246+v242))) = v213
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v251 = v249 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	if base.Ui32(v251) <= base.Ui32(v253) {
		v266 = v213
		goto L10
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v230 = F__emscripten_memcpy_bulkmem(m, v227, l1, l2)
	mBase = m.M
	v231 = v230
	goto L42
L44:
	;
	if int32(1073741822) < v253 {
		v266 = v213
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_luaS_resize(m, l0, v253<<(uint(int32(1))%32))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v266 = v213
	goto L10
}
func F_luaS_newudata(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	if base.Ui32(l1) < base.Ui32(int32(-26)) {
		v12 = int32(0)
		v16 = F_luaM_realloc_(m, l0, v12, v12, l1+int32(24))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
			v21 = int32(7)
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)) = uint8(v21)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
			v27 = v19 & int32(3)
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)) = uint8(v27)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v16
			return v16
		}
	} else {
		v8 = F_luaM_toobig(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(0)
			v16 = F_luaM_realloc_(m, l0, v12, v12, l1+int32(24))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
				v21 = int32(7)
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)) = uint8(v21)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
				v27 = v19 & int32(3)
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)) = uint8(v27)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = v16
				return v16
			}
		}
	}
}
func F_luaT_gettmbyobj(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	switch v5 + int32(-5) {
	case 0:
		goto L4
	default:
		goto L2
	case 2:
		goto L3
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = m.G398
	if v21 == int32(0) {
		v55 = v22
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = v14 + v5<<(uint(int32(2))%32) + int32(152)
	goto L1
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = v11 + int32(8)
	goto L1
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = v8 + int32(16)
	goto L1
L5:
	;
	return v55
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+l2<<(uint(int32(2))%32))+188))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v33 = int32(-1)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)))
	v42 = v31 + v32&(v33<<(uint(v34)%32)^v33)<<(uint(int32(5))%32)
	goto L8
L7:
	;
	v55 = v54
	goto L5
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if v45 != int32(4) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v54 = v51
	goto L7
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v51 = m.G398
	if v50 != 0 {
		v42 = v50
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v48 != v29 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v54 = v42
	goto L7
L13:
	;
	goto L9
}
func F_luaZ_fill(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, v11, v12, v8+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v32 = v10
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v22 == int32(0) {
				v32 = v10
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v22 + int32(-1)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				v32 = v31
			}
		}
		m.G0 = v8 + int32(16)
		return v32
	}
}
func F_lwCreateCanvas(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v6 = F_valkey_malloc(m, int32(12))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v12 = l1 * l0
		v13 = F_valkey_malloc(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
			v17 = F__emscripten_memset_bulkmem(m, v13, base.I32_extend8_s(l2), v12)
			mBase = m.M
			return v6
		}
	}
}
func F_lwDrawPixel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = l3
	if l1 < int32(0) {
	} else {
		if l2 < int32(0) {
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v10 <= l1 {
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v12 <= l2 {
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*uint8)(unsafe.Add(mBase, uint32(v14+l1+v10*l2))) = uint8(v4)
				}
			}
		}
	}
	return
}
func F_lwDrawSchotter(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 float32
	_ = v32
	var v33 float32
	_ = v33
	var v36 float32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 float32
	_ = v54
	var v55 int32
	_ = v55
	var v57 float32
	_ = v57
	var v68 int32
	_ = v68
	var v82 float32
	_ = v82
	var v85 float32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v123 float32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v166 int64
	_ = v166
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v184 float32
	_ = v184
	var v187 float32
	_ = v187
	var v192 float32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v208 float32
	_ = v208
	var v213 float32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v229 float32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v238 int64
	_ = v238
	var v245 float32
	_ = v245
	var v249 float32
	_ = v249
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 float32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 float32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v295 int32
	_ = v295
	v24 = l0 << (uint(int32(1)) % 32)
	v25 = int32(2)
	v26 = base.B2i32(v25 < l0)
	v28 = v26 << (uint(v25) % 32)
	v32 = base.F32_div(base.F32_convert_i32_s(v24-v28), base.F32_convert_i32_s(l1))
	v33 = base.F32_convert_i32_s(l2)
	v36 = base.F32_add(base.F32_mul(v32, v33), base.F32_convert_i32_u(v28))
	if base.F32_lt(base.F32_abs(v36), float32(2.1474836e+09)) == int32(0) {
		v44 = int32(-2147483648)
	} else {
		v42 = base.I32_trunc_f32_s(v36)
		v44 = v42
	}
	v45 = int32(0)
	v47 = F_lwCreateCanvas(m, v24, v44, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return int32(0)
	} else {
		if l2 < int32(1) {
		} else {
			v54 = base.F32_mul(v32, float32(0.5))
			v55 = int32(1)
			v57 = base.F32_convert_i32_u(v26 << (uint(v55) % 32))
			v68 = v45
			for {
				if l1 < v55 {
				} else {
					v82 = base.F32_convert_i32_u(v68)
					v85 = base.F32_add(base.F32_add(base.F32_mul(v82, v32), v54), v57)
					if base.F32_lt(base.F32_abs(v85), float32(2.1474836e+09)) == int32(0) {
						v93 = int32(-2147483648)
					} else {
						v91 = base.I32_trunc_f32_s(v85)
						v93 = v91
					}
					v98 = int32(0)
					for {
						v123 = base.F32_add(base.F32_add(base.F32_mul(base.F32_convert_i32_u(v98), v32), v54), v57)
						if base.F32_lt(base.F32_abs(v123), float32(2.1474836e+09)) == int32(0) {
							v131 = int32(-2147483648)
						} else {
							v129 = base.I32_trunc_f32_s(v123)
							v131 = v129
						}
						if base.B2i32(base.Ui32(v68) < base.Ui32(int32(2))) == int32(0) {
							v136 = int32(0)
							v138 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v142 = v138*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v142
							v148 = int32(0)
							v150 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v154 = v150*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v154
							v160 = int32(0)
							v162 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v166 = v162*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v166
							v172 = int32(0)
							v174 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v178 = v174*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v178
							v184 = float32(4.656613e-10)
							v187 = base.F32_mul(base.F32_div(base.F32_mul(base.F32_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v142)>>(uint(int64(33))%64)))), v184), v33), v82)
							v192 = base.F32_mul(base.F32_div(base.F32_mul(base.F32_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v154)>>(uint(int64(33))%64)))), v184), v33), v82)
							v195 = int32(0)
							v197 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v201 = v197*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v201
							if base.I32_wrap_i64(int64(base.Ui64(v201)>>(uint(int64(33))%64)))&int32(1) != 0 {
								v208 = base.F32_neg(v192)
							} else {
								v208 = v192
							}
							v213 = base.F32_add(base.F32_div(base.F32_mul(v32, v208), float32(3)), base.F32_convert_i32_s(v131))
							if base.F32_lt(base.F32_abs(v213), float32(2.1474836e+09)) == int32(0) {
								v221 = int32(-2147483648)
							} else {
								v219 = base.I32_trunc_f32_s(v213)
								v221 = v219
							}
							v229 = base.F32_mul(base.F32_div(base.F32_mul(base.F32_convert_i32_s(base.I32_wrap_i64(int64(base.Ui64(v166)>>(uint(int64(33))%64)))), float32(4.656613e-10)), v33), v82)
							v232 = int32(0)
							v234 = *(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0]))
							v238 = v234*int64(6364136223846793005) + int64(1)
							*(*int64)(unsafe.Add(mBase, _c_F_lwDrawSchotter[0])) = v238
							if base.I32_wrap_i64(int64(base.Ui64(v238)>>(uint(int64(33))%64)))&int32(1) != 0 {
								v245 = base.F32_neg(v229)
							} else {
								v245 = v229
							}
							v249 = base.F32_add(base.F32_div(base.F32_mul(v32, v245), float32(3)), base.F32_convert_i32_s(v93))
							if base.F32_lt(base.F32_abs(v249), float32(2.1474836e+09)) == int32(0) {
								v257 = int32(-2147483648)
							} else {
								v255 = base.I32_trunc_f32_s(v249)
								v257 = v255
							}
							if base.I32_wrap_i64(int64(base.Ui64(v178)>>(uint(int64(33))%64)))&int32(1) != 0 {
								v258 = base.F32_neg(v187)
							} else {
								v258 = v187
							}
							v259 = v221
							v260 = v257
							v261 = v258
						} else {
							v259 = v131
							v260 = v93
							v261 = float32(0)
						}
						v267 = int32(1)
						F_lwDrawSquare(m, v47, v259, v260, v32, v261, v267)
						mBase = m.M
						v270 = v98 + v267
						if v270 != l1 {
							v98 = v270
							continue
						} else {
							break
						}
						break
					}
				}
				v295 = v68 + int32(1)
				if v295 != l2 {
					v68 = v295
					continue
				} else {
					break
				}
				break
			}
		}
		return v47
	}
}
func F_lwDrawSquare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float32, l4 float32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 float32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 float32
	_ = v46
	var v51 float32
	_ = v51
	var v59 float32
	_ = v59
	var v66 float32
	_ = v66
	var v70 float32
	_ = v70
	var v71 float32
	_ = v71
	var v75 float32
	_ = v75
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 float64
	_ = v100
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v142 float64
	_ = v142
	var v147 float64
	_ = v147
	var v155 float64
	_ = v155
	var v162 float64
	_ = v162
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v171 float64
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v200 float64
	_ = v200
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v216 float64
	_ = v216
	var v219 float64
	_ = v219
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v229 int64
	_ = v229
	var v234 int32
	_ = v234
	var v241 float64
	_ = v241
	var v246 float64
	_ = v246
	var v254 float64
	_ = v254
	var v261 float64
	_ = v261
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v270 float64
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 float64
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v303 float64
	_ = v303
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v316 float64
	_ = v316
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v327 float64
	_ = v327
	var v331 int64
	_ = v331
	var v336 int32
	_ = v336
	var v343 float64
	_ = v343
	var v348 float64
	_ = v348
	var v356 float64
	_ = v356
	var v363 float64
	_ = v363
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v372 float64
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v401 float64
	_ = v401
	var v405 int32
	_ = v405
	var v406 float64
	_ = v406
	var v407 float64
	_ = v407
	var v410 float64
	_ = v410
	var v412 float64
	_ = v412
	var v414 float64
	_ = v414
	var v417 float64
	_ = v417
	var v420 float64
	_ = v420
	var v425 float64
	_ = v425
	var v429 int64
	_ = v429
	var v434 int32
	_ = v434
	var v441 float64
	_ = v441
	var v446 float64
	_ = v446
	var v454 float64
	_ = v454
	var v461 float64
	_ = v461
	var v465 float64
	_ = v465
	var v466 float64
	_ = v466
	var v470 float64
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 float64
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v503 float64
	_ = v503
	var v507 int32
	_ = v507
	var v508 float64
	_ = v508
	var v509 float64
	_ = v509
	var v513 float64
	_ = v513
	var v514 float64
	_ = v514
	var v516 float64
	_ = v516
	var v518 float64
	_ = v518
	var v520 float64
	_ = v520
	var v527 float64
	_ = v527
	var v531 int64
	_ = v531
	var v536 int32
	_ = v536
	var v543 float64
	_ = v543
	var v548 float64
	_ = v548
	var v556 float64
	_ = v556
	var v563 float64
	_ = v563
	var v567 float64
	_ = v567
	var v568 float64
	_ = v568
	var v572 float64
	_ = v572
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v601 float64
	_ = v601
	var v605 int32
	_ = v605
	var v606 float64
	_ = v606
	var v607 float64
	_ = v607
	var v610 float64
	_ = v610
	var v612 float64
	_ = v612
	var v614 float64
	_ = v614
	var v617 float64
	_ = v617
	var v620 float64
	_ = v620
	var v625 float64
	_ = v625
	var v629 int64
	_ = v629
	var v634 int32
	_ = v634
	var v641 float64
	_ = v641
	var v646 float64
	_ = v646
	var v654 float64
	_ = v654
	var v661 float64
	_ = v661
	var v665 float64
	_ = v665
	var v666 float64
	_ = v666
	var v670 float64
	_ = v670
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 float64
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v703 float64
	_ = v703
	var v707 int32
	_ = v707
	var v708 float64
	_ = v708
	var v709 float64
	_ = v709
	var v713 float64
	_ = v713
	var v714 float64
	_ = v714
	var v716 float64
	_ = v716
	var v718 float64
	_ = v718
	var v720 float64
	_ = v720
	var v727 float64
	_ = v727
	var v731 int64
	_ = v731
	var v736 int32
	_ = v736
	var v743 float64
	_ = v743
	var v748 float64
	_ = v748
	var v756 float64
	_ = v756
	var v763 float64
	_ = v763
	var v767 float64
	_ = v767
	var v768 float64
	_ = v768
	var v772 float64
	_ = v772
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v801 float64
	_ = v801
	var v805 int32
	_ = v805
	var v806 float64
	_ = v806
	var v807 float64
	_ = v807
	var v810 float64
	_ = v810
	var v812 float64
	_ = v812
	var v814 float64
	_ = v814
	var v817 float64
	_ = v817
	var v820 float64
	_ = v820
	var v825 float64
	_ = v825
	var v829 int64
	_ = v829
	var v834 int32
	_ = v834
	var v841 float64
	_ = v841
	var v846 float64
	_ = v846
	var v854 float64
	_ = v854
	var v861 float64
	_ = v861
	var v865 float64
	_ = v865
	var v866 float64
	_ = v866
	var v870 float64
	_ = v870
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v893 int32
	_ = v893
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v958 int32
	_ = v958
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	v6 = l5
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	v31 = base.F32_demote_f64(base.F64_div(base.F64_promote_f32(l3), float64(1.4142135623)))
	v35 = base.I32_reinterpret_f32(v31)
	v39 = int32(base.Ui32(v35)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(int32(149)) < base.Ui32(v39) {
		v71 = v31
		v75 = v71
	} else {
		if base.Ui32(int32(125)) < base.Ui32(v39) {
			v46 = base.F32_abs(v31)
			v51 = base.F32_sub(base.F32_add(base.F32_add(v46, float32(8.388608e+06)), float32(-8.388608e+06)), v46)
			if base.F32_gt(v51, float32(0.5)) == int32(0) {
				v59 = base.F32_add(v46, v51)
				if base.F32_le(v51, float32(-0.5)) == int32(0) {
					v66 = v59
				} else {
					v66 = base.F32_add(v59, float32(1))
				}
			} else {
				v66 = base.F32_add(base.F32_add(v46, v51), float32(-1))
			}
			if v35 < int32(0) {
				v70 = base.F32_neg(v66)
			} else {
				v70 = v66
			}
			v71 = v70
			v75 = v71
		} else {
			v75 = base.F32_mul(v31, float32(0))
		}
	}
	v80 = base.F64_promote_f32(base.F32_demote_f64(base.F64_add(base.F64_promote_f32(l4), float64(0.7853981633974483))))
	v84 = m.G0
	v86 = v84 - int32(16)
	m.G0 = v86
	v93 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v80))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v93) {
		if base.Ui32(v93) < base.Ui32(int32(2146435072)) {
			v104 = F___rem_pio2(m, v80, v86)
			mBase = m.M
			v105 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
			v106 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
			switch v104 & int32(3) {
			default:
				v110 = F___sin(m, v106, v105, int32(1))
				mBase = m.M
				v117 = v110
			case 1:
				v111 = F___cos(m, v106, v105)
				mBase = m.M
				v117 = v111
			case 2:
				v113 = F___sin(m, v106, v105, int32(1))
				mBase = m.M
				v117 = base.F64_neg(v113)
			case 3:
				v115 = F___cos(m, v106, v105)
				mBase = m.M
				v117 = base.F64_neg(v115)
			}
		} else {
			v117 = base.F64_sub(v80, v80)
		}
	} else {
		if base.Ui32(v93) < base.Ui32(int32(1045430272)) {
			v117 = v80
		} else {
			v100 = F___sin(m, v80, float64(0), int32(0))
			mBase = m.M
			v117 = v100
		}
	}
	m.G0 = v86 + int32(16)
	v123 = base.F64_promote_f32(v75)
	v125 = base.F64_convert_i32_s(l1)
	v126 = base.F64_add(base.F64_mul(v117, v123), v125)
	v130 = base.I64_reinterpret_f64(v126)
	v135 = base.I32_wrap_i64(int64(base.Ui64(v130)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v135) {
		v167 = v126
		v171 = v167
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v135) {
			v142 = base.F64_abs(v126)
			v147 = base.F64_sub(base.F64_add(base.F64_add(v142, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v142)
			if base.F64_gt(v147, float64(0.5)) == int32(0) {
				v155 = base.F64_add(v142, v147)
				if base.F64_le(v147, float64(-0.5)) == int32(0) {
					v162 = v155
				} else {
					v162 = base.F64_add(v155, float64(1))
				}
			} else {
				v162 = base.F64_add(base.F64_add(v142, v147), float64(-1))
			}
			if v130 < int64(0) {
				v166 = base.F64_neg(v162)
			} else {
				v166 = v162
			}
			v167 = v166
			v171 = v167
		} else {
			v171 = base.F64_mul(v126, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v171), float64(2.147483648e+09)) == int32(0) {
		v179 = int32(-2147483648)
	} else {
		v177 = base.I32_trunc_f64_s(v171)
		v179 = v177
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v179
	v184 = m.G0
	v186 = v184 - int32(16)
	m.G0 = v186
	v193 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v80))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v193) {
		if base.Ui32(v193) < base.Ui32(int32(2146435072)) {
			v204 = F___rem_pio2(m, v80, v186)
			mBase = m.M
			v205 = *(*float64)(unsafe.Add(mBase, uint32(v186)+8))
			v206 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
			switch v204 & int32(3) {
			default:
				v209 = F___cos(m, v206, v205)
				mBase = m.M
				v219 = v209
			case 1:
				v211 = F___sin(m, v206, v205, int32(1))
				mBase = m.M
				v219 = base.F64_neg(v211)
			case 2:
				v213 = F___cos(m, v206, v205)
				mBase = m.M
				v219 = base.F64_neg(v213)
			case 3:
				v216 = F___sin(m, v206, v205, int32(1))
				mBase = m.M
				v219 = v216
			}
		} else {
			v219 = base.F64_sub(v80, v80)
		}
	} else {
		if base.Ui32(v193) < base.Ui32(int32(1044816030)) {
			v219 = float64(1)
		} else {
			v200 = F___cos(m, v80, float64(0))
			mBase = m.M
			v219 = v200
		}
	}
	m.G0 = v186 + int32(16)
	v224 = base.F64_convert_i32_s(l2)
	v225 = base.F64_add(base.F64_mul(v219, v123), v224)
	v229 = base.I64_reinterpret_f64(v225)
	v234 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v234) {
		v266 = v225
		v270 = v266
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v234) {
			v241 = base.F64_abs(v225)
			v246 = base.F64_sub(base.F64_add(base.F64_add(v241, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v241)
			if base.F64_gt(v246, float64(0.5)) == int32(0) {
				v254 = base.F64_add(v241, v246)
				if base.F64_le(v246, float64(-0.5)) == int32(0) {
					v261 = v254
				} else {
					v261 = base.F64_add(v254, float64(1))
				}
			} else {
				v261 = base.F64_add(base.F64_add(v241, v246), float64(-1))
			}
			if v229 < int64(0) {
				v265 = base.F64_neg(v261)
			} else {
				v265 = v261
			}
			v266 = v265
			v270 = v266
		} else {
			v270 = base.F64_mul(v225, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v270), float64(2.147483648e+09)) == int32(0) {
		v278 = int32(-2147483648)
	} else {
		v276 = base.I32_trunc_f64_s(v270)
		v278 = v276
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v278
	v283 = base.F64_promote_f32(base.F32_demote_f64(base.F64_add(v80, float64(1.5707963267948966))))
	v287 = m.G0
	v289 = v287 - int32(16)
	m.G0 = v289
	v296 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v283))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v296) {
		if base.Ui32(v296) < base.Ui32(int32(2146435072)) {
			v307 = F___rem_pio2(m, v283, v289)
			mBase = m.M
			v308 = *(*float64)(unsafe.Add(mBase, uint32(v289)+8))
			v309 = *(*float64)(unsafe.Add(mBase, uint32(v289)))
			switch v307 & int32(3) {
			default:
				v313 = F___sin(m, v309, v308, int32(1))
				mBase = m.M
				v320 = v313
			case 1:
				v314 = F___cos(m, v309, v308)
				mBase = m.M
				v320 = v314
			case 2:
				v316 = F___sin(m, v309, v308, int32(1))
				mBase = m.M
				v320 = base.F64_neg(v316)
			case 3:
				v318 = F___cos(m, v309, v308)
				mBase = m.M
				v320 = base.F64_neg(v318)
			}
		} else {
			v320 = base.F64_sub(v283, v283)
		}
	} else {
		if base.Ui32(v296) < base.Ui32(int32(1045430272)) {
			v320 = v283
		} else {
			v303 = F___sin(m, v283, float64(0), int32(0))
			mBase = m.M
			v320 = v303
		}
	}
	m.G0 = v289 + int32(16)
	v327 = base.F64_add(base.F64_mul(v320, v123), v125)
	v331 = base.I64_reinterpret_f64(v327)
	v336 = base.I32_wrap_i64(int64(base.Ui64(v331)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v336) {
		v368 = v327
		v372 = v368
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v336) {
			v343 = base.F64_abs(v327)
			v348 = base.F64_sub(base.F64_add(base.F64_add(v343, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v343)
			if base.F64_gt(v348, float64(0.5)) == int32(0) {
				v356 = base.F64_add(v343, v348)
				if base.F64_le(v348, float64(-0.5)) == int32(0) {
					v363 = v356
				} else {
					v363 = base.F64_add(v356, float64(1))
				}
			} else {
				v363 = base.F64_add(base.F64_add(v343, v348), float64(-1))
			}
			if v331 < int64(0) {
				v367 = base.F64_neg(v363)
			} else {
				v367 = v363
			}
			v368 = v367
			v372 = v368
		} else {
			v372 = base.F64_mul(v327, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v372), float64(2.147483648e+09)) == int32(0) {
		v380 = int32(-2147483648)
	} else {
		v378 = base.I32_trunc_f64_s(v372)
		v380 = v378
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v380
	v385 = m.G0
	v387 = v385 - int32(16)
	m.G0 = v387
	v394 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v283))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v394) {
		if base.Ui32(v394) < base.Ui32(int32(2146435072)) {
			v405 = F___rem_pio2(m, v283, v387)
			mBase = m.M
			v406 = *(*float64)(unsafe.Add(mBase, uint32(v387)+8))
			v407 = *(*float64)(unsafe.Add(mBase, uint32(v387)))
			switch v405 & int32(3) {
			default:
				v410 = F___cos(m, v407, v406)
				mBase = m.M
				v420 = v410
			case 1:
				v412 = F___sin(m, v407, v406, int32(1))
				mBase = m.M
				v420 = base.F64_neg(v412)
			case 2:
				v414 = F___cos(m, v407, v406)
				mBase = m.M
				v420 = base.F64_neg(v414)
			case 3:
				v417 = F___sin(m, v407, v406, int32(1))
				mBase = m.M
				v420 = v417
			}
		} else {
			v420 = base.F64_sub(v283, v283)
		}
	} else {
		if base.Ui32(v394) < base.Ui32(int32(1044816030)) {
			v420 = float64(1)
		} else {
			v401 = F___cos(m, v283, float64(0))
			mBase = m.M
			v420 = v401
		}
	}
	m.G0 = v387 + int32(16)
	v425 = base.F64_add(base.F64_mul(v420, v123), v224)
	v429 = base.I64_reinterpret_f64(v425)
	v434 = base.I32_wrap_i64(int64(base.Ui64(v429)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v434) {
		v466 = v425
		v470 = v466
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v434) {
			v441 = base.F64_abs(v425)
			v446 = base.F64_sub(base.F64_add(base.F64_add(v441, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v441)
			if base.F64_gt(v446, float64(0.5)) == int32(0) {
				v454 = base.F64_add(v441, v446)
				if base.F64_le(v446, float64(-0.5)) == int32(0) {
					v461 = v454
				} else {
					v461 = base.F64_add(v454, float64(1))
				}
			} else {
				v461 = base.F64_add(base.F64_add(v441, v446), float64(-1))
			}
			if v429 < int64(0) {
				v465 = base.F64_neg(v461)
			} else {
				v465 = v461
			}
			v466 = v465
			v470 = v466
		} else {
			v470 = base.F64_mul(v425, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v470), float64(2.147483648e+09)) == int32(0) {
		v478 = int32(-2147483648)
	} else {
		v476 = base.I32_trunc_f64_s(v470)
		v478 = v476
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v478
	v483 = base.F64_promote_f32(base.F32_demote_f64(base.F64_add(v283, float64(1.5707963267948966))))
	v487 = m.G0
	v489 = v487 - int32(16)
	m.G0 = v489
	v496 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v483))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v496) {
		if base.Ui32(v496) < base.Ui32(int32(2146435072)) {
			v507 = F___rem_pio2(m, v483, v489)
			mBase = m.M
			v508 = *(*float64)(unsafe.Add(mBase, uint32(v489)+8))
			v509 = *(*float64)(unsafe.Add(mBase, uint32(v489)))
			switch v507 & int32(3) {
			default:
				v513 = F___sin(m, v509, v508, int32(1))
				mBase = m.M
				v520 = v513
			case 1:
				v514 = F___cos(m, v509, v508)
				mBase = m.M
				v520 = v514
			case 2:
				v516 = F___sin(m, v509, v508, int32(1))
				mBase = m.M
				v520 = base.F64_neg(v516)
			case 3:
				v518 = F___cos(m, v509, v508)
				mBase = m.M
				v520 = base.F64_neg(v518)
			}
		} else {
			v520 = base.F64_sub(v483, v483)
		}
	} else {
		if base.Ui32(v496) < base.Ui32(int32(1045430272)) {
			v520 = v483
		} else {
			v503 = F___sin(m, v483, float64(0), int32(0))
			mBase = m.M
			v520 = v503
		}
	}
	m.G0 = v489 + int32(16)
	v527 = base.F64_add(base.F64_mul(v520, v123), v125)
	v531 = base.I64_reinterpret_f64(v527)
	v536 = base.I32_wrap_i64(int64(base.Ui64(v531)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v536) {
		v568 = v527
		v572 = v568
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v536) {
			v543 = base.F64_abs(v527)
			v548 = base.F64_sub(base.F64_add(base.F64_add(v543, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v543)
			if base.F64_gt(v548, float64(0.5)) == int32(0) {
				v556 = base.F64_add(v543, v548)
				if base.F64_le(v548, float64(-0.5)) == int32(0) {
					v563 = v556
				} else {
					v563 = base.F64_add(v556, float64(1))
				}
			} else {
				v563 = base.F64_add(base.F64_add(v543, v548), float64(-1))
			}
			if v531 < int64(0) {
				v567 = base.F64_neg(v563)
			} else {
				v567 = v563
			}
			v568 = v567
			v572 = v568
		} else {
			v572 = base.F64_mul(v527, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v572), float64(2.147483648e+09)) == int32(0) {
		v580 = int32(-2147483648)
	} else {
		v578 = base.I32_trunc_f64_s(v572)
		v580 = v578
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v580
	v585 = m.G0
	v587 = v585 - int32(16)
	m.G0 = v587
	v594 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v483))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v594) {
		if base.Ui32(v594) < base.Ui32(int32(2146435072)) {
			v605 = F___rem_pio2(m, v483, v587)
			mBase = m.M
			v606 = *(*float64)(unsafe.Add(mBase, uint32(v587)+8))
			v607 = *(*float64)(unsafe.Add(mBase, uint32(v587)))
			switch v605 & int32(3) {
			default:
				v610 = F___cos(m, v607, v606)
				mBase = m.M
				v620 = v610
			case 1:
				v612 = F___sin(m, v607, v606, int32(1))
				mBase = m.M
				v620 = base.F64_neg(v612)
			case 2:
				v614 = F___cos(m, v607, v606)
				mBase = m.M
				v620 = base.F64_neg(v614)
			case 3:
				v617 = F___sin(m, v607, v606, int32(1))
				mBase = m.M
				v620 = v617
			}
		} else {
			v620 = base.F64_sub(v483, v483)
		}
	} else {
		if base.Ui32(v594) < base.Ui32(int32(1044816030)) {
			v620 = float64(1)
		} else {
			v601 = F___cos(m, v483, float64(0))
			mBase = m.M
			v620 = v601
		}
	}
	m.G0 = v587 + int32(16)
	v625 = base.F64_add(base.F64_mul(v620, v123), v224)
	v629 = base.I64_reinterpret_f64(v625)
	v634 = base.I32_wrap_i64(int64(base.Ui64(v629)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v634) {
		v666 = v625
		v670 = v666
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v634) {
			v641 = base.F64_abs(v625)
			v646 = base.F64_sub(base.F64_add(base.F64_add(v641, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v641)
			if base.F64_gt(v646, float64(0.5)) == int32(0) {
				v654 = base.F64_add(v641, v646)
				if base.F64_le(v646, float64(-0.5)) == int32(0) {
					v661 = v654
				} else {
					v661 = base.F64_add(v654, float64(1))
				}
			} else {
				v661 = base.F64_add(base.F64_add(v641, v646), float64(-1))
			}
			if v629 < int64(0) {
				v665 = base.F64_neg(v661)
			} else {
				v665 = v661
			}
			v666 = v665
			v670 = v666
		} else {
			v670 = base.F64_mul(v625, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v670), float64(2.147483648e+09)) == int32(0) {
		v678 = int32(-2147483648)
	} else {
		v676 = base.I32_trunc_f64_s(v670)
		v678 = v676
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v678
	v683 = base.F64_promote_f32(base.F32_demote_f64(base.F64_add(v483, float64(1.5707963267948966))))
	v687 = m.G0
	v689 = v687 - int32(16)
	m.G0 = v689
	v696 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v683))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v696) {
		if base.Ui32(v696) < base.Ui32(int32(2146435072)) {
			v707 = F___rem_pio2(m, v683, v689)
			mBase = m.M
			v708 = *(*float64)(unsafe.Add(mBase, uint32(v689)+8))
			v709 = *(*float64)(unsafe.Add(mBase, uint32(v689)))
			switch v707 & int32(3) {
			default:
				v713 = F___sin(m, v709, v708, int32(1))
				mBase = m.M
				v720 = v713
			case 1:
				v714 = F___cos(m, v709, v708)
				mBase = m.M
				v720 = v714
			case 2:
				v716 = F___sin(m, v709, v708, int32(1))
				mBase = m.M
				v720 = base.F64_neg(v716)
			case 3:
				v718 = F___cos(m, v709, v708)
				mBase = m.M
				v720 = base.F64_neg(v718)
			}
		} else {
			v720 = base.F64_sub(v683, v683)
		}
	} else {
		if base.Ui32(v696) < base.Ui32(int32(1045430272)) {
			v720 = v683
		} else {
			v703 = F___sin(m, v683, float64(0), int32(0))
			mBase = m.M
			v720 = v703
		}
	}
	m.G0 = v689 + int32(16)
	v727 = base.F64_add(base.F64_mul(v720, v123), v125)
	v731 = base.I64_reinterpret_f64(v727)
	v736 = base.I32_wrap_i64(int64(base.Ui64(v731)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v736) {
		v768 = v727
		v772 = v768
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v736) {
			v743 = base.F64_abs(v727)
			v748 = base.F64_sub(base.F64_add(base.F64_add(v743, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v743)
			if base.F64_gt(v748, float64(0.5)) == int32(0) {
				v756 = base.F64_add(v743, v748)
				if base.F64_le(v748, float64(-0.5)) == int32(0) {
					v763 = v756
				} else {
					v763 = base.F64_add(v756, float64(1))
				}
			} else {
				v763 = base.F64_add(base.F64_add(v743, v748), float64(-1))
			}
			if v731 < int64(0) {
				v767 = base.F64_neg(v763)
			} else {
				v767 = v763
			}
			v768 = v767
			v772 = v768
		} else {
			v772 = base.F64_mul(v727, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v772), float64(2.147483648e+09)) == int32(0) {
		v780 = int32(-2147483648)
	} else {
		v778 = base.I32_trunc_f64_s(v772)
		v780 = v778
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v780
	v785 = m.G0
	v787 = v785 - int32(16)
	m.G0 = v787
	v794 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v683))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072243195)) < base.Ui32(v794) {
		if base.Ui32(v794) < base.Ui32(int32(2146435072)) {
			v805 = F___rem_pio2(m, v683, v787)
			mBase = m.M
			v806 = *(*float64)(unsafe.Add(mBase, uint32(v787)+8))
			v807 = *(*float64)(unsafe.Add(mBase, uint32(v787)))
			switch v805 & int32(3) {
			default:
				v810 = F___cos(m, v807, v806)
				mBase = m.M
				v820 = v810
			case 1:
				v812 = F___sin(m, v807, v806, int32(1))
				mBase = m.M
				v820 = base.F64_neg(v812)
			case 2:
				v814 = F___cos(m, v807, v806)
				mBase = m.M
				v820 = base.F64_neg(v814)
			case 3:
				v817 = F___sin(m, v807, v806, int32(1))
				mBase = m.M
				v820 = v817
			}
		} else {
			v820 = base.F64_sub(v683, v683)
		}
	} else {
		if base.Ui32(v794) < base.Ui32(int32(1044816030)) {
			v820 = float64(1)
		} else {
			v801 = F___cos(m, v683, float64(0))
			mBase = m.M
			v820 = v801
		}
	}
	m.G0 = v787 + int32(16)
	v825 = base.F64_add(base.F64_mul(v820, v123), v224)
	v829 = base.I64_reinterpret_f64(v825)
	v834 = base.I32_wrap_i64(int64(base.Ui64(v829)>>(uint(int64(52))%64))) & int32(2047)
	if base.Ui32(int32(1074)) < base.Ui32(v834) {
		v866 = v825
		v870 = v866
	} else {
		if base.Ui32(int32(1021)) < base.Ui32(v834) {
			v841 = base.F64_abs(v825)
			v846 = base.F64_sub(base.F64_add(base.F64_add(v841, float64(4.503599627370496e+15)), float64(-4.503599627370496e+15)), v841)
			if base.F64_gt(v846, float64(0.5)) == int32(0) {
				v854 = base.F64_add(v841, v846)
				if base.F64_le(v846, float64(-0.5)) == int32(0) {
					v861 = v854
				} else {
					v861 = base.F64_add(v854, float64(1))
				}
			} else {
				v861 = base.F64_add(base.F64_add(v841, v846), float64(-1))
			}
			if v829 < int64(0) {
				v865 = base.F64_neg(v861)
			} else {
				v865 = v861
			}
			v866 = v865
			v870 = v866
		} else {
			v870 = base.F64_mul(v825, float64(0))
		}
	}
	if base.F64_lt(base.F64_abs(v870), float64(2.147483648e+09)) == int32(0) {
		v878 = int32(-2147483648)
	} else {
		v876 = base.I32_trunc_f64_s(v870)
		v878 = v876
	}
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v878
	v893 = int32(0)
	for {
		v904 = int32(2)
		v905 = v893 << (uint(v904) % 32)
		v907 = v26 + int32(16)
		v908 = int32(1)
		v909 = v893 + v908
		v913 = v909 & int32(3) << (uint(v904) % 32)
		v915 = *(*int32)(unsafe.Add(mBase, uint32(v907|v913)))
		v919 = *(*int32)(unsafe.Add(mBase, uint32(v905+v907)))
		v920 = v915 - v919
		v921 = int32(31)
		v922 = v920 >> (uint(v921) % 32)
		v924 = v920 ^ v922 - v922
		v926 = *(*int32)(unsafe.Add(mBase, uint32(v26|v913)))
		v928 = *(*int32)(unsafe.Add(mBase, uint32(v26+v905)))
		v929 = v926 - v928
		v931 = v929 >> (uint(v921) % 32)
		v933 = v929 ^ v931 - v931
		if v928 < v926 {
			v940 = v908
		} else {
			v940 = int32(-1)
		}
		if v919 < v915 {
			v944 = int32(1)
		} else {
			v944 = int32(-1)
		}
		v946 = v919
		v947 = v928
		v958 = v924 - v933
		for {
			if v946 < int32(0) {
			} else {
				if v947 < int32(0) {
				} else {
					v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v972 <= v946 {
					} else {
						v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v974 <= v947 {
						} else {
							v976 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*uint8)(unsafe.Add(mBase, uint32(v976+v946+v972*v947))) = uint8(v6)
						}
					}
				}
			}
			if v947 != v926 {
				v986 = v958 << (uint(int32(1)) % 32)
				v987 = base.B2i32(int32(0)-v933 < v986)
				if int32(0)-v933 < v986 {
					v988 = v933
				} else {
					v988 = int32(0)
				}
				v991 = base.B2i32(v986 < v924)
				if v986 < v924 {
					v992 = v924
				} else {
					v992 = int32(0)
				}
				if v986 < v924 {
					v995 = v940
				} else {
					v995 = int32(0)
				}
				if int32(0)-v933 < v986 {
					v998 = v944
				} else {
					v998 = int32(0)
				}
				v946 = v998 + v946
				v947 = v995 + v947
				v958 = v958 - v988 + v992
				continue
			} else {
			}
			if v946 == v915 {
				break
			} else {
				v986 = v958 << (uint(int32(1)) % 32)
				v987 = base.B2i32(int32(0)-v933 < v986)
				if int32(0)-v933 < v986 {
					v988 = v933
				} else {
					v988 = int32(0)
				}
				v991 = base.B2i32(v986 < v924)
				if v986 < v924 {
					v992 = v924
				} else {
					v992 = int32(0)
				}
				if v986 < v924 {
					v995 = v940
				} else {
					v995 = int32(0)
				}
				if int32(0)-v933 < v986 {
					v998 = v944
				} else {
					v998 = int32(0)
				}
				v946 = v998 + v946
				v947 = v995 + v947
				v958 = v958 - v988 + v992
				continue
			}
			break
		}
		if v909 != int32(4) {
			v893 = v909
			continue
		} else {
			break
		}
		break
	}
	m.G0 = v26 + int32(32)
	return
}
func F_lzf_compress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
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
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(262144)
	m.G0 = v19
	if l1 == v5 {
		v507 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(262144)
	return v507
L2:
	;
	if l3 == int32(0) {
		v507 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = l2 + l3
	v28 = l2 + int32(1)
	v30 = l0 + l1
	v32 = v30 + int32(-2)
	if base.Ui32(l0) < base.Ui32(v32) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if base.Ui32(v26) < base.Ui32(v380+int32(3)) {
		v507 = int32(0)
		goto L1
	} else {
		goto L75
	}
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v42 = l0
	v44 = v28
	v50 = int32(0)
	v51 = v35<<(uint(int32(8))%32) | v38
	goto L7
L6:
	;
	v378 = l0
	v380 = v28
	v386 = int32(0)
	v389 = v5
	goto L4
L7:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)))
	v60 = v51<<(uint(int32(8))%32) | v59
	v68 = v19 + (v60*int32(65531)+v51)&int32(65535)<<(uint(int32(2))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v42
	if base.Ui32(v69) <= base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v378 = v361
	v380 = v363
	v386 = v369
	v389 = v372
	goto L4
L9:
	;
	if base.Ui32(v361) < base.Ui32(v32) {
		v42 = v361
		v44 = v363
		v50 = v369
		v51 = v370
		goto L7
	} else {
		goto L74
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v303))) = uint8(v74)
	v318 = v303 + int32(2)
	v319 = v42 + v309
	if base.Ui32(v319) < base.Ui32(v32) {
		goto L72
	} else {
		goto L73
	}
L11:
	;
	v291 = v283 + int32(-9)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)) = uint8(v291)
	v296 = int32(base.Ui32(v74)>>(uint(int32(8))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v296)
	v303 = v101 + int32(2)
	v309 = v283
	goto L10
L12:
	;
	if base.Ui32(v44) < base.Ui32(v26) {
		goto L68
	} else {
		goto L69
	}
L13:
	;
	v74 = v69 ^ int32(-1) + v42
	if base.Ui32(int32(8191)) < base.Ui32(v74) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+2)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)))
	if v77 != v78 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69))))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
	if v80 != v81 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v44+int32(4)) < base.Ui32(v26) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v93 = int32(-1)
	v97 = v50 + v93
	*(*uint8)(unsafe.Add(mBase, uint32(v44+(v50^v93)))) = uint8(v97)
	v101 = v44 - base.B2i32(v50 == int32(0))
	v104 = v30 - v42 + int32(-2)
	if base.Ui32(int32(17)) <= base.Ui32(v104) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if base.Ui32(v44-base.B2i32(v50 == int32(0))+int32(4)) < base.Ui32(v26) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v507 = int32(0)
	goto L1
L20:
	;
	v251 = v241<<(uint(int32(5))%32) | int32(base.Ui32(v74)>>(uint(int32(8))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v251)
	v303 = v101 + int32(1)
	v309 = v240
	goto L10
L21:
	;
	v191 = int32(264)
	if base.Ui32(v104) < base.Ui32(v191) {
		goto L55
	} else {
		goto L56
	}
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+3)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+3)))
	if v108 == v109 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v190 = int32(2)
	goto L21
L24:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v115 == v116 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v240 = int32(3)
	v241 = int32(1)
	goto L20
L26:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+5)))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+5)))
	if v122 == v123 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v240 = int32(4)
	v241 = int32(2)
	goto L20
L28:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+6)))
	if v129 == v130 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v240 = int32(5)
	v241 = int32(3)
	goto L20
L30:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+7)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+7)))
	if v136 == v137 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v240 = int32(6)
	v241 = int32(4)
	goto L20
L32:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+8)))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+8)))
	if v143 == v144 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v240 = int32(7)
	v241 = int32(5)
	goto L20
L34:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+9)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+9)))
	if v150 == v151 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v240 = int32(8)
	v241 = int32(6)
	goto L20
L36:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+10)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+10)))
	if v154 == v155 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v283 = int32(9)
	goto L11
L38:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+11)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+11)))
	if v158 == v159 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v283 = int32(10)
	goto L11
L40:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+12)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)))
	if v162 == v163 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v283 = int32(11)
	goto L11
L42:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+13)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+13)))
	if v166 == v167 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v283 = int32(12)
	goto L11
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+14)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+14)))
	if v170 == v171 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v283 = int32(13)
	goto L11
L46:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+15)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+15)))
	if v174 == v175 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v283 = int32(14)
	goto L11
L48:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+16)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)))
	if v178 == v179 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v283 = int32(15)
	goto L11
L50:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+17)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+17)))
	if v182 == v183 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v283 = int32(16)
	goto L11
L52:
	;
	v186 = int32(18)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+18)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+18)))
	if v187 != v188 {
		v283 = v186
		goto L11
	} else {
		goto L54
	}
L53:
	;
	v283 = int32(17)
	goto L11
L54:
	;
	v190 = v186
	goto L21
L55:
	;
	v194 = v104
	goto L57
L56:
	;
	v194 = v191
	goto L57
L57:
	;
	v196 = v190 | int32(1)
	if base.Ui32(v196) < base.Ui32(v194) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v198 = v194
	goto L60
L59:
	;
	v198 = v196
	goto L60
L60:
	;
	v210 = v190
	goto L62
L61:
	;
	v228 = v226 + int32(-1)
	if base.Ui32(int32(6)) < base.Ui32(v228) {
		v283 = v225
		goto L11
	} else {
		goto L67
	}
L62:
	;
	v218 = v210 + int32(1)
	if base.Ui32(v218) < base.Ui32(v194) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v225 = v218
	v226 = v210
	goto L61
L64:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v218))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v218))))
	if v221 == v223 {
		v210 = v218
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v225 = v198
	v226 = v198 + int32(-1)
	goto L61
L66:
	;
	goto L63
L67:
	;
	v240 = v225
	v241 = v228
	goto L20
L68:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v258)
	v260 = int32(1)
	v261 = v42 + v260
	v263 = v50 + v260
	if v263 == int32(32) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v507 = int32(0)
	goto L1
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(-32)))) = uint8(v50)
	v361 = v261
	v363 = v44 + int32(2)
	v369 = int32(0)
	v370 = v60
	v372 = v263
	goto L9
L71:
	;
	v361 = v261
	v363 = v44 + int32(1)
	v369 = v263
	v370 = v60
	v372 = v263
	goto L9
L72:
	;
	v323 = v319 + int32(-1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v325 = int32(8)
	v328 = v319 + int32(-2)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	v332 = v324<<(uint(v325)%32) | v329<<(uint(int32(16))%32)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v334 = v332 | v333
	v335 = int32(65531)
	v340 = int32(65535)
	v342 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19+(v334*v335+int32(base.Ui32(v332)>>(uint(v325)%32)))&v340<<(uint(v342)%32)))) = v328
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+1)))
	v349 = v334<<(uint(v325)%32) | v348
	*(*int32)(unsafe.Add(mBase, uint32(v19+(v349*v335+v334)&v340<<(uint(v342)%32)))) = v323
	v361 = v319
	v363 = v318
	v369 = int32(0)
	v370 = v349
	v372 = v323
	goto L9
L73:
	;
	v378 = v319
	v380 = v318
	v386 = int32(0)
	v389 = v74
	goto L4
L74:
	;
	goto L8
L75:
	;
	if base.Ui32(v378) < base.Ui32(v30) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v492 = int32(-1)
	v496 = v488 + v492
	*(*uint8)(unsafe.Add(mBase, uint32(v479+(v488^v492)))) = uint8(v496)
	v507 = v479 - base.B2i32(v488 == int32(0)) - l2
	goto L1
L77:
	;
	v397 = int32(1)
	v399 = v30 - v378
	if v399&v397 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v479 = v380
	v488 = v386
	goto L76
L79:
	;
	if v30 == v378+v397 {
		v479 = v419
		v488 = v422
		goto L76
	} else {
		goto L84
	}
L80:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	*(*uint8)(unsafe.Add(mBase, uint32(v380))) = uint8(v402)
	v404 = int32(1)
	v405 = v378 + v404
	v407 = v386 + v404
	if v407 == int32(32) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v419 = v380
	v420 = v386
	v421 = v378
	v422 = v389
	goto L79
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v380+int32(-32)))) = uint8(v386)
	v417 = int32(0)
	v419 = v380 + int32(2)
	v420 = v417
	v421 = v405
	v422 = v417
	goto L79
L83:
	;
	v419 = v380 + int32(1)
	v420 = v407
	v421 = v405
	v422 = v407
	goto L79
L84:
	;
	v428 = v419
	v435 = v421
	v437 = v420
	goto L85
L85:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v441)
	v444 = v437 + int32(1)
	if v444 == int32(32) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v479 = v471
	v488 = v472
	goto L76
L87:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v457)
	v460 = v456 + int32(1)
	if v460 == int32(32) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v428+int32(-32)))) = uint8(v437)
	v455 = v428 + int32(2)
	v456 = int32(0)
	goto L87
L89:
	;
	v455 = v428 + int32(1)
	v456 = v444
	goto L87
L90:
	;
	v474 = v435 + int32(2)
	if v474 != v378+v399 {
		v428 = v471
		v435 = v474
		v437 = v472
		goto L85
	} else {
		goto L93
	}
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v455+int32(-32)))) = uint8(v456)
	v471 = v455 + int32(2)
	v472 = int32(0)
	goto L90
L92:
	;
	v471 = v455 + int32(1)
	v472 = v460
	goto L90
L93:
	;
	goto L86
}
