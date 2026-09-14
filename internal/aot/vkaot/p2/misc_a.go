package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F___acquire_ptc(m *base.Module) {
	return
}
func F___asctime_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v148 int64
	_ = v148
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v12 + int32(131072)
	if v14 != int32(14) {
		v26 = v14 >> (uint(int32(16)) % 32)
		v27 = int32(65535)
		v28 = v14 & v27
		if v28 != v27 {
			v41 = int32(_a139)
			switch v26 + int32(-1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v28) {
					v68 = v41
					v74 = v68
				} else {
					v53 = int32(_a2163)
					if v28 != 0 {
						v54 = v53
						v57 = v28
						for {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
							v61 = v54 + int32(1)
							if v59 != 0 {
								v54 = v61
								continue
							} else {
							}
							v63 = v57 + int32(-1)
							if v63 != 0 {
								v54 = v61
								v57 = v63
								continue
							} else {
								break
							}
							break
						}
						v68 = v61
						v74 = v68
					} else {
						v74 = v53
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v28) {
					v68 = v41
					v74 = v68
				} else {
					v53 = int32(_a2164)
					if v28 != 0 {
						v54 = v53
						v57 = v28
						for {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
							v61 = v54 + int32(1)
							if v59 != 0 {
								v54 = v61
								continue
							} else {
							}
							v63 = v57 + int32(-1)
							if v63 != 0 {
								v54 = v61
								v57 = v63
								continue
							} else {
								break
							}
							break
						}
						v68 = v61
						v74 = v68
					} else {
						v74 = v53
					}
				}
			default:
				v68 = v41
				v74 = v68
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v28) {
					v68 = v41
					v74 = v68
				} else {
					v53 = int32(_a2165)
					if v28 != 0 {
						v54 = v53
						v57 = v28
						for {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
							v61 = v54 + int32(1)
							if v59 != 0 {
								v54 = v61
								continue
							} else {
							}
							v63 = v57 + int32(-1)
							if v63 != 0 {
								v54 = v61
								v57 = v63
								continue
							} else {
								break
							}
							break
						}
						v68 = v61
						v74 = v68
					} else {
						v74 = v53
					}
				}
			}
		} else {
			if int32(5) < v26 {
				v41 = int32(_a139)
				switch v26 + int32(-1) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(v28) {
						v68 = v41
						v74 = v68
					} else {
						v53 = int32(_a2163)
						if v28 != 0 {
							v54 = v53
							v57 = v28
							for {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
								v61 = v54 + int32(1)
								if v59 != 0 {
									v54 = v61
									continue
								} else {
								}
								v63 = v57 + int32(-1)
								if v63 != 0 {
									v54 = v61
									v57 = v63
									continue
								} else {
									break
								}
								break
							}
							v68 = v61
							v74 = v68
						} else {
							v74 = v53
						}
					}
				case 1:
					if base.Ui32(int32(49)) < base.Ui32(v28) {
						v68 = v41
						v74 = v68
					} else {
						v53 = int32(_a2164)
						if v28 != 0 {
							v54 = v53
							v57 = v28
							for {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
								v61 = v54 + int32(1)
								if v59 != 0 {
									v54 = v61
									continue
								} else {
								}
								v63 = v57 + int32(-1)
								if v63 != 0 {
									v54 = v61
									v57 = v63
									continue
								} else {
									break
								}
								break
							}
							v68 = v61
							v74 = v68
						} else {
							v74 = v53
						}
					}
				default:
					v68 = v41
					v74 = v68
				case 4:
					if base.Ui32(int32(3)) < base.Ui32(v28) {
						v68 = v41
						v74 = v68
					} else {
						v53 = int32(_a2165)
						if v28 != 0 {
							v54 = v53
							v57 = v28
							for {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
								v61 = v54 + int32(1)
								if v59 != 0 {
									v54 = v61
									continue
								} else {
								}
								v63 = v57 + int32(-1)
								if v63 != 0 {
									v54 = v61
									v57 = v63
									continue
								} else {
									break
								}
								break
							}
							v68 = v61
							v74 = v68
						} else {
							v74 = v53
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v26<<(uint(int32(2))%32))+uint32(_consts[1035])))
				if v36 != 0 {
					v40 = v36 + int32(8)
				} else {
					v40 = int32(_a2166)
				}
				v74 = v40
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[1035]))
		if v23 != 0 {
			v24 = int32(_a2167)
		} else {
			v24 = int32(_a2168)
		}
		v74 = v24
	}
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v77 = v75 + int32(131086)
	if v77 != int32(14) {
		v89 = v77 >> (uint(int32(16)) % 32)
		v90 = int32(65535)
		v91 = v77 & v90
		if v91 != v90 {
			v104 = int32(_a139)
			switch v89 + int32(-1) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(v91) {
					v131 = v104
					v137 = v131
				} else {
					v116 = int32(_a2163)
					if v91 != 0 {
						v117 = v116
						v120 = v91
						for {
							v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
							v124 = v117 + int32(1)
							if v122 != 0 {
								v117 = v124
								continue
							} else {
							}
							v126 = v120 + int32(-1)
							if v126 != 0 {
								v117 = v124
								v120 = v126
								continue
							} else {
								break
							}
							break
						}
						v131 = v124
						v137 = v131
					} else {
						v137 = v116
					}
				}
			case 1:
				if base.Ui32(int32(49)) < base.Ui32(v91) {
					v131 = v104
					v137 = v131
				} else {
					v116 = int32(_a2164)
					if v91 != 0 {
						v117 = v116
						v120 = v91
						for {
							v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
							v124 = v117 + int32(1)
							if v122 != 0 {
								v117 = v124
								continue
							} else {
							}
							v126 = v120 + int32(-1)
							if v126 != 0 {
								v117 = v124
								v120 = v126
								continue
							} else {
								break
							}
							break
						}
						v131 = v124
						v137 = v131
					} else {
						v137 = v116
					}
				}
			default:
				v131 = v104
				v137 = v131
			case 4:
				if base.Ui32(int32(3)) < base.Ui32(v91) {
					v131 = v104
					v137 = v131
				} else {
					v116 = int32(_a2165)
					if v91 != 0 {
						v117 = v116
						v120 = v91
						for {
							v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
							v124 = v117 + int32(1)
							if v122 != 0 {
								v117 = v124
								continue
							} else {
							}
							v126 = v120 + int32(-1)
							if v126 != 0 {
								v117 = v124
								v120 = v126
								continue
							} else {
								break
							}
							break
						}
						v131 = v124
						v137 = v131
					} else {
						v137 = v116
					}
				}
			}
		} else {
			if int32(5) < v89 {
				v104 = int32(_a139)
				switch v89 + int32(-1) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(v91) {
						v131 = v104
						v137 = v131
					} else {
						v116 = int32(_a2163)
						if v91 != 0 {
							v117 = v116
							v120 = v91
							for {
								v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
								v124 = v117 + int32(1)
								if v122 != 0 {
									v117 = v124
									continue
								} else {
								}
								v126 = v120 + int32(-1)
								if v126 != 0 {
									v117 = v124
									v120 = v126
									continue
								} else {
									break
								}
								break
							}
							v131 = v124
							v137 = v131
						} else {
							v137 = v116
						}
					}
				case 1:
					if base.Ui32(int32(49)) < base.Ui32(v91) {
						v131 = v104
						v137 = v131
					} else {
						v116 = int32(_a2164)
						if v91 != 0 {
							v117 = v116
							v120 = v91
							for {
								v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
								v124 = v117 + int32(1)
								if v122 != 0 {
									v117 = v124
									continue
								} else {
								}
								v126 = v120 + int32(-1)
								if v126 != 0 {
									v117 = v124
									v120 = v126
									continue
								} else {
									break
								}
								break
							}
							v131 = v124
							v137 = v131
						} else {
							v137 = v116
						}
					}
				default:
					v131 = v104
					v137 = v131
				case 4:
					if base.Ui32(int32(3)) < base.Ui32(v91) {
						v131 = v104
						v137 = v131
					} else {
						v116 = int32(_a2165)
						if v91 != 0 {
							v117 = v116
							v120 = v91
							for {
								v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
								v124 = v117 + int32(1)
								if v122 != 0 {
									v117 = v124
									continue
								} else {
								}
								v126 = v120 + int32(-1)
								if v126 != 0 {
									v117 = v124
									v120 = v126
									continue
								} else {
									break
								}
								break
							}
							v131 = v124
							v137 = v131
						} else {
							v137 = v116
						}
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[1035])))
				if v99 != 0 {
					v103 = v99 + int32(8)
				} else {
					v103 = int32(_a2166)
				}
				v137 = v103
			}
		}
	} else {
		v86 = *(*int32)(unsafe.Add(mBase, _consts[1035]))
		if v86 != 0 {
			v87 = int32(_a2167)
		} else {
			v87 = int32(_a2168)
		}
		v137 = v87
	}
	v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v142 + int32(1900)
	v148 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(16)))) = base.I64_rotl(v139, v148)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v74
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = base.I64_rotl(v138, v148)
	v158 = F_snprintf(m, l1, int32(26), int32(_a2169), v10)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		return int32(0)
	} else {
		if v158 < int32(26) {
			m.G0 = v10 + int32(32)
			return l1
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F___ashlti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	if l3&int32(64) == int32(0) {
		if l3 == int32(0) {
			v25 = l1
			v26 = l2
		} else {
			v21 = base.I64_extend_i32_u(l3)
			v25 = l1 << (uint(v21) % 64)
			v26 = int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(int32(64)-l3))%64)) | l2<<(uint(v21)%64)
		}
	} else {
		v25 = int64(0)
		v26 = l1 << (uint(base.I64_extend_i32_u(l3+int32(-64))) % 64)
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v25
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v26
	return
}
func F_a_cas_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
	return v3
}
func F_a_ctz_32(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	if l0 != 0 {
		v4 = base.I32_ctz(l0)
	} else {
		v4 = int32(0)
	}
	return v4
}
func F_a_swap_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return v3
}
func F_a_swap_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return v3
}
func F_abort(m *base.Module) {
	m.Env.X_abort_js(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_abortFailover(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v10 == int32(0) {
		m.G0 = v7 + int32(32)
		return
	} else {
		v13 = int32(_a44)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v16 = *(*int32)(unsafe.Add(mBase, _consts[713]))
		if v16 == int32(0) {
			if int32(2) < v14 {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[658]))
				if v40 != int32(2) {
					v45 = int32(_a44)
					*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
					*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
					v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
					F_valkey_free(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
						F_unpauseActions(m, int32(2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				} else {
					F_replicationUnsetPrimary(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
						v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
						F_valkey_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
							F_unpauseActions(m, int32(2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F__serverLog(m, int32(2), int32(_a1287), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[658]))
					if v40 != int32(2) {
						v45 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
						v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
						F_valkey_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
							F_unpauseActions(m, int32(2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					} else {
						F_replicationUnsetPrimary(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v45 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
							v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
							F_valkey_free(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
								F_unpauseActions(m, int32(2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		} else {
			if int32(2) < v14 {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[658]))
				if v40 != int32(2) {
					v45 = int32(_a44)
					*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
					*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
					v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
					F_valkey_free(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
						F_unpauseActions(m, int32(2))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				} else {
					F_replicationUnsetPrimary(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
						v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
						F_valkey_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
							F_unpauseActions(m, int32(2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l0
				v24 = *(*int32)(unsafe.Add(mBase, _consts[716]))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v24
				F__serverLog(m, int32(2), int32(_a1288), v7+int32(16))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[658]))
					if v40 != int32(2) {
						v45 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
						v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
						F_valkey_free(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
							F_unpauseActions(m, int32(2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					} else {
						F_replicationUnsetPrimary(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v45 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[714])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[715])) = int64(0)
							v52 = *(*int32)(unsafe.Add(mBase, _consts[713]))
							F_valkey_free(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[658])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[713])) = int64(0)
								F_unpauseActions(m, int32(2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
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
func F_abortShutdown(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v3 = *(*int64)(unsafe.Add(mBase, _consts[840]))
	if v3 == int64(0) {
		v23 = *(*int32)(unsafe.Add(mBase, _consts[841]))
		if v23 != 0 {
			v26 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[841])) = v26
			v29 = int32(0)
			v31 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if int32(2) < v31 {
				v40 = v29
				return v40
			} else {
				v34 = int32(0)
				F__serverLog(m, int32(2), int32(_a1613), v34)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = v34
					return v40
				}
			}
		} else {
			return int32(-1)
		}
	} else {
		v6 = int32(0)
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, _consts[840])) = v7
		*(*int64)(unsafe.Add(mBase, _consts[844])) = v7
		*(*int32)(unsafe.Add(mBase, _consts[841])) = v6
		F_replyToClientsBlockedOnShutdown(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_unpauseActions(m, int32(1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v29 = int32(0)
				v31 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				if int32(2) < v31 {
					v40 = v29
					return v40
				} else {
					v34 = int32(0)
					F__serverLog(m, int32(2), int32(_a1613), v34)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = v34
						return v40
					}
				}
			}
		}
	}
}
func F_acceptCommonHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	v7 = m.G0
	v9 = v7 - int32(480)
	m.G0 = v9
	v16 = F__emscripten_memset_bulkmem(m, v9+int32(208), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L1
L1:
	;
	v22 = F__emscripten_memset_bulkmem(m, v9+int32(80), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L2
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v105 == int32(2) {
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	if v66 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L6:
	;
	v35 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, v9+int32(352), int32(128), v9+int32(348), int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v35 < int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v41 = int32(58)
	v42 = F___strchrnul(m, v9+int32(352), v41)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v44 == v41 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v9 + int32(352)
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v48 = v42
	goto L13
L12:
	;
	v48 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v59 = int32(_a1004)
	goto L16
L15:
	;
	v59 = int32(_a1005)
	goto L16
L16:
	;
	v62 = F_snprintf(m, v9+int32(208), int32(128), v59, v9+int32(64))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L5
L18:
	;
	v75 = m.T0[v66].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, v9+int32(352), int32(128), v9+int32(348), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	if v75 < int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v81 = int32(58)
	v82 = F___strchrnul(m, v9+int32(352), v81)
	mBase = m.M
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v84 == v81 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v9 + int32(352)
	if v88 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v88 = v82
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	goto L21
L25:
	;
	v99 = int32(_a1004)
	goto L27
L26:
	;
	v99 = int32(_a1005)
	goto L27
L27:
	;
	v102 = F_snprintf(m, v9+int32(80), int32(128), v99, v9+int32(48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	m.G0 = v9 + int32(480)
	return
L30:
	;
	v134 = int32(_a44)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	v139 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v139 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(1) < v110 {
		v130 = v108
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+52))
	m.T0[v131].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L36
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+88))
	v114 = m.T0[v113].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v9 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v9 + int32(208)
	F__serverLog(m, int32(1), int32(_a1003), v9+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = v129
	goto L32
L36:
	;
	goto L29
L37:
	;
	v179 = F_createClient(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L51
	}
L38:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	if base.Ui32(v136+v152) < base.Ui32(v155) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v152 = (v144+v145)<<(uint(int32(1))%32) + int32(-2)
	goto L38
L40:
	;
	v152 = int32(0)
	goto L38
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v160 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v161 = int32(_a1001)
	goto L44
L43:
	;
	v161 = int32(_a1002)
	goto L44
L44:
	;
	if v160 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v164 = int32(58)
	goto L47
L46:
	;
	v164 = int32(36)
	goto L47
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+68))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32) int32)(m, l0, v161, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v169 = int32(_a44)
	v171 = *(*int64)(unsafe.Add(mBase, _consts[594]))
	*(*int64)(unsafe.Add(mBase, _consts[594])) = v171 + int64(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+52))
	m.T0[v176].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	goto L29
L50:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v204&int32(8) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	if v179 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v182 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+52))
	m.T0[v201].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L57
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+88))
	v187 = m.T0[v186].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v9 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(208)
	F__serverLog(m, int32(3), int32(_a1000), v9)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	goto L29
L58:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+64))
	v216 = m.T0[v215].(func(*base.Module, int32, int32) int32)(m, l0, int32(955))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v179)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+200)) = v209 | int32(2048)
	goto L58
L60:
	;
	if v216 != int32(-1) {
		goto L29
	} else {
		goto L61
	}
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v220 != int32(5) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v247 = F_freeClient(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L69
	}
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v224 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+88))
	v229 = m.T0[v228].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v231 = F_getClientPeerId(m, v179)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v233 = F_getClientSockname(m, v179)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v229
	F__serverLog(m, int32(3), int32(_a999), v9+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	goto L62
L69:
	;
	goto L29
}
func F_action_abort(m *base.Module, l0 int32) {
	F_abort(m)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_activeExpireCycleJob(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v212 int32
	_ = v212
	var v225 int32
	_ = v225
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v271 int64
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v321 int64
	_ = v321
	var v322 int32
	_ = v322
	var v325 int64
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v411 float64
	_ = v411
	var v414 float64
	_ = v414
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int64
	_ = v447
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v540 int32
	_ = v540
	var v541 float64
	_ = v541
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v572 float64
	_ = v572
	var v578 int32
	_ = v578
	var v581 float64
	_ = v581
	var v590 int64
	_ = v590
	v5 = int64(0)
	v38 = m.G0
	v40 = v38 - int32(48)
	m.G0 = v40
	if l2 < int64(1) {
		v590 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v40 + int32(48)
	return v590
L2:
	;
	v46 = l0 << (uint(int32(3)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[417]))
	v52 = int32(11) - v51
	v54 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v55 = m.T0[v54].(func(*base.Module) int64)(m)
	mBase = m.M
	if l1 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v70 < int32(16) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))))
	if v58 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[419])))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v63)))
	if base.F64_lt(v64, base.F64_convert_i32_u(v52)) != 0 {
		v590 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v568 = m.T0[v567].(func(*base.Module) int64)(m)
	mBase = m.M
	if v540 != 0 {
		goto L81
	} else {
		goto L82
	}
L8:
	;
	v90 = l0 & int32(1)
	if v90 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))) = uint8(v78)
	if int32(1) <= v70 {
		v85 = v70
		goto L8
	} else {
		goto L14
	}
L10:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))))
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))) = uint8(v74)
	if v73 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v77 = v70
	goto L13
L12:
	;
	v77 = int32(16)
	goto L13
L13:
	;
	v85 = v77
	goto L8
L14:
	;
	v540 = v78
	v541 = float64(0)
	goto L7
L15:
	;
	v91 = int32(8)
	goto L17
L16:
	;
	v91 = int32(4)
	goto L17
L17:
	;
	if v90 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v94 = int32(525)
	goto L20
L19:
	;
	v94 = int32(526)
	goto L20
L20:
	;
	if v90 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = int32(0)
	goto L23
L22:
	;
	v97 = int32(15)
	goto L23
L23:
	;
	v99 = l0 << (uint(int32(4)) % 32)
	v103 = v51*int32(5) + int32(15)
	v114 = int32(0)
	v130 = v114
	v142 = v114
	v143 = v114
	v144 = v114
	v145 = v114
	goto L25
L24:
	;
	v540 = v502
	v541 = base.F64_convert_i32_s(v517)
	goto L7
L25:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))))
	if v156 != 0 {
		v502 = v130
		v517 = v145
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v502 = v462
	v517 = v477
	goto L24
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if v158 <= v142 {
		v502 = v130
		v517 = v145
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v160 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40+int32(40)))) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v40+int32(32)))) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v40+int32(24)))) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v40+int32(16)))) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v40+int32(8)))) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = v160
	v175 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[420])))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[420]))) = v176 + int32(1)
	v180 = base.I32_rem_u_s(v176, v158)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v175+v180<<(uint(int32(2))%32))))
	if v184 == int32(0) {
		v462 = v130
		v475 = v143
		v476 = v144
		v477 = v145
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v476 < v85 {
		v130 = v462
		v142 = v142 + int32(1)
		v143 = v475
		v144 = v476
		v145 = v477
		goto L25
	} else {
		goto L79
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v103 << (uint(int32(2)) % 32)
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v184
	v193 = v184 + int32(32)
	v196 = v184 + v99 + int32(40)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v184+v91)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if v199 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v212 = v144 + base.B2i32(v209 != int64(0))
	v225 = v130
	v238 = v143
	v240 = v145
	v242 = v188
	v245 = int32(0)
	goto L36
L32:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v204 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v198)+40))
	v209 = v202
	goto L31
L34:
	;
	v206 = F_hashtableSize(m, v204)
	mBase = m.M
	v209 = base.I64_extend_i32_u(v206)
	goto L31
L35:
	;
	v209 = int64(0)
	goto L31
L36:
	;
	v251 = int32(1)
	v252 = v238 + v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if v253 == v251 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v462 = v435
	v475 = v252
	v476 = v212
	v477 = v436
	goto L29
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+16)) = int64(0)
	v271 = *(*int64)(unsafe.Add(mBase, _consts[35]))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v103 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v264 = base.I32_wrap_i64(v263)
	if v264 != 0 {
		goto L38
	} else {
		goto L44
	}
L40:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v258 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v198)+40))
	v263 = v256
	goto L39
L42:
	;
	v260 = F_hashtableSize(m, v258)
	mBase = m.M
	v263 = base.I64_extend_i32_u(v260)
	goto L39
L43:
	;
	v263 = int64(0)
	goto L39
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v193+v99))) = int64(0)
	v462 = v225
	v475 = v252
	v476 = v212
	v477 = v240
	goto L29
L45:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v380 = v245 + base.B2i32(v274 < v378)
	if v368 != 0 {
		v396 = int32(0)
		goto L63
	} else {
		goto L64
	}
L46:
	;
	if base.Ui32(v103) < base.Ui32(v264) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v367 = int32(0)
	v368 = v242
	goto L45
L48:
	;
	v285 = int32(0)
	goto L53
L49:
	;
	v278 = v103
	goto L51
L50:
	;
	v278 = v264
	goto L51
L51:
	;
	v280 = v278 * int32(10)
	if int32(1) <= v280 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v367 = int32(0)
	v368 = v242
	goto L45
L53:
	;
	v321 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v196))))
	v322 = int32(-1)
	v325 = F_kvstoreScan(m, v198, v321, v322, v322, v94, int32(527), v40)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v367 = v335
	v368 = v242
	goto L45
L55:
	;
	return int64(0)
L56:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+40)))
	if v329 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if base.Ui32(v278) <= base.Ui32(v335) {
		v367 = v335
		v368 = v242
		goto L45
	} else {
		goto L60
	}
L58:
	;
	v330 = base.I32_wrap_i64(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v330
	if v330 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v367 = v333
	v368 = int32(1)
	goto L45
L60:
	;
	v338 = v285 + int32(1)
	if v338 < v280 {
		v285 = v338
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	v435 = v367 + v225
	v436 = v377 + v240
	if v252&v97 != 0 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	if l0 != 0 {
		v432 = v380
		v433 = v396
		goto L62
	} else {
		goto L69
	}
L64:
	;
	if v367 == int32(0) {
		v388 = int32(1)
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v252&int32(15) == int32(0) {
		v396 = v388
		goto L63
	} else {
		goto L67
	}
L66:
	;
	v386 = base.I32_div_u_s(v377*int32(100), v367)
	v388 = base.B2i32(base.Ui32(v52) < base.Ui32(v386))
	goto L65
L67:
	;
	if v388 == int32(0) {
		v396 = v388
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v432 = v380
	v433 = int32(1)
	goto L62
L69:
	;
	if v378 == int32(0) {
		v432 = v380
		v433 = v396
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v40)+24))
	v401 = base.I64_div_s(v399, base.I64_extend_i32_s(v378))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	if v402 == int64(0) {
		v422 = v401
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = v422
	v425 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v425
	*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = int64(0)
	v432 = v425
	v433 = v396
	goto L62
L72:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v380<<(uint(int32(3))%32))+uint32(_consts[419])))
	v414 = base.F64_add(base.F64_mul(base.F64_convert_i64_s(v402-v401), v411), base.F64_convert_i64_s(v401))
	if base.F64_lt(base.F64_abs(v414), float64(9.223372036854776e+18)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v422 = int64(-9223372036854775807 - 1)
	goto L71
L74:
	;
	v420 = base.I64_trunc_f64_s(v414)
	v422 = v420
	goto L71
L75:
	;
	if v433 != 0 {
		v225 = v435
		v238 = v252
		v240 = v436
		v242 = v368
		v245 = v432
		goto L36
	} else {
		goto L78
	}
L76:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v440 = m.T0[v439].(func(*base.Module) int64)(m)
	mBase = m.M
	if base.Ui64(v440-v55) <= base.Ui64(l2) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[418]))) = uint8(v443)
	v445 = int32(_a44)
	v447 = *(*int64)(unsafe.Add(mBase, _consts[421]))
	*(*int64)(unsafe.Add(mBase, _consts[421])) = v447 + int64(1)
	v462 = v435
	v475 = v252
	v476 = v212
	v477 = v436
	goto L29
L78:
	;
	goto L37
L79:
	;
	goto L26
L80:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[419])))
	v581 = *(*float64)(unsafe.Add(mBase, uint32(v578)))
	*(*float64)(unsafe.Add(mBase, uint32(v578))) = base.F64_add(base.F64_mul(v572, float64(0.05)), base.F64_mul(v581, float64(0.95)))
	v590 = v568 - v55
	goto L1
L81:
	;
	v572 = base.F64_div(v541, base.F64_convert_i32_s(v540))
	goto L80
L82:
	;
	v572 = float64(0)
	goto L80
}
func F_addBulkStringToReplyIOV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v16 = l0
	v17 = l1
	v20 = v15
	goto L3
L3:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		goto L11
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		v49 = int32(0)
		goto L6
	}
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v54 = v50 + v51*int32(24)
	v55 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
	v58 = v54 + int32(1)
	v60 = base.I64_extend_i32_u(v49)
	if v60 <= int64(-1) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
	v49 = v48
	goto L6
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v49 = v45
	goto L6
L9:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
	v49 = v42
	goto L6
L10:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v49 = v39
	goto L6
L11:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v105 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v54+v101+int32(1)))) = uint16(v105)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v107 != 0 {
		v222 = v107
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v101 = int32(0)
	goto L12
L15:
	;
	v82 = F_ull2string(m, v78, v79, v80)
	mBase = m.M
	if v82 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L16:
	;
	goto L18
L17:
	;
	v78 = v58
	v79 = int32(21)
	v80 = v60
	v81 = int32(0)
	goto L15
L18:
	;
	v69 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v69)
	v78 = v54 + int32(2)
	v79 = int32(20)
	v80 = int64(0) - v60
	v81 = int32(1)
	goto L15
L19:
	;
	v101 = v82 + v81
	goto L12
L21:
	;
	v233 = v17 + int32(-8)
	if v233 != 0 {
		v16 = v16 + int32(8)
		v17 = v233
		v20 = v222
		goto L3
	} else {
		goto L34
	}
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v108 == v109 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v219 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v219
	v222 = v219
	goto L21
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	v115 = v101 + int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v113 + base.I64_extend_i32_u(v115)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v115) <= base.Ui32(v119) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v145 == v148 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v119 - v115
	v145 = v108
	goto L25
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v124 = v121 + v108<<(uint(int32(3))%32)
	v125 = v115 - v119
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v112 + v111*int32(24) + v119
	v132 = int32(1)
	v133 = v108 + v132
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v111 + v132
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v140 + v125
	v145 = v133
	goto L25
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	v152 = v151 + v60
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v49) <= base.Ui32(v154) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v196 = v193 + v189<<(uint(int32(3))%32)
	v198 = int32(2) - v190
	*(*int32)(unsafe.Add(mBase, uint32(v196)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v191 + v190
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v189 + int32(1)
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v208 + v198
	v222 = v205
	goto L21
L30:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v152 + int64(2)
	v182 = v154 - v49
	if base.Ui32(v182) < base.Ui32(int32(2)) {
		v189 = v145
		v190 = v182
		v191 = v178
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v159 = v156 + v145<<(uint(int32(3))%32)
	v160 = v49 - v154
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v150 + v154
	v165 = v145 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v165
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v169 + v160
	if v165 == v148 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v152 + int64(2)
	v189 = v165
	v190 = int32(0)
	v191 = v173
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v182 + int32(-2)
	v222 = int32(0)
	goto L21
L34:
	;
	goto L4
}
func F_addEncodedBufferToReplyIOV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = l0
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v24 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v26 = v14 + int32(12)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)))
	if v27&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v73 = v26 + v67
	if base.Ui32(v73) < base.Ui32(l0+l1) {
		v14 = v73
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	F_addBulkStringToReplyIOV(m, v26, v61, l2, l3)
	mBase = m.M
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	v64 = v63 - v60
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)) = uint32(v64)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v67 = v66
	goto L6
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v31 != v32 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v36 + base.I64_extend_i32_u(v30)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v40) < base.Ui32(v30) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(1)
	v67 = v30
	goto L6
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v47 = v44 + v31<<(uint(int32(3))%32)
	v48 = v30 - v40
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v26 + v40
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v57 + v48
	v67 = v30
	goto L6
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v40 - v30
	v67 = v30
	goto L6
L13:
	;
	goto L4
}
func F_addInfoSectionsToDict(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = l1
	v9 = v4
	goto L3
L3:
	;
	v10 = F_sdsnew(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L1
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v19 != 0 {
		v8 = v8 + int32(4)
		v9 = v19
		goto L3
	} else {
		goto L11
	}
L6:
	;
	return
L7:
	;
	v13 = F_dictAdd(m, l0, v10, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v13 != int32(1) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_sdsfree(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	goto L4
}
func F_adjustOpenFilesLimit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v115 int64
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v14 = int32(7)
	v16 = v10 + int32(96)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v28 = F___syscall_prlimit64(m, v1, v14, v1, v16)
	mBase = m.M
	v29 = F___syscall_ret(m, v28)
	mBase = m.M
	if v29 == v1 {
		v66 = v1
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v10 + int32(112)
	return
L2:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v10)+96))
	v91 = base.I64_extend_i32_u(v13 + int32(32))
	if base.Ui64(v91) <= base.Ui64(v88) {
		goto L1
	} else {
		goto L25
	}
L3:
	;
	if v66 != int32(-1) {
		goto L2
	} else {
		goto L18
	}
L4:
	;
	m.G0 = v23 + int32(16)
	goto L3
L5:
	;
	v32 = F___errno_location(m)
	mBase = m.M
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 != int32(52) {
		v66 = v29
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = F___syscall_ugetrlimit(m, v14, v23+int32(8))
	mBase = m.M
	v39 = F___syscall_ret(m, v38)
	mBase = m.M
	if int32(0) <= v39 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+8)))
	if v44 == int64(4294967295) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v66 = int32(-1)
	goto L4
L9:
	;
	v47 = int64(-1)
	goto L11
L10:
	;
	v47 = v44
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v50 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = int64(-1)
	goto L14
L13:
	;
	v54 = base.I64_extend_i32_u(v50)
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v54
	if v44 != int64(4294967295) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = int32(0)
	if v50 != int32(-1) {
		v66 = v60
		goto L4
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(-1)
	goto L15
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(-1)
	v66 = v60
	goto L4
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v74 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[278])) = int32(992)
	goto L1
L20:
	;
	goto L21
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v79 = F___strerror_l(m, v78, v78)
	mBase = m.M
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v79
	F__serverLog(m, int32(3), int32(_a1581), v10)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return
L24:
	;
	goto L19
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v91
	v99 = F_setrlimit(m, int32(7), v10+int32(96))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L23
	} else {
		goto L27
	}
L26:
	;
	if base.Ui64(v88) < base.Ui64(v130) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	if v99 != int32(-1) {
		v127 = int32(0)
		v130 = v91
		goto L26
	} else {
		goto L28
	}
L28:
	;
	goto L29
L29:
	;
	v108 = v91
	goto L31
L30:
	;
	v127 = v111
	v130 = v88
	goto L26
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if base.Ui64(v108) <= base.Ui64(int64(15)) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v115 = v108 + int64(-16)
	if base.Ui64(v115) <= base.Ui64(v88) {
		v127 = v111
		v130 = v115
		goto L26
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v115
	*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v115
	v122 = F_setrlimit(m, int32(7), v10+int32(96))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L35
	}
L35:
	;
	if v122 != int32(-1) {
		v127 = v111
		v130 = v115
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v108 = v115
	goto L31
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(2) < v197 {
		goto L1
	} else {
		goto L54
	}
L38:
	;
	v134 = v130
	goto L40
L39:
	;
	v134 = v88
	goto L40
L40:
	;
	if base.Ui64(v91) <= base.Ui64(v134) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	*(*int32)(unsafe.Add(mBase, _consts[278])) = base.I32_wrap_i64(v134) + int32(-32)
	v144 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if base.Ui64(int64(32)) < base.Ui64(v134) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if int32(3) < v144 {
		goto L1
	} else {
		goto L47
	}
L43:
	;
	if int32(3) < v144 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v88
	F__serverLog(m, int32(3), int32(_a1582), v10+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v137
	F__serverLog(m, int32(3), int32(_a1583), v10+int32(64))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v170 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v173 = F___strerror_l(m, v127, v127)
	mBase = m.M
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v173
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v91
	F__serverLog(m, int32(3), int32(_a1584), v10+int32(48))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v183 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v134
	v188 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v188
	F__serverLog(m, int32(3), int32(_a1585), v10+int32(32))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	goto L1
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v91
	F__serverLog(m, int32(2), int32(_a1586), v10+int32(80))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	goto L1
}
func F_afterErrorReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int64
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int64
	_ = v429
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int64
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int64
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+203)))
	if v14&int32(64) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v460
	if v467 != 0 {
		goto L158
	} else {
		goto L159
	}
L2:
	;
	m.G0 = v12 + int32(48)
	return
L3:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v19 != 0 {
		v25 = v19
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = F_sdsnewlen(m, l1, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L6:
	;
	v20 = F_listCreate(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(3)
	v25 = v20
	goto L5
L9:
	;
	v28 = F_listAddNodeTail(m, v25, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L2
L11:
	;
	if l3&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v391 = v389 & int32(1)
	if v391 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L13:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v381)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v381)+128)) = v382 + int64(1)
	goto L12
L14:
	;
	v35 = int32(_a44)
	v37 = *(*int64)(unsafe.Add(mBase, _consts[94]))
	*(*int64)(unsafe.Add(mBase, _consts[94])) = v37 + int64(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v41 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if l3&int32(2) == int32(0) {
		v377 = v163
		v378 = v164
		goto L52
	} else {
		goto L53
	}
L16:
	;
	v49 = int32(32)
	if base.Ui32(l2) < base.Ui32(v49) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v163 = int32(_a978)
	v164 = int32(3)
	goto L15
L18:
	;
	v53 = l2
	goto L20
L19:
	;
	v53 = v49
	goto L20
L20:
	;
	v54 = int32(0)
	v57 = base.B2i32(v53 != v54)
	if l1&int32(3) == v54 {
		v83 = l1
		v85 = v53
		v86 = v57
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v156 != 0 {
		goto L46
	} else {
		goto L47
	}
L22:
	;
	v156 = int32(0)
	goto L21
L23:
	;
	v134 = v127
	v136 = v129
	goto L41
L24:
	;
	if v86 == int32(0) {
		goto L22
	} else {
		goto L32
	}
L25:
	;
	if v53 == int32(0) {
		v83 = l1
		v85 = v53
		v86 = v57
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v66 = l1
	v68 = v53
	goto L27
L27:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 == int32(32) {
		v127 = v66
		v129 = v68
		goto L23
	} else {
		goto L29
	}
L28:
	;
	v83 = v78
	v85 = v74
	v86 = v76
	goto L24
L29:
	;
	v74 = v68 + int32(-1)
	v75 = int32(0)
	v76 = base.B2i32(v74 != v75)
	v78 = v66 + int32(1)
	if v78&int32(3) == v75 {
		v83 = v78
		v85 = v74
		v86 = v76
		goto L24
	} else {
		goto L30
	}
L30:
	;
	if v74 != 0 {
		v66 = v78
		v68 = v74
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v90 == int32(32) {
		v120 = v83
		v122 = v85
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v122 == int32(0) {
		goto L22
	} else {
		goto L40
	}
L34:
	;
	if base.Ui32(v85) < base.Ui32(int32(4)) {
		v120 = v83
		v122 = v85
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v100 = v83
	v102 = v85
	goto L36
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v107 = v106 ^ int32(538976288)
	v110 = int32(-2139062144)
	if (int32(16843008)-v107|v107)&v110 != v110 {
		v127 = v100
		v129 = v102
		goto L23
	} else {
		goto L38
	}
L37:
	;
	v120 = v115
	v122 = v117
	goto L33
L38:
	;
	v115 = v100 + int32(4)
	v117 = v102 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v117) {
		v100 = v115
		v102 = v117
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v127 = v120
	v129 = v122
	goto L23
L41:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v139 != int32(32) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L22
L43:
	;
	v144 = v136 + int32(-1)
	if v144 != 0 {
		v134 = v134 + int32(1)
		v136 = v144
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v156 = v134
	goto L21
L45:
	;
	goto L42
L46:
	;
	v157 = l1 + int32(1)
	goto L48
L47:
	;
	v157 = int32(_a978)
	goto L48
L48:
	;
	if v156 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v162 = v156 + (l1 ^ int32(-1))
	goto L51
L50:
	;
	v162 = int32(3)
	goto L51
L51:
	;
	v163 = v157
	v164 = v162
	goto L15
L52:
	;
	F_incrementErrorCount(m, v377, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L100
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[591]))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
	goto L54
L54:
	;
	if base.Ui64(v171) < base.Ui64(int64(128)) {
		v377 = v163
		v378 = v164
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[591]))
	v177 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if v164 == v177 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v371 != 0 {
		goto L94
	} else {
		goto L95
	}
L57:
	;
	if v330 != v164 {
		v371 = v177
		goto L84
	} else {
		goto L85
	}
L58:
	;
	v321 = int32(0)
	v328 = v187
	v330 = v321
	v334 = v321
	goto L57
L59:
	;
	if base.Ui32(v187) < base.Ui32(int32(8)) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v198 = v186
	v199 = v187
	v201 = int32(0)
	goto L62
L61:
	;
	v328 = v312
	v330 = v314
	v334 = base.B2i32(v317 != int32(0))
	goto L57
L62:
	;
	v207 = int32(base.Ui32(v199) >> (uint(int32(3)) % 32))
	v208 = int32(4)
	v209 = v198 + v208
	if v199&v208 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v312 = v303
	v314 = v287
	v317 = v292
	goto L61
L64:
	;
	v292 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v209+v207+(v292-v207)&int32(3)+v280<<(uint(int32(2))%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if base.Ui32(v303) < base.Ui32(int32(8)) {
		v312 = v303
		v314 = v287
		v317 = v292
		goto L61
	} else {
		goto L82
	}
L65:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v201))))
	v258 = int32(0)
	goto L76
L66:
	;
	v214 = int32(0)
	if base.Ui32(v164) <= base.Ui32(v201) {
		v247 = v201
		v250 = v214
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v250 == v207 {
		v280 = v214
		v287 = v247
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v224 = v201
	v227 = v214
	goto L69
L69:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v227))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v224))))
	if v230 != v232 {
		v247 = v224
		v250 = v227
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v247 = v235
	v250 = v237
	goto L67
L71:
	;
	v234 = int32(1)
	v235 = v224 + v234
	v237 = v227 + v234
	if base.Ui32(v207) <= base.Ui32(v237) {
		v247 = v235
		v250 = v237
		goto L67
	} else {
		goto L72
	}
L72:
	;
	if base.Ui32(v235) < base.Ui32(v164) {
		v224 = v235
		v227 = v237
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v312 = v199
	v314 = v247
	v317 = v250
	goto L61
L75:
	;
	if v258 != v207 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+v258))))
	if v271 == v255&int32(255) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v273 = int32(1)
	v275 = v258 + v273
	if v275 != v207 {
		v258 = v275
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v328 = v199
	v330 = v201
	v334 = v273
	goto L57
L80:
	;
	v280 = v258
	v287 = v201 + int32(1)
	goto L64
L81:
	;
	v312 = v199
	v314 = v201
	v317 = v207
	goto L61
L82:
	;
	if base.Ui32(v287) < base.Ui32(v164) {
		v198 = v302
		v199 = v303
		v201 = v287
		goto L62
	} else {
		goto L83
	}
L83:
	;
	goto L63
L84:
	;
	goto L56
L85:
	;
	v336 = int32(0)
	if v328&int32(1) == v336 {
		v371 = v336
		goto L84
	} else {
		goto L86
	}
L86:
	;
	if v334&base.B2i32(v328&int32(4) != int32(0)) != 0 {
		v371 = v336
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v371 = int32(1)
	goto L84
L94:
	;
	v373 = v163
	goto L96
L95:
	;
	v373 = int32(_a979)
	goto L96
L96:
	;
	if v371 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v375 = v164
	goto L99
L98:
	;
	v375 = int32(19)
	goto L99
L99:
	;
	v377 = v373
	v378 = v375
	goto L52
L100:
	;
	goto L12
L101:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v464 != 0 {
		goto L135
	} else {
		goto L136
	}
L102:
	;
	v404 = int32(2)
	if v389&v404 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v398 = base.B2i32(v396 == int64(-1))
	if v396 == int64(-1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v399 = int32(_a980)
	goto L106
L105:
	;
	v399 = int32(_a268)
	goto L106
L106:
	;
	if v396 == int64(-1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v402 = int32(_a981)
	goto L109
L108:
	;
	v402 = int32(_a982)
	goto L109
L109:
	;
	v460 = v399
	v461 = v402
	v462 = int32(3)
	goto L101
L110:
	;
	v437 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v439 = base.B2i32(v437 == int64(-1))
	if v439 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L111:
	;
	if v389&int32(262144) != 0 {
		v428 = v404
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v389&int32(4) == int32(0) {
		v434 = int32(1)
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v429 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v429 != int64(-1) {
		goto L2
	} else {
		goto L121
	}
L115:
	;
	v417 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v418 == v417 {
		v428 = v417
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	goto L117
L117:
	;
	if v423 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v426 = int32(4)
	goto L120
L119:
	;
	v426 = int32(5)
	goto L120
L120:
	;
	v434 = v426
	goto L110
L121:
	;
	v434 = v428
	goto L110
L122:
	;
	switch v434 + int32(-1) {
	case 0:
		v460 = int32(_a982)
		v461 = int32(_a268)
		v462 = v434
		goto L101
	default:
		goto L131
	case 3:
		goto L130
	case 4:
		goto L132
	}
L123:
	;
	if v437 == int64(-1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v444 = int32(_a980)
	goto L126
L125:
	;
	v444 = int32(_a268)
	goto L126
L126:
	;
	if v437 == int64(-1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v447 = int32(_a981)
	goto L129
L128:
	;
	v447 = int32(_a982)
	goto L129
L129:
	;
	v460 = v444
	v461 = v447
	v462 = v434
	goto L101
L130:
	;
	v460 = int32(_a983)
	v461 = int32(_a984)
	v462 = v434
	goto L101
L131:
	;
	F__serverAssert(m, int32(_a195), int32(_a977), int32(941))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L7
	} else {
		goto L133
	}
L132:
	;
	v460 = int32(_a985)
	v461 = int32(_a986)
	v462 = v434
	goto L101
L133:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if int32(3) < v469 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v464)+140))
	v467 = v466
	goto L134
L136:
	;
	v467 = int32(0)
	goto L134
L137:
	;
	if v391 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	if v467 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v475 = v467
	goto L141
L140:
	;
	v475 = int32(_a301)
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(32)))) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
	v478 = int32(4096)
	if base.Ui32(l2) < base.Ui32(v478) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v481 = l2
	goto L144
L143:
	;
	v481 = v478
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v460
	F__serverLog(m, int32(3), int32(_a987), v12+int32(16))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	goto L137
L146:
	;
	v503 = int32(_a44)
	v505 = *(*int64)(unsafe.Add(mBase, _consts[306]))
	*(*int64)(unsafe.Add(mBase, _consts[306])) = v505 + int64(1)
	v510 = *(*int32)(unsafe.Add(mBase, _consts[592]))
	v511 = int32(0)
	if v391 == v511 {
		v522 = v511
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	if v494 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v494)+16))
	if v497 < int64(1) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	F_showLatestBacklog(m)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	goto L146
L151:
	;
	if v522 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	if v515 == int32(0) {
		v522 = v511
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v522 = base.B2i32(base.Ui32(v510+int32(-1)) < base.Ui32(int32(2)))
	goto L151
L154:
	;
	v523 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v523 == int64(-1))&base.B2i32(v510 == int32(1)) != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	if v462&int32(6) != int32(4) {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	F_clusterHandleSlotMigrationErrorResponse(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	goto L2
L158:
	;
	v549 = v467
	goto L160
L159:
	;
	v549 = int32(_a301)
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v549
	F__serverPanic_1(m, int32(_a977), int32(968), int32(_a988), v12)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_aggregateClientOutputBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v13 == int32(0) {
		v19 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v22 = v7 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23
	goto L6
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v17 = F_sdscatlen(m, v9, v16, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = v17
	goto L3
L6:
	;
	v28 = v7 + int32(8)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v7 + int32(16)
	return v69
L8:
	;
	if v30 == int32(0) {
		v69 = v19
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30+base.B2i32(v33 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v39
	goto L9
L11:
	;
	v43 = v30
	v45 = v19
	goto L12
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v51 = F_sdscatlen(m, v45, v47+int32(13), v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v69 = v51
	goto L7
L14:
	;
	v54 = v7 + int32(8)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v56 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v56 != 0 {
		v43 = v56
		v45 = v51
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56+base.B2i32(v59 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v65
	goto L16
L18:
	;
	goto L13
}
func F_allowProtectedAction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	switch l0 + int32(-1) {
	case 0:
		v22 = l0
		return v22
	case 1:
		v6 = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 == v6 {
			v22 = v6
			return v22
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			if v11 == int32(0) {
				v22 = v6
				return v22
			} else {
				v14 = m.T0[v11].(func(*base.Module, int32) int32)(m, v7)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v14 == int32(1))
				}
			}
		}
	default:
		v22 = int32(0)
		return v22
	}
}
func F_arg_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if base.Ui32(int32(1)) < base.Ui32(l1) {
		v15 = l0 + l1<<(uint(int32(2))%32) + int32(-4)
	} else {
		v15 = l0
	}
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v15 + int32(4)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	return v19
}
func F_atoll(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	v6 = l0
	for {
		v12 = v6 + int32(1)
		v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6))))
		if base.B2i32(v13 == int32(32))|base.B2i32(base.Ui32(v13+int32(-9)) < base.Ui32(int32(5))) != 0 {
			v6 = v12
			continue
		} else {
			break
		}
		break
	}
	v21 = int32(1)
	switch v13&int32(255) + int32(-43) {
	case 0:
		v27 = v21
		v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
		v29 = v12
		v30 = v28
		v31 = v27
	default:
		v29 = v6
		v30 = v13
		v31 = v21
	case 2:
		v27 = int32(0)
		v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
		v29 = v12
		v30 = v28
		v31 = v27
	}
	v34 = v30 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v34) {
		v58 = int64(0)
	} else {
		v38 = v34
		v39 = v29
		v42 = int64(0)
		for {
			v46 = v42*int64(10) - base.I64_extend_i32_u(v38)
			v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+1)))
			v51 = v47 + int32(-48)
			if base.Ui32(v51) < base.Ui32(int32(10)) {
				v38 = v51
				v39 = v39 + int32(1)
				v42 = v46
				continue
			} else {
				break
			}
			break
		}
		v58 = v46
	}
	if v31 != 0 {
		v61 = int64(0) - v58
	} else {
		v61 = v58
	}
	return v61
}
func F_authRequired(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v6&int32(6) == int32(4) {
		v18 = v2
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v18 = int32(base.Ui32(v11^int32(-1))>>(uint(int32(23))%32)) & int32(1)
	}
	return v18
}
func F_autoMemoryCollect(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v5&int32(1) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = v5 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 < int32(1) {
		v48 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48 | int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_valkey_free(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L22
	}
L4:
	;
	v18 = int32(0)
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = v21 + v18<<(uint(int32(3))%32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	switch v26 {
	case 0:
		goto L11
	case 1:
		goto L13
	case 2:
		goto L12
	default:
		goto L7
	case 4:
		goto L10
	case 5:
		goto L9
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v48 = v46
	goto L3
L7:
	;
	v43 = v18 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v43 < v44 {
		v18 = v43
		goto L5
	} else {
		goto L21
	}
L8:
	;
	F_valkey_free(m, v25)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L14
	} else {
		goto L20
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	F_raxFreeWithCallback(m, v36, int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L19
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	F_raxFree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L18
	}
L11:
	;
	F_VM_CloseKey(m, v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L14
	} else {
		goto L17
	}
L12:
	;
	F_VM_FreeCallReply(m, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	F_decrRefCount(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	goto L7
L16:
	;
	goto L7
L17:
	;
	goto L7
L18:
	;
	goto L8
L19:
	;
	goto L8
L20:
	;
	goto L7
L21:
	;
	goto L6
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
	goto L1
}
func F_auxupvalue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	v6 = F_luaL_checkinteger(m, l0, int32(2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_luaL_checktype(m, l0, int32(1), int32(6))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = v20 + int32(0)
			v26 = m.G398
			if base.Ui32(v25) < base.Ui32(v19) {
				v28 = v25
			} else {
				v28 = v26
			}
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
			if v71 != int32(6) {
				v78 = int32(0)
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
				v78 = base.B2i32(v75 != int32(0))
			}
			if v78 != 0 {
				v409 = v14
				return v409
			} else {
				if l1 == int32(0) {
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v203 = v198 + int32(0)
					v204 = m.G398
					if base.Ui32(v203) < base.Ui32(v197) {
						v206 = v203
					} else {
						v206 = v204
					}
					v248 = int32(0)
					v249 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
					if v249 != int32(6) {
						v317 = v248
					} else {
						v252 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
						v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+6)))
						if v253 == int32(0) {
							if v6 < int32(1) {
								v317 = v248
							} else {
								v270 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
								v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+36))
								if v271 < v6 {
									v317 = v248
								} else {
									v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
									v277 = v6<<(uint(int32(2))%32) + int32(-4)
									v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+v277)))
									v283 = *(*int32)(unsafe.Add(mBase, uint32(v252+v277)+20))
									v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
									v285 = v284
									v286 = v279 + int32(16)
									v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v290 = v288 + int32(-16)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v290
									v292 = *(*int64)(unsafe.Add(mBase, uint32(v290)))
									*(*int64)(unsafe.Add(mBase, uint32(v285))) = v292
									v296 = *(*int32)(unsafe.Add(mBase, uint32(v288+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v285)+8)) = v296
									v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
									if v299 < int32(4) {
										v317 = v286
									} else {
										v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
										v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+5)))
										if v303&int32(3) == int32(0) {
											v317 = v286
										} else {
											v308 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+5)))
											if v309&int32(4) == int32(0) {
												v317 = v286
											} else {
												F_luaC_barrierf(m, l0, v308, v302)
												mBase = m.M
												v317 = v286
											}
										}
									}
								}
							}
						} else {
							if v6 < int32(1) {
								v317 = v248
							} else {
								v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+7)))
								if base.Ui32(v258) < base.Ui32(v6) {
									v317 = v248
								} else {
									v260 = m.G3
									v285 = v6<<(uint(int32(4))%32) + v252 + int32(8)
									v286 = v260 + int32(_a139)
									v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v290 = v288 + int32(-16)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v290
									v292 = *(*int64)(unsafe.Add(mBase, uint32(v290)))
									*(*int64)(unsafe.Add(mBase, uint32(v285))) = v292
									v296 = *(*int32)(unsafe.Add(mBase, uint32(v288+int32(-8))))
									*(*int32)(unsafe.Add(mBase, uint32(v285)+8)) = v296
									v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+8))
									if v299 < int32(4) {
										v317 = v286
									} else {
										v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
										v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+5)))
										if v303&int32(3) == int32(0) {
											v317 = v286
										} else {
											v308 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+5)))
											if v309&int32(4) == int32(0) {
												v317 = v286
											} else {
												F_luaC_barrierf(m, l0, v308, v302)
												mBase = m.M
												v317 = v286
											}
										}
									}
								}
							}
						}
					}
					v320 = v317
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v92 = v87 + int32(0)
					v93 = m.G398
					if base.Ui32(v92) < base.Ui32(v86) {
						v95 = v92
					} else {
						v95 = v93
					}
					v137 = int32(0)
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
					if v138 != int32(6) {
						v189 = v137
					} else {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+6)))
						if v142 == int32(0) {
							if v6 < int32(1) {
								v189 = v137
							} else {
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+36))
								if v160 < v6 {
									v189 = v137
								} else {
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
									v166 = v6<<(uint(int32(2))%32) + int32(-4)
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v162+v166)))
									v172 = *(*int32)(unsafe.Add(mBase, uint32(v141+v166)+20))
									v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
									v174 = v173
									v176 = v168 + int32(16)
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v179 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
									*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v181
									v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v183 + int32(16)
									v189 = v176
								}
							}
						} else {
							if v6 < int32(1) {
								v189 = v137
							} else {
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+7)))
								if base.Ui32(v147) < base.Ui32(v6) {
									v189 = v137
								} else {
									v149 = m.G3
									v174 = v6<<(uint(int32(4))%32) + v141 + int32(8)
									v176 = v149 + int32(_a139)
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v179 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
									*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v181
									v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v183 + int32(16)
									v189 = v176
								}
							}
						}
					}
					v320 = v189
				}
				if v320 == int32(0) {
					v409 = v14
					return v409
				} else {
					F_lua_pushstring(m, l0, v320)
					mBase = m.M
					v324 = m.ExcPending
					if v324 != 0 {
						return int32(0)
					} else {
						v326 = l1 ^ int32(-1)
						if v326 < int32(1) {
							if v326 < int32(-9999) {
								switch v326 + int32(10002) {
								case 0:
									v381 = l0 + int32(72)
								case 1:
									v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
									v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
									v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(5)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v355
									v381 = l0 + int32(88)
								case 2:
									v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v381 = v349 + int32(96)
								default:
									v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
									v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
									v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+7)))
									v367 = m.G398
									if base.Ui32(v366) < base.Ui32(int32(-10002)-v326) {
										v378 = v367
									} else {
										v378 = v365 + (int32(-10003)-v326)<<(uint(int32(4))%32) + int32(24)
									}
									v381 = v378
								}
							} else {
								v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v381 = v343 + v326<<(uint(int32(4))%32)
							}
						} else {
							v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v337 = v332 + v326<<(uint(int32(4))%32) + int32(-16)
							v338 = m.G398
							if base.Ui32(v337) < base.Ui32(v331) {
								v340 = v337
							} else {
								v340 = v338
							}
							v381 = v340
						}
						v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if base.Ui32(v382) <= base.Ui32(v381) {
							v399 = v382
						} else {
							v385 = v382
							for {
								v389 = v385 + int32(-16)
								v390 = *(*int64)(unsafe.Add(mBase, uint32(v389)))
								*(*int64)(unsafe.Add(mBase, uint32(v385))) = v390
								v394 = *(*int32)(unsafe.Add(mBase, uint32(v385+int32(-8))))
								*(*int32)(unsafe.Add(mBase, uint32(v385)+8)) = v394
								if base.Ui32(v381) < base.Ui32(v389) {
									v385 = v389
									continue
								} else {
									break
								}
								break
							}
							v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v399 = v397
						}
						v402 = *(*int64)(unsafe.Add(mBase, uint32(v399)))
						*(*int64)(unsafe.Add(mBase, uint32(v381))) = v402
						v404 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v381)+8)) = v404
						v409 = l1 + int32(1)
						return v409
					}
				}
			}
		}
	}
}
