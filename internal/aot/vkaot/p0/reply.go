package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__addReplyLongLongWithPrefix(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	if l1 < int64(10) {
		v15 = int32(4)
	} else {
		v15 = int32(5)
	}
	v17 = base.B2i32(base.Ui64(int64(31)) < base.Ui64(l1))
	if base.Ui64(int64(31)) < base.Ui64(l1) {
		if base.Ui64(int64(31)) < base.Ui64(l1) {
			v41 = base.B2i32(base.Ui64(int64(31)) < base.Ui64(l1))
			if base.Ui64(int64(31)) < base.Ui64(l1) {
				if base.Ui64(int64(31)) < base.Ui64(l1) {
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
					v66 = v9 | int32(1)
					if l1 <= int64(-1) {
						v76 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
						v80 = int32(1)
						v85 = v66 + v80
						v86 = int32(126)
						v87 = int64(0) - l1
						v88 = v80
					} else {
						v85 = v66
						v86 = int32(127)
						v87 = l1
						v88 = int32(0)
					}
					v89 = F_ull2string(m, v85, v86, v87)
					mBase = m.M
					if v89 == int32(0) {
						v108 = int32(0)
					} else {
						v108 = v89 + v88
					}
					v112 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
					F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						m.G0 = v9 + int32(128)
						return
					}
				} else {
					if v3 != int32(126) {
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
						v66 = v9 | int32(1)
						if l1 <= int64(-1) {
							v76 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
							v80 = int32(1)
							v85 = v66 + v80
							v86 = int32(126)
							v87 = int64(0) - l1
							v88 = v80
						} else {
							v85 = v66
							v86 = int32(127)
							v87 = l1
							v88 = int32(0)
						}
						v89 = F_ull2string(m, v85, v86, v87)
						mBase = m.M
						if v89 == int32(0) {
							v108 = int32(0)
						} else {
							v108 = v89 + v88
						}
						v112 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
						F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
						v61 = F_objectGetVal(m, v60)
						mBase = m.M
						F__addReplyToBufferOrList(m, l0, v61, v15)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					}
				}
			} else {
				if v3 != int32(37) {
					if base.Ui64(int64(31)) < base.Ui64(l1) {
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
						v66 = v9 | int32(1)
						if l1 <= int64(-1) {
							v76 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
							v80 = int32(1)
							v85 = v66 + v80
							v86 = int32(126)
							v87 = int64(0) - l1
							v88 = v80
						} else {
							v85 = v66
							v86 = int32(127)
							v87 = l1
							v88 = int32(0)
						}
						v89 = F_ull2string(m, v85, v86, v87)
						mBase = m.M
						if v89 == int32(0) {
							v108 = int32(0)
						} else {
							v108 = v89 + v88
						}
						v112 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
						F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					} else {
						if v3 != int32(126) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
							v61 = F_objectGetVal(m, v60)
							mBase = m.M
							F__addReplyToBufferOrList(m, l0, v61, v15)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[1])))
					v50 = F_objectGetVal(m, v49)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v50, v15)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						m.G0 = v9 + int32(128)
						return
					}
				}
			}
		} else {
			if v3 != int32(36) {
				v41 = base.B2i32(base.Ui64(int64(31)) < base.Ui64(l1))
				if base.Ui64(int64(31)) < base.Ui64(l1) {
					if base.Ui64(int64(31)) < base.Ui64(l1) {
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
						v66 = v9 | int32(1)
						if l1 <= int64(-1) {
							v76 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
							v80 = int32(1)
							v85 = v66 + v80
							v86 = int32(126)
							v87 = int64(0) - l1
							v88 = v80
						} else {
							v85 = v66
							v86 = int32(127)
							v87 = l1
							v88 = int32(0)
						}
						v89 = F_ull2string(m, v85, v86, v87)
						mBase = m.M
						if v89 == int32(0) {
							v108 = int32(0)
						} else {
							v108 = v89 + v88
						}
						v112 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
						F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					} else {
						if v3 != int32(126) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
							v61 = F_objectGetVal(m, v60)
							mBase = m.M
							F__addReplyToBufferOrList(m, l0, v61, v15)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				} else {
					if v3 != int32(37) {
						if base.Ui64(int64(31)) < base.Ui64(l1) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							if v3 != int32(126) {
								*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
								v66 = v9 | int32(1)
								if l1 <= int64(-1) {
									v76 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
									v80 = int32(1)
									v85 = v66 + v80
									v86 = int32(126)
									v87 = int64(0) - l1
									v88 = v80
								} else {
									v85 = v66
									v86 = int32(127)
									v87 = l1
									v88 = int32(0)
								}
								v89 = F_ull2string(m, v85, v86, v87)
								mBase = m.M
								if v89 == int32(0) {
									v108 = int32(0)
								} else {
									v108 = v89 + v88
								}
								v112 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
								F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
								v61 = F_objectGetVal(m, v60)
								mBase = m.M
								F__addReplyToBufferOrList(m, l0, v61, v15)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[1])))
						v50 = F_objectGetVal(m, v49)
						mBase = m.M
						F__addReplyToBufferOrList(m, l0, v50, v15)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[2])))
				v37 = F_objectGetVal(m, v36)
				mBase = m.M
				F__addReplyToBufferOrList(m, l0, v37, v15)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v9 + int32(128)
					return
				}
			}
		}
	} else {
		if v3 != int32(42) {
			if base.Ui64(int64(31)) < base.Ui64(l1) {
				v41 = base.B2i32(base.Ui64(int64(31)) < base.Ui64(l1))
				if base.Ui64(int64(31)) < base.Ui64(l1) {
					if base.Ui64(int64(31)) < base.Ui64(l1) {
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
						v66 = v9 | int32(1)
						if l1 <= int64(-1) {
							v76 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
							v80 = int32(1)
							v85 = v66 + v80
							v86 = int32(126)
							v87 = int64(0) - l1
							v88 = v80
						} else {
							v85 = v66
							v86 = int32(127)
							v87 = l1
							v88 = int32(0)
						}
						v89 = F_ull2string(m, v85, v86, v87)
						mBase = m.M
						if v89 == int32(0) {
							v108 = int32(0)
						} else {
							v108 = v89 + v88
						}
						v112 = int32(2573)
						*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
						F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					} else {
						if v3 != int32(126) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
							v61 = F_objectGetVal(m, v60)
							mBase = m.M
							F__addReplyToBufferOrList(m, l0, v61, v15)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				} else {
					if v3 != int32(37) {
						if base.Ui64(int64(31)) < base.Ui64(l1) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							if v3 != int32(126) {
								*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
								v66 = v9 | int32(1)
								if l1 <= int64(-1) {
									v76 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
									v80 = int32(1)
									v85 = v66 + v80
									v86 = int32(126)
									v87 = int64(0) - l1
									v88 = v80
								} else {
									v85 = v66
									v86 = int32(127)
									v87 = l1
									v88 = int32(0)
								}
								v89 = F_ull2string(m, v85, v86, v87)
								mBase = m.M
								if v89 == int32(0) {
									v108 = int32(0)
								} else {
									v108 = v89 + v88
								}
								v112 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
								F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
								v61 = F_objectGetVal(m, v60)
								mBase = m.M
								F__addReplyToBufferOrList(m, l0, v61, v15)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[1])))
						v50 = F_objectGetVal(m, v49)
						mBase = m.M
						F__addReplyToBufferOrList(m, l0, v50, v15)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							m.G0 = v9 + int32(128)
							return
						}
					}
				}
			} else {
				if v3 != int32(36) {
					v41 = base.B2i32(base.Ui64(int64(31)) < base.Ui64(l1))
					if base.Ui64(int64(31)) < base.Ui64(l1) {
						if base.Ui64(int64(31)) < base.Ui64(l1) {
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
							v66 = v9 | int32(1)
							if l1 <= int64(-1) {
								v76 = int32(45)
								*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
								v80 = int32(1)
								v85 = v66 + v80
								v86 = int32(126)
								v87 = int64(0) - l1
								v88 = v80
							} else {
								v85 = v66
								v86 = int32(127)
								v87 = l1
								v88 = int32(0)
							}
							v89 = F_ull2string(m, v85, v86, v87)
							mBase = m.M
							if v89 == int32(0) {
								v108 = int32(0)
							} else {
								v108 = v89 + v88
							}
							v112 = int32(2573)
							*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
							F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						} else {
							if v3 != int32(126) {
								*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
								v66 = v9 | int32(1)
								if l1 <= int64(-1) {
									v76 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
									v80 = int32(1)
									v85 = v66 + v80
									v86 = int32(126)
									v87 = int64(0) - l1
									v88 = v80
								} else {
									v85 = v66
									v86 = int32(127)
									v87 = l1
									v88 = int32(0)
								}
								v89 = F_ull2string(m, v85, v86, v87)
								mBase = m.M
								if v89 == int32(0) {
									v108 = int32(0)
								} else {
									v108 = v89 + v88
								}
								v112 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
								F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
								v61 = F_objectGetVal(m, v60)
								mBase = m.M
								F__addReplyToBufferOrList(m, l0, v61, v15)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					} else {
						if v3 != int32(37) {
							if base.Ui64(int64(31)) < base.Ui64(l1) {
								*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
								v66 = v9 | int32(1)
								if l1 <= int64(-1) {
									v76 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
									v80 = int32(1)
									v85 = v66 + v80
									v86 = int32(126)
									v87 = int64(0) - l1
									v88 = v80
								} else {
									v85 = v66
									v86 = int32(127)
									v87 = l1
									v88 = int32(0)
								}
								v89 = F_ull2string(m, v85, v86, v87)
								mBase = m.M
								if v89 == int32(0) {
									v108 = int32(0)
								} else {
									v108 = v89 + v88
								}
								v112 = int32(2573)
								*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
								F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							} else {
								if v3 != int32(126) {
									*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v3)
									v66 = v9 | int32(1)
									if l1 <= int64(-1) {
										v76 = int32(45)
										*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
										v80 = int32(1)
										v85 = v66 + v80
										v86 = int32(126)
										v87 = int64(0) - l1
										v88 = v80
									} else {
										v85 = v66
										v86 = int32(127)
										v87 = l1
										v88 = int32(0)
									}
									v89 = F_ull2string(m, v85, v86, v87)
									mBase = m.M
									if v89 == int32(0) {
										v108 = int32(0)
									} else {
										v108 = v89 + v88
									}
									v112 = int32(2573)
									*(*uint16)(unsafe.Add(mBase, uint32(v108+v9+int32(1)))) = uint16(v112)
									F__addReplyToBufferOrList(m, l0, v9, v108+int32(3))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										m.G0 = v9 + int32(128)
										return
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[0])))
									v61 = F_objectGetVal(m, v60)
									mBase = m.M
									F__addReplyToBufferOrList(m, l0, v61, v15)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v9 + int32(128)
										return
									}
								}
							}
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[1])))
							v50 = F_objectGetVal(m, v49)
							mBase = m.M
							F__addReplyToBufferOrList(m, l0, v50, v15)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[2])))
					v37 = F_objectGetVal(m, v36)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v37, v15)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(128)
						return
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l1)<<(uint(int32(2))%32))+uint32(_c_F__addReplyLongLongWithPrefix[3])))
			v26 = F_objectGetVal(m, v25)
			mBase = m.M
			F__addReplyToBufferOrList(m, l0, v26, v15)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				m.G0 = v9 + int32(128)
				return
			}
		}
	}
}
func F_addReplyArrayLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	if l1 <= int32(-1) {
		F__serverAssert(m, int32(_a_F_addReplyArrayLen_0), int32(_a_F_addReplyArrayLen_1), int32(1420))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v10 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 != 0 {
				m.G0 = v6 + int32(128)
				return
			} else {
				if base.Ui32(int32(31)) < base.Ui32(l1) {
					v27 = int32(42)
					*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v27)
					v30 = v6 | int32(1)
					v32 = base.I64_extend_i32_u(l1)
					if v32 <= int64(-1) {
						v41 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v41)
						v45 = int32(1)
						v50 = v30 + v45
						v51 = int32(126)
						v52 = int64(0) - v32
						v53 = v45
					} else {
						v50 = v30
						v51 = int32(127)
						v52 = v32
						v53 = int32(0)
					}
					v54 = F_ull2string(m, v50, v51, v52)
					mBase = m.M
					if v54 == int32(0) {
						v73 = int32(0)
					} else {
						v73 = v54 + v53
					}
					v77 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v73+v6+int32(1)))) = uint16(v77)
					F__addReplyToBufferOrList(m, l0, v6, v73+int32(3))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32))+uint32(_c_F_addReplyArrayLen[0])))
					v19 = F_objectGetVal(m, v18)
					mBase = m.M
					if base.Ui32(l1) < base.Ui32(int32(10)) {
						v24 = int32(4)
					} else {
						v24 = int32(5)
					}
					F__addReplyToBufferOrList(m, l0, v19, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				}
			}
		}
	}
}
func F_addReplyAttributeLen(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v8) <= base.Ui32(int32(2)) {
		F__serverAssert(m, int32(_a_F_addReplyAttributeLen_0), int32(_a_F_addReplyAttributeLen_1), int32(1454))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if l1 <= int32(-1) {
			F__serverAssert(m, int32(_a_F_addReplyAttributeLen_2), int32(_a_F_addReplyAttributeLen_1), int32(1420))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v13 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				if v13 != 0 {
					m.G0 = v6 + int32(128)
					return
				} else {
					v15 = int32(124)
					*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v15)
					v18 = v6 | int32(1)
					v20 = base.I64_extend_i32_u(l1)
					if v20 <= int64(-1) {
						v29 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v29)
						v33 = int32(1)
						v38 = v18 + v33
						v39 = int32(126)
						v40 = int64(0) - v20
						v41 = v33
					} else {
						v38 = v18
						v39 = int32(127)
						v40 = v20
						v41 = int32(0)
					}
					v42 = F_ull2string(m, v38, v39, v40)
					mBase = m.M
					if v42 == int32(0) {
						v61 = int32(0)
					} else {
						v61 = v42 + v41
					}
					v65 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v61+v6+int32(1)))) = uint16(v65)
					F__addReplyToBufferOrList(m, l0, v6, v61+int32(3))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						m.G0 = v6 + int32(128)
						return
					}
				}
			}
		}
	}
}
func F_addReplyBigNum(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v24 int32
	_ = v24
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v4 != int32(2) {
		v9 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if v9 != 0 {
				v15 = F_prepareClientToWrite(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if v15 != 0 {
						v19 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							if v19 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBigNum_0), int32(2))
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F__addReplyToBufferOrList(m, l0, l1, l2)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v19 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								if v19 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBigNum_0), int32(2))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBigNum_1), int32(1))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						if v15 != 0 {
							v19 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								if v19 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBigNum_0), int32(2))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F__addReplyToBufferOrList(m, l0, l1, l2)
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return
							} else {
								v19 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									if v19 != 0 {
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBigNum_0), int32(2))
										mBase = m.M
										v24 = m.ExcPending
										if v24 != 0 {
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
		F_addReplyBulkCBuffer(m, l0, l1, l2)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_addReplyBulk(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int64
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v14&int32(268435456) != 0 {
		v115 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_addReplyBulk_0), int32(_a_F_addReplyBulk_1), int32(1516))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L38
	} else {
		goto L107
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	F_addReplyBulkLen(m, l0, l1)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L38
	} else {
		goto L102
	}
L4:
	;
	if v120 == int32(0) {
		goto L3
	} else {
		goto L37
	}
L5:
	;
	v120 = v115
	goto L4
L6:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
	if v17 != int64(-1) {
		v115 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v14&int32(131072) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v48&int32(1) != 0 {
		v115 = v3
		goto L5
	} else {
		goto L19
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[0]))
	if l0 != v25 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[1]))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	if v31 == int32(0) {
		v115 = v3
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
	if v34 == int32(288) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v34 == int32(286) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v34 == int32(284) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v34 == int32(282) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v34 == int32(287) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v34 != int32(289) {
		v115 = v3
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	if v48&int32(2) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v48&int32(262144) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v48&int32(4) == int32(0) {
		v115 = v3
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if l1 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v61 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v64 = F_isImportSlotMigrationJob(m, v61)
	mBase = m.M
	v120 = int32(0)
	goto L4
L26:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v75&int32(240) != 0 {
		v115 = v3
		goto L5
	} else {
		goto L28
	}
L27:
	;
	v67 = int32(_a_F_addReplyBulk_2)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[2]))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[3]))
	v120 = base.B2i32(v68 != int32(0)) & base.B2i32(v68 <= v72)
	goto L4
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v78&int32(-8) == int32(-16) {
		v115 = v3
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v83 = int32(_a_F_addReplyBulk_2)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[3]))
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[2]))
	if v86 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v84 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v86 <= v84 {
		v115 = int32(1)
		goto L5
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[4]))
	if v106 == v104 {
		v115 = v104
		goto L5
	} else {
		goto L36
	}
L34:
	;
	v94 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[5]))
	if v96 == v94 {
		v115 = v94
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v99 = F_objectGetVal(m, l1)
	mBase = m.M
	v100 = F_sdslen_7(m, v99)
	mBase = m.M
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[5]))
	v120 = base.B2i32(base.Ui32(v102) <= base.Ui32(v100))
	goto L4
L36:
	;
	v109 = F_objectGetVal(m, l1)
	mBase = m.M
	v110 = F_sdslen_7(m, v109)
	mBase = m.M
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[4]))
	v115 = base.B2i32(base.Ui32(v112) <= base.Ui32(v110))
	goto L5
L37:
	;
	v123 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	if v123 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v125&int32(64) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = *(*int64)(unsafe.Add(mBase, _c_F_addReplyBulk[6]))
	if v199 == int64(-1) {
		goto L2
	} else {
		goto L62
	}
L42:
	;
	F_incrRefCount(m, l1)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
	v131 = F_objectGetVal(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v131
	v133 = int32(8)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v139&int32(16777216) != 0 {
		v147 = v139
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v187 != 0 {
		goto L41
	} else {
		goto L60
	}
L45:
	;
	goto L44
L46:
	;
	v149 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyBulk[7]))
	if v151 != 0 {
		v187 = v149
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v143 != 0 {
		v187 = int32(0)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v145 = v139 | int32(16777216)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v145
	v147 = v145
	goto L46
L49:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v153 != 0 {
		v187 = v149
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v156 = v154 - v155
	if v147&int32(16777216) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v173 == int32(0) {
		v187 = v149
		goto L45
	} else {
		goto L57
	}
L52:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v169 = *(*int64)(unsafe.Add(mBase, _c_F_addReplyBulk[6]))
	v172 = F_upsertPayloadHeader(m, v161, l0+int32(180), l0+int32(184), int32(1), v133, v167, base.B2i32(v169 != int64(-1)), v156)
	mBase = m.M
	v173 = v172
	goto L51
L53:
	;
	if base.Ui32(v156) < base.Ui32(v133) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v160 = v156
	goto L56
L55:
	;
	v160 = v133
	goto L56
L56:
	;
	v173 = v160
	goto L51
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v179 = F___memcpy(m, v176+v177, v9+v133, v173)
	mBase = m.M
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v181 = v180 + v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if base.Ui32(v181) <= base.Ui32(v183) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v187 = v173
	goto L45
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v181
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v192 = int32(8)
	F__addReplyPayloadToList(m, l0, v191, v9+v192, v192, int32(1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	goto L41
L62:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v202&int32(240) != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v206 = F_objectGetVal(m, l1)
	mBase = m.M
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+int32(-1)))))
	switch v209 & int32(7) {
	case 0:
		goto L69
	case 1:
		goto L68
	case 2:
		goto L67
	case 3:
		goto L66
	case 4:
		goto L65
	default:
		v226 = int32(0)
		goto L64
	}
L64:
	;
	v227 = base.I64_extend_i32_u(v226)
	v228 = int32(0)
	if base.Ui64(v227) < base.Ui64(int64(10)) {
		v287 = v228
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(-17))))
	v226 = v225
	goto L64
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v206+int32(-9))))
	v226 = v222
	goto L64
L67:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206+int32(-5)))))
	v226 = v219
	goto L64
L68:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+int32(-3)))))
	v226 = v216
	goto L64
L69:
	;
	v226 = int32(base.Ui32(v209) >> (uint(int32(3)) % 32))
	goto L64
L70:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+264)) = v227 + v295 + base.I64_extend_i32_u(v294+int32(3)) + int64(2)
	goto L2
L71:
	;
	v294 = int32(1) + v287
	goto L70
L72:
	;
	v233 = v227
	v234 = v228
	goto L73
L73:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v233) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v287 = v281
	goto L71
L75:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v233) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v294 = int32(2) + v234
	goto L70
L77:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v233) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v294 = int32(3) + v234
	goto L70
L79:
	;
	v281 = v234 + int32(12)
	v285 = base.I64_div_u_s(v233, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v233) {
		v233 = v285
		v234 = v281
		goto L73
	} else {
		goto L101
	}
L80:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v233) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v233) {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v233) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v233) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v233) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v233) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v294 = int32(4) + v234
	goto L70
L87:
	;
	v258 = int32(6)
	goto L89
L88:
	;
	v258 = int32(5)
	goto L89
L89:
	;
	v294 = v258 + v234
	goto L70
L90:
	;
	v264 = int32(8)
	goto L92
L91:
	;
	v264 = int32(7)
	goto L92
L92:
	;
	v294 = v264 + v234
	goto L70
L93:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v233) {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v233) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v272 = int32(10)
	goto L97
L96:
	;
	v272 = int32(9)
	goto L97
L97:
	;
	v294 = v272 + v234
	goto L70
L98:
	;
	v278 = int32(12)
	goto L100
L99:
	;
	v278 = int32(11)
	goto L100
L100:
	;
	v294 = v278 + v234
	goto L70
L101:
	;
	goto L74
L102:
	;
	F_addReply(m, l0, l1)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L38
	} else {
		goto L103
	}
L103:
	;
	v308 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L38
	} else {
		goto L104
	}
L104:
	;
	if v308 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulk_3), int32(2))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L38
	} else {
		goto L106
	}
L106:
	;
	goto L2
L107:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addReplyBulkSds(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v11 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			v18 = l1 + int32(-1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			switch v19 & int32(7) {
			case 0:
				if base.Ui32(v19) < base.Ui32(int32(80)) {
					v28 = int32(4)
				} else {
					v28 = int32(5)
				}
				v51 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
				v52 = v28
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
				v58 = F_objectGetVal(m, v57)
				mBase = m.M
				F__addReplyToBufferOrList(m, l0, v58, v52)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					switch v120 & int32(7) {
					case 0:
						v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
					case 1:
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v137 = v127
					case 2:
						v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v137 = v130
					case 3:
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v137 = v133
					case 4:
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v137 = v136
					default:
						v137 = int32(0)
					}
					F__addReplyToBufferOrList(m, l0, l1, v137)
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return
					} else {
						F_sdsfree(m, l1)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				}
			case 1:
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
				v43 = v33
				if base.Ui32(int32(31)) < base.Ui32(v43) {
					v61 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v61)
					v64 = v9 | int32(1)
					v66 = base.I64_extend_i32_u(v43)
					if v66 <= int64(-1) {
						v75 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v75)
						v79 = int32(1)
						v84 = v64 + v79
						v85 = int32(126)
						v86 = int64(0) - v66
						v87 = v79
					} else {
						v84 = v64
						v85 = int32(127)
						v86 = v66
						v87 = int32(0)
					}
					v88 = F_ull2string(m, v84, v85, v86)
					mBase = m.M
					if v88 == int32(0) {
						v107 = int32(0)
					} else {
						v107 = v88 + v87
					}
					v111 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v107+v9+int32(1)))) = uint16(v111)
					F__addReplyToBufferOrList(m, l0, v9, v107+int32(3))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				} else {
					if base.Ui32(v43) < base.Ui32(int32(10)) {
						v50 = int32(4)
					} else {
						v50 = int32(5)
					}
					v51 = v43
					v52 = v50
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
					v58 = F_objectGetVal(m, v57)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v58, v52)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				}
			case 2:
				v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
				v43 = v36
				if base.Ui32(int32(31)) < base.Ui32(v43) {
					v61 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v61)
					v64 = v9 | int32(1)
					v66 = base.I64_extend_i32_u(v43)
					if v66 <= int64(-1) {
						v75 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v75)
						v79 = int32(1)
						v84 = v64 + v79
						v85 = int32(126)
						v86 = int64(0) - v66
						v87 = v79
					} else {
						v84 = v64
						v85 = int32(127)
						v86 = v66
						v87 = int32(0)
					}
					v88 = F_ull2string(m, v84, v85, v86)
					mBase = m.M
					if v88 == int32(0) {
						v107 = int32(0)
					} else {
						v107 = v88 + v87
					}
					v111 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v107+v9+int32(1)))) = uint16(v111)
					F__addReplyToBufferOrList(m, l0, v9, v107+int32(3))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				} else {
					if base.Ui32(v43) < base.Ui32(int32(10)) {
						v50 = int32(4)
					} else {
						v50 = int32(5)
					}
					v51 = v43
					v52 = v50
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
					v58 = F_objectGetVal(m, v57)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v58, v52)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				}
			case 3:
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
				v43 = v39
				if base.Ui32(int32(31)) < base.Ui32(v43) {
					v61 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v61)
					v64 = v9 | int32(1)
					v66 = base.I64_extend_i32_u(v43)
					if v66 <= int64(-1) {
						v75 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v75)
						v79 = int32(1)
						v84 = v64 + v79
						v85 = int32(126)
						v86 = int64(0) - v66
						v87 = v79
					} else {
						v84 = v64
						v85 = int32(127)
						v86 = v66
						v87 = int32(0)
					}
					v88 = F_ull2string(m, v84, v85, v86)
					mBase = m.M
					if v88 == int32(0) {
						v107 = int32(0)
					} else {
						v107 = v88 + v87
					}
					v111 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v107+v9+int32(1)))) = uint16(v111)
					F__addReplyToBufferOrList(m, l0, v9, v107+int32(3))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				} else {
					if base.Ui32(v43) < base.Ui32(int32(10)) {
						v50 = int32(4)
					} else {
						v50 = int32(5)
					}
					v51 = v43
					v52 = v50
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
					v58 = F_objectGetVal(m, v57)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v58, v52)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				}
			case 4:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
				v43 = v42
				if base.Ui32(int32(31)) < base.Ui32(v43) {
					v61 = int32(36)
					*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v61)
					v64 = v9 | int32(1)
					v66 = base.I64_extend_i32_u(v43)
					if v66 <= int64(-1) {
						v75 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v75)
						v79 = int32(1)
						v84 = v64 + v79
						v85 = int32(126)
						v86 = int64(0) - v66
						v87 = v79
					} else {
						v84 = v64
						v85 = int32(127)
						v86 = v66
						v87 = int32(0)
					}
					v88 = F_ull2string(m, v84, v85, v86)
					mBase = m.M
					if v88 == int32(0) {
						v107 = int32(0)
					} else {
						v107 = v88 + v87
					}
					v111 = int32(2573)
					*(*uint16)(unsafe.Add(mBase, uint32(v107+v9+int32(1)))) = uint16(v111)
					F__addReplyToBufferOrList(m, l0, v9, v107+int32(3))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				} else {
					if base.Ui32(v43) < base.Ui32(int32(10)) {
						v50 = int32(4)
					} else {
						v50 = int32(5)
					}
					v51 = v43
					v52 = v50
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
					v58 = F_objectGetVal(m, v57)
					mBase = m.M
					F__addReplyToBufferOrList(m, l0, v58, v52)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						switch v120 & int32(7) {
						case 0:
							v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
						case 1:
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v137 = v127
						case 2:
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v137 = v130
						case 3:
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v137 = v133
						case 4:
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v137 = v136
						default:
							v137 = int32(0)
						}
						F__addReplyToBufferOrList(m, l0, l1, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							F_sdsfree(m, l1)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									m.G0 = v9 + int32(128)
									return
								}
							}
						}
					}
				}
			default:
				v51 = int32(0)
				v52 = int32(4)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32))+uint32(_c_F_addReplyBulkSds[0])))
				v58 = F_objectGetVal(m, v57)
				mBase = m.M
				F__addReplyToBufferOrList(m, l0, v58, v52)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					switch v120 & int32(7) {
					case 0:
						v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
					case 1:
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v137 = v127
					case 2:
						v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v137 = v130
					case 3:
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v137 = v133
					case 4:
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v137 = v136
					default:
						v137 = int32(0)
					}
					F__addReplyToBufferOrList(m, l0, l1, v137)
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return
					} else {
						F_sdsfree(m, l1)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyBulkSds_0), int32(2))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								m.G0 = v9 + int32(128)
								return
							}
						}
					}
				}
			}
		} else {
			F_sdsfree(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				m.G0 = v9 + int32(128)
				return
			}
		}
	}
}
func F_addReplyClusterLinkDescription(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_addReplyMapLen(m, l0, int32(6))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			if v18 != 0 {
				v19 = int32(_a_F_addReplyClusterLinkDescription_1)
			} else {
				v19 = int32(_a_F_addReplyClusterLinkDescription_2)
			}
			F_addReplyBulkCString(m, l0, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v22 == int32(0) {
					F__serverAssert(m, int32(_a_F_addReplyClusterLinkDescription_3), int32(_a_F_addReplyClusterLinkDescription_4), int32(7104))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v28 = F_sdsnewlen(m, v22+int32(8), int32(40))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_5))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_addReplyBulkCString(m, l0, v28)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								F_sdsfree(m, v28)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_6))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										v40 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										F_addReplyLongLong(m, l0, v40)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											if v43 != 0 {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
												if v46 != 0 {
													v49 = int32(114)
													*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v49)
													v53 = v8 + int32(14)
												} else {
													v53 = v8 + int32(13)
												}
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
												if v54 == int32(0) {
													v61 = v53
												} else {
													v57 = int32(119)
													*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
													v61 = v53 + int32(1)
												}
											} else {
												v61 = v8 + int32(13)
											}
											v62 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v62)
											F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_7))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												F_addReplyBulkCString(m, l0, v8+int32(13))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return
												} else {
													F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_8))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														v74 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
														F_addReplyLongLong(m, l0, v74)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return
														} else {
															F_addReplyBulkCString(m, l0, int32(_a_F_addReplyClusterLinkDescription_9))
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return
															} else {
																v80 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
																F_addReplyLongLong(m, l0, v80)
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
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
func F_addReplyDouble(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	v11 = m.G0
	v13 = v11 - int32(5152)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v15 != int32(3) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(5152)
	return
L2:
	;
	v42 = F_d2string(m, v13|int32(7), int32(5145), l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L11
	}
L3:
	;
	v18 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v18)
	v23 = F_d2string(m, v13|int32(1), int32(130), l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v26 = v23 + int32(3)
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v26))) = uint8(v28)
	v33 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+v13+int32(1)))) = uint16(v33)
	v35 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F__addReplyToBufferOrList(m, l0, v13, v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v168 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+5)) = uint16(v168)
	v171 = v42 + int32(9)
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+v171))) = uint8(v173)
	*(*uint16)(unsafe.Add(mBase, uint32(v42+v13+int32(7)))) = uint16(v168)
	v180 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L52
	}
L10:
	;
	F__serverAssert(m, int32(_a_F_addReplyDouble_0), int32(_a_F_addReplyDouble_1), int32(1332))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L51
	}
L11:
	;
	v44 = base.I64_extend_i32_s(v42)
	v45 = int32(0)
	if base.Ui64(v44) < base.Ui64(int64(10)) {
		v104 = v45
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if int32(5) <= v111 {
		goto L10
	} else {
		goto L44
	}
L13:
	;
	v111 = int32(1) + v104
	goto L12
L14:
	;
	v50 = v44
	v51 = v45
	goto L15
L15:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v50) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v104 = v98
	goto L13
L17:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v50) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v111 = int32(2) + v51
	goto L12
L19:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v50) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v111 = int32(3) + v51
	goto L12
L21:
	;
	v98 = v51 + int32(12)
	v102 = base.I64_div_u_s(v50, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v50) {
		v50 = v102
		v51 = v98
		goto L15
	} else {
		goto L43
	}
L22:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v50) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v50) {
		goto L35
	} else {
		goto L36
	}
L24:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v50) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v50) {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v50) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v50) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v111 = int32(4) + v51
	goto L12
L29:
	;
	v75 = int32(6)
	goto L31
L30:
	;
	v75 = int32(5)
	goto L31
L31:
	;
	v111 = v75 + v51
	goto L12
L32:
	;
	v81 = int32(8)
	goto L34
L33:
	;
	v81 = int32(7)
	goto L34
L34:
	;
	v111 = v81 + v51
	goto L12
L35:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v50) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v50) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v89 = int32(10)
	goto L39
L38:
	;
	v89 = int32(9)
	goto L39
L39:
	;
	v111 = v89 + v51
	goto L12
L40:
	;
	v95 = int32(12)
	goto L42
L41:
	;
	v95 = int32(11)
	goto L42
L42:
	;
	v111 = v95 + v51
	goto L12
L43:
	;
	goto L16
L44:
	;
	v115 = int32(4) - v111
	v116 = v13 + v115
	v117 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v117)
	if v42 == int32(0) {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	if v111 <= int32(0) {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v126 = v111
	v127 = v42
	goto L47
L47:
	;
	v135 = int32(10)
	v136 = base.I32_div_s(v127, v135)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127-v136*v135)+uint32(_c_F_addReplyDouble[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13+(v126+v115)))) = uint8(v142)
	if base.Ui32(int32(-20)) < base.Ui32(v127+int32(-10)) {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(int32(1)) < base.Ui32(v126) {
		v126 = v126 + int32(-1)
		v127 = v136
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L9
L51:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	if v180 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F__addReplyToBufferOrList(m, l0, v116, v171-v115)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L1
}
func F_addReplyDoubleDistance(m *base.Module, l0 int32, l1 float64) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = m.G0
	v5 = int32(128)
	v6 = v4 - v5
	m.G0 = v6
	v10 = F_fixedpoint_d2string(m, v6, v5, l1, int32(4))
	F_addReplyBulkCBuffer(m, l0, v6, v10)
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v6 + int32(128)
		return
	}
}
func F_addReplyErrorArity(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v8
	F_addReplyErrorFormat(m, l0, int32(_a_F_addReplyErrorArity_0), v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_addReplyErrorLength(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	if l2 == int32(0) {
		v9 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if v9 != 0 {
				v15 = F_prepareClientToWrite(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					if v15 != 0 {
						v19 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							if v19 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F__addReplyToBufferOrList(m, l0, l1, l2)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v19 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								if v19 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_1), int32(5))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						if v15 != 0 {
							v19 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								if v19 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F__addReplyToBufferOrList(m, l0, l1, l2)
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return
							} else {
								v19 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									if v19 != 0 {
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
										mBase = m.M
										v24 = m.ExcPending
										if v24 != 0 {
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
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v6 == int32(45) {
			v15 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				if v15 != 0 {
					v19 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						if v19 != 0 {
							return
						} else {
							F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, l1, l2)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							if v19 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			v9 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				if v9 != 0 {
					v15 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						if v15 != 0 {
							v19 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								if v19 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F__addReplyToBufferOrList(m, l0, l1, l2)
							mBase = m.M
							v18 = m.ExcPending
							if v18 != 0 {
								return
							} else {
								v19 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									if v19 != 0 {
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
										mBase = m.M
										v24 = m.ExcPending
										if v24 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_1), int32(5))
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return
					} else {
						v15 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v16 = m.ExcPending
						if v16 != 0 {
							return
						} else {
							if v15 != 0 {
								v19 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									if v19 != 0 {
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
										mBase = m.M
										v24 = m.ExcPending
										if v24 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								F__addReplyToBufferOrList(m, l0, l1, l2)
								mBase = m.M
								v18 = m.ExcPending
								if v18 != 0 {
									return
								} else {
									v19 = F_prepareClientToWrite(m, l0)
									mBase = m.M
									v20 = m.ExcPending
									if v20 != 0 {
										return
									} else {
										if v19 != 0 {
											return
										} else {
											F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyErrorLength_0), int32(2))
											mBase = m.M
											v24 = m.ExcPending
											if v24 != 0 {
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
func F_addReplyErrorObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	F_addReply(m, l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = F_objectGetVal(m, l1)
		mBase = m.M
		v10 = F_objectGetVal(m, l1)
		mBase = m.M
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-1)))))
		switch v13 & int32(7) {
		case 0:
			v30 = int32(base.Ui32(v13) >> (uint(int32(3)) % 32))
		case 1:
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
			v30 = v20
		case 2:
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-5)))))
			v30 = v23
		case 3:
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
			v30 = v26
		case 4:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-17))))
			v30 = v29
		default:
			v30 = int32(0)
		}
		F_afterErrorReply(m, l0, v8, v30+int32(-2), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			return
		}
	}
}
func F_addReplyErrorSds(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	F_addReplyErrorSdsEx(m, l0, l1, int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_addReplyErrorSdsEx(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v4 = int32(0)
	v11 = l1 + int32(-1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	switch v12 & int32(7) {
	case 0:
		v29 = int32(base.Ui32(v12) >> (uint(int32(3)) % 32))
	case 1:
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
		v29 = v19
	case 2:
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
		v29 = v22
	case 3:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
		v29 = v25
	case 4:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
		v29 = v28
	default:
		v29 = v4
	}
	F_addReplyErrorLength(m, l0, l1, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return
	} else {
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		switch v32 & int32(7) {
		case 0:
			v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
		case 1:
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
			v49 = v39
		case 2:
			v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
			v49 = v42
		case 3:
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
			v49 = v45
		case 4:
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
			v49 = v48
		default:
			v49 = v4
		}
		F_afterErrorReply(m, l0, l1, v49, l2)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return
		} else {
			F_sdsfree(m, l1)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_addReplyFlagsForCommand(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	v6 = base.I32_wrap_i64(v5)
	v7 = int32(1)
	F_addReplySetLen(m, l0, v6&v7+int32(base.Ui32(v6)>>(uint(v7)%32))&v7+int32(base.Ui32(v6)>>(uint(int32(2))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(3))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(4))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(5))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(6))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(8))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(9))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(10))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(11))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(12))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(13))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(14))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(15))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(19))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(23))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(24))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(25))%32))&v7+int32(base.Ui32(v6)>>(uint(int32(26))%32))&v7)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v109 = int32(0)
	v111 = int32(_a_F_addReplyFlagsForCommand_0)
	goto L3
L3:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v109<<(uint(int32(4))%32))+uint32(_c_F_addReplyFlagsForCommand[0])))
	if v116&v5 == int64(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v123 = v109 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(4))%32))+uint32(_c_F_addReplyFlagsForCommand[1])))
	if v123 != int32(20) {
		v109 = v123
		v111 = v128
		goto L3
	} else {
		goto L8
	}
L6:
	;
	F_addReplyStatus(m, l0, v111)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	goto L4
}
func F_addReplyFlagsForKeyArgs(m *base.Module, l0 int32, l1 int64) {
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v41 int64
	_ = v41
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	v13 = int64(1)
	v14 = l1 & v13
	v16 = int64(2)
	v17 = l1 & v16
	v22 = int64(4)
	v23 = l1 & v22
	v28 = int64(8)
	v29 = l1 & v28
	v35 = l1 & int64(16)
	v41 = l1 & int64(32)
	v47 = l1 & int64(64)
	v53 = l1 & int64(128)
	v59 = l1 & int64(256)
	v65 = l1 & int64(512)
	v71 = l1 & int64(1024)
	F_addReplySetLen(m, l0, base.I32_wrap_i64(v14)+base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(v13)%64)))+base.I32_wrap_i64(int64(base.Ui64(v23)>>(uint(v16)%64)))+base.I32_wrap_i64(int64(base.Ui64(v29)>>(uint(int64(3))%64)))+base.I32_wrap_i64(int64(base.Ui64(v35)>>(uint(v22)%64)))+base.I32_wrap_i64(int64(base.Ui64(v41)>>(uint(int64(5))%64)))+base.I32_wrap_i64(int64(base.Ui64(v47)>>(uint(int64(6))%64)))+base.I32_wrap_i64(int64(base.Ui64(v53)>>(uint(int64(7))%64)))+base.I32_wrap_i64(int64(base.Ui64(v59)>>(uint(v28)%64)))+base.I32_wrap_i64(int64(base.Ui64(v65)>>(uint(int64(9))%64)))+base.I32_wrap_i64(int64(base.Ui64(v71)>>(uint(int64(10))%64))))
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v14 == int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v17 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_7))
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	if v23 == int64(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_8))
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if v29 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_9))
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if v35 == int64(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_10))
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v41 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_6))
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	if v47 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_5))
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if v53 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_4))
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v59 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_3))
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v65 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_2))
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v71 == int64(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_1))
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	return
L34:
	;
	F_addReplyStatus(m, l0, int32(_a_F_addReplyFlagsForKeyArgs_0))
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L33
}
func F_addReplyNull(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v5 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v4 != int32(2) {
			if v5 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyNull_0), int32(3))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			if v5 != 0 {
				return
			} else {
				F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyNull_1), int32(5))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_addReplyPubsubMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v6 | int32(131072)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v10 != int32(2) {
		F_addReplyPushLen(m, l0, int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_addReply(m, l0, l3)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if l2 == int32(0) {
						if v6&int32(131072) != 0 {
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v30 & int32(-131073)
						}
						return
					} else {
						F_addReplyBulk(m, l0, l2)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v6&int32(131072) != 0 {
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v30 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_addReplyPubsubMessage[0]))
		F_addReply(m, l0, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_addReply(m, l0, l3)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_addReplyBulk(m, l0, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if l2 == int32(0) {
						if v6&int32(131072) != 0 {
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v30 & int32(-131073)
						}
						return
					} else {
						F_addReplyBulk(m, l0, l2)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v6&int32(131072) != 0 {
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v30 & int32(-131073)
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_addReplySds(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v5 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v5 != 0 {
			F_sdsfree(m, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				return
			}
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
			switch v10 & int32(7) {
			case 0:
				v27 = int32(base.Ui32(v10) >> (uint(int32(3)) % 32))
			case 1:
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
				v27 = v17
			case 2:
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
				v27 = v20
			case 3:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
				v27 = v23
			case 4:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
				v27 = v26
			default:
				v27 = int32(0)
			}
			F__addReplyToBufferOrList(m, l0, l1, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_sdsfree(m, l1)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_addReplySetLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	if l1 <= int32(-1) {
		F__serverAssert(m, int32(_a_F_addReplySetLen_0), int32(_a_F_addReplySetLen_1), int32(1420))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
		v7 = F_prepareClientToWrite(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			if v7 != 0 {
				return
			} else {
				if v6&int32(255) == int32(2) {
					v16 = int32(42)
				} else {
					v16 = int32(126)
				}
				F__addReplyLongLongWithPrefix(m, l0, base.I64_extend_i32_u(l1), v16)
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
func F_addReplyStatusLength(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = F_prepareClientToWrite(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			v10 = F_prepareClientToWrite(m, l0)
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				if v10 != 0 {
					v14 = F_prepareClientToWrite(m, l0)
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						if v14 != 0 {
							return
						} else {
							F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyStatusLength_0), int32(2))
							v19 = m.ExcPending
							if v19 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, l1, l2)
					v13 = m.ExcPending
					if v13 != 0 {
						return
					} else {
						v14 = F_prepareClientToWrite(m, l0)
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							if v14 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyStatusLength_0), int32(2))
								v19 = m.ExcPending
								if v19 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		} else {
			F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyStatusLength_1), int32(1))
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v10 = F_prepareClientToWrite(m, l0)
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					if v10 != 0 {
						v14 = F_prepareClientToWrite(m, l0)
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							if v14 != 0 {
								return
							} else {
								F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyStatusLength_0), int32(2))
								v19 = m.ExcPending
								if v19 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F__addReplyToBufferOrList(m, l0, l1, l2)
						v13 = m.ExcPending
						if v13 != 0 {
							return
						} else {
							v14 = F_prepareClientToWrite(m, l0)
							v15 = m.ExcPending
							if v15 != 0 {
								return
							} else {
								if v14 != 0 {
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyStatusLength_0), int32(2))
									v19 = m.ExcPending
									if v19 != 0 {
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
func F_addReplyVerbatim(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v15 != int32(2) {
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2 + int32(4)
		v24 = v13 + int32(16)
		v29 = F_snprintf(m, v24, int32(32), int32(_a_F_addReplyVerbatim_0), v13)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			v31 = v24 + v29
			v33 = v31 + int32(-4)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
			if v34 != 0 {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v34)
				if v37 != 0 {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)))
					*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-3)))) = uint8(v37)
					if v44 != 0 {
						v51 = v44
					} else {
						v51 = int32(32)
					}
				} else {
					v42 = int32(32)
					*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-3)))) = uint8(v42)
					v51 = int32(32)
				}
			} else {
				v35 = int32(32)
				*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v35)
				v42 = int32(32)
				*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-3)))) = uint8(v42)
				v51 = int32(32)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v31+int32(-2)))) = uint8(v51)
			v56 = F_prepareClientToWrite(m, l0)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				if v56 != 0 {
					v62 = F_prepareClientToWrite(m, l0)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						if v62 != 0 {
							v66 = F_prepareClientToWrite(m, l0)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								if v66 != 0 {
									m.G0 = v13 + int32(48)
									return
								} else {
									F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyVerbatim_1), int32(2))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										m.G0 = v13 + int32(48)
										return
									}
								}
							}
						} else {
							F__addReplyToBufferOrList(m, l0, l1, l2)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v66 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									if v66 != 0 {
										m.G0 = v13 + int32(48)
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyVerbatim_1), int32(2))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											m.G0 = v13 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F__addReplyToBufferOrList(m, l0, v13+int32(16), v29)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = F_prepareClientToWrite(m, l0)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v62 != 0 {
								v66 = F_prepareClientToWrite(m, l0)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									if v66 != 0 {
										m.G0 = v13 + int32(48)
										return
									} else {
										F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyVerbatim_1), int32(2))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											m.G0 = v13 + int32(48)
											return
										}
									}
								}
							} else {
								F__addReplyToBufferOrList(m, l0, l1, l2)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v66 = F_prepareClientToWrite(m, l0)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										if v66 != 0 {
											m.G0 = v13 + int32(48)
											return
										} else {
											F__addReplyToBufferOrList(m, l0, int32(_a_F_addReplyVerbatim_1), int32(2))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												m.G0 = v13 + int32(48)
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
	} else {
		F_addReplyBulkCBuffer(m, l0, l1, l2)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v13 + int32(48)
			return
		}
	}
}
func F_callReplyBigNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v12 | int32(4)
	return
}
func F_callReplyCreatePromise(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v4 = F_valkey_malloc(m, int32(48))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v11 | int32(3)
		return v4
	}
}
func F_callReplyDeferredErrorList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return v2
}
func F_callReplyError(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(1)
	return
}
func F_callReplyGetArrayElement(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v63 = int32(0)
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v64 != int32(3) {
			v73 = v63
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v67) <= base.Ui32(l1) {
				v73 = v63
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v73 = v69 + l1*int32(48)
			}
		}
		m.G0 = v7 + int32(80)
		return v73
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetArrayElement[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v63 = int32(0)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v64 != int32(3) {
				v73 = v63
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if base.Ui32(v67) <= base.Ui32(l1) {
					v73 = v63
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v73 = v69 + l1*int32(48)
				}
			}
			m.G0 = v7 + int32(80)
			return v73
		}
	}
}
func F_callReplyGetAttributeElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v13&int32(2) != 0 {
		v67 = int32(-1)
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v68 != int32(5) {
			v107 = v67
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v71) <= base.Ui32(l1) {
				v107 = v67
			} else {
				if l2 == int32(0) {
				} else {
					v76 = int32(1)
					v77 = l1 << (uint(v76) % 32)
					if base.Ui32(v71<<(uint(v76)%32)) <= base.Ui32(v77) {
						v85 = int32(0)
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v85 = v81 + v77*int32(48)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v85
				}
				if l3 != 0 {
					v90 = int32(0)
					v92 = int32(1)
					v95 = l1<<(uint(v92)%32) | v92
					if base.Ui32(v71<<(uint(v92)%32)) <= base.Ui32(v95) {
						v103 = v90
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v103 = v99 + v95*int32(48)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v103
					v107 = v90
				} else {
					v107 = int32(0)
				}
			}
		}
		m.G0 = v11 + int32(80)
		return v107
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v19 = int32(0)
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)))) = v20
		v25 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(32)))) = v25
		v30 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(40)))) = v30
		v35 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v35
		v40 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(56)))) = v40
		v45 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(64)))) = v45
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v11+int32(72)))) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v16
		v54 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetAttributeElement[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v54
		v58 = F_parseReply(m, v11+int32(12), l0)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 | int32(2)
			v67 = int32(-1)
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v68 != int32(5) {
				v107 = v67
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if base.Ui32(v71) <= base.Ui32(l1) {
					v107 = v67
				} else {
					if l2 == int32(0) {
					} else {
						v76 = int32(1)
						v77 = l1 << (uint(v76) % 32)
						if base.Ui32(v71<<(uint(v76)%32)) <= base.Ui32(v77) {
							v85 = int32(0)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v85 = v81 + v77*int32(48)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v85
					}
					if l3 != 0 {
						v90 = int32(0)
						v92 = int32(1)
						v95 = l1<<(uint(v92)%32) | v92
						if base.Ui32(v71<<(uint(v92)%32)) <= base.Ui32(v95) {
							v103 = v90
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v103 = v99 + v95*int32(48)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v103
						v107 = v90
					} else {
						v107 = int32(0)
					}
				}
			}
			m.G0 = v11 + int32(80)
			return v107
		}
	}
}
func F_callReplyGetBool(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v8&int32(2) != 0 {
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v63 != int32(7) {
			v67 = int32(-2147483648)
		} else {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v67 = v66
		}
		m.G0 = v6 + int32(80)
		return v67
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = int32(0)
		v15 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v15
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v20
		v25 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(40)))) = v25
		v30 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v30
		v35 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v35
		v40 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(64)))) = v40
		v45 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(72)))) = v45
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v11
		v49 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetBool[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v49
		v53 = F_parseReply(m, v6+int32(12), l0)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v57 | int32(2)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v63 != int32(7) {
				v67 = int32(-2147483648)
			} else {
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v67 = v66
			}
			m.G0 = v6 + int32(80)
			return v67
		}
	}
}
func F_callReplyGetDouble(m *base.Module, l0 int32) float64 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v64 != int32(8) {
			v68 = float64(-9.223372036854776e+18)
		} else {
			v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
			v68 = v67
		}
		m.G0 = v7 + int32(80)
		return v68
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetDouble[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return float64(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v64 != int32(8) {
				v68 = float64(-9.223372036854776e+18)
			} else {
				v67 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
				v68 = v67
			}
			m.G0 = v7 + int32(80)
			return v68
		}
	}
}
func F_callReplyGetLen(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v63 = int32(0)
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(int32(11)) < base.Ui32(v64) {
			v74 = v63
		} else {
			if int32(1)<<(uint(v64)%32)&int32(2155) == int32(0) {
				v74 = v63
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v74 = v73
			}
		}
		m.G0 = v7 + int32(80)
		return v74
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetLen[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v63 = int32(0)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if base.Ui32(int32(11)) < base.Ui32(v64) {
				v74 = v63
			} else {
				if int32(1)<<(uint(v64)%32)&int32(2155) == int32(0) {
					v74 = v63
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v74 = v73
				}
			}
			m.G0 = v7 + int32(80)
			return v74
		}
	}
}
func F_callReplyGetPrivateData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_callReplyGetSetElement(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9&int32(2) != 0 {
		v63 = int32(0)
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v64 != int32(6) {
			v73 = v63
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v67) <= base.Ui32(l1) {
				v73 = v63
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v73 = v69 + l1*int32(48)
			}
		}
		m.G0 = v7 + int32(80)
		return v73
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v16
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = v21
		v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v26
		v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(48)))) = v31
		v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(56)))) = v36
		v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(64)))) = v41
		v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v46
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v12
		v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyGetSetElement[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v50
		v54 = F_parseReply(m, v7+int32(12), l0)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
			v63 = int32(0)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v64 != int32(6) {
				v73 = v63
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if base.Ui32(v67) <= base.Ui32(l1) {
					v73 = v63
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v73 = v69 + l1*int32(48)
				}
			}
			m.G0 = v7 + int32(80)
			return v73
		}
	}
}
func F_callReplyNullBulkString(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(4)
	return
}
func F_callReplyParseError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	return
}
func F_callReplySet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(6)
	v13 = F_valkey_calloc(m, l2*int32(48))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v13
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v64 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v63 - l3
	return
L4:
	;
	v23 = v13
	v24 = int32(0)
	goto L5
L5:
	;
	v27 = v24 * int32(48)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+v27))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v33 = F_parseReply(m, l0, v31+v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v36 = v35 + v27
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v37 | int32(2)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v27)+20)))
	if v43&int32(4) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = v24 + int32(1)
	if v53 != l2 {
		v23 = v41
		v24 = v53
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v48 | int32(4)
	goto L8
L10:
	;
	goto L6
}
func F_callReplySimpleStr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10 << (uint(int32(28)) % 32) >> (uint(int32(31)) % 32) & int32(13)
	return
}
func F_callReplyType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int64
	_ = v21
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	if l0 != 0 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v9&int32(2) != 0 {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v64 = v63
			m.G0 = v6 + int32(80)
			return v64
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = int32(0)
			v16 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v16
			v21 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v21
			v26 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[2]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(40)))) = v26
			v31 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[3]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v31
			v36 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[4]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v36
			v41 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[5]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(64)))) = v41
			v46 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[6]))
			*(*int64)(unsafe.Add(mBase, uint32(v6+int32(72)))) = v46
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v12
			v50 = *(*int64)(unsafe.Add(mBase, _c_F_callReplyType[7]))
			*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v50
			v54 = F_parseReply(m, v6+int32(12), l0)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 | int32(2)
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v64 = v63
				m.G0 = v6 + int32(80)
				return v64
			}
		}
	} else {
		v64 = int32(-1)
		m.G0 = v6 + int32(80)
		return v64
	}
}
func F_callReplyVerbatimString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v14 | int32(4)
	return
}
func F_freeCallReplyInternal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	switch v4 + int32(-3) {
	case 0, 3:
		goto L2
	default:
		v32 = v4
		goto L1
	}
L1:
	;
	switch v32 + int32(-5) {
	case 0, 6:
		goto L12
	default:
		goto L11
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v7 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_valkey_free(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L4:
	;
	v12 = int32(0)
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_freeCallReplyInternal(m, v14+v12*int32(48))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	return
L8:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v21) < base.Ui32(v22) {
		v12 = v21
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v32 = v30
	goto L1
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v68 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_valkey_free(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L20
	}
L14:
	;
	v41 = int32(0)
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v45 = v41 * int32(96)
	F_freeCallReplyInternal(m, v43+v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L13
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_freeCallReplyInternal(m, v49+v45+int32(48))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v56 = v41 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(v56) < base.Ui32(v57) {
		v41 = v56
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	goto L11
L21:
	;
	return
L22:
	;
	F_freeCallReplyInternal(m, v68)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_valkey_free(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	goto L21
}
func F_replyHandlersBool(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	if v7 == int32(0) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v7].(func(*base.Module, int32, int32))(m, v10, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_replyHandlersMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	m.T0[v7].(func(*base.Module, int32, int32))(m, v10, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	goto L1
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v20 = int32(0)
	goto L7
L7:
	;
	v21 = F_parseReply(m, l0, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v23 = F_parseReply(m, l0, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v26 = v20 + int32(1)
	if v26 != l2 {
		v20 = v26
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	m.T0[v34].(func(*base.Module, int32))(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_replyHandlersSimpleStr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v8 == int32(0) {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		m.T0[v8].(func(*base.Module, int32, int32, int32))(m, v11, l1, l2)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
