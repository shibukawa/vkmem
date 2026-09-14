package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___uflow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(-1)
	v9 = F___toread(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v22 = v8
			m.G0 = v6 + int32(16)
			return v22
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v17 = m.T0[v16].(func(*base.Module, int32, int32, int32) int32)(m, l0, v6+int32(15), int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != int32(1) {
					v22 = v8
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
					v22 = v21
				}
				m.G0 = v6 + int32(16)
				return v22
			}
		}
	}
}
func F___unlock(m *base.Module, l0 int32) {
	return
}
func F___unlockfile(m *base.Module, l0 int32) {
	return
}
func F_uintCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return v3 - v4
}
func F_unblockClient(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	switch v6 + int32(-1) {
	case 0, 3, 4:
		F_unblockClientWaitingData(m, l0)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
			if v40&int32(2) != 0 {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v50&int32(64) != 0 {
				} else {
					v53 = int32(_a20)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
					*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
				}
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v63 = v61 << (uint(int32(2)) % 32)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
				*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
				v73 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
				F_removeClientFromTimeoutTable(m, l0)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					if l1 == int32(0) {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v81&int32(128) != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
							v89 = F_listAddNodeTail(m, v88, l0)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v44 == int32(7) {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v50&int32(64) != 0 {
					} else {
						v53 = int32(_a20)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
					}
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v63 = v61 << (uint(int32(2)) % 32)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
					F_removeClientFromTimeoutTable(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v81&int32(128) != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
								v89 = F_listAddNodeTail(m, v88, l0)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					F_resetClient(m, l0)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v50&int32(64) != 0 {
						} else {
							v53 = int32(_a20)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
						}
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v63 = v61 << (uint(int32(2)) % 32)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
						F_removeClientFromTimeoutTable(m, l0)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if l1 == int32(0) {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v81&int32(128) != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
									v89 = F_listAddNodeTail(m, v88, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
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
	case 1:
		F_unblockClientWaitingReplicas(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
			if v40&int32(2) != 0 {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v50&int32(64) != 0 {
				} else {
					v53 = int32(_a20)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
					*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
				}
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v63 = v61 << (uint(int32(2)) % 32)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
				*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
				v73 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
				F_removeClientFromTimeoutTable(m, l0)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					if l1 == int32(0) {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v81&int32(128) != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
							v89 = F_listAddNodeTail(m, v88, l0)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				if v44 == int32(7) {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v50&int32(64) != 0 {
					} else {
						v53 = int32(_a20)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
					}
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v63 = v61 << (uint(int32(2)) % 32)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
					F_removeClientFromTimeoutTable(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v81&int32(128) != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
								v89 = F_listAddNodeTail(m, v88, l0)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					F_resetClient(m, l0)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v50&int32(64) != 0 {
						} else {
							v53 = int32(_a20)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
						}
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v63 = v61 << (uint(int32(2)) % 32)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
						F_removeClientFromTimeoutTable(m, l0)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if l1 == int32(0) {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v81&int32(128) != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
									v89 = F_listAddNodeTail(m, v88, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
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
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
		if v13 == int32(0) {
			F_unblockClientFromModule(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
				if v40&int32(2) != 0 {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v50&int32(64) != 0 {
					} else {
						v53 = int32(_a20)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
					}
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v63 = v61 << (uint(int32(2)) % 32)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
					F_removeClientFromTimeoutTable(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v81&int32(128) != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
								v89 = F_listAddNodeTail(m, v88, l0)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					if v44 == int32(7) {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v50&int32(64) != 0 {
						} else {
							v53 = int32(_a20)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
						}
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v63 = v61 << (uint(int32(2)) % 32)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
						F_removeClientFromTimeoutTable(m, l0)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if l1 == int32(0) {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v81&int32(128) != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
									v89 = F_listAddNodeTail(m, v88, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						F_resetClient(m, l0)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
							if v50&int32(64) != 0 {
							} else {
								v53 = int32(_a20)
								v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
								*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
							}
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							v63 = v61 << (uint(int32(2)) % 32)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
							*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
							F_removeClientFromTimeoutTable(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								if l1 == int32(0) {
									return
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v81&int32(128) != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
										v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
										v89 = F_listAddNodeTail(m, v88, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
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
		} else {
			F_unblockClientWaitingData(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_unblockClientFromModule(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
					if v40&int32(2) != 0 {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v50&int32(64) != 0 {
						} else {
							v53 = int32(_a20)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
						}
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v63 = v61 << (uint(int32(2)) % 32)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
						F_removeClientFromTimeoutTable(m, l0)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if l1 == int32(0) {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v81&int32(128) != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
									v89 = F_listAddNodeTail(m, v88, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v44 == int32(7) {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
							if v50&int32(64) != 0 {
							} else {
								v53 = int32(_a20)
								v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
								*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
							}
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							v63 = v61 << (uint(int32(2)) % 32)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
							*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
							F_removeClientFromTimeoutTable(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								if l1 == int32(0) {
									return
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v81&int32(128) != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
										v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
										v89 = F_listAddNodeTail(m, v88, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							F_resetClient(m, l0)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
								if v50&int32(64) != 0 {
								} else {
									v53 = int32(_a20)
									v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
									*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
								}
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
								v63 = v61 << (uint(int32(2)) % 32)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
								*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
								v73 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
								F_removeClientFromTimeoutTable(m, l0)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									if l1 == int32(0) {
										return
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v81&int32(128) != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
											v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
											v89 = F_listAddNodeTail(m, v88, l0)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
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
	case 5:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		if v20 == int32(0) {
			F__serverAssert(m, int32(_a185), int32(_a184), int32(226))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[76]))
			F_listDelNode(m, v24, v20)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = int32(0)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
				if v40&int32(2) != 0 {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v50&int32(64) != 0 {
					} else {
						v53 = int32(_a20)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
					}
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v63 = v61 << (uint(int32(2)) % 32)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
					F_removeClientFromTimeoutTable(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v81&int32(128) != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
								v89 = F_listAddNodeTail(m, v88, l0)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					if v44 == int32(7) {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
						if v50&int32(64) != 0 {
						} else {
							v53 = int32(_a20)
							v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
							*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
						}
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v63 = v61 << (uint(int32(2)) % 32)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
						v73 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
						F_removeClientFromTimeoutTable(m, l0)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if l1 == int32(0) {
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v81&int32(128) != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
									v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
									v89 = F_listAddNodeTail(m, v88, l0)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						F_resetClient(m, l0)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
							if v50&int32(64) != 0 {
							} else {
								v53 = int32(_a20)
								v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
								*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
							}
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
							v63 = v61 << (uint(int32(2)) % 32)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
							*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
							F_removeClientFromTimeoutTable(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								if l1 == int32(0) {
									return
								} else {
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v81&int32(128) != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
										v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
										v89 = F_listAddNodeTail(m, v88, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
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
	case 6:
		v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
		if v40&int32(2) != 0 {
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
			if v50&int32(64) != 0 {
			} else {
				v53 = int32(_a20)
				v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
				*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
			}
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v63 = v61 << (uint(int32(2)) % 32)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
			*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
			v73 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
			*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
			F_removeClientFromTimeoutTable(m, l0)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				if l1 == int32(0) {
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					if v81&int32(128) != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
						v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
						v89 = F_listAddNodeTail(m, v88, l0)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
			if v44 == int32(7) {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
				if v50&int32(64) != 0 {
				} else {
					v53 = int32(_a20)
					v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
					*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
				}
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v63 = v61 << (uint(int32(2)) % 32)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
				*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
				v73 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
				F_removeClientFromTimeoutTable(m, l0)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					if l1 == int32(0) {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						if v81&int32(128) != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
							v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
							v89 = F_listAddNodeTail(m, v88, l0)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_resetClient(m, l0)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
					if v50&int32(64) != 0 {
					} else {
						v53 = int32(_a20)
						v55 = *(*int32)(unsafe.Add(mBase, _consts[73]))
						*(*int32)(unsafe.Add(mBase, _consts[73])) = v55 + int32(-1)
					}
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
					v63 = v61 << (uint(int32(2)) % 32)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74])))
					*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[74]))) = v65 + int32(-1)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v69 & int32(-17)
					v73 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v73
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v73
					F_removeClientFromTimeoutTable(m, l0)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if l1 == int32(0) {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							if v81&int32(128) != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v81 | int32(128)
								v88 = *(*int32)(unsafe.Add(mBase, _consts[75]))
								v89 = F_listAddNodeTail(m, v88, l0)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
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
	default:
		F__serverPanic_1(m, int32(_a184), int32(232), int32(_a186), int32(0))
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
	}
}
func F_unblockClientOnError(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v7 = int32(0)
		F_updateStatsOnUnblock(m, l0, v7, v7, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
			if v12&int32(2) == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v12 & int32(-3)
			}
			F_unblockClient(m, l0, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_addReplyError(m, l0, l1)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = int32(0)
			F_updateStatsOnUnblock(m, l0, v7, v7, int32(1))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
				if v12&int32(2) == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v12 & int32(-3)
				}
				F_unblockClient(m, l0, int32(1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_unblockClientWaitingReplicas(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	if v4 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[80]))
		F_listDelNode(m, v12, v4)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v16
			F_updateStatsOnUnblock(m, l0, v16, v16, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F__serverAssert(m, int32(_a2009), int32(_a1913), int32(5106))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
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
func F_unblockPostponedClients(m *base.Module) {
	mBase := m.M
	_ = mBase
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v10 = v5 + int32(8)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
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
	m.G0 = v5 + int32(16)
	return
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
	F_unblockClient(m, v33, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	return
L10:
	;
	v38 = v5 + int32(8)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v40 != 0 {
		v32 = v40
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40+base.B2i32(v43 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v49
	goto L12
L14:
	;
	goto L8
}
func F_unlinkClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v261&int32(4194304) == int32(0) {
		v279 = v261
		goto L58
	} else {
		goto L59
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v19 != l0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[14])) = int32(0)
	goto L2
L6:
	;
	goto L2
L7:
	;
	v82 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v87 == v82 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v30 = int64(56)
	v32 = int64(65280)
	v34 = int64(40)
	v37 = int64(16711680)
	v39 = int64(24)
	v41 = int64(4278190080)
	v43 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v29<<(uint(v30)%64) | v29&v32<<(uint(v34)%64) | (v29&v37<<(uint(v39)%64) | v29&v41<<(uint(v43)%64)) | (int64(base.Ui64(v29)>>(uint(v43)%64))&v41 | int64(base.Ui64(v29)>>(uint(v39)%64))&v37 | (int64(base.Ui64(v29)>>(uint(v34)%64))&v32 | int64(base.Ui64(v29)>>(uint(v30)%64))))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v68 = int32(8)
	v72 = F_raxRemove(m, v67, v13+v68, v68, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	F_listDelNode(m, v75, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = int32(0)
	goto L7
L12:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v115 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	goto L12
L14:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v90 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	v96 = int32(0)
	goto L17
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(0)
	goto L13
L17:
	;
	v102 = v93 + v96<<(uint(int32(2))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v103 == l0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v106 = v96 + int32(1)
	if v106 == v90 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v96 = v106
	goto L17
L21:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[222]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v203 != v204 {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	if v119 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v122&int32(2) == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v127 != int32(7) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	if v131 < int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v182 {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v134 = int32(0)
	v138 = v134
	v140 = v119
	v141 = v131
	v142 = v134
	goto L28
L28:
	;
	v147 = v138 << (uint(int32(2)) % 32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v140+v147)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v149 == v150 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v167 != 0 {
		goto L21
	} else {
		goto L35
	}
L30:
	;
	v167 = v142 + v166
	v169 = v138 + int32(1)
	if v169 < v165 {
		v138 = v169
		v140 = v164
		v141 = v165
		v142 = v167
		goto L28
	} else {
		goto L34
	}
L31:
	;
	F_rdbPipeWriteHandlerConnRemoved(m, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v164 = v140
	v165 = v141
	v166 = base.B2i32(v149 != int32(0))
	goto L30
L33:
	;
	v156 = int32(0)
	v157 = int32(_a20)
	v158 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	*(*int32)(unsafe.Add(mBase, uint32(v158+v147))) = v156
	v163 = *(*int32)(unsafe.Add(mBase, _consts[477]))
	v164 = v158
	v165 = v163
	v166 = v156
	goto L30
L34:
	;
	goto L29
L35:
	;
	goto L26
L36:
	;
	F_killRDBChild(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L39
	}
L37:
	;
	F__serverLog(m, int32(2), int32(_a1632), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L21
L40:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v222 = v220 & int32(67108864)
	v224 = *(*int32)(unsafe.Add(mBase, _consts[60]))
	if v224 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v206 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[222])) = int32(0)
	v210 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v210 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_killSlotMigrationChild(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	F__serverLog(m, int32(2), int32(_a1633), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L40
L46:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+52))
	m.T0[v246].(func(*base.Module, int32))(m, v244)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L54
	}
L47:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v237 = int32(0)
	v241 = F___syscall_shutdown(m, v235, int32(2), v237, v237, v237, v237)
	mBase = m.M
	v242 = F___syscall_ret(m, v241)
	mBase = m.M
	goto L53
L48:
	;
	if v222 == int32(0) {
		goto L46
	} else {
		goto L52
	}
L49:
	;
	if v222 != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+48))
	m.T0[v229].(func(*base.Module, int32))(m, v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	goto L46
L52:
	;
	goto L47
L53:
	;
	goto L46
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	goto L1
L55:
	;
	F__serverAssert(m, int32(_a1634), int32(_a1630), int32(2073))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L74
	}
L56:
	;
	F__serverAssert(m, int32(_a1635), int32(_a1630), int32(2067))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L9
	} else {
		goto L73
	}
L57:
	;
	F__serverAssert(m, int32(_a1636), int32(_a1630), int32(2062))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L72
	}
L58:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v280 == int32(1) {
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _consts[478]))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	if v268 == int32(0) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	F_listUnlinkNode(m, v267, l0+int32(168))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v277 = v275 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v277
	v279 = v277
	goto L58
L62:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v283 == int32(1) {
		goto L56
	} else {
		goto L63
	}
L63:
	;
	if v279&int32(128) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v305&int32(4) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v292 = F_listSearchKey(m, v291, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	if v292 == int32(0) {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	F_listDelNode(m, v297, v292)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v300 & int32(-129)
	goto L64
L69:
	;
	m.G0 = v13 + int32(16)
	return
L70:
	;
	F_disableTracking(m, l0)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unsetenv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	goto L8
L1:
	;
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
	if v122 == v120 {
		v201 = v120
		goto L28
	} else {
		goto L29
	}
L2:
	;
	goto L27
L3:
	;
	if v97 == l0 {
		goto L2
	} else {
		goto L25
	}
L4:
	;
	goto L3
L5:
	;
	v88 = v83
	goto L21
L6:
	;
	v83 = v74
	goto L5
L8:
	;
	if l0&int32(3) == int32(0) {
		v34 = l0
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 != v43 {
		v74 = v34
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v21 = l0
	goto L11
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v26 == int32(0) {
		v97 = v21
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v34 = v31
	goto L9
L13:
	;
	if v26 == int32(61) {
		v97 = v21
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v31 = v21 + int32(1)
	if v31&int32(3) != 0 {
		v21 = v31
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v49 = v34
	v52 = v40
	goto L17
L17:
	;
	v55 = v52 ^ int32(1027423549)
	v58 = int32(-2139062144)
	if (int32(16843008)-v55|v55)&v58 != v58 {
		v74 = v49
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v64 = v49 + int32(4)
	v68 = int32(-2139062144)
	if (v62|(int32(16843008)-v62))&v68 == v68 {
		v49 = v64
		v52 = v62
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v83 = v64
	goto L5
L21:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 == int32(0) {
		v97 = v88
		goto L4
	} else {
		goto L23
	}
L22:
	;
	v97 = v88
	goto L4
L23:
	;
	if v89 != int32(61) {
		v88 = v88 + int32(1)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v109 = v97 - l0
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v109))))
	if v111 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L2
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(28)
	return int32(-1)
L28:
	;
	return v201
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v125 == int32(0) {
		v201 = v120
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v131 = v122
	v132 = v125
	v133 = v122
	goto L31
L31:
	;
	if v109 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v196 = int32(0)
	if v191 == v194 {
		v201 = v196
		goto L28
	} else {
		goto L54
	}
L33:
	;
	v194 = v133 + int32(4)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v195 != 0 {
		v131 = v191
		v132 = v195
		v133 = v194
		goto L31
	} else {
		goto L53
	}
L34:
	;
	if v131 == v133 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	if v178 != 0 {
		goto L34
	} else {
		goto L48
	}
L36:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v137 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v178 = int32(0)
	goto L35
L38:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v178 = v166 - v171
	goto L35
L39:
	;
	v139 = l0
	v140 = v132
	v141 = v109
	v142 = v137
	goto L42
L40:
	;
	v166 = int32(0)
	v167 = v132
	goto L38
L41:
	;
	v166 = v163 & int32(255)
	v167 = v161
	goto L38
L42:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v142&int32(255) != v146 {
		v161 = v140
		v163 = v142
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v161 = v155
	v163 = int32(0)
	goto L41
L44:
	;
	if v146 == int32(0) {
		v161 = v140
		v163 = v142
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v151 = v141 + int32(-1)
	if v151 == int32(0) {
		v161 = v140
		v163 = v142
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v154 = int32(1)
	v155 = v140 + v154
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	if v156 != 0 {
		v139 = v139 + v154
		v140 = v155
		v141 = v151
		v142 = v156
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v109))))
	if v181 != int32(61) {
		goto L34
	} else {
		goto L49
	}
L49:
	;
	goto L50
L50:
	;
	v191 = v131
	goto L33
L51:
	;
	v191 = v131 + int32(4)
	goto L33
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v187
	goto L51
L53:
	;
	goto L32
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = int32(0)
	v201 = v196
	goto L28
}
func F_unwatchAllKeys(m *base.Module, l0 int32) {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1626), int32(_a1627), int32(420))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L14
	} else {
		goto L35
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = v8 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L5
L5:
	;
	v25 = v8 + int32(8)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+48))
	if v87 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L7:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27+base.B2i32(v30 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	goto L8
L10:
	;
	v43 = v27
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v46 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L6
L13:
	;
	F_listUnlinkNode(m, v46, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_listDelNode(m, v57+int32(24), v43)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+24))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v55 = F_dictDelete(m, v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	F_decrRefCount(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_valkey_free(m, v45)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v68 = v8 + int32(8)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v70 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v70 != 0 {
		v43 = v70
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70+base.B2i32(v73 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v79
	goto L23
L25:
	;
	goto L12
L26:
	;
	v122 = int32(_a20)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[473]))
	*(*int32)(unsafe.Add(mBase, _consts[473])) = v124 + int32(-1)
	goto L2
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if v91 < int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = int32(0)
	v98 = v91
	goto L29
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v97<<(uint(int32(2))%32))))
	if v105 == int32(0) {
		v113 = v98
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v115 = v97 + int32(1)
	if v115 < v113 {
		v97 = v115
		v98 = v113
		goto L29
	} else {
		goto L34
	}
L32:
	;
	F_hashtableEmpty(m, v105, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v113 = v112
	goto L31
L34:
	;
	goto L30
L35:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unwatchCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	F_unwatchAllKeys(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v4 & int32(-33)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[27]))
		F_addReply(m, l0, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
func F_updateAppendFsync(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	if v3 != int32(1) {
		return int32(1)
	} else {
		F_bioDrainWorker(m, int32(1))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_updateDefragConfiguration(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[256])) = v3
	return v3
}
func F_updateFailoverStatus(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int64
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	if v11 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(96)
	return
L2:
	;
	v15 = *(*int64)(unsafe.Add(mBase, _consts[639]))
	if v15 == int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v19 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	if v19 < v15 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[641]))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_abortFailover(m, int32(_a2010))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L13
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v26 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[638])) = int32(2)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v45 = int32(0)
	F_replicationSetPrimary(m, v42, v44, v45, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L12
	}
L9:
	;
	v30 = *(*int64)(unsafe.Add(mBase, _consts[640]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v30
	F__serverLog(m, int32(2), int32(_a2011), v8+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	goto L1
L14:
	;
	if v142 == int32(0) {
		goto L1
	} else {
		goto L39
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v63 = v8 + int32(88)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64
	goto L18
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v58 = F_findReplica(m, v53, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v142 = v58
	goto L14
L18:
	;
	v69 = v8 + int32(88)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v71 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v85 = v71
	goto L24
L20:
	;
	if v71 != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71+base.B2i32(v74 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v80
	goto L21
L23:
	;
	v142 = int32(0)
	goto L14
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+104))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v89)+64))
	v92 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	if v90 != v92 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v142 = v88
	goto L14
L26:
	;
	v126 = v8 + int32(88)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v128 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+152))
	if v94 != 0 {
		v113 = v94
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = F_zstrdup(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L34
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	if v95 == int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	if v99 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v107 = m.T0[v99].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v95, v8+int32(32), int32(46), int32(0), int32(1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v107 == int32(-1) {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v113 = v8 + int32(32)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[640])) = v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+104))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+148))
	*(*int32)(unsafe.Add(mBase, _consts[642])) = v121
	v142 = v88
	goto L14
L35:
	;
	if v128 != 0 {
		v85 = v128
		goto L24
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128+base.B2i32(v131 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v137
	goto L36
L38:
	;
	goto L25
L39:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+104))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v146)+64))
	v149 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	if v147 != v149 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v151 = int32(_a20)
	v152 = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[638])) = v152
	v155 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v157 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v159 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v152 < v159 {
		v172 = v155
		v173 = v157
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = int32(0)
	F_replicationSetPrimary(m, v173, v172, v174, v174)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L10
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v157
	F__serverLog(m, int32(2), int32(_a2012), v8)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v168 = int32(_a20)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v171 = *(*int32)(unsafe.Add(mBase, _consts[640]))
	v172 = v169
	v173 = v171
	goto L41
L44:
	;
	goto L1
}
func F_updateHZ(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	v4 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	if v6 < v4 {
		v12 = v4
		*(*int32)(unsafe.Add(mBase, _consts[251])) = v12
	} else {
		if base.Ui32(v6) < base.Ui32(int32(501)) {
		} else {
			v12 = int32(500)
			*(*int32)(unsafe.Add(mBase, _consts[251])) = v12
		}
	}
	return int32(1)
}
func F_updatePausedActions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(_a20)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, _consts[496])) = v1
	v16 = *(*int64)(unsafe.Add(mBase, _consts[497]))
	v18 = *(*int64)(unsafe.Add(mBase, _consts[54]))
	if v16 <= v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = *(*int64)(unsafe.Add(mBase, _consts[498]))
	if v18 < v32 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[497])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[499])) = int32(0)
	v30 = v1
	goto L1
L3:
	;
	v20 = int32(_a20)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	*(*int32)(unsafe.Add(mBase, _consts[496])) = v22
	v30 = v22
	goto L1
L4:
	;
	v47 = *(*int64)(unsafe.Add(mBase, _consts[500]))
	if v18 < v47 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v40 = int32(_a20)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[501]))
	v43 = v30 | v42
	*(*int32)(unsafe.Add(mBase, _consts[496])) = v43
	v45 = v43
	goto L4
L6:
	;
	v34 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[498])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[501])) = int32(0)
	v45 = v30
	goto L4
L7:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _consts[502]))
	if v18 < v62 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v55 = int32(_a20)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[503]))
	v58 = v45 | v57
	*(*int32)(unsafe.Add(mBase, _consts[496])) = v58
	v60 = v58
	goto L7
L9:
	;
	v49 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[500])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[503])) = int32(0)
	v60 = v45
	goto L7
L10:
	;
	v76 = int32(3)
	if base.Ui32(v10&v76) <= base.Ui32(v75&v76) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v70 = int32(_a20)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	v73 = v60 | v72
	*(*int32)(unsafe.Add(mBase, _consts[496])) = v73
	v75 = v73
	goto L10
L12:
	;
	v64 = int32(_a20)
	*(*int64)(unsafe.Add(mBase, _consts[502])) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[228])) = int32(0)
	v75 = v60
	goto L10
L13:
	;
	m.G0 = v7 + int32(16)
	return
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[76]))
	v84 = v7 + int32(8)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v85
	goto L15
L15:
	;
	v90 = v7 + int32(8)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v92 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v92 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v92+base.B2i32(v95 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v101
	goto L17
L19:
	;
	v107 = v92
	goto L20
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	F_unblockClient(m, v109, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L13
L22:
	;
	return
L23:
	;
	v114 = v7 + int32(8)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v116 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v116 != 0 {
		v107 = v116
		goto L20
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116+base.B2i32(v119 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v125
	goto L25
L27:
	;
	goto L21
}
func F_updatePort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v5 = F_listenerByType(m, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			F__serverAssert(m, int32(_a524), int32(_a473), int32(2564))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = int32(_a20)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = int32(_a261)
			v16 = *(*int32)(unsafe.Add(mBase, _consts[135]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v16
			v19 = *(*int32)(unsafe.Add(mBase, _consts[137]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = v19
			v22 = F_connectionByType(m, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+80)) = v22
				F_clusterUpdateMyselfAnnouncedPorts(m)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = F_changeListener(m, v5)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						if v28 != int32(-1) {
							v35 = int32(1)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a525)
							v35 = int32(0)
						}
						return v35
					}
				}
			}
		}
	}
}
func F_updateSharedObjectsWithCompat(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v1 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[296]))
	v6 = v4 << (uint(int32(2)) % 32)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[736])))
	*(*int32)(unsafe.Add(mBase, _consts[741])) = v9
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[737])))
	*(*int32)(unsafe.Add(mBase, _consts[334])) = v14
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[738])))
	*(*int32)(unsafe.Add(mBase, _consts[742])) = v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[739])))
	*(*int32)(unsafe.Add(mBase, _consts[743])) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[740])))
	*(*int32)(unsafe.Add(mBase, _consts[744])) = v29
	return
}
func F_usUntilEarliestTimer(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v5 != 0 {
		v9 = v5
		v10 = int32(0)
		for {
			if v10 == int32(0) {
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
				if v18 == int64(-1) {
					v21 = v10
				} else {
					v21 = v9
				}
				v22 = v21
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
				v16 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(v16) <= base.Ui64(v15) {
					v22 = v10
				} else {
					v18 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
					if v18 == int64(-1) {
						v21 = v10
					} else {
						v21 = v9
					}
					v22 = v21
				}
			}
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
			if v23 != 0 {
				v9 = v23
				v10 = v22
				continue
			} else {
				break
			}
			break
		}
		v25 = *(*int32)(unsafe.Add(mBase, _consts[34]))
		v26 = m.T0[v25].(func(*base.Module) int64)(m)
		mBase = m.M
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
		v29 = v28 - v26
		if base.Ui64(v28) < base.Ui64(v29) {
			v31 = int64(0)
		} else {
			v31 = v29
		}
		return v31
	} else {
		return int64(-1)
	}
}
func F_ustime(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[34]))
	if v9 == v1 {
		v30 = int32(0)
		v31 = F___gettimeofday(m, v6, v30)
		mBase = m.M
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
		v37 = v33*int64(1000000) + v36
		*(*int64)(unsafe.Add(mBase, _consts[901])) = v37
		v42 = v37
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[34]))
		if base.B2i32(v13 != int32(948)) != int32(1) {
			v30 = int32(0)
			v31 = F___gettimeofday(m, v6, v30)
			mBase = m.M
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
			v37 = v33*int64(1000000) + v36
			*(*int64)(unsafe.Add(mBase, _consts[901])) = v37
			v42 = v37
		} else {
			v18 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _consts[34]))
			v20 = m.T0[v19].(func(*base.Module) int64)(m)
			mBase = m.M
			v22 = *(*int64)(unsafe.Add(mBase, _consts[902]))
			v23 = v20 - v22
			if base.Ui64(v23) < base.Ui64(int64(1000)) {
				v40 = *(*int64)(unsafe.Add(mBase, _consts[901]))
				v42 = v40 + v23
			} else {
				*(*int64)(unsafe.Add(mBase, _consts[902])) = v20
				v30 = int32(0)
				v31 = F___gettimeofday(m, v6, v30)
				mBase = m.M
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+8)))
				v37 = v33*int64(1000000) + v36
				*(*int64)(unsafe.Add(mBase, _consts[901])) = v37
				v42 = v37
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v42
}
