package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_mp_encode_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	v4 = l3
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	if base.Ui32(int32(31)) < base.Ui32(v4) {
		if base.Ui32(int32(255)) < base.Ui32(v4) {
			if base.Ui32(int32(65535)) < base.Ui32(v4) {
				v37 = int32(219)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v37)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+7)) = uint8(v4)
				v41 = int32(base.Ui32(v4) >> (uint(int32(8)) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+6)) = uint8(v41)
				v44 = int32(base.Ui32(v4) >> (uint(int32(16)) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v44)
				v47 = int32(base.Ui32(v4) >> (uint(int32(24)) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v47)
				v50 = int32(5)
			} else {
				v30 = int32(218)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v30)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v4)
				v34 = int32(base.Ui32(v4) >> (uint(int32(8)) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v34)
				v50 = int32(3)
			}
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v4)
			v25 = int32(217)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v25)
			v50 = int32(2)
		}
	} else {
		v19 = v4 | int32(160)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v19)
		v50 = int32(1)
	}
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v51) < base.Ui32(v50) {
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v56 = v55 + v50
		if base.Ui32(v56) < base.Ui32(v55) {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			if base.Ui32(int32(2147483647)) <= base.Ui32(v56) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v62 = v14 + int32(8)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v62 == int32(0) {
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v62))) = v66
				}
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v72 = v56 << (uint(int32(1)) % 32)
				v73 = m.T0[v68].(func(*base.Module, int32, int32, int32, int32) int32)(m, v69, v60, v55+v51, v72)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v73
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v72 - v76
					v79 = v76
					v80 = v73
					if v50 == int32(0) {
					} else {
						v89 = F__emscripten_memcpy_bulkmem(m, v80+v79, v14+int32(3), v50)
						mBase = m.M
					}
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v92 = v91 + v50
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v92
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v95 = v94 - v50
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v95
					if base.Ui32(v95) < base.Ui32(v4) {
						v99 = v92 + v4
						if base.Ui32(v99) < base.Ui32(v92) {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						} else {
							if base.Ui32(int32(2147483647)) <= base.Ui32(v99) {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v105 = v14 + int32(12)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v105 == int32(0) {
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v105))) = v109
								}
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
								v115 = v99 << (uint(int32(1)) % 32)
								v116 = m.T0[v111].(func(*base.Module, int32, int32, int32, int32) int32)(m, v112, v103, v94+v91, v115)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v115 - v119
									v123 = v116
									v124 = v119
									if v4 == int32(0) {
									} else {
										v129 = F__emscripten_memcpy_bulkmem(m, v123+v124, l2, v4)
										mBase = m.M
									}
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v131 + v4
									v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v134 - v4
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v123 = v98
						v124 = v92
						if v4 == int32(0) {
						} else {
							v129 = F__emscripten_memcpy_bulkmem(m, v123+v124, l2, v4)
							mBase = m.M
						}
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v131 + v4
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v134 - v4
						m.G0 = v14 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v79 = v53
		v80 = v54
		if v50 == int32(0) {
		} else {
			v89 = F__emscripten_memcpy_bulkmem(m, v80+v79, v14+int32(3), v50)
			mBase = m.M
		}
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v92 = v91 + v50
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v92
		v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v95 = v94 - v50
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v95
		if base.Ui32(v95) < base.Ui32(v4) {
			v99 = v92 + v4
			if base.Ui32(v99) < base.Ui32(v92) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				if base.Ui32(int32(2147483647)) <= base.Ui32(v99) {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v105 = v14 + int32(12)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v105 == int32(0) {
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v105))) = v109
					}
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					v115 = v99 << (uint(int32(1)) % 32)
					v116 = m.T0[v111].(func(*base.Module, int32, int32, int32, int32) int32)(m, v112, v103, v94+v91, v115)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v115 - v119
						v123 = v116
						v124 = v119
						if v4 == int32(0) {
						} else {
							v129 = F__emscripten_memcpy_bulkmem(m, v123+v124, l2, v4)
							mBase = m.M
						}
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v131 + v4
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v134 - v4
						m.G0 = v14 + int32(16)
						return
					}
				}
			}
		} else {
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v123 = v98
			v124 = v92
			if v4 == int32(0) {
			} else {
				v129 = F__emscripten_memcpy_bulkmem(m, v123+v124, l2, v4)
				mBase = m.M
			}
			v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v131 + v4
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v134 - v4
			m.G0 = v14 + int32(16)
			return
		}
	}
}
func F_mp_encode_map(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
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
	var v53 int32
	_ = v53
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
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if base.Ui64(int64(15)) < base.Ui64(v3) {
		if base.Ui64(int64(65535)) < base.Ui64(v3) {
			v30 = int32(223)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v30)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v3)
			v34 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v34)
			v37 = int64(base.Ui64(v3) >> (uint(int64(16)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v37)
			v40 = int64(base.Ui64(v3) >> (uint(int64(24)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v40)
			v43 = int32(5)
		} else {
			v23 = int32(222)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v23)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v3)
			v27 = int64(base.Ui64(v3) >> (uint(int64(8)) % 64))
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v27)
			v43 = int32(3)
		}
	} else {
		v18 = base.I32_wrap_i64(v3) | int32(128)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v18)
		v43 = int32(1)
	}
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v44) < base.Ui32(v43) {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v49 = v48 + v43
		if base.Ui32(v49) < base.Ui32(v48) {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		} else {
			if base.Ui32(int32(2147483647)) <= base.Ui32(v49) {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v55 = v12 + int32(12)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v55 == int32(0) {
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59
				}
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v65 = v49 << (uint(int32(1)) % 32)
				v66 = m.T0[v61].(func(*base.Module, int32, int32, int32, int32) int32)(m, v62, v53, v48+v44, v65)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v65 - v69
					v72 = v69
					v73 = v66
					if v43 == int32(0) {
					} else {
						v82 = F__emscripten_memcpy_bulkmem(m, v73+v72, v12+int32(7), v43)
						mBase = m.M
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84 + v43
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v87 - v43
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v72 = v46
		v73 = v47
		if v43 == int32(0) {
		} else {
			v82 = F__emscripten_memcpy_bulkmem(m, v73+v72, v12+int32(7), v43)
			mBase = m.M
		}
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84 + v43
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v87 - v43
		m.G0 = v12 + int32(16)
		return
	}
}
func F_mp_unpack_full(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v16 = F_luaL_checklstring(m, l0, int32(1), v11+int32(44))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v21 = l2 | l1
	if int32(-1) < v21 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	m.G0 = v11 + int32(48)
	return v209
L5:
	;
	if base.Ui32(l2) <= base.Ui32(v20) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	v26 = m.G3
	v31 = F_luaL_error(m, l0, v26+int32(_a2754), v11+int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v209 = v31
	goto L4
L8:
	;
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v41
	v44 = v20 - l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v16 + l2
	if v44 == v41 {
		v92 = v44
		v94 = v41
		goto L11
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	v36 = m.G3
	v39 = F_luaL_error(m, l0, v36+int32(_a2755), v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v209 = v39
	goto L4
L11:
	;
	if v21 == int32(0) {
		v209 = v94
		goto L4
	} else {
		goto L27
	}
L12:
	;
	if v21 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = l1
	goto L15
L14:
	;
	v51 = int32(2147483647)
	goto L15
L15:
	;
	if v51 < int32(1) {
		v92 = v44
		v94 = v41
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v61 = v41
	goto L17
L17:
	;
	F_mp_decode_to_lua_type(m, l0, v11+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v92 = v83
	v94 = v82
	goto L11
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	switch v66 + int32(-1) {
	case 0:
		goto L22
	case 1:
		goto L21
	default:
		goto L20
	}
L20:
	;
	v82 = v61 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	if v83 == int32(0) {
		v92 = v83
		v94 = v82
		goto L11
	} else {
		goto L25
	}
L21:
	;
	v75 = m.G3
	v79 = F_luaL_error(m, l0, v75+int32(_a2756), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v69 = m.G3
	v73 = F_luaL_error(m, l0, v69+int32(_a2757), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v209 = v73
	goto L4
L24:
	;
	v209 = v79
	goto L4
L25:
	;
	if v82 < v51 {
		v61 = v82
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v98 = v97 - v92
	if v98 <= int32(-1) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v102 = m.G3
	F_luaL_checkstack(m, l0, int32(1), v102+int32(_a2758))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	if v108 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v109 = v98
	goto L32
L31:
	;
	v109 = int32(-1)
	goto L32
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v111))) = base.F64_convert_i32_s(v109)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v116 + int32(16)
	goto L33
L33:
	;
	goto L37
L34:
	;
	v209 = v94 + int32(1)
	goto L4
L35:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v176) <= base.Ui32(v134) {
		v193 = v176
		goto L50
	} else {
		goto L51
	}
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = v126 + int32(16)
	v132 = m.G398
	if base.Ui32(v131) < base.Ui32(v125) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v134 = v131
	goto L40
L39:
	;
	v134 = v132
	goto L40
L40:
	;
	goto L35
L50:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	*(*int64)(unsafe.Add(mBase, uint32(v134))) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v198
	goto L34
L51:
	;
	v179 = v176
	goto L52
L52:
	;
	v183 = v179 + int32(-16)
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, uint32(v179))) = v184
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v188
	if base.Ui32(v134) < base.Ui32(v183) {
		v179 = v183
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v193 = v191
	goto L50
L54:
	;
	goto L53
}
func F_mp_unpack_limit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v5 = F_luaL_checkinteger(m, l0, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = F_luaL_optinteger(m, l0, int32(3), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = int32(0) - (v14-v15)>>(uint(int32(4))%32)
			if v19 < int32(0) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v44 = v37 + v19<<(uint(int32(4))%32) + int32(16)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v27 = v24 + v19<<(uint(int32(4))%32)
				if base.Ui32(v27) <= base.Ui32(v23) {
					v44 = v27
				} else {
					v31 = v23
					for {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(0)
						v35 = v31 + int32(16)
						if base.Ui32(v35) < base.Ui32(v27) {
							v31 = v35
							continue
						} else {
							break
						}
						break
					}
					v44 = v27
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
			v47 = F_mp_unpack_full(m, l0, v5, v11)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				return v47
			}
		}
	}
}
func F_mp_unpack_one(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = F_luaL_optinteger(m, l0, int32(2), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v15 = int32(0) - (v10-v11)>>(uint(int32(4))%32)
		if v15 < int32(0) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v40 = v33 + v15<<(uint(int32(4))%32) + int32(16)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = v20 + v15<<(uint(int32(4))%32)
			if base.Ui32(v23) <= base.Ui32(v19) {
				v40 = v23
			} else {
				v27 = v19
				for {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = int32(0)
					v31 = v27 + int32(16)
					if base.Ui32(v31) < base.Ui32(v23) {
						v27 = v31
						continue
					} else {
						break
					}
					break
				}
				v40 = v23
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v40
		v44 = F_mp_unpack_full(m, l0, int32(1), v5)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			return v44
		}
	}
}
