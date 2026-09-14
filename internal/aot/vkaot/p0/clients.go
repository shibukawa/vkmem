package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clientsCronResizeOutputBuffer(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	if v8 == v3 {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
			if v12&int32(1) != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				if base.Ui32(v15) < base.Ui32(int32(2048)) {
					v35 = int32(0)
					v36 = int32(1)
					v38 = v15 << (uint(v36) % 32)
					if base.Ui32(int32(32767)) < base.Ui32(v38) {
						v62 = v35
						v63 = v36
						v65 = int32(0)
						v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
						if v66 < v65 {
						} else {
							v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
							if l1-v69 < base.I64_extend_i32_u(v66) {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
							}
						}
						if v63 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
							v79 = F_zmalloc_usable(m, v62, l0+int32(128))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								if v84 == int32(0) {
								} else {
									v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
									mBase = m.M
								}
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
								if v76 != v89 {
								} else {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
								}
								F_zfree_with_size(m, v76, v15)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
						if v41 != v15 {
							v62 = v35
							v63 = v36
							v65 = int32(0)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
							if v66 < v65 {
							} else {
								v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
								if l1-v69 < base.I64_extend_i32_u(v66) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
								}
							}
							if v63 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v79 = F_zmalloc_usable(m, v62, l0+int32(128))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v84 == int32(0) {
									} else {
										v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
										mBase = m.M
									}
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v76 != v89 {
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
									}
									F_zfree_with_size(m, v76, v15)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v43 = int32(0)
							v46 = *(*int64)(unsafe.Add(mBase, _consts[683]))
							*(*int64)(unsafe.Add(mBase, _consts[683])) = v46 + int64(1)
							if v38 == v43 {
								v62 = v43
								v63 = v36
								v65 = int32(0)
								v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
								if v66 < v65 {
								} else {
									v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
									if l1-v69 < base.I64_extend_i32_u(v66) {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
									}
								}
								if v63 != 0 {
									return int32(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v79 = F_zmalloc_usable(m, v62, l0+int32(128))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v84 == int32(0) {
										} else {
											v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
											mBase = m.M
										}
										v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										if v76 != v89 {
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
										}
										F_zfree_with_size(m, v76, v15)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								v52 = int32(16384)
								if base.Ui32(v38) < base.Ui32(v52) {
									v55 = v38
								} else {
									v55 = v52
								}
								v56 = v55
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
								if base.Ui32(v56) < base.Ui32(v60) {
									F__serverAssertWithInfo(m, l0, int32(0), int32(_a1239), int32(_a1240), int32(984))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v62 = v56
									v63 = int32(0)
									v65 = int32(0)
									v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
									if v66 < v65 {
									} else {
										v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
										if l1-v69 < base.I64_extend_i32_u(v66) {
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
										}
									}
									if v63 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v79 = F_zmalloc_usable(m, v62, l0+int32(128))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v84 == int32(0) {
											} else {
												v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
												mBase = m.M
											}
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											if v76 != v89 {
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
											}
											F_zfree_with_size(m, v76, v15)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												return int32(0)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
					if base.Ui32(int32(base.Ui32(v15)>>(uint(int32(1))%32))) <= base.Ui32(v18) {
						v35 = int32(0)
						v36 = int32(1)
						v38 = v15 << (uint(v36) % 32)
						if base.Ui32(int32(32767)) < base.Ui32(v38) {
							v62 = v35
							v63 = v36
							v65 = int32(0)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
							if v66 < v65 {
							} else {
								v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
								if l1-v69 < base.I64_extend_i32_u(v66) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
								}
							}
							if v63 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v79 = F_zmalloc_usable(m, v62, l0+int32(128))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v84 == int32(0) {
									} else {
										v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
										mBase = m.M
									}
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v76 != v89 {
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
									}
									F_zfree_with_size(m, v76, v15)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
							if v41 != v15 {
								v62 = v35
								v63 = v36
								v65 = int32(0)
								v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
								if v66 < v65 {
								} else {
									v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
									if l1-v69 < base.I64_extend_i32_u(v66) {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
									}
								}
								if v63 != 0 {
									return int32(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v79 = F_zmalloc_usable(m, v62, l0+int32(128))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
										if v84 == int32(0) {
										} else {
											v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
											mBase = m.M
										}
										v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
										if v76 != v89 {
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
										}
										F_zfree_with_size(m, v76, v15)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								}
							} else {
								v43 = int32(0)
								v46 = *(*int64)(unsafe.Add(mBase, _consts[683]))
								*(*int64)(unsafe.Add(mBase, _consts[683])) = v46 + int64(1)
								if v38 == v43 {
									v62 = v43
									v63 = v36
									v65 = int32(0)
									v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
									if v66 < v65 {
									} else {
										v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
										if l1-v69 < base.I64_extend_i32_u(v66) {
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
										}
									}
									if v63 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										v79 = F_zmalloc_usable(m, v62, l0+int32(128))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
											if v84 == int32(0) {
											} else {
												v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
												mBase = m.M
											}
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
											if v76 != v89 {
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
											}
											F_zfree_with_size(m, v76, v15)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												return int32(0)
											}
										}
									}
								} else {
									v52 = int32(16384)
									if base.Ui32(v38) < base.Ui32(v52) {
										v55 = v38
									} else {
										v55 = v52
									}
									v56 = v55
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if base.Ui32(v56) < base.Ui32(v60) {
										F__serverAssertWithInfo(m, l0, int32(0), int32(_a1239), int32(_a1240), int32(984))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v62 = v56
										v63 = int32(0)
										v65 = int32(0)
										v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
										if v66 < v65 {
										} else {
											v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
											if l1-v69 < base.I64_extend_i32_u(v66) {
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
											}
										}
										if v63 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
											v79 = F_zmalloc_usable(m, v62, l0+int32(128))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
												if v84 == int32(0) {
												} else {
													v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
													mBase = m.M
												}
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
												if v76 != v89 {
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
												}
												F_zfree_with_size(m, v76, v15)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													return int32(0)
												}
											}
										}
									}
								}
							}
						}
					} else {
						v22 = int32(0)
						v24 = *(*int64)(unsafe.Add(mBase, _consts[684]))
						*(*int64)(unsafe.Add(mBase, _consts[684])) = v24 + int64(1)
						v28 = int32(1023)
						if base.Ui32(v28) < base.Ui32(v18) {
							v31 = v18
						} else {
							v31 = v28
						}
						v56 = v31 + int32(1)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
						if base.Ui32(v56) < base.Ui32(v60) {
							F__serverAssertWithInfo(m, l0, int32(0), int32(_a1239), int32(_a1240), int32(984))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v62 = v56
							v63 = int32(0)
							v65 = int32(0)
							v66 = *(*int32)(unsafe.Add(mBase, _consts[682]))
							if v66 < v65 {
							} else {
								v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+312))
								if l1-v69 < base.I64_extend_i32_u(v66) {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+312)) = l1
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v74
								}
							}
							if v63 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v79 = F_zmalloc_usable(m, v62, l0+int32(128))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v79
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v84 == int32(0) {
									} else {
										v87 = F__emscripten_memcpy_bulkmem(m, v79, v76, v84)
										mBase = m.M
									}
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
									if v76 != v89 {
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v91
									}
									F_zfree_with_size(m, v76, v15)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										return int32(0)
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
func F_freeClientsInAsyncFreeQueue(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
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
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v14 = v9 + int32(40)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15
	goto L1
L1:
	;
	v19 = int32(0)
	v21 = v9 + int32(40)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v23 == v19 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return v157
L3:
	;
	if v23 == int32(0) {
		v157 = v19
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
	v37 = v19
	v38 = v23
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+207)))
	if v43&int32(2) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v157 = v139
	goto L2
L9:
	;
	v143 = v9 + int32(40)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v145 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
	if v117 < int32(0) {
		v139 = v37
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+176))
	if v49 != int64(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v86 = int32(_a69)
	v87 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+176))
	v92 = int64(*(*int32)(unsafe.Add(mBase, _consts[429])))
	if v87-v89 <= v92 {
		v139 = v37
		goto L9
	} else {
		goto L21
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v52 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v63 = int32(_a69)
	v64 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	*(*int64)(unsafe.Add(mBase, uint32(v62)+176)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(1) < v67 {
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+84))
	v58 = m.T0[v57].(func(*base.Module, int32, int32) int32)(m, v52, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	goto L14
L18:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v71 = F_replicationGetReplicaName(m, v42)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v71
	v76 = *(*int32)(unsafe.Add(mBase, _consts[429]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v76
	F__serverLog(m, int32(1), int32(_a801), v9+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v95 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v42)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+204)) = v111 & int32(-33554433)
	goto L10
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)+176))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v100
	v103 = *(*int64)(unsafe.Add(mBase, _consts[37]))
	v104 = v103 - v99
	*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v104)
	F__serverLog(m, int32(2), int32(_a802), v9)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+222)))
	if v122 != 0 {
		v126 = int32(1)
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v126 != 0 {
		v139 = v37
		goto L9
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+223)))
	v126 = base.B2i32(v123 != int32(0))
	goto L27
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v42)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+200)) = v127 & int32(-1025)
	v131 = F_freeClient(m, v42)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	F_listDelNode(m, v134, v38)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v139 = v37 + int32(1)
	goto L9
L32:
	;
	if v145 != 0 {
		v37 = v139
		v38 = v145
		goto L7
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v145+base.B2i32(v148 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v154
	goto L33
L35:
	;
	goto L8
}
func F_processClientsWaitingReplicas(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[595]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v18
	goto L1
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v14 + int32(16)
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23+base.B2i32(v26 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v32
	goto L4
L6:
	;
	v36 = int32(0)
	v37 = int64(0)
	v41 = v23
	v42 = v36
	v43 = v37
	v44 = v37
	v45 = v36
	goto L7
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v55 = base.B2i32(v53 != int32(20))
	if v53 != int32(20) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	goto L2
L9:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v286 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L10:
	;
	F_unblockClient(m, v51, int32(1))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L21
	} else {
		goto L66
	}
L11:
	;
	F_addReplyLongLong(m, v51, base.I64_extend_i32_s(v215))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L21
	} else {
		goto L65
	}
L12:
	;
	v245 = *(*int64)(unsafe.Add(mBase, _consts[524]))
	v246 = base.B2i32(v242 <= v245)
	if v241 <= v246 {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v207)+32))
	v237 = v210
	v241 = v232
	v242 = v210
	v243 = v206
	goto L12
L14:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v51)+204))
	if v224&int32(2) == int32(0) {
		goto L11
	} else {
		goto L59
	}
L15:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+28))
	if v206 < v208 {
		v276 = v42
		v277 = v43
		v278 = v44
		v279 = v45
		goto L9
	} else {
		goto L57
	}
L16:
	;
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v56)+40))
	if v44 == int64(0) {
		goto L40
	} else {
		goto L41
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)+40))
	if v43 == int64(0) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	if v61 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_addReplyError(m, v51, int32(_a1055))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v262 = v42
	v263 = v43
	v264 = v44
	v265 = v45
	goto L10
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v75 = v14 + int32(8)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76
	goto L27
L24:
	;
	if v43 < v66 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	if v70 <= v42 {
		v215 = v42
		v216 = v43
		goto L14
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v80 = int32(0)
	v82 = v14 + int32(8)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v84 == v80 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v84 == int32(0) {
		v206 = v80
		goto L15
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84+base.B2i32(v87 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v93
	goto L29
L31:
	;
	v98 = v84
	v107 = v80
	goto L32
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+104))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v110 != int32(9) {
		v116 = v107
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v118 = v14 + int32(8)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v120 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v109)+64))
	v116 = v107 + base.B2i32(v66 <= v113)
	goto L34
L36:
	;
	if v120 != 0 {
		v98 = v120
		v107 = v116
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120+base.B2i32(v123 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v129
	goto L37
L39:
	;
	v206 = v116
	goto L15
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v140 = v14 + int32(8)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v141
	goto L44
L41:
	;
	if v44 < v131 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	if v45 < v135 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v237 = v44
	v241 = v57
	v242 = v131
	v243 = v45
	goto L12
L44:
	;
	v145 = int32(0)
	v147 = v14 + int32(8)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v149 == v145 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v149 == int32(0) {
		v206 = v145
		goto L15
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v149+base.B2i32(v152 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v158
	goto L46
L48:
	;
	v163 = v149
	v172 = v145
	goto L49
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+104))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 != int32(9) {
		v181 = v172
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v206 = v181
	goto L15
L51:
	;
	v183 = v14 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v185 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v174)+72))
	v181 = v172 + base.B2i32(v131 <= v178)
	goto L51
L53:
	;
	if v185 != 0 {
		v163 = v185
		v172 = v181
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185+base.B2i32(v188 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v194
	goto L54
L56:
	;
	goto L50
L57:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v207)+40))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v215 = v206
	v216 = v210
	goto L14
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+204)) = v224 | int32(4194304)
	v262 = v215
	v263 = v216
	v264 = v44
	v265 = v45
	goto L10
L60:
	;
	F_addReplyArrayLen(m, v51, int32(2))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L62
	}
L61:
	;
	v276 = v42
	v277 = v43
	v278 = v237
	v279 = v243
	goto L9
L62:
	;
	F_addReplyLongLong(m, v51, base.I64_extend_i32_u(v246))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	F_addReplyLongLong(m, v51, base.I64_extend_i32_s(v243))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L64
	}
L64:
	;
	v262 = v42
	v263 = v43
	v264 = v237
	v265 = v243
	goto L10
L65:
	;
	v262 = v215
	v263 = v216
	v264 = v44
	v265 = v45
	goto L10
L66:
	;
	v276 = v262
	v277 = v263
	v278 = v264
	v279 = v265
	goto L9
L67:
	;
	if v286 != 0 {
		v41 = v286
		v42 = v276
		v43 = v277
		v44 = v278
		v45 = v279
		goto L7
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v286+base.B2i32(v289 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v295
	goto L68
L70:
	;
	goto L8
}
