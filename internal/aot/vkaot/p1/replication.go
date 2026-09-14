package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createReplicationBacklog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int64
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v4 == int32(0) {
		v15 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[293])) = v15
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(0)
			v20 = F_raxNew(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(_a20)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[293]))
				*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v20
				v28 = *(*int64)(unsafe.Add(mBase, _consts[47]))
				*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v28 + int64(1)
				return
			}
		}
	} else {
		F__serverAssert(m, int32(_a1912), int32(_a1913), int32(136))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
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
func F_replicationAttachToNewPrimary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v3 != 0 {
		F__serverAssert(m, int32(_a1950), int32(_a1913), int32(2259))
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
		v5 = *(*int32)(unsafe.Add(mBase, _consts[168]))
		if v5 == int32(0) {
			F_clusterCleanSlotImportsOnFullSync(m)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_disconnectReplicas(m)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_freeReplicationBacklog(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(2) < v9 {
				v19 = v5
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v20 & int32(-2)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[168]))
				v26 = F_freeClient(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[168])) = int32(0)
					F_clusterCleanSlotImportsOnFullSync(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_disconnectReplicas(m)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_freeReplicationBacklog(m)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F__serverLog(m, int32(2), int32(_a1939), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[168]))
					v19 = v18
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+200))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+200)) = v20 & int32(-2)
					v25 = *(*int32)(unsafe.Add(mBase, _consts[168]))
					v26 = F_freeClient(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[168])) = int32(0)
						F_clusterCleanSlotImportsOnFullSync(m)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_disconnectReplicas(m)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_freeReplicationBacklog(m)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
func F_replicationCachePrimary(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v4 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v4 == int32(0) {
		F__serverAssert(m, int32(_a2004), int32(_a1913), int32(4756))
		mBase = m.M
		v124 = m.ExcPending
		if v124 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[168]))
		if v8 != 0 {
			F__serverAssert(m, int32(_a2004), int32(_a1913), int32(4756))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(2) < v10 {
				F_unlinkClient(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v25 = v22 + int32(-1)
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
					switch v26 & int32(7) {
					case 0:
						v29 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v29)
					case 1:
						v33 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))) = uint8(v33)
					case 2:
						v37 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))) = uint16(v37)
					case 3:
						*(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9)))) = int32(0)
					case 4:
						*(*int64)(unsafe.Add(mBase, uint32(v22+int32(-17)))) = int64(0)
					default:
					}
					v47 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v47)
					v50 = *(*int32)(unsafe.Add(mBase, _consts[167]))
					v51 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+56)) = int64(0)
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+40)) = v56
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
					if v58&int32(8) == v51 {
						F_releaseReplyReferences(m, l0)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
							F_listEmpty(m, v71)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
								F_resetClient(m, l0)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									v80 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v80)
									v82 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v82
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v80
									*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v82
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v80)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v90 & int32(-3)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
									F_discardCommandQueue(m, l0)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										v100 = int32(_a20)
										v102 = *(*int32)(unsafe.Add(mBase, _consts[167]))
										*(*int32)(unsafe.Add(mBase, _consts[168])) = v102
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
										if v104 == int32(0) {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
											if v111 == int32(0) {
												F_replicationHandlePrimaryDisconnection(m)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													return
												}
											} else {
												F_sdsfree(m, v111)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
													F_replicationHandlePrimaryDisconnection(m)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														return
													}
												}
											}
										} else {
											F_sdsfree(m, v104)
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = int32(0)
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
												if v111 == int32(0) {
													F_replicationHandlePrimaryDisconnection(m)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														return
													}
												} else {
													F_sdsfree(m, v111)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
														F_replicationHandlePrimaryDisconnection(m)
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
					} else {
						F_discardTransaction(m, l0)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_releaseReplyReferences(m, l0)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								F_listEmpty(m, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
									F_resetClient(m, l0)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v80)
										v82 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v82
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v80
										*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v82
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v80)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v90 & int32(-3)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
										F_discardCommandQueue(m, l0)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = int32(_a20)
											v102 = *(*int32)(unsafe.Add(mBase, _consts[167]))
											*(*int32)(unsafe.Add(mBase, _consts[168])) = v102
											v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
											if v104 == int32(0) {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
												if v111 == int32(0) {
													F_replicationHandlePrimaryDisconnection(m)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														return
													}
												} else {
													F_sdsfree(m, v111)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
														F_replicationHandlePrimaryDisconnection(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															return
														}
													}
												}
											} else {
												F_sdsfree(m, v104)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = int32(0)
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
													if v111 == int32(0) {
														F_replicationHandlePrimaryDisconnection(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															return
														}
													} else {
														F_sdsfree(m, v111)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
															F_replicationHandlePrimaryDisconnection(m)
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
						}
					}
				}
			} else {
				F__serverLog(m, int32(2), int32(_a2005), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_unlinkClient(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
						v25 = v22 + int32(-1)
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
						switch v26 & int32(7) {
						case 0:
							v29 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v29)
						case 1:
							v33 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-3)))) = uint8(v33)
						case 2:
							v37 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-5)))) = uint16(v37)
						case 3:
							*(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9)))) = int32(0)
						case 4:
							*(*int64)(unsafe.Add(mBase, uint32(v22+int32(-17)))) = int64(0)
						default:
						}
						v47 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v47)
						v50 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						v51 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v53)+56)) = int64(0)
						v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v53)+40)) = v56
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
						if v58&int32(8) == v51 {
							F_releaseReplyReferences(m, l0)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								F_listEmpty(m, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
									F_resetClient(m, l0)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v80)
										v82 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v82
										*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v80
										*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v82
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v80)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v90 & int32(-3)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
										F_discardCommandQueue(m, l0)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v100 = int32(_a20)
											v102 = *(*int32)(unsafe.Add(mBase, _consts[167]))
											*(*int32)(unsafe.Add(mBase, _consts[168])) = v102
											v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
											if v104 == int32(0) {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
												if v111 == int32(0) {
													F_replicationHandlePrimaryDisconnection(m)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														return
													}
												} else {
													F_sdsfree(m, v111)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
														F_replicationHandlePrimaryDisconnection(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															return
														}
													}
												}
											} else {
												F_sdsfree(m, v104)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = int32(0)
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
													if v111 == int32(0) {
														F_replicationHandlePrimaryDisconnection(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															return
														}
													} else {
														F_sdsfree(m, v111)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
															F_replicationHandlePrimaryDisconnection(m)
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
						} else {
							F_discardTransaction(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_releaseReplyReferences(m, l0)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = int64(0)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									F_listEmpty(m, v71)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = int64(0)
										F_resetClient(m, l0)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v80)
											v82 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v82
											*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v80
											*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v82
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v80)
											v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v90 & int32(-3)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
											F_discardCommandQueue(m, l0)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v100 = int32(_a20)
												v102 = *(*int32)(unsafe.Add(mBase, _consts[167]))
												*(*int32)(unsafe.Add(mBase, _consts[168])) = v102
												v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
												if v104 == int32(0) {
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
													if v111 == int32(0) {
														F_replicationHandlePrimaryDisconnection(m)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															return
														}
													} else {
														F_sdsfree(m, v111)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
															F_replicationHandlePrimaryDisconnection(m)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																return
															}
														}
													}
												} else {
													F_sdsfree(m, v104)
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = int32(0)
														v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
														if v111 == int32(0) {
															F_replicationHandlePrimaryDisconnection(m)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																return
															}
														} else {
															F_sdsfree(m, v111)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = int32(0)
																F_replicationHandlePrimaryDisconnection(m)
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
							}
						}
					}
				}
			}
		}
	}
}
func F_replicationEmptyDbCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	if v4 != int32(13) {
		return
	} else {
		v7 = int32(0)
		v8 = F___time(m, v7)
		mBase = m.M
		v10 = *(*int64)(unsafe.Add(mBase, _consts[598]))
		if v8 == v10 {
			return
		} else {
			v12 = int32(0)
			v14 = F___time(m, v12)
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, _consts[598])) = v14
			v17 = *(*int32)(unsafe.Add(mBase, _consts[592]))
			if v17 == v12 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
				v24 = m.T0[v23].(func(*base.Module, int32, int32, int32) int32)(m, v17, int32(_a26), int32(1))
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
func F_replicationFeedReplicas(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1924), int32(_a1913), int32(565))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L84
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if l0 < int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if v18 <= l0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
L6:
	;
	m.G0 = v11 + int32(48)
	return
L7:
	;
	v22 = int32(_a20)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v25 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = v11 + int32(16)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v41
	goto L14
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F__serverAssert(m, int32(_a1925), int32(_a1913), int32(585))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v27 = int32(_a20)
	v29 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[47])) = v29 + int64(1)
	goto L6
L12:
	;
	return
L13:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v46 = v11 + int32(16)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if l0 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L16:
	;
	if v48 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48+base.B2i32(v51 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v57
	goto L17
L19:
	;
	v65 = v48
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+205)))
	if v70&int32(32) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L15
L22:
	;
	v80 = v11 + int32(16)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+104))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 == int32(6) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v77 = F_prepareClientToWrite(m, v69)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	if v82 != 0 {
		v65 = v82
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82+base.B2i32(v85 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v91
	goto L27
L29:
	;
	goto L21
L30:
	;
	v207 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v207)
	v212 = v11 + int32(16) | int32(1)
	v214 = base.I64_extend_i32_s(l2)
	if v214 <= int64(-1) {
		goto L61
	} else {
		goto L62
	}
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[294]))
	if v104 == l0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(int32(9)) < base.Ui32(l0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_feedReplicationBufferWithObject(m, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L48
	}
L34:
	;
	v114 = v11 + int32(16)
	v116 = base.I64_extend_i32_s(l0)
	if v116 <= int64(-1) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[570])))
	v170 = v112
	goto L33
L36:
	;
	v158 = F_sdsempty(m)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L45
	}
L37:
	;
	v157 = int32(0)
	goto L36
L39:
	;
	v138 = F_ull2string(m, v134, v135, v136)
	mBase = m.M
	if v138 == int32(0) {
		goto L37
	} else {
		goto L43
	}
L40:
	;
	goto L42
L41:
	;
	v134 = v114
	v135 = int32(21)
	v136 = v116
	v137 = int32(0)
	goto L39
L42:
	;
	v125 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v125)
	v134 = v11 + int32(17)
	v135 = int32(20)
	v136 = int64(0) - v116
	v137 = int32(1)
	goto L39
L43:
	;
	v157 = v138 + v137
	goto L36
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v11 + int32(16)
	v166 = F_sdscatfmt(m, v158, int32(_a1926), v11)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	v168 = F_createObject(m, int32(0), v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v170 = v168
	goto L33
L48:
	;
	v175 = F_objectGetVal(m, v170)
	mBase = m.M
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-1)))))
	switch v178 & int32(7) {
	case 0:
		goto L54
	case 1:
		goto L53
	case 2:
		goto L52
	case 3:
		goto L51
	case 4:
		goto L50
	default:
		v195 = int32(0)
		goto L49
	}
L49:
	;
	F_clusterSlotStatsDecrNetworkBytesOutForReplication(m, base.I64_extend_i32_u(v195))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L55
	}
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-17))))
	v195 = v194
	goto L49
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v175+int32(-9))))
	v195 = v191
	goto L49
L52:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(-5)))))
	v195 = v188
	goto L49
L53:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(-3)))))
	v195 = v185
	goto L49
L54:
	;
	v195 = int32(base.Ui32(v178) >> (uint(int32(3)) % 32))
	goto L49
L55:
	;
	F_decrRefCount(m, v170)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[294])) = l0
	goto L30
L57:
	;
	v257 = v11 + int32(16)
	v261 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v257+int32(1)))) = uint16(v261)
	F_feedReplicationBuffer(m, v257, v255+int32(3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L66
	}
L58:
	;
	v255 = int32(0)
	goto L57
L60:
	;
	v236 = F_ull2string(m, v232, v233, v234)
	mBase = m.M
	if v236 == int32(0) {
		goto L58
	} else {
		goto L64
	}
L61:
	;
	goto L63
L62:
	;
	v232 = v212
	v233 = int32(23)
	v234 = v214
	v235 = int32(0)
	goto L60
L63:
	;
	v223 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v223)
	v227 = int32(1)
	v232 = v212 + v227
	v233 = int32(22)
	v234 = int64(0) - v214
	v235 = v227
	goto L60
L64:
	;
	v255 = v236 + v235
	goto L57
L66:
	;
	if l2 < int32(1) {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v276 = int32(0)
	goto L68
L68:
	;
	v282 = l1 + v276<<(uint(int32(2))%32)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v284 = F_stringObjectLen(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L70
	}
L69:
	;
	goto L6
L70:
	;
	v286 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)) = uint8(v286)
	v289 = base.I64_extend_i32_s(v284)
	if v289 <= int64(-1) {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v332 = v11 + int32(16)
	v336 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v330+v332+int32(1)))) = uint16(v336)
	F_feedReplicationBuffer(m, v332, v330+int32(3))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L12
	} else {
		goto L80
	}
L72:
	;
	v330 = int32(0)
	goto L71
L74:
	;
	v311 = F_ull2string(m, v307, v308, v309)
	mBase = m.M
	if v311 == int32(0) {
		goto L72
	} else {
		goto L78
	}
L75:
	;
	goto L77
L76:
	;
	v307 = v212
	v308 = int32(23)
	v309 = v289
	v310 = int32(0)
	goto L74
L77:
	;
	v298 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v298)
	v302 = int32(1)
	v307 = v212 + v302
	v308 = int32(22)
	v309 = int64(0) - v289
	v310 = v302
	goto L74
L78:
	;
	v330 = v311 + v310
	goto L71
L80:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	F_feedReplicationBufferWithObject(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L81
	}
L81:
	;
	F_feedReplicationBuffer(m, v212+v330, int32(2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	v352 = v276 + int32(1)
	if v352 != l2 {
		v276 = v352
		goto L68
	} else {
		goto L83
	}
L83:
	;
	goto L69
L84:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replicationRequestAckFromReplicas(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[637])) = int32(1)
	return
}
func F_replicationSendNewlineToPrimary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v1 = int32(0)
	v3 = F___time(m, v1)
	mBase = m.M
	v5 = *(*int64)(unsafe.Add(mBase, _consts[598]))
	if v3 == v5 {
		return
	} else {
		v7 = int32(0)
		v9 = F___time(m, v7)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[598])) = v9
		v12 = *(*int32)(unsafe.Add(mBase, _consts[592]))
		if v12 == v7 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
			v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, v12, int32(_a26), int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_replicationSetupReplicaForFullResync(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(7)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = l1
	*(*int32)(unsafe.Add(mBase, _consts[294])) = int32(-1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v19&int32(1) != 0 {
		v46 = v3
		m.G0 = v9 + int32(144)
		return v46
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a1818)
		v31 = F_snprintf(m, v9+int32(16), int32(128), int32(_a1935), v9)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
			v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int32)(m, v35, v9+int32(16), v31)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v31 == v40 {
					v46 = v3
					m.G0 = v9 + int32(144)
					return v46
				} else {
					F_freeClientAsync(m, l0)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v46 = int32(-1)
						m.G0 = v9 + int32(144)
						return v46
					}
				}
			}
		}
	}
}
func F_replicationUnsetPrimary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int64
	_ = v80
	var v88 int64
	_ = v88
	var v96 int64
	_ = v96
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int64
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(32)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[423]))
	if v13 != int32(14) {
		v23 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_sdsfree(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L4:
	;
	F_moduleFireServerEvent(m, int64(7), int32(1), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v23 = v22
	goto L3
L7:
	;
	v26 = int32(_a20)
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[166])) = v27
	v30 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if v30 == v27 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v33 = F_freeClient(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v64 = F_cancelReplicationHandshake(m, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v40 {
		v50 = v36
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+200)) = v51 & int32(-2)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v57 = F_freeClient(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	F__serverLog(m, int32(2), int32(_a1939), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v50 = v49
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[168])) = int32(0)
	goto L11
L17:
	;
	v66 = int32(_a20)
	v71 = int32(_a1936)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _consts[571])))
	*(*uint8)(unsafe.Add(mBase, _consts[572])) = uint8(v72)
	v80 = *(*int64)(unsafe.Add(mBase, _consts[573]))
	*(*int64)(unsafe.Add(mBase, _consts[574])) = v80
	v88 = *(*int64)(unsafe.Add(mBase, _consts[575]))
	*(*int64)(unsafe.Add(mBase, _consts[576])) = v88
	v96 = *(*int64)(unsafe.Add(mBase, _consts[577]))
	*(*int64)(unsafe.Add(mBase, _consts[578])) = v96
	v104 = *(*int64)(unsafe.Add(mBase, _consts[579]))
	*(*int64)(unsafe.Add(mBase, _consts[580])) = v104
	v108 = *(*int64)(unsafe.Add(mBase, _consts[581]))
	*(*int64)(unsafe.Add(mBase, _consts[582])) = v108
	v112 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, _consts[583])) = v112 + int64(1)
	v118 = int32(_a1818)
	F_getRandomHexChars(m, v118, int32(40))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[571])) = uint8(v122)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v125 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_disconnectReplicas(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(16)))) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a1937)
	v136 = *(*int64)(unsafe.Add(mBase, _consts[583]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v136
	F__serverLog(m, int32(2), int32(_a1938), v6)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v144 = int32(_a20)
	v145 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _consts[294])) = v145
	*(*int32)(unsafe.Add(mBase, _consts[423])) = int32(0)
	v151 = F_setOOMScoreAdj(m, v145)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v153 = int32(_a20)
	v154 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[584])) = v154
	v158 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	*(*int64)(unsafe.Add(mBase, _consts[585])) = v158
	v161 = int32(0)
	F_moduleFireServerEvent(m, v154, v161, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[586]))
	if v166 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_clusterCleanSlotImportsOnPromotion(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L29
	}
L26:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v170 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_restartAOFAfterSYNC(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L1
}
