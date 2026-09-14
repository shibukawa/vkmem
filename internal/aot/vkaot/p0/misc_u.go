package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ull2string(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v4 = int32(0)
	v8 = int32(1)
	if base.Ui64(l2) < base.Ui64(int64(10)) {
		v65 = v8
		v66 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v69 = v65 + v66
	if base.Ui32(l1) <= base.Ui32(v69) {
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v16 = v4
	v17 = l2
	goto L3
L3:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v17) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v65 = v8
	v66 = v57
	goto L1
L5:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v17) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v65 = int32(2)
	v66 = v16
	goto L1
L7:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v17) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v65 = int32(3)
	v66 = v16
	goto L1
L9:
	;
	v57 = v16 + int32(12)
	v61 = base.I64_div_u_s(v17, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v17) {
		v16 = v57
		v17 = v61
		goto L3
	} else {
		goto L31
	}
L10:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v17) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v17) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v17) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v17) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v17) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v17) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v65 = int32(4)
	v66 = v16
	goto L1
L17:
	;
	v38 = int32(6)
	goto L19
L18:
	;
	v38 = int32(5)
	goto L19
L19:
	;
	v65 = v38
	v66 = v16
	goto L1
L20:
	;
	v43 = int32(8)
	goto L22
L21:
	;
	v43 = int32(7)
	goto L22
L22:
	;
	v65 = v43
	v66 = v16
	goto L1
L23:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v17) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v17) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v50 = int32(10)
	goto L27
L26:
	;
	v50 = int32(9)
	goto L27
L27:
	;
	v65 = v50
	v66 = v16
	goto L1
L28:
	;
	v55 = int32(12)
	goto L30
L29:
	;
	v55 = int32(11)
	goto L30
L30:
	;
	v65 = v55
	v66 = v16
	goto L1
L31:
	;
	goto L4
L32:
	;
	if l1 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v69))) = uint8(v72)
	v75 = v69 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(l2) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v111 = l0 + v108
	if base.Ui64(int64(9)) < base.Ui64(v109) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v82 = l2
	v84 = v75
	goto L37
L36:
	;
	v108 = v75
	v109 = l2
	goto L34
L37:
	;
	v88 = int64(100)
	v89 = base.I64_div_u_s(v82, v88)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v82-v89*v88)<<(uint(int32(1))%32))+uint32(_c_F_ull2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-1)+v84))) = uint16(v98)
	v101 = v84 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v82) {
		v82 = v89
		v84 = v101
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v108 = v101
	v109 = v89
	goto L34
L39:
	;
	goto L38
L40:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v109)<<(uint(int32(1))%32))+uint32(_c_F_ull2string[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(-1)))) = uint16(v126)
	return v69
L41:
	;
	v116 = base.I32_wrap_i64(v109) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v116)
	return v69
L42:
	;
	return int32(0)
L43:
	;
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v131)
	goto L42
}
func F_umask(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_umask[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_umask[0])) = l0
	if base.Ui32(v4) < base.Ui32(int32(-4095)) {
		v14 = v4
	} else {
		v9 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0) - v4
		v14 = int32(-1)
	}
	return v14
}
func F_uname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	if l0 != 0 {
		v5 = int32(0)
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_uname[0]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(7)))) = v6
		v9 = *(*int64)(unsafe.Add(mBase, _c_F_uname[1]))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
		v12 = *(*int64)(unsafe.Add(mBase, _c_F_uname[2]))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+65)) = v12
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_uname[3]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(72)))) = v17
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_uname[4]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+130)) = v20
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_uname[5]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(133)))) = v25
		v28 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_uname[6])))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+195)) = uint16(v28)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_uname[7])))
		*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(197)))) = uint8(v33)
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_uname[8]))
		*(*int32)(unsafe.Add(mBase, uint32(l0+int32(263)))) = v38
		v41 = *(*int32)(unsafe.Add(mBase, _c_F_uname[9]))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v41
		v44 = v5
	} else {
		v44 = int32(-21)
	}
	if base.Ui32(v44) < base.Ui32(int32(-4095)) {
		v52 = v44
	} else {
		v47 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(0) - v44
		v52 = int32(-1)
	}
	return v52
}
func F_ungetc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	v1 = l0
	v6 = int32(-1)
	if v1 == v6 {
		v43 = v6
		return v43
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		if int32(0) <= v9 {
			v16 = int32(0)
		} else {
			v16 = int32(1)
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v17 != 0 {
			v25 = v17
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			if base.Ui32(v26+int32(-8)) < base.Ui32(v25) {
				v34 = v25 + int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
				*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v1)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37 & int32(-17)
				if v16 != 0 {
				} else {
				}
				v43 = v1 & int32(255)
				return v43
			} else {
				if v16 != 0 {
					v43 = v6
					return v43
				} else {
					return int32(-1)
				}
			}
		} else {
			v18 = F___toread(m, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v22 == int32(0) {
					if v16 != 0 {
						v43 = v6
						return v43
					} else {
						return int32(-1)
					}
				} else {
					v25 = v22
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
					if base.Ui32(v26+int32(-8)) < base.Ui32(v25) {
						v34 = v25 + int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
						*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v1)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37 & int32(-17)
						if v16 != 0 {
						} else {
						}
						v43 = v1 & int32(255)
						return v43
					} else {
						if v16 != 0 {
							v43 = v6
							return v43
						} else {
							return int32(-1)
						}
					}
				}
			}
		}
	}
}
func F_unlinkCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_delGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_unprepareCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = int32(-1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v6 & int32(-1966081)
	return
}
func F_unprotectClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if int32(-1) < v4 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v4 & int32(2147483647)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 == int32(0) {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
			v16 = m.T0[v15].(func(*base.Module, int32, int32) int32)(m, v10, int32(107))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = F_clientHasPendingReplies(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v22&int32(4194304) != 0 {
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							if v25 == int32(0) {
								v33 = int32(1)
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v34 == int32(0) {
									v41 = v33
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									if v37 != 0 {
										v41 = v33
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+156))
										v41 = base.B2i32(v38 != int32(13))
									}
								}
								if v41 == int32(0) {
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v44 | int32(4194304)
									v49 = *(*int32)(unsafe.Add(mBase, _c_F_unprotectClient[0]))
									v51 = l0 + int32(168)
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
									if v54 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
										*(*int32)(unsafe.Add(mBase, uint32(v61))) = v51
										v63 = v61
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
										v56 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
										v63 = v56
									}
									*(*int32)(unsafe.Add(mBase, uint32(v49))) = v51
									*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v54 + int32(1)
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
								switch v28 {
								case 0:
									v33 = int32(1)
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
									if v34 == int32(0) {
										v41 = v33
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
										if v37 != 0 {
											v41 = v33
										} else {
											v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+156))
											v41 = base.B2i32(v38 != int32(13))
										}
									}
									if v41 == int32(0) {
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v44 | int32(4194304)
										v49 = *(*int32)(unsafe.Add(mBase, _c_F_unprotectClient[0]))
										v51 = l0 + int32(168)
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
										if v54 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
											*(*int32)(unsafe.Add(mBase, uint32(v61))) = v51
											v63 = v61
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
											v56 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
											v63 = v56
										}
										*(*int32)(unsafe.Add(mBase, uint32(v49))) = v51
										*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
										*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v54 + int32(1)
									}
								default:
								case 9, 11:
									if v22&int32(1024) != 0 {
									} else {
										v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										if v31 != 0 {
										} else {
											v33 = int32(1)
											v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
											if v34 == int32(0) {
												v41 = v33
											} else {
												v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
												if v37 != 0 {
													v41 = v33
												} else {
													v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+156))
													v41 = base.B2i32(v38 != int32(13))
												}
											}
											if v41 == int32(0) {
											} else {
												v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v44 | int32(4194304)
												v49 = *(*int32)(unsafe.Add(mBase, _c_F_unprotectClient[0]))
												v51 = l0 + int32(168)
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
												if v54 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
													*(*int32)(unsafe.Add(mBase, uint32(v61))) = v51
													v63 = v61
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
													v56 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56
													v63 = v56
												}
												*(*int32)(unsafe.Add(mBase, uint32(v49))) = v51
												*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v63
												*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v54 + int32(1)
											}
										}
									}
								}
							}
						}
					}
					return
				}
			}
		}
	}
}
func F_updateCachedTime(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(64)
	m.G0 = v7
	v10 = F_ustime(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_updateCachedTime[0])) = v10
	v13 = int64(1000)
	v14 = base.I64_div_s(v10, v13)
	*(*int64)(unsafe.Add(mBase, _c_F_updateCachedTime[1])) = v14
	v18 = base.I64_div_s(v10, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _c_F_updateCachedTime[2])) = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_updateCachedTime[3]))
	v25 = int32(base.Ui32(v21&int32(2)) >> (uint(int32(1)) % 32))
	*(*uint8)(unsafe.Add(mBase, _c_F_updateCachedTime[4])) = uint8(v25)
	v30 = base.I64_div_s(v14, int64(60000))
	*(*uint16)(unsafe.Add(mBase, _c_F_updateCachedTime[5])) = uint16(v30)
	v34 = base.I64_div_s(v14, v13)
	*(*int32)(unsafe.Add(mBase, _c_F_updateCachedTime[6])) = base.I32_wrap_i64(v34) & int32(16777215)
	if l0 == int32(0) {
	} else {
		v41 = int32(0)
		v42 = *(*int64)(unsafe.Add(mBase, _c_F_updateCachedTime[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v42
		v48 = F___localtime_r(m, v7+int32(8), v7+int32(20))
		mBase = m.M
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
		*(*int32)(unsafe.Add(mBase, _c_F_updateCachedTime[7])) = v50
	}
	m.G0 = v7 + int32(64)
	return
}
func F_updateGoodReplicas(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_updateGoodReplicas[0]))
	if v10 == int32(0) {
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_updateGoodReplicas[1]))
		if v14 == int32(0) {
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_updateGoodReplicas[2]))
			v20 = v7 + int32(8)
			F_listRewind(m, v18, v20)
			mBase = m.M
			v22 = int32(0)
			v25 = F_listNext(m, v20)
			mBase = m.M
			if v25 == v22 {
				v49 = v22
			} else {
				v29 = v22
				v30 = v25
				for {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					if v33 != int32(9) {
						v44 = v29
					} else {
						v36 = int32(_a_F_updateGoodReplicas_0)
						v37 = *(*int64)(unsafe.Add(mBase, _c_F_updateGoodReplicas[3]))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v32)+80))
						v41 = int64(*(*int32)(unsafe.Add(mBase, _c_F_updateGoodReplicas[1])))
						v44 = v29 + base.B2i32(v37-v38 <= v41)
					}
					v47 = F_listNext(m, v7+int32(8))
					mBase = m.M
					if v47 != 0 {
						v29 = v44
						v30 = v47
						continue
					} else {
						break
					}
					break
				}
				v49 = v44
			}
			*(*int32)(unsafe.Add(mBase, _c_F_updateGoodReplicas[4])) = v49
		}
	}
	m.G0 = v7 + int32(16)
	return int32(1)
}
func F_updateJemallocBgThread(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F_updateProcTitleTemplate(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F_updateReplicasWaitingBgsave(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int64
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[0]))
	v15 = v10 + int32(152)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16
	goto L1
L1:
	;
	v21 = v10 + int32(152)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v10 + int32(160)
	return
L3:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23+base.B2i32(v26 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
	goto L4
L6:
	;
	v41 = v23
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 != int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v182 = v10 + int32(152)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v184 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L10:
	;
	if l0 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l1 != int32(2) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+204)) = v52 & int32(-33554433)
	F_freeClientAsync(m, v45)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[1]))
	if int32(3) < v59 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F__serverLog(m, int32(3), int32(_a_F_updateReplicasWaitingBgsave_0), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	F_freeClientAsync(m, v45)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L49
	}
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[2]))
	v87 = int32(0)
	v89 = F_open(m, v86, v87, v87)
	mBase = m.M
	if v89 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[1]))
	if int32(2) < v68 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v78 = F_replicaPutOnline(m, v45)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L24
	}
L21:
	;
	v71 = F_replicationGetReplicaName(m, v45)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v71
	F__serverLog(m, int32(2), int32(_a_F_updateReplicasWaitingBgsave_1), v10)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if v78 == int32(0) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = int32(1)
	goto L9
L26:
	;
	if int32(-1) < v89 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	F_freeClientAsync(m, v45)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[1]))
	if int32(3) < v95 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L30
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[3]))
	v100 = F___strerror_l(m, v99, v99)
	mBase = m.M
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v100
	F__serverLog(m, int32(3), int32(_a_F_updateReplicasWaitingBgsave_2), v10+int32(16))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L9
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v89
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v10)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v141
	v145 = F_sdsempty(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L44
	}
L34:
	;
	if v117 != int32(-1) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v116 = F___fstatat(m, v89, int32(_a_F_updateReplicasWaitingBgsave_3), v10+int32(56), int32(4096))
	mBase = m.M
	v117 = v116
	goto L34
L36:
	;
	v113 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v117 = v113
	goto L34
L37:
	;
	F_freeClientAsync(m, v45)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[1]))
	if int32(3) < v123 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v136 = F_close(m, v89)
	mBase = m.M
	goto L9
L40:
	;
	goto L41
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_updateReplicasWaitingBgsave[3]))
	v128 = F___strerror_l(m, v127, v127)
	mBase = m.M
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v128
	F__serverLog(m, int32(3), int32(_a_F_updateReplicasWaitingBgsave_4), v10+int32(32))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v147)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v148
	v153 = F_sdscatprintf(m, v145, int32(_a_F_updateReplicasWaitingBgsave_5), v10+int32(48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v153
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v158 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+80))
	v162 = m.T0[v161].(func(*base.Module, int32, int32, int32) int32)(m, v157, v158, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+80))
	v169 = m.T0[v168].(func(*base.Module, int32, int32, int32) int32)(m, v164, int32(973), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v169 != int32(-1) {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	goto L17
L49:
	;
	goto L9
L50:
	;
	if v184 != 0 {
		v41 = v184
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v184+base.B2i32(v187 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v193
	goto L51
L53:
	;
	goto L8
}
func F_updateRequirePass(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_updateRequirePass[0]))
	F_ACLUpdateDefaultUserPassword(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_updateSighandlerEnabled(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
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
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_updateSighandlerEnabled[0]))
	if v3 == int32(0) {
		v67 = int32(0)
		v68 = m.G0
		v69 = int32(144)
		v70 = v68 - v69
		m.G0 = v70
		v72 = int32(4)
		v73 = v70 + v72
		v111 = int32(8)
		v76 = F_sigemptyset(m, v70+v111)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v67
		*(*int32)(unsafe.Add(mBase, uint32(v70)+136)) = int32(-1073741824)
		v85 = F___sigaction(m, int32(11), v73, v67)
		mBase = m.M
		v90 = F___sigaction(m, int32(7), v73, v67)
		mBase = m.M
		v95 = F___sigaction(m, v111, v73, v67)
		mBase = m.M
		v100 = F___sigaction(m, v72, v73, v67)
		mBase = m.M
		v105 = F___sigaction(m, int32(6), v73, v67)
		mBase = m.M
		m.G0 = v70 + v69
		return int32(1)
	} else {
		v7 = m.G0
		v9 = v7 - int32(144)
		m.G0 = v9
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_updateSighandlerEnabled[1]))
		if v12 != 0 {
		} else {
			v13 = int32(_a_F_updateSighandlerEnabled_0)
			v14 = F_pthread_mutexattr_init(m, v13)
			mBase = m.M
			v17 = F_pthread_mutexattr_settype(m, v13, int32(2))
			mBase = m.M
			v20 = F_pthread_mutex_init(m, int32(_a_F_updateSighandlerEnabled_1), v13)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _c_F_updateSighandlerEnabled[1])) = int32(1)
		}
		v28 = F_sigemptyset(m, v9+int32(8))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(519)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = int32(1073741828)
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_updateSighandlerEnabled[0]))
		if v34 == int32(0) {
		} else {
			v38 = int32(4)
			v39 = v9 + v38
			v40 = int32(0)
			v41 = F___sigaction(m, int32(11), v39, v40)
			mBase = m.M
			v46 = F___sigaction(m, int32(7), v39, v40)
			mBase = m.M
			v51 = F___sigaction(m, int32(8), v39, v40)
			mBase = m.M
			v56 = F___sigaction(m, v38, v39, v40)
			mBase = m.M
			v61 = F___sigaction(m, int32(6), v39, v40)
			mBase = m.M
		}
		m.G0 = v9 + int32(144)
		return int32(1)
	}
}
func F_upsertPayloadHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	v9 = int32(0)
	v15 = int32(1)
	if l3 == v15 {
		v18 = l4
	} else {
		v18 = v15
	}
	if base.Ui32(l7) < base.Ui32(v18) {
		v100 = v9
		return v100
	} else {
		v20 = int32(-1)
		v23 = int32(_a_F_upsertPayloadHeader_0)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_upsertPayloadHeader[0]))
		v25 = int32(0)
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_upsertPayloadHeader[1]))
		if base.B2i32(l5 != v20)&(base.B2i32(v24 != v25)&base.B2i32(v28 != v25)) != 0 {
			v33 = l5
		} else {
			v33 = v20
		}
		if base.Ui32(l7) < base.Ui32(l4) {
			v35 = l7
		} else {
			v35 = l4
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v36 == int32(0) {
			if base.Ui32(l7) < base.Ui32(v18+int32(12)) {
				v100 = v9
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v61 = l0 + v60
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
				v68 = v63&int32(254) | l3&int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)) = uint8(v68)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v33)
				v73 = l7 + int32(-12)
				if base.Ui32(v73) < base.Ui32(l4) {
					v75 = v73
				} else {
					v75 = v35
				}
				*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
				v77 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
				v84 = v79&int32(253) | l6<<(uint(int32(1))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)) = uint8(v84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*uint8)(unsafe.Add(mBase, uint32(v86)+11)) = uint8(v77)
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
				v91 = v89 & int32(3)
				*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)) = uint8(v91)
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93 + int32(12)
				v100 = v75
			}
			return v100
		} else {
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+10)))
			if v39&int32(1) != l3 {
				if base.Ui32(l7) < base.Ui32(v18+int32(12)) {
					v100 = v9
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v61 = l0 + v60
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
					v68 = v63&int32(254) | l3&int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)) = uint8(v68)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v33)
					v73 = l7 + int32(-12)
					if base.Ui32(v73) < base.Ui32(l4) {
						v75 = v73
					} else {
						v75 = v35
					}
					*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
					v77 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
					v84 = v79&int32(253) | l6<<(uint(int32(1))%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)) = uint8(v84)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					*(*uint8)(unsafe.Add(mBase, uint32(v86)+11)) = uint8(v77)
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
					v91 = v89 & int32(3)
					*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)) = uint8(v91)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93 + int32(12)
					v100 = v75
				}
				return v100
			} else {
				v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+8)))
				if v33 != v43 {
					if base.Ui32(l7) < base.Ui32(v18+int32(12)) {
						v100 = v9
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v61 = l0 + v60
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
						v68 = v63&int32(254) | l3&int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)) = uint8(v68)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v33)
						v73 = l7 + int32(-12)
						if base.Ui32(v73) < base.Ui32(l4) {
							v75 = v73
						} else {
							v75 = v35
						}
						*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
						v84 = v79&int32(253) | l6<<(uint(int32(1))%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)) = uint8(v84)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						*(*uint8)(unsafe.Add(mBase, uint32(v86)+11)) = uint8(v77)
						v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
						v91 = v89 & int32(3)
						*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)) = uint8(v91)
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93 + int32(12)
						v100 = v75
					}
					return v100
				} else {
					v45 = int32(1)
					if l6 != int32(base.Ui32(v39)>>(uint(v45)%32))&v45 {
						if base.Ui32(l7) < base.Ui32(v18+int32(12)) {
							v100 = v9
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v61 = l0 + v60
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
							v68 = v63&int32(254) | l3&int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)) = uint8(v68)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v33)
							v73 = l7 + int32(-12)
							if base.Ui32(v73) < base.Ui32(l4) {
								v75 = v73
							} else {
								v75 = v35
							}
							*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
							v84 = v79&int32(253) | l6<<(uint(int32(1))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)) = uint8(v84)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*uint8)(unsafe.Add(mBase, uint32(v86)+11)) = uint8(v77)
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
							v91 = v89 & int32(3)
							*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)) = uint8(v91)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93 + int32(12)
							v100 = v75
						}
						return v100
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+11)))
						if v50 != 0 {
							if base.Ui32(l7) < base.Ui32(v18+int32(12)) {
								v100 = v9
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v61 = l0 + v60
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
								v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
								v68 = v63&int32(254) | l3&int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)) = uint8(v68)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v33)
								v73 = l7 + int32(-12)
								if base.Ui32(v73) < base.Ui32(l4) {
									v75 = v73
								} else {
									v75 = v35
								}
								*(*int32)(unsafe.Add(mBase, uint32(v70))) = v75
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v77
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
								v84 = v79&int32(253) | l6<<(uint(int32(1))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)) = uint8(v84)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								*(*uint8)(unsafe.Add(mBase, uint32(v86)+11)) = uint8(v77)
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
								v91 = v89 & int32(3)
								*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)) = uint8(v91)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93 + int32(12)
								v100 = v75
							}
							return v100
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v52 + v35
							return v35
						}
					}
				}
			}
		}
	}
}
