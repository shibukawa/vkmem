package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_compareStringObjects(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_compareStringObjectsWithFlags(m, l0, l1, int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_createStringObjectFromLongLongForValue(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_createStringObjectFromLongLongWithOptions(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_createStringObject_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v34 = F_sdsnewlen(m, l0, l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
			v38 = int32(12)
			v41 = F_zmalloc_usable(m, v38, v6+v38)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(34359738368)
				v47 = v41
				m.G0 = v6 + int32(16)
				return v47
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[249]))
		if base.Ui32(int32(128)) < base.Ui32(l1+v21+int32(9)) {
			v34 = F_sdsnewlen(m, l0, l1)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
				v38 = int32(12)
				v41 = F_zmalloc_usable(m, v38, v6+v38)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v34
					*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(34359738368)
					v47 = v41
					m.G0 = v6 + int32(16)
					return v47
				}
			}
		} else {
			v30 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, l1, int32(0), int64(-1))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v47 = v30
				m.G0 = v6 + int32(16)
				return v47
			}
		}
	}
}
func F_createStringObject_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = m.G4
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(48))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v100 = int32(0)
			return v100
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v6
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if base.Ui32(int32(14)) < base.Ui32(v18) {
				v102 = m.G3
				m.Env.X__assert_fail(m, v102+int32(_a1710), v102+int32(_a1707), int32(128), v102+int32(_a1711))
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				if int32(1)<<(uint(v18)%32)&int32(8290) != 0 {
					v58 = m.G4
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, l2+int32(1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 == int32(0) {
							F_freeReplyObject(m, v11)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v100 = int32(0)
								return v100
							}
						} else {
							if l2 == int32(0) {
								v67 = v60
							} else {
								v66 = F__emscripten_memcpy_bulkmem(m, v60, l1, l2)
								mBase = m.M
								v67 = v66
							}
							v69 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v67+l2))) = uint8(v69)
							v71 = l2
							v72 = v60
							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v71
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v75 == int32(0) {
								v100 = v11
								return v100
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
								if base.Ui32(v79+int32(-9)) < base.Ui32(int32(4)) {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v11
									return v11
								} else {
									if v79 != int32(2) {
										v111 = m.G3
										m.Env.X__assert_fail(m, v111+int32(_a1712), v111+int32(_a1707), int32(158), v111+int32(_a1711))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v11
										return v11
									}
								}
							}
						}
					}
				} else {
					if v18 != int32(14) {
						v102 = m.G3
						m.Env.X__assert_fail(m, v102+int32(_a1710), v102+int32(_a1707), int32(128), v102+int32(_a1711))
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					} else {
						v29 = m.G4
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l2+int32(-3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								F_freeReplyObject(m, v11)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									v100 = int32(0)
									return v100
								}
							} else {
								v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
								*(*uint16)(unsafe.Add(mBase, uint32(v11)+32)) = uint16(v35)
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(2)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(34)))) = uint8(v41)
								v43 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v11)+35)) = uint8(v43)
								v48 = l2 + int32(-4)
								if v48 == v43 {
									v52 = v31
								} else {
									v51 = F__emscripten_memcpy_bulkmem(m, v31, l1+int32(4), v48)
									mBase = m.M
									v52 = v51
								}
								v54 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v52+v48))) = uint8(v54)
								v71 = v48
								v72 = v31
								*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v71
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v75 == int32(0) {
									v100 = v11
									return v100
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
									if base.Ui32(v79+int32(-9)) < base.Ui32(int32(4)) {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v11
										return v11
									} else {
										if v79 != int32(2) {
											v111 = m.G3
											m.Env.X__assert_fail(m, v111+int32(_a1712), v111+int32(_a1707), int32(158), v111+int32(_a1711))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
											v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v11
											return v11
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
func F_getStringObjectSdsUsedMemory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4&int32(15) != 0 {
		F__serverAssertWithInfo(m, int32(0), l0, int32(_a479), int32(_a774), int32(173))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v4&int32(240) == int32(16) {
			v53 = int32(0)
		} else {
			v12 = F_objectGetVal(m, l0)
			mBase = m.M
			v19 = v12 + int32(-1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v22 = v20 & int32(7)
			switch v22 {
			case 0:
				v23 = F_zmalloc_usable_size(m, v19)
				mBase = m.M
				v52 = v23
			case 1:
				v28 = int32(4)
				switch v22 {
				case 0:
					v52 = v28 + int32(base.Ui32(v20)>>(uint(int32(3))%32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-2)))))
					v52 = v28 + v35
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
					v52 = v28 + v39
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-5))))
					v52 = v28 + v43
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
					v48 = v47
					v52 = v28 + v48
				default:
					v48 = int32(0)
					v52 = v28 + v48
				}
			case 2:
				v28 = int32(6)
				switch v22 {
				case 0:
					v52 = v28 + int32(base.Ui32(v20)>>(uint(int32(3))%32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-2)))))
					v52 = v28 + v35
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
					v52 = v28 + v39
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-5))))
					v52 = v28 + v43
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
					v48 = v47
					v52 = v28 + v48
				default:
					v48 = int32(0)
					v52 = v28 + v48
				}
			case 3:
				v28 = int32(10)
				switch v22 {
				case 0:
					v52 = v28 + int32(base.Ui32(v20)>>(uint(int32(3))%32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-2)))))
					v52 = v28 + v35
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
					v52 = v28 + v39
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-5))))
					v52 = v28 + v43
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
					v48 = v47
					v52 = v28 + v48
				default:
					v48 = int32(0)
					v52 = v28 + v48
				}
			case 4:
				v28 = int32(18)
				switch v22 {
				case 0:
					v52 = v28 + int32(base.Ui32(v20)>>(uint(int32(3))%32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-2)))))
					v52 = v28 + v35
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
					v52 = v28 + v39
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-5))))
					v52 = v28 + v43
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
					v48 = v47
					v52 = v28 + v48
				default:
					v48 = int32(0)
					v52 = v28 + v48
				}
			default:
				v28 = int32(1)
				switch v22 {
				case 0:
					v52 = v28 + int32(base.Ui32(v20)>>(uint(int32(3))%32))
				case 1:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-2)))))
					v52 = v28 + v35
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
					v52 = v28 + v39
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-5))))
					v52 = v28 + v43
				case 4:
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
					v48 = v47
					v52 = v28 + v48
				default:
					v48 = int32(0)
					v52 = v28 + v48
				}
			}
			v53 = v52
		}
		return v53
	}
}
func F_stringConfigGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(_a188)
	}
	v6 = F_sdsnew(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_stringConfigInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v5 == int32(0) {
		v12 = F_zstrdup(m, v4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = v12
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
			return
		}
	} else {
		v8 = int32(0)
		if v4 == v8 {
			v14 = v8
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
			return
		} else {
			v12 = F_zstrdup(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = v12
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v14
				return
			}
		}
	}
}
func F_stringConfigSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v19 == int32(0) {
		v24 = v16
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, v9, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v10 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	return int32(0)
L6:
	;
	if v24 == v18 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v22 != 0 {
		v24 = v16
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(0)
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v74
	F_valkey_free(m, v18)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L30
	}
L10:
	;
	v71 = F_zstrdup(m, v24)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L29
	}
L11:
	;
	v65 = int32(2)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v66&v65 != 0 {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	if v24 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v24 != 0 {
		goto L10
	} else {
		goto L25
	}
L14:
	;
	if v18 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v33 == int32(0) {
		v56 = v32
		v57 = v33
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v57-v56&int32(255) == int32(0) {
		goto L11
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v33 != v32&int32(255) {
		v56 = v32
		v57 = v33
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v39 = v18
	v40 = v24
	goto L20
L20:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v44 == int32(0) {
		v56 = v43
		v57 = v44
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v56 = v43
	v57 = v44
	goto L17
L22:
	;
	v47 = int32(1)
	if v44 == v43&int32(255) {
		v39 = v39 + v47
		v40 = v40 + v47
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L10
L25:
	;
	v74 = int32(0)
	v75 = v17
	goto L9
L26:
	;
	v69 = int32(1)
	goto L28
L27:
	;
	v69 = v65
	goto L28
L28:
	;
	return v69
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v74 = v71
	v75 = v73
	goto L9
L30:
	;
	return int32(1)
}
func F_trimStringObjectIfNeeded(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7&int32(240) != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
		v16 = v14 & int32(7)
		switch v16 {
		case 0:
			v31 = int32(base.Ui32(v14) >> (uint(int32(3)) % 32))
		case 1:
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
			v31 = v21
		case 2:
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
			v31 = v24
		case 3:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
			v31 = v27
		case 4:
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
			v31 = v30
		default:
			v31 = int32(0)
		}
		if l1 != 0 {
			switch v16 + int32(-1) {
			case 0:
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-2)))))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
				v77 = v50 - v53
				v79 = base.I32_div_u_s(v31, int32(10))
				if base.Ui32(v77) <= base.Ui32(v79) {
					return
				} else {
					v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
						return
					}
				}
			case 1:
				v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
				v77 = v57 - v60
				v79 = base.I32_div_u_s(v31, int32(10))
				if base.Ui32(v77) <= base.Ui32(v79) {
					return
				} else {
					v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
						return
					}
				}
			case 2:
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-5))))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
				v77 = v64 - v67
				v79 = base.I32_div_u_s(v31, int32(10))
				if base.Ui32(v77) <= base.Ui32(v79) {
					return
				} else {
					v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
						return
					}
				}
			case 3:
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-9))))
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-17))))
				v77 = base.I32_wrap_i64(v71 - v74)
				v79 = base.I32_div_u_s(v31, int32(10))
				if base.Ui32(v77) <= base.Ui32(v79) {
					return
				} else {
					v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
						return
					}
				}
			default:
				return
			}
		} else {
			if base.Ui32(int32(32767)) < base.Ui32(v31) {
				switch v16 + int32(-1) {
				case 0:
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-2)))))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
					v77 = v50 - v53
					v79 = base.I32_div_u_s(v31, int32(10))
					if base.Ui32(v77) <= base.Ui32(v79) {
						return
					} else {
						v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
							return
						}
					}
				case 1:
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
					v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
					v77 = v57 - v60
					v79 = base.I32_div_u_s(v31, int32(10))
					if base.Ui32(v77) <= base.Ui32(v79) {
						return
					} else {
						v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
							return
						}
					}
				case 2:
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-5))))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
					v77 = v64 - v67
					v79 = base.I32_div_u_s(v31, int32(10))
					if base.Ui32(v77) <= base.Ui32(v79) {
						return
					} else {
						v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
							return
						}
					}
				case 3:
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-9))))
					v74 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-17))))
					v77 = base.I32_wrap_i64(v71 - v74)
					v79 = base.I32_div_u_s(v31, int32(10))
					if base.Ui32(v77) <= base.Ui32(v79) {
						return
					} else {
						v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
							return
						}
					}
				default:
					return
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[381]))
				if v35 == int32(0) {
					return
				} else {
					if base.Ui32(int32(63)) < base.Ui32(v31) {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+200))
						if v40&int32(256) == int32(0) {
							return
						} else {
							switch v16 + int32(-1) {
							case 0:
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-2)))))
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
								v77 = v50 - v53
								v79 = base.I32_div_u_s(v31, int32(10))
								if base.Ui32(v77) <= base.Ui32(v79) {
									return
								} else {
									v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
										return
									}
								}
							case 1:
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
								v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
								v77 = v57 - v60
								v79 = base.I32_div_u_s(v31, int32(10))
								if base.Ui32(v77) <= base.Ui32(v79) {
									return
								} else {
									v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
										return
									}
								}
							case 2:
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-5))))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
								v77 = v64 - v67
								v79 = base.I32_div_u_s(v31, int32(10))
								if base.Ui32(v77) <= base.Ui32(v79) {
									return
								} else {
									v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
										return
									}
								}
							case 3:
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-9))))
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(-17))))
								v77 = base.I32_wrap_i64(v71 - v74)
								v79 = base.I32_div_u_s(v31, int32(10))
								if base.Ui32(v77) <= base.Ui32(v79) {
									return
								} else {
									v82 = F_sdsRemoveFreeSpace(m, v11, int32(0))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v82
										return
									}
								}
							default:
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_tryCreateStringObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v34 = F_sdstrynewlen(m, l0, l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			if v34 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
				v39 = int32(12)
				v42 = F_zmalloc_usable(m, v39, v6+v39)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v34
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = int64(34359738368)
					v48 = v42
					m.G0 = v6 + int32(16)
					return v48
				}
			} else {
				v48 = int32(0)
				m.G0 = v6 + int32(16)
				return v48
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[249]))
		if base.Ui32(int32(128)) < base.Ui32(l1+v21+int32(9)) {
			v34 = F_sdstrynewlen(m, l0, l1)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if v34 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
					v39 = int32(12)
					v42 = F_zmalloc_usable(m, v39, v6+v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v34
						*(*int64)(unsafe.Add(mBase, uint32(v42))) = int64(34359738368)
						v48 = v42
						m.G0 = v6 + int32(16)
						return v48
					}
				} else {
					v48 = int32(0)
					m.G0 = v6 + int32(16)
					return v48
				}
			}
		} else {
			v30 = F_createEmbeddedStringObjectWithKeyAndExpire(m, l0, l1, int32(0), int64(-1))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v48 = v30
				m.G0 = v6 + int32(16)
				return v48
			}
		}
	}
}
