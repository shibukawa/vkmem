package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_scriptingEngineCallDebuggerEnd(m *base.Module, l0 int32, l1 int32) {
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
		F__serverAssert(m, int32(_a1097), int32(_a1093), int32(486))
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
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v9 == int32(0) {
			F__serverAssert(m, int32(_a1098), int32(_a1093), int32(487))
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
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
func F_scriptingEngineCallGetFunctionMemoryOverhead(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a1092), int32(_a1093), int32(261))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v7 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v10 = int32(0)
			F_moduleScriptingEngineInitContext(m, v9, v7, v10, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = v9
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v18 = m.T0[v17].(func(*base.Module, int32, int32) int32)(m, v16, l1)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v20 == int32(0) {
						return v18
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return v18
						}
					}
				}
			}
		} else {
			v16 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v18 = m.T0[v17].(func(*base.Module, int32, int32) int32)(m, v16, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v20 == int32(0) {
					return v18
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					F_moduleFreeContext(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v18
					}
				}
			}
		}
	}
}
func F_scriptingEngineCallGetMemoryInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	if l1 == int32(0) {
		F__serverAssert(m, int32(_a1092), int32(_a1093), int32(261))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			v11 = int32(0)
			F_moduleScriptingEngineInitContext(m, v10, v8, v11, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = v10
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				m.T0[v17].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15, v16, l2)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v20 == int32(0) {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
						F_moduleFreeContext(m, v23)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v15 = int32(0)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			m.T0[v17].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15, v16, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v20 == int32(0) {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
					F_moduleFreeContext(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_scriptingEngineCallResetEnvFunc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	if l0 == int32(0) {
		F__serverAssert(m, int32(_a1092), int32(_a1093), int32(261))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v11 = int32(0)
			F_moduleScriptingEngineInitContext(m, v10, v8, v11, v11)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = v10
				v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				if base.Ui64(int64(2)) < base.Ui64(v18) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v28 = m.T0[v27].(func(*base.Module, int32, int32, int32, int32) int32)(m, v17, v26, l1, l2)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = v28
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v31 == int32(0) {
							return v30
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				} else {
					if l1 != 0 {
						v30 = int32(0)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v31 == int32(0) {
							return v30
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, v17, v22, l2)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v30 = v24
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v31 == int32(0) {
								return v30
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								F_moduleFreeContext(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									return v30
								}
							}
						}
					}
				}
			}
		} else {
			v17 = int32(0)
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			if base.Ui64(int64(2)) < base.Ui64(v18) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v28 = m.T0[v27].(func(*base.Module, int32, int32, int32, int32) int32)(m, v17, v26, l1, l2)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = v28
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v31 == int32(0) {
						return v30
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							return v30
						}
					}
				}
			} else {
				if l1 != 0 {
					v30 = int32(0)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v31 == int32(0) {
						return v30
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_moduleFreeContext(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							return v30
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, v17, v22, l2)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v30 = v24
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v31 == int32(0) {
							return v30
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_moduleFreeContext(m, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				}
			}
		}
	}
}
func F_scriptingEngineDebuggerEnable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = F_scriptingEngineCallDebuggerEnable(m, l1, int32(0), int32(_a1099), int32(_a1100))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v69
L2:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[614])) = l1
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v38 | int32(268435456)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	goto L12
L3:
	;
	v25 = F_sdsempty(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L9
	}
L4:
	;
	v16 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	switch v12 {
	case 0:
		goto L4
	default:
		goto L2
	case 2:
		goto L3
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v18
	v21 = F_sdscatfmt(m, v16, int32(_a1101), v7)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
	v69 = int32(-1)
	goto L1
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v27
	v32 = F_sdscatfmt(m, v25, int32(_a1102), v7+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
	v69 = int32(-1)
	goto L1
L11:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[612])) = v55
	v58 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	F_sdsfree(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v48 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_listDelNode(m, v43, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v61 = F_sdsempty(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v63 = int32(0)
	*(*int64)(unsafe.Add(mBase, _consts[619])) = int64(256)
	*(*int32)(unsafe.Add(mBase, _consts[617])) = v61
	v69 = v53
	goto L1
}
func F_scriptingEngineDebuggerLogRespReplyStr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_sdsnew(m, int32(_a1116))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v8
		v13 = F_debugScriptRespToHuman(m, v5+int32(12), l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v17 = F_createObject(m, int32(0), v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_scriptingEngineDebuggerLogWithMaxLen(m, v17)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	}
}
func F_scriptingEngineDebuggerLogWithMaxLen(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	if v7 == v2 {
		v150 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v155 = F_listAddNodeTail(m, v154, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L44
	} else {
		goto L47
	}
L2:
	;
	v11 = F_objectGetVal(m, l0)
	mBase = m.M
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-1)))))
	switch v14 & int32(7) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		v150 = int32(0)
		goto L1
	}
L3:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	if base.Ui32(v31) <= base.Ui32(v34) {
		v150 = v32
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-17))))
	v31 = v30
	goto L3
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(-9))))
	v31 = v27
	goto L3
L6:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11+int32(-5)))))
	v31 = v24
	goto L3
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-3)))))
	v31 = v21
	goto L3
L8:
	;
	v31 = int32(base.Ui32(v14) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	v36 = F_objectGetVal(m, l0)
	mBase = m.M
	v37 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	v40 = int32(-1)
	v41 = v39 + v40
	v49 = v36 + v40
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v52 = v50 & int32(7)
	switch v52 {
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
		goto L11
	}
L10:
	;
	v142 = F_objectGetVal(m, l0)
	mBase = m.M
	v145 = F_sdscatlen(m, v142, int32(_a1104), int32(4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L44
	} else {
		goto L45
	}
L11:
	;
	goto L10
L12:
	;
	if v67 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v67 = v66
	goto L12
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v67 = v63
	goto L12
L15:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v67 = v60
	goto L12
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v67 = v57
	goto L12
L17:
	;
	v67 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
	goto L12
L18:
	;
	v73 = v41>>(uint(int32(31))%32)&v67 + v41
	v77 = int32(0)&v67 + v37
	v80 = v73 - v77 + int32(1)
	switch v52 {
	default:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	}
L19:
	;
	v96 = int32(0)
	v98 = base.B2i32(base.Ui32(v77) < base.Ui32(v95))
	if base.Ui32(v77) < base.Ui32(v95) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-17))))
	v95 = v94
	goto L19
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9))))
	v95 = v91
	goto L19
L22:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))))
	v95 = v88
	goto L19
L23:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))))
	v95 = v85
	goto L19
L24:
	;
	v95 = int32(base.Ui32(v50) >> (uint(int32(3)) % 32))
	goto L19
L25:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36+v106))) = uint8(v112)
	switch v52 {
	default:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	case 4:
		goto L39
	}
L26:
	;
	v99 = v77
	goto L28
L27:
	;
	v99 = v96
	goto L28
L28:
	;
	v100 = v95 - v99
	if base.Ui32(v80) < base.Ui32(v100) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v102 = v80
	goto L31
L30:
	;
	v102 = v100
	goto L31
L31:
	;
	if v73 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v104 = v96
	goto L34
L33:
	;
	v104 = v102
	goto L34
L34:
	;
	if base.Ui32(v77) < base.Ui32(v95) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v106 = v104
	goto L37
L36:
	;
	v106 = int32(0)
	goto L37
L37:
	;
	if v106 == int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v110 = F_memmove(m, v36, v36+v99, v106)
	mBase = m.M
	goto L25
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36+int32(-17)))) = base.I64_extend_i32_u(v106)
	goto L11
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(-9)))) = v106
	goto L10
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v36+int32(-5)))) = uint16(v106)
	goto L10
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v36+int32(-3)))) = uint8(v106)
	goto L10
L43:
	;
	v115 = v106 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v115)
	goto L10
L44:
	;
	return
L45:
	;
	F_objectSetVal(m, l0, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v150 = int32(1)
	goto L1
L47:
	;
	if v150 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[620]))
	if v160 != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v161 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[620])) = int32(1)
	v166 = F_sdsnew(m, int32(_a1103))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v168 = F_createObject(m, v161, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v172 = F_listAddNodeTail(m, v171, v168)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L48
}
func F_scriptingEngineDebuggerProcessCommands(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1222 int32
	_ = v1222
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1403 int32
	_ = v1403
	var v1417 int32
	_ = v1417
	var v1424 int32
	_ = v1424
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	v20 = m.G0
	v22 = v20 - int32(1024)
	m.G0 = v22
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a583), int32(_a1093), int32(1061))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L28
	} else {
		goto L354
	}
L2:
	;
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
	v34 = v26
	goto L5
L3:
	;
	F__serverAssert(m, int32(_a1108), int32(_a1093), int32(759))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L28
	} else {
		goto L353
	}
L4:
	;
	m.G0 = v22 + int32(1024)
	return
L5:
	;
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v1417 = int32(0)
	if v516 == v1417 {
		goto L344
	} else {
		goto L345
	}
L8:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+32))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+36))
	if v1389 != 0 {
		v1393 = v1389
		goto L341
	} else {
		goto L342
	}
L9:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v807 = F_objectGetVal(m, v806)
	mBase = m.M
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807+int32(-1)))))
	switch v810 & int32(7) {
	case 0:
		goto L211
	case 1:
		goto L212
	case 2:
		goto L213
	case 3:
		goto L214
	case 4:
		goto L215
	default:
		v827 = int32(0)
		goto L210
	}
L10:
	;
	F__serverAssert(m, int32(_a1109), int32(_a1093), int32(1090))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L28
	} else {
		goto L209
	}
L11:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v777 == int32(0) {
		v34 = v447
		goto L5
	} else {
		goto L208
	}
L12:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	F_sdsfree(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L28
	} else {
		goto L145
	}
L13:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v463 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L14:
	;
	F_valkey_free(m, v428)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L28
	} else {
		goto L129
	}
L15:
	;
	F__serverAssert(m, int32(_a1110), int32(_a1093), int32(760))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L28
	} else {
		goto L128
	}
L16:
	;
	v50 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
	switch v55 & int32(7) {
	case 0:
		goto L22
	case 1:
		goto L21
	case 2:
		goto L20
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		v447 = v50
		goto L13
	}
L17:
	;
	v73 = int32(0)
	if v72 == v73 {
		v447 = v73
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
	v72 = v71
	goto L17
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
	v72 = v68
	goto L17
L20:
	;
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
	v72 = v65
	goto L17
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
	v72 = v62
	goto L17
L22:
	;
	v72 = int32(base.Ui32(v55) >> (uint(int32(3)) % 32))
	goto L17
L23:
	;
	v76 = F_sdsdup(m, v52)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	F_sdsfree(m, v76)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L28
	} else {
		goto L126
	}
L25:
	;
	v382 = int32(0)
	if v368 == v382 {
		v424 = v366
		v428 = v370
		goto L14
	} else {
		goto L121
	}
L26:
	;
	v360 = F_createStringObject_1(m, int32(_a1111), int32(14))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L28
	} else {
		goto L120
	}
L27:
	;
	v89 = int32(_a727)
	v92 = int32(*(*int8)(unsafe.Add(mBase, _consts[623])))
	if v92 != 0 {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	return
L29:
	;
	v78 = int32(42)
	v79 = F___strchrnul(m, v76, v78)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v81 == v78 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v85 != 0 {
		goto L27
	} else {
		goto L34
	}
L31:
	;
	v85 = v79
	goto L33
L32:
	;
	v85 = int32(0)
	goto L33
L33:
	;
	goto L30
L34:
	;
	v86 = int32(0)
	v342 = v86
	v344 = v86
	v346 = v86
	goto L26
L35:
	;
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	v128 = v85 + int32(1)
	goto L55
L36:
	;
	if v117 != 0 {
		goto L35
	} else {
		goto L52
	}
L37:
	;
	v93 = int32(0)
	v94 = F_strchr(m, v85, v92)
	mBase = m.M
	if v94 == v93 {
		v114 = v93
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v117 = v85
	goto L36
L39:
	;
	v117 = v114
	goto L36
L40:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[624])))
	if v97 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	if v98 == int32(0) {
		v114 = v93
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v117 = v94
	goto L36
L43:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _consts[625])))
	if v101 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
	if v103 == int32(0) {
		v114 = v93
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v102 = F_twobyte_strstr(m, v94, v89)
	mBase = m.M
	v117 = v102
	goto L36
L46:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _consts[626])))
	if v106 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+3)))
	if v108 == int32(0) {
		v114 = v93
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v107 = F_threebyte_strstr(m, v94, v89)
	mBase = m.M
	v117 = v107
	goto L36
L49:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _consts[627])))
	if v111 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v113 = F_twoway_strstr(m, v94, v89)
	mBase = m.M
	v114 = v113
	goto L39
L51:
	;
	v112 = F_fourbyte_strstr(m, v94, v89)
	mBase = m.M
	v117 = v112
	goto L36
L52:
	;
	v118 = int32(0)
	v424 = v118
	v428 = v118
	goto L14
L53:
	;
	v189 = int32(2)
	v194 = F_valkey_malloc(m, v173<<(uint(v189)%32))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L28
	} else {
		goto L70
	}
L54:
	;
	if base.Ui32(int32(-1024)) <= base.Ui32(v173+int32(-1025)) {
		goto L53
	} else {
		goto L69
	}
L55:
	;
	v133 = v128 + int32(1)
	v134 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
	v135 = F___isspace_1(m, v134)
	mBase = m.M
	if v135 != 0 {
		v128 = v133
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v136 = int32(1)
	switch v134&int32(255) + int32(-43) {
	case 0:
		v142 = v136
		goto L59
	default:
		v144 = v128
		v145 = v134
		v146 = v136
		goto L58
	case 2:
		goto L60
	}
L57:
	;
	goto L56
L58:
	;
	v149 = v145 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v149) {
		v167 = int32(0)
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v143 = int32(*(*int8)(unsafe.Add(mBase, uint32(v133))))
	v144 = v133
	v145 = v143
	v146 = v142
	goto L58
L60:
	;
	v142 = int32(0)
	goto L59
L61:
	;
	if v146 != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v153 = int32(0)
	v154 = v144
	v155 = v149
	goto L63
L63:
	;
	v157 = int32(10)
	v159 = v153*v157 - v155
	v160 = int32(*(*int8)(unsafe.Add(mBase, uint32(v154)+1)))
	v164 = v160 + int32(-48)
	if base.Ui32(v164) < base.Ui32(v157) {
		v153 = v159
		v154 = v154 + int32(1)
		v155 = v164
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v167 = v159
	goto L61
L65:
	;
	goto L64
L66:
	;
	v173 = int32(0) - v167
	goto L68
L67:
	;
	v173 = v167
	goto L68
L68:
	;
	goto L54
L69:
	;
	v342 = v173
	v344 = v120
	v346 = int32(0)
	goto L26
L70:
	;
	v200 = v117 + v189
	v201 = int32(0)
	goto L71
L71:
	;
	if v201 == v173 {
		goto L24
	} else {
		goto L73
	}
L72:
	;
	v342 = v173
	v344 = v332
	v346 = v194
	goto L26
L73:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v216 == int32(36) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v219 = int32(_a727)
	v222 = int32(*(*int8)(unsafe.Add(mBase, _consts[623])))
	if v222 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	if v216 != 0 {
		v342 = v173
		v344 = v201
		v346 = v194
		goto L26
	} else {
		goto L76
	}
L76:
	;
	v366 = v173
	v368 = v201
	v370 = v194
	goto L25
L77:
	;
	if v247 == int32(0) {
		v366 = v173
		v368 = v201
		v370 = v194
		goto L25
	} else {
		goto L93
	}
L78:
	;
	v223 = int32(0)
	v224 = F_strchr(m, v200, v222)
	mBase = m.M
	if v224 == v223 {
		v244 = v223
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v247 = v200
	goto L77
L80:
	;
	v247 = v244
	goto L77
L81:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, _consts[624])))
	if v227 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v228 == int32(0) {
		v244 = v223
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v247 = v224
	goto L77
L84:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, _consts[625])))
	if v231 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+2)))
	if v233 == int32(0) {
		v244 = v223
		goto L80
	} else {
		goto L87
	}
L86:
	;
	v232 = F_twobyte_strstr(m, v224, v219)
	mBase = m.M
	v247 = v232
	goto L77
L87:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _consts[626])))
	if v236 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+3)))
	if v238 == int32(0) {
		v244 = v223
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v237 = F_threebyte_strstr(m, v224, v219)
	mBase = m.M
	v247 = v237
	goto L77
L90:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, _consts[627])))
	if v241 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v243 = F_twoway_strstr(m, v224, v219)
	mBase = m.M
	v244 = v243
	goto L80
L92:
	;
	v242 = F_fourbyte_strstr(m, v224, v219)
	mBase = m.M
	v247 = v242
	goto L77
L93:
	;
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v250)
	v257 = v200 + int32(1)
	goto L95
L94:
	;
	if base.Ui32(v302+int32(-1025)) < base.Ui32(int32(-1024)) {
		v342 = v173
		v344 = v201
		v346 = v194
		goto L26
	} else {
		goto L109
	}
L95:
	;
	v262 = v257 + int32(1)
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(v257))))
	v264 = F___isspace_1(m, v263)
	mBase = m.M
	if v264 != 0 {
		v257 = v262
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v265 = int32(1)
	switch v263&int32(255) + int32(-43) {
	case 0:
		v271 = v265
		goto L99
	default:
		v273 = v257
		v274 = v263
		v275 = v265
		goto L98
	case 2:
		goto L100
	}
L97:
	;
	goto L96
L98:
	;
	v278 = v274 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v278) {
		v296 = int32(0)
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v272 = int32(*(*int8)(unsafe.Add(mBase, uint32(v262))))
	v273 = v262
	v274 = v272
	v275 = v271
	goto L98
L100:
	;
	v271 = int32(0)
	goto L99
L101:
	;
	if v275 != 0 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v282 = int32(0)
	v283 = v273
	v284 = v278
	goto L103
L103:
	;
	v286 = int32(10)
	v288 = v282*v286 - v284
	v289 = int32(*(*int8)(unsafe.Add(mBase, uint32(v283)+1)))
	v293 = v289 + int32(-48)
	if base.Ui32(v293) < base.Ui32(v286) {
		v282 = v288
		v283 = v283 + int32(1)
		v284 = v293
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v296 = v288
	goto L101
L105:
	;
	goto L104
L106:
	;
	v302 = int32(0) - v296
	goto L108
L107:
	;
	v302 = v296
	goto L108
L108:
	;
	goto L94
L109:
	;
	v307 = int32(2)
	v308 = v247 + v307
	v309 = v308 + v302
	v311 = v309 + v307
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+int32(-1)))))
	switch v314 & int32(7) {
	case 0:
		goto L115
	case 1:
		goto L114
	case 2:
		goto L113
	case 3:
		goto L112
	case 4:
		goto L111
	default:
		v323 = int32(0)
		goto L110
	}
L110:
	;
	if base.Ui32(v323) < base.Ui32(v311-v76) {
		v366 = v173
		v368 = v201
		v370 = v194
		goto L25
	} else {
		goto L116
	}
L111:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(-17))))
	v323 = v322
	goto L110
L112:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v76+int32(-9))))
	v323 = v321
	goto L110
L113:
	;
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76+int32(-5)))))
	v323 = v320
	goto L110
L114:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+int32(-3)))))
	v323 = v319
	goto L110
L115:
	;
	v323 = int32(base.Ui32(v314) >> (uint(int32(3)) % 32))
	goto L110
L116:
	;
	v328 = F_createStringObject_1(m, v308, v302)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L28
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194+v201<<(uint(int32(2))%32)))) = v328
	v332 = v201 + int32(1)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v333 != int32(13) {
		v342 = v173
		v344 = v332
		v346 = v194
		goto L26
	} else {
		goto L118
	}
L118:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v336 == int32(10) {
		v200 = v311
		v201 = v332
		goto L71
	} else {
		goto L119
	}
L119:
	;
	goto L72
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v360
	v366 = v342
	v368 = v344
	v370 = v346
	goto L25
L121:
	;
	v389 = v382
	goto L122
L122:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v370+v389<<(uint(int32(2))%32))))
	F_decrRefCount(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L28
	} else {
		goto L124
	}
L124:
	;
	v411 = v389 + int32(1)
	if v411 != v368 {
		v389 = v411
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v424 = v366
	v428 = v370
	goto L14
L126:
	;
	if v194 != 0 {
		v516 = v173
		v520 = v194
		goto L12
	} else {
		goto L127
	}
L127:
	;
	v447 = v173
	goto L13
L128:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_sdsfree(m, v76)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L28
	} else {
		goto L130
	}
L130:
	;
	v447 = v424
	goto L13
L131:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _consts[612]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+76))
	v472 = m.T0[v471].(func(*base.Module, int32, int32, int32) int32)(m, v468, v22, int32(1024))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L28
	} else {
		goto L134
	}
L132:
	;
	v516 = v447
	v520 = int32(0)
	goto L12
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v505 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L134:
	;
	if v472 < int32(1) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v476 = int32(0)
	v478 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	v479 = F_sdscatlen(m, v478, v22, v472)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L28
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[617])) = v479
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479+int32(-1)))))
	switch v484&int32(7) + int32(-3) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		goto L11
	}
L137:
	;
	if base.Ui32(v495) < base.Ui32(int32(1048577)) {
		goto L11
	} else {
		goto L140
	}
L138:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v479+int32(-17))))
	v495 = v494
	goto L137
L139:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v479+int32(-9))))
	v495 = v491
	goto L137
L140:
	;
	v500 = F_createStringObject_1(m, int32(_a1112), int32(25))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L28
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v500
	goto L4
L142:
	;
	v509 = int32(0)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v510 == v509 {
		goto L10
	} else {
		goto L144
	}
L143:
	;
	v516 = v447
	v520 = int32(0)
	goto L12
L144:
	;
	v516 = v447
	v520 = v509
	goto L12
L145:
	;
	v537 = F_sdsempty(m)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L28
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[617])) = v537
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v540 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v541 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v544 = F_objectGetVal(m, v543)
	mBase = m.M
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+int32(-1)))))
	switch v547 & int32(7) {
	case 0:
		goto L154
	case 1:
		goto L153
	case 2:
		goto L152
	case 3:
		goto L151
	case 4:
		goto L150
	default:
		v564 = int32(0)
		goto L149
	}
L149:
	;
	v566 = v516 + int32(-1)
	v568 = *(*int32)(unsafe.Add(mBase, _consts[628]))
	if v564 != v568 {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v544+int32(-17))))
	v564 = v563
	goto L149
L151:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v544+int32(-9))))
	v564 = v560
	goto L149
L152:
	;
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544+int32(-5)))))
	v564 = v557
	goto L149
L153:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+int32(-3)))))
	v564 = v554
	goto L149
L154:
	;
	v564 = int32(base.Ui32(v547) >> (uint(int32(3)) % 32))
	goto L149
L155:
	;
	v672 = int32(0)
	v673 = *(*int32)(unsafe.Add(mBase, _consts[629]))
	if v673 == v672 {
		goto L187
	} else {
		goto L188
	}
L156:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[630]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v633 = F_objectGetVal(m, v632)
	mBase = m.M
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
	if v636 != 0 {
		goto L176
	} else {
		goto L177
	}
L157:
	;
	v570 = int32(0)
	v571 = *(*int32)(unsafe.Add(mBase, _consts[630]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v573 = F_objectGetVal(m, v572)
	mBase = m.M
	v575 = *(*int32)(unsafe.Add(mBase, _consts[628]))
	if v575 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v627 == int32(0) {
		goto L155
	} else {
		goto L173
	}
L159:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v579 != 0 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v627 = int32(0)
	goto L158
L161:
	;
	v618 = F_tolower(m, v613)
	mBase = m.M
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v620 = F_tolower(m, v619)
	mBase = m.M
	v627 = v618 - v620
	goto L158
L162:
	;
	v581 = v571
	v582 = v573
	v583 = v575
	v584 = v579
	goto L165
L163:
	;
	v613 = int32(0)
	v614 = v573
	goto L161
L164:
	;
	v613 = v610 & int32(255)
	v614 = v608
	goto L161
L165:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	if v586 == int32(0) {
		v608 = v582
		v610 = v584
		goto L164
	} else {
		goto L167
	}
L166:
	;
	v608 = v602
	v610 = int32(0)
	goto L164
L167:
	;
	v590 = v583 + int32(-1)
	if v590 == int32(0) {
		v608 = v582
		v610 = v584
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v594 = v584 & int32(255)
	if v594 == v586 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v601 = int32(1)
	v602 = v582 + v601
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581)+1)))
	if v603 != 0 {
		v581 = v581 + v601
		v582 = v602
		v583 = v590
		v584 = v603
		goto L165
	} else {
		goto L172
	}
L170:
	;
	v596 = F_tolower(m, v594)
	mBase = m.M
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v598 = F_tolower(m, v597)
	mBase = m.M
	if v596 == v598 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	v608 = v582
	v610 = v600
	goto L164
L172:
	;
	goto L166
L173:
	;
	goto L156
L174:
	;
	if v668-v670 != 0 {
		goto L9
	} else {
		goto L186
	}
L175:
	;
	v668 = F_tolower(m, v664)
	mBase = m.M
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v670 = F_tolower(m, v669)
	mBase = m.M
	goto L174
L176:
	;
	v638 = v631
	v639 = v633
	v640 = v636
	goto L179
L177:
	;
	v664 = int32(0)
	v665 = v633
	goto L175
L178:
	;
	v664 = v661 & int32(255)
	v665 = v660
	goto L175
L179:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	if v642 == int32(0) {
		v660 = v639
		v661 = v640
		goto L178
	} else {
		goto L181
	}
L180:
	;
	v660 = v654
	v661 = int32(0)
	goto L178
L181:
	;
	v646 = v640 & int32(255)
	if v646 == v642 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v653 = int32(1)
	v654 = v639 + v653
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638)+1)))
	if v655 != 0 {
		v638 = v638 + v653
		v639 = v654
		v640 = v655
		goto L179
	} else {
		goto L185
	}
L183:
	;
	v648 = F_tolower(m, v646)
	mBase = m.M
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639))))
	v650 = F_tolower(m, v649)
	mBase = m.M
	if v648 == v650 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	v660 = v639
	v661 = v652
	goto L178
L185:
	;
	goto L180
L186:
	;
	goto L155
L187:
	;
	if v566 != 0 {
		goto L9
	} else {
		goto L207
	}
L188:
	;
	v676 = int32(1)
	v678 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, _consts[631]))
	if v673 == v676 {
		v740 = v678
		v741 = v678
		v742 = v678
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if v673&v676 == int32(0) {
		v768 = v741
		v769 = v742
		goto L197
	} else {
		goto L198
	}
L190:
	;
	v687 = int32(0)
	v695 = v687
	v696 = v687
	v697 = v687
	v705 = v687
	goto L191
L191:
	;
	v710 = int32(1)
	v713 = int32(12)
	v715 = v680 + (v695|v710)*v713
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+8))
	v719 = v680 + v695*v713
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+8))
	if v716|v720 != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v740 = v732
	v741 = v722
	v742 = v730
	goto L189
L193:
	;
	v722 = v710
	goto L195
L194:
	;
	v722 = v696
	goto L195
L195:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	v724 = int32(0)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	v730 = v697 + base.B2i32(v723 == v724) + base.B2i32(v727 == v724)
	v731 = int32(2)
	v732 = v695 + v731
	v734 = v705 + v731
	if v734 != v673&int32(-2) {
		v695 = v732
		v696 = v722
		v697 = v730
		v705 = v734
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v770 = int32(_a1113)
	if v566 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v760 = v680 + v740*int32(12)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)+8))
	if v761 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v762 = int32(1)
	goto L201
L200:
	;
	v762 = v741
	goto L201
L201:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	v768 = v762
	v769 = v742 + base.B2i32(v763 == int32(0))
	goto L197
L202:
	;
	if base.Ui32(v673) < base.Ui32(v566) {
		goto L9
	} else {
		goto L205
	}
L203:
	;
	if v768 != 0 {
		v1382 = v770
		goto L8
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	if base.Ui32(v769) <= base.Ui32(v566) {
		v1382 = v770
		goto L8
	} else {
		goto L206
	}
L206:
	;
	goto L9
L207:
	;
	v1382 = int32(_a1113)
	goto L8
L208:
	;
	goto L3
L209:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	if v827 != v829 {
		goto L218
	} else {
		goto L219
	}
L211:
	;
	v827 = int32(base.Ui32(v810) >> (uint(int32(3)) % 32))
	goto L210
L212:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807+int32(-3)))))
	v827 = v824
	goto L210
L213:
	;
	v821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v807+int32(-5)))))
	v827 = v821
	goto L210
L214:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v807+int32(-9))))
	v827 = v818
	goto L210
L215:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v807+int32(-17))))
	v827 = v815
	goto L210
L216:
	;
	v1056 = int32(0)
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[633]))
	if v1057 == v1056 {
		goto L270
	} else {
		goto L271
	}
L217:
	;
	v933 = int32(0)
	v934 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	if v934 == v933 {
		goto L249
	} else {
		goto L250
	}
L218:
	;
	v892 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v894 = F_objectGetVal(m, v893)
	mBase = m.M
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	if v897 != 0 {
		goto L238
	} else {
		goto L239
	}
L219:
	;
	v831 = int32(0)
	v832 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v834 = F_objectGetVal(m, v833)
	mBase = m.M
	v836 = *(*int32)(unsafe.Add(mBase, _consts[632]))
	if v836 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v888 == int32(0) {
		goto L217
	} else {
		goto L235
	}
L221:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832))))
	if v840 != 0 {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	v888 = int32(0)
	goto L220
L223:
	;
	v879 = F_tolower(m, v874)
	mBase = m.M
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	v881 = F_tolower(m, v880)
	mBase = m.M
	v888 = v879 - v881
	goto L220
L224:
	;
	v842 = v832
	v843 = v834
	v844 = v836
	v845 = v840
	goto L227
L225:
	;
	v874 = int32(0)
	v875 = v834
	goto L223
L226:
	;
	v874 = v871 & int32(255)
	v875 = v869
	goto L223
L227:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843))))
	if v847 == int32(0) {
		v869 = v843
		v871 = v845
		goto L226
	} else {
		goto L229
	}
L228:
	;
	v869 = v863
	v871 = int32(0)
	goto L226
L229:
	;
	v851 = v844 + int32(-1)
	if v851 == int32(0) {
		v869 = v843
		v871 = v845
		goto L226
	} else {
		goto L230
	}
L230:
	;
	v855 = v845 & int32(255)
	if v855 == v847 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v862 = int32(1)
	v863 = v843 + v862
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+1)))
	if v864 != 0 {
		v842 = v842 + v862
		v843 = v863
		v844 = v851
		v845 = v864
		goto L227
	} else {
		goto L234
	}
L232:
	;
	v857 = F_tolower(m, v855)
	mBase = m.M
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843))))
	v859 = F_tolower(m, v858)
	mBase = m.M
	if v857 == v859 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	v869 = v843
	v871 = v861
	goto L226
L234:
	;
	goto L228
L235:
	;
	goto L218
L236:
	;
	if v929-v931 != 0 {
		goto L216
	} else {
		goto L248
	}
L237:
	;
	v929 = F_tolower(m, v925)
	mBase = m.M
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	v931 = F_tolower(m, v930)
	mBase = m.M
	goto L236
L238:
	;
	v899 = v892
	v900 = v894
	v901 = v897
	goto L241
L239:
	;
	v925 = int32(0)
	v926 = v894
	goto L237
L240:
	;
	v925 = v922 & int32(255)
	v926 = v921
	goto L237
L241:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
	if v903 == int32(0) {
		v921 = v900
		v922 = v901
		goto L240
	} else {
		goto L243
	}
L242:
	;
	v921 = v915
	v922 = int32(0)
	goto L240
L243:
	;
	v907 = v901 & int32(255)
	if v907 == v903 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v914 = int32(1)
	v915 = v900 + v914
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899)+1)))
	if v916 != 0 {
		v899 = v899 + v914
		v900 = v915
		v901 = v916
		goto L241
	} else {
		goto L247
	}
L245:
	;
	v909 = F_tolower(m, v907)
	mBase = m.M
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
	v911 = F_tolower(m, v910)
	mBase = m.M
	if v909 == v911 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899))))
	v921 = v900
	v922 = v913
	goto L240
L247:
	;
	goto L242
L248:
	;
	goto L217
L249:
	;
	if v566 != 0 {
		goto L216
	} else {
		goto L269
	}
L250:
	;
	v937 = int32(1)
	v939 = int32(0)
	v941 = *(*int32)(unsafe.Add(mBase, _consts[636]))
	if v934 == v937 {
		v1001 = v939
		v1002 = v939
		v1003 = v939
		goto L251
	} else {
		goto L252
	}
L251:
	;
	if v934&v937 == int32(0) {
		v1029 = v1002
		v1030 = v1003
		goto L259
	} else {
		goto L260
	}
L252:
	;
	v948 = int32(0)
	v956 = v948
	v957 = v948
	v958 = v948
	v966 = v948
	goto L253
L253:
	;
	v971 = int32(1)
	v974 = int32(12)
	v976 = v941 + (v956|v971)*v974
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+8))
	v980 = v941 + v956*v974
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)+8))
	if v977|v981 != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1001 = v993
	v1002 = v983
	v1003 = v991
	goto L251
L255:
	;
	v983 = v971
	goto L257
L256:
	;
	v983 = v957
	goto L257
L257:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	v985 = int32(0)
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	v991 = v958 + base.B2i32(v984 == v985) + base.B2i32(v988 == v985)
	v992 = int32(2)
	v993 = v956 + v992
	v995 = v966 + v992
	if v995 != v934&int32(-2) {
		v956 = v993
		v957 = v983
		v958 = v991
		v966 = v995
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L254
L259:
	;
	v1031 = int32(_a1114)
	if v566 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L260:
	;
	v1021 = v941 + v1001*int32(12)
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+8))
	if v1022 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1023 = int32(1)
	goto L263
L262:
	;
	v1023 = v1002
	goto L263
L263:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1029 = v1023
	v1030 = v1003 + base.B2i32(v1024 == int32(0))
	goto L259
L264:
	;
	if base.Ui32(v934) < base.Ui32(v566) {
		goto L216
	} else {
		goto L267
	}
L265:
	;
	if v1029 != 0 {
		v1382 = v1031
		goto L8
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	if base.Ui32(v1030) <= base.Ui32(v566) {
		v1382 = v1031
		goto L8
	} else {
		goto L268
	}
L268:
	;
	goto L216
L269:
	;
	v1382 = int32(_a1114)
	goto L8
L270:
	;
	v1358 = F_sdsnew(m, int32(_a1115))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L28
	} else {
		goto L337
	}
L271:
	;
	v1079 = int32(0)
	goto L272
L272:
	;
	v1080 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	v1084 = v1081 + v1079*int32(40)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v1087 = F_objectGetVal(m, v1086)
	mBase = m.M
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+int32(-1)))))
	switch v1090 & int32(7) {
	case 0:
		goto L279
	case 1:
		goto L278
	case 2:
		goto L277
	case 3:
		goto L276
	case 4:
		goto L275
	default:
		v1107 = v1080
		goto L274
	}
L273:
	;
	goto L270
L274:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+12))
	if v1107 != v1108 {
		goto L282
	} else {
		goto L283
	}
L275:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1087+int32(-17))))
	v1107 = v1106
	goto L274
L276:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1087+int32(-9))))
	v1107 = v1103
	goto L274
L277:
	;
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087+int32(-5)))))
	v1107 = v1100
	goto L274
L278:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+int32(-3)))))
	v1107 = v1097
	goto L274
L279:
	;
	v1107 = int32(base.Ui32(v1090) >> (uint(int32(3)) % 32))
	goto L274
L280:
	;
	v1333 = v1079 + int32(1)
	v1335 = *(*int32)(unsafe.Add(mBase, _consts[633]))
	if base.Ui32(v1333) < base.Ui32(v1335) {
		v1079 = v1333
		goto L272
	} else {
		goto L336
	}
L281:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+20))
	if v1209 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L282:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+8))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v1170 = F_objectGetVal(m, v1169)
	mBase = m.M
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	if v1173 != 0 {
		goto L302
	} else {
		goto L303
	}
L283:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+8))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	v1112 = F_objectGetVal(m, v1111)
	mBase = m.M
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+12))
	if v1113 != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	if v1165 == int32(0) {
		goto L281
	} else {
		goto L299
	}
L285:
	;
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	if v1117 != 0 {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	v1165 = int32(0)
	goto L284
L287:
	;
	v1156 = F_tolower(m, v1151)
	mBase = m.M
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152))))
	v1158 = F_tolower(m, v1157)
	mBase = m.M
	v1165 = v1156 - v1158
	goto L284
L288:
	;
	v1119 = v1110
	v1120 = v1112
	v1121 = v1113
	v1122 = v1117
	goto L291
L289:
	;
	v1151 = int32(0)
	v1152 = v1112
	goto L287
L290:
	;
	v1151 = v1148 & int32(255)
	v1152 = v1146
	goto L287
L291:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	if v1124 == int32(0) {
		v1146 = v1120
		v1148 = v1122
		goto L290
	} else {
		goto L293
	}
L292:
	;
	v1146 = v1140
	v1148 = int32(0)
	goto L290
L293:
	;
	v1128 = v1121 + int32(-1)
	if v1128 == int32(0) {
		v1146 = v1120
		v1148 = v1122
		goto L290
	} else {
		goto L294
	}
L294:
	;
	v1132 = v1122 & int32(255)
	if v1132 == v1124 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1139 = int32(1)
	v1140 = v1120 + v1139
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+1)))
	if v1141 != 0 {
		v1119 = v1119 + v1139
		v1120 = v1140
		v1121 = v1128
		v1122 = v1141
		goto L291
	} else {
		goto L298
	}
L296:
	;
	v1134 = F_tolower(m, v1132)
	mBase = m.M
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	v1136 = F_tolower(m, v1135)
	mBase = m.M
	if v1134 == v1136 {
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119))))
	v1146 = v1120
	v1148 = v1138
	goto L290
L298:
	;
	goto L292
L299:
	;
	goto L282
L300:
	;
	if v1205-v1207 != 0 {
		goto L280
	} else {
		goto L312
	}
L301:
	;
	v1205 = F_tolower(m, v1201)
	mBase = m.M
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1202))))
	v1207 = F_tolower(m, v1206)
	mBase = m.M
	goto L300
L302:
	;
	v1175 = v1168
	v1176 = v1170
	v1177 = v1173
	goto L305
L303:
	;
	v1201 = int32(0)
	v1202 = v1170
	goto L301
L304:
	;
	v1201 = v1198 & int32(255)
	v1202 = v1197
	goto L301
L305:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176))))
	if v1179 == int32(0) {
		v1197 = v1176
		v1198 = v1177
		goto L304
	} else {
		goto L307
	}
L306:
	;
	v1197 = v1191
	v1198 = int32(0)
	goto L304
L307:
	;
	v1183 = v1177 & int32(255)
	if v1183 == v1179 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1190 = int32(1)
	v1191 = v1176 + v1190
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+1)))
	if v1192 != 0 {
		v1175 = v1175 + v1190
		v1176 = v1191
		v1177 = v1192
		goto L305
	} else {
		goto L311
	}
L309:
	;
	v1185 = F_tolower(m, v1183)
	mBase = m.M
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176))))
	v1187 = F_tolower(m, v1186)
	mBase = m.M
	if v1185 == v1187 {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175))))
	v1197 = v1176
	v1198 = v1189
	goto L304
L311:
	;
	goto L306
L312:
	;
	goto L281
L313:
	;
	if v566 == int32(0) {
		v1382 = v1084
		goto L8
	} else {
		goto L335
	}
L314:
	;
	v1212 = int32(1)
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+16))
	v1215 = int32(0)
	if v1209 == v1212 {
		v1275 = v1215
		v1276 = v1215
		v1277 = v1215
		goto L315
	} else {
		goto L316
	}
L315:
	;
	if v1209&v1212 == int32(0) {
		v1303 = v1276
		v1304 = v1277
		goto L323
	} else {
		goto L324
	}
L316:
	;
	v1222 = int32(0)
	v1230 = v1222
	v1231 = v1222
	v1232 = v1222
	v1240 = v1222
	goto L317
L317:
	;
	v1245 = int32(1)
	v1248 = int32(12)
	v1250 = v1214 + (v1230|v1245)*v1248
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	v1254 = v1214 + v1230*v1248
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+8))
	if v1251|v1255 != 0 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v1275 = v1267
	v1276 = v1257
	v1277 = v1265
	goto L315
L319:
	;
	v1257 = v1245
	goto L321
L320:
	;
	v1257 = v1231
	goto L321
L321:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+4))
	v1259 = int32(0)
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+4))
	v1265 = v1232 + base.B2i32(v1258 == v1259) + base.B2i32(v1262 == v1259)
	v1266 = int32(2)
	v1267 = v1230 + v1266
	v1269 = v1240 + v1266
	if v1269 != v1209&int32(-2) {
		v1230 = v1267
		v1231 = v1257
		v1232 = v1265
		v1240 = v1269
		goto L317
	} else {
		goto L322
	}
L322:
	;
	goto L318
L323:
	;
	if v566 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	v1295 = v1214 + v1275*int32(12)
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+8))
	if v1296 != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1297 = int32(1)
	goto L327
L326:
	;
	v1297 = v1276
	goto L327
L327:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	v1303 = v1297
	v1304 = v1277 + base.B2i32(v1298 == int32(0))
	goto L323
L328:
	;
	if v1081 == int32(0) {
		goto L270
	} else {
		goto L334
	}
L329:
	;
	if base.Ui32(v1209) < base.Ui32(v566) {
		goto L280
	} else {
		goto L332
	}
L330:
	;
	if v1303 != 0 {
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	if base.Ui32(v566) < base.Ui32(v1304) {
		goto L280
	} else {
		goto L333
	}
L333:
	;
	goto L328
L334:
	;
	v1382 = v1084
	goto L8
L335:
	;
	goto L280
L336:
	;
	goto L273
L337:
	;
	v1360 = F_createObject(m, int32(0), v1358)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L28
	} else {
		goto L338
	}
L338:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v1364 = F_listAddNodeTail(m, v1363, v1360)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L28
	} else {
		goto L339
	}
L339:
	;
	F_scriptingEngineDebuggerFlushLogs(m)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L28
	} else {
		goto L340
	}
L340:
	;
	v1403 = int32(1)
	goto L7
L341:
	;
	v1394 = m.T0[v1388].(func(*base.Module, int32, int32, int32) int32)(m, v520, v516, v1393)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L28
	} else {
		goto L343
	}
L342:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, _consts[614]))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1391)+8))
	v1393 = v1392
	goto L341
L343:
	;
	v1403 = base.B2i32(v1394 == int32(1))
	goto L7
L344:
	;
	F_valkey_free(m, v520)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L28
	} else {
		goto L350
	}
L345:
	;
	v1424 = v1417
	goto L346
L346:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v520+v1424<<(uint(int32(2))%32))))
	F_decrRefCount(m, v1442)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L28
	} else {
		goto L348
	}
L347:
	;
	goto L344
L348:
	;
	v1446 = v1424 + int32(1)
	if v1446 != v516 {
		v1424 = v1446
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	if v1403 == int32(0) {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1472 != 0 {
		goto L3
	} else {
		goto L352
	}
L352:
	;
	v34 = int32(0)
	goto L5
L353:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scriptingEngineDebuggerStartSession(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v12 = v10 & int32(536870912)
	*(*int32)(unsafe.Add(mBase, _consts[621])) = int32(base.Ui32(v12)>>(uint(int32(29))%32)) ^ int32(1)
	if v12 != 0 {
		v111 = *(*int32)(unsafe.Add(mBase, _consts[15]))
		if int32(2) < v111 {
			v124 = *(*int32)(unsafe.Add(mBase, _consts[612]))
			v125 = F_connBlock(m, v124)
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				v128 = *(*int32)(unsafe.Add(mBase, _consts[612]))
				v130 = F_connSendTimeout(m, v128, int64(5000))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					v132 = int32(1)
					v133 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[622])) = v132
					v137 = *(*int32)(unsafe.Add(mBase, _consts[614]))
					v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
					F_scriptingEngineCallDebuggerStart(m, v137, v133, v140)
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						v150 = v132
						m.G0 = v7 + int32(176)
						return v150
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a184)
			F__serverLog(m, int32(2), int32(_a1105), v7+int32(32))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, _consts[612]))
				v125 = F_connBlock(m, v124)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					v128 = *(*int32)(unsafe.Add(mBase, _consts[612]))
					v130 = F_connSendTimeout(m, v128, int64(5000))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						v132 = int32(1)
						v133 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[622])) = v132
						v137 = *(*int32)(unsafe.Add(mBase, _consts[614]))
						v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
						F_scriptingEngineCallDebuggerStart(m, v137, v133, v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v150 = v132
							m.G0 = v7 + int32(176)
							return v150
						}
					}
				}
			}
		}
	} else {
		v19 = F_serverFork(m, int32(3))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			switch v19 + int32(1) {
			case 0:
				v26 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v27 = F___strerror_l(m, v26, v26)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
				F_addReplyErrorFormat(m, l0, int32(_a1106), v7)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v150 = int32(0)
					m.G0 = v7 + int32(176)
					return v150
				}
			case 1:
				*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(-2)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+168)) = int32(0)
				v44 = v7 + int32(36)
				if v44 == int32(0) {
				} else {
					v67 = F___memcpy(m, int32(9119284), v44, int32(140))
					mBase = m.M
				}
				v72 = v7 + int32(36)
				if v72 == int32(0) {
				} else {
					v95 = F___memcpy(m, int32(9117464), v72, int32(140))
					mBase = m.M
				}
				v99 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(2) < v99 {
					v124 = *(*int32)(unsafe.Add(mBase, _consts[612]))
					v125 = F_connBlock(m, v124)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v128 = *(*int32)(unsafe.Add(mBase, _consts[612]))
						v130 = F_connSendTimeout(m, v128, int64(5000))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							v132 = int32(1)
							v133 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[622])) = v132
							v137 = *(*int32)(unsafe.Add(mBase, _consts[614]))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
							F_scriptingEngineCallDebuggerStart(m, v137, v133, v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v150 = v132
								m.G0 = v7 + int32(176)
								return v150
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a184)
					F__serverLog(m, int32(2), int32(_a1107), v7+int32(16))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v124 = *(*int32)(unsafe.Add(mBase, _consts[612]))
						v125 = F_connBlock(m, v124)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v128 = *(*int32)(unsafe.Add(mBase, _consts[612]))
							v130 = F_connSendTimeout(m, v128, int64(5000))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								v132 = int32(1)
								v133 = int32(0)
								*(*int32)(unsafe.Add(mBase, _consts[622])) = v132
								v137 = *(*int32)(unsafe.Add(mBase, _consts[614]))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
								F_scriptingEngineCallDebuggerStart(m, v137, v133, v140)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									v150 = v132
									m.G0 = v7 + int32(176)
									return v150
								}
							}
						}
					}
				}
			default:
				v143 = int32(0)
				v145 = *(*int32)(unsafe.Add(mBase, _consts[616]))
				v146 = F_listAddNodeTail(m, v145, v19)
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					F_freeClientAsync(m, l0)
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						v150 = v143
						m.G0 = v7 + int32(176)
						return v150
					}
				}
			}
		}
	}
}
func F_scriptingEngineManagerGetNumEngines(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	return v4 + v5
}
func F_scriptingEngineManagerInit(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = F_dictCreate(m, int32(_a1091))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		*(*int64)(unsafe.Add(mBase, _consts[612])) = int64(0)
		*(*int32)(unsafe.Add(mBase, _consts[613])) = v3
		*(*int32)(unsafe.Add(mBase, _consts[614])) = v7
		v16 = F_listCreate(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[615])) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1000)
			v22 = F_listCreate(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[616])) = v22
				v26 = F_sdsempty(m)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[617])) = v26
					return int32(0)
				}
			}
		}
	}
}
func F_scriptingEngineManagerUnregister(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	v12 = F_dictUnlink(m, v11, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			F_functionsRemoveLibFromEngine(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_evalRemoveScriptsFromEngine(m, v26)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v26 == int32(0) {
						F__serverAssert(m, int32(_a1092), int32(_a1093), int32(261))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						if v33 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
							v36 = int32(0)
							F_moduleScriptingEngineInitContext(m, v35, v33, v36, v36)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = v35
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
								m.T0[v45].(func(*base.Module, int32, int32, int32, int32))(m, v8+int32(16), v40, v43, int32(2))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
									if v48 == int32(0) {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8))))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
										v68 = v61 + int32(-1)
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
										v71 = v69 & int32(7)
										switch v71 {
										case 0:
											v72 = F_zmalloc_usable_size(m, v68)
											mBase = m.M
											v101 = v72
										case 1:
											v77 = int32(4)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 2:
											v77 = int32(6)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 3:
											v77 = int32(10)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 4:
											v77 = int32(18)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										default:
											v77 = int32(1)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										}
										v102 = int32(0)
										v104 = *(*int32)(unsafe.Add(mBase, _consts[618]))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										*(*int32)(unsafe.Add(mBase, _consts[618])) = v104 - (v101 + (v56&int32(2147483647) + int32(8)) + v106)
										v111 = F___pthread_mutex_unlock(m, int32(_a1094))
										mBase = m.M
										F_bioDrainWorker(m, int32(2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											v116 = F___pthread_mutex_lock(m, int32(_a1094))
											mBase = m.M
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
											F_sdsfree(m, v117)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
												if v120 == int32(0) {
													F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													F_valkey_free(m, v120)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int32(0)
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
														if v125 == int32(0) {
															F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															F_valkey_free(m, v125)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return int32(0)
															} else {
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
																if v130 == int32(0) {
																	F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		F_abort(m)
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	F_valkey_free(m, v130)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int32(0)
																	} else {
																		F_valkey_free(m, v26)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			v137 = int32(0)
																			v139 = *(*int32)(unsafe.Add(mBase, _consts[613]))
																			F_dictFreeUnlinkedEntry(m, v139, v12)
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v143 = v137
																				m.G0 = v8 + int32(32)
																				return v143
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
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
										F_moduleFreeContext(m, v51)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8))))
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
											v68 = v61 + int32(-1)
											v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
											v71 = v69 & int32(7)
											switch v71 {
											case 0:
												v72 = F_zmalloc_usable_size(m, v68)
												mBase = m.M
												v101 = v72
											case 1:
												v77 = int32(4)
												switch v71 {
												case 0:
													v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
												case 1:
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
													v101 = v77 + v84
												case 2:
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
													v101 = v77 + v88
												case 3:
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
													v101 = v77 + v92
												case 4:
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
													v97 = v96
													v101 = v77 + v97
												default:
													v97 = int32(0)
													v101 = v77 + v97
												}
											case 2:
												v77 = int32(6)
												switch v71 {
												case 0:
													v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
												case 1:
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
													v101 = v77 + v84
												case 2:
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
													v101 = v77 + v88
												case 3:
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
													v101 = v77 + v92
												case 4:
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
													v97 = v96
													v101 = v77 + v97
												default:
													v97 = int32(0)
													v101 = v77 + v97
												}
											case 3:
												v77 = int32(10)
												switch v71 {
												case 0:
													v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
												case 1:
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
													v101 = v77 + v84
												case 2:
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
													v101 = v77 + v88
												case 3:
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
													v101 = v77 + v92
												case 4:
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
													v97 = v96
													v101 = v77 + v97
												default:
													v97 = int32(0)
													v101 = v77 + v97
												}
											case 4:
												v77 = int32(18)
												switch v71 {
												case 0:
													v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
												case 1:
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
													v101 = v77 + v84
												case 2:
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
													v101 = v77 + v88
												case 3:
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
													v101 = v77 + v92
												case 4:
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
													v97 = v96
													v101 = v77 + v97
												default:
													v97 = int32(0)
													v101 = v77 + v97
												}
											default:
												v77 = int32(1)
												switch v71 {
												case 0:
													v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
												case 1:
													v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
													v101 = v77 + v84
												case 2:
													v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
													v101 = v77 + v88
												case 3:
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
													v101 = v77 + v92
												case 4:
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
													v97 = v96
													v101 = v77 + v97
												default:
													v97 = int32(0)
													v101 = v77 + v97
												}
											}
											v102 = int32(0)
											v104 = *(*int32)(unsafe.Add(mBase, _consts[618]))
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
											*(*int32)(unsafe.Add(mBase, _consts[618])) = v104 - (v101 + (v56&int32(2147483647) + int32(8)) + v106)
											v111 = F___pthread_mutex_unlock(m, int32(_a1094))
											mBase = m.M
											F_bioDrainWorker(m, int32(2))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												v116 = F___pthread_mutex_lock(m, int32(_a1094))
												mBase = m.M
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
												F_sdsfree(m, v117)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
													if v120 == int32(0) {
														F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														F_valkey_free(m, v120)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
															if v125 == int32(0) {
																F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return int32(0)
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																F_valkey_free(m, v125)
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
																	if v130 == int32(0) {
																		F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return int32(0)
																		} else {
																			F_abort(m)
																			mBase = m.M
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		F_valkey_free(m, v130)
																		mBase = m.M
																		v134 = m.ExcPending
																		if v134 != 0 {
																			return int32(0)
																		} else {
																			F_valkey_free(m, v26)
																			mBase = m.M
																			v136 = m.ExcPending
																			if v136 != 0 {
																				return int32(0)
																			} else {
																				v137 = int32(0)
																				v139 = *(*int32)(unsafe.Add(mBase, _consts[613]))
																				F_dictFreeUnlinkedEntry(m, v139, v12)
																				mBase = m.M
																				v141 = m.ExcPending
																				if v141 != 0 {
																					return int32(0)
																				} else {
																					v143 = v137
																					m.G0 = v8 + int32(32)
																					return v143
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
							v40 = int32(0)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
							m.T0[v45].(func(*base.Module, int32, int32, int32, int32))(m, v8+int32(16), v40, v43, int32(2))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
								if v48 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8))))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
									v68 = v61 + int32(-1)
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
									v71 = v69 & int32(7)
									switch v71 {
									case 0:
										v72 = F_zmalloc_usable_size(m, v68)
										mBase = m.M
										v101 = v72
									case 1:
										v77 = int32(4)
										switch v71 {
										case 0:
											v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
										case 1:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
											v101 = v77 + v84
										case 2:
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
											v101 = v77 + v88
										case 3:
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
											v101 = v77 + v92
										case 4:
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
											v97 = v96
											v101 = v77 + v97
										default:
											v97 = int32(0)
											v101 = v77 + v97
										}
									case 2:
										v77 = int32(6)
										switch v71 {
										case 0:
											v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
										case 1:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
											v101 = v77 + v84
										case 2:
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
											v101 = v77 + v88
										case 3:
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
											v101 = v77 + v92
										case 4:
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
											v97 = v96
											v101 = v77 + v97
										default:
											v97 = int32(0)
											v101 = v77 + v97
										}
									case 3:
										v77 = int32(10)
										switch v71 {
										case 0:
											v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
										case 1:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
											v101 = v77 + v84
										case 2:
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
											v101 = v77 + v88
										case 3:
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
											v101 = v77 + v92
										case 4:
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
											v97 = v96
											v101 = v77 + v97
										default:
											v97 = int32(0)
											v101 = v77 + v97
										}
									case 4:
										v77 = int32(18)
										switch v71 {
										case 0:
											v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
										case 1:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
											v101 = v77 + v84
										case 2:
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
											v101 = v77 + v88
										case 3:
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
											v101 = v77 + v92
										case 4:
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
											v97 = v96
											v101 = v77 + v97
										default:
											v97 = int32(0)
											v101 = v77 + v97
										}
									default:
										v77 = int32(1)
										switch v71 {
										case 0:
											v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
										case 1:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
											v101 = v77 + v84
										case 2:
											v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
											v101 = v77 + v88
										case 3:
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
											v101 = v77 + v92
										case 4:
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
											v97 = v96
											v101 = v77 + v97
										default:
											v97 = int32(0)
											v101 = v77 + v97
										}
									}
									v102 = int32(0)
									v104 = *(*int32)(unsafe.Add(mBase, _consts[618]))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
									*(*int32)(unsafe.Add(mBase, _consts[618])) = v104 - (v101 + (v56&int32(2147483647) + int32(8)) + v106)
									v111 = F___pthread_mutex_unlock(m, int32(_a1094))
									mBase = m.M
									F_bioDrainWorker(m, int32(2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v116 = F___pthread_mutex_lock(m, int32(_a1094))
										mBase = m.M
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
										F_sdsfree(m, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
											if v120 == int32(0) {
												F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_valkey_free(m, v120)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
													if v125 == int32(0) {
														F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														F_valkey_free(m, v125)
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
															if v130 == int32(0) {
																F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return int32(0)
																} else {
																	F_abort(m)
																	mBase = m.M
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																F_valkey_free(m, v130)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	F_valkey_free(m, v26)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		v137 = int32(0)
																		v139 = *(*int32)(unsafe.Add(mBase, _consts[613]))
																		F_dictFreeUnlinkedEntry(m, v139, v12)
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return int32(0)
																		} else {
																			v143 = v137
																			m.G0 = v8 + int32(32)
																			return v143
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
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
									F_moduleFreeContext(m, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(-8))))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
										v68 = v61 + int32(-1)
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
										v71 = v69 & int32(7)
										switch v71 {
										case 0:
											v72 = F_zmalloc_usable_size(m, v68)
											mBase = m.M
											v101 = v72
										case 1:
											v77 = int32(4)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 2:
											v77 = int32(6)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 3:
											v77 = int32(10)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										case 4:
											v77 = int32(18)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										default:
											v77 = int32(1)
											switch v71 {
											case 0:
												v101 = v77 + int32(base.Ui32(v69)>>(uint(int32(3))%32))
											case 1:
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(-2)))))
												v101 = v77 + v84
											case 2:
												v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61+int32(-3)))))
												v101 = v77 + v88
											case 3:
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-5))))
												v101 = v77 + v92
											case 4:
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(-9))))
												v97 = v96
												v101 = v77 + v97
											default:
												v97 = int32(0)
												v101 = v77 + v97
											}
										}
										v102 = int32(0)
										v104 = *(*int32)(unsafe.Add(mBase, _consts[618]))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
										*(*int32)(unsafe.Add(mBase, _consts[618])) = v104 - (v101 + (v56&int32(2147483647) + int32(8)) + v106)
										v111 = F___pthread_mutex_unlock(m, int32(_a1094))
										mBase = m.M
										F_bioDrainWorker(m, int32(2))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											v116 = F___pthread_mutex_lock(m, int32(_a1094))
											mBase = m.M
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
											F_sdsfree(m, v117)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
												if v120 == int32(0) {
													F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													F_valkey_free(m, v120)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int32(0)
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
														if v125 == int32(0) {
															F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															F_valkey_free(m, v125)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return int32(0)
															} else {
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
																if v130 == int32(0) {
																	F__serverAssert(m, int32(_a1095), int32(_a1093), int32(206))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		F_abort(m)
																		mBase = m.M
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	F_valkey_free(m, v130)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int32(0)
																	} else {
																		F_valkey_free(m, v26)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			v137 = int32(0)
																			v139 = *(*int32)(unsafe.Add(mBase, _consts[613]))
																			F_dictFreeUnlinkedEntry(m, v139, v12)
																			mBase = m.M
																			v141 = m.ExcPending
																			if v141 != 0 {
																				return int32(0)
																			} else {
																				v143 = v137
																				m.G0 = v8 + int32(32)
																				return v143
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
		} else {
			v16 = int32(-1)
			v18 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v18 {
				v143 = v16
				m.G0 = v8 + int32(32)
				return v143
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F__serverLog(m, int32(3), int32(_a1096), v8)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v143 = v16
					m.G0 = v8 + int32(32)
					return v143
				}
			}
		}
	}
}
