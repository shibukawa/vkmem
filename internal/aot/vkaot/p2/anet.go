package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_anetKeepAlive(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v9
	v12 = int32(9)
	v14 = v7 + int32(44)
	v15 = int32(4)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = F___syscall_setsockopt(m, l1, v9, v12, v14, v15, int32(0))
	mBase = m.M
	if v24 != int32(-50) {
		v71 = v24
		v73 = F___syscall_ret(m, v71)
		mBase = m.M
		v74 = v73
	} else {
		switch int32(-54) {
		case 0, 1:
			v70 = F___syscall_setsockopt(m, l1, int32(1), v12, v14, v15, int32(0))
			mBase = m.M
			v71 = v70
			v73 = F___syscall_ret(m, v71)
			mBase = m.M
			v74 = v73
		default:
			v71 = int32(-50)
			v73 = F___syscall_ret(m, v71)
			mBase = m.M
			v74 = v73
		case 3, 4:
			v35 = F___syscall_ret(m, int32(-28))
			mBase = m.M
			v74 = v35
		}
	}
	m.G0 = v21 + int32(16)
	if v74 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = l2
		v94 = int32(4)
		v101 = m.G0
		v103 = v101 - int32(16)
		m.G0 = v103
		v106 = F___syscall_setsockopt(m, l1, int32(6), v94, v7+int32(40), v94, int32(0))
		mBase = m.M
		v155 = F___syscall_ret(m, v106)
		mBase = m.M
		m.G0 = v103 + int32(16)
		if v155 == int32(0) {
			v166 = base.I32_div_s(l2, int32(3))
			if base.Ui32(l2+int32(2)) < base.Ui32(int32(5)) {
				v171 = int32(1)
			} else {
				v171 = v166
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v171
			v181 = m.G0
			v183 = v181 - int32(16)
			m.G0 = v183
			v186 = F___syscall_setsockopt(m, l1, int32(6), int32(5), v7+int32(36), int32(4), int32(0))
			mBase = m.M
			v235 = F___syscall_ret(m, v186)
			mBase = m.M
			m.G0 = v183 + int32(16)
			if v235 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(3)
				v246 = int32(6)
				v254 = m.G0
				v256 = v254 - int32(16)
				m.G0 = v256
				v259 = F___syscall_setsockopt(m, l1, v246, v246, v7+int32(32), int32(4), int32(0))
				mBase = m.M
				v308 = F___syscall_ret(m, v259)
				mBase = m.M
				m.G0 = v256 + int32(16)
				if v308 == int32(0) {
					v327 = int32(0)
					m.G0 = v7 + int32(48)
					return v327
				} else {
					v317 = int32(_a72)
					v319 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v320 = F___strerror_l(m, v319, v319)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v320
					F_anetSetError(m, l0, v317, v7)
					mBase = m.M
					v323 = m.ExcPending
					if v323 != 0 {
						return int32(0)
					} else {
						v327 = int32(-1)
						m.G0 = v7 + int32(48)
						return v327
					}
				}
			} else {
				v317 = int32(_a73)
				v319 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v320 = F___strerror_l(m, v319, v319)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v320
				F_anetSetError(m, l0, v317, v7)
				mBase = m.M
				v323 = m.ExcPending
				if v323 != 0 {
					return int32(0)
				} else {
					v327 = int32(-1)
					m.G0 = v7 + int32(48)
					return v327
				}
			}
		} else {
			v317 = int32(_a74)
			v319 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			v320 = F___strerror_l(m, v319, v319)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v320
			F_anetSetError(m, l0, v317, v7)
			mBase = m.M
			v323 = m.ExcPending
			if v323 != 0 {
				return int32(0)
			} else {
				v327 = int32(-1)
				m.G0 = v7 + int32(48)
				return v327
			}
		}
	} else {
		v82 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v83 = F___strerror_l(m, v82, v82)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v83
		F_anetSetError(m, l0, int32(_a75), v7+int32(16))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return int32(0)
		} else {
			v327 = int32(-1)
			m.G0 = v7 + int32(48)
			return v327
		}
	}
}
func F_anetListen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v13 = F_bind(m, l1, l2, l3)
	mBase = m.M
	if v13 != int32(-1) {
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
		if l5 == int32(0) {
			v36 = v25
		} else {
			if v25&int32(65535) != int32(1) {
				v36 = v25
			} else {
				v34 = F_chmod(m, l2+int32(2), l5)
				mBase = m.M
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
				v36 = v35
			}
		}
		if l6 == int32(0) {
			v76 = F_listen(m, l1, l4)
			mBase = m.M
			if v76 != int32(-1) {
				v92 = int32(0)
				m.G0 = v11 + int32(64)
				return v92
			} else {
				v80 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v81 = F___strerror_l(m, v80, v80)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v81
				F_anetSetError(m, l0, int32(_a82), v11+int32(48))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v90 = F_close(m, l1)
					mBase = m.M
					v92 = int32(-1)
					m.G0 = v11 + int32(64)
					return v92
				}
			}
		} else {
			if v36&int32(65535) != int32(1) {
				v76 = F_listen(m, l1, l4)
				mBase = m.M
				if v76 != int32(-1) {
					v92 = int32(0)
					m.G0 = v11 + int32(64)
					return v92
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v81 = F___strerror_l(m, v80, v80)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v81
					F_anetSetError(m, l0, int32(_a82), v11+int32(48))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v90 = F_close(m, l1)
						mBase = m.M
						v92 = int32(-1)
						m.G0 = v11 + int32(64)
						return v92
					}
				}
			} else {
				v43 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(44)
				v48 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v49 = F___strerror_l(m, v48, v48)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l6
				F_anetSetError(m, l0, int32(_a83), v11+int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v90 = F_close(m, l1)
					mBase = m.M
					v92 = int32(-1)
					m.G0 = v11 + int32(64)
					return v92
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v18 = F___strerror_l(m, v17, v17)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
		F_anetSetError(m, l0, int32(_a81), v11)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v90 = F_close(m, l1)
			mBase = m.M
			v92 = int32(-1)
			m.G0 = v11 + int32(64)
			return v92
		}
	}
}
func F_anetPipe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v11 = F_pipe(m, l0)
	mBase = m.M
	if v11 != 0 {
		v68 = int32(-1)
	} else {
		if l1&int32(524288) == int32(0) {
			if l2&int32(524288) == int32(0) {
				v37 = l1 & int32(-524289)
				if v37 == int32(0) {
					v47 = int32(0)
					v49 = l2 & int32(-524289)
					if v49 == v47 {
						v68 = v47
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
						v55 = F_fcntl(m, v52, int32(4), v8)
						mBase = m.M
						if v55 == int32(0) {
							v68 = v47
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v62 = F_close(m, v61)
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v64 = F_close(m, v63)
							mBase = m.M
							v68 = int32(-1)
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v37
					v45 = F_fcntl(m, v40, int32(4), v8+int32(16))
					mBase = m.M
					if v45 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v62 = F_close(m, v61)
						mBase = m.M
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v64 = F_close(m, v63)
						mBase = m.M
						v68 = int32(-1)
					} else {
						v47 = int32(0)
						v49 = l2 & int32(-524289)
						if v49 == v47 {
							v68 = v47
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							v55 = F_fcntl(m, v52, int32(4), v8)
							mBase = m.M
							if v55 == int32(0) {
								v68 = v47
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = F_close(m, v61)
								mBase = m.M
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v64 = F_close(m, v63)
								mBase = m.M
								v68 = int32(-1)
							}
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(1)
				v34 = F_fcntl(m, v28, int32(2), v8+int32(32))
				mBase = m.M
				if v34 != 0 {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v62 = F_close(m, v61)
					mBase = m.M
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v64 = F_close(m, v63)
					mBase = m.M
					v68 = int32(-1)
				} else {
					v37 = l1 & int32(-524289)
					if v37 == int32(0) {
						v47 = int32(0)
						v49 = l2 & int32(-524289)
						if v49 == v47 {
							v68 = v47
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							v55 = F_fcntl(m, v52, int32(4), v8)
							mBase = m.M
							if v55 == int32(0) {
								v68 = v47
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = F_close(m, v61)
								mBase = m.M
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v64 = F_close(m, v63)
								mBase = m.M
								v68 = int32(-1)
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v37
						v45 = F_fcntl(m, v40, int32(4), v8+int32(16))
						mBase = m.M
						if v45 != 0 {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v62 = F_close(m, v61)
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v64 = F_close(m, v63)
							mBase = m.M
							v68 = int32(-1)
						} else {
							v47 = int32(0)
							v49 = l2 & int32(-524289)
							if v49 == v47 {
								v68 = v47
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
								v55 = F_fcntl(m, v52, int32(4), v8)
								mBase = m.M
								if v55 == int32(0) {
									v68 = v47
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v62 = F_close(m, v61)
									mBase = m.M
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v64 = F_close(m, v63)
									mBase = m.M
									v68 = int32(-1)
								}
							}
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(1)
			v22 = F_fcntl(m, v16, int32(2), v8+int32(48))
			mBase = m.M
			if v22 != 0 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v62 = F_close(m, v61)
				mBase = m.M
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v64 = F_close(m, v63)
				mBase = m.M
				v68 = int32(-1)
			} else {
				if l2&int32(524288) == int32(0) {
					v37 = l1 & int32(-524289)
					if v37 == int32(0) {
						v47 = int32(0)
						v49 = l2 & int32(-524289)
						if v49 == v47 {
							v68 = v47
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							v55 = F_fcntl(m, v52, int32(4), v8)
							mBase = m.M
							if v55 == int32(0) {
								v68 = v47
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = F_close(m, v61)
								mBase = m.M
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v64 = F_close(m, v63)
								mBase = m.M
								v68 = int32(-1)
							}
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v37
						v45 = F_fcntl(m, v40, int32(4), v8+int32(16))
						mBase = m.M
						if v45 != 0 {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v62 = F_close(m, v61)
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v64 = F_close(m, v63)
							mBase = m.M
							v68 = int32(-1)
						} else {
							v47 = int32(0)
							v49 = l2 & int32(-524289)
							if v49 == v47 {
								v68 = v47
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
								v55 = F_fcntl(m, v52, int32(4), v8)
								mBase = m.M
								if v55 == int32(0) {
									v68 = v47
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v62 = F_close(m, v61)
									mBase = m.M
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v64 = F_close(m, v63)
									mBase = m.M
									v68 = int32(-1)
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(1)
					v34 = F_fcntl(m, v28, int32(2), v8+int32(32))
					mBase = m.M
					if v34 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v62 = F_close(m, v61)
						mBase = m.M
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v64 = F_close(m, v63)
						mBase = m.M
						v68 = int32(-1)
					} else {
						v37 = l1 & int32(-524289)
						if v37 == int32(0) {
							v47 = int32(0)
							v49 = l2 & int32(-524289)
							if v49 == v47 {
								v68 = v47
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
								v55 = F_fcntl(m, v52, int32(4), v8)
								mBase = m.M
								if v55 == int32(0) {
									v68 = v47
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v62 = F_close(m, v61)
									mBase = m.M
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v64 = F_close(m, v63)
									mBase = m.M
									v68 = int32(-1)
								}
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v37
							v45 = F_fcntl(m, v40, int32(4), v8+int32(16))
							mBase = m.M
							if v45 != 0 {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = F_close(m, v61)
								mBase = m.M
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v64 = F_close(m, v63)
								mBase = m.M
								v68 = int32(-1)
							} else {
								v47 = int32(0)
								v49 = l2 & int32(-524289)
								if v49 == v47 {
									v68 = v47
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
									v55 = F_fcntl(m, v52, int32(4), v8)
									mBase = m.M
									if v55 == int32(0) {
										v68 = v47
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v62 = F_close(m, v61)
										mBase = m.M
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v64 = F_close(m, v63)
										mBase = m.M
										v68 = int32(-1)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(64)
	return v68
}
func F_anetRetryAcceptOnError(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(13))
}
func F_anetSendTimeout(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int64(1000)
	v12 = base.I64_div_s(l2, v11)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = base.I32_wrap_i64(l2-v12*v11) * int32(1000)
	v23 = int32(67)
	v24 = int32(16)
	v25 = v9 + v24
	v30 = m.G0
	v32 = v30 - v24
	m.G0 = v32
	v35 = F___syscall_setsockopt(m, l1, int32(1), v23, v25, v24, v4)
	mBase = m.M
	if v35 != int32(-50) {
		v82 = v35
		v84 = F___syscall_ret(m, v82)
		mBase = m.M
		v85 = v84
	} else {
		switch int32(4) {
		case 0, 1:
			v81 = F___syscall_setsockopt(m, l1, int32(1), v23, v25, v24, int32(0))
			mBase = m.M
			v82 = v81
			v84 = F___syscall_ret(m, v82)
			mBase = m.M
			v85 = v84
		default:
			v82 = int32(-50)
			v84 = F___syscall_ret(m, v82)
			mBase = m.M
			v85 = v84
		case 3, 4:
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
			if base.Ui64(v47+int64(2147483648)) < base.Ui64(int64(4294967296)) {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v54
				*(*uint32)(unsafe.Add(mBase, uint32(v32)+8)) = uint32(v47)
				v66 = int32(8)
				v70 = F___syscall_setsockopt(m, l1, int32(1), int32(21), v32+v66, v66, int32(0))
				mBase = m.M
				v82 = v70
				v84 = F___syscall_ret(m, v82)
				mBase = m.M
				v85 = v84
			} else {
				v53 = F___syscall_ret(m, int32(-138))
				mBase = m.M
				v85 = v53
			}
		}
	}
	m.G0 = v32 + int32(16)
	if v85 != int32(-1) {
		v102 = v4
		m.G0 = v9 + int32(32)
		return v102
	} else {
		v93 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v94 = F___strerror_l(m, v93, v93)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
		F_anetSetError(m, l0, int32(_a76), v9)
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			v102 = int32(-1)
			m.G0 = v9 + int32(32)
			return v102
		}
	}
}
func F_anetTcp6Server(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F__anetTcpServer(m, l0, l1, l2, int32(10), l3, l4)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anetTcpAccept(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(128)
	v17 = F_anetGenericAccept(m, l0, l1, v9+int32(12), v9+int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(-1) {
			m.G0 = v9 + int32(144)
			return v17
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
			if v23 != int32(2) {
				if l2 == int32(0) {
					if l4 == int32(0) {
					} else {
						v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
						v43 = F___bswap_16_2(m, v42)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v43
					}
					v45 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v45
					v55 = m.G0
					v57 = v55 - int32(16)
					m.G0 = v57
					v60 = F___syscall_setsockopt(m, v17, int32(6), v45, v9+int32(140), int32(4), int32(0))
					mBase = m.M
					v109 = F___syscall_ret(m, v60)
					mBase = m.M
					m.G0 = v57 + int32(16)
					if v109 != int32(-1) {
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v119 = F___strerror_l(m, v118, v118)
						mBase = m.M
					}
					m.G0 = v9 + int32(144)
					return v17
				} else {
					v38 = F_inet_ntop(m, int32(10), v9+int32(20), l2, l3)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						if l4 == int32(0) {
						} else {
							v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
							v43 = F___bswap_16_2(m, v42)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v43
						}
						v45 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v45
						v55 = m.G0
						v57 = v55 - int32(16)
						m.G0 = v57
						v60 = F___syscall_setsockopt(m, v17, int32(6), v45, v9+int32(140), int32(4), int32(0))
						mBase = m.M
						v109 = F___syscall_ret(m, v60)
						mBase = m.M
						m.G0 = v57 + int32(16)
						if v109 != int32(-1) {
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							v119 = F___strerror_l(m, v118, v118)
							mBase = m.M
						}
						m.G0 = v9 + int32(144)
						return v17
					}
				}
			} else {
				if l2 == int32(0) {
					if l4 != 0 {
						v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
						v43 = F___bswap_16_2(m, v42)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v43
					} else {
					}
					v45 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v45
					v55 = m.G0
					v57 = v55 - int32(16)
					m.G0 = v57
					v60 = F___syscall_setsockopt(m, v17, int32(6), v45, v9+int32(140), int32(4), int32(0))
					mBase = m.M
					v109 = F___syscall_ret(m, v60)
					mBase = m.M
					m.G0 = v57 + int32(16)
					if v109 != int32(-1) {
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v119 = F___strerror_l(m, v118, v118)
						mBase = m.M
					}
					m.G0 = v9 + int32(144)
					return v17
				} else {
					v31 = F_inet_ntop(m, int32(2), v9+int32(16), l2, l3)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if l4 != 0 {
							v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
							v43 = F___bswap_16_2(m, v42)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v43
						} else {
						}
						v45 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v45
						v55 = m.G0
						v57 = v55 - int32(16)
						m.G0 = v57
						v60 = F___syscall_setsockopt(m, v17, int32(6), v45, v9+int32(140), int32(4), int32(0))
						mBase = m.M
						v109 = F___syscall_ret(m, v60)
						mBase = m.M
						m.G0 = v57 + int32(16)
						if v109 != int32(-1) {
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							v119 = F___strerror_l(m, v118, v118)
							mBase = m.M
						}
						m.G0 = v9 + int32(144)
						return v17
					}
				}
			}
		}
	}
}
func F_anetTcpGenericConnect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l2
	v24 = F_snprintf(m, v15+int32(118), int32(6), int32(_a77), v15+int32(64))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = v15 + int32(88)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(104)))) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v15+int32(96)))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v30
	v50 = m.Env.Getaddrinfo(m, l1, v15+int32(118), v15+int32(80), v15+int32(76))
	mBase = m.M
	if v50 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v15 + int32(128)
	return v409
L4:
	;
	v363 = int32(_a78)
	v365 = v50 + int32(1)
	if v365 == int32(0) {
		v385 = v363
		goto L93
	} else {
		goto L94
	}
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	if v51 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+20))
	F_emscripten_builtin_free(m, v347)
	mBase = m.M
	F_emscripten_builtin_free(m, v346)
	mBase = m.M
	goto L87
L7:
	;
	v345 = int32(-1)
	goto L6
L8:
	;
	v320 = F_close(m, v78)
	mBase = m.M
	goto L7
L9:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	F_emscripten_builtin_free(m, v296)
	mBase = m.M
	F_emscripten_builtin_free(m, v290)
	mBase = m.M
	goto L83
L10:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v290 = v283
	goto L9
L11:
	;
	v208 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v208
	v218 = m.G0
	v220 = v218 - int32(16)
	m.G0 = v220
	v223 = F___syscall_setsockopt(m, v78, int32(6), v208, v15+int32(124), int32(4), int32(0))
	mBase = m.M
	goto L59
L12:
	;
	goto L54
L13:
	;
	v57 = l4 & int32(1)
	if v57 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = int32(7)
	goto L16
L15:
	;
	v58 = int32(5)
	goto L16
L16:
	;
	if base.Ui32(l4) < base.Ui32(int32(4)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = int32(6)
	goto L19
L18:
	;
	v63 = int32(262)
	goto L19
L19:
	;
	v71 = v51
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v78 = F_anetCreateSocket(m, l0, v76, v77, v63, v58)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L12
L22:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	if v188 != 0 {
		v71 = v188
		goto L20
	} else {
		goto L53
	}
L23:
	;
	if v78 == int32(-1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if l3 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	v166 = F_connect(m, v78, v164, v165)
	mBase = m.M
	if v166 != int32(-1) {
		goto L11
	} else {
		goto L48
	}
L26:
	;
	v89 = m.Env.Getaddrinfo(m, l3, int32(0), v15+int32(80), v15+int32(72))
	mBase = m.M
	if v89 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v134 = v90
	goto L43
L28:
	;
	v94 = int32(_a78)
	v96 = v89 + int32(1)
	if v96 == int32(0) {
		v116 = v94
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	if v90 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v290 = int32(0)
	goto L9
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v116 + base.B2i32(v118 == int32(0))
	F_anetSetError(m, l0, int32(_a79), v15+int32(32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L41
	}
L32:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	goto L31
L33:
	;
	v100 = v94
	v101 = v96
	goto L34
L34:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v102 == int32(0) {
		v116 = v100
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v116 = v112
	goto L32
L36:
	;
	v106 = v100
	goto L37
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v110 != 0 {
		v106 = v106 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v112 = v106 + int32(2)
	v114 = v101 + int32(1)
	if v114 != 0 {
		v100 = v112
		v101 = v114
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L35
L41:
	;
	goto L8
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	F_emscripten_builtin_free(m, v149)
	mBase = m.M
	F_emscripten_builtin_free(m, v148)
	mBase = m.M
	goto L47
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	v142 = F_bind(m, v78, v140, v141)
	mBase = m.M
	if v142 != int32(-1) {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+28))
	if v145 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v134 = v145
	goto L43
L47:
	;
	goto L25
L48:
	;
	goto L49
L49:
	;
	if v57 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v175 = F_close(m, v78)
	mBase = m.M
	goto L22
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v172 == int32(26) {
		v345 = v78
		goto L6
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L21
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v203 = F___strerror_l(m, v202, v202)
	mBase = m.M
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v203
	F_anetSetError(m, l0, int32(_a80), v15)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L7
L57:
	;
	if v272 != int32(-1) {
		v345 = v78
		goto L6
	} else {
		goto L80
	}
L58:
	;
	m.G0 = v220 + int32(16)
	goto L57
L59:
	;
	v272 = F___syscall_ret(m, v223)
	mBase = m.M
	goto L58
L80:
	;
	goto L81
L81:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v282 = F___strerror_l(m, v281, v281)
	mBase = m.M
	goto L82
L82:
	;
	v345 = v78
	goto L6
L83:
	;
	goto L84
L84:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v301 = F___strerror_l(m, v300, v300)
	mBase = m.M
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v301
	F_anetSetError(m, l0, int32(_a81), v15+int32(16))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L8
L87:
	;
	if l4&int32(2) == int32(0) {
		v409 = v345
		goto L3
	} else {
		goto L88
	}
L88:
	;
	if l3 == int32(0) {
		v409 = v345
		goto L3
	} else {
		goto L89
	}
L89:
	;
	if v345 != int32(-1) {
		v409 = v345
		goto L3
	} else {
		goto L90
	}
L90:
	;
	v359 = F_anetTcpGenericConnect(m, l0, l1, l2, int32(0), l4)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v409 = v359
	goto L3
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v385 + base.B2i32(v387 == int32(0))
	F_anetSetError(m, l0, int32(_a79), v15+int32(48))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L102
	}
L93:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	goto L92
L94:
	;
	v369 = v363
	v370 = v365
	goto L95
L95:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	if v371 == int32(0) {
		v385 = v369
		goto L93
	} else {
		goto L97
	}
L96:
	;
	v385 = v381
	goto L93
L97:
	;
	v375 = v369
	goto L98
L98:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+1)))
	if v379 != 0 {
		v375 = v375 + int32(1)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v381 = v375 + int32(2)
	v383 = v370 + int32(1)
	if v383 != 0 {
		v369 = v381
		v370 = v383
		goto L95
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	goto L96
L102:
	;
	v409 = int32(-1)
	goto L3
}
func F_anetTcpNonBlockBestEffortBindConnect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	if l4 != 0 {
		v8 = int32(7)
	} else {
		v8 = int32(3)
	}
	v9 = F_anetTcpGenericConnect(m, l0, l1, l2, l3, v8)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
