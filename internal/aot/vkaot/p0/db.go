package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_copyDbIdArgs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 < int32(5) {
		v227 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v227
L2:
	;
	v24 = int32(3)
	v25 = int32(0)
	goto L3
L3:
	;
	v30 = l0 + v24<<(uint(int32(2))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = F_objectGetVal(m, v31)
	mBase = m.M
	v33 = int32(_a138)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v141 != 0 {
		goto L39
	} else {
		goto L40
	}
L5:
	;
	v145 = v140 + int32(1)
	if v145 < l1 {
		v24 = v145
		v25 = v141
		goto L3
	} else {
		goto L38
	}
L6:
	;
	if v68-v70 == int32(0) {
		v140 = v24
		v141 = v25
		goto L5
	} else {
		goto L18
	}
L7:
	;
	v68 = F_tolower(m, v64)
	mBase = m.M
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v70 = F_tolower(m, v69)
	mBase = m.M
	goto L6
L8:
	;
	v38 = v32
	v39 = v33
	v40 = v36
	goto L11
L9:
	;
	v64 = int32(0)
	v65 = v33
	goto L7
L10:
	;
	v64 = v61 & int32(255)
	v65 = v60
	goto L7
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v42 == int32(0) {
		v60 = v39
		v61 = v40
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v60 = v54
	v61 = int32(0)
	goto L10
L13:
	;
	v46 = v40 & int32(255)
	if v46 == v42 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = int32(1)
	v54 = v39 + v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v55 != 0 {
		v38 = v38 + v53
		v39 = v54
		v40 = v55
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v48 = F_tolower(m, v46)
	mBase = m.M
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v50 = F_tolower(m, v49)
	mBase = m.M
	if v48 == v50 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v60 = v39
	v61 = v52
	goto L10
L17:
	;
	goto L12
L18:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v76 = F_objectGetVal(m, v75)
	mBase = m.M
	v77 = int32(_a506)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if int32(-2) < v24-l1 {
		v227 = v74
		goto L1
	} else {
		goto L31
	}
L20:
	;
	v112 = F_tolower(m, v108)
	mBase = m.M
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v114 = F_tolower(m, v113)
	mBase = m.M
	goto L19
L21:
	;
	v82 = v76
	v83 = v77
	v84 = v80
	goto L24
L22:
	;
	v108 = int32(0)
	v109 = v77
	goto L20
L23:
	;
	v108 = v105 & int32(255)
	v109 = v104
	goto L20
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v86 == int32(0) {
		v104 = v83
		v105 = v84
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v104 = v98
	v105 = int32(0)
	goto L23
L26:
	;
	v90 = v84 & int32(255)
	if v90 == v86 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v97 = int32(1)
	v98 = v83 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v99 != 0 {
		v82 = v82 + v97
		v83 = v98
		v84 = v99
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v92 = F_tolower(m, v90)
	mBase = m.M
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v94 = F_tolower(m, v93)
	mBase = m.M
	if v92 == v94 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v104 = v83
	v105 = v96
	goto L23
L30:
	;
	goto L25
L31:
	;
	if v112-v114 != 0 {
		v227 = v74
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v120 = v24 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0+v120<<(uint(int32(2))%32))))
	v127 = F_getLongLongFromObject(m, v124, v12+int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	if v127 != 0 {
		v227 = v74
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v131 < int64(0) {
		v227 = v74
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v135 = int64(*(*int32)(unsafe.Add(mBase, _consts[10])))
	if v135 <= v131 {
		v227 = v74
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v140 = v120
	v141 = v25 + int32(1)
	goto L5
L38:
	;
	goto L4
L39:
	;
	v150 = F_valkey_malloc(m, v141<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L33
	} else {
		goto L41
	}
L40:
	;
	v227 = int32(0)
	goto L1
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v160 = int32(3)
	goto L42
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0+v160<<(uint(int32(2))%32))))
	v168 = F_objectGetVal(m, v167)
	mBase = m.M
	v169 = int32(_a506)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v172 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v227 = v150
	goto L1
L44:
	;
	v221 = v218 + int32(1)
	if v221 < l1 {
		v160 = v221
		goto L42
	} else {
		goto L58
	}
L45:
	;
	if v204-v206 != 0 {
		v218 = v160
		goto L44
	} else {
		goto L57
	}
L46:
	;
	v204 = F_tolower(m, v200)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	goto L45
L47:
	;
	v174 = v168
	v175 = v169
	v176 = v172
	goto L50
L48:
	;
	v200 = int32(0)
	v201 = v169
	goto L46
L49:
	;
	v200 = v197 & int32(255)
	v201 = v196
	goto L46
L50:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v178 == int32(0) {
		v196 = v175
		v197 = v176
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v196 = v190
	v197 = int32(0)
	goto L49
L52:
	;
	v182 = v176 & int32(255)
	if v182 == v178 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v189 = int32(1)
	v190 = v175 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v191 != 0 {
		v174 = v174 + v189
		v175 = v190
		v176 = v191
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v184 = F_tolower(m, v182)
	mBase = m.M
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v186 = F_tolower(m, v185)
	mBase = m.M
	if v184 == v186 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v196 = v175
	v197 = v188
	goto L49
L56:
	;
	goto L51
L57:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v209 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v208 + v209
	v216 = v160 + v209
	*(*int32)(unsafe.Add(mBase, uint32(v150+v208<<(uint(int32(2))%32)))) = v216
	v218 = v216
	goto L44
L58:
	;
	goto L43
}
func F_dbAddRDBLoad(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v14 == v4 {
		v21 = v4
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = F_kvstoreHashtableFindPositionForInsert(m, v22, v21, l1, v9, int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v92 = v4
				m.G0 = v9 + int32(16)
				return v92
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v30 = F_objectSetKeyAndExpire(m, v28, l1, int64(-1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_kvstoreHashtableInsertAtPosition(m, v32, v21, v30, v9)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
						if base.Ui32(int32(-9)) < base.Ui32(v36) {
						} else {
							v39 = F_lrulfu_init(m)
							mBase = m.M
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
							*(*int32)(unsafe.Add(mBase, uint32(v30))) = v40 | v39<<(uint(int32(8))%32)
						}
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						if v46&int32(15) != int32(4) {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
							v92 = int32(1)
							m.G0 = v9 + int32(16)
							return v92
						} else {
							v51 = F_hashTypeHasVolatileFields(m, v30)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
									v92 = int32(1)
									m.G0 = v9 + int32(16)
									return v92
								} else {
									v55 = int32(0)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
									if v58&int32(2) == v55 {
										v78 = v55
									} else {
										v72 = v30 + (v58&int32(4) ^ int32(12)) + v58<<(uint(int32(3))%32)&int32(8)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
										v78 = v72 + v73 + int32(1)
									}
									v80 = *(*int32)(unsafe.Add(mBase, _consts[63]))
									if v80 != 0 {
										v82 = F_getKeySlot(m, v78)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											v84 = v82
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v86 = F_kvstoreHashtableAdd(m, v85, v84, v30)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
												v92 = int32(1)
												m.G0 = v9 + int32(16)
												return v92
											}
										}
									} else {
										v84 = int32(0)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v86 = F_kvstoreHashtableAdd(m, v85, v84, v30)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
											v92 = int32(1)
											m.G0 = v9 + int32(16)
											return v92
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
		v17 = F_getKeySlot(m, l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v24 = F_kvstoreHashtableFindPositionForInsert(m, v22, v21, l1, v9, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v92 = v4
					m.G0 = v9 + int32(16)
					return v92
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v30 = F_objectSetKeyAndExpire(m, v28, l1, int64(-1))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_kvstoreHashtableInsertAtPosition(m, v32, v21, v30, v9)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							if base.Ui32(int32(-9)) < base.Ui32(v36) {
							} else {
								v39 = F_lrulfu_init(m)
								mBase = m.M
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
								*(*int32)(unsafe.Add(mBase, uint32(v30))) = v40 | v39<<(uint(int32(8))%32)
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
							if v46&int32(15) != int32(4) {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
								v92 = int32(1)
								m.G0 = v9 + int32(16)
								return v92
							} else {
								v51 = F_hashTypeHasVolatileFields(m, v30)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v51 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
										v92 = int32(1)
										m.G0 = v9 + int32(16)
										return v92
									} else {
										v55 = int32(0)
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
										if v58&int32(2) == v55 {
											v78 = v55
										} else {
											v72 = v30 + (v58&int32(4) ^ int32(12)) + v58<<(uint(int32(3))%32)&int32(8)
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
											v78 = v72 + v73 + int32(1)
										}
										v80 = *(*int32)(unsafe.Add(mBase, _consts[63]))
										if v80 != 0 {
											v82 = F_getKeySlot(m, v78)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = v82
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v86 = F_kvstoreHashtableAdd(m, v85, v84, v30)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
													v92 = int32(1)
													m.G0 = v9 + int32(16)
													return v92
												}
											}
										} else {
											v84 = int32(0)
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v86 = F_kvstoreHashtableAdd(m, v85, v84, v30)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
												v92 = int32(1)
												m.G0 = v9 + int32(16)
												return v92
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
func F_dbExpandExpires(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v7 == int32(0) {
		v35 = F_kvstoreExpand(m, v5, l1, l2, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = v35
			return v37 + int32(-1)
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[90]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)))
		if v14&int32(2) == int32(0) {
			v21 = v13
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2160))
			v24 = v22
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2172))
			if v19 != 0 {
				v21 = v19
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2160))
				v24 = v22
			} else {
				v24 = int32(0)
			}
		}
		if v24 != 0 {
			v28 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v24))
			v30 = F_kvstoreExpand(m, v5, v28, l2, int32(517))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v37 = v30
				return v37 + int32(-1)
			}
		} else {
			return int32(0)
		}
	}
}
func F_dbExpandSkipSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = F_clusterNodeGetPrimary(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = base.I32_div_s(l0, int32(8))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v10+int32(104)))))
		return base.B2i32(int32(base.Ui32(v14)>>(uint(l0&int32(7))%32))&int32(1) == int32(0))
	}
}
func F_dbScan(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(-1)
	v9 = F_kvstoreScan(m, v5, l1, v6, v6, l2, int32(0), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_dbSetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l4 != 0 {
		v25 = l4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a477), int32(_a474), int32(370))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L8
	} else {
		goto L81
	}
L2:
	;
	F__serverAssertWithInfo(m, int32(0), l1, int32(_a478), int32(_a474), int32(325))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L80
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if l3 == int32(0) {
		v44 = v27
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v12 = F_objectGetVal(m, l1)
	mBase = m.M
	v14 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = F_objectGetVal(m, l1)
	mBase = m.M
	v21 = F_kvstoreHashtableFindRef(m, v19, v18, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L10
	}
L6:
	;
	v16 = F_getKeySlot(m, v12)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v18 = int32(0)
	goto L5
L8:
	;
	return
L9:
	;
	v18 = v16
	goto L5
L10:
	;
	if v21 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v25 = v21
	goto L3
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v46&int32(-8) != int32(8) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	F_incrRefCount(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_moduleNotifyKeyUnlink(m, l1, v27, v32, int32(8))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	F_signalDeletedKeyAsReady(m, l0, l1, v36&int32(15))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_decrRefCount(m, v27)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v44 = v43
	goto L12
L18:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v146&int32(15) != int32(4) {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v99 | v45&int32(-256)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v107&int32(1) == int32(0) {
		v118 = int64(-1)
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v54 = int32(base.Ui32(v45)>>(uint(int32(4))%32)) & int32(15)
	if v54 == int32(8) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v57&int32(-8) != int32(8) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v62&int32(240) == int32(128) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v67 = F_objectGetVal(m, v44)
	mBase = m.M
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v74 = v68&int32(-16) | v71&int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v74
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v74&int32(-241) | v78&int32(240)
	v83 = F_objectGetVal(m, v11)
	mBase = m.M
	F_objectSetVal(m, v44, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v45&int32(15) | v88&int32(-256) | v54<<(uint(int32(4))%32)
	F_objectSetVal(m, v11, v67)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v140 = v44
	v142 = v11
	goto L18
L26:
	;
	v119 = F_objectGetVal(m, l1)
	mBase = m.M
	v120 = F_objectSetKeyAndExpire(m, v11, v119, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v44+(v107&int32(4)^int32(12)))))
	v118 = v117
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v120
	if v118 < int64(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v140 = v120
	v142 = v44
	goto L18
L31:
	;
	v125 = F_objectGetVal(m, l1)
	mBase = m.M
	v127 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v127 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = F_objectGetVal(m, l1)
	mBase = m.M
	v134 = F_kvstoreHashtableFindRef(m, v132, v131, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L36
	}
L33:
	;
	v129 = F_getKeySlot(m, v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	v131 = int32(0)
	goto L32
L35:
	;
	v131 = v129
	goto L32
L36:
	;
	if v134 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v120
	goto L30
L38:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v218&int32(15) != int32(4) {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v151 = F_hashTypeHasVolatileFields(m, v142)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	if v151 == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v155&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v158 = v142
	goto L44
L43:
	;
	v158 = v140
	goto L44
L44:
	;
	v159 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v162&int32(2) == v159 {
		v182 = v159
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v184 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L45
L47:
	;
	v176 = v158 + (v162&int32(4) ^ int32(12)) + v162<<(uint(int32(3))%32)&int32(8)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v182 = v176 + v177 + int32(1)
	goto L46
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v190 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v193&int32(2) == v190 {
		v213 = v190
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v186 = F_getKeySlot(m, v182)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L51
	}
L50:
	;
	v188 = int32(0)
	goto L48
L51:
	;
	v188 = v186
	goto L48
L52:
	;
	v214 = F_kvstoreHashtableDelete(m, v189, v188, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v207 = v158 + (v193&int32(4) ^ int32(12)) + v193<<(uint(int32(3))%32)&int32(8)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v213 = v207 + v208 + int32(1)
	goto L53
L55:
	;
	goto L38
L56:
	;
	v262 = int32(-1)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	if v264 < int32(2) {
		v290 = v262
		goto L70
	} else {
		goto L71
	}
L57:
	;
	v223 = F_hashTypeHasVolatileFields(m, v140)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	if v223 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v227 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v230&int32(2) == v227 {
		v250 = v227
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v252 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L60
L62:
	;
	v244 = v140 + (v230&int32(4) ^ int32(12)) + v230<<(uint(int32(3))%32)&int32(8)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v250 = v244 + v245 + int32(1)
	goto L61
L63:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v258 = F_kvstoreHashtableAdd(m, v257, v256, v140)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L67
	}
L64:
	;
	v254 = F_getKeySlot(m, v250)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L66
	}
L65:
	;
	v256 = int32(0)
	goto L63
L66:
	;
	v256 = v254
	goto L63
L67:
	;
	goto L56
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v140
	return
L69:
	;
	if v290 == int32(0) {
		goto L68
	} else {
		goto L75
	}
L70:
	;
	goto L69
L71:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if base.Ui32(int32(15)) < base.Ui32(v267) {
		v290 = v262
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v270 != 0 {
		v290 = v262
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v274 = F_spmcEnqueue(m, int32(_a476), v142|int32(3))
	mBase = m.M
	if v274 == int32(0) {
		v290 = v262
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v277 = int32(0)
	v280 = *(*int32)(unsafe.Add(mBase, _consts[220]))
	*(*int32)(unsafe.Add(mBase, _consts[220])) = v280 + int32(1)
	v284 = int32(_a69)
	v286 = *(*int64)(unsafe.Add(mBase, _consts[221]))
	*(*int64)(unsafe.Add(mBase, _consts[221])) = v286 + int64(1)
	v290 = v277
	goto L70
L75:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	if v294 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_decrRefCount(m, v142)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L8
	} else {
		goto L79
	}
L77:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_freeObjAsync(m, l1, v142, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v140
	return
L79:
	;
	goto L68
L80:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dbSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	if v3 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v8 != 0 {
			v10 = F_hashtableSize(m, v8)
			mBase = m.M
			v13 = base.I64_extend_i32_u(v10)
		} else {
			v13 = int64(0)
		}
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+40))
		v13 = v6
	}
	return v13
}
func F_dbSwapDatabases(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v55 int64
	_ = v55
	var v63 int64
	_ = v63
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v117 int64
	_ = v117
	var v123 int64
	_ = v123
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	v14 = int32(-1)
	if l0 < int32(0) {
		v137 = v14
		m.G0 = v12 + int32(64)
		return v137
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		if v18 <= l1 {
			v137 = v14
			m.G0 = v12 + int32(64)
			return v137
		} else {
			if l1 < int32(0) {
				v137 = v14
				m.G0 = v12 + int32(64)
				return v137
			} else {
				if v18 <= l0 {
					v137 = v14
					m.G0 = v12 + int32(64)
					return v137
				} else {
					v23 = int32(0)
					if l0 == l1 {
						v137 = v23
						m.G0 = v12 + int32(64)
						return v137
					} else {
						v25 = F_createDatabaseIfNeeded(m, l0)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = F_createDatabaseIfNeeded(m, l1)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								v31 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
								v37 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(20))))
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(16)))) = v37
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(28))))
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)))) = v43
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(36))))
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v49
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(44))))
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(40)))) = v55
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(52))))
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(48)))) = v63
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(60))))
								*(*int32)(unsafe.Add(mBase, uint32(v12+int32(56)))) = v71
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v25)+12))
								*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v73
								F_touchAllWatchedKeysInDb(m, v25, v29)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_touchAllWatchedKeysInDb(m, v29, v25)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										F_scanDatabaseForDeletedKeys(m, v25, v29)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											F_scanDatabaseForDeletedKeys(m, v29, v25)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
												*(*int32)(unsafe.Add(mBase, uint32(v25))) = v83
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v85
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v87
												v89 = int32(40)
												v92 = v29 + v89
												v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
												*(*int64)(unsafe.Add(mBase, uint32(v25+v89))) = v93
												v95 = int32(48)
												v98 = v29 + v95
												v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
												*(*int64)(unsafe.Add(mBase, uint32(v25+v95))) = v99
												v101 = int32(56)
												v104 = v29 + v101
												v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
												*(*int64)(unsafe.Add(mBase, uint32(v25+v101))) = v105
												v107 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
												*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v107
												*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v32
												*(*int64)(unsafe.Add(mBase, uint32(v29))) = v31
												v111 = *(*int64)(unsafe.Add(mBase, uint32(v12)+28))
												*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v111
												v117 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(36))))
												*(*int64)(unsafe.Add(mBase, uint32(v92))) = v117
												v123 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(44))))
												*(*int64)(unsafe.Add(mBase, uint32(v98))) = v123
												v129 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(52))))
												*(*int64)(unsafe.Add(mBase, uint32(v104))) = v129
												F_scanDatabaseForReadyKeys(m, v25)
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_scanDatabaseForReadyKeys(m, v29)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return int32(0)
													} else {
														v137 = v23
														m.G0 = v12 + int32(64)
														return v137
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
func F_dbSyncDelete(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v5 = F_objectGetVal(m, l1)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v8 == v3 {
		v15 = v3
		v18 = F_dbGenericDeleteWithDictIndex(m, l0, l1, int32(0), int32(1), v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return v18
		}
	} else {
		v11 = F_getKeySlot(m, v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			v18 = F_dbGenericDeleteWithDictIndex(m, l0, l1, int32(0), int32(1), v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_dbUnshareStringValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v50 int32
	_ = v50
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
	var v75 int32
	_ = v75
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v13&int32(15) != 0 {
		F__serverAssert(m, int32(_a479), int32(_a474), int32(584))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v13&int32(240) != 0 {
			v23 = F_getDecodedObject(m, l2)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_objectGetVal(m, v23)
				mBase = m.M
				v29 = F_objectGetVal(m, v23)
				mBase = m.M
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
				switch v32 & int32(7) {
				case 0:
					v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
				case 1:
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
					v49 = v39
				case 2:
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
					v49 = v42
				case 3:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
					v49 = v45
				case 4:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
					v49 = v48
				default:
					v49 = int32(0)
				}
				v50 = F_createRawStringObject(m, v27, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v50
					F_decrRefCount(m, v23)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v57 = int32(0)
						F_dbSetValue(m, l0, l1, v11+int32(12), v57, v57)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v62 = v61
							m.G0 = v11 + int32(16)
							return v62
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v18&int32(-8) == int32(8) {
				v62 = l2
				m.G0 = v11 + int32(16)
				return v62
			} else {
				v23 = F_getDecodedObject(m, l2)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_objectGetVal(m, v23)
					mBase = m.M
					v29 = F_objectGetVal(m, v23)
					mBase = m.M
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
					switch v32 & int32(7) {
					case 0:
						v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
					case 1:
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
						v49 = v39
					case 2:
						v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
						v49 = v42
					case 3:
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
						v49 = v45
					case 4:
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
						v49 = v48
					default:
						v49 = int32(0)
					}
					v50 = F_createRawStringObject(m, v27, v49)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v50
						F_decrRefCount(m, v23)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v57 = int32(0)
							F_dbSetValue(m, l0, l1, v11+int32(12), v57, v57)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								v62 = v61
								m.G0 = v11 + int32(16)
								return v62
							}
						}
					}
				}
			}
		}
	}
}
func F_dbUntrackKeyWithVolatileItems(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v7&int32(2) == v3 {
		v27 = v3
	} else {
		v21 = l1 + (v7&int32(4) ^ int32(12)) + v7<<(uint(int32(3))%32)&int32(8)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		v27 = v21 + v22 + int32(1)
	}
	v29 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v29 != 0 {
		v31 = F_getKeySlot(m, v27)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v33 = v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v35 = int32(0)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v38&int32(2) == v35 {
				v58 = v35
			} else {
				v52 = l1 + (v38&int32(4) ^ int32(12)) + v38<<(uint(int32(3))%32)&int32(8)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				v58 = v52 + v53 + int32(1)
			}
			v59 = F_kvstoreHashtableDelete(m, v34, v33, v58)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v33 = int32(0)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v35 = int32(0)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v38&int32(2) == v35 {
			v58 = v35
		} else {
			v52 = l1 + (v38&int32(4) ^ int32(12)) + v38<<(uint(int32(3))%32)&int32(8)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
			v58 = v52 + v53 + int32(1)
		}
		v59 = F_kvstoreHashtableDelete(m, v34, v33, v58)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			return
		}
	}
}
func F_db_sethook(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
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
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int64
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v11 + int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v16) < base.Ui32(v17) {
		v59 = m.G398
		if v16 != v59 {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v65 = v62
		} else {
			v65 = int32(-1)
		}
	} else {
		v65 = int32(-1)
	}
	if v65 != int32(8) {
		v130 = int32(0)
		v131 = l0
	} else {
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v78 = v73 + int32(0)
		v79 = m.G398
		if base.Ui32(v78) < base.Ui32(v72) {
			v81 = v78
		} else {
			v81 = v79
		}
		v124 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
		if v124 != int32(8) {
			v128 = int32(0)
		} else {
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
			v128 = v127
		}
		v130 = int32(1)
		v131 = v128
	}
	v132 = int32(1)
	v133 = v130 + v132
	if v133 < v132 {
		if v133 < int32(-9999) {
			switch v130 + int32(10003) {
			case 0:
				v183 = l0 + int32(72)
				v184 = m.G398
				if v183 != v184 {
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
					v190 = v187
				} else {
					v190 = int32(-1)
				}
			case 1:
				v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
				v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v156
				v183 = l0 + int32(88)
				v184 = m.G398
				if v183 != v184 {
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
					v190 = v187
				} else {
					v190 = int32(-1)
				}
			case 2:
				v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v183 = v179 + int32(96)
				v184 = m.G398
				if v183 != v184 {
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
					v190 = v187
				} else {
					v190 = int32(-1)
				}
			default:
				v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
				v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
				v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+7)))
				if base.Ui32(v169) < base.Ui32(int32(-10002)-v133) {
					v190 = int32(-1)
				} else {
					v183 = v168 + (int32(-10003)-v133)<<(uint(int32(4))%32) + int32(24)
					v184 = m.G398
					if v183 != v184 {
						v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
						v190 = v187
					} else {
						v190 = int32(-1)
					}
				}
			}
		} else {
			v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v183 = v147 + v133<<(uint(int32(4))%32)
			v184 = m.G398
			if v183 != v184 {
				v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
				v190 = v187
			} else {
				v190 = int32(-1)
			}
		}
	} else {
		v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v141 = v136 + v133<<(uint(int32(4))%32) + int32(-16)
		v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui32(v141) < base.Ui32(v142) {
			v183 = v141
			v184 = m.G398
			if v183 != v184 {
				v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
				v190 = v187
			} else {
				v190 = int32(-1)
			}
		} else {
			v190 = int32(-1)
		}
	}
	if int32(0) < v190 {
		v226 = F_luaL_checklstring(m, l0, v130|int32(2), int32(0))
		mBase = m.M
		v229 = m.ExcPending
		if v229 != 0 {
			return int32(0)
		} else {
			F_luaL_checktype(m, l0, v133, int32(6))
			mBase = m.M
			v232 = m.ExcPending
			if v232 != 0 {
				return int32(0)
			} else {
				v236 = F_luaL_optinteger(m, l0, v130+int32(3), int32(0))
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return int32(0)
				} else {
					v238 = int32(99)
					v239 = F___strchrnul(m, v226, v238)
					mBase = m.M
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
					if v241 == v238 {
						v245 = v239
					} else {
						v245 = int32(0)
					}
					v246 = int32(0)
					v247 = base.B2i32(v245 != v246)
					v250 = int32(114)
					v251 = F___strchrnul(m, v226, v250)
					mBase = m.M
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
					if v253 == v250 {
						v257 = v251
					} else {
						v257 = v246
					}
					if v257 != 0 {
						v258 = v247 | int32(2)
					} else {
						v258 = v247
					}
					v261 = int32(108)
					v262 = F___strchrnul(m, v226, v261)
					mBase = m.M
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
					if v264 == v261 {
						v268 = v262
					} else {
						v268 = int32(0)
					}
					if v268 != 0 {
						v269 = v258 | int32(4)
					} else {
						v269 = v258
					}
					if int32(0) < v236 {
						v274 = v269 | int32(8)
					} else {
						v274 = v269
					}
					v275 = m.G5
					v278 = v236
					v279 = v275 + int32(1274)
					v280 = v274
					F_gethooktable(m, l0)
					mBase = m.M
					v282 = m.ExcPending
					if v282 != 0 {
						return int32(0)
					} else {
						v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v284)+8)) = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v284))) = v131
						v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v288 + int32(16)
						if v133 < int32(1) {
							if v133 < int32(-9999) {
								switch v130 + int32(10003) {
								case 0:
									v344 = l0 + int32(72)
								case 1:
									v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
									v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
									v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v320
									v344 = l0 + int32(88)
								case 2:
									v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v344 = v314 + int32(96)
								default:
									v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
									v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
									v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+7)))
									v332 = m.G398
									if base.Ui32(v331) < base.Ui32(int32(-10002)-v133) {
										v343 = v332
									} else {
										v343 = v330 + (int32(-10003)-v133)<<(uint(int32(4))%32) + int32(24)
									}
									v344 = v343
								}
							} else {
								v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v344 = v308 + v133<<(uint(int32(4))%32)
							}
						} else {
							v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v302 = v297 + v133<<(uint(int32(4))%32) + int32(-16)
							v303 = m.G398
							if base.Ui32(v302) < base.Ui32(v296) {
								v305 = v302
							} else {
								v305 = v303
							}
							v344 = v305
						}
						v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v348 = *(*int64)(unsafe.Add(mBase, uint32(v344)))
						*(*int64)(unsafe.Add(mBase, uint32(v347))) = v348
						v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = v350
						v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v352 + int32(16)
						F_lua_rawset(m, l0, int32(-3))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
							return int32(0)
						} else {
							v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v377 + int32(-16)
							*(*int32)(unsafe.Add(mBase, uint32(v131)+64)) = v278
							*(*int32)(unsafe.Add(mBase, uint32(v131)+60)) = v278
							if v280 != 0 {
								v390 = v279
							} else {
								v390 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v131)+68)) = v390
							if v279 != 0 {
								v393 = v280
							} else {
								v393 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v131)+56)) = uint8(v393)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		if v133 < int32(0) {
			v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v217 = v210 + v133<<(uint(int32(4))%32) + int32(16)
		} else {
			v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v200 = v197 + v133<<(uint(int32(4))%32)
			if base.Ui32(v200) <= base.Ui32(v196) {
				v217 = v200
			} else {
				v204 = v196
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v204)+8)) = int32(0)
					v208 = v204 + int32(16)
					if base.Ui32(v208) < base.Ui32(v200) {
						v204 = v208
						continue
					} else {
						break
					}
					break
				}
				v217 = v200
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v217
		v220 = int32(0)
		v278 = v220
		v279 = v220
		v280 = v220
		F_gethooktable(m, l0)
		mBase = m.M
		v282 = m.ExcPending
		if v282 != 0 {
			return int32(0)
		} else {
			v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v284)+8)) = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v284))) = v131
			v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v288 + int32(16)
			if v133 < int32(1) {
				if v133 < int32(-9999) {
					switch v130 + int32(10003) {
					case 0:
						v344 = l0 + int32(72)
					case 1:
						v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
						v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
						v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v320
						v344 = l0 + int32(88)
					case 2:
						v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v344 = v314 + int32(96)
					default:
						v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
						v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
						v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+7)))
						v332 = m.G398
						if base.Ui32(v331) < base.Ui32(int32(-10002)-v133) {
							v343 = v332
						} else {
							v343 = v330 + (int32(-10003)-v133)<<(uint(int32(4))%32) + int32(24)
						}
						v344 = v343
					}
				} else {
					v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v344 = v308 + v133<<(uint(int32(4))%32)
				}
			} else {
				v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v302 = v297 + v133<<(uint(int32(4))%32) + int32(-16)
				v303 = m.G398
				if base.Ui32(v302) < base.Ui32(v296) {
					v305 = v302
				} else {
					v305 = v303
				}
				v344 = v305
			}
			v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v348 = *(*int64)(unsafe.Add(mBase, uint32(v344)))
			*(*int64)(unsafe.Add(mBase, uint32(v347))) = v348
			v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v347)+8)) = v350
			v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v352 + int32(16)
			F_lua_rawset(m, l0, int32(-3))
			mBase = m.M
			v358 = m.ExcPending
			if v358 != 0 {
				return int32(0)
			} else {
				v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v377 + int32(-16)
				*(*int32)(unsafe.Add(mBase, uint32(v131)+64)) = v278
				*(*int32)(unsafe.Add(mBase, uint32(v131)+60)) = v278
				if v280 != 0 {
					v390 = v279
				} else {
					v390 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v131)+68)) = v390
				if v279 != 0 {
					v393 = v280
				} else {
					v393 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v131)+56)) = uint8(v393)
				return int32(0)
			}
		}
	}
}
func F_db_setlocal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
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
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	goto L5
L1:
	;
	v136 = v133 + int32(1)
	v137 = F_luaL_checkinteger(m, l0, v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	if v68 != int32(8) {
		v133 = int32(0)
		v134 = l0
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v62 = m.G398
	if v19 != v62 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v14 + int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v19) < base.Ui32(v20) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v68 = int32(-1)
	goto L2
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v68 = v65
	goto L2
L16:
	;
	v68 = int32(-1)
	goto L2
L17:
	;
	goto L21
L18:
	;
	v133 = int32(1)
	v134 = v131
	goto L1
L19:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v127 != int32(8) {
		v131 = int32(0)
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = v76 + int32(0)
	v82 = m.G398
	if base.Ui32(v81) < base.Ui32(v75) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v81
	goto L24
L23:
	;
	v84 = v82
	goto L24
L24:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v131 = v130
	goto L34
L36:
	;
	m.G0 = v8 + int32(112)
	return v422
L37:
	;
	v200 = v133 + int32(3)
	F_luaL_checkany(m, l0, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L38
	} else {
		goto L55
	}
L38:
	;
	return int32(0)
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v137 < int32(1) {
		v171 = v137
		v173 = v146
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v193 != 0 {
		goto L37
	} else {
		goto L53
	}
L41:
	;
	goto L40
L42:
	;
	if v171 != 0 {
		v184 = int32(0)
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v152 = v137
	v154 = v146
	goto L44
L44:
	;
	if base.Ui32(v154) <= base.Ui32(v149) {
		v193 = int32(0)
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v171 = v165
	v173 = v167
	goto L42
L46:
	;
	v159 = v152 + int32(-1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+6)))
	if v162 != 0 {
		v165 = v159
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v167 = v154 + int32(-24)
	if int32(0) < v165 {
		v152 = v165
		v154 = v167
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v165 = v159 - v163
	goto L47
L49:
	;
	goto L45
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12))+96)) = v184
	v193 = int32(1)
	goto L41
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	if base.Ui32(v173) <= base.Ui32(v178) {
		v193 = int32(0)
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v182 = base.I32_div_s(v173-v178, int32(24))
	v184 = v182
	goto L50
L53:
	;
	v194 = m.G3
	v197 = F_luaL_argerror(m, l0, v136, v194+int32(_a2314))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v422 = v197
	goto L36
L55:
	;
	if v200 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if l0 == v134 {
		goto L65
	} else {
		goto L66
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v227
	goto L56
L58:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v227 = v220 + v200<<(uint(int32(4))%32) + int32(16)
	goto L57
L59:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v210 = v207 + v200<<(uint(int32(4))%32)
	if base.Ui32(v210) <= base.Ui32(v206) {
		v227 = v210
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v214 = v206
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = int32(0)
	v218 = v214 + int32(16)
	if base.Ui32(v218) < base.Ui32(v210) {
		v214 = v218
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v227 = v210
	goto L57
L64:
	;
	v331 = F_luaL_checkinteger(m, l0, v133|int32(2))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L38
	} else {
		goto L74
	}
L65:
	;
	goto L64
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v239 - int32(16)
	goto L67
L67:
	;
	goto L68
L68:
	;
	goto L73
L73:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v307 + int32(16)
	v313 = v306 + int32(0)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v313)))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+8)) = v316
	goto L65
L74:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(12))+96))
	v340 = v336 + v337*int32(24)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+8))
	if v342 != int32(6) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	F_lua_pushstring(m, l0, v413)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L38
	} else {
		goto L91
	}
L76:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v416 + int32(-16)
	goto L75
L77:
	;
	v397 = v391 + v331<<(uint(int32(4))%32)
	v398 = int32(-16)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v400+v398)))
	*(*int64)(unsafe.Add(mBase, uint32(v397+v398))) = v403
	v405 = int32(-8)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v400+v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v397+v405))) = v409
	v413 = v392
	goto L76
L78:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v340 == v376 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+6)))
	if v346 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
	if v347 == int32(0) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v340 == v350 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v365 = F_luaF_getlocalname(m, v347, v331, (v358-v359)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v365 == int32(0) {
		goto L78
	} else {
		goto L85
	}
L83:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+12)) = v353
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v357 = v356
	v358 = v353
	goto L82
L84:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v357 = v347
	v358 = v352
	goto L82
L85:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v391 = v368
	v392 = v365
	goto L77
L86:
	;
	v378 = v134 + int32(8)
	goto L88
L87:
	;
	v378 = v340 + int32(28)
	goto L88
L88:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v381 = m.G3
	v382 = int32(0)
	if v331 < int32(1) {
		v413 = v382
		goto L76
	} else {
		goto L89
	}
L89:
	;
	if (v379-v380)>>(uint(int32(4))%32) < v331 {
		v413 = v382
		goto L76
	} else {
		goto L90
	}
L90:
	;
	v391 = v380
	v392 = v381 + int32(_a2242)
	goto L77
L91:
	;
	v422 = int32(1)
	goto L36
}
func F_db_setupvalue(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	F_luaL_checkany(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_auxupvalue(m, l0, int32(0))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_selectDb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = int32(-1)
	if l1 < int32(0) {
		v16 = v4
		return v16
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		if v8 <= l1 {
			v16 = v4
			return v16
		} else {
			v10 = F_createDatabaseIfNeeded(m, l1)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v10
				v16 = int32(0)
				return v16
			}
		}
	}
}
