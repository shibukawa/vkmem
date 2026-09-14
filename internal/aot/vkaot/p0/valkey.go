package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_valkeyAeAttach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+220))
	if v4 != 0 {
		return
	} else {
		v6 = F_valkey_malloc(m, int32(20))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+240)) = int32(1009)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = int32(1010)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+232)) = int32(1011)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+228)) = int32(1012)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+224)) = int32(1013)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+220)) = v6
			return
		}
	}
}
func F_valkeyAeCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v12 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_aeDeleteFileEvent(m, v17, v18, int32(2))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_aeDeleteFileEvent(m, v7, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v12 == int32(0) {
				F_valkey_free(m, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_aeDeleteFileEvent(m, v17, v18, int32(2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_valkeyAeReadEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_valkeyAsyncHandleRead(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_valkeyAeWriteEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_valkeyAsyncHandleWrite(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_valkeyAsyncHandleConnect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	v2 = int32(0)
	v7 = m.G0
	v8 = int32(16)
	v9 = v7 - v8
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v2
	v15 = v9 + int32(12)
	v20 = m.G0
	v22 = v20 - v8
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v27 = F_connect(m, v24, v25, v26)
	mBase = m.M
	if v27 != 0 {
		v31 = int32(26)
		v32 = F___errno_location(m)
		mBase = m.M
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		if v33 != v31 {
			v52 = v33
			switch v52 + int32(-6) {
			case 0, 1:
				v59 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v59
				v65 = v59
			default:
				v65 = int32(-1)
			case 24:
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
				v65 = int32(0)
			}
		} else {
			v36 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			v45 = F_getsockopt(m, v38, int32(1), v36, v22+int32(12), v22+int32(8))
			mBase = m.M
			if v45 != 0 {
				v50 = v31
				v52 = v50
				switch v52 + int32(-6) {
				case 0, 1:
					v59 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v59
					v65 = v59
				default:
					v65 = int32(-1)
				case 24:
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
					v65 = int32(0)
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				if v46 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
					v65 = int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v46
					v50 = v46
					v52 = v50
					switch v52 + int32(-6) {
					case 0, 1:
						v59 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v59
						v65 = v59
					default:
						v65 = int32(-1)
					case 24:
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
						v65 = int32(0)
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
		v65 = int32(0)
	}
	m.G0 = v22 + int32(16)
	if v65 != int32(-1) {
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		if v87 != int32(1) {
			v164 = v2
			m.G0 = v9 + int32(16)
			return v164
		} else {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
			if v90 != 0 {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				v97 = v95 | int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v97
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
				if v99 == int32(0) {
					v118 = v97
					if v118&int32(4) == int32(0) {
						if v118&int32(8) == int32(0) {
							v164 = v2
							m.G0 = v9 + int32(16)
							return v164
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
							v153 = int32(-1)
							if v118&int32(16) != 0 {
								v164 = v153
								m.G0 = v9 + int32(16)
								return v164
							} else {
								v156 = v153
								F_valkeyAsyncFreeInternal(m, l0)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									v164 = v156
									m.G0 = v9 + int32(16)
									return v164
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
						v126 = int32(-1)
						if v118&int32(16) != 0 {
							v164 = v126
							m.G0 = v9 + int32(16)
							return v164
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
							if v129 != 0 {
								v164 = v126
								m.G0 = v9 + int32(16)
								return v164
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
								v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
								if v135 == int32(0) {
									v156 = v126
									F_valkeyAsyncFreeInternal(m, l0)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										v164 = v156
										m.G0 = v9 + int32(16)
										return v164
									}
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
									m.T0[v135].(func(*base.Module, int32))(m, v138)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										v141 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
										v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
										if v143&int32(512) == v141 {
											v156 = v126
											F_valkeyAsyncFreeInternal(m, l0)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v164 = v156
												m.G0 = v9 + int32(16)
												return v164
											}
										} else {
											v164 = v126
											m.G0 = v9 + int32(16)
											return v164
										}
									}
								}
							}
						}
					}
				} else {
					if v95&int32(16) != 0 {
						m.T0[v99].(func(*base.Module, int32, int32))(m, l0, int32(0))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							v118 = v117
							if v118&int32(4) == int32(0) {
								if v118&int32(8) == int32(0) {
									v164 = v2
									m.G0 = v9 + int32(16)
									return v164
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
									v153 = int32(-1)
									if v118&int32(16) != 0 {
										v164 = v153
										m.G0 = v9 + int32(16)
										return v164
									} else {
										v156 = v153
										F_valkeyAsyncFreeInternal(m, l0)
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											v164 = v156
											m.G0 = v9 + int32(16)
											return v164
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
								v126 = int32(-1)
								if v118&int32(16) != 0 {
									v164 = v126
									m.G0 = v9 + int32(16)
									return v164
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
									if v129 != 0 {
										v164 = v126
										m.G0 = v9 + int32(16)
										return v164
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
										if v135 == int32(0) {
											v156 = v126
											F_valkeyAsyncFreeInternal(m, l0)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v164 = v156
												m.G0 = v9 + int32(16)
												return v164
											}
										} else {
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
											m.T0[v135].(func(*base.Module, int32))(m, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v141 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
												v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
												if v143&int32(512) == v141 {
													v156 = v126
													F_valkeyAsyncFreeInternal(m, l0)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v164 = v156
														m.G0 = v9 + int32(16)
														return v164
													}
												} else {
													v164 = v126
													m.G0 = v9 + int32(16)
													return v164
												}
											}
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v95 | int32(18)
						m.T0[v99].(func(*base.Module, int32, int32))(m, l0, int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							v112 = v110 & int32(-17)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v112
							v118 = v112
							if v118&int32(4) == int32(0) {
								if v118&int32(8) == int32(0) {
									v164 = v2
									m.G0 = v9 + int32(16)
									return v164
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
									v153 = int32(-1)
									if v118&int32(16) != 0 {
										v164 = v153
										m.G0 = v9 + int32(16)
										return v164
									} else {
										v156 = v153
										F_valkeyAsyncFreeInternal(m, l0)
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											v164 = v156
											m.G0 = v9 + int32(16)
											return v164
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
								v126 = int32(-1)
								if v118&int32(16) != 0 {
									v164 = v126
									m.G0 = v9 + int32(16)
									return v164
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
									if v129 != 0 {
										v164 = v126
										m.G0 = v9 + int32(16)
										return v164
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
										if v135 == int32(0) {
											v156 = v126
											F_valkeyAsyncFreeInternal(m, l0)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v164 = v156
												m.G0 = v9 + int32(16)
												return v164
											}
										} else {
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
											m.T0[v135].(func(*base.Module, int32))(m, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v141 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
												v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
												if v143&int32(512) == v141 {
													v156 = v126
													F_valkeyAsyncFreeInternal(m, l0)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v164 = v156
														m.G0 = v9 + int32(16)
														return v164
													}
												} else {
													v164 = v126
													m.G0 = v9 + int32(16)
													return v164
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
				v91 = F_valkeySetTcpNoDelay(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					if v91 == int32(-1) {
						F_valkeyAsyncHandleConnectFailure(m, l0)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							v164 = int32(-1)
							m.G0 = v9 + int32(16)
							return v164
						}
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
						v97 = v95 | int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v97
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
						if v99 == int32(0) {
							v118 = v97
							if v118&int32(4) == int32(0) {
								if v118&int32(8) == int32(0) {
									v164 = v2
									m.G0 = v9 + int32(16)
									return v164
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
									v153 = int32(-1)
									if v118&int32(16) != 0 {
										v164 = v153
										m.G0 = v9 + int32(16)
										return v164
									} else {
										v156 = v153
										F_valkeyAsyncFreeInternal(m, l0)
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											v164 = v156
											m.G0 = v9 + int32(16)
											return v164
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
								v126 = int32(-1)
								if v118&int32(16) != 0 {
									v164 = v126
									m.G0 = v9 + int32(16)
									return v164
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
									if v129 != 0 {
										v164 = v126
										m.G0 = v9 + int32(16)
										return v164
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
										v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
										if v135 == int32(0) {
											v156 = v126
											F_valkeyAsyncFreeInternal(m, l0)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v164 = v156
												m.G0 = v9 + int32(16)
												return v164
											}
										} else {
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
											m.T0[v135].(func(*base.Module, int32))(m, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v141 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
												v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
												if v143&int32(512) == v141 {
													v156 = v126
													F_valkeyAsyncFreeInternal(m, l0)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v164 = v156
														m.G0 = v9 + int32(16)
														return v164
													}
												} else {
													v164 = v126
													m.G0 = v9 + int32(16)
													return v164
												}
											}
										}
									}
								}
							}
						} else {
							if v95&int32(16) != 0 {
								m.T0[v99].(func(*base.Module, int32, int32))(m, l0, int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
									v118 = v117
									if v118&int32(4) == int32(0) {
										if v118&int32(8) == int32(0) {
											v164 = v2
											m.G0 = v9 + int32(16)
											return v164
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
											v153 = int32(-1)
											if v118&int32(16) != 0 {
												v164 = v153
												m.G0 = v9 + int32(16)
												return v164
											} else {
												v156 = v153
												F_valkeyAsyncFreeInternal(m, l0)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v164 = v156
													m.G0 = v9 + int32(16)
													return v164
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
										v126 = int32(-1)
										if v118&int32(16) != 0 {
											v164 = v126
											m.G0 = v9 + int32(16)
											return v164
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
											if v129 != 0 {
												v164 = v126
												m.G0 = v9 + int32(16)
												return v164
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
												v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
												if v135 == int32(0) {
													v156 = v126
													F_valkeyAsyncFreeInternal(m, l0)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v164 = v156
														m.G0 = v9 + int32(16)
														return v164
													}
												} else {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
													m.T0[v135].(func(*base.Module, int32))(m, v138)
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														v141 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
														v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
														if v143&int32(512) == v141 {
															v156 = v126
															F_valkeyAsyncFreeInternal(m, l0)
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return int32(0)
															} else {
																v164 = v156
																m.G0 = v9 + int32(16)
																return v164
															}
														} else {
															v164 = v126
															m.G0 = v9 + int32(16)
															return v164
														}
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v95 | int32(18)
								m.T0[v99].(func(*base.Module, int32, int32))(m, l0, int32(0))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
									v112 = v110 & int32(-17)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v112
									v118 = v112
									if v118&int32(4) == int32(0) {
										if v118&int32(8) == int32(0) {
											v164 = v2
											m.G0 = v9 + int32(16)
											return v164
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118
											v153 = int32(-1)
											if v118&int32(16) != 0 {
												v164 = v153
												m.G0 = v9 + int32(16)
												return v164
											} else {
												v156 = v153
												F_valkeyAsyncFreeInternal(m, l0)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v164 = v156
													m.G0 = v9 + int32(16)
													return v164
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v118 & int32(-513)
										v126 = int32(-1)
										if v118&int32(16) != 0 {
											v164 = v126
											m.G0 = v9 + int32(16)
											return v164
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
											if v129 != 0 {
												v164 = v126
												m.G0 = v9 + int32(16)
												return v164
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v133
												v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
												if v135 == int32(0) {
													v156 = v126
													F_valkeyAsyncFreeInternal(m, l0)
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v164 = v156
														m.G0 = v9 + int32(16)
														return v164
													}
												} else {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
													m.T0[v135].(func(*base.Module, int32))(m, v138)
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														v141 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v141
														v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
														if v143&int32(512) == v141 {
															v156 = v126
															F_valkeyAsyncFreeInternal(m, l0)
															mBase = m.M
															v159 = m.ExcPending
															if v159 != 0 {
																return int32(0)
															} else {
																v164 = v156
																m.G0 = v9 + int32(16)
																return v164
															}
														} else {
															v164 = v126
															m.G0 = v9 + int32(16)
															return v164
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
		v74 = F_valkeyCheckSocketError(m, l0)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			if l0 == int32(0) {
			} else {
				if v74 != int32(-1) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v85
				}
			}
			F_valkeyAsyncHandleConnectFailure(m, l0)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return int32(0)
			} else {
				v164 = int32(-1)
				m.G0 = v9 + int32(16)
				return v164
			}
		}
	}
}
func F_valkeyAsyncHandleRead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v3&int32(16) != 0 {
		v19 = m.G3
		m.Env.X__assert_fail(m, v19+int32(_a_F_valkeyAsyncHandleRead_0), v19+int32(_a_F_valkeyAsyncHandleRead_1), int32(725), v19+int32(_a_F_valkeyAsyncHandleRead_2))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		if v3&int32(2) != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
			m.T0[v16].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		} else {
			v8 = F_valkeyAsyncHandleConnect(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				if v8 != 0 {
					return
				} else {
					v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
					if v10&int32(2) == int32(0) {
						return
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
						m.T0[v16].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
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
func F_valkeyAsyncHandleWrite(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v3&int32(16) != 0 {
		v19 = m.G3
		m.Env.X__assert_fail(m, v19+int32(_a_F_valkeyAsyncHandleWrite_0), v19+int32(_a_F_valkeyAsyncHandleWrite_1), int32(760), v19+int32(_a_F_valkeyAsyncHandleWrite_2))
		mBase = m.M
		base.Wasm_trap_unreachable()
		for {
		}
	} else {
		if v3&int32(2) != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			m.T0[v16].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		} else {
			v8 = F_valkeyAsyncHandleConnect(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				if v8 != 0 {
					return
				} else {
					v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
					if v10&int32(2) == int32(0) {
						return
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						m.T0[v16].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
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
func F_valkeyAsyncWrite(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
	v14 = F_valkeyBufferWrite(m, l0, v8+int32(44))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return
L2:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	return
L4:
	;
	if v14 != int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if l0 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v23
	v25 = v23
	goto L6
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, 204))
	v25 = v19
	goto L6
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v48 | int32(4)
	goto L9
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v26 != v31 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = m.G4
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	m.T0[v36].(func(*base.Module, int32))(m, v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = int32(0)
	goto L13
L15:
	;
	v39 = m.G3
	m.Env.X__assert_fail(m, v39+int32(_a_F_valkeyAsyncWrite_0), v39+int32(_a_F_valkeyAsyncWrite_1), int32(423), v39+int32(_a_F_valkeyAsyncWrite_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = int32(0)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
	if v61&int32(2) != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v53].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_valkeyAsyncFreeInternal(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L1
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v127&int32(2) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v117 == int32(0) {
		goto L21
	} else {
		goto L39
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v68&int32(2) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v111 == int32(0) {
		goto L21
	} else {
		goto L37
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v93+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v93)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v103
	m.T0[v67].(func(*base.Module, int32, int32))(m, v94, v8+int32(24))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L36
	}
L26:
	;
	if v67 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L27:
	;
	if v67 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v75 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	if v78 != int64(0) {
		v93 = v75
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	if v81 != 0 {
		v93 = v75
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v84 == int32(0) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	if v87 != int64(0) {
		v93 = v84
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v90 == int32(0) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v93 = v84
	goto L25
L36:
	;
	goto L24
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v111].(func(*base.Module, int32))(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	goto L21
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v117].(func(*base.Module, int32))(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	goto L21
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v170 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v154 = int32(8)
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v152+v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v162
	m.T0[v126].(func(*base.Module, int32, int32))(m, v153, v8+v154)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L53
	}
L43:
	;
	if v126 == int32(0) {
		goto L41
	} else {
		goto L49
	}
L44:
	;
	if v126 == int32(0) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v134 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
	if v137 != int64(0) {
		v152 = v134
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	if v140 != 0 {
		v152 = v134
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v143 == int32(0) {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
	if v146 != int64(0) {
		v152 = v143
		goto L42
	} else {
		goto L51
	}
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v149 == int32(0) {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v152 = v143
	goto L42
L53:
	;
	goto L41
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	m.T0[v170].(func(*base.Module, int32))(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L1
}
func F_valkeyCheckConnectDone(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v14 = F_connect(m, v11, v12, v13)
	mBase = m.M
	if v14 != 0 {
		v19 = int32(9116376)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_valkeyCheckConnectDone[0]))
		if v20 != int32(26) {
			v39 = v20
			switch v39 + int32(-6) {
			case 0, 1:
				v46 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
				v52 = v46
			default:
				v52 = int32(-1)
			case 24:
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
				v52 = int32(0)
			}
		} else {
			v23 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			v32 = F_getsockopt(m, v25, int32(1), v23, v9+int32(12), v9+int32(8))
			mBase = m.M
			if v32 != 0 {
				v37 = int32(26)
				v39 = v37
				switch v39 + int32(-6) {
				case 0, 1:
					v46 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
					v52 = v46
				default:
					v52 = int32(-1)
				case 24:
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
					v52 = int32(0)
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				if v33 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
					v52 = int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_valkeyCheckConnectDone[0])) = v33
					v37 = v33
					v39 = v37
					switch v39 + int32(-6) {
					case 0, 1:
						v46 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v46
						v52 = v46
					default:
						v52 = int32(-1)
					case 24:
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
						v52 = int32(0)
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
		v52 = int32(0)
	}
	m.G0 = v9 + int32(16)
	return v52
}
func F_valkeyCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v100 int32
	_ = v100
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v125 = m.G3
	v131 = m.G8
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	m.T0[v132].(func(*base.Module, int32, int32, int32))(m, v125+int32(_a_F_valkeyCommandHandler_0), v125+int32(_a_F_valkeyCommandHandler_1), int32(622))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L29
	}
L2:
	;
	v13 = int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v17 = F_lua_checkstack(m, v14, l1+v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v116 = m.G14
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	m.T0[v117].(func(*base.Module))(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L28
	}
L4:
	;
	v29 = m.G3
	F_lua_getfield(m, v14, int32(-10002), v29+int32(_a_F_valkeyCommandHandler_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v21 = m.G3
	v24 = m.G19
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	m.T0[v25].(func(*base.Module, int32))(m, v21+int32(_a_F_valkeyCommandHandler_3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	F_lua_pushstring(m, v14, v29+int32(_a_F_valkeyCommandHandler_4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	F_lua_gettable(m, v14, int32(-2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if l1 < int32(2) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = m.G6
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+264)) = v72
	v78 = F_lua_pcall(m, v14, l1+int32(-1), v72, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L19
	}
L13:
	;
	v47 = v13
	goto L14
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47<<(uint(int32(2))%32))))
	v55 = m.G7
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = m.T0[v56].(func(*base.Module, int32, int32) int32)(m, v52, v9+int32(12))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	F_lua_pushlstring(m, v14, v57, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v63 = v47 + int32(1)
	if v63 != l1 {
		v47 = v63
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+264)) = int32(0)
	goto L22
L20:
	;
	goto L3
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v100 + int32(-32)
	goto L20
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	goto L21
L28:
	;
	m.G0 = v9 + int32(16)
	return int32(1)
L29:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_valkeyConnectWithOptions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v185 int64
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int64
	_ = v225
	var v231 int64
	_ = v231
	var v237 int32
	_ = v237
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(3) < v11 {
		v237 = v2
		m.G0 = v8 + int32(16)
		return v237
	} else {
		v16 = m.G4
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v18 = m.T0[v17].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(204))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v237 = v2
				m.G0 = v8 + int32(16)
				return v237
			} else {
				v24 = F_sdsempty(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v24
					v27 = m.G3
					v30 = F_valkeyReaderCreateWithFunctions(m, v27+int32(_a_F_valkeyConnectWithOptions_0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+148)) = v30
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+144))
						if v35 == int32(0) {
							F_valkeyFree(m, v18)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v237 = v2
								m.G0 = v8 + int32(16)
								return v237
							}
						} else {
							if v30 != 0 {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v40&int32(1) != 0 {
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v43 | int32(1)
								}
								if v40&int32(2) == int32(0) {
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v51 | int32(128)
								}
								if v40&int32(4) == int32(0) {
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v59 | int32(512)
								}
								if v40&int32(16) == int32(0) {
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v67 | int32(1024)
								}
								if v40&int32(32) == int32(0) {
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v75 | int32(2048)
								}
								if v40&int32(64) == int32(0) {
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v83 | int32(4096)
								}
								if v40&int32(128) == int32(0) {
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
									*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v125 | int32(8192)
								}
								v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v129 != 0 {
									v136 = v129
									*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v136
								} else {
									v130 = m.G5
									v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
									if v131&int32(8) != 0 {
									} else {
										v136 = v130 + int32(1150)
										*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v136
									}
								}
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v139
								v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v141
								v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v18)+152)) = v143
								F_valkeyContextSetFuncs(m, v18)
								mBase = m.M
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v18)+156))
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v146 == v147 {
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v167 == v168 {
										v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
										v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
										v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
											if v192 != 0 {
												v237 = v18
												m.G0 = v8 + int32(16)
												return v237
											} else {
												v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												if v193 == int32(-1) {
													v237 = v18
													m.G0 = v8 + int32(16)
													return v237
												} else {
													v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													if v196 == int32(0) {
														v237 = v18
														m.G0 = v8 + int32(16)
														return v237
													} else {
														v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
														if v199&int32(1) == int32(0) {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
															v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
															v206 = int32(8)
															v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
															*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
															v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
															*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
															v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															}
														}
													}
												}
											}
										}
									} else {
										if v167 != 0 {
											v178 = v167
											v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
											*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
											v181 = int32(8)
											v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
											*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
											v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
											v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												if v192 != 0 {
													v237 = v18
													m.G0 = v8 + int32(16)
													return v237
												} else {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
													if v193 == int32(-1) {
														v237 = v18
														m.G0 = v8 + int32(16)
														return v237
													} else {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														if v196 == int32(0) {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
															if v199&int32(1) == int32(0) {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																v206 = int32(8)
																v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																}
															}
														}
													}
												}
											}
										} else {
											v171 = m.G4
											v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
											v173 = m.T0[v172].(func(*base.Module, int32) int32)(m, int32(16))
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v173
												if v173 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(5)
													v220 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)) = uint8(v220)
													v222 = m.G3
													v225 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[0])))
													*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v225
													v231 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[1])))
													*(*int64)(unsafe.Add(mBase, uint32(v18+int32(13)))) = v231
													v237 = v18
													m.G0 = v8 + int32(16)
													return v237
												} else {
													v178 = v173
													v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
													*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
													v181 = int32(8)
													v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
													*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
													v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
													v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
													v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int32(0)
													} else {
														v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
														if v192 != 0 {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
															if v193 == int32(-1) {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																if v196 == int32(0) {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																} else {
																	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																	if v199&int32(1) == int32(0) {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	} else {
																		v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																		v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																		v206 = int32(8)
																		v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																		v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																		v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			v237 = v18
																			m.G0 = v8 + int32(16)
																			return v237
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
									if v146 != 0 {
										v157 = v146
										v158 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
										*(*int64)(unsafe.Add(mBase, uint32(v157))) = v158
										v160 = int32(8)
										v164 = *(*int64)(unsafe.Add(mBase, uint32(v147+v160)))
										*(*int64)(unsafe.Add(mBase, uint32(v157+v160))) = v164
										v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v167 == v168 {
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
											v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
											v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												if v192 != 0 {
													v237 = v18
													m.G0 = v8 + int32(16)
													return v237
												} else {
													v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
													if v193 == int32(-1) {
														v237 = v18
														m.G0 = v8 + int32(16)
														return v237
													} else {
														v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														if v196 == int32(0) {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
															if v199&int32(1) == int32(0) {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																v206 = int32(8)
																v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																}
															}
														}
													}
												}
											}
										} else {
											if v167 != 0 {
												v178 = v167
												v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
												*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
												v181 = int32(8)
												v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
												*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
												v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
												v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
												v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return int32(0)
												} else {
													v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
													if v192 != 0 {
														v237 = v18
														m.G0 = v8 + int32(16)
														return v237
													} else {
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
														if v193 == int32(-1) {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															if v196 == int32(0) {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																if v199&int32(1) == int32(0) {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																} else {
																	v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																	v206 = int32(8)
																	v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																	*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																	v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																	v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return int32(0)
																	} else {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	}
																}
															}
														}
													}
												}
											} else {
												v171 = m.G4
												v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
												v173 = m.T0[v172].(func(*base.Module, int32) int32)(m, int32(16))
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v173
													if v173 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(5)
														v220 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)) = uint8(v220)
														v222 = m.G3
														v225 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[0])))
														*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v225
														v231 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[1])))
														*(*int64)(unsafe.Add(mBase, uint32(v18+int32(13)))) = v231
														v237 = v18
														m.G0 = v8 + int32(16)
														return v237
													} else {
														v178 = v173
														v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
														*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
														v181 = int32(8)
														v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
														*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
														v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
														v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
															if v192 != 0 {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
																if v193 == int32(-1) {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																} else {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	if v196 == int32(0) {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	} else {
																		v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																		if v199&int32(1) == int32(0) {
																			v237 = v18
																			m.G0 = v8 + int32(16)
																			return v237
																		} else {
																			v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																			v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																			v206 = int32(8)
																			v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																			v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																			v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return int32(0)
																			} else {
																				v237 = v18
																				m.G0 = v8 + int32(16)
																				return v237
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
										v150 = m.G4
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
										v152 = m.T0[v151].(func(*base.Module, int32) int32)(m, int32(16))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v18)+156)) = v152
											if v152 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(5)
												v220 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)) = uint8(v220)
												v222 = m.G3
												v225 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[0])))
												*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v225
												v231 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[1])))
												*(*int64)(unsafe.Add(mBase, uint32(v18+int32(13)))) = v231
												v237 = v18
												m.G0 = v8 + int32(16)
												return v237
											} else {
												v157 = v152
												v158 = *(*int64)(unsafe.Add(mBase, uint32(v147)))
												*(*int64)(unsafe.Add(mBase, uint32(v157))) = v158
												v160 = int32(8)
												v164 = *(*int64)(unsafe.Add(mBase, uint32(v147+v160)))
												*(*int64)(unsafe.Add(mBase, uint32(v157+v160))) = v164
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+160))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v167 == v168 {
													v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
													v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
													v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int32(0)
													} else {
														v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
														if v192 != 0 {
															v237 = v18
															m.G0 = v8 + int32(16)
															return v237
														} else {
															v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
															if v193 == int32(-1) {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																if v196 == int32(0) {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																} else {
																	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																	if v199&int32(1) == int32(0) {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	} else {
																		v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																		v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																		v206 = int32(8)
																		v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																		v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																		v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			v237 = v18
																			m.G0 = v8 + int32(16)
																			return v237
																		}
																	}
																}
															}
														}
													}
												} else {
													if v167 != 0 {
														v178 = v167
														v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
														*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
														v181 = int32(8)
														v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
														*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
														v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
														v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
															if v192 != 0 {
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
																if v193 == int32(-1) {
																	v237 = v18
																	m.G0 = v8 + int32(16)
																	return v237
																} else {
																	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	if v196 == int32(0) {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	} else {
																		v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																		if v199&int32(1) == int32(0) {
																			v237 = v18
																			m.G0 = v8 + int32(16)
																			return v237
																		} else {
																			v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																			v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																			v206 = int32(8)
																			v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																			v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																			v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return int32(0)
																			} else {
																				v237 = v18
																				m.G0 = v8 + int32(16)
																				return v237
																			}
																		}
																	}
																}
															}
														}
													} else {
														v171 = m.G4
														v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
														v173 = m.T0[v172].(func(*base.Module, int32) int32)(m, int32(16))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v173
															if v173 == int32(0) {
																*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(5)
																v220 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)) = uint8(v220)
																v222 = m.G3
																v225 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[0])))
																*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v225
																v231 = *(*int64)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_valkeyConnectWithOptions[1])))
																*(*int64)(unsafe.Add(mBase, uint32(v18+int32(13)))) = v231
																v237 = v18
																m.G0 = v8 + int32(16)
																return v237
															} else {
																v178 = v173
																v179 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
																*(*int64)(unsafe.Add(mBase, uint32(v178))) = v179
																v181 = int32(8)
																v185 = *(*int64)(unsafe.Add(mBase, uint32(v168+v181)))
																*(*int64)(unsafe.Add(mBase, uint32(v178+v181))) = v185
																v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
																v190 = m.T0[v189].(func(*base.Module, int32, int32) int32)(m, v18, l0)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																	if v192 != 0 {
																		v237 = v18
																		m.G0 = v8 + int32(16)
																		return v237
																	} else {
																		v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
																		if v193 == int32(-1) {
																			v237 = v18
																			m.G0 = v8 + int32(16)
																			return v237
																		} else {
																			v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																			if v196 == int32(0) {
																				v237 = v18
																				m.G0 = v8 + int32(16)
																				return v237
																			} else {
																				v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)))
																				if v199&int32(1) == int32(0) {
																					v237 = v18
																					m.G0 = v8 + int32(16)
																					return v237
																				} else {
																					v204 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
																					v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
																					v206 = int32(8)
																					v210 = *(*int64)(unsafe.Add(mBase, uint32(v196+v206)))
																					*(*int64)(unsafe.Add(mBase, uint32(v8+v206))) = v210
																					v212 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
																					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v212
																					v214 = m.T0[v205].(func(*base.Module, int32, int32) int32)(m, v18, v8)
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						v237 = v18
																						m.G0 = v8 + int32(16)
																						return v237
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
								F_valkeyFree(m, v18)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v237 = v2
									m.G0 = v8 + int32(16)
									return v237
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_valkeyContextConnectUserfd(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v3
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v5 | int32(2)
	return int32(0)
}
func F_valkeyContextRegisterUserfdFuncs(m *base.Module) {
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	v1 = m.G3
	v5 = F_valkeyContextRegisterFuncs(m, v1+int32(_a_F_valkeyContextRegisterUserfdFuncs_0), int32(2))
	return
}
func F_valkeyGetSubscribeCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
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
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v253 int64
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int64
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int64
	_ = v416
	var v421 int32
	_ = v421
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int64
	_ = v438
	var v442 int64
	_ = v442
	var v446 int64
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int64
	_ = v481
	var v483 int32
	_ = v483
	var v487 int64
	_ = v487
	var v489 int32
	_ = v489
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v16 + int32(-2) {
	case 0:
		goto L8
	default:
		goto L6
	case 10:
		goto L7
	}
L1:
	;
	v543 = m.G3
	m.Env.X__assert_fail(m, v543+int32(_a_F_valkeyGetSubscribeCallback_0), v543+int32(_a_F_valkeyGetSubscribeCallback_1), int32(511), v543+int32(_a_F_valkeyGetSubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	v534 = m.G3
	m.Env.X__assert_fail(m, v534+int32(_a_F_valkeyGetSubscribeCallback_3), v534+int32(_a_F_valkeyGetSubscribeCallback_1), int32(500), v534+int32(_a_F_valkeyGetSubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	v525 = m.G3
	m.Env.X__assert_fail(m, v525+int32(_a_F_valkeyGetSubscribeCallback_4), v525+int32(_a_F_valkeyGetSubscribeCallback_1), int32(477), v525+int32(_a_F_valkeyGetSubscribeCallback_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	m.G0 = v14 + int32(32)
	return
L5:
	;
	v500 = m.G3
	F_valkeySetError(m, l0, int32(5), v500+int32(_a_F_valkeyGetSubscribeCallback_5))
	mBase = m.M
	if l0 == int32(0) {
		goto L4
	} else {
		goto L139
	}
L6:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v472 == int32(0) {
		goto L4
	} else {
		goto L135
	}
L7:
	;
	v25 = int32(1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 != v25 {
		goto L3
	} else {
		goto L11
	}
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+141)))
	if v19&int32(1) != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if base.Ui32(v22) < base.Ui32(int32(3)) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
	if base.Ui32(v32+int32(-65)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v41 = base.B2i32(v39 == int32(112))
	v42 = m.G3
	goto L18
L13:
	;
	v39 = v32 | int32(32)
	goto L15
L14:
	;
	v39 = v32
	goto L15
L15:
	;
	goto L12
L16:
	;
	if v216 != 0 {
		goto L64
	} else {
		goto L65
	}
L17:
	;
	if v88-v90 == int32(0) {
		v216 = v25
		goto L16
	} else {
		goto L32
	}
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_valkeyGetSubscribeCallback[0]))))
	if v49 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v88 = F_tolower(m, v83)
	mBase = m.M
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v90 = F_tolower(m, v89)
	mBase = m.M
	goto L17
L21:
	;
	v51 = v42 + int32(_a_F_valkeyGetSubscribeCallback_6)
	v52 = v31
	v53 = int32(2)
	v54 = v49
	goto L24
L22:
	;
	v83 = int32(0)
	v84 = v31
	goto L20
L23:
	;
	v83 = v80 & int32(255)
	v84 = v78
	goto L20
L24:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == int32(0) {
		v78 = v52
		v80 = v54
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v78 = v72
	v80 = int32(0)
	goto L23
L26:
	;
	v60 = v53 + int32(-1)
	if v60 == int32(0) {
		v78 = v52
		v80 = v54
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v64 = v54 & int32(255)
	if v64 == v56 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v71 = int32(1)
	v72 = v52 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v73 != 0 {
		v51 = v51 + v71
		v52 = v72
		v53 = v60
		v54 = v73
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v66 = F_tolower(m, v64)
	mBase = m.M
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v68 = F_tolower(m, v67)
	mBase = m.M
	if v66 == v68 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v78 = v52
	v80 = v70
	goto L23
L31:
	;
	goto L25
L32:
	;
	v100 = m.G3
	goto L34
L33:
	;
	if v146-v148 == int32(0) {
		v216 = v25
		goto L16
	} else {
		goto L48
	}
L34:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_valkeyGetSubscribeCallback[1]))))
	if v107 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v146 = F_tolower(m, v141)
	mBase = m.M
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v148 = F_tolower(m, v147)
	mBase = m.M
	goto L33
L37:
	;
	v109 = v100 + int32(_a_F_valkeyGetSubscribeCallback_7)
	v110 = v31
	v111 = int32(2)
	v112 = v107
	goto L40
L38:
	;
	v141 = int32(0)
	v142 = v31
	goto L36
L39:
	;
	v141 = v138 & int32(255)
	v142 = v136
	goto L36
L40:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 == int32(0) {
		v136 = v110
		v138 = v112
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v136 = v130
	v138 = int32(0)
	goto L39
L42:
	;
	v118 = v111 + int32(-1)
	if v118 == int32(0) {
		v136 = v110
		v138 = v112
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v122 = v112 & int32(255)
	if v122 == v114 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v129 = int32(1)
	v130 = v110 + v129
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v131 != 0 {
		v109 = v109 + v129
		v110 = v130
		v111 = v118
		v112 = v131
		goto L40
	} else {
		goto L47
	}
L45:
	;
	v124 = F_tolower(m, v122)
	mBase = m.M
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v126 = F_tolower(m, v125)
	mBase = m.M
	if v124 == v126 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v136 = v110
	v138 = v128
	goto L39
L47:
	;
	goto L41
L48:
	;
	v158 = m.G3
	goto L50
L49:
	;
	v216 = base.B2i32(v204-v206 == int32(0))
	goto L16
L50:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+uint32(_c_F_valkeyGetSubscribeCallback[2]))))
	if v165 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v204 = F_tolower(m, v199)
	mBase = m.M
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v206 = F_tolower(m, v205)
	mBase = m.M
	goto L49
L53:
	;
	v167 = v158 + int32(_a_F_valkeyGetSubscribeCallback_8)
	v168 = v31
	v169 = int32(3)
	v170 = v165
	goto L56
L54:
	;
	v199 = int32(0)
	v200 = v31
	goto L52
L55:
	;
	v199 = v196 & int32(255)
	v200 = v194
	goto L52
L56:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v172 == int32(0) {
		v194 = v168
		v196 = v170
		goto L55
	} else {
		goto L58
	}
L57:
	;
	v194 = v188
	v196 = int32(0)
	goto L55
L58:
	;
	v176 = v169 + int32(-1)
	if v176 == int32(0) {
		v194 = v168
		v196 = v170
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v180 = v170 & int32(255)
	if v180 == v172 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v187 = int32(1)
	v188 = v168 + v187
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v189 != 0 {
		v167 = v167 + v187
		v168 = v188
		v169 = v176
		v170 = v189
		goto L56
	} else {
		goto L63
	}
L61:
	;
	v182 = F_tolower(m, v180)
	mBase = m.M
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v184 = F_tolower(m, v183)
	mBase = m.M
	if v182 == v184 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v194 = v168
	v196 = v186
	goto L55
L63:
	;
	goto L57
L64:
	;
	v220 = int32(288)
	goto L66
L65:
	;
	v220 = int32(280)
	goto L66
L66:
	;
	if v39 == int32(112) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v221 = int32(284)
	goto L69
L68:
	;
	v221 = v220
	goto L69
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0+v221)))
	v224 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v226 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v259 = v31 + v41 + v216
	v260 = m.G3
	v262 = v260 + int32(_a_F_valkeyGetSubscribeCallback_9)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v265 != 0 {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+28))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225)+24))
	v232 = F_sdsnewlen(m, v230, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v255 = int32(0)
	v256 = v224
	goto L70
L73:
	;
	return
L74:
	;
	if v232 == int32(0) {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v236 = F_dictFind(m, v223, v232)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	if v236 == int32(0) {
		v255 = v232
		v256 = v224
		goto L70
	} else {
		goto L77
	}
L77:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236)+8))
	goto L78
L78:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v240)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v241
	v243 = int32(16)
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v240+v243)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v243))) = v247
	v249 = int32(8)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v240+v249)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v249))) = v253
	v255 = v232
	v256 = v240
	goto L70
L79:
	;
	v311 = m.G3
	v313 = v311 + int32(_a_F_valkeyGetSubscribeCallback_10)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v316 != 0 {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	if v297-v299 != 0 {
		goto L79
	} else {
		goto L92
	}
L81:
	;
	v297 = F_tolower(m, v293)
	mBase = m.M
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v299 = F_tolower(m, v298)
	mBase = m.M
	goto L80
L82:
	;
	v267 = v259
	v268 = v262
	v269 = v265
	goto L85
L83:
	;
	v293 = int32(0)
	v294 = v262
	goto L81
L84:
	;
	v293 = v290 & int32(255)
	v294 = v289
	goto L81
L85:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v271 == int32(0) {
		v289 = v268
		v290 = v269
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v289 = v283
	v290 = int32(0)
	goto L84
L87:
	;
	v275 = v269 & int32(255)
	if v275 == v271 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v282 = int32(1)
	v283 = v268 + v282
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	if v284 != 0 {
		v267 = v267 + v282
		v268 = v283
		v269 = v284
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v277 = F_tolower(m, v275)
	mBase = m.M
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v279 = F_tolower(m, v278)
	mBase = m.M
	if v277 == v279 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v289 = v268
	v290 = v281
	goto L84
L91:
	;
	goto L86
L92:
	;
	if v256 == int32(0) {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+20)) = int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v305 + int32(-1)
	F_sdsfree(m, v255)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L73
	} else {
		goto L94
	}
L94:
	;
	goto L4
L95:
	;
	F_sdsfree(m, v255)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L73
	} else {
		goto L134
	}
L96:
	;
	if v348-v350 != 0 {
		goto L95
	} else {
		goto L108
	}
L97:
	;
	v348 = F_tolower(m, v344)
	mBase = m.M
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v350 = F_tolower(m, v349)
	mBase = m.M
	goto L96
L98:
	;
	v318 = v259
	v319 = v313
	v320 = v316
	goto L101
L99:
	;
	v344 = int32(0)
	v345 = v313
	goto L97
L100:
	;
	v344 = v341 & int32(255)
	v345 = v340
	goto L97
L101:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v322 == int32(0) {
		v340 = v319
		v341 = v320
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v340 = v334
	v341 = int32(0)
	goto L100
L103:
	;
	v326 = v320 & int32(255)
	if v326 == v322 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v333 = int32(1)
	v334 = v319 + v333
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v335 != 0 {
		v318 = v318 + v333
		v319 = v334
		v320 = v335
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v328 = F_tolower(m, v326)
	mBase = m.M
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	v330 = F_tolower(m, v329)
	mBase = m.M
	if v328 == v330 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v340 = v319
	v341 = v332
	goto L100
L107:
	;
	goto L102
L108:
	;
	if v256 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	if v361 != int32(3) {
		goto L1
	} else {
		goto L114
	}
L110:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if v356 != 0 {
		goto L109
	} else {
		goto L112
	}
L111:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+292)) = v352 + int32(-1)
	goto L109
L112:
	;
	v357 = F_dictDelete(m, v223, v255)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L73
	} else {
		goto L113
	}
L113:
	;
	goto L109
L114:
	;
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v360)+8))
	if v364 != int64(0) {
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	if v368 != int32(0)-v370 {
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if v374 != int32(0)-v376 {
		goto L95
	} else {
		goto L117
	}
L117:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+16))
	if v380 != int32(0)-v382 {
		goto L95
	} else {
		goto L118
	}
L118:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v385 != 0 {
		goto L95
	} else {
		goto L119
	}
L119:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v386 & int32(-33)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v390 == int32(0) {
		goto L95
	} else {
		goto L120
	}
L120:
	;
	v394 = v390
	goto L121
L121:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v394 != v406 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L95
L123:
	;
	v412 = int32(16)
	v413 = v14 + int32(24)
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v394+v412)))
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = v416
	v421 = v14 + v412
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v394+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v426
	v428 = m.G4
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	m.T0[v429].(func(*base.Module, int32))(m, v394)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L73
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = int32(0)
	goto L123
L125:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	v434 = m.T0[v433].(func(*base.Module, int32) int32)(m, int32(24))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L73
	} else {
		goto L127
	}
L126:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v458 != 0 {
		v394 = v458
		goto L121
	} else {
		goto L133
	}
L127:
	;
	if v434 == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = v438
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v413)))
	*(*int64)(unsafe.Add(mBase, uint32(v434+int32(16)))) = v442
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v421)))
	*(*int64)(unsafe.Add(mBase, uint32(v434+int32(8)))) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v434))) = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v450 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v452 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v434
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v434
	goto L126
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v434
	goto L131
L133:
	;
	goto L122
L134:
	;
	goto L4
L135:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v472 != v477 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v472)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v481
	v483 = int32(16)
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v472+v483)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v483))) = v487
	v489 = int32(8)
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v472+v489)))
	*(*int64)(unsafe.Add(mBase, uint32(l2+v489))) = v493
	v495 = m.G4
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+16))
	m.T0[v496].(func(*base.Module, int32))(m, v472)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L73
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = int32(0)
	goto L136
L138:
	;
	goto L4
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = l0 + int32(8)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v509
	goto L4
}
func F_valkeyHasMptcp(m *base.Module) int32 {
	return int32(1)
}
func F_valkeyNetRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v5 = int32(0)
	v8 = F_recvfrom(m, v4, l1, l2, v5, v5, v5)
	mBase = m.M
	switch v8 + int32(1) {
	case 0:
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_valkeyNetRead[0]))
		if v12 == int32(73) {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
			if v27&int32(1) == int32(0) {
				v44 = F___strerror_l(m, v12, v12)
				mBase = m.M
				F_valkeySetError(m, l0, int32(1), v44)
				mBase = m.M
			} else {
				v33 = m.G3
				F_valkeySetError(m, l0, int32(6), v33+int32(_a_F_valkeyNetRead_0))
				mBase = m.M
			}
			v50 = int32(-1)
			return v50
		} else {
			if v12 != int32(27) {
				if v12 != int32(6) {
					v44 = F___strerror_l(m, v12, v12)
					mBase = m.M
					F_valkeySetError(m, l0, int32(1), v44)
					mBase = m.M
					v50 = int32(-1)
				} else {
					v21 = int32(0)
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
					if v22&int32(1) == v21 {
						v50 = v21
					} else {
						v44 = F___strerror_l(m, v12, v12)
						mBase = m.M
						F_valkeySetError(m, l0, int32(1), v44)
						mBase = m.M
						v50 = int32(-1)
					}
				}
				return v50
			} else {
				return int32(0)
			}
		}
	case 1:
		v38 = m.G3
		F_valkeySetError(m, l0, int32(3), v38+int32(_a_F_valkeyNetRead_1))
		mBase = m.M
		v50 = int32(-1)
		return v50
	default:
		v50 = v8
		return v50
	}
}
func F_valkeyPushAutoFree(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_freeReplyObject(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_valkeyReaderCreateWithFunctions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v7 = m.G4
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(184))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	return int32(0)
L3:
	;
	if v9 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v15 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v15
	if v15 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v22 = m.G4
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, int32(9), int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v24
	if v24 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+164))
	if int32(8) < v29 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = int32(16384)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = l0
	return v9
L11:
	;
	goto L12
L12:
	;
	v38 = m.G4
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = m.T0[v39].(func(*base.Module, int32, int32) int32)(m, int32(1), int32(32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L14
	}
L13:
	;
	goto L10
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+160))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v42+v43<<(uint(int32(2))%32)))) = v40
	if v40 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v43 + int32(1)
	if v43 < int32(8) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+160))
	if v83 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+176))
	if v74 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v77 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	m.T0[v77].(func(*base.Module, int32))(m, v71)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+132))
	F_sdsfree(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+164))
	if v87 <= v86 {
		v111 = v83
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v112 = m.G4
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	m.T0[v113].(func(*base.Module, int32))(m, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L30
	}
L25:
	;
	v92 = v86
	goto L26
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+160))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v92<<(uint(int32(2))%32))))
	v99 = m.G4
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	m.T0[v100].(func(*base.Module, int32))(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L28
	}
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+160))
	v111 = v107
	goto L24
L28:
	;
	v104 = v92 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v9)+164))
	if v104 < v105 {
		v92 = v104
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L22
L31:
	;
	v123 = m.G4
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	m.T0[v124].(func(*base.Module, int32))(m, v9)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L1
}
func F_valkeySetError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v8 = l0 + int32(8)
	if l2&int32(3) == int32(0) {
		v30 = l2
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v64 = int32(127)
	if base.Ui32(v63) < base.Ui32(v64) {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v63 = v55 - l2
	goto L3
L5:
	;
	v34 = v30
	goto L13
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = l2
	goto L9
L8:
	;
	v63 = l2 - l2
	goto L3
L9:
	;
	v23 = v19 + int32(1)
	if v23&int32(3) == int32(0) {
		v30 = v23
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v28 != 0 {
		v19 = v23
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v55 = v23
	goto L4
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v43 = int32(-2139062144)
	if (int32(16843008)-v40|v40)&v43 == v43 {
		v34 = v34 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v49 = v34
	goto L16
L15:
	;
	goto L14
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 != 0 {
		v49 = v49 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v55 = v49
	goto L4
L18:
	;
	goto L17
L19:
	;
	v67 = v63
	goto L21
L20:
	;
	v67 = v64
	goto L21
L21:
	;
	if v67 == int32(0) {
		v71 = v8
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v67))) = uint8(v73)
	return
L23:
	;
	goto L22
L24:
	;
	v70 = F__emscripten_memcpy_bulkmem(m, v8, l2, v67)
	mBase = m.M
	v71 = v70
	goto L23
L25:
	;
	v104 = m.G3
	m.Env.X__assert_fail(m, v104+int32(_a_F_valkeySetError_0), v104+int32(_a_F_valkeySetError_1), int32(707), v104+int32(_a_F_valkeySetError_2))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_valkeySetError[0]))
	v80 = l0 + int32(8)
	v83 = F_strerror(m, v78)
	mBase = m.M
	v84 = F_strlen(m, v83)
	mBase = m.M
	if base.Ui32(v84) < base.Ui32(int32(128)) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	return
L29:
	;
	goto L28
L30:
	;
	v98 = F___memcpy(m, v80, v83, v84+int32(1))
	mBase = m.M
	goto L29
L31:
	;
	goto L32
L32:
	;
	v91 = F___memcpy(m, v80, v83, int32(127))
	mBase = m.M
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(135)))) = uint8(v93)
	goto L28
}
func F_valkey_realloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = F_ztryrealloc_usable_internal(m, l0, l1, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if l1 == int32(0) {
			return v4
		} else {
			if v4 != 0 {
				return v4
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, _c_F_valkey_realloc[0]))
				m.T0[v11].(func(*base.Module, int32))(m, l1)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					return v4
				}
			}
		}
	}
}
func F_valkey_strtod_sds(m *base.Module, l0 int32, l1 int32) float64 {
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
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 float64
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(9116376)
	v13 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_valkey_strtod_sds[0])) = v13
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v18 & int32(7) {
	case 0:
		v35 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
	case 1:
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v35 = v25
	case 2:
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v35 = v28
	case 3:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v35 = v31
	case 4:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v35 = v34
	default:
		v35 = v13
	}
	v38 = int32(0)
	v39 = *(*int64)(unsafe.Add(mBase, _c_F_valkey_strtod_sds[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_valkey_strtod_sds[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
	F_ffc_from_chars_double_options(m, v10+int32(16), l0, l0+v35, v10+int32(24), v10)
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v52 == v38 {
	} else {
		if v52 == int32(2) {
			v59 = int32(68)
		} else {
			v59 = int32(28)
		}
		*(*int32)(unsafe.Add(mBase, _c_F_valkey_strtod_sds[0])) = v59
	}
	if l1 == int32(0) {
	} else {
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v63
	}
	v65 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
	m.G0 = v10 + int32(32)
	return v65
}
