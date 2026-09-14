package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_freeReplicationBacklog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v4 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	if v5 != 0 {
		F__serverAssert(m, int32(_a1189), int32(_a1190), int32(160))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[370]))
		if v7 == int32(0) {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if v10 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[217]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				F_freeReplicationBacklogRefMemAsync(m, v21, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(_a44)
					*(*int32)(unsafe.Add(mBase, _consts[288])) = int32(0)
					v29 = F_listCreate(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[217])) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(102)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[370]))
						F_valkey_free(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[370])) = int32(0)
							return
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v14 != int32(1) {
					F__serverAssert(m, int32(_a1191), int32(_a1190), int32(166))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(0)
					v21 = *(*int32)(unsafe.Add(mBase, _consts[217]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					F_freeReplicationBacklogRefMemAsync(m, v21, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[288])) = int32(0)
						v29 = F_listCreate(m)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[217])) = v29
							*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(102)
							v35 = *(*int32)(unsafe.Add(mBase, _consts[370]))
							F_valkey_free(m, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[370])) = int32(0)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_replicationAbortDualChannelSyncTransfer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	v3 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v3 == int32(0) {
		F__serverAssert(m, int32(_a1256), int32(_a1190), int32(2983))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		if int32(2) < v7 {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[699]))
			if v16 == int32(0) {
				F_cleanupTransferResources(m)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = int32(_a44)
					v29 = int64(0)
					*(*int64)(unsafe.Add(mBase, _consts[679])) = v29
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[700])) = v32
					*(*int32)(unsafe.Add(mBase, _consts[697])) = v32
					*(*int32)(unsafe.Add(mBase, _consts[681])) = int32(-1)
					*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
					*(*int64)(unsafe.Add(mBase, _consts[688])) = v29
					v49 = *(*int32)(unsafe.Add(mBase, _consts[695]))
					if v49 == v32 {
						v54 = int32(_a44)
						*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
						*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
						return
					} else {
						F_freePendingReplDataBufAsync(m, v49)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v54 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
							return
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
				m.T0[v20].(func(*base.Module, int32))(m, v16)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[699])) = int32(0)
					F_cleanupTransferResources(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = int32(_a44)
						v29 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[679])) = v29
						v32 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[700])) = v32
						*(*int32)(unsafe.Add(mBase, _consts[697])) = v32
						*(*int32)(unsafe.Add(mBase, _consts[681])) = int32(-1)
						*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
						*(*int64)(unsafe.Add(mBase, _consts[688])) = v29
						v49 = *(*int32)(unsafe.Add(mBase, _consts[695]))
						if v49 == v32 {
							v54 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
							return
						} else {
							F_freePendingReplDataBufAsync(m, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v54 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
								return
							}
						}
					}
				}
			}
		} else {
			F__serverLog(m, int32(2), int32(_a1257), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[699]))
				if v16 == int32(0) {
					F_cleanupTransferResources(m)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = int32(_a44)
						v29 = int64(0)
						*(*int64)(unsafe.Add(mBase, _consts[679])) = v29
						v32 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[700])) = v32
						*(*int32)(unsafe.Add(mBase, _consts[697])) = v32
						*(*int32)(unsafe.Add(mBase, _consts[681])) = int32(-1)
						*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
						*(*int64)(unsafe.Add(mBase, _consts[688])) = v29
						v49 = *(*int32)(unsafe.Add(mBase, _consts[695]))
						if v49 == v32 {
							v54 = int32(_a44)
							*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
							*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
							return
						} else {
							F_freePendingReplDataBufAsync(m, v49)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v54 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
								return
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
					m.T0[v20].(func(*base.Module, int32))(m, v16)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[699])) = int32(0)
						F_cleanupTransferResources(m)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = int32(_a44)
							v29 = int64(0)
							*(*int64)(unsafe.Add(mBase, _consts[679])) = v29
							v32 = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[700])) = v32
							*(*int32)(unsafe.Add(mBase, _consts[697])) = v32
							*(*int32)(unsafe.Add(mBase, _consts[681])) = int32(-1)
							*(*int64)(unsafe.Add(mBase, _consts[698])) = int64(-1)
							*(*int64)(unsafe.Add(mBase, _consts[688])) = v29
							v49 = *(*int32)(unsafe.Add(mBase, _consts[695]))
							if v49 == v32 {
								v54 = int32(_a44)
								*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
								*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
								return
							} else {
								F_freePendingReplDataBufAsync(m, v49)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = int32(_a44)
									*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
									*(*int64)(unsafe.Add(mBase, _consts[695])) = int64(0)
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
func F_replicationCountAcksByOffset(m *base.Module, l0 int64) int32 {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v12 = v7 + int32(8)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
	v17 = int32(0)
	v19 = v7 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == v17 {
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	}
	if v21 == int32(0) {
		v63 = v17
	} else {
		v36 = v17
		v37 = v21
		for {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			if v40 != int32(9) {
				v46 = v36
			} else {
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v39)+64))
				v46 = v36 + base.B2i32(l0 <= v43)
			}
			v48 = v7 + int32(8)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if v50 == int32(0) {
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v50+base.B2i32(v53 == int32(0))<<(uint(int32(2))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
			}
			if v50 != 0 {
				v36 = v46
				v37 = v50
				continue
			} else {
				break
			}
			break
		}
		v63 = v46
	}
	m.G0 = v7 + int32(16)
	return v63
}
func F_replicationCreatePrimaryClientWithHandler(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v71 int32
	_ = v71
	var v78 int64
	_ = v78
	var v85 int64
	_ = v85
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v7 = F_createClient(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[188])) = v7
		if l0 == int32(0) {
			v19 = v7
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v20 | int32(1)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[188]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+328)) = int32(0)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+204))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+204)) = v30&int32(-25165825) | int32(_a0) | int32(16777216)
			v45 = F_sdsempty(m)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _consts[188]))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v45
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+104))
				if v50 != 0 {
					v58 = v48
					v59 = v50
					v60 = int32(_a44)
					v61 = *(*int64)(unsafe.Add(mBase, _consts[672]))
					*(*int64)(unsafe.Add(mBase, uint32(v59)+40)) = v61
					*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v61
					*(*int32)(unsafe.Add(mBase, uint32(v58)+328)) = int32(0)
					v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[673])))
					*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(144)))) = uint8(v71)
					v78 = *(*int64)(unsafe.Add(mBase, _consts[674]))
					*(*int64)(unsafe.Add(mBase, uint32(v59+int32(136)))) = v78
					v85 = *(*int64)(unsafe.Add(mBase, _consts[675]))
					*(*int64)(unsafe.Add(mBase, uint32(v59+int32(128)))) = v85
					v92 = *(*int64)(unsafe.Add(mBase, _consts[676]))
					*(*int64)(unsafe.Add(mBase, uint32(v59+int32(120)))) = v92
					v99 = *(*int64)(unsafe.Add(mBase, _consts[677]))
					*(*int64)(unsafe.Add(mBase, uint32(v59+int32(112)))) = v99
					v102 = *(*int64)(unsafe.Add(mBase, _consts[678]))
					*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v102
					v105 = *(*int32)(unsafe.Add(mBase, _consts[188]))
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+48))
					if v107 != int64(-1) {
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+200))
						*(*int32)(unsafe.Add(mBase, uint32(v105)+200)) = v110 | int32(65536)
					}
					if l1 == int32(-1) {
						return
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, _consts[188]))
						v118 = F_selectDb(m, v117, l1)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v52 = F_valkey_calloc(m, int32(200))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+104)) = v52
						v56 = *(*int32)(unsafe.Add(mBase, _consts[188]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+104))
						v58 = v56
						v59 = v57
						v60 = int32(_a44)
						v61 = *(*int64)(unsafe.Add(mBase, _consts[672]))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+40)) = v61
						*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v58)+328)) = int32(0)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[673])))
						*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(144)))) = uint8(v71)
						v78 = *(*int64)(unsafe.Add(mBase, _consts[674]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(136)))) = v78
						v85 = *(*int64)(unsafe.Add(mBase, _consts[675]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(128)))) = v85
						v92 = *(*int64)(unsafe.Add(mBase, _consts[676]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(120)))) = v92
						v99 = *(*int64)(unsafe.Add(mBase, _consts[677]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(112)))) = v99
						v102 = *(*int64)(unsafe.Add(mBase, _consts[678]))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v102
						v105 = *(*int32)(unsafe.Add(mBase, _consts[188]))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
						v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+48))
						if v107 != int64(-1) {
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v105)+200)) = v110 | int32(65536)
						}
						if l1 == int32(-1) {
							return
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, _consts[188]))
							v118 = F_selectDb(m, v117, l1)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
			v15 = m.T0[v14].(func(*base.Module, int32, int32) int32)(m, v12, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[188]))
				v19 = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v20 | int32(1)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[188]))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+328)) = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+204)) = v30&int32(-25165825) | int32(_a0) | int32(16777216)
				v45 = F_sdsempty(m)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _consts[188]))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v45
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+104))
					if v50 != 0 {
						v58 = v48
						v59 = v50
						v60 = int32(_a44)
						v61 = *(*int64)(unsafe.Add(mBase, _consts[672]))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+40)) = v61
						*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v58)+328)) = int32(0)
						v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[673])))
						*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(144)))) = uint8(v71)
						v78 = *(*int64)(unsafe.Add(mBase, _consts[674]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(136)))) = v78
						v85 = *(*int64)(unsafe.Add(mBase, _consts[675]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(128)))) = v85
						v92 = *(*int64)(unsafe.Add(mBase, _consts[676]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(120)))) = v92
						v99 = *(*int64)(unsafe.Add(mBase, _consts[677]))
						*(*int64)(unsafe.Add(mBase, uint32(v59+int32(112)))) = v99
						v102 = *(*int64)(unsafe.Add(mBase, _consts[678]))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v102
						v105 = *(*int32)(unsafe.Add(mBase, _consts[188]))
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
						v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+48))
						if v107 != int64(-1) {
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+200))
							*(*int32)(unsafe.Add(mBase, uint32(v105)+200)) = v110 | int32(65536)
						}
						if l1 == int32(-1) {
							return
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, _consts[188]))
							v118 = F_selectDb(m, v117, l1)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v52 = F_valkey_calloc(m, int32(200))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+104)) = v52
							v56 = *(*int32)(unsafe.Add(mBase, _consts[188]))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+104))
							v58 = v56
							v59 = v57
							v60 = int32(_a44)
							v61 = *(*int64)(unsafe.Add(mBase, _consts[672]))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+40)) = v61
							*(*int64)(unsafe.Add(mBase, uint32(v59)+48)) = v61
							*(*int32)(unsafe.Add(mBase, uint32(v58)+328)) = int32(0)
							v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[673])))
							*(*uint8)(unsafe.Add(mBase, uint32(v59+int32(144)))) = uint8(v71)
							v78 = *(*int64)(unsafe.Add(mBase, _consts[674]))
							*(*int64)(unsafe.Add(mBase, uint32(v59+int32(136)))) = v78
							v85 = *(*int64)(unsafe.Add(mBase, _consts[675]))
							*(*int64)(unsafe.Add(mBase, uint32(v59+int32(128)))) = v85
							v92 = *(*int64)(unsafe.Add(mBase, _consts[676]))
							*(*int64)(unsafe.Add(mBase, uint32(v59+int32(120)))) = v92
							v99 = *(*int64)(unsafe.Add(mBase, _consts[677]))
							*(*int64)(unsafe.Add(mBase, uint32(v59+int32(112)))) = v99
							v102 = *(*int64)(unsafe.Add(mBase, _consts[678]))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v102
							v105 = *(*int32)(unsafe.Add(mBase, _consts[188]))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+104))
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+48))
							if v107 != int64(-1) {
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+200))
								*(*int32)(unsafe.Add(mBase, uint32(v105)+200)) = v110 | int32(65536)
							}
							if l1 == int32(-1) {
								return
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, _consts[188]))
								v118 = F_selectDb(m, v117, l1)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
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
func F_replicationGetReplicaOffset(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v5 == int32(0) {
		v19 = int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[188]))
		if v9 != 0 {
			v16 = v9
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
			v19 = v18
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[189]))
			if v12 == int32(0) {
				v19 = int64(0)
			} else {
				v16 = v12
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+104))
				v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
				v19 = v18
			}
		}
	}
	v21 = int64(0)
	if v21 < v19 {
		v24 = v19
	} else {
		v24 = v21
	}
	return v24
}
func F_replicationStartPendingFork(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v25 int32
	_ = v25
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v6 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v6
	v18 = F_shouldStartChildReplication(m, v4+int32(12), v4+int32(8), v4+int32(4))
	mBase = m.M
	if v18 == int32(0) {
		m.G0 = v4 + int32(16)
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v24 = F_startBgsaveForReplication(m, v21, v22, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			m.G0 = v4 + int32(16)
			return
		}
	}
}
func F_resetReplicationBuffer(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = int32(_a44)
	*(*int32)(unsafe.Add(mBase, _consts[288])) = int32(0)
	v6 = F_listCreate(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[217])) = v6
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(102)
		return
	}
}
