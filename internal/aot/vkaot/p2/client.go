package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clearClientConnectionState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(4) == int32(0) {
		v23 = v4
		if v23&int32(3) != 0 {
			F__serverAssert(m, int32(_a_F_clearClientConnectionState_0), int32(_a_F_clearClientConnectionState_1), int32(2098))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			if v26 != 0 {
				F__serverAssert(m, int32(_a_F_clearClientConnectionState_0), int32(_a_F_clearClientConnectionState_1), int32(2098))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
				if v27&int32(4) == int32(0) {
					v35 = F_selectDb(m, l0, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = int32(2)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)) = uint8(v37)
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[0]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v40
						v45 = int32(1)
						v51 = int32(base.Ui32(v41^int32(-1))>>(uint(v45)%32)) & int32(base.Ui32(v41)>>(uint(v37)%32)) & v45
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						if v51 != 0 {
							v63 = v51<<(uint(int32(23))%32) | v54&int32(-25165825) | v51<<(uint(int32(24))%32)
						} else {
							v63 = v54 & int32(-8388609)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
						F_moduleNotifyUserChanged(m, l0)
						mBase = m.M
						F_discardTransaction(m, l0)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_freeClientPubSubData(m, l0)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
								if v70 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
									return
								} else {
									F_decrRefCount(m, v70)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = int32(0)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
										return
									}
								}
							}
						}
					}
				} else {
					F_disableTracking(m, l0)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v35 = F_selectDb(m, l0, int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)) = uint8(v37)
							v40 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[0]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v40
							v45 = int32(1)
							v51 = int32(base.Ui32(v41^int32(-1))>>(uint(v45)%32)) & int32(base.Ui32(v41)>>(uint(v37)%32)) & v45
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							if v51 != 0 {
								v63 = v51<<(uint(int32(23))%32) | v54&int32(-25165825) | v51<<(uint(int32(24))%32)
							} else {
								v63 = v54 & int32(-8388609)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
							F_moduleNotifyUserChanged(m, l0)
							mBase = m.M
							F_discardTransaction(m, l0)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_freeClientPubSubData(m, l0)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
									if v70 == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
										return
									} else {
										F_decrRefCount(m, v70)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = int32(0)
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
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
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[1]))
		v11 = F_listSearchKey(m, v10, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 == int32(0) {
				F__serverAssert(m, int32(_a_F_clearClientConnectionState_2), int32(_a_F_clearClientConnectionState_1), int32(2091))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[1]))
				F_listDelNode(m, v16, v11)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					v21 = v19 & int32(-7)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v21
					v23 = v21
					if v23&int32(3) != 0 {
						F__serverAssert(m, int32(_a_F_clearClientConnectionState_0), int32(_a_F_clearClientConnectionState_1), int32(2098))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v26 != 0 {
							F__serverAssert(m, int32(_a_F_clearClientConnectionState_0), int32(_a_F_clearClientConnectionState_1), int32(2098))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
							if v27&int32(4) == int32(0) {
								v35 = F_selectDb(m, l0, int32(0))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v37 = int32(2)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)) = uint8(v37)
									v40 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[0]))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v40
									v45 = int32(1)
									v51 = int32(base.Ui32(v41^int32(-1))>>(uint(v45)%32)) & int32(base.Ui32(v41)>>(uint(v37)%32)) & v45
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
									if v51 != 0 {
										v63 = v51<<(uint(int32(23))%32) | v54&int32(-25165825) | v51<<(uint(int32(24))%32)
									} else {
										v63 = v54 & int32(-8388609)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
									F_moduleNotifyUserChanged(m, l0)
									mBase = m.M
									F_discardTransaction(m, l0)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_freeClientPubSubData(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
											if v70 == int32(0) {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
												return
											} else {
												F_decrRefCount(m, v70)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = int32(0)
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
													return
												}
											}
										}
									}
								}
							} else {
								F_disableTracking(m, l0)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v35 = F_selectDb(m, l0, int32(0))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return
									} else {
										v37 = int32(2)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)) = uint8(v37)
										v40 = *(*int32)(unsafe.Add(mBase, _c_F_clearClientConnectionState[0]))
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v40
										v45 = int32(1)
										v51 = int32(base.Ui32(v41^int32(-1))>>(uint(v45)%32)) & int32(base.Ui32(v41)>>(uint(v37)%32)) & v45
										v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
										if v51 != 0 {
											v63 = v51<<(uint(int32(23))%32) | v54&int32(-25165825) | v51<<(uint(int32(24))%32)
										} else {
											v63 = v54 & int32(-8388609)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
										F_moduleNotifyUserChanged(m, l0)
										mBase = m.M
										F_discardTransaction(m, l0)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_freeClientPubSubData(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
												if v70 == int32(0) {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
													return
												} else {
													F_decrRefCount(m, v70)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = int32(0)
														v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v77 & int32(-100794881)
														v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v81 & int32(-536952833)
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
func F_clientCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = F_objectGetVal(m, v10)
	mBase = m.M
	v12 = F_sdsnew(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-1)))))
		switch v19 & int32(7) {
		case 0:
			v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 1:
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(-3)))))
			v36 = v26
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 2:
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(-5)))))
			v36 = v29
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 3:
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-9))))
			v36 = v32
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		case 4:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(-17))))
			v36 = v35
			if v36 == int32(0) {
			} else {
				v41 = int32(0)
				for {
					v44 = v12 + v41
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44))))
					v46 = F_toupper(m, v45)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v46)
					v49 = v41 + int32(1)
					if v49 != v36 {
						v41 = v49
						continue
					} else {
						break
					}
					break
				}
			}
		default:
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
		v57 = F_objectGetVal(m, v56)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
		F_addReplyErrorFormat(m, l0, int32(_a_F_clientCommand_0), v7)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			F_sdsfree(m, v12)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_clientCommandArgShouldBeRedacted(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	if int32(1) <= l1 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
		if base.Ui32(l1) < base.Ui32(int32(32)) {
			return int32(base.Ui32(v7)>>(uint(l1)%32)) & int32(1)
		} else {
			return v7 & int32(1)
		}
	} else {
		return int32(0)
	}
}
func F_clientConnPostponeMask(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	v5 = base.B2i32(v3 != int32(0))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v8 != 0 {
		v9 = v5 | int32(2)
	} else {
		v9 = v5
	}
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	return v9 | base.B2i32(base.Ui32(v10) < base.Ui32(v11))
}
func F_clientGetredirCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v8&int32(4) == int32(0) {
		v17 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 != 0 {
				m.G0 = v6 + int32(128)
				return
			} else {
				v19 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v19)
				v22 = v6 | int32(1)
				v33 = int32(45)
				*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v33)
				v37 = int32(1)
				v46 = F_ull2string(m, v22+v37, int32(126), int64(1))
				mBase = m.M
				if v46 == int32(0) {
					v65 = int32(0)
				} else {
					v65 = v46 + v37
				}
				v69 = int32(2573)
				*(*uint16)(unsafe.Add(mBase, uint32(v65+v6+int32(1)))) = uint16(v69)
				F__addReplyToBufferOrList(m, l0, v6, v65+int32(3))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					m.G0 = v6 + int32(128)
					return
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+16))
		F_addReplyLongLong(m, l0, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v6 + int32(128)
			return
		}
	}
}
func F_clientHasPendingIO(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v4 != 0 {
		v8 = int32(1)
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
		v8 = base.B2i32(v5 != int32(0))
	}
	return v8
}
func F_clientIDCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_addReplyLongLong(m, l0, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_clientMatchesIpFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	v6 = F_getClientPeerId(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v15 = v6 + base.B2i32(v12 == int32(91))
	v16 = int32(0)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v20 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v37 = v16
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	if v6 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	return int32(0)
L5:
	;
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v37 = v36
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v37 = v33
	goto L5
L8:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v37 = v30
	goto L5
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v37 = v27
	goto L5
L10:
	;
	v37 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	return v95
L12:
	;
	if v82 != 0 {
		v95 = v16
		goto L11
	} else {
		goto L25
	}
L13:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v82 = int32(0)
	goto L12
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v82 = v70 - v75
	goto L12
L16:
	;
	v43 = v15
	v44 = l1
	v45 = v37
	v46 = v41
	goto L19
L17:
	;
	v70 = int32(0)
	v71 = l1
	goto L15
L18:
	;
	v70 = v67 & int32(255)
	v71 = v65
	goto L15
L19:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v46&int32(255) != v50 {
		v65 = v44
		v67 = v46
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v65 = v59
	v67 = int32(0)
	goto L18
L21:
	;
	if v50 == int32(0) {
		v65 = v44
		v67 = v46
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v55 = v45 + int32(-1)
	if v55 == int32(0) {
		v65 = v44
		v67 = v46
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v58 = int32(1)
	v59 = v44 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v60 != 0 {
		v43 = v43 + v58
		v44 = v59
		v45 = v55
		v46 = v60
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v83 = v15 + v37
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v87 = v83 + base.B2i32(v84 == int32(93))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != int32(58) {
		v95 = v16
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v95 = base.B2i32(v91 != int32(48))
	goto L11
}
func F_clientPauseCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v10 != int32(4) {
		v109 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v115 = F_getTimeoutFromObjectOrReply(m, l0, v111, v7+int32(8), int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L34
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v16 = int32(_a_F_clientPauseCommand_0)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a_F_clientPauseCommand_1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if v51-v53 != 0 {
		goto L4
	} else {
		goto L17
	}
L6:
	;
	v51 = F_tolower(m, v47)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v53 = F_tolower(m, v52)
	mBase = m.M
	goto L5
L7:
	;
	v21 = v15
	v22 = v16
	v23 = v19
	goto L10
L8:
	;
	v47 = int32(0)
	v48 = v16
	goto L6
L9:
	;
	v47 = v44 & int32(255)
	v48 = v43
	goto L6
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		v43 = v22
		v44 = v23
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v43 = v37
	v44 = int32(0)
	goto L9
L12:
	;
	v29 = v23 & int32(255)
	if v29 == v25 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = int32(1)
	v37 = v22 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v38 != 0 {
		v21 = v21 + v36
		v22 = v37
		v23 = v38
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v31 = F_tolower(m, v29)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v33 = F_tolower(m, v32)
	mBase = m.M
	if v31 == v33 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v43 = v22
	v44 = v35
	goto L9
L16:
	;
	goto L11
L17:
	;
	v109 = int32(1)
	goto L2
L18:
	;
	if v94-v96 == int32(0) {
		v109 = v2
		goto L2
	} else {
		goto L30
	}
L19:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L18
L20:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L23
L21:
	;
	v90 = int32(0)
	v91 = v59
	goto L19
L22:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L19
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v86 = v80
	v87 = int32(0)
	goto L22
L25:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L22
L29:
	;
	goto L24
L30:
	;
	F_addReplyErrorLength(m, l0, int32(_a_F_clientPauseCommand_2), int32(38))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return
L32:
	;
	F_afterErrorReply(m, l0, int32(_a_F_clientPauseCommand_2), int32(38), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	goto L1
L34:
	;
	if v115 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v117 = int32(_a_F_clientPauseCommand_3)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_clientPauseCommand[0]))
	if v121&int32(2) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = int32(30)
	goto L38
L37:
	;
	v124 = int32(29)
	goto L38
L38:
	;
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v126 = v124
	goto L41
L40:
	;
	v126 = int32(30)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clientPauseCommand[0])) = v126
	v129 = *(*int64)(unsafe.Add(mBase, _c_F_clientPauseCommand[1]))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	if v130 <= v129 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_updatePausedActions(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L31
	} else {
		goto L44
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_clientPauseCommand[1])) = v130
	goto L42
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_clientPauseCommand[2]))
	if v137 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_clientPauseCommand[3]))
	F_addReply(m, l0, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L31
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_clientPauseCommand[4])) = int32(1)
	goto L45
L47:
	;
	goto L1
}
func F_clientSetName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v112
L2:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v100 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v89 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v8 = F_objectGetVal(m, l1)
	mBase = m.M
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-1)))))
	switch v11 & int32(7) {
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
		goto L5
	}
L5:
	;
	v63 = F_objectGetVal(m, l1)
	mBase = m.M
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
	switch v66 & int32(7) {
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
		goto L3
	}
L6:
	;
	if v28 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-17))))
	v28 = v27
	goto L6
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-9))))
	v28 = v24
	goto L6
L9:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(-5)))))
	v28 = v21
	goto L6
L10:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-3)))))
	v28 = v18
	goto L6
L11:
	;
	v28 = int32(base.Ui32(v11) >> (uint(int32(3)) % 32))
	goto L6
L12:
	;
	v31 = F_objectGetVal(m, l1)
	mBase = m.M
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v38 = v32
	v39 = v31
	goto L15
L14:
	;
	if l2 == int32(0) {
		v112 = int32(-1)
		goto L1
	} else {
		goto L19
	}
L15:
	;
	if base.Ui32((v38+int32(-127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	if v46 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v38 = v46
	v39 = v39 + int32(1)
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a_F_clientSetName_0)
	return int32(-1)
L20:
	;
	if v83 != 0 {
		goto L2
	} else {
		goto L26
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
	v83 = v82
	goto L20
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
	v83 = v79
	goto L20
L23:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
	v83 = v76
	goto L20
L24:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
	v83 = v73
	goto L20
L25:
	;
	v83 = int32(base.Ui32(v66) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	goto L3
L27:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = v96
	return v96
L28:
	;
	F_decrRefCount(m, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+348)) = l1
	F_incrRefCount(m, l1)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	F_decrRefCount(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v112 = int32(0)
	goto L1
}
func F_clientSetUser(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v10 = l2 & int32(1)
	if l2 != 0 {
		v19 = v6&int32(-25165825) | v10<<(uint(int32(23))%32) | v10<<(uint(int32(24))%32)
	} else {
		v19 = v6 & int32(-8388609)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v19
	return
}
func F_clientSetinfoCommand(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = F_objectGetVal(m, v14)
	mBase = m.M
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v18 = F_objectGetVal(m, v17)
	mBase = m.M
	v19 = int32(_a_F_clientSetinfoCommand_0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
	F_addReplyErrorFormat(m, l0, int32(_a_F_clientSetinfoCommand_1), v11+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L38
	} else {
		goto L54
	}
L3:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v100 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v59 = int32(_a_F_clientSetinfoCommand_2)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if v54-v56 != 0 {
		goto L4
	} else {
		goto L17
	}
L6:
	;
	v54 = F_tolower(m, v50)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	goto L5
L7:
	;
	v24 = v15
	v25 = v19
	v26 = v22
	goto L10
L8:
	;
	v50 = int32(0)
	v51 = v19
	goto L6
L9:
	;
	v50 = v47 & int32(255)
	v51 = v46
	goto L6
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == int32(0) {
		v46 = v25
		v47 = v26
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v46 = v40
	v47 = int32(0)
	goto L9
L12:
	;
	v32 = v26 & int32(255)
	if v32 == v28 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = int32(1)
	v40 = v25 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v41 != 0 {
		v24 = v24 + v39
		v25 = v40
		v26 = v41
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v34 = F_tolower(m, v32)
	mBase = m.M
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = F_tolower(m, v35)
	mBase = m.M
	if v34 == v36 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v46 = v25
	v47 = v38
	goto L9
L16:
	;
	goto L11
L17:
	;
	v99 = int32(352)
	goto L3
L18:
	;
	if v94-v96 != 0 {
		goto L2
	} else {
		goto L30
	}
L19:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L18
L20:
	;
	v64 = v15
	v65 = v59
	v66 = v62
	goto L23
L21:
	;
	v90 = int32(0)
	v91 = v59
	goto L19
L22:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L19
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v86 = v80
	v87 = int32(0)
	goto L22
L25:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L22
L29:
	;
	goto L24
L30:
	;
	v99 = int32(356)
	goto L3
L31:
	;
	v134 = l0 + v99
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v109 = v100
	v110 = v18
	goto L34
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_addReplyErrorFormat(m, l0, int32(_a_F_clientSetinfoCommand_3), v11)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	if base.Ui32((v109+int32(-127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	if v117 == int32(0) {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v109 = v117
	v110 = v110 + int32(1)
	goto L34
L38:
	;
	return
L39:
	;
	goto L1
L40:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-1)))))
	switch v142 & int32(7) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	case 3:
		goto L47
	case 4:
		goto L46
	default:
		goto L44
	}
L41:
	;
	F_decrRefCount(m, v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_clientSetinfoCommand[0]))
	F_addReply(m, l0, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L38
	} else {
		goto L53
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = int32(0)
	goto L43
L45:
	;
	if v159 == int32(0) {
		goto L44
	} else {
		goto L51
	}
L46:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-17))))
	v159 = v158
	goto L45
L47:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(-9))))
	v159 = v155
	goto L45
L48:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+int32(-5)))))
	v159 = v152
	goto L45
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(-3)))))
	v159 = v149
	goto L45
L50:
	;
	v159 = int32(base.Ui32(v142) >> (uint(int32(3)) % 32))
	goto L45
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v17
	F_incrRefCount(m, v17)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L38
	} else {
		goto L52
	}
L52:
	;
	goto L43
L53:
	;
	goto L1
L54:
	;
	goto L1
}
func F_clientSupportStandAloneRedirect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_clientSupportStandAloneRedirect[0]))
	if v5 != 0 {
		v13 = v2
	} else {
		v6 = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_clientSupportStandAloneRedirect[1]))
		if v7 == v6 {
			v13 = v2
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
			v13 = v10 & int32(1)
		}
	}
	return v13
}
func F_clientUnblockCommand(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v11 != int32(4) {
		v101 = v10
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	F_addReplyErrorLength(m, l0, int32(_a_F_clientUnblockCommand_0), int32(48))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L31
	} else {
		goto L89
	}
L3:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v107 = F_getLongLongFromObjectOrReply(m, l0, v103, v8+int32(8), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v16 = F_objectGetVal(m, v15)
	mBase = m.M
	v17 = int32(_a_F_clientUnblockCommand_1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v52-v54 == int32(0) {
		v101 = v10
		goto L3
	} else {
		goto L17
	}
L6:
	;
	v52 = F_tolower(m, v48)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v54 = F_tolower(m, v53)
	mBase = m.M
	goto L5
L7:
	;
	v22 = v16
	v23 = v17
	v24 = v20
	goto L10
L8:
	;
	v48 = int32(0)
	v49 = v17
	goto L6
L9:
	;
	v48 = v45 & int32(255)
	v49 = v44
	goto L6
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v26 == int32(0) {
		v44 = v23
		v45 = v24
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v44 = v38
	v45 = int32(0)
	goto L9
L12:
	;
	v30 = v24 & int32(255)
	if v30 == v26 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(1)
	v38 = v23 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v39 != 0 {
		v22 = v22 + v37
		v23 = v38
		v24 = v39
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v32 = F_tolower(m, v30)
	mBase = m.M
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v34 = F_tolower(m, v33)
	mBase = m.M
	if v32 == v34 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v44 = v23
	v45 = v36
	goto L9
L16:
	;
	goto L11
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = F_objectGetVal(m, v59)
	mBase = m.M
	v61 = int32(_a_F_clientUnblockCommand_2)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v64 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v96-v98 != 0 {
		goto L2
	} else {
		goto L30
	}
L19:
	;
	v96 = F_tolower(m, v92)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v98 = F_tolower(m, v97)
	mBase = m.M
	goto L18
L20:
	;
	v66 = v60
	v67 = v61
	v68 = v64
	goto L23
L21:
	;
	v92 = int32(0)
	v93 = v61
	goto L19
L22:
	;
	v92 = v89 & int32(255)
	v93 = v88
	goto L19
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 == int32(0) {
		v88 = v67
		v89 = v68
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v88 = v82
	v89 = int32(0)
	goto L22
L25:
	;
	v74 = v68 & int32(255)
	if v74 == v70 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v81 = int32(1)
	v82 = v67 + v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v83 != 0 {
		v66 = v66 + v81
		v67 = v82
		v68 = v83
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v76 = F_tolower(m, v74)
	mBase = m.M
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v78 = F_tolower(m, v77)
	mBase = m.M
	if v76 == v78 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v88 = v67
	v89 = v80
	goto L22
L29:
	;
	goto L24
L30:
	;
	v101 = int32(0)
	goto L3
L31:
	;
	return
L32:
	;
	if v107 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	v110 = int64(56)
	v112 = int64(65280)
	v114 = int64(40)
	v117 = int64(16711680)
	v119 = int64(24)
	v121 = int64(4278190080)
	v123 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v109<<(uint(v110)%64) | v109&v112<<(uint(v114)%64) | (v109&v117<<(uint(v119)%64) | v109&v121<<(uint(v123)%64)) | (int64(base.Ui64(v109)>>(uint(v123)%64))&v121 | int64(base.Ui64(v109)>>(uint(v119)%64))&v117 | (int64(base.Ui64(v109)>>(uint(v114)%64))&v112 | int64(base.Ui64(v109)>>(uint(v110)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_clientUnblockCommand[0]))
	v151 = v8 + int32(24)
	v152 = int32(8)
	v154 = v8 + int32(20)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	goto L37
L34:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v350 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L35:
	;
	if v307 != v152 {
		goto L62
	} else {
		goto L63
	}
L36:
	;
	v298 = int32(0)
	v304 = v163
	v305 = v164
	v307 = v298
	v311 = v298
	goto L35
L37:
	;
	if base.Ui32(v164) < base.Ui32(int32(8)) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v175 = v163
	v176 = v164
	v178 = int32(0)
	goto L40
L39:
	;
	v304 = v288
	v305 = v289
	v307 = v291
	v311 = base.B2i32(v294 != int32(0))
	goto L35
L40:
	;
	v184 = int32(base.Ui32(v176) >> (uint(int32(3)) % 32))
	v185 = int32(4)
	v186 = v175 + v185
	if v176&v185 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v288 = v279
	v289 = v280
	v291 = v264
	v294 = v269
	goto L39
L42:
	;
	v269 = int32(0)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v186+v184+(v269-v184)&int32(3)+v257<<(uint(int32(2))%32))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if base.Ui32(v280) < base.Ui32(int32(8)) {
		v288 = v279
		v289 = v280
		v291 = v264
		v294 = v269
		goto L39
	} else {
		goto L60
	}
L43:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v178))))
	v235 = int32(0)
	goto L54
L44:
	;
	v191 = int32(0)
	if base.Ui32(v152) <= base.Ui32(v178) {
		v224 = v178
		v227 = v191
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v227 == v184 {
		v257 = v191
		v264 = v224
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v201 = v178
	v204 = v191
	goto L47
L47:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v204))))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v201))))
	if v207 != v209 {
		v224 = v201
		v227 = v204
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v224 = v212
	v227 = v214
	goto L45
L49:
	;
	v211 = int32(1)
	v212 = v201 + v211
	v214 = v204 + v211
	if base.Ui32(v184) <= base.Ui32(v214) {
		v224 = v212
		v227 = v214
		goto L45
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v212) < base.Ui32(v152) {
		v201 = v212
		v204 = v214
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v288 = v175
	v289 = v176
	v291 = v224
	v294 = v227
	goto L39
L53:
	;
	if v235 != v184 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v235))))
	if v248 == v232&int32(255) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v250 = int32(1)
	v252 = v235 + v250
	if v252 != v184 {
		v235 = v252
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v304 = v175
	v305 = v176
	v307 = v178
	v311 = v250
	goto L35
L58:
	;
	v257 = v235
	v264 = v178 + int32(1)
	goto L42
L59:
	;
	v288 = v175
	v289 = v176
	v291 = v178
	v294 = v184
	goto L39
L60:
	;
	if base.Ui32(v264) < base.Ui32(v152) {
		v175 = v279
		v176 = v280
		v178 = v264
		goto L40
	} else {
		goto L61
	}
L61:
	;
	goto L41
L62:
	;
	goto L34
L63:
	;
	if v305&int32(1) == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v319 = v305 & int32(4)
	if v311&base.B2i32(v319 != int32(0)) != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	if v154 == int32(0) {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	if v305&int32(2) != 0 {
		v345 = int32(0)
		goto L67
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v345
	goto L62
L68:
	;
	v329 = int32(3)
	v330 = int32(base.Ui32(v305) >> (uint(v329) % 32))
	if v319 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v340 = int32(4)
	goto L71
L70:
	;
	v340 = v330 << (uint(int32(2)) % 32)
	goto L71
L71:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v304+v330+(int32(0)-v330)&v329+v340+int32(4))))
	v345 = v344
	goto L67
L72:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_clientUnblockCommand[1]))
	F_addReply(m, l0, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L31
	} else {
		goto L88
	}
L73:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+200)))
	if v353&int32(16) == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v350)+116))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	if base.Ui32(int32(5)) < base.Ui32(v361) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v376 == int32(0) {
		goto L72
	} else {
		goto L81
	}
L76:
	;
	v376 = v374
	goto L75
L77:
	;
	v374 = int32(0)
	goto L76
L78:
	;
	v364 = int32(1)
	if v364<<(uint(v361)%32)&int32(54) != 0 {
		v374 = v364
		goto L76
	} else {
		goto L79
	}
L79:
	;
	if v361 != int32(3) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v371 = F_moduleBlockedClientMayTimeout(m, v350)
	mBase = m.M
	v376 = v371
	goto L75
L81:
	;
	if v101 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_clientUnblockCommand[2]))
	F_addReply(m, l0, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L31
	} else {
		goto L87
	}
L83:
	;
	F_unblockClientOnTimeout(m, v350)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L31
	} else {
		goto L86
	}
L84:
	;
	F_unblockClientOnError(m, v350, int32(_a_F_clientUnblockCommand_3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L31
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	goto L82
L87:
	;
	goto L1
L88:
	;
	goto L1
L89:
	;
	F_afterErrorReply(m, l0, int32(_a_F_clientUnblockCommand_0), int32(48), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L31
	} else {
		goto L90
	}
L90:
	;
	goto L1
}
func F_freeClientFilter(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_valkey_free(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	goto L1
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_sdsfree(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(0)
	goto L5
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v24 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_sdsfree(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	goto L8
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_sdsfree(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	goto L11
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v38 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_decrRefCount(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(0)
	goto L14
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	F_decrRefCount(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	goto L17
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v52 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_valkey_free(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L20
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v59 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_sdsfree(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	goto L23
L26:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v66 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	F_sdsfree(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(0)
	goto L26
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v73 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_sdsfree(m, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(0)
	goto L29
L32:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v80 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	F_decrRefCount(m, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(0)
	goto L32
L35:
	;
	return
L36:
	;
	F_decrRefCount(m, v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(0)
	goto L35
}
func F_freeClientModuleData(m *base.Module, l0 int32) {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_valkey_free(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
			F_valkey_free(m, v9)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(0)
				return
			}
		}
	}
}
func F_freeClientMultiState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 < int32(1) {
		v64 = v8
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_valkey_free(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L16
	}
L4:
	;
	v16 = v8
	v17 = int32(0)
	goto L5
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v26 = v23 + v17*int32(20)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v27 < int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v64 = v60
	goto L3
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	F_valkey_free(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L11
	} else {
		goto L14
	}
L8:
	;
	v33 = int32(0)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v33<<(uint(int32(2))%32))))
	F_decrRefCount(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	return
L12:
	;
	v45 = v33 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v45 < v46 {
		v33 = v45
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v59 = v17 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v59 < v61 {
		v16 = v60
		v17 = v59
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(0)
	F_unwatchAllKeys(m, l0)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v79 == v78 {
		v132 = v78
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_valkey_free(m, v132)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L11
	} else {
		goto L31
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)+48))
	if v82 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientMultiState[0]))
	if v85 <= v83 {
		v119 = v82
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v132 = v79
	goto L18
L22:
	;
	F_valkey_free(m, v119)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L30
	}
L23:
	;
	v89 = v79
	v91 = v83
	v92 = v85
	goto L24
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v97 = v91 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v97)))
	if v99 == int32(0) {
		v111 = v89
		v112 = v95
		v113 = v92
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v119 = v112
	goto L22
L26:
	;
	v115 = v91 + int32(1)
	if v115 < v113 {
		v89 = v111
		v91 = v115
		v92 = v113
		goto L24
	} else {
		goto L29
	}
L27:
	;
	F_hashtableRelease(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v105+v97))) = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientMultiState[0]))
	v111 = v104
	v112 = v105
	v113 = v110
	goto L26
L29:
	;
	goto L25
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+48)) = int32(0)
	v132 = v126
	goto L18
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	goto L1
}
func F_freeClientReplicationData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_freeClientReplicationData_0), int32(_a_F_freeClientReplicationData_1), int32(1327))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L79
	}
L2:
	;
	m.G0 = v7 + int32(16)
	return
L3:
	;
	F_freeReplicaReferencedReplBuffer(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v14&int32(2) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v247&int32(1) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[0]))
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != int32(8) {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 != int32(7) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[1]))
	if v26 != int32(1) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[2]))
	if v30 != int32(1) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[3]))
	v36 = v7 + int32(8)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v37
	goto L13
L13:
	;
	v42 = v7 + int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v44 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[4]))
	if int32(2) < v86 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44+base.B2i32(v47 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v53
	goto L16
L18:
	;
	v59 = v44
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v61 == l0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L14
L21:
	;
	v68 = v7 + int32(8)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v70 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+104))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 == int32(7) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v70 != 0 {
		v59 = v70
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70+base.B2i32(v73 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v79
	goto L25
L27:
	;
	goto L20
L28:
	;
	F_killRDBChild(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	F__serverLog(m, int32(2), int32(_a_F_freeClientReplicationData_2), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L8
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v120&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v104 == int32(-1) {
		v109 = v100
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	if v110 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v107 = F_close(m, v104)
	mBase = m.M
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v109 = v108
	goto L34
L36:
	;
	F_sdsfree(m, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	goto L32
L38:
	;
	v123 = int32(792)
	goto L40
L39:
	;
	v123 = int32(788)
	goto L40
L40:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+uint32(_c_F_freeClientReplicationData[5])))
	v126 = F_listSearchKey(m, v125, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v126 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_listDelNode(m, v125, v126)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v132&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[6]))
	if v160 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L45:
	;
	if v132&int32(2) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[3]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v153 != 0 {
		goto L44
	} else {
		goto L53
	}
L47:
	;
	if v132&int32(262144) != 0 {
		goto L44
	} else {
		goto L50
	}
L48:
	;
	if v132&int32(4) == int32(0) {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v145 == int32(0) {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L52
L52:
	;
	goto L44
L53:
	;
	v154 = int32(_a_F_freeClientReplicationData_3)
	v156 = *(*int64)(unsafe.Add(mBase, _c_F_freeClientReplicationData[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_freeClientReplicationData[8])) = v156
	goto L44
L54:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v235 != int32(9) {
		goto L6
	} else {
		goto L71
	}
L55:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[9]))
	if v164 == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[3]))
	v170 = v7 + int32(8)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v171
	goto L57
L57:
	;
	v175 = int32(0)
	v177 = v7 + int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v179 == v175 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[10])) = v227
	goto L54
L59:
	;
	if v179 == int32(0) {
		v227 = v175
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179+base.B2i32(v182 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v188
	goto L60
L62:
	;
	v194 = v179
	v195 = v175
	goto L63
L63:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+104))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v198 != int32(9) {
		v209 = v195
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v227 = v209
	goto L58
L65:
	;
	v211 = v7 + int32(8)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v213 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v201 = int32(_a_F_freeClientReplicationData_3)
	v202 = *(*int64)(unsafe.Add(mBase, _c_F_freeClientReplicationData[7]))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v197)+80))
	v206 = int64(*(*int32)(unsafe.Add(mBase, _c_F_freeClientReplicationData[9])))
	v209 = v195 + base.B2i32(v202-v203 <= v206)
	goto L65
L67:
	;
	if v213 != 0 {
		v194 = v213
		v195 = v209
		goto L63
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213+base.B2i32(v216 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v222
	goto L68
L70:
	;
	goto L64
L71:
	;
	F_moduleFireServerEvent(m, int64(6), int32(1), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	goto L6
L73:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+152))
	F_sdsfree(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	F_replicationHandlePrimaryDisconnection(m)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+192))
	F_sdsfree(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	F_valkey_free(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	goto L2
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_freeClientReplyValue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	if l0 == int32(0) {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_zfree_with_size(m, l0, v4+int32(16))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_getClientOutputBufferMemoryUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(1) != 0 {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
		v29 = v25*int32(28) + v28
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		if v30 == int64(-1) {
			v40 = v29
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
			v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
		}
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		return v41 + v40
	} else {
		if v4&int32(2) == int32(0) {
			if v4&int32(262144) != 0 {
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v17 == int32(0) {
				} else {
				}
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			v29 = v25*int32(28) + v28
			v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v30 == int64(-1) {
				v40 = v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
				v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
			}
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
			return v41 + v40
		} else {
			if v4&int32(4) == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+184))
				if v45 != 0 {
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_getClientOutputBufferMemoryUsage[0]))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
					v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v51)+24)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
					v63 = int32(44)
					return base.I32_wrap_i64(v52+v53-v56) + base.I32_wrap_i64(v59-v60)*v63 + v63
				} else {
					return int32(0)
				}
			} else {
				if v4&int32(262144) != 0 {
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v17 == int32(0) {
					} else {
					}
				}
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				v29 = v25*int32(28) + v28
				v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v30 == int64(-1) {
					v40 = v29
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
					v40 = v29 + base.I32_wrap_i64(v30) + v36*int32(28)
				}
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				return v41 + v40
			}
		}
	}
}
func F_getClientPubSubChannels(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
func F_getClientTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(int32(5)) < base.Ui32(l0) {
		v11 = int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_getClientTypeName[0])))
		v11 = v10
	}
	return v11
}
func F_initClientPubSubData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v3 != 0 {
		return
	} else {
		v5 = F_valkey_malloc(m, int32(32))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v5
			v9 = F_hashtableCreate(m, int32(_a_F_initClientPubSubData_0))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v9
				v14 = F_hashtableCreate(m, int32(_a_F_initClientPubSubData_0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v14
					v19 = F_hashtableCreate(m, int32(_a_F_initClientPubSubData_0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v19
						return
					}
				}
			}
		}
	}
}
func F_initClientReplicationData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2 != 0 {
		return
	} else {
		v4 = F_valkey_calloc(m, int32(200))
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v4
			return
		}
	}
}
func F_linkClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_linkClient[0]))
	v10 = F_listAddNodeTail(m, v9, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = int32(_a_F_linkClient_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_linkClient[0]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v17 = int64(56)
		v19 = int64(65280)
		v21 = int64(40)
		v24 = int64(16711680)
		v26 = int64(24)
		v28 = int64(4278190080)
		v30 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v16<<(uint(v17)%64) | v16&v19<<(uint(v21)%64) | (v16&v24<<(uint(v26)%64) | v16&v28<<(uint(v30)%64)) | (int64(base.Ui64(v16)>>(uint(v30)%64))&v28 | int64(base.Ui64(v16)>>(uint(v26)%64))&v24 | (int64(base.Ui64(v16)>>(uint(v21)%64))&v19 | int64(base.Ui64(v16)>>(uint(v17)%64))))
		v54 = *(*int32)(unsafe.Add(mBase, _c_F_linkClient[1]))
		v55 = int32(8)
		v59 = F_raxInsert(m, v54, v6+v55, v55, l0, int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_prepareClientForFutureWrites(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_prepareClientToWrite(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v7 = int32(0)
		} else {
			v7 = l0
		}
		return v7
	}
}
func F_processClientIOReadsDone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v8 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_processClientIOReadsDone_0), int32(_a_F_processClientIOReadsDone_1), int32(6548))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L37
	}
L2:
	;
	v11 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_processClientIOReadsDone[0]))
	if v13 == v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v22 & int32(-8388609)
	if v22&int32(-2147482624) != 0 {
		v92 = v11
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_processClientsCommandsBatch(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L3
L7:
	;
	return v92
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 != int32(2) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v46 == int32(0) {
		v53 = v29
		v54 = v28
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v35 = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	if v36 == v35 {
		v44 = v30
		v45 = v35
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v44 = v30
	v45 = int32(0)
	goto L9
L12:
	;
	v39 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = int32(3)
	goto L15
L14:
	;
	v43 = v39
	goto L15
L15:
	;
	v44 = v39
	v45 = v43
	goto L9
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v55 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	m.T0[v46].(func(*base.Module, int32, int32))(m, v28, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = v52
	v54 = v51
	goto L16
L19:
	;
	if v31 == int32(2) {
		v92 = int32(0)
		goto L7
	} else {
		goto L22
	}
L20:
	;
	m.T0[v55].(func(*base.Module, int32))(m, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v63 = F_handleReadResult(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L24
	}
L23:
	;
	v92 = v44
	goto L7
L24:
	;
	if v63 == int32(-1) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+289)))
	if v67&int32(128) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_beforeNextClient(m, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L36
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v75 < int32(1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v70 = F_handleParseResults(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	switch v70 + int32(2) {
	case 0:
		goto L26
	case 1:
		v92 = v44
		goto L7
	default:
		goto L27
	}
L30:
	;
	v82 = F_addCommandToBatchAndProcessIfFull(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v78 | int32(2)
	goto L30
L32:
	;
	if v82 != int32(-1) {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v86 = F_processPendingCommandAndInputBuffer(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v86 != 0 {
		goto L23
	} else {
		goto L35
	}
L35:
	;
	goto L26
L36:
	;
	goto L23
L37:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_processClientIOWriteDone(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v382 int32
	_ = v382
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_processClientIOWriteDone_0), int32(_a_F_processClientIOWriteDone_1), int32(3307))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L14
	} else {
		goto L88
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
	if v14 != int32(2) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v17)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v19&int32(4) != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v22 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v31 = int32(1)
	goto L6
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v26 != int32(2) {
		v31 = int32(0)
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v44 = F_postWriteToClient(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L16
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v35 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+112))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	m.T0[v38].(func(*base.Module, int32, int32))(m, v32, v31)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	goto L10
L16:
	;
	if v44 == int32(-1) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+116))
	if v56 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+120))
	v54 = base.B2i32(v51 == int32(0))
	goto L18
L20:
	;
	v54 = int32(1)
	goto L18
L21:
	;
	if v54 != 0 {
		v307 = l0
		goto L24
	} else {
		goto L25
	}
L22:
	;
	m.T0[v56].(func(*base.Module, int32))(m, v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v308 = F_clientHasPendingReplies(m, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L14
	} else {
		goto L66
	}
L25:
	;
	v61 = int64(56)
	v63 = int64(65280)
	v65 = int64(40)
	v68 = int64(16711680)
	v70 = int64(24)
	v72 = int64(4278190080)
	v74 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v55<<(uint(v61)%64) | v55&v63<<(uint(v65)%64) | (v55&v68<<(uint(v70)%64) | v55&v72<<(uint(v74)%64)) | (int64(base.Ui64(v55)>>(uint(v74)%64))&v72 | int64(base.Ui64(v55)>>(uint(v70)%64))&v68 | (int64(base.Ui64(v55)>>(uint(v65)%64))&v63 | int64(base.Ui64(v55)>>(uint(v61)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_processClientIOWriteDone[0]))
	v101 = int32(8)
	v102 = v9 + v101
	v105 = v9 + int32(4)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	goto L29
L26:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v301 == int32(0) {
		goto L2
	} else {
		goto L64
	}
L27:
	;
	if v258 != v101 {
		goto L54
	} else {
		goto L55
	}
L28:
	;
	v249 = int32(0)
	v255 = v114
	v256 = v115
	v258 = v249
	v262 = v249
	goto L27
L29:
	;
	if base.Ui32(v115) < base.Ui32(int32(8)) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v126 = v114
	v127 = v115
	v129 = int32(0)
	goto L32
L31:
	;
	v255 = v239
	v256 = v240
	v258 = v242
	v262 = base.B2i32(v245 != int32(0))
	goto L27
L32:
	;
	v135 = int32(base.Ui32(v127) >> (uint(int32(3)) % 32))
	v136 = int32(4)
	v137 = v126 + v136
	if v127&v136 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v239 = v230
	v240 = v231
	v242 = v215
	v245 = v220
	goto L31
L34:
	;
	v220 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v137+v135+(v220-v135)&int32(3)+v208<<(uint(int32(2))%32))))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if base.Ui32(v231) < base.Ui32(int32(8)) {
		v239 = v230
		v240 = v231
		v242 = v215
		v245 = v220
		goto L31
	} else {
		goto L52
	}
L35:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v129))))
	v186 = int32(0)
	goto L46
L36:
	;
	v142 = int32(0)
	if base.Ui32(v101) <= base.Ui32(v129) {
		v175 = v129
		v178 = v142
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v178 == v135 {
		v208 = v142
		v215 = v175
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v152 = v129
	v155 = v142
	goto L39
L39:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v155))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v152))))
	if v158 != v160 {
		v175 = v152
		v178 = v155
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v175 = v163
	v178 = v165
	goto L37
L41:
	;
	v162 = int32(1)
	v163 = v152 + v162
	v165 = v155 + v162
	if base.Ui32(v135) <= base.Ui32(v165) {
		v175 = v163
		v178 = v165
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(v163) < base.Ui32(v101) {
		v152 = v163
		v155 = v165
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v239 = v126
	v240 = v127
	v242 = v175
	v245 = v178
	goto L31
L45:
	;
	if v186 != v135 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v186))))
	if v199 == v183&int32(255) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v201 = int32(1)
	v203 = v186 + v201
	if v203 != v135 {
		v186 = v203
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v255 = v126
	v256 = v127
	v258 = v129
	v262 = v201
	goto L27
L50:
	;
	v208 = v186
	v215 = v129 + int32(1)
	goto L34
L51:
	;
	v239 = v126
	v240 = v127
	v242 = v129
	v245 = v135
	goto L31
L52:
	;
	if base.Ui32(v215) < base.Ui32(v101) {
		v126 = v230
		v127 = v231
		v129 = v215
		goto L32
	} else {
		goto L53
	}
L53:
	;
	goto L33
L54:
	;
	goto L26
L55:
	;
	if v256&int32(1) == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v270 = v256 & int32(4)
	if v262&base.B2i32(v270 != int32(0)) != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	if v105 == int32(0) {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	if v256&int32(2) != 0 {
		v296 = int32(0)
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v296
	goto L54
L60:
	;
	v280 = int32(3)
	v281 = int32(base.Ui32(v256) >> (uint(v280) % 32))
	if v270 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v291 = int32(4)
	goto L63
L62:
	;
	v291 = v281 << (uint(int32(2)) % 32)
	goto L63
L63:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v255+v281+(int32(0)-v281)&v280+v291+int32(4))))
	v296 = v295
	goto L59
L64:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	if v304 == int32(0) {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v307 = v301
	goto L24
L66:
	;
	if v308 == int32(0) {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+220)))
	if v312&int32(1) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v319 = F_trySendWriteToIOThreads(m, v307)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L14
	} else {
		goto L71
	}
L69:
	;
	F_installClientWriteHandler(m, v307)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	goto L2
L71:
	;
	if v319 == int32(0) {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v307)+200))
	if v323&int32(4194304) != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v307)+104))
	if v326 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v334 = int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v307)+216))
	if v335 == int32(0) {
		v342 = v334
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	switch v329 {
	case 0:
		goto L74
	default:
		goto L2
	case 9, 11:
		goto L76
	}
L76:
	;
	if v323&int32(1024) != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v332 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	if v342 == int32(0) {
		goto L2
	} else {
		goto L83
	}
L80:
	;
	goto L79
L81:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v338 != 0 {
		v342 = v334
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+156))
	v342 = base.B2i32(v339 != int32(13))
	goto L80
L83:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v307)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+200)) = v345 | int32(4194304)
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_processClientIOWriteDone[1]))
	v352 = v307 + int32(168)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v350)+20))
	if v355 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L2
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v350)+20)) = v355 + int32(1)
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = int32(0)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v352
	v364 = v362
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v352
	v357 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = v357
	v364 = v357
	goto L85
L88:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_putClientInPendingWriteQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(4194304) != 0 {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v7 == int32(0) {
			v15 = int32(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
			if v16 == int32(0) {
				v23 = v15
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				if v19 != 0 {
					v23 = v15
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
					v23 = base.B2i32(v20 != int32(13))
				}
			}
			if v23 == int32(0) {
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v26 | int32(4194304)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_putClientInPendingWriteQueue[0]))
				v33 = l0 + int32(168)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
				if v36 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v43))) = v33
					v45 = v43
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v33
					v38 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = v38
					v45 = v38
				}
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v45
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v36 + int32(1)
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			switch v10 {
			case 0:
				v15 = int32(1)
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v16 == int32(0) {
					v23 = v15
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					if v19 != 0 {
						v23 = v15
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
						v23 = base.B2i32(v20 != int32(13))
					}
				}
				if v23 == int32(0) {
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v26 | int32(4194304)
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_putClientInPendingWriteQueue[0]))
					v33 = l0 + int32(168)
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					if v36 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						*(*int32)(unsafe.Add(mBase, uint32(v43))) = v33
						v45 = v43
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v33
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = v38
						v45 = v38
					}
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v45
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v36 + int32(1)
				}
			default:
			case 9, 11:
				if v4&int32(1024) != 0 {
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					if v13 != 0 {
					} else {
						v15 = int32(1)
						v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v16 == int32(0) {
							v23 = v15
						} else {
							v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							if v19 != 0 {
								v23 = v15
							} else {
								v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+156))
								v23 = base.B2i32(v20 != int32(13))
							}
						}
						if v23 == int32(0) {
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v26 | int32(4194304)
							v31 = *(*int32)(unsafe.Add(mBase, _c_F_putClientInPendingWriteQueue[0]))
							v33 = l0 + int32(168)
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
							if v36 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v33
								v45 = v43
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v33
								v38 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v33))) = v38
								v45 = v38
							}
							*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v36 + int32(1)
						}
					}
				}
			}
		}
	}
	return
}
func F_resetClient(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	v5 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v7 == int32(0) {
		v18 = v5
		v19 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v22&int32(1) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	v12 = base.B2i32(v10 != int32(175))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+204))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	v18 = v12
	v19 = base.B2i32(v15 != int32(180))
	goto L1
L4:
	;
	v18 = v12
	v19 = int32(1)
	goto L1
L5:
	;
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = int32(-1)
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v140
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v144 & int32(1069547518)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v148 != 0 {
		goto L35
	} else {
		goto L36
	}
L6:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v73
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v77
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v83 == v73 {
		goto L5
	} else {
		goto L22
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v43 < int32(1) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if v21 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	if v21 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v27
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v31
	goto L5
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = F_tryOffloadFreeArgvToIOThreads(m, l0, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v39 != int32(-1) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_valkey_free(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L21
	}
L16:
	;
	v50 = int32(0)
	goto L17
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v50<<(uint(int32(2))%32))))
	F_decrRefCount(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	goto L15
L19:
	;
	v59 = v50 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 < v60 {
		v50 = v59
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L6
L22:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v86&int32(1) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v94 = F_tryOffloadFreeArgvToIOThreads(m, l0, v93, v83)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	goto L5
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	goto L5
L26:
	;
	if v94 != int32(-1) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v99 <= v98 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_valkey_free(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L34
	}
L29:
	;
	v105 = v98
	goto L30
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v105<<(uint(int32(2))%32))))
	F_decrRefCount(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	goto L28
L32:
	;
	v114 = v105 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v114 < v115 {
		v105 = v114
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L25
L35:
	;
	F__serverAssert(m, int32(_a_F_resetClient_0), int32(_a_F_resetClient_1), int32(3412))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L12
	} else {
		goto L52
	}
L36:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v149 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = int32(0)
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L40
	}
L38:
	;
	F_listRelease(m, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v18&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	v181 = v179 & int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v181)
	if v167&int32(67108864) != 0 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v164 = v159 & int32(-521)
	goto L44
L43:
	;
	v164 = v159
	goto L44
L44:
	;
	if v159&int32(8) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v167 = v159
	goto L47
L46:
	;
	v167 = v164
	goto L47
L47:
	;
	if base.B2i32(v167&int32(8) == int32(0))&v19 != int32(1) {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v175 & int32(-129)
	goto L41
L49:
	;
	v191 = v167&int32(-201326593) | int32(134217728)
	goto L51
L50:
	;
	v191 = v167 & int32(-134217729)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v191
	return
L52:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_updateClientMemUsageAndBucket(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	v2 = int32(0)
	v5 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[0]))
	if v7 == v2 {
		v38 = v5
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
		if v40 == int32(0) {
			if v38 != 0 {
				v97 = int32(0)
				return v97
			} else {
				F_updateClientMemoryUsage(m, l0)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
					v66 = int32(17)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
					v69 = v66 - base.I32_clz(v67)
					if base.Ui32(v66) < base.Ui32(v69) {
						v72 = v63
					} else {
						v72 = v69
					}
					v75 = v64 + v72<<(uint(int32(3))%32)
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
					v79 = int32(1)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
					if v75 == v80 {
						v97 = v79
						return v97
					} else {
						if v80 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
							v90 = F_listAddNodeTail(m, v89, l0)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
								v97 = v79
								return v97
							}
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
							F_listDelNode(m, v84, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v90 = F_listAddNodeTail(m, v89, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
									v97 = v79
									return v97
								}
							}
						}
					}
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
			*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v43 - v44
			if v38 == int32(0) {
				F_updateClientMemoryUsage(m, l0)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
					v66 = int32(17)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
					v69 = v66 - base.I32_clz(v67)
					if base.Ui32(v66) < base.Ui32(v69) {
						v72 = v63
					} else {
						v72 = v69
					}
					v75 = v64 + v72<<(uint(int32(3))%32)
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
					v79 = int32(1)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
					if v75 == v80 {
						v97 = v79
						return v97
					} else {
						if v80 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
							v90 = F_listAddNodeTail(m, v89, l0)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
								v97 = v79
								return v97
							}
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
							F_listDelNode(m, v84, v85)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v90 = F_listAddNodeTail(m, v89, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
									v97 = v79
									return v97
								}
							}
						}
					}
				}
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
				F_listDelNode(m, v49, v50)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = int64(0)
					return int32(0)
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		if v10&int32(268451840) != 0 {
			v38 = v5
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
			if v40 == int32(0) {
				if v38 != 0 {
					v97 = int32(0)
					return v97
				} else {
					F_updateClientMemoryUsage(m, l0)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
						v66 = int32(17)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
						v69 = v66 - base.I32_clz(v67)
						if base.Ui32(v66) < base.Ui32(v69) {
							v72 = v63
						} else {
							v72 = v69
						}
						v75 = v64 + v72<<(uint(int32(3))%32)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
						v79 = int32(1)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
						if v75 == v80 {
							v97 = v79
							return v97
						} else {
							if v80 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v90 = F_listAddNodeTail(m, v89, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
									v97 = v79
									return v97
								}
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
								F_listDelNode(m, v84, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v90 = F_listAddNodeTail(m, v89, l0)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
										v97 = v79
										return v97
									}
								}
							}
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
				*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v43 - v44
				if v38 == int32(0) {
					F_updateClientMemoryUsage(m, l0)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
						v66 = int32(17)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
						v69 = v66 - base.I32_clz(v67)
						if base.Ui32(v66) < base.Ui32(v69) {
							v72 = v63
						} else {
							v72 = v69
						}
						v75 = v64 + v72<<(uint(int32(3))%32)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
						v79 = int32(1)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
						if v75 == v80 {
							v97 = v79
							return v97
						} else {
							if v80 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
								v90 = F_listAddNodeTail(m, v89, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
									v97 = v79
									return v97
								}
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
								F_listDelNode(m, v84, v85)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v90 = F_listAddNodeTail(m, v89, l0)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
										v97 = v79
										return v97
									}
								}
							}
						}
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
					F_listDelNode(m, v49, v50)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = int64(0)
						return int32(0)
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v13 == int32(0) {
				F__serverAssert(m, int32(_a_F_updateClientMemUsageAndBucket_0), int32(_a_F_updateClientMemUsageAndBucket_1), int32(1085))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v16 = int32(1)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v17&v16 != 0 {
					v38 = v16
				} else {
					if v17&int32(2) == int32(0) {
						v28 = int32(0)
						if v17&int32(262144) != 0 {
							v38 = v28
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v31 == int32(0) {
								v38 = v28
							} else {
								v38 = int32(1)
							}
						}
					} else {
						if v17&int32(4) == int32(0) {
							v38 = v16
						} else {
							v28 = int32(0)
							if v17&int32(262144) != 0 {
								v38 = v28
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v31 == int32(0) {
									v38 = v28
								} else {
									v38 = int32(1)
								}
							}
						}
					}
				}
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
				if v40 == int32(0) {
					if v38 != 0 {
						v97 = int32(0)
						return v97
					} else {
						F_updateClientMemoryUsage(m, l0)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(0)
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
							v66 = int32(17)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
							v69 = v66 - base.I32_clz(v67)
							if base.Ui32(v66) < base.Ui32(v69) {
								v72 = v63
							} else {
								v72 = v69
							}
							v75 = v64 + v72<<(uint(int32(3))%32)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
							v79 = int32(1)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
							if v75 == v80 {
								v97 = v79
								return v97
							} else {
								if v80 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v90 = F_listAddNodeTail(m, v89, l0)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
										v97 = v79
										return v97
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
									F_listDelNode(m, v84, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v90 = F_listAddNodeTail(m, v89, l0)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
											v97 = v79
											return v97
										}
									}
								}
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v43 - v44
					if v38 == int32(0) {
						F_updateClientMemoryUsage(m, l0)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(0)
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_updateClientMemUsageAndBucket[1]))
							v66 = int32(17)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
							v69 = v66 - base.I32_clz(v67)
							if base.Ui32(v66) < base.Ui32(v69) {
								v72 = v63
							} else {
								v72 = v69
							}
							v75 = v64 + v72<<(uint(int32(3))%32)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 + v67
							v79 = int32(1)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
							if v75 == v80 {
								v97 = v79
								return v97
							} else {
								if v80 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									v90 = F_listAddNodeTail(m, v89, l0)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
										v97 = v79
										return v97
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
									F_listDelNode(m, v84, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v75
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
										v90 = F_listAddNodeTail(m, v89, l0)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v93
											v97 = v79
											return v97
										}
									}
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
						F_listDelNode(m, v49, v50)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = int64(0)
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_validateClientFlagFilter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
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
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v31 = int32(0)
wl1:
	for {
		switch v20 & int32(7) {
		case 0:
			v40 = int32(base.Ui32(v20) >> (uint(int32(3)) % 32))
		case 1:
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v40 = v36
		case 2:
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v40 = v37
		case 3:
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v40 = v38
		case 4:
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v40 = v39
		default:
			v40 = int32(0)
		}
		if base.Ui32(v40) <= base.Ui32(v31) {
			break
		} else {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v31))))
			switch v43 + int32(-65) {
			case 0, 1, 4, 8, 12, 13, 14, 15, 17, 18, 19, 20, 33, 34, 35, 36, 40, 49, 51, 52, 55:
				v31 = v31 + int32(1)
				continue
			default:
				break wl1
			}
			break
		}
		break
	}
	if base.Ui32(v31) < base.Ui32(v40) {
		v51 = int32(-1)
	} else {
		v51 = int32(0)
	}
	return v51
}
