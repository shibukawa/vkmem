package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___wasi_fd_is_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v5+int32(8))
	mBase = m.M
	if v9 != 0 {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v9
		v14 = int32(0)
	} else {
		v14 = int32(1)
	}
	m.G0 = v5 + int32(32)
	return v14
}
func F___wasi_syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	if l0 != 0 {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = l0
		return int32(-1)
	} else {
		return int32(0)
	}
}
func F___wasm_call_ctors(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	m.G2 = int32(_a0)
	m.G1 = int32(0)
	F___emscripten_environ_constructor(m)
	mBase = m.M
	v10 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(9116516)
	v14 = F_getpid(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v14
	return
}
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = int32(1)
	if base.Ui32(v3) < base.Ui32(l1) {
		v6 = l1
	} else {
		v6 = v3
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l0
	{
		m.ExcTag = uint32(int32(0))
		m.ExcVals[0] = uint64(uint32(l0 + int32(8)))
		m.ExcPending = 1
	}
	return
}
func F___wasm_setjmp(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
	return
}
func F_wangHash64(m *base.Module, l0 int64) int64 {
	var v2 int64
	_ = v2
	var v6 int64
	_ = v6
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	v2 = int64(21)
	v6 = l0<<(uint(v2)%64) + (l0 ^ int64(-1))
	v11 = (int64(base.Ui64(v6)>>(uint(int64(24))%64)) ^ v6) * int64(265)
	v16 = (int64(base.Ui64(v11)>>(uint(int64(14))%64)) ^ v11) * v2
	return (int64(base.Ui64(v16)>>(uint(int64(28))%64)) ^ v16) * int64(2147483649)
}
func F_watchCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v3&int32(32) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	F_addReply(m, l0, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L12
	}
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v6 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15 < int32(2) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v8 = F_valkey_calloc(m, int32(56))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v12
	goto L3
L7:
	;
	v20 = int32(1)
	goto L8
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v20<<(uint(int32(2))%32))))
	F_watchForKey(m, l0, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	goto L1
L10:
	;
	v29 = v20 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v29 < v30 {
		v20 = v29
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return
}
func F_watchForKey(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	if v7 != 0 {
	} else {
		v8 = int32(_a44)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[588]))
		*(*int32)(unsafe.Add(mBase, _consts[588])) = v10 + int32(1)
	}
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	if v14 != 0 {
		v23 = v14
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v25<<(uint(int32(2))%32))))
		if v29 != 0 {
			v41 = v29
			v43 = F_hashtableFind(m, v41, l1, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				if v43 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
					v47 = F_dictFetchValue(m, v46, l1)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 != 0 {
							v57 = v47
							v59 = F_valkey_malloc(m, int32(28))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
								v65 = F_keyIsExpired(m, v63, l1)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
									v72 = v67&int32(254) | v65&int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
									F_incrRefCount(m, l1)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										v79 = F_listAddNodeTail(m, v76+int32(24), v59)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
											if v84 != 0 {
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
												*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
												v95 = v90
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
												*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
												v87 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
												v95 = v87
											}
											*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
											*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
											v108 = F_hashtableAdd(m, v107, v59)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						} else {
							v49 = F_listCreate(m)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
								v53 = F_dictAdd(m, v52, l1, v49)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									F_incrRefCount(m, l1)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v57 = v49
										v59 = F_valkey_malloc(m, int32(28))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
											v65 = F_keyIsExpired(m, v63, l1)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
												v72 = v67&int32(254) | v65&int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
												F_incrRefCount(m, l1)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v79 = F_listAddNodeTail(m, v76+int32(24), v59)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
														if v84 != 0 {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
															*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
															v95 = v90
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
															*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
															v87 = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
															v95 = v87
														}
														*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
														*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
														v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
														v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
														v108 = F_hashtableAdd(m, v107, v59)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
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
				}
			}
		} else {
			v31 = F_hashtableCreate(m, int32(_a975))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32)))) = v31
				v41 = v31
				v43 = F_hashtableFind(m, v41, l1, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
						v47 = F_dictFetchValue(m, v46, l1)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							if v47 != 0 {
								v57 = v47
								v59 = F_valkey_malloc(m, int32(28))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
									v65 = F_keyIsExpired(m, v63, l1)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
										v72 = v67&int32(254) | v65&int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
										F_incrRefCount(m, l1)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v79 = F_listAddNodeTail(m, v76+int32(24), v59)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
												if v84 != 0 {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
													v95 = v90
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
													v95 = v87
												}
												*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
												*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
												v108 = F_hashtableAdd(m, v107, v59)
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							} else {
								v49 = F_listCreate(m)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
									v53 = F_dictAdd(m, v52, l1, v49)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										F_incrRefCount(m, l1)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											v57 = v49
											v59 = F_valkey_malloc(m, int32(28))
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
												*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
												v65 = F_keyIsExpired(m, v63, l1)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
													v72 = v67&int32(254) | v65&int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
													F_incrRefCount(m, l1)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v79 = F_listAddNodeTail(m, v76+int32(24), v59)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
															if v84 != 0 {
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
																*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																v95 = v90
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
																*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																v87 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
																v95 = v87
															}
															*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
															v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
															v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
															v108 = F_hashtableAdd(m, v107, v59)
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
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
					}
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[65]))
		v19 = F_valkey_calloc(m, v16<<(uint(int32(2))%32))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v19
			v23 = v19
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v25<<(uint(int32(2))%32))))
			if v29 != 0 {
				v41 = v29
				v43 = F_hashtableFind(m, v41, l1, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
						v47 = F_dictFetchValue(m, v46, l1)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							if v47 != 0 {
								v57 = v47
								v59 = F_valkey_malloc(m, int32(28))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
									v65 = F_keyIsExpired(m, v63, l1)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
										v72 = v67&int32(254) | v65&int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
										F_incrRefCount(m, l1)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											v79 = F_listAddNodeTail(m, v76+int32(24), v59)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
												if v84 != 0 {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
													v95 = v90
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
													v87 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
													v95 = v87
												}
												*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
												*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
												v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
												v108 = F_hashtableAdd(m, v107, v59)
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							} else {
								v49 = F_listCreate(m)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
									v53 = F_dictAdd(m, v52, l1, v49)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										F_incrRefCount(m, l1)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											v57 = v49
											v59 = F_valkey_malloc(m, int32(28))
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
												*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
												v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
												v65 = F_keyIsExpired(m, v63, l1)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
													v72 = v67&int32(254) | v65&int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
													F_incrRefCount(m, l1)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
														v79 = F_listAddNodeTail(m, v76+int32(24), v59)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
															if v84 != 0 {
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
																*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																v95 = v90
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
																*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																v87 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
																v95 = v87
															}
															*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
															v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
															v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
															v108 = F_hashtableAdd(m, v107, v59)
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
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
					}
				}
			} else {
				v31 = F_hashtableCreate(m, int32(_a975))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32)))) = v31
					v41 = v31
					v43 = F_hashtableFind(m, v41, l1, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v43 != 0 {
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
							v47 = F_dictFetchValue(m, v46, l1)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								if v47 != 0 {
									v57 = v47
									v59 = F_valkey_malloc(m, int32(28))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
										v65 = F_keyIsExpired(m, v63, l1)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
											v72 = v67&int32(254) | v65&int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
											F_incrRefCount(m, l1)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
												v79 = F_listAddNodeTail(m, v76+int32(24), v59)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
													if v84 != 0 {
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
														*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
														v95 = v90
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
														*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
														v87 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
														v95 = v87
													}
													*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
													v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
													v108 = F_hashtableAdd(m, v107, v59)
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														return
													}
												}
											}
										}
									}
								} else {
									v49 = F_listCreate(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
										v53 = F_dictAdd(m, v52, l1, v49)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											F_incrRefCount(m, l1)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return
											} else {
												v57 = v49
												v59 = F_valkey_malloc(m, int32(28))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
													v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
													*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v63
													v65 = F_keyIsExpired(m, v63, l1)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)))
														v72 = v67&int32(254) | v65&int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v72)
														F_incrRefCount(m, l1)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return
														} else {
															v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
															v79 = F_listAddNodeTail(m, v76+int32(24), v59)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v57
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
																if v84 != 0 {
																	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v59
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																	v95 = v90
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v59
																	v87 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v87
																	v95 = v87
																}
																*(*int32)(unsafe.Add(mBase, uint32(v59))) = v95
																*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v84 + int32(1)
																v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
																v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
																v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
																v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+28))
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(2))%32))))
																v108 = F_hashtableAdd(m, v107, v59)
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
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
						}
					}
				}
			}
		}
	}
}
func F_wcrtomb(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	v2 = l1
	if l0 == int32(0) {
		v93 = int32(1)
		return v93
	} else {
		if base.Ui32(v2) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
			return int32(1)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v12 != 0 {
				if base.Ui32(int32(2047)) < base.Ui32(v2) {
					if base.Ui32(v2) < base.Ui32(int32(55296)) {
						v40 = int32(63)
						v42 = int32(128)
						v43 = v2&v40 | v42
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
						v48 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
						v55 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v40 | v42
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v55)
						return int32(3)
					} else {
						if v2&int32(-8192) != int32(57344) {
							if base.Ui32(int32(1048575)) < base.Ui32(v2+int32(-65536)) {
								*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(25)
								v93 = int32(-1)
								return v93
							} else {
								v63 = int32(63)
								v65 = int32(128)
								v66 = v2&v63 | v65
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v66)
								v71 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v71)
								v78 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v63 | v65
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v78)
								v85 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v63 | v65
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v85)
								return int32(4)
							}
						} else {
							v40 = int32(63)
							v42 = int32(128)
							v43 = v2&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
							v48 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
							v55 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v55)
							return int32(3)
						}
					}
				} else {
					v25 = v2&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v25)
					v30 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v30)
					return int32(2)
				}
			} else {
				if v2&int32(-128) == int32(57216) {
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
					return int32(1)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(25)
					v93 = int32(-1)
					return v93
				}
			}
		}
	}
}
func F_wctomb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	v2 = l1
	if l0 != 0 {
		if l0 == int32(0) {
			v92 = int32(1)
			v96 = v92
		} else {
			if base.Ui32(v2) <= base.Ui32(int32(127)) {
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
				v96 = int32(1)
			} else {
				v12 = F___get_tp(m)
				mBase = m.M
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v14 != 0 {
					if base.Ui32(int32(2047)) < base.Ui32(v2) {
						if base.Ui32(v2) < base.Ui32(int32(55296)) {
							v41 = int32(63)
							v43 = int32(128)
							v44 = v2&v41 | v43
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v44)
							v49 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v49)
							v56 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v41 | v43
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v56)
							v96 = int32(3)
						} else {
							if v2&int32(-8192) != int32(57344) {
								if base.Ui32(int32(1048575)) < base.Ui32(v2+int32(-65536)) {
									v88 = F___errno_location(m)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(25)
									v92 = int32(-1)
									v96 = v92
								} else {
									v63 = int32(63)
									v65 = int32(128)
									v66 = v2&v63 | v65
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v66)
									v71 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
									*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v71)
									v78 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v63 | v65
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v78)
									v85 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v63 | v65
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v85)
									v96 = int32(4)
								}
							} else {
								v41 = int32(63)
								v43 = int32(128)
								v44 = v2&v41 | v43
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v44)
								v49 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v49)
								v56 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v41 | v43
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v56)
								v96 = int32(3)
							}
						}
					} else {
						v27 = v2&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v27)
						v32 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v32)
						v96 = int32(2)
					}
				} else {
					if v2&int32(-128) == int32(57216) {
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
						v96 = int32(1)
					} else {
						v19 = F___errno_location(m)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(25)
						v92 = int32(-1)
						v96 = v92
					}
				}
			}
		}
		return v96
	} else {
		return int32(0)
	}
}
func F_wctype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v9 = int32(_a2171)
	v10 = int32(97)
	v11 = int32(1)
	goto L2
L1:
	;
	return v59
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 != v10&int32(255) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v59 = int32(0)
	goto L1
L4:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+6)))
	v53 = v11 + int32(1)
	if v53 != int32(13) {
		v9 = v9 + int32(6)
		v10 = v49
		v11 = v53
		goto L2
	} else {
		goto L15
	}
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v19 == int32(0) {
		v42 = v18
		v43 = v19
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v43-v42&int32(255) == int32(0) {
		v59 = v11
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v19 != v18&int32(255) {
		v42 = v18
		v43 = v19
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v25 = l0
	v26 = v9
	goto L10
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v42 = v29
		v43 = v30
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v42 = v29
	v43 = v30
	goto L7
L12:
	;
	v33 = int32(1)
	if v30 == v29&int32(255) {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L4
L15:
	;
	goto L3
}
func F_writeCommandsDeniedByDiskError(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[903]))
	if v3 == v1 {
		v16 = int32(0)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if v18 == v16 {
			v35 = v16
		} else {
			v21 = int32(0)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
			if v23 == int32(-1) {
				v35 = int32(1)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[76]))
				if v27 != int32(-1) {
					v35 = v21
				} else {
					v30 = int32(0)
					v32 = *(*int32)(unsafe.Add(mBase, _consts[78]))
					*(*int32)(unsafe.Add(mBase, _consts[49])) = v32
					v35 = int32(1)
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[229]))
		if v7 < int32(1) {
			v16 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if v18 == v16 {
				v35 = v16
			} else {
				v21 = int32(0)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
				if v23 == int32(-1) {
					v35 = int32(1)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[76]))
					if v27 != int32(-1) {
						v35 = v21
					} else {
						v30 = int32(0)
						v32 = *(*int32)(unsafe.Add(mBase, _consts[78]))
						*(*int32)(unsafe.Add(mBase, _consts[49])) = v32
						v35 = int32(1)
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[847]))
			if v12 == int32(-1) {
				v35 = int32(2)
			} else {
				v16 = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if v18 == v16 {
					v35 = v16
				} else {
					v21 = int32(0)
					v23 = *(*int32)(unsafe.Add(mBase, _consts[46]))
					if v23 == int32(-1) {
						v35 = int32(1)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _consts[76]))
						if v27 != int32(-1) {
							v35 = v21
						} else {
							v30 = int32(0)
							v32 = *(*int32)(unsafe.Add(mBase, _consts[78]))
							*(*int32)(unsafe.Add(mBase, _consts[49])) = v32
							v35 = int32(1)
						}
					}
				}
			}
		}
	}
	return v35
}
func F_writeCommandsGetDiskErrorMessage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 != int32(2) {
		v16 = F_sdsempty(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[49]))
			v20 = F___strerror_l(m, v19, v19)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20
			v23 = F_sdscatfmt(m, v16, int32(_a1612), v5)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				m.G0 = v5 + int32(16)
				return v25
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[271]))
		v11 = F_objectGetVal(m, v10)
		mBase = m.M
		v12 = F_sdsdup(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v25 = v12
			m.G0 = v5 + int32(16)
			return v25
		}
	}
}
func F_writev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = m.Wasi_snapshot_preview1.Fd_write(m, l0, l1, l2, v7+int32(12))
	mBase = m.M
	if v11 != 0 {
		v13 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v11
	} else {
	}
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	m.G0 = v7 + int32(16)
	if v11 != 0 {
		v22 = int32(-1)
	} else {
		v22 = v17
	}
	return v22
}
func F_writevToClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int64
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int64
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v644 int64
	_ = v644
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int64
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v810 int64
	_ = v810
	var v814 int32
	_ = v814
	var v818 int64
	_ = v818
	var v819 int64
	_ = v819
	var v827 int32
	_ = v827
	var v836 int64
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int64
	_ = v854
	var v855 int64
	_ = v855
	var v860 int32
	_ = v860
	var v880 int32
	_ = v880
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+20)))
	v22 = int32(1024)
	if base.Ui32(v21) < base.Ui32(v22) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = v21
	goto L3
L2:
	;
	v25 = v22
	goto L3
L3:
	;
	v26 = int32(3)
	v32 = v18 - (v25<<(uint(v26)%32)+int32(15))&int32(32752)
	m.G0 = v32
	v35 = base.I32_div_u_s(v25, v26)
	v42 = v32 - (v35*int32(24)+int32(39))&int32(65520)
	m.G0 = v42
	v44 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+46)) = uint16(v44)
	v46 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	goto L6
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+v59)))
	v63 = int32(1)
	if v60 == int32(0) {
		v71 = v63
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v57 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if base.B2i32(v47 == v46) == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v59 = int32(180)
	v60 = v53
	goto L4
L8:
	;
	v58 = int32(180)
	goto L10
L9:
	;
	v58 = int32(140)
	goto L10
L10:
	;
	v59 = v58
	v60 = v57
	goto L4
L11:
	;
	if v71 < v25 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v71 = v68 + int32(1)
	goto L11
L13:
	;
	v73 = v71
	goto L15
L14:
	;
	v73 = v25
	goto L15
L15:
	;
	v80 = v42 - (v73*int32(24)+int32(15))&int32(-16)
	m.G0 = v80
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v32
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v18 + int32(46)
	if v62 == v82 {
		v251 = v82
		v252 = v82
		v253 = v63
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v60 == int32(0) {
		v508 = v251
		v509 = v252
		goto L46
	} else {
		goto L47
	}
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v102&int32(16777216) == int32(0) {
		v113 = v102
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = int64(0)
	if v113&int32(16777216) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_trackBufReferences(m, v107, v62, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v113 = v112
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v114
	v251 = int32(1)
	v252 = v242
	v253 = base.B2i32(v241 == int32(0))
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v231
	if v233 < int32(65537) {
		v241 = v232
		v242 = v233
		goto L22
	} else {
		goto L45
	}
L24:
	;
	if v21 != 0 {
		goto L42
	} else {
		goto L43
	}
L25:
	;
	v122 = v18 + int32(8)
	if v62 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v231 = base.B2i32(v203 == int32(0))
	v232 = v203
	v233 = v206
	goto L23
L27:
	;
	goto L26
L28:
	;
	v132 = v114
	goto L29
L29:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	if v142 != 0 {
		goto L27
	} else {
		goto L31
	}
L30:
	;
	goto L27
L31:
	;
	v144 = v132 + int32(12)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+10)))
	if v145&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v191 = v144 + v185
	if base.Ui32(v191) < base.Ui32(v114+v62) {
		v132 = v191
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	F_addBulkStringToReplyIOV(m, v144, v179, v122, v80)
	mBase = m.M
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
	v182 = v181 - v178
	*(*uint32)(unsafe.Add(mBase, uint32(v132)+4)) = uint32(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v185 = v184
	goto L32
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v149 != v150 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v154 + base.I64_extend_i32_u(v148)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if base.Ui32(v158) < base.Ui32(v148) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = int32(1)
	v185 = v148
	goto L32
L37:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v165 = v162 + v149<<(uint(int32(3))%32)
	v166 = v148 - v158
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v144 + v158
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v149 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v175 + v166
	v185 = v148
	goto L32
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v158 - v148
	v185 = v148
	goto L32
L39:
	;
	goto L30
L40:
	;
	v220 = v62 - v91
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v114 + v91
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v220
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v225
	v228 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v228
	v231 = v225
	v232 = v228
	v233 = v220
	goto L23
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = int32(1)
	v241 = v216
	v242 = int32(0)
	goto L22
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = base.I64_extend_i32_u(v62)
	if base.Ui32(v91) < base.Ui32(v62) {
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v207 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v207
	v216 = v207
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v91 - v62
	v216 = int32(0)
	goto L41
L45:
	;
	v237 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v237
	v241 = v237
	v242 = v233
	goto L22
L46:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+72))
	v523 = m.T0[v522].(func(*base.Module, int32, int32, int32) int32)(m, v518, v519, v520)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L20
	} else {
		goto L102
	}
L47:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v257
	goto L48
L48:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v262 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if base.B2i32(v262 != int32(0))&v253 != int32(1) {
		v508 = v251
		v509 = v252
		goto L46
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262+base.B2i32(v265 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v271
	goto L50
L52:
	;
	v281 = v262
	v284 = v251
	v285 = v252
	goto L53
L53:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v295 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	goto L55
L54:
	;
	v508 = v481
	v509 = v482
	goto L46
L55:
	;
	if v281 != v60 {
		v301 = v294
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if v301 != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	if v296 == v295 {
		v301 = v294
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v301 = v300
	goto L56
L59:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v487 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L60:
	;
	if v281 == v60 {
		v508 = v472
		v509 = v473
		goto L46
	} else {
		goto L92
	}
L61:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+12)))
	if v303&int32(1) == int32(0) {
		v313 = v303
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v469 = int32(0)
	v472 = v284
	v473 = v285
	goto L60
L63:
	;
	v316 = v80 + v284*int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v316)+8)) = int64(0)
	v320 = v293 + int32(13)
	if v313&int32(1) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	F_trackBufReferences(m, v293+int32(13), v301, l0)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+12)))
	v313 = v312
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+16)) = v444
	if v445 < int32(65537) {
		v454 = v443
		goto L87
	} else {
		goto L88
	}
L67:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v411 != v412 {
		goto L83
	} else {
		goto L84
	}
L68:
	;
	v326 = v18 + int32(8)
	if v301 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v443 = v407
	v444 = base.B2i32(v407 == int32(0))
	v445 = v410
	goto L66
L70:
	;
	goto L69
L71:
	;
	v336 = v320
	goto L72
L72:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v326)+20))
	if v346 != 0 {
		goto L70
	} else {
		goto L74
	}
L73:
	;
	goto L70
L74:
	;
	v348 = v336 + int32(12)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+10)))
	if v349&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v395 = v348 + v389
	if base.Ui32(v395) < base.Ui32(v320+v301) {
		v336 = v395
		goto L72
	} else {
		goto L82
	}
L76:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v316)+8))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	F_addBulkStringToReplyIOV(m, v348, v383, v326, v316)
	mBase = m.M
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v316)+8))
	v386 = v385 - v382
	*(*uint32)(unsafe.Add(mBase, uint32(v336)+4)) = uint32(v386)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v389 = v388
	goto L75
L77:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v353 != v354 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v316)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v316)+8)) = v358 + base.I64_extend_i32_u(v352)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v326)+16))
	if base.Ui32(v362) < base.Ui32(v352) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326)+20)) = int32(1)
	v389 = v352
	goto L75
L80:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v369 = v366 + v353<<(uint(int32(3))%32)
	v370 = v352 - v362
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v348 + v362
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v326)+16)) = int32(0)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+12)) = v379 + v370
	v389 = v352
	goto L75
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v326)+16)) = v362 - v352
	v389 = v352
	goto L75
L82:
	;
	goto L73
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v316)+8)) = base.I64_extend_i32_u(v301)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if base.Ui32(v420) < base.Ui32(v301) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v414 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v414
	v443 = v414
	v444 = v414
	v445 = v285
	goto L66
L85:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v429 = v426 + v411<<(uint(int32(3))%32)
	v430 = v301 - v420
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v320 + v420
	v434 = v285 + v430
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v434
	v436 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v411 + v436
	v440 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v440
	v443 = v440
	v444 = v436
	v445 = v434
	goto L66
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v420 - v301
	v443 = int32(0)
	v444 = int32(1)
	v445 = v285
	goto L66
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v320
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v316)+8))
	if v457 == int64(0) {
		v508 = v284
		v509 = v445
		goto L46
	} else {
		goto L89
	}
L88:
	;
	v451 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v451
	v454 = v451
	goto L87
L89:
	;
	v461 = v284 + int32(1)
	if v281 == v60 {
		v508 = v461
		v509 = v445
		goto L46
	} else {
		goto L90
	}
L90:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v463 != v464 {
		v469 = v454
		v472 = v461
		v473 = v445
		goto L60
	} else {
		goto L91
	}
L91:
	;
	v466 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v466
	v478 = v466
	v481 = v461
	v482 = v445
	goto L59
L92:
	;
	v478 = v469
	v481 = v472
	v482 = v473
	goto L59
L93:
	;
	if v487 == int32(0) {
		v508 = v481
		v509 = v482
		goto L46
	} else {
		goto L96
	}
L94:
	;
	goto L93
L95:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v487+base.B2i32(v490 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v496
	goto L94
L96:
	;
	if v478 == int32(0) {
		v281 = v487
		v284 = v481
		v285 = v482
		goto L53
	} else {
		goto L97
	}
L97:
	;
	goto L54
L98:
	;
	F__serverAssert(m, int32(_a1009), int32(_a977), int32(2728))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L20
	} else {
		goto L154
	}
L99:
	;
	m.G0 = v18 + int32(48)
	return v860
L100:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v760 != v765 {
		goto L140
	} else {
		goto L141
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v523
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)))
	v747 = v745 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v747)
	v860 = int32(-1)
	goto L99
L102:
	;
	if v523 <= int32(0) {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v528 = v523
	v530 = v519
	v536 = v520
	v537 = int32(0)
	goto L106
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v509
	if int32(1) <= v509 {
		v760 = v509
		goto L100
	} else {
		goto L139
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v542
	v760 = v542
	goto L100
L106:
	;
	v542 = v528 + v537
	if v542 == v509 {
		goto L104
	} else {
		goto L108
	}
L107:
	;
	v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)))
	v722 = v720 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+220)) = uint16(v722)
	goto L105
L108:
	;
	if base.Ui32(v542) < base.Ui32(int32(65537)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	if base.Ui32(v528) < base.Ui32(v648) {
		v675 = v528
		v677 = v530
		v679 = v648
		v683 = v536
		goto L131
	} else {
		goto L132
	}
L110:
	;
	v547 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if v547 == int64(0) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v550 = int32(0)
	v559 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v559 < int32(261) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v644 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	if base.Ui64(base.I64_extend_i32_u(v636)) <= base.Ui64(v644) {
		goto L105
	} else {
		goto L128
	}
L113:
	;
	goto L112
L114:
	;
	v570 = v568 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v568) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	if v559 < int32(1) {
		v636 = v550
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	v567 = v563
	v568 = int32(260)
	goto L114
L117:
	;
	v567 = v550
	v568 = v559
	goto L114
L118:
	;
	if v570 == int32(0) {
		v636 = v609
		goto L113
	} else {
		goto L124
	}
L119:
	;
	v577 = int32(0)
	v579 = v567
	v580 = v577
	v584 = v577
	goto L121
L120:
	;
	v609 = v567
	v610 = int32(0)
	goto L118
L121:
	;
	v587 = v580 << (uint(int32(2)) % 32)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v587)+uint32(_consts[282])))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v587)+uint32(_consts[283])))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v587)+uint32(_consts[284])))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v587)+uint32(_consts[285])))
	v603 = v590 + (v593 + (v596 + (v599 + v579)))
	v604 = int32(4)
	v605 = v580 + v604
	v607 = v584 + v604
	if v607 != v568&int32(2147483644) {
		v579 = v603
		v580 = v605
		v584 = v607
		goto L121
	} else {
		goto L123
	}
L122:
	;
	v609 = v603
	v610 = v605
	goto L118
L123:
	;
	goto L122
L124:
	;
	v618 = v609
	v619 = v610
	v621 = int32(0)
	goto L125
L125:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v619<<(uint(int32(2))%32))+uint32(_consts[285])))
	v630 = v629 + v618
	v631 = int32(1)
	v634 = v621 + v631
	if v634 != v570 {
		v618 = v630
		v619 = v619 + v631
		v621 = v634
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v636 = v630
	goto L113
L127:
	;
	goto L126
L128:
	;
	goto L109
L129:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+72))
	v716 = m.T0[v715].(func(*base.Module, int32, int32, int32) int32)(m, v713, v701, v707)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L20
	} else {
		goto L137
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v668
	v701 = v668
	v707 = v666
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v683
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v677
	*(*int32)(unsafe.Add(mBase, uint32(v677)+4)) = v679 - v675
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	*(*int32)(unsafe.Add(mBase, uint32(v677))) = v693 + v675
	v701 = v677
	v707 = v683
	goto L129
L132:
	;
	v651 = v528
	v653 = v530
	v655 = v648
	v659 = v536
	goto L133
L133:
	;
	v666 = v659 + int32(-1)
	v668 = v653 + int32(8)
	v669 = v651 - v655
	if v669 == int32(0) {
		goto L130
	} else {
		goto L135
	}
L134:
	;
	v675 = v669
	v677 = v668
	v679 = v672
	v683 = v666
	goto L131
L135:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v653)+12))
	if base.Ui32(v672) <= base.Ui32(v669) {
		v651 = v669
		v653 = v668
		v655 = v672
		v659 = v666
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	if int32(1) <= v716 {
		v528 = v716
		v530 = v701
		v536 = v707
		v537 = v542
		goto L106
	} else {
		goto L138
	}
L138:
	;
	goto L107
L139:
	;
	v860 = int32(-1)
	goto L99
L140:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v791 = v790 + v760
	if v791 != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v769 = v80 + v508*int32(24)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v769+int32(-24))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v772
	v774 = int32(0)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v769+int32(-8))))
	if v778 == v774 {
		v784 = v774
		goto L142
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v784
	v788 = *(*int64)(unsafe.Add(mBase, uint32(v769+int32(-16))))
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+152)) = uint32(v788)
	v860 = v774
	goto L99
L143:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v769+int32(-20))))
	v784 = v783
	goto L142
L144:
	;
	if v508 <= v827 {
		goto L98
	} else {
		goto L150
	}
L145:
	;
	v799 = int32(-1)
	v810 = base.I64_extend_i32_u(v791)
	goto L147
L146:
	;
	v827 = int32(-1)
	v836 = int64(0)
	goto L144
L147:
	;
	v814 = v799 + int32(1)
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v799*int32(24)+v80+int32(32))))
	v819 = v810 - v818
	if int64(0) < v819 {
		v799 = v814
		v810 = v819
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v827 = v814
	v836 = v819
	goto L144
L149:
	;
	goto L148
L150:
	;
	v840 = v80 + v827*int32(24)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v841
	v843 = int32(0)
	if v836 != int64(0) {
		v852 = v843
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v852
	v854 = *(*int64)(unsafe.Add(mBase, uint32(v840)+8))
	v855 = v854 + v836
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+152)) = uint32(v855)
	v860 = v843
	goto L99
L152:
	;
	v847 = int32(0)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v840)+16))
	if v848 == v847 {
		v852 = v847
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	v852 = v851
	goto L151
L154:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
