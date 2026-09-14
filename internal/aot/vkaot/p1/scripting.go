package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_collectScriptingEngineInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
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
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_scriptingEngineCallGetMemoryInfo(m, v10+int32(32), l0, int32(2))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v13 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v24 = v23
		} else {
			v24 = int32(_a_F_collectScriptingEngineInfo_0)
		}
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
		*(*uint32)(unsafe.Add(mBase, uint32(v10)+12)) = uint32(v14)
		v34 = F_sdscatprintf(m, v21, int32(_a_F_collectScriptingEngineInfo_1), v10)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37 + int32(1)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v41 + v42
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v45 + v46
			m.G0 = v10 + int32(48)
			return
		}
	}
}
func F_scriptingEngineCallCompileCode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	if base.Ui32(int32(2)) <= base.Ui32(l1) {
		F__serverAssert(m, int32(_a_F_scriptingEngineCallCompileCode_0), int32(_a_F_scriptingEngineCallCompileCode_1), int32(287))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l0 == int32(0) {
			F__serverAssert(m, int32(_a_F_scriptingEngineCallCompileCode_2), int32(_a_F_scriptingEngineCallCompileCode_1), int32(261))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v18 = int32(0)
				F_moduleScriptingEngineInitContext(m, v17, v15, v18, v18)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = v17
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					if v27 != int64(1) {
						v32 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, l1, l2, l3, l4, l5, l6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v35 == int32(0) {
								return v34
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								F_moduleFreeContext(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									return v34
								}
							}
						}
					} else {
						v30 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, l1, l2, l4, l5, l6)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v34 = v30
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v35 == int32(0) {
								return v34
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								F_moduleFreeContext(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									return v34
								}
							}
						}
					}
				}
			} else {
				v24 = int32(0)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 != int64(1) {
					v32 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, l1, l2, l3, l4, l5, l6)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = v32
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v35 == int32(0) {
							return v34
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v34
							}
						}
					}
				} else {
					v30 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v24, v25, l1, l2, l4, l5, l6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v34 = v30
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v35 == int32(0) {
							return v34
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_scriptingEngineCallDebuggerEnable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(int64(3)) < base.Ui64(v14) {
		v31 = int32(0)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v32 == v31 {
			v64 = v31
			m.G0 = v12 + int32(16)
			return v64
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			if v35 == int32(0) {
				v64 = v31
				m.G0 = v12 + int32(16)
				return v64
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v38 == int32(0) {
					v64 = v31
					m.G0 = v12 + int32(16)
					return v64
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					if v41 == int32(0) {
						v64 = v31
						m.G0 = v12 + int32(16)
						return v64
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v44 != 0 {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v47 = int32(0)
							F_moduleScriptingEngineInitContext(m, v46, v44, v47, v47)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v52 = v46
								v53 = v51
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v55 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v52, v54, l1, l2, l3)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v57 == int32(0) {
										v64 = v55
										m.G0 = v12 + int32(16)
										return v64
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										F_moduleFreeContext(m, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v64 = v55
											m.G0 = v12 + int32(16)
											return v64
										}
									}
								}
							}
						} else {
							v52 = int32(0)
							v53 = v32
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v55 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v52, v54, l1, l2, l3)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v57 == int32(0) {
									v64 = v55
									m.G0 = v12 + int32(16)
									return v64
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									F_moduleFreeContext(m, v60)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v64 = v55
										m.G0 = v12 + int32(16)
										return v64
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v17 = int32(0)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineCallDebuggerEnable[0]))
		if int32(3) < v19 {
			v64 = v17
			m.G0 = v12 + int32(16)
			return v64
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v14)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v22
			F__serverLog(m, int32(3), int32(_a_F_scriptingEngineCallDebuggerEnable_0), v12)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v64 = v17
				m.G0 = v12 + int32(16)
				return v64
			}
		}
	}
}
func F_scriptingEngineDebuggerDisable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerDisable[0]))
	if v4 == v2 {
		return
	} else {
		v7 = int32(0)
		*(*int64)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerDisable[1])) = int64(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v10 & int32(-805306369)
		F_scriptingEngineCallDebuggerDisable(m, v4, v7)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_scriptingEngineDebuggerEndSession(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[0]))
	if v9 == v2 {
		F__serverAssert(m, int32(_a_F_scriptingEngineDebuggerEndSession_0), int32(_a_F_scriptingEngineDebuggerEndSession_1), int32(695))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v14 = F_sdsnew(m, int32(_a_F_scriptingEngineDebuggerEndSession_2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = F_createObject(m, int32(0), v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[1]))
				v20 = F_listAddNodeTail(m, v19, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_scriptingEngineDebuggerFlushLogs(m)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = int32(0)
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[2]))
						if v25 == v24 {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[3]))
							if int32(2) < v42 {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[4]))
								v53 = F_connNonBlock(m, v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[4]))
									v58 = F_connSendTimeout(m, v56, int64(0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v60 | int32(64)
										v64 = int32(0)
										v65 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[5]))
										F_scriptingEngineCallDebuggerEnd(m, v65, v64)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											m.G0 = v6 + int32(16)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_scriptingEngineDebuggerEndSession_3)
								F__serverLog(m, int32(2), int32(_a_F_scriptingEngineDebuggerEndSession_4), v6)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[4]))
									v53 = F_connNonBlock(m, v52)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[4]))
										v58 = F_connSendTimeout(m, v56, int64(0))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v60 | int32(64)
											v64 = int32(0)
											v65 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[5]))
											F_scriptingEngineCallDebuggerEnd(m, v65, v64)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												m.G0 = v6 + int32(16)
												return
											}
										}
									}
								}
							}
						} else {
							v28 = F_writeToClient(m, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerEndSession[3]))
								if int32(2) < v31 {
									F__exit(m, int32(0))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									F__serverLog(m, int32(2), int32(_a_F_scriptingEngineDebuggerEndSession_5), int32(0))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										F__exit(m, int32(0))
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
}
func F_scriptingEngineDebuggerFlushLogs(m *base.Module) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerFlushLogs[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
	v18 = F_sdscatfmt(m, v11, int32(_a_F_scriptingEngineDebuggerFlushLogs_0), v8)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerFlushLogs[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v22 == v20 {
		v140 = v18
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerFlushLogs[1]))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-1)))))
	switch v146 & int32(7) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	default:
		v163 = int32(0)
		goto L32
	}
L5:
	;
	v27 = v21
	v28 = v18
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v34 = F_sdscatlen(m, v28, int32(_a_F_scriptingEngineDebuggerFlushLogs_1), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v140 = v128
	goto L4
L8:
	;
	v36 = F_objectGetVal(m, v31)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-1)))))
	switch v47 & int32(7) {
	case 0:
		goto L16
	case 1:
		goto L15
	case 2:
		goto L14
	case 3:
		goto L13
	case 4:
		goto L12
	default:
		goto L10
	}
L9:
	;
	v123 = F_objectGetVal(m, v31)
	mBase = m.M
	v124 = F_sdscatsds(m, v34, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L28
	}
L10:
	;
	goto L9
L11:
	;
	if v64 == int32(0) {
		goto L10
	} else {
		goto L17
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v64 = v63
	goto L11
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v64 = v60
	goto L11
L14:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v64 = v57
	goto L11
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v64 = v54
	goto L11
L16:
	;
	v64 = int32(base.Ui32(v47) >> (uint(int32(3)) % 32))
	goto L11
L17:
	;
	v74 = int32(0)
	goto L18
L18:
	;
	goto L21
L19:
	;
	goto L10
L20:
	;
	v112 = v74 + int32(1)
	if v112 != v64 {
		v74 = v112
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v80 = v36 + v74
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v88 = int32(0)
	goto L22
L22:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_scriptingEngineDebuggerFlushLogs[2]))))
	if v81&int32(255) != v94 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L20
L24:
	;
	v100 = v88 + int32(1)
	if v100 != int32(2) {
		v88 = v100
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_scriptingEngineDebuggerFlushLogs[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v97)
	goto L20
L26:
	;
	goto L23
L27:
	;
	goto L19
L28:
	;
	v128 = F_sdscatlen(m, v124, int32(_a_F_scriptingEngineDebuggerFlushLogs_2), int32(2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerFlushLogs[0]))
	F_listDelNode(m, v131, v30)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerFlushLogs[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	if v136 != 0 {
		v27 = v135
		v28 = v128
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L7
L32:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+68))
	v166 = m.T0[v165].(func(*base.Module, int32, int32, int32) int32)(m, v143, v140, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-17))))
	v163 = v162
	goto L32
L34:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v140+int32(-9))))
	v163 = v159
	goto L32
L35:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+int32(-5)))))
	v163 = v156
	goto L32
L36:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(-3)))))
	v163 = v153
	goto L32
L37:
	;
	v163 = int32(base.Ui32(v146) >> (uint(int32(3)) % 32))
	goto L32
L38:
	;
	F_sdsfree(m, v140)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v8 + int32(16)
	return
}
func F_scriptingEngineDebuggerLog(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerLog[0]))
	v4 = F_listAddNodeTail(m, v3, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_scriptingEngineDebuggerRemoveChild(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerRemoveChild[0]))
	v6 = F_listSearchKey(m, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v17 = v2
			return v17
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineDebuggerRemoveChild[0]))
			F_listDelNode(m, v13, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = int32(1)
				return v17
			}
		}
	}
}
func F_scriptingEngineGetModule(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_scriptingEngineGetName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_scriptingEngineManagerFind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineManagerFind[0]))
	v6 = F_dictFind(m, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v13 = v2
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v13 = v12
		}
		return v13
	}
}
func F_scriptingEngineManagerForEachEngine(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineManagerForEachEngine[0]))
	v7 = F_dictGetIterator(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	F_dictReleaseIterator(m, v7)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L35
	}
L4:
	;
	v20 = v7 + int32(20)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if v116 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L7:
	;
	v27 = v20
	v28 = v24
	goto L10
L8:
	;
	v24 = int32(1)
	goto L7
L9:
	;
	v24 = int32(0)
	goto L7
L10:
	;
	switch v28 {
	case 0:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v28 = int32(0)
	goto L10
L13:
	;
	goto L6
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v108
	if v108 == int32(0) {
		goto L12
	} else {
		goto L31
	}
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v32 != int32(-1) {
		v71 = v32
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = int32(1)
	v73 = v71 + v72
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v73
	v75 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v79+int32(26)))))
	if v83 == int32(255) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v36 != 0 {
		v71 = int32(-1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v38 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v65 != int32(-1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+16)))
	v46 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+27)))
	v47 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+8)))
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+12)))
	v49 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+26)))
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+4)))
	v51 = F_wangHash64(m, v50)
	mBase = m.M
	v53 = F_wangHash64(m, v49+v51)
	mBase = m.M
	v55 = F_wangHash64(m, v48+v53)
	mBase = m.M
	v57 = F_wangHash64(m, v47+v55)
	mBase = m.M
	v59 = F_wangHash64(m, v46+v57)
	mBase = m.M
	v61 = F_wangHash64(m, v45+v59)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v64 = v63
	goto L19
L21:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)))
	v43 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+24)) = uint16(v43)
	v64 = v37
	goto L19
L22:
	;
	v71 = v65 + int32(-1)
	goto L16
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v71 = v68
	goto L16
L24:
	;
	v98 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78+v96<<(uint(v98)%32)+int32(4))))
	v27 = v103 + v97<<(uint(v98)%32)
	v28 = int32(1)
	goto L10
L25:
	;
	v87 = v75
	goto L27
L26:
	;
	v87 = v72 << (uint(v83) % 32)
	goto L27
L27:
	;
	if v73 < v87 {
		v96 = v79
		v97 = v73
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v79 != 0 {
		v116 = v75
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if v89 == int32(-1) {
		v116 = v75
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967296)
	v96 = int32(1)
	v97 = int32(0)
	goto L24
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v112
	v116 = v108
	goto L13
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	goto L33
L33:
	;
	m.T0[l0].(func(*base.Module, int32, int32))(m, v122, l1)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L4
L35:
	;
	return
}
func F_scriptingEngineManagerGetMemoryUsage(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v1 = int32(0)
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_scriptingEngineManagerGetMemoryUsage[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+27)))
	if v6 == int32(255) {
		v10 = v1
	} else {
		v10 = int32(1) << (uint(v6) % 32)
	}
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+26)))
	if v13 == int32(255) {
		v17 = int32(0)
	} else {
		v17 = int32(1) << (uint(v13) % 32)
	}
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	return (v10+v17)<<(uint(int32(2))%32) + (v21+v22)*int32(24) + int32(8)
}
