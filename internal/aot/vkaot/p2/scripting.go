package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_scriptingEngineCallDebuggerDisable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v6) <= base.Ui64(int64(3)) {
		F__serverAssert(m, int32(_a1316), int32(_a1310), int32(458))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v9 == int32(0) {
			F__serverAssert(m, int32(_a1317), int32(_a1310), int32(459))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v15 = int32(0)
				F_moduleScriptingEngineInitContext(m, v14, v12, v15, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v20 = v19
					v21 = v14
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					m.T0[v20].(func(*base.Module, int32, int32, int32))(m, v21, v22, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v25 == int32(0) {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v20 = v9
				v21 = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				m.T0[v20].(func(*base.Module, int32, int32, int32))(m, v21, v22, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v25 == int32(0) {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
func F_scriptingEngineCallDebuggerStart(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v7) <= base.Ui64(int64(3)) {
		F__serverAssert(m, int32(_a1316), int32(_a1310), int32(472))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v10 == int32(0) {
			F__serverAssert(m, int32(_a1318), int32(_a1310), int32(473))
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
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v13 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v16 = int32(0)
				F_moduleScriptingEngineInitContext(m, v15, v13, v16, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v21 = v20
					v22 = v15
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					m.T0[v21].(func(*base.Module, int32, int32, int32, int32))(m, v22, v23, l1, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v26 == int32(0) {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v21 = v10
				v22 = int32(0)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				m.T0[v21].(func(*base.Module, int32, int32, int32, int32))(m, v22, v23, l1, l2)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v26 == int32(0) {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
func F_scriptingEngineCallFreeFunction(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	if base.Ui32(int32(2)) <= base.Ui32(l1) {
		F__serverAssert(m, int32(_a1313), int32(_a1310), int32(321))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[719]))
		v10 = F___get_tp(m)
		mBase = m.M
		if v9 == v10 {
			if l0 == int32(0) {
				F__serverAssert(m, int32(_a1314), int32(_a1310), int32(261))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v38 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v41 = int32(0)
					F_moduleScriptingEngineInitContext(m, v40, v38, v41, v41)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = v40
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						m.T0[v47].(func(*base.Module, int32, int32, int32, int32))(m, v45, v46, l1, l2)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v50 == int32(0) {
								return
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								F_moduleFreeContext(m, v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v45 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					m.T0[v47].(func(*base.Module, int32, int32, int32, int32))(m, v45, v46, l1, l2)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v50 == int32(0) {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							F_moduleFreeContext(m, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v13 = F___pthread_mutex_lock(m, int32(_a1315))
			mBase = m.M
			if l0 == int32(0) {
				F__serverAssert(m, int32(_a1314), int32(_a1310), int32(261))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v16 != 0 {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v19 = int32(0)
					F_moduleScriptingEngineInitContext(m, v18, v16, v19, v19)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = v18
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						m.T0[v25].(func(*base.Module, int32, int32, int32, int32))(m, v23, v24, l1, l2)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v28 == int32(0) {
								v35 = F___pthread_mutex_unlock(m, int32(_a1315))
								mBase = m.M
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								F_moduleFreeContext(m, v31)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v35 = F___pthread_mutex_unlock(m, int32(_a1315))
									mBase = m.M
									return
								}
							}
						}
					}
				} else {
					v23 = int32(0)
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					m.T0[v25].(func(*base.Module, int32, int32, int32, int32))(m, v23, v24, l1, l2)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v28 == int32(0) {
							v35 = F___pthread_mutex_unlock(m, int32(_a1315))
							mBase = m.M
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							F_moduleFreeContext(m, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v35 = F___pthread_mutex_unlock(m, int32(_a1315))
								mBase = m.M
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_scriptingEngineCallFunction(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		F__serverAssert(m, int32(_a1313), int32(_a1310), int32(352))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l0 == int32(0) {
			F__serverAssert(m, int32(_a1314), int32(_a1310), int32(261))
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
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				F_moduleScriptingEngineInitContext(m, v18, v16, int32(1), l2)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					v22 = v18
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					m.T0[v24].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, v22, v23, l1, l3, l4, l5, l6, l7, l8)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v27 == int32(0) {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v22 = int32(0)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				m.T0[v24].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, v22, v23, l1, l3, l4, l5, l6, l7, l8)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v27 == int32(0) {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
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
func F_scriptingEngineDebuggerKillForkedSessions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v1 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v10 = v5 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v1
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	goto L1
L1:
	;
	v16 = v5 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	F_listRelease(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L18
	}
L3:
	;
	if v18 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18+base.B2i32(v21 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	goto L4
L6:
	;
	v32 = v18
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v35 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v44 = F_kill(m, v33, int32(9))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v33
	F__serverLog(m, int32(2), int32(_a1319), v5)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	goto L9
L13:
	;
	v47 = v5 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v49 != 0 {
		v32 = v49
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49+base.B2i32(v52 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v58
	goto L15
L17:
	;
	goto L8
L18:
	;
	v67 = F_listCreate(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[720])) = v67
	m.G0 = v5 + int32(16)
	return
}
func F_scriptingEngineDebuggerPendingChildren(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	return v3
}
func F_scriptingEngineGetAbiVersion(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	return v2
}
func F_scriptingEngineManagerGetTotalMemoryOverhead(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	return v2
}
func F_scriptingEngineManagerRegister(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v107 int64
	_ = v107
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a1309), int32(_a1310), int32(140))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L46
	}
L2:
	;
	v14 = F_sdsnew(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	v20 = F_dictFetchValue(m, v19, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v10 + int32(48)
	return v210
L6:
	;
	v39 = F_valkey_malloc(m, int32(80))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L13
	}
L7:
	;
	if v20 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v25 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_sdsfree(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	F__serverLog(m, int32(3), int32(_a1311), v10+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v210 = int32(-1)
	goto L5
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v14
	v49 = F__emscripten_memset_bulkmem(m, v39+int32(12), base.I32_extend8_s(int32(0)), int32(68))
	mBase = m.M
	goto L14
L14:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui64(int64(3)) < base.Ui64(v50) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v115 = F_moduleAllocateContext(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L21
	}
L16:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v83
	v87 = int32(40)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l3+v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(56)))) = v89
	v93 = int32(32)
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l3+v93)))
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(48)))) = v95
	v99 = int32(24)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l3+v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v87))) = v101
	v107 = *(*int64)(unsafe.Add(mBase, uint32(l3+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v93))) = v107
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l3+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v99))) = v113
	goto L15
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v54 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v63
	v67 = int32(24)
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l3+v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(40)))) = v69
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l3+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v39+int32(32)))) = v75
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l3+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v67))) = v81
	goto L15
L19:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
	F__serverLog(m, int32(3), int32(_a1312), v10)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v115
	v118 = F_moduleAllocateContext(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v118
	v121 = F_moduleAllocateContext(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v121
	v124 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	v127 = F_dictAdd(m, v126, v14, v39)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v129 == int32(0) {
		v137 = v124
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v39)+44))
	m.T0[v142].(func(*base.Module, int32, int32, int32, int32))(m, v10+int32(32), v137, v140, int32(2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	v133 = int32(0)
	F_moduleScriptingEngineInitContext(m, v132, v129, v133, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v137 = v132
	goto L25
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v145 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v151 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(-8))))
	goto L32
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	F_moduleFreeContext(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v167 = v160 + int32(-1)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v170 = v168 & int32(7)
	switch v170 {
	case 0:
		goto L39
	case 1:
		v176 = int32(4)
		goto L34
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		goto L35
	}
L33:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v205 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v155&int32(2147483647) + int32(8) + v200 + v202 + v205
	v210 = v151
	goto L5
L34:
	;
	switch v170 {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	case 3:
		goto L42
	case 4:
		goto L41
	default:
		v196 = int32(0)
		goto L40
	}
L35:
	;
	v176 = int32(1)
	goto L34
L36:
	;
	v176 = int32(18)
	goto L34
L37:
	;
	v176 = int32(10)
	goto L34
L38:
	;
	v176 = int32(6)
	goto L34
L39:
	;
	v171 = F_zmalloc_usable_size(m, v167)
	mBase = m.M
	v200 = v171
	goto L33
L40:
	;
	v200 = v176 + v196
	goto L33
L41:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v160+int32(-9))))
	v196 = v195
	goto L40
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v160+int32(-5))))
	v200 = v176 + v191
	goto L33
L43:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160+int32(-3)))))
	v200 = v176 + v187
	goto L33
L44:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+int32(-2)))))
	v200 = v176 + v183
	goto L33
L45:
	;
	v200 = v176 + int32(base.Ui32(v168)>>(uint(int32(3))%32))
	goto L33
L46:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
