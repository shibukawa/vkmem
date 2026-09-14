package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_IOThreadsBeforeSleep(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v76 int64
	_ = v76
	var v84 int64
	_ = v84
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	v5 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v5 == int32(1) {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[346]))
		if v9 != 0 {
			F__serverAssert(m, int32(_a671), int32(_a672), int32(118))
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[347]))
			if v11 < int32(2) {
				v68 = v11
				if v68 != int32(1) {
				} else {
					v76 = *(*int64)(unsafe.Add(mBase, _consts[348]))
					if l0-v76 < int64(50000) {
					} else {
						*(*int64)(unsafe.Add(mBase, _consts[348])) = l0
						v84 = *(*int64)(unsafe.Add(mBase, _consts[349]))
						v91 = int32(_a673)
						v92 = *(*int64)(unsafe.Add(mBase, _consts[350]))
						if v92 < int64(1) {
						} else {
							v95 = l0 - v92
							if int64(1) <= v95 {
								v99 = *(*int64)(unsafe.Add(mBase, _consts[351]))
								v102 = base.I64_div_s((v84-v99)*int64(1000000), v95)
								v103 = v102
							} else {
								v103 = int64(0)
							}
							v104 = *(*int32)(unsafe.Add(mBase, _consts[352]))
							*(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[353]))) = v103
							v114 = base.I32_rem_s(v104+int32(1), int32(16))
							*(*int32)(unsafe.Add(mBase, _consts[352])) = v114
						}
						*(*int64)(unsafe.Add(mBase, _consts[351])) = v84
						*(*int64)(unsafe.Add(mBase, _consts[350])) = l0
					}
				}
			} else {
				v17 = int32(1)
				for {
					v19 = v17 * int32(192)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[354])))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[355])))
					if v23 == v24 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[355]))) = v23
					}
					v28 = v17 + int32(1)
					v30 = *(*int32)(unsafe.Add(mBase, _consts[347]))
					if v28 < v30 {
						v17 = v28
						continue
					} else {
						break
					}
					break
				}
				v33 = *(*int32)(unsafe.Add(mBase, _consts[356]))
				if v33 == int32(0) {
					v68 = v30
					if v68 != int32(1) {
					} else {
						v76 = *(*int64)(unsafe.Add(mBase, _consts[348]))
						if l0-v76 < int64(50000) {
						} else {
							*(*int64)(unsafe.Add(mBase, _consts[348])) = l0
							v84 = *(*int64)(unsafe.Add(mBase, _consts[349]))
							v91 = int32(_a673)
							v92 = *(*int64)(unsafe.Add(mBase, _consts[350]))
							if v92 < int64(1) {
							} else {
								v95 = l0 - v92
								if int64(1) <= v95 {
									v99 = *(*int64)(unsafe.Add(mBase, _consts[351]))
									v102 = base.I64_div_s((v84-v99)*int64(1000000), v95)
									v103 = v102
								} else {
									v103 = int64(0)
								}
								v104 = *(*int32)(unsafe.Add(mBase, _consts[352]))
								*(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[353]))) = v103
								v114 = base.I32_rem_s(v104+int32(1), int32(16))
								*(*int32)(unsafe.Add(mBase, _consts[352])) = v114
							}
							*(*int64)(unsafe.Add(mBase, _consts[351])) = v84
							*(*int64)(unsafe.Add(mBase, _consts[350])) = l0
						}
					}
				} else {
					if v30 < int32(2) {
						v68 = v30
						if v68 != int32(1) {
						} else {
							v76 = *(*int64)(unsafe.Add(mBase, _consts[348]))
							if l0-v76 < int64(50000) {
							} else {
								*(*int64)(unsafe.Add(mBase, _consts[348])) = l0
								v84 = *(*int64)(unsafe.Add(mBase, _consts[349]))
								v91 = int32(_a673)
								v92 = *(*int64)(unsafe.Add(mBase, _consts[350]))
								if v92 < int64(1) {
								} else {
									v95 = l0 - v92
									if int64(1) <= v95 {
										v99 = *(*int64)(unsafe.Add(mBase, _consts[351]))
										v102 = base.I64_div_s((v84-v99)*int64(1000000), v95)
										v103 = v102
									} else {
										v103 = int64(0)
									}
									v104 = *(*int32)(unsafe.Add(mBase, _consts[352]))
									*(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[353]))) = v103
									v114 = base.I32_rem_s(v104+int32(1), int32(16))
									*(*int32)(unsafe.Add(mBase, _consts[352])) = v114
								}
								*(*int64)(unsafe.Add(mBase, _consts[351])) = v84
								*(*int64)(unsafe.Add(mBase, _consts[350])) = l0
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[347]))
						v40 = int32(0)
						v41 = *(*int32)(unsafe.Add(mBase, _consts[357]))
						v43 = *(*int32)(unsafe.Add(mBase, _consts[358]))
						if v41 != v43 {
							v68 = v39
							if v68 != int32(1) {
							} else {
								v76 = *(*int64)(unsafe.Add(mBase, _consts[348]))
								if l0-v76 < int64(50000) {
								} else {
									*(*int64)(unsafe.Add(mBase, _consts[348])) = l0
									v84 = *(*int64)(unsafe.Add(mBase, _consts[349]))
									v91 = int32(_a673)
									v92 = *(*int64)(unsafe.Add(mBase, _consts[350]))
									if v92 < int64(1) {
									} else {
										v95 = l0 - v92
										if int64(1) <= v95 {
											v99 = *(*int64)(unsafe.Add(mBase, _consts[351]))
											v102 = base.I64_div_s((v84-v99)*int64(1000000), v95)
											v103 = v102
										} else {
											v103 = int64(0)
										}
										v104 = *(*int32)(unsafe.Add(mBase, _consts[352]))
										*(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[353]))) = v103
										v114 = base.I32_rem_s(v104+int32(1), int32(16))
										*(*int32)(unsafe.Add(mBase, _consts[352])) = v114
									}
									*(*int64)(unsafe.Add(mBase, _consts[351])) = v84
									*(*int64)(unsafe.Add(mBase, _consts[350])) = l0
								}
							}
						} else {
							v45 = int32(1)
							if v39 <= v45 {
							} else {
								v50 = v45
								for {
									v57 = v50 + int32(1)
									v59 = *(*int32)(unsafe.Add(mBase, _consts[347]))
									if v57 < v59 {
										v50 = v57
										continue
									} else {
										break
									}
									break
								}
							}
							*(*int32)(unsafe.Add(mBase, _consts[347])) = int32(1)
							v76 = *(*int64)(unsafe.Add(mBase, _consts[348]))
							if l0-v76 < int64(50000) {
							} else {
								*(*int64)(unsafe.Add(mBase, _consts[348])) = l0
								v84 = *(*int64)(unsafe.Add(mBase, _consts[349]))
								v91 = int32(_a673)
								v92 = *(*int64)(unsafe.Add(mBase, _consts[350]))
								if v92 < int64(1) {
								} else {
									v95 = l0 - v92
									if int64(1) <= v95 {
										v99 = *(*int64)(unsafe.Add(mBase, _consts[351]))
										v102 = base.I64_div_s((v84-v99)*int64(1000000), v95)
										v103 = v102
									} else {
										v103 = int64(0)
									}
									v104 = *(*int32)(unsafe.Add(mBase, _consts[352]))
									*(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[353]))) = v103
									v114 = base.I32_rem_s(v104+int32(1), int32(16))
									*(*int32)(unsafe.Add(mBase, _consts[352])) = v114
								}
								*(*int64)(unsafe.Add(mBase, _consts[351])) = v84
								*(*int64)(unsafe.Add(mBase, _consts[350])) = l0
							}
						}
					}
				}
			}
			return
		}
	}
}
func F_getIOThreadPollResults(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, _consts[363]))
	switch v5 + int32(-1) {
	case 0:
		v39 = int32(0)
		return v39
	case 1:
		v16 = int32(_a20)
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[363])) = v17
		v21 = *(*int64)(unsafe.Add(mBase, _consts[364]))
		*(*int64)(unsafe.Add(mBase, _consts[364])) = v21 + int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v17
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v28&int32(-33) | int32(0)
		v38 = *(*int32)(unsafe.Add(mBase, _consts[365]))
		v39 = v38
		return v39
	default:
		F__serverAssert(m, int32(_a678), int32(_a672), int32(732))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_ioThreadReadQueryFromClient(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v2 != int32(1) {
		F__serverAssert(m, int32(_a1756), int32(_a1630), int32(6670))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v5 = F_readToQueryBuf(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
			if v7&int32(4) != 0 {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+289)))
				if v22&int32(64) != 0 {
					v27 = int32(2)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
					v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
					F_sendToMainThread(m, l0, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						return
					}
				} else {
					F_trimClientQueryBuffer(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = int32(2)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
						v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
						F_sendToMainThread(m, l0, int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
				if v10 < int32(1) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+289)))
					if v22&int32(64) != 0 {
						v27 = int32(2)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
						v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
						F_sendToMainThread(m, l0, int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							return
						}
					} else {
						F_trimClientQueryBuffer(m, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
							v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
							F_sendToMainThread(m, l0, int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+288)))
					if v13&int32(32769) != 0 {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+289)))
						if v22&int32(64) != 0 {
							v27 = int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
							v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
							F_sendToMainThread(m, l0, int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						} else {
							F_trimClientQueryBuffer(m, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								v27 = int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
								v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
								F_sendToMainThread(m, l0, int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_parseInputBuffer(m, l0)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							F_trimCommandQueue(m, l0)
							mBase = m.M
							v19 = m.ExcPending
							if v19 != 0 {
								return
							} else {
								F_prepareCommandQueue(m, l0)
								mBase = m.M
								v21 = m.ExcPending
								if v21 != 0 {
									return
								} else {
									v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+289)))
									if v22&int32(64) != 0 {
										v27 = int32(2)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
										v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
										F_sendToMainThread(m, l0, int32(0))
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											return
										}
									} else {
										F_trimClientQueryBuffer(m, l0)
										mBase = m.M
										v26 = m.ExcPending
										if v26 != 0 {
											return
										} else {
											v27 = int32(2)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v27)
											v30 = *(*int32)(unsafe.Add(mBase, _consts[346]))
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+225)) = uint8(v30)
											F_sendToMainThread(m, l0, int32(0))
											mBase = m.M
											v34 = m.ExcPending
											if v34 != 0 {
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
func F_updateIOThreads(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v71 == v74 {
		v217 = int32(1)
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v71 = v66 + int32(1)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a671), int32(_a672), int32(438))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v15 = int32(255)
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_consts[359])))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v15 + int32(-1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26<<(uint(int32(2))%32))+uint32(_consts[359])))
	if v31 != 0 {
		v66 = v26
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v66 = v15
	goto L2
L9:
	;
	v33 = v15 + int32(-2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32))+uint32(_consts[359])))
	if v38 != 0 {
		v66 = v33
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v40 = v15 + int32(-3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40<<(uint(int32(2))%32))+uint32(_consts[359])))
	if v45 != 0 {
		v66 = v40
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v47 = v15 + int32(-4)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_consts[359])))
	if v52 != 0 {
		v66 = v47
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v55 = int32(1)
	if base.Ui32(v55) < base.Ui32(v47) {
		v15 = v15 + int32(-5)
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v71 = v55
	goto L1
L14:
	;
	return int32(0)
L15:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	m.G0 = v8 + int32(16)
	return v217
L17:
	;
	v76 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	v79 = int32(_a20)
	v80 = *(*int64)(unsafe.Add(mBase, _consts[361]))
	v82 = *(*int64)(unsafe.Add(mBase, _consts[362]))
	if base.Ui32(base.I32_wrap_i64(v80+v82)) <= base.Ui32(v78) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v91 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if l0 == int32(0) {
		v217 = v76
		goto L16
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a676)
	v217 = v76
	goto L16
L21:
	;
	F_initIOThreads(m, v71)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L14
	} else {
		goto L54
	}
L22:
	;
	F__serverAssert(m, int32(_a671), int32(_a672), int32(85))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L14
	} else {
		goto L53
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v103 < int32(2) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v71
	F__serverLog(m, int32(2), int32(_a677), v8)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if v101 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, _consts[357]))
	v134 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	if v132 == v134 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v109 = int32(1)
	goto L29
L29:
	;
	v113 = v109 * int32(192)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[354])))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[355])))
	if v117 == v118 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L27
L31:
	;
	v122 = v109 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v122 < v124 {
		v109 = v122
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[355]))) = v117
	goto L32
L34:
	;
	goto L30
L35:
	;
	v151 = int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v153 <= v151 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	goto L37
L37:
	;
	v141 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[357]))
	v144 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	if v142 != v144 {
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L35
L39:
	;
	goto L38
L40:
	;
	v176 = int32(1)
	v177 = int32(_a20)
	*(*int32)(unsafe.Add(mBase, _consts[347])) = v176
	v181 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v71 < v181 {
		goto L21
	} else {
		goto L46
	}
L41:
	;
	v158 = v151
	goto L42
L42:
	;
	goto L44
L43:
	;
	goto L40
L44:
	;
	v167 = v158 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, _consts[347]))
	if v167 < v169 {
		v158 = v167
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v71 <= v181 {
		v217 = v176
		goto L16
	} else {
		goto L47
	}
L47:
	;
	v188 = v71
	goto L48
L48:
	;
	v190 = v188 + int32(-1)
	goto L50
L50:
	;
	F_shutdownIOThread(m, v190)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190<<(uint(int32(2))%32))+uint32(_consts[359]))) = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, _consts[345]))
	if v205 < v190 {
		v188 = v190
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v217 = v176
	goto L16
L53:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v217 = v176
	goto L16
}
