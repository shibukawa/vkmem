package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_bioCreateSaveRDBToDiskJob(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v5 = F_allocBioJob(m, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(4)
		v11 = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[79]))
		*(*int32)(unsafe.Add(mBase, _consts[79])) = v13 + int32(1)
		F_bioExecuteJob(m, v5)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bioDrainWorker(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = l0 << (uint(int32(2)) % 32)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[77])))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v11 = F_usleep(m, int32(100))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[77])))
	if v13 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_bioExecuteJob(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		switch v10 + int32(-1) {
		case 0, 2:
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v67 = F_fsync(m, v66)
			mBase = m.M
			if v67 != int32(-1) {
				v98 = int32(_a44)
				*(*int32)(unsafe.Add(mBase, _consts[76])) = int32(0)
				v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, _consts[32])) = v102
				v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
				if v106&int32(2) == int32(0) {
				} else {
				}
				if v10 != int32(3) {
				} else {
					v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v134 = F_close(m, v133)
					mBase = m.M
				}
				F_valkey_free(m, l0)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					v163 = v10 << (uint(int32(2)) % 32)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
					*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
					m.G0 = v8 + int32(64)
					return
				}
			} else {
				v70 = int32(9116376)
				v71 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				if v71 == int32(8) {
					v98 = int32(_a44)
					*(*int32)(unsafe.Add(mBase, _consts[76])) = int32(0)
					v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, _consts[32])) = v102
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					if v106&int32(2) == int32(0) {
					} else {
					}
					if v10 != int32(3) {
					} else {
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v134 = F_close(m, v133)
						mBase = m.M
					}
					F_valkey_free(m, l0)
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						v163 = v10 << (uint(int32(2)) % 32)
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
						*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
						m.G0 = v8 + int32(64)
						return
					}
				} else {
					if v71 == int32(28) {
						v98 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[76])) = int32(0)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, _consts[32])) = v102
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						if v106&int32(2) == int32(0) {
						} else {
						}
						if v10 != int32(3) {
						} else {
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v134 = F_close(m, v133)
							mBase = m.M
						}
						F_valkey_free(m, l0)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							v163 = v10 << (uint(int32(2)) % 32)
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
							m.G0 = v8 + int32(64)
							return
						}
					} else {
						v76 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[78])) = v71
						v79 = *(*int32)(unsafe.Add(mBase, _consts[76]))
						*(*int32)(unsafe.Add(mBase, _consts[76])) = int32(-1)
						if v79 != 0 {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
							if v106&int32(2) == int32(0) {
							} else {
							}
							if v10 != int32(3) {
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v134 = F_close(m, v133)
								mBase = m.M
							}
							F_valkey_free(m, l0)
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return
							} else {
								v163 = v10 << (uint(int32(2)) % 32)
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
								*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
								m.G0 = v8 + int32(64)
								return
							}
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							if int32(3) < v84 {
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
								if v106&int32(2) == int32(0) {
								} else {
								}
								if v10 != int32(3) {
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v134 = F_close(m, v133)
									mBase = m.M
								}
								F_valkey_free(m, l0)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									v163 = v10 << (uint(int32(2)) % 32)
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
									*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
									m.G0 = v8 + int32(64)
									return
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, _consts[5]))
								v88 = F___strerror_l(m, v87, v87)
								mBase = m.M
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v88
								F__serverLog(m, int32(3), int32(_a113), v8+int32(48))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
									if v106&int32(2) == int32(0) {
									} else {
									}
									if v10 != int32(3) {
									} else {
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v134 = F_close(m, v133)
										mBase = m.M
									}
									F_valkey_free(m, l0)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										v163 = v10 << (uint(int32(2)) % 32)
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
										*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
										m.G0 = v8 + int32(64)
										return
									}
								}
							}
						}
					}
				}
			}
		case 1:
			v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			m.T0[v155].(func(*base.Module, int32))(m, l0+int32(8))
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					v163 = v10 << (uint(int32(2)) % 32)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
					*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
					m.G0 = v8 + int32(64)
					return
				}
			}
		case 3:
			v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_replicaReceiveRDBFromPrimaryToDisk(m, v135, v136)
			mBase = m.M
			v138 = m.ExcPending
			if v138 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					v163 = v10 << (uint(int32(2)) % 32)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
					*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
					m.G0 = v8 + int32(64)
					return
				}
			}
		case 4:
			F__serverPanic_1(m, int32(_a177), int32(323), int32(_a178), int32(0))
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		default:
			F__serverPanic_1(m, int32(_a177), int32(326), int32(_a179), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v11&int32(1) == int32(0) {
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v39&int32(2) == int32(0) {
			} else {
			}
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v63 = F_close(m, v62)
			mBase = m.M
			F_valkey_free(m, l0)
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return
			} else {
				v163 = v10 << (uint(int32(2)) % 32)
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
				*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
				m.G0 = v8 + int32(64)
				return
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = F_fsync(m, v16)
			mBase = m.M
			if v17 != int32(-1) {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
				if v39&int32(2) == int32(0) {
				} else {
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v63 = F_close(m, v62)
				mBase = m.M
				F_valkey_free(m, l0)
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					v163 = v10 << (uint(int32(2)) % 32)
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
					*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
					m.G0 = v8 + int32(64)
					return
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				if v21 == int32(8) {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					if v39&int32(2) == int32(0) {
					} else {
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v63 = F_close(m, v62)
					mBase = m.M
					F_valkey_free(m, l0)
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						v163 = v10 << (uint(int32(2)) % 32)
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
						*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
						m.G0 = v8 + int32(64)
						return
					}
				} else {
					if v21 == int32(28) {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
						if v39&int32(2) == int32(0) {
						} else {
						}
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v63 = F_close(m, v62)
						mBase = m.M
						F_valkey_free(m, l0)
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							v163 = v10 << (uint(int32(2)) % 32)
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
							m.G0 = v8 + int32(64)
							return
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, _consts[6]))
						if int32(3) < v27 {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
							if v39&int32(2) == int32(0) {
							} else {
							}
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v63 = F_close(m, v62)
							mBase = m.M
							F_valkey_free(m, l0)
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return
							} else {
								v163 = v10 << (uint(int32(2)) % 32)
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
								*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
								m.G0 = v8 + int32(64)
								return
							}
						} else {
							v30 = F___strerror_l(m, v21, v21)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v30
							F__serverLog(m, int32(3), int32(_a113), v8+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
								if v39&int32(2) == int32(0) {
								} else {
								}
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v63 = F_close(m, v62)
								mBase = m.M
								F_valkey_free(m, l0)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									v163 = v10 << (uint(int32(2)) % 32)
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77])))
									*(*int32)(unsafe.Add(mBase, uint32(v163)+uint32(_consts[77]))) = v166 + int32(-1)
									m.G0 = v8 + int32(64)
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
