package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_dbAdd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_dbAddInternal(m, l0, l1, l2, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_dbAsyncDelete(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = F_objectGetVal(m, l1)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v6 != 0 {
		v8 = F_getKeySlot(m, v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = v8
			v13 = int32(1)
			v15 = F_dbGenericDeleteWithDictIndex(m, l0, l1, v13, v13, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	} else {
		v12 = int32(0)
		v13 = int32(1)
		v15 = F_dbGenericDeleteWithDictIndex(m, l0, l1, v13, v13, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_dbExpand(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v7 == int32(0) {
		v35 = F_kvstoreExpand(m, v5, l1, l2, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = v35
			return v37 + int32(-1)
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[136]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)))
		if v14&int32(2) == int32(0) {
			v21 = v13
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2160))
			v24 = v22
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+2172))
			if v19 != 0 {
				v21 = v19
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+2160))
				v24 = v22
			} else {
				v24 = int32(0)
			}
		}
		if v24 != 0 {
			v28 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v24))
			v30 = F_kvstoreExpand(m, v5, v28, l2, int32(517))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v37 = v30
				return v37 + int32(-1)
			}
		} else {
			return int32(0)
		}
	}
}
func F_dbFindExpiresWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_kvstoreHashtableFind(m, v11, l2, l1, v7+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		m.G0 = v7 + int32(16)
		return v18
	}
}
func F_dbGenericDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v6 = F_objectGetVal(m, l1)
	mBase = m.M
	v8 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v8 != 0 {
		v15 = F_getKeySlot(m, v6)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_dbGenericDeleteWithDictIndex(m, l0, l1, l2, l3, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v17
			}
		}
	} else {
		v10 = F_dbGenericDeleteWithDictIndex(m, l0, l1, l2, l3, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_dbGenericDeleteWithDictIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_objectGetVal(m, l1)
	mBase = m.M
	v15 = F_kvstoreHashtableTwoPhasePopFindRef(m, v13, l4, v14, v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			F_incrRefCount(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				F_moduleNotifyKeyUnlink(m, l1, v20, v23, l3)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					F_signalDeletedKeyAsReady(m, l0, l1, v26&int32(15))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_decrRefCount(m, v20)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_kvstoreHashtableTwoPhasePopDelete(m, v34, l4, v11)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
								if v40&int32(1) == int32(0) {
									v51 = int64(-1)
								} else {
									v50 = *(*int64)(unsafe.Add(mBase, uint32(v33+(v40&int32(4)^int32(12)))))
									v51 = v50
								}
								if v51 == int64(-1) {
									v65 = *(*int32)(unsafe.Add(mBase, _consts[88]))
									if v65 == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
										if v72&int32(15) != int32(4) {
											if l2 == int32(0) {
												F_decrRefCount(m, v33)
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													v147 = int32(1)
													m.G0 = v11 + int32(16)
													return v147
												}
											} else {
												v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_freeObjAsync(m, l1, v33, v141)
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return int32(0)
												} else {
													v147 = int32(1)
													m.G0 = v11 + int32(16)
													return v147
												}
											}
										} else {
											v77 = F_hashTypeHasVolatileFields(m, v33)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v77 == int32(0) {
													if l2 == int32(0) {
														F_decrRefCount(m, v33)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															v147 = int32(1)
															m.G0 = v11 + int32(16)
															return v147
														}
													} else {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														F_freeObjAsync(m, l1, v33, v141)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															v147 = int32(1)
															m.G0 = v11 + int32(16)
															return v147
														}
													}
												} else {
													v81 = int32(0)
													v84 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
													if v84&int32(2) == v81 {
														v104 = v81
													} else {
														v98 = v33 + (v84&int32(4) ^ int32(12)) + v84<<(uint(int32(3))%32)&int32(8)
														v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
														v104 = v98 + v99 + int32(1)
													}
													v106 = *(*int32)(unsafe.Add(mBase, _consts[139]))
													if v106 != 0 {
														v108 = F_getKeySlot(m, v104)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int32(0)
														} else {
															v110 = v108
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v112 = int32(0)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
															if v115&int32(2) == v112 {
																v135 = v112
															} else {
																v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
																v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
																v135 = v129 + v130 + int32(1)
															}
															v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																if l2 == int32(0) {
																	F_decrRefCount(m, v33)
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int32(0)
																	} else {
																		v147 = int32(1)
																		m.G0 = v11 + int32(16)
																		return v147
																	}
																} else {
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	F_freeObjAsync(m, l1, v33, v141)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		v147 = int32(1)
																		m.G0 = v11 + int32(16)
																		return v147
																	}
																}
															}
														}
													} else {
														v110 = int32(0)
														v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v112 = int32(0)
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
														if v115&int32(2) == v112 {
															v135 = v112
														} else {
															v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
															v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
															v135 = v129 + v130 + int32(1)
														}
														v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															if l2 == int32(0) {
																F_decrRefCount(m, v33)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int32(0)
																} else {
																	v147 = int32(1)
																	m.G0 = v11 + int32(16)
																	return v147
																}
															} else {
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																F_freeObjAsync(m, l1, v33, v141)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	v147 = int32(1)
																	m.G0 = v11 + int32(16)
																	return v147
																}
															}
														}
													}
												}
											}
										}
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v69 = F_objectGetVal(m, l1)
										mBase = m.M
										v70 = F_kvstoreHashtableDelete(m, v68, l4, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											if v70 != 0 {
												F__serverAssert(m, int32(_a554), int32(_a550), int32(499))
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
												if v72&int32(15) != int32(4) {
													if l2 == int32(0) {
														F_decrRefCount(m, v33)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															v147 = int32(1)
															m.G0 = v11 + int32(16)
															return v147
														}
													} else {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														F_freeObjAsync(m, l1, v33, v141)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															v147 = int32(1)
															m.G0 = v11 + int32(16)
															return v147
														}
													}
												} else {
													v77 = F_hashTypeHasVolatileFields(m, v33)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														if v77 == int32(0) {
															if l2 == int32(0) {
																F_decrRefCount(m, v33)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return int32(0)
																} else {
																	v147 = int32(1)
																	m.G0 = v11 + int32(16)
																	return v147
																}
															} else {
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																F_freeObjAsync(m, l1, v33, v141)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	v147 = int32(1)
																	m.G0 = v11 + int32(16)
																	return v147
																}
															}
														} else {
															v81 = int32(0)
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
															if v84&int32(2) == v81 {
																v104 = v81
															} else {
																v98 = v33 + (v84&int32(4) ^ int32(12)) + v84<<(uint(int32(3))%32)&int32(8)
																v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
																v104 = v98 + v99 + int32(1)
															}
															v106 = *(*int32)(unsafe.Add(mBase, _consts[139]))
															if v106 != 0 {
																v108 = F_getKeySlot(m, v104)
																mBase = m.M
																v109 = m.ExcPending
																if v109 != 0 {
																	return int32(0)
																} else {
																	v110 = v108
																	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	v112 = int32(0)
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
																	if v115&int32(2) == v112 {
																		v135 = v112
																	} else {
																		v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
																		v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
																		v135 = v129 + v130 + int32(1)
																	}
																	v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return int32(0)
																	} else {
																		if l2 == int32(0) {
																			F_decrRefCount(m, v33)
																			mBase = m.M
																			v145 = m.ExcPending
																			if v145 != 0 {
																				return int32(0)
																			} else {
																				v147 = int32(1)
																				m.G0 = v11 + int32(16)
																				return v147
																			}
																		} else {
																			v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																			F_freeObjAsync(m, l1, v33, v141)
																			mBase = m.M
																			v143 = m.ExcPending
																			if v143 != 0 {
																				return int32(0)
																			} else {
																				v147 = int32(1)
																				m.G0 = v11 + int32(16)
																				return v147
																			}
																		}
																	}
																}
															} else {
																v110 = int32(0)
																v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v112 = int32(0)
																v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
																if v115&int32(2) == v112 {
																	v135 = v112
																} else {
																	v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
																	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
																	v135 = v129 + v130 + int32(1)
																}
																v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	if l2 == int32(0) {
																		F_decrRefCount(m, v33)
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return int32(0)
																		} else {
																			v147 = int32(1)
																			m.G0 = v11 + int32(16)
																			return v147
																		}
																	} else {
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		F_freeObjAsync(m, l1, v33, v141)
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			v147 = int32(1)
																			m.G0 = v11 + int32(16)
																			return v147
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
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v55 = F_objectGetVal(m, l1)
									mBase = m.M
									v56 = F_kvstoreHashtableDelete(m, v54, l4, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
											if v72&int32(15) != int32(4) {
												if l2 == int32(0) {
													F_decrRefCount(m, v33)
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int32(0)
													} else {
														v147 = int32(1)
														m.G0 = v11 + int32(16)
														return v147
													}
												} else {
													v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													F_freeObjAsync(m, l1, v33, v141)
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														v147 = int32(1)
														m.G0 = v11 + int32(16)
														return v147
													}
												}
											} else {
												v77 = F_hashTypeHasVolatileFields(m, v33)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if v77 == int32(0) {
														if l2 == int32(0) {
															F_decrRefCount(m, v33)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int32(0)
															} else {
																v147 = int32(1)
																m.G0 = v11 + int32(16)
																return v147
															}
														} else {
															v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															F_freeObjAsync(m, l1, v33, v141)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																v147 = int32(1)
																m.G0 = v11 + int32(16)
																return v147
															}
														}
													} else {
														v81 = int32(0)
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
														if v84&int32(2) == v81 {
															v104 = v81
														} else {
															v98 = v33 + (v84&int32(4) ^ int32(12)) + v84<<(uint(int32(3))%32)&int32(8)
															v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
															v104 = v98 + v99 + int32(1)
														}
														v106 = *(*int32)(unsafe.Add(mBase, _consts[139]))
														if v106 != 0 {
															v108 = F_getKeySlot(m, v104)
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return int32(0)
															} else {
																v110 = v108
																v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																v112 = int32(0)
																v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
																if v115&int32(2) == v112 {
																	v135 = v112
																} else {
																	v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
																	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
																	v135 = v129 + v130 + int32(1)
																}
																v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	if l2 == int32(0) {
																		F_decrRefCount(m, v33)
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return int32(0)
																		} else {
																			v147 = int32(1)
																			m.G0 = v11 + int32(16)
																			return v147
																		}
																	} else {
																		v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		F_freeObjAsync(m, l1, v33, v141)
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			v147 = int32(1)
																			m.G0 = v11 + int32(16)
																			return v147
																		}
																	}
																}
															}
														} else {
															v110 = int32(0)
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v112 = int32(0)
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
															if v115&int32(2) == v112 {
																v135 = v112
															} else {
																v129 = v33 + (v115&int32(4) ^ int32(12)) + v115<<(uint(int32(3))%32)&int32(8)
																v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
																v135 = v129 + v130 + int32(1)
															}
															v136 = F_kvstoreHashtableDelete(m, v111, v110, v135)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																if l2 == int32(0) {
																	F_decrRefCount(m, v33)
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int32(0)
																	} else {
																		v147 = int32(1)
																		m.G0 = v11 + int32(16)
																		return v147
																	}
																} else {
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	F_freeObjAsync(m, l1, v33, v141)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		v147 = int32(1)
																		m.G0 = v11 + int32(16)
																		return v147
																	}
																}
															}
														}
													}
												}
											}
										} else {
											F__serverAssert(m, int32(_a555), int32(_a550), int32(497))
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
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v147 = int32(0)
			m.G0 = v11 + int32(16)
			return v147
		}
	}
}
func F_dbRandomKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v18 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v30 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	v28 = v21
	goto L1
L4:
	;
	v25 = F_hashtableSize(m, v23)
	mBase = m.M
	v28 = base.I64_extend_i32_u(v25)
	goto L1
L5:
	;
	v28 = int64(0)
	goto L1
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = F_kvstoreGetFairRandomHashtableIndex(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	v40 = v33
	goto L6
L9:
	;
	v37 = F_hashtableSize(m, v35)
	mBase = m.M
	v40 = base.I64_extend_i32_u(v37)
	goto L6
L10:
	;
	v40 = int64(0)
	goto L6
L11:
	;
	m.G0 = v15 + int32(16)
	return v197
L12:
	;
	v197 = int32(0)
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	if v42 == int32(-1) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v54 = v42
	v55 = int32(100)
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = F_kvstoreHashtableFairRandomEntry(m, v62, v54, v15+int32(12))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L18
	}
L17:
	;
	goto L12
L18:
	;
	if v65 == int32(0) {
		v197 = int32(0)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74&int32(2) == v69 {
		v94 = v69
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v115 = F_createStringObject_1(m, v94, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L29
	}
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-17))))
	v114 = v113
	goto L20
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-9))))
	v114 = v110
	goto L20
L23:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-5)))))
	v114 = v107
	goto L20
L24:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-3)))))
	v114 = v104
	goto L20
L25:
	;
	v114 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-1)))))
	switch v97 & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	default:
		v114 = v69
		goto L20
	}
L27:
	;
	goto L26
L28:
	;
	v88 = v70 + (v74&int32(4) ^ int32(12)) + v74<<(uint(int32(3))%32)&int32(8)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v94 = v88 + v89 + int32(1)
	goto L27
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v118 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v124&int32(1) == int32(0) {
		v135 = int64(-1)
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v197 = v115
	goto L11
L32:
	;
	v142 = int32(_a44)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v145 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v145 != 0 {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	if int64(0) <= v135 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L33
L35:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v70+(v124&int32(4)^int32(12)))))
	v135 = v134
	goto L34
L36:
	;
	if v141 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	v139 = F_commandTimeSnapshot(m)
	mBase = m.M
	v141 = base.B2i32(v135 < v139)
	goto L36
L38:
	;
	v141 = int32(0)
	goto L36
L39:
	;
	v197 = v115
	goto L11
L40:
	;
	if v28 != v40 {
		v167 = v55
		goto L45
	} else {
		goto L46
	}
L41:
	;
	if v143 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v149 == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+207)))
	if v152&int32(32) == int32(0) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v197 = v115
	goto L11
L45:
	;
	v169 = F_expireIfNeededWithDictIndex(m, l0, v115, v70, int32(0), v54)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L54
	}
L46:
	;
	if v145 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v55 + int32(-1)
	if v166 != 0 {
		v167 = v166
		goto L45
	} else {
		goto L52
	}
L48:
	;
	if v143 != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	goto L50
L50:
	;
	if v161&int32(4) == int32(0) {
		v167 = v55
		goto L45
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v197 = v115
	goto L11
L53:
	;
	F_decrRefCount(m, v115)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L13
	} else {
		goto L56
	}
L54:
	;
	if v169 != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v197 = v115
	goto L11
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v174 = F_kvstoreGetFairRandomHashtableIndex(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	if v174 != int32(-1) {
		v54 = v174
		v55 = v167
		goto L16
	} else {
		goto L58
	}
L58:
	;
	goto L17
}
func F_dbReclaimExpiredFields(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v144 int64
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
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
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	v14 = m.G0
	v16 = v14 - int32(8208)
	m.G0 = v16
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(8208)
	return v387
L2:
	;
	v27 = l3
	v30 = int32(0)
	goto L4
L3:
	;
	v387 = int32(0)
	goto L1
L4:
	;
	v37 = int32(1024)
	if base.Ui32(v27) < base.Ui32(v37) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v387 = v377
	goto L1
L6:
	;
	v40 = v27
	goto L8
L7:
	;
	v40 = v37
	goto L8
L8:
	;
	v41 = F_hashTypeDeleteExpiredFields(m, l0, l2, v40, v16)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v41 == int32(0) {
		v387 = v30
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v47 = F_hashTypeHasVolatileFields(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v107 = F_hashTypeLength(m, l0)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L26
	}
L13:
	;
	if v47 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v49 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52&int32(2) == v49 {
		v72 = v49
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v74 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L15
L17:
	;
	v66 = l0 + (v52&int32(4) ^ int32(12)) + v52<<(uint(int32(3))%32)&int32(8)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v72 = v66 + v67 + int32(1)
	goto L16
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v80 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83&int32(2) == v80 {
		v103 = v80
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v76 = F_getKeySlot(m, v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	v78 = int32(0)
	goto L18
L21:
	;
	v78 = v76
	goto L18
L22:
	;
	v104 = F_kvstoreHashtableDelete(m, v79, v78, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v97 = l0 + (v83&int32(4) ^ int32(12)) + v83<<(uint(int32(3))%32)&int32(8)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v103 = v97 + v98 + int32(1)
	goto L23
L25:
	;
	goto L12
L26:
	;
	v111 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v115 + int32(1)
	goto L29
L27:
	;
	v148 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v151&int32(2) == v148 {
		v171 = v148
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	if v115 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	goto L32
L31:
	;
	v125 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[96])) = v123
	v129 = base.I64_div_s(v123, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[35])) = v129
	v133 = base.I64_div_s(v123, int64(1000000))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v133
	v136 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	F_lrulfu_updateClockAndPolicy(m, v129, int32(base.Ui32(v136&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v144 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, _consts[98])) = v144
	goto L28
L32:
	;
	v123 = F_ustime(m)
	mBase = m.M
	goto L31
L33:
	;
	v172 = F_createStringObjectFromSds(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v165 = l0 + (v151&int32(4) ^ int32(12)) + v151<<(uint(int32(3))%32)&int32(8)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v171 = v165 + v166 + int32(1)
	goto L34
L36:
	;
	v174 = int32(_a44)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	*(*int32)(unsafe.Add(mBase, _consts[349])) = int32(1)
	v179 = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v182&int32(2) == v179 {
		v202 = v179
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v203 = F_createStringObjectFromSds(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v196 = l0 + (v182&int32(4) ^ int32(12)) + v182<<(uint(int32(3))%32)&int32(8)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v202 = v196 + v197 + int32(1)
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[352]))) = v203
	v207 = *(*int32)(unsafe.Add(mBase, _consts[353]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[354]))) = v207
	v209 = int32(1024)
	if base.Ui32(v41) < base.Ui32(v209) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v212 = v41
	goto L43
L42:
	;
	v212 = v209
	goto L43
L43:
	;
	v214 = v212 << (uint(int32(2)) % 32)
	if v214 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v223 = v212 + int32(2)
	F_alsoPropagate(m, v219, v16+int32(4096), v223, int32(3), l4)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v217 = F__emscripten_memcpy_bulkmem(m, v16+int32(4096)|int32(8), v16, v214)
	mBase = m.M
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v175
	v239 = int32(0)
	goto L48
L48:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(4096)+v239<<(uint(int32(2))%32))))
	F_decrRefCount(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L50
	}
L49:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_notifyKeyspaceEvent(m, int32(256), int32(_a584), v172, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	v252 = v239 + int32(1)
	if v252 != v223 {
		v239 = v252
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v107 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_touchWatchedKey(m, l1, v172)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L79
	}
L54:
	;
	v300 = F_hashTypeHasVolatileFields(m, l0)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L9
	} else {
		goto L66
	}
L55:
	;
	v259 = int32(_a44)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v261 = F_objectGetVal(m, v172)
	mBase = m.M
	v263 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v263 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v269 = F_dbGenericDeleteWithDictIndex(m, l1, v172, v260, int32(1), v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L60
	}
L57:
	;
	v265 = F_getKeySlot(m, v261)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L59
	}
L58:
	;
	v267 = int32(0)
	goto L56
L59:
	;
	v267 = v265
	goto L56
L60:
	;
	v271 = int32(_a44)
	v272 = *(*int32)(unsafe.Add(mBase, _consts[349]))
	*(*int32)(unsafe.Add(mBase, _consts[349])) = int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[352]))) = v172
	if v277 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v282 = int32(244)
	goto L63
L62:
	;
	v282 = int32(240)
	goto L63
L63:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[84])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[354]))) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_alsoPropagate(m, v286, v16+int32(4096), int32(2), int32(3), l4)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[349])) = v272
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a213), v172, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	goto L53
L66:
	;
	if v300 != 0 {
		goto L53
	} else {
		goto L67
	}
L67:
	;
	v302 = int32(0)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v305&int32(2) == v302 {
		v325 = v302
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v327 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	goto L68
L70:
	;
	v319 = l0 + (v305&int32(4) ^ int32(12)) + v305<<(uint(int32(3))%32)&int32(8)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v325 = v319 + v320 + int32(1)
	goto L69
L71:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v333 = int32(0)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v336&int32(2) == v333 {
		v356 = v333
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v329 = F_getKeySlot(m, v325)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L9
	} else {
		goto L74
	}
L73:
	;
	v331 = int32(0)
	goto L71
L74:
	;
	v331 = v329
	goto L71
L75:
	;
	v357 = F_kvstoreHashtableDelete(m, v332, v331, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v350 = l0 + (v336&int32(4) ^ int32(12)) + v336<<(uint(int32(3))%32)&int32(8)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	v356 = v350 + v351 + int32(1)
	goto L76
L78:
	;
	goto L53
L79:
	;
	F_trackingInvalidateKey(m, int32(0), v172, int32(1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v367 = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, _consts[95])) = v369 + int32(-1)
	goto L81
L81:
	;
	F_postExecutionUnitOperations(m)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	F_decrRefCount(m, v172)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	v377 = v41 + v30
	if v107 == int32(0) {
		v387 = v377
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v380 = v27 - v41
	if v380 != 0 {
		v27 = v380
		v30 = v377
		goto L4
	} else {
		goto L85
	}
L85:
	;
	goto L5
}
func F_dbUpdateObjectWithVolatileItemsTracking(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4&int32(15) != int32(4) {
		return
	} else {
		v9 = F_hashTypeHasVolatileFields(m, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if v9 == int32(0) {
				v55 = int32(0)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v58&int32(2) == v55 {
					v78 = v55
				} else {
					v72 = l1 + (v58&int32(4) ^ int32(12)) + v58<<(uint(int32(3))%32)&int32(8)
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
					v78 = v72 + v73 + int32(1)
				}
				v80 = *(*int32)(unsafe.Add(mBase, _consts[139]))
				if v80 != 0 {
					v82 = F_getKeySlot(m, v78)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v84 = v82
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v86 = int32(0)
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v89&int32(2) == v86 {
							v109 = v86
						} else {
							v103 = l1 + (v89&int32(4) ^ int32(12)) + v89<<(uint(int32(3))%32)&int32(8)
							v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
							v109 = v103 + v104 + int32(1)
						}
						v110 = F_kvstoreHashtableDelete(m, v85, v84, v109)
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v84 = int32(0)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v86 = int32(0)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v89&int32(2) == v86 {
						v109 = v86
					} else {
						v103 = l1 + (v89&int32(4) ^ int32(12)) + v89<<(uint(int32(3))%32)&int32(8)
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
						v109 = v103 + v104 + int32(1)
					}
					v110 = F_kvstoreHashtableDelete(m, v85, v84, v109)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v13&int32(15) != int32(4) {
					return
				} else {
					v18 = F_hashTypeHasVolatileFields(m, l1)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						if v18 == int32(0) {
							return
						} else {
							v22 = int32(0)
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							if v25&int32(2) == v22 {
								v45 = v22
							} else {
								v39 = l1 + (v25&int32(4) ^ int32(12)) + v25<<(uint(int32(3))%32)&int32(8)
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
								v45 = v39 + v40 + int32(1)
							}
							v47 = *(*int32)(unsafe.Add(mBase, _consts[139]))
							if v47 != 0 {
								v49 = F_getKeySlot(m, v45)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = v49
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v53 = F_kvstoreHashtableAdd(m, v52, v51, l1)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								v51 = int32(0)
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v53 = F_kvstoreHashtableAdd(m, v52, v51, l1)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
func F_db_debug(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
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
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v10 = m.G3
	v15 = m.G397
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_fwrite(m, v10+int32(_a2079), int32(11), int32(1), v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = m.G402
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = F_fgets(m, v8, int32(250), v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(256)
	return int32(0)
L4:
	;
	if v24 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	goto L6
L6:
	;
	v33 = m.G3
	v35 = v33 + int32(_a2080)
	v36 = int32(6)
	goto L12
L7:
	;
	goto L3
L8:
	;
	if v100 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L9:
	;
	v100 = int32(0)
	goto L8
L10:
	;
	v72 = v67
	v73 = v68
	v74 = v69
	goto L20
L11:
	;
	if v57 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	if (v35|v8)&int32(3) != 0 {
		v67 = v8
		v68 = v35
		v69 = v36
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v44 = v8
	v45 = v35
	v46 = v36
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v49 != v50 {
		v67 = v44
		v68 = v45
		v69 = v46
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v52 = int32(4)
	v53 = v45 + v52
	v55 = v44 + v52
	v57 = v46 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v57) {
		v44 = v55
		v45 = v53
		v46 = v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v67 = v55
	v68 = v53
	v69 = v57
	goto L10
L19:
	;
	v100 = v77 - v78
	goto L8
L20:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != v78 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v80 = int32(1)
	v85 = v74 + int32(-1)
	if v85 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v72 = v72 + v80
	v73 = v73 + v80
	v74 = v85
	goto L20
L24:
	;
	v103 = m.G3
	if v8&int32(3) == int32(0) {
		v125 = v8
		goto L29
	} else {
		goto L30
	}
L25:
	;
	goto L53
L26:
	;
	v172 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L47
	}
L27:
	;
	v161 = F_luaL_loadbuffer(m, l0, v8, v158, v103+int32(_a2081))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L43
	}
L28:
	;
	v158 = v150 - v8
	goto L27
L29:
	;
	v129 = v125
	goto L37
L30:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = v8
	goto L33
L32:
	;
	v158 = v8 - v8
	goto L27
L33:
	;
	v118 = v114 + int32(1)
	if v118&int32(3) == int32(0) {
		v125 = v118
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v123 != 0 {
		v114 = v118
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v150 = v118
	goto L28
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v138 = int32(-2139062144)
	if (int32(16843008)-v135|v135)&v138 == v138 {
		v129 = v129 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v144 = v129
	goto L40
L39:
	;
	goto L38
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 != 0 {
		v144 = v144 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v150 = v144
	goto L28
L42:
	;
	goto L41
L43:
	;
	if v161 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v163 = int32(0)
	v166 = F_lua_pcall(m, l0, v163, v163, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v166 == int32(0) {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	goto L26
L47:
	;
	v174 = F_fputs(m, v172, v16)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v177 = F_fputc(m, int32(10), v16)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L25
L50:
	;
	v207 = m.G3
	v212 = F_fwrite(m, v207+int32(_a2079), int32(11), int32(1), v16)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v187
	goto L50
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v187 = v184 + int32(0)
	if base.Ui32(v187) <= base.Ui32(v183) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v191 = v183
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = int32(0)
	v195 = v191 + int32(16)
	if base.Ui32(v195) < base.Ui32(v187) {
		v191 = v195
		goto L55
	} else {
		goto L57
	}
L57:
	;
	goto L51
L58:
	;
	v215 = F_fgets(m, v8, int32(250), v23)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v215 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	goto L7
}
func F_db_getinfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v356 int32
	_ = v356
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v433 int32
	_ = v433
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v676 int32
	_ = v676
	var v677 int64
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v702 int32
	_ = v702
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var __phi746 int32
	_ = __phi746
	var v747 int32
	_ = v747
	var __phi747 int32
	_ = __phi747
	var v749 int64
	_ = v749
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v772 int32
	_ = v772
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v847 int64
	_ = v847
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v934 int32
	_ = v934
	var v935 int64
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v960 int32
	_ = v960
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var __phi1004 int32
	_ = __phi1004
	var v1005 int32
	_ = v1005
	var __phi1005 int32
	_ = __phi1005
	var v1007 int64
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1030 int32
	_ = v1030
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	goto L5
L1:
	;
	v137 = v134 | int32(2)
	v138 = m.G3
	v142 = F_luaL_optlstring(m, l0, v137, v138+int32(_a2082), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	if v69 != int32(8) {
		v134 = int32(0)
		v135 = l0
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v63 = m.G398
	if v20 != v63 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = v15 + int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v20) < base.Ui32(v21) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v69 = int32(-1)
	goto L2
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v69 = v66
	goto L2
L16:
	;
	v69 = int32(-1)
	goto L2
L17:
	;
	goto L21
L18:
	;
	v134 = int32(1)
	v135 = v132
	goto L1
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if v128 != int32(8) {
		v132 = int32(0)
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v82 = v77 + int32(0)
	v83 = m.G398
	if base.Ui32(v82) < base.Ui32(v76) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = v82
	goto L24
L23:
	;
	v85 = v83
	goto L24
L24:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v132 = v131
	goto L34
L36:
	;
	return int32(0)
L37:
	;
	v147 = v134 + int32(1)
	v148 = F_lua_isnumber(m, l0, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L36
	} else {
		goto L42
	}
L38:
	;
	m.G0 = v9 + int32(112)
	return v1129
L39:
	;
	v1124 = m.G3
	v1127 = F_luaL_argerror(m, l0, v147, v1124+int32(_a2083))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L36
	} else {
		goto L264
	}
L40:
	;
	v447 = F_lua_getinfo(m, v135, v444, v9+int32(12))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L36
	} else {
		goto L105
	}
L41:
	;
	if v147 < int32(1) {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	if v148 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v152 = F_lua_tointeger(m, l0, v147)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	if v152 < int32(1) {
		v184 = v152
		v186 = v159
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v206 != 0 {
		v444 = v142
		goto L40
	} else {
		goto L58
	}
L46:
	;
	goto L45
L47:
	;
	if v184 != 0 {
		v197 = int32(0)
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	v165 = v152
	v167 = v159
	goto L49
L49:
	;
	if base.Ui32(v167) <= base.Ui32(v162) {
		v206 = int32(0)
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v184 = v178
	v186 = v180
	goto L47
L51:
	;
	v172 = v165 + int32(-1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
	if v175 != 0 {
		v178 = v172
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v180 = v167 + int32(-24)
	if int32(0) < v178 {
		v165 = v178
		v167 = v180
		goto L49
	} else {
		goto L54
	}
L53:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v178 = v172 - v176
	goto L52
L54:
	;
	goto L50
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12))+96)) = v197
	v206 = int32(1)
	goto L46
L56:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	if base.Ui32(v186) <= base.Ui32(v191) {
		v206 = int32(0)
		goto L46
	} else {
		goto L57
	}
L57:
	;
	v195 = base.I32_div_s(v186-v191, int32(24))
	v197 = v195
	goto L55
L58:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v208 + int32(16)
	goto L59
L59:
	;
	v1129 = int32(1)
	goto L38
L60:
	;
	if v271 != int32(6) {
		goto L39
	} else {
		goto L75
	}
L61:
	;
	v265 = m.G398
	if v264 != v265 {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if v147 < int32(-9999) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v222 = v217 + v147<<(uint(int32(4))%32) + int32(-16)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v222) < base.Ui32(v223) {
		v264 = v222
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v271 = int32(-1)
	goto L60
L65:
	;
	switch v134 + int32(10003) {
	case 0:
		goto L69
	case 1:
		goto L70
	case 2:
		goto L67
	default:
		goto L68
	}
L66:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v264 = v228 + v147<<(uint(int32(4))%32)
	goto L61
L67:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v264 = v260 + int32(96)
	goto L61
L68:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+7)))
	if base.Ui32(v250) < base.Ui32(int32(-10002)-v147) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v264 = l0 + int32(72)
	goto L61
L70:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v237
	v264 = l0 + int32(88)
	goto L61
L71:
	;
	v271 = int32(-1)
	goto L60
L72:
	;
	v264 = v249 + (int32(-10003)-v147)<<(uint(int32(4))%32) + int32(24)
	goto L61
L73:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v271 = v268
	goto L60
L74:
	;
	v271 = int32(-1)
	goto L60
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v142
	v275 = m.G3
	v278 = F_lua_pushfstring(m, l0, v275+int32(_a2084), v9)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L36
	} else {
		goto L76
	}
L76:
	;
	v282 = F_lua_tolstring(m, l0, int32(-1), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L36
	} else {
		goto L77
	}
L77:
	;
	if v147 < int32(1) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if l0 == v135 {
		goto L95
	} else {
		goto L96
	}
L79:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v336)))
	*(*int64)(unsafe.Add(mBase, uint32(v339))) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v336)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+8)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v344 + int32(16)
	goto L78
L80:
	;
	if v147 < int32(-9999) {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v294 = v289 + v147<<(uint(int32(4))%32) + int32(-16)
	v295 = m.G398
	if base.Ui32(v294) < base.Ui32(v288) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v297 = v294
	goto L84
L83:
	;
	v297 = v295
	goto L84
L84:
	;
	v336 = v297
	goto L79
L85:
	;
	switch v134 + int32(10003) {
	case 0:
		goto L88
	case 1:
		goto L89
	case 2:
		goto L90
	default:
		goto L87
	}
L86:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v336 = v300 + v147<<(uint(int32(4))%32)
	goto L79
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+7)))
	v324 = m.G398
	if base.Ui32(v323) < base.Ui32(int32(-10002)-v147) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v336 = l0 + int32(72)
	goto L79
L89:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v312
	v336 = l0 + int32(88)
	goto L79
L90:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v336 = v306 + int32(96)
	goto L79
L91:
	;
	v335 = v324
	goto L93
L92:
	;
	v335 = v322 + (int32(-10003)-v147)<<(uint(int32(4))%32) + int32(24)
	goto L93
L93:
	;
	v336 = v335
	goto L79
L94:
	;
	v444 = v282
	goto L40
L95:
	;
	goto L94
L96:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v356 - int32(16)
	goto L97
L97:
	;
	goto L98
L98:
	;
	goto L103
L103:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v424 + int32(16)
	v430 = v423 + int32(0)
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v430)))
	*(*int64)(unsafe.Add(mBase, uint32(v424))) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = v433
	goto L95
L104:
	;
	F_lua_createtable(m, l0, int32(0), int32(2))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L36
	} else {
		goto L108
	}
L105:
	;
	if v447 != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v449 = m.G3
	v452 = F_luaL_argerror(m, l0, v137, v449+int32(_a2085))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L36
	} else {
		goto L107
	}
L107:
	;
	v1129 = v452
	goto L38
L108:
	;
	v458 = int32(83)
	v459 = F___strchrnul(m, v444, v458)
	mBase = m.M
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	if v461 == v458 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v527 = int32(108)
	v528 = F___strchrnul(m, v444, v527)
	mBase = m.M
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v530 == v527 {
		goto L127
	} else {
		goto L128
	}
L110:
	;
	if v465 == int32(0) {
		goto L109
	} else {
		goto L114
	}
L111:
	;
	v465 = v459
	goto L113
L112:
	;
	v465 = int32(0)
	goto L113
L113:
	;
	goto L110
L114:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_lua_pushstring(m, l0, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L36
	} else {
		goto L115
	}
L115:
	;
	v472 = m.G3
	F_lua_setfield(m, l0, int32(-2), v472+int32(_a2086))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L36
	} else {
		goto L116
	}
L116:
	;
	F_lua_pushstring(m, l0, v9+int32(48))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L36
	} else {
		goto L117
	}
L117:
	;
	F_lua_setfield(m, l0, int32(-2), v472+int32(_a2087))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L36
	} else {
		goto L118
	}
L118:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v488)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v488))) = base.F64_convert_i32_s(v486)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v493 + int32(16)
	goto L119
L119:
	;
	F_lua_setfield(m, l0, int32(-2), v472+int32(_a2088))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L36
	} else {
		goto L120
	}
L120:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v504))) = base.F64_convert_i32_s(v502)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v509 + int32(16)
	goto L121
L121:
	;
	F_lua_setfield(m, l0, int32(-2), v472+int32(_a2089))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L36
	} else {
		goto L122
	}
L122:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	F_lua_pushstring(m, l0, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L36
	} else {
		goto L123
	}
L123:
	;
	F_lua_setfield(m, l0, int32(-2), v472+int32(_a2090))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L36
	} else {
		goto L124
	}
L124:
	;
	goto L109
L125:
	;
	v554 = int32(117)
	v555 = F___strchrnul(m, v444, v554)
	mBase = m.M
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	if v557 == v554 {
		goto L135
	} else {
		goto L136
	}
L126:
	;
	if v534 == int32(0) {
		goto L125
	} else {
		goto L130
	}
L127:
	;
	v534 = v528
	goto L129
L128:
	;
	v534 = int32(0)
	goto L129
L129:
	;
	goto L126
L130:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v539)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v539))) = base.F64_convert_i32_s(v537)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v544 + int32(16)
	goto L131
L131:
	;
	v549 = m.G3
	F_lua_setfield(m, l0, int32(-2), v549+int32(_a2091))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L36
	} else {
		goto L132
	}
L132:
	;
	goto L125
L133:
	;
	v581 = int32(110)
	v582 = F___strchrnul(m, v444, v581)
	mBase = m.M
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	if v584 == v581 {
		goto L143
	} else {
		goto L144
	}
L134:
	;
	if v561 == int32(0) {
		goto L133
	} else {
		goto L138
	}
L135:
	;
	v561 = v555
	goto L137
L136:
	;
	v561 = int32(0)
	goto L137
L137:
	;
	goto L134
L138:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v566)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v566))) = base.F64_convert_i32_s(v564)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v571 + int32(16)
	goto L139
L139:
	;
	v576 = m.G3
	F_lua_setfield(m, l0, int32(-2), v576+int32(_a2092))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L36
	} else {
		goto L140
	}
L140:
	;
	goto L133
L141:
	;
	v609 = int32(76)
	v610 = F___strchrnul(m, v444, v609)
	mBase = m.M
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if v612 == v609 {
		goto L153
	} else {
		goto L154
	}
L142:
	;
	if v588 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L143:
	;
	v588 = v582
	goto L145
L144:
	;
	v588 = int32(0)
	goto L145
L145:
	;
	goto L142
L146:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_lua_pushstring(m, l0, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L36
	} else {
		goto L147
	}
L147:
	;
	v595 = m.G3
	F_lua_setfield(m, l0, int32(-2), v595+int32(_a807))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L36
	} else {
		goto L148
	}
L148:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	F_lua_pushstring(m, l0, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L36
	} else {
		goto L149
	}
L149:
	;
	F_lua_setfield(m, l0, int32(-2), v595+int32(_a2093))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L36
	} else {
		goto L150
	}
L150:
	;
	goto L141
L151:
	;
	v866 = int32(1)
	v867 = int32(102)
	v868 = F___strchrnul(m, v444, v867)
	mBase = m.M
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868))))
	if v870 == v867 {
		goto L209
	} else {
		goto L210
	}
L152:
	;
	if v616 == int32(0) {
		goto L151
	} else {
		goto L156
	}
L153:
	;
	v616 = v610
	goto L155
L154:
	;
	v616 = int32(0)
	goto L155
L155:
	;
	goto L152
L156:
	;
	if l0 != v135 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v861 = m.G3
	F_lua_setfield(m, l0, int32(-2), v861+int32(_a2094))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L36
	} else {
		goto L207
	}
L158:
	;
	if v135 == l0 {
		goto L198
	} else {
		goto L199
	}
L159:
	;
	goto L162
L160:
	;
	goto L178
L161:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v640)))
	*(*int64)(unsafe.Add(mBase, uint32(v676))) = v677
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v640)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v676)+8)) = v679
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v681 + int32(16)
	goto L160
L162:
	;
	goto L168
L168:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v640 = v637 + int32(-32)
	goto L161
L176:
	;
	goto L157
L177:
	;
	v742 = v702 + int32(-32)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v743) <= base.Ui32(v742) {
		v760 = v743
		goto L192
	} else {
		goto L193
	}
L178:
	;
	goto L184
L184:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L177
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v760 + int32(-16)
	goto L176
L193:
	;
	__phi746 = v702 + int32(-48)
	__phi747 = v742
	v746 = __phi746
	v747 = __phi747
	goto L194
L194:
	;
	v749 = *(*int64)(unsafe.Add(mBase, uint32(v746)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v746))) = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v746)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v746)+8)) = v751
	v754 = v747 + int32(16)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v754) < base.Ui32(v755) {
		__phi746 = v747
		__phi747 = v754
		v746 = __phi746
		v747 = __phi747
		goto L194
	} else {
		goto L196
	}
L195:
	;
	v760 = v755
	goto L192
L196:
	;
	goto L195
L197:
	;
	goto L157
L198:
	;
	goto L197
L199:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v772 - int32(16)
	goto L200
L200:
	;
	goto L201
L201:
	;
	goto L206
L206:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v840 + int32(16)
	v846 = v839 + int32(0)
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v846)))
	*(*int64)(unsafe.Add(mBase, uint32(v840))) = v847
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v846)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v840)+8)) = v849
	goto L198
L207:
	;
	goto L151
L208:
	;
	if v874 == int32(0) {
		v1129 = v866
		goto L38
	} else {
		goto L212
	}
L209:
	;
	v874 = v868
	goto L211
L210:
	;
	v874 = int32(0)
	goto L211
L211:
	;
	goto L208
L212:
	;
	if l0 != v135 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1119 = m.G3
	F_lua_setfield(m, l0, int32(-2), v1119+int32(_a2095))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L36
	} else {
		goto L263
	}
L214:
	;
	if v135 == l0 {
		goto L254
	} else {
		goto L255
	}
L215:
	;
	goto L218
L216:
	;
	goto L234
L217:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v935 = *(*int64)(unsafe.Add(mBase, uint32(v898)))
	*(*int64)(unsafe.Add(mBase, uint32(v934))) = v935
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v898)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v934)+8)) = v937
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v939 + int32(16)
	goto L216
L218:
	;
	goto L224
L224:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v898 = v895 + int32(-32)
	goto L217
L232:
	;
	goto L213
L233:
	;
	v1000 = v960 + int32(-32)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1001) <= base.Ui32(v1000) {
		v1018 = v1001
		goto L248
	} else {
		goto L249
	}
L234:
	;
	goto L240
L240:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L233
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1018 + int32(-16)
	goto L232
L249:
	;
	__phi1004 = v960 + int32(-48)
	__phi1005 = v1000
	v1004 = __phi1004
	v1005 = __phi1005
	goto L250
L250:
	;
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(v1004)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1004))) = v1007
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1004)+8)) = v1009
	v1012 = v1005 + int32(16)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v1012) < base.Ui32(v1013) {
		__phi1004 = v1005
		__phi1005 = v1012
		v1004 = __phi1004
		v1005 = __phi1005
		goto L250
	} else {
		goto L252
	}
L251:
	;
	v1018 = v1013
	goto L248
L252:
	;
	goto L251
L253:
	;
	goto L213
L254:
	;
	goto L253
L255:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v1030 - int32(16)
	goto L256
L256:
	;
	goto L257
L257:
	;
	goto L262
L262:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1098 + int32(16)
	v1104 = v1097 + int32(0)
	v1105 = *(*int64)(unsafe.Add(mBase, uint32(v1104)))
	*(*int64)(unsafe.Add(mBase, uint32(v1098))) = v1105
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+8)) = v1107
	goto L254
L263:
	;
	v1129 = v866
	goto L38
L264:
	;
	v1129 = v1127
	goto L38
}
func F_db_getmetatable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	F_luaL_checkany(m, l0, int32(1))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = v13 + int32(0)
		v19 = m.G398
		if base.Ui32(v18) < base.Ui32(v12) {
			v21 = v18
		} else {
			v21 = v19
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
		switch v63 + int32(-5) {
		case 0:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v78 = v66 + int32(16)
		default:
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v78 = v72 + v63<<(uint(int32(2))%32) + int32(152)
		case 2:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v78 = v69 + int32(8)
		}
		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
		if v79 != 0 {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(v81))) = v79
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v85 + int32(16)
		} else {
		}
		if v79 != 0 {
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93 + int32(16)
		}
		return int32(1)
	}
}
func F_db_getupvalue(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_auxupvalue(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_db_setfenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	F_luaL_checktype(m, l0, int32(2), int32(5))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v16 = v13 + int32(32)
		if base.Ui32(v16) <= base.Ui32(v12) {
		} else {
			v20 = v12
			for {
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(0)
				v24 = v20 + int32(16)
				if base.Ui32(v24) < base.Ui32(v16) {
					v20 = v24
					continue
				} else {
					break
				}
				break
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v16
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = v42 + int32(0)
		v48 = m.G398
		if base.Ui32(v47) < base.Ui32(v41) {
			v50 = v47
		} else {
			v50 = v48
		}
		v93 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
		switch v93 + int32(-6) {
		case 0:
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+int32(-16))))
			*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v100
			v118 = int32(1)
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
			if v123&int32(3) == int32(0) {
				v136 = v118
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
				if v129&int32(4) == int32(0) {
					v136 = v118
				} else {
					F_luaC_barrierf(m, l0, v128, v122)
					mBase = m.M
					v136 = v118
				}
			}
		case 1:
			v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(-16))))
			*(*int32)(unsafe.Add(mBase, uint32(v102)+12)) = v106
			v118 = int32(1)
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
			if v123&int32(3) == int32(0) {
				v136 = v118
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
				if v129&int32(4) == int32(0) {
					v136 = v118
				} else {
					F_luaC_barrierf(m, l0, v128, v122)
					mBase = m.M
					v136 = v118
				}
			}
		case 2:
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v108+int32(-16))))
			v112 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			*(*int32)(unsafe.Add(mBase, uint32(v112)+80)) = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(v112)+72)) = v111
			v118 = int32(1)
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+int32(-16))))
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
			if v123&int32(3) == int32(0) {
				v136 = v118
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
				if v129&int32(4) == int32(0) {
					v136 = v118
				} else {
					F_luaC_barrierf(m, l0, v128, v122)
					mBase = m.M
					v136 = v118
				}
			}
		default:
			v136 = int32(0)
		}
		v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v138 + int32(-16)
		if v136 != 0 {
			return int32(1)
		} else {
			v142 = m.G3
			v146 = F_luaL_error(m, l0, v142+int32(_a2096), int32(0))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}
	}
}
