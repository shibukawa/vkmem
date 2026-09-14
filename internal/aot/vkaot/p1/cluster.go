package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clusterAddNodeToShard(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	v5 = F_sdsnewlen(m, l0, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
		v10 = F_dictFind(m, v9, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v30 = F_listSearchKey(m, v29, l1)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					if v30 != 0 {
						F_sdsfree(m, v5)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							return
						}
					} else {
						v32 = F_listAddNodeTail(m, v29, l1)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_sdsfree(m, v5)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v12 = F_listCreate(m)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v14 = F_listAddNodeTail(m, v12, l1)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
						v19 = F_dictAdd(m, v18, v5, v12)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							if v19 == int32(0) {
								return
							} else {
								F__serverAssert(m, int32(_a246), int32(_a247), int32(2313))
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
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
	}
}
func F_clusterBlacklistAddNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	v5 = F_sdsnewlen(m, l0+int32(8), int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_clusterBlacklistCleanup(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[111]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
			v13 = F_dictAdd(m, v11, v5, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				if v13 != 0 {
					v17 = v5
					v19 = *(*int32)(unsafe.Add(mBase, _consts[111]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
					v21 = F_dictFind(m, v20, v17)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = F___time(m, int32(0))
						mBase = m.M
						v26 = int64(*(*uint32)(unsafe.Add(mBase, _consts[169])))
						*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v24 + v26
						F_sdsfree(m, v17)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v15 = F_sdsdup(m, v5)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						v17 = v15
						v19 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						v21 = F_dictFind(m, v20, v17)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							v24 = F___time(m, int32(0))
							mBase = m.M
							v26 = int64(*(*uint32)(unsafe.Add(mBase, _consts[169])))
							*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v24 + v26
							F_sdsfree(m, v17)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
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
func F_clusterBlacklistCleanup(m *base.Module) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
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
	var v130 int32
	_ = v130
	v4 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+40))
	v6 = F_dictGetSafeIterator(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v6)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L4:
	;
	v17 = v6 + int32(20)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v113 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v24 = v17
	v25 = v21
	goto L10
L8:
	;
	v21 = int32(1)
	goto L7
L9:
	;
	v21 = int32(0)
	goto L7
L10:
	;
	switch v25 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v25 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v105
	if v105 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v29 != int32(-1) {
		v68 = v29
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = int32(1)
	v70 = v68 + v69
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v70
	v72 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v76+int32(26)))))
	if v80 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v33 != 0 {
		v68 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v35 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	if v62 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+16)))
	v43 = int64(*(*int8)(unsafe.Add(mBase, uint32(v34)+27)))
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+8)))
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+12)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v34)+26)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+4)))
	v48 = F_wangHash64(m, v47)
	mBase = m.M
	v50 = F_wangHash64(m, v46+v48)
	mBase = m.M
	v52 = F_wangHash64(m, v45+v50)
	mBase = m.M
	v54 = F_wangHash64(m, v44+v52)
	mBase = m.M
	v56 = F_wangHash64(m, v43+v54)
	mBase = m.M
	v58 = F_wangHash64(m, v42+v56)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v61 = v60
	goto L19
L21:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)))
	v40 = v38 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)) = uint16(v40)
	v61 = v34
	goto L19
L22:
	;
	v68 = v62 + int32(-1)
	goto L16
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v68 = v65
	goto L16
L24:
	;
	v95 = int32(2)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v75+v93<<(uint(v95)%32)+int32(4))))
	v24 = v100 + v94<<(uint(v95)%32)
	v25 = int32(1)
	goto L10
L25:
	;
	v84 = v72
	goto L27
L26:
	;
	v84 = v69 << (uint(v80) % 32)
	goto L27
L27:
	;
	if v70 < v84 {
		v93 = v76
		v94 = v70
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v76 != 0 {
		v113 = v72
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v86 == int32(-1) {
		v113 = v72
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(4294967296)
	v93 = int32(1)
	v94 = int32(0)
	goto L24
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v109
	v113 = v105
	goto L13
L32:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v113)+8))
	goto L33
L33:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	if v121 <= v119 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+40))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	goto L35
L35:
	;
	v127 = F_dictDelete(m, v125, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	return
}
func F_clusterBlacklistExists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = F_sdsnewlen(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_clusterBlacklistCleanup(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[111]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
			v12 = F_dictFind(m, v11, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_sdsfree(m, v3)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v12 != int32(0))
				}
			}
		}
	}
}
func F_clusterBumpConfigEpochWithoutConsensus(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v34 int32
	_ = v34
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v16 = F_dictGetSafeIterator(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int64(0)
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v16)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L4:
	;
	v34 = v16 + int32(20)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v130 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v41 = v34
	v42 = v38
	goto L10
L8:
	;
	v38 = int32(1)
	goto L7
L9:
	;
	v38 = int32(0)
	goto L7
L10:
	;
	switch v42 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v42 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v122
	if v122 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v46 != int32(-1) {
		v85 = v46
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v86 = int32(1)
	v87 = v85 + v86
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v87
	v89 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93+int32(26)))))
	if v97 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v50 != 0 {
		v85 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v52 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v79 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+16)))
	v60 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+27)))
	v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+8)))
	v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+12)))
	v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v51)+26)))
	v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
	v65 = F_wangHash64(m, v64)
	mBase = m.M
	v67 = F_wangHash64(m, v63+v65)
	mBase = m.M
	v69 = F_wangHash64(m, v62+v67)
	mBase = m.M
	v71 = F_wangHash64(m, v61+v69)
	mBase = m.M
	v73 = F_wangHash64(m, v60+v71)
	mBase = m.M
	v75 = F_wangHash64(m, v59+v73)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v78 = v77
	goto L19
L21:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)))
	v57 = v55 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+24)) = uint16(v57)
	v78 = v51
	goto L19
L22:
	;
	v85 = v79 + int32(-1)
	goto L16
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v85 = v82
	goto L16
L24:
	;
	v112 = int32(2)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v92+v110<<(uint(v112)%32)+int32(4))))
	v41 = v117 + v111<<(uint(v112)%32)
	v42 = int32(1)
	goto L10
L25:
	;
	v101 = v89
	goto L27
L26:
	;
	v101 = v86 << (uint(v97) % 32)
	goto L27
L27:
	;
	if v87 < v101 {
		v110 = v93
		v111 = v87
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v93 != 0 {
		v130 = v89
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	if v103 == int32(-1) {
		v130 = v89
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4294967296)
	v110 = int32(1)
	v111 = int32(0)
	goto L24
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v126
	v130 = v122
	goto L13
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	goto L33
L33:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v136)+96))
	if base.Ui64(v21) < base.Ui64(v137) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v139 = v137
	goto L36
L35:
	;
	v139 = v21
	goto L36
L36:
	;
	v21 = v139
	goto L4
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
	v146 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v146)+96))
	if v147 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	m.G0 = v10 + int32(16)
	return v182
L39:
	;
	v156 = v144 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v143)+8)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v146)+96)) = v156
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L40:
	;
	if base.Ui64(v144) < base.Ui64(v21) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v152 = v21
	goto L43
L42:
	;
	v152 = v144
	goto L43
L43:
	;
	if v147 == v152 {
		v182 = int32(-1)
		goto L38
	} else {
		goto L44
	}
L44:
	;
	goto L39
L45:
	;
	v161 = int32(_a20)
	v162 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+uint32(_consts[122]))) = v163 | int32(44)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v169 {
		v182 = int32(0)
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v172 = int32(0)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v174)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v175
	F__serverLog(m, int32(2), int32(_a288), v10)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v182 = v172
	goto L38
}
func F_clusterCleanSlotImportsBeforeLoad(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v76 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[171])))
	v18 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L4
L4:
	;
	v24 = v6 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v29 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v35
	goto L6
L8:
	;
	v40 = v26
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v59 = v6 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+156))
	if base.Ui32(int32(20)) < base.Ui32(v46) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_finishSlotMigrationJob(m, v42, int32(18), int32(_a391))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if int32(1)<<(uint(v46)%32)&int32(1835040) != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	goto L11
L18:
	;
	if v61 != 0 {
		v40 = v61
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+base.B2i32(v64 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v70
	goto L19
L21:
	;
	goto L10
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v85 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F__serverLog(m, int32(2), int32(_a392), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	m.G0 = v6 + int32(16)
	return
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v89 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_consts[171])))
	v94 = v6 + int32(8)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95
	goto L28
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_consts[171])))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	if v102 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L30
L30:
	;
	v109 = v6 + int32(8)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v111 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L25
L32:
	;
	if v111 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111+base.B2i32(v114 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v120
	goto L33
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+156))
	if base.Ui32(int32(2)) < base.Ui32(v125+int32(-18)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[171])))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v138 != 0 {
		goto L30
	} else {
		goto L39
	}
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[171])))
	F_listDelNode(m, v132, v111)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L31
}
func F_clusterCleanSlotImportsOnFullSync(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[171])))
	v18 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L4
L4:
	;
	v24 = v6 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v29 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v35
	goto L6
L8:
	;
	v40 = v26
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v59 = v6 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+156))
	if base.Ui32(int32(20)) < base.Ui32(v46) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_finishSlotMigrationJob(m, v42, int32(18), int32(_a390))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if int32(1)<<(uint(v46)%32)&int32(1835040) != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	goto L11
L18:
	;
	if v61 != 0 {
		v40 = v61
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+base.B2i32(v64 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v70
	goto L19
L21:
	;
	goto L10
}
func F_clusterCleanSlotImportsOnPromotion(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+uint32(_consts[171])))
	v18 = v6 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
	goto L4
L4:
	;
	v24 = v6 + int32(8)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+base.B2i32(v29 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v35
	goto L6
L8:
	;
	v40 = v26
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L1
L11:
	;
	v59 = v6 + int32(8)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+156))
	if base.Ui32(int32(20)) < base.Ui32(v46) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_finishSlotMigrationJob(m, v42, int32(18), int32(_a389))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if int32(1)<<(uint(v46)%32)&int32(1835040) != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return
L17:
	;
	goto L11
L18:
	;
	if v61 != 0 {
		v40 = v61
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61+base.B2i32(v64 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v70
	goto L19
L21:
	;
	goto L10
}
func F_clusterCloseAllSlots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+44))
	F_dictEmpty(m, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		F_dictEmpty(m, v9, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
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
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int64
	_ = v583
	var v586 int64
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int64
	_ = v603
	var v604 int64
	_ = v604
	var v606 int64
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
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
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
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
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v899 int32
	_ = v899
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a208), int32(_a203), int32(941))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L292
	}
L2:
	;
	m.G0 = v13 + int32(128)
	return
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v20 != int32(2) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_addReplyError(m, l0, int32(_a209))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L2
L7:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = F_objectGetVal(m, v79)
	mBase = m.M
	v81 = int32(_a210)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v84 != 0 {
		goto L30
	} else {
		goto L31
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = F_objectGetVal(m, v24)
	mBase = m.M
	v26 = int32(_a98)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v61-v63 != 0 {
		goto L7
	} else {
		goto L21
	}
L10:
	;
	v61 = F_tolower(m, v57)
	mBase = m.M
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v63 = F_tolower(m, v62)
	mBase = m.M
	goto L9
L11:
	;
	v31 = v25
	v32 = v26
	v33 = v29
	goto L14
L12:
	;
	v57 = int32(0)
	v58 = v26
	goto L10
L13:
	;
	v57 = v54 & int32(255)
	v58 = v53
	goto L10
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v35 == int32(0) {
		v53 = v32
		v54 = v33
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v53 = v47
	v54 = int32(0)
	goto L13
L16:
	;
	v39 = v33 & int32(255)
	if v39 == v35 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = int32(1)
	v47 = v32 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v48 != 0 {
		v31 = v31 + v46
		v32 = v47
		v33 = v48
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v41 = F_tolower(m, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v43 = F_tolower(m, v42)
	mBase = m.M
	if v41 == v43 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v53 = v32
	v54 = v45
	goto L13
L20:
	;
	goto L15
L21:
	;
	goto L24
L22:
	;
	goto L25
L23:
	;
	goto L22
L24:
	;
	v71 = F__emscripten_memcpy_bulkmem(m, v13+int32(16), int32(_a211), int32(100))
	mBase = m.M
	goto L23
L25:
	;
	F_addExtendedReplyHelp(m, l0, v13+int32(16), int32(_a212))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	goto L2
L27:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v169 = F_objectGetVal(m, v168)
	mBase = m.M
	v170 = int32(_a213)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v173 != 0 {
		goto L59
	} else {
		goto L60
	}
L28:
	;
	if v116-v118 != 0 {
		goto L27
	} else {
		goto L40
	}
L29:
	;
	v116 = F_tolower(m, v112)
	mBase = m.M
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v118 = F_tolower(m, v117)
	mBase = m.M
	goto L28
L30:
	;
	v86 = v80
	v87 = v81
	v88 = v84
	goto L33
L31:
	;
	v112 = int32(0)
	v113 = v81
	goto L29
L32:
	;
	v112 = v109 & int32(255)
	v113 = v108
	goto L29
L33:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v90 == int32(0) {
		v108 = v87
		v109 = v88
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v108 = v102
	v109 = int32(0)
	goto L32
L35:
	;
	v94 = v88 & int32(255)
	if v94 == v90 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v101 = int32(1)
	v102 = v87 + v101
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v103 != 0 {
		v86 = v86 + v101
		v87 = v102
		v88 = v103
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v96 = F_tolower(m, v94)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v98 = F_tolower(m, v97)
	mBase = m.M
	if v96 == v98 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v108 = v87
	v109 = v100
	goto L32
L39:
	;
	goto L34
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v120 != int32(2) {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v124 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v138 = int32(0)
	v140 = F_clusterGenNodesDescription(m, l0, v138, v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L53
	}
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	v137 = v136
	goto L42
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v127 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v131 = F_connectionTypeTls(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v137 = base.B2i32(v130 == v131)
	goto L42
L47:
	;
	F_addReplyVerbatim(m, l0, v140, v161, int32(_a214))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L54
	}
L48:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-17))))
	v161 = v160
	goto L47
L49:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-9))))
	v161 = v157
	goto L47
L50:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+int32(-5)))))
	v161 = v154
	goto L47
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-3)))))
	v161 = v151
	goto L47
L52:
	;
	v161 = int32(base.Ui32(v144) >> (uint(int32(3)) % 32))
	goto L47
L53:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-1)))))
	switch v144 & int32(7) {
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
		v161 = v138
		goto L47
	}
L54:
	;
	F_sdsfree(m, v140)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	goto L2
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	v227 = F_objectGetVal(m, v226)
	mBase = m.M
	v228 = int32(_a215)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v231 != 0 {
		goto L80
	} else {
		goto L81
	}
L57:
	;
	if v205-v207 != 0 {
		goto L56
	} else {
		goto L69
	}
L58:
	;
	v205 = F_tolower(m, v201)
	mBase = m.M
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v207 = F_tolower(m, v206)
	mBase = m.M
	goto L57
L59:
	;
	v175 = v169
	v176 = v170
	v177 = v173
	goto L62
L60:
	;
	v201 = int32(0)
	v202 = v170
	goto L58
L61:
	;
	v201 = v198 & int32(255)
	v202 = v197
	goto L58
L62:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v179 == int32(0) {
		v197 = v176
		v198 = v177
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v197 = v191
	v198 = int32(0)
	goto L61
L64:
	;
	v183 = v177 & int32(255)
	if v183 == v179 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v190 = int32(1)
	v191 = v176 + v190
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v192 != 0 {
		v175 = v175 + v190
		v176 = v191
		v177 = v192
		goto L62
	} else {
		goto L68
	}
L66:
	;
	v185 = F_tolower(m, v183)
	mBase = m.M
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v187 = F_tolower(m, v186)
	mBase = m.M
	if v185 == v187 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v197 = v176
	v198 = v189
	goto L61
L68:
	;
	goto L63
L69:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v209 != int32(2) {
		goto L56
	} else {
		goto L70
	}
L70:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	goto L72
L71:
	;
	F_addReplyError(m, l0, int32(_a216))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L76
	}
L72:
	;
	v216 = v214 + int32(8)
	goto L73
L73:
	;
	if v216 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	F_addReplyBulkCBuffer(m, l0, v216, int32(40))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	goto L2
L76:
	;
	goto L2
L77:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v285 = F_objectGetVal(m, v284)
	mBase = m.M
	v286 = int32(_a217)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v289 != 0 {
		goto L101
	} else {
		goto L102
	}
L78:
	;
	if v263-v265 != 0 {
		goto L77
	} else {
		goto L90
	}
L79:
	;
	v263 = F_tolower(m, v259)
	mBase = m.M
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v265 = F_tolower(m, v264)
	mBase = m.M
	goto L78
L80:
	;
	v233 = v227
	v234 = v228
	v235 = v231
	goto L83
L81:
	;
	v259 = int32(0)
	v260 = v228
	goto L79
L82:
	;
	v259 = v256 & int32(255)
	v260 = v255
	goto L79
L83:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v237 == int32(0) {
		v255 = v234
		v256 = v235
		goto L82
	} else {
		goto L85
	}
L84:
	;
	v255 = v249
	v256 = int32(0)
	goto L82
L85:
	;
	v241 = v235 & int32(255)
	if v241 == v237 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v248 = int32(1)
	v249 = v234 + v248
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if v250 != 0 {
		v233 = v233 + v248
		v234 = v249
		v235 = v250
		goto L83
	} else {
		goto L89
	}
L87:
	;
	v243 = F_tolower(m, v241)
	mBase = m.M
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	v245 = F_tolower(m, v244)
	mBase = m.M
	if v243 == v245 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v255 = v234
	v256 = v247
	goto L82
L89:
	;
	goto L84
L90:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v267 != int32(2) {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	goto L93
L92:
	;
	F_addReplyError(m, l0, int32(_a218))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L97
	}
L93:
	;
	v274 = v272 + int32(48)
	goto L94
L94:
	;
	if v274 == int32(0) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	F_addReplyBulkCBuffer(m, l0, v274, int32(40))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	goto L2
L97:
	;
	goto L2
L98:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v332 = F_objectGetVal(m, v331)
	mBase = m.M
	v333 = int32(_a219)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v336 != 0 {
		goto L117
	} else {
		goto L118
	}
L99:
	;
	if v321-v323 != 0 {
		goto L98
	} else {
		goto L111
	}
L100:
	;
	v321 = F_tolower(m, v317)
	mBase = m.M
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v323 = F_tolower(m, v322)
	mBase = m.M
	goto L99
L101:
	;
	v291 = v285
	v292 = v286
	v293 = v289
	goto L104
L102:
	;
	v317 = int32(0)
	v318 = v286
	goto L100
L103:
	;
	v317 = v314 & int32(255)
	v318 = v313
	goto L100
L104:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v295 == int32(0) {
		v313 = v292
		v314 = v293
		goto L103
	} else {
		goto L106
	}
L105:
	;
	v313 = v307
	v314 = int32(0)
	goto L103
L106:
	;
	v299 = v293 & int32(255)
	if v299 == v295 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v306 = int32(1)
	v307 = v292 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v308 != 0 {
		v291 = v291 + v306
		v292 = v307
		v293 = v308
		goto L104
	} else {
		goto L110
	}
L108:
	;
	v301 = F_tolower(m, v299)
	mBase = m.M
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	v303 = F_tolower(m, v302)
	mBase = m.M
	if v301 == v303 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	v313 = v292
	v314 = v305
	goto L103
L110:
	;
	goto L105
L111:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v325 != int32(2) {
		goto L98
	} else {
		goto L112
	}
L112:
	;
	F_clusterCommandSlots(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	goto L2
L114:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	v379 = F_objectGetVal(m, v378)
	mBase = m.M
	v380 = int32(_a220)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v383 != 0 {
		goto L133
	} else {
		goto L134
	}
L115:
	;
	if v368-v370 != 0 {
		goto L114
	} else {
		goto L127
	}
L116:
	;
	v368 = F_tolower(m, v364)
	mBase = m.M
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v370 = F_tolower(m, v369)
	mBase = m.M
	goto L115
L117:
	;
	v338 = v332
	v339 = v333
	v340 = v336
	goto L120
L118:
	;
	v364 = int32(0)
	v365 = v333
	goto L116
L119:
	;
	v364 = v361 & int32(255)
	v365 = v360
	goto L116
L120:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v342 == int32(0) {
		v360 = v339
		v361 = v340
		goto L119
	} else {
		goto L122
	}
L121:
	;
	v360 = v354
	v361 = int32(0)
	goto L119
L122:
	;
	v346 = v340 & int32(255)
	if v346 == v342 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v353 = int32(1)
	v354 = v339 + v353
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+1)))
	if v355 != 0 {
		v338 = v338 + v353
		v339 = v354
		v340 = v355
		goto L120
	} else {
		goto L126
	}
L124:
	;
	v348 = F_tolower(m, v346)
	mBase = m.M
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v350 = F_tolower(m, v349)
	mBase = m.M
	if v348 == v350 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v360 = v339
	v361 = v352
	goto L119
L126:
	;
	goto L121
L127:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v372 != int32(2) {
		goto L114
	} else {
		goto L128
	}
L128:
	;
	F_clusterCommandShards(m, l0)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	goto L2
L130:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v454 = F_objectGetVal(m, v453)
	mBase = m.M
	v455 = int32(_a221)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v458 != 0 {
		goto L158
	} else {
		goto L159
	}
L131:
	;
	if v415-v417 != 0 {
		goto L130
	} else {
		goto L143
	}
L132:
	;
	v415 = F_tolower(m, v411)
	mBase = m.M
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	v417 = F_tolower(m, v416)
	mBase = m.M
	goto L131
L133:
	;
	v385 = v379
	v386 = v380
	v387 = v383
	goto L136
L134:
	;
	v411 = int32(0)
	v412 = v380
	goto L132
L135:
	;
	v411 = v408 & int32(255)
	v412 = v407
	goto L132
L136:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v389 == int32(0) {
		v407 = v386
		v408 = v387
		goto L135
	} else {
		goto L138
	}
L137:
	;
	v407 = v401
	v408 = int32(0)
	goto L135
L138:
	;
	v393 = v387 & int32(255)
	if v393 == v389 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v400 = int32(1)
	v401 = v386 + v400
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+1)))
	if v402 != 0 {
		v385 = v385 + v400
		v386 = v401
		v387 = v402
		goto L136
	} else {
		goto L142
	}
L140:
	;
	v395 = F_tolower(m, v393)
	mBase = m.M
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	v397 = F_tolower(m, v396)
	mBase = m.M
	if v395 == v397 {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v407 = v386
	v408 = v399
	goto L135
L142:
	;
	goto L137
L143:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v419 != int32(2) {
		goto L130
	} else {
		goto L144
	}
L144:
	;
	v423 = F_sdsempty(m)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L151
	}
L145:
	;
	F_addReplyVerbatim(m, l0, v425, v446, int32(_a214))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L153
	}
L146:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v425+int32(-17))))
	v446 = v445
	goto L145
L147:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v425+int32(-9))))
	v446 = v442
	goto L145
L148:
	;
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v425+int32(-5)))))
	v446 = v439
	goto L145
L149:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+int32(-3)))))
	v446 = v436
	goto L145
L150:
	;
	v446 = int32(base.Ui32(v429) >> (uint(int32(3)) % 32))
	goto L145
L151:
	;
	v425 = F_genClusterInfoString(m, v423)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+int32(-1)))))
	switch v429 & int32(7) {
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
		v446 = int32(0)
		goto L145
	}
L153:
	;
	F_sdsfree(m, v425)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	goto L2
L155:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v526 = F_objectGetVal(m, v525)
	mBase = m.M
	v527 = int32(_a222)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if v530 != 0 {
		goto L182
	} else {
		goto L183
	}
L156:
	;
	if v490-v492 != 0 {
		goto L155
	} else {
		goto L168
	}
L157:
	;
	v490 = F_tolower(m, v486)
	mBase = m.M
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	v492 = F_tolower(m, v491)
	mBase = m.M
	goto L156
L158:
	;
	v460 = v454
	v461 = v455
	v462 = v458
	goto L161
L159:
	;
	v486 = int32(0)
	v487 = v455
	goto L157
L160:
	;
	v486 = v483 & int32(255)
	v487 = v482
	goto L157
L161:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v464 == int32(0) {
		v482 = v461
		v483 = v462
		goto L160
	} else {
		goto L163
	}
L162:
	;
	v482 = v476
	v483 = int32(0)
	goto L160
L163:
	;
	v468 = v462 & int32(255)
	if v468 == v464 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v475 = int32(1)
	v476 = v461 + v475
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	if v477 != 0 {
		v460 = v460 + v475
		v461 = v476
		v462 = v477
		goto L161
	} else {
		goto L167
	}
L165:
	;
	v470 = F_tolower(m, v468)
	mBase = m.M
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v472 = F_tolower(m, v471)
	mBase = m.M
	if v470 == v472 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v482 = v461
	v483 = v474
	goto L160
L167:
	;
	goto L162
L168:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v494 != int32(3) {
		goto L155
	} else {
		goto L169
	}
L169:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+8))
	v502 = F_getLongLongFromObjectOrReply(m, l0, v498, v13+int32(16), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L170
	}
L170:
	;
	if v502 != 0 {
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if base.Ui64(v504) < base.Ui64(int64(16384)) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v511)+8))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513+base.I32_wrap_i64(v504)<<(uint(int32(2))%32))))
	if v517 != 0 {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	F_addReplyError(m, l0, int32(_a223))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	goto L2
L175:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v520))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L5
	} else {
		goto L178
	}
L176:
	;
	v519 = F_hashtableSize(m, v517)
	mBase = m.M
	v520 = v519
	goto L175
L177:
	;
	v520 = int32(0)
	goto L175
L178:
	;
	goto L2
L179:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	v700 = F_objectGetVal(m, v699)
	mBase = m.M
	v701 = int32(_a224)
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700))))
	if v704 != 0 {
		goto L233
	} else {
		goto L234
	}
L180:
	;
	if v562-v564 != 0 {
		goto L179
	} else {
		goto L192
	}
L181:
	;
	v562 = F_tolower(m, v558)
	mBase = m.M
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559))))
	v564 = F_tolower(m, v563)
	mBase = m.M
	goto L180
L182:
	;
	v532 = v526
	v533 = v527
	v534 = v530
	goto L185
L183:
	;
	v558 = int32(0)
	v559 = v527
	goto L181
L184:
	;
	v558 = v555 & int32(255)
	v559 = v554
	goto L181
L185:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v536 == int32(0) {
		v554 = v533
		v555 = v534
		goto L184
	} else {
		goto L187
	}
L186:
	;
	v554 = v548
	v555 = int32(0)
	goto L184
L187:
	;
	v540 = v534 & int32(255)
	if v540 == v536 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v547 = int32(1)
	v548 = v533 + v547
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+1)))
	if v549 != 0 {
		v532 = v532 + v547
		v533 = v548
		v534 = v549
		goto L185
	} else {
		goto L191
	}
L189:
	;
	v542 = F_tolower(m, v540)
	mBase = m.M
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	v544 = F_tolower(m, v543)
	mBase = m.M
	if v542 == v544 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v554 = v533
	v555 = v546
	goto L184
L191:
	;
	goto L186
L192:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v566 != int32(4) {
		goto L179
	} else {
		goto L193
	}
L193:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+8))
	v574 = F_getLongLongFromObjectOrReply(m, l0, v570, v13+int32(8), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	if v574 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	v581 = F_getLongLongFromObjectOrReply(m, l0, v577, v13+int32(16), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	if v581 != 0 {
		goto L2
	} else {
		goto L197
	}
L197:
	;
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if base.Ui64(int64(16383)) < base.Ui64(v583) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595+base.I32_wrap_i64(v583)<<(uint(int32(2))%32))))
	if v599 != 0 {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	F_addReplyError(m, l0, int32(_a225))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L5
	} else {
		goto L202
	}
L200:
	;
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	if int64(-1) < v586 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	goto L2
L203:
	;
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
	v604 = base.I64_extend_i32_u(v602)
	if v603 < v604 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v601 = F_hashtableSize(m, v599)
	mBase = m.M
	v602 = v601
	goto L203
L205:
	;
	v602 = int32(0)
	goto L203
L206:
	;
	v606 = v603
	goto L208
L207:
	;
	v606 = v604
	goto L208
L208:
	;
	v607 = base.I32_wrap_i64(v606)
	F_addReplyArrayLen(m, l0, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v610 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v615 = F_kvstoreGetHashtableIterator(m, v612, v613, v610)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	if v607 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	F_kvstoreReleaseHashtableIterator(m, v615)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L228
	}
L212:
	;
	v623 = v610
	goto L213
L213:
	;
	v631 = F_kvstoreHashtableIteratorNext(m, v615, v13+int32(4))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L215
	}
L214:
	;
	goto L211
L215:
	;
	if v631 == int32(0) {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v635 = int32(0)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v640&int32(2) == v635 {
		v660 = v635
		goto L224
	} else {
		goto L225
	}
L217:
	;
	F_addReplyBulkCBuffer(m, l0, v660, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L226
	}
L218:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v660+int32(-17))))
	v680 = v679
	goto L217
L219:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v660+int32(-9))))
	v680 = v676
	goto L217
L220:
	;
	v673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v660+int32(-5)))))
	v680 = v673
	goto L217
L221:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660+int32(-3)))))
	v680 = v670
	goto L217
L222:
	;
	v680 = int32(base.Ui32(v663) >> (uint(int32(3)) % 32))
	goto L217
L223:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660+int32(-1)))))
	switch v663 & int32(7) {
	case 0:
		goto L222
	case 1:
		goto L221
	case 2:
		goto L220
	case 3:
		goto L219
	case 4:
		goto L218
	default:
		v680 = v635
		goto L217
	}
L224:
	;
	goto L223
L225:
	;
	v654 = v636 + (v640&int32(4) ^ int32(12)) + v640<<(uint(int32(3))%32)&int32(8)
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	v660 = v654 + v655 + int32(1)
	goto L224
L226:
	;
	v684 = v623 + int32(1)
	if v684 != v607 {
		v623 = v684
		goto L213
	} else {
		goto L227
	}
L227:
	;
	goto L214
L228:
	;
	goto L2
L229:
	;
	v878 = F_clusterCommandSpecial(m, l0)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L5
	} else {
		goto L289
	}
L230:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v784 != int32(3) {
		goto L229
	} else {
		goto L257
	}
L231:
	;
	if v736-v738 == int32(0) {
		goto L230
	} else {
		goto L243
	}
L232:
	;
	v736 = F_tolower(m, v732)
	mBase = m.M
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733))))
	v738 = F_tolower(m, v737)
	mBase = m.M
	goto L231
L233:
	;
	v706 = v700
	v707 = v701
	v708 = v704
	goto L236
L234:
	;
	v732 = int32(0)
	v733 = v701
	goto L232
L235:
	;
	v732 = v729 & int32(255)
	v733 = v728
	goto L232
L236:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	if v710 == int32(0) {
		v728 = v707
		v729 = v708
		goto L235
	} else {
		goto L238
	}
L237:
	;
	v728 = v722
	v729 = int32(0)
	goto L235
L238:
	;
	v714 = v708 & int32(255)
	if v714 == v710 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v721 = int32(1)
	v722 = v707 + v721
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
	if v723 != 0 {
		v706 = v706 + v721
		v707 = v722
		v708 = v723
		goto L236
	} else {
		goto L242
	}
L240:
	;
	v716 = F_tolower(m, v714)
	mBase = m.M
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	v718 = F_tolower(m, v717)
	mBase = m.M
	if v716 == v718 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	v728 = v707
	v729 = v720
	goto L235
L242:
	;
	goto L237
L243:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+4))
	v744 = F_objectGetVal(m, v743)
	mBase = m.M
	v745 = int32(_a226)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v748 != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	if v780-v782 != 0 {
		goto L229
	} else {
		goto L256
	}
L245:
	;
	v780 = F_tolower(m, v776)
	mBase = m.M
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	v782 = F_tolower(m, v781)
	mBase = m.M
	goto L244
L246:
	;
	v750 = v744
	v751 = v745
	v752 = v748
	goto L249
L247:
	;
	v776 = int32(0)
	v777 = v745
	goto L245
L248:
	;
	v776 = v773 & int32(255)
	v777 = v772
	goto L245
L249:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	if v754 == int32(0) {
		v772 = v751
		v773 = v752
		goto L248
	} else {
		goto L251
	}
L250:
	;
	v772 = v766
	v773 = int32(0)
	goto L248
L251:
	;
	v758 = v752 & int32(255)
	if v758 == v754 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v765 = int32(1)
	v766 = v751 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	if v767 != 0 {
		v750 = v750 + v765
		v751 = v766
		v752 = v767
		goto L249
	} else {
		goto L255
	}
L253:
	;
	v760 = F_tolower(m, v758)
	mBase = m.M
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	v762 = F_tolower(m, v761)
	mBase = m.M
	if v760 == v762 {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v772 = v751
	v773 = v764
	goto L248
L255:
	;
	goto L250
L256:
	;
	goto L230
L257:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)+8))
	v789 = F_objectGetVal(m, v788)
	mBase = m.M
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+8))
	v793 = F_objectGetVal(m, v792)
	mBase = m.M
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793+int32(-1)))))
	switch v796 & int32(7) {
	case 0:
		goto L263
	case 1:
		goto L262
	case 2:
		goto L261
	case 3:
		goto L260
	case 4:
		goto L259
	default:
		v813 = int32(0)
		goto L258
	}
L258:
	;
	v814 = F_clusterLookupNode(m, v789, v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L265
	}
L259:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v793+int32(-17))))
	v813 = v812
	goto L258
L260:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v793+int32(-9))))
	v813 = v809
	goto L258
L261:
	;
	v806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v793+int32(-5)))))
	v813 = v806
	goto L258
L262:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793+int32(-3)))))
	v813 = v803
	goto L258
L263:
	;
	v813 = int32(base.Ui32(v796) >> (uint(int32(3)) % 32))
	goto L258
L264:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v814)+88))
	goto L269
L265:
	;
	if v814 != 0 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+8))
	v818 = F_objectGetVal(m, v817)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v818
	F_addReplyErrorFormat(m, l0, int32(_a227), v13)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	goto L2
L268:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v814)+2164))
	goto L272
L269:
	;
	if v823&int32(2) == int32(0) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	F_addReplyError(m, l0, int32(_a228))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	goto L2
L272:
	;
	F_addReplyArrayLen(m, l0, v831)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v814)+2164))
	goto L274
L274:
	;
	if v834 < int32(1) {
		goto L2
	} else {
		goto L275
	}
L275:
	;
	v840 = int32(0)
	goto L276
L276:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v814)+2168))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v848+v840<<(uint(int32(2))%32))))
	goto L278
L278:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v854 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v868 = F_clusterGenNodeDescription(m, l0, v852, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L5
	} else {
		goto L284
	}
L280:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	v867 = v866
	goto L279
L281:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v854)+8))
	if v857 == int32(0) {
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v857)))
	v861 = F_connectionTypeTls(m)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L5
	} else {
		goto L283
	}
L283:
	;
	v867 = base.B2i32(v860 == v861)
	goto L279
L284:
	;
	F_addReplyBulkCString(m, l0, v868)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L285
	}
L285:
	;
	F_sdsfree(m, v868)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	v875 = v840 + int32(1)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v814)+2164))
	goto L287
L287:
	;
	if v875 < v876 {
		v840 = v875
		goto L276
	} else {
		goto L288
	}
L288:
	;
	goto L2
L289:
	;
	if v878 != 0 {
		goto L2
	} else {
		goto L290
	}
L290:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	goto L2
L292:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterCommandGetSlotMigrations(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[171])))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_addReplyArrayLen(m, l0, v14)
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[171])))
	v21 = v9 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	goto L3
L3:
	;
	v27 = v9 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return
L5:
	;
	if v29 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29+base.B2i32(v32 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v38
	goto L6
L8:
	;
	v44 = v29
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+196)))
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v52 = int32(10)
	goto L13
L12:
	;
	v52 = int32(12)
	goto L13
L13:
	;
	F_addReplyMapLen(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_addReplyBulkCString(m, l0, int32(_a373))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_addReplyBulkCBuffer(m, l0, v50+int32(112), int32(40))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_addReplyBulkCString(m, l0, int32(_a423))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v68 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = int32(_a424)
	goto L20
L19:
	;
	v71 = int32(_a425)
	goto L20
L20:
	;
	F_addReplyBulkCString(m, l0, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_addReplyBulkCString(m, l0, int32(_a426))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+168))
	F_addReplyBulkCString(m, l0, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+196)))
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_addReplyBulkCString(m, l0, int32(_a427))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	F_addReplyBulkCString(m, l0, int32(_a428))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_addReplyBulkCBuffer(m, l0, v50+int32(32), int32(40))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_addReplyBulkCString(m, l0, int32(_a429))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_addReplyBulkCBuffer(m, l0, v50+int32(72), int32(40))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	F_addReplyLongLong(m, l0, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_addReplyBulkCString(m, l0, int32(_a430))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v50)+16))
	F_addReplyLongLong(m, l0, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_addReplyBulkCString(m, l0, int32(_a431))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
	F_addReplyLongLong(m, l0, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_addReplyBulkCString(m, l0, int32(_a432))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v50)+156))
	if base.Ui32(int32(20)) < base.Ui32(v119) {
		v127 = int32(_a86)
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_addReplyBulkCString(m, l0, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119<<(uint(int32(2))%32))+uint32(_consts[213])))
	v127 = v126
	goto L37
L39:
	;
	F_addReplyBulkCString(m, l0, int32(_a433))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v50)+160))
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v135 = v133
	goto L43
L42:
	;
	v135 = int32(_a320)
	goto L43
L43:
	;
	F_addReplyBulkCString(m, l0, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_addReplyBulkCString(m, l0, int32(_a434))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v141 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v50)+192)))
	F_addReplyLongLong(m, l0, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_addReplyBulkCString(m, l0, int32(_a435))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v147 = int64(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v148 != 0 {
		v154 = v147
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_addReplyLongLong(m, l0, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v50)+152))
	if v149 == int32(0) {
		v154 = v147
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v149)+160))
	v154 = v152
	goto L48
L51:
	;
	v158 = v9 + int32(8)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v160 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v160 != 0 {
		v44 = v160
		goto L9
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160+base.B2i32(v163 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v169
	goto L53
L55:
	;
	goto L10
}
func F_clusterCommandSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v29 = F_isClientConnIpV6(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	v26 = base.B2i32(v23 != int32(0))
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v18 = F_connectionTypeTls(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v26 = base.B2i32(v17 == v18)
	goto L1
L7:
	;
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v26 | int32(2)
	goto L10
L9:
	;
	v31 = v26
	goto L10
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v34 == int32(3) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v31 | int32(4)
	goto L13
L12:
	;
	v37 = v31
	goto L13
L13:
	;
	v42 = m.G0
	v44 = v42 - int32(32)
	m.G0 = v44
	v47 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	F_dictInitIterator(m, v44, v48)
	mBase = m.M
	v52 = int32(0)
	goto L17
L14:
	;
	v73 = v37 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[112])))
	if v75 != 0 {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	if v52 == int32(0) {
		goto L14
	} else {
		goto L21
	}
L16:
	;
	m.G0 = v44 + int32(32)
	goto L15
L17:
	;
	v55 = F_dictNext(m, v44)
	mBase = m.M
	if v55 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v58 = F_dictGetVal(m, v55)
	mBase = m.M
	v59 = F_isNodeAvailable(m, v58)
	mBase = m.M
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+2356))
	if v59 == v60 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+2356)) = v59
	v52 = int32(1)
	goto L17
L21:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a229), int32(_a203), int32(1585))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L63
	}
L24:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-1)))))
	switch v162 & int32(7) {
	case 0:
		goto L61
	case 1:
		goto L60
	case 2:
		goto L59
	case 3:
		goto L58
	case 4:
		goto L57
	default:
		v179 = int32(0)
		goto L56
	}
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v83 == int32(0) {
		v156 = v75
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v79 = F_generateClusterSlotResponse(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[112]))) = v79
	v156 = v79
	goto L24
L28:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v87 = F_generateClusterSlotResponse(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	F_sdsfree(m, v87)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L54
	}
L30:
	;
	v89 = int32(0)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+int32(-1)))))
	switch v96 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v113 = v89
		goto L32
	}
L31:
	;
	if v139 == int32(0) {
		goto L29
	} else {
		goto L50
	}
L32:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-1)))))
	switch v116 & int32(7) {
	case 0:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	default:
		v133 = v89
		goto L38
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v87+int32(-17))))
	v113 = v112
	goto L32
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v87+int32(-9))))
	v113 = v109
	goto L32
L35:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87+int32(-5)))))
	v113 = v106
	goto L32
L36:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+int32(-3)))))
	v113 = v103
	goto L32
L37:
	;
	v113 = int32(base.Ui32(v96) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	v134 = base.B2i32(base.Ui32(v113) < base.Ui32(v133))
	if base.Ui32(v113) < base.Ui32(v133) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-17))))
	v133 = v132
	goto L38
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-9))))
	v133 = v129
	goto L38
L41:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-5)))))
	v133 = v126
	goto L38
L42:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-3)))))
	v133 = v123
	goto L38
L43:
	;
	v133 = int32(base.Ui32(v116) >> (uint(int32(3)) % 32))
	goto L38
L44:
	;
	v135 = v113
	goto L46
L45:
	;
	v135 = v133
	goto L46
L46:
	;
	v136 = F_memcmp(m, v87, v75, v135)
	mBase = m.M
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v139 = v136
	goto L49
L48:
	;
	v139 = base.B2i32(base.Ui32(v133) < base.Ui32(v113)) - v134
	goto L49
L49:
	;
	goto L31
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v143 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v87
	F__serverLog(m, int32(3), int32(_a230), v8)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_sdsfree(m, v87)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	goto L23
L54:
	;
	if v139 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v156 = v75
	goto L24
L56:
	;
	F_addReplyProto(m, l0, v156, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L62
	}
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-17))))
	v179 = v178
	goto L56
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-9))))
	v179 = v175
	goto L56
L59:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+int32(-5)))))
	v179 = v172
	goto L56
L60:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-3)))))
	v179 = v169
	goto L56
L61:
	;
	v179 = int32(base.Ui32(v162) >> (uint(int32(3)) % 32))
	goto L56
L62:
	;
	m.G0 = v8 + int32(16)
	return
L63:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterCommandSyncSlots(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
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
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v11 = F_objectGetVal(m, v10)
	mBase = m.M
	v12 = int32(_a436)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v55 = F_objectGetVal(m, v54)
	mBase = m.M
	v56 = int32(_a437)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	if v47-v49 != 0 {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v47 = F_tolower(m, v43)
	mBase = m.M
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v49 = F_tolower(m, v48)
	mBase = m.M
	goto L3
L5:
	;
	v17 = v11
	v18 = v12
	v19 = v15
	goto L8
L6:
	;
	v43 = int32(0)
	v44 = v12
	goto L4
L7:
	;
	v43 = v40 & int32(255)
	v44 = v39
	goto L4
L8:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v21 == int32(0) {
		v39 = v18
		v40 = v19
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v39 = v33
	v40 = int32(0)
	goto L7
L10:
	;
	v25 = v19 & int32(255)
	if v25 == v21 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(1)
	v33 = v18 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v34 != 0 {
		v17 = v17 + v32
		v18 = v33
		v19 = v34
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v27 = F_tolower(m, v25)
	mBase = m.M
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v29 = F_tolower(m, v28)
	mBase = m.M
	if v27 == v29 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v39 = v18
	v40 = v31
	goto L7
L14:
	;
	goto L9
L15:
	;
	F_clusterCommandSyncSlotsEstablish(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L1
L18:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v97&int32(1) != 0 {
		goto L1
	} else {
		goto L33
	}
L19:
	;
	if v91-v93 != 0 {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v91 = F_tolower(m, v87)
	mBase = m.M
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v93 = F_tolower(m, v92)
	mBase = m.M
	goto L19
L21:
	;
	v61 = v55
	v62 = v56
	v63 = v59
	goto L24
L22:
	;
	v87 = int32(0)
	v88 = v56
	goto L20
L23:
	;
	v87 = v84 & int32(255)
	v88 = v83
	goto L20
L24:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == int32(0) {
		v83 = v62
		v84 = v63
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v83 = v77
	v84 = int32(0)
	goto L23
L26:
	;
	v69 = v63 & int32(255)
	if v69 == v65 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v76 = int32(1)
	v77 = v62 + v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v78 != 0 {
		v61 = v61 + v76
		v62 = v77
		v63 = v78
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v71 = F_tolower(m, v69)
	mBase = m.M
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v73 = F_tolower(m, v72)
	mBase = m.M
	if v71 == v73 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v83 = v62
	v84 = v75
	goto L23
L30:
	;
	goto L25
L31:
	;
	F_clusterCommandSyncSlotsFinish(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	goto L1
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = F_objectGetVal(m, v101)
	mBase = m.M
	v103 = int32(_a438)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v146 = F_objectGetVal(m, v145)
	mBase = m.M
	v147 = int32(_a439)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 != 0 {
		goto L52
	} else {
		goto L53
	}
L35:
	;
	if v138-v140 != 0 {
		goto L34
	} else {
		goto L47
	}
L36:
	;
	v138 = F_tolower(m, v134)
	mBase = m.M
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v140 = F_tolower(m, v139)
	mBase = m.M
	goto L35
L37:
	;
	v108 = v102
	v109 = v103
	v110 = v106
	goto L40
L38:
	;
	v134 = int32(0)
	v135 = v103
	goto L36
L39:
	;
	v134 = v131 & int32(255)
	v135 = v130
	goto L36
L40:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v112 == int32(0) {
		v130 = v109
		v131 = v110
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v130 = v124
	v131 = int32(0)
	goto L39
L42:
	;
	v116 = v110 & int32(255)
	if v116 == v112 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v123 = int32(1)
	v124 = v109 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v125 != 0 {
		v108 = v108 + v123
		v109 = v124
		v110 = v125
		goto L40
	} else {
		goto L46
	}
L44:
	;
	v118 = F_tolower(m, v116)
	mBase = m.M
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v120 = F_tolower(m, v119)
	mBase = m.M
	if v118 == v120 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v130 = v109
	v131 = v122
	goto L39
L46:
	;
	goto L41
L47:
	;
	F_clusterCommandSyncSlotsSnapshotEof(m, l0)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	goto L1
L49:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v190 = F_objectGetVal(m, v189)
	mBase = m.M
	v191 = int32(_a440)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 != 0 {
		goto L67
	} else {
		goto L68
	}
L50:
	;
	if v182-v184 != 0 {
		goto L49
	} else {
		goto L62
	}
L51:
	;
	v182 = F_tolower(m, v178)
	mBase = m.M
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v184 = F_tolower(m, v183)
	mBase = m.M
	goto L50
L52:
	;
	v152 = v146
	v153 = v147
	v154 = v150
	goto L55
L53:
	;
	v178 = int32(0)
	v179 = v147
	goto L51
L54:
	;
	v178 = v175 & int32(255)
	v179 = v174
	goto L51
L55:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v156 == int32(0) {
		v174 = v153
		v175 = v154
		goto L54
	} else {
		goto L57
	}
L56:
	;
	v174 = v168
	v175 = int32(0)
	goto L54
L57:
	;
	v160 = v154 & int32(255)
	if v160 == v156 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v167 = int32(1)
	v168 = v153 + v167
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	if v169 != 0 {
		v152 = v152 + v167
		v153 = v168
		v154 = v169
		goto L55
	} else {
		goto L61
	}
L59:
	;
	v162 = F_tolower(m, v160)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	if v162 == v164 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v174 = v153
	v175 = v166
	goto L54
L61:
	;
	goto L56
L62:
	;
	F_clusterCommandSyncSlotsRequestPause(m, l0)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	goto L1
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v234 = F_objectGetVal(m, v233)
	mBase = m.M
	v235 = int32(_a441)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v238 != 0 {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	if v226-v228 != 0 {
		goto L64
	} else {
		goto L77
	}
L66:
	;
	v226 = F_tolower(m, v222)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	goto L65
L67:
	;
	v196 = v190
	v197 = v191
	v198 = v194
	goto L70
L68:
	;
	v222 = int32(0)
	v223 = v191
	goto L66
L69:
	;
	v222 = v219 & int32(255)
	v223 = v218
	goto L66
L70:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v200 == int32(0) {
		v218 = v197
		v219 = v198
		goto L69
	} else {
		goto L72
	}
L71:
	;
	v218 = v212
	v219 = int32(0)
	goto L69
L72:
	;
	v204 = v198 & int32(255)
	if v204 == v200 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v211 = int32(1)
	v212 = v197 + v211
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v213 != 0 {
		v196 = v196 + v211
		v197 = v212
		v198 = v213
		goto L70
	} else {
		goto L76
	}
L74:
	;
	v206 = F_tolower(m, v204)
	mBase = m.M
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v208 = F_tolower(m, v207)
	mBase = m.M
	if v206 == v208 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v218 = v197
	v219 = v210
	goto L69
L76:
	;
	goto L71
L77:
	;
	F_clusterCommandSyncSlotsPaused(m, l0)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	goto L1
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	v278 = F_objectGetVal(m, v277)
	mBase = m.M
	v279 = int32(_a411)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v282 != 0 {
		goto L97
	} else {
		goto L98
	}
L80:
	;
	if v270-v272 != 0 {
		goto L79
	} else {
		goto L92
	}
L81:
	;
	v270 = F_tolower(m, v266)
	mBase = m.M
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v272 = F_tolower(m, v271)
	mBase = m.M
	goto L80
L82:
	;
	v240 = v234
	v241 = v235
	v242 = v238
	goto L85
L83:
	;
	v266 = int32(0)
	v267 = v235
	goto L81
L84:
	;
	v266 = v263 & int32(255)
	v267 = v262
	goto L81
L85:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v244 == int32(0) {
		v262 = v241
		v263 = v242
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v262 = v256
	v263 = int32(0)
	goto L84
L87:
	;
	v248 = v242 & int32(255)
	if v248 == v244 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v255 = int32(1)
	v256 = v241 + v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v257 != 0 {
		v240 = v240 + v255
		v241 = v256
		v242 = v257
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v250 = F_tolower(m, v248)
	mBase = m.M
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	v252 = F_tolower(m, v251)
	mBase = m.M
	if v250 == v252 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v262 = v241
	v263 = v254
	goto L84
L91:
	;
	goto L86
L92:
	;
	F_clusterCommandSyncSlotsRequestFailover(m, l0)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L93
	}
L93:
	;
	goto L1
L94:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	v322 = F_objectGetVal(m, v321)
	mBase = m.M
	v323 = int32(_a442)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v326 != 0 {
		goto L112
	} else {
		goto L113
	}
L95:
	;
	if v314-v316 != 0 {
		goto L94
	} else {
		goto L107
	}
L96:
	;
	v314 = F_tolower(m, v310)
	mBase = m.M
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v316 = F_tolower(m, v315)
	mBase = m.M
	goto L95
L97:
	;
	v284 = v278
	v285 = v279
	v286 = v282
	goto L100
L98:
	;
	v310 = int32(0)
	v311 = v279
	goto L96
L99:
	;
	v310 = v307 & int32(255)
	v311 = v306
	goto L96
L100:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v288 == int32(0) {
		v306 = v285
		v307 = v286
		goto L99
	} else {
		goto L102
	}
L101:
	;
	v306 = v300
	v307 = int32(0)
	goto L99
L102:
	;
	v292 = v286 & int32(255)
	if v292 == v288 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v299 = int32(1)
	v300 = v285 + v299
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v301 != 0 {
		v284 = v284 + v299
		v285 = v300
		v286 = v301
		goto L100
	} else {
		goto L106
	}
L104:
	;
	v294 = F_tolower(m, v292)
	mBase = m.M
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v296 = F_tolower(m, v295)
	mBase = m.M
	if v294 == v296 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	v306 = v285
	v307 = v298
	goto L99
L106:
	;
	goto L101
L107:
	;
	F_clusterCommandSyncSlotsFailoverGranted(m, l0)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	goto L1
L109:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v366 = F_objectGetVal(m, v365)
	mBase = m.M
	v367 = int32(_a443)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v370 != 0 {
		goto L126
	} else {
		goto L127
	}
L110:
	;
	if v358-v360 != 0 {
		goto L109
	} else {
		goto L122
	}
L111:
	;
	v358 = F_tolower(m, v354)
	mBase = m.M
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	v360 = F_tolower(m, v359)
	mBase = m.M
	goto L110
L112:
	;
	v328 = v322
	v329 = v323
	v330 = v326
	goto L115
L113:
	;
	v354 = int32(0)
	v355 = v323
	goto L111
L114:
	;
	v354 = v351 & int32(255)
	v355 = v350
	goto L111
L115:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	if v332 == int32(0) {
		v350 = v329
		v351 = v330
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v350 = v344
	v351 = int32(0)
	goto L114
L117:
	;
	v336 = v330 & int32(255)
	if v336 == v332 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v343 = int32(1)
	v344 = v329 + v343
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
	if v345 != 0 {
		v328 = v328 + v343
		v329 = v344
		v330 = v345
		goto L115
	} else {
		goto L121
	}
L119:
	;
	v338 = F_tolower(m, v336)
	mBase = m.M
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	v340 = F_tolower(m, v339)
	mBase = m.M
	if v338 == v340 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	v350 = v329
	v351 = v342
	goto L114
L121:
	;
	goto L116
L122:
	;
	F_clusterCommandSyncSlotsAck(m, l0)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L123
	}
L123:
	;
	goto L1
L124:
	;
	if v402-v404 == int32(0) {
		goto L1
	} else {
		goto L136
	}
L125:
	;
	v402 = F_tolower(m, v398)
	mBase = m.M
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v404 = F_tolower(m, v403)
	mBase = m.M
	goto L124
L126:
	;
	v372 = v366
	v373 = v367
	v374 = v370
	goto L129
L127:
	;
	v398 = int32(0)
	v399 = v367
	goto L125
L128:
	;
	v398 = v395 & int32(255)
	v399 = v394
	goto L125
L129:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v376 == int32(0) {
		v394 = v373
		v395 = v374
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v394 = v388
	v395 = int32(0)
	goto L128
L131:
	;
	v380 = v374 & int32(255)
	if v380 == v376 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v387 = int32(1)
	v388 = v373 + v387
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+1)))
	if v389 != 0 {
		v372 = v372 + v387
		v373 = v388
		v374 = v389
		goto L129
	} else {
		goto L135
	}
L133:
	;
	v382 = F_tolower(m, v380)
	mBase = m.M
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v384 = F_tolower(m, v383)
	mBase = m.M
	if v382 == v384 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v394 = v373
	v395 = v386
	goto L128
L135:
	;
	goto L130
L136:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v408 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L146
	}
L138:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v408)+156))
	if base.Ui32(int32(20)) < base.Ui32(v411) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v419 {
		v429 = v408
		goto L142
	} else {
		goto L143
	}
L140:
	;
	if int32(1)<<(uint(v411)%32)&int32(1835040) != 0 {
		goto L137
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	F_finishSlotMigrationJob(m, v429, int32(18), int32(_a444))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L16
	} else {
		goto L145
	}
L143:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v408)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v422
	F__serverLog(m, int32(3), int32(_a445), v7)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L16
	} else {
		goto L144
	}
L144:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v429 = v428
	goto L142
L145:
	;
	goto L1
L146:
	;
	goto L1
}
func F_clusterCommandSyncSlotsEstablish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
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
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v17 != int64(-1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v74 = int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v75&v74 != 0 {
		v81 = v74
		goto L24
	} else {
		goto L25
	}
L3:
	;
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L4:
	;
	v21 = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v22&v21 != 0 {
		v29 = v21
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v32 = int32(1)
	goto L3
L6:
	;
	v32 = v29
	goto L3
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = F_isImportSlotMigrationJob(m, v25)
	mBase = m.M
	v29 = v27
	goto L6
L9:
	;
	v32 = int32(0)
	goto L3
L10:
	;
	v33 = F_moduleVerifyAllAllowAtomicSlotMigrationOrReply(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v33 == int32(-1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+88)))
	if v40&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	goto L18
L15:
	;
	F_addReplyError(m, l0, int32(_a369))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	goto L21
L18:
	;
	if base.B2i32(v50 != int32(0)-v52) == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_addReplyError(m, l0, int32(_a370))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	if base.B2i32(v64 != int32(0)-v66) == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_addReplyError(m, l0, int32(_a371))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	v83 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v86 < int32(4) {
		v487 = int32(3)
		v488 = v83
		v489 = v83
		v490 = v83
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v81 = base.B2i32(v78 == int64(-1))
	goto L24
L26:
	;
	F_listRelease(m, v557)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L11
	} else {
		goto L165
	}
L27:
	;
	if v489 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L28:
	;
	if v512 == int32(0) {
		goto L1
	} else {
		goto L153
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v487
	v494 = int32(0)
	v498 = base.B2i32(v490 != v494) & base.B2i32(v488 != v494)
	if v81 != 0 {
		goto L148
	} else {
		goto L149
	}
L30:
	;
	v89 = int32(0)
	v96 = int32(3)
	v97 = v89
	v98 = v89
	v99 = v89
	goto L31
L31:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v104 = v96 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v104)))
	v107 = F_objectGetVal(m, v106)
	mBase = m.M
	v108 = int32(_a372)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v487 = v480
	v488 = v97
	v489 = v481
	v490 = v99
	goto L29
L33:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v480 < v482 {
		v96 = v480
		v98 = v481
		goto L31
	} else {
		goto L146
	}
L34:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213+v104)))
	v216 = F_objectGetVal(m, v215)
	mBase = m.M
	v217 = int32(_a373)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v220 != 0 {
		goto L70
	} else {
		goto L71
	}
L35:
	;
	if v143-v145 != 0 {
		goto L34
	} else {
		goto L47
	}
L36:
	;
	v143 = F_tolower(m, v139)
	mBase = m.M
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v145 = F_tolower(m, v144)
	mBase = m.M
	goto L35
L37:
	;
	v113 = v107
	v114 = v108
	v115 = v111
	goto L40
L38:
	;
	v139 = int32(0)
	v140 = v108
	goto L36
L39:
	;
	v139 = v136 & int32(255)
	v140 = v135
	goto L36
L40:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 == int32(0) {
		v135 = v114
		v136 = v115
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v135 = v129
	v136 = int32(0)
	goto L39
L42:
	;
	v121 = v115 & int32(255)
	if v121 == v117 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v128 = int32(1)
	v129 = v114 + v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v130 != 0 {
		v113 = v113 + v128
		v114 = v129
		v115 = v130
		goto L40
	} else {
		goto L46
	}
L44:
	;
	v123 = F_tolower(m, v121)
	mBase = m.M
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v125 = F_tolower(m, v124)
	mBase = m.M
	if v123 == v125 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v135 = v114
	v136 = v127
	goto L39
L46:
	;
	goto L41
L47:
	;
	if v98 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v81 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v96
	v182 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L58
	}
L50:
	;
	v148 = v96 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v149 <= v148 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v148<<(uint(int32(2))%32))))
	v156 = F_objectGetVal(m, v155)
	mBase = m.M
	v157 = int32(-1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v157))))
	switch v159&int32(7) + v157 {
	case 0:
		goto L56
	case 1:
		goto L55
	case 2:
		goto L54
	case 3:
		goto L53
	default:
		goto L49
	}
L52:
	;
	if v176 == int32(40) {
		goto L48
	} else {
		goto L57
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-17))))
	v176 = v175
	goto L52
L54:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v156+int32(-9))))
	v176 = v172
	goto L52
L55:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+int32(-5)))))
	v176 = v169
	goto L52
L56:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(-3)))))
	v176 = v166
	goto L52
L57:
	;
	goto L49
L58:
	;
	v512 = v97
	goto L28
L59:
	;
	v191 = v96 + int32(2)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v104+int32(4))))
	v197 = F_objectGetVal(m, v196)
	mBase = m.M
	v199 = F_clusterLookupNode(m, v197, int32(40))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L11
	} else {
		goto L62
	}
L60:
	;
	v480 = v96 + int32(2)
	v481 = int32(0)
	goto L33
L61:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v199 != v207 {
		v480 = v191
		v481 = v199
		goto L33
	} else {
		goto L65
	}
L62:
	;
	if v199 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v191
	F_addReplyError(m, l0, int32(_a374))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v512 = v97
	goto L28
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v191
	F_addReplyError(m, l0, int32(_a375))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v512 = v97
	goto L28
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v96
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v305+v104)))
	v308 = F_objectGetVal(m, v307)
	mBase = m.M
	v309 = int32(_a376)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v312 != 0 {
		goto L96
	} else {
		goto L97
	}
L68:
	;
	if v252-v254 != 0 {
		goto L67
	} else {
		goto L80
	}
L69:
	;
	v252 = F_tolower(m, v248)
	mBase = m.M
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v254 = F_tolower(m, v253)
	mBase = m.M
	goto L68
L70:
	;
	v222 = v216
	v223 = v217
	v224 = v220
	goto L73
L71:
	;
	v248 = int32(0)
	v249 = v217
	goto L69
L72:
	;
	v248 = v245 & int32(255)
	v249 = v244
	goto L69
L73:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v226 == int32(0) {
		v244 = v223
		v245 = v224
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v244 = v238
	v245 = int32(0)
	goto L72
L75:
	;
	v230 = v224 & int32(255)
	if v230 == v226 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v237 = int32(1)
	v238 = v223 + v237
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v239 != 0 {
		v222 = v222 + v237
		v223 = v238
		v224 = v239
		goto L73
	} else {
		goto L79
	}
L77:
	;
	v232 = F_tolower(m, v230)
	mBase = m.M
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v234 = F_tolower(m, v233)
	mBase = m.M
	if v232 == v234 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v244 = v223
	v245 = v236
	goto L72
L79:
	;
	goto L74
L80:
	;
	if v99 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v104+int32(4))))
	v299 = F_objectGetVal(m, v298)
	mBase = m.M
	v301 = v96 + int32(2)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v301 < v302 {
		v96 = v301
		v99 = v299
		goto L31
	} else {
		goto L92
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v96
	v291 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L11
	} else {
		goto L91
	}
L83:
	;
	v257 = v96 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v258 <= v257 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v257<<(uint(int32(2))%32))))
	v265 = F_objectGetVal(m, v264)
	mBase = m.M
	v266 = int32(-1)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v266))))
	switch v268&int32(7) + v266 {
	case 0:
		goto L89
	case 1:
		goto L88
	case 2:
		goto L87
	case 3:
		goto L86
	default:
		goto L82
	}
L85:
	;
	if v285 == int32(40) {
		goto L81
	} else {
		goto L90
	}
L86:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-17))))
	v285 = v284
	goto L85
L87:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v265+int32(-9))))
	v285 = v281
	goto L85
L88:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265+int32(-5)))))
	v285 = v278
	goto L85
L89:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+int32(-3)))))
	v285 = v275
	goto L85
L90:
	;
	goto L82
L91:
	;
	v512 = v97
	goto L28
L92:
	;
	v487 = v301
	v488 = v97
	v489 = v98
	v490 = v299
	goto L29
L93:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L11
	} else {
		goto L145
	}
L94:
	;
	if v344-v346 != 0 {
		goto L93
	} else {
		goto L106
	}
L95:
	;
	v344 = F_tolower(m, v340)
	mBase = m.M
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	v346 = F_tolower(m, v345)
	mBase = m.M
	goto L94
L96:
	;
	v314 = v308
	v315 = v309
	v316 = v312
	goto L99
L97:
	;
	v340 = int32(0)
	v341 = v309
	goto L95
L98:
	;
	v340 = v337 & int32(255)
	v341 = v336
	goto L95
L99:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v318 == int32(0) {
		v336 = v315
		v337 = v316
		goto L98
	} else {
		goto L101
	}
L100:
	;
	v336 = v330
	v337 = int32(0)
	goto L98
L101:
	;
	v322 = v316 & int32(255)
	if v322 == v318 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v329 = int32(1)
	v330 = v315 + v329
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v331 != 0 {
		v314 = v314 + v329
		v315 = v330
		v316 = v331
		goto L99
	} else {
		goto L105
	}
L103:
	;
	v324 = F_tolower(m, v322)
	mBase = m.M
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v326 = F_tolower(m, v325)
	mBase = m.M
	if v324 == v326 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	v336 = v315
	v337 = v328
	goto L98
L105:
	;
	goto L100
L106:
	;
	if v97 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v360 = F_parseSlotRangesOrReply(m, l0, v96+int32(1), v12+int32(8), v12+int32(12))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L11
	} else {
		goto L110
	}
L108:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	v557 = v97
	goto L26
L110:
	;
	if v360 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v364
	goto L112
L112:
	;
	goto L114
L113:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v473 < v474 {
		v96 = v473
		v97 = v360
		goto L31
	} else {
		goto L144
	}
L114:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v378 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	F_addReplyError(m, l0, int32(_a377))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L143
	}
L116:
	;
	if v378 == int32(0) {
		goto L113
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v378+base.B2i32(v381 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v387
	goto L117
L119:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v393 < v392 {
		goto L114
	} else {
		goto L120
	}
L120:
	;
	v398 = v392
	goto L122
L121:
	;
	goto L115
L122:
	;
	v404 = int32(0)
	v407 = m.G0
	v409 = v407 - int32(16)
	m.G0 = v409
	v412 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+uint32(_consts[171])))
	F_listRewind(m, v413, v409)
	mBase = m.M
	v416 = F_listNext(m, v409)
	mBase = m.M
	if v416 == v404 {
		v461 = v404
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v461 != 0 {
		goto L121
	} else {
		goto L141
	}
L125:
	;
	m.G0 = v409 + int32(16)
	goto L124
L126:
	;
	v422 = v416
	goto L128
L127:
	;
	v461 = int32(1)
	goto L125
L128:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	if v424 != int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v457 = F_listNext(m, v409)
	mBase = m.M
	if v457 != 0 {
		v422 = v457
		goto L128
	} else {
		goto L140
	}
L131:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423)+156))
	if base.Ui32(v427+int32(-18)) < base.Ui32(int32(3)) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v423)+164))
	v434 = v409 + int32(8)
	F_listRewind(m, v432, v434)
	mBase = m.M
	v438 = F_listNext(m, v434)
	mBase = m.M
	if v438 == int32(0) {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v444 = v438
	goto L134
L134:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v398 < v446 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L130
L136:
	;
	v452 = F_listNext(m, v409+int32(8))
	mBase = m.M
	if v452 != 0 {
		v444 = v452
		goto L134
	} else {
		goto L139
	}
L137:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v398 <= v448 {
		goto L127
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L135
L140:
	;
	v461 = v404
	goto L125
L141:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v466 <= v398 {
		goto L114
	} else {
		goto L142
	}
L142:
	;
	v398 = v398 + int32(1)
	goto L122
L143:
	;
	v512 = v360
	goto L28
L144:
	;
	v487 = v473
	v488 = v360
	v489 = v98
	v490 = v99
	goto L29
L145:
	;
	v512 = v97
	goto L28
L146:
	;
	goto L32
L147:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L11
	} else {
		goto L152
	}
L148:
	;
	if v498 != 0 {
		goto L27
	} else {
		goto L151
	}
L149:
	;
	v499 = int32(0)
	if v498&base.B2i32(v489 != v499) == v499 {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L27
L151:
	;
	goto L147
L152:
	;
	v512 = v488
	goto L28
L153:
	;
	v557 = v512
	goto L26
L154:
	;
	v526 = F_createSlotImportJob(m, l0, v489, v490, v488)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L158
	}
L155:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v489 == v521 {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	F_addReplyError(m, l0, int32(_a378))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L157
	}
L157:
	;
	v557 = v488
	goto L26
L158:
	;
	F_fireModuleSlotMigrationEvent(m, v526, int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L159
	}
L159:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+uint32(_consts[171])))
	v534 = F_listAddNodeHead(m, v533, v526)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L11
	} else {
		goto L160
	}
L160:
	;
	F_clusterDoBeforeSleep(m, int32(64))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L161
	}
L161:
	;
	F_forceCommandPropagation(m, l0, int32(3))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L11
	} else {
		goto L163
	}
L163:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v526)+152))
	if v546 == int32(0) {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v546)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v546)+200)) = v549 | int32(33554432)
	goto L1
L165:
	;
	goto L1
}
func F_clusterCommandSyncSlotsRequestFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v11 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
		switch v15 + int32(-5) {
		case 0, 13, 14, 15:
			F__serverAssert(m, int32(_a385), int32(_a386), int32(1522))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		default:
			v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v25 {
				v35 = v11
				F_finishSlotMigrationJob(m, v35, int32(18), int32(_a383))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
				F__serverLog(m, int32(3), int32(_a407), v9)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v35 = v34
					F_finishSlotMigrationJob(m, v35, int32(18), int32(_a383))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		case 11:
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v11)+176))
			v41 = F_mstime(m)
			mBase = m.M
			if v40 < v41 {
				v50 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(3) < v50 {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					F_finishSlotMigrationJob(m, v62, int32(18), int32(_a408))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
					F__serverLog(m, int32(3), int32(_a409), v9+int32(16))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						F_finishSlotMigrationJob(m, v62, int32(18), int32(_a408))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _consts[228]))
				if v48 != 0 {
					v67 = F_mstime(m)
					mBase = m.M
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)+176))
					v71 = v67 + int64(2000)
					if v71 <= v69 {
						v79 = v68
						F_sendSyncSlotsMessage(m, v79, int32(_a410))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v85 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(2) < v85 {
								v113 = *(*int64)(unsafe.Add(mBase, _consts[109]))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+156)) = int32(17)
								*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v113
								m.G0 = v9 + int32(48)
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+188))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+156))
								if base.Ui32(int32(20)) < base.Ui32(v90) {
									v98 = int32(_a86)
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[213])))
									v98 = v97
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(_a411)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
								F__serverLog(m, int32(2), int32(_a381), v9+int32(32))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v113 = *(*int64)(unsafe.Add(mBase, _consts[109]))
									*(*int32)(unsafe.Add(mBase, uint32(v83)+156)) = int32(17)
									*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v113
									m.G0 = v9 + int32(48)
									return
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v68)+176)) = v71
						F_pauseActions(m, int32(3), v71, int32(29))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							v79 = v78
							F_sendSyncSlotsMessage(m, v79, int32(_a410))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								v85 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(2) < v85 {
									v113 = *(*int64)(unsafe.Add(mBase, _consts[109]))
									*(*int32)(unsafe.Add(mBase, uint32(v83)+156)) = int32(17)
									*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v113
									m.G0 = v9 + int32(48)
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+188))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+156))
									if base.Ui32(int32(20)) < base.Ui32(v90) {
										v98 = int32(_a86)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[213])))
										v98 = v97
									}
									*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(_a411)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
									F__serverLog(m, int32(2), int32(_a381), v9+int32(32))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										v113 = *(*int64)(unsafe.Add(mBase, _consts[109]))
										*(*int32)(unsafe.Add(mBase, uint32(v83)+156)) = int32(17)
										*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v113
										m.G0 = v9 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(3) < v50 {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						F_finishSlotMigrationJob(m, v62, int32(18), int32(_a408))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
						F__serverLog(m, int32(3), int32(_a409), v9+int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							F_finishSlotMigrationJob(m, v62, int32(18), int32(_a408))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								m.G0 = v9 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a412))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v9 + int32(48)
			return
		}
	}
}
func F_clusterCommandSyncSlotsSnapshotEof(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v11 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
		switch v15 + int32(-1) {
		case 0:
			v41 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(2) < v41 {
				v53 = v11
				F_sendSyncSlotsMessage(m, v53, int32(_a379))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v59 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(2) < v59 {
						v87 = *(*int64)(unsafe.Add(mBase, _consts[109]))
						*(*int32)(unsafe.Add(mBase, uint32(v57)+156)) = int32(2)
						*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v87
						m.G0 = v9 + int32(48)
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+156))
						if base.Ui32(int32(20)) < base.Ui32(v64) {
							v72 = int32(_a86)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(int32(2))%32))+uint32(_consts[213])))
							v72 = v71
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a380)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v62
						F__serverLog(m, int32(2), int32(_a381), v9+int32(16))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v87 = *(*int64)(unsafe.Add(mBase, _consts[109]))
							*(*int32)(unsafe.Add(mBase, uint32(v57)+156)) = int32(2)
							*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v87
							m.G0 = v9 + int32(48)
							return
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v44
				F__serverLog(m, int32(2), int32(_a382), v9+int32(32))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v53 = v52
					F_sendSyncSlotsMessage(m, v53, int32(_a379))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						v59 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(2) < v59 {
							v87 = *(*int64)(unsafe.Add(mBase, _consts[109]))
							*(*int32)(unsafe.Add(mBase, uint32(v57)+156)) = int32(2)
							*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v87
							m.G0 = v9 + int32(48)
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+156))
							if base.Ui32(int32(20)) < base.Ui32(v64) {
								v72 = int32(_a86)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v64<<(uint(int32(2))%32))+uint32(_consts[213])))
								v72 = v71
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a380)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v62
							F__serverLog(m, int32(2), int32(_a381), v9+int32(16))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v87 = *(*int64)(unsafe.Add(mBase, _consts[109]))
								*(*int32)(unsafe.Add(mBase, uint32(v57)+156)) = int32(2)
								*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v87
								m.G0 = v9 + int32(48)
								return
							}
						}
					}
				}
			}
		default:
			v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v25 {
				v35 = v11
				F_finishSlotMigrationJob(m, v35, int32(18), int32(_a383))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v9 + int32(48)
					return
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
				F__serverLog(m, int32(3), int32(_a384), v9)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					v35 = v34
					F_finishSlotMigrationJob(m, v35, int32(18), int32(_a383))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		case 4, 17, 18, 19:
			F__serverAssert(m, int32(_a385), int32(_a386), int32(664))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a387))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v9 + int32(48)
			return
		}
	}
}
func F_clusterConnAcceptHandler(m *base.Module, l0 int32) {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 == int32(3) {
		v30 = F_valkey_malloc(m, int32(56))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v32 = F_mstime(m)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v30))) = v32
			v34 = F_listCreate(m)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(63)
				v39 = int32(1024)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(24)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(0)
				v46 = F_valkey_malloc(m, v39)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v46
					v51 = int32(_a20)
					v53 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v53 + int32(1048)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v30)+44)) = int64(4294967296)
					*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v30
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
					v66 = m.T0[v65].(func(*base.Module, int32, int32) int32)(m, l0, int32(64))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(1) < v13 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
			m.T0[v26].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
			v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v18
				F__serverLog(m, int32(1), int32(_a267), v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
					m.T0[v26].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_clusterDebugCommandExtendedHelp(m *base.Module) int32 {
	return int32(_a347)
}
func F_clusterDecodeOpenSlotsAuxField(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v297 int32
	_ = v297
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l0&int32(2) == v3 {
		v281 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a258), int32(_a247), int32(2220))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L52
	} else {
		goto L85
	}
L2:
	;
	m.G0 = v14 + int32(48)
	return v281
L3:
	;
	if l1 == int32(0) {
		v281 = v3
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v24 == int32(0) {
		v281 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v27 == int32(0) {
		v281 = v3
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v30 = v27
	v31 = l1
	goto L7
L7:
	;
	v41 = int32(-1)
	v45 = v31
	goto L10
L8:
	;
	v281 = int32(0)
	goto L2
L9:
	;
	if base.Ui32(int32(16383)) < base.Ui32(v90) {
		v281 = v41
		goto L2
	} else {
		goto L24
	}
L10:
	;
	v50 = v45 + int32(1)
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
	v52 = F___isspace_1(m, v51)
	mBase = m.M
	if v52 != 0 {
		v45 = v50
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v53 = int32(1)
	switch v51&int32(255) + int32(-43) {
	case 0:
		v59 = v53
		goto L14
	default:
		v61 = v45
		v62 = v51
		v63 = v53
		goto L13
	case 2:
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v66 = v62 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v66) {
		v84 = int32(0)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
	v61 = v50
	v62 = v60
	v63 = v59
	goto L13
L15:
	;
	v59 = int32(0)
	goto L14
L16:
	;
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v70 = int32(0)
	v71 = v61
	v72 = v66
	goto L18
L18:
	;
	v74 = int32(10)
	v76 = v70*v74 - v72
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+1)))
	v81 = v77 + int32(-48)
	if base.Ui32(v81) < base.Ui32(v74) {
		v70 = v76
		v71 = v71 + int32(1)
		v72 = v81
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v84 = v76
	goto L16
L20:
	;
	goto L19
L21:
	;
	v90 = int32(0) - v84
	goto L23
L22:
	;
	v90 = v84
	goto L23
L23:
	;
	goto L9
L24:
	;
	v93 = v30
	v94 = v31
	goto L26
L25:
	;
	v113 = int32(0)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v114 == int32(44) {
		v153 = v113
		v154 = v94
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v105 = v93 & int32(255)
	switch v105 + int32(-60) {
	case 0, 2:
		goto L25
	case 1:
		goto L28
	default:
		goto L29
	}
L28:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v93 = v110
	v94 = v94 + int32(1)
	goto L26
L29:
	;
	if v105 == int32(0) {
		v281 = v41
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v153 == int32(0) {
		v281 = v41
		goto L2
	} else {
		goto L40
	}
L32:
	;
	if v114 == int32(0) {
		v281 = v41
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v121 = v113
	v127 = v114
	v128 = v94 + int32(1)
	goto L34
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v121))) = uint8(v127)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	if v134 == int32(0) {
		v281 = v41
		goto L2
	} else {
		goto L36
	}
L35:
	;
	v153 = base.B2i32(v138 == int32(40)) & v140
	v154 = v128
	goto L31
L36:
	;
	v138 = v121 + int32(1)
	v140 = base.B2i32(v134 == int32(44))
	if v134 == int32(44) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	if base.B2i32(base.Ui32(int32(38)) < base.Ui32(v121)) == int32(0) {
		v121 = v138
		v127 = v134
		v128 = v128 + int32(1)
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L45
L41:
	;
	v233 = v154 + int32(2)
	v235 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v105 != int32(60) {
		goto L65
	} else {
		goto L66
	}
L42:
	;
	v218 = F_createClusterNode(m, v14, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L52
	} else {
		goto L59
	}
L43:
	;
	if int32(0)-v189 != 0 {
		goto L42
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	v174 = int32(0)
	goto L47
L46:
	;
	goto L44
L47:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v174))))
	v179 = int32(255)
	v189 = base.B2i32(base.Ui32((v176+int32(-123))&v179) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v176+int32(-58))&v179) < base.Ui32(int32(246)))
	if v189 != 0 {
		goto L46
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v191 = v174 + int32(1)
	if v191 != int32(40) {
		v174 = v191
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v201 = F_sdsnewlen(m, v14, int32(40))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(0)
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v208 = F_dictFind(m, v207, v201)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	F_sdsfree(m, v201)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	if v208 == int32(0) {
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	goto L57
L57:
	;
	if v214 != 0 {
		v230 = v214
		goto L41
	} else {
		goto L58
	}
L58:
	;
	goto L42
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+32))
	v226 = F_sdsnewlen(m, v218+int32(8), int32(40))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v228 = F_dictAdd(m, v222, v226, v218)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L52
	} else {
		goto L61
	}
L61:
	;
	if v228 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v230 = v218
	goto L41
L63:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v277 != 0 {
		v30 = v277
		v31 = v233
		goto L7
	} else {
		goto L84
	}
L64:
	;
	v272 = F_dictAdd(m, v250, v90, v230)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L52
	} else {
		goto L83
	}
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v235)+44))
	v255 = F_dictFind(m, v254, v90)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L52
	} else {
		goto L74
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)+48))
	v239 = F_dictFind(m, v238, v90)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L52
	} else {
		goto L67
	}
L67:
	;
	if v230 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+48))
	if v239 == int32(0) {
		goto L64
	} else {
		goto L72
	}
L69:
	;
	if v239 == int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	v246 = F_dictDelete(m, v245, v90)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L52
	} else {
		goto L71
	}
L71:
	;
	goto L63
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+8)) = v230
	goto L73
L73:
	;
	goto L63
L74:
	;
	if v230 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+44))
	if v255 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v255 == int32(0) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+44))
	v262 = F_dictDelete(m, v261, v90)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L52
	} else {
		goto L78
	}
L78:
	;
	goto L63
L79:
	;
	v270 = F_dictAdd(m, v266, v90, v230)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L52
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+8)) = v230
	goto L81
L81:
	;
	goto L63
L82:
	;
	goto L63
L83:
	;
	goto L63
L84:
	;
	goto L8
L85:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterDelNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int64
	_ = v217
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	v6 = m.G0
	v8 = v6 - int32(320)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a270), int32(_a247), int32(2236))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L108
	}
L2:
	;
	v12 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v12 < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v31 = v12
	v32 = v28
	goto L8
L4:
	;
	v17 = F_humanNodename(m, l0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0 + int32(8)
	F__serverLog(m, int32(0), int32(_a271), v8)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v35 = F_dictFind(m, v34, v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L12
	}
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	v91 = F_dictGetSafeIterator(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L35
	}
L10:
	;
	if v39 != l0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	goto L14
L12:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v39 = int32(0)
	goto L10
L14:
	;
	v39 = v38
	goto L10
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	v56 = F_dictFind(m, v55, v31)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L22
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v44 = F_dictFind(m, v43, v31)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v44 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+48))
	v51 = F_dictDelete(m, v50, v31)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if v60 != l0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	goto L24
L22:
	;
	if v56 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v60 = int32(0)
	goto L20
L24:
	;
	v60 = v59
	goto L20
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v31<<(uint(int32(2))%32))+52))
	if v79 != l0 {
		v85 = v75
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+44))
	v65 = F_dictFind(m, v64, v31)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v65 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v72 = F_dictDelete(m, v71, v31)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v87 = v31 + int32(1)
	if v87 != int32(16384) {
		v31 = v87
		v32 = v85
		goto L8
	} else {
		goto L33
	}
L31:
	;
	v81 = F_clusterDelSlot(m, v31)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v85 = v84
	goto L30
L33:
	;
	goto L9
L34:
	;
	F_dictReleaseIterator(m, v91)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L105
	}
L35:
	;
	v100 = v91 + int32(20)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v101 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v196 == int32(0) {
		goto L34
	} else {
		goto L62
	}
L37:
	;
	v107 = v100
	v108 = v104
	goto L40
L38:
	;
	v104 = int32(1)
	goto L37
L39:
	;
	v104 = int32(0)
	goto L37
L40:
	;
	switch v108 {
	case 0:
		goto L45
	default:
		goto L44
	}
L42:
	;
	v108 = int32(0)
	goto L40
L43:
	;
	goto L36
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v188
	if v188 == int32(0) {
		goto L42
	} else {
		goto L61
	}
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v112 != int32(-1) {
		v151 = v112
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v152 = int32(1)
	v153 = v151 + v152
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v153
	v155 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v159+int32(26)))))
	if v163 == int32(255) {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	if v116 != 0 {
		v151 = int32(-1)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v118 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v145 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v125 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+16)))
	v126 = int64(*(*int8)(unsafe.Add(mBase, uint32(v117)+27)))
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+8)))
	v128 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+12)))
	v129 = int64(*(*int8)(unsafe.Add(mBase, uint32(v117)+26)))
	v130 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117)+4)))
	v131 = F_wangHash64(m, v130)
	mBase = m.M
	v133 = F_wangHash64(m, v129+v131)
	mBase = m.M
	v135 = F_wangHash64(m, v128+v133)
	mBase = m.M
	v137 = F_wangHash64(m, v127+v135)
	mBase = m.M
	v139 = F_wangHash64(m, v126+v137)
	mBase = m.M
	v141 = F_wangHash64(m, v125+v139)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v144 = v143
	goto L49
L51:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+24)))
	v123 = v121 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+24)) = uint16(v123)
	v144 = v117
	goto L49
L52:
	;
	v151 = v145 + int32(-1)
	goto L46
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v151 = v148
	goto L46
L54:
	;
	v178 = int32(2)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v158+v176<<(uint(v178)%32)+int32(4))))
	v107 = v183 + v177<<(uint(v178)%32)
	v108 = int32(1)
	goto L40
L55:
	;
	v167 = v155
	goto L57
L56:
	;
	v167 = v152 << (uint(v163) % 32)
	goto L57
L57:
	;
	if v153 < v167 {
		v176 = v159
		v177 = v153
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if v159 != 0 {
		v196 = v155
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	if v169 == int32(-1) {
		v196 = v155
		goto L43
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = int64(4294967296)
	v176 = int32(1)
	v177 = int32(0)
	goto L54
L61:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v192
	v196 = v188
	goto L43
L62:
	;
	v204 = v196
	goto L63
L63:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	goto L66
L64:
	;
	goto L34
L65:
	;
	v272 = v91 + int32(20)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v273 != 0 {
		goto L80
	} else {
		goto L81
	}
L66:
	;
	if v207 == l0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v210 = v8 + int32(16)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+2352))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+20)) = int32(128)
	v217 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v210)+12)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v210)+296)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v210)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v8 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+156)) = v8 + int32(184)
	goto L68
L68:
	;
	v232 = int32(0)
	v234 = F_raxSeek(m, v8+int32(16), int32(_a67), v232, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	goto L71
L70:
	;
	F_raxStop(m, v8+int32(16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L77
	}
L71:
	;
	v243 = F_raxNext(m, v8+int32(16))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L73
	}
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v207)+2352))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v253 = F_raxRemove(m, v250, v247, v251, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L76
	}
L73:
	;
	if v243 == int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	if l0 != v248 {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L70
L77:
	;
	goto L65
L78:
	;
	if v368 != 0 {
		v204 = v368
		goto L63
	} else {
		goto L104
	}
L79:
	;
	v279 = v272
	v280 = v276
	goto L82
L80:
	;
	v276 = int32(1)
	goto L79
L81:
	;
	v276 = int32(0)
	goto L79
L82:
	;
	switch v280 {
	case 0:
		goto L87
	default:
		goto L86
	}
L84:
	;
	v280 = int32(0)
	goto L82
L85:
	;
	goto L78
L86:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v360
	if v360 == int32(0) {
		goto L84
	} else {
		goto L103
	}
L87:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v284 != int32(-1) {
		v323 = v284
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v324 = int32(1)
	v325 = v323 + v324
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v325
	v327 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v331+int32(26)))))
	if v335 == int32(255) {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	if v288 != 0 {
		v323 = int32(-1)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v290 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+20))
	if v317 != int32(-1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v289)+16)))
	v298 = int64(*(*int8)(unsafe.Add(mBase, uint32(v289)+27)))
	v299 = int64(*(*int32)(unsafe.Add(mBase, uint32(v289)+8)))
	v300 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v289)+12)))
	v301 = int64(*(*int8)(unsafe.Add(mBase, uint32(v289)+26)))
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v289)+4)))
	v303 = F_wangHash64(m, v302)
	mBase = m.M
	v305 = F_wangHash64(m, v301+v303)
	mBase = m.M
	v307 = F_wangHash64(m, v300+v305)
	mBase = m.M
	v309 = F_wangHash64(m, v299+v307)
	mBase = m.M
	v311 = F_wangHash64(m, v298+v309)
	mBase = m.M
	v313 = F_wangHash64(m, v297+v311)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v316 = v315
	goto L91
L93:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+24)))
	v295 = v293 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+24)) = uint16(v295)
	v316 = v289
	goto L91
L94:
	;
	v323 = v317 + int32(-1)
	goto L88
L95:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v323 = v320
	goto L88
L96:
	;
	v350 = int32(2)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v330+v348<<(uint(v350)%32)+int32(4))))
	v279 = v355 + v349<<(uint(v350)%32)
	v280 = int32(1)
	goto L82
L97:
	;
	v339 = v327
	goto L99
L98:
	;
	v339 = v324 << (uint(v335) % 32)
	goto L99
L99:
	;
	if v325 < v339 {
		v348 = v331
		v349 = v325
		goto L96
	} else {
		goto L100
	}
L100:
	;
	if v331 != 0 {
		v368 = v327
		goto L85
	} else {
		goto L101
	}
L101:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v330)+20))
	if v341 == int32(-1) {
		v368 = v327
		goto L85
	} else {
		goto L102
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = int64(4294967296)
	v348 = int32(1)
	v349 = int32(0)
	goto L96
L103:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v364
	v368 = v360
	goto L85
L104:
	;
	goto L64
L105:
	;
	F_clusterRemoveNodeFromShard(m, l0)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_freeClusterNode(m, l0)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	m.G0 = v8 + int32(320)
	return
L108:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterDelSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
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
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v92 int32
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+l0<<(uint(int32(2))%32)+int32(52))))
	if v14 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[25]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+l0<<(uint(int32(2))%32))))
		if v23 != 0 {
			v25 = F_hashtableSize(m, v23)
			mBase = m.M
			v26 = v25
		} else {
			v26 = int32(0)
		}
		if v26 == int32(0) {
			v36 = int32(1) << (uint(l0&int32(7)) % 32)
			v38 = base.I32_div_s(l0, int32(8))
			v41 = v14 + v38 + int32(104)
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
			if v36&v42 == int32(0) {
				F__serverAssert(m, int32(_a269), int32(_a247), int32(6527))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v46 = int32(-1)
				v47 = v36 ^ v46
				v48 = v42 & v47
				*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v48)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+2160))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+2160)) = v50 + v46
				v54 = int32(_a20)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[111]))
				*(*int32)(unsafe.Add(mBase, uint32(v55+l0<<(uint(int32(2))%32)+int32(52)))) = int32(0)
				v63 = v55 + v38
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[162]))))
				v67 = v66 & v47
				*(*uint8)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[162]))) = uint8(v67)
				v70 = *(*int32)(unsafe.Add(mBase, _consts[111]))
				v73 = v70 + l0*int32(24)
				v76 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[163]))) = v76
				*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[164]))) = v76
				*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[165]))) = v76
				return int32(0)
			}
		} else {
			F_pubsubShardUnsubscribeAllChannelsInSlot(m, l0)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v36 = int32(1) << (uint(l0&int32(7)) % 32)
				v38 = base.I32_div_s(l0, int32(8))
				v41 = v14 + v38 + int32(104)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				if v36&v42 == int32(0) {
					F__serverAssert(m, int32(_a269), int32(_a247), int32(6527))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v46 = int32(-1)
					v47 = v36 ^ v46
					v48 = v42 & v47
					*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v48)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+2160))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+2160)) = v50 + v46
					v54 = int32(_a20)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[111]))
					*(*int32)(unsafe.Add(mBase, uint32(v55+l0<<(uint(int32(2))%32)+int32(52)))) = int32(0)
					v63 = v55 + v38
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[162]))))
					v67 = v66 & v47
					*(*uint8)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[162]))) = uint8(v67)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[111]))
					v73 = v70 + l0*int32(24)
					v76 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[163]))) = v76
					*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[164]))) = v76
					*(*int64)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[165]))) = v76
					return int32(0)
				}
			}
		}
	} else {
		return int32(-1)
	}
}
func F_clusterFeedSlotExportJobs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v123 int32
	_ = v123
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l3 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a419), int32(_a386), int32(1768))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L35
	} else {
		goto L54
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[171])))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v19
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v24 == int32(0) {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24+base.B2i32(v27 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v33
	goto L6
L8:
	;
	v45 = v24
	v46 = int32(0)
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != 0 {
		v157 = v46
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v157 == int32(0) {
		goto L2
	} else {
		goto L52
	}
L11:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v161 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+152))
	if v51 == int32(0) {
		v157 = v46
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+156))
	if base.Ui32(int32(20)) < base.Ui32(v54) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui32(v54) < base.Ui32(int32(13)) {
		v157 = v46
		goto L11
	} else {
		goto L17
	}
L15:
	;
	if int32(1)<<(uint(v54)%32)&int32(1835040) != 0 {
		v157 = v46
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v49)+164))
	v65 = v12 + int32(8)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
	goto L18
L18:
	;
	v71 = v12 + int32(8)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v73 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v73 == int32(0) {
		v157 = v46
		goto L11
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73+base.B2i32(v76 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v82
	goto L20
L22:
	;
	v94 = v73
	goto L24
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v49)+152))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+96))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+28))
	if l0 == v116 {
		v127 = v46
		v128 = v114
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if l3 < v96 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v101 = v12 + int32(8)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v103 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if l3 <= v98 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v103 != 0 {
		v94 = v103
		goto L24
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103+base.B2i32(v106 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v112
	goto L30
L32:
	;
	v157 = v46
	goto L11
L33:
	;
	F_addReplyArrayLen(m, v128, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L35
	} else {
		goto L42
	}
L34:
	;
	v118 = F_selectDb(m, v114, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return
L36:
	;
	if v118 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v46 != 0 {
		v122 = v46
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v49)+152))
	F_addReply(m, v123, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v120 = F_generateSelectCommand(m, l0)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v122 = v120
	goto L38
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v49)+152))
	v127 = v122
	v128 = v126
	goto L33
L42:
	;
	if l2 < int32(1) {
		v157 = v127
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v140 = int32(0)
	goto L44
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v49)+152))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1+v140<<(uint(int32(2))%32))))
	F_addReplyBulk(m, v141, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L35
	} else {
		goto L46
	}
L45:
	;
	v157 = v127
	goto L11
L46:
	;
	v149 = v140 + int32(1)
	if v149 != l2 {
		v140 = v149
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if v161 != 0 {
		v45 = v161
		v46 = v157
		goto L9
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161+base.B2i32(v164 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v170
	goto L49
L51:
	;
	goto L10
L52:
	;
	F_decrRefCount(m, v157)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	goto L2
L54:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterGetMessageTypeString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(int32(10)) < base.Ui32(l0) {
		v11 = int32(_a86)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[181])))
		v11 = v10
	}
	return v11
}
func F_clusterGetNodesInMyShard(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v6 = F_sdsnewlen(m, l0+int32(48), int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
		v13 = F_dictFind(m, v12, v6)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v6)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v13 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					return v19
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_clusterGetReplicaRank(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v168 int32
	_ = v168
	v1 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+88)))
	if v13&int32(2) == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a314), int32(_a247), int32(5436))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2172))
	if v18 == int32(0) {
		v152 = v1
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v152
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v25 == int32(0) {
		v39 = int64(0)
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2164))
	if v45 < int32(1) {
		v152 = v1
		goto L3
	} else {
		goto L14
	}
L6:
	;
	v41 = int64(0)
	if v41 < v39 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v29 != 0 {
		v36 = v29
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+48))
	v39 = v38
	goto L6
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	if v32 == int32(0) {
		v39 = int64(0)
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v36 = v32
	goto L8
L11:
	;
	v44 = v39
	goto L13
L12:
	;
	v44 = v41
	goto L13
L13:
	;
	goto L5
L14:
	;
	v48 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v52 = v50 + int32(8)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2168))
	v55 = v48
	v56 = v48
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v53+v56<<(uint(int32(2))%32))))
	if v67 == v50 {
		v147 = v55
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v152 = v147
	goto L3
L17:
	;
	v150 = v56 + int32(1)
	if v150 != v45 {
		v55 = v147
		v56 = v150
		goto L15
	} else {
		goto L39
	}
L18:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+89)))
	if v69&int32(2) != 0 {
		v147 = v55
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v67)+2248))
	if v72 <= v44 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v72 != v44 {
		v147 = v55
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v147 = v55 + int32(1)
	goto L17
L22:
	;
	v78 = v67 + int32(8)
	v79 = int32(40)
	goto L27
L23:
	;
	v147 = int32(base.Ui32(v143)>>(uint(int32(31))%32)) + v55
	goto L17
L24:
	;
	v143 = int32(0)
	goto L23
L25:
	;
	v115 = v110
	v116 = v111
	v117 = v112
	goto L35
L26:
	;
	if v100 == int32(0) {
		goto L24
	} else {
		goto L33
	}
L27:
	;
	if (v52|v78)&int32(3) != 0 {
		v110 = v78
		v111 = v52
		v112 = v79
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v87 = v78
	v88 = v52
	v89 = v79
	goto L29
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v92 != v93 {
		v110 = v87
		v111 = v88
		v112 = v89
		goto L25
	} else {
		goto L31
	}
L30:
	;
	goto L26
L31:
	;
	v95 = int32(4)
	v96 = v88 + v95
	v98 = v87 + v95
	v100 = v89 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v87 = v98
		v88 = v96
		v89 = v100
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v110 = v98
	v111 = v96
	v112 = v100
	goto L25
L34:
	;
	v143 = v120 - v121
	goto L23
L35:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 != v121 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v123 = int32(1)
	v128 = v117 + int32(-1)
	if v128 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v115 = v115 + v123
	v116 = v116 + v123
	v117 = v128
	goto L35
L39:
	;
	goto L16
L40:
	;
	return int32(0)
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterHandleConfigEpochCollision(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+96))
	if v11 != v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v16&int32(1) == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)))
	if v21&int32(1) == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(8)
	v27 = l0 + v26
	v29 = v13 + v26
	v30 = int32(40)
	goto L9
L5:
	;
	if v94 < int32(1) {
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v94 = int32(0)
	goto L5
L7:
	;
	v66 = v61
	v67 = v62
	v68 = v63
	goto L17
L8:
	;
	if v51 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	if (v29|v27)&int32(3) != 0 {
		v61 = v27
		v62 = v29
		v63 = v30
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v38 = v27
	v39 = v29
	v40 = v30
	goto L11
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v43 != v44 {
		v61 = v38
		v62 = v39
		v63 = v40
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v46 = int32(4)
	v47 = v39 + v46
	v49 = v38 + v46
	v51 = v40 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v51) {
		v38 = v49
		v39 = v47
		v40 = v51
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v61 = v49
	v62 = v47
	v63 = v51
	goto L7
L16:
	;
	v94 = v71 - v72
	goto L5
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 != v72 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v74 = int32(1)
	v79 = v68 + int32(-1)
	if v79 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v66 = v66 + v74
	v67 = v67 + v74
	v68 = v79
	goto L17
L21:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)+8))
	v101 = v99 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v101
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v101
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	v106 = int32(_a20)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[122]))) = v108 | int32(44)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v113 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v116 = F_humanNodename(m, l0)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v120
	F__serverLog(m, int32(2), int32(_a289), v9)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L1
}
func F_clusterHandleManualFailover(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[147])))
	if v13 == int64(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[146])))
		if v16 != 0 {
			m.G0 = v9 + int32(16)
			return
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[148])))
			if v17 == int64(-1) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[166]))
				if v24 == int32(0) {
					v38 = int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					if v28 != 0 {
						v35 = v28
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+48))
						v38 = v37
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[168]))
						if v31 == int32(0) {
							v38 = int64(0)
						} else {
							v35 = v31
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+104))
							v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+48))
							v38 = v37
						}
					}
				}
				v40 = int64(0)
				if v40 < v38 {
					v43 = v38
				} else {
					v43 = v40
				}
				v45 = *(*int32)(unsafe.Add(mBase, _consts[111]))
				if v17 != v43 {
					v84 = v45
					v85 = int32(16)
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
					*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
					m.G0 = v9 + int32(16)
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[146])))
					if v48 != 0 {
						F__serverAssert(m, int32(_a315), int32(_a247), int32(5986))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[130])))
						if v49 == int64(0) {
							v69 = v45
							v71 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							v72 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[146]))) = v72
							if int32(2) < v71 {
								v84 = v69
								v85 = v72
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
								*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
								m.G0 = v9 + int32(16)
								return
							} else {
								F__serverLog(m, int32(2), int32(_a316), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, _consts[111]))
									v84 = v83
									v85 = v72
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
									*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[130]))) = int64(0)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if v55 < int32(4) {
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[127])))
								*(*int64)(unsafe.Add(mBase, uint32(v9))) = v61
								F__serverLog(m, int32(3), int32(_a317), v9)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, _consts[111]))
									v69 = v68
									v71 = *(*int32)(unsafe.Add(mBase, _consts[28]))
									v72 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[146]))) = v72
									if int32(2) < v71 {
										v84 = v69
										v85 = v72
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
										*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
										m.G0 = v9 + int32(16)
										return
									} else {
										F__serverLog(m, int32(2), int32(_a316), int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, _consts[111]))
											v84 = v83
											v85 = v72
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
											*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							} else {
								v58 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[146]))) = v58
								v84 = v45
								v85 = v58
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122])))
								*(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[122]))) = v87 | v85
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
func F_clusterHandleServerShutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	if l0 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(2) < v7 {
			v16 = F_clusterSaveConfig(m, int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[132]))
				if v19 == int32(-1) {
				} else {
				}
				return
			}
		} else {
			F__serverLog(m, int32(2), int32(_a268), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v16 = F_clusterSaveConfig(m, int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, _consts[132]))
					if v19 == int32(-1) {
					} else {
					}
					return
				}
			}
		}
	} else {
		F_clusterAutoFailoverOnShutdown(m)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(2) < v7 {
				v16 = F_clusterSaveConfig(m, int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, _consts[132]))
					if v19 == int32(-1) {
					} else {
					}
					return
				}
			} else {
				F__serverLog(m, int32(2), int32(_a268), int32(0))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v16 = F_clusterSaveConfig(m, int32(1))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, _consts[132]))
						if v19 == int32(-1) {
						} else {
						}
						return
					}
				}
			}
		}
	}
}
func F_clusterHandleSlotMigrationClientOOM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(1) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
		if base.Ui32(int32(20)) < base.Ui32(v11) {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v19 {
				F_finishSlotMigrationJob(m, l0, int32(18), int32(_a421))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
				F__serverLog(m, int32(3), int32(_a422), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_finishSlotMigrationJob(m, l0, int32(18), int32(_a421))
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
		} else {
			if int32(1)<<(uint(v11)%32)&int32(1835040) != 0 {
				m.G0 = v6 + int32(16)
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(3) < v19 {
					F_finishSlotMigrationJob(m, l0, int32(18), int32(_a421))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v22
					F__serverLog(m, int32(3), int32(_a422), v6)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_finishSlotMigrationJob(m, l0, int32(18), int32(_a421))
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
	}
}
func F_clusterHandleSlotMigrationErrorResponse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
		if base.Ui32(int32(20)) < base.Ui32(v5) {
			F_finishSlotMigrationJob(m, l0, int32(18), int32(_a420))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		} else {
			if int32(1)<<(uint(v5)%32)&int32(1835040) != 0 {
				return
			} else {
				F_finishSlotMigrationJob(m, l0, int32(18), int32(_a420))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_clusterInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int64
	_ = v268
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_valkey_malloc(m, int32(461184))
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
	*(*int32)(unsafe.Add(mBase, _consts[111])) = v12
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[122]))) = v15
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	v26 = F_dictCreate(m, int32(_a249))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v26
	v32 = F_dictCreate(m, int32(_a250))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v32
	v38 = F_dictCreate(m, int32(_a251))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v38
	v44 = F_dictCreate(m, int32(_a252))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v44
	v50 = F_dictCreate(m, int32(_a252))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[125]))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+48)) = v50
	v62 = F__emscripten_memset_bulkmem(m, v53+int32(52), base.I32_extend8_s(int32(0)), int32(65536))
	mBase = m.M
	goto L8
L8:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[126]))) = v65
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[127]))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[128]))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[129]))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[130]))) = v69
	v86 = F__emscripten_memset_bulkmem(m, v53+int32(65680), base.I32_extend8_s(v65), int32(176))
	mBase = m.M
	goto L9
L9:
	;
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[117]))) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[131]))) = v89
	v94 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	F_dictEmpty(m, v95, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	F_dictEmpty(m, v101, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v112 = F__emscripten_memset_bulkmem(m, v106+int32(65920), base.I32_extend8_s(int32(0)), int32(2048))
	mBase = m.M
	goto L12
L12:
	;
	v113 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[132])) = int32(-1)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v118 = F_clusterLockConfig(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	F__serverAssert(m, int32(_a254), int32(_a247), int32(1516))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L82
	}
L14:
	;
	F__serverAssert(m, int32(_a258), int32(_a247), int32(2220))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L81
	}
L15:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	if v118 == int32(-1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v124 = F_clusterLoadConfig(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v189 != 0 {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	if v124 != int32(-1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v130 = F_createClusterNode(m, int32(0), int32(17))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v132 = int32(_a20)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v130
	*(*int32)(unsafe.Add(mBase, _consts[119])) = v130
	v138 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v138 {
		v154 = v130
		v155 = v133
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	v160 = F_sdsnewlen(m, v154+int32(8), int32(40))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v130 + int32(8)
	F__serverLog(m, int32(2), int32(_a260), v8+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v153 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v154 = v153
	v155 = v151
	goto L22
L25:
	;
	v162 = F_dictAdd(m, v156, v160, v154)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v162 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	F_clusterAddNodeToShard(m, v165+int32(48), v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v171 = F_clusterSaveConfig(m, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v171 != int32(-1) {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v176 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F__serverLog(m, int32(3), int32(_a259), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	if v217 != 0 {
		goto L43
	} else {
		goto L44
	}
L35:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v194 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v195 = int32(244)
	goto L38
L37:
	;
	v195 = int32(240)
	goto L38
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v195)+uint32(_consts[154])))
	if v197 < int32(55536) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v201 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v204 = int32(_a256)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v204
	F__serverLog(m, int32(3), int32(_a257), v8)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v232 = F_rdbRegisterAuxField(m, int32(_a253), int32(59), int32(60))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v219 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F__serverLog(m, int32(3), int32(_a255), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	if v232 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	F_initClusterSlotMigrationJobList(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v238 = int32(_a20)
	v239 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v241 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v239 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v242 = v239
	goto L53
L52:
	;
	v242 = v241
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+2324)) = v242
	v244 = int32(_a20)
	v245 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v247 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v245 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v248 = v245
	goto L56
L55:
	;
	v248 = v247
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+2328)) = v248
	v251 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v251 != 0 {
		v259 = v251
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+2332)) = v259
	v261 = int32(_a20)
	v262 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+2336)) = v262
	v265 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+2340)) = v265
	v268 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[112])) = v268
	*(*int64)(unsafe.Add(mBase, _consts[143])) = v268
	*(*int64)(unsafe.Add(mBase, _consts[144])) = v268
	*(*int64)(unsafe.Add(mBase, _consts[145])) = v268
	v286 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v287 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[146]))) = v287
	*(*int64)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[147]))) = v268
	*(*int64)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[148]))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+uint32(_consts[149]))) = v287
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v237)+88))
	v299 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v306 = v295&int32(-7681) | base.B2i32(v299 != v287)<<(uint(int32(9))%32) | int32(7168)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+88)) = v306
	if v306 == v295 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v253 != 0 {
		v259 = v253
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v255 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v256 = v247
	goto L62
L61:
	;
	v256 = v241
	goto L62
L62:
	;
	v259 = v256 + int32(10000)
	goto L57
L63:
	;
	F_clusterUpdateMyselfIp(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+uint32(_consts[122]))) = v313 | int32(6)
	goto L63
L66:
	;
	v320 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v321 == v320 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v373 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	F_updateSdsExtensionField(m, v321+int32(2304), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v330 = int32(0)
	v331 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v331 == v330 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_updateSdsExtensionField(m, v331+int32(2308), v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v340 = int32(0)
	v341 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v341 == v340 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	F_updateSdsExtensionField(m, v341+int32(2312), v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v350 = int32(0)
	v351 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v351 == v350 {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	F_updateSdsExtensionField(m, v351+int32(2316), v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v360 = int32(0)
	v361 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v361 == v360 {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	F_updateSdsExtensionField(m, v361+int32(2320), v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L67
L78:
	;
	m.G0 = v8 + int32(32)
	return
L79:
	;
	goto L78
L80:
	;
	F_clusterSlotStatResetAll(m)
	mBase = m.M
	v378 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int64)(unsafe.Add(mBase, uint32(v378)+uint32(_consts[117]))) = int64(0)
	v385 = F___memset(m, v378+int32(65680), int32(0), int32(224))
	mBase = m.M
	goto L79
L81:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterInitLast(m *base.Module) {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = F_connTypeOfCluster(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v11 = m.T0[v10].(func(*base.Module) int32)(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = F_connectionByType(m, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				if v13 != 0 {
					v42 = int32(_a20)
					*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[156])) = int32(_a261)
					v52 = *(*int32)(unsafe.Add(mBase, _consts[135]))
					*(*int32)(unsafe.Add(mBase, _consts[157])) = v52
					v56 = *(*int32)(unsafe.Add(mBase, _consts[134]))
					v61 = *(*int32)(unsafe.Add(mBase, _consts[107]))
					if v61 != 0 {
						v62 = int32(244)
					} else {
						v62 = int32(240)
					}
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_consts[154])))
					if v56 != 0 {
						v67 = v56
					} else {
						v67 = v64 + int32(10000)
					}
					*(*int32)(unsafe.Add(mBase, _consts[158])) = v67
					v70 = F_connTypeOfCluster(m)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[159])) = v70
						v75 = int32(_a262)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+32))
						v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v75)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if v77 != int32(-1) {
								v97 = F_createSocketAcceptHandler(m, v75, int32(61))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									if v97 == int32(0) {
										m.G0 = v6 + int32(32)
										return
									} else {
										F__serverPanic_1(m, int32(_a247), int32(1563), int32(_a263), int32(0))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
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
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(3) < v82 {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, _consts[158]))
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v86
									F__serverLog(m, int32(3), int32(_a264), v6+int32(16))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										m.Env.Exit(m, int32(1))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(3) < v16 {
						m.Env.Exit(m, int32(1))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v20 = F_connTypeOfCluster(m)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							v23 = m.T0[v22].(func(*base.Module) int32)(m)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								if base.Ui32(int32(3)) < base.Ui32(v23) {
									v32 = int32(_a265)
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_consts[160])))
									v32 = v31
								}
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v32
								F__serverLog(m, int32(3), int32(_a266), v6)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.Env.Exit(m, int32(1))
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
	}
}
func F_clusterIsAnySlotExporting(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v10 == v1 {
		v67 = v1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v67
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v14 == int32(0) {
		v67 = v1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[171])))
	v19 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L4
L4:
	;
	v24 = int32(0)
	v26 = v6 + int32(8)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == v24 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v28 == int32(0) {
		v67 = v24
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28+base.B2i32(v31 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v37
	goto L6
L8:
	;
	v43 = v28
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v67 = v24
	goto L1
L11:
	;
	v53 = v6 + int32(8)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v55 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+156))
	if base.Ui32(v46+int32(-18)) <= base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v67 = int32(1)
	goto L1
L14:
	;
	if v55 != 0 {
		v43 = v55
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v55+base.B2i32(v58 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v64
	goto L15
L17:
	;
	goto L10
}
func F_clusterManualFailoverTimeLimit(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)+uint32(_consts[147])))
	return v3
}
func F_clusterMarkImportingSlotsInDb(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[171])))
	v12 = v6 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	goto L1
L1:
	;
	v18 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return
L3:
	;
	if v20 == int32(0) {
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
	v35 = v20
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v50 = v6 + int32(8)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v52 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+156))
	if base.Ui32(v40+int32(-18)) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+164))
	F_setSlotImportingStateInDb(m, l0, v45, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	goto L9
L14:
	;
	if v52 != 0 {
		v35 = v52
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52+base.B2i32(v55 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v61
	goto L15
L17:
	;
	goto L8
}
func F_clusterNodeAddReplica(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	if v8 <= v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = F_valkey_realloc(m, v7, v8<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v13 = v3
	goto L4
L3:
	;
	return int32(-1)
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7+v13<<(uint(int32(2))%32))))
	if v19 == l1 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v22 = v13 + int32(1)
	if v22 == v8 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v13 = v22
	goto L4
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2168)) = v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	*(*int32)(unsafe.Add(mBase, uint32(v35+v40<<(uint(int32(2))%32)))) = l1
	v46 = v40 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2164)) = v46
	F_qsort(m, v35, v46, int32(4), int32(58))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v52 | int32(256)
	return int32(0)
}
func F_clusterNodeGetPrimary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	v5 = l0
	goto L2
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v13 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+2172))
	if v7 == int32(0) {
		v11 = v5
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v11 = v7
	goto L1
L4:
	;
	if v7 != l0 {
		v5 = v7
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return v11
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2172))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F__serverAssert(m, int32(_a310), int32(_a247), int32(7623))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterNodeGetReplica(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3+l1<<(uint(int32(2))%32))))
	return v7
}
func F_clusterNodeIp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	if l1 != 0 {
		v7 = F_isClientConnIpV6(m, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2304))
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+int32(-1)))))
				switch v46 & int32(7) {
				case 0:
					v63 = int32(base.Ui32(v46) >> (uint(int32(3)) % 32))
					if v63 != 0 {
						v68 = v43
					} else {
						v68 = l0 + int32(2256)
					}
				case 1:
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+int32(-3)))))
					v63 = v53
					if v63 != 0 {
						v68 = v43
					} else {
						v68 = l0 + int32(2256)
					}
				case 2:
					v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43+int32(-5)))))
					v63 = v56
					if v63 != 0 {
						v68 = v43
					} else {
						v68 = l0 + int32(2256)
					}
				case 3:
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v43+int32(-9))))
					v63 = v59
					if v63 != 0 {
						v68 = v43
					} else {
						v68 = l0 + int32(2256)
					}
				case 4:
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v43+int32(-17))))
					v63 = v62
					if v63 != 0 {
						v68 = v43
					} else {
						v68 = l0 + int32(2256)
					}
				default:
					v68 = l0 + int32(2256)
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2308))
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
				switch v16 & int32(7) {
				case 0:
					if int32(base.Ui32(v16)>>(uint(int32(3))%32)) == int32(0) {
						v68 = l0 + int32(2256)
					} else {
						v68 = v13
					}
				case 1:
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
					if v25 == int32(0) {
						v68 = l0 + int32(2256)
					} else {
						v68 = v13
					}
				case 2:
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
					if v30 == int32(0) {
						v68 = l0 + int32(2256)
					} else {
						v68 = v13
					}
				case 3:
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
					if v35 == int32(0) {
						v68 = l0 + int32(2256)
					} else {
						v68 = v13
					}
				case 4:
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
					if v40 == int32(0) {
						v68 = l0 + int32(2256)
					} else {
						v68 = v13
					}
				default:
					v68 = l0 + int32(2256)
				}
			}
			return v68
		}
	} else {
		return l0 + int32(2256)
	}
}
func F_clusterNodeIsFailing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(8)
}
func F_clusterNodeIsNoFailover(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(512)
}
func F_clusterNodeIsPrimary(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(1)
}
func F_clusterNodePreferredEndpoint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	switch v6 {
	case 0:
		v7 = F_clusterNodeIp(m, l0, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	case 1:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2312))
		if v12 != 0 {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v16 != 0 {
				v17 = v12
			} else {
				v17 = int32(_a319)
			}
			return v17
		} else {
			return int32(_a319)
		}
	case 2:
		v20 = int32(_a320)
		return v20
	default:
		v20 = int32(_a86)
		return v20
	}
}
func F_clusterNodeRemoveReplica(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	v9 = int32(-1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	if v10 < int32(1) {
		v202 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v202
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2168))
	v20 = int32(0)
	goto L3
L3:
	;
	v24 = v20 + int32(1)
	v27 = v13 + v20<<(uint(int32(2))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v202 = v9
	goto L1
L5:
	;
	if v24 != v10 {
		v20 = v24
		goto L3
	} else {
		goto L51
	}
L6:
	;
	if v10 <= v24 {
		v188 = v10
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v190 = v188 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2164)) = v190
	if v190 != 0 {
		v202 = int32(0)
		goto L1
	} else {
		goto L50
	}
L8:
	;
	v31 = int32(2)
	v33 = v13 + v24<<(uint(v31)%32)
	v38 = (v10 + (v20 ^ int32(-1))) << (uint(v31) % 32)
	if v27 == v33 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2164))
	v188 = v187
	goto L7
L10:
	;
	goto L9
L11:
	;
	v42 = v38 + v27
	if base.Ui32(int32(0)-v38<<(uint(int32(1))%32)) < base.Ui32(v33-v42) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = (v33 ^ v27) & int32(3)
	if base.Ui32(v33) <= base.Ui32(v27) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v49 = F___memcpy(m, v27, v33, v38)
	mBase = m.M
	goto L9
L14:
	;
	if v158 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L15:
	;
	if base.Ui32(v136) <= base.Ui32(int32(3)) {
		v157 = v135
		v158 = v136
		v159 = v137
		goto L14
	} else {
		goto L42
	}
L16:
	;
	if v52 != 0 {
		v118 = v38
		goto L26
	} else {
		goto L27
	}
L17:
	;
	if v52 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v27&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v157 = v33
	v158 = v38
	v159 = v27
	goto L14
L20:
	;
	v59 = v33
	v60 = v38
	v61 = v27
	goto L22
L21:
	;
	v135 = v33
	v136 = v38
	v137 = v27
	goto L15
L22:
	;
	if v60 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v65)
	v67 = int32(1)
	v68 = v59 + v67
	v70 = v60 + int32(-1)
	v72 = v61 + v67
	if v72&int32(3) == int32(0) {
		v135 = v68
		v136 = v70
		v137 = v72
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v59 = v68
	v60 = v70
	v61 = v72
	goto L22
L26:
	;
	if v118 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L27:
	;
	if v42&int32(3) == int32(0) {
		v98 = v38
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.Ui32(v98) <= base.Ui32(int32(3)) {
		v118 = v98
		goto L26
	} else {
		goto L34
	}
L29:
	;
	v83 = v38
	goto L30
L30:
	;
	if v83 == int32(0) {
		goto L10
	} else {
		goto L32
	}
L31:
	;
	v98 = v89
	goto L28
L32:
	;
	v89 = v83 + int32(-1)
	v90 = v27 + v89
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v89))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v92)
	if v90&int32(3) != 0 {
		v83 = v89
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v105 = v98
	goto L35
L35:
	;
	v109 = v105 + int32(-4)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v33+v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v27+v109))) = v112
	if base.Ui32(int32(3)) < base.Ui32(v109) {
		v105 = v109
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v118 = v109
	goto L26
L37:
	;
	goto L36
L38:
	;
	v125 = v118
	goto L39
L39:
	;
	v129 = v125 + int32(-1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v129))) = uint8(v132)
	if v129 != 0 {
		v125 = v129
		goto L39
	} else {
		goto L41
	}
L41:
	;
	goto L10
L42:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L43
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v146
	v148 = int32(4)
	v149 = v142 + v148
	v151 = v144 + v148
	v153 = v143 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v142 = v149
		v143 = v153
		v144 = v151
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v157 = v149
	v158 = v153
	v159 = v151
	goto L14
L45:
	;
	goto L44
L46:
	;
	v164 = v157
	v165 = v158
	v166 = v159
	goto L47
L47:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v168)
	v170 = int32(1)
	v175 = v165 + int32(-1)
	if v175 != 0 {
		v164 = v164 + v170
		v165 = v175
		v166 = v166 + v170
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L10
L49:
	;
	goto L48
L50:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v193 & int32(-257)
	return int32(0)
L51:
	;
	goto L4
}
func F_clusterNodeTimedOut(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	return v2 & int32(4)
}
func F_clusterParseSetSlotCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
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
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
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
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
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
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
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
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int64
	_ = v733
	var v739 int64
	_ = v739
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v17 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L1
L1:
	;
	v19 = v17 + int64(2000)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v21&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F__serverAssert(m, int32(_a348), int32(_a247), int32(7684))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L7
	} else {
		goto L210
	}
L3:
	;
	m.G0 = v14 + int32(112)
	return v755
L4:
	;
	v38 = int32(0)
	v41 = m.G0
	v43 = v41 - int32(16)
	m.G0 = v43
	v47 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v47 == v38 {
		v81 = v38
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v24 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+88)))
	if v27&int32(2) == v24 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_addReplyError(m, l0, int32(_a349))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v755 = v24
	goto L3
L9:
	;
	v92 = int32(0)
	v95 = m.G0
	v97 = v95 - int32(16)
	m.G0 = v97
	v101 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v101 == v92 {
		v133 = v92
		goto L25
	} else {
		goto L26
	}
L10:
	;
	if v81 == int32(0) {
		goto L9
	} else {
		goto L21
	}
L11:
	;
	m.G0 = v43 + int32(16)
	goto L10
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v51 == int32(0) {
		v81 = v38
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+uint32(_consts[171])))
	v56 = v43 + int32(8)
	F_listRewind(m, v54, v56)
	mBase = m.M
	v58 = int32(0)
	v61 = F_listNext(m, v56)
	mBase = m.M
	if v61 == v58 {
		v81 = v58
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v66 = v61
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v81 = v58
	goto L11
L17:
	;
	v79 = F_listNext(m, v43+int32(8))
	mBase = m.M
	if v79 != 0 {
		v66 = v79
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+156))
	if base.Ui32(v71+int32(-18)) <= base.Ui32(int32(2)) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v81 = int32(1)
	goto L11
L20:
	;
	goto L16
L21:
	;
	F_addReplyError(m, l0, int32(_a350))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v755 = int32(0)
	goto L3
L23:
	;
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+88)))
	if v146&int32(2) == v144 {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	if v133 == int32(0) {
		goto L23
	} else {
		goto L35
	}
L25:
	;
	m.G0 = v97 + int32(16)
	goto L24
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v105 == int32(0) {
		v133 = v92
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_consts[171])))
	v110 = v97 + int32(8)
	F_listRewind(m, v108, v110)
	mBase = m.M
	v112 = int32(0)
	v115 = F_listNext(m, v110)
	mBase = m.M
	if v115 == v112 {
		v133 = v112
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v120 = v115
	goto L29
L29:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v122 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v133 = v112
	goto L25
L31:
	;
	v131 = F_listNext(m, v97+int32(8))
	mBase = m.M
	if v131 != 0 {
		v120 = v131
		goto L29
	} else {
		goto L34
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+156))
	if base.Ui32(v123+int32(-18)) <= base.Ui32(int32(2)) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v133 = int32(1)
	goto L25
L34:
	;
	goto L30
L35:
	;
	F_addReplyError(m, l0, int32(_a351))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v755 = int32(0)
	goto L3
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v158 = F_getLongLongFromObject(m, v155, v14+int32(104))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L42
	}
L38:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if l0 != v152 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v168 = base.I32_wrap_i64(v160)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v171 = F_objectGetVal(m, v170)
	mBase = m.M
	v172 = int32(_a352)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v175 != 0 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	F_addReplyError(m, l0, int32(_a353))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L45
	}
L42:
	;
	if v158 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v14)+104))
	if base.Ui64(v160) < base.Ui64(int64(16384)) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v755 = int32(0)
	goto L3
L46:
	;
	if v650 <= v652 {
		v739 = v19
		goto L186
	} else {
		goto L187
	}
L47:
	;
	v644 = int32(5)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v644 < v646 {
		goto L183
	} else {
		goto L184
	}
L48:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v294 = F_objectGetVal(m, v293)
	mBase = m.M
	v295 = int32(_a354)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v298 != 0 {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	if v207-v209 != 0 {
		goto L48
	} else {
		goto L61
	}
L50:
	;
	v207 = F_tolower(m, v203)
	mBase = m.M
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	v209 = F_tolower(m, v208)
	mBase = m.M
	goto L49
L51:
	;
	v177 = v171
	v178 = v172
	v179 = v175
	goto L54
L52:
	;
	v203 = int32(0)
	v204 = v172
	goto L50
L53:
	;
	v203 = v200 & int32(255)
	v204 = v199
	goto L50
L54:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v181 == int32(0) {
		v199 = v178
		v200 = v179
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v199 = v193
	v200 = int32(0)
	goto L53
L56:
	;
	v185 = v179 & int32(255)
	if v185 == v181 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v192 = int32(1)
	v193 = v178 + v192
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v194 != 0 {
		v177 = v177 + v192
		v178 = v193
		v179 = v194
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v187 = F_tolower(m, v185)
	mBase = m.M
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v189 = F_tolower(m, v188)
	mBase = m.M
	if v187 == v189 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v199 = v178
	v200 = v191
	goto L53
L60:
	;
	goto L55
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v211 < int32(5) {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v214 = int32(0)
	v215 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+88)))
	if v216&int32(1) == v214 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+16))
	v237 = F_objectGetVal(m, v236)
	mBase = m.M
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v240 = F_objectGetVal(m, v239)
	mBase = m.M
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-1)))))
	switch v246 & int32(7) {
	case 0:
		goto L74
	case 1:
		goto L73
	case 2:
		goto L72
	case 3:
		goto L71
	case 4:
		goto L70
	default:
		v263 = int32(0)
		goto L69
	}
L64:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222+v168<<(uint(int32(2))%32))+52))
	if v226 == v215 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v168
	F_addReplyErrorFormat(m, l0, int32(_a355), v14+int32(16))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v755 = int32(0)
	goto L3
L67:
	;
	v276 = int32(0)
	v278 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v266 != v278 {
		goto L78
	} else {
		goto L79
	}
L68:
	;
	v266 = F_clusterLookupNode(m, v237, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L75
	}
L69:
	;
	v265 = v263
	goto L68
L70:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-17))))
	v263 = v262
	goto L69
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v240+int32(-9))))
	v265 = v259
	goto L68
L72:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240+int32(-5)))))
	v265 = v256
	goto L68
L73:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+int32(-3)))))
	v265 = v253
	goto L68
L74:
	;
	v265 = int32(base.Ui32(v246) >> (uint(int32(3)) % 32))
	goto L68
L75:
	;
	if v266 != 0 {
		goto L67
	} else {
		goto L76
	}
L76:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	v270 = F_objectGetVal(m, v269)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v270
	F_addReplyErrorFormat(m, l0, int32(_a356), v14)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v755 = int32(0)
	goto L3
L78:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+88)))
	if v283&int32(2) == int32(0) {
		goto L47
	} else {
		goto L81
	}
L79:
	;
	F_addReplyError(m, l0, int32(_a357))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	v755 = v276
	goto L3
L81:
	;
	F_addReplyError(m, l0, int32(_a358))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	v755 = int32(0)
	goto L3
L83:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v420 = F_objectGetVal(m, v419)
	mBase = m.M
	v421 = int32(_a359)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	if v424 != 0 {
		goto L124
	} else {
		goto L125
	}
L84:
	;
	if v330-v332 != 0 {
		goto L83
	} else {
		goto L96
	}
L85:
	;
	v330 = F_tolower(m, v326)
	mBase = m.M
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v332 = F_tolower(m, v331)
	mBase = m.M
	goto L84
L86:
	;
	v300 = v294
	v301 = v295
	v302 = v298
	goto L89
L87:
	;
	v326 = int32(0)
	v327 = v295
	goto L85
L88:
	;
	v326 = v323 & int32(255)
	v327 = v322
	goto L85
L89:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v304 == int32(0) {
		v322 = v301
		v323 = v302
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v322 = v316
	v323 = int32(0)
	goto L88
L91:
	;
	v308 = v302 & int32(255)
	if v308 == v304 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v315 = int32(1)
	v316 = v301 + v315
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v317 != 0 {
		v300 = v300 + v315
		v301 = v316
		v302 = v317
		goto L89
	} else {
		goto L95
	}
L93:
	;
	v310 = F_tolower(m, v308)
	mBase = m.M
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v312 = F_tolower(m, v311)
	mBase = m.M
	if v310 == v312 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v322 = v301
	v323 = v314
	goto L88
L95:
	;
	goto L90
L96:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v334 < int32(5) {
		goto L83
	} else {
		goto L97
	}
L97:
	;
	v337 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v168<<(uint(int32(2))%32))+52))
	v345 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v343 != v345 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+16))
	v355 = F_objectGetVal(m, v354)
	mBase = m.M
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	v358 = F_objectGetVal(m, v357)
	mBase = m.M
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358+int32(-1)))))
	switch v364 & int32(7) {
	case 0:
		goto L108
	case 1:
		goto L107
	case 2:
		goto L106
	case 3:
		goto L105
	case 4:
		goto L104
	default:
		v381 = int32(0)
		goto L103
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v168
	F_addReplyErrorFormat(m, l0, int32(_a360), v14+int32(32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	v755 = v337
	goto L3
L101:
	;
	v396 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v384 != v398 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v384 = F_clusterLookupNode(m, v355, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L109
	}
L103:
	;
	v383 = v381
	goto L102
L104:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v358+int32(-17))))
	v381 = v380
	goto L103
L105:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v358+int32(-9))))
	v383 = v377
	goto L102
L106:
	;
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358+int32(-5)))))
	v383 = v374
	goto L102
L107:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358+int32(-3)))))
	v383 = v371
	goto L102
L108:
	;
	v383 = int32(base.Ui32(v364) >> (uint(int32(3)) % 32))
	goto L102
L109:
	;
	if v384 != 0 {
		goto L101
	} else {
		goto L110
	}
L110:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	v388 = F_objectGetVal(m, v387)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v388
	F_addReplyErrorFormat(m, l0, int32(_a356), v14+int32(48))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v755 = int32(0)
	goto L3
L112:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+88)))
	if v403&int32(2) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	F_addReplyError(m, l0, int32(_a357))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v755 = v396
	goto L3
L115:
	;
	v412 = int32(5)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v412 < v414 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	F_addReplyError(m, l0, int32(_a358))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v755 = int32(0)
	goto L3
L118:
	;
	v417 = v412
	goto L120
L119:
	;
	v417 = int32(0)
	goto L120
L120:
	;
	v650 = v414
	v651 = v384
	v652 = v417
	goto L46
L121:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+12))
	v471 = F_objectGetVal(m, v470)
	mBase = m.M
	v472 = int32(_a361)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if v475 != 0 {
		goto L139
	} else {
		goto L140
	}
L122:
	;
	if v456-v458 != 0 {
		goto L121
	} else {
		goto L134
	}
L123:
	;
	v456 = F_tolower(m, v452)
	mBase = m.M
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	v458 = F_tolower(m, v457)
	mBase = m.M
	goto L122
L124:
	;
	v426 = v420
	v427 = v421
	v428 = v424
	goto L127
L125:
	;
	v452 = int32(0)
	v453 = v421
	goto L123
L126:
	;
	v452 = v449 & int32(255)
	v453 = v448
	goto L123
L127:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v430 == int32(0) {
		v448 = v427
		v449 = v428
		goto L126
	} else {
		goto L129
	}
L128:
	;
	v448 = v442
	v449 = int32(0)
	goto L126
L129:
	;
	v434 = v428 & int32(255)
	if v434 == v430 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v441 = int32(1)
	v442 = v427 + v441
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	if v443 != 0 {
		v426 = v426 + v441
		v427 = v442
		v428 = v443
		goto L127
	} else {
		goto L133
	}
L131:
	;
	v436 = F_tolower(m, v434)
	mBase = m.M
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	v438 = F_tolower(m, v437)
	mBase = m.M
	if v436 == v438 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	v448 = v427
	v449 = v440
	goto L126
L133:
	;
	goto L128
L134:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v460 < int32(4) {
		goto L121
	} else {
		goto L135
	}
L135:
	;
	v650 = v460
	v651 = int32(0)
	v652 = base.B2i32(v460 != int32(4)) << (uint(int32(2)) % 32)
	goto L46
L136:
	;
	F_addReplyError(m, l0, int32(_a362))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L7
	} else {
		goto L182
	}
L137:
	;
	if v507-v509 != 0 {
		goto L136
	} else {
		goto L149
	}
L138:
	;
	v507 = F_tolower(m, v503)
	mBase = m.M
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	v509 = F_tolower(m, v508)
	mBase = m.M
	goto L137
L139:
	;
	v477 = v471
	v478 = v472
	v479 = v475
	goto L142
L140:
	;
	v503 = int32(0)
	v504 = v472
	goto L138
L141:
	;
	v503 = v500 & int32(255)
	v504 = v499
	goto L138
L142:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v481 == int32(0) {
		v499 = v478
		v500 = v479
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v499 = v493
	v500 = int32(0)
	goto L141
L144:
	;
	v485 = v479 & int32(255)
	if v485 == v481 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v492 = int32(1)
	v493 = v478 + v492
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+1)))
	if v494 != 0 {
		v477 = v477 + v492
		v478 = v493
		v479 = v494
		goto L142
	} else {
		goto L148
	}
L146:
	;
	v487 = F_tolower(m, v485)
	mBase = m.M
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v489 = F_tolower(m, v488)
	mBase = m.M
	if v487 == v489 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v499 = v478
	v500 = v491
	goto L141
L148:
	;
	goto L143
L149:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v511 < int32(5) {
		goto L136
	} else {
		goto L150
	}
L150:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)+16))
	v516 = F_objectGetVal(m, v515)
	mBase = m.M
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+16))
	v519 = F_objectGetVal(m, v518)
	mBase = m.M
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+int32(-1)))))
	switch v525 & int32(7) {
	case 0:
		goto L158
	case 1:
		goto L157
	case 2:
		goto L156
	case 3:
		goto L155
	case 4:
		goto L154
	default:
		v542 = int32(0)
		goto L153
	}
L151:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+88)))
	if v557&int32(2) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L152:
	;
	v545 = F_clusterLookupNode(m, v516, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L7
	} else {
		goto L159
	}
L153:
	;
	v544 = v542
	goto L152
L154:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v519+int32(-17))))
	v542 = v541
	goto L153
L155:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v519+int32(-9))))
	v544 = v538
	goto L152
L156:
	;
	v535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v519+int32(-5)))))
	v544 = v535
	goto L152
L157:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+int32(-3)))))
	v544 = v532
	goto L152
L158:
	;
	v544 = int32(base.Ui32(v525) >> (uint(int32(3)) % 32))
	goto L152
L159:
	;
	if v545 != 0 {
		goto L151
	} else {
		goto L160
	}
L160:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+16))
	v549 = F_objectGetVal(m, v548)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v549
	F_addReplyErrorFormat(m, l0, int32(_a227), v14+int32(64))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	v755 = int32(0)
	goto L3
L162:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567+v168<<(uint(int32(2))%32))+52))
	v573 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v571 != v573 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	F_addReplyError(m, l0, int32(_a358))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L164
	}
L164:
	;
	v755 = int32(0)
	goto L3
L165:
	;
	v634 = int32(5)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v634 < v636 {
		goto L179
	} else {
		goto L180
	}
L166:
	;
	if v545 == v573 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v583 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(1) <= v583 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	if v620 == int32(0) {
		goto L165
	} else {
		goto L177
	}
L169:
	;
	goto L168
L170:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v589 = int32(0)
	v592 = v583
	v593 = v589
	v594 = v588
	v595 = v589
	goto L172
L171:
	;
	v620 = int32(0)
	goto L169
L172:
	;
	v598 = int32(0)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v594+v595<<(uint(int32(2))%32))))
	if v602 == v598 {
		v611 = v592
		v612 = v594
		v613 = v598
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v620 = v614
	goto L169
L174:
	;
	v614 = v613 + v593
	v616 = v595 + int32(1)
	if v616 < v611 {
		v592 = v611
		v593 = v614
		v594 = v612
		v595 = v616
		goto L172
	} else {
		goto L176
	}
L175:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	v606 = F_kvstoreHashtableSize(m, v605, v168)
	mBase = m.M
	v607 = int32(_a20)
	v608 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v610 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v611 = v608
	v612 = v610
	v613 = v606
	goto L174
L176:
	;
	goto L173
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v168
	F_addReplyErrorFormat(m, l0, int32(_a363), v14+int32(80))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L7
	} else {
		goto L178
	}
L178:
	;
	v755 = int32(0)
	goto L3
L179:
	;
	v639 = v634
	goto L181
L180:
	;
	v639 = int32(0)
	goto L181
L181:
	;
	v650 = v636
	v651 = v545
	v652 = v639
	goto L46
L182:
	;
	v755 = int32(0)
	goto L3
L183:
	;
	v649 = v644
	goto L185
L184:
	;
	v649 = int32(0)
	goto L185
L185:
	;
	v650 = v646
	v651 = v266
	v652 = v649
	goto L46
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v168
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v651
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v739
	v755 = int32(1)
	goto L3
L187:
	;
	v664 = v652
	goto L188
L188:
	;
	v666 = v664 + int32(1)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v667+v664<<(uint(int32(2))%32))))
	v672 = F_objectGetVal(m, v671)
	mBase = m.M
	v673 = int32(_a364)
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
	if v676 != 0 {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v733 = *(*int64)(unsafe.Add(mBase, uint32(v14)+96))
	v739 = v733
	goto L186
L190:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v666 < v731 {
		v664 = v666
		goto L188
	} else {
		goto L209
	}
L191:
	;
	if v708-v710 != 0 {
		goto L190
	} else {
		goto L203
	}
L192:
	;
	v708 = F_tolower(m, v704)
	mBase = m.M
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	v710 = F_tolower(m, v709)
	mBase = m.M
	goto L191
L193:
	;
	v678 = v672
	v679 = v673
	v680 = v676
	goto L196
L194:
	;
	v704 = int32(0)
	v705 = v673
	goto L192
L195:
	;
	v704 = v701 & int32(255)
	v705 = v700
	goto L192
L196:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if v682 == int32(0) {
		v700 = v679
		v701 = v680
		goto L195
	} else {
		goto L198
	}
L197:
	;
	v700 = v694
	v701 = int32(0)
	goto L195
L198:
	;
	v686 = v680 & int32(255)
	if v686 == v682 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v693 = int32(1)
	v694 = v679 + v693
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+1)))
	if v695 != 0 {
		v678 = v678 + v693
		v679 = v694
		v680 = v695
		goto L196
	} else {
		goto L202
	}
L200:
	;
	v688 = F_tolower(m, v686)
	mBase = m.M
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	v690 = F_tolower(m, v689)
	mBase = m.M
	if v688 == v690 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678))))
	v700 = v679
	v701 = v692
	goto L195
L202:
	;
	goto L197
L203:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v666 < v712 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v718+v666<<(uint(int32(2))%32))))
	v726 = F_getTimeoutFromObjectOrReply(m, l0, v722, v14+int32(96), int32(1))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L7
	} else {
		goto L207
	}
L205:
	;
	F_addReplyError(m, l0, int32(_a365))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L7
	} else {
		goto L206
	}
L206:
	;
	v755 = int32(0)
	goto L3
L207:
	;
	if v726 == int32(0) {
		goto L190
	} else {
		goto L208
	}
L208:
	;
	v755 = int32(0)
	goto L3
L209:
	;
	goto L189
L210:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterProcessGossipSection(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int64
	_ = v253
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v341 int32
	_ = v341
	var v344 int64
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v362 int64
	_ = v362
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
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
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int64
	_ = v428
	var v434 int64
	_ = v434
	var v440 int64
	_ = v440
	var v446 int64
	_ = v446
	var v452 int64
	_ = v452
	var v458 int64
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int64
	_ = v541
	var v547 int64
	_ = v547
	var v553 int64
	_ = v553
	var v559 int64
	_ = v559
	var v565 int64
	_ = v565
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v631 int32
	_ = v631
	var v640 int32
	_ = v640
	v14 = m.G0
	v16 = v14 - int32(416)
	m.G0 = v16
	v19 = l0 + int32(2256)
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v21 = F___bswap_16_2(m, v20)
	mBase = m.M
	goto L1
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v22 != 0 {
		v75 = v22
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v76 = F_verifyGossipSectionNodeIds(m, v19, v21)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L21
	}
L3:
	;
	v23 = int32(0)
	v25 = l0 + int32(40)
	goto L6
L4:
	;
	if int32(0)-v49 != 0 {
		v75 = v23
		goto L2
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	v34 = int32(0)
	goto L8
L7:
	;
	goto L5
L8:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v34))))
	v39 = int32(255)
	v49 = base.B2i32(base.Ui32((v36+int32(-123))&v39) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v36+int32(-58))&v39) < base.Ui32(int32(246)))
	if v49 != 0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v51 = v34 + int32(1)
	if v51 != int32(40) {
		v34 = v51
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v61 = F_sdsnewlen(m, v25, int32(40))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
	v66 = F_dictFind(m, v65, v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_sdsfree(m, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v66 == int32(0) {
		v75 = v23
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	goto L18
L18:
	;
	v75 = v72
	goto L2
L19:
	;
	m.G0 = v16 + int32(416)
	return
L20:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v75 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L21:
	;
	if v76 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v21 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v81 = v75 + int32(8)
	v89 = v19
	v90 = v21
	goto L24
L24:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+98)))
	v100 = F___bswap_16_2(m, v99)
	mBase = m.M
	goto L26
L25:
	;
	F__serverAssert(m, int32(_a258), int32(_a247), int32(2220))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L13
	} else {
		goto L155
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v102 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v138 != 0 {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	v103 = F_sdsempty(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v105 = F_representClusterNodeFlags(m, v103, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v108 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_sdsfree(m, v105)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L36
	}
L32:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+94)))
	v112 = F___bswap_16_2(m, v111)
	mBase = m.M
	goto L33
L33:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+96)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(64)))) = v105
	v115 = F___bswap_16_2(m, v113)
	mBase = m.M
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v112
	v118 = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v89 + v118
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v89
	F__serverLog(m, int32(0), int32(_a291), v16+v118)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	goto L27
L37:
	;
	v139 = int32(94)
	goto L39
L38:
	;
	v139 = int32(100)
	goto L39
L39:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89+v139))))
	v142 = F___bswap_16_2(m, v141)
	mBase = m.M
	goto L40
L40:
	;
	v143 = int32(1)
	if v138 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v146 = int32(100)
	goto L43
L42:
	;
	v146 = int32(94)
	goto L43
L43:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89+v146))))
	v149 = F___bswap_16_2(m, v148)
	mBase = m.M
	goto L44
L44:
	;
	goto L50
L45:
	;
	goto L25
L46:
	;
	v604 = v90 + int32(-1)
	if v604&int32(65535) != 0 {
		v89 = v89 + int32(104)
		v90 = v604
		goto L24
	} else {
		goto L154
	}
L47:
	;
	if v514 == int32(0) {
		goto L46
	} else {
		goto L140
	}
L48:
	;
	if int32(0)-v173 != 0 {
		v514 = v143
		goto L47
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	v158 = int32(0)
	goto L52
L51:
	;
	goto L49
L52:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v158))))
	v163 = int32(255)
	v173 = base.B2i32(base.Ui32((v160+int32(-123))&v163) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v160+int32(-58))&v163) < base.Ui32(int32(246)))
	if v173 != 0 {
		goto L51
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v175 = v158 + int32(1)
	if v175 != int32(40) {
		v158 = v175
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v185 = F_sdsnewlen(m, v89, int32(40))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+32))
	v190 = F_dictFind(m, v189, v185)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	F_sdsfree(m, v185)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	if v190 == int32(0) {
		v514 = v143
		goto L47
	} else {
		goto L60
	}
L60:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	goto L61
L61:
	;
	v198 = base.B2i32(v196 == int32(0))
	if v196 == int32(0) {
		v514 = v198
		goto L47
	} else {
		goto L62
	}
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v196 == v202 {
		v514 = v198
		goto L47
	} else {
		goto L63
	}
L63:
	;
	if v75 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v100&int32(12) != 0 {
		goto L93
	} else {
		goto L94
	}
L65:
	;
	if v100&int32(12) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v246 = v16 + int32(112)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2352))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+20)) = int32(128)
	v253 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v246)+12)) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v246)+296)) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v246)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v16 + int32(136)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+156)) = v16 + int32(280)
	goto L78
L67:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+88)))
	if v210&int32(1) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_markNodeAsFailingIfNeeded(m, v196)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L13
	} else {
		goto L77
	}
L69:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2160))
	if v215 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v218 = F_clusterNodeAddFailureReport(m, v196, v75)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	if v218 == int32(0) {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(1) < v223 {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v226 = F_humanNodename(m, v75)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	v228 = F_humanNodename(m, v196)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v196 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v81
	F__serverLog(m, int32(1), int32(_a292), v16+int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	goto L68
L77:
	;
	goto L64
L78:
	;
	v268 = int32(0)
	v270 = F_raxSeek(m, v16+int32(112), int32(_a67), v268, v268)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	goto L81
L80:
	;
	F_raxStop(m, v16+int32(112))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L13
	} else {
		goto L87
	}
L81:
	;
	v287 = F_raxNext(m, v16+int32(112))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L13
	} else {
		goto L83
	}
L82:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2352))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v16)+128))
	v297 = F_raxRemove(m, v294, v291, v295, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L13
	} else {
		goto L86
	}
L83:
	;
	if v287 == int32(0) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v16)+120))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	if v75 != v292 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	goto L80
L87:
	;
	if v287 == int32(0) {
		goto L64
	} else {
		goto L88
	}
L88:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(1) < v307 {
		goto L64
	} else {
		goto L89
	}
L89:
	;
	v310 = F_humanNodename(m, v75)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	v312 = F_humanNodename(m, v196)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v196 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v81
	F__serverLog(m, int32(1), int32(_a293), v16+int32(16))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	goto L64
L93:
	;
	if v100&int32(76) != 0 {
		goto L46
	} else {
		goto L103
	}
L94:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+88)))
	if v341&int32(172) != 0 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v196)+2184))
	if v344 != int64(0) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	F_clusterNodeCleanupFailureReports(m, v196)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2352))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v349)+8))
	goto L98
L98:
	;
	if base.I32_wrap_i64(v350) != 0 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v89)+44))
	v353 = F___bswap_32_2(m, v352)
	mBase = m.M
	goto L100
L100:
	;
	v356 = base.I64_extend_i32_u(v353) * int64(1000)
	v358 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	if v358+int64(500) < v356 {
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v196)+2192))
	if v356 <= v362 {
		goto L93
	} else {
		goto L102
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v196)+2192)) = v356
	goto L93
L103:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v196)+88))
	if v368&int32(12) == int32(0) {
		goto L46
	} else {
		goto L104
	}
L104:
	;
	v374 = v196 + int32(2256)
	v376 = v89 + int32(48)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v379 != 0 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2344))
	if v423 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L106:
	;
	if v411-v413 != 0 {
		goto L105
	} else {
		goto L118
	}
L107:
	;
	v411 = F_tolower(m, v407)
	mBase = m.M
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	v413 = F_tolower(m, v412)
	mBase = m.M
	goto L106
L108:
	;
	v381 = v374
	v382 = v376
	v383 = v379
	goto L111
L109:
	;
	v407 = int32(0)
	v408 = v376
	goto L107
L110:
	;
	v407 = v404 & int32(255)
	v408 = v403
	goto L107
L111:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v385 == int32(0) {
		v403 = v382
		v404 = v383
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v403 = v397
	v404 = int32(0)
	goto L110
L113:
	;
	v389 = v383 & int32(255)
	if v389 == v385 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v396 = int32(1)
	v397 = v382 + v396
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
	if v398 != 0 {
		v381 = v381 + v396
		v382 = v397
		v383 = v398
		goto L111
	} else {
		goto L117
	}
L115:
	;
	v391 = F_tolower(m, v389)
	mBase = m.M
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	v393 = F_tolower(m, v392)
	mBase = m.M
	if v391 == v393 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v403 = v382
	v404 = v395
	goto L110
L117:
	;
	goto L112
L118:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2328))
	if v415 != v142 {
		goto L105
	} else {
		goto L119
	}
L119:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2324))
	if v417 != v149 {
		goto L105
	} else {
		goto L120
	}
L120:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v196)+2332))
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+96)))
	v421 = F___bswap_16_2(m, v420)
	mBase = m.M
	goto L121
L121:
	;
	if v419 == v421 {
		goto L46
	} else {
		goto L122
	}
L122:
	;
	goto L105
L123:
	;
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	*(*int64)(unsafe.Add(mBase, uint32(v374))) = v428
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(86))))
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(2294)))) = v434
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(2288)))) = v440
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(2280)))) = v446
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(2272)))) = v452
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(56))))
	*(*int64)(unsafe.Add(mBase, uint32(v196+int32(2264)))) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v196)+2328)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v196)+2324)) = v149
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+96)))
	v463 = F___bswap_16_2(m, v462)
	mBase = m.M
	goto L126
L124:
	;
	F_freeClusterLink(m, v423)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196)+2332)) = v463
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v196)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+88)) = v465 & int32(-65)
	v470 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v470 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v494 = int32(0)
	v495 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+88)))
	if v496&int32(2) == v494 {
		goto L46
	} else {
		goto L134
	}
L128:
	;
	v473 = F_humanNodename(m, v196)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L13
	} else {
		goto L129
	}
L129:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v478 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v479 = int32(2328)
	goto L132
L131:
	;
	v479 = int32(2324)
	goto L132
L132:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v196+v479)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v196 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v481
	F__serverLog(m, int32(2), int32(_a294), v16)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	goto L127
L134:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v495)+2172))
	if v501 != v196 {
		goto L46
	} else {
		goto L135
	}
L135:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	if v506 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v507 = int32(2328)
	goto L138
L137:
	;
	v507 = int32(2324)
	goto L138
L138:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v196+v507)))
	v510 = int32(0)
	F_replicationSetPrimary(m, v374, v509, v510, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	goto L46
L140:
	;
	if v75 == int32(0) {
		goto L46
	} else {
		goto L141
	}
L141:
	;
	if v100&int32(64) != 0 {
		goto L46
	} else {
		goto L142
	}
L142:
	;
	v524 = F_sdsnewlen(m, v89, int32(40))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	F_clusterBlacklistCleanup(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)+40))
	v531 = F_dictFind(m, v530, v524)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	F_sdsfree(m, v524)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	if v531 != 0 {
		goto L46
	} else {
		goto L147
	}
L147:
	;
	v535 = F_createClusterNode(m, v89, v100)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(86))))
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(2294)))) = v541
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(2288)))) = v547
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(2280)))) = v553
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(64))))
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(2272)))) = v559
	v565 = *(*int64)(unsafe.Add(mBase, uint32(v89+int32(56))))
	*(*int64)(unsafe.Add(mBase, uint32(v535+int32(2264)))) = v565
	v567 = *(*int64)(unsafe.Add(mBase, uint32(v89)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v535)+2256)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(v535)+2328)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v535)+2324)) = v149
	v571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+96)))
	v572 = F___bswap_16_2(m, v571)
	mBase = m.M
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+2332)) = v572
	v575 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)+32))
	v580 = F_sdsnewlen(m, v535+int32(8), int32(40))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	v582 = F_dictAdd(m, v576, v580, v535)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	if v582 != 0 {
		goto L45
	} else {
		goto L152
	}
L152:
	;
	F_clusterAddNodeToShard(m, v535+int32(48), v535)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	goto L46
L154:
	;
	goto L19
L155:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	if int32(3) < v614 {
		goto L19
	} else {
		goto L161
	}
L157:
	;
	if int32(3) < v614 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	v619 = F_humanNodename(m, v75)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v619
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v75 + int32(8)
	F__serverLog(m, int32(3), int32(_a295), v16+int32(96))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	goto L19
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v76
	F__serverLog(m, int32(3), int32(_a296), v16+int32(80))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	goto L19
}
func F_clusterProcessModulePacket(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		return
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = F___bswap_32_2(m, v9)
		mBase = m.M
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		v15 = F_sdsnewlen(m, l1+int32(8), int32(40))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_moduleCallClusterReceivers(m, v15, v8, v11, l0+int32(13), v10)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_sdsfree(m, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_clusterPromoteSelfToPrimary(m *base.Module) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	F_replicationUnsetPrimary(m)
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		v3 = F_verifyClusterConfigWithData(m)
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterRDBLoadSlotImport(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_listCreate(m)
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(102)
	v20 = F_rdbLoadStringObject(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v12 + int32(16)
	return v162
L4:
	;
	v162 = int32(-1)
	goto L3
L5:
	;
	F_listRelease(m, v14)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L45
	}
L6:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = F_objectGetVal(m, v20)
	mBase = m.M
	v25 = int32(-1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v25))))
	switch v27&int32(7) + v25 {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	default:
		goto L10
	}
L8:
	;
	F_decrRefCount(m, v20)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L9:
	;
	v58 = F_rdbLoadLen(m, l0, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v49 {
		goto L8
	} else {
		goto L17
	}
L11:
	;
	if v44 == int32(40) {
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
	v44 = v43
	goto L11
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
	v44 = v40
	goto L11
L14:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
	v44 = v37
	goto L11
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
	v44 = v34
	goto L11
L16:
	;
	goto L10
L17:
	;
	F__serverLog(m, int32(3), int32(_a367), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	v123 = int32(0)
	v126 = F_objectGetVal(m, v20)
	mBase = m.M
	v127 = F_createSlotImportJob(m, v123, v123, v126, v14)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L41
	}
L20:
	;
	switch base.I32_wrap_i64(v61) {
	default:
		goto L8
	case 1:
		goto L19
	}
L21:
	;
	v60 = int64(1)
	v61 = v58 + v60
	if base.Ui64(v61) <= base.Ui64(v60) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v72 = int64(0)
	goto L23
L23:
	;
	v75 = F_rdbLoadLen(m, l0, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v104 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L38
	}
L26:
	;
	F_decrRefCount(m, v20)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L36
	}
L27:
	;
	if v75 == int64(-1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v80 = F_rdbLoadLen(m, l0, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v80 == int64(-1) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	if base.Ui64(v80) < base.Ui64(v75) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v89 {
		goto L26
	} else {
		goto L34
	}
L32:
	;
	if base.Ui64(v80|v75) < base.Ui64(int64(16384)) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v75
	F__serverLog(m, int32(3), int32(_a368), v12)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L26
L36:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L5
L38:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v104)+4)) = uint32(v80)
	*(*uint32)(unsafe.Add(mBase, uint32(v104))) = uint32(v75)
	v108 = F_listAddNodeTail(m, v14, v104)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v111 = v72 + int64(1)
	if v111 != v58 {
		v72 = v111
		goto L23
	} else {
		goto L40
	}
L40:
	;
	goto L19
L41:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_consts[171])))
	v132 = F_listAddNodeTail(m, v131, v127)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_decrRefCount(m, v20)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v162 = v123
	goto L3
L44:
	;
	goto L5
L45:
	;
	goto L4
}
func F_clusterRDBSaveSlotImports(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v13 == v3 {
		v166 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v166
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[171])))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v19 == int32(0) {
		v166 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v9 + int32(24)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v24
	goto L4
L4:
	;
	v29 = v9 + int32(24)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v31 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v166 = int32(0)
	goto L1
L6:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.B2i32(v34 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v40
	goto L7
L9:
	;
	v50 = v31
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+156))
	if base.Ui32(v53+int32(-18)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v143 = v9 + int32(24)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v145 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v58 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if int32(79) < l1 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = F_rdbSaveType(m, l0, int32(243))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L21
	}
L16:
	;
	v61 = int32(-1)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v63 {
		v166 = v61
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F__serverLog(m, int32(3), int32(_a366), v9)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v166 = v61
	goto L1
L20:
	;
	v82 = F_rdbSaveRawString(m, l0, v52+int32(112), int32(40))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L24
	}
L21:
	;
	if int32(0) <= v74 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v166 = int32(-1)
	goto L1
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+164))
	v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v87)+20)))
	v89 = F_rdbSaveLen(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L27
	}
L24:
	;
	if int32(0) <= v82 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v166 = int32(-1)
	goto L1
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v52)+164))
	v96 = v9 + int32(16)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v97
	goto L29
L27:
	;
	if int32(0) <= v89 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v166 = int32(-1)
	goto L1
L29:
	;
	goto L30
L30:
	;
	v108 = v9 + int32(16)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v110 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v110 == int32(0) {
		goto L12
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v110+base.B2i32(v113 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v119
	goto L33
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v124 = int64(*(*int32)(unsafe.Add(mBase, uint32(v123))))
	v125 = F_rdbSaveLen(m, l0, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L18
	} else {
		goto L37
	}
L36:
	;
	v131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v123)+4)))
	v132 = F_rdbSaveLen(m, l0, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	if int32(0) <= v125 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v166 = int32(-1)
	goto L1
L39:
	;
	if int32(-1) < v132 {
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v166 = int32(-1)
	goto L1
L41:
	;
	if v145 != 0 {
		v50 = v145
		goto L10
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v145+base.B2i32(v148 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v154
	goto L42
L44:
	;
	goto L11
}
func F_clusterReadHandler(m *base.Module, l0 int32) {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
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
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	v9 = m.G0
	v11 = v9 - int32(4432)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v18 = v14
	goto L3
L1:
	;
	m.G0 = v11 + int32(4432)
	return
L2:
	;
	F_freeClusterLink(m, v13)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L22
	} else {
		goto L72
	}
L3:
	;
	if base.Ui32(int32(13)) < base.Ui32(v18) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+76))
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32) int32)(m, l0, v11+int32(80), v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L22
	} else {
		goto L36
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v18 == int32(14) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v93 = int32(14) - v18
	goto L5
L8:
	;
	v87 = F___bswap_32_2(m, v85)
	mBase = m.M
	goto L29
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v31 != int32(1651327826) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v85 = v30
	goto L8
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v48 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	v35 = F___bswap_16_2(m, v34)
	mBase = m.M
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v37 = F___bswap_32_2(m, v36)
	mBase = m.M
	goto L14
L14:
	;
	if int32(-1) < base.I32_extend16_s(v35) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(2256)
	goto L17
L16:
	;
	v43 = int32(16)
	goto L17
L17:
	;
	if base.Ui32(v43) <= base.Ui32(v37) {
		v85 = v36
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v71 {
		goto L2
	} else {
		goto L27
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v62 {
		goto L2
	} else {
		goto L25
	}
L21:
	;
	v57 = m.T0[v48].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, v11+int32(32), int32(46), v11+int32(28), int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return
L23:
	;
	if v57 != int32(-1) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	F__serverLog(m, int32(3), int32(_a284), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L2
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(32)
	F__serverLog(m, int32(3), int32(_a285), v11+int32(16))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L2
L29:
	;
	v88 = v87 - v18
	v89 = int32(4352)
	if base.Ui32(v88) < base.Ui32(v89) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v92 = v88
	goto L32
L31:
	;
	v92 = v89
	goto L32
L32:
	;
	v93 = v92
	goto L5
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if base.Ui32(v99) <= base.Ui32(v146-v147) {
		v170 = v145
		v171 = v147
		goto L54
	} else {
		goto L55
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v113 {
		goto L2
	} else {
		goto L40
	}
L35:
	;
	if int32(0) < v99 {
		goto L33
	} else {
		goto L39
	}
L36:
	;
	if v99 != int32(-1) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v104 == int32(3) {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v111 = int32(0)
	goto L34
L39:
	;
	v111 = base.B2i32(v99 == int32(0))
	goto L34
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v117 == int32(0) {
		v125 = int32(_a277)
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v116 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v123 = F_humanNodename(m, v117)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	v125 = v123
	goto L41
L44:
	;
	v128 = int32(_a275)
	goto L46
L45:
	;
	v128 = int32(_a276)
	goto L46
L46:
	;
	if v117 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v130 = v117 + int32(8)
	goto L49
L48:
	;
	v130 = int32(_a277)
	goto L49
L49:
	;
	if v111 != 0 {
		v136 = int32(_a286)
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v130
	F__serverLog(m, int32(0), int32(_a287), v11)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L22
	} else {
		goto L53
	}
L51:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
	v134 = m.T0[v133].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L22
	} else {
		goto L52
	}
L52:
	;
	v136 = v134
	goto L50
L53:
	;
	goto L2
L54:
	;
	if v99 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v150 = v147 + v99
	v153 = int32(1048576)
	if base.Ui32(v150) < base.Ui32(v153) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v157 = v150 << (uint(int32(1)) % 32)
	goto L58
L57:
	;
	v157 = v150 + v153
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v157
	v159 = F_valkey_realloc(m, v145, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v159
	v162 = int32(_a20)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v166 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v163 - v146 + v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v170 = v159
	v171 = v169
	goto L54
L60:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	v180 = v179 + v99
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v180
	v182 = v99 + v18
	if base.Ui32(v182) < base.Ui32(int32(14)) {
		v18 = v180
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v177 = F__emscripten_memcpy_bulkmem(m, v170+v171, v11+int32(80), v99)
	mBase = m.M
	goto L61
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v187 = F___bswap_32_2(m, v186)
	mBase = m.M
	goto L64
L64:
	;
	if v182 != v187 {
		v18 = v180
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v189 = F_clusterProcessPacket(m, v13)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L22
	} else {
		goto L66
	}
L66:
	;
	if v189 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if base.Ui32(v193) < base.Ui32(int32(1025)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v212
	v18 = v212
	goto L3
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	F_valkey_free(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	v199 = int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v199
	v202 = F_valkey_malloc(m, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v202
	v205 = int32(_a20)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v209 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v206 - v193 + v209
	goto L68
L72:
	;
	goto L1
}
func F_clusterRedirectClient(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	switch l3 + int32(-1) {
	case 0:
		F_addReplyError(m, l0, int32(_a233))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	case 1:
		F_addReplyError(m, l0, int32(_a234))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	default:
		if base.Ui32(int32(1)) < base.Ui32(l3+int32(-3)) {
			F__serverPanic_1(m, int32(_a203), int32(1342), int32(_a235), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, _consts[14]))
			if v31 == int32(0) {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[107]))
				v44 = v43
				if v44 == int32(0) {
					if v44 == int32(0) {
						if l0 == int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
							v61 = v60
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
							if v58 != 0 {
								v61 = v58
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
								v61 = v60
							}
						}
						v63 = v61
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
						v63 = v55
					}
				} else {
					if l0 == int32(0) {
						if v44 == int32(0) {
							if l0 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
								v61 = v60
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
								if v58 != 0 {
									v61 = v58
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
									v61 = v60
								}
							}
							v63 = v61
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
							v63 = v55
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
						if v49 == int32(0) {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
							v63 = v55
						} else {
							v61 = v49
							v63 = v61
						}
					}
				}
				v64 = F_sdsempty(m)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = F_clusterNodePreferredEndpoint(m, l1, l0)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
						if l3 == int32(3) {
							v75 = int32(_a236)
						} else {
							v75 = int32(_a237)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v75
						v78 = F_sdscatprintf(m, v64, int32(_a238), v10)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_addReplyErrorSds(m, l0, v78)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v34 == int32(0) {
					v43 = *(*int32)(unsafe.Add(mBase, _consts[107]))
					v44 = v43
					if v44 == int32(0) {
						if v44 == int32(0) {
							if l0 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
								v61 = v60
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
								if v58 != 0 {
									v61 = v58
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
									v61 = v60
								}
							}
							v63 = v61
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
							v63 = v55
						}
					} else {
						if l0 == int32(0) {
							if v44 == int32(0) {
								if l0 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
									v61 = v60
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
									if v58 != 0 {
										v61 = v58
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
										v61 = v60
									}
								}
								v63 = v61
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
								v63 = v55
							}
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
							if v49 == int32(0) {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
								v63 = v55
							} else {
								v61 = v49
								v63 = v61
							}
						}
					}
					v64 = F_sdsempty(m)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = F_clusterNodePreferredEndpoint(m, l1, l0)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
							if l3 == int32(3) {
								v75 = int32(_a236)
							} else {
								v75 = int32(_a237)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v75
							v78 = F_sdscatprintf(m, v64, int32(_a238), v10)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								F_addReplyErrorSds(m, l0, v78)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v38 = F_connectionTypeTls(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v44 = base.B2i32(v37 == v38)
						if v44 == int32(0) {
							if v44 == int32(0) {
								if l0 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
									v61 = v60
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
									if v58 != 0 {
										v61 = v58
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
										v61 = v60
									}
								}
								v63 = v61
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
								v63 = v55
							}
						} else {
							if l0 == int32(0) {
								if v44 == int32(0) {
									if l0 == int32(0) {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
										v61 = v60
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2336))
										if v58 != 0 {
											v61 = v58
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2324))
											v61 = v60
										}
									}
									v63 = v61
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
									v63 = v55
								}
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2340))
								if v49 == int32(0) {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+2328))
									v63 = v55
								} else {
									v61 = v49
									v63 = v61
								}
							}
						}
						v64 = F_sdsempty(m)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v66 = F_clusterNodePreferredEndpoint(m, l1, l0)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
								if l3 == int32(3) {
									v75 = int32(_a236)
								} else {
									v75 = int32(_a237)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v75
								v78 = F_sdscatprintf(m, v64, int32(_a238), v10)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									F_addReplyErrorSds(m, l0, v78)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	case 4:
		F_addReplyError(m, l0, int32(_a239))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	case 5:
		F_addReplyError(m, l0, int32(_a240))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	case 6:
		F_addReplyError(m, l0, int32(_a241))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_clusterSendFail(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v18 int64
	_ = v18
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = F_createClusterMsgSendBlock(m, int32(3), int32(2296))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(32))))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(2296)))) = v12
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24))))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(2288)))) = v18
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(16))))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(2280)))) = v24
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(2272)))) = v30
		v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+2264)) = v32
		F_clusterBroadcastMessage(m, v6)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v38 = v36 + int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v38
			if v36 <= int32(0) {
				F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if v38 != 0 {
					return
				} else {
					v42 = int32(_a20)
					v44 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v44 - v45
					F_valkey_free(m, v6)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_clusterSendFailoverAuth(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
	if v4 == int32(0) {
		return
	} else {
		v9 = F_createClusterMsgSendBlock(m, int32(6), int32(2256))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2344))
			F_clusterSendMessage(m, v11, v9)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v16 = v14 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
				if v14 <= int32(0) {
					F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if v16 != 0 {
						return
					} else {
						v20 = int32(_a20)
						v22 = *(*int32)(unsafe.Add(mBase, _consts[161]))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						*(*int32)(unsafe.Add(mBase, _consts[161])) = v22 - v23
						F_valkey_free(m, v9)
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
func F_clusterSendMessage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
		if v7 != 0 {
			v19 = v6
			v20 = F_listAddNodeTail(m, v19, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 + int32(1)
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v28 = int32(12)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v26 + base.I64_extend_i32_u(v27+v28)
				v33 = int32(_a20)
				v35 = *(*int32)(unsafe.Add(mBase, _consts[161]))
				*(*int32)(unsafe.Add(mBase, _consts[161])) = v35 + v28
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
				v40 = F___bswap_16_2(m, v39)
				mBase = m.M
				v42 = v40 & int32(32767)
				if base.Ui32(int32(10)) < base.Ui32(v42) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[111]))
					v49 = v46 + v42<<(uint(int32(3))%32)
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182])))
					*(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182]))) = v50 + int64(1)
				}
				return
			}
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v8 == int32(0) {
				v19 = v6
				v20 = F_listAddNodeTail(m, v19, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 + int32(1)
					v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v28 = int32(12)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v26 + base.I64_extend_i32_u(v27+v28)
					v33 = int32(_a20)
					v35 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v35 + v28
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
					v40 = F___bswap_16_2(m, v39)
					mBase = m.M
					v42 = v40 & int32(32767)
					if base.Ui32(int32(10)) < base.Ui32(v42) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v49 = v46 + v42<<(uint(int32(3))%32)
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182])))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182]))) = v50 + int64(1)
					}
					return
				}
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
				v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, v11, int32(65), int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = v18
					v20 = F_listAddNodeTail(m, v19, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22 + int32(1)
						v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v28 = int32(12)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v26 + base.I64_extend_i32_u(v27+v28)
						v33 = int32(_a20)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[161]))
						*(*int32)(unsafe.Add(mBase, _consts[161])) = v35 + v28
						v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
						v40 = F___bswap_16_2(m, v39)
						mBase = m.M
						v42 = v40 & int32(32767)
						if base.Ui32(int32(10)) < base.Ui32(v42) {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, _consts[111]))
							v49 = v46 + v42<<(uint(int32(3))%32)
							v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182])))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[182]))) = v50 + int64(1)
						}
						return
					}
				}
			}
		}
	}
}
func F_clusterSendModule(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
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
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	v3 = l2
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v38 = F_clusterNodeIterNext(m, v14)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v24
	v27 = v14 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v27)+4)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = int64(1)
	goto L4
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(2)
	goto L1
L4:
	;
	goto L1
L5:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	switch v123 {
	case 0:
		goto L37
	default:
		goto L35
	case 2:
		goto L36
	}
L6:
	;
	return
L7:
	;
	if v38 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v44 = v38
	goto L9
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	if v55&int32(48) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v110 = F_clusterNodeIterNext(m, v14)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L33
	}
L12:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2344))
	if v59 == v58 {
		v69 = v58
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v74 = v14 + int32(40) | v69<<(uint(int32(2))%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v75 != 0 {
		v101 = v75
		v103 = v59
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v44)+2192))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	if v62 < v63 {
		v69 = v58
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v69 = int32(base.Ui32(v55)>>(uint(int32(12))%32)) & int32(1)
	goto L13
L16:
	;
	F_clusterSendMessage(m, v103, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L32
	}
L17:
	;
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = int32(32777)
	goto L20
L19:
	;
	v78 = int32(9)
	goto L20
L20:
	;
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = int32(16)
	goto L23
L22:
	;
	v81 = int32(2256)
	goto L23
L23:
	;
	v83 = F_createClusterMsgSendBlock(m, v78, l4+int32(13)+v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	if v69 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v87 = int32(24)
	goto L27
L26:
	;
	v87 = int32(2264)
	goto L27
L27:
	;
	v88 = v83 + v87
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+12)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = l1
	v91 = F___bswap_32_1(m, l4)
	mBase = m.M
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v91
	if l4 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v83
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2344))
	v101 = v83
	v103 = v100
	goto L16
L30:
	;
	goto L29
L31:
	;
	v97 = F__emscripten_memcpy_bulkmem(m, v88+int32(13), l3, l4)
	mBase = m.M
	goto L30
L32:
	;
	goto L11
L33:
	;
	if v110 != 0 {
		v44 = v110
		goto L9
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v130 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(0)
	goto L35
L37:
	;
	F_dictResetIterator(m, v14+int32(8))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L50
	}
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v149 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v135 = v133 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v135
	if v133 <= int32(0) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	if v135 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(_a20)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v141 - v142
	F_valkey_free(m, v130)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	m.G0 = v14 + int32(48)
	return
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v154 = v152 + int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v154
	if v152 < int32(1) {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	if v154 != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v158 = int32(_a20)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v160 - v161
	F_valkey_free(m, v149)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterSendUpdate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	if l0 == int32(0) {
		return
	} else {
		v12 = F_createClusterMsgSendBlock(m, int32(7), int32(4352))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(40))))
			*(*int64)(unsafe.Add(mBase, uint32(v12+int32(2304)))) = v18
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(32))))
			*(*int64)(unsafe.Add(mBase, uint32(v12+int32(2296)))) = v24
			v30 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(24))))
			*(*int64)(unsafe.Add(mBase, uint32(v12+int32(2288)))) = v30
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l1+int32(16))))
			*(*int64)(unsafe.Add(mBase, uint32(v12+int32(2280)))) = v36
			v38 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v12)+2272)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
			v41 = int64(56)
			v43 = int64(65280)
			v45 = int64(40)
			v48 = int64(16711680)
			v50 = int64(24)
			v52 = int64(4278190080)
			v54 = int64(8)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+2264)) = v40<<(uint(v41)%64) | v40&v43<<(uint(v45)%64) | (v40&v48<<(uint(v50)%64) | v40&v52<<(uint(v54)%64)) | (int64(base.Ui64(v40)>>(uint(v54)%64))&v52 | int64(base.Ui64(v40)>>(uint(v50)%64))&v48 | (int64(base.Ui64(v40)>>(uint(v45)%64))&v43 | int64(base.Ui64(v40)>>(uint(v41)%64))))
			v84 = F__emscripten_memcpy_bulkmem(m, v12+int32(2312), l1+int32(104), int32(2048))
			mBase = m.M
			v88 = int32(0)
			for {
				v94 = v84 + v88
				v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				v96 = int32(_a20)
				v97 = *(*int32)(unsafe.Add(mBase, _consts[111]))
				v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v88)+uint32(_consts[162]))))
				v100 = int32(-1)
				v102 = v95 & (v99 ^ v100)
				*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v102)
				v105 = v88 | int32(1)
				v106 = v84 + v105
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
				v109 = *(*int32)(unsafe.Add(mBase, _consts[111]))
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v105)+uint32(_consts[162]))))
				v114 = v107 & (v111 ^ v100)
				*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v114)
				v117 = v88 + int32(2)
				if v117 != int32(2048) {
					v88 = v117
					continue
				} else {
					break
				}
				break
			}
			v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
			if v121 != 0 {
				v133 = v120
				v134 = F_listAddNodeTail(m, v133, v12)
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v139 = int32(12)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v137 + base.I64_extend_i32_u(v138+v139)
					v144 = int32(_a20)
					v146 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					v148 = v146 + v139
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v148
					v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
					v151 = F___bswap_16_2(m, v150)
					mBase = m.M
					v153 = v151 & int32(32767)
					if base.Ui32(int32(10)) < base.Ui32(v153) {
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v160 = v157 + v153<<(uint(int32(3))%32)
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182])))
						*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182]))) = v161 + int64(1)
					}
					if v136 <= int32(-1) {
						F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if v136 != 0 {
							return
						} else {
							v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							*(*int32)(unsafe.Add(mBase, _consts[161])) = v148 - v169
							F_valkey_free(m, v12)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				if v122 == int32(0) {
					v133 = v120
					v134 = F_listAddNodeTail(m, v133, v12)
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v139 = int32(12)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v137 + base.I64_extend_i32_u(v138+v139)
						v144 = int32(_a20)
						v146 = *(*int32)(unsafe.Add(mBase, _consts[161]))
						v148 = v146 + v139
						*(*int32)(unsafe.Add(mBase, _consts[161])) = v148
						v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
						v151 = F___bswap_16_2(m, v150)
						mBase = m.M
						v153 = v151 & int32(32767)
						if base.Ui32(int32(10)) < base.Ui32(v153) {
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, _consts[111]))
							v160 = v157 + v153<<(uint(int32(3))%32)
							v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182])))
							*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182]))) = v161 + int64(1)
						}
						if v136 <= int32(-1) {
							F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v136 != 0 {
								return
							} else {
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								*(*int32)(unsafe.Add(mBase, _consts[161])) = v148 - v169
								F_valkey_free(m, v12)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+80))
					v130 = m.T0[v129].(func(*base.Module, int32, int32, int32) int32)(m, v125, int32(65), int32(1))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v133 = v132
						v134 = F_listAddNodeTail(m, v133, v12)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v139 = int32(12)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v137 + base.I64_extend_i32_u(v138+v139)
							v144 = int32(_a20)
							v146 = *(*int32)(unsafe.Add(mBase, _consts[161]))
							v148 = v146 + v139
							*(*int32)(unsafe.Add(mBase, _consts[161])) = v148
							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)))
							v151 = F___bswap_16_2(m, v150)
							mBase = m.M
							v153 = v151 & int32(32767)
							if base.Ui32(int32(10)) < base.Ui32(v153) {
							} else {
								v157 = *(*int32)(unsafe.Add(mBase, _consts[111]))
								v160 = v157 + v153<<(uint(int32(3))%32)
								v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182])))
								*(*int64)(unsafe.Add(mBase, uint32(v160)+uint32(_consts[182]))) = v161 + int64(1)
							}
							if v136 <= int32(-1) {
								F__serverAssert(m, int32(_a290), int32(_a247), int32(1756))
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if v136 != 0 {
									return
								} else {
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									*(*int32)(unsafe.Add(mBase, _consts[161])) = v148 - v169
									F_valkey_free(m, v12)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
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
func F_clusterSetGossipEntry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v18 int64
	_ = v18
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v78 int64
	_ = v78
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v7 = l0 + l1*int32(104)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2288)))) = v12
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2280)))) = v18
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2272)))) = v24
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2264)))) = v30
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2256)))) = v34
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l2)+2184))
	v40 = base.I64_div_s(v38, int64(1000))
	v42 = F___bswap_32_1(m, base.I32_wrap_i64(v40))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(2296)))) = v42
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l2)+2192))
	v48 = base.I64_div_s(v46, int64(1000))
	v50 = F___bswap_32_1(m, base.I32_wrap_i64(v48))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(2300)))) = v50
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+2256))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2304)))) = v54
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2264))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2312)))) = v60
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2272))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2320)))) = v66
	v68 = int32(2328)
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2280))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+v68))) = v72
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2288))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2336)))) = v78
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(2294))))
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(2342)))) = v84
	v89 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v89 != 0 {
		v90 = v68
	} else {
		v90 = int32(2324)
	}
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v90))))
	v93 = F___bswap_16_1(m, v92)
	mBase = m.M
	if v89 != 0 {
		v98 = int32(2324)
	} else {
		v98 = int32(2328)
	}
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+v98))))
	v101 = F___bswap_16_1(m, v100)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(2356)))) = uint16(v101)
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(2350)))) = uint16(v93)
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2332)))
	v109 = F___bswap_16_1(m, v108)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(2352)))) = uint16(v109)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+88)))
	v114 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(2358)))) = uint16(v114)
	v118 = F___bswap_16_1(m, v111)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(2354)))) = uint16(v118)
	return
}
func F_clusterSlotByCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(2064)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v12)+4)) = int64(1099511627776)
	v21 = F_getKeysFromCommand(m, l0, l1, l2, v12+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_getKeysFreeResult(m, v12+int32(4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L42
	}
L2:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v153 | v149
	v164 = int32(-1)
	goto L1
L3:
	;
	v26 = int32(-1)
	if v21 < int32(1) {
		v164 = v26
		goto L1
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	if v21 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v149 = int32(524288)
	goto L2
L7:
	;
	v34 = v5
	v36 = v26
	goto L8
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(3))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1+v43<<(uint(int32(2))%32))))
	v48 = F_objectGetVal(m, v47)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-1)))))
	switch v51 & int32(7) {
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
		v68 = int32(0)
		goto L10
	}
L10:
	;
	v69 = int32(0)
	if v68 < int32(1) {
		v89 = v69
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-17))))
	v68 = v67
	goto L10
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(-9))))
	v68 = v64
	goto L10
L13:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48+int32(-5)))))
	v68 = v61
	goto L10
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+int32(-3)))))
	v68 = v58
	goto L10
L15:
	;
	v68 = int32(base.Ui32(v51) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	if v36 != int32(-1) {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	v135 = v131 & int32(16383)
	goto L16
L18:
	;
	v100 = v89 + int32(1)
	if v68 <= v100 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v98 = F_crc16(m, v48, v68)
	mBase = m.M
	v131 = v98
	goto L17
L20:
	;
	if v89 != v68 {
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v77 = v69
	goto L22
L22:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v77))))
	if v81 == int32(123) {
		v89 = v77
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v85 = v77 + int32(1)
	if v85 != v68 {
		v77 = v85
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L19
L26:
	;
	goto L19
L27:
	;
	v128 = F_crc16(m, v48+v89+int32(1), v106+(v89^int32(-1)))
	mBase = m.M
	v131 = v128
	goto L17
L28:
	;
	v121 = F_crc16(m, v48, v68)
	mBase = m.M
	v131 = v121
	goto L17
L29:
	;
	v106 = v100
	goto L31
L30:
	;
	if v106 == v68 {
		goto L28
	} else {
		goto L35
	}
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v106))))
	if v108 == int32(125) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v112 = v106 + int32(1)
	if v112 != v68 {
		v106 = v112
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	if v106 != v100 {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	v142 = v34 + int32(1)
	if v142 != v21 {
		v34 = v142
		v36 = v140
		goto L8
	} else {
		goto L41
	}
L38:
	;
	if v135 == v36 {
		v140 = v36
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v140 = v135
	goto L37
L40:
	;
	v149 = int32(1048576)
	goto L2
L41:
	;
	v164 = v140
	goto L1
L42:
	;
	m.G0 = v12 + int32(2064)
	return v164
}
func F_clusterSlotStatsAddCpuDuration(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v4 == int32(-1) {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[229]))
		if v8 == int32(0) {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[110]))
			if v12 == int32(0) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[177]))
				if v16 == int32(0) {
					if base.Ui32(int32(16384)) <= base.Ui32(v4) {
						F__serverAssert(m, int32(_a446), int32(_a447), int32(227))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v31 = v28 + v4*int32(24)
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[165])))
						*(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[165]))) = v32 + l1
						return
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+57)))
					if v20&int32(1) == int32(0) {
						return
					} else {
						if base.Ui32(int32(16384)) <= base.Ui32(v4) {
							F__serverAssert(m, int32(_a446), int32(_a447), int32(227))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, _consts[111]))
							v31 = v28 + v4*int32(24)
							v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[165])))
							*(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_consts[165]))) = v32 + l1
							return
						}
					}
				}
			}
		}
	}
}
func F_clusterSlotStatsAddNetworkBytesInForUserClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v4 == int32(-1) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[229]))
		if v8 == int32(0) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[110]))
			if v12 == int32(0) {
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				v19 = *(*int32)(unsafe.Add(mBase, _consts[62]))
				if v15&int32(16)|v19 != 0 {
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
					if v23 != int32(17) {
						v29 = v21
					} else {
						v27 = v21 + int64(15)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v27
						v29 = v27
					}
					v31 = *(*int32)(unsafe.Add(mBase, _consts[111]))
					v34 = v31 + v4*int32(24)
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[164])))
					*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[164]))) = v37 + v29
				}
			}
		}
	}
	return
}
func F_clusterSlotStatsCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
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
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v640 int32
	_ = v640
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v697 int64
	_ = v697
	var v699 int32
	_ = v699
	var v705 int64
	_ = v705
	var v707 int32
	_ = v707
	var v713 int64
	_ = v713
	var v714 int64
	_ = v714
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	v13 = m.G0
	v15 = v13 - int32(262160)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(262160)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v22 != int32(5) {
		v162 = v22
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a209))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	if v162 < int32(4) {
		goto L47
	} else {
		goto L48
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v27 = F_objectGetVal(m, v26)
	mBase = m.M
	v28 = int32(_a376)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v162 = v161
	goto L6
L9:
	;
	if v63-v65 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v63 = F_tolower(m, v59)
	mBase = m.M
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v65 = F_tolower(m, v64)
	mBase = m.M
	goto L9
L11:
	;
	v33 = v27
	v34 = v28
	v35 = v31
	goto L14
L12:
	;
	v59 = int32(0)
	v60 = v28
	goto L10
L13:
	;
	v59 = v56 & int32(255)
	v60 = v55
	goto L10
L14:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v37 == int32(0) {
		v55 = v34
		v56 = v35
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v55 = v49
	v56 = int32(0)
	goto L13
L16:
	;
	v41 = v35 & int32(255)
	if v41 == v37 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v48 = int32(1)
	v49 = v34 + v48
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v50 != 0 {
		v33 = v33 + v48
		v34 = v49
		v35 = v50
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v43 = F_tolower(m, v41)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v45 = F_tolower(m, v44)
	mBase = m.M
	if v43 == v45 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v55 = v34
	v56 = v47
	goto L13
L20:
	;
	goto L15
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = F_getSlotOrReply(m, l0, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	if v69 == int32(-1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = F_getSlotOrReply(m, l0, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v75 == int32(-1) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v69 <= v75 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = int32(0)
	v91 = F__emscripten_memset_bulkmem(m, v15+int32(16), base.I32_extend8_s(v85), int32(16384))
	mBase = m.M
	goto L29
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v69
	F_addReplyErrorFormat(m, l0, int32(_a450), v15)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L1
L29:
	;
	v96 = v85
	v97 = v69
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	goto L33
L31:
	;
	F_addReplyArrayLen(m, l0, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L38
	}
L32:
	;
	if v97 != v75 {
		v96 = v131
		v97 = v97 + int32(1)
		goto L30
	} else {
		goto L37
	}
L33:
	;
	v107 = F_clusterNodeGetPrimary(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v110 = base.I32_div_s(v97, int32(8))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v110+int32(104)))))
	goto L35
L35:
	;
	if int32(base.Ui32(v114)>>(uint(v97&int32(7))%32))&int32(1) == int32(0) {
		v131 = v96
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v124 = v15 + int32(16) + v97
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v126 = int32(1)
	v127 = v125 + v126
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v127)
	v131 = v96 + v126
	goto L32
L37:
	;
	goto L31
L38:
	;
	v140 = v69
	goto L39
L39:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(16)+v140))))
	if v153 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v140 != v75 {
		v140 = v140 + int32(1)
		goto L39
	} else {
		goto L44
	}
L42:
	;
	F_addReplySlotStat(m, l0, v140)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L1
L45:
	;
	v781 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L4
	} else {
		goto L217
	}
L46:
	;
	F_addReplyError(m, l0, int32(_a451))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L4
	} else {
		goto L216
	}
L47:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L215
	}
L48:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v167 = F_objectGetVal(m, v166)
	mBase = m.M
	v168 = int32(_a452)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v203-v205 != 0 {
		goto L47
	} else {
		goto L61
	}
L50:
	;
	v203 = F_tolower(m, v199)
	mBase = m.M
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v205 = F_tolower(m, v204)
	mBase = m.M
	goto L49
L51:
	;
	v173 = v167
	v174 = v168
	v175 = v171
	goto L54
L52:
	;
	v199 = int32(0)
	v200 = v168
	goto L50
L53:
	;
	v199 = v196 & int32(255)
	v200 = v195
	goto L50
L54:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v177 == int32(0) {
		v195 = v174
		v196 = v175
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v195 = v189
	v196 = int32(0)
	goto L53
L56:
	;
	v181 = v175 & int32(255)
	if v181 == v177 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v188 = int32(1)
	v189 = v174 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	if v190 != 0 {
		v173 = v173 + v188
		v174 = v189
		v175 = v190
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v183 = F_tolower(m, v181)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v185 = F_tolower(m, v184)
	mBase = m.M
	if v183 == v185 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v195 = v174
	v196 = v187
	goto L53
L60:
	;
	goto L55
L61:
	;
	v207 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v211 = F_objectGetVal(m, v210)
	mBase = m.M
	v212 = int32(_a453)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != 0 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v395 = int32(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v395
	v397 = int32(4)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v400 <= v397 {
		v596 = v395
		v599 = int32(109)
		goto L120
	} else {
		goto L121
	}
L63:
	;
	if v247-v249 == int32(0) {
		v394 = v207
		goto L62
	} else {
		goto L75
	}
L64:
	;
	v247 = F_tolower(m, v243)
	mBase = m.M
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v249 = F_tolower(m, v248)
	mBase = m.M
	goto L63
L65:
	;
	v217 = v211
	v218 = v212
	v219 = v215
	goto L68
L66:
	;
	v243 = int32(0)
	v244 = v212
	goto L64
L67:
	;
	v243 = v240 & int32(255)
	v244 = v239
	goto L64
L68:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v221 == int32(0) {
		v239 = v218
		v240 = v219
		goto L67
	} else {
		goto L70
	}
L69:
	;
	v239 = v233
	v240 = int32(0)
	goto L67
L70:
	;
	v225 = v219 & int32(255)
	if v225 == v221 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v232 = int32(1)
	v233 = v218 + v232
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v234 != 0 {
		v217 = v217 + v232
		v218 = v233
		v219 = v234
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v227 = F_tolower(m, v225)
	mBase = m.M
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	v229 = F_tolower(m, v228)
	mBase = m.M
	if v227 == v229 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v239 = v218
	v240 = v231
	goto L67
L74:
	;
	goto L69
L75:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v255 = F_objectGetVal(m, v254)
	mBase = m.M
	v256 = int32(_a454)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v259 != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v302 = F_objectGetVal(m, v301)
	mBase = m.M
	v303 = int32(_a455)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v306 != 0 {
		goto L94
	} else {
		goto L95
	}
L77:
	;
	if v291-v293 != 0 {
		goto L76
	} else {
		goto L89
	}
L78:
	;
	v291 = F_tolower(m, v287)
	mBase = m.M
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v293 = F_tolower(m, v292)
	mBase = m.M
	goto L77
L79:
	;
	v261 = v255
	v262 = v256
	v263 = v259
	goto L82
L80:
	;
	v287 = int32(0)
	v288 = v256
	goto L78
L81:
	;
	v287 = v284 & int32(255)
	v288 = v283
	goto L78
L82:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v265 == int32(0) {
		v283 = v262
		v284 = v263
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v283 = v277
	v284 = int32(0)
	goto L81
L84:
	;
	v269 = v263 & int32(255)
	if v269 == v265 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v276 = int32(1)
	v277 = v262 + v276
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v278 != 0 {
		v261 = v261 + v276
		v262 = v277
		v263 = v278
		goto L82
	} else {
		goto L88
	}
L86:
	;
	v271 = F_tolower(m, v269)
	mBase = m.M
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v273 = F_tolower(m, v272)
	mBase = m.M
	if v271 == v273 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v283 = v262
	v284 = v275
	goto L81
L88:
	;
	goto L83
L89:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if v296 == int32(0) {
		goto L76
	} else {
		goto L90
	}
L90:
	;
	v394 = int32(1)
	goto L62
L91:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v349 = F_objectGetVal(m, v348)
	mBase = m.M
	v350 = int32(_a456)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v353 != 0 {
		goto L108
	} else {
		goto L109
	}
L92:
	;
	if v338-v340 != 0 {
		goto L91
	} else {
		goto L104
	}
L93:
	;
	v338 = F_tolower(m, v334)
	mBase = m.M
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v340 = F_tolower(m, v339)
	mBase = m.M
	goto L92
L94:
	;
	v308 = v302
	v309 = v303
	v310 = v306
	goto L97
L95:
	;
	v334 = int32(0)
	v335 = v303
	goto L93
L96:
	;
	v334 = v331 & int32(255)
	v335 = v330
	goto L93
L97:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v312 == int32(0) {
		v330 = v309
		v331 = v310
		goto L96
	} else {
		goto L99
	}
L98:
	;
	v330 = v324
	v331 = int32(0)
	goto L96
L99:
	;
	v316 = v310 & int32(255)
	if v316 == v312 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v323 = int32(1)
	v324 = v309 + v323
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v325 != 0 {
		v308 = v308 + v323
		v309 = v324
		v310 = v325
		goto L97
	} else {
		goto L103
	}
L101:
	;
	v318 = F_tolower(m, v316)
	mBase = m.M
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v320 = F_tolower(m, v319)
	mBase = m.M
	if v318 == v320 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v330 = v309
	v331 = v322
	goto L96
L103:
	;
	goto L98
L104:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if v343 == int32(0) {
		goto L91
	} else {
		goto L105
	}
L105:
	;
	v394 = int32(2)
	goto L62
L106:
	;
	if v385-v387 != 0 {
		goto L46
	} else {
		goto L118
	}
L107:
	;
	v385 = F_tolower(m, v381)
	mBase = m.M
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	v387 = F_tolower(m, v386)
	mBase = m.M
	goto L106
L108:
	;
	v355 = v349
	v356 = v350
	v357 = v353
	goto L111
L109:
	;
	v381 = int32(0)
	v382 = v350
	goto L107
L110:
	;
	v381 = v378 & int32(255)
	v382 = v377
	goto L107
L111:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	if v359 == int32(0) {
		v377 = v356
		v378 = v357
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v377 = v371
	v378 = int32(0)
	goto L110
L113:
	;
	v363 = v357 & int32(255)
	if v363 == v359 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v370 = int32(1)
	v371 = v356 + v370
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+1)))
	if v372 != 0 {
		v355 = v355 + v370
		v356 = v371
		v357 = v372
		goto L111
	} else {
		goto L117
	}
L115:
	;
	v365 = F_tolower(m, v363)
	mBase = m.M
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356))))
	v367 = F_tolower(m, v366)
	mBase = m.M
	if v365 == v367 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	v377 = v356
	v378 = v369
	goto L110
L117:
	;
	goto L112
L118:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	if v390 == int32(0) {
		goto L46
	} else {
		goto L119
	}
L119:
	;
	v394 = int32(3)
	goto L62
L120:
	;
	v604 = int32(0)
	v608 = v604
	v611 = v604
	goto L179
L121:
	;
	v407 = v397
	v408 = v207
	v409 = int32(0)
	v413 = v400
	v414 = int32(1)
	goto L123
L122:
	;
	if v576 != 0 {
		goto L176
	} else {
		goto L177
	}
L123:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v419 = v407 << (uint(int32(2)) % 32)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417+v419)))
	v422 = F_objectGetVal(m, v421)
	mBase = m.M
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v424 = int32(_a457)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422))))
	if v427 != 0 {
		goto L129
	} else {
		goto L130
	}
L124:
	;
	F_addReplyError(m, l0, int32(_a458))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L175
	}
L125:
	;
	if int32(1) < v574 {
		goto L171
	} else {
		goto L172
	}
L126:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v423+v419)))
	v482 = F_objectGetVal(m, v481)
	mBase = m.M
	v483 = int32(_a459)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v486 != 0 {
		goto L147
	} else {
		goto L148
	}
L127:
	;
	if v459-v461 != 0 {
		goto L126
	} else {
		goto L139
	}
L128:
	;
	v459 = F_tolower(m, v455)
	mBase = m.M
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	v461 = F_tolower(m, v460)
	mBase = m.M
	goto L127
L129:
	;
	v429 = v422
	v430 = v424
	v431 = v427
	goto L132
L130:
	;
	v455 = int32(0)
	v456 = v424
	goto L128
L131:
	;
	v455 = v452 & int32(255)
	v456 = v451
	goto L128
L132:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v433 == int32(0) {
		v451 = v430
		v452 = v431
		goto L131
	} else {
		goto L134
	}
L133:
	;
	v451 = v445
	v452 = int32(0)
	goto L131
L134:
	;
	v437 = v431 & int32(255)
	if v437 == v433 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v444 = int32(1)
	v445 = v430 + v444
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	if v446 != 0 {
		v429 = v429 + v444
		v430 = v445
		v431 = v446
		goto L132
	} else {
		goto L138
	}
L136:
	;
	v439 = F_tolower(m, v437)
	mBase = m.M
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v441 = F_tolower(m, v440)
	mBase = m.M
	if v439 == v441 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	v451 = v430
	v452 = v443
	goto L131
L138:
	;
	goto L133
L139:
	;
	v464 = v407 + int32(1)
	if v413 <= v464 {
		goto L126
	} else {
		goto L140
	}
L140:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v423+v464<<(uint(int32(2))%32))))
	v475 = F_getRangeLongFromObjectOrReply(m, l0, v469, int32(1), int32(16384), v15+int32(12), int32(_a460))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	if v475 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v573 = v408
	v574 = v409 + int32(1)
	v575 = v464
	v576 = v414
	goto L125
L143:
	;
	v573 = v571
	v574 = v409
	v575 = v407
	v576 = v572
	goto L125
L144:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v525+v419)))
	v528 = F_objectGetVal(m, v527)
	mBase = m.M
	v529 = int32(_a461)
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v532 != 0 {
		goto L160
	} else {
		goto L161
	}
L145:
	;
	if v518-v520 != 0 {
		goto L144
	} else {
		goto L157
	}
L146:
	;
	v518 = F_tolower(m, v514)
	mBase = m.M
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	v520 = F_tolower(m, v519)
	mBase = m.M
	goto L145
L147:
	;
	v488 = v482
	v489 = v483
	v490 = v486
	goto L150
L148:
	;
	v514 = int32(0)
	v515 = v483
	goto L146
L149:
	;
	v514 = v511 & int32(255)
	v515 = v510
	goto L146
L150:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	if v492 == int32(0) {
		v510 = v489
		v511 = v490
		goto L149
	} else {
		goto L152
	}
L151:
	;
	v510 = v504
	v511 = int32(0)
	goto L149
L152:
	;
	v496 = v490 & int32(255)
	if v496 == v492 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v503 = int32(1)
	v504 = v489 + v503
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
	if v505 != 0 {
		v488 = v488 + v503
		v489 = v504
		v490 = v505
		goto L150
	} else {
		goto L156
	}
L154:
	;
	v498 = F_tolower(m, v496)
	mBase = m.M
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	v500 = F_tolower(m, v499)
	mBase = m.M
	if v498 == v500 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	v510 = v489
	v511 = v502
	goto L149
L156:
	;
	goto L151
L157:
	;
	v571 = v408 + int32(1)
	v572 = int32(0)
	goto L143
L158:
	;
	if v564-v566 != 0 {
		goto L45
	} else {
		goto L170
	}
L159:
	;
	v564 = F_tolower(m, v560)
	mBase = m.M
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	v566 = F_tolower(m, v565)
	mBase = m.M
	goto L158
L160:
	;
	v534 = v528
	v535 = v529
	v536 = v532
	goto L163
L161:
	;
	v560 = int32(0)
	v561 = v529
	goto L159
L162:
	;
	v560 = v557 & int32(255)
	v561 = v556
	goto L159
L163:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v538 == int32(0) {
		v556 = v535
		v557 = v536
		goto L162
	} else {
		goto L165
	}
L164:
	;
	v556 = v550
	v557 = int32(0)
	goto L162
L165:
	;
	v542 = v536 & int32(255)
	if v542 == v538 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v549 = int32(1)
	v550 = v535 + v549
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	if v551 != 0 {
		v534 = v534 + v549
		v535 = v550
		v536 = v551
		goto L163
	} else {
		goto L169
	}
L167:
	;
	v544 = F_tolower(m, v542)
	mBase = m.M
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	v546 = F_tolower(m, v545)
	mBase = m.M
	if v544 == v546 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	v556 = v535
	v557 = v548
	goto L162
L169:
	;
	goto L164
L170:
	;
	v568 = int32(1)
	v571 = v408 + v568
	v572 = v568
	goto L143
L171:
	;
	goto L124
L172:
	;
	if int32(2) <= v573 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v582 = v575 + int32(1)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v583 <= v582 {
		goto L122
	} else {
		goto L174
	}
L174:
	;
	v407 = v582
	v408 = v573
	v409 = v574
	v413 = v583
	v414 = v576
	goto L123
L175:
	;
	goto L1
L176:
	;
	v590 = int32(109)
	goto L178
L177:
	;
	v590 = int32(110)
	goto L178
L178:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v596 = v591
	v599 = v590
	goto L120
L179:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	goto L182
L180:
	;
	v725 = int32(16)
	F_qsort(m, v15+v725, v719, v725, v599)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L201
	}
L181:
	;
	v722 = v608 + int32(1)
	if v722 != int32(16384) {
		v608 = v722
		v611 = v719
		goto L179
	} else {
		goto L200
	}
L182:
	;
	v621 = F_clusterNodeGetPrimary(m, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v624 = base.I32_div_s(v608, int32(8))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v624+int32(104)))))
	goto L184
L184:
	;
	if int32(base.Ui32(v628)>>(uint(v608&int32(7))%32))&int32(1) == int32(0) {
		v719 = v611
		goto L181
	} else {
		goto L185
	}
L185:
	;
	v640 = v15 + int32(16) + v611<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v608
	switch v394 {
	default:
		goto L190
	case 1:
		goto L189
	case 2:
		goto L188
	case 3:
		goto L187
	}
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v640)+8)) = v714
	v719 = v611 + int32(1)
	goto L181
L187:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v713 = *(*int64)(unsafe.Add(mBase, uint32(v707+v608*int32(24))+uint32(_consts[163])))
	v714 = v713
	goto L186
L188:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v699+v608*int32(24))+uint32(_consts[164])))
	v714 = v705
	goto L186
L189:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v693+v608*int32(24))+uint32(_consts[165])))
	v714 = v697
	goto L186
L190:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(1) <= v649 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v714 = base.I64_extend_i32_u(v686)
	goto L186
L192:
	;
	goto L191
L193:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v655 = int32(0)
	v658 = v649
	v659 = v655
	v660 = v654
	v661 = v655
	goto L195
L194:
	;
	v686 = int32(0)
	goto L192
L195:
	;
	v664 = int32(0)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v660+v661<<(uint(int32(2))%32))))
	if v668 == v664 {
		v677 = v658
		v678 = v660
		v679 = v664
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v686 = v680
	goto L192
L197:
	;
	v680 = v679 + v659
	v682 = v661 + int32(1)
	if v682 < v677 {
		v658 = v677
		v659 = v680
		v660 = v678
		v661 = v682
		goto L195
	} else {
		goto L199
	}
L198:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	v672 = F_kvstoreHashtableSize(m, v671, v608)
	mBase = m.M
	v673 = int32(_a20)
	v674 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v676 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v677 = v674
	v678 = v676
	v679 = v672
	goto L197
L199:
	;
	goto L196
L200:
	;
	goto L180
L201:
	;
	v732 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v733)+88)))
	if v734&int32(2) == int32(0) {
		v741 = v733
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v596 < v744 {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+2160))
	v744 = v742
	goto L202
L204:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v733)+2172))
	if v739 != 0 {
		v741 = v739
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v744 = int32(0)
	goto L202
L206:
	;
	v746 = v596
	goto L208
L207:
	;
	v746 = v744
	goto L208
L208:
	;
	F_addReplyArrayLen(m, l0, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L209
	}
L209:
	;
	if v746 < int32(1) {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v754 = int32(0)
	goto L211
L211:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(16)+v754<<(uint(int32(4))%32))))
	F_addReplySlotStat(m, l0, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	v773 = v754 + int32(1)
	if v773 != v746 {
		v754 = v773
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L1
L215:
	;
	goto L1
L216:
	;
	goto L1
L217:
	;
	goto L1
}
func F_clusterSlotStatsEnabled(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(_a20)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	v6 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	return base.B2i32(l0 != int32(-1)) & (base.B2i32(v5 != v6) & base.B2i32(v9 != v6))
}
func F_clusterSlotStatsUpdateNetworkBytesOutForReplication(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v5 == int32(0) {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+292))
		if v8 == int32(-1) {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[229]))
			if v12 == int32(0) {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[110]))
				if v16 == int32(0) {
					return
				} else {
					if base.Ui32(int32(16384)) <= base.Ui32(v8) {
						F__serverAssert(m, int32(_a446), int32(_a447), int32(156))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, _consts[111]))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+88)))
						if v24&int32(1) == int32(0) {
							F__serverAssert(m, int32(_a448), int32(_a447), int32(157))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v30 = v8 * int32(24)
							v32 = *(*int32)(unsafe.Add(mBase, _consts[78]))
							v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+20)))
							v34 = l0 * v33
							if int64(-1) < v34 {
								v44 = v22 + v30
								v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[163])))
								*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[163]))) = v47 + v34
								return
							} else {
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v22+v30)+uint32(_consts[163])))
								if base.Ui64(v40) < base.Ui64(int64(0)-v34) {
									F__serverAssert(m, int32(_a449), int32(_a447), int32(160))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v44 = v22 + v30
									v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[163])))
									*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[163]))) = v47 + v34
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
func F_clusterStartHandshake(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
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
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int64
	_ = v377
	var v383 int64
	_ = v383
	var v389 int64
	_ = v389
	var v395 int64
	_ = v395
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v444 int32
	_ = v444
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v12 = v8 + int32(4)
	v13 = F_inet_pton(m, int32(2), l0, v12)
	mBase = m.M
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(176)
	return v444
L2:
	;
	goto L101
L3:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	v81 = F_dictGetSafeIterator(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L13
	}
L4:
	;
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(144)))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(152)))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(160)))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(166)))) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v8)+128)) = v51
	v67 = int32(10)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v67)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+136)) = v51
	v75 = F_inet_ntop(m, v67, v16, v8+int32(128), int32(46))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L10
	}
L5:
	;
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(144)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(152)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(160)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(166)))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8)+128)) = v21
	v37 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v37)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+136)) = v21
	v45 = F_inet_ntop(m, v37, v12, v8+int32(128), int32(46))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v16 = v8 + int32(8)
	v17 = F_inet_pton(m, int32(10), l0, v16)
	mBase = m.M
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v432 = int32(28)
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	goto L3
L10:
	;
	goto L3
L11:
	;
	F_dictReleaseIterator(m, v81)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L8
	} else {
		goto L100
	}
L12:
	;
	F_dictReleaseIterator(m, v81)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L91
	}
L13:
	;
	v90 = v81 + int32(20)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v91 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v186 == int32(0) {
		goto L12
	} else {
		goto L40
	}
L15:
	;
	v97 = v90
	v98 = v94
	goto L18
L16:
	;
	v94 = int32(1)
	goto L15
L17:
	;
	v94 = int32(0)
	goto L15
L18:
	;
	switch v98 {
	case 0:
		goto L23
	default:
		goto L22
	}
L20:
	;
	v98 = int32(0)
	goto L18
L21:
	;
	goto L14
L22:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v178
	if v178 == int32(0) {
		goto L20
	} else {
		goto L39
	}
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v102 != int32(-1) {
		v141 = v102
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v142 = int32(1)
	v143 = v141 + v142
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v143
	v145 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v149+int32(26)))))
	if v153 == int32(255) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v106 != 0 {
		v141 = int32(-1)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v108 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	if v135 != int32(-1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v107)+16)))
	v116 = int64(*(*int8)(unsafe.Add(mBase, uint32(v107)+27)))
	v117 = int64(*(*int32)(unsafe.Add(mBase, uint32(v107)+8)))
	v118 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v107)+12)))
	v119 = int64(*(*int8)(unsafe.Add(mBase, uint32(v107)+26)))
	v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v107)+4)))
	v121 = F_wangHash64(m, v120)
	mBase = m.M
	v123 = F_wangHash64(m, v119+v121)
	mBase = m.M
	v125 = F_wangHash64(m, v118+v123)
	mBase = m.M
	v127 = F_wangHash64(m, v117+v125)
	mBase = m.M
	v129 = F_wangHash64(m, v116+v127)
	mBase = m.M
	v131 = F_wangHash64(m, v115+v129)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v134 = v133
	goto L27
L29:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+24)))
	v113 = v111 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+24)) = uint16(v113)
	v134 = v107
	goto L27
L30:
	;
	v141 = v135 + int32(-1)
	goto L24
L31:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v141 = v138
	goto L24
L32:
	;
	v168 = int32(2)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v148+v166<<(uint(v168)%32)+int32(4))))
	v97 = v173 + v167<<(uint(v168)%32)
	v98 = int32(1)
	goto L18
L33:
	;
	v157 = v145
	goto L35
L34:
	;
	v157 = v142 << (uint(v153) % 32)
	goto L35
L35:
	;
	if v143 < v157 {
		v166 = v149
		v167 = v143
		goto L32
	} else {
		goto L36
	}
L36:
	;
	if v149 != 0 {
		v186 = v145
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	if v159 == int32(-1) {
		v186 = v145
		goto L21
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+4)) = int64(4294967296)
	v166 = int32(1)
	v167 = int32(0)
	goto L32
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v182
	v186 = v178
	goto L21
L40:
	;
	v192 = v186
	goto L41
L41:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	goto L44
L42:
	;
	goto L12
L43:
	;
	v262 = v81 + int32(20)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v263 != 0 {
		goto L66
	} else {
		goto L67
	}
L44:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+88)))
	if v198&int32(32) == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v204 = v197 + int32(2256)
	v206 = v8 + int32(128)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v209 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v241-v243 != 0 {
		goto L43
	} else {
		goto L58
	}
L47:
	;
	v241 = F_tolower(m, v237)
	mBase = m.M
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v243 = F_tolower(m, v242)
	mBase = m.M
	goto L46
L48:
	;
	v211 = v204
	v212 = v206
	v213 = v209
	goto L51
L49:
	;
	v237 = int32(0)
	v238 = v206
	goto L47
L50:
	;
	v237 = v234 & int32(255)
	v238 = v233
	goto L47
L51:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v215 == int32(0) {
		v233 = v212
		v234 = v213
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v233 = v227
	v234 = int32(0)
	goto L50
L53:
	;
	v219 = v213 & int32(255)
	if v219 == v215 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v226 = int32(1)
	v227 = v212 + v226
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	if v228 != 0 {
		v211 = v211 + v226
		v212 = v227
		v213 = v228
		goto L51
	} else {
		goto L57
	}
L55:
	;
	v221 = F_tolower(m, v219)
	mBase = m.M
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v223 = F_tolower(m, v222)
	mBase = m.M
	if v221 == v223 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v233 = v212
	v234 = v225
	goto L50
L57:
	;
	goto L52
L58:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v248 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v249 = int32(2328)
	goto L61
L60:
	;
	v249 = int32(2324)
	goto L61
L61:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v197+v249)))
	if v251 != l1 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v197)+2332))
	if v253 == l2 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	goto L43
L64:
	;
	if v358 != 0 {
		v192 = v358
		goto L41
	} else {
		goto L90
	}
L65:
	;
	v269 = v262
	v270 = v266
	goto L68
L66:
	;
	v266 = int32(1)
	goto L65
L67:
	;
	v266 = int32(0)
	goto L65
L68:
	;
	switch v270 {
	case 0:
		goto L73
	default:
		goto L72
	}
L70:
	;
	v270 = int32(0)
	goto L68
L71:
	;
	goto L64
L72:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v350
	if v350 == int32(0) {
		goto L70
	} else {
		goto L89
	}
L73:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v274 != int32(-1) {
		v313 = v274
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v314 = int32(1)
	v315 = v313 + v314
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v315
	v317 = int32(0)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v321+int32(26)))))
	if v325 == int32(255) {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v278 != 0 {
		v313 = int32(-1)
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v280 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+20))
	if v307 != int32(-1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v287 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v279)+16)))
	v288 = int64(*(*int8)(unsafe.Add(mBase, uint32(v279)+27)))
	v289 = int64(*(*int32)(unsafe.Add(mBase, uint32(v279)+8)))
	v290 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v279)+12)))
	v291 = int64(*(*int8)(unsafe.Add(mBase, uint32(v279)+26)))
	v292 = int64(*(*int32)(unsafe.Add(mBase, uint32(v279)+4)))
	v293 = F_wangHash64(m, v292)
	mBase = m.M
	v295 = F_wangHash64(m, v291+v293)
	mBase = m.M
	v297 = F_wangHash64(m, v290+v295)
	mBase = m.M
	v299 = F_wangHash64(m, v289+v297)
	mBase = m.M
	v301 = F_wangHash64(m, v288+v299)
	mBase = m.M
	v303 = F_wangHash64(m, v287+v301)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v306 = v305
	goto L77
L79:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+24)))
	v285 = v283 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279)+24)) = uint16(v285)
	v306 = v279
	goto L77
L80:
	;
	v313 = v307 + int32(-1)
	goto L74
L81:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v313 = v310
	goto L74
L82:
	;
	v340 = int32(2)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v320+v338<<(uint(v340)%32)+int32(4))))
	v269 = v345 + v339<<(uint(v340)%32)
	v270 = int32(1)
	goto L68
L83:
	;
	v329 = v317
	goto L85
L84:
	;
	v329 = v314 << (uint(v325) % 32)
	goto L85
L85:
	;
	if v315 < v329 {
		v338 = v321
		v339 = v315
		goto L82
	} else {
		goto L86
	}
L86:
	;
	if v321 != 0 {
		v358 = v317
		goto L71
	} else {
		goto L87
	}
L87:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	if v331 == int32(-1) {
		v358 = v317
		goto L71
	} else {
		goto L88
	}
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+4)) = int64(4294967296)
	v338 = int32(1)
	v339 = int32(0)
	goto L82
L89:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v354
	v358 = v350
	goto L71
L90:
	;
	goto L42
L91:
	;
	v371 = F_createClusterNode(m, int32(0), int32(160))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(166))))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(2294)))) = v377
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(160))))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(2288)))) = v383
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(152))))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(2280)))) = v389
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v8+int32(144))))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(2272)))) = v395
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v8)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(2264)))) = v399
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v8)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v371)+2256)) = v401
	v404 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v404 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+2332)) = l2
	v412 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+32))
	v417 = F_sdsnewlen(m, v371+int32(8), int32(40))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+2324)) = l1
	goto L93
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+2328)) = l1
	goto L93
L96:
	;
	v419 = F_dictAdd(m, v413, v417, v371)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	if v419 == int32(0) {
		v444 = int32(1)
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F__serverAssert(m, int32(_a258), int32(_a247), int32(2220))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v432 = int32(6)
	goto L2
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v432
	v444 = int32(0)
	goto L1
}
func F_clusterUpdateMyselfClientIpV6(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v3 == v1 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[124]))
		F_updateSdsExtensionField(m, v3+int32(2308), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterUpdateMyselfHumanNodename(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v3 == v1 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[123]))
		F_updateSdsExtensionField(m, v3+int32(2316), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_clusterUpdateMyselfIp(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v4 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v10 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L39
	}
L4:
	;
	F_valkey_free(m, v10)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L23
	} else {
		goto L38
	}
L5:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[121])) = v8
	v57 = F_zstrdup(m, v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L23
	} else {
		goto L25
	}
L6:
	;
	if v10 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v8 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if v10 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	if v8 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v8 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v22 == int32(0) {
		v45 = v21
		v46 = v22
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v46-v45&int32(255) == int32(0) {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v22 != v21&int32(255) {
		v45 = v21
		v46 = v22
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v28 = v10
	v29 = v8
	goto L18
L18:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v33 == int32(0) {
		v45 = v32
		v46 = v33
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v45 = v32
	v46 = v33
	goto L15
L20:
	;
	v36 = int32(1)
	if v33 == v32&int32(255) {
		v28 = v28 + v36
		v29 = v29 + v36
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_valkey_free(m, v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	goto L5
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[121])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v65 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	goto L29
L26:
	;
	goto L3
L27:
	;
	goto L26
L28:
	;
	v96 = v75
	goto L35
L29:
	;
	v71 = v61 + int32(2256)
	v73 = int32(46)
	v75 = v65
	goto L31
L30:
	;
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v86)
	goto L28
L31:
	;
	v77 = v73 + int32(-1)
	if v77 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v80)
	v82 = int32(1)
	if v80 != 0 {
		v71 = v71 + v82
		v73 = v77
		v75 = v75 + v82
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v98 != 0 {
		v96 = v96 + int32(1)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L27
L37:
	;
	goto L36
L38:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[121])) = v111
	v115 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+2256)) = uint8(v111)
	goto L3
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_consts[122]))) = v122 | int32(4)
	goto L1
}
func F_clusterUpdateSlots(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	v12 = int32(0)
	goto L2
L1:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a318), int32(_a247), int32(7219))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L34
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v12))))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v141 = v12 + int32(1)
	if v141 != int32(16384) {
		v12 = v141
		goto L2
	} else {
		goto L33
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v23 = F_dictFind(m, v22, v12)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	return
L8:
	;
	if v23 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	goto L10
L10:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	v33 = F_dictFind(m, v32, v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v40 = F_dictDelete(m, v39, v12)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	v135 = F_clusterDelSlot(m, v12)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L31
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v45 = v12 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v45)+52))
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v53 = m.G0
	v55 = v53 - int32(32)
	m.G0 = v55
	v60 = int32(1) << (uint(v12&int32(7)) % 32)
	v62 = base.I32_div_s(v12, int32(8))
	v65 = v49 + v62 + int32(104)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v60&v66 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v104 = int32(_a20)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v105+v45)+52)) = v49
	v110 = v105 + int32(base.Ui32(v12)>>(uint(int32(3))%32))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[162]))))
	v116 = v111 & base.I32_rotl(int32(-2), v12&int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_consts[162]))) = uint8(v116)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v122 = v119 + v12*int32(24)
	v125 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[163]))) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[164]))) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[165]))) = v125
	goto L30
L19:
	;
	m.G0 = v55 + int32(32)
	goto L18
L20:
	;
	v68 = v66 | v60
	*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v49)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+2160)) = v70 + int32(1)
	if v70 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
	F_dictInitIterator(m, v55, v76)
	mBase = m.M
	v78 = F_dictNext(m, v55)
	mBase = m.M
	if v78 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v82 = v78
	goto L24
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v49)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+88)) = v92 | int32(256)
	goto L19
L24:
	;
	v86 = F_dictGetVal(m, v82)
	mBase = m.M
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+88)))
	if v87&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v91 = F_dictNext(m, v55)
	mBase = m.M
	if v91 != 0 {
		v82 = v91
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+2164))
	if v90 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L19
L30:
	;
	goto L4
L31:
	;
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L4
L33:
	;
	goto L3
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_clusterUpdateSlotsConfigWith(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
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
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int64
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
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
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int64
	_ = v677
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int64
	_ = v1176
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1250 int32
	_ = v1250
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1364 int32
	_ = v1364
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	v32 = m.G0
	v34 = v32 - int32(32960)
	m.G0 = v34
	v37 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if l0 != v37 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v34 + int32(32960)
	return
L2:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+88)))
	if v51&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v40 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F__serverLog(m, int32(2), int32(_a297), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v54 = int32(_a298)
	goto L9
L8:
	;
	v54 = v37 + int32(2172)
	goto L9
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v57 = l0 + int32(8)
	v58 = int32(48)
	v59 = l0 + v58
	v61 = v37 + v58
	v62 = int32(40)
	goto L14
L10:
	;
	v135 = int32(0)
	v138 = int32(-1)
	v149 = v135
	v158 = v135
	v159 = v135
	v160 = v138
	v161 = v138
	v162 = v135
	v163 = v135
	v164 = v135
	v165 = v135
	v166 = v135
	goto L26
L11:
	;
	v126 = int32(0)
	goto L10
L12:
	;
	v98 = v93
	v99 = v94
	v100 = v95
	goto L22
L13:
	;
	if v83 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	if (v61|v59)&int32(3) != 0 {
		v93 = v59
		v94 = v61
		v95 = v62
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v70 = v59
	v71 = v61
	v72 = v62
	goto L16
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v75 != v76 {
		v93 = v70
		v94 = v71
		v95 = v72
		goto L12
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v78 = int32(4)
	v79 = v71 + v78
	v81 = v70 + v78
	v83 = v72 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v83) {
		v70 = v81
		v71 = v79
		v72 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v93 = v81
	v94 = v79
	v95 = v83
	goto L12
L21:
	;
	v126 = v103 - v104
	goto L10
L22:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v103 != v104 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v106 = int32(1)
	v111 = v100 + int32(-1)
	if v111 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v98 = v98 + v106
	v99 = v99 + v106
	v100 = v111
	goto L22
L26:
	;
	v179 = int32(1) << (uint(v149&int32(7)) % 32)
	v181 = int32(base.Ui32(v149) >> (uint(int32(3)) % 32))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v181))))
	if v179&v183 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	if v1210 == int32(-1) {
		goto L291
	} else {
		goto L292
	}
L28:
	;
	v1226 = v149 + int32(1)
	if v1226 != int32(16384) {
		v149 = v1226
		v158 = v1207
		v159 = v1208
		v160 = v1209
		v161 = v1210
		v162 = v1211
		v163 = v1212
		v164 = v1213
		v165 = v1214
		v166 = v1215
		goto L26
	} else {
		goto L290
	}
L29:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v706 != v711 {
		v775 = v162
		goto L165
	} else {
		goto L166
	}
L30:
	;
	v706 = v704
	v707 = v161
	v708 = v160
	v709 = v159
	goto L29
L31:
	;
	v704 = int32(0)
	goto L30
L32:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v404 = v149 << (uint(int32(2)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402+v404)+52))
	if v406 != l0 {
		v414 = v402
		goto L88
	} else {
		goto L89
	}
L33:
	;
	v188 = v164 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v192 = v149 << (uint(int32(2)) % 32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v192)+52))
	if v194 != l0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v194 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v196 = v190 + v181
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+uint32(_consts[162]))))
	v200 = v197 & (v179 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+uint32(_consts[162]))) = uint8(v200)
	v1207 = v158
	v1208 = v159
	v1209 = v160
	v1210 = v161
	v1211 = v162
	v1212 = v163
	v1213 = v188
	v1214 = v165
	v1215 = v166
	goto L28
L36:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v181)+uint32(_consts[162]))))
	if v179&v205 != 0 {
		v286 = v190
		v287 = v194
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v181)+uint32(_consts[162]))))
	if v179&v289 != 0 {
		v704 = v287
		goto L30
	} else {
		goto L61
	}
L38:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v194)+96))
	if base.Ui64(v207) < base.Ui64(l1) {
		v286 = v190
		v287 = v194
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v209 = int32(0)
	v213 = m.G0
	v215 = v213 - int32(16)
	m.G0 = v215
	v218 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+uint32(_consts[171])))
	F_listRewind(m, v219, v215)
	mBase = m.M
	v222 = F_listNext(m, v215)
	mBase = m.M
	if v222 == v209 {
		v272 = v209
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v272 == int32(0) {
		v1207 = v158
		v1208 = v159
		v1209 = v160
		v1210 = v161
		v1211 = v162
		v1212 = v163
		v1213 = v188
		v1214 = v165
		v1215 = v166
		goto L28
	} else {
		goto L59
	}
L41:
	;
	m.G0 = v215 + int32(16)
	goto L40
L42:
	;
	v228 = v222
	goto L44
L43:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v230)+156))
	v272 = base.B2i32(v267 == int32(17))
	goto L41
L44:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v231 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v266 = F_listNext(m, v215)
	mBase = m.M
	if v266 != 0 {
		v228 = v266
		goto L44
	} else {
		goto L58
	}
L47:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)+156))
	if base.Ui32(int32(20)) < base.Ui32(v232) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v230)+164))
	v241 = v215 + int32(8)
	F_listRewind(m, v239, v241)
	mBase = m.M
	v245 = F_listNext(m, v241)
	mBase = m.M
	if v245 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	if int32(1)<<(uint(v232)%32)&int32(1835040) != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v251 = v245
	goto L52
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v149 < v254 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L46
L54:
	;
	v260 = F_listNext(m, v215+int32(8))
	mBase = m.M
	if v260 != 0 {
		v251 = v260
		goto L52
	} else {
		goto L57
	}
L55:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v149 <= v256 {
		goto L43
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L53
L58:
	;
	v272 = v209
	goto L41
L59:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281+v192)+52))
	if v283 == int32(0) {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v286 = v281
	v287 = v283
	goto L37
L61:
	;
	v292 = v287 + int32(48)
	v293 = int32(40)
	goto L66
L62:
	;
	if v357 == int32(0) {
		v704 = v287
		goto L30
	} else {
		goto L78
	}
L63:
	;
	v357 = int32(0)
	goto L62
L64:
	;
	v329 = v324
	v330 = v325
	v331 = v326
	goto L74
L65:
	;
	if v314 == int32(0) {
		goto L63
	} else {
		goto L72
	}
L66:
	;
	if (v59|v292)&int32(3) != 0 {
		v324 = v292
		v325 = v59
		v326 = v293
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v301 = v292
	v302 = v59
	v303 = v293
	goto L68
L68:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v306 != v307 {
		v324 = v301
		v325 = v302
		v326 = v303
		goto L64
	} else {
		goto L70
	}
L69:
	;
	goto L65
L70:
	;
	v309 = int32(4)
	v310 = v302 + v309
	v312 = v301 + v309
	v314 = v303 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v314) {
		v301 = v312
		v302 = v310
		v303 = v314
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v324 = v312
	v325 = v310
	v326 = v314
	goto L64
L73:
	;
	v357 = v334 - v335
	goto L62
L74:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v334 != v335 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v337 = int32(1)
	v342 = v331 + int32(-1)
	if v342 == int32(0) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v329 = v329 + v337
	v330 = v330 + v337
	v331 = v342
	goto L74
L78:
	;
	if v161 == int32(-1) {
		v397 = v287
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v706 = v397
	v707 = v149
	v708 = v149
	v709 = v397
	goto L29
L80:
	;
	if v159 != v287 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v367 {
		v706 = v287
		v707 = v149
		v708 = v149
		v709 = v287
		goto L29
	} else {
		goto L84
	}
L82:
	;
	if v149 != v160+int32(1) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v706 = v159
	v707 = v161
	v708 = v149
	v709 = v159
	goto L29
L84:
	;
	v370 = F_humanNodename(m, v159)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v372 = F_humanNodename(m, l0)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(188)))) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(184)))) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(180)))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(176)))) = v159 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+172)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v34)+168)) = v159 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+164)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v161
	F__serverLog(m, int32(2), int32(_a299), v34+int32(160))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v393+v192)+52))
	v397 = v395
	goto L79
L88:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)+44))
	v416 = F_dictFind(m, v415, v149)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L91
	}
L89:
	;
	v408 = v402 + v181
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+uint32(_consts[162]))))
	v410 = v409 | v179
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+uint32(_consts[162]))) = uint8(v410)
	v413 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v414 = v413
	goto L88
L90:
	;
	v546 = int32(0)
	v547 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+88)))
	if v548&int32(1) == v546 {
		v1207 = v158
		v1208 = v159
		v1209 = v160
		v1210 = v161
		v1211 = v162
		v1212 = v163
		v1213 = v164
		v1214 = v165
		v1215 = v166
		goto L28
	} else {
		goto L131
	}
L91:
	;
	if v416 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416)+8))
	goto L93
L93:
	;
	if v420 == int32(0) {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	if v420 == l0 {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v420)+96))
	if base.Ui64(v424) < base.Ui64(l1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v432 = v420 + int32(48)
	v433 = int32(40)
	goto L103
L97:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+88)))
	if v426&int32(2) == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	if v497 != 0 {
		goto L90
	} else {
		goto L115
	}
L100:
	;
	v497 = int32(0)
	goto L99
L101:
	;
	v469 = v464
	v470 = v465
	v471 = v466
	goto L111
L102:
	;
	if v454 == int32(0) {
		goto L100
	} else {
		goto L109
	}
L103:
	;
	if (v59|v432)&int32(3) != 0 {
		v464 = v432
		v465 = v59
		v466 = v433
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v441 = v432
	v442 = v59
	v443 = v433
	goto L105
L105:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v446 != v447 {
		v464 = v441
		v465 = v442
		v466 = v443
		goto L101
	} else {
		goto L107
	}
L106:
	;
	goto L102
L107:
	;
	v449 = int32(4)
	v450 = v442 + v449
	v452 = v441 + v449
	v454 = v443 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v454) {
		v441 = v452
		v442 = v450
		v443 = v454
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v464 = v452
	v465 = v450
	v466 = v454
	goto L101
L110:
	;
	v497 = v474 - v475
	goto L99
L111:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	if v474 != v475 {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v477 = int32(1)
	v482 = v471 + int32(-1)
	if v482 == int32(0) {
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v469 = v469 + v477
	v470 = v470 + v477
	v471 = v482
	goto L111
L115:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(1) < v499 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+44))
	v518 = F_dictFind(m, v517, v149)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L5
	} else {
		goto L120
	}
L117:
	;
	v502 = F_humanNodename(m, l0)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+108)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+104)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v34)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v149
	F__serverLog(m, int32(1), int32(_a300), v34+int32(96))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	goto L116
L120:
	;
	if l0 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L5
	} else {
		goto L130
	}
L122:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+44))
	if v518 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	if v518 == int32(0) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+44))
	v525 = F_dictDelete(m, v524, v149)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	goto L121
L126:
	;
	v533 = F_dictAdd(m, v529, v149, l0)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+8)) = l0
	goto L128
L128:
	;
	goto L121
L129:
	;
	goto L121
L130:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v539)+uint32(_consts[122]))) = v540 | int32(14)
	goto L90
L131:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+48))
	v556 = F_dictFind(m, v555, v149)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L5
	} else {
		goto L134
	}
L132:
	;
	if v560 != l0 {
		v1207 = v158
		v1208 = v159
		v1209 = v160
		v1210 = v161
		v1211 = v162
		v1212 = v163
		v1213 = v164
		v1214 = v165
		v1215 = v166
		goto L28
	} else {
		goto L137
	}
L133:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	goto L136
L134:
	;
	if v556 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v560 = int32(0)
	goto L132
L136:
	;
	v560 = v559
	goto L132
L137:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v563 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+48))
	v582 = F_dictFind(m, v581, v149)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L5
	} else {
		goto L143
	}
L139:
	;
	v566 = F_humanNodename(m, l0)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+92)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+88)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v34)+84)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v149
	F__serverLog(m, int32(2), int32(_a301), v34+int32(80))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	v592 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v592+v404)+52))
	v596 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v594 == v596 {
		v1207 = v158
		v1208 = v159
		v1209 = v160
		v1210 = v161
		v1211 = v162
		v1212 = v163
		v1213 = v164
		v1214 = v165
		v1215 = v166
		goto L28
	} else {
		goto L146
	}
L143:
	;
	if v582 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v587 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+48))
	v589 = F_dictDelete(m, v588, v149)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L5
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	v598 = F_clusterDelSlot(m, v149)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L5
	} else {
		goto L147
	}
L147:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v601+v404)+52))
	if v603 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v690 = F_clusterBumpConfigEpochWithoutConsensus(m)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L5
	} else {
		goto L163
	}
L149:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v609 = m.G0
	v611 = v609 - int32(32)
	m.G0 = v611
	v616 = int32(1) << (uint(v149&int32(7)) % 32)
	v618 = base.I32_div_s(v149, int32(8))
	v621 = v605 + v618 + int32(104)
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v616&v622 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v660 = int32(_a20)
	v661 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v661+v404)+52)) = v605
	v664 = v661 + v181
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+uint32(_consts[162]))))
	v668 = v665 & (v179 ^ int32(-1))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+uint32(_consts[162]))) = uint8(v668)
	v671 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v674 = v671 + v149*int32(24)
	v677 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v674)+uint32(_consts[163]))) = v677
	*(*int64)(unsafe.Add(mBase, uint32(v674)+uint32(_consts[164]))) = v677
	*(*int64)(unsafe.Add(mBase, uint32(v674)+uint32(_consts[165]))) = v677
	goto L162
L151:
	;
	m.G0 = v611 + int32(32)
	goto L150
L152:
	;
	v624 = v622 | v616
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v624)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v605)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(v605)+2160)) = v626 + int32(1)
	if v626 != 0 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+32))
	F_dictInitIterator(m, v611, v632)
	mBase = m.M
	v634 = F_dictNext(m, v611)
	mBase = m.M
	if v634 == int32(0) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v638 = v634
	goto L156
L155:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v605)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v605)+88)) = v648 | int32(256)
	goto L151
L156:
	;
	v642 = F_dictGetVal(m, v638)
	mBase = m.M
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+88)))
	if v643&int32(2) != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v647 = F_dictNext(m, v611)
	mBase = m.M
	if v647 != 0 {
		v638 = v647
		goto L156
	} else {
		goto L161
	}
L159:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v642)+2164))
	if v646 != 0 {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	goto L151
L162:
	;
	goto L148
L163:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v695)+uint32(_consts[122]))) = v696 | int32(14)
	v1207 = v158
	v1208 = v159
	v1209 = v160
	v1210 = v161
	v1211 = v162
	v1212 = v163
	v1213 = v164
	v1214 = v165
	v1215 = v166
	goto L28
L165:
	;
	v776 = int32(0)
	v779 = m.G0
	v781 = v779 - int32(16)
	m.G0 = v781
	v784 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)+uint32(_consts[171])))
	F_listRewind(m, v785, v781)
	mBase = m.M
	v788 = F_listNext(m, v781)
	mBase = m.M
	if v788 == v776 {
		v831 = v776
		goto L179
	} else {
		goto L180
	}
L166:
	;
	v720 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(1) <= v720 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	if v757 == int32(0) {
		v775 = v162
		goto L165
	} else {
		goto L176
	}
L168:
	;
	goto L167
L169:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v726 = int32(0)
	v729 = v720
	v730 = v726
	v731 = v725
	v732 = v726
	goto L171
L170:
	;
	v757 = int32(0)
	goto L168
L171:
	;
	v735 = int32(0)
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v731+v732<<(uint(int32(2))%32))))
	if v739 == v735 {
		v748 = v729
		v749 = v731
		v750 = v735
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v757 = v751
	goto L168
L173:
	;
	v751 = v750 + v730
	v753 = v732 + int32(1)
	if v753 < v748 {
		v729 = v748
		v730 = v751
		v731 = v749
		v732 = v753
		goto L171
	} else {
		goto L175
	}
L174:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	v743 = F_kvstoreHashtableSize(m, v742, v149)
	mBase = m.M
	v744 = int32(_a20)
	v745 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v747 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v748 = v745
	v749 = v747
	v750 = v743
	goto L173
L175:
	;
	goto L172
L176:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if l0 == v765 {
		v775 = v162
		goto L165
	} else {
		goto L177
	}
L177:
	;
	v769 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(192)+v162<<(uint(v769)%32)))) = uint16(v149)
	v775 = v162 + v769
	goto L165
L178:
	;
	v836 = int32(0)
	v839 = m.G0
	v841 = v839 - int32(16)
	m.G0 = v841
	v844 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+uint32(_consts[171])))
	F_listRewind(m, v845, v841)
	mBase = m.M
	v848 = F_listNext(m, v841)
	mBase = m.M
	if v848 == v836 {
		v893 = v836
		goto L196
	} else {
		goto L197
	}
L179:
	;
	m.G0 = v781 + int32(16)
	goto L178
L180:
	;
	v794 = v788
	goto L182
L181:
	;
	v831 = int32(1)
	goto L179
L182:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+8))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	if v796 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v827 = F_listNext(m, v781)
	mBase = m.M
	if v827 != 0 {
		v794 = v827
		goto L182
	} else {
		goto L194
	}
L185:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v795)+156))
	if base.Ui32(v797+int32(-18)) < base.Ui32(int32(3)) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v795)+164))
	v804 = v781 + int32(8)
	F_listRewind(m, v802, v804)
	mBase = m.M
	v808 = F_listNext(m, v804)
	mBase = m.M
	if v808 == int32(0) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v814 = v808
	goto L188
L188:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+8))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	if v149 < v816 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L184
L190:
	;
	v822 = F_listNext(m, v781+int32(8))
	mBase = m.M
	if v822 != 0 {
		v814 = v822
		goto L188
	} else {
		goto L193
	}
L191:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v815)+4))
	if v149 <= v818 {
		goto L181
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	goto L189
L194:
	;
	v831 = v776
	goto L179
L195:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v899+v192)+52))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v899)+44))
	v903 = F_dictFind(m, v902, v149)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L213
	}
L196:
	;
	m.G0 = v841 + int32(16)
	goto L195
L197:
	;
	v854 = v848
	goto L199
L198:
	;
	v893 = int32(1)
	goto L196
L199:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+8))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)))
	if v856 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v889 = F_listNext(m, v841)
	mBase = m.M
	if v889 != 0 {
		v854 = v889
		goto L199
	} else {
		goto L211
	}
L202:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v855)+156))
	if base.Ui32(v859+int32(-18)) < base.Ui32(int32(3)) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v855)+164))
	v866 = v841 + int32(8)
	F_listRewind(m, v864, v866)
	mBase = m.M
	v870 = F_listNext(m, v866)
	mBase = m.M
	if v870 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v876 = v870
	goto L205
L205:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)+8))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	if v149 < v878 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	goto L201
L207:
	;
	v884 = F_listNext(m, v841+int32(8))
	mBase = m.M
	if v884 != 0 {
		v876 = v884
		goto L205
	} else {
		goto L210
	}
L208:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v877)+4))
	if v149 <= v880 {
		goto L198
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	goto L206
L211:
	;
	v893 = v836
	goto L196
L212:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+48))
	v950 = F_dictFind(m, v949, v149)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L226
	}
L213:
	;
	if v903 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v903)+8))
	goto L215
L215:
	;
	if v907 == int32(0) {
		goto L212
	} else {
		goto L216
	}
L216:
	;
	if v126 == int32(0) {
		goto L212
	} else {
		goto L217
	}
L217:
	;
	v913 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v913 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)+44))
	v936 = F_dictFind(m, v935, v149)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L5
	} else {
		goto L222
	}
L219:
	;
	v916 = F_humanNodename(m, v907)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+156)) = v907 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+152)) = v916
	*(*int32)(unsafe.Add(mBase, uint32(v34)+148)) = v907 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+144)) = v149
	F__serverLog(m, int32(2), int32(_a302), v34+int32(144))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	goto L218
L222:
	;
	if v936 == int32(0) {
		goto L212
	} else {
		goto L223
	}
L223:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v941)+44))
	v943 = F_dictDelete(m, v942, v149)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	goto L212
L225:
	;
	v1094 = base.B2i32(v901 == v55)
	v1095 = F_clusterDelSlot(m, v149)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L5
	} else {
		goto L269
	}
L226:
	;
	if v950 == int32(0) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v950)+8))
	goto L228
L228:
	;
	if v954 == int32(0) {
		goto L225
	} else {
		goto L229
	}
L229:
	;
	if v954 == l0 {
		goto L225
	} else {
		goto L230
	}
L230:
	;
	v959 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v961 = v954 + int32(48)
	v962 = int32(40)
	goto L236
L231:
	;
	if int32(2) < v959 {
		goto L262
	} else {
		goto L263
	}
L232:
	;
	if v1026 != 0 {
		goto L231
	} else {
		goto L248
	}
L233:
	;
	v1026 = int32(0)
	goto L232
L234:
	;
	v998 = v993
	v999 = v994
	v1000 = v995
	goto L244
L235:
	;
	if v983 == int32(0) {
		goto L233
	} else {
		goto L242
	}
L236:
	;
	if (v961|v59)&int32(3) != 0 {
		v993 = v59
		v994 = v961
		v995 = v962
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v970 = v59
	v971 = v961
	v972 = v962
	goto L238
L238:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v970)))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v971)))
	if v975 != v976 {
		v993 = v970
		v994 = v971
		v995 = v972
		goto L234
	} else {
		goto L240
	}
L239:
	;
	goto L235
L240:
	;
	v978 = int32(4)
	v979 = v971 + v978
	v981 = v970 + v978
	v983 = v972 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v983) {
		v970 = v981
		v971 = v979
		v972 = v983
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v993 = v981
	v994 = v979
	v995 = v983
	goto L234
L243:
	;
	v1026 = v1003 - v1004
	goto L232
L244:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999))))
	if v1003 != v1004 {
		goto L243
	} else {
		goto L246
	}
L246:
	;
	v1006 = int32(1)
	v1011 = v1000 + int32(-1)
	if v1011 == int32(0) {
		goto L233
	} else {
		goto L247
	}
L247:
	;
	v998 = v998 + v1006
	v999 = v999 + v1006
	v1000 = v1011
	goto L244
L248:
	;
	if int32(1) < v959 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+48))
	v1045 = F_dictFind(m, v1044, v149)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L5
	} else {
		goto L253
	}
L250:
	;
	v1029 = F_humanNodename(m, l0)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L5
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+124)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+120)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v34)+116)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v149
	F__serverLog(m, int32(1), int32(_a303), v34+int32(112))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L5
	} else {
		goto L252
	}
L252:
	;
	goto L249
L253:
	;
	if l0 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+48))
	if v1045 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	if v1045 == int32(0) {
		goto L225
	} else {
		goto L256
	}
L256:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+48))
	v1052 = F_dictDelete(m, v1051, v149)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L5
	} else {
		goto L257
	}
L257:
	;
	goto L225
L258:
	;
	v1060 = F_dictAdd(m, v1056, v149, l0)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L5
	} else {
		goto L261
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1045)+8)) = l0
	goto L260
L260:
	;
	goto L225
L261:
	;
	goto L225
L262:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+48))
	v1082 = F_dictFind(m, v1081, v149)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L5
	} else {
		goto L266
	}
L263:
	;
	v1064 = F_humanNodename(m, v954)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+140)) = v961
	*(*int32)(unsafe.Add(mBase, uint32(v34)+136)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(v34)+132)) = v954 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v149
	F__serverLog(m, int32(2), int32(_a304), v34+int32(128))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	goto L262
L266:
	;
	if v1082 == int32(0) {
		goto L225
	} else {
		goto L267
	}
L267:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1087)+48))
	v1089 = F_dictDelete(m, v1088, v149)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	goto L225
L269:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1098+v192)+52))
	if v1100 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	if v901 == v55 {
		goto L286
	} else {
		goto L287
	}
L271:
	;
	v1108 = m.G0
	v1110 = v1108 - int32(32)
	m.G0 = v1110
	v1115 = int32(1) << (uint(v149&int32(7)) % 32)
	v1117 = base.I32_div_s(v149, int32(8))
	v1120 = l0 + v1117 + int32(104)
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	if v1115&v1121 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1188 = v179 ^ int32(-1)
	v1189 = v1098
	goto L270
L273:
	;
	v1159 = int32(_a20)
	v1160 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	*(*int32)(unsafe.Add(mBase, uint32(v1160+v192)+52)) = l0
	v1163 = v1160 + v181
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+uint32(_consts[162]))))
	v1166 = v179 ^ int32(-1)
	v1167 = v1164 & v1166
	*(*uint8)(unsafe.Add(mBase, uint32(v1163)+uint32(_consts[162]))) = uint8(v1167)
	v1170 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1173 = v1170 + v149*int32(24)
	v1176 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1173)+uint32(_consts[163]))) = v1176
	*(*int64)(unsafe.Add(mBase, uint32(v1173)+uint32(_consts[164]))) = v1176
	*(*int64)(unsafe.Add(mBase, uint32(v1173)+uint32(_consts[165]))) = v1176
	goto L285
L274:
	;
	m.G0 = v1110 + int32(32)
	goto L273
L275:
	;
	v1123 = v1121 | v1115
	*(*uint8)(unsafe.Add(mBase, uint32(v1120))) = uint8(v1123)
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2160))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+2160)) = v1125 + int32(1)
	if v1125 != 0 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+32))
	F_dictInitIterator(m, v1110, v1131)
	mBase = m.M
	v1133 = F_dictNext(m, v1110)
	mBase = m.M
	if v1133 == int32(0) {
		goto L274
	} else {
		goto L277
	}
L277:
	;
	v1137 = v1133
	goto L279
L278:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v1147 | int32(256)
	goto L274
L279:
	;
	v1141 = F_dictGetVal(m, v1137)
	mBase = m.M
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1141)+88)))
	if v1142&int32(2) != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1146 = F_dictNext(m, v1110)
	mBase = m.M
	if v1146 != 0 {
		v1137 = v1146
		goto L279
	} else {
		goto L284
	}
L282:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+2164))
	if v1145 != 0 {
		goto L278
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	goto L274
L285:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1188 = v1166
	v1189 = v1187
	goto L270
L286:
	;
	v1191 = l0
	goto L288
L287:
	;
	v1191 = v158
	goto L288
L288:
	;
	v1195 = v1189 + v181
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1195)+uint32(_consts[162]))))
	v1197 = v1196 & v1188
	*(*uint8)(unsafe.Add(mBase, uint32(v1195)+uint32(_consts[162]))) = uint8(v1197)
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v1202)+uint32(_consts[122]))) = v1203 | int32(14)
	v1207 = v1191
	v1208 = v709
	v1209 = v708
	v1210 = v707
	v1211 = v775
	v1212 = v163 + v1094
	v1213 = v188
	v1214 = v165 + v893
	v1215 = v166 + v831
	goto L28
L290:
	;
	goto L27
L291:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, _consts[114])))
	if v1268&int32(4) != 0 {
		goto L1
	} else {
		goto L297
	}
L292:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1232 {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1235 = F_humanNodename(m, v1208)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	v1237 = F_humanNodename(m, l0)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(76)))) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(72)))) = v1237
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(68)))) = v57
	v1250 = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(64)))) = v1208 + v1250
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v1235
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v1208 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v1210
	F__serverLog(m, int32(2), int32(_a299), v34+v1250)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	goto L291
L297:
	;
	if v1207 != 0 {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	if v1215 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L299:
	;
	v1345 = base.B2i32(v1211 != int32(0))
	goto L298
L300:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v55)+2160))
	if v1283 != 0 {
		goto L299
	} else {
		goto L311
	}
L301:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+2172))
	if v1273 == l0 {
		goto L299
	} else {
		goto L302
	}
L302:
	;
	if v1213 != 0 {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+2160))
	if v1275 != 0 {
		goto L299
	} else {
		goto L304
	}
L304:
	;
	if l0 == int32(0) {
		goto L299
	} else {
		goto L305
	}
L305:
	;
	if v1273 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1278 = v1273
	goto L308
L307:
	;
	v1278 = v1272
	goto L308
L308:
	;
	v1279 = *(*int64)(unsafe.Add(mBase, uint32(v1278)+96))
	if base.Ui64(l1) <= base.Ui64(v1279) {
		goto L299
	} else {
		goto L309
	}
L309:
	;
	if v126 != 0 {
		goto L299
	} else {
		goto L310
	}
L310:
	;
	goto L300
L311:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v1285 != 0 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	v1315 = int32(0)
	v1317 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317)+88)))
	if v1318&int32(1) == v1315 {
		v1345 = v1315
		goto L298
	} else {
		goto L322
	}
L313:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1287 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	if v126 != 0 {
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	v1302 = int32(0)
	v1304 = base.B2i32(v126 != v1302)
	F_clusterSetPrimary(m, l0, v1304, v1304)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L5
	} else {
		goto L320
	}
L317:
	;
	v1290 = F_humanNodename(m, l0)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v57
	F__serverLog(m, int32(2), int32(_a305), v34+int32(16))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	goto L316
L320:
	;
	F_clearCachedClusterSlotsResponse(m)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1310)+uint32(_consts[122])))
	*(*int32)(unsafe.Add(mBase, uint32(v1310)+uint32(_consts[122]))) = v1311 | int32(46)
	v1345 = v1302
	goto L298
L322:
	;
	if v1213 < v1212 {
		v1345 = v1315
		goto L298
	} else {
		goto L323
	}
L323:
	;
	v1324 = int32(1)
	v1326 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1326 {
		v1345 = v1324
		goto L298
	} else {
		goto L324
	}
L324:
	;
	v1329 = F_humanNodename(m, l0)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v57
	F__serverLog(m, int32(2), int32(_a306), v34+int32(32))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1345 = v1324
	goto L298
L327:
	;
	if v1214 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L328:
	;
	F_clusterUpdateSlotExportsOnOwnershipChange(m)
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	v1354 = int32(0)
	if v1345&base.B2i32(v1354 < v1211) != int32(1) {
		goto L1
	} else {
		goto L333
	}
L331:
	;
	F_clusterUpdateSlotImportsOnOwnershipChange(m)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	v1364 = v1354
	goto L334
L334:
	;
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(192)+v1364<<(uint(int32(1))%32)))))
	v1398 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v1398 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	goto L1
L336:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1425 = F_delKeysInSlot(m, v1396, v1422, int32(1), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L5
	} else {
		goto L340
	}
L337:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v1403 = F_humanNodename(m, v1402)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v1402 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v1403
	v1411 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v1411 + int32(48)
	F__serverLog(m, int32(2), int32(_a307), v34)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	goto L336
L340:
	;
	v1428 = v1364 + int32(1)
	if v1428 != v1211 {
		v1364 = v1428
		goto L334
	} else {
		goto L341
	}
L341:
	;
	goto L335
}
func F_clusterWriteHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int64
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = int32(0)
	goto L4
L1:
	;
	F__serverAssert(m, int32(_a311), int32(_a247), int32(4539))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L33
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return
L3:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+80))
	v144 = m.T0[v143].(func(*base.Module, int32, int32, int32) int32)(m, v139, v140, v140)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L32
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	if v34 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v128 != 0 {
		goto L2
	} else {
		goto L31
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v44 = F___bswap_32_2(m, v43)
	mBase = m.M
	goto L8
L7:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+20)))
	v81 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[183])))
	v83 = base.I64_extend_i32_u(v48)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[183]))) = v82 + v83
	v87 = F___bswap_16_2(m, v79)
	mBase = m.M
	goto L25
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int32)(m, l0, v38+v39+int32(8), v44-v39)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if int32(0) < v48 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	if v48 != int32(-1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_freeClusterLink(m, v18)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L21
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v70
	F__serverLog(m, int32(0), int32(_a312), v16)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v66 {
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v54 == int32(3) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(0) < v58 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+88))
	v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v70 = v63
	goto L13
L19:
	;
	v70 = int32(_a313)
	goto L13
L20:
	;
	goto L12
L21:
	;
	goto L2
L22:
	;
	v100 = v48 + v39
	if base.Ui32(v44) <= base.Ui32(v100) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v94 = v81 + v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v95 + v83
	goto L22
L24:
	;
	v93 = int32(65888)
	goto L23
L25:
	;
	switch v87&int32(32767) + int32(-4) {
	case 0, 6:
		v93 = int32(65872)
		goto L23
	default:
		goto L22
	case 5:
		goto L24
	}
L26:
	;
	if v100 != v44 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v102 + v48
	goto L2
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	F_listDelNode(m, v109, v37)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v112 = int32(_a20)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v114 + int32(-12)
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v118 - base.I64_extend_i32_u(v108+int32(12))
	v124 = v48 + v23
	if base.Ui32(v124) < base.Ui32(int32(65536)) {
		v23 = v124
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L5
L31:
	;
	goto L3
L32:
	;
	goto L2
L33:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_createClusterMsgSendBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v67 int64
	_ = v67
	var v73 int64
	_ = v73
	var v79 int64
	_ = v79
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int64
	_ = v189
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int64
	_ = v212
	var v218 int64
	_ = v218
	var v224 int64
	_ = v224
	var v230 int64
	_ = v230
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v301 int64
	_ = v301
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v414 int32
	_ = v414
	var v415 int64
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	v12 = l1 + int32(8)
	v13 = F_valkey_calloc(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(1)
	v20 = int32(_a20)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[161]))
	*(*int32)(unsafe.Add(mBase, _consts[161])) = v22 + v12
	if l0&int32(32768) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v430 = F___bswap_32_1(m, l1)
	mBase = m.M
	goto L69
L4:
	;
	v42 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+88)))
	if v44&int32(2) == v42 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(1651327826)
	v32 = F___bswap_16_1(m, int32(1))
	mBase = m.M
	goto L6
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)) = uint16(v32)
	v34 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+22)) = uint16(v34)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)) = uint16(v34)
	v40 = F___bswap_16_1(m, l0&int32(65535))
	mBase = m.M
	goto L7
L7:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v40)
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(1651327826)
	v55 = F___bswap_16_1(m, int32(1))
	mBase = m.M
	goto L12
L9:
	;
	v51 = v43
	goto L8
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2172))
	if v49 != 0 {
		v51 = v49
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)) = uint16(v55)
	v59 = F___bswap_16_1(m, l0&int32(65535))
	mBase = m.M
	goto L13
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v59)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v61
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v43+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(56)))) = v67
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v43+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(64)))) = v73
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v43+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(72)))) = v79
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v43+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(80)))) = v85
	v87 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+2176)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2184)))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2192)))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2200)))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2208)))) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2214)))) = v87
	v110 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	if v110 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v158 = int32(_a20)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v161 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v163 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v165 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v167 != 0 {
		v175 = v167
		goto L28
	} else {
		goto L29
	}
L15:
	;
	goto L19
L16:
	;
	goto L14
L17:
	;
	goto L16
L18:
	;
	v145 = v124
	goto L25
L19:
	;
	v120 = v13 + int32(2176)
	v122 = int32(46)
	v124 = v110
	goto L21
L20:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v135)
	goto L18
L21:
	;
	v126 = v122 + int32(-1)
	if v126 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v129)
	v131 = int32(1)
	if v129 != 0 {
		v120 = v120 + v131
		v122 = v126
		v124 = v124 + v131
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v147 != 0 {
		v145 = v145 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L17
L27:
	;
	goto L26
L28:
	;
	if v161 != 0 {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v169 != 0 {
		v175 = v169
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v171 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v172 = v159
	goto L33
L32:
	;
	v172 = v163
	goto L33
L33:
	;
	v175 = v172 + int32(10000)
	goto L28
L34:
	;
	v176 = v161
	goto L36
L35:
	;
	v176 = v159
	goto L36
L36:
	;
	if v165 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v177 = v165
	goto L39
L38:
	;
	v177 = v163
	goto L39
L39:
	;
	goto L42
L40:
	;
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2168)))) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2160)))) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2152)))) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2144)))) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v13)+2136)) = v189
	v205 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+2172))
	if v207 == v205 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	v185 = F__emscripten_memcpy_bulkmem(m, v13+int32(88), v51+int32(104), int32(2048))
	mBase = m.M
	goto L41
L43:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	if v240 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2136)))) = v212
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(40))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2168)))) = v218
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2160)))) = v224
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(24))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2152)))) = v230
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v207+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(2144)))) = v236
	goto L43
L45:
	;
	v241 = v176
	goto L47
L46:
	;
	v241 = v177
	goto L47
L47:
	;
	v244 = F___bswap_16_1(m, v241&int32(65535))
	mBase = m.M
	goto L48
L48:
	;
	if v240 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v245 = v177
	goto L51
L50:
	;
	v245 = v176
	goto L51
L51:
	;
	v248 = F___bswap_16_1(m, v245&int32(65535))
	mBase = m.M
	goto L52
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+2254)) = uint16(v248)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+18)) = uint16(v244)
	v253 = F___bswap_16_1(m, v175&int32(65535))
	mBase = m.M
	goto L53
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+2256)) = uint16(v253)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v206)+88))
	v258 = F___bswap_16_1(m, v255&int32(65535))
	mBase = m.M
	goto L54
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+2258)) = uint16(v258)
	v261 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+2260)) = uint8(v262)
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v261)+8))
	v265 = int64(56)
	v267 = int64(65280)
	v269 = int64(40)
	v272 = int64(16711680)
	v274 = int64(24)
	v276 = int64(4278190080)
	v278 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v264<<(uint(v265)%64) | v264&v267<<(uint(v269)%64) | (v264&v272<<(uint(v274)%64) | v264&v276<<(uint(v278)%64)) | (int64(base.Ui64(v264)>>(uint(v278)%64))&v276 | int64(base.Ui64(v264)>>(uint(v274)%64))&v272 | (int64(base.Ui64(v264)>>(uint(v269)%64))&v267 | int64(base.Ui64(v264)>>(uint(v265)%64))))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v51)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v301<<(uint(v265)%64) | v301&v267<<(uint(v269)%64) | (v301&v272<<(uint(v274)%64) | v301&v276<<(uint(v278)%64)) | (int64(base.Ui64(v301)>>(uint(v278)%64))&v276 | int64(base.Ui64(v301)>>(uint(v274)%64))&v272 | (int64(base.Ui64(v301)>>(uint(v269)%64))&v267 | int64(base.Ui64(v301)>>(uint(v265)%64))))
	if v255&int32(2) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v373 = int64(56)
	v375 = int64(65280)
	v377 = int64(40)
	v380 = int64(16711680)
	v382 = int64(24)
	v384 = int64(4278190080)
	v386 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v372<<(uint(v373)%64) | v372&v375<<(uint(v377)%64) | (v372&v380<<(uint(v382)%64) | v372&v384<<(uint(v386)%64)) | (int64(base.Ui64(v372)>>(uint(v386)%64))&v384 | int64(base.Ui64(v372)>>(uint(v382)%64))&v380 | (int64(base.Ui64(v372)>>(uint(v377)%64))&v375 | int64(base.Ui64(v372)>>(uint(v373)%64))))
	if v371&int32(1) == int32(0) {
		goto L3
	} else {
		goto L67
	}
L56:
	;
	v370 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v371 = v255
	v372 = v370
	goto L55
L57:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v346 == int32(0) {
		v360 = int64(0)
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+88))
	v371 = v368
	v372 = v365
	goto L55
L59:
	;
	v362 = int64(0)
	if v362 < v360 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v350 != 0 {
		v357 = v350
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+104))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v358)+48))
	v360 = v359
	goto L59
L62:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	if v353 == int32(0) {
		v360 = int64(0)
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v357 = v353
	goto L61
L64:
	;
	v365 = v360
	goto L66
L65:
	;
	v365 = v362
	goto L66
L66:
	;
	goto L58
L67:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v414)+uint32(_consts[147])))
	if v415 == int64(0) {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2261)))
	v420 = v418 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+2261)) = uint8(v420)
	goto L3
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v430
	return v13
}
func F_freeClusterLink(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a272), int32(_a247), int32(1788))
		mBase = m.M
		v103 = m.ExcPending
		if v103 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(0) < v15 {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v44 == int32(0) {
				v53 = int32(_a20)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
				v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
				*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
				F_listRelease(m, v54)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = int32(_a20)
					v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					F_valkey_free(m, v72)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v75 == int32(0) {
							F_valkey_free(m, l0)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
							if v78 != l0 {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
								if v83 != l0 {
									F_valkey_free(m, l0)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v85 == int32(0) {
										F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
										v90 = F_mstime(m)
										mBase = m.M
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
										F_valkey_free(m, l0)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								if v80 != 0 {
									F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
									F_valkey_free(m, l0)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
				m.T0[v48].(func(*base.Module, int32))(m, v44)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v53 = int32(_a20)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
					v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
					*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
					F_listRelease(m, v54)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v66 = int32(_a20)
						v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						F_valkey_free(m, v72)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v75 == int32(0) {
								F_valkey_free(m, l0)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
								if v78 != l0 {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
									if v83 != l0 {
										F_valkey_free(m, l0)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										if v85 == int32(0) {
											F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
											v90 = F_mstime(m)
											mBase = m.M
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
											*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
											F_valkey_free(m, l0)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									if v80 != 0 {
										F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
										F_valkey_free(m, l0)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
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
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v20 != 0 {
				v21 = int32(_a275)
			} else {
				v21 = int32(_a276)
			}
			v22 = int32(_a277)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v23 != 0 {
				v27 = v23 + int32(8)
			} else {
				v27 = v22
			}
			if v23 == int32(0) {
				v32 = v22
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
				F__serverLog(m, int32(0), int32(_a278), v10)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v44 == int32(0) {
						v53 = int32(_a20)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
						v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
						*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
						F_listRelease(m, v54)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v66 = int32(_a20)
							v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							F_valkey_free(m, v72)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v75 == int32(0) {
									F_valkey_free(m, l0)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
									if v78 != l0 {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
										if v83 != l0 {
											F_valkey_free(m, l0)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v85 == int32(0) {
												F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
												v90 = F_mstime(m)
												mBase = m.M
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
												*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
												F_valkey_free(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									} else {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										if v80 != 0 {
											F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
											F_valkey_free(m, l0)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
						m.T0[v48].(func(*base.Module, int32))(m, v44)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							v53 = int32(_a20)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
							v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
							*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
							F_listRelease(m, v54)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = int32(_a20)
								v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								F_valkey_free(m, v72)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v75 == int32(0) {
										F_valkey_free(m, l0)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
										if v78 != l0 {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
											if v83 != l0 {
												F_valkey_free(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												if v85 == int32(0) {
													F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
													v90 = F_mstime(m)
													mBase = m.M
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
													F_valkey_free(m, l0)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v80 != 0 {
												F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
												F_valkey_free(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
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
				v30 = F_humanNodename(m, v23)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v32 = v30
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v27
					F__serverLog(m, int32(0), int32(_a278), v10)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v44 == int32(0) {
							v53 = int32(_a20)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
							v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
							*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
							F_listRelease(m, v54)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = int32(_a20)
								v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								F_valkey_free(m, v72)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									if v75 == int32(0) {
										F_valkey_free(m, l0)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
										if v78 != l0 {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
											if v83 != l0 {
												F_valkey_free(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												if v85 == int32(0) {
													F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
													v90 = F_mstime(m)
													mBase = m.M
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
													*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
													F_valkey_free(m, l0)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										} else {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											if v80 != 0 {
												F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
												F_valkey_free(m, l0)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
							m.T0[v48].(func(*base.Module, int32))(m, v44)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								v53 = int32(_a20)
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
								v59 = *(*int32)(unsafe.Add(mBase, _consts[161]))
								*(*int32)(unsafe.Add(mBase, _consts[161])) = v55*int32(-12) + v59 + int32(-24)
								F_listRelease(m, v54)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v66 = int32(_a20)
									v68 = *(*int32)(unsafe.Add(mBase, _consts[161]))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									*(*int32)(unsafe.Add(mBase, _consts[161])) = v68 - v69
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									F_valkey_free(m, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
										if v75 == int32(0) {
											F_valkey_free(m, l0)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2344))
											if v78 != l0 {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+2348))
												if v83 != l0 {
													F_valkey_free(m, l0)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												} else {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													if v85 == int32(0) {
														F__serverAssert(m, int32(_a273), int32(_a247), int32(1807))
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v75)+2348)) = int32(0)
														v90 = F_mstime(m)
														mBase = m.M
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														*(*int64)(unsafe.Add(mBase, uint32(v91)+2240)) = v90
														F_valkey_free(m, l0)
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											} else {
												v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												if v80 != 0 {
													F__serverAssert(m, int32(_a274), int32(_a247), int32(1804))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v75)+2344)) = int32(0)
													F_valkey_free(m, l0)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
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
func F_genClusterInfoString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
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
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
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
	var v358 int32
	_ = v358
	var v359 int64
	_ = v359
	var v360 int32
	_ = v360
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v459 int64
	_ = v459
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v475 int32
	_ = v475
	var v476 int64
	_ = v476
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v492 int32
	_ = v492
	var v493 int64
	_ = v493
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int64
	_ = v525
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int64
	_ = v559
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int64
	_ = v600
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int64
	_ = v617
	var v618 int32
	_ = v618
	var v619 int64
	_ = v619
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v635 int32
	_ = v635
	var v636 int64
	_ = v636
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int64
	_ = v651
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v670 int64
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int64
	_ = v685
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int64
	_ = v702
	var v703 int32
	_ = v703
	var v704 int64
	_ = v704
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int64
	_ = v719
	var v720 int32
	_ = v720
	var v721 int64
	_ = v721
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v754 int32
	_ = v754
	var v755 int64
	_ = v755
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int64
	_ = v770
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int64
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int64
	_ = v794
	var v795 int64
	_ = v795
	var v796 int64
	_ = v796
	var v797 int64
	_ = v797
	var v798 int64
	_ = v798
	var v801 int64
	_ = v801
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int64
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(528)
	m.G0 = v22
	v26 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	v36 = F_dictGetIterator(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+2172))
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v31 = int64(0)
	goto L1
L4:
	;
	v29 = v28
	goto L6
L5:
	;
	v29 = v26
	goto L6
L6:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+96))
	v31 = v30
	goto L1
L7:
	;
	F_dictReleaseIterator(m, v36)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L79
	}
L8:
	;
	v154 = int32(0)
	v163 = v2
	v164 = v143
	v168 = v154
	v169 = v154
	v170 = v154
	v171 = v154
	v172 = v154
	v173 = v154
	v174 = v154
	goto L38
L9:
	;
	return int32(0)
L10:
	;
	v47 = v36 + int32(20)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if v143 != 0 {
		goto L8
	} else {
		goto L37
	}
L12:
	;
	v54 = v47
	v55 = v51
	goto L15
L13:
	;
	v51 = int32(1)
	goto L12
L14:
	;
	v51 = int32(0)
	goto L12
L15:
	;
	switch v55 {
	case 0:
		goto L20
	default:
		goto L19
	}
L17:
	;
	v55 = int32(0)
	goto L15
L18:
	;
	goto L11
L19:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v135
	if v135 == int32(0) {
		goto L17
	} else {
		goto L36
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v59 != int32(-1) {
		v98 = v59
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v99 = int32(1)
	v100 = v98 + v99
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v100
	v102 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v106+int32(26)))))
	if v110 == int32(255) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v63 != 0 {
		v98 = int32(-1)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v65 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v92 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v72 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v64)+16)))
	v73 = int64(*(*int8)(unsafe.Add(mBase, uint32(v64)+27)))
	v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v75 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v64)+12)))
	v76 = int64(*(*int8)(unsafe.Add(mBase, uint32(v64)+26)))
	v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+4)))
	v78 = F_wangHash64(m, v77)
	mBase = m.M
	v80 = F_wangHash64(m, v76+v78)
	mBase = m.M
	v82 = F_wangHash64(m, v75+v80)
	mBase = m.M
	v84 = F_wangHash64(m, v74+v82)
	mBase = m.M
	v86 = F_wangHash64(m, v73+v84)
	mBase = m.M
	v88 = F_wangHash64(m, v72+v86)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v91 = v90
	goto L24
L26:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+24)))
	v70 = v68 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+24)) = uint16(v70)
	v91 = v64
	goto L24
L27:
	;
	v98 = v92 + int32(-1)
	goto L21
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v98 = v95
	goto L21
L29:
	;
	v125 = int32(2)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v105+v123<<(uint(v125)%32)+int32(4))))
	v54 = v130 + v124<<(uint(v125)%32)
	v55 = int32(1)
	goto L15
L30:
	;
	v114 = v102
	goto L32
L31:
	;
	v114 = v99 << (uint(v110) % 32)
	goto L32
L32:
	;
	if v100 < v114 {
		v123 = v106
		v124 = v100
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v106 != 0 {
		v143 = v102
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v116 == int32(-1) {
		v143 = v102
		goto L18
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(4294967296)
	v123 = int32(1)
	v124 = int32(0)
	goto L29
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v139
	v143 = v135
	goto L18
L37:
	;
	v147 = int32(0)
	v334 = v2
	v339 = v147
	v340 = v147
	v341 = v147
	v342 = v147
	v343 = v147
	v344 = v147
	v345 = v147
	goto L7
L38:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	goto L42
L39:
	;
	v334 = v223
	v339 = v197
	v340 = v198
	v341 = v199
	v342 = v200
	v343 = v211
	v344 = v224
	v345 = v212
	goto L7
L40:
	;
	if v196&int32(4) == int32(0) {
		v211 = v172
		v212 = v174
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v183 = v181 + v171
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+88))
	if v184&int32(8) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+2160))
	if v181 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+88))
	v196 = v182
	v197 = v168
	v198 = v169
	v199 = v170
	v200 = v171
	goto L40
L44:
	;
	if v184&int32(4) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v196 = v184
	v197 = v181 + v168
	v198 = v169
	v199 = v170
	v200 = v183
	goto L40
L46:
	;
	v196 = v184
	v197 = v168
	v198 = v169
	v199 = v181 + v170
	v200 = v183
	goto L40
L47:
	;
	v196 = v184
	v197 = v168
	v198 = v181 + v169
	v199 = v170
	v200 = v183
	goto L40
L48:
	;
	if v196&int32(8) == int32(0) {
		v223 = v163
		v224 = v173
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v211 = v172 + int32(1)
	v212 = v174 + base.B2i32(v181 != int32(0))&v196
	goto L48
L50:
	;
	v232 = v36 + int32(20)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v233 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v223 = v163 + base.B2i32(v181 != int32(0))&v196
	v224 = v173 + int32(1)
	goto L50
L52:
	;
	if v328 != 0 {
		v163 = v223
		v164 = v328
		v168 = v197
		v169 = v198
		v170 = v199
		v171 = v200
		v172 = v211
		v173 = v224
		v174 = v212
		goto L38
	} else {
		goto L78
	}
L53:
	;
	v239 = v232
	v240 = v236
	goto L56
L54:
	;
	v236 = int32(1)
	goto L53
L55:
	;
	v236 = int32(0)
	goto L53
L56:
	;
	switch v240 {
	case 0:
		goto L61
	default:
		goto L60
	}
L58:
	;
	v240 = int32(0)
	goto L56
L59:
	;
	goto L52
L60:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v320
	if v320 == int32(0) {
		goto L58
	} else {
		goto L77
	}
L61:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v244 != int32(-1) {
		v283 = v244
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v284 = int32(1)
	v285 = v283 + v284
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v285
	v287 = int32(0)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v291+int32(26)))))
	if v295 == int32(255) {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v248 != 0 {
		v283 = int32(-1)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v250 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+20))
	if v277 != int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v257 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v249)+16)))
	v258 = int64(*(*int8)(unsafe.Add(mBase, uint32(v249)+27)))
	v259 = int64(*(*int32)(unsafe.Add(mBase, uint32(v249)+8)))
	v260 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v249)+12)))
	v261 = int64(*(*int8)(unsafe.Add(mBase, uint32(v249)+26)))
	v262 = int64(*(*int32)(unsafe.Add(mBase, uint32(v249)+4)))
	v263 = F_wangHash64(m, v262)
	mBase = m.M
	v265 = F_wangHash64(m, v261+v263)
	mBase = m.M
	v267 = F_wangHash64(m, v260+v265)
	mBase = m.M
	v269 = F_wangHash64(m, v259+v267)
	mBase = m.M
	v271 = F_wangHash64(m, v258+v269)
	mBase = m.M
	v273 = F_wangHash64(m, v257+v271)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v276 = v275
	goto L65
L67:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+24)))
	v255 = v253 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+24)) = uint16(v255)
	v276 = v249
	goto L65
L68:
	;
	v283 = v277 + int32(-1)
	goto L62
L69:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v283 = v280
	goto L62
L70:
	;
	v310 = int32(2)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v290+v308<<(uint(v310)%32)+int32(4))))
	v239 = v315 + v309<<(uint(v310)%32)
	v240 = int32(1)
	goto L56
L71:
	;
	v299 = v287
	goto L73
L72:
	;
	v299 = v284 << (uint(v295) % 32)
	goto L73
L73:
	;
	if v285 < v299 {
		v308 = v291
		v309 = v285
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v291 != 0 {
		v328 = v287
		goto L59
	} else {
		goto L75
	}
L75:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	if v301 == int32(-1) {
		v328 = v287
		goto L59
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(4294967296)
	v308 = int32(1)
	v309 = int32(0)
	goto L70
L77:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v324
	v328 = v320
	goto L59
L78:
	;
	goto L39
L79:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354)+32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v354)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(464)))) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(468)))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(472)))) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(476)))) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(480)))) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(496)))) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(504)))) = v359
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(512)))) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(488)))) = base.I64_extend_i32_u(v357 + v358)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+452)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v22)+456)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v22)+460)) = v340
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v355<<(uint(int32(2))%32))+uint32(_consts[186])))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+448)) = v397
	v402 = F_sdscatfmt(m, l0, int32(_a321), v22+int32(448))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v405)+uint32(_consts[182])))
	if base.B2i32(v406 == int64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v422)+uint32(_consts[187])))
	if v425 == int64(0) {
		v439 = v422
		v440 = v423
		v441 = v424
		goto L85
	} else {
		goto L86
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+440)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v22)+432)) = int32(_a322)
	v418 = F_sdscatfmt(m, v402, int32(_a323), v22+int32(432))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L9
	} else {
		goto L84
	}
L83:
	;
	v422 = v405
	v423 = int64(0)
	v424 = v402
	goto L81
L84:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v422 = v421
	v423 = v406
	v424 = v418
	goto L81
L85:
	;
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v439)+uint32(_consts[188])))
	if v442 == int64(0) {
		v456 = v439
		v457 = v440
		v458 = v441
		goto L88
	} else {
		goto L89
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+416)) = int32(_a324)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+424)) = v425
	v435 = F_sdscatfmt(m, v424, int32(_a323), v22+int32(416))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v439 = v438
	v440 = v425 + v423
	v441 = v435
	goto L85
L88:
	;
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v456)+uint32(_consts[189])))
	if v459 == int64(0) {
		v473 = v456
		v474 = v457
		v475 = v458
		goto L91
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+400)) = int32(_a325)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+408)) = v442
	v452 = F_sdscatfmt(m, v441, int32(_a323), v22+int32(400))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v456 = v455
	v457 = v442 + v440
	v458 = v452
	goto L88
L91:
	;
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v473)+uint32(_consts[190])))
	if v476 == int64(0) {
		v490 = v473
		v491 = v474
		v492 = v475
		goto L94
	} else {
		goto L95
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+384)) = int32(_a326)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+392)) = v459
	v469 = F_sdscatfmt(m, v458, int32(_a323), v22+int32(384))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v473 = v472
	v474 = v459 + v457
	v475 = v469
	goto L91
L94:
	;
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v490)+uint32(_consts[191])))
	if v493 == int64(0) {
		v507 = v490
		v508 = v491
		v509 = v492
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+368)) = int32(_a327)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+376)) = v476
	v486 = F_sdscatfmt(m, v475, int32(_a323), v22+int32(368))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v490 = v489
	v491 = v476 + v474
	v492 = v486
	goto L94
L97:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v507)+uint32(_consts[192])))
	if v510 == int64(0) {
		v524 = v507
		v525 = v508
		v526 = v509
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+352)) = int32(_a328)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+360)) = v493
	v503 = F_sdscatfmt(m, v492, int32(_a323), v22+int32(352))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v507 = v506
	v508 = v493 + v491
	v509 = v503
	goto L97
L100:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v524)+uint32(_consts[193])))
	if v527 == int64(0) {
		v541 = v524
		v542 = v525
		v543 = v526
		goto L103
	} else {
		goto L104
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = int32(_a329)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+344)) = v510
	v520 = F_sdscatfmt(m, v509, int32(_a323), v22+int32(336))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v524 = v523
	v525 = v510 + v508
	v526 = v520
	goto L100
L103:
	;
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v541)+uint32(_consts[194])))
	if v544 == int64(0) {
		v558 = v541
		v559 = v542
		v560 = v543
		goto L106
	} else {
		goto L107
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = int32(_a330)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+328)) = v527
	v537 = F_sdscatfmt(m, v526, int32(_a323), v22+int32(320))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v541 = v540
	v542 = v527 + v525
	v543 = v537
	goto L103
L106:
	;
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v558)+uint32(_consts[195])))
	if v561 == int64(0) {
		v575 = v558
		v576 = v559
		v577 = v560
		goto L109
	} else {
		goto L110
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+304)) = int32(_a331)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+312)) = v544
	v554 = F_sdscatfmt(m, v543, int32(_a323), v22+int32(304))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v558 = v557
	v559 = v544 + v542
	v560 = v554
	goto L106
L109:
	;
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v575)+uint32(_consts[196])))
	if v578 == int64(0) {
		v590 = v576
		v591 = v577
		goto L112
	} else {
		goto L113
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = int32(_a332)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+296)) = v561
	v571 = F_sdscatfmt(m, v560, int32(_a323), v22+int32(288))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v575 = v574
	v576 = v561 + v559
	v577 = v571
	goto L109
L112:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+256)) = v590
	v596 = F_sdscatfmt(m, v591, int32(_a333), v22+int32(256))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L9
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(_a334)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+280)) = v578
	v588 = F_sdscatfmt(m, v577, int32(_a323), v22+int32(272))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	v590 = v578 + v576
	v591 = v588
	goto L112
L115:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v599)+uint32(_consts[197])))
	if base.B2i32(v600 == int64(0)) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v616)+uint32(_consts[198])))
	if v619 == int64(0) {
		v633 = v616
		v634 = v617
		v635 = v618
		goto L120
	} else {
		goto L121
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+248)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = int32(_a322)
	v612 = F_sdscatfmt(m, v596, int32(_a335), v22+int32(240))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L9
	} else {
		goto L119
	}
L118:
	;
	v616 = v599
	v617 = int64(0)
	v618 = v596
	goto L116
L119:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v616 = v615
	v617 = v600
	v618 = v612
	goto L116
L120:
	;
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v633)+uint32(_consts[199])))
	if v636 == int64(0) {
		v650 = v633
		v651 = v634
		v652 = v635
		goto L123
	} else {
		goto L124
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = int32(_a324)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = v619
	v629 = F_sdscatfmt(m, v618, int32(_a335), v22+int32(224))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v633 = v632
	v634 = v619 + v617
	v635 = v629
	goto L120
L123:
	;
	v653 = *(*int64)(unsafe.Add(mBase, uint32(v650)+uint32(_consts[200])))
	if v653 == int64(0) {
		v667 = v650
		v668 = v651
		v669 = v652
		goto L126
	} else {
		goto L127
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = int32(_a325)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+216)) = v636
	v646 = F_sdscatfmt(m, v635, int32(_a335), v22+int32(208))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v650 = v649
	v651 = v636 + v634
	v652 = v646
	goto L123
L126:
	;
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v667)+uint32(_consts[201])))
	if v670 == int64(0) {
		v684 = v667
		v685 = v668
		v686 = v669
		goto L129
	} else {
		goto L130
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = int32(_a326)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+200)) = v653
	v663 = F_sdscatfmt(m, v652, int32(_a335), v22+int32(192))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	v666 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v667 = v666
	v668 = v653 + v651
	v669 = v663
	goto L126
L129:
	;
	v687 = *(*int64)(unsafe.Add(mBase, uint32(v684)+uint32(_consts[202])))
	if v687 == int64(0) {
		v701 = v684
		v702 = v685
		v703 = v686
		goto L132
	} else {
		goto L133
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = int32(_a327)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v670
	v680 = F_sdscatfmt(m, v669, int32(_a335), v22+int32(176))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L9
	} else {
		goto L131
	}
L131:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v684 = v683
	v685 = v670 + v668
	v686 = v680
	goto L129
L132:
	;
	v704 = *(*int64)(unsafe.Add(mBase, uint32(v701)+uint32(_consts[203])))
	if v704 == int64(0) {
		v718 = v701
		v719 = v702
		v720 = v703
		goto L135
	} else {
		goto L136
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = int32(_a328)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+168)) = v687
	v697 = F_sdscatfmt(m, v686, int32(_a335), v22+int32(160))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v701 = v700
	v702 = v687 + v685
	v703 = v697
	goto L132
L135:
	;
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v718)+uint32(_consts[204])))
	if v721 == int64(0) {
		v735 = v718
		v736 = v719
		v737 = v720
		goto L138
	} else {
		goto L139
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = int32(_a329)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v704
	v714 = F_sdscatfmt(m, v703, int32(_a335), v22+int32(144))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v718 = v717
	v719 = v704 + v702
	v720 = v714
	goto L135
L138:
	;
	v738 = *(*int64)(unsafe.Add(mBase, uint32(v735)+uint32(_consts[205])))
	if v738 == int64(0) {
		v752 = v735
		v753 = v736
		v754 = v737
		goto L141
	} else {
		goto L142
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = int32(_a330)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+136)) = v721
	v731 = F_sdscatfmt(m, v720, int32(_a335), v22+int32(128))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v735 = v734
	v736 = v721 + v719
	v737 = v731
	goto L138
L141:
	;
	v755 = *(*int64)(unsafe.Add(mBase, uint32(v752)+uint32(_consts[206])))
	if v755 == int64(0) {
		v769 = v752
		v770 = v753
		v771 = v754
		goto L144
	} else {
		goto L145
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = int32(_a331)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v738
	v748 = F_sdscatfmt(m, v737, int32(_a335), v22+int32(112))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v752 = v751
	v753 = v738 + v736
	v754 = v748
	goto L141
L144:
	;
	v772 = *(*int64)(unsafe.Add(mBase, uint32(v769)+uint32(_consts[207])))
	if v772 == int64(0) {
		v784 = v770
		v785 = v771
		goto L147
	} else {
		goto L148
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(_a332)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v755
	v765 = F_sdscatfmt(m, v754, int32(_a335), v22+int32(96))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v769 = v768
	v770 = v755 + v753
	v771 = v765
	goto L144
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v784
	v790 = F_sdscatfmt(m, v785, int32(_a336), v22+int32(64))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L9
	} else {
		goto L150
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(_a334)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v772
	v782 = F_sdscatfmt(m, v771, int32(_a335), v22+int32(80))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	v784 = v772 + v770
	v785 = v782
	goto L147
L150:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[208])))
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[183])))
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[209])))
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[210])))
	v798 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[211])))
	v801 = *(*int64)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[212])))
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(32)))) = v801
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(40)))) = v798
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(48)))) = v797
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(56)))) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v795
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v794
	v817 = F_sdscatfmt(m, v790, int32(_a337), v22+int32(16))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	v820 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v821 = *(*int64)(unsafe.Add(mBase, uint32(v820)+uint32(_consts[117])))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v821
	v824 = F_sdscatfmt(m, v817, int32(_a338), v22)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	m.G0 = v22 + int32(528)
	return v824
}
func F_getClusterConnectionsCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v3 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		return (v9+v10)<<(uint(int32(1))%32) + int32(-2)
	} else {
		return int32(0)
	}
}
func F_initClusterSlotMigrationJobList(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = F_listCreate(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+uint32(_consts[171]))) = v2
		*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(103)
		return
	}
}
func F_resetClusterStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v3 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		v12 = F___memset(m, v7+int32(67968), int32(0), int32(393216))
		mBase = m.M
		v14 = *(*int32)(unsafe.Add(mBase, _consts[111]))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[117]))) = int64(0)
		v22 = F__emscripten_memset_bulkmem(m, v14+int32(65680), base.I32_extend8_s(int32(0)), int32(224))
		mBase = m.M
	}
	return
}
func F_setClusterNodeToInboundClusterLink(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v9 != 0 {
		F__serverAssert(m, int32(_a279), int32(_a247), int32(1816))
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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		if v10 == int32(0) {
			F__serverAssert(m, int32(_a273), int32(_a247), int32(1817))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2348))
			if v13 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+2348)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				v44 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(1) < v44 {
					m.G0 = v7 + int32(96)
					return
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v47 == int32(0) {
						v89 = F_humanNodename(m, l0)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
							F__serverLog(m, int32(1), int32(_a280), v7)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								m.G0 = v7 + int32(96)
								return
							}
						}
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
						if v51 == int32(0) {
							v89 = F_humanNodename(m, l0)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
								F__serverLog(m, int32(1), int32(_a280), v7)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return
								} else {
									m.G0 = v7 + int32(96)
									return
								}
							}
						} else {
							v60 = m.T0[v51].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v47, v7+int32(48), int32(46), v7+int32(44), int32(1))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if v60 == int32(-1) {
									if int32(1) < v63 {
										m.G0 = v7 + int32(96)
										return
									} else {
										v89 = F_humanNodename(m, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
											F__serverLog(m, int32(1), int32(_a280), v7)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v7 + int32(96)
												return
											}
										}
									}
								} else {
									if int32(1) < v63 {
										m.G0 = v7 + int32(96)
										return
									} else {
										v68 = F_humanNodename(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0 + int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v68
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v74
											*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v7 + int32(48)
											F__serverLog(m, int32(1), int32(_a281), v7+int32(16))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												m.G0 = v7 + int32(96)
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
				v17 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(0) < v17 {
					v36 = v13
					F_freeClusterLink(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2348))
						if v39 != 0 {
							F__serverAssert(m, int32(_a282), int32(_a247), int32(1830))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+2348)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
							v44 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(1) < v44 {
								m.G0 = v7 + int32(96)
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v47 == int32(0) {
									v89 = F_humanNodename(m, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
										F__serverLog(m, int32(1), int32(_a280), v7)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											m.G0 = v7 + int32(96)
											return
										}
									}
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
									if v51 == int32(0) {
										v89 = F_humanNodename(m, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
											F__serverLog(m, int32(1), int32(_a280), v7)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v7 + int32(96)
												return
											}
										}
									} else {
										v60 = m.T0[v51].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v47, v7+int32(48), int32(46), v7+int32(44), int32(1))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
											if v60 == int32(-1) {
												if int32(1) < v63 {
													m.G0 = v7 + int32(96)
													return
												} else {
													v89 = F_humanNodename(m, l0)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
														*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
														F__serverLog(m, int32(1), int32(_a280), v7)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v7 + int32(96)
															return
														}
													}
												}
											} else {
												if int32(1) < v63 {
													m.G0 = v7 + int32(96)
													return
												} else {
													v68 = F_humanNodename(m, l0)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0 + int32(8)
														*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v68
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v74
														*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v7 + int32(48)
														F__serverLog(m, int32(1), int32(_a281), v7+int32(16))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															m.G0 = v7 + int32(96)
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
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l0 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v21
					F__serverLog(m, int32(0), int32(_a283), v7+int32(32))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2348))
						v36 = v35
						F_freeClusterLink(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2348))
							if v39 != 0 {
								F__serverAssert(m, int32(_a282), int32(_a247), int32(1830))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+2348)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
								v44 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(1) < v44 {
									m.G0 = v7 + int32(96)
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									if v47 == int32(0) {
										v89 = F_humanNodename(m, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
											F__serverLog(m, int32(1), int32(_a280), v7)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v7 + int32(96)
												return
											}
										}
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
										if v51 == int32(0) {
											v89 = F_humanNodename(m, l0)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
												F__serverLog(m, int32(1), int32(_a280), v7)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													m.G0 = v7 + int32(96)
													return
												}
											}
										} else {
											v60 = m.T0[v51].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v47, v7+int32(48), int32(46), v7+int32(44), int32(1))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
												if v60 == int32(-1) {
													if int32(1) < v63 {
														m.G0 = v7 + int32(96)
														return
													} else {
														v89 = F_humanNodename(m, l0)
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v89
															*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0 + int32(8)
															F__serverLog(m, int32(1), int32(_a280), v7)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																m.G0 = v7 + int32(96)
																return
															}
														}
													}
												} else {
													if int32(1) < v63 {
														m.G0 = v7 + int32(96)
														return
													} else {
														v68 = F_humanNodename(m, l0)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0 + int32(8)
															*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v68
															v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
															*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v74
															*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v7 + int32(48)
															F__serverLog(m, int32(1), int32(_a281), v7+int32(16))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																m.G0 = v7 + int32(96)
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
func F_updateClusterAnnouncedPort(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfAnnouncedPorts(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateClusterClientIpV4(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfClientIpV4(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateClusterFlags(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfFlags(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateClusterHumanNodename(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfHumanNodename(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateClusterIp(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	F_clusterUpdateMyselfIp(m)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_verifyClusterNodeId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	if l1 != int32(40) {
		v35 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v35
L2:
	;
	v10 = int32(0)
	goto L4
L3:
	;
	v35 = int32(0) - v25
	goto L1
L4:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v10))))
	v15 = int32(255)
	v25 = base.B2i32(base.Ui32((v12+int32(-123))&v15) < base.Ui32(int32(230))) & base.B2i32(base.Ui32((v12+int32(-58))&v15) < base.Ui32(int32(246)))
	if v25 != 0 {
		goto L3
	} else {
		goto L6
	}
L5:
	;
	goto L3
L6:
	;
	v27 = v10 + int32(1)
	if v27 != int32(40) {
		v10 = v27
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
