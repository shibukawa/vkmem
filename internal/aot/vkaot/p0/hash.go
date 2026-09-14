package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addHashIteratorCursorToReply(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v9 + int32(-2) {
	case 0:
		v32 = F_hashTypeCurrentFromHashTable(m, l1, l2, v7)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			F_addWritePreparedReplyBulkCBuffer(m, l0, v32, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	default:
		F__serverPanic_1(m, int32(_a_F_addHashIteratorCursorToReply_0), int32(1251), int32(_a_F_addHashIteratorCursorToReply_1), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
		if l2&int32(1) != 0 {
			v20 = int32(12)
		} else {
			v20 = int32(16)
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1+v20)))
		v25 = F_lpGetValue(m, v22, v7+int32(12), v7)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 == int32(0) {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
				F_addWritePreparedReplyBulkLongLong(m, l0, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				F_addWritePreparedReplyBulkCBuffer(m, l0, v25, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_freeHashObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v5)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		F_hashTypeFreeVolatileSet(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v14&int32(4) == int32(0) {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_hashtableRelease(m, v75)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					return
				}
			} else {
				if v14&int32(1) != 0 {
					v23 = int32(16)
				} else {
					v23 = int32(8)
				}
				v24 = l0 + v23
				if v14&int32(2) == int32(0) {
					v55 = v24
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					v30 = v24 + v29
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
					switch v34 & int32(7) {
					case 0:
						v51 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
					case 1:
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(-2)))))
						v51 = v41
					case 2:
						v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(-4)))))
						v51 = v44
					case 3:
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-8))))
						v51 = v47
					case 4:
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(-16))))
						v51 = v50
					default:
						v51 = int32(0)
					}
					v55 = v30 + int32(1) + v51 + int32(1)
				}
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_freeHashObject[0]))
				F_hashtableRelease(m, v55+v70)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					return
				}
			}
		}
	default:
		F__serverPanic_1(m, int32(_a_F_freeHashObject_0), int32(601), int32(_a_F_freeHashObject_1), int32(0))
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v78&int32(4) == int32(0) {
			v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_lpFree(m, v139)
			mBase = m.M
			v141 = m.ExcPending
			if v141 != 0 {
				return
			} else {
				return
			}
		} else {
			if v78&int32(1) != 0 {
				v87 = int32(16)
			} else {
				v87 = int32(8)
			}
			v88 = l0 + v87
			if v78&int32(2) == int32(0) {
				v119 = v88
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
				v94 = v88 + v93
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
				switch v98 & int32(7) {
				case 0:
					v115 = int32(base.Ui32(v98) >> (uint(int32(3)) % 32))
				case 1:
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-2)))))
					v115 = v105
				case 2:
					v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-4)))))
					v115 = v108
				case 3:
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-8))))
					v115 = v111
				case 4:
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-16))))
					v115 = v114
				default:
					v115 = int32(0)
				}
				v119 = v94 + int32(1) + v115 + int32(1)
			}
			v134 = *(*int32)(unsafe.Add(mBase, _c_F_freeHashObject[0]))
			F_lpFree(m, v119+v134)
			mBase = m.M
			v138 = m.ExcPending
			if v138 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getHashSeedFromString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = int32(32)
	v11 = v8 + v10
	F_sha256_init(m, v11)
	F_sha256_update(m, v11, l2, l3)
	F_sha256_final(m, v11, v8)
	if base.Ui32(l1) < base.Ui32(v10) {
		v22 = l1
	} else {
		v22 = v10
	}
	if v22 == int32(0) {
	} else {
		v25 = F__emscripten_memcpy_bulkmem(m, l0, v8, v22)
	}
	m.G0 = v8 + int32(144)
	return
}
func F_hashHashtableTypeDestructor(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_entryFree(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_hashHashtableTypeGetKey(m *base.Module, l0 int32) int32 {
	return l0
}
func F_hashTypeCurrentFromListpack(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(11) {
		if l1&int32(1) != 0 {
			v19 = int32(12)
		} else {
			v19 = int32(16)
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+v19)))
		v22 = F_lpGetValue(m, v21, l3, l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
			return
		}
	} else {
		F__serverAssert(m, int32(_a_F_hashTypeCurrentFromListpack_0), int32(_a_F_hashTypeCurrentFromListpack_1), int32(716))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
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
func F_hashTypeCurrentObjectNewSds(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 + int32(-2) {
	case 0:
		v33 = F_hashTypeCurrentFromHashTable(m, l0, l1, v6+int32(8))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v36 = F_sdsnewlen(m, v33, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v48 = v36
				m.G0 = v6 + int32(16)
				return v48
			}
		}
	default:
		F__serverPanic_1(m, int32(_a_F_hashTypeCurrentObjectNewSds_0), int32(755), int32(_a_F_hashTypeCurrentObjectNewSds_1), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		if l1&int32(1) != 0 {
			v15 = int32(12)
		} else {
			v15 = int32(16)
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0+v15)))
		v22 = F_lpGetValue(m, v17, v6+int32(4), v6+int32(8))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
				v46 = F_sdsfromlonglong(m, v45)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = v46
					m.G0 = v6 + int32(16)
					return v48
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				v29 = F_sdsnewlen(m, v22, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v48 = v29
					m.G0 = v6 + int32(16)
					return v48
				}
			}
		}
	}
}
func F_hashTypeExists(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(9223372036854775807)
	v17 = F_hashTypeGetValue(m, l0, l1, v6+int32(12), v6+int32(8), v6, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v17 == int32(0))
	}
}
func F_hashTypeFreeVolatileSet(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = F_objectGetVal(m, l0)
	mBase = m.M
	v5 = v3 + int32(44)
	if v5 == int32(0) {
		v19 = int32(0)
	} else {
		v9 = int32(1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		switch v10 + v9 {
		case 0:
			v19 = v9
		case 1:
			v19 = int32(0)
		default:
			if v10&int32(7) != 0 {
				v19 = v9
			} else {
				v19 = int32(0)
			}
		}
	}
	if v19 == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v24&int32(240) != int32(32) {
		} else {
			v29 = F_objectGetVal(m, l0)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(_a_F_hashTypeFreeVolatileSet_0)
		}
		return
	} else {
		F_vsetRelease(m, v5)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v24&int32(240) != int32(32) {
			} else {
				v29 = F_objectGetVal(m, l0)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(_a_F_hashTypeFreeVolatileSet_0)
			}
			return
		}
	}
}
func F_hashTypeGetExpiry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v10)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		v35 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v35
		v38 = F_objectGetVal(m, l0)
		mBase = m.M
		v39 = F_hashtableFind(m, v38, l1, v8)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if l2 == int32(0) {
			} else {
				if v39 == int32(0) {
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v49 = int64(-1)
					v51 = v45 + int32(-1)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					if v52&int32(7) == int32(0) {
						v87 = v49
					} else {
						if v52&int32(8) == int32(0) {
							v87 = v49
						} else {
							v61 = F_sdsAllocPtr(m, v45)
							mBase = m.M
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							if int32(base.Ui32(v64&int32(16))>>(uint(int32(4))%32)) != 0 {
								v69 = int32(-4)
							} else {
								v69 = int32(0)
							}
							v72 = v64 & int32(7)
							if v72 != 0 {
								v73 = v69
							} else {
								v73 = int32(0)
							}
							if int32(base.Ui32(v64&int32(8))>>(uint(int32(3))%32)) != 0 {
								v81 = int32(-8)
							} else {
								v81 = int32(0)
							}
							if v72 != 0 {
								v83 = v81
							} else {
								v83 = int32(0)
							}
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v61+v73+v83)))
							v87 = v85
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v87
				}
			}
			if v39 != 0 {
				v94 = v35
			} else {
				v94 = int32(-1)
			}
			m.G0 = v8 + int32(16)
			return v94
		}
	default:
		F__serverPanic_1(m, int32(_a_F_hashTypeGetExpiry_0), int32(276), int32(_a_F_hashTypeGetExpiry_1), int32(0))
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
		v21 = int32(0)
		v27 = F_hashTypeGetValue(m, l0, l1, v8+int32(12), v8+int32(8), v8, v21)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v27 != 0 {
				v94 = int32(-1)
			} else {
				if l2 == int32(0) {
					v94 = v21
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(-1)
					v94 = v21
				}
			}
			m.G0 = v8 + int32(16)
			return v94
		}
	}
}
func F_hashTypeGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v13)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
		v33 = F_objectGetVal(m, l0)
		mBase = m.M
		v36 = F_hashtableFind(m, v33, l1, v11+int32(12))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			if v38 != 0 {
				v40 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v40
				v43 = v11 + int32(8)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-1)))))
				v51 = v49 & int32(7)
				if v51 == v40 {
					switch v51 {
					case 0:
						v71 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
					case 1:
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-3)))))
						v71 = v61
					case 2:
						v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(-5)))))
						v71 = v64
					case 3:
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-9))))
						v71 = v67
					case 4:
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-17))))
						v71 = v70
					default:
						v71 = int32(0)
					}
					v73 = int32(1)
					v74 = F_sdsHdrSize(m, v73)
					mBase = m.M
					v75 = v38 + v71 + v74
					v77 = v75 + v73
					if v43 == int32(0) {
						v145 = v77
						v153 = v145
					} else {
						v80 = int32(0)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v80))))
						switch v83 & int32(7) {
						case 0:
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(base.Ui32(v83) >> (uint(int32(3)) % 32))
							v153 = v77
						case 1:
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-2)))))
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v91
							v153 = v77
						case 2:
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-4)))))
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v95
							v153 = v77
						case 3:
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-8))))
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v99
							v153 = v77
						case 4:
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-16))))
							v104 = v103
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v104
							v153 = v77
						default:
							v104 = v80
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v104
							v153 = v77
						}
					}
				} else {
					if v49&int32(16) != 0 {
						v106 = F_sdsAllocPtr(m, v38)
						mBase = m.M
						v108 = v106 + int32(-4)
						if v49&int32(32) == int32(0) {
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							if v43 == int32(0) {
								v145 = v120
							} else {
								v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+int32(-1)))))
								switch v126 & int32(7) {
								case 0:
									v143 = int32(base.Ui32(v126) >> (uint(int32(3)) % 32))
								case 1:
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+int32(-3)))))
									v143 = v133
								case 2:
									v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120+int32(-5)))))
									v143 = v136
								case 3:
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(-9))))
									v143 = v139
								case 4:
									v142 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(-17))))
									v143 = v142
								default:
									v143 = int32(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v143
								v145 = v120
							}
							v153 = v145
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
							if v113 != 0 {
								if v43 == int32(0) {
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v43))) = v117
								}
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
								v153 = v119
							} else {
								v153 = int32(0)
							}
						}
					} else {
						switch v51 {
						case 0:
							v71 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
						case 1:
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+int32(-3)))))
							v71 = v61
						case 2:
							v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(-5)))))
							v71 = v64
						case 3:
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-9))))
							v71 = v67
						case 4:
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(-17))))
							v71 = v70
						default:
							v71 = int32(0)
						}
						v73 = int32(1)
						v74 = F_sdsHdrSize(m, v73)
						mBase = m.M
						v75 = v38 + v71 + v74
						v77 = v75 + v73
						if v43 == int32(0) {
							v145 = v77
							v153 = v145
						} else {
							v80 = int32(0)
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v80))))
							switch v83 & int32(7) {
							case 0:
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(base.Ui32(v83) >> (uint(int32(3)) % 32))
								v153 = v77
							case 1:
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+int32(-2)))))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v91
								v153 = v77
							case 2:
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75+int32(-4)))))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v95
								v153 = v77
							case 3:
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-8))))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v99
								v153 = v77
							case 4:
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(-16))))
								v104 = v103
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v104
								v153 = v77
							default:
								v104 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v104
								v153 = v77
							}
						}
					}
				}
				if v153 == int32(0) {
					F__serverAssert(m, int32(_a_F_hashTypeGetValue_0), int32(_a_F_hashTypeGetValue_1), int32(247))
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v153
					v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v157
					if l5 == int32(0) {
						v224 = int32(0)
					} else {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						v165 = int64(-1)
						v167 = v161 + int32(-1)
						v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
						if v168&int32(7) == int32(0) {
							v203 = v165
						} else {
							if v168&int32(8) == int32(0) {
								v203 = v165
							} else {
								v177 = F_sdsAllocPtr(m, v161)
								mBase = m.M
								v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
								if int32(base.Ui32(v180&int32(16))>>(uint(int32(4))%32)) != 0 {
									v185 = int32(-4)
								} else {
									v185 = int32(0)
								}
								v188 = v180 & int32(7)
								if v188 != 0 {
									v189 = v185
								} else {
									v189 = int32(0)
								}
								if int32(base.Ui32(v180&int32(8))>>(uint(int32(3))%32)) != 0 {
									v197 = int32(-8)
								} else {
									v197 = int32(0)
								}
								if v188 != 0 {
									v199 = v197
								} else {
									v199 = int32(0)
								}
								v201 = *(*int64)(unsafe.Add(mBase, uint32(v177+v189+v199)))
								v203 = v201
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(l5))) = v203
						v224 = int32(0)
					}
					m.G0 = v11 + int32(16)
					return v224
				}
			} else {
				v224 = int32(-1)
				m.G0 = v11 + int32(16)
				return v224
			}
		}
	default:
		F__serverPanic_1(m, int32(_a_F_hashTypeGetValue_1), int32(254), int32(_a_F_hashTypeGetValue_2), int32(0))
		mBase = m.M
		v212 = m.ExcPending
		if v212 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
		v23 = F_hashTypeGetFromListpack(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v23 != 0 {
				v224 = int32(-1)
			} else {
				if l5 == int32(0) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l5))) = int64(-1)
				}
				v224 = int32(0)
			}
			m.G0 = v11 + int32(16)
			return v224
		}
	}
}
func F_hashTypeGetValueObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_hashTypeGetValue(m, l0, l1, v7+int32(12), v7+int32(8), v7, v3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(-1) {
			v30 = v3
			m.G0 = v7 + int32(16)
			return v30
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			if v21 == int32(0) {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
				v28 = F_createStringObjectFromLongLong(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = v28
					m.G0 = v7 + int32(16)
					return v30
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v25 = F_createStringObject_1(m, v21, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v30 = v25
					m.G0 = v7 + int32(16)
					return v30
				}
			}
		}
	}
}
func F_hashTypeHasStringRef(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5&int32(240) == int32(176) {
		v34 = int32(0)
		return v34
	} else {
		v10 = F_objectGetVal(m, l0)
		mBase = m.M
		v11 = F_hashtableFindRef(m, v10, l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v16 = int32(0)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
			if v20&int32(7) == v16 {
				v33 = v16
			} else {
				if v20&int32(16) == int32(0) {
					v33 = v16
				} else {
					v33 = int32(base.Ui32(v20&int32(32)) >> (uint(int32(5)) % 32))
				}
			}
			v34 = v33
			return v34
		}
	}
}
func F_hashTypeInitVolatileIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v6)
	v11 = int32(base.Ui32(v5)>>(uint(int32(4))%32)) & int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v11
	switch v11 + int32(-2) {
	case 0:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v15&int32(240) != int32(32) {
			F__serverAssert(m, int32(_a_F_hashTypeInitVolatileIterator_0), int32(_a_F_hashTypeInitVolatileIterator_1), int32(65))
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
			v20 = F_objectGetVal(m, l0)
			mBase = m.M
			v22 = v20 + int32(44)
			v23 = int32(0)
			if v22 == v23 {
				v37 = int32(0)
			} else {
				v27 = int32(1)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				switch v28 + v27 {
				case 0:
					v37 = v27
				case 1:
					v37 = int32(0)
				default:
					if v28&int32(7) != 0 {
						v37 = v27
					} else {
						v37 = int32(0)
					}
				}
			}
			if v37 != 0 {
				v38 = v22
			} else {
				v38 = v23
			}
			v40 = l1 + int32(72)
			v41 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v40)+376)) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+368)) = int64(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v40)+356)) = v43
			*(*int32)(unsafe.Add(mBase, uint32(v40)+352)) = v41
			return
		}
	default:
		F__serverPanic_1(m, int32(_a_F_hashTypeInitVolatileIterator_1), int32(652), int32(_a_F_hashTypeInitVolatileIterator_2), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		return
	}
}
func F_hashTypeLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v2)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		return v10 + v11
	default:
		F__serverPanic_1(m, int32(_a_F_hashTypeLength_0), int32(622), int32(_a_F_hashTypeLength_1), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 9:
		v23 = F_objectGetVal(m, l0)
		mBase = m.M
		v24 = F_lpLength(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return int32(base.Ui32(v24) >> (uint(int32(1)) % 32))
		}
	}
}
func F_hashTypeResetIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2 != int32(2) {
		return
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v5 != 0 {
			F_vsetResetIterator(m, l0+int32(72))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			F_hashtableCleanupIterator(m, l0+int32(24))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_hashTypeSetExpire(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v127 int32
	_ = v127
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == int32(0) {
		v173 = int32(-2)
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_hashTypeSetExpire_1), int32(_a_F_hashTypeSetExpire_2), int32(500))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L16
	} else {
		goto L66
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_hashTypeSetExpire_3), int32(_a_F_hashTypeSetExpire_2), int32(492))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L65
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v173
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeSetExpire[0]))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v40&int32(240) != int32(176) {
		v70 = v40
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v27 = int32(0)
	v28 = F_commandTimeSnapshot(m)
	mBase = m.M
	if v28 < l2 {
		v39 = v27
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+216))
	if v25 != 0 {
		v39 = int32(0)
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v30 = int32(_a_F_hashTypeSetExpire_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeSetExpire[1]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeSetExpire[2]))
	if v33 != 0 {
		v39 = v27
		goto L6
	} else {
		goto L11
	}
L11:
	;
	if v31 != 0 {
		v39 = v27
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeSetExpire[3]))
	v39 = base.B2i32(v35 == int32(0))
	goto L6
L13:
	;
	if v70&int32(240) != int32(32) {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	v50 = F_hashTypeGetFromListpack(m, l0, l1, v13+int32(12), v13+int32(8), v13)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if l3&int32(6) != 0 {
		v173 = int32(0)
		goto L3
	} else {
		goto L19
	}
L16:
	;
	return int32(0)
L17:
	;
	if int32(0) <= v50 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v173 = int32(-2)
	goto L3
L19:
	;
	if v39 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L24
	}
L21:
	;
	v61 = F_hashTypeDelete(m, l0, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v61 == int32(0) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v173 = int32(2)
	goto L3
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = v69
	goto L13
L25:
	;
	v76 = F_objectGetVal(m, l0)
	mBase = m.M
	v77 = F_hashtableFindRef(m, v76, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v77 == int32(0) {
		v173 = int32(-2)
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v85 = int64(-1)
	v87 = v81 + int32(-1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88&int32(7) == int32(0) {
		v123 = v85
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if l3 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L29:
	;
	goto L28
L30:
	;
	if v88&int32(8) == int32(0) {
		v123 = v85
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v97 = F_sdsAllocPtr(m, v81)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if int32(base.Ui32(v100&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v105 = int32(-4)
	goto L34
L33:
	;
	v105 = int32(0)
	goto L34
L34:
	;
	v108 = v100 & int32(7)
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v109 = v105
	goto L37
L36:
	;
	v109 = int32(0)
	goto L37
L37:
	;
	if int32(base.Ui32(v100&int32(8))>>(uint(int32(3))%32)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = int32(-8)
	goto L40
L39:
	;
	v117 = int32(0)
	goto L40
L40:
	;
	if v108 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v119 = v117
	goto L43
L42:
	;
	v119 = int32(0)
	goto L43
L43:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v97+v109+v119)))
	v123 = v121
	goto L29
L44:
	;
	if v39 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L45:
	;
	v127 = int32(0)
	if l3&int32(1) == v127 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if l3&int32(2) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v123 != int64(-1) {
		v173 = v127
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	if l3&int32(4) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v123 == int64(-1) {
		v173 = v127
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if l3&int32(8) == int32(0) {
		goto L44
	} else {
		goto L55
	}
L53:
	;
	if base.B2i32(l2 <= v123)|base.B2i32(v123 == int64(-1)) != 0 {
		v173 = v127
		goto L3
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v123 == int64(-1) {
		goto L44
	} else {
		goto L56
	}
L56:
	;
	if v123 <= l2 {
		v173 = v127
		goto L3
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v167 = F_entrySetExpiry(m, v81, l2)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L16
	} else {
		goto L63
	}
L59:
	;
	v159 = F_hashTypeDelete(m, l0, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	if v159 != 0 {
		v173 = int32(2)
		goto L3
	} else {
		goto L61
	}
L61:
	;
	F__serverAssert(m, int32(_a_F_hashTypeSetExpire_3), int32(_a_F_hashTypeSetExpire_2), int32(543))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v167
	F_hashTypeTrackUpdateEntry(m, l0, v81, v167, v123, l2)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v173 = int32(1)
	goto L3
L65:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashTypeTrackEntry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v3 = F_hashTypeHasVolatileFields(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(0) {
			v31 = F_hashTypeGetOrcreateVolatileSet(m, l0)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = v31
				v35 = F_vsetAddEntry(m, v33, int32(1087), l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v35 == int32(0) {
						F__serverAssert(m, int32(_a_F_hashTypeTrackEntry_0), int32(_a_F_hashTypeTrackEntry_1), int32(119))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return
					}
				}
			}
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v7&int32(240) != int32(32) {
				F__serverAssert(m, int32(_a_F_hashTypeTrackEntry_2), int32(_a_F_hashTypeTrackEntry_1), int32(65))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v12 = F_objectGetVal(m, l0)
				mBase = m.M
				v14 = v12 + int32(44)
				v15 = int32(0)
				if v14 == v15 {
					v29 = int32(0)
				} else {
					v19 = int32(1)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					switch v20 + v19 {
					case 0:
						v29 = v19
					case 1:
						v29 = int32(0)
					default:
						if v20&int32(7) != 0 {
							v29 = v19
						} else {
							v29 = int32(0)
						}
					}
				}
				if v29 != 0 {
					v30 = v14
				} else {
					v30 = v15
				}
				v33 = v30
				v35 = F_vsetAddEntry(m, v33, int32(1087), l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v35 == int32(0) {
						F__serverAssert(m, int32(_a_F_hashTypeTrackEntry_0), int32(_a_F_hashTypeTrackEntry_1), int32(119))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return
					}
				}
			}
		}
	}
}
func F_hashTypeTrackUpdateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v12 = base.B2i32(l1 != int32(0)) & base.B2i32(l3 != int64(-1))
	if v12 != 0 {
		v20 = F_hashTypeGetOrcreateVolatileSet(m, l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeTrackUpdateEntry[0]))
			if v12^int32(1)|base.B2i32(v25 == int32(0)) != 0 {
				v32 = F_vsetUpdateEntry(m, v20, int32(1087), l1, l2, l3, l4)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_0), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(142))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v36 = F_vsetIsEmpty(m, v20)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if v36 == int32(0) {
								return
							} else {
								v40 = F_objectGetVal(m, l0)
								mBase = m.M
								v42 = v40 + int32(44)
								if v42 == int32(0) {
									v56 = int32(0)
								} else {
									v46 = int32(1)
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
									switch v47 + v46 {
									case 0:
										v56 = v46
									case 1:
										v56 = int32(0)
									default:
										if v47&int32(7) != 0 {
											v56 = v46
										} else {
											v56 = int32(0)
										}
									}
								}
								if v56 == int32(0) {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v61&int32(240) != int32(32) {
									} else {
										v66 = F_objectGetVal(m, l0)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
									}
									return
								} else {
									F_vsetRelease(m, v42)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v61&int32(240) != int32(32) {
										} else {
											v66 = F_objectGetVal(m, l0)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
										}
										return
									}
								}
							}
						}
					}
				}
			} else {
				v29 = F_vsetIsEmpty(m, v20)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					if v29 != 0 {
						F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_3), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(140))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v32 = F_vsetUpdateEntry(m, v20, int32(1087), l1, l2, l3, l4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							if v32 == int32(0) {
								F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_0), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(142))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v36 = F_vsetIsEmpty(m, v20)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									if v36 == int32(0) {
										return
									} else {
										v40 = F_objectGetVal(m, l0)
										mBase = m.M
										v42 = v40 + int32(44)
										if v42 == int32(0) {
											v56 = int32(0)
										} else {
											v46 = int32(1)
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
											switch v47 + v46 {
											case 0:
												v56 = v46
											case 1:
												v56 = int32(0)
											default:
												if v47&int32(7) != 0 {
													v56 = v46
												} else {
													v56 = int32(0)
												}
											}
										}
										if v56 == int32(0) {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											if v61&int32(240) != int32(32) {
											} else {
												v66 = F_objectGetVal(m, l0)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
											}
											return
										} else {
											F_vsetRelease(m, v42)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v61&int32(240) != int32(32) {
												} else {
													v66 = F_objectGetVal(m, l0)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
												}
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
		v13 = int32(0)
		if base.B2i32(l2 != v13)&base.B2i32(l4 != int64(-1)) == v13 {
			return
		} else {
			v20 = F_hashTypeGetOrcreateVolatileSet(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeTrackUpdateEntry[0]))
				if v12^int32(1)|base.B2i32(v25 == int32(0)) != 0 {
					v32 = F_vsetUpdateEntry(m, v20, int32(1087), l1, l2, l3, l4)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if v32 == int32(0) {
							F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_0), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(142))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v36 = F_vsetIsEmpty(m, v20)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								if v36 == int32(0) {
									return
								} else {
									v40 = F_objectGetVal(m, l0)
									mBase = m.M
									v42 = v40 + int32(44)
									if v42 == int32(0) {
										v56 = int32(0)
									} else {
										v46 = int32(1)
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										switch v47 + v46 {
										case 0:
											v56 = v46
										case 1:
											v56 = int32(0)
										default:
											if v47&int32(7) != 0 {
												v56 = v46
											} else {
												v56 = int32(0)
											}
										}
									}
									if v56 == int32(0) {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										if v61&int32(240) != int32(32) {
										} else {
											v66 = F_objectGetVal(m, l0)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
										}
										return
									} else {
										F_vsetRelease(m, v42)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											if v61&int32(240) != int32(32) {
											} else {
												v66 = F_objectGetVal(m, l0)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
											}
											return
										}
									}
								}
							}
						}
					}
				} else {
					v29 = F_vsetIsEmpty(m, v20)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 != 0 {
							F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_3), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(140))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v32 = F_vsetUpdateEntry(m, v20, int32(1087), l1, l2, l3, l4)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								if v32 == int32(0) {
									F__serverAssert(m, int32(_a_F_hashTypeTrackUpdateEntry_0), int32(_a_F_hashTypeTrackUpdateEntry_1), int32(142))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v36 = F_vsetIsEmpty(m, v20)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										if v36 == int32(0) {
											return
										} else {
											v40 = F_objectGetVal(m, l0)
											mBase = m.M
											v42 = v40 + int32(44)
											if v42 == int32(0) {
												v56 = int32(0)
											} else {
												v46 = int32(1)
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
												switch v47 + v46 {
												case 0:
													v56 = v46
												case 1:
													v56 = int32(0)
												default:
													if v47&int32(7) != 0 {
														v56 = v46
													} else {
														v56 = int32(0)
													}
												}
											}
											if v56 == int32(0) {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v61&int32(240) != int32(32) {
												} else {
													v66 = F_objectGetVal(m, l0)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
												}
												return
											} else {
												F_vsetRelease(m, v42)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													if v61&int32(240) != int32(32) {
													} else {
														v66 = F_objectGetVal(m, l0)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(_a_F_hashTypeTrackUpdateEntry_2)
													}
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
func F_hashTypeTryConversion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9&int32(240) != int32(176) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L28
	}
L2:
	;
	return
L3:
	;
	v18 = base.I32_div_s(l3-l2+int32(1), int32(2))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeTryConversion[0]))
	if base.Ui32(v20) < base.Ui32(v18) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(0)
	if l3 < l2 {
		v83 = v22
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v86 = F_objectGetVal(m, l0)
	mBase = m.M
	if v86 != 0 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v26 = l2
	v29 = v22
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	switch int32(base.Ui32(v36)>>(uint(int32(4))%32)) & int32(15) {
	case 0, 8:
		goto L10
	default:
		v72 = v29
		goto L9
	}
L8:
	;
	v83 = v72
	goto L5
L9:
	;
	if v26 != l3 {
		v26 = v26 + int32(1)
		v29 = v72
		goto L7
	} else {
		goto L21
	}
L10:
	;
	v42 = F_objectGetVal(m, v35)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
	switch v45 & int32(7) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		v69 = int32(0)
		goto L11
	}
L11:
	;
	v72 = v69 + v29
	goto L9
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_hashTypeTryConversion[1]))
	if base.Ui32(v62) <= base.Ui32(v64) {
		v69 = v62
		goto L11
	} else {
		goto L18
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
	v62 = v61
	goto L12
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
	v62 = v58
	goto L12
L15:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
	v62 = v55
	goto L12
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
	v62 = v52
	goto L12
L17:
	;
	v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	return
L21:
	;
	goto L8
L22:
	;
	if base.Ui32(v89+v83) < base.Ui32(int32(1073741825)) {
		goto L2
	} else {
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v89 = v88
	goto L23
L25:
	;
	v89 = int32(0)
	goto L23
L26:
	;
	F_hashTypeConvert(m, l0, int32(2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L2
L28:
	;
	v107 = F_objectGetVal(m, l0)
	mBase = m.M
	v108 = F_hashtableExpand(m, v107, v18)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	return
}
func F_hash_pointer(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(16)
	v6 = (int32(base.Ui32(l0)>>(uint(v2)%32)) ^ l0) * int32(-2048144789)
	v11 = (int32(base.Ui32(v6)>>(uint(int32(13))%32)) ^ v6) * int32(-1028477387)
	return base.I64_extend_i32_u(int32(base.Ui32(v11)>>(uint(v2)%32)) ^ v11)
}
