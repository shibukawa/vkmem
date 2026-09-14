package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zzlDeleteRangeByScore(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v84 float64
	_ = v84
	var v88 int64
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v15 = F_zzlFirstInRange(m, l0, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v15
	if v15 == int32(0) {
		v112 = l0
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v11 + int32(32)
	return v112
L4:
	;
	v23 = l0
	v27 = v15
	v28 = int32(0)
	goto L6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
	v112 = v106
	goto L3
L6:
	;
	v31 = F_lpNext(m, v23, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v106 = v103
	v108 = v99
	goto L5
L8:
	;
	if v31 == int32(0) {
		v106 = v23
		v108 = v28
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v39 = F_lpGetValue(m, v31, v11+int32(28), v11+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v94 != 0 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	v90 = base.F64_convert_i64_s(v88)
	goto L10
L12:
	;
	if v39 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v44 = int32(0)
	v48 = m.G0
	v50 = v48 - int32(32)
	m.G0 = v50
	v52 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v44
	v58 = *(*int64)(unsafe.Add(mBase, _consts[1101]))
	*(*int64)(unsafe.Add(mBase, uint32(v50+int32(8)))) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = int64(0)
	v63 = *(*int64)(unsafe.Add(mBase, _consts[1102]))
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v63
	F_ffc_from_chars_double_options(m, v50+int32(16), v39, v39+v43, v50+int32(24), v50)
	mBase = m.M
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v71 == v44 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v90 = v84
	goto L10
L15:
	;
	goto L20
L16:
	;
	if v71 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = int32(68)
	goto L19
L18:
	;
	v78 = int32(28)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v78
	goto L15
L20:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v50)+24))
	m.G0 = v50 + int32(32)
	goto L14
L22:
	;
	v95 = base.F64_lt(v90, v91)
	goto L24
L23:
	;
	v95 = base.F64_le(v90, v91)
	goto L24
L24:
	;
	if v95 != int32(1) {
		v106 = v23
		v108 = v28
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v99 = v28 + int32(1)
	v103 = F_lpDeleteRangeWithEntry(m, v23, v11+int32(12), int32(2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v105 != 0 {
		v23 = v103
		v27 = v105
		v28 = v99
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L7
}
func F_zzlFind(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v99 float64
	_ = v99
	var v103 int64
	_ = v103
	var v105 float64
	_ = v105
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_lpFirst(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v19 = int32(0)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
			switch v23 & int32(7) {
			case 0:
				v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
			case 1:
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
				v40 = v30
			case 2:
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
				v40 = v33
			case 3:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
				v40 = v36
			case 4:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
				v40 = v39
			default:
				v40 = v19
			}
			v42 = F_lpFind(m, l0, v14, l1, v40, int32(1))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v111 = v19
					m.G0 = v12 + int32(16)
					return v111
				} else {
					v46 = F_lpNext(m, l0, v42)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						if v46 == int32(0) {
							F__serverAssert(m, int32(_a1608), int32(_a1609), int32(1140))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if l2 == int32(0) {
								v111 = v42
								m.G0 = v12 + int32(16)
								return v111
							} else {
								v54 = F_lpGetValue(m, v46, v12+int32(12), v12)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 == int32(0) {
										v103 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
										v105 = base.F64_convert_i64_s(v103)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										v59 = int32(0)
										v63 = m.G0
										v65 = v63 - int32(32)
										m.G0 = v65
										v67 = F___errno_location(m)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v67))) = v59
										v73 = *(*int64)(unsafe.Add(mBase, _consts[1101]))
										*(*int64)(unsafe.Add(mBase, uint32(v65+int32(8)))) = v73
										*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = int64(0)
										v78 = *(*int64)(unsafe.Add(mBase, _consts[1102]))
										*(*int64)(unsafe.Add(mBase, uint32(v65))) = v78
										F_ffc_from_chars_double_options(m, v65+int32(16), v54, v54+v58, v65+int32(24), v65)
										mBase = m.M
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
										if v86 == v59 {
										} else {
											if v86 == int32(2) {
												v93 = int32(68)
											} else {
												v93 = int32(28)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v67))) = v93
										}
										v99 = *(*float64)(unsafe.Add(mBase, uint32(v65)+24))
										m.G0 = v65 + int32(32)
										v105 = v99
									}
									*(*float64)(unsafe.Add(mBase, uint32(l2))) = v105
									v111 = v42
									m.G0 = v12 + int32(16)
									return v111
								}
							}
						}
					}
				}
			}
		} else {
			v111 = int32(0)
			m.G0 = v12 + int32(16)
			return v111
		}
	}
}
func F_zzlFirstInLexRange(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	v5 = F_lpSeek(m, l0, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = F_zzlIsInLexRange(m, l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F__serverAssert(m, int32(_a1608), int32(_a1609), int32(1099))
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	return int32(0)
L5:
	;
	if v9 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v5 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v17 = v5
	goto L8
L8:
	;
	v18 = F_zzlLexValueGteMin(m, v17, l1)
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L4
L10:
	;
	v27 = F_lpNext(m, l0, v17)
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v18 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v23 = F_zzlLexValueLteMax(m, v17, l1)
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v23 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v25 = v17
	goto L16
L15:
	;
	v25 = int32(0)
	goto L16
L16:
	;
	return v25
L17:
	;
	if v27 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v31 = F_lpNext(m, l0, v27)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v31 != 0 {
		v17 = v31
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	F_abort(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zzlGetScore(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v59 float64
	_ = v59
	var v63 int64
	_ = v63
	var v65 float64
	_ = v65
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a1608), int32(_a1609), int32(857))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return float64(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v12 = F_lpGetValue(m, l0, v6+int32(12), v6)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return float64(0)
		} else {
			if v12 == int32(0) {
				v63 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v65 = base.F64_convert_i64_s(v63)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v19 = int32(0)
				v23 = m.G0
				v25 = v23 - int32(32)
				m.G0 = v25
				v27 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v19
				v33 = *(*int64)(unsafe.Add(mBase, _consts[1101]))
				*(*int64)(unsafe.Add(mBase, uint32(v25+int32(8)))) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = int64(0)
				v38 = *(*int64)(unsafe.Add(mBase, _consts[1102]))
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = v38
				F_ffc_from_chars_double_options(m, v25+int32(16), v12, v12+v18, v25+int32(24), v25)
				mBase = m.M
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				if v46 == v19 {
				} else {
					if v46 == int32(2) {
						v53 = int32(68)
					} else {
						v53 = int32(28)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v53
				}
				v59 = *(*float64)(unsafe.Add(mBase, uint32(v25)+24))
				m.G0 = v25 + int32(32)
				v65 = v59
			}
			m.G0 = v6 + int32(16)
			return v65
		}
	}
}
func F_zzlLexValueGteMin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a1608), int32(_a1609), int32(894))
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = F_lpGetValue(m, l0, v9+int32(12), v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
				v25 = F_sdsfromlonglong(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v29 == int32(0) {
						if v27 != v28 {
							v46 = int32(0)
							v48 = *(*int32)(unsafe.Add(mBase, _consts[376]))
							if v27 == v48 {
								v113 = v46
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[377]))
								if v28 == v51 {
									v113 = v46
								} else {
									if v28 != v48 {
										if v27 == v51 {
											v113 = int32(1)
										} else {
											v58 = int32(-1)
											v61 = int32(0)
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v68 & int32(7) {
											case 0:
												v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
											case 1:
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v85 = v75
											case 2:
												v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v85 = v78
											case 3:
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v85 = v81
											case 4:
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v85 = v84
											default:
												v85 = v61
											}
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v88 & int32(7) {
											case 0:
												v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
											case 1:
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v105 = v95
											case 2:
												v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v105 = v98
											case 3:
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v105 = v101
											case 4:
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v105 = v104
											default:
												v105 = v61
											}
											v106 = base.B2i32(base.Ui32(v85) < base.Ui32(v105))
											if base.Ui32(v85) < base.Ui32(v105) {
												v107 = v85
											} else {
												v107 = v105
											}
											v108 = F_memcmp(m, v27, v28, v107)
											mBase = m.M
											if v108 != 0 {
												v111 = v108
											} else {
												v111 = base.B2i32(base.Ui32(v105) < base.Ui32(v85)) - v106
											}
											v113 = base.B2i32(v58 < v111)
										}
									} else {
										v113 = int32(1)
									}
								}
							}
						} else {
							v113 = int32(1)
						}
					} else {
						v32 = int32(0)
						if v27 == v28 {
							v113 = v32
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _consts[376]))
							if v27 == v35 {
								v113 = v32
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _consts[377]))
								if v28 == v38 {
									v113 = v32
								} else {
									v40 = int32(1)
									if v28 == v35 {
										v113 = v40
									} else {
										if v27 == v38 {
											v113 = v40
										} else {
											v58 = int32(0)
											v61 = int32(0)
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v68 & int32(7) {
											case 0:
												v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
											case 1:
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v85 = v75
											case 2:
												v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v85 = v78
											case 3:
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v85 = v81
											case 4:
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v85 = v84
											default:
												v85 = v61
											}
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v88 & int32(7) {
											case 0:
												v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
											case 1:
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v105 = v95
											case 2:
												v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v105 = v98
											case 3:
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v105 = v101
											case 4:
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v105 = v104
											default:
												v105 = v61
											}
											v106 = base.B2i32(base.Ui32(v85) < base.Ui32(v105))
											if base.Ui32(v85) < base.Ui32(v105) {
												v107 = v85
											} else {
												v107 = v105
											}
											v108 = F_memcmp(m, v27, v28, v107)
											mBase = m.M
											if v108 != 0 {
												v111 = v108
											} else {
												v111 = base.B2i32(base.Ui32(v105) < base.Ui32(v85)) - v106
											}
											v113 = base.B2i32(v58 < v111)
										}
									}
								}
							}
						}
					}
					F_sdsfree(m, v27)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v113
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v22 = F_sdsnewlen(m, v15, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v27 = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v29 == int32(0) {
						if v27 != v28 {
							v46 = int32(0)
							v48 = *(*int32)(unsafe.Add(mBase, _consts[376]))
							if v27 == v48 {
								v113 = v46
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[377]))
								if v28 == v51 {
									v113 = v46
								} else {
									if v28 != v48 {
										if v27 == v51 {
											v113 = int32(1)
										} else {
											v58 = int32(-1)
											v61 = int32(0)
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v68 & int32(7) {
											case 0:
												v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
											case 1:
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v85 = v75
											case 2:
												v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v85 = v78
											case 3:
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v85 = v81
											case 4:
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v85 = v84
											default:
												v85 = v61
											}
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v88 & int32(7) {
											case 0:
												v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
											case 1:
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v105 = v95
											case 2:
												v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v105 = v98
											case 3:
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v105 = v101
											case 4:
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v105 = v104
											default:
												v105 = v61
											}
											v106 = base.B2i32(base.Ui32(v85) < base.Ui32(v105))
											if base.Ui32(v85) < base.Ui32(v105) {
												v107 = v85
											} else {
												v107 = v105
											}
											v108 = F_memcmp(m, v27, v28, v107)
											mBase = m.M
											if v108 != 0 {
												v111 = v108
											} else {
												v111 = base.B2i32(base.Ui32(v105) < base.Ui32(v85)) - v106
											}
											v113 = base.B2i32(v58 < v111)
										}
									} else {
										v113 = int32(1)
									}
								}
							}
						} else {
							v113 = int32(1)
						}
					} else {
						v32 = int32(0)
						if v27 == v28 {
							v113 = v32
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, _consts[376]))
							if v27 == v35 {
								v113 = v32
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _consts[377]))
								if v28 == v38 {
									v113 = v32
								} else {
									v40 = int32(1)
									if v28 == v35 {
										v113 = v40
									} else {
										if v27 == v38 {
											v113 = v40
										} else {
											v58 = int32(0)
											v61 = int32(0)
											v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-1)))))
											switch v68 & int32(7) {
											case 0:
												v85 = int32(base.Ui32(v68) >> (uint(int32(3)) % 32))
											case 1:
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(-3)))))
												v85 = v75
											case 2:
												v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(-5)))))
												v85 = v78
											case 3:
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-9))))
												v85 = v81
											case 4:
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-17))))
												v85 = v84
											default:
												v85 = v61
											}
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
											switch v88 & int32(7) {
											case 0:
												v105 = int32(base.Ui32(v88) >> (uint(int32(3)) % 32))
											case 1:
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
												v105 = v95
											case 2:
												v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
												v105 = v98
											case 3:
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
												v105 = v101
											case 4:
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
												v105 = v104
											default:
												v105 = v61
											}
											v106 = base.B2i32(base.Ui32(v85) < base.Ui32(v105))
											if base.Ui32(v85) < base.Ui32(v105) {
												v107 = v85
											} else {
												v107 = v105
											}
											v108 = F_memcmp(m, v27, v28, v107)
											mBase = m.M
											if v108 != 0 {
												v111 = v108
											} else {
												v111 = base.B2i32(base.Ui32(v105) < base.Ui32(v85)) - v106
											}
											v113 = base.B2i32(v58 < v111)
										}
									}
								}
							}
						}
					}
					F_sdsfree(m, v27)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v113
					}
				}
			}
		}
	}
}
