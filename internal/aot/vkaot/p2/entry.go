package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_entryGetExpiryVsetFunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	v5 = int64(-1)
	v7 = l0 + int32(-1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8&int32(7) == int32(0) {
		v43 = v5
	} else {
		if v8&int32(8) == int32(0) {
			v43 = v5
		} else {
			v17 = F_sdsAllocPtr(m, l0)
			mBase = m.M
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if int32(base.Ui32(v20&int32(16))>>(uint(int32(4))%32)) != 0 {
				v25 = int32(-4)
			} else {
				v25 = int32(0)
			}
			v28 = v20 & int32(7)
			if v28 != 0 {
				v29 = v25
			} else {
				v29 = int32(0)
			}
			if int32(base.Ui32(v20&int32(8))>>(uint(int32(3))%32)) != 0 {
				v37 = int32(-8)
			} else {
				v37 = int32(0)
			}
			if v28 != 0 {
				v39 = v37
			} else {
				v39 = int32(0)
			}
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v17+v29+v39)))
			v43 = v41
		}
	}
	return v43
}
func F_entryGetField(m *base.Module, l0 int32) int32 {
	return l0
}
func F_entrySetExpiry(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v7 = l0 + int32(-1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v10 = v8 & int32(7)
	if l1 == int64(-1) {
		if v10 == int32(0) {
			v95 = F_entryUpdate(m, l0, int32(0), l1)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				return v95
			}
		} else {
			v63 = int32(48)
			if v8&v63 != v63 {
				v95 = F_entryUpdate(m, l0, int32(0), l1)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					return v95
				}
			} else {
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
				v74 = v72 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v74) {
					v82 = int32(0)
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v74<<(uint(int32(2))%32))+uint32(_consts[412])))
					v82 = v81
				}
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0+v82+int32(-4))))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
				v89 = F_entryUpdateAsStringRef(m, l0, v87, v88, l1)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					return v89
				}
			}
		}
	} else {
		if v10 == int32(0) {
			if v10 == int32(0) {
				v95 = F_entryUpdate(m, l0, int32(0), l1)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					return v95
				}
			} else {
				v63 = int32(48)
				if v8&v63 != v63 {
					v95 = F_entryUpdate(m, l0, int32(0), l1)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						return v95
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
					v74 = v72 & int32(7)
					if base.Ui32(int32(4)) < base.Ui32(v74) {
						v82 = int32(0)
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v74<<(uint(int32(2))%32))+uint32(_consts[412])))
						v82 = v81
					}
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0+v82+int32(-4))))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v89 = F_entryUpdateAsStringRef(m, l0, v87, v88, l1)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						return v89
					}
				}
			}
		} else {
			if v8&int32(8) == int32(0) {
				if v10 == int32(0) {
					v95 = F_entryUpdate(m, l0, int32(0), l1)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						return v95
					}
				} else {
					v63 = int32(48)
					if v8&v63 != v63 {
						v95 = F_entryUpdate(m, l0, int32(0), l1)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return v95
						}
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
						v74 = v72 & int32(7)
						if base.Ui32(int32(4)) < base.Ui32(v74) {
							v82 = int32(0)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v74<<(uint(int32(2))%32))+uint32(_consts[412])))
							v82 = v81
						}
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0+v82+int32(-4))))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
						v89 = F_entryUpdateAsStringRef(m, l0, v87, v88, l1)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							return v89
						}
					}
				}
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
				v26 = v24 & int32(7)
				if base.Ui32(int32(4)) < base.Ui32(v26) {
					v34 = int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v26<<(uint(int32(2))%32))+uint32(_consts[412])))
					v34 = v33
				}
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
				if int32(base.Ui32(v38&int32(16))>>(uint(int32(4))%32)) != 0 {
					v43 = int32(-4)
				} else {
					v43 = int32(0)
				}
				v46 = v38 & int32(7)
				if v46 != 0 {
					v47 = v43
				} else {
					v47 = int32(0)
				}
				if int32(base.Ui32(v38&int32(8))>>(uint(int32(3))%32)) != 0 {
					v55 = int32(-8)
				} else {
					v55 = int32(0)
				}
				if v46 != 0 {
					v57 = v55
				} else {
					v57 = int32(0)
				}
				*(*int64)(unsafe.Add(mBase, uint32(l0+v34+v47+v57))) = l1
				return l0
			}
		}
	}
}
func F_removeEntryFromRaxBucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	v7 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v7)
	switch l2 + int32(1) {
	case 0:
		F__serverPanic_1(m, int32(_a1861), int32(1422), int32(_a1866), int32(0))
		mBase = m.M
		v123 = m.ExcPending
		if v123 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 1:
		F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	default:
		if l2&int32(1) != 0 {
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(base.B2i32(l2 == l1))
			if l2 != l1 {
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
				m.G0 = v10 + int32(16)
				return v158
			} else {
				switch l0 + int32(1) {
				case 0:
					F__serverAssert(m, int32(_a1867), int32(_a1861), int32(807))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 1:
					F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				default:
					if l0&int32(7) != int32(6) {
						F__serverAssert(m, int32(_a1867), int32(_a1861), int32(807))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v48 = F_raxRemove(m, l0&int32(-8), l3, l4, int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							m.G0 = v10 + int32(16)
							return v158
						}
					}
				}
			}
		} else {
			switch l2&int32(6) + int32(-2) {
			case 0:
				v53 = F_removeFromBucket_VECTOR(m, l2, l1, v10+int32(15), int32(1))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					if v53 == l2 {
						v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						m.G0 = v10 + int32(16)
						return v158
					} else {
						switch v53 + int32(1) {
						case 0:
							switch l0 + int32(1) {
							case 0:
								F__serverAssert(m, int32(_a1867), int32(_a1861), int32(807))
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							case 1:
								F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							default:
								if l0&int32(7) != int32(6) {
									F__serverAssert(m, int32(_a1867), int32(_a1861), int32(807))
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v79 = F_raxRemove(m, l0&int32(-8), l3, l4, int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										m.G0 = v10 + int32(16)
										return v158
									}
								}
							}
						case 1:
							F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						default:
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
							if v53 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v127 | int32(3)
							} else {
								v130 = int32(3)
								v131 = int32(base.Ui32(v127) >> (uint(v130) % 32))
								v138 = int32(4)
								if v127&v138 != 0 {
									v143 = v138
								} else {
									v143 = v131 << (uint(int32(2)) % 32)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l5+v131+(int32(0)-v131)&v130+v143+int32(4)))) = v53
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v127&int32(-4) | int32(1)
							}
							v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							m.G0 = v10 + int32(16)
							return v158
						}
					}
				}
			default:
				F__serverPanic_1(m, int32(_a1861), int32(1422), int32(_a1866), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 2:
				v83 = F_removeFromBucket_HASHTABLE(m, l2, l1, v10+int32(15))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					if v83 == l2 {
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						if v83 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v88 | int32(3)
						} else {
							v91 = int32(3)
							v92 = int32(base.Ui32(v88) >> (uint(v91) % 32))
							v99 = int32(4)
							if v88&v99 != 0 {
								v104 = v99
							} else {
								v104 = v92 << (uint(int32(2)) % 32)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l5+v92+(int32(0)-v92)&v91+v104+int32(4)))) = v83
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v88&int32(-4) | int32(1)
						}
					}
					v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					m.G0 = v10 + int32(16)
					return v158
				}
			}
		}
	}
}
