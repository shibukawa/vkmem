package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_aofDelHistoryFiles(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	if v11 == v1 {
		v111 = v1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a85), int32(_a68), int32(670))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return v111
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	if v15 == int32(1) {
		v111 = v1
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v19 == int32(0) {
		v111 = v1
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = v7 + int32(8)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	goto L6
L6:
	;
	v29 = v7 + int32(8)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = int32(1)
	v100 = F_getAofManifestAsString(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L27
	}
L8:
	;
	if v31 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L9
L11:
	;
	v45 = v31
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v49 != int32(104) {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L7
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v53 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v67 = F_makePath(m, v65, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L19
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
	F__serverLog(m, int32(2), int32(_a86), v7)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L15
L19:
	;
	v69 = F_bg_unlink(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_sdsfree(m, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	F_listDelNode(m, v75, v45)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v79 = v7 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v81 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v81 != 0 {
		v45 = v81
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81+base.B2i32(v84 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v90
	goto L24
L26:
	;
	goto L13
L27:
	;
	v102 = F_writeAofManifestFile(m, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	F_sdsfree(m, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	if v102 != 0 {
		v111 = int32(-1)
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v107
	v111 = v107
	goto L2
L31:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_aofInfoFormat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	switch v21 & int32(7) {
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
		v72 = v3
		goto L4
	}
L1:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(28)))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)))) = int32(_a63)
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(16)))) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(_a64)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a65)
	v134 = F_sdscatprintf(m, l0, int32(_a66), v11)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L23
	} else {
		goto L33
	}
L2:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v113 = int32(0)
	v114 = v110
	goto L1
L3:
	;
	if v72 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L4:
	;
	goto L3
L5:
	;
	if v38 == int32(0) {
		v72 = v3
		goto L4
	} else {
		goto L11
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
	v38 = v37
	goto L5
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
	v38 = v34
	goto L5
L8:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
	v38 = v31
	goto L5
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
	v38 = v28
	goto L5
L10:
	;
	v38 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v41 = v13
	v44 = v38
	goto L13
L12:
	;
	v72 = int32(1)
	goto L4
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v48 = v46 + int32(-7)
	if base.Ui32(int32(27)) < base.Ui32(v48) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v46 == int32(92) {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	if int32(1)<<(uint(v48)%32)&int32(134217807) != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	if v46 == int32(32) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(base.I32_extend8_s(v46)+int32(-127)) <= base.Ui32(int32(-96)) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v67 = v44 + int32(-1)
	if v67 == int32(0) {
		v72 = v3
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v41 = v41 + int32(1)
	v44 = v67
	goto L13
L22:
	;
	v78 = F_sdsempty(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-1)))))
	switch v86 & int32(7) {
	case 0:
		goto L30
	case 1:
		goto L29
	case 2:
		goto L28
	case 3:
		goto L27
	case 4:
		goto L26
	default:
		v103 = int32(0)
		goto L25
	}
L25:
	;
	v104 = F_sdscatrepr(m, v78, v83, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L23
	} else {
		goto L31
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-17))))
	v103 = v102
	goto L25
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+int32(-9))))
	v103 = v99
	goto L25
L28:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+int32(-5)))))
	v103 = v96
	goto L25
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+int32(-3)))))
	v103 = v93
	goto L25
L30:
	;
	v103 = int32(base.Ui32(v86) >> (uint(int32(3)) % 32))
	goto L25
L31:
	;
	if v104 != 0 {
		v113 = v104
		v114 = v104
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L2
L33:
	;
	F_sdsfree(m, v113)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	m.G0 = v11 + int32(32)
	return v134
}
func F_aofListDup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	if l0 != 0 {
		v12 = F_valkey_calloc(m, int32(24))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v15 = F_sdsdup(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
				v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v20
				return v12
			}
		}
	} else {
		F__serverAssert(m, int32(_a67), int32(_a68), int32(112))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
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
func F_aofLoadManifestFromDisk(m *base.Module) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_valkey_calloc(m, int32(40))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_listCreate(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v15
			v18 = F_listCreate(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v18
				v21 = int32(15)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v21
				v23 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v23
				v29 = int32(_a69)
				*(*int32)(unsafe.Add(mBase, _consts[12])) = v13
				v32 = *(*int32)(unsafe.Add(mBase, _consts[13]))
				v35 = m.G0
				v36 = int32(96)
				v37 = v35 - v36
				m.G0 = v37
				v39 = F_stat(m, v32, v37)
				mBase = m.M
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
				m.G0 = v37 + v36
				if base.B2i32(v39 == int32(0))&base.B2i32(v40&int32(61440) == int32(16384)) != 0 {
					v62 = F_sdsempty(m)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a70)
						v67 = *(*int32)(unsafe.Add(mBase, _consts[14]))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v67
						v72 = F_sdscatprintf(m, v62, int32(_a71), v10+int32(32))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, _consts[13]))
							v76 = F_makePath(m, v75, v72)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								v80 = m.G0
								v81 = int32(96)
								v82 = v80 - v81
								m.G0 = v82
								v84 = F_stat(m, v76, v82)
								mBase = m.M
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
								m.G0 = v82 + v81
								if base.B2i32(v84 == int32(0))&base.B2i32(v85&int32(61440) == int32(32768)) != 0 {
									v107 = F_aofLoadManifestFromFile(m, v76)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										if v107 == int32(0) {
											F_sdsfree(m, v72)
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												F_sdsfree(m, v76)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, _consts[12]))
											if v112 == int32(0) {
												*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
												F_sdsfree(m, v72)
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													F_sdsfree(m, v76)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return
													} else {
														m.G0 = v10 + int32(48)
														return
													}
												}
											} else {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
												if v115 == int32(0) {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
													if v126 == int32(0) {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
														if v131 == int32(0) {
															F_valkey_free(m, v112)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																F_sdsfree(m, v72)
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return
																} else {
																	F_sdsfree(m, v76)
																	mBase = m.M
																	v149 = m.ExcPending
																	if v149 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(48)
																		return
																	}
																}
															}
														} else {
															F_listRelease(m, v131)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return
															} else {
																F_valkey_free(m, v112)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																	F_sdsfree(m, v72)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return
																	} else {
																		F_sdsfree(m, v76)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													} else {
														F_listRelease(m, v126)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
															if v131 == int32(0) {
																F_valkey_free(m, v112)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																	F_sdsfree(m, v72)
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return
																	} else {
																		F_sdsfree(m, v76)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(48)
																			return
																		}
																	}
																}
															} else {
																F_listRelease(m, v131)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return
																} else {
																	F_valkey_free(m, v112)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																		F_sdsfree(m, v72)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return
																		} else {
																			F_sdsfree(m, v76)
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(48)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
													if v118 == int32(0) {
														F_valkey_free(m, v115)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															v126 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
															if v126 == int32(0) {
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
																if v131 == int32(0) {
																	F_valkey_free(m, v112)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																		F_sdsfree(m, v72)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return
																		} else {
																			F_sdsfree(m, v76)
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(48)
																				return
																			}
																		}
																	}
																} else {
																	F_listRelease(m, v131)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return
																	} else {
																		F_valkey_free(m, v112)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																			F_sdsfree(m, v72)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return
																			} else {
																				F_sdsfree(m, v76)
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(48)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																F_listRelease(m, v126)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
																	if v131 == int32(0) {
																		F_valkey_free(m, v112)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																			F_sdsfree(m, v72)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return
																			} else {
																				F_sdsfree(m, v76)
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(48)
																					return
																				}
																			}
																		}
																	} else {
																		F_listRelease(m, v131)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return
																		} else {
																			F_valkey_free(m, v112)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																				F_sdsfree(m, v72)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return
																				} else {
																					F_sdsfree(m, v76)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(48)
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
														F_sdsfree(m, v118)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return
														} else {
															F_valkey_free(m, v115)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																v126 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
																if v126 == int32(0) {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
																	if v131 == int32(0) {
																		F_valkey_free(m, v112)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																			F_sdsfree(m, v72)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return
																			} else {
																				F_sdsfree(m, v76)
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(48)
																					return
																				}
																			}
																		}
																	} else {
																		F_listRelease(m, v131)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return
																		} else {
																			F_valkey_free(m, v112)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																				F_sdsfree(m, v72)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return
																				} else {
																					F_sdsfree(m, v76)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(48)
																						return
																					}
																				}
																			}
																		}
																	}
																} else {
																	F_listRelease(m, v126)
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
																		return
																	} else {
																		v131 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
																		if v131 == int32(0) {
																			F_valkey_free(m, v112)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																				F_sdsfree(m, v72)
																				mBase = m.M
																				v147 = m.ExcPending
																				if v147 != 0 {
																					return
																				} else {
																					F_sdsfree(m, v76)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return
																					} else {
																						m.G0 = v10 + int32(48)
																						return
																					}
																				}
																			}
																		} else {
																			F_listRelease(m, v131)
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return
																			} else {
																				F_valkey_free(m, v112)
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, _consts[12])) = v107
																					F_sdsfree(m, v72)
																					mBase = m.M
																					v147 = m.ExcPending
																					if v147 != 0 {
																						return
																					} else {
																						F_sdsfree(m, v76)
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return
																						} else {
																							m.G0 = v10 + int32(48)
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
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(0) < v97 {
										F_sdsfree(m, v72)
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											F_sdsfree(m, v76)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												m.G0 = v10 + int32(48)
												return
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v72
										F__serverLog(m, int32(0), int32(_a72), v10+int32(16))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											F_sdsfree(m, v72)
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												F_sdsfree(m, v76)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
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
					v52 = *(*int32)(unsafe.Add(mBase, _consts[15]))
					if int32(0) < v52 {
						m.G0 = v10 + int32(48)
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _consts[13]))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v56
						F__serverLog(m, int32(0), int32(_a73), v10)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			}
		}
	}
}
func F_aofManifestFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v15 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v20 == int32(0) {
				F_valkey_free(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			} else {
				F_listRelease(m, v20)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			F_listRelease(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v20 == int32(0) {
					F_valkey_free(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				} else {
					F_listRelease(m, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		if v7 == int32(0) {
			F_valkey_free(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v15 == int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v20 == int32(0) {
						F_valkey_free(m, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							return
						}
					} else {
						F_listRelease(m, v20)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							F_valkey_free(m, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_listRelease(m, v15)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v20 == int32(0) {
							F_valkey_free(m, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								return
							}
						} else {
							F_listRelease(m, v20)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								F_valkey_free(m, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			F_sdsfree(m, v7)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_valkey_free(m, v4)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v15 == int32(0) {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v20 == int32(0) {
							F_valkey_free(m, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								return
							}
						} else {
							F_listRelease(m, v20)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								F_valkey_free(m, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_listRelease(m, v15)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return
						} else {
							v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v20 == int32(0) {
								F_valkey_free(m, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									return
								}
							} else {
								F_listRelease(m, v20)
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									F_valkey_free(m, l0)
									mBase = m.M
									v26 = m.ExcPending
									if v26 != 0 {
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
func F_aofRemoveTempFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v59 int32
	_ = v59
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
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	v4 = m.G0
	v6 = v4 - int32(544)
	m.G0 = v6
	v9 = base.I64_extend_i32_s(l0)
	if v9 <= int64(-1) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	goto L13
L2:
	;
	goto L1
L4:
	;
	v31 = F_ull2string(m, v27, v28, v29)
	mBase = m.M
	if v31 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	goto L7
L6:
	;
	v27 = v6
	v28 = int32(32)
	v29 = v9
	goto L4
L7:
	;
	v18 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v18)
	v27 = v6 + int32(1)
	v28 = int32(31)
	v29 = int64(0) - v9
	goto L4
L8:
	;
	goto L1
L10:
	;
	v98 = v6 + int32(288)
	v99 = int32(256)
	goto L24
L11:
	;
	goto L10
L12:
	;
	v84 = v63
	goto L19
L13:
	;
	v59 = v6 + int32(288)
	v61 = int32(256)
	v63 = int32(_a94)
	goto L15
L14:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v74)
	goto L12
L15:
	;
	v65 = v61 + int32(-1)
	if v65 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v68)
	v70 = int32(1)
	if v68 != 0 {
		v59 = v59 + v70
		v61 = v65
		v63 = v63 + v70
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 != 0 {
		v84 = v84 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L11
L21:
	;
	goto L20
L22:
	;
	v174 = v6 + int32(288)
	v175 = int32(_a75)
	v176 = int32(256)
	goto L42
L23:
	;
	v128 = v124 - v98
	if v99 == v128 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v110 = v98
	v112 = v99
	goto L25
L25:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == int32(0) {
		v124 = v110
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v124 = v6 + int32(544)
	goto L23
L27:
	;
	v120 = v112 + int32(-1)
	if v120 != 0 {
		v110 = v110 + int32(1)
		v112 = v120
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v161)
	goto L22
L30:
	;
	v136 = v130
	v138 = v99 + (v128 ^ int32(-1))
	v139 = v124
	v141 = v6
	goto L34
L31:
	;
	v131 = F_strlen(m, v6)
	mBase = m.M
	goto L22
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v130 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v157 = v124
	goto L29
L34:
	;
	if v138 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v157 = v150
	goto L29
L36:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v151 != 0 {
		v136 = v151
		v138 = v149
		v139 = v150
		v141 = v141 + int32(1)
		goto L34
	} else {
		goto L39
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v136)
	v149 = v138 + int32(-1)
	v150 = v139 + int32(1)
	goto L36
L38:
	;
	v149 = int32(0)
	v150 = v139
	goto L36
L39:
	;
	goto L35
L40:
	;
	goto L61
L41:
	;
	v205 = v201 - v174
	if v176 == v205 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v187 = v174
	v189 = v176
	goto L43
L43:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v191 == int32(0) {
		v201 = v187
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v201 = v6 + int32(544)
	goto L41
L45:
	;
	v197 = v189 + int32(-1)
	if v197 != 0 {
		v187 = v187 + int32(1)
		v189 = v197
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v238)
	goto L40
L48:
	;
	v213 = v207
	v215 = v176 + (v205 ^ int32(-1))
	v216 = v201
	v218 = v175
	goto L52
L49:
	;
	v208 = F_strlen(m, v175)
	mBase = m.M
	goto L40
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v207 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v234 = v201
	goto L47
L52:
	;
	if v215 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v234 = v227
	goto L47
L54:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v228 != 0 {
		v213 = v228
		v215 = v226
		v216 = v227
		v218 = v218 + int32(1)
		goto L52
	} else {
		goto L57
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v213)
	v226 = v215 + int32(-1)
	v227 = v216 + int32(1)
	goto L54
L56:
	;
	v226 = int32(0)
	v227 = v216
	goto L54
L57:
	;
	goto L53
L58:
	;
	v297 = v6 + int32(32)
	v298 = int32(256)
	goto L72
L59:
	;
	goto L58
L60:
	;
	v283 = v262
	goto L67
L61:
	;
	v258 = v6 + int32(32)
	v260 = int32(256)
	v262 = int32(_a95)
	goto L63
L62:
	;
	v273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v273)
	goto L60
L63:
	;
	v264 = v260 + int32(-1)
	if v264 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v267)
	v269 = int32(1)
	if v267 != 0 {
		v258 = v258 + v269
		v260 = v264
		v262 = v262 + v269
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v285 != 0 {
		v283 = v283 + int32(1)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L59
L69:
	;
	goto L68
L70:
	;
	v373 = v6 + int32(32)
	v374 = int32(_a75)
	v375 = int32(256)
	goto L90
L71:
	;
	v327 = v323 - v297
	if v298 == v327 {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v309 = v297
	v311 = v298
	goto L73
L73:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v313 == int32(0) {
		v323 = v309
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v323 = v6 + int32(288)
	goto L71
L75:
	;
	v319 = v311 + int32(-1)
	if v319 != 0 {
		v309 = v309 + int32(1)
		v311 = v319
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356))) = uint8(v360)
	goto L70
L78:
	;
	v335 = v329
	v337 = v298 + (v327 ^ int32(-1))
	v338 = v323
	v340 = v6
	goto L82
L79:
	;
	v330 = F_strlen(m, v6)
	mBase = m.M
	goto L70
L80:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v329 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v356 = v323
	goto L77
L82:
	;
	if v337 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v356 = v349
	goto L77
L84:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if v350 != 0 {
		v335 = v350
		v337 = v348
		v338 = v349
		v340 = v340 + int32(1)
		goto L82
	} else {
		goto L87
	}
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v335)
	v348 = v337 + int32(-1)
	v349 = v338 + int32(1)
	goto L84
L86:
	;
	v348 = int32(0)
	v349 = v338
	goto L84
L87:
	;
	goto L83
L88:
	;
	if l1 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L89:
	;
	v404 = v400 - v373
	if v375 == v404 {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	v386 = v373
	v388 = v375
	goto L91
L91:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v390 == int32(0) {
		v400 = v386
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v400 = v6 + int32(288)
	goto L89
L93:
	;
	v396 = v388 + int32(-1)
	if v396 != 0 {
		v386 = v386 + int32(1)
		v388 = v396
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v437 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v433))) = uint8(v437)
	goto L88
L96:
	;
	v412 = v406
	v414 = v375 + (v404 ^ int32(-1))
	v415 = v400
	v417 = v374
	goto L100
L97:
	;
	v407 = F_strlen(m, v374)
	mBase = m.M
	goto L88
L98:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v406 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v433 = v400
	goto L95
L100:
	;
	if v414 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v433 = v426
	goto L95
L102:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)))
	if v427 != 0 {
		v412 = v427
		v414 = v425
		v415 = v426
		v417 = v417 + int32(1)
		goto L100
	} else {
		goto L105
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v415))) = uint8(v412)
	v425 = v414 + int32(-1)
	v426 = v415 + int32(1)
	goto L102
L104:
	;
	v425 = int32(0)
	v426 = v415
	goto L102
L105:
	;
	goto L101
L106:
	;
	m.G0 = v6 + int32(544)
	return
L107:
	;
	v469 = F_bg_unlink(m, v6+int32(288))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v452 = v6 + int32(288)
	v453 = int32(2048)
	v454 = int32(0)
	v455 = F_open(m, v452, v453, v454)
	mBase = m.M
	v457 = v6 + int32(32)
	v460 = F_open(m, v457, v453, v454)
	mBase = m.M
	v463 = F_unlink(m, v452)
	mBase = m.M
	v466 = F_unlink(m, v457)
	mBase = m.M
	goto L106
L109:
	;
	return
L110:
	;
	v473 = F_bg_unlink(m, v6+int32(32))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	goto L106
}
func F_aofRewriteLimited(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int64)(unsafe.Add(mBase, _consts[34]))
	if int64(2) < v10 {
		v21 = *(*int64)(unsafe.Add(mBase, _consts[35]))
		if v21 == int64(0) {
			v34 = int32(1)
			v35 = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, _consts[36]))
			v39 = v37 << (uint(v34) % 32)
			v40 = int32(60)
			if v39 < v40 {
				v43 = v39
			} else {
				v43 = v40
			}
			if v37 != 0 {
				v45 = v43
			} else {
				v45 = int32(1)
			}
			*(*int32)(unsafe.Add(mBase, _consts[36])) = v45
			v48 = int32(_a69)
			v49 = *(*int64)(unsafe.Add(mBase, _consts[37]))
			*(*int64)(unsafe.Add(mBase, _consts[35])) = v49 + base.I64_extend_i32_s(v45*int32(60))
			v56 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v56 {
				v66 = v34
				m.G0 = v7 + int32(16)
				return v66
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v45
				F__serverLog(m, int32(3), int32(_a92), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					v66 = v34
					m.G0 = v7 + int32(16)
					return v66
				}
			}
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, _consts[37]))
			v28 = *(*int64)(unsafe.Add(mBase, _consts[35]))
			if v26 < v28 {
				v66 = int32(1)
			} else {
				v30 = int32(0)
				*(*int64)(unsafe.Add(mBase, _consts[35])) = int64(0)
				v66 = v30
			}
			m.G0 = v7 + int32(16)
			return v66
		}
	} else {
		v13 = int32(0)
		*(*int64)(unsafe.Add(mBase, _consts[35])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _consts[36])) = v13
		v66 = v13
		m.G0 = v7 + int32(16)
		return v66
	}
}
func F_writeAofManifestFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v14 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+116)) = int32(_a70)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v21
	v26 = F_sdscatprintf(m, v14, int32(_a71), v12+int32(112))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v30 = F_makePath(m, v29, v26)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = F_sdsempty(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(_a78)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(_a70)
	v39 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v39
	v44 = F_sdscatprintf(m, v32, int32(_a79), v12+int32(96))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v48 = F_makePath(m, v47, v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(438)
	v52 = int32(-1)
	v56 = F_open(m, v48, int32(577), v12+int32(80))
	mBase = m.M
	if v56 != v52 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_sdsfree(m, v26)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L65
	}
L9:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v74 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L16
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v60 {
		v262 = v52
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L12
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v65 = F___strerror_l(m, v64, v64)
	mBase = m.M
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v44
	F__serverLog(m, int32(3), int32(_a80), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v262 = v52
	goto L8
L15:
	;
	v254 = F_close(m, v56)
	mBase = m.M
	v262 = v252
	goto L8
L16:
	;
	v135 = int32(-1)
	v136 = F_fsync(m, v56)
	mBase = m.M
	if v136 != v135 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	if v91 == int32(0) {
		goto L16
	} else {
		goto L23
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v91 = v90
	goto L17
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v91 = v87
	goto L17
L20:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v91 = v84
	goto L17
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v91 = v81
	goto L17
L22:
	;
	v91 = int32(base.Ui32(v74) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	v94 = l0
	v101 = v91
	goto L24
L24:
	;
	v103 = F_write(m, v56, v94, v101)
	mBase = m.M
	if int32(-1) < v103 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L16
L26:
	;
	v125 = v101 - v103
	if v125 != 0 {
		v94 = v94 + v103
		v101 = v125
		goto L24
	} else {
		goto L33
	}
L27:
	;
	goto L28
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	if v107 == int32(27) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v110 = int32(-1)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v112 {
		v252 = v110
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v115 = F___strerror_l(m, v107, v107)
	mBase = m.M
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v44
	F__serverLog(m, int32(3), int32(_a81), v12+int32(64))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v252 = v110
	goto L15
L33:
	;
	goto L25
L34:
	;
	v154 = F_rename(m, v48, v30)
	mBase = m.M
	if v154 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v140 {
		v252 = v135
		goto L15
	} else {
		goto L36
	}
L36:
	;
	goto L37
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v145 = F___strerror_l(m, v144, v144)
	mBase = m.M
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v44
	F__serverLog(m, int32(3), int32(_a82), v12+int32(16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v252 = v135
	goto L15
L40:
	;
	v178 = m.G0
	v180 = v178 - int32(4112)
	m.G0 = v180
	v182 = F_strlen(m, v30)
	mBase = m.M
	if base.Ui32(v182) < base.Ui32(int32(4097)) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v157 = int32(-1)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v159 {
		v252 = v157
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L43
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v164 = F___strerror_l(m, v163, v163)
	mBase = m.M
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v44
	F__serverLog(m, int32(3), int32(_a83), v12+int32(48))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v252 = v157
	goto L15
L46:
	;
	if v221 != int32(-1) {
		v252 = int32(0)
		goto L15
	} else {
		goto L59
	}
L47:
	;
	m.G0 = v180 + int32(4112)
	goto L46
L48:
	;
	v191 = F___memcpy(m, v180, v30, v182+int32(1))
	mBase = m.M
	v192 = F_dirname(m, v191)
	mBase = m.M
	v193 = int32(0)
	v195 = F_open(m, v192, v193, v193)
	mBase = m.M
	if v195 != int32(-1) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v185 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(37)
	v221 = int32(-1)
	goto L47
L50:
	;
	v205 = F_fsync(m, v195)
	mBase = m.M
	if v205 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v200 = F___errno_location(m)
	mBase = m.M
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	if v201 != int32(31) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v204 = int32(-1)
	goto L54
L53:
	;
	v204 = int32(0)
	goto L54
L54:
	;
	v221 = v204
	goto L47
L55:
	;
	v219 = F_close(m, v195)
	mBase = m.M
	v221 = int32(0)
	goto L47
L56:
	;
	v208 = F___errno_location(m)
	mBase = m.M
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v209 == int32(8) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if v209 == int32(28) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v214 = F_close(m, v195)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v209
	v221 = int32(-1)
	goto L47
L59:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v230 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v252 = int32(-1)
	goto L15
L61:
	;
	goto L62
L62:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v235 = F___strerror_l(m, v234, v234)
	mBase = m.M
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v30
	F__serverLog(m, int32(3), int32(_a84), v12+int32(32))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	F_sdsfree(m, v30)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_sdsfree(m, v44)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_sdsfree(m, v48)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	m.G0 = v12 + int32(128)
	return v262
}
