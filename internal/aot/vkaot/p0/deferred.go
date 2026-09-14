package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_deferredAfterErrorReply(m *base.Module, l0 int32, l1 int32) {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = v8 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v17 = v8 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return
L3:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+base.B2i32(v22 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v28
	goto L4
L6:
	;
	v35 = v19
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))))
	switch v41 & int32(7) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v58 = int32(0)
		goto L9
	}
L8:
	;
	goto L2
L9:
	;
	F_afterErrorReply(m, l0, v38, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-17))))
	v58 = v57
	goto L9
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-9))))
	v58 = v54
	goto L9
L12:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(-5)))))
	v58 = v51
	goto L9
L13:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-3)))))
	v58 = v48
	goto L9
L14:
	;
	v58 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	return
L16:
	;
	v63 = v8 + int32(8)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v65 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v65 != 0 {
		v35 = v65
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65+base.B2i32(v68 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v74
	goto L18
L20:
	;
	goto L8
}
func F_setDeferredAggregateLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	v4 = l3
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	if l2 <= int32(-1) {
		F__serverAssert(m, int32(_a789), int32(_a774), int32(1246))
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
		if l1 == int32(0) {
			m.G0 = v10 + int32(128)
			return
		} else {
			if base.Ui32(l2) < base.Ui32(int32(10)) {
				v20 = int32(4)
			} else {
				v20 = int32(5)
			}
			v22 = base.B2i32(base.Ui32(int32(31)) < base.Ui32(l2))
			if base.Ui32(int32(31)) < base.Ui32(l2) {
				if base.Ui32(int32(31)) < base.Ui32(l2) {
					if base.Ui32(int32(31)) < base.Ui32(l2) {
						*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
						v57 = v10 | int32(1)
						v59 = base.I64_extend_i32_u(l2)
						if v59 <= int64(-1) {
							v68 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
							v72 = int32(1)
							v77 = v57 + v72
							v78 = int32(126)
							v79 = int64(0) - v59
							v80 = v72
						} else {
							v77 = v57
							v78 = int32(127)
							v79 = v59
							v80 = int32(0)
						}
						v81 = F_ull2string(m, v77, v78, v79)
						mBase = m.M
						if v81 == int32(0) {
							v100 = int32(0)
						} else {
							v100 = v81 + v80
						}
						v104 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
						F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					} else {
						if v4 != int32(126) {
							*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
							v57 = v10 | int32(1)
							v59 = base.I64_extend_i32_u(l2)
							if v59 <= int64(-1) {
								v68 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
								v72 = int32(1)
								v77 = v57 + v72
								v78 = int32(126)
								v79 = int64(0) - v59
								v80 = v72
							} else {
								v77 = v57
								v78 = int32(127)
								v79 = v59
								v80 = int32(0)
							}
							v81 = F_ull2string(m, v77, v78, v79)
							mBase = m.M
							if v81 == int32(0) {
								v100 = int32(0)
							} else {
								v100 = v81 + v80
							}
							v104 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
							F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[422])))
							v52 = F_objectGetVal(m, v51)
							mBase = m.M
							F_setDeferredReply(m, l0, l1, v52, v20)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						}
					}
				} else {
					if v4 != int32(37) {
						if base.Ui32(int32(31)) < base.Ui32(l2) {
							*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
							v57 = v10 | int32(1)
							v59 = base.I64_extend_i32_u(l2)
							if v59 <= int64(-1) {
								v68 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
								v72 = int32(1)
								v77 = v57 + v72
								v78 = int32(126)
								v79 = int64(0) - v59
								v80 = v72
							} else {
								v77 = v57
								v78 = int32(127)
								v79 = v59
								v80 = int32(0)
							}
							v81 = F_ull2string(m, v77, v78, v79)
							mBase = m.M
							if v81 == int32(0) {
								v100 = int32(0)
							} else {
								v100 = v81 + v80
							}
							v104 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
							F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						} else {
							if v4 != int32(126) {
								*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
								v57 = v10 | int32(1)
								v59 = base.I64_extend_i32_u(l2)
								if v59 <= int64(-1) {
									v68 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
									v72 = int32(1)
									v77 = v57 + v72
									v78 = int32(126)
									v79 = int64(0) - v59
									v80 = v72
								} else {
									v77 = v57
									v78 = int32(127)
									v79 = v59
									v80 = int32(0)
								}
								v81 = F_ull2string(m, v77, v78, v79)
								mBase = m.M
								if v81 == int32(0) {
									v100 = int32(0)
								} else {
									v100 = v81 + v80
								}
								v104 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
								F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									m.G0 = v10 + int32(128)
									return
								}
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[422])))
								v52 = F_objectGetVal(m, v51)
								mBase = m.M
								F_setDeferredReply(m, l0, l1, v52, v20)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									m.G0 = v10 + int32(128)
									return
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[423])))
						v40 = F_objectGetVal(m, v39)
						mBase = m.M
						F_setDeferredReply(m, l0, l1, v40, v20)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v10 + int32(128)
							return
						}
					}
				}
			} else {
				if v4 != int32(42) {
					if base.Ui32(int32(31)) < base.Ui32(l2) {
						if base.Ui32(int32(31)) < base.Ui32(l2) {
							*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
							v57 = v10 | int32(1)
							v59 = base.I64_extend_i32_u(l2)
							if v59 <= int64(-1) {
								v68 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
								v72 = int32(1)
								v77 = v57 + v72
								v78 = int32(126)
								v79 = int64(0) - v59
								v80 = v72
							} else {
								v77 = v57
								v78 = int32(127)
								v79 = v59
								v80 = int32(0)
							}
							v81 = F_ull2string(m, v77, v78, v79)
							mBase = m.M
							if v81 == int32(0) {
								v100 = int32(0)
							} else {
								v100 = v81 + v80
							}
							v104 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
							F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						} else {
							if v4 != int32(126) {
								*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
								v57 = v10 | int32(1)
								v59 = base.I64_extend_i32_u(l2)
								if v59 <= int64(-1) {
									v68 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
									v72 = int32(1)
									v77 = v57 + v72
									v78 = int32(126)
									v79 = int64(0) - v59
									v80 = v72
								} else {
									v77 = v57
									v78 = int32(127)
									v79 = v59
									v80 = int32(0)
								}
								v81 = F_ull2string(m, v77, v78, v79)
								mBase = m.M
								if v81 == int32(0) {
									v100 = int32(0)
								} else {
									v100 = v81 + v80
								}
								v104 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
								F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									m.G0 = v10 + int32(128)
									return
								}
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[422])))
								v52 = F_objectGetVal(m, v51)
								mBase = m.M
								F_setDeferredReply(m, l0, l1, v52, v20)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									m.G0 = v10 + int32(128)
									return
								}
							}
						}
					} else {
						if v4 != int32(37) {
							if base.Ui32(int32(31)) < base.Ui32(l2) {
								*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
								v57 = v10 | int32(1)
								v59 = base.I64_extend_i32_u(l2)
								if v59 <= int64(-1) {
									v68 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
									v72 = int32(1)
									v77 = v57 + v72
									v78 = int32(126)
									v79 = int64(0) - v59
									v80 = v72
								} else {
									v77 = v57
									v78 = int32(127)
									v79 = v59
									v80 = int32(0)
								}
								v81 = F_ull2string(m, v77, v78, v79)
								mBase = m.M
								if v81 == int32(0) {
									v100 = int32(0)
								} else {
									v100 = v81 + v80
								}
								v104 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
								F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									m.G0 = v10 + int32(128)
									return
								}
							} else {
								if v4 != int32(126) {
									*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v4)
									v57 = v10 | int32(1)
									v59 = base.I64_extend_i32_u(l2)
									if v59 <= int64(-1) {
										v68 = int32(45)
										*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v68)
										v72 = int32(1)
										v77 = v57 + v72
										v78 = int32(126)
										v79 = int64(0) - v59
										v80 = v72
									} else {
										v77 = v57
										v78 = int32(127)
										v79 = v59
										v80 = int32(0)
									}
									v81 = F_ull2string(m, v77, v78, v79)
									mBase = m.M
									if v81 == int32(0) {
										v100 = int32(0)
									} else {
										v100 = v81 + v80
									}
									v104 = int32(2573)
									*(*uint16)(unsafe.Add(mBase, uint32(v100+v10+int32(1)))) = uint16(v104)
									F_setDeferredReply(m, l0, l1, v10, v100+int32(3))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v10 + int32(128)
										return
									}
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[422])))
									v52 = F_objectGetVal(m, v51)
									mBase = m.M
									F_setDeferredReply(m, l0, l1, v52, v20)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										m.G0 = v10 + int32(128)
										return
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[423])))
							v40 = F_objectGetVal(m, v39)
							mBase = m.M
							F_setDeferredReply(m, l0, l1, v40, v20)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								m.G0 = v10 + int32(128)
								return
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[424])))
					v30 = F_objectGetVal(m, v29)
					mBase = m.M
					F_setDeferredReply(m, l0, l1, v30, v20)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v10 + int32(128)
						return
					}
				}
			}
		}
	}
}
func F_setDeferredReplyBulkSds(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_sdsempty(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(0)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
		switch v19 & int32(7) {
		case 0:
			v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
		case 1:
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
			v36 = v26
		case 2:
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
			v36 = v29
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
			v36 = v32
		case 4:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
			v36 = v35
		default:
			v36 = v15
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v36
		v40 = F_sdscatprintf(m, v13, int32(_a796), v11)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-1)))))
			switch v44 & int32(7) {
			case 0:
				v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
			case 1:
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+int32(-3)))))
				v61 = v51
			case 2:
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40+int32(-5)))))
				v61 = v54
			case 3:
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-9))))
				v61 = v57
			case 4:
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(-17))))
				v61 = v60
			default:
				v61 = v15
			}
			F_setDeferredReply(m, l0, l1, v40, v61)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_sdsfree(m, v40)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_sdsfree(m, l2)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
func F_setDeferredReplyStreamID(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int64
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v16 = int32(1)
	if base.Ui64(v11) < base.Ui64(int64(10)) {
		v73 = v16
		v74 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v148 = v8 + v147
	v149 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v149)
	v151 = int32(1)
	v152 = v148 + v151
	v154 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v155 = int32(0)
	if base.Ui64(v154) < base.Ui64(int64(10)) {
		v216 = v151
		v217 = v155
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v77 = v73 + v74
	if base.Ui32(int32(21)) <= base.Ui32(v77) {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	v24 = v4
	v25 = v11
	goto L4
L4:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v25) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v73 = v16
	v74 = v65
	goto L2
L6:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v25) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = int32(2)
	v74 = v24
	goto L2
L8:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v25) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v73 = int32(3)
	v74 = v24
	goto L2
L10:
	;
	v65 = v24 + int32(12)
	v69 = base.I64_div_u_s(v25, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v25) {
		v24 = v65
		v25 = v69
		goto L4
	} else {
		goto L32
	}
L11:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v25) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v25) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v25) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v25) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v25) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v25) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v73 = int32(4)
	v74 = v24
	goto L2
L18:
	;
	v46 = int32(6)
	goto L20
L19:
	;
	v46 = int32(5)
	goto L20
L20:
	;
	v73 = v46
	v74 = v24
	goto L2
L21:
	;
	v51 = int32(8)
	goto L23
L22:
	;
	v51 = int32(7)
	goto L23
L23:
	;
	v73 = v51
	v74 = v24
	goto L2
L24:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v25) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v25) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v58 = int32(10)
	goto L28
L27:
	;
	v58 = int32(9)
	goto L28
L28:
	;
	v73 = v58
	v74 = v24
	goto L2
L29:
	;
	v63 = int32(12)
	goto L31
L30:
	;
	v63 = int32(11)
	goto L31
L31:
	;
	v73 = v63
	v74 = v24
	goto L2
L32:
	;
	goto L5
L33:
	;
	goto L44
L34:
	;
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8+v77))) = uint8(v80)
	v83 = v77 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v11) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v119 = v8 + v116
	if base.Ui64(int64(9)) < base.Ui64(v117) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v90 = v11
	v92 = v83
	goto L38
L37:
	;
	v116 = v83
	v117 = v11
	goto L35
L38:
	;
	v96 = int64(100)
	v97 = base.I64_div_u_s(v90, v96)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v90-v97*v96)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-1)+v92))) = uint16(v106)
	v109 = v92 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v90) {
		v90 = v97
		v92 = v109
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v116 = v109
	v117 = v97
	goto L35
L40:
	;
	goto L39
L41:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v117)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v119+int32(-1)))) = uint16(v133)
	v147 = v77
	goto L1
L42:
	;
	v124 = base.I32_wrap_i64(v117) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v124)
	v147 = v77
	goto L1
L43:
	;
	v147 = int32(0)
	goto L1
L44:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v137)
	goto L43
L45:
	;
	v293 = F_sdsnewlen(m, v8, v152+v290-v8)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v220 = v216 + v217
	if base.Ui32(int32(21)) <= base.Ui32(v220) {
		goto L77
	} else {
		goto L78
	}
L47:
	;
	v167 = v155
	v168 = v154
	goto L48
L48:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v168) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v216 = v151
	v217 = v208
	goto L46
L50:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v168) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v216 = int32(2)
	v217 = v167
	goto L46
L52:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v168) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v216 = int32(3)
	v217 = v167
	goto L46
L54:
	;
	v208 = v167 + int32(12)
	v212 = base.I64_div_u_s(v168, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v168) {
		v167 = v208
		v168 = v212
		goto L48
	} else {
		goto L76
	}
L55:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v168) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v168) {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v168) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v168) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v168) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v168) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v216 = int32(4)
	v217 = v167
	goto L46
L62:
	;
	v189 = int32(6)
	goto L64
L63:
	;
	v189 = int32(5)
	goto L64
L64:
	;
	v216 = v189
	v217 = v167
	goto L46
L65:
	;
	v194 = int32(8)
	goto L67
L66:
	;
	v194 = int32(7)
	goto L67
L67:
	;
	v216 = v194
	v217 = v167
	goto L46
L68:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v168) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v168) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v201 = int32(10)
	goto L72
L71:
	;
	v201 = int32(9)
	goto L72
L72:
	;
	v216 = v201
	v217 = v167
	goto L46
L73:
	;
	v206 = int32(12)
	goto L75
L74:
	;
	v206 = int32(11)
	goto L75
L75:
	;
	v216 = v206
	v217 = v167
	goto L46
L76:
	;
	goto L49
L77:
	;
	goto L88
L78:
	;
	v223 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v152+v220))) = uint8(v223)
	v226 = v220 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v154) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v262 = v152 + v259
	if base.Ui64(int64(9)) < base.Ui64(v260) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v233 = v154
	v235 = v226
	goto L82
L81:
	;
	v259 = v226
	v260 = v154
	goto L79
L82:
	;
	v239 = int64(100)
	v240 = base.I64_div_u_s(v233, v239)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v233-v240*v239)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v148+int32(0)+v235))) = uint16(v249)
	v252 = v235 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v233) {
		v233 = v240
		v235 = v252
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v259 = v252
	v260 = v240
	goto L79
L84:
	;
	goto L83
L85:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v260)<<(uint(int32(1))%32))+uint32(_consts[185]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v262+int32(-1)))) = uint16(v276)
	v290 = v220
	goto L45
L86:
	;
	v267 = base.I32_wrap_i64(v260) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v267)
	v290 = v220
	goto L45
L87:
	;
	v290 = int32(0)
	goto L45
L88:
	;
	v280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v152))) = uint8(v280)
	goto L87
L89:
	;
	return
L90:
	;
	F_setDeferredReplyBulkSds(m, l0, l1, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v8 + int32(48)
	return
}
