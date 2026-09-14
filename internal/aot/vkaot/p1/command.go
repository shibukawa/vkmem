package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addCommandToBatchAndProcessIfFull(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_addCommandToBatchAndProcessIfFull[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12 + int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v12<<(uint(int32(2))%32)))) = l0
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v21 == int32(0) {
		v40 = v9
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(-1)
L3:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v43) <= base.Ui32(v42) {
		v86 = v40
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v24&int32(262144) != 0 {
		v40 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v24 | int32(2097152)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_addCommandToBatch(m, v21, v30, v31, v32, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_addCommandToBatchAndProcessIfFull[0]))
	v40 = v39
	goto L3
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	if v92 == v93 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v46 = v40
	v47 = v42
	v49 = v43
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if base.Ui32(v53) <= base.Ui32(v52) {
		v86 = v46
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v86 = v77
	goto L8
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v58 = v55 + v47*int32(40)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+32))
	if v59 == int32(0) {
		v77 = v46
		v78 = v49
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = v47 + int32(1)
	if base.Ui32(v81) < base.Ui32(v78&int32(65535)) {
		v46 = v77
		v47 = v81
		v49 = v78
		goto L10
	} else {
		goto L17
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62&int32(262144) != 0 {
		v77 = v46
		v78 = v49
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v62 | int32(2097152)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	F_addCommandToBatch(m, v59, v68, v69, v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_addCommandToBatchAndProcessIfFull[0]))
	v77 = v76
	v78 = v74
	goto L13
L17:
	;
	goto L11
L18:
	;
	return int32(0)
L19:
	;
	F_processClientsCommandsBatch(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	if v95 != v93 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L18
}
func F_canParseCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v4 != 0 {
		v37 = v2
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v5&int32(144) != 0 {
			v37 = v2
		} else {
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
			if v8&int32(2) != 0 {
				v37 = v2
			} else {
				v11 = F_scriptIsTimedout(m)
				mBase = m.M
				v12 = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, _c_F_canParseCommand[0]))
				if base.B2i32(v11|v13 != v12) == int32(0) {
					v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
					v37 = base.B2i32(v32&int32(1088) == int32(0))
				} else {
					v20 = int32(1)
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
					if v21&v20 != 0 {
						v28 = v20
						v31 = v28
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v24 != 0 {
							v26 = F_isImportSlotMigrationJob(m, v24)
							mBase = m.M
							v28 = v26
							v31 = v28
						} else {
							v31 = int32(0)
						}
					}
					if v31 != 0 {
						v37 = v2
					} else {
						v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+200)))
						v37 = base.B2i32(v32&int32(1088) == int32(0))
					}
				}
			}
		}
	}
	return v37
}
func F_commandAddSubcommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = l0
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
		v10 = F_ACLGetCommandID(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = v10
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			v14 = F_hashtableAdd(m, v13, l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				if v14 != 0 {
					return
				} else {
					F__serverAssert(m, int32(_a_F_commandAddSubcommand_0), int32(_a_F_commandAddSubcommand_1), int32(3309))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
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
	} else {
		v5 = F_hashtableCreate(m, int32(_a_F_commandAddSubcommand_2))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v5
			*(*int32)(unsafe.Add(mBase, uint32(l1)+204)) = l0
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
			v10 = F_ACLGetCommandID(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = v10
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				v14 = F_hashtableAdd(m, v13, l1)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					if v14 != 0 {
						return
					} else {
						F__serverAssert(m, int32(_a_F_commandAddSubcommand_0), int32(_a_F_commandAddSubcommand_1), int32(3309))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
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
		}
	}
}
func F_commandCheckExistence(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v300 int32
	_ = v300
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v300
L2:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v300 = int32(1)
	goto L1
L4:
	;
	v300 = int32(0)
	goto L1
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = F_objectGetVal(m, v20)
	mBase = m.M
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_commandCheckExistence[0]))
	v26 = F_hashtableFind(m, v23, v21, v11+int32(44))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v190
	F_sdsfree(m, v193)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L49
	}
L7:
	;
	v99 = F_sdsempty(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L28
	}
L8:
	;
	return int32(0)
L9:
	;
	if v26 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 < int32(2) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = F_objectGetVal(m, v40)
	mBase = m.M
	v42 = F_sdsnew(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
	switch v49 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		goto L15
	}
L14:
	;
	v86 = F_sdsnew(m, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L26
	}
L15:
	;
	goto L14
L16:
	;
	if v66 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
	v66 = v65
	goto L16
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
	v66 = v62
	goto L16
L19:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
	v66 = v59
	goto L16
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
	v66 = v56
	goto L16
L21:
	;
	v66 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v71 = int32(0)
	goto L23
L23:
	;
	v74 = v42 + v71
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v74))))
	v76 = F_toupper(m, v75)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v76)
	v79 = v71 + int32(1)
	if v79 != v66 {
		v71 = v79
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L15
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = F_objectGetVal(m, v90)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v91
	v97 = F_sdscatprintf(m, v86, int32(_a_F_commandCheckExistence_0), v11+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v190 = v97
	v193 = v42
	goto L6
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v101 < int32(2) {
		v173 = v99
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v179 = F_sdsnew(m, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L47
	}
L30:
	;
	v108 = v99
	v109 = int32(1)
	goto L31
L31:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-1)))))
	v118 = v116 & int32(7)
	switch v118 {
	case 0:
		goto L38
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	default:
		v150 = int32(0)
		goto L33
	}
L32:
	;
	v173 = v164
	goto L29
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v109<<(uint(int32(2))%32))))
	v156 = F_objectGetVal(m, v155)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(128) - v150
	v164 = F_sdscatprintf(m, v108, int32(_a_F_commandCheckExistence_1), v11+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L45
	}
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-17))))
	v150 = v149
	goto L33
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-9))))
	v150 = v146
	goto L33
L36:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+int32(-5)))))
	v150 = v143
	goto L33
L37:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-3)))))
	v150 = v140
	goto L33
L38:
	;
	v150 = int32(base.Ui32(v116) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	if base.Ui32(int32(127)) < base.Ui32(v131) {
		v173 = v108
		goto L29
	} else {
		goto L44
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-17))))
	v131 = v130
	goto L39
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-9))))
	v131 = v127
	goto L39
L42:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+int32(-5)))))
	v131 = v124
	goto L39
L43:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+int32(-3)))))
	v131 = v121
	goto L39
L44:
	;
	switch v118 + int32(-1) {
	default:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	}
L45:
	;
	v167 = v109 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v167 < v168 {
		v108 = v164
		v109 = v167
		goto L31
	} else {
		goto L46
	}
L46:
	;
	goto L32
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v184 = F_objectGetVal(m, v183)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v184
	v188 = F_sdscatprintf(m, v179, int32(_a_F_commandCheckExistence_2), v11)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v190 = v188
	v193 = v173
	goto L6
L49:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-1)))))
	switch v212 & int32(7) {
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
		goto L51
	}
L50:
	;
	goto L4
L51:
	;
	goto L50
L52:
	;
	if v229 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L53:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-17))))
	v229 = v228
	goto L52
L54:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-9))))
	v229 = v225
	goto L52
L55:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+int32(-5)))))
	v229 = v222
	goto L52
L56:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201+int32(-3)))))
	v229 = v219
	goto L52
L57:
	;
	v229 = int32(base.Ui32(v212) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v239 = int32(0)
	goto L59
L59:
	;
	goto L62
L60:
	;
	goto L51
L61:
	;
	v277 = v239 + int32(1)
	if v277 != v229 {
		v239 = v277
		goto L59
	} else {
		goto L68
	}
L62:
	;
	v245 = v201 + v239
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v253 = int32(0)
	goto L63
L63:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+uint32(_c_F_commandCheckExistence[1]))))
	if v246&int32(255) != v259 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L61
L65:
	;
	v265 = v253 + int32(1)
	if v265 != int32(2) {
		v253 = v265
		goto L63
	} else {
		goto L67
	}
L66:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+uint32(_c_F_commandCheckExistence[2]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v262)
	goto L61
L67:
	;
	goto L64
L68:
	;
	goto L60
}
func F_commandCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v14 = base.B2i32(v10 == int32(3)) << (uint(int32(2)) % 32)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_commandCommand[0])))
	if v17 != 0 {
		v21 = int32(0)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_commandCommand[1]))
		if v22 == v21 {
			v96 = v17
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-1)))))
			switch v100 & int32(7) {
			case 0:
				v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
			case 1:
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-3)))))
				v117 = v107
			case 2:
				v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+int32(-5)))))
				v117 = v110
			case 3:
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-9))))
				v117 = v113
			case 4:
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-17))))
				v117 = v116
			default:
				v117 = int32(0)
			}
			F_addReplyProto(m, l0, v96, v117)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			v25 = F_generateCommandResponse(m, v10)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = int32(0)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-1)))))
				switch v34 & int32(7) {
				case 0:
					v51 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
				case 1:
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+int32(-3)))))
					v51 = v41
				case 2:
					v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+int32(-5)))))
					v51 = v44
				case 3:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-9))))
					v51 = v47
				case 4:
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(-17))))
					v51 = v50
				default:
					v51 = v27
				}
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-1)))))
				switch v54 & int32(7) {
				case 0:
					v71 = int32(base.Ui32(v54) >> (uint(int32(3)) % 32))
				case 1:
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(-3)))))
					v71 = v61
				case 2:
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(-5)))))
					v71 = v64
				case 3:
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-9))))
					v71 = v67
				case 4:
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(-17))))
					v71 = v70
				default:
					v71 = v27
				}
				v72 = base.B2i32(base.Ui32(v51) < base.Ui32(v71))
				if base.Ui32(v51) < base.Ui32(v71) {
					v73 = v51
				} else {
					v73 = v71
				}
				v74 = F_memcmp(m, v25, v17, v73)
				mBase = m.M
				if v74 != 0 {
					v77 = v74
				} else {
					v77 = base.B2i32(base.Ui32(v71) < base.Ui32(v51)) - v72
				}
				if v77 == int32(0) {
					F_sdsfree(m, v25)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						if v77 != 0 {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_commandCommand_0), int32(_a_F_commandCommand_1), int32(5585))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v96 = v17
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-1)))))
							switch v100 & int32(7) {
							case 0:
								v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
							case 1:
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-3)))))
								v117 = v107
							case 2:
								v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+int32(-5)))))
								v117 = v110
							case 3:
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-9))))
								v117 = v113
							case 4:
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-17))))
								v117 = v116
							default:
								v117 = int32(0)
							}
							F_addReplyProto(m, l0, v96, v117)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, _c_F_commandCommand[2]))
					if int32(3) < v81 {
						F_sdsfree(m, v25)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							if v77 != 0 {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_commandCommand_0), int32(_a_F_commandCommand_1), int32(5585))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v96 = v17
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-1)))))
								switch v100 & int32(7) {
								case 0:
									v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
								case 1:
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-3)))))
									v117 = v107
								case 2:
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+int32(-5)))))
									v117 = v110
								case 3:
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-9))))
									v117 = v113
								case 4:
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-17))))
									v117 = v116
								default:
									v117 = int32(0)
								}
								F_addReplyProto(m, l0, v96, v117)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
						F__serverLog(m, int32(3), int32(_a_F_commandCommand_2), v8)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							F_sdsfree(m, v25)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_commandCommand_0), int32(_a_F_commandCommand_1), int32(5585))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
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
				}
			}
		}
	} else {
		v18 = F_generateCommandResponse(m, v10)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_commandCommand[0]))) = v18
			v96 = v18
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-1)))))
			switch v100 & int32(7) {
			case 0:
				v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
			case 1:
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(-3)))))
				v117 = v107
			case 2:
				v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96+int32(-5)))))
				v117 = v110
			case 3:
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-9))))
				v117 = v113
			case 4:
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(-17))))
				v117 = v116
			default:
				v117 = int32(0)
			}
			F_addReplyProto(m, l0, v96, v117)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_commandFlagsFromString(m *base.Module, l0 int32) int64 {
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
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
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
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
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
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
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
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
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
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
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
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
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int64
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v998 int64
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1009 int64
	_ = v1009
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
	v70 = F_sdssplitlen(m, l0, v65, int32(_a_F_commandFlagsFromString_0), int32(1), v9+int32(12))
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
		v995 = v75
		v998 = v74
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_sdsfreesplitres(m, v70, v76)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L17
	} else {
		goto L331
	}
L20:
	;
	v79 = v75
	v82 = v74
	goto L21
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70+v79<<(uint(int32(2))%32))))
	v89 = int32(_a_F_commandFlagsFromString_1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v995 = v76
	v998 = v991
	goto L19
L23:
	;
	v993 = v79 + int32(1)
	if v993 != v76 {
		v79 = v993
		v82 = v991
		goto L21
	} else {
		goto L330
	}
L24:
	;
	v130 = int32(_a_F_commandFlagsFromString_2)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	if v124-v126 != 0 {
		goto L24
	} else {
		goto L37
	}
L26:
	;
	v124 = F_tolower(m, v120)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	goto L25
L27:
	;
	v94 = v88
	v95 = v89
	v96 = v92
	goto L30
L28:
	;
	v120 = int32(0)
	v121 = v89
	goto L26
L29:
	;
	v120 = v117 & int32(255)
	v121 = v116
	goto L26
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v98 == int32(0) {
		v116 = v95
		v117 = v96
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v116 = v110
	v117 = int32(0)
	goto L29
L32:
	;
	v102 = v96 & int32(255)
	if v102 == v98 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = int32(1)
	v110 = v95 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v111 != 0 {
		v94 = v94 + v109
		v95 = v110
		v96 = v111
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v104 = F_tolower(m, v102)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	if v104 == v106 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v116 = v95
	v117 = v108
	goto L29
L36:
	;
	goto L31
L37:
	;
	v991 = v82 | int64(1)
	goto L23
L38:
	;
	v171 = int32(_a_F_commandFlagsFromString_3)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v174 != 0 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	if v165-v167 != 0 {
		goto L38
	} else {
		goto L51
	}
L40:
	;
	v165 = F_tolower(m, v161)
	mBase = m.M
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v167 = F_tolower(m, v166)
	mBase = m.M
	goto L39
L41:
	;
	v135 = v88
	v136 = v130
	v137 = v133
	goto L44
L42:
	;
	v161 = int32(0)
	v162 = v130
	goto L40
L43:
	;
	v161 = v158 & int32(255)
	v162 = v157
	goto L40
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v139 == int32(0) {
		v157 = v136
		v158 = v137
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v157 = v151
	v158 = int32(0)
	goto L43
L46:
	;
	v143 = v137 & int32(255)
	if v143 == v139 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v150 = int32(1)
	v151 = v136 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v152 != 0 {
		v135 = v135 + v150
		v136 = v151
		v137 = v152
		goto L44
	} else {
		goto L50
	}
L48:
	;
	v145 = F_tolower(m, v143)
	mBase = m.M
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v147 = F_tolower(m, v146)
	mBase = m.M
	if v145 == v147 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v157 = v136
	v158 = v149
	goto L43
L50:
	;
	goto L45
L51:
	;
	v991 = v82 | int64(2)
	goto L23
L52:
	;
	v212 = int32(_a_F_commandFlagsFromString_4)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v215 != 0 {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	if v206-v208 != 0 {
		goto L52
	} else {
		goto L65
	}
L54:
	;
	v206 = F_tolower(m, v202)
	mBase = m.M
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v208 = F_tolower(m, v207)
	mBase = m.M
	goto L53
L55:
	;
	v176 = v88
	v177 = v171
	v178 = v174
	goto L58
L56:
	;
	v202 = int32(0)
	v203 = v171
	goto L54
L57:
	;
	v202 = v199 & int32(255)
	v203 = v198
	goto L54
L58:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v180 == int32(0) {
		v198 = v177
		v199 = v178
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v198 = v192
	v199 = int32(0)
	goto L57
L60:
	;
	v184 = v178 & int32(255)
	if v184 == v180 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v191 = int32(1)
	v192 = v177 + v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	if v193 != 0 {
		v176 = v176 + v191
		v177 = v192
		v178 = v193
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v186 = F_tolower(m, v184)
	mBase = m.M
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v188 = F_tolower(m, v187)
	mBase = m.M
	if v186 == v188 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v198 = v177
	v199 = v190
	goto L57
L64:
	;
	goto L59
L65:
	;
	v991 = v82 | int64(16)
	goto L23
L66:
	;
	v253 = int32(_a_F_commandFlagsFromString_5)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v256 != 0 {
		goto L83
	} else {
		goto L84
	}
L67:
	;
	if v247-v249 != 0 {
		goto L66
	} else {
		goto L79
	}
L68:
	;
	v247 = F_tolower(m, v243)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v249 = F_tolower(m, v248)
	mBase = m.M
	goto L67
L69:
	;
	v217 = v88
	v218 = v212
	v219 = v215
	goto L72
L70:
	;
	v243 = int32(0)
	v244 = v212
	goto L68
L71:
	;
	v243 = v240 & int32(255)
	v244 = v239
	goto L68
L72:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v221 == int32(0) {
		v239 = v218
		v240 = v219
		goto L71
	} else {
		goto L74
	}
L73:
	;
	v239 = v233
	v240 = int32(0)
	goto L71
L74:
	;
	v225 = v219 & int32(255)
	if v225 == v221 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v232 = int32(1)
	v233 = v218 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v234 != 0 {
		v217 = v217 + v232
		v218 = v233
		v219 = v234
		goto L72
	} else {
		goto L78
	}
L76:
	;
	v227 = F_tolower(m, v225)
	mBase = m.M
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v229 = F_tolower(m, v228)
	mBase = m.M
	if v227 == v229 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v239 = v218
	v240 = v231
	goto L71
L78:
	;
	goto L73
L79:
	;
	v991 = v82 | int64(4)
	goto L23
L80:
	;
	v294 = int32(_a_F_commandFlagsFromString_6)
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v297 != 0 {
		goto L97
	} else {
		goto L98
	}
L81:
	;
	if v288-v290 != 0 {
		goto L80
	} else {
		goto L93
	}
L82:
	;
	v288 = F_tolower(m, v284)
	mBase = m.M
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v290 = F_tolower(m, v289)
	mBase = m.M
	goto L81
L83:
	;
	v258 = v88
	v259 = v253
	v260 = v256
	goto L86
L84:
	;
	v284 = int32(0)
	v285 = v253
	goto L82
L85:
	;
	v284 = v281 & int32(255)
	v285 = v280
	goto L82
L86:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v262 == int32(0) {
		v280 = v259
		v281 = v260
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v280 = v274
	v281 = int32(0)
	goto L85
L88:
	;
	v266 = v260 & int32(255)
	if v266 == v262 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v273 = int32(1)
	v274 = v259 + v273
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	if v275 != 0 {
		v258 = v258 + v273
		v259 = v274
		v260 = v275
		goto L86
	} else {
		goto L92
	}
L90:
	;
	v268 = F_tolower(m, v266)
	mBase = m.M
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v270 = F_tolower(m, v269)
	mBase = m.M
	if v268 == v270 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v280 = v259
	v281 = v272
	goto L85
L92:
	;
	goto L87
L93:
	;
	v991 = v82 | int64(64)
	goto L23
L94:
	;
	v335 = int32(_a_F_commandFlagsFromString_7)
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v338 != 0 {
		goto L111
	} else {
		goto L112
	}
L95:
	;
	if v329-v331 != 0 {
		goto L94
	} else {
		goto L107
	}
L96:
	;
	v329 = F_tolower(m, v325)
	mBase = m.M
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	v331 = F_tolower(m, v330)
	mBase = m.M
	goto L95
L97:
	;
	v299 = v88
	v300 = v294
	v301 = v297
	goto L100
L98:
	;
	v325 = int32(0)
	v326 = v294
	goto L96
L99:
	;
	v325 = v322 & int32(255)
	v326 = v321
	goto L96
L100:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v303 == int32(0) {
		v321 = v300
		v322 = v301
		goto L99
	} else {
		goto L102
	}
L101:
	;
	v321 = v315
	v322 = int32(0)
	goto L99
L102:
	;
	v307 = v301 & int32(255)
	if v307 == v303 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v314 = int32(1)
	v315 = v300 + v314
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+1)))
	if v316 != 0 {
		v299 = v299 + v314
		v300 = v315
		v301 = v316
		goto L100
	} else {
		goto L106
	}
L104:
	;
	v309 = F_tolower(m, v307)
	mBase = m.M
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v311 = F_tolower(m, v310)
	mBase = m.M
	if v309 == v311 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v321 = v300
	v322 = v313
	goto L99
L106:
	;
	goto L101
L107:
	;
	v991 = v82 | int64(512)
	goto L23
L108:
	;
	v376 = int32(_a_F_commandFlagsFromString_8)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v379 != 0 {
		goto L124
	} else {
		goto L125
	}
L109:
	;
	if v370-v372 != 0 {
		goto L108
	} else {
		goto L121
	}
L110:
	;
	v370 = F_tolower(m, v366)
	mBase = m.M
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	v372 = F_tolower(m, v371)
	mBase = m.M
	goto L109
L111:
	;
	v340 = v88
	v341 = v335
	v342 = v338
	goto L114
L112:
	;
	v366 = int32(0)
	v367 = v335
	goto L110
L113:
	;
	v366 = v363 & int32(255)
	v367 = v362
	goto L110
L114:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v344 == int32(0) {
		v362 = v341
		v363 = v342
		goto L113
	} else {
		goto L116
	}
L115:
	;
	v362 = v356
	v363 = int32(0)
	goto L113
L116:
	;
	v348 = v342 & int32(255)
	if v348 == v344 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v355 = int32(1)
	v356 = v341 + v355
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if v357 != 0 {
		v340 = v340 + v355
		v341 = v356
		v342 = v357
		goto L114
	} else {
		goto L120
	}
L118:
	;
	v350 = F_tolower(m, v348)
	mBase = m.M
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	v352 = F_tolower(m, v351)
	mBase = m.M
	if v350 == v352 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	v362 = v341
	v363 = v354
	goto L113
L120:
	;
	goto L115
L121:
	;
	v991 = v82 | int64(32)
	goto L23
L122:
	;
	if v411-v413 == int32(0) {
		v991 = v82
		goto L23
	} else {
		goto L134
	}
L123:
	;
	v411 = F_tolower(m, v407)
	mBase = m.M
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	v413 = F_tolower(m, v412)
	mBase = m.M
	goto L122
L124:
	;
	v381 = v88
	v382 = v376
	v383 = v379
	goto L127
L125:
	;
	v407 = int32(0)
	v408 = v376
	goto L123
L126:
	;
	v407 = v404 & int32(255)
	v408 = v403
	goto L123
L127:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v385 == int32(0) {
		v403 = v382
		v404 = v383
		goto L126
	} else {
		goto L129
	}
L128:
	;
	v403 = v397
	v404 = int32(0)
	goto L126
L129:
	;
	v389 = v383 & int32(255)
	if v389 == v385 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v396 = int32(1)
	v397 = v382 + v396
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
	if v398 != 0 {
		v381 = v381 + v396
		v382 = v397
		v383 = v398
		goto L127
	} else {
		goto L133
	}
L131:
	;
	v391 = F_tolower(m, v389)
	mBase = m.M
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	v393 = F_tolower(m, v392)
	mBase = m.M
	if v391 == v393 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v403 = v382
	v404 = v395
	goto L126
L133:
	;
	goto L128
L134:
	;
	v417 = int32(_a_F_commandFlagsFromString_9)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v420 != 0 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v458 = int32(_a_F_commandFlagsFromString_10)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v461 != 0 {
		goto L152
	} else {
		goto L153
	}
L136:
	;
	if v452-v454 != 0 {
		goto L135
	} else {
		goto L148
	}
L137:
	;
	v452 = F_tolower(m, v448)
	mBase = m.M
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v454 = F_tolower(m, v453)
	mBase = m.M
	goto L136
L138:
	;
	v422 = v88
	v423 = v417
	v424 = v420
	goto L141
L139:
	;
	v448 = int32(0)
	v449 = v417
	goto L137
L140:
	;
	v448 = v445 & int32(255)
	v449 = v444
	goto L137
L141:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	if v426 == int32(0) {
		v444 = v423
		v445 = v424
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v444 = v438
	v445 = int32(0)
	goto L140
L143:
	;
	v430 = v424 & int32(255)
	if v430 == v426 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v437 = int32(1)
	v438 = v423 + v437
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	if v439 != 0 {
		v422 = v422 + v437
		v423 = v438
		v424 = v439
		goto L141
	} else {
		goto L147
	}
L145:
	;
	v432 = F_tolower(m, v430)
	mBase = m.M
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	v434 = F_tolower(m, v433)
	mBase = m.M
	if v432 == v434 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422))))
	v444 = v423
	v445 = v436
	goto L140
L147:
	;
	goto L142
L148:
	;
	v991 = v82 | int64(256)
	goto L23
L149:
	;
	v499 = int32(_a_F_commandFlagsFromString_11)
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v502 != 0 {
		goto L166
	} else {
		goto L167
	}
L150:
	;
	if v493-v495 != 0 {
		goto L149
	} else {
		goto L162
	}
L151:
	;
	v493 = F_tolower(m, v489)
	mBase = m.M
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490))))
	v495 = F_tolower(m, v494)
	mBase = m.M
	goto L150
L152:
	;
	v463 = v88
	v464 = v458
	v465 = v461
	goto L155
L153:
	;
	v489 = int32(0)
	v490 = v458
	goto L151
L154:
	;
	v489 = v486 & int32(255)
	v490 = v485
	goto L151
L155:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	if v467 == int32(0) {
		v485 = v464
		v486 = v465
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v485 = v479
	v486 = int32(0)
	goto L154
L157:
	;
	v471 = v465 & int32(255)
	if v471 == v467 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v478 = int32(1)
	v479 = v464 + v478
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	if v480 != 0 {
		v463 = v463 + v478
		v464 = v479
		v465 = v480
		goto L155
	} else {
		goto L161
	}
L159:
	;
	v473 = F_tolower(m, v471)
	mBase = m.M
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	v475 = F_tolower(m, v474)
	mBase = m.M
	if v473 == v475 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v485 = v464
	v486 = v477
	goto L154
L161:
	;
	goto L156
L162:
	;
	v991 = v82 | int64(1024)
	goto L23
L163:
	;
	v540 = int32(_a_F_commandFlagsFromString_12)
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v543 != 0 {
		goto L180
	} else {
		goto L181
	}
L164:
	;
	if v534-v536 != 0 {
		goto L163
	} else {
		goto L176
	}
L165:
	;
	v534 = F_tolower(m, v530)
	mBase = m.M
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	v536 = F_tolower(m, v535)
	mBase = m.M
	goto L164
L166:
	;
	v504 = v88
	v505 = v499
	v506 = v502
	goto L169
L167:
	;
	v530 = int32(0)
	v531 = v499
	goto L165
L168:
	;
	v530 = v527 & int32(255)
	v531 = v526
	goto L165
L169:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	if v508 == int32(0) {
		v526 = v505
		v527 = v506
		goto L168
	} else {
		goto L171
	}
L170:
	;
	v526 = v520
	v527 = int32(0)
	goto L168
L171:
	;
	v512 = v506 & int32(255)
	if v512 == v508 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v519 = int32(1)
	v520 = v505 + v519
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504)+1)))
	if v521 != 0 {
		v504 = v504 + v519
		v505 = v520
		v506 = v521
		goto L169
	} else {
		goto L175
	}
L173:
	;
	v514 = F_tolower(m, v512)
	mBase = m.M
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v516 = F_tolower(m, v515)
	mBase = m.M
	if v514 == v516 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	v526 = v505
	v527 = v518
	goto L168
L175:
	;
	goto L170
L176:
	;
	v991 = v82 | int64(2048)
	goto L23
L177:
	;
	v581 = int32(_a_F_commandFlagsFromString_13)
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v584 != 0 {
		goto L194
	} else {
		goto L195
	}
L178:
	;
	if v575-v577 != 0 {
		goto L177
	} else {
		goto L190
	}
L179:
	;
	v575 = F_tolower(m, v571)
	mBase = m.M
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	v577 = F_tolower(m, v576)
	mBase = m.M
	goto L178
L180:
	;
	v545 = v88
	v546 = v540
	v547 = v543
	goto L183
L181:
	;
	v571 = int32(0)
	v572 = v540
	goto L179
L182:
	;
	v571 = v568 & int32(255)
	v572 = v567
	goto L179
L183:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if v549 == int32(0) {
		v567 = v546
		v568 = v547
		goto L182
	} else {
		goto L185
	}
L184:
	;
	v567 = v561
	v568 = int32(0)
	goto L182
L185:
	;
	v553 = v547 & int32(255)
	if v553 == v549 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v560 = int32(1)
	v561 = v546 + v560
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)))
	if v562 != 0 {
		v545 = v545 + v560
		v546 = v561
		v547 = v562
		goto L183
	} else {
		goto L189
	}
L187:
	;
	v555 = F_tolower(m, v553)
	mBase = m.M
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	v557 = F_tolower(m, v556)
	mBase = m.M
	if v555 == v557 {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	v567 = v546
	v568 = v559
	goto L182
L189:
	;
	goto L184
L190:
	;
	v991 = v82 | int64(4096)
	goto L23
L191:
	;
	v622 = int32(_a_F_commandFlagsFromString_14)
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v625 != 0 {
		goto L208
	} else {
		goto L209
	}
L192:
	;
	if v616-v618 != 0 {
		goto L191
	} else {
		goto L204
	}
L193:
	;
	v616 = F_tolower(m, v612)
	mBase = m.M
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v618 = F_tolower(m, v617)
	mBase = m.M
	goto L192
L194:
	;
	v586 = v88
	v587 = v581
	v588 = v584
	goto L197
L195:
	;
	v612 = int32(0)
	v613 = v581
	goto L193
L196:
	;
	v612 = v609 & int32(255)
	v613 = v608
	goto L193
L197:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v590 == int32(0) {
		v608 = v587
		v609 = v588
		goto L196
	} else {
		goto L199
	}
L198:
	;
	v608 = v602
	v609 = int32(0)
	goto L196
L199:
	;
	v594 = v588 & int32(255)
	if v594 == v590 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v601 = int32(1)
	v602 = v587 + v601
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+1)))
	if v603 != 0 {
		v586 = v586 + v601
		v587 = v602
		v588 = v603
		goto L197
	} else {
		goto L203
	}
L201:
	;
	v596 = F_tolower(m, v594)
	mBase = m.M
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	v598 = F_tolower(m, v597)
	mBase = m.M
	if v596 == v598 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	v608 = v587
	v609 = v600
	goto L196
L203:
	;
	goto L198
L204:
	;
	v991 = v82 | int64(4096)
	goto L23
L205:
	;
	v663 = int32(_a_F_commandFlagsFromString_15)
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v666 != 0 {
		goto L222
	} else {
		goto L223
	}
L206:
	;
	if v657-v659 != 0 {
		goto L205
	} else {
		goto L218
	}
L207:
	;
	v657 = F_tolower(m, v653)
	mBase = m.M
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	v659 = F_tolower(m, v658)
	mBase = m.M
	goto L206
L208:
	;
	v627 = v88
	v628 = v622
	v629 = v625
	goto L211
L209:
	;
	v653 = int32(0)
	v654 = v622
	goto L207
L210:
	;
	v653 = v650 & int32(255)
	v654 = v649
	goto L207
L211:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v631 == int32(0) {
		v649 = v628
		v650 = v629
		goto L210
	} else {
		goto L213
	}
L212:
	;
	v649 = v643
	v650 = int32(0)
	goto L210
L213:
	;
	v635 = v629 & int32(255)
	if v635 == v631 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v642 = int32(1)
	v643 = v628 + v642
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+1)))
	if v644 != 0 {
		v627 = v627 + v642
		v628 = v643
		v629 = v644
		goto L211
	} else {
		goto L217
	}
L215:
	;
	v637 = F_tolower(m, v635)
	mBase = m.M
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	v639 = F_tolower(m, v638)
	mBase = m.M
	if v637 == v639 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	v649 = v628
	v650 = v641
	goto L210
L217:
	;
	goto L212
L218:
	;
	v991 = v82 | int64(16384)
	goto L23
L219:
	;
	v704 = int32(_a_F_commandFlagsFromString_16)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v707 != 0 {
		goto L236
	} else {
		goto L237
	}
L220:
	;
	if v698-v700 != 0 {
		goto L219
	} else {
		goto L232
	}
L221:
	;
	v698 = F_tolower(m, v694)
	mBase = m.M
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	v700 = F_tolower(m, v699)
	mBase = m.M
	goto L220
L222:
	;
	v668 = v88
	v669 = v663
	v670 = v666
	goto L225
L223:
	;
	v694 = int32(0)
	v695 = v663
	goto L221
L224:
	;
	v694 = v691 & int32(255)
	v695 = v690
	goto L221
L225:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669))))
	if v672 == int32(0) {
		v690 = v669
		v691 = v670
		goto L224
	} else {
		goto L227
	}
L226:
	;
	v690 = v684
	v691 = int32(0)
	goto L224
L227:
	;
	v676 = v670 & int32(255)
	if v676 == v672 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v683 = int32(1)
	v684 = v669 + v683
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+1)))
	if v685 != 0 {
		v668 = v668 + v683
		v669 = v684
		v670 = v685
		goto L225
	} else {
		goto L231
	}
L229:
	;
	v678 = F_tolower(m, v676)
	mBase = m.M
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669))))
	v680 = F_tolower(m, v679)
	mBase = m.M
	if v678 == v680 {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	v690 = v669
	v691 = v682
	goto L224
L231:
	;
	goto L226
L232:
	;
	v991 = v82 | int64(32768)
	goto L23
L233:
	;
	v745 = int32(_a_F_commandFlagsFromString_17)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v748 != 0 {
		goto L250
	} else {
		goto L251
	}
L234:
	;
	if v739-v741 != 0 {
		goto L233
	} else {
		goto L246
	}
L235:
	;
	v739 = F_tolower(m, v735)
	mBase = m.M
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736))))
	v741 = F_tolower(m, v740)
	mBase = m.M
	goto L234
L236:
	;
	v709 = v88
	v710 = v704
	v711 = v707
	goto L239
L237:
	;
	v735 = int32(0)
	v736 = v704
	goto L235
L238:
	;
	v735 = v732 & int32(255)
	v736 = v731
	goto L235
L239:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
	if v713 == int32(0) {
		v731 = v710
		v732 = v711
		goto L238
	} else {
		goto L241
	}
L240:
	;
	v731 = v725
	v732 = int32(0)
	goto L238
L241:
	;
	v717 = v711 & int32(255)
	if v717 == v713 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v724 = int32(1)
	v725 = v710 + v724
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709)+1)))
	if v726 != 0 {
		v709 = v709 + v724
		v710 = v725
		v711 = v726
		goto L239
	} else {
		goto L245
	}
L243:
	;
	v719 = F_tolower(m, v717)
	mBase = m.M
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710))))
	v721 = F_tolower(m, v720)
	mBase = m.M
	if v719 == v721 {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	v731 = v710
	v732 = v723
	goto L238
L245:
	;
	goto L240
L246:
	;
	v991 = v82 | int64(65536)
	goto L23
L247:
	;
	v786 = int32(_a_F_commandFlagsFromString_18)
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v789 != 0 {
		goto L264
	} else {
		goto L265
	}
L248:
	;
	if v780-v782 != 0 {
		goto L247
	} else {
		goto L260
	}
L249:
	;
	v780 = F_tolower(m, v776)
	mBase = m.M
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	v782 = F_tolower(m, v781)
	mBase = m.M
	goto L248
L250:
	;
	v750 = v88
	v751 = v745
	v752 = v748
	goto L253
L251:
	;
	v776 = int32(0)
	v777 = v745
	goto L249
L252:
	;
	v776 = v773 & int32(255)
	v777 = v772
	goto L249
L253:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	if v754 == int32(0) {
		v772 = v751
		v773 = v752
		goto L252
	} else {
		goto L255
	}
L254:
	;
	v772 = v766
	v773 = int32(0)
	goto L252
L255:
	;
	v758 = v752 & int32(255)
	if v758 == v754 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v765 = int32(1)
	v766 = v751 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	if v767 != 0 {
		v750 = v750 + v765
		v751 = v766
		v752 = v767
		goto L253
	} else {
		goto L259
	}
L257:
	;
	v760 = F_tolower(m, v758)
	mBase = m.M
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	v762 = F_tolower(m, v761)
	mBase = m.M
	if v760 == v762 {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v772 = v751
	v773 = v764
	goto L252
L259:
	;
	goto L254
L260:
	;
	v991 = v82 | int64(2097152)
	goto L23
L261:
	;
	v827 = int32(_a_F_commandFlagsFromString_19)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v830 != 0 {
		goto L278
	} else {
		goto L279
	}
L262:
	;
	if v821-v823 != 0 {
		goto L261
	} else {
		goto L274
	}
L263:
	;
	v821 = F_tolower(m, v817)
	mBase = m.M
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818))))
	v823 = F_tolower(m, v822)
	mBase = m.M
	goto L262
L264:
	;
	v791 = v88
	v792 = v786
	v793 = v789
	goto L267
L265:
	;
	v817 = int32(0)
	v818 = v786
	goto L263
L266:
	;
	v817 = v814 & int32(255)
	v818 = v813
	goto L263
L267:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792))))
	if v795 == int32(0) {
		v813 = v792
		v814 = v793
		goto L266
	} else {
		goto L269
	}
L268:
	;
	v813 = v807
	v814 = int32(0)
	goto L266
L269:
	;
	v799 = v793 & int32(255)
	if v799 == v795 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v806 = int32(1)
	v807 = v792 + v806
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+1)))
	if v808 != 0 {
		v791 = v791 + v806
		v792 = v807
		v793 = v808
		goto L267
	} else {
		goto L273
	}
L271:
	;
	v801 = F_tolower(m, v799)
	mBase = m.M
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792))))
	v803 = F_tolower(m, v802)
	mBase = m.M
	if v801 == v803 {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791))))
	v813 = v792
	v814 = v805
	goto L266
L273:
	;
	goto L268
L274:
	;
	v991 = v82 | int64(134217728)
	goto L23
L275:
	;
	v868 = int32(_a_F_commandFlagsFromString_20)
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v871 != 0 {
		goto L292
	} else {
		goto L293
	}
L276:
	;
	if v862-v864 != 0 {
		goto L275
	} else {
		goto L288
	}
L277:
	;
	v862 = F_tolower(m, v858)
	mBase = m.M
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	v864 = F_tolower(m, v863)
	mBase = m.M
	goto L276
L278:
	;
	v832 = v88
	v833 = v827
	v834 = v830
	goto L281
L279:
	;
	v858 = int32(0)
	v859 = v827
	goto L277
L280:
	;
	v858 = v855 & int32(255)
	v859 = v854
	goto L277
L281:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	if v836 == int32(0) {
		v854 = v833
		v855 = v834
		goto L280
	} else {
		goto L283
	}
L282:
	;
	v854 = v848
	v855 = int32(0)
	goto L280
L283:
	;
	v840 = v834 & int32(255)
	if v840 == v836 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v847 = int32(1)
	v848 = v833 + v847
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+1)))
	if v849 != 0 {
		v832 = v832 + v847
		v833 = v848
		v834 = v849
		goto L281
	} else {
		goto L287
	}
L285:
	;
	v842 = F_tolower(m, v840)
	mBase = m.M
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	v844 = F_tolower(m, v843)
	mBase = m.M
	if v842 == v844 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832))))
	v854 = v833
	v855 = v846
	goto L280
L287:
	;
	goto L282
L288:
	;
	v991 = v82 | int64(4194304)
	goto L23
L289:
	;
	v909 = int32(_a_F_commandFlagsFromString_21)
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v912 != 0 {
		goto L306
	} else {
		goto L307
	}
L290:
	;
	if v903-v905 != 0 {
		goto L289
	} else {
		goto L302
	}
L291:
	;
	v903 = F_tolower(m, v899)
	mBase = m.M
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
	v905 = F_tolower(m, v904)
	mBase = m.M
	goto L290
L292:
	;
	v873 = v88
	v874 = v868
	v875 = v871
	goto L295
L293:
	;
	v899 = int32(0)
	v900 = v868
	goto L291
L294:
	;
	v899 = v896 & int32(255)
	v900 = v895
	goto L291
L295:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874))))
	if v877 == int32(0) {
		v895 = v874
		v896 = v875
		goto L294
	} else {
		goto L297
	}
L296:
	;
	v895 = v889
	v896 = int32(0)
	goto L294
L297:
	;
	v881 = v875 & int32(255)
	if v881 == v877 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v888 = int32(1)
	v889 = v874 + v888
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+1)))
	if v890 != 0 {
		v873 = v873 + v888
		v874 = v889
		v875 = v890
		goto L295
	} else {
		goto L301
	}
L299:
	;
	v883 = F_tolower(m, v881)
	mBase = m.M
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874))))
	v885 = F_tolower(m, v884)
	mBase = m.M
	if v883 == v885 {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	v895 = v874
	v896 = v887
	goto L294
L301:
	;
	goto L296
L302:
	;
	v991 = v82 | int64(524288)
	goto L23
L303:
	;
	v950 = int32(_a_F_commandFlagsFromString_22)
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v953 != 0 {
		goto L319
	} else {
		goto L320
	}
L304:
	;
	if v944-v946 != 0 {
		goto L303
	} else {
		goto L316
	}
L305:
	;
	v944 = F_tolower(m, v940)
	mBase = m.M
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941))))
	v946 = F_tolower(m, v945)
	mBase = m.M
	goto L304
L306:
	;
	v914 = v88
	v915 = v909
	v916 = v912
	goto L309
L307:
	;
	v940 = int32(0)
	v941 = v909
	goto L305
L308:
	;
	v940 = v937 & int32(255)
	v941 = v936
	goto L305
L309:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915))))
	if v918 == int32(0) {
		v936 = v915
		v937 = v916
		goto L308
	} else {
		goto L311
	}
L310:
	;
	v936 = v930
	v937 = int32(0)
	goto L308
L311:
	;
	v922 = v916 & int32(255)
	if v922 == v918 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v929 = int32(1)
	v930 = v915 + v929
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914)+1)))
	if v931 != 0 {
		v914 = v914 + v929
		v915 = v930
		v916 = v931
		goto L309
	} else {
		goto L315
	}
L313:
	;
	v924 = F_tolower(m, v922)
	mBase = m.M
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915))))
	v926 = F_tolower(m, v925)
	mBase = m.M
	if v924 == v926 {
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	v936 = v915
	v937 = v928
	goto L308
L315:
	;
	goto L310
L316:
	;
	v991 = v82 | int64(67108864)
	goto L23
L317:
	;
	if v985-v987 != 0 {
		v995 = v79
		v998 = v82
		goto L19
	} else {
		goto L329
	}
L318:
	;
	v985 = F_tolower(m, v981)
	mBase = m.M
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	v987 = F_tolower(m, v986)
	mBase = m.M
	goto L317
L319:
	;
	v955 = v88
	v956 = v950
	v957 = v953
	goto L322
L320:
	;
	v981 = int32(0)
	v982 = v950
	goto L318
L321:
	;
	v981 = v978 & int32(255)
	v982 = v977
	goto L318
L322:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	if v959 == int32(0) {
		v977 = v956
		v978 = v957
		goto L321
	} else {
		goto L324
	}
L323:
	;
	v977 = v971
	v978 = int32(0)
	goto L321
L324:
	;
	v963 = v957 & int32(255)
	if v963 == v959 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v970 = int32(1)
	v971 = v956 + v970
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+1)))
	if v972 != 0 {
		v955 = v955 + v970
		v956 = v971
		v957 = v972
		goto L322
	} else {
		goto L328
	}
L326:
	;
	v965 = F_tolower(m, v963)
	mBase = m.M
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	v967 = F_tolower(m, v966)
	mBase = m.M
	if v965 == v967 {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955))))
	v977 = v956
	v978 = v969
	goto L321
L328:
	;
	goto L323
L329:
	;
	v991 = v82 | int64(536870912)
	goto L23
L330:
	;
	goto L22
L331:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	m.G0 = v9 + int32(16)
	if v995 == v1003 {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1009 = v998
	goto L334
L333:
	;
	v1009 = int64(-1)
	goto L334
L334:
	;
	return v1009
}
func F_commandInfoCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(64)
	return
L2:
	;
	F_addReplyArrayLen(m, l0, v9+int32(-2))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L20
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_commandInfoCommand[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	goto L4
L4:
	;
	F_addReplyArrayLen(m, l0, v14+v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v20 = v7 + int32(16)
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_commandInfoCommand[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(-1)
	if v22 == v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v45 = F_hashtableNext(m, v7+int32(16), v7+int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L12
	}
L8:
	;
	goto L7
L9:
	;
	goto L8
L11:
	;
	F_hashtableCleanupIterator(m, v7+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L19
	}
L12:
	;
	if v45 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_addReplyCommandInfo(m, l0, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v60 = F_hashtableNext(m, v7+int32(16), v7+int32(12))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v60 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L1
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v74 < int32(3) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v80 = int32(2)
	goto L22
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v80<<(uint(int32(2))%32))))
	v87 = F_objectGetVal(m, v86)
	mBase = m.M
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_commandInfoCommand[0]))
	v90 = F_lookupCommandBySdsLogic(m, v89, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L24
	}
L23:
	;
	goto L1
L24:
	;
	F_addReplyCommandInfo(m, l0, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v95 = v80 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v95 < v96 {
		v80 = v95
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
}
func F_commandListCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int64
	_ = v242
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(2) < v17 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(64)
	return
L2:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	F_setDeferredArrayLen(m, l0, v267, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L73
	}
L3:
	;
	v33 = v17
	v34 = int32(2)
	goto L8
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(0)
	v22 = F_addReplyDeferredLen(m, l0)
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_commandListCommand[0]))
	F_commandListWithoutFilter(m, l0, v25, v9+int32(36))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v267 = v22
	goto L2
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = v34 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v39)))
	v42 = F_objectGetVal(m, v41)
	mBase = m.M
	v43 = int32(_a_F_commandListCommand_0)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(0)
	v232 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L71
	}
L10:
	;
	if v34-v33 != int32(-3) {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v78 = F_tolower(m, v74)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v80 = F_tolower(m, v79)
	mBase = m.M
	goto L10
L12:
	;
	v48 = v42
	v49 = v43
	v50 = v46
	goto L15
L13:
	;
	v74 = int32(0)
	v75 = v43
	goto L11
L14:
	;
	v74 = v71 & int32(255)
	v75 = v70
	goto L11
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v52 == int32(0) {
		v70 = v49
		v71 = v50
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v70 = v64
	v71 = int32(0)
	goto L14
L17:
	;
	v56 = v50 & int32(255)
	if v56 == v52 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = int32(1)
	v64 = v49 + v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v65 != 0 {
		v48 = v48 + v63
		v49 = v64
		v50 = v65
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v58 = F_tolower(m, v56)
	mBase = m.M
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v60 = F_tolower(m, v59)
	mBase = m.M
	if v58 == v60 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v70 = v49
	v71 = v62
	goto L14
L21:
	;
	goto L16
L22:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v220+v39)+8))
	v223 = F_objectGetVal(m, v222)
	mBase = m.M
	v225 = v34 + int32(3)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v225 < v226 {
		v33 = v226
		v34 = v225
		goto L8
	} else {
		goto L70
	}
L23:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_commandListCommand[1]))
	F_addReplyErrorObject(m, l0, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L5
	} else {
		goto L69
	}
L24:
	;
	if v78-v80 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v39+int32(4))))
	v90 = F_objectGetVal(m, v89)
	mBase = m.M
	v91 = int32(_a_F_commandListCommand_1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v131 = int32(_a_F_commandListCommand_2)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v134 != 0 {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	if v126-v128 != 0 {
		goto L26
	} else {
		goto L39
	}
L28:
	;
	v126 = F_tolower(m, v122)
	mBase = m.M
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v128 = F_tolower(m, v127)
	mBase = m.M
	goto L27
L29:
	;
	v96 = v90
	v97 = v91
	v98 = v94
	goto L32
L30:
	;
	v122 = int32(0)
	v123 = v91
	goto L28
L31:
	;
	v122 = v119 & int32(255)
	v123 = v118
	goto L28
L32:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v100 == int32(0) {
		v118 = v97
		v119 = v98
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v118 = v112
	v119 = int32(0)
	goto L31
L34:
	;
	v104 = v98 & int32(255)
	if v104 == v100 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v111 = int32(1)
	v112 = v97 + v111
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v113 != 0 {
		v96 = v96 + v111
		v97 = v112
		v98 = v113
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v106 = F_tolower(m, v104)
	mBase = m.M
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v108 = F_tolower(m, v107)
	mBase = m.M
	if v106 == v108 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v118 = v97
	v119 = v110
	goto L31
L38:
	;
	goto L33
L39:
	;
	v219 = int32(0)
	goto L22
L40:
	;
	v171 = int32(_a_F_commandListCommand_3)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v174 != 0 {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	if v166-v168 != 0 {
		goto L40
	} else {
		goto L53
	}
L42:
	;
	v166 = F_tolower(m, v162)
	mBase = m.M
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v168 = F_tolower(m, v167)
	mBase = m.M
	goto L41
L43:
	;
	v136 = v90
	v137 = v131
	v138 = v134
	goto L46
L44:
	;
	v162 = int32(0)
	v163 = v131
	goto L42
L45:
	;
	v162 = v159 & int32(255)
	v163 = v158
	goto L42
L46:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v140 == int32(0) {
		v158 = v137
		v159 = v138
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v158 = v152
	v159 = int32(0)
	goto L45
L48:
	;
	v144 = v138 & int32(255)
	if v144 == v140 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v151 = int32(1)
	v152 = v137 + v151
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v153 != 0 {
		v136 = v136 + v151
		v137 = v152
		v138 = v153
		goto L46
	} else {
		goto L52
	}
L50:
	;
	v146 = F_tolower(m, v144)
	mBase = m.M
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v148 = F_tolower(m, v147)
	mBase = m.M
	if v146 == v148 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v158 = v137
	v159 = v150
	goto L45
L52:
	;
	goto L47
L53:
	;
	v219 = int32(1)
	goto L22
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_commandListCommand[1]))
	F_addReplyErrorObject(m, l0, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L68
	}
L55:
	;
	if v206-v208 != 0 {
		goto L54
	} else {
		goto L67
	}
L56:
	;
	v206 = F_tolower(m, v202)
	mBase = m.M
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v208 = F_tolower(m, v207)
	mBase = m.M
	goto L55
L57:
	;
	v176 = v90
	v177 = v171
	v178 = v174
	goto L60
L58:
	;
	v202 = int32(0)
	v203 = v171
	goto L56
L59:
	;
	v202 = v199 & int32(255)
	v203 = v198
	goto L56
L60:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v180 == int32(0) {
		v198 = v177
		v199 = v178
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v198 = v192
	v199 = int32(0)
	goto L59
L62:
	;
	v184 = v178 & int32(255)
	if v184 == v180 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v191 = int32(1)
	v192 = v177 + v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	if v193 != 0 {
		v176 = v176 + v191
		v177 = v192
		v178 = v193
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v186 = F_tolower(m, v184)
	mBase = m.M
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v188 = F_tolower(m, v187)
	mBase = m.M
	if v186 == v188 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v198 = v177
	v199 = v190
	goto L59
L66:
	;
	goto L61
L67:
	;
	v219 = int32(2)
	goto L22
L68:
	;
	goto L1
L69:
	;
	goto L1
L70:
	;
	goto L9
L71:
	;
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v242
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(56))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v254
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_commandListCommand[0]))
	F_commandListWithFilter(m, l0, v257, v9+int32(8), v9+int32(36))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v267 = v232
	goto L2
L73:
	;
	goto L1
}
func F_commandListWithoutFilter(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v13 = v10 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(-1)
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v10+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L23
	}
L6:
	;
	return
L7:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+140))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(-1)))))
	switch v52 & int32(7) {
	case 0:
		goto L16
	case 1:
		goto L15
	case 2:
		goto L14
	case 3:
		goto L13
	case 4:
		goto L12
	default:
		v69 = int32(0)
		goto L11
	}
L10:
	;
	goto L5
L11:
	;
	F_addReplyBulkCBuffer(m, l0, v49, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L17
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(-17))))
	v69 = v68
	goto L11
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(-9))))
	v69 = v65
	goto L11
L14:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49+int32(-5)))))
	v69 = v62
	goto L11
L15:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(-3)))))
	v69 = v59
	goto L11
L16:
	;
	v69 = int32(base.Ui32(v52) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v72 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+200))
	if v76 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = F_hashtableNext(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	F_commandListWithoutFilter(m, l0, v76, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if v85 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	goto L10
L23:
	;
	m.G0 = v10 + int32(64)
	return
}
func F_dumpCommand(m *base.Module, l0 int32) {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = F_lookupKeyRead(m, v8, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
			F_createDumpPayload(m, v6, v11, v16, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
				F_addReplyBulkSds(m, l0, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					m.G0 = v6 + int32(80)
					return
				}
			}
		} else {
			F_addReplyNull(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				m.G0 = v6 + int32(80)
				return
			}
		}
	}
}
func F_evalGetCommandFlags(m *base.Module, l0 int32, l1 int64) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
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
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v210 int64
	_ = v210
	var v217 int64
	_ = v217
	var v221 int64
	_ = v221
	v7 = m.G0
	v9 = v7 - int32(64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v17 = base.B2i32(v12 == int32(292)) | base.B2i32(v12 == int32(293))
	if v17 == int32(0) {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
		v48 = F_objectGetVal(m, v47)
		mBase = m.M
		v50 = v9 + int32(16)
		v54 = m.G0
		v56 = v54 - int32(112)
		m.G0 = v56
		if v17 == int32(0) {
			v101 = int32(0)
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
			switch v105 & int32(7) {
			case 0:
				v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
			case 1:
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
				v122 = v112
			case 2:
				v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
				v122 = v115
			case 3:
				v118 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
				v122 = v118
			case 4:
				v121 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
				v122 = v121
			default:
				v122 = v101
			}
			v124 = v56 + int32(20)
			F_SHA1Init(m, v124)
			mBase = m.M
			F_SHA1Update(m, v124, v48, v122)
			mBase = m.M
			F_SHA1Final(m, v56, v124)
			mBase = m.M
			v132 = v101
			for {
				v138 = int32(1)
				v140 = v50 + v132<<(uint(v138)%32)
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v132))))
				v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144&int32(15))+uint32(_c_F_evalGetCommandFlags[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v140+v138))) = uint8(v149)
				v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v144)>>(uint(int32(4))%32)))+uint32(_c_F_evalGetCommandFlags[0]))))
				*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
				v158 = v132 + v138
				if v158 != int32(20) {
					v132 = v158
					continue
				} else {
					break
				}
				break
			}
			v161 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v161)
		} else {
			v61 = int32(0)
			for {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v61))))
				if base.Ui32((v69+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
					v78 = v69 + int32(32)
				} else {
					v78 = v69
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v50+v61))) = uint8(v78)
				v81 = v61 | int32(1)
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v81))))
				if base.Ui32((v84+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
					v93 = v84 + int32(32)
				} else {
					v93 = v84
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v50+v81))) = uint8(v93)
				v96 = v61 + int32(2)
				if v96 != int32(40) {
					v61 = v96
					continue
				} else {
					break
				}
				break
			}
			v99 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v99)
		}
		m.G0 = v56 + int32(112)
		v173 = *(*int32)(unsafe.Add(mBase, _c_F_evalGetCommandFlags[1]))
		v176 = F_dictFind(m, v173, v9+int32(16))
		mBase = m.M
		v179 = m.ExcPending
		if v179 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v176
			if v176 != 0 {
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
				v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v195
				v197 = v195
				if v197&int64(16) != int64(0) {
					v221 = l1
				} else {
					v210 = l1 & int64(-66566)
					if v197&int64(3) == int64(0) {
						v217 = v210 | int64(4)
					} else {
						v217 = v210
					}
					v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
				}
				m.G0 = v9 + int32(64)
				return v221
			} else {
				if v17 != 0 {
					v221 = l1
					m.G0 = v9 + int32(64)
					return v221
				} else {
					v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
					v183 = F_objectGetVal(m, v182)
					mBase = m.M
					v184 = int32(0)
					v189 = F_evalExtractShebangFlags(m, v183, v184, v9+int32(8), v184, v184)
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int64(0)
					} else {
						if v189 == int32(-1) {
							v221 = l1
						} else {
							v193 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
							v197 = v193
							if v197&int64(16) != int64(0) {
								v221 = l1
							} else {
								v210 = l1 & int64(-66566)
								if v197&int64(3) == int64(0) {
									v217 = v210 | int64(4)
								} else {
									v217 = v210
								}
								v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
							}
						}
						m.G0 = v9 + int32(64)
						return v221
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		v22 = F_objectGetVal(m, v21)
		mBase = m.M
		v23 = int32(-1)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23))))
		switch v25&int32(7) + v23 {
		case 0:
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
			v42 = v32
			if v42 != int32(40) {
				v221 = l1
				m.G0 = v9 + int32(64)
				return v221
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
				v48 = F_objectGetVal(m, v47)
				mBase = m.M
				v50 = v9 + int32(16)
				v54 = m.G0
				v56 = v54 - int32(112)
				m.G0 = v56
				if v17 == int32(0) {
					v101 = int32(0)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
					switch v105 & int32(7) {
					case 0:
						v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
					case 1:
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
						v122 = v112
					case 2:
						v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
						v122 = v115
					case 3:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
						v122 = v118
					case 4:
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
						v122 = v121
					default:
						v122 = v101
					}
					v124 = v56 + int32(20)
					F_SHA1Init(m, v124)
					mBase = m.M
					F_SHA1Update(m, v124, v48, v122)
					mBase = m.M
					F_SHA1Final(m, v56, v124)
					mBase = m.M
					v132 = v101
					for {
						v138 = int32(1)
						v140 = v50 + v132<<(uint(v138)%32)
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v132))))
						v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144&int32(15))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140+v138))) = uint8(v149)
						v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v144)>>(uint(int32(4))%32)))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
						v158 = v132 + v138
						if v158 != int32(20) {
							v132 = v158
							continue
						} else {
							break
						}
						break
					}
					v161 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v161)
				} else {
					v61 = int32(0)
					for {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v61))))
						if base.Ui32((v69+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v78 = v69 + int32(32)
						} else {
							v78 = v69
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v61))) = uint8(v78)
						v81 = v61 | int32(1)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v81))))
						if base.Ui32((v84+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v93 = v84 + int32(32)
						} else {
							v93 = v84
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v81))) = uint8(v93)
						v96 = v61 + int32(2)
						if v96 != int32(40) {
							v61 = v96
							continue
						} else {
							break
						}
						break
					}
					v99 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v99)
				}
				m.G0 = v56 + int32(112)
				v173 = *(*int32)(unsafe.Add(mBase, _c_F_evalGetCommandFlags[1]))
				v176 = F_dictFind(m, v173, v9+int32(16))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v176
					if v176 != 0 {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
						v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v195
						v197 = v195
						if v197&int64(16) != int64(0) {
							v221 = l1
						} else {
							v210 = l1 & int64(-66566)
							if v197&int64(3) == int64(0) {
								v217 = v210 | int64(4)
							} else {
								v217 = v210
							}
							v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
						}
						m.G0 = v9 + int32(64)
						return v221
					} else {
						if v17 != 0 {
							v221 = l1
							m.G0 = v9 + int32(64)
							return v221
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
							v183 = F_objectGetVal(m, v182)
							mBase = m.M
							v184 = int32(0)
							v189 = F_evalExtractShebangFlags(m, v183, v184, v9+int32(8), v184, v184)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int64(0)
							} else {
								if v189 == int32(-1) {
									v221 = l1
								} else {
									v193 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
									v197 = v193
									if v197&int64(16) != int64(0) {
										v221 = l1
									} else {
										v210 = l1 & int64(-66566)
										if v197&int64(3) == int64(0) {
											v217 = v210 | int64(4)
										} else {
											v217 = v210
										}
										v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
									}
								}
								m.G0 = v9 + int32(64)
								return v221
							}
						}
					}
				}
			}
		case 1:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))))
			v42 = v35
			if v42 != int32(40) {
				v221 = l1
				m.G0 = v9 + int32(64)
				return v221
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
				v48 = F_objectGetVal(m, v47)
				mBase = m.M
				v50 = v9 + int32(16)
				v54 = m.G0
				v56 = v54 - int32(112)
				m.G0 = v56
				if v17 == int32(0) {
					v101 = int32(0)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
					switch v105 & int32(7) {
					case 0:
						v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
					case 1:
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
						v122 = v112
					case 2:
						v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
						v122 = v115
					case 3:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
						v122 = v118
					case 4:
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
						v122 = v121
					default:
						v122 = v101
					}
					v124 = v56 + int32(20)
					F_SHA1Init(m, v124)
					mBase = m.M
					F_SHA1Update(m, v124, v48, v122)
					mBase = m.M
					F_SHA1Final(m, v56, v124)
					mBase = m.M
					v132 = v101
					for {
						v138 = int32(1)
						v140 = v50 + v132<<(uint(v138)%32)
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v132))))
						v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144&int32(15))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140+v138))) = uint8(v149)
						v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v144)>>(uint(int32(4))%32)))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
						v158 = v132 + v138
						if v158 != int32(20) {
							v132 = v158
							continue
						} else {
							break
						}
						break
					}
					v161 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v161)
				} else {
					v61 = int32(0)
					for {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v61))))
						if base.Ui32((v69+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v78 = v69 + int32(32)
						} else {
							v78 = v69
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v61))) = uint8(v78)
						v81 = v61 | int32(1)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v81))))
						if base.Ui32((v84+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v93 = v84 + int32(32)
						} else {
							v93 = v84
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v81))) = uint8(v93)
						v96 = v61 + int32(2)
						if v96 != int32(40) {
							v61 = v96
							continue
						} else {
							break
						}
						break
					}
					v99 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v99)
				}
				m.G0 = v56 + int32(112)
				v173 = *(*int32)(unsafe.Add(mBase, _c_F_evalGetCommandFlags[1]))
				v176 = F_dictFind(m, v173, v9+int32(16))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v176
					if v176 != 0 {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
						v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v195
						v197 = v195
						if v197&int64(16) != int64(0) {
							v221 = l1
						} else {
							v210 = l1 & int64(-66566)
							if v197&int64(3) == int64(0) {
								v217 = v210 | int64(4)
							} else {
								v217 = v210
							}
							v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
						}
						m.G0 = v9 + int32(64)
						return v221
					} else {
						if v17 != 0 {
							v221 = l1
							m.G0 = v9 + int32(64)
							return v221
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
							v183 = F_objectGetVal(m, v182)
							mBase = m.M
							v184 = int32(0)
							v189 = F_evalExtractShebangFlags(m, v183, v184, v9+int32(8), v184, v184)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int64(0)
							} else {
								if v189 == int32(-1) {
									v221 = l1
								} else {
									v193 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
									v197 = v193
									if v197&int64(16) != int64(0) {
										v221 = l1
									} else {
										v210 = l1 & int64(-66566)
										if v197&int64(3) == int64(0) {
											v217 = v210 | int64(4)
										} else {
											v217 = v210
										}
										v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
									}
								}
								m.G0 = v9 + int32(64)
								return v221
							}
						}
					}
				}
			}
		case 2:
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
			v42 = v38
			if v42 != int32(40) {
				v221 = l1
				m.G0 = v9 + int32(64)
				return v221
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
				v48 = F_objectGetVal(m, v47)
				mBase = m.M
				v50 = v9 + int32(16)
				v54 = m.G0
				v56 = v54 - int32(112)
				m.G0 = v56
				if v17 == int32(0) {
					v101 = int32(0)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
					switch v105 & int32(7) {
					case 0:
						v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
					case 1:
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
						v122 = v112
					case 2:
						v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
						v122 = v115
					case 3:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
						v122 = v118
					case 4:
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
						v122 = v121
					default:
						v122 = v101
					}
					v124 = v56 + int32(20)
					F_SHA1Init(m, v124)
					mBase = m.M
					F_SHA1Update(m, v124, v48, v122)
					mBase = m.M
					F_SHA1Final(m, v56, v124)
					mBase = m.M
					v132 = v101
					for {
						v138 = int32(1)
						v140 = v50 + v132<<(uint(v138)%32)
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v132))))
						v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144&int32(15))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140+v138))) = uint8(v149)
						v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v144)>>(uint(int32(4))%32)))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
						v158 = v132 + v138
						if v158 != int32(20) {
							v132 = v158
							continue
						} else {
							break
						}
						break
					}
					v161 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v161)
				} else {
					v61 = int32(0)
					for {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v61))))
						if base.Ui32((v69+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v78 = v69 + int32(32)
						} else {
							v78 = v69
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v61))) = uint8(v78)
						v81 = v61 | int32(1)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v81))))
						if base.Ui32((v84+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v93 = v84 + int32(32)
						} else {
							v93 = v84
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v81))) = uint8(v93)
						v96 = v61 + int32(2)
						if v96 != int32(40) {
							v61 = v96
							continue
						} else {
							break
						}
						break
					}
					v99 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v99)
				}
				m.G0 = v56 + int32(112)
				v173 = *(*int32)(unsafe.Add(mBase, _c_F_evalGetCommandFlags[1]))
				v176 = F_dictFind(m, v173, v9+int32(16))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v176
					if v176 != 0 {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
						v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v195
						v197 = v195
						if v197&int64(16) != int64(0) {
							v221 = l1
						} else {
							v210 = l1 & int64(-66566)
							if v197&int64(3) == int64(0) {
								v217 = v210 | int64(4)
							} else {
								v217 = v210
							}
							v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
						}
						m.G0 = v9 + int32(64)
						return v221
					} else {
						if v17 != 0 {
							v221 = l1
							m.G0 = v9 + int32(64)
							return v221
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
							v183 = F_objectGetVal(m, v182)
							mBase = m.M
							v184 = int32(0)
							v189 = F_evalExtractShebangFlags(m, v183, v184, v9+int32(8), v184, v184)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int64(0)
							} else {
								if v189 == int32(-1) {
									v221 = l1
								} else {
									v193 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
									v197 = v193
									if v197&int64(16) != int64(0) {
										v221 = l1
									} else {
										v210 = l1 & int64(-66566)
										if v197&int64(3) == int64(0) {
											v217 = v210 | int64(4)
										} else {
											v217 = v210
										}
										v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
									}
								}
								m.G0 = v9 + int32(64)
								return v221
							}
						}
					}
				}
			}
		case 3:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-17))))
			v42 = v41
			if v42 != int32(40) {
				v221 = l1
				m.G0 = v9 + int32(64)
				return v221
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
				v48 = F_objectGetVal(m, v47)
				mBase = m.M
				v50 = v9 + int32(16)
				v54 = m.G0
				v56 = v54 - int32(112)
				m.G0 = v56
				if v17 == int32(0) {
					v101 = int32(0)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
					switch v105 & int32(7) {
					case 0:
						v122 = int32(base.Ui32(v105) >> (uint(int32(3)) % 32))
					case 1:
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
						v122 = v112
					case 2:
						v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
						v122 = v115
					case 3:
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
						v122 = v118
					case 4:
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
						v122 = v121
					default:
						v122 = v101
					}
					v124 = v56 + int32(20)
					F_SHA1Init(m, v124)
					mBase = m.M
					F_SHA1Update(m, v124, v48, v122)
					mBase = m.M
					F_SHA1Final(m, v56, v124)
					mBase = m.M
					v132 = v101
					for {
						v138 = int32(1)
						v140 = v50 + v132<<(uint(v138)%32)
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v132))))
						v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144&int32(15))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140+v138))) = uint8(v149)
						v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v144)>>(uint(int32(4))%32)))+uint32(_c_F_evalGetCommandFlags[0]))))
						*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v155)
						v158 = v132 + v138
						if v158 != int32(20) {
							v132 = v158
							continue
						} else {
							break
						}
						break
					}
					v161 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v161)
				} else {
					v61 = int32(0)
					for {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v61))))
						if base.Ui32((v69+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v78 = v69 + int32(32)
						} else {
							v78 = v69
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v61))) = uint8(v78)
						v81 = v61 | int32(1)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v81))))
						if base.Ui32((v84+int32(-65))&int32(255)) < base.Ui32(int32(26)) {
							v93 = v84 + int32(32)
						} else {
							v93 = v84
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v50+v81))) = uint8(v93)
						v96 = v61 + int32(2)
						if v96 != int32(40) {
							v61 = v96
							continue
						} else {
							break
						}
						break
					}
					v99 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v50)+40)) = uint8(v99)
				}
				m.G0 = v56 + int32(112)
				v173 = *(*int32)(unsafe.Add(mBase, _c_F_evalGetCommandFlags[1]))
				v176 = F_dictFind(m, v173, v9+int32(16))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v176
					if v176 != 0 {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
						v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v195
						v197 = v195
						if v197&int64(16) != int64(0) {
							v221 = l1
						} else {
							v210 = l1 & int64(-66566)
							if v197&int64(3) == int64(0) {
								v217 = v210 | int64(4)
							} else {
								v217 = v210
							}
							v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
						}
						m.G0 = v9 + int32(64)
						return v221
					} else {
						if v17 != 0 {
							v221 = l1
							m.G0 = v9 + int32(64)
							return v221
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
							v183 = F_objectGetVal(m, v182)
							mBase = m.M
							v184 = int32(0)
							v189 = F_evalExtractShebangFlags(m, v183, v184, v9+int32(8), v184, v184)
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int64(0)
							} else {
								if v189 == int32(-1) {
									v221 = l1
								} else {
									v193 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
									v197 = v193
									if v197&int64(16) != int64(0) {
										v221 = l1
									} else {
										v210 = l1 & int64(-66566)
										if v197&int64(3) == int64(0) {
											v217 = v210 | int64(4)
										} else {
											v217 = v210
										}
										v221 = v197<<(uint(int64(8))%64)&int64(1024) | v197&int64(1) | v217 ^ int64(1)
									}
								}
								m.G0 = v9 + int32(64)
								return v221
							}
						}
					}
				}
			}
		default:
			v221 = l1
			m.G0 = v9 + int32(64)
			return v221
		}
	}
}
func F_execCommandAbort(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_resetClientMultiState(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v10 & int32(-4137)
		F_unwatchAllKeys(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1 + base.B2i32(v16 == int32(45))
			F_addReplyErrorFormat(m, l0, int32(_a_F_execCommandAbort_0), v6)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_execCommandAbort[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_replicationFeedMonitors(m, l0, v25, v27, v28, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
func F_fillCommandCDF(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v70 int64
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	F_addReplyMapLen(m, l0, int32(2))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_fillCommandCDF_0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	F_addReplyLongLong(m, l0, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_addReplyBulkCString(m, l0, int32(_a_F_fillCommandCDF_1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = v9 + int32(8)
	v27 = int64(1024)
	v29 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = l1
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(32)))) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(80)))) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v26)+88)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v26)+96)) = v27
	*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = float64(2)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = int32(1159)
	v70 = base.I64_extend_i32_u(int32(63) - (v60 + base.I32_wrap_i64(base.I64_clz(v61|v27))))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+104)) = base.I64_extend32_s(v27>>(uint(v70)%64)) << (uint(v70) % 64)
	goto L7
L7:
	;
	v78 = F_hdr_iter_next(m, v9+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_setDeferredMapLen(m, l0, v23, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	if v78 == int32(0) {
		v107 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v85 = int32(0)
	v88 = int64(0)
	goto L11
L11:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	if v90 <= v88 {
		v101 = v85
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v107 = v101
	goto L8
L13:
	;
	v104 = F_hdr_iter_next(m, v9+int32(8))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
	v94 = base.I64_div_s(v92, int64(1000))
	F_addReplyLongLong(m, l0, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_addReplyLongLong(m, l0, v90)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v101 = v85 + int32(1)
	goto L13
L17:
	;
	if v104 != 0 {
		v85 = v101
		v88 = v90
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	m.G0 = v9 + int32(128)
	return
}
func F_generateCommandInfoResponse(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v9 = F_createCachedResponseClient(m, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_addReplyArrayLen(m, v9, int32(10))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v19 < v17 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v14 = int64(0)
	v28 = v14
	v29 = v14
	v30 = v14
	goto L3
L6:
	;
	v22 = v17
	goto L8
L7:
	;
	v22 = v18
	goto L8
L8:
	;
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+188)))
	v28 = base.I64_extend_i32_s(v22 + v19)
	v29 = v25
	v30 = base.I64_extend_i32_s(v18)
	goto L3
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-1)))))
	switch v39 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		v56 = int32(0)
		goto L10
	}
L10:
	;
	F_addReplyBulkCBuffer(m, v9, v36, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v56 = v55
	goto L10
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v56 = v52
	goto L10
L13:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v56 = v49
	goto L10
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v56 = v46
	goto L10
L15:
	;
	v56 = int32(base.Ui32(v39) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+52)))
	F_addReplyLongLong(m, v9, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_addReplyFlagsForCommand(m, v9, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_addReplyLongLong(m, v9, v30)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_addReplyLongLong(m, v9, v28)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_addReplyLongLong(m, v9, v29)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_addReplyCommandCategories(m, v9, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_addReplySetLen(m, v9, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v75 < int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_addReplyCommandKeySpecs(m, v9, l0)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v80 = int32(0)
	goto L26
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v80<<(uint(int32(2))%32))))
	F_addReplyBulkCString(m, v9, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	v95 = v80 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v95 < v96 {
		v80 = v95
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	F_addReplyCommandSubCommands(m, v9, l0, int32(1029), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v112 = F_aggregateClientOutputBuffer(m, v9)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_deleteCachedResponseClient(m, v9)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	return v112
}
func F_getCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(2))%32))+uint32(_c_F_getCommand[0])))
	v11 = F_lookupKeyReadOrReply(m, l0, v4, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			return
		} else {
			v16 = F_checkType(m, l0, v11, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if v16 != 0 {
					return
				} else {
					F_addReplyBulk(m, l0, v11)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_lookupCommand(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_lookupCommand[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v18 = F_hashtableFind(m, v11, v15, v8+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		if v18 == int32(0) {
			v41 = v22
			m.G0 = v8 + int32(16)
			return v41
		} else {
			if l1 == int32(1) {
				v41 = v22
				m.G0 = v8 + int32(16)
				return v41
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
				if v27 == int32(0) {
					v41 = v22
					m.G0 = v8 + int32(16)
					return v41
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v31 = F_objectGetVal(m, v30)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
					v37 = F_hashtableFind(m, v34, v31, v8+int32(12))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v41 = v39
						m.G0 = v8 + int32(16)
						return v41
					}
				}
			}
		}
	}
}
func F_lookupCommandBySds(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_lookupCommandBySds[0]))
	v4 = F_lookupCommandBySdsLogic(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lookupCommandBySdsLogic(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v22 & int32(7) {
	case 0:
		v39 = int32(base.Ui32(v22) >> (uint(int32(3)) % 32))
	case 1:
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v39 = v29
	case 2:
		v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v39 = v32
	case 3:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v39 = v35
	case 4:
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v39 = v38
	default:
		v39 = v3
	}
	v44 = F_sdssplitlen(m, l1, v39, int32(_a_F_lookupCommandBySdsLogic_0), int32(1), v16+int32(4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		return int32(0)
	} else {
		if v44 == int32(0) {
			v208 = v3
			m.G0 = v16 + int32(16)
			return v208
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v51 = int32(-3)
			if base.Ui32(v51) < base.Ui32(v50+v51) {
				v59 = int32(15)
				v61 = int32(-16)
				v63 = v16 - (v50*int32(12)+v59)&v61
				m.G0 = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				v66 = int32(1)
				v74 = v63 - (v65<<(uint(int32(2))%32)+v59)&v61
				m.G0 = v74
				if v65 == v66 {
					v136 = int32(0)
				} else {
					v81 = int32(0)
					v84 = v81
					v94 = v81
					for {
						v96 = int32(12)
						v98 = v63 + v84*v96
						v99 = int32(-16)
						*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v99
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
						v102 = int32(-256)
						*(*int32)(unsafe.Add(mBase, uint32(v98))) = v101 & v102
						v105 = int32(2)
						v106 = v84 << (uint(v105) % 32)
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v44+v106)))
						*(*int32)(unsafe.Add(mBase, uint32(v98)+8)) = v108
						*(*int32)(unsafe.Add(mBase, uint32(v74+v106))) = v98
						v113 = v84 | int32(1)
						v116 = v63 + v113*v96
						*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v99
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
						*(*int32)(unsafe.Add(mBase, uint32(v116))) = v119 & v102
						v124 = v113 << (uint(v105) % 32)
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v44+v124)))
						*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v74+v124))) = v116
						v131 = v84 + v105
						v133 = v94 + v105
						if v133 != v65&int32(2147483646) {
							v84 = v131
							v94 = v133
							continue
						} else {
							break
						}
						break
					}
					v136 = v131
				}
				if v65&v66 == int32(0) {
				} else {
					v152 = v63 + v136*int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = int32(-16)
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
					*(*int32)(unsafe.Add(mBase, uint32(v152))) = v155 & int32(-256)
					v160 = v136 << (uint(int32(2)) % 32)
					v162 = *(*int32)(unsafe.Add(mBase, uint32(v44+v160)))
					*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v74+v160))) = v152
				}
				v168 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(0)
				v171 = F_objectGetVal(m, v168)
				mBase = m.M
				v174 = F_hashtableFind(m, l0, v171, v16+int32(8))
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return int32(0)
				} else {
					v176 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					if v174 == int32(0) {
						if v65 == int32(1) {
							v185 = v176
						} else {
							v185 = int32(0)
						}
						v199 = v185
						v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
						F_sdsfreesplitres(m, v44, v201)
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return int32(0)
						} else {
							v208 = v199
							m.G0 = v16 + int32(16)
							return v208
						}
					} else {
						if v65 == int32(1) {
							if v65 == int32(1) {
								v185 = v176
							} else {
								v185 = int32(0)
							}
							v199 = v185
							v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
							F_sdsfreesplitres(m, v44, v201)
							mBase = m.M
							v203 = m.ExcPending
							if v203 != 0 {
								return int32(0)
							} else {
								v208 = v199
								m.G0 = v16 + int32(16)
								return v208
							}
						} else {
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+200))
							if v181 != 0 {
								if v65 != int32(2) {
									v199 = int32(0)
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
									F_sdsfreesplitres(m, v44, v201)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										v208 = v199
										m.G0 = v16 + int32(16)
										return v208
									}
								} else {
									v189 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
									v190 = F_objectGetVal(m, v189)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
									v193 = *(*int32)(unsafe.Add(mBase, uint32(v176)+200))
									v196 = F_hashtableFind(m, v193, v190, v16+int32(12))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
										v199 = v198
										v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										F_sdsfreesplitres(m, v44, v201)
										mBase = m.M
										v203 = m.ExcPending
										if v203 != 0 {
											return int32(0)
										} else {
											v208 = v199
											m.G0 = v16 + int32(16)
											return v208
										}
									}
								}
							} else {
								if v65 == int32(1) {
									v185 = v176
								} else {
									v185 = int32(0)
								}
								v199 = v185
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
								F_sdsfreesplitres(m, v44, v201)
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return int32(0)
								} else {
									v208 = v199
									m.G0 = v16 + int32(16)
									return v208
								}
							}
						}
					}
				}
			} else {
				F_sdsfreesplitres(m, v44, v50)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v208 = v3
					m.G0 = v16 + int32(16)
					return v208
				}
			}
		}
	}
}
func F_prepareCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_prepareCommandGeneric(m, v2, v3, l0+int32(288), l0+int32(80), l0+int32(292))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_prepareCommandGeneric(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
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
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == int32(0) {
		m.G0 = v11 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v15&int32(8192) == int32(0) {
			m.G0 = v11 + int32(16)
			return
		} else {
			v20 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
			if v21 == v20 {
				v27 = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v27
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v32 = F_objectGetVal(m, v31)
				mBase = m.M
				v35 = F_hashtableFind(m, v28, v32, v11+int32(8))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					if v35 == int32(0) {
						v54 = v37
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v54
						if v54 != 0 {
							v67 = v54
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
							if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
								m.G0 = v11 + int32(16)
								return
							} else {
								if int32(0)-v69 <= l1 {
									v81 = int32(0)
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
									if v82 == v81 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v85 = int32(0)
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
										if v86 == v85 {
											v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
												m.G0 = v11 + int32(16)
												return
											}
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											if v89 != int32(-1) {
												F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
												v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
												if v92&int32(24) != 0 {
													F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
													v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
									m.G0 = v11 + int32(16)
									return
								}
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57 | int32(131072)
							m.G0 = v11 + int32(16)
							return
						}
					} else {
						if l1 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
							v67 = v37
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
							if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
								m.G0 = v11 + int32(16)
								return
							} else {
								if int32(0)-v69 <= l1 {
									v81 = int32(0)
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
									if v82 == v81 {
										m.G0 = v11 + int32(16)
										return
									} else {
										v85 = int32(0)
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
										if v86 == v85 {
											v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
												m.G0 = v11 + int32(16)
												return
											}
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
											if v89 != int32(-1) {
												F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
												v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
												if v92&int32(24) != 0 {
													F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
													v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
									m.G0 = v11 + int32(16)
									return
								}
							}
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
							if v42 != 0 {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v45 = F_objectGetVal(m, v44)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
								v51 = F_hashtableFind(m, v48, v45, v11+int32(12))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v54 = v53
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v54
									if v54 != 0 {
										v67 = v54
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
										if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
											m.G0 = v11 + int32(16)
											return
										} else {
											if int32(0)-v69 <= l1 {
												v81 = int32(0)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
												if v82 == v81 {
													m.G0 = v11 + int32(16)
													return
												} else {
													v85 = int32(0)
													v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
													if v86 == v85 {
														v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
															m.G0 = v11 + int32(16)
															return
														}
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														if v89 != int32(-1) {
															F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
															v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
															if v92&int32(24) != 0 {
																F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
																v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																	m.G0 = v11 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
												m.G0 = v11 + int32(16)
												return
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57 | int32(131072)
										m.G0 = v11 + int32(16)
										return
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
								v67 = v37
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
								if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
									m.G0 = v11 + int32(16)
									return
								} else {
									if int32(0)-v69 <= l1 {
										v81 = int32(0)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
										if v82 == v81 {
											m.G0 = v11 + int32(16)
											return
										} else {
											v85 = int32(0)
											v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
											if v86 == v85 {
												v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
													m.G0 = v11 + int32(16)
													return
												}
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												if v89 != int32(-1) {
													F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
													v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
													if v92&int32(24) != 0 {
														F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
														v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
										m.G0 = v11 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				if v15&int32(131072) != 0 {
					F__serverAssert(m, int32(_a_F_prepareCommandGeneric_2), int32(_a_F_prepareCommandGeneric_1), int32(4229))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v26 != 0 {
						F__serverAssert(m, int32(_a_F_prepareCommandGeneric_2), int32(_a_F_prepareCommandGeneric_1), int32(4229))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v27 = int32(0)
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v27
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v32 = F_objectGetVal(m, v31)
						mBase = m.M
						v35 = F_hashtableFind(m, v28, v32, v11+int32(8))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							if v35 == int32(0) {
								v54 = v37
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v54
								if v54 != 0 {
									v67 = v54
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
									if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
										m.G0 = v11 + int32(16)
										return
									} else {
										if int32(0)-v69 <= l1 {
											v81 = int32(0)
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
											if v82 == v81 {
												m.G0 = v11 + int32(16)
												return
											} else {
												v85 = int32(0)
												v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
												if v86 == v85 {
													v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
														m.G0 = v11 + int32(16)
														return
													}
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													if v89 != int32(-1) {
														F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
														v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
														if v92&int32(24) != 0 {
															F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
															v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																m.G0 = v11 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
											m.G0 = v11 + int32(16)
											return
										}
									}
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57 | int32(131072)
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								if l1 == int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
									v67 = v37
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
									if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
										m.G0 = v11 + int32(16)
										return
									} else {
										if int32(0)-v69 <= l1 {
											v81 = int32(0)
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
											if v82 == v81 {
												m.G0 = v11 + int32(16)
												return
											} else {
												v85 = int32(0)
												v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
												if v86 == v85 {
													v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
														m.G0 = v11 + int32(16)
														return
													}
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													if v89 != int32(-1) {
														F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
														v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
														if v92&int32(24) != 0 {
															F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
															v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																m.G0 = v11 + int32(16)
																return
															}
														}
													}
												}
											}
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
											m.G0 = v11 + int32(16)
											return
										}
									}
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
									if v42 != 0 {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v45 = F_objectGetVal(m, v44)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+200))
										v51 = F_hashtableFind(m, v48, v45, v11+int32(12))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v54 = v53
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v54
											if v54 != 0 {
												v67 = v54
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
												if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
													m.G0 = v11 + int32(16)
													return
												} else {
													if int32(0)-v69 <= l1 {
														v81 = int32(0)
														v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
														if v82 == v81 {
															m.G0 = v11 + int32(16)
															return
														} else {
															v85 = int32(0)
															v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
															if v86 == v85 {
																v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																	m.G0 = v11 + int32(16)
																	return
																}
															} else {
																v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
																if v89 != int32(-1) {
																	F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
																	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
																	if v92&int32(24) != 0 {
																		F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
																		v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																			m.G0 = v11 + int32(16)
																			return
																		}
																	}
																}
															}
														}
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
														*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
														m.G0 = v11 + int32(16)
														return
													}
												}
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v57 | int32(131072)
												m.G0 = v11 + int32(16)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v37
										v67 = v37
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+52))
										if base.B2i32(int32(0) < v69)&base.B2i32(v69 != l1) != 0 {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
											m.G0 = v11 + int32(16)
											return
										} else {
											if int32(0)-v69 <= l1 {
												v81 = int32(0)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[2]))
												if v82 == v81 {
													m.G0 = v11 + int32(16)
													return
												} else {
													v85 = int32(0)
													v86 = *(*int32)(unsafe.Add(mBase, _c_F_prepareCommandGeneric[0]))
													if v86 == v85 {
														v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
															m.G0 = v11 + int32(16)
															return
														}
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														if v89 != int32(-1) {
															F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
															v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
															if v92&int32(24) != 0 {
																F__serverAssert(m, int32(_a_F_prepareCommandGeneric_0), int32(_a_F_prepareCommandGeneric_1), int32(4238))
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
																v95 = F_clusterSlotByCommand(m, v67, l0, l1, l2)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
																	m.G0 = v11 + int32(16)
																	return
																}
															}
														}
													}
												}
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = v77 | int32(262144)
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
func F_printCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
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
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v411 int32
	_ = v411
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v837 int32
	_ = v837
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v920 int32
	_ = v920
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v960 = m.G3
	v966 = m.G8
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	m.T0[v967].(func(*base.Module, int32, int32, int32))(m, v960+int32(_a_F_printCommandHandler_0), v960+int32(_a_F_printCommandHandler_1), int32(585))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L6
	} else {
		goto L241
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l1 != int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v951 = m.G14
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	m.T0[v952].(func(*base.Module))(m)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L6
	} else {
		goto L240
	}
L4:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	goto L142
L5:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = m.G7
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, v18, v17)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	goto L11
L8:
	;
	v428 = m.G3
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+uint32(_c_F_printCommandHandler[0]))))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v434 == int32(0) {
		v457 = v433
		v458 = v434
		goto L108
	} else {
		goto L109
	}
L9:
	;
	if v79 == int32(0) {
		goto L8
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	goto L20
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96)) = v68
	v79 = int32(1)
	goto L10
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if base.Ui32(v32) <= base.Ui32(v64) {
		v79 = int32(0)
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v68 = base.I32_div_s(v32-v64, int32(24))
	goto L19
L22:
	;
	v86 = v17
	goto L24
L23:
	;
	v375 = m.G3
	v376 = m.G13
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v378 = m.G12
	v383 = m.T0[v377].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v375+int32(_a_F_printCommandHandler_2), int32(8))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L94
	}
L24:
	;
	v89 = int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96))
	v100 = v96 + v97*int32(24)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	if v102 != int32(6) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v318 = int32(1)
	v319 = v86 + v318
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v319 < v318 {
		v350 = v319
		v352 = v325
		goto L82
	} else {
		goto L83
	}
L27:
	;
	if v162 == int32(0) {
		goto L26
	} else {
		goto L43
	}
L28:
	;
	goto L27
L29:
	;
	F_luaA_pushobject(m, v14, v151+int32(0))
	mBase = m.M
	v162 = v152
	goto L28
L30:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v100 == v136 {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+6)))
	if v106 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	if v107 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v100 == v110 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v125 = F_luaF_getlocalname(m, v107, v89, (v118-v119)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v125 == int32(0) {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v117 = v116
	v118 = v113
	goto L34
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v117 = v107
	v118 = v112
	goto L34
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v151 = v128
	v152 = v125
	goto L29
L38:
	;
	v138 = v14 + int32(8)
	goto L40
L39:
	;
	v138 = v100 + int32(28)
	goto L40
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v141 = m.G3
	goto L41
L41:
	;
	if (v139-v140)>>(uint(int32(4))%32) < v89 {
		v162 = int32(0)
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v151 = v140
	v152 = v141 + int32(_a_F_printCommandHandler_3)
	goto L29
L43:
	;
	v167 = v89
	v168 = v162
	goto L44
L44:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v177 == int32(0) {
		v200 = v176
		v201 = v177
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L26
L46:
	;
	if v201-v200&int32(255) == int32(0) {
		goto L23
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v177 != v176&int32(255) {
		v200 = v176
		v201 = v177
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v183 = v22
	v184 = v168
	goto L50
L50:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v188 == int32(0) {
		v200 = v187
		v201 = v188
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v200 = v187
	v201 = v188
	goto L47
L52:
	;
	v191 = int32(1)
	if v188 == v187&int32(255) {
		v183 = v183 + v191
		v184 = v184 + v191
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L57
L55:
	;
	v238 = v167 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96))
	v246 = v242 + v243*int32(24)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	if v248 != int32(6) {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v225 + int32(-16)
	goto L55
L57:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L56
L63:
	;
	if v308 != 0 {
		v167 = v238
		v168 = v308
		goto L44
	} else {
		goto L79
	}
L64:
	;
	goto L63
L65:
	;
	F_luaA_pushobject(m, v14, v297+v238<<(uint(int32(4))%32)+int32(-16))
	mBase = m.M
	v308 = v298
	goto L64
L66:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v246 == v282 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+6)))
	if v252 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v253 == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v246 == v256 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v271 = F_luaF_getlocalname(m, v253, v238, (v264-v265)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v271 == int32(0) {
		goto L66
	} else {
		goto L73
	}
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	v263 = v262
	v264 = v259
	goto L70
L72:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v263 = v253
	v264 = v258
	goto L70
L73:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v297 = v274
	v298 = v271
	goto L65
L74:
	;
	v284 = v14 + int32(8)
	goto L76
L75:
	;
	v284 = v246 + int32(28)
	goto L76
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v287 = m.G3
	v288 = int32(0)
	if v238 < int32(1) {
		v308 = v288
		goto L64
	} else {
		goto L77
	}
L77:
	;
	if (v285-v286)>>(uint(int32(4))%32) < v238 {
		v308 = v288
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v297 = v286
	v298 = v287 + int32(_a_F_printCommandHandler_3)
	goto L65
L79:
	;
	goto L45
L80:
	;
	if v372 == int32(0) {
		goto L8
	} else {
		goto L93
	}
L81:
	;
	goto L80
L82:
	;
	if v350 != 0 {
		v363 = int32(0)
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v331 = v319
	v333 = v325
	goto L84
L84:
	;
	if base.Ui32(v333) <= base.Ui32(v328) {
		v372 = int32(0)
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v350 = v344
	v352 = v346
	goto L82
L86:
	;
	v338 = v331 + int32(-1)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+6)))
	if v341 != 0 {
		v344 = v338
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v346 = v333 + int32(-24)
	if int32(0) < v344 {
		v331 = v344
		v333 = v346
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v333)+20))
	v344 = v338 - v342
	goto L87
L89:
	;
	goto L85
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96)) = v363
	v372 = int32(1)
	goto L81
L91:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if base.Ui32(v352) <= base.Ui32(v357) {
		v372 = int32(0)
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v361 = base.I32_div_s(v352-v357, int32(24))
	v363 = v361
	goto L90
L93:
	;
	v86 = v319
	goto L24
L94:
	;
	v387 = F_ldbCatStackValueRec(m, v383, v14, int32(-1), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	m.T0[v390].(func(*base.Module, int32, int32))(m, v387, int32(1))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	goto L99
L97:
	;
	goto L3
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v411 + int32(-16)
	goto L97
L99:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L98
L105:
	;
	v547 = m.G3
	v548 = m.G13
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v550 = m.G12
	v555 = m.T0[v549].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v547+int32(_a_F_printCommandHandler_4), int32(17))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L6
	} else {
		goto L137
	}
L106:
	;
	F_lua_getfield(m, v14, int32(-10002), v22)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L125
	}
L107:
	;
	if v458-v457&int32(255) == int32(0) {
		goto L106
	} else {
		goto L115
	}
L108:
	;
	goto L107
L109:
	;
	if v434 != v433&int32(255) {
		v457 = v433
		v458 = v434
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v440 = v22
	v441 = v428 + int32(_a_F_printCommandHandler_5)
	goto L111
L111:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441)+1)))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)))
	if v445 == int32(0) {
		v457 = v444
		v458 = v445
		goto L108
	} else {
		goto L113
	}
L112:
	;
	v457 = v444
	v458 = v445
	goto L108
L113:
	;
	v448 = int32(1)
	if v445 == v444&int32(255) {
		v440 = v440 + v448
		v441 = v441 + v448
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v464 = m.G3
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+uint32(_c_F_printCommandHandler[1]))))
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v470 == int32(0) {
		v493 = v469
		v494 = v470
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v494-v493&int32(255) != 0 {
		goto L105
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v470 != v469&int32(255) {
		v493 = v469
		v494 = v470
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v476 = v22
	v477 = v464 + int32(_a_F_printCommandHandler_6)
	goto L120
L120:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+1)))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+1)))
	if v481 == int32(0) {
		v493 = v480
		v494 = v481
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v493 = v480
	v494 = v481
	goto L117
L122:
	;
	v484 = int32(1)
	if v481 == v480&int32(255) {
		v476 = v476 + v484
		v477 = v477 + v484
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	goto L106
L125:
	;
	v501 = m.G3
	v502 = m.G13
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v504 = m.G12
	v509 = m.T0[v503].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v501+int32(_a_F_printCommandHandler_2), int32(8))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	v513 = F_ldbCatStackValueRec(m, v509, v14, int32(-1), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	m.T0[v516].(func(*base.Module, int32, int32))(m, v513, int32(1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	goto L131
L129:
	;
	goto L3
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v537 + int32(-16)
	goto L129
L131:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L130
L137:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	m.T0[v558].(func(*base.Module, int32, int32))(m, v555, int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L3
L139:
	;
	v930 = m.G3
	v931 = m.G13
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)))
	v933 = m.G12
	v938 = m.T0[v932].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v930+int32(_a_F_printCommandHandler_7), int32(42))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L6
	} else {
		goto L238
	}
L140:
	;
	if v615 == int32(0) {
		goto L139
	} else {
		goto L153
	}
L141:
	;
	goto L140
L142:
	;
	goto L151
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96)) = v604
	v615 = int32(1)
	goto L141
L151:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if base.Ui32(v568) <= base.Ui32(v600) {
		v615 = int32(0)
		goto L141
	} else {
		goto L152
	}
L152:
	;
	v604 = base.I32_div_s(v568-v600, int32(24))
	goto L150
L153:
	;
	v618 = int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96))
	v629 = v625 + v626*int32(24)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+8))
	if v631 != int32(6) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	if v691 == int32(0) {
		goto L139
	} else {
		goto L170
	}
L155:
	;
	goto L154
L156:
	;
	F_luaA_pushobject(m, v14, v680+int32(0))
	mBase = m.M
	v691 = v681
	goto L155
L157:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v629 == v665 {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+6)))
	if v635 != 0 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v634)+16))
	if v636 == int32(0) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v629 == v639 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	v654 = F_luaF_getlocalname(m, v636, v618, (v647-v648)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v654 == int32(0) {
		goto L157
	} else {
		goto L164
	}
L162:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v629)+12)) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v644)+16))
	v646 = v645
	v647 = v642
	goto L161
L163:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v629)+12))
	v646 = v636
	v647 = v641
	goto L161
L164:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	v680 = v657
	v681 = v654
	goto L156
L165:
	;
	v667 = v14 + int32(8)
	goto L167
L166:
	;
	v667 = v629 + int32(28)
	goto L167
L167:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	v670 = m.G3
	goto L168
L168:
	;
	if (v668-v669)>>(uint(int32(4))%32) < v618 {
		v691 = int32(0)
		goto L155
	} else {
		goto L169
	}
L169:
	;
	v680 = v669
	v681 = v670 + int32(_a_F_printCommandHandler_3)
	goto L156
L170:
	;
	v696 = v618
	v697 = v691
	v701 = int32(0)
	goto L171
L171:
	;
	v703 = m.G3
	v705 = v703 + int32(_a_F_printCommandHandler_3)
	v708 = int32(*(*int8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_printCommandHandler[2]))))
	if v708 != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	if v817 != 0 {
		goto L3
	} else {
		goto L237
	}
L173:
	;
	goto L214
L174:
	;
	if v733 != 0 {
		v817 = v701
		goto L173
	} else {
		goto L190
	}
L175:
	;
	v709 = int32(0)
	v710 = F_strchr(m, v697, v708)
	mBase = m.M
	if v710 == v709 {
		v730 = v709
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v733 = v697
	goto L174
L177:
	;
	v733 = v730
	goto L174
L178:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_printCommandHandler[3]))))
	if v713 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+1)))
	if v714 == int32(0) {
		v730 = v709
		goto L177
	} else {
		goto L181
	}
L180:
	;
	v733 = v710
	goto L174
L181:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_printCommandHandler[4]))))
	if v717 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+2)))
	if v719 == int32(0) {
		v730 = v709
		goto L177
	} else {
		goto L184
	}
L183:
	;
	v718 = F_twobyte_strstr(m, v710, v705)
	mBase = m.M
	v733 = v718
	goto L174
L184:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_printCommandHandler[5]))))
	if v722 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710)+3)))
	if v724 == int32(0) {
		v730 = v709
		goto L177
	} else {
		goto L187
	}
L186:
	;
	v723 = F_threebyte_strstr(m, v710, v705)
	mBase = m.M
	v733 = v723
	goto L174
L187:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_printCommandHandler[6]))))
	if v727 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v729 = F_twoway_strstr(m, v710, v705)
	mBase = m.M
	v730 = v729
	goto L177
L189:
	;
	v728 = F_fourbyte_strstr(m, v710, v705)
	mBase = m.M
	v733 = v728
	goto L174
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v697
	v735 = m.G3
	v736 = m.G13
	v739 = F_lm_asprintf(m, v735+int32(_a_F_printCommandHandler_8), v10)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	v742 = m.G12
	v743 = int32(0)
	if v739&int32(3) == v743 {
		v765 = v739
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v799 = m.T0[v741].(func(*base.Module, int32, int32, int32) int32)(m, v743, v739, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L208
	}
L193:
	;
	v798 = v790 - v739
	goto L192
L194:
	;
	v769 = v765
	goto L202
L195:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	if v751 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v754 = v739
	goto L198
L197:
	;
	v798 = v739 - v739
	goto L192
L198:
	;
	v758 = v754 + int32(1)
	if v758&int32(3) == int32(0) {
		v765 = v758
		goto L194
	} else {
		goto L200
	}
L200:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758))))
	if v763 != 0 {
		v754 = v758
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v790 = v758
	goto L193
L202:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v769)))
	v778 = int32(-2139062144)
	if (int32(16843008)-v775|v775)&v778 == v778 {
		v769 = v769 + int32(4)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v784 = v769
	goto L205
L204:
	;
	goto L203
L205:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	if v788 != 0 {
		v784 = v784 + int32(1)
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v790 = v784
	goto L193
L207:
	;
	goto L206
L208:
	;
	v803 = F_ldbCatStackValueRec(m, v799, v14, int32(-1), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	m.T0[v806].(func(*base.Module, int32, int32))(m, v803, int32(1))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v809 = m.G11
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	m.T0[v810].(func(*base.Module, int32))(m, v739)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L211
	}
L211:
	;
	v817 = v701 + int32(1)
	goto L173
L212:
	;
	v850 = v696 + int32(1)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(12))+96))
	v858 = v854 + v855*int32(24)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+8))
	if v860 != int32(6) {
		goto L223
	} else {
		goto L224
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v837 + int32(-16)
	goto L212
L214:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L213
L220:
	;
	if v920 != 0 {
		v696 = v850
		v697 = v920
		v701 = v817
		goto L171
	} else {
		goto L236
	}
L221:
	;
	goto L220
L222:
	;
	F_luaA_pushobject(m, v14, v909+v850<<(uint(int32(4))%32)+int32(-16))
	mBase = m.M
	v920 = v910
	goto L221
L223:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v858 == v894 {
		goto L231
	} else {
		goto L232
	}
L224:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863)+6)))
	if v864 != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v863)+16))
	if v865 == int32(0) {
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v858 == v868 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v875)+12))
	v883 = F_luaF_getlocalname(m, v865, v850, (v876-v877)>>(uint(int32(2))%32)+int32(-1))
	mBase = m.M
	if v883 == int32(0) {
		goto L223
	} else {
		goto L230
	}
L228:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v858)+12)) = v871
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	v875 = v874
	v876 = v871
	goto L227
L229:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	v875 = v865
	v876 = v870
	goto L227
L230:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v909 = v886
	v910 = v883
	goto L222
L231:
	;
	v896 = v14 + int32(8)
	goto L233
L232:
	;
	v896 = v858 + int32(28)
	goto L233
L233:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v899 = m.G3
	v900 = int32(0)
	if v850 < int32(1) {
		v920 = v900
		goto L221
	} else {
		goto L234
	}
L234:
	;
	if (v897-v898)>>(uint(int32(4))%32) < v850 {
		v920 = v900
		goto L221
	} else {
		goto L235
	}
L235:
	;
	v909 = v898
	v910 = v899 + int32(_a_F_printCommandHandler_3)
	goto L222
L236:
	;
	goto L172
L237:
	;
	goto L139
L238:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	m.T0[v941].(func(*base.Module, int32, int32))(m, v938, int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L6
	} else {
		goto L239
	}
L239:
	;
	goto L3
L240:
	;
	m.G0 = v10 + int32(112)
	return int32(1)
L241:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_printCommandHelp(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v101 == int32(0) {
		v135 = v97
		goto L27
	} else {
		goto L28
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v17
	v95 = F_sdscatfmt(m, v15, int32(_a_F_printCommandHelp_0), v13+int32(64))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L26
	}
L5:
	;
	if v17&int32(3) == int32(0) {
		v42 = v17
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if base.Ui32(v75) <= base.Ui32(v18) {
		goto L4
	} else {
		goto L22
	}
L7:
	;
	v75 = v67 - v17
	goto L6
L8:
	;
	v46 = v42
	goto L16
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v17
	goto L12
L11:
	;
	v75 = v17 - v17
	goto L6
L12:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v67 = v35
	goto L7
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v61 = v46
	goto L19
L18:
	;
	goto L17
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v67 = v61
	goto L7
L21:
	;
	goto L20
L22:
	;
	v77 = F_sdsnewlen(m, v17, v18)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v80 + v79
	v87 = F_sdscatfmt(m, v15, int32(_a_F_printCommandHelp_1), v13+int32(80))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_sdsfree(m, v77)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v97 = v87
	goto L3
L26:
	;
	v97 = v95
	goto L3
L27:
	;
	v143 = int32(0)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(-1)))))
	switch v147 & int32(7) {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	default:
		v164 = v143
		goto L36
	}
L28:
	;
	v107 = v97
	v108 = int32(0)
	goto L29
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v118 = v115 + v108*int32(12)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v120
	if v119 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v135 = v127
	goto L27
L31:
	;
	v124 = int32(_a_F_printCommandHelp_2)
	goto L33
L32:
	;
	v124 = int32(_a_F_printCommandHelp_3)
	goto L33
L33:
	;
	v127 = F_sdscatfmt(m, v107, v124, v13+int32(48))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v130 = v108 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v130) < base.Ui32(v131) {
		v107 = v127
		v108 = v130
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = int32(_a_F_printCommandHelp_4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v164 + int32(-20)
	v173 = F_sdscatprintf(m, v135, int32(_a_F_printCommandHelp_5), v13+int32(32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L49
	}
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v135+int32(-17))))
	v164 = v163
	goto L36
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v135+int32(-9))))
	v164 = v160
	goto L36
L39:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135+int32(-5)))))
	v164 = v157
	goto L36
L40:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+int32(-3)))))
	v164 = v154
	goto L36
L41:
	;
	v164 = int32(base.Ui32(v147) >> (uint(int32(3)) % 32))
	goto L36
L42:
	;
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v208&int32(3) == v207 {
		v230 = v208
		goto L56
	} else {
		goto L57
	}
L43:
	;
	if v194 <= int32(21) {
		v206 = v143
		goto L42
	} else {
		goto L50
	}
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v173+int32(-17))))
	v194 = v193
	goto L43
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v173+int32(-9))))
	v194 = v190
	goto L43
L46:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173+int32(-5)))))
	v194 = v187
	goto L43
L47:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+int32(-3)))))
	v194 = v184
	goto L43
L48:
	;
	v194 = int32(base.Ui32(v177) >> (uint(int32(3)) % 32))
	goto L43
L49:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+int32(-1)))))
	switch v177 & int32(7) {
	case 0:
		goto L48
	case 1:
		goto L47
	case 2:
		goto L46
	case 3:
		goto L45
	case 4:
		goto L44
	default:
		v206 = v143
		goto L42
	}
L50:
	;
	v198 = F_createObject(m, int32(0), v173)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_printCommandHelp[0]))
	v202 = F_listAddNodeTail(m, v201, v198)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v206 = int32(1)
	goto L42
L53:
	;
	F_valkey_free(m, v458)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L121
	}
L54:
	;
	if v263 == int32(0) {
		v458 = v207
		goto L53
	} else {
		goto L70
	}
L55:
	;
	v263 = v255 - v208
	goto L54
L56:
	;
	v234 = v230
	goto L64
L57:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v216 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v219 = v208
	goto L60
L59:
	;
	v263 = v208 - v208
	goto L54
L60:
	;
	v223 = v219 + int32(1)
	if v223&int32(3) == int32(0) {
		v230 = v223
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v228 != 0 {
		v219 = v223
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v255 = v223
	goto L55
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v243 = int32(-2139062144)
	if (int32(16843008)-v240|v240)&v243 == v243 {
		v234 = v234 + int32(4)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v249 = v234
	goto L67
L66:
	;
	goto L65
L67:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v253 != 0 {
		v249 = v249 + int32(1)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v255 = v249
	goto L55
L69:
	;
	goto L68
L70:
	;
	v270 = v208
	v271 = v207
	v272 = int32(0)
	goto L71
L71:
	;
	if v270&int32(3) == int32(0) {
		v298 = v270
		goto L75
	} else {
		goto L76
	}
L72:
	;
	if v370 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L73:
	;
	v333 = F_valkey_malloc(m, int32(50))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L89
	}
L74:
	;
	v331 = v323 - v270
	goto L73
L75:
	;
	v302 = v298
	goto L83
L76:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v284 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v287 = v270
	goto L79
L78:
	;
	v331 = v270 - v270
	goto L73
L79:
	;
	v291 = v287 + int32(1)
	if v291&int32(3) == int32(0) {
		v298 = v291
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v296 != 0 {
		v287 = v291
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v323 = v291
	goto L74
L83:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v311 = int32(-2139062144)
	if (int32(16843008)-v308|v308)&v311 == v311 {
		v302 = v302 + int32(4)
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v317 = v302
	goto L86
L85:
	;
	goto L84
L86:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v321 != 0 {
		v317 = v317 + int32(1)
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v323 = v317
	goto L74
L88:
	;
	goto L87
L89:
	;
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v333)+49)) = uint8(v335)
	v338 = F___stpncpy(m, v333, v270, int32(49))
	mBase = m.M
	goto L90
L90:
	;
	if base.Ui32(v331) < base.Ui32(int32(50)) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v358 = v271 << (uint(int32(2)) % 32)
	v361 = F_valkey_realloc(m, v272, v358+int32(4))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L97
	}
L92:
	;
	v355 = v270 + v331
	goto L91
L93:
	;
	v342 = F_strlen(m, v333)
	mBase = m.M
	v345 = F___memrchr(m, v333, int32(32), v342+int32(1))
	mBase = m.M
	goto L95
L94:
	;
	v355 = v270 + (v345 - v333) + int32(1)
	goto L91
L95:
	;
	if v345 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v348 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v348)
	goto L94
L97:
	;
	v364 = F_sdsnew(m, v333)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361+v358))) = v364
	F_valkey_free(m, v333)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v370 = v271 + int32(1)
	if base.Ui32(v355-v208) < base.Ui32(v263) {
		v270 = v355
		v271 = v370
		v272 = v361
		goto L71
	} else {
		goto L100
	}
L100:
	;
	goto L72
L101:
	;
	v458 = v361
	goto L53
L102:
	;
	if v206 != 0 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v393 = F_createObject(m, int32(0), v391)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L109
	}
L104:
	;
	v378 = F_sdsempty(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v376 = F_sdscatsds(m, v173, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v391 = v376
	goto L103
L107:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(_a_F_printCommandHelp_4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(21)
	v389 = F_sdscatprintf(m, v378, int32(_a_F_printCommandHelp_6), v13+int32(16))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v391 = v389
	goto L103
L109:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_printCommandHelp[0]))
	v397 = F_listAddNodeTail(m, v396, v393)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	F_sdsfree(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v271 == int32(0) {
		goto L101
	} else {
		goto L112
	}
L112:
	;
	v408 = int32(1)
	goto L113
L113:
	;
	v415 = F_sdsempty(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L115
	}
L114:
	;
	goto L101
L115:
	;
	v419 = v361 + v408<<(uint(int32(2))%32)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(_a_F_printCommandHelp_4)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(21)
	v428 = F_sdscatprintf(m, v415, int32(_a_F_printCommandHelp_6), v13)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v430 = F_createObject(m, int32(0), v428)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_printCommandHelp[0]))
	v434 = F_listAddNodeTail(m, v433, v430)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	F_sdsfree(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if base.B2i32(v408 == v271) == int32(0) {
		v408 = v408 + int32(1)
		goto L113
	} else {
		goto L120
	}
L120:
	;
	goto L114
L121:
	;
	m.G0 = v13 + int32(96)
	return
}
func F_processCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int64
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int64
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v549 int64
	_ = v549
	var v550 int64
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int64
	_ = v616
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int64
	_ = v715
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int64
	_ = v741
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int64
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int64
	_ = v780
	var v782 int64
	_ = v782
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v956 int32
	_ = v956
	var v969 int32
	_ = v969
	var v974 int64
	_ = v974
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1135 int32
	_ = v1135
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1343 int32
	_ = v1343
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1381 int32
	_ = v1381
	var v1394 int32
	_ = v1394
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[0]))
	if v20 == v2 {
		v28 = v2
		goto L13
	} else {
		goto L14
	}
L1:
	;
	m.G0 = v14 + int32(112)
	return v1511
L2:
	;
	F_blockPostponeClient(m, l0)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L24
	} else {
		goto L455
	}
L3:
	;
	v1493 = int32(0)
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[1]))
	F_rejectCommand(m, l0, v1495, int32(1))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L24
	} else {
		goto L454
	}
L4:
	;
	v1403 = F_writeCommandsGetDiskErrorMessage(m, v693)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L24
	} else {
		goto L421
	}
L5:
	;
	F__serverPanic_1(m, int32(_a_F_processCommand_0), int32(4529), int32(_a_F_processCommand_1), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L24
	} else {
		goto L420
	}
L6:
	;
	v1287 = F_sdsempty(m)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L24
	} else {
		goto L392
	}
L7:
	;
	F__serverAssert(m, int32(_a_F_processCommand_2), int32(_a_F_processCommand_0), int32(4343))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L24
	} else {
		goto L391
	}
L8:
	;
	F__serverAssert(m, int32(_a_F_processCommand_3), int32(_a_F_processCommand_0), int32(4320))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L24
	} else {
		goto L390
	}
L9:
	;
	F__serverAssert(m, int32(_a_F_processCommand_4), int32(_a_F_processCommand_0), int32(4285))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L24
	} else {
		goto L389
	}
L10:
	;
	F__serverAssert(m, int32(_a_F_processCommand_5), int32(_a_F_processCommand_0), int32(4284))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L24
	} else {
		goto L388
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v35 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	if v28 != 0 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v28 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & int32(1)
	goto L13
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[2]))
	if v30 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v31 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[0]))
	goto L17
L17:
	;
	if v32 != v31 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v322)+56))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322)+48))
	if v325 == int32(295) {
		goto L118
	} else {
		goto L119
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v49
	v153 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[3]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v157&int32(6) == int32(4) {
		v169 = v153
		goto L64
	} else {
		goto L65
	}
L21:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+290)))
	if v145&int32(2) == int32(0) {
		goto L8
	} else {
		goto L61
	}
L22:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[4]))
	if v140 == v139 {
		v322 = v35
		goto L19
	} else {
		goto L59
	}
L23:
	;
	F_moduleCallCommandFilters(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L26
L26:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[4]))
	if v42 == v41 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v49 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	if v42&int32(2) == int32(0) {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = F_objectGetVal(m, v51)
	mBase = m.M
	v53 = int32(_a_F_processCommand_6)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	F_securityWarningCommand(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L24
	} else {
		goto L58
	}
L32:
	;
	if v88-v90 == int32(0) {
		goto L31
	} else {
		goto L44
	}
L33:
	;
	v88 = F_tolower(m, v84)
	mBase = m.M
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v90 = F_tolower(m, v89)
	mBase = m.M
	goto L32
L34:
	;
	v58 = v52
	v59 = v53
	v60 = v56
	goto L37
L35:
	;
	v84 = int32(0)
	v85 = v53
	goto L33
L36:
	;
	v84 = v81 & int32(255)
	v85 = v80
	goto L33
L37:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v62 == int32(0) {
		v80 = v59
		v81 = v60
		goto L36
	} else {
		goto L39
	}
L38:
	;
	v80 = v74
	v81 = int32(0)
	goto L36
L39:
	;
	v66 = v60 & int32(255)
	if v66 == v62 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v73 = int32(1)
	v74 = v59 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v75 != 0 {
		v58 = v58 + v73
		v59 = v74
		v60 = v75
		goto L37
	} else {
		goto L43
	}
L41:
	;
	v68 = F_tolower(m, v66)
	mBase = m.M
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v70 = F_tolower(m, v69)
	mBase = m.M
	if v68 == v70 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v80 = v59
	v81 = v72
	goto L36
L43:
	;
	goto L38
L44:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = F_objectGetVal(m, v95)
	mBase = m.M
	v97 = int32(_a_F_processCommand_7)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v132-v134 != 0 {
		goto L21
	} else {
		goto L57
	}
L46:
	;
	v132 = F_tolower(m, v128)
	mBase = m.M
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v134 = F_tolower(m, v133)
	mBase = m.M
	goto L45
L47:
	;
	v102 = v96
	v103 = v97
	v104 = v100
	goto L50
L48:
	;
	v128 = int32(0)
	v129 = v97
	goto L46
L49:
	;
	v128 = v125 & int32(255)
	v129 = v124
	goto L46
L50:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v106 == int32(0) {
		v124 = v103
		v125 = v104
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v124 = v118
	v125 = int32(0)
	goto L49
L52:
	;
	v110 = v104 & int32(255)
	if v110 == v106 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v117 = int32(1)
	v118 = v103 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v119 != 0 {
		v102 = v102 + v117
		v103 = v118
		v104 = v119
		goto L50
	} else {
		goto L56
	}
L54:
	;
	v112 = F_tolower(m, v110)
	mBase = m.M
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v114 = F_tolower(m, v113)
	mBase = m.M
	if v112 == v114 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v124 = v103
	v125 = v116
	goto L49
L56:
	;
	goto L51
L57:
	;
	goto L31
L58:
	;
	v1511 = int32(-1)
	goto L1
L59:
	;
	if v140&int32(2) != 0 {
		v322 = v35
		goto L19
	} else {
		goto L60
	}
L60:
	;
	goto L2
L61:
	;
	goto L20
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v205 & int32(-1073741825)
	v211 = F_commandCheckExistence(m, l0, v14+int32(108))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L24
	} else {
		goto L79
	}
L63:
	;
	if v169 == int32(0) {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v169 = int32(base.Ui32(v162^int32(-1))>>(uint(int32(23))%32)) & int32(1)
	goto L64
L66:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v172 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[5]))
	F_flagTransaction(m, l0)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L24
	} else {
		goto L70
	}
L68:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+57)))
	if v175&int32(128) != 0 {
		goto L62
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v182 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v184 == v182 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	F_moduleFireCommandACLRejectedEvent(m, l0, int64(0), int32(-1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L24
	} else {
		goto L77
	}
L72:
	;
	F_addReplyErrorObject(m, l0, v179)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L24
	} else {
		goto L76
	}
L73:
	;
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v184)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v184)+120)) = v187 + int64(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+48))
	if v191 != int32(17) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v194 = F_objectGetVal(m, v179)
	mBase = m.M
	F_execCommandAbort(m, l0, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L24
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	goto L71
L77:
	;
	v1511 = int32(0)
	goto L1
L78:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+290)))
	if v242&int32(4) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L79:
	;
	if v211 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	F_flagTransaction(m, l0)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v218 == v216 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_moduleFireCommandRejectedEvent(m, l0, v213)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L24
	} else {
		goto L84
	}
L83:
	;
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v218)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v218)+120)) = v221 + int64(1)
	goto L82
L84:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v227 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_addReplyErrorSds(m, l0, v213)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L24
	} else {
		goto L90
	}
L86:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)+48))
	if v230 != int32(17) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	F_execCommandAbort(m, l0, v213)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	F_sdsfree(m, v213)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
	;
	v1511 = int32(0)
	goto L1
L90:
	;
	v1511 = int32(0)
	goto L1
L91:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+58)))
	if v295&int32(16) == int32(0) {
		v322 = v241
		goto L19
	} else {
		goto L108
	}
L92:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)+52))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(int32(0) < v247)&base.B2i32(v247 != v250) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v257 = F_sdsnew(m, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L24
	} else {
		goto L96
	}
L94:
	;
	if int32(0)-v247 <= v250 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v257
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v241)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v260
	v265 = F_sdscatprintf(m, v257, int32(_a_F_processCommand_8), v14+int32(96))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L24
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v265
	F_flagTransaction(m, l0)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v272 == v270 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_moduleFireCommandRejectedEvent(m, l0, v265)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L24
	} else {
		goto L101
	}
L100:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v272)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+120)) = v275 + int64(1)
	goto L99
L101:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v281 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_addReplyErrorSds(m, l0, v265)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L24
	} else {
		goto L107
	}
L103:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v281)+48))
	if v284 != int32(17) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	F_execCommandAbort(m, l0, v265)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	F_sdsfree(m, v265)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	v1511 = int32(0)
	goto L1
L107:
	;
	v1511 = int32(0)
	goto L1
L108:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v241)+48))
	if v300 != int32(302) {
		v311 = v241
		v312 = v300
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v312 != int32(163) {
		v322 = v311
		goto L19
	} else {
		goto L113
	}
L110:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[6]))
	v305 = F_allowProtectedAction(m, v304, l0)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	if v305 == int32(0) {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+48))
	v311 = v309
	v312 = v310
	goto L109
L113:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[7]))
	v317 = F_allowProtectedAction(m, v316, l0)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L24
	} else {
		goto L114
	}
L114:
	;
	if v317 == int32(0) {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v322 = v321
	goto L19
L116:
	;
	v343 = int64(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v344 != 0 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	if v325 == int32(293) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v330 = F_fcallGetCommandFlags(m, l0, v324)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L24
	} else {
		goto L121
	}
L119:
	;
	if v325 != int32(297) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v342 = v330
	goto L116
L122:
	;
	v340 = F_evalGetCommandFlags(m, l0, v324)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L24
	} else {
		goto L127
	}
L123:
	;
	if v325 == int32(292) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	if v325 == int32(290) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	if v325 != int32(294) {
		v342 = v324
		goto L116
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v342 = v340
	goto L116
L128:
	;
	v356 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v356 == int64(-1) {
		v369 = int32(1)
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+48))
	if v348 != int32(17) {
		v353 = v343
		v354 = int64(0)
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v353 = v343
	v354 = int64(0)
	goto L128
L131:
	;
	v351 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v344)+12)))
	v352 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v344)+8)))
	v353 = v352
	v354 = v351
	goto L128
L132:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v370&int32(8) == int32(0) {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	v359 = int32(1)
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v360&v359 != 0 {
		v369 = v359
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v363 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	goto L137
L136:
	;
	v369 = int32(0)
	goto L132
L137:
	;
	v369 = base.B2i32(v365 == int32(1))
	goto L132
L138:
	;
	v506 = F_ACLCheckAllPerm(m, l0, v14+int32(108))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L24
	} else {
		goto L173
	}
L139:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+59)))
	if v376&int32(1) == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v381 = F_sdsempty(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L24
	} else {
		goto L141
	}
L141:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v384
	v389 = F_sdscatprintf(m, v381, int32(_a_F_processCommand_9), v14+int32(64))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L24
	} else {
		goto L142
	}
L142:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+int32(-1)))))
	switch v401 & int32(7) {
	case 0:
		goto L150
	case 1:
		goto L149
	case 2:
		goto L148
	case 3:
		goto L147
	case 4:
		goto L146
	default:
		goto L144
	}
L143:
	;
	F_flagTransaction(m, l0)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L24
	} else {
		goto L162
	}
L144:
	;
	goto L143
L145:
	;
	if v418 == int32(0) {
		goto L144
	} else {
		goto L151
	}
L146:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v389+int32(-17))))
	v418 = v417
	goto L145
L147:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v389+int32(-9))))
	v418 = v414
	goto L145
L148:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389+int32(-5)))))
	v418 = v411
	goto L145
L149:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+int32(-3)))))
	v418 = v408
	goto L145
L150:
	;
	v418 = int32(base.Ui32(v401) >> (uint(int32(3)) % 32))
	goto L145
L151:
	;
	v428 = int32(0)
	goto L152
L152:
	;
	goto L155
L153:
	;
	goto L144
L154:
	;
	v466 = v428 + int32(1)
	if v466 != v418 {
		v428 = v466
		goto L152
	} else {
		goto L161
	}
L155:
	;
	v434 = v389 + v428
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v442 = int32(0)
	goto L156
L156:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+uint32(_c_F_processCommand[8]))))
	if v435&int32(255) != v448 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L154
L158:
	;
	v454 = v442 + int32(1)
	if v454 != int32(2) {
		v442 = v454
		goto L156
	} else {
		goto L160
	}
L159:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+uint32(_c_F_processCommand[9]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v451)
	goto L154
L160:
	;
	goto L157
L161:
	;
	goto L153
L162:
	;
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v479
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v481 == v479 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	F_moduleFireCommandRejectedEvent(m, l0, v389)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L24
	} else {
		goto L165
	}
L164:
	;
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v481)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v481)+120)) = v484 + int64(1)
	goto L163
L165:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v490 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_addReplyErrorSds(m, l0, v389)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L24
	} else {
		goto L171
	}
L167:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v490)+48))
	if v493 != int32(17) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	F_execCommandAbort(m, l0, v389)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L24
	} else {
		goto L169
	}
L169:
	;
	F_sdsfree(m, v389)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L24
	} else {
		goto L170
	}
L170:
	;
	v1511 = int32(0)
	goto L1
L171:
	;
	v1511 = int32(0)
	goto L1
L172:
	;
	v554 = v353 | v342
	v555 = base.I32_wrap_i64(v554)
	v556 = int32(1)
	v557 = v555 & v556
	v558 = int32(0)
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[10]))
	if (base.B2i32(v559 == v558)|v369)&v556 != 0 {
		v585 = v559
		goto L189
	} else {
		goto L190
	}
L173:
	;
	if v506 == int32(0) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v510 = int32(0)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v512 = int32(2)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	F_addACLLogEntry(m, l0, v506, int32(base.Ui32(v511)>>(uint(v512)%32))&v512, v516, v510, v510)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L24
	} else {
		goto L175
	}
L175:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v523+v524<<(uint(int32(2))%32))))
	v529 = F_objectGetVal(m, v528)
	mBase = m.M
	v531 = F_getAclErrorMessage(m, v506, v521, v522, v529, int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L24
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v531
	F_rejectCommandFormat(m, l0, int32(0), int32(_a_F_processCommand_10), v14+int32(48))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L24
	} else {
		goto L177
	}
L177:
	;
	F_sdsfree(m, v531)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L24
	} else {
		goto L178
	}
L178:
	;
	if base.Ui32(int32(5)) < base.Ui32(v506) {
		v550 = int64(1)
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v14)+108))
	F_moduleFireCommandACLRejectedEvent(m, l0, v550, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L24
	} else {
		goto L181
	}
L180:
	;
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v506<<(uint(int32(3))%32))+uint32(_c_F_processCommand[11])))
	v550 = v549
	goto L179
L181:
	;
	v1511 = v510
	goto L1
L182:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[12]))
	if v823 != 0 {
		goto L266
	} else {
		goto L267
	}
L183:
	;
	if v557 != 0 {
		goto L262
	} else {
		goto L263
	}
L184:
	;
	if v369 == int32(0) {
		goto L4
	} else {
		goto L255
	}
L185:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v760 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L186:
	;
	v722 = int32(0)
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[13]))
	if v724 != int32(2) {
		goto L239
	} else {
		goto L240
	}
L187:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+48))
	if v700 != int32(17) {
		goto L233
	} else {
		goto L234
	}
L188:
	;
	F_evictClients(m)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L24
	} else {
		goto L205
	}
L189:
	;
	if v585 != 0 {
		goto L188
	} else {
		goto L199
	}
L190:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+59)))
	if v566&int32(2) != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v575 = F_getNodeByQuery(m, l0, v14+int32(104))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L24
	} else {
		goto L195
	}
L192:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v565)+80))
	if v569 != 0 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v565)+48))
	if v570 != int32(17) {
		goto L188
	} else {
		goto L194
	}
L194:
	;
	goto L191
L195:
	;
	if v575 == int32(0) {
		goto L187
	} else {
		goto L196
	}
L196:
	;
	v579 = F_getMyClusterNode(m)
	mBase = m.M
	goto L197
L197:
	;
	if base.B2i32(v575 == v579) == int32(0) {
		goto L187
	} else {
		goto L198
	}
L198:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[10]))
	v585 = v584
	goto L189
L199:
	;
	v586 = int32(0)
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[12]))
	if v587 == v586 {
		goto L188
	} else {
		goto L200
	}
L200:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
	v591 = int32(1)
	if (base.B2i32(v590&v591 == int32(0))|v369)&v591 != 0 {
		goto L188
	} else {
		goto L201
	}
L201:
	;
	if v557 != 0 {
		goto L186
	} else {
		goto L202
	}
L202:
	;
	if v555&int32(2) == int32(0) {
		goto L188
	} else {
		goto L203
	}
L203:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v602&int32(2) == int32(0) {
		goto L186
	} else {
		goto L204
	}
L204:
	;
	goto L188
L205:
	;
	v610 = int32(-1)
	v611 = int32(0)
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[14]))
	if v612 == v611 {
		v1511 = v610
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v616 = *(*int64)(unsafe.Add(mBase, _c_F_processCommand[15]))
	if v616 == int64(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v656 = int32(0)
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[16]))
	if v657 == v656 {
		goto L219
	} else {
		goto L220
	}
L208:
	;
	v619 = int32(0)
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[0]))
	if v623 == v619 {
		v631 = v619
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[4]))
	if v631|v633 != 0 {
		goto L207
	} else {
		goto L212
	}
L210:
	;
	goto L209
L211:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v623)+16))
	v631 = int32(base.Ui32(v626)>>(uint(int32(3))%32)) & int32(1)
	goto L210
L212:
	;
	v635 = F_performEvictions(m)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L24
	} else {
		goto L213
	}
L213:
	;
	F_trackingHandlePendingKeyInvalidations(m)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L24
	} else {
		goto L214
	}
L214:
	;
	v639 = int32(0)
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[14]))
	if v640 == v639 {
		v1511 = v610
		goto L1
	} else {
		goto L215
	}
L215:
	;
	if v635 != int32(2) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_processCommand[17])) = base.B2i32(v635 == int32(2))
	goto L207
L217:
	;
	if base.B2i32(v554&int64(4) == int64(0)) == int32(0) {
		goto L185
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v662 = int32(0)
	v663 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[18]))
	if v663 == v662 {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	F_trackingLimitUsedSlots(m)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	if v557 != 0 {
		goto L184
	} else {
		goto L230
	}
L223:
	;
	v676 = int32(0)
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[19]))
	if v677 == v676 {
		goto L183
	} else {
		goto L227
	}
L224:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[20]))
	if v667 < int32(1) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v672 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[21]))
	if v672 == int32(-1) {
		v693 = int32(2)
		goto L222
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	v680 = int32(1)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[22]))
	if v682 == int32(-1) {
		v693 = v680
		goto L222
	} else {
		goto L228
	}
L228:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[23]))
	if v686 != int32(-1) {
		goto L183
	} else {
		goto L229
	}
L229:
	;
	v689 = int32(0)
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_processCommand[25])) = v691
	v693 = v680
	goto L222
L230:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+48))
	if v695 == int32(183) {
		goto L184
	} else {
		goto L231
	}
L231:
	;
	v823 = int32(1)
	goto L182
L232:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	F_clusterRedirectClient(m, l0, v575, v707, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L24
	} else {
		goto L237
	}
L233:
	;
	F_flagTransaction(m, l0)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L24
	} else {
		goto L236
	}
L234:
	;
	F_discardTransaction(m, l0)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L24
	} else {
		goto L235
	}
L235:
	;
	goto L232
L236:
	;
	goto L232
L237:
	;
	v711 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v711
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v715 = *(*int64)(unsafe.Add(mBase, uint32(v714)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v714)+120)) = v715 + int64(1)
	F_moduleFireCommandRejectedEvent(m, l0, v711)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L24
	} else {
		goto L238
	}
L238:
	;
	v1511 = v711
	goto L1
L239:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)+48))
	if v730 != int32(17) {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	F_blockPostponeClient(m, l0)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L24
	} else {
		goto L241
	}
L241:
	;
	v1511 = v722
	goto L1
L242:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v737
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v741 = *(*int64)(unsafe.Add(mBase, uint32(v740)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v740)+120)) = v741 + int64(1)
	F_moduleFireCommandRejectedEvent(m, l0, int32(_a_F_processCommand_11))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L24
	} else {
		goto L247
	}
L243:
	;
	F_flagTransaction(m, l0)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L24
	} else {
		goto L246
	}
L244:
	;
	F_discardTransaction(m, l0)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L24
	} else {
		goto L245
	}
L245:
	;
	goto L242
L246:
	;
	goto L242
L247:
	;
	v748 = F_sdsempty(m)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L24
	} else {
		goto L248
	}
L248:
	;
	v751 = *(*int64)(unsafe.Add(mBase, _c_F_processCommand[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v751
	v756 = F_sdscatprintf(m, v748, int32(_a_F_processCommand_12), v14+int32(32))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L24
	} else {
		goto L249
	}
L249:
	;
	F_addReplyErrorSds(m, l0, v756)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L24
	} else {
		goto L250
	}
L250:
	;
	v1511 = v737
	goto L1
L251:
	;
	v765 = int32(0)
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[26]))
	F_rejectCommand(m, l0, v767, int32(1))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L24
	} else {
		goto L254
	}
L252:
	;
	F_clusterHandleSlotMigrationClientOOM(m, v760)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L24
	} else {
		goto L253
	}
L253:
	;
	v1511 = v610
	goto L1
L254:
	;
	v1511 = v765
	goto L1
L255:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[27]))
	if v774 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v779 = int32(0)
	v780 = *(*int64)(unsafe.Add(mBase, _c_F_processCommand[28]))
	v782 = *(*int64)(unsafe.Add(mBase, _c_F_processCommand[29]))
	if v780 <= v782+int64(10000) {
		goto L183
	} else {
		goto L259
	}
L257:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)+48))
	if v776 != int32(183) {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v786 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_processCommand[29])) = v780
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[30]))
	if int32(3) < v789 {
		goto L183
	} else {
		goto L260
	}
L260:
	;
	F__serverLog(m, int32(3), int32(_a_F_processCommand_13), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L24
	} else {
		goto L261
	}
L261:
	;
	goto L183
L262:
	;
	v800 = int32(0)
	v801 = int32(_a_F_processCommand_14)
	v802 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[31]))
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[32]))
	v811 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[12]))
	v816 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[33]))
	goto L264
L263:
	;
	v823 = int32(1)
	goto L182
L264:
	;
	if base.B2i32(v802 == v800)|base.B2i32(v806 == v800)|base.B2i32(v811 != v800)|base.B2i32(v806 <= v816) == int32(0) {
		goto L3
	} else {
		goto L265
	}
L265:
	;
	v823 = int32(0)
	goto L182
L266:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v843&int32(4) == int32(0) {
		goto L270
	} else {
		goto L271
	}
L267:
	;
	v827 = int32(0)
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[34]))
	if (base.B2i32(v826 == v827)|base.B2i32(v830 == v827)|v369)&int32(1) != 0 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v837 = int32(0)
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[35]))
	F_rejectCommand(m, l0, v839, int32(1))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L24
	} else {
		goto L269
	}
L269:
	;
	v1511 = v837
	goto L1
L270:
	;
	v974 = v354 | (v342 ^ int64(-1))
	if v826 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L271:
	;
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v848 != int32(2) {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v851)+48))
	if v852 == int32(185) {
		goto L270
	} else {
		goto L273
	}
L273:
	;
	if v852 == int32(184) {
		goto L270
	} else {
		goto L274
	}
L274:
	;
	if v852 == int32(284) {
		goto L270
	} else {
		goto L275
	}
L275:
	;
	if v852 == int32(282) {
		goto L270
	} else {
		goto L276
	}
L276:
	;
	if v852 == int32(288) {
		goto L270
	} else {
		goto L277
	}
L277:
	;
	if v852 == int32(289) {
		goto L270
	} else {
		goto L278
	}
L278:
	;
	if v852 == int32(286) {
		goto L270
	} else {
		goto L279
	}
L279:
	;
	if v852 == int32(183) {
		goto L270
	} else {
		goto L280
	}
L280:
	;
	if v852 == int32(287) {
		goto L270
	} else {
		goto L281
	}
L281:
	;
	v871 = F_sdsempty(m)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L24
	} else {
		goto L282
	}
L282:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v874
	v879 = F_sdscatprintf(m, v871, int32(_a_F_processCommand_15), v14+int32(16))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L24
	} else {
		goto L283
	}
L283:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879+int32(-1)))))
	switch v891 & int32(7) {
	case 0:
		goto L291
	case 1:
		goto L290
	case 2:
		goto L289
	case 3:
		goto L288
	case 4:
		goto L287
	default:
		goto L285
	}
L284:
	;
	F_rejectCommandSds(m, l0, v879, int32(1))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L24
	} else {
		goto L303
	}
L285:
	;
	goto L284
L286:
	;
	if v908 == int32(0) {
		goto L285
	} else {
		goto L292
	}
L287:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v879+int32(-17))))
	v908 = v907
	goto L286
L288:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v879+int32(-9))))
	v908 = v904
	goto L286
L289:
	;
	v901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v879+int32(-5)))))
	v908 = v901
	goto L286
L290:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879+int32(-3)))))
	v908 = v898
	goto L286
L291:
	;
	v908 = int32(base.Ui32(v891) >> (uint(int32(3)) % 32))
	goto L286
L292:
	;
	v918 = int32(0)
	goto L293
L293:
	;
	goto L296
L294:
	;
	goto L285
L295:
	;
	v956 = v918 + int32(1)
	if v956 != v908 {
		v918 = v956
		goto L293
	} else {
		goto L302
	}
L296:
	;
	v924 = v879 + v918
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924))))
	v932 = int32(0)
	goto L297
L297:
	;
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+uint32(_c_F_processCommand[8]))))
	if v925&int32(255) != v938 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	goto L295
L299:
	;
	v944 = v932 + int32(1)
	if v944 != int32(2) {
		v932 = v944
		goto L297
	} else {
		goto L301
	}
L300:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+uint32(_c_F_processCommand[9]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v924))) = uint8(v941)
	goto L295
L301:
	;
	goto L298
L302:
	;
	goto L294
L303:
	;
	v1511 = int32(0)
	goto L1
L304:
	;
	v993 = int32(0)
	v994 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[36]))
	if v994 == v993 {
		goto L310
	} else {
		goto L311
	}
L305:
	;
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[37]))
	if v978 == int32(14) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v982 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[38]))
	if v982 != 0 {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	if v974&int64(1024) == int64(0) {
		goto L304
	} else {
		goto L308
	}
L308:
	;
	v987 = int32(0)
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[39]))
	F_rejectCommand(m, l0, v989, int32(1))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L24
	} else {
		goto L309
	}
L309:
	;
	v1511 = v987
	goto L1
L310:
	;
	v1009 = int32(0)
	v1010 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[40]))
	if v1010 == v1009 {
		goto L315
	} else {
		goto L316
	}
L311:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[40]))
	if v998 != 0 {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	if v974&int64(512) == int64(0) {
		goto L310
	} else {
		goto L313
	}
L313:
	;
	v1003 = int32(0)
	v1005 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[41]))
	F_rejectCommand(m, l0, v1005, int32(1))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L24
	} else {
		goto L314
	}
L314:
	;
	v1511 = v1003
	goto L1
L315:
	;
	v1023 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[0]))
	if v1027 == v1023 {
		v1035 = v1023
		goto L321
	} else {
		goto L322
	}
L316:
	;
	if v554&int64(8388608) == int64(0) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1017 = int32(0)
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[41]))
	F_rejectCommand(m, l0, v1019, int32(1))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L24
	} else {
		goto L318
	}
L318:
	;
	v1511 = v1017
	goto L1
L319:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v1173&int32(2) == int32(0) {
		goto L359
	} else {
		goto L360
	}
L320:
	;
	v1036 = int32(0)
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[4]))
	if v1035|v1037 == v1036 {
		goto L319
	} else {
		goto L323
	}
L321:
	;
	goto L320
L322:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+16))
	v1035 = int32(base.Ui32(v1030)>>(uint(int32(3))%32)) & int32(1)
	goto L321
L323:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041)+59)))
	if v1042&int32(4) != 0 {
		goto L319
	} else {
		goto L324
	}
L324:
	;
	if v1037 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	if v1037 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L326:
	;
	v1047 = int32(0)
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[42]))
	if v1048 == v1047 {
		goto L325
	} else {
		goto L327
	}
L327:
	;
	v1052 = F_sdsempty(m)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L24
	} else {
		goto L328
	}
L328:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v1055
	v1058 = F_sdscatprintf(m, v1052, int32(_a_F_processCommand_16), v14)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L24
	} else {
		goto L329
	}
L329:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058+int32(-1)))))
	switch v1070 & int32(7) {
	case 0:
		goto L337
	case 1:
		goto L336
	case 2:
		goto L335
	case 3:
		goto L334
	case 4:
		goto L333
	default:
		goto L331
	}
L330:
	;
	F_rejectCommandSds(m, l0, v1058, int32(1))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L24
	} else {
		goto L349
	}
L331:
	;
	goto L330
L332:
	;
	if v1087 == int32(0) {
		goto L331
	} else {
		goto L338
	}
L333:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1058+int32(-17))))
	v1087 = v1086
	goto L332
L334:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1058+int32(-9))))
	v1087 = v1083
	goto L332
L335:
	;
	v1080 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058+int32(-5)))))
	v1087 = v1080
	goto L332
L336:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058+int32(-3)))))
	v1087 = v1077
	goto L332
L337:
	;
	v1087 = int32(base.Ui32(v1070) >> (uint(int32(3)) % 32))
	goto L332
L338:
	;
	v1097 = int32(0)
	goto L339
L339:
	;
	goto L342
L340:
	;
	goto L331
L341:
	;
	v1135 = v1097 + int32(1)
	if v1135 != v1087 {
		v1097 = v1135
		goto L339
	} else {
		goto L348
	}
L342:
	;
	v1103 = v1058 + v1097
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103))))
	v1111 = int32(0)
	goto L343
L343:
	;
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111)+uint32(_c_F_processCommand[8]))))
	if v1104&int32(255) != v1117 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	goto L341
L345:
	;
	v1123 = v1111 + int32(1)
	if v1123 != int32(2) {
		v1111 = v1123
		goto L343
	} else {
		goto L347
	}
L346:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111)+uint32(_c_F_processCommand[9]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1103))) = uint8(v1120)
	goto L341
L347:
	;
	goto L344
L348:
	;
	goto L340
L349:
	;
	v1511 = int32(0)
	goto L1
L350:
	;
	v1157 = F_scriptIsEval(m)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L24
	} else {
		goto L354
	}
L351:
	;
	v1151 = int32(0)
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[43]))
	F_rejectCommand(m, l0, v1153, int32(1))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L24
	} else {
		goto L352
	}
L352:
	;
	v1511 = v1151
	goto L1
L353:
	;
	v1167 = int32(0)
	v1169 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[44]))
	F_rejectCommand(m, l0, v1169, int32(1))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L24
	} else {
		goto L357
	}
L354:
	;
	if v1157 == int32(0) {
		goto L353
	} else {
		goto L355
	}
L355:
	;
	v1161 = int32(0)
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[45]))
	F_rejectCommand(m, l0, v1163, int32(1))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L24
	} else {
		goto L356
	}
L356:
	;
	v1511 = v1161
	goto L1
L357:
	;
	v1511 = v1167
	goto L1
L358:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v1210&int32(8) == int32(0) {
		goto L373
	} else {
		goto L374
	}
L359:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v1189 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L360:
	;
	if v555&int32(65539) == int32(0) {
		goto L358
	} else {
		goto L361
	}
L361:
	;
	v1183 = F_sdsnew(m, int32(_a_F_processCommand_17))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L24
	} else {
		goto L362
	}
L362:
	;
	F_rejectCommandSds(m, l0, v1183, int32(1))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L24
	} else {
		goto L363
	}
L363:
	;
	v1511 = int32(0)
	goto L1
L364:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[46]))
	goto L368
L365:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	goto L366
L366:
	;
	if base.B2i32(v1192 == int32(1)) == int32(0) {
		goto L358
	} else {
		goto L367
	}
L367:
	;
	goto L364
L368:
	;
	if v1199&int32(2) != 0 {
		goto L2
	} else {
		goto L369
	}
L369:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[46]))
	goto L370
L370:
	;
	if v1203&int32(1) == int32(0) {
		goto L358
	} else {
		goto L371
	}
L371:
	;
	if v555&int32(65537) != 0 {
		goto L2
	} else {
		goto L372
	}
L372:
	;
	goto L358
L373:
	;
	F_call(m, l0, int32(3))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L24
	} else {
		goto L381
	}
L374:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1215)+48))
	if v1216 == int32(185) {
		goto L373
	} else {
		goto L375
	}
L375:
	;
	if v1216 == int32(184) {
		goto L373
	} else {
		goto L376
	}
L376:
	;
	if v1216 == int32(17) {
		goto L373
	} else {
		goto L377
	}
L377:
	;
	if v1216 == int32(412) {
		goto L373
	} else {
		goto L378
	}
L378:
	;
	F_queueMultiCommand(m, l0, v342)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L24
	} else {
		goto L379
	}
L379:
	;
	v1227 = int32(0)
	v1229 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[47]))
	F_addReply(m, l0, v1229)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L24
	} else {
		goto L380
	}
L380:
	;
	v1511 = v1227
	goto L1
L381:
	;
	v1236 = int32(0)
	v1238 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[48]))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+20))
	if v1239 == v1236 {
		v1511 = v1236
		goto L1
	} else {
		goto L382
	}
L382:
	;
	v1242 = int32(0)
	v1247 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[0]))
	if v1247 == v1242 {
		v1255 = v1242
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, _c_F_processCommand[4]))
	if v1255|v1257 != 0 {
		v1511 = v1242
		goto L1
	} else {
		goto L386
	}
L384:
	;
	goto L383
L385:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+16))
	v1255 = int32(base.Ui32(v1250)>>(uint(int32(3))%32)) & int32(1)
	goto L384
L386:
	;
	F_handleClientsBlockedOnKeys(m)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L24
	} else {
		goto L387
	}
L387:
	;
	v1511 = v1242
	goto L1
L388:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L390:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L391:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+48))
	v1294 = base.B2i32(v1292 == int32(302))
	if v1292 == int32(302) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1295 = int32(_a_F_processCommand_18)
	goto L395
L394:
	;
	v1295 = int32(_a_F_processCommand_19)
	goto L395
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v1295
	if v1292 == int32(302) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1299 = int32(_a_F_processCommand_20)
	goto L398
L397:
	;
	v1299 = int32(_a_F_processCommand_21)
	goto L398
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v1299
	v1304 = F_sdscatprintf(m, v1287, int32(_a_F_processCommand_22), v14+int32(80))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L24
	} else {
		goto L399
	}
L399:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304+int32(-1)))))
	switch v1316 & int32(7) {
	case 0:
		goto L407
	case 1:
		goto L406
	case 2:
		goto L405
	case 3:
		goto L404
	case 4:
		goto L403
	default:
		goto L401
	}
L400:
	;
	F_rejectCommandSds(m, l0, v1304, int32(1))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L24
	} else {
		goto L419
	}
L401:
	;
	goto L400
L402:
	;
	if v1333 == int32(0) {
		goto L401
	} else {
		goto L408
	}
L403:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1304+int32(-17))))
	v1333 = v1332
	goto L402
L404:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1304+int32(-9))))
	v1333 = v1329
	goto L402
L405:
	;
	v1326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304+int32(-5)))))
	v1333 = v1326
	goto L402
L406:
	;
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304+int32(-3)))))
	v1333 = v1323
	goto L402
L407:
	;
	v1333 = int32(base.Ui32(v1316) >> (uint(int32(3)) % 32))
	goto L402
L408:
	;
	v1343 = int32(0)
	goto L409
L409:
	;
	goto L412
L410:
	;
	goto L401
L411:
	;
	v1381 = v1343 + int32(1)
	if v1381 != v1333 {
		v1343 = v1381
		goto L409
	} else {
		goto L418
	}
L412:
	;
	v1349 = v1304 + v1343
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349))))
	v1357 = int32(0)
	goto L413
L413:
	;
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357)+uint32(_c_F_processCommand[8]))))
	if v1350&int32(255) != v1363 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	goto L411
L415:
	;
	v1369 = v1357 + int32(1)
	if v1369 != int32(2) {
		v1357 = v1369
		goto L413
	} else {
		goto L417
	}
L416:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357)+uint32(_c_F_processCommand[9]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1349))) = uint8(v1366)
	goto L411
L417:
	;
	goto L414
L418:
	;
	goto L410
L419:
	;
	v1511 = int32(0)
	goto L1
L420:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L421:
	;
	v1405 = int32(0)
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403+int32(-1)))))
	switch v1412 & int32(7) {
	case 0:
		goto L428
	case 1:
		goto L427
	case 2:
		goto L426
	case 3:
		goto L425
	case 4:
		goto L424
	default:
		v1429 = v1405
		goto L423
	}
L422:
	;
	v1433 = v1431 + int32(-2)
	v1440 = v1403 + int32(-1)
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1440))))
	v1443 = v1441 & int32(7)
	switch v1443 {
	case 0:
		goto L435
	case 1:
		goto L434
	case 2:
		goto L433
	case 3:
		goto L432
	case 4:
		goto L431
	default:
		v1458 = int32(0)
		goto L430
	}
L423:
	;
	v1431 = v1429
	goto L422
L424:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1403+int32(-17))))
	v1429 = v1428
	goto L423
L425:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1403+int32(-9))))
	v1431 = v1425
	goto L422
L426:
	;
	v1422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1403+int32(-5)))))
	v1431 = v1422
	goto L422
L427:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403+int32(-3)))))
	v1431 = v1419
	goto L422
L428:
	;
	v1431 = int32(base.Ui32(v1412) >> (uint(int32(3)) % 32))
	goto L422
L429:
	;
	F_rejectCommandSds(m, l0, v1403, int32(1))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L24
	} else {
		goto L453
	}
L430:
	;
	v1460 = base.B2i32(base.Ui32(v1405) < base.Ui32(v1458))
	if base.Ui32(v1405) < base.Ui32(v1458) {
		goto L437
	} else {
		goto L438
	}
L431:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1403+int32(-17))))
	v1458 = v1457
	goto L430
L432:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1403+int32(-9))))
	v1458 = v1454
	goto L430
L433:
	;
	v1451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1403+int32(-5)))))
	v1458 = v1451
	goto L430
L434:
	;
	v1448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403+int32(-3)))))
	v1458 = v1448
	goto L430
L435:
	;
	v1458 = int32(base.Ui32(v1441) >> (uint(int32(3)) % 32))
	goto L430
L436:
	;
	v1472 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1403+v1466))) = uint8(v1472)
	switch v1443 {
	case 0:
		goto L452
	case 1:
		goto L451
	case 2:
		goto L450
	case 3:
		goto L449
	case 4:
		goto L448
	default:
		goto L447
	}
L437:
	;
	v1461 = v1405
	goto L439
L438:
	;
	v1461 = int32(0)
	goto L439
L439:
	;
	v1462 = v1458 - v1461
	if base.Ui32(v1433) < base.Ui32(v1462) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1464 = v1433
	goto L442
L441:
	;
	v1464 = v1462
	goto L442
L442:
	;
	if base.Ui32(v1405) < base.Ui32(v1458) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1466 = v1464
	goto L445
L444:
	;
	v1466 = int32(0)
	goto L445
L445:
	;
	if v1466 == int32(0) {
		goto L436
	} else {
		goto L446
	}
L446:
	;
	v1470 = F_memmove(m, v1403, v1403+v1461, v1466)
	mBase = m.M
	goto L436
L447:
	;
	goto L429
L448:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1403+int32(-17)))) = base.I64_extend_i32_u(v1466)
	goto L447
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403+int32(-9)))) = v1466
	goto L429
L450:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1403+int32(-5)))) = uint16(v1466)
	goto L429
L451:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1403+int32(-3)))) = uint8(v1466)
	goto L429
L452:
	;
	v1475 = v1466 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1440))) = uint8(v1475)
	goto L429
L453:
	;
	v1511 = v1405
	goto L1
L454:
	;
	v1511 = v1493
	goto L1
L455:
	;
	v1511 = int32(0)
	goto L1
}
func F_processCommandAndResetClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = int32(_a_F_processCommandAndResetClient_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0])) = l0
	v7 = F_processCommand(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v18 = int32(_a_F_processCommandAndResetClient_0)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0])) = v4
			if v19 != 0 {
				v24 = int32(0)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			F_commandProcessed(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v13 == int32(0) {
					v18 = int32(_a_F_processCommandAndResetClient_0)
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0])) = v4
					if v19 != 0 {
						v24 = int32(0)
					} else {
						v24 = int32(-1)
					}
					return v24
				} else {
					v16 = F_updateClientMemUsageAndBucket(m, l0)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						v18 = int32(_a_F_processCommandAndResetClient_0)
						v19 = *(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_processCommandAndResetClient[0])) = v4
						if v19 != 0 {
							v24 = int32(0)
						} else {
							v24 = int32(-1)
						}
						return v24
					}
				}
			}
		}
	}
}
func F_resetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v5&int32(4) != 0 {
		v8 = int32(1073741825)
	} else {
		v8 = int32(1073741827)
	}
	if v8&v5 == int32(0) {
		F_clearClientConnectionState(m, l0)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_addReplyStatusLength(m, l0, int32(_a_F_resetCommand_0), int32(5))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_addReplyErrorLength(m, l0, int32(_a_F_resetCommand_1), int32(40))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_afterErrorReply(m, l0, int32(_a_F_resetCommand_1), int32(40), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_resetCommandTableStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(64)
	m.G0 = v6
	v9 = v6 + int32(16)
	v10 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(-1)
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v9
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v6+int32(16))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L19
	}
L6:
	;
	return
L7:
	;
	if v32 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+104)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(128)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(120)))) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(112)))) = v40
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+148))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+200))
	if v61 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_hdr_close(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+148)) = int32(0)
	goto L11
L14:
	;
	v70 = F_hashtableNext(m, v6+int32(16), v6+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	F_resetCommandTableStats(m, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v70 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	m.G0 = v6 + int32(64)
	return
}
func F_scanCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v84 int64
	_ = v84
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int64
	_ = v106
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v127 int64
	_ = v127
	var v138 int64
	_ = v138
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int64
	_ = v184
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = F_objectGetVal(m, v13)
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v17 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v34 = int32(0)
		goto L1
	}
L1:
	;
	v42 = m.G0
	v44 = v42 - int32(16)
	m.G0 = v44
	if base.Ui32(v34+int32(-21)) < base.Ui32(int32(-20)) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v34 = v33
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v34 = v30
	goto L1
L4:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v34 = v27
	goto L1
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v34 = v24
	goto L1
L6:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	m.G0 = v9 + int32(48)
	return
L8:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v207 = int32(0)
	v212 = F_parseScanOptionsOrReply(m, l0, v207, int32(2), v207, v9+int32(8))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L37
	} else {
		goto L39
	}
L9:
	;
	if v191 != 0 {
		goto L8
	} else {
		goto L36
	}
L10:
	;
	m.G0 = v44 + int32(16)
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v184
	v191 = int32(1)
	goto L10
L12:
	;
	v155 = int32(0)
	v156 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v155
	v164 = F_strtoull(m, v14, v44+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v166 == int32(28) {
		v191 = v155
		goto L10
	} else {
		goto L33
	}
L13:
	;
	v50 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v34 != v50 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v51&int32(255) != int32(45) {
		v71 = v50
		v72 = v51
		v73 = v14
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v55 = v51 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v55&int32(255)) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v184 = base.I64_extend_i32_u(v55) & int64(255)
	goto L11
L17:
	;
	if base.Ui32(int32(8)) < base.Ui32((v72+int32(-49))&int32(255)) {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v71 = int32(2)
	v72 = v69
	v73 = v14 + int32(1)
	goto L17
L19:
	;
	v84 = base.I64_extend_i32_u(v72+int32(-48)) & int64(255)
	if base.Ui32(v34) <= base.Ui32(v71) {
		v127 = v84
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v51&int32(255) != int32(45) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v90 = v71
	v92 = v84
	v94 = v73
	goto L22
L22:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if base.Ui32((v96+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L12
	} else {
		goto L24
	}
L23:
	;
	v127 = v117
	goto L20
L24:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v92) {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v106 = v92 * int64(10)
	v111 = base.I64_extend_i32_u(v96+int32(-48)) & int64(255)
	if base.Ui64(v111^int64(-1)) < base.Ui64(v106) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v115 = int32(1)
	v117 = v106 + v111
	v119 = v90 + v115
	if v119 != v34 {
		v90 = v119
		v92 = v117
		v94 = v94 + v115
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	if int64(0) <= v127 {
		v184 = v127
		goto L11
	} else {
		goto L32
	}
L29:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v127) {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v138 = int64(-1)
	if v138 < v127+v138 {
		v191 = int32(0)
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v184 = int64(0)
	goto L11
L32:
	;
	goto L12
L33:
	;
	if v166 == int32(68) {
		v191 = v155
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v171 == int32(0) {
		v191 = v155
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v191 = base.B2i32(v175 == int32(0))
	goto L10
L36:
	;
	F_addReplyError(m, l0, int32(_a_F_scanCommand_0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return
L38:
	;
	goto L7
L39:
	;
	if v212 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v214 = int32(0)
	F_scanGenericCommandWithOptions(m, l0, v214, v206, v9+int32(8), v214)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L7
}
func F_sortCommandGeneric(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
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
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
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
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
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
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
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
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v970 int32
	_ = v970
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1051 int32
	_ = v1051
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1126 int32
	_ = v1126
	var v1134 int32
	_ = v1134
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var __phi1238 int32
	_ = __phi1238
	var v1242 int32
	_ = v1242
	var __phi1242 int32
	_ = __phi1242
	var v1245 int32
	_ = v1245
	var __phi1245 int32
	_ = __phi1245
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1509 int32
	_ = v1509
	var v1530 int32
	_ = v1530
	var v1536 int32
	_ = v1536
	var v1559 int32
	_ = v1559
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1601 int32
	_ = v1601
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int64
	_ = v1679
	var v1684 int64
	_ = v1684
	var v1692 int32
	_ = v1692
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1705 float64
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1728 int32
	_ = v1728
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1771 int32
	_ = v1771
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1821 int32
	_ = v1821
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1858 int32
	_ = v1858
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1890 int32
	_ = v1890
	var v1902 int32
	_ = v1902
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1951 int32
	_ = v1951
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1994 int32
	_ = v1994
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2037 int32
	_ = v2037
	var v2049 int32
	_ = v2049
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2096 int32
	_ = v2096
	var v2148 int32
	_ = v2148
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2183 int64
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int64
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2224 int32
	_ = v2224
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2289 int32
	_ = v2289
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2326 int32
	_ = v2326
	var v2357 int32
	_ = v2357
	var v2364 int32
	_ = v2364
	var v2370 int32
	_ = v2370
	var v2376 int32
	_ = v2376
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(80)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = int32(-1)
	v33 = F_listCreate(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = int32(102)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v45 = v43 & int32(8)
	if v45 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v46 = int32(112)
	goto L5
L4:
	;
	v46 = int32(96)
	goto L5
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+v46)))
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = int32(52)
	goto L8
L7:
	;
	v51 = int32(28)
	goto L8
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48+v51)))
	v55 = F_ACLUserCheckCmdWithUnrestrictedKeyAccess(m, v37, v38, v39, v40, v53, int32(16))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(3) <= v57 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	F__serverAssertWithInfo(m, l0, v838, int32(_a_F_sortCommandGeneric_0), int32(_a_F_sortCommandGeneric_1), int32(581))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L556
	}
L11:
	;
	F__serverAssertWithInfo(m, l0, v838, int32(_a_F_sortCommandGeneric_2), int32(_a_F_sortCommandGeneric_1), int32(466))
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L1
	} else {
		goto L555
	}
L12:
	;
	F__serverPanic_1(m, int32(_a_F_sortCommandGeneric_1), int32(464), int32(_a_F_sortCommandGeneric_3), int32(0))
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L1
	} else {
		goto L554
	}
L13:
	;
	F__serverAssertWithInfo(m, l0, v838, int32(_a_F_sortCommandGeneric_4), int32(_a_F_sortCommandGeneric_1), int32(438))
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L1
	} else {
		goto L553
	}
L14:
	;
	m.G0 = v26 + int32(80)
	return
L15:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v841 = v839 & int32(15)
	if v781|base.B2i32(v841 != int32(2)) != 0 {
		goto L235
	} else {
		goto L236
	}
L16:
	;
	v832 = int32(_a_F_sortCommandGeneric_5)
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[0]))
	v835 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[1]))
	v836 = F_createQuicklistObject(m, v833, v835)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L233
	}
L17:
	;
	F_incrRefCount(m, v802)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L232
	}
L18:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L231
	}
L19:
	;
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[2]))
	F_addReplyErrorObject(m, l0, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L230
	}
L20:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v800)+4))
	v802 = F_lookupKeyRead(m, v799, v801)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L225
	}
L21:
	;
	v65 = int32(0)
	v75 = v65
	v77 = int32(2)
	v79 = v57
	v80 = v65
	v81 = v65
	v82 = v65
	v83 = v65
	v84 = v65
	goto L23
L22:
	;
	v61 = int32(0)
	v779 = v3
	v781 = int32(1)
	v784 = v61
	v785 = v61
	v786 = v61
	v787 = v61
	goto L20
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v97 = v77 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v97)))
	v100 = F_objectGetVal(m, v99)
	mBase = m.M
	v101 = int32(_a_F_sortCommandGeneric_6)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v104 != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v779 = v759
	v781 = base.B2i32(v766 == int32(0))
	v784 = v762
	v785 = v763
	v786 = v764
	v787 = v765
	goto L20
L25:
	;
	v771 = v760 + int32(1)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v771 < v772 {
		v75 = v759
		v77 = v771
		v79 = v772
		v80 = v762
		v81 = v763
		v82 = v764
		v83 = v765
		v84 = v766
		goto L23
	} else {
		goto L224
	}
L26:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141+v97)))
	v144 = F_objectGetVal(m, v143)
	mBase = m.M
	v145 = int32(_a_F_sortCommandGeneric_7)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 != 0 {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	if v136-v138 != 0 {
		goto L26
	} else {
		goto L39
	}
L28:
	;
	v136 = F_tolower(m, v132)
	mBase = m.M
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v138 = F_tolower(m, v137)
	mBase = m.M
	goto L27
L29:
	;
	v106 = v100
	v107 = v101
	v108 = v104
	goto L32
L30:
	;
	v132 = int32(0)
	v133 = v101
	goto L28
L31:
	;
	v132 = v129 & int32(255)
	v133 = v128
	goto L28
L32:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 == int32(0) {
		v128 = v107
		v129 = v108
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v128 = v122
	v129 = int32(0)
	goto L31
L34:
	;
	v114 = v108 & int32(255)
	if v114 == v110 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v121 = int32(1)
	v122 = v107 + v121
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v123 != 0 {
		v106 = v106 + v121
		v107 = v122
		v108 = v123
		goto L32
	} else {
		goto L38
	}
L36:
	;
	v116 = F_tolower(m, v114)
	mBase = m.M
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v118 = F_tolower(m, v117)
	mBase = m.M
	if v116 == v118 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v128 = v107
	v129 = v120
	goto L31
L38:
	;
	goto L33
L39:
	;
	v759 = v75
	v760 = v77
	v762 = v80
	v763 = v81
	v764 = v82
	v765 = int32(0)
	v766 = v84
	goto L25
L40:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185+v97)))
	v188 = F_objectGetVal(m, v187)
	mBase = m.M
	v189 = int32(_a_F_sortCommandGeneric_8)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 != 0 {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	if v180-v182 != 0 {
		goto L40
	} else {
		goto L53
	}
L42:
	;
	v180 = F_tolower(m, v176)
	mBase = m.M
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v182 = F_tolower(m, v181)
	mBase = m.M
	goto L41
L43:
	;
	v150 = v144
	v151 = v145
	v152 = v148
	goto L46
L44:
	;
	v176 = int32(0)
	v177 = v145
	goto L42
L45:
	;
	v176 = v173 & int32(255)
	v177 = v172
	goto L42
L46:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 == int32(0) {
		v172 = v151
		v173 = v152
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v172 = v166
	v173 = int32(0)
	goto L45
L48:
	;
	v158 = v152 & int32(255)
	if v158 == v154 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v165 = int32(1)
	v166 = v151 + v165
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v167 != 0 {
		v150 = v150 + v165
		v151 = v166
		v152 = v167
		goto L46
	} else {
		goto L52
	}
L50:
	;
	v160 = F_tolower(m, v158)
	mBase = m.M
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v162 = F_tolower(m, v161)
	mBase = m.M
	if v160 == v162 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v172 = v151
	v173 = v164
	goto L45
L52:
	;
	goto L47
L53:
	;
	v759 = v75
	v760 = v77
	v762 = v80
	v763 = v81
	v764 = v82
	v765 = int32(1)
	v766 = v84
	goto L25
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229+v97)))
	v232 = F_objectGetVal(m, v231)
	mBase = m.M
	v233 = int32(_a_F_sortCommandGeneric_9)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v236 != 0 {
		goto L70
	} else {
		goto L71
	}
L55:
	;
	if v224-v226 != 0 {
		goto L54
	} else {
		goto L67
	}
L56:
	;
	v224 = F_tolower(m, v220)
	mBase = m.M
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v226 = F_tolower(m, v225)
	mBase = m.M
	goto L55
L57:
	;
	v194 = v188
	v195 = v189
	v196 = v192
	goto L60
L58:
	;
	v220 = int32(0)
	v221 = v189
	goto L56
L59:
	;
	v220 = v217 & int32(255)
	v221 = v216
	goto L56
L60:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v198 == int32(0) {
		v216 = v195
		v217 = v196
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v216 = v210
	v217 = int32(0)
	goto L59
L62:
	;
	v202 = v196 & int32(255)
	if v202 == v198 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v209 = int32(1)
	v210 = v195 + v209
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v211 != 0 {
		v194 = v194 + v209
		v195 = v210
		v196 = v211
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v204 = F_tolower(m, v202)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	if v204 == v206 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v216 = v195
	v217 = v208
	goto L59
L66:
	;
	goto L61
L67:
	;
	v759 = v75
	v760 = v77
	v762 = v80
	v763 = v81
	v764 = int32(1)
	v765 = v83
	v766 = v84
	goto L25
L68:
	;
	v274 = v79 + (v77 ^ int32(-1))
	if v274 < int32(2) {
		goto L80
	} else {
		goto L81
	}
L69:
	;
	v268 = F_tolower(m, v264)
	mBase = m.M
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	v270 = F_tolower(m, v269)
	mBase = m.M
	goto L68
L70:
	;
	v238 = v232
	v239 = v233
	v240 = v236
	goto L73
L71:
	;
	v264 = int32(0)
	v265 = v233
	goto L69
L72:
	;
	v264 = v261 & int32(255)
	v265 = v260
	goto L69
L73:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v242 == int32(0) {
		v260 = v239
		v261 = v240
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v260 = v254
	v261 = int32(0)
	goto L72
L75:
	;
	v246 = v240 & int32(255)
	if v246 == v242 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v253 = int32(1)
	v254 = v239 + v253
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	if v255 != 0 {
		v238 = v238 + v253
		v239 = v254
		v240 = v255
		goto L73
	} else {
		goto L79
	}
L77:
	;
	v248 = F_tolower(m, v246)
	mBase = m.M
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v250 = F_tolower(m, v249)
	mBase = m.M
	if v248 == v250 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v260 = v239
	v261 = v252
	goto L72
L79:
	;
	goto L74
L80:
	;
	if l1 != 0 {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	if v268-v270 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277+v97+int32(4))))
	v285 = F_getLongFromObjectOrReply(m, l0, v281, v26+int32(76), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	if v285 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v288 = int32(2)
	v289 = v77 + v288
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287+v289<<(uint(v288)%32))))
	v297 = F_getLongFromObjectOrReply(m, l0, v293, v26+int32(72), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v297 == int32(0) {
		v759 = v75
		v760 = v289
		v762 = v80
		v763 = v81
		v764 = v82
		v765 = v83
		v766 = v84
		goto L25
	} else {
		goto L86
	}
L86:
	;
	goto L18
L87:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354+v97)))
	v357 = F_objectGetVal(m, v356)
	mBase = m.M
	v358 = int32(_a_F_sortCommandGeneric_10)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	if v361 != 0 {
		goto L105
	} else {
		goto L106
	}
L88:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301+v97)))
	v304 = F_objectGetVal(m, v303)
	mBase = m.M
	v305 = int32(_a_F_sortCommandGeneric_11)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v308 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	if v274 < int32(1) {
		goto L87
	} else {
		goto L101
	}
L90:
	;
	v340 = F_tolower(m, v336)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v342 = F_tolower(m, v341)
	mBase = m.M
	goto L89
L91:
	;
	v310 = v304
	v311 = v305
	v312 = v308
	goto L94
L92:
	;
	v336 = int32(0)
	v337 = v305
	goto L90
L93:
	;
	v336 = v333 & int32(255)
	v337 = v332
	goto L90
L94:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v314 == int32(0) {
		v332 = v311
		v333 = v312
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v332 = v326
	v333 = int32(0)
	goto L93
L96:
	;
	v318 = v312 & int32(255)
	if v318 == v314 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v325 = int32(1)
	v326 = v311 + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	if v327 != 0 {
		v310 = v310 + v325
		v311 = v326
		v312 = v327
		goto L94
	} else {
		goto L100
	}
L98:
	;
	v320 = F_tolower(m, v318)
	mBase = m.M
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v322 = F_tolower(m, v321)
	mBase = m.M
	if v320 == v322 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	v332 = v311
	v333 = v324
	goto L93
L100:
	;
	goto L95
L101:
	;
	if v340-v342 != 0 {
		goto L87
	} else {
		goto L102
	}
L102:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v348 = v77 + int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v346+v348<<(uint(int32(2))%32))))
	v759 = v352
	v760 = v348
	v762 = v80
	v763 = v81
	v764 = v82
	v765 = v83
	v766 = v84
	goto L25
L103:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v399 = base.B2i32(v274 < int32(1))
	if v274 < int32(1) {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	v393 = F_tolower(m, v389)
	mBase = m.M
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390))))
	v395 = F_tolower(m, v394)
	mBase = m.M
	goto L103
L105:
	;
	v363 = v357
	v364 = v358
	v365 = v361
	goto L108
L106:
	;
	v389 = int32(0)
	v390 = v358
	goto L104
L107:
	;
	v389 = v386 & int32(255)
	v390 = v385
	goto L104
L108:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	if v367 == int32(0) {
		v385 = v364
		v386 = v365
		goto L107
	} else {
		goto L110
	}
L109:
	;
	v385 = v379
	v386 = int32(0)
	goto L107
L110:
	;
	v371 = v365 & int32(255)
	if v371 == v367 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v378 = int32(1)
	v379 = v364 + v378
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
	if v380 != 0 {
		v363 = v363 + v378
		v364 = v379
		v365 = v380
		goto L108
	} else {
		goto L114
	}
L112:
	;
	v373 = F_tolower(m, v371)
	mBase = m.M
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	v375 = F_tolower(m, v374)
	mBase = m.M
	if v373 == v375 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	v385 = v364
	v386 = v377
	goto L107
L114:
	;
	goto L109
L115:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v397+v97)))
	v550 = F_objectGetVal(m, v549)
	mBase = m.M
	v551 = int32(_a_F_sortCommandGeneric_12)
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v554 != 0 {
		goto L166
	} else {
		goto L167
	}
L116:
	;
	if v393-v395 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v401 = v77 + int32(1)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v397+v401<<(uint(int32(2))%32))))
	v406 = F_objectGetVal(m, v405)
	mBase = m.M
	v407 = int32(42)
	v408 = F___strchrnul(m, v406, v407)
	mBase = m.M
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if v410 == v407 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[3]))
	if v417 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	if v414 != 0 {
		goto L118
	} else {
		goto L123
	}
L120:
	;
	v414 = v408
	goto L122
L121:
	;
	v414 = int32(0)
	goto L122
L122:
	;
	goto L119
L123:
	;
	v759 = v75
	v760 = v401
	v762 = v405
	v763 = v81
	v764 = v82
	v765 = v83
	v766 = int32(1)
	goto L25
L124:
	;
	if v55 != 0 {
		v759 = v75
		v760 = v401
		v762 = v405
		v763 = v81
		v764 = v82
		v765 = v83
		v766 = v84
		goto L25
	} else {
		goto L161
	}
L125:
	;
	v420 = F_objectGetVal(m, v405)
	mBase = m.M
	v422 = F_objectGetVal(m, v405)
	mBase = m.M
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+int32(-1)))))
	switch v425 & int32(7) {
	case 0:
		goto L131
	case 1:
		goto L130
	case 2:
		goto L129
	case 3:
		goto L128
	case 4:
		goto L127
	default:
		v442 = int32(0)
		goto L126
	}
L126:
	;
	if v442 < int32(1) {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(-17))))
	v442 = v441
	goto L126
L128:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(-9))))
	v442 = v438
	goto L126
L129:
	;
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v422+int32(-5)))))
	v442 = v435
	goto L126
L130:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422+int32(-3)))))
	v442 = v432
	goto L126
L131:
	;
	v442 = int32(base.Ui32(v425) >> (uint(int32(3)) % 32))
	goto L126
L132:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	v530 = F_objectGetVal(m, v529)
	mBase = m.M
	v531 = F_getKeySlot(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L157
	}
L133:
	;
	v516 = F_crc16(m, v420, v442)
	mBase = m.M
	v527 = v516 & int32(16383)
	goto L132
L134:
	;
	v451 = int32(-1)
	v456 = v451
	v457 = int32(0)
	goto L135
L135:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v457))))
	v465 = v463 + int32(-63)
	if base.Ui32(int32(29)) < base.Ui32(v465) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L133
L137:
	;
	v506 = v457 + int32(1)
	if v506 != v442 {
		v456 = v503
		v457 = v506
		goto L135
	} else {
		goto L156
	}
L138:
	;
	v527 = v501
	goto L132
L139:
	;
	if v463 == int32(42) {
		v501 = v451
		goto L138
	} else {
		goto L142
	}
L140:
	;
	if int32(1)<<(uint(v465)%32)&int32(805306369) != 0 {
		v501 = v451
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	if v456 != int32(-1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if int32(0) <= v456 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	if v463 == int32(123) {
		v503 = v457
		goto L137
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v484 = base.B2i32(v457 == v456+int32(1))
	if v457 == v456+int32(1) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v503 = v456
	goto L137
L148:
	;
	v485 = int32(-2)
	goto L150
L149:
	;
	v485 = v456
	goto L150
L150:
	;
	if v463 == int32(125) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v488 = v485
	goto L153
L152:
	;
	v488 = v456
	goto L153
L153:
	;
	if v463 != int32(125) {
		v503 = v488
		goto L137
	} else {
		goto L154
	}
L154:
	;
	if v457 == v456+int32(1) {
		v503 = v488
		goto L137
	} else {
		goto L155
	}
L155:
	;
	v497 = F_crc16(m, v420+v456+int32(1), v457+(v456^int32(-1)))
	mBase = m.M
	v501 = v497 & int32(16383)
	goto L138
L156:
	;
	goto L136
L157:
	;
	if v527 == v531 {
		goto L124
	} else {
		goto L158
	}
L158:
	;
	F_addReplyError(m, l0, int32(_a_F_sortCommandGeneric_13))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	goto L14
L161:
	;
	F_addReplyError(m, l0, int32(_a_F_sortCommandGeneric_14))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L14
L164:
	;
	if v274 < int32(1) {
		goto L19
	} else {
		goto L176
	}
L165:
	;
	v586 = F_tolower(m, v582)
	mBase = m.M
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	v588 = F_tolower(m, v587)
	mBase = m.M
	goto L164
L166:
	;
	v556 = v550
	v557 = v551
	v558 = v554
	goto L169
L167:
	;
	v582 = int32(0)
	v583 = v551
	goto L165
L168:
	;
	v582 = v579 & int32(255)
	v583 = v578
	goto L165
L169:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v560 == int32(0) {
		v578 = v557
		v579 = v558
		goto L168
	} else {
		goto L171
	}
L170:
	;
	v578 = v572
	v579 = int32(0)
	goto L168
L171:
	;
	v564 = v558 & int32(255)
	if v564 == v560 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v571 = int32(1)
	v572 = v557 + v571
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	if v573 != 0 {
		v556 = v556 + v571
		v557 = v572
		v558 = v573
		goto L169
	} else {
		goto L175
	}
L173:
	;
	v566 = F_tolower(m, v564)
	mBase = m.M
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v568 = F_tolower(m, v567)
	mBase = m.M
	if v566 == v568 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v578 = v557
	v579 = v570
	goto L168
L175:
	;
	goto L170
L176:
	;
	if v586-v588 != 0 {
		goto L19
	} else {
		goto L177
	}
L177:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[3]))
	if v591 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	if v55 != 0 {
		goto L218
	} else {
		goto L219
	}
L179:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v598 = (v77 + int32(1)) << (uint(int32(2)) % 32)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v594+v598)))
	v601 = F_objectGetVal(m, v600)
	mBase = m.M
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	if v602 != int32(35) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v608+v598)))
	v611 = F_objectGetVal(m, v610)
	mBase = m.M
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v613+v598)))
	v616 = F_objectGetVal(m, v615)
	mBase = m.M
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616+int32(-1)))))
	switch v619 & int32(7) {
	case 0:
		goto L188
	case 1:
		goto L187
	case 2:
		goto L186
	case 3:
		goto L185
	case 4:
		goto L184
	default:
		v636 = int32(0)
		goto L183
	}
L181:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	if v605 == int32(0) {
		goto L178
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	if v636 < int32(1) {
		goto L190
	} else {
		goto L191
	}
L184:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v616+int32(-17))))
	v636 = v635
	goto L183
L185:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v616+int32(-9))))
	v636 = v632
	goto L183
L186:
	;
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v616+int32(-5)))))
	v636 = v629
	goto L183
L187:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616+int32(-3)))))
	v636 = v626
	goto L183
L188:
	;
	v636 = int32(base.Ui32(v619) >> (uint(int32(3)) % 32))
	goto L183
L189:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	v724 = F_objectGetVal(m, v723)
	mBase = m.M
	v725 = F_getKeySlot(m, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L214
	}
L190:
	;
	v710 = F_crc16(m, v611, v636)
	mBase = m.M
	v721 = v710 & int32(16383)
	goto L189
L191:
	;
	v645 = int32(-1)
	v650 = v645
	v651 = int32(0)
	goto L192
L192:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611+v651))))
	v659 = v657 + int32(-63)
	if base.Ui32(int32(29)) < base.Ui32(v659) {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L190
L194:
	;
	v700 = v651 + int32(1)
	if v700 != v636 {
		v650 = v697
		v651 = v700
		goto L192
	} else {
		goto L213
	}
L195:
	;
	v721 = v695
	goto L189
L196:
	;
	if v657 == int32(42) {
		v695 = v645
		goto L195
	} else {
		goto L199
	}
L197:
	;
	if int32(1)<<(uint(v659)%32)&int32(805306369) != 0 {
		v695 = v645
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	if v650 != int32(-1) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if int32(0) <= v650 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	if v657 == int32(123) {
		v697 = v651
		goto L194
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	v678 = base.B2i32(v651 == v650+int32(1))
	if v651 == v650+int32(1) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v697 = v650
	goto L194
L205:
	;
	v679 = int32(-2)
	goto L207
L206:
	;
	v679 = v650
	goto L207
L207:
	;
	if v657 == int32(125) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v682 = v679
	goto L210
L209:
	;
	v682 = v650
	goto L210
L210:
	;
	if v657 != int32(125) {
		v697 = v682
		goto L194
	} else {
		goto L211
	}
L211:
	;
	if v651 == v650+int32(1) {
		v697 = v682
		goto L194
	} else {
		goto L212
	}
L212:
	;
	v691 = F_crc16(m, v611+v650+int32(1), v651+(v650^int32(-1)))
	mBase = m.M
	v695 = v691 & int32(16383)
	goto L195
L213:
	;
	goto L193
L214:
	;
	if v721 == v725 {
		goto L178
	} else {
		goto L215
	}
L215:
	;
	F_addReplyError(m, l0, int32(_a_F_sortCommandGeneric_15))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L14
L218:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v744 = v77 + int32(1)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v742+v744<<(uint(int32(2))%32))))
	v750 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L222
	}
L219:
	;
	F_addReplyError(m, l0, int32(_a_F_sortCommandGeneric_16))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	goto L14
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v750)+4)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v750))) = int32(0)
	v755 = F_listAddNodeTail(m, v33, v750)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v759 = v75
	v760 = v744
	v762 = v80
	v763 = v81 + int32(1)
	v764 = v82
	v765 = v83
	v766 = v84
	goto L25
L224:
	;
	goto L24
L225:
	;
	if v802 == int32(0) {
		goto L16
	} else {
		goto L226
	}
L226:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v802)))
	if base.Ui32(v806&int32(15)+int32(-1)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L227
	}
L227:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[4]))
	F_addReplyErrorObject(m, l0, v816)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	goto L14
L230:
	;
	goto L18
L231:
	;
	goto L14
L232:
	;
	v838 = v802
	goto L15
L233:
	;
	v838 = v836
	goto L15
L234:
	;
	switch v869&int32(15) + int32(-1) {
	case 0:
		goto L243
	case 1:
		goto L246
	case 2:
		goto L245
	default:
		goto L244
	}
L235:
	;
	v859 = v781 ^ int32(1)
	if v841 != int32(3) {
		v866 = v781
		v867 = v784
		v868 = v786
		v869 = v839
		v870 = v859
		goto L234
	} else {
		goto L240
	}
L236:
	;
	if v779 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v854 = int32(0)
	v855 = int32(1)
	v866 = v855
	v867 = v854
	v868 = v855
	v869 = v839
	v870 = v854
	goto L234
L238:
	;
	v845 = int32(1)
	v846 = int32(0)
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v847&v845 == v846 {
		v866 = v846
		v867 = v784
		v868 = v786
		v869 = v839
		v870 = v845
		goto L234
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	F_zsetConvert(m, v838, int32(7))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v866 = v781
	v867 = v784
	v868 = v786
	v869 = v865
	v870 = v859
	goto L234
L242:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	v894 = int32(-1)
	if v894 < v893 {
		goto L251
	} else {
		goto L252
	}
L243:
	;
	v889 = F_listTypeLength(m, v838)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L250
	}
L244:
	;
	F__serverPanic_1(m, int32(_a_F_sortCommandGeneric_1), int32(335), int32(_a_F_sortCommandGeneric_17), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L249
	}
L245:
	;
	v877 = F_objectGetVal(m, v838)
	mBase = m.M
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v878)+20))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v878)+16))
	goto L248
L246:
	;
	v875 = F_setTypeSize(m, v838)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v891 = v875
	goto L242
L248:
	;
	v891 = v879 + v880
	goto L242
L249:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	v891 = v889
	goto L242
L251:
	;
	v897 = v893
	goto L253
L252:
	;
	v897 = v894
	goto L253
L253:
	;
	if v897 < v891 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v899 = v897
	goto L256
L255:
	;
	v899 = v891
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v899
	v901 = int32(0)
	if v901 < v892 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v904 = v892
	goto L259
L258:
	;
	v904 = v901
	goto L259
L259:
	;
	v905 = base.B2i32(v904 < v891)
	if v904 < v891 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v906 = v904
	goto L262
L261:
	;
	v906 = v891
	goto L262
L262:
	;
	if v899 < int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v910 = v891
	goto L265
L264:
	;
	v910 = v899 + v906
	goto L265
L265:
	;
	if v904 < v891 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v915 = v910 + int32(-1)
	goto L268
L267:
	;
	v915 = v891 + int32(-2)
	goto L268
L268:
	;
	v917 = v891 + int32(-1)
	if v915 < v891 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v919 = v915
	goto L271
L270:
	;
	v919 = v917
	goto L271
L271:
	;
	if v904 < v891 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v920 = v906
	goto L274
L273:
	;
	v920 = v917
	goto L274
L274:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v923 = v921 & int32(15)
	if v923 == int32(3) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v936 = F_valkey_malloc(m, v933<<(uint(int32(4))%32))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L284
	}
L276:
	;
	if v920 != 0 {
		goto L281
	} else {
		goto L282
	}
L277:
	;
	if v866 != 0 {
		v933 = v891
		goto L275
	} else {
		goto L280
	}
L278:
	;
	if v870&base.B2i32(v923 == int32(1)) != 0 {
		goto L276
	} else {
		goto L279
	}
L279:
	;
	v933 = v891
	goto L275
L280:
	;
	goto L276
L281:
	;
	v933 = v919 - v920 + int32(1)
	goto L275
L282:
	;
	if v919 == v917 {
		v933 = v891
		goto L275
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	v938 = int32(1)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v942 = v940 & int32(15)
	v944 = base.B2i32(v942 != v938)
	if v870^v938|v944 != 0 {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	if v1559 != v933 {
		goto L11
	} else {
		goto L399
	}
L286:
	;
	v1559 = v1536
	v1566 = v919
	v1569 = v920
	goto L285
L287:
	;
	if v942 != v938 {
		goto L303
	} else {
		goto L304
	}
L288:
	;
	if v919 < v920 {
		v1536 = int32(0)
		goto L286
	} else {
		goto L289
	}
L289:
	;
	if v787 == int32(0) {
		v955 = v920
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v956 = int32(0)
	v959 = F_listTypeInitIterator(m, v838, v955, base.B2i32(v787 == v956))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	v950 = F_listTypeLength(m, v838)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v955 = v950 + (v920 ^ int32(-1))
	goto L290
L293:
	;
	if v933 < int32(1) {
		v1014 = int32(0)
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_listTypeReleaseIterator(m, v959)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L302
	}
L295:
	;
	v970 = int32(0)
	goto L296
L296:
	;
	v990 = F_listTypeNext(m, v959, v26+int32(16))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L298
	}
L297:
	;
	v1014 = v933
	goto L294
L298:
	;
	if v990 == int32(0) {
		v1014 = v970
		goto L294
	} else {
		goto L299
	}
L299:
	;
	v996 = F_listTypeGet(m, v26+int32(16))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1000 = v936 + v970<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1000)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1000))) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v1000)+8)) = int32(0)
	v1007 = v970 + int32(1)
	if v1007 != v933 {
		v970 = v1007
		goto L296
	} else {
		goto L301
	}
L301:
	;
	goto L297
L302:
	;
	v1559 = v1014
	v1566 = v919 - v920
	v1569 = v956
	goto L285
L303:
	;
	if v942 != int32(2) {
		goto L315
	} else {
		goto L316
	}
L304:
	;
	v1035 = int32(0)
	v1038 = F_listTypeInitIterator(m, v838, v1035, int32(1))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L306
	}
L305:
	;
	F_listTypeReleaseIterator(m, v1038)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L314
	}
L306:
	;
	v1042 = F_listTypeNext(m, v1038, v26+int32(16))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	if v1042 == int32(0) {
		v1092 = v1035
		goto L305
	} else {
		goto L308
	}
L308:
	;
	v1051 = v1035
	goto L309
L309:
	;
	v1071 = F_listTypeGet(m, v26+int32(16))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L311
	}
L310:
	;
	v1092 = v1082
	goto L305
L311:
	;
	v1075 = v936 + v1051<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1075)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1075))) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+8)) = int32(0)
	v1082 = v1051 + int32(1)
	v1085 = F_listTypeNext(m, v1038, v26+int32(16))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	if v1085 != 0 {
		v1051 = v1082
		goto L309
	} else {
		goto L313
	}
L313:
	;
	goto L310
L314:
	;
	v1536 = v1092
	goto L286
L315:
	;
	v1187 = base.B2i32(v942 != int32(3))
	if v870^int32(1)|v1187 != 0 {
		goto L327
	} else {
		goto L328
	}
L316:
	;
	v1114 = int32(0)
	v1115 = F_setTypeInitIterator(m, v838)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L318
	}
L317:
	;
	F_setTypeReleaseIterator(m, v1115)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L326
	}
L318:
	;
	v1117 = F_setTypeNextObject(m, v1115)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	if v1117 == int32(0) {
		v1164 = v1114
		goto L317
	} else {
		goto L320
	}
L320:
	;
	v1126 = v1114
	v1134 = v1117
	goto L321
L321:
	;
	v1145 = F_createObject(m, int32(0), v1134)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L323
	}
L322:
	;
	v1164 = v1156
	goto L317
L323:
	;
	v1149 = v936 + v1126<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1149)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1149))) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+8)) = int32(0)
	v1156 = v1126 + int32(1)
	v1157 = F_setTypeNextObject(m, v1115)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	if v1157 != 0 {
		v1126 = v1156
		v1134 = v1157
		goto L321
	} else {
		goto L325
	}
L325:
	;
	goto L322
L326:
	;
	v1536 = v1164
	goto L286
L327:
	;
	if v942 != int32(3) {
		goto L12
	} else {
		goto L378
	}
L328:
	;
	v1189 = F_objectGetVal(m, v838)
	mBase = m.M
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+4))
	if v787 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	if v933 == int32(0) {
		goto L361
	} else {
		goto L362
	}
L330:
	;
	v1209 = int32(0)
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+16))
	if v1217 <= v1209 {
		v1278 = v1209
		goto L340
	} else {
		goto L341
	}
L331:
	;
	goto L336
L332:
	;
	v1193 = F_objectGetVal(m, v838)
	mBase = m.M
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+20))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+16))
	goto L333
L333:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+8))
	goto L334
L334:
	;
	if v920 < int32(1) {
		v1296 = v1198
		goto L329
	} else {
		goto L335
	}
L335:
	;
	v1208 = v1195 + v1196 - v920
	goto L330
L336:
	;
	if int32(1) <= v920 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1208 = v920 + int32(1)
	goto L330
L338:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+12))
	v1296 = v1204
	goto L329
L339:
	;
	v1296 = v1294
	goto L329
L340:
	;
	v1294 = v1278
	goto L339
L341:
	;
	v1221 = v1190
	v1224 = v1217
	v1225 = int32(0)
	goto L343
L342:
	;
	if v1242 == v1208 {
		goto L358
	} else {
		goto L359
	}
L343:
	;
	v1231 = v1224 + int32(-1)
	v1233 = v1231 << (uint(int32(3)) % 32)
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1221+v1233)+12))
	if v1235 == int32(0) {
		v1261 = v1221
		v1265 = v1225
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1294 = v1261
	goto L339
L345:
	;
	if v1265 == v1208 {
		goto L355
	} else {
		goto L356
	}
L346:
	;
	__phi1238 = v1221
	__phi1242 = v1225
	__phi1245 = v1235
	v1238 = __phi1238
	v1242 = __phi1242
	v1245 = __phi1245
	goto L347
L347:
	;
	if v1231 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L348:
	;
	v1261 = v1245
	v1265 = v1258
	goto L345
L349:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1245+v1233)+12))
	if v1260 != 0 {
		__phi1238 = v1245
		__phi1242 = v1258
		__phi1245 = v1260
		v1238 = __phi1238
		v1242 = __phi1242
		v1245 = __phi1245
		goto L347
	} else {
		goto L354
	}
L350:
	;
	v1256 = v1242 + int32(1)
	if base.Ui32(v1208) < base.Ui32(v1256) {
		goto L342
	} else {
		goto L353
	}
L351:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1238+v1233+int32(16))))
	v1253 = v1252 + v1242
	if base.Ui32(v1253) <= base.Ui32(v1208) {
		v1258 = v1253
		goto L349
	} else {
		goto L352
	}
L352:
	;
	v1261 = v1238
	v1265 = v1242
	goto L345
L353:
	;
	v1258 = v1256
	goto L349
L354:
	;
	goto L348
L355:
	;
	goto L344
L356:
	;
	if v1224 < int32(2) {
		v1278 = v1209
		goto L340
	} else {
		goto L357
	}
L357:
	;
	v1221 = v1261
	v1224 = v1231
	v1225 = v1265
	goto L343
L358:
	;
	v1275 = v1238
	goto L360
L359:
	;
	v1275 = int32(0)
	goto L360
L360:
	;
	v1278 = v1275
	goto L340
L361:
	;
	v1559 = v933
	v1566 = v919 - v920
	v1569 = int32(0)
	goto L285
L362:
	;
	if v787 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1301 = int32(8)
	goto L365
L364:
	;
	v1301 = int32(12)
	goto L365
L365:
	;
	v1308 = v1296
	v1316 = int32(0)
	v1321 = v933
	goto L366
L366:
	;
	if v1308 == int32(0) {
		goto L13
	} else {
		goto L368
	}
L367:
	;
	goto L361
L368:
	;
	v1330 = v1308 + int32(16)
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1334 = v1330 + v1331<<(uint(int32(3))%32)
	v1335 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1334))))
	v1336 = v1334 + v1335
	goto L375
L369:
	;
	v1359 = F_createStringObject_1(m, v1336+int32(1), v1358)
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L376
	}
L370:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1336+int32(-16))))
	v1358 = v1357
	goto L369
L371:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1336+int32(-8))))
	v1358 = v1354
	goto L369
L372:
	;
	v1351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1336+int32(-4)))))
	v1358 = v1351
	goto L369
L373:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336+int32(-2)))))
	v1358 = v1348
	goto L369
L374:
	;
	v1358 = int32(base.Ui32(v1341) >> (uint(int32(3)) % 32))
	goto L369
L375:
	;
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336+int32(0)))))
	switch v1341 & int32(7) {
	case 0:
		goto L374
	case 1:
		goto L373
	case 2:
		goto L372
	case 3:
		goto L371
	case 4:
		goto L370
	default:
		v1358 = int32(0)
		goto L369
	}
L376:
	;
	v1363 = v936 + v1316<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1363)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1363))) = v1359
	*(*int32)(unsafe.Add(mBase, uint32(v1363)+8)) = int32(0)
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1308+v1301)))
	v1374 = v1321 + int32(-1)
	if v1374 != 0 {
		v1308 = v1372
		v1316 = v1316 + int32(1)
		v1321 = v1374
		goto L366
	} else {
		goto L377
	}
L377:
	;
	goto L367
L378:
	;
	v1400 = int32(0)
	v1402 = v26 + int32(16)
	v1403 = F_objectGetVal(m, v838)
	mBase = m.M
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1403)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1402)+14)) = uint8(v1400)
	*(*int32)(unsafe.Add(mBase, uint32(v1402))) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+24)) = v1400
	*(*uint8)(unsafe.Add(mBase, uint32(v1402)+15)) = uint8(v1400)
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+8)) = int32(-1)
	if v1404 == v1400 {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	v1427 = F_hashtableNext(m, v26+int32(16), v26+int32(12))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L384
	}
L380:
	;
	goto L379
L381:
	;
	goto L380
L383:
	;
	F_hashtableCleanupIterator(m, v26+int32(16))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L398
	}
L384:
	;
	if v1427 == int32(0) {
		v1509 = v1400
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1436 = v1400
	goto L386
L386:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v1457 = v1455 + int32(16)
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1457)))
	v1461 = v1457 + v1458<<(uint(int32(3))%32)
	v1462 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1461))))
	v1463 = v1461 + v1462
	goto L394
L387:
	;
	v1509 = v1497
	goto L383
L388:
	;
	v1486 = F_createStringObject_1(m, v1463+int32(1), v1485)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L1
	} else {
		goto L395
	}
L389:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1463+int32(-16))))
	v1485 = v1484
	goto L388
L390:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1463+int32(-8))))
	v1485 = v1481
	goto L388
L391:
	;
	v1478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1463+int32(-4)))))
	v1485 = v1478
	goto L388
L392:
	;
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463+int32(-2)))))
	v1485 = v1475
	goto L388
L393:
	;
	v1485 = int32(base.Ui32(v1468) >> (uint(int32(3)) % 32))
	goto L388
L394:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463+int32(0)))))
	switch v1468 & int32(7) {
	case 0:
		goto L393
	case 1:
		goto L392
	case 2:
		goto L391
	case 3:
		goto L390
	case 4:
		goto L389
	default:
		v1485 = int32(0)
		goto L388
	}
L395:
	;
	v1490 = v936 + v1436<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1490)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1490))) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+8)) = int32(0)
	v1497 = v1436 + int32(1)
	v1502 = F_hashtableNext(m, v26+int32(16), v26+int32(12))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	if v1502 != 0 {
		v1436 = v1497
		goto L386
	} else {
		goto L397
	}
L397:
	;
	goto L387
L398:
	;
	v1536 = v1509
	goto L286
L399:
	;
	if v870 != 0 {
		v1821 = int32(0)
		goto L400
	} else {
		goto L401
	}
L400:
	;
	if v1821 == int32(0) {
		goto L453
	} else {
		goto L454
	}
L401:
	;
	if v933 < int32(1) {
		v1771 = int32(0)
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1777 = int32(_a_F_sortCommandGeneric_5)
	*(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[5])) = v868
	*(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[6])) = v787
	v1782 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[7])) = base.B2i32(v779 != v1782)
	*(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[8])) = base.B2i32(v867 != v1782)
	if v933 == v1782 {
		v1821 = v1771
		goto L400
	} else {
		goto L444
	}
L403:
	;
	v1582 = int32(0)
	v1589 = v1582
	v1601 = v1582
	goto L404
L404:
	;
	if v867 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L405:
	;
	v1771 = v1749
	goto L402
L406:
	;
	v1752 = v1589 + int32(1)
	if v1752 != v933 {
		v1589 = v1752
		v1601 = v1749
		goto L404
	} else {
		goto L443
	}
L407:
	;
	F_decrRefCount(m, v1742)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L442
	}
L408:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1628)))
	switch int32(base.Ui32(v1629)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L419
	case 1:
		goto L417
	default:
		goto L418
	}
L409:
	;
	if v868 != 0 {
		v1749 = v1601
		goto L406
	} else {
		goto L415
	}
L410:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1612 = v936 + v1589<<(uint(int32(4))%32)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1612)))
	v1614 = F_lookupKeyByPattern(m, v1609, v867, v1613)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	if v1614 == int32(0) {
		v1749 = v1601
		goto L406
	} else {
		goto L412
	}
L412:
	;
	if v868 == int32(0) {
		v1628 = v1614
		goto L408
	} else {
		goto L413
	}
L413:
	;
	v1620 = F_getDecodedObject(m, v1614)
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+8)) = v1620
	v1742 = v1614
	v1743 = v1601
	goto L407
L415:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v936+v1589<<(uint(int32(4))%32))))
	v1628 = v1626
	goto L408
L416:
	;
	if v867 == int32(0) {
		v1749 = v1737
		goto L406
	} else {
		goto L441
	}
L417:
	;
	v1733 = F_objectGetVal(m, v1628)
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v936+v1589<<(uint(int32(4))%32))+8)) = base.F64_convert_i32_s(v1733)
	v1737 = v1601
	goto L416
L418:
	;
	F__serverAssertWithInfo(m, l0, v838, int32(_a_F_sortCommandGeneric_18), int32(_a_F_sortCommandGeneric_1), int32(497))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L1
	} else {
		goto L440
	}
L419:
	;
	v1634 = int32(9116376)
	goto L420
L420:
	;
	v1635 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[9])) = v1635
	v1640 = F_objectGetVal(m, v1628)
	mBase = m.M
	v1642 = v26 + int32(16)
	v1648 = m.G0
	v1650 = v1648 - int32(32)
	m.G0 = v1650
	v1652 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1652))) = v1635
	v1658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1640+int32(-1)))))
	switch v1658 & int32(7) {
	case 0:
		goto L427
	case 1:
		goto L426
	case 2:
		goto L425
	case 3:
		goto L424
	case 4:
		goto L423
	default:
		v1675 = v1635
		goto L422
	}
L421:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v936+v1589<<(uint(int32(4))%32))+8)) = v1705
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v1711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1710))))
	if v1711 != 0 {
		goto L435
	} else {
		goto L436
	}
L422:
	;
	v1678 = int32(0)
	v1679 = *(*int64)(unsafe.Add(mBase, _c_F_sortCommandGeneric[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v1650+int32(8)))) = v1679
	*(*int64)(unsafe.Add(mBase, uint32(v1650)+24)) = int64(0)
	v1684 = *(*int64)(unsafe.Add(mBase, _c_F_sortCommandGeneric[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v1650))) = v1684
	F_ffc_from_chars_double_options(m, v1650+int32(16), v1640, v1640+v1675, v1650+int32(24), v1650)
	mBase = m.M
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+20))
	if v1692 == v1678 {
		goto L428
	} else {
		goto L429
	}
L423:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1640+int32(-17))))
	v1675 = v1674
	goto L422
L424:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1640+int32(-9))))
	v1675 = v1671
	goto L422
L425:
	;
	v1668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1640+int32(-5)))))
	v1675 = v1668
	goto L422
L426:
	;
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1640+int32(-3)))))
	v1675 = v1665
	goto L422
L427:
	;
	v1675 = int32(base.Ui32(v1658) >> (uint(int32(3)) % 32))
	goto L422
L428:
	;
	if v1642 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L429:
	;
	if v1692 == int32(2) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1699 = int32(68)
	goto L432
L431:
	;
	v1699 = int32(28)
	goto L432
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1652))) = v1699
	goto L428
L433:
	;
	v1705 = *(*float64)(unsafe.Add(mBase, uint32(v1650)+24))
	m.G0 = v1650 + int32(32)
	goto L421
L434:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1642))) = v1703
	goto L433
L435:
	;
	v1737 = int32(1)
	goto L416
L436:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[9]))
	if v1712 == int32(28) {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	if v1712 == int32(68) {
		goto L435
	} else {
		goto L438
	}
L438:
	;
	if base.Ui64(base.I64_reinterpret_f64(v1705)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v1737 = v1601
		goto L416
	} else {
		goto L439
	}
L439:
	;
	goto L435
L440:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	v1742 = v1628
	v1743 = v1737
	goto L407
L442:
	;
	v1749 = v1743
	goto L406
L443:
	;
	goto L405
L444:
	;
	if v867 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	F_qsort(m, v936, v933, int32(16), int32(1086))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L1
	} else {
		goto L451
	}
L446:
	;
	if v1569 != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	F_pqsort(m, v936, v933, int32(16), int32(1086), v1569, v1566)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L450
	}
L448:
	;
	if v1566 == v933+int32(-1) {
		goto L445
	} else {
		goto L449
	}
L449:
	;
	goto L447
L450:
	;
	v1821 = v1771
	goto L400
L451:
	;
	v1821 = v1771
	goto L400
L452:
	;
	if v933 < int32(1) {
		goto L535
	} else {
		goto L536
	}
L453:
	;
	v1832 = int32(1)
	v1833 = v1566 + v1832
	if base.Ui32(v1832) < base.Ui32(v785) {
		goto L456
	} else {
		goto L457
	}
L454:
	;
	F_addReplyError(m, l0, int32(_a_F_sortCommandGeneric_19))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	goto L452
L456:
	;
	v1838 = v785
	goto L458
L457:
	;
	v1838 = v1832
	goto L458
L458:
	;
	v1839 = (v1833 - v1569) * v1838
	if v779 != 0 {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1981 = int32(_a_F_sortCommandGeneric_5)
	v1982 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[0]))
	v1984 = *(*int32)(unsafe.Add(mBase, _c_F_sortCommandGeneric[1]))
	v1985 = F_createQuicklistObject(m, v1982, v1984)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L491
	}
L460:
	;
	F_addReplyArrayLen(m, l0, v1839)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L461
	}
L461:
	;
	if v1566 < v1569 {
		goto L452
	} else {
		goto L462
	}
L462:
	;
	v1858 = v1569
	goto L463
L463:
	;
	if v785 != 0 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v1873 = v26 + int32(16)
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v1873)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1873))) = v1874
	goto L468
L466:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v936+v1858<<(uint(int32(4))%32))))
	F_addReplyBulk(m, l0, v1869)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	goto L465
L468:
	;
	v1879 = v26 + int32(16)
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1879)))
	if v1881 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L469:
	;
	if base.B2i32(v1858 == v1566) == int32(0) {
		v1858 = v1858 + int32(1)
		goto L463
	} else {
		goto L490
	}
L470:
	;
	if v1881 == int32(0) {
		goto L469
	} else {
		goto L473
	}
L471:
	;
	goto L470
L472:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1879)+4))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1881+base.B2i32(v1884 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1879))) = v1890
	goto L471
L473:
	;
	v1902 = v1881
	goto L474
L474:
	;
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1902)+8))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+4))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v936+v1858<<(uint(int32(4))%32))))
	v1924 = F_lookupKeyByPattern(m, v1920, v1922, v1923)
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L476
	}
L475:
	;
	goto L469
L476:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1921)))
	if v1926 != 0 {
		goto L479
	} else {
		goto L480
	}
L477:
	;
	v1940 = v26 + int32(16)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1940)))
	if v1942 == int32(0) {
		goto L487
	} else {
		goto L488
	}
L478:
	;
	F_addReplyBulk(m, l0, v1924)
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L1
	} else {
		goto L484
	}
L479:
	;
	F__serverAssertWithInfo(m, l0, v838, int32(_a_F_sortCommandGeneric_0), int32(_a_F_sortCommandGeneric_1), int32(549))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L483
	}
L480:
	;
	if v1924 != 0 {
		goto L478
	} else {
		goto L481
	}
L481:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	goto L477
L483:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L484:
	;
	F_decrRefCount(m, v1924)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	goto L477
L486:
	;
	if v1942 != 0 {
		v1902 = v1942
		goto L474
	} else {
		goto L489
	}
L487:
	;
	goto L486
L488:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+4))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1942+base.B2i32(v1945 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1940))) = v1951
	goto L487
L489:
	;
	goto L475
L490:
	;
	goto L452
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v1985
	if v1566 < v1569 {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	if v1833 == v1569 {
		goto L521
	} else {
		goto L522
	}
L493:
	;
	v1994 = v1569
	goto L494
L494:
	;
	if v785 != 0 {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	goto L492
L496:
	;
	if v1994 != v1566 {
		v1994 = v1994 + int32(1)
		goto L494
	} else {
		goto L518
	}
L497:
	;
	v2020 = v26 + int32(16)
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v2020)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2020))) = v2021
	goto L500
L498:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v936+v1994<<(uint(int32(4))%32))))
	F_listTypePush(m, v1985, v2015, int32(1))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	goto L496
L500:
	;
	v2026 = v26 + int32(16)
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2026)))
	if v2028 == int32(0) {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v2028 == int32(0) {
		goto L496
	} else {
		goto L504
	}
L502:
	;
	goto L501
L503:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2026)+4))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2028+base.B2i32(v2031 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2026))) = v2037
	goto L502
L504:
	;
	v2049 = v2028
	goto L505
L505:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v2049)+8))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+4))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v936+v1994<<(uint(int32(4))%32))))
	v2071 = F_lookupKeyByPattern(m, v2067, v2069, v2070)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L1
	} else {
		goto L507
	}
L506:
	;
	goto L496
L507:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2068)))
	if v2073 != 0 {
		goto L10
	} else {
		goto L508
	}
L508:
	;
	if v2071 != 0 {
		v2078 = v2071
		goto L509
	} else {
		goto L510
	}
L509:
	;
	F_listTypePush(m, v1985, v2078, int32(1))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L1
	} else {
		goto L512
	}
L510:
	;
	v2076 = F_createStringObject_1(m, int32(_a_F_sortCommandGeneric_20), int32(0))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	v2078 = v2076
	goto L509
L512:
	;
	F_decrRefCount(m, v2078)
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	v2085 = v26 + int32(16)
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2085)))
	if v2087 == int32(0) {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	if v2087 != 0 {
		v2049 = v2087
		goto L505
	} else {
		goto L517
	}
L515:
	;
	goto L514
L516:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2085)+4))
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2087+base.B2i32(v2090 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2085))) = v2096
	goto L515
L517:
	;
	goto L506
L518:
	;
	goto L495
L519:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v2190 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L520:
	;
	v2184 = int32(_a_F_sortCommandGeneric_5)
	v2186 = *(*int64)(unsafe.Add(mBase, _c_F_sortCommandGeneric[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_sortCommandGeneric[12])) = v2186 + v2183
	goto L519
L521:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2169 = F_dbDelete(m, v2168, v779)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L1
	} else {
		goto L526
	}
L522:
	;
	v2148 = int32(0)
	F_listTypeTryConversion(m, v1985, v2148, v2148, v2148)
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v2153, v779, v26+int32(12), int32(0))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v2163)+28))
	F_notifyKeyspaceEvent(m, int32(16), int32(_a_F_sortCommandGeneric_21), v779, v2164)
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v2183 = base.I64_extend_i32_u(v1839)
	goto L520
L526:
	;
	if v2169 == int32(0) {
		goto L519
	} else {
		goto L527
	}
L527:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v2173, v779)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v2178)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_sortCommandGeneric_22), v779, v2179)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v2183 = int64(1)
	goto L520
L530:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v1839))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L1
	} else {
		goto L533
	}
L531:
	;
	F_decrRefCount(m, v2190)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	goto L530
L533:
	;
	goto L452
L534:
	;
	F_valkey_free(m, v936)
	mBase = m.M
	v2326 = m.ExcPending
	if v2326 != 0 {
		goto L1
	} else {
		goto L552
	}
L535:
	;
	F_decrRefCount(m, v838)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L550
	}
L536:
	;
	v2224 = int32(0)
	goto L537
L537:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v936+v2224<<(uint(int32(4))%32))))
	F_decrRefCount(m, v2250)
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L1
	} else {
		goto L539
	}
L538:
	;
	F_decrRefCount(m, v838)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L1
	} else {
		goto L541
	}
L539:
	;
	v2254 = v2224 + int32(1)
	if v2254 != v933 {
		v2224 = v2254
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v2261 = int32(0)
	goto L543
L543:
	;
	if v868 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2296 = v2261 + int32(1)
	if v2296 != v933 {
		v2261 = v2296
		goto L543
	} else {
		goto L549
	}
L546:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v936+v2261<<(uint(int32(4))%32))+8))
	if v2289 == int32(0) {
		goto L545
	} else {
		goto L547
	}
L547:
	;
	F_decrRefCount(m, v2289)
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	goto L545
L549:
	;
	goto L534
L550:
	;
	F_listRelease(m, v33)
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	goto L534
L552:
	;
	goto L14
L553:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L554:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L556:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_waitCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_waitCommand[0]))
	if base.B2i32(v11 != v2) == int32(0) {
		v18 = l0
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_waitCommand[1]))
		if v20 == int32(0) {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			v32 = F_getLongFromObjectOrReply(m, l0, v28, v8+int32(12), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 != 0 {
					m.G0 = v8 + int32(32)
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v39 = F_getTimeoutFromObjectOrReply(m, l0, v35, v8+int32(16), int32(1))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						if v39 != 0 {
							m.G0 = v8 + int32(32)
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_waitCommand[2]))
							v44 = v8 + int32(24)
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v44))) = v45
							v49 = int32(0)
							v51 = v8 + int32(24)
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							if v53 == v49 {
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v53+base.B2i32(v56 == int32(0))<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = v62
							}
							if v53 == int32(0) {
								v98 = v49
							} else {
								v68 = v53
								v70 = v49
								for {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									if v73 != int32(9) {
										v79 = v70
									} else {
										v76 = *(*int64)(unsafe.Add(mBase, uint32(v72)+64))
										v79 = v70 + base.B2i32(v26 <= v76)
									}
									v81 = v8 + int32(24)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
									if v83 == int32(0) {
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+base.B2i32(v86 == int32(0))<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v81))) = v92
									}
									if v83 != 0 {
										v68 = v83
										v70 = v79
										continue
									} else {
										break
									}
									break
								}
								v98 = v79
							}
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							if v99 <= v98 {
								F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v98))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							} else {
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
								if v101&int32(16) == int32(0) {
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
									F_blockClientForReplicaAck(m, l0, v109, v26, v99, int32(0))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_waitCommand[3])) = int32(1)
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v98))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_addReplyError(m, l0, int32(_a_F_waitCommand_0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	} else {
		v16 = F_scriptGetCaller(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = v16
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_waitCommand[1]))
			if v20 == int32(0) {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v32 = F_getLongFromObjectOrReply(m, l0, v28, v8+int32(12), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 != 0 {
						m.G0 = v8 + int32(32)
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
						v39 = F_getTimeoutFromObjectOrReply(m, l0, v35, v8+int32(16), int32(1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							if v39 != 0 {
								m.G0 = v8 + int32(32)
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, _c_F_waitCommand[2]))
								v44 = v8 + int32(24)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
								*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = v45
								v49 = int32(0)
								v51 = v8 + int32(24)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								if v53 == v49 {
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v53+base.B2i32(v56 == int32(0))<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v62
								}
								if v53 == int32(0) {
									v98 = v49
								} else {
									v68 = v53
									v70 = v49
									for {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+104))
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										if v73 != int32(9) {
											v79 = v70
										} else {
											v76 = *(*int64)(unsafe.Add(mBase, uint32(v72)+64))
											v79 = v70 + base.B2i32(v26 <= v76)
										}
										v81 = v8 + int32(24)
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
										if v83 == int32(0) {
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+base.B2i32(v86 == int32(0))<<(uint(int32(2))%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v81))) = v92
										}
										if v83 != 0 {
											v68 = v83
											v70 = v79
											continue
										} else {
											break
										}
										break
									}
									v98 = v79
								}
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								if v99 <= v98 {
									F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v98))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
									if v101&int32(16) == int32(0) {
										v109 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
										F_blockClientForReplicaAck(m, l0, v109, v26, v99, int32(0))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_waitCommand[3])) = int32(1)
											m.G0 = v8 + int32(32)
											return
										}
									} else {
										F_addReplyLongLong(m, l0, base.I64_extend_i32_s(v98))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_addReplyError(m, l0, int32(_a_F_waitCommand_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			}
		}
	}
}
