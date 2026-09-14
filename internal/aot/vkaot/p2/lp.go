package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__lpEntryValidation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v10 != 0 {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v24 == int32(0) {
			v33 = F_lpGet(m, l0, v8+int32(24), v8)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
				v36 = F_sdsnewlen(m, v33, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v39 = F_hashtableAdd(m, v38, v36)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v46 = v44
							v47 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
							v51 = v47
							m.G0 = v8 + int32(32)
							return v51
						} else {
							F_sdsfree(m, v36)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v51 = int32(0)
								m.G0 = v8 + int32(32)
								return v51
							}
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v27&int32(1) != 0 {
				v46 = v27
				v47 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
				v51 = v47
				m.G0 = v8 + int32(32)
				return v51
			} else {
				v33 = F_lpGet(m, l0, v8+int32(24), v8)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
					v36 = F_sdsnewlen(m, v33, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v39 = F_hashtableAdd(m, v38, v36)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v39 != 0 {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								v46 = v44
								v47 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
								v51 = v47
								m.G0 = v8 + int32(32)
								return v51
							} else {
								F_sdsfree(m, v36)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v51 = int32(0)
									m.G0 = v8 + int32(32)
									return v51
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = F_hashtableCreate(m, int32(_a1130))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v12
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v21 = F_hashtableExpand(m, v12, int32(base.Ui32(l1)>>(uint(base.B2i32(v17 != int32(0)))%32)))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				if v24 == int32(0) {
					v33 = F_lpGet(m, l0, v8+int32(24), v8)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
						v36 = F_sdsnewlen(m, v33, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v39 = F_hashtableAdd(m, v38, v36)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								if v39 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v46 = v44
									v47 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
									v51 = v47
									m.G0 = v8 + int32(32)
									return v51
								} else {
									F_sdsfree(m, v36)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v51 = int32(0)
										m.G0 = v8 + int32(32)
										return v51
									}
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v27&int32(1) != 0 {
						v46 = v27
						v47 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
						v51 = v47
						m.G0 = v8 + int32(32)
						return v51
					} else {
						v33 = F_lpGet(m, l0, v8+int32(24), v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
							v36 = F_sdsnewlen(m, v33, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								v39 = F_hashtableAdd(m, v38, v36)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									if v39 != 0 {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										v46 = v44
										v47 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46 + v47
										v51 = v47
										m.G0 = v8 + int32(32)
										return v51
									} else {
										F_sdsfree(m, v36)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											v51 = int32(0)
											m.G0 = v8 + int32(32)
											return v51
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
func F_lpAppendInteger(m *base.Module, l0 int32, l1 int64) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(0)
	v9 = F_lpInsertInteger(m, l0, l1, l0+v3+int32(-1), v7, v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_lpDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(0)
	v8 = F_lpInsert(m, l0, v4, v4, v4, l1, int32(2), l2)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_lpDup(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = F_zmalloc_usable(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			v12 = v5
		} else {
			v11 = F__emscripten_memcpy_bulkmem(m, v5, l0, v3)
			mBase = m.M
			v12 = v11
		}
		return v12
	}
}
func F_lpEstimateBytesRepeatedInteger(m *base.Module, l0 int64, l1 int32) int32 {
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if base.Ui64(int64(128)) <= base.Ui64(l0) {
		if base.Ui64(int64(8192)) <= base.Ui64(l0+int64(4096)) {
			if base.Ui64(int64(65536)) <= base.Ui64(l0+int64(32768)) {
				if base.Ui64(int64(16777216)) <= base.Ui64(l0+int64(8388608)) {
					if base.Ui64(l0+int64(2147483648)) < base.Ui64(int64(4294967296)) {
						v28 = int32(6)
					} else {
						v28 = int32(10)
					}
					v29 = v28
				} else {
					v29 = int32(5)
				}
			} else {
				v29 = int32(4)
			}
		} else {
			v29 = int32(3)
		}
	} else {
		v29 = int32(2)
	}
	return v29*l1 + int32(7)
}
func F_lpFirst(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v3 == int32(255) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v9 == int32(7) {
			v20 = int32(0)
			return v20
		} else {
			F__serverAssert(m, int32(_a888), int32(_a889), int32(426))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v20 = l0 + int32(6)
		return v20
	}
}
func F_lpGet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_lpGetWithSize(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_lpGetWithSize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v29 int64
	_ = v29
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v118 int64
	_ = v118
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v157 int64
	_ = v157
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a890), int32(_a889), int32(504))
		mBase = m.M
		v214 = m.ExcPending
		if v214 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v13 = base.I64_extend_i32_u(v12)
		v14 = base.I32_extend8_s(v12)
		if v14 < int32(0) {
			if v12&int32(192) != int32(128) {
				if v12&int32(224) != int32(192) {
					switch v12 + int32(-241) {
					case 0:
						v58 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+1)))
						v59 = int64(-65535)
						v60 = int64(32768)
						if l3 == int32(0) {
							v150 = v58
							v151 = v59
							v152 = v60
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(4)
							v150 = v58
							v151 = v59
							v152 = v60
						}
						if base.Ui64(v150) < base.Ui64(v152) {
							v157 = v150
						} else {
							v157 = v150 + v151 + int64(-1)
						}
						if l2 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
							return int32(0)
						} else {
							if v157 <= int64(-1) {
								v169 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
								v173 = int32(1)
								v178 = l2 + v173
								v179 = int32(20)
								v180 = int64(0) - v157
								v181 = v173
							} else {
								v178 = l2
								v179 = int32(21)
								v180 = v157
								v181 = int32(0)
							}
							v182 = F_ull2string(m, v178, v179, v180)
							mBase = m.M
							if v182 == int32(0) {
								v201 = int32(0)
							} else {
								v201 = v182 + v181
							}
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
							return l2
						}
					case 1:
						v65 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+1)))
						v66 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
						v69 = v65 | v66<<(uint(int64(16))%64)
						v70 = int64(-16777215)
						v71 = int64(8388608)
						if l3 == int32(0) {
							v150 = v69
							v151 = v70
							v152 = v71
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(5)
							v150 = v69
							v151 = v70
							v152 = v71
						}
						if base.Ui64(v150) < base.Ui64(v152) {
							v157 = v150
						} else {
							v157 = v150 + v151 + int64(-1)
						}
						if l2 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
							return int32(0)
						} else {
							if v157 <= int64(-1) {
								v169 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
								v173 = int32(1)
								v178 = l2 + v173
								v179 = int32(20)
								v180 = int64(0) - v157
								v181 = v173
							} else {
								v178 = l2
								v179 = int32(21)
								v180 = v157
								v181 = int32(0)
							}
							v182 = F_ull2string(m, v178, v179, v180)
							mBase = m.M
							if v182 == int32(0) {
								v201 = int32(0)
							} else {
								v201 = v182 + v181
							}
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
							return l2
						}
					case 2:
						v76 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+1)))
						v77 = int64(-4294967295)
						v78 = int64(2147483648)
						if l3 == int32(0) {
							v150 = v76
							v151 = v77
							v152 = v78
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(6)
							v150 = v76
							v151 = v77
							v152 = v78
						}
						if base.Ui64(v150) < base.Ui64(v152) {
							v157 = v150
						} else {
							v157 = v150 + v151 + int64(-1)
						}
						if l2 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
							return int32(0)
						} else {
							if v157 <= int64(-1) {
								v169 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
								v173 = int32(1)
								v178 = l2 + v173
								v179 = int32(20)
								v180 = int64(0) - v157
								v181 = v173
							} else {
								v178 = l2
								v179 = int32(21)
								v180 = v157
								v181 = int32(0)
							}
							v182 = F_ull2string(m, v178, v179, v180)
							mBase = m.M
							if v182 == int32(0) {
								v201 = int32(0)
							} else {
								v201 = v182 + v181
							}
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
							return l2
						}
					case 3:
						v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1))
						v84 = int64(1)
						v85 = int64(-9223372036854775807 - 1)
						if l3 == int32(0) {
							v150 = v83
							v151 = v84
							v152 = v85
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(10)
							v150 = v83
							v151 = v84
							v152 = v85
						}
						if base.Ui64(v150) < base.Ui64(v152) {
							v157 = v150
						} else {
							v157 = v150 + v151 + int64(-1)
						}
						if l2 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
							return int32(0)
						} else {
							if v157 <= int64(-1) {
								v169 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
								v173 = int32(1)
								v178 = l2 + v173
								v179 = int32(20)
								v180 = int64(0) - v157
								v181 = v173
							} else {
								v178 = l2
								v179 = int32(21)
								v180 = v157
								v181 = int32(0)
							}
							v182 = F_ull2string(m, v178, v179, v180)
							mBase = m.M
							if v182 == int32(0) {
								v201 = int32(0)
							} else {
								v201 = v182 + v181
							}
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
							return l2
						}
					default:
						if v12&int32(240) != int32(224) {
							if v14 != int32(-16) {
								v150 = v13 | int64(12345678900000000)
								v151 = int64(0)
								v152 = int64(-1)
								if base.Ui64(v150) < base.Ui64(v152) {
									v157 = v150
								} else {
									v157 = v150 + v151 + int64(-1)
								}
								if l2 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
									return int32(0)
								} else {
									if v157 <= int64(-1) {
										v169 = int32(45)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
										v173 = int32(1)
										v178 = l2 + v173
										v179 = int32(20)
										v180 = int64(0) - v157
										v181 = v173
									} else {
										v178 = l2
										v179 = int32(21)
										v180 = v157
										v181 = int32(0)
									}
									v182 = F_ull2string(m, v178, v179, v180)
									mBase = m.M
									if v182 == int32(0) {
										v201 = int32(0)
									} else {
										v201 = v182 + v181
									}
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
									return l2
								}
							} else {
								v118 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+1)))
								*(*int64)(unsafe.Add(mBase, uint32(l1))) = v118
								if l3 == int32(0) {
								} else {
									if base.Ui64(int64(123)) <= base.Ui64(v118) {
										if base.Ui64(int64(16379)) <= base.Ui64(v118) {
											if base.Ui64(int64(2097147)) <= base.Ui64(v118) {
												if base.Ui64(v118) < base.Ui64(int64(268435451)) {
													v137 = int64(4)
												} else {
													v137 = int64(5)
												}
												v138 = v137
											} else {
												v138 = int64(3)
											}
										} else {
											v138 = int64(2)
										}
									} else {
										v138 = int64(1)
									}
									*(*int64)(unsafe.Add(mBase, uint32(l3))) = v118 + int64(5) + v138
								}
								return l0 + int32(5)
							}
						} else {
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v99 = v12<<(uint(int32(8))%32)&int32(3840) | v98
							v100 = base.I64_extend_i32_u(v99)
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v100
							if l3 == int32(0) {
							} else {
								if base.Ui32(v99) < base.Ui32(int32(126)) {
									v108 = int64(1)
								} else {
									v108 = int64(2)
								}
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v100 + v108 + int64(2)
							}
							return l0 + int32(2)
						}
					}
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					v49 = base.I64_extend_i32_u(v12<<(uint(int32(8))%32)&int32(7936) | v47)
					v50 = int64(-8191)
					v51 = int64(4096)
					if l3 == int32(0) {
						v150 = v49
						v151 = v50
						v152 = v51
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(3)
						v150 = v49
						v151 = v50
						v152 = v51
					}
					if base.Ui64(v150) < base.Ui64(v152) {
						v157 = v150
					} else {
						v157 = v150 + v151 + int64(-1)
					}
					if l2 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
						return int32(0)
					} else {
						if v157 <= int64(-1) {
							v169 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
							v173 = int32(1)
							v178 = l2 + v173
							v179 = int32(20)
							v180 = int64(0) - v157
							v181 = v173
						} else {
							v178 = l2
							v179 = int32(21)
							v180 = v157
							v181 = int32(0)
						}
						v182 = F_ull2string(m, v178, v179, v180)
						mBase = m.M
						if v182 == int32(0) {
							v201 = int32(0)
						} else {
							v201 = v182 + v181
						}
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
						return l2
					}
				}
			} else {
				v29 = base.I64_extend_i32_u(v12 & int32(63))
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v29
				if l3 == int32(0) {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v29 + int64(2)
				}
				return l0 + int32(1)
			}
		} else {
			v17 = int64(0)
			v18 = int64(-1)
			if l3 == int32(0) {
				v150 = v13
				v151 = v17
				v152 = v18
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(2)
				v150 = v13
				v151 = v17
				v152 = v18
			}
			if base.Ui64(v150) < base.Ui64(v152) {
				v157 = v150
			} else {
				v157 = v150 + v151 + int64(-1)
			}
			if l2 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v157
				return int32(0)
			} else {
				if v157 <= int64(-1) {
					v169 = int32(45)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v169)
					v173 = int32(1)
					v178 = l2 + v173
					v179 = int32(20)
					v180 = int64(0) - v157
					v181 = v173
				} else {
					v178 = l2
					v179 = int32(21)
					v180 = v157
					v181 = int32(0)
				}
				v182 = F_ull2string(m, v178, v179, v180)
				mBase = m.M
				if v182 == int32(0) {
					v201 = int32(0)
				} else {
					v201 = v182 + v181
				}
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = base.I64_extend_i32_s(v201)
				return l2
			}
		}
	}
}
func F_lpInsertString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_lpInsert(m, l0, l1, int32(0), l2, l3, l4, l5)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_lpLast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v70 int64
	_ = v70
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int32(7) {
		v94 = int32(0)
	} else {
		v10 = int32(-1)
		v11 = l0 + v7
		v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11+int32(-2)))))
		v19 = base.I64_extend_i32_u(v16 & int32(127))
		if v10 < v16 {
			v87 = v10
			v89 = v19
		} else {
			v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
			v30 = base.I64_extend_i32_u(v24&int32(127))<<(uint(int64(7))%64) | v19
			if int32(-1) < v24 {
				v70 = v30
				if base.Ui64(int64(128)) <= base.Ui64(v70) {
					if base.Ui64(int64(16384)) <= base.Ui64(v70) {
						if base.Ui64(int64(2097152)) <= base.Ui64(v70) {
							if base.Ui64(v70) < base.Ui64(int64(268435456)) {
								v84 = int32(-4)
							} else {
								v84 = int32(-5)
							}
							v87 = v84
							v89 = v70
						} else {
							v87 = int32(-3)
							v89 = v70
						}
					} else {
						v87 = int32(-2)
						v89 = v70
					}
				} else {
					v87 = int32(-1)
					v89 = v70
				}
			} else {
				v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11+int32(-4)))))
				v41 = base.I64_extend_i32_u(v35&int32(127))<<(uint(int64(14))%64) | v30
				if int32(-1) < v35 {
					v70 = v41
					if base.Ui64(int64(128)) <= base.Ui64(v70) {
						if base.Ui64(int64(16384)) <= base.Ui64(v70) {
							if base.Ui64(int64(2097152)) <= base.Ui64(v70) {
								if base.Ui64(v70) < base.Ui64(int64(268435456)) {
									v84 = int32(-4)
								} else {
									v84 = int32(-5)
								}
								v87 = v84
								v89 = v70
							} else {
								v87 = int32(-3)
								v89 = v70
							}
						} else {
							v87 = int32(-2)
							v89 = v70
						}
					} else {
						v87 = int32(-1)
						v89 = v70
					}
				} else {
					v44 = int32(-5)
					v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11+v44))))
					v53 = base.I64_extend_i32_u(v47&int32(127))<<(uint(int64(21))%64) | v41
					if int32(-1) < v47 {
						v70 = v53
						if base.Ui64(int64(128)) <= base.Ui64(v70) {
							if base.Ui64(int64(16384)) <= base.Ui64(v70) {
								if base.Ui64(int64(2097152)) <= base.Ui64(v70) {
									if base.Ui64(v70) < base.Ui64(int64(268435456)) {
										v84 = int32(-4)
									} else {
										v84 = int32(-5)
									}
									v87 = v84
									v89 = v70
								} else {
									v87 = int32(-3)
									v89 = v70
								}
							} else {
								v87 = int32(-2)
								v89 = v70
							}
						} else {
							v87 = int32(-1)
							v89 = v70
						}
					} else {
						v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11+int32(-6)))))
						if v58 < int32(0) {
							v87 = v44
							v89 = int64(-1)
						} else {
							v70 = base.I64_extend_i32_u(v58&int32(127))<<(uint(int64(28))%64) | v53
							if base.Ui64(int64(128)) <= base.Ui64(v70) {
								if base.Ui64(int64(16384)) <= base.Ui64(v70) {
									if base.Ui64(int64(2097152)) <= base.Ui64(v70) {
										if base.Ui64(v70) < base.Ui64(int64(268435456)) {
											v84 = int32(-4)
										} else {
											v84 = int32(-5)
										}
										v87 = v84
										v89 = v70
									} else {
										v87 = int32(-3)
										v89 = v70
									}
								} else {
									v87 = int32(-2)
									v89 = v70
								}
							} else {
								v87 = int32(-1)
								v89 = v70
							}
						}
					}
				}
			}
		}
		v94 = v11 + v10 + (v87 - base.I32_wrap_i64(v89))
	}
	return v94
}
func F_lpNew(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v2 = int32(7)
	if base.Ui32(v2) < base.Ui32(l0) {
		v5 = l0
	} else {
		v5 = v2
	}
	v7 = F_zmalloc_usable(m, v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+1)) = int32(0)
			v15 = int32(7)
			*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v15)
			v19 = int32(65280)
			*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(5)))) = uint16(v19)
		}
		return v7
	}
}
func F_lpPrepend(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 == int32(255) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v10 != int32(7) {
			F__serverAssert(m, int32(_a888), int32(_a889), int32(426))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v15 = l0 + int32(6)
			v16 = int32(0)
			v19 = F_lpInsert(m, l0, l1, v16, l2, v15, v16, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	} else {
		v15 = l0 + int32(6)
		v16 = int32(0)
		v19 = F_lpInsert(m, l0, l1, v16, l2, v15, v16, v16)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v19
		}
	}
}
func F_lpRandomPairsUnique(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v291 int32
	_ = v291
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = F_lpLength(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v23
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v26 != int32(255) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F__serverAssert(m, int32(_a888), int32(_a889), int32(400))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L92
	}
L4:
	;
	F__serverAssert(m, int32(_a888), int32(_a889), int32(400))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L91
	}
L5:
	;
	F__serverAssert(m, int32(_a891), int32(_a889), int32(1438))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L90
	}
L6:
	;
	m.G0 = v17 + int32(16)
	return v291
L7:
	;
	v39 = int32(base.Ui32(v19) >> (uint(int32(1)) % 32))
	if base.Ui32(l1) < base.Ui32(v39) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 == int32(7) {
		v291 = v23
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F__serverAssert(m, int32(_a888), int32(_a889), int32(426))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	v41 = l1
	goto L13
L12:
	;
	v41 = v39
	goto L13
L13:
	;
	if v41 == int32(0) {
		v291 = v23
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v47 = int32(0)
	v50 = l0 + int32(6)
	v54 = v47
	v55 = v47
	v57 = int64(0)
	v58 = v41
	goto L15
L15:
	;
	v66 = F_lpNextRandom(m, l0, v50, v17+int32(4), v58, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v291 = v279
	goto L6
L17:
	;
	if v66 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(0)
	v74 = F_lpGetWithSize(m, v66, v17+int32(8), v72, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v77 = v55 << (uint(int32(4)) % 32)
	v78 = l2 + v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = v57
	goto L22
L21:
	;
	v80 = v79
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v74
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v84 = base.I32_wrap_i64(v79)
	goto L25
L24:
	;
	v84 = v54
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v84
	v89 = int32(1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v91 = base.I32_extend8_s(v90)
	if v91 <= int32(-1) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if l3 == int32(0) {
		v188 = v84
		v189 = v80
		goto L53
	} else {
		goto L54
	}
L27:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v157 != int32(255) {
		goto L26
	} else {
		goto L50
	}
L28:
	;
	v156 = v66 + v154 + v153
	goto L27
L29:
	;
	if v90&int32(192) != int32(128) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v153 = v89
	v154 = int32(1)
	goto L28
L31:
	;
	if v90&int32(224) != int32(192) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v99 = int32(1)
	v153 = v99
	v154 = v90&int32(63) + v99
	goto L28
L33:
	;
	v112 = (v91 + int32(15)) & int32(255)
	if base.Ui32(v112) < base.Ui32(int32(4)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v153 = v89
	v154 = int32(2)
	goto L28
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v112<<(uint(int32(2))%32))+uint32(_consts[505])))
	v153 = v89
	v154 = v152
	goto L28
L36:
	;
	if v90&int32(240) != int32(224) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(v134) < base.Ui32(int32(128)) {
		v153 = v89
		v154 = v134
		goto L28
	} else {
		goto L42
	}
L38:
	;
	switch v90 + int32(-240) {
	case 0:
		goto L40
	default:
		goto L41
	case 15:
		v153 = v89
		v154 = int32(1)
		goto L28
	}
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	v134 = v119 | v90<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L37
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v66)+1))
	v134 = v131 + int32(5)
	goto L37
L41:
	;
	v153 = v89
	v154 = int32(0)
	goto L28
L42:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v134) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v134) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v153 = int32(2)
	v154 = v134
	goto L28
L45:
	;
	if base.Ui32(v134) < base.Ui32(int32(268435456)) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v153 = int32(3)
	v154 = v134
	goto L28
L47:
	;
	v147 = int32(4)
	goto L49
L48:
	;
	v147 = int32(5)
	goto L49
L49:
	;
	v153 = v147
	v154 = v134
	goto L28
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v156+int32(1) != l0+v162 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F__serverAssert(m, int32(_a892), int32(_a889), int32(1441))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v196 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v198 = base.I32_extend8_s(v197)
	if v198 <= int32(-1) {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	v176 = int32(0)
	v178 = F_lpGetWithSize(m, v156, v17+int32(8), v176, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v180 = l3 + v77
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	if v178 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v182 = v80
	goto L58
L57:
	;
	v182 = v181
	goto L58
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v178
	if v178 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v186 = base.I32_wrap_i64(v181)
	goto L61
L60:
	;
	v186 = v84
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v186
	v188 = v186
	v189 = v182
	goto L53
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v171 + int32(2)
	v279 = v55 + int32(1)
	if v273 == int32(0) {
		v291 = v279
		goto L6
	} else {
		goto L88
	}
L63:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v264 != int32(255) {
		v273 = v263
		goto L62
	} else {
		goto L86
	}
L64:
	;
	v263 = v156 + v261 + v260
	goto L63
L65:
	;
	if v197&int32(192) != int32(128) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v260 = v196
	v261 = int32(1)
	goto L64
L67:
	;
	if v197&int32(224) != int32(192) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v206 = int32(1)
	v260 = v206
	v261 = v197&int32(63) + v206
	goto L64
L69:
	;
	v219 = (v198 + int32(15)) & int32(255)
	if base.Ui32(v219) < base.Ui32(int32(4)) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v260 = v196
	v261 = int32(2)
	goto L64
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v219<<(uint(int32(2))%32))+uint32(_consts[505])))
	v260 = v196
	v261 = v259
	goto L64
L72:
	;
	if v197&int32(240) != int32(224) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if base.Ui32(v241) < base.Ui32(int32(128)) {
		v260 = v196
		v261 = v241
		goto L64
	} else {
		goto L78
	}
L74:
	;
	switch v197 + int32(-240) {
	case 0:
		goto L76
	default:
		goto L77
	case 15:
		v260 = v196
		v261 = int32(1)
		goto L64
	}
L75:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	v241 = v226 | v197<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L73
L76:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v156)+1))
	v241 = v238 + int32(5)
	goto L73
L77:
	;
	v260 = v196
	v261 = int32(0)
	goto L64
L78:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v241) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v241) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v260 = int32(2)
	v261 = v241
	goto L64
L81:
	;
	if base.Ui32(v241) < base.Ui32(int32(268435456)) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v260 = int32(3)
	v261 = v241
	goto L64
L83:
	;
	v254 = int32(4)
	goto L85
L84:
	;
	v254 = int32(5)
	goto L85
L85:
	;
	v260 = v254
	v261 = v241
	goto L64
L86:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v263+int32(1) != l0+v270 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	v273 = int32(0)
	goto L62
L88:
	;
	if base.Ui32(v279) < base.Ui32(v41) {
		v50 = v273
		v54 = v188
		v55 = v279
		v57 = v189
		v58 = v58 + int32(-1)
		goto L15
	} else {
		goto L89
	}
L89:
	;
	goto L16
L90:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpReplaceInteger(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = F_lpInsertInteger(m, l0, l2, v4, int32(2), l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_lpRepr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = F_lpLength(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v21
	v29 = F_iprintf(m, int32(_a893), v19+int32(48))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v31 != int32(255) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F__serverAssert(m, int32(_a888), int32(_a889), int32(400))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L95
	}
L5:
	;
	v366 = F_puts(m, int32(_a894))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L94
	}
L6:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	v57 = l0 + int32(6)
	v62 = v45
	goto L10
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 == int32(7) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F__serverAssert(m, int32(_a888), int32(_a889), int32(426))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v71 = int32(192)
	v72 = v70 & v71
	v73 = base.I32_extend8_s(v70)
	v77 = base.B2i32(v70&int32(224) == v71)
	if v70&int32(224) == v71 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v340+int32(1) != l0+v346 {
		goto L4
	} else {
		goto L93
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(40)))) = v164 - v165
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(36)))) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(32)))) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v62
	v172 = v166 + v164
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v57 - l0
	v180 = F_iprintf(m, int32(_a895), v19+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L46
	}
L13:
	;
	if v72 != int32(128) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v99 = int32(1)
	if int32(-1) < v73 {
		v164 = v99
		v165 = v99
		v166 = v99
		goto L12
	} else {
		goto L23
	}
L15:
	;
	if int32(-1) < v73 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32((v70+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if v72 == int32(128) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v70&int32(240) == int32(224) {
		v104 = int32(2)
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v73 == int32(-16) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v98 = int32(5)
	goto L22
L21:
	;
	v98 = base.B2i32(v73 == int32(-1))
	goto L22
L22:
	;
	v104 = v98
	goto L13
L23:
	;
	v104 = v99
	goto L13
L24:
	;
	if v77 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v109 = int32(1)
	v164 = v70&int32(63) + v109
	v165 = v104
	v166 = v109
	goto L12
L26:
	;
	switch v70 + int32(-241) {
	case 0:
		v164 = int32(3)
		v165 = v104
		v166 = int32(1)
		goto L12
	case 1:
		goto L31
	case 2:
		goto L30
	case 3:
		goto L29
	default:
		goto L28
	}
L27:
	;
	v164 = int32(2)
	v165 = v104
	v166 = int32(1)
	goto L12
L28:
	;
	if v70&int32(240) != int32(224) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v164 = int32(9)
	v165 = v104
	v166 = int32(1)
	goto L12
L30:
	;
	v164 = int32(5)
	v165 = v104
	v166 = int32(1)
	goto L12
L31:
	;
	v164 = int32(4)
	v165 = v104
	v166 = int32(1)
	goto L12
L32:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v148) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v140 = int32(1)
	switch v70 + int32(-240) {
	case 0:
		goto L35
	default:
		goto L36
	case 15:
		v164 = v140
		v165 = v104
		v166 = v140
		goto L12
	}
L34:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v148 = v132 | v70<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L32
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v57)+1))
	v148 = v145 + int32(5)
	goto L32
L36:
	;
	v164 = int32(0)
	v165 = v104
	v166 = v140
	goto L12
L37:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v148) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v164 = v148
	v165 = v104
	v166 = int32(1)
	goto L12
L39:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v148) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v164 = v148
	v165 = v104
	v166 = int32(2)
	goto L12
L41:
	;
	if base.Ui32(v148) < base.Ui32(int32(268435456)) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v164 = v148
	v165 = v104
	v166 = int32(3)
	goto L12
L43:
	;
	v163 = int32(4)
	goto L45
L44:
	;
	v163 = int32(5)
	goto L45
L45:
	;
	v164 = v148
	v165 = v104
	v166 = v163
	goto L12
L46:
	;
	v182 = int32(0)
	v185 = F_iprintf(m, int32(_a896), v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v172 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v231 = F_putchar(m, int32(10))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	v191 = v182
	goto L50
L50:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v191))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v206
	v209 = F_iprintf(m, int32(_a897), v19)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	goto L48
L52:
	;
	v212 = v191 + int32(1)
	if v212 != v172 {
		v191 = v212
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v238 = F_lpGetWithSize(m, v57, v19+int32(88), v19+int32(64), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v242 = F_iprintf(m, int32(_a898), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v19)+88))
	if v244 <= int64(40) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v266 = F_puts(m, int32(_a899))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L68
	}
L58:
	;
	v260 = F_fwrite(m, v238, base.I32_wrap_i64(v244), int32(1), v46)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L65
	}
L59:
	;
	v249 = F_fwrite(m, v238, int32(40), int32(1), v46)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	v256 = F_iprintf(m, int32(_a900), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	if v249 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	F_perror(m, int32(_a901))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L57
L65:
	;
	if v260 != 0 {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	F_perror(m, int32(_a901))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L57
L68:
	;
	v268 = int32(1)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v275 = base.I32_extend8_s(v274)
	if v275 <= int32(-1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	if v341 != int32(255) {
		v57 = v340
		v62 = v62 + v268
		goto L10
	} else {
		goto L92
	}
L70:
	;
	v340 = v57 + v338 + v337
	goto L69
L71:
	;
	if v274&int32(192) != int32(128) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v337 = v268
	v338 = int32(1)
	goto L70
L73:
	;
	if v274&int32(224) != int32(192) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v283 = int32(1)
	v337 = v283
	v338 = v274&int32(63) + v283
	goto L70
L75:
	;
	v296 = (v275 + int32(15)) & int32(255)
	if base.Ui32(v296) < base.Ui32(int32(4)) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v337 = v268
	v338 = int32(2)
	goto L70
L77:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v296<<(uint(int32(2))%32))+uint32(_consts[505])))
	v337 = v268
	v338 = v336
	goto L70
L78:
	;
	if v274&int32(240) != int32(224) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if base.Ui32(v318) < base.Ui32(int32(128)) {
		v337 = v268
		v338 = v318
		goto L70
	} else {
		goto L84
	}
L80:
	;
	switch v274 + int32(-240) {
	case 0:
		goto L82
	default:
		goto L83
	case 15:
		v337 = v268
		v338 = int32(1)
		goto L70
	}
L81:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v318 = v303 | v274<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L79
L82:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v57)+1))
	v318 = v315 + int32(5)
	goto L79
L83:
	;
	v337 = v268
	v338 = int32(0)
	goto L70
L84:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v318) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v318) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v337 = int32(2)
	v338 = v318
	goto L70
L87:
	;
	if base.Ui32(v318) < base.Ui32(int32(268435456)) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v337 = int32(3)
	v338 = v318
	goto L70
L89:
	;
	v331 = int32(4)
	goto L91
L90:
	;
	v331 = int32(5)
	goto L91
L91:
	;
	v337 = v331
	v338 = v318
	goto L70
L92:
	;
	goto L11
L93:
	;
	goto L5
L94:
	;
	m.G0 = v19 + int32(96)
	return
L95:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lpSkip(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = int32(1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v7 = base.I32_extend8_s(v6)
	if v7 <= int32(-1) {
		if v6&int32(192) != int32(128) {
			if v6&int32(224) != int32(192) {
				v28 = (v7 + int32(15)) & int32(255)
				if base.Ui32(v28) < base.Ui32(int32(4)) {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32))+uint32(_consts[505])))
					v69 = v5
					v70 = v68
				} else {
					if v6&int32(240) != int32(224) {
						switch v6 + int32(-240) {
						case 0:
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1))
							v50 = v47 + int32(5)
							if base.Ui32(v50) < base.Ui32(int32(128)) {
								v69 = v5
								v70 = v50
							} else {
								if base.Ui32(int32(16384)) <= base.Ui32(v50) {
									if base.Ui32(int32(2097152)) <= base.Ui32(v50) {
										if base.Ui32(v50) < base.Ui32(int32(268435456)) {
											v63 = int32(4)
										} else {
											v63 = int32(5)
										}
										v69 = v63
										v70 = v50
									} else {
										v69 = int32(3)
										v70 = v50
									}
								} else {
									v69 = int32(2)
									v70 = v50
								}
							}
						default:
							v69 = v5
							v70 = int32(0)
						case 15:
							v69 = v5
							v70 = int32(1)
						}
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						v50 = v35 | v6<<(uint(int32(8))%32)&int32(3840) + int32(2)
						if base.Ui32(v50) < base.Ui32(int32(128)) {
							v69 = v5
							v70 = v50
						} else {
							if base.Ui32(int32(16384)) <= base.Ui32(v50) {
								if base.Ui32(int32(2097152)) <= base.Ui32(v50) {
									if base.Ui32(v50) < base.Ui32(int32(268435456)) {
										v63 = int32(4)
									} else {
										v63 = int32(5)
									}
									v69 = v63
									v70 = v50
								} else {
									v69 = int32(3)
									v70 = v50
								}
							} else {
								v69 = int32(2)
								v70 = v50
							}
						}
					}
				}
			} else {
				v69 = v5
				v70 = int32(2)
			}
		} else {
			v15 = int32(1)
			v69 = v15
			v70 = v6&int32(63) + v15
		}
	} else {
		v69 = v5
		v70 = int32(1)
	}
	return l0 + v70 + v69
}
