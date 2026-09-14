package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F__ziplistPairsEntryConvertAndValidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v10 != 0 {
		v22 = int32(0)
		v27 = F_ziplistGet(m, l0, v8+int32(12), v8+int32(8), v8)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v27 == int32(0) {
				v70 = v22
				m.G0 = v8 + int32(16)
				return v70
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v31&int32(1) != 0 {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					if v52 == int32(0) {
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
						v59 = F_lpAppendInteger(m, v51, v58)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = v59
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
							v64 = int32(1)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
							v70 = v64
							m.G0 = v8 + int32(16)
							return v70
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						v56 = F_lpAppend(m, v51, v52, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v61 = v56
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
							v64 = int32(1)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
							v70 = v64
							m.G0 = v8 + int32(16)
							return v70
						}
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					if v34 == int32(0) {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
						v41 = F_sdsfromlonglong(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = v41
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v45 = F_hashtableAdd(m, v44, v43)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								if v45 != 0 {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									if v52 == int32(0) {
										v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
										v59 = F_lpAppendInteger(m, v51, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = v59
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
											v64 = int32(1)
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
											v70 = v64
											m.G0 = v8 + int32(16)
											return v70
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										v56 = F_lpAppend(m, v51, v52, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v61 = v56
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
											v64 = int32(1)
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
											v70 = v64
											m.G0 = v8 + int32(16)
											return v70
										}
									}
								} else {
									F_sdsfree(m, v43)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v70 = v22
										m.G0 = v8 + int32(16)
										return v70
									}
								}
							}
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						v38 = F_sdsnewlen(m, v34, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v43 = v38
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v45 = F_hashtableAdd(m, v44, v43)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								if v45 != 0 {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									if v52 == int32(0) {
										v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
										v59 = F_lpAppendInteger(m, v51, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v61 = v59
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
											v64 = int32(1)
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
											v70 = v64
											m.G0 = v8 + int32(16)
											return v70
										}
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										v56 = F_lpAppend(m, v51, v52, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v61 = v56
											v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
											v64 = int32(1)
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
											v70 = v64
											m.G0 = v8 + int32(16)
											return v70
										}
									}
								} else {
									F_sdsfree(m, v43)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v70 = v22
										m.G0 = v8 + int32(16)
										return v70
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = F_hashtableCreate(m, int32(_a_F__ziplistPairsEntryConvertAndValidate_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v12
			v19 = F_hashtableExpand(m, v12, int32(base.Ui32(l1)>>(uint(int32(1))%32)))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = int32(0)
				v27 = F_ziplistGet(m, l0, v8+int32(12), v8+int32(8), v8)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v70 = v22
						m.G0 = v8 + int32(16)
						return v70
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
						if v31&int32(1) != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							if v52 == int32(0) {
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
								v59 = F_lpAppendInteger(m, v51, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = v59
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
									v64 = int32(1)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
									v70 = v64
									m.G0 = v8 + int32(16)
									return v70
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								v56 = F_lpAppend(m, v51, v52, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v61 = v56
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
									v64 = int32(1)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
									v70 = v64
									m.G0 = v8 + int32(16)
									return v70
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							if v34 == int32(0) {
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
								v41 = F_sdsfromlonglong(m, v40)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = v41
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v45 = F_hashtableAdd(m, v44, v43)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										if v45 != 0 {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											if v52 == int32(0) {
												v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
												v59 = F_lpAppendInteger(m, v51, v58)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v61 = v59
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
													v64 = int32(1)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
													v70 = v64
													m.G0 = v8 + int32(16)
													return v70
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v56 = F_lpAppend(m, v51, v52, v55)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v61 = v56
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
													v64 = int32(1)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
													v70 = v64
													m.G0 = v8 + int32(16)
													return v70
												}
											}
										} else {
											F_sdsfree(m, v43)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												v70 = v22
												m.G0 = v8 + int32(16)
												return v70
											}
										}
									}
								}
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
								v38 = F_sdsnewlen(m, v34, v37)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v43 = v38
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v45 = F_hashtableAdd(m, v44, v43)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										if v45 != 0 {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											if v52 == int32(0) {
												v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
												v59 = F_lpAppendInteger(m, v51, v58)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v61 = v59
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
													v64 = int32(1)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
													v70 = v64
													m.G0 = v8 + int32(16)
													return v70
												}
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v56 = F_lpAppend(m, v51, v52, v55)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v61 = v56
													v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v62))) = v61
													v64 = int32(1)
													v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + v64
													v70 = v64
													m.G0 = v8 + int32(16)
													return v70
												}
											}
										} else {
											F_sdsfree(m, v43)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int32(0)
											} else {
												v70 = v22
												m.G0 = v8 + int32(16)
												return v70
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
func F_zcalloc_usable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v80 = int32(0)
		v82 = *(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[0]))
		m.T0[v82].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return int32(0)
		} else {
			v89 = int32(0)
			v91 = v80
			if l1 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v89
			}
			return v91
		}
	} else {
		v8 = int32(1)
		if l0 != 0 {
			v10 = l0
		} else {
			v10 = int32(4)
		}
		v12 = v10 + int32(8)
		v18 = base.I64_extend_i32_u(v8) * base.I64_extend_i32_u(v12)
		v19 = base.I32_wrap_i64(v18)
		if base.Ui32(v12|v8) < base.Ui32(int32(65536)) {
			v30 = v19
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) != int32(0) {
				v29 = int32(-1)
			} else {
				v29 = v19
			}
			v30 = v29
		}
		v32 = F_emscripten_builtin_malloc(m, v30)
		mBase = m.M
		if v32 == int32(0) {
		} else {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+int32(-4)))))
			if v37&int32(3) == int32(0) {
			} else {
				v43 = F___memset(m, v32, int32(0), v30)
				mBase = m.M
			}
		}
		if v32 == int32(0) {
			v80 = int32(0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[0]))
			m.T0[v82].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				v89 = int32(0)
				v91 = v80
				if l1 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v89
				}
				return v91
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v32))) = v10
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[1]))
			if v48 != int32(-1) {
				v59 = v48
			} else {
				v51 = int32(0)
				v53 = *(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[1])) = v53
				*(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[2])) = v53 + int32(1)
				v59 = v53
			}
			if v59 < int32(260) {
				v68 = v59 << (uint(int32(2)) % 32)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_zcalloc_usable[3])))
				*(*int32)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_zcalloc_usable[3]))) = v71 + v12
			} else {
				v62 = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[4]))
				*(*int32)(unsafe.Add(mBase, _c_F_zcalloc_usable[4])) = v64 + v12
			}
			v89 = v10
			v91 = v32 + int32(8)
			if l1 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v89
			}
			return v91
		}
	}
}
func F_zdiffCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v2 = int32(0)
	v3 = int32(1)
	F_zunionInterDiffGenericCommand(m, l0, v2, v3, v3, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_zincrbyCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zaddGenericCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zinterCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(0)
	F_zunionInterDiffGenericCommand(m, l0, v2, int32(1), int32(2), v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_zinterstoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = int32(2)
	F_zunionInterDiffGenericCommand(m, l0, v3, v4, v4, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_zipEntry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v9) < base.Ui32(int32(254)) {
		v12 = int32(1)
	} else {
		v12 = int32(5)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
	if base.Ui32(int32(253)) < base.Ui32(v9) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(1))))
		v20 = v19
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v20 = v16
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v20
	v22 = l0 + v12
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v24 = int32(192)
	if base.Ui32(v23) < base.Ui32(v24) {
		v28 = v23 & v24
	} else {
		v28 = v23
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v28)
	if base.Ui32(int32(191)) < base.Ui32(v23) {
		v69 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v69
		switch v23 + int32(-208) {
		case 0:
			v113 = int32(4)
			v114 = v69
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
			v118 = v114
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
			return
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
			v86 = int32(1)
			if base.Ui32(int32(241)) < base.Ui32((v23+v86)&int32(255)) {
				v118 = v86
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
				return
			} else {
				v95 = l1 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
				F__serverAssert(m, int32(_a_F_zipEntry_0), int32(_a_F_zipEntry_1), int32(619))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		case 16:
			v113 = int32(8)
			v114 = v69
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
			v118 = v114
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
			return
		default:
			switch v23 + int32(-240) {
			case 0:
				v113 = int32(3)
				v114 = v69
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
				v118 = v114
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
				return
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
				v86 = int32(1)
				if base.Ui32(int32(241)) < base.Ui32((v23+v86)&int32(255)) {
					v118 = v86
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
					return
				} else {
					v95 = l1 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
					F__serverAssert(m, int32(_a_F_zipEntry_0), int32(_a_F_zipEntry_1), int32(619))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			case 14:
				v113 = int32(1)
				v114 = v69
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
				v118 = v114
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
				return
			default:
				if v23 != int32(192) {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
					v86 = int32(1)
					if base.Ui32(int32(241)) < base.Ui32((v23+v86)&int32(255)) {
						v118 = v86
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
						return
					} else {
						v95 = l1 + int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
						F__serverAssert(m, int32(_a_F_zipEntry_0), int32(_a_F_zipEntry_1), int32(619))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					v113 = int32(2)
					v114 = v69
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
					v118 = v114
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
					return
				}
			}
		}
	} else {
		v32 = int32(2)
		switch int32(base.Ui32(v23)>>(uint(int32(6))%32)) ^ v32 {
		default:
			v46 = int32(5)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v46
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+1))
			v50 = int32(24)
			v52 = int32(65280)
			v54 = int32(8)
			v113 = v49<<(uint(v50)%32) | v49&v52<<(uint(v54)%32) | (int32(base.Ui32(v49)>>(uint(v54)%32))&v52 | int32(base.Ui32(v49)>>(uint(v50)%32)))
			v114 = v46
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
			v118 = v114
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
			return
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
			v95 = l1 + int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
			F__serverAssert(m, int32(_a_F_zipEntry_0), int32(_a_F_zipEntry_1), int32(619))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 2:
			v106 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v106
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			v113 = v109 & int32(63)
			v114 = v106
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
			v118 = v114
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
			return
		case 3:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
			v113 = v39&int32(63)<<(uint(int32(8))%32) | v44
			v114 = v32
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v113
			v118 = v114
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v118 + v12
			return
		}
	}
}
func F_zlexcountCommand(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
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
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v18 = F_zsetParseLexRange(m, v14, v15, v10+int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverPanic_1(m, int32(_a_F_zlexcountCommand_0), int32(3480), int32(_a_F_zlexcountCommand_1), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L66
	}
L2:
	;
	F__serverAssertWithInfo(m, l0, v27, int32(_a_F_zlexcountCommand_2), int32(_a_F_zlexcountCommand_0), int32(3444))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L65
	}
L3:
	;
	m.G0 = v10 + int32(32)
	return
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[0]))
	v27 = F_lookupKeyReadOrReply(m, l0, v13, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L11
	}
L5:
	;
	return
L6:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_addReplyError(m, l0, int32(_a_F_zlexcountCommand_3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	switch int32(base.Ui32(v56)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L23
	default:
		goto L1
	case 4:
		goto L24
	}
L10:
	;
	v36 = int32(_a_F_zlexcountCommand_4)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[1]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[2]))
	if v38 == v40 {
		v49 = v40
		v50 = v37
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v27 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = F_checkType(m, l0, v27, int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v32 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v51 == v49 {
		goto L3
	} else {
		goto L19
	}
L16:
	;
	if v38 == v37 {
		v49 = v40
		v50 = v37
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_sdsfree(m, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v45 = int32(_a_F_zlexcountCommand_4)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[1]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[2]))
	v49 = v48
	v50 = v46
	goto L15
L19:
	;
	if v51 == v50 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	F_sdsfree(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v211 = int32(_a_F_zlexcountCommand_4)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[1]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[2]))
	if v213 == v215 {
		v224 = v212
		v225 = v215
		goto L56
	} else {
		goto L57
	}
L23:
	;
	v109 = int32(0)
	v110 = F_objectGetVal(m, v27)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v115 = F_zslNthInLexRange(m, v111, v10+int32(16), v109)
	mBase = m.M
	if v115 == v109 {
		v207 = v109
		goto L22
	} else {
		goto L39
	}
L24:
	;
	v63 = F_objectGetVal(m, v27)
	mBase = m.M
	v66 = F_zzlFirstInLexRange(m, v63, v10+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v66
	if v66 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = F_lpNext(m, v63, v66)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L30
	}
L27:
	;
	F_zsetFreeLexRange(m, v10+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[0]))
	F_addReply(m, l0, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L3
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v77
	v82 = F_zzlLexValueLteMax(m, v66, v10+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	if v82 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v89 = v66
	v90 = int32(0)
	goto L33
L33:
	;
	v96 = F_zzlLexValueLteMax(m, v89, v10+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v96 == int32(0) {
		v207 = v90
		goto L22
	} else {
		goto L36
	}
L36:
	;
	F_zzlNext(m, v63, v10+int32(12), v10+int32(8))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v107 = v90 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v108 != 0 {
		v89 = v108
		v90 = v107
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v207 = v107
	goto L22
L39:
	;
	v120 = v115
	v121 = v109
	goto L40
L40:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v127 = v125 + int32(-1)
	if v127 < int32(1) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v151 = v149 + int32(1)
	v155 = F_zslNthInLexRange(m, v111, v10+int32(16), int32(-1))
	mBase = m.M
	if v155 == int32(0) {
		v207 = v151
		goto L22
	} else {
		goto L46
	}
L42:
	;
	v149 = v147 + v121
	if v146 != 0 {
		v120 = v146
		v121 = v149
		goto L40
	} else {
		goto L45
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v120+v127<<(uint(int32(3))%32)+int32(12))))
	v146 = v143
	v147 = base.B2i32(v143 != int32(0))
	goto L42
L44:
	;
	v133 = v127 << (uint(int32(3)) % 32)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v120+v133)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v120+int32(16)+v133)))
	v146 = v135
	v147 = v137
	goto L42
L45:
	;
	goto L41
L46:
	;
	v158 = int32(0)
	if v155 == v158 {
		v196 = v158
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v207 = v199 - v196 + v151 - v202
	goto L22
L48:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	goto L47
L49:
	;
	v165 = v155
	v166 = v158
	goto L50
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v171 = v169 + int32(-1)
	if v171 < int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v196 = v193
	goto L48
L52:
	;
	v193 = v191 + v166
	if v190 != 0 {
		v165 = v190
		v166 = v193
		goto L50
	} else {
		goto L55
	}
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v165+v171<<(uint(int32(3))%32)+int32(12))))
	v190 = v187
	v191 = base.B2i32(v187 != int32(0))
	goto L52
L54:
	;
	v177 = v171 << (uint(int32(3)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165+v177)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v165+int32(16)+v177)))
	v190 = v179
	v191 = v181
	goto L52
L55:
	;
	goto L51
L56:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v226 == v225 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v213 == v212 {
		v224 = v212
		v225 = v215
		goto L56
	} else {
		goto L58
	}
L58:
	;
	F_sdsfree(m, v213)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v220 = int32(_a_F_zlexcountCommand_4)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[1]))
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_zlexcountCommand[2]))
	v224 = v221
	v225 = v223
	goto L56
L60:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v207))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L64
	}
L61:
	;
	if v226 == v224 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	F_sdsfree(m, v226)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L3
L65:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zmpopCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_zmpopGenericCommand(m, l0, int32(1), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_zpopminCommand(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 < int32(4) {
		v16 = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
		if v9 != int32(3) {
			v29 = v16
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v34 = int32(0)
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
			F_genericZpopCommand(m, l0, v30+int32(4), int32(1), v34, v34, v29, base.B2i32(v29 != int32(-1))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v38)), v34, v34)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			v26 = F_getPositiveLongFromObjectOrReply(m, l0, v22, v7+int32(12), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if v26 != 0 {
					m.G0 = v7 + int32(16)
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v29 = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v34 = int32(0)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
					F_genericZpopCommand(m, l0, v30+int32(4), int32(1), v34, v34, v29, base.B2i32(v29 != int32(-1))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v38)), v34, v34)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_zpopminCommand[0]))
		F_addReplyErrorObject(m, l0, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_zrandmemberWithCountCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
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
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
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
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 float64
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int64
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v655 int32
	_ = v655
	var v656 float64
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v686 int32
	_ = v686
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	v11 = m.G0
	v13 = v11 - int32(176)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_zrandmemberWithCountCommand[0]))
	v19 = F_lookupKeyReadOrReply(m, l0, v16, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_zrandmemberWithCountCommand_0), int32(_a_F_zrandmemberWithCountCommand_1), int32(4282))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L5
	} else {
		goto L225
	}
L2:
	;
	F__serverAssert(m, int32(_a_F_zrandmemberWithCountCommand_2), int32(_a_F_zrandmemberWithCountCommand_1), int32(4252))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L5
	} else {
		goto L224
	}
L3:
	;
	F__serverAssert(m, int32(_a_F_zrandmemberWithCountCommand_3), int32(_a_F_zrandmemberWithCountCommand_1), int32(4185))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L223
	}
L4:
	;
	m.G0 = v13 + int32(176)
	return
L5:
	;
	return
L6:
	;
	if v19 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = F_checkType(m, l0, v19, int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v26 = F_zsetLength(m, v19)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if l1 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = l1 >> (uint(int32(31)) % 32)
	v35 = l1 ^ v33 - v33
	if l1 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_zrandmemberWithCountCommand[0]))
	F_addReply(m, l0, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L4
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v19
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v202 = int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = v201 & v202
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = int32(base.Ui32(v201)>>(uint(int32(4))%32)) & v202
	F_zuiInitIterator(m, v13+int32(136))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L70
	}
L15:
	;
	if l2 == int32(0) {
		v46 = v35
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v35 != int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_addReplyArrayLen(m, l0, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v46 = v35 << (uint(base.B2i32(v42 == int32(2))) % 32)
	goto L18
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch int32(base.Ui32(v49)>>(uint(int32(4))%32))&int32(15) + int32(-7) {
	case 0:
		goto L22
	default:
		goto L4
	case 4:
		goto L21
	}
L21:
	;
	v162 = int32(1000)
	if base.Ui32(v35) < base.Ui32(v162) {
		goto L50
	} else {
		goto L51
	}
L22:
	;
	v56 = F_objectGetVal(m, v19)
	mBase = m.M
	v63 = v35
	goto L23
L23:
	;
	if v63 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v72 = F_hashtableFairRandomEntry(m, v69, v13+int32(72))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v72 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if l2 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v157&int32(1024) == int32(0) {
		v63 = v63 + int32(-1)
		goto L23
	} else {
		goto L49
	}
L29:
	;
	v120 = v76 + int32(16)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v124 = v120 + v121<<(uint(int32(3))%32)
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v124))))
	v126 = v124 + v125
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	switch v130 & int32(7) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v147 = int32(0)
		goto L42
	}
L30:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v79) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v86 = v76 + int32(16)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v90 = v86 + v87<<(uint(int32(3))%32)
	v91 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90))))
	v92 = v90 + v91
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	switch v96 & int32(7) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		v113 = int32(0)
		goto L34
	}
L32:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_addReplyBulkCBuffer(m, l0, v92+int32(1), v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L40
	}
L35:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-16))))
	v113 = v112
	goto L34
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v92+int32(-8))))
	v113 = v109
	goto L34
L37:
	;
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+int32(-4)))))
	v113 = v106
	goto L34
L38:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+int32(-2)))))
	v113 = v103
	goto L34
L39:
	;
	v113 = int32(base.Ui32(v96) >> (uint(int32(3)) % 32))
	goto L34
L40:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
	F_addReplyDouble(m, l0, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L28
L42:
	;
	F_addReplyBulkCBuffer(m, l0, v126+int32(1), v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L48
	}
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(-16))))
	v147 = v146
	goto L42
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(-8))))
	v147 = v143
	goto L42
L45:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126+int32(-4)))))
	v147 = v140
	goto L42
L46:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+int32(-2)))))
	v147 = v137
	goto L42
L47:
	;
	v147 = int32(base.Ui32(v130) >> (uint(int32(3)) % 32))
	goto L42
L48:
	;
	goto L28
L49:
	;
	goto L4
L50:
	;
	v165 = v35
	goto L52
L51:
	;
	v165 = v162
	goto L52
L52:
	;
	v167 = v165 << (uint(int32(4)) % 32)
	v168 = F_valkey_malloc(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	if l2 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v180 = v35
	goto L59
L55:
	;
	v171 = F_valkey_malloc(m, v167)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	v173 = int32(0)
	goto L54
L57:
	;
	v173 = v171
	goto L54
L58:
	;
	F_valkey_free(m, v168)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L68
	}
L59:
	;
	v184 = F_objectGetVal(m, v19)
	mBase = m.M
	if base.Ui32(v180) < base.Ui32(v165) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L58
L61:
	;
	v186 = v180
	goto L63
L62:
	;
	v186 = v165
	goto L63
L63:
	;
	F_lpRandomPairs(m, v184, v186, v168, v173)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_zrandmemberReplyWithListpack(m, l0, v186, v168, v173)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v191&int32(4) != 0 {
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v194 = v180 - v186
	if v194 != 0 {
		v180 = v194
		goto L59
	} else {
		goto L67
	}
L67:
	;
	goto L60
L68:
	;
	F_valkey_free(m, v173)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	goto L4
L70:
	;
	v216 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(128)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(120)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(112)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(104)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(96)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(88)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(80)))) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v216
	v244 = base.B2i32(base.Ui32(l1) < base.Ui32(v26))
	if base.Ui32(l1) < base.Ui32(v26) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v245 = l1
	goto L73
L72:
	;
	v245 = v26
	goto L73
L73:
	;
	if l2 == int32(0) {
		v252 = v245
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_addReplyArrayLen(m, l0, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L76
	}
L75:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v252 = v245 << (uint(base.B2i32(v248 == int32(2))) % 32)
	goto L74
L76:
	;
	if base.Ui32(l1) < base.Ui32(v26) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	F_zuiClearIterator(m, v13+int32(136))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L222
	}
L78:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v346&int32(240) != int32(176) {
		goto L115
	} else {
		goto L116
	}
L79:
	;
	v259 = F_zuiNext(m, v13+int32(136), v13+int32(72))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	if v259 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L82
L82:
	;
	if l2 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v344 = F_zuiNext(m, v13+int32(136), v13+int32(72))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L113
	}
L85:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if v312&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L86:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v275) < base.Ui32(int32(3)) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	if v282&int32(1) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	F_addReplyBulkSds(m, l0, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L100
	}
L91:
	;
	if v281 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v282 & int32(-2)
	v305 = v281
	goto L90
L93:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v296 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v294 = F_sdsdup(m, v281)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v305 = v294
	goto L90
L96:
	;
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	v303 = F_sdsfromlonglong(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L99
	}
L97:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v13)+116))
	v300 = F_sdsnewlen(m, v296, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v305 = v300
	goto L90
L99:
	;
	v305 = v303
	goto L90
L100:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v13)+128))
	F_addReplyDouble(m, l0, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	goto L84
L102:
	;
	F_addReplyBulkSds(m, l0, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L112
	}
L103:
	;
	if v311 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v312 & int32(-2)
	v335 = v311
	goto L102
L105:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v326 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v324 = F_sdsdup(m, v311)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v335 = v324
	goto L102
L108:
	;
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	v333 = F_sdsfromlonglong(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v13)+116))
	v330 = F_sdsnewlen(m, v326, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v335 = v330
	goto L102
L111:
	;
	v335 = v333
	goto L102
L112:
	;
	goto L84
L113:
	;
	if v344 != 0 {
		goto L82
	} else {
		goto L114
	}
L114:
	;
	goto L77
L115:
	;
	if base.Ui32(l1*int32(3)) <= base.Ui32(v26) {
		goto L127
	} else {
		goto L128
	}
L116:
	;
	v352 = l1 << (uint(int32(4)) % 32)
	v353 = F_valkey_malloc(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if l2 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v359 = F_objectGetVal(m, v19)
	mBase = m.M
	v360 = F_lpRandomPairsUnique(m, v359, l1, v353, v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L122
	}
L119:
	;
	v356 = F_valkey_malloc(m, v352)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L121
	}
L120:
	;
	v358 = int32(0)
	goto L118
L121:
	;
	v358 = v356
	goto L118
L122:
	;
	if v360 != l1 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_zrandmemberReplyWithListpack(m, l0, l1, v353, v358)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_valkey_free(m, v353)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	F_valkey_free(m, v358)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	goto L77
L127:
	;
	v602 = F_hashtableCreate(m, int32(_a_F_zrandmemberWithCountCommand_4))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L5
	} else {
		goto L187
	}
L128:
	;
	v373 = F_hashtableCreate(m, int32(_a_F_zrandmemberWithCountCommand_5))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	v375 = F_hashtableExpand(m, v373, v26)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	v378 = v13 + int32(16)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v13)+136))
	v380 = F_objectGetVal(m, v379)
	mBase = m.M
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v382 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v378)+14)) = uint8(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v378)+24)) = v382
	*(*uint8)(unsafe.Add(mBase, uint32(v378)+15)) = uint8(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v378)+8)) = int32(-1)
	if v381 == v382 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L136
L132:
	;
	goto L131
L133:
	;
	goto L132
L135:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v373)+20))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	goto L143
L136:
	;
	v412 = F_hashtableNext(m, v13+int32(16), v13)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L138
	}
L137:
	;
	F__serverAssert(m, int32(_a_F_zrandmemberWithCountCommand_6), int32(_a_F_zrandmemberWithCountCommand_1), int32(4280))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L142
	}
L138:
	;
	if v412 == int32(0) {
		goto L135
	} else {
		goto L139
	}
L139:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v417 = F_hashtableAdd(m, v373, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	if v417 != 0 {
		goto L136
	} else {
		goto L141
	}
L141:
	;
	goto L137
L142:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	if v425+v426 != v26 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	if base.Ui32(v26) <= base.Ui32(v35) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	F_hashtableCleanupIterator(m, v13+int32(16))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L5
	} else {
		goto L152
	}
L146:
	;
	v435 = v26
	goto L147
L147:
	;
	v442 = F_hashtableFairRandomEntry(m, v373, v13+int32(12))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L5
	} else {
		goto L149
	}
L148:
	;
	goto L145
L149:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v446 = v444 + int32(16)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v450 = v446 + v447<<(uint(int32(3))%32)
	v451 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450))))
	v455 = F_hashtableDelete(m, v373, v450+v451+int32(1))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v458 = v435 + int32(-1)
	if base.Ui32(v35) < base.Ui32(v458) {
		v435 = v458
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v475 = v13 + int32(16)
	v476 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v475)+14)) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(v475))) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v475)+24)) = v476
	*(*uint8)(unsafe.Add(mBase, uint32(v475)+15)) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(v475)+8)) = int32(-1)
	if v373 == v476 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v498 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L158
	}
L154:
	;
	goto L153
L155:
	;
	goto L154
L157:
	;
	F_hashtableCleanupIterator(m, v13+int32(16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L5
	} else {
		goto L185
	}
L158:
	;
	if v498 == int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	goto L160
L160:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v514 = v512 + int32(16)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v518 = v514 + v515<<(uint(int32(3))%32)
	v519 = int32(*(*int8)(unsafe.Add(mBase, uint32(v518))))
	v520 = v518 + v519
	v522 = v520 + int32(1)
	if l2 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	goto L157
L162:
	;
	v583 = F_hashtableNext(m, v13+int32(16), v13+int32(12))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L183
	}
L163:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	switch v556 & int32(7) {
	case 0:
		goto L181
	case 1:
		goto L180
	case 2:
		goto L179
	case 3:
		goto L178
	case 4:
		goto L177
	default:
		v573 = int32(0)
		goto L176
	}
L164:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v525) < base.Ui32(int32(3)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	switch v532 & int32(7) {
	case 0:
		goto L173
	case 1:
		goto L172
	case 2:
		goto L171
	case 3:
		goto L170
	case 4:
		goto L169
	default:
		v549 = int32(0)
		goto L168
	}
L166:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	F_addReplyBulkCBuffer(m, l0, v522, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L174
	}
L169:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v520+int32(-16))))
	v549 = v548
	goto L168
L170:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v520+int32(-8))))
	v549 = v545
	goto L168
L171:
	;
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v520+int32(-4)))))
	v549 = v542
	goto L168
L172:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520+int32(-2)))))
	v549 = v539
	goto L168
L173:
	;
	v549 = int32(base.Ui32(v532) >> (uint(int32(3)) % 32))
	goto L168
L174:
	;
	v552 = *(*float64)(unsafe.Add(mBase, uint32(v512)))
	F_addReplyDouble(m, l0, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	goto L162
L176:
	;
	F_addReplyBulkCBuffer(m, l0, v522, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L5
	} else {
		goto L182
	}
L177:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v520+int32(-16))))
	v573 = v572
	goto L176
L178:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v520+int32(-8))))
	v573 = v569
	goto L176
L179:
	;
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v520+int32(-4)))))
	v573 = v566
	goto L176
L180:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520+int32(-2)))))
	v573 = v563
	goto L176
L181:
	;
	v573 = int32(base.Ui32(v556) >> (uint(int32(3)) % 32))
	goto L176
L182:
	;
	goto L162
L183:
	;
	if v583 != 0 {
		goto L160
	} else {
		goto L184
	}
L184:
	;
	goto L161
L185:
	;
	F_hashtableRelease(m, v373)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L5
	} else {
		goto L186
	}
L186:
	;
	goto L77
L187:
	;
	v604 = F_hashtableExpand(m, v602, l1)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v606 = int32(0)
	if l2 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v608 = v13
	goto L191
L190:
	;
	v608 = v606
	goto L191
L191:
	;
	v617 = v606
	goto L192
L192:
	;
	F_zsetTypeRandomElement(m, v19, v26, v13+int32(16), v608)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L5
	} else {
		goto L194
	}
L193:
	;
	F_hashtableRelease(m, v602)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L5
	} else {
		goto L221
	}
L194:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v623 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v633 = F_hashtableAdd(m, v602, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L5
	} else {
		goto L202
	}
L196:
	;
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	v630 = F_sdsfromlonglong(m, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L5
	} else {
		goto L199
	}
L197:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v627 = F_sdsnewlen(m, v623, v626)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v632 = v627
	goto L195
L199:
	;
	v632 = v630
	goto L195
L200:
	;
	if base.Ui32(v669) < base.Ui32(v35) {
		v617 = v669
		goto L192
	} else {
		goto L220
	}
L201:
	;
	v638 = v617 + int32(1)
	if l2 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	if v633 != 0 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	F_sdsfree(m, v632)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v669 = v617
	goto L200
L205:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v659 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L206:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v641) < base.Ui32(int32(3)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v647 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v656 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
	F_addReplyDouble(m, l0, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L5
	} else {
		goto L215
	}
L211:
	;
	v653 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	F_addReplyBulkLongLong(m, l0, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L5
	} else {
		goto L214
	}
L212:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_addReplyBulkCBuffer(m, l0, v647, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	goto L210
L215:
	;
	v669 = v638
	goto L200
L216:
	;
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	F_addReplyBulkLongLong(m, l0, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L5
	} else {
		goto L219
	}
L217:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_addReplyBulkCBuffer(m, l0, v659, v662)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	v669 = v638
	goto L200
L219:
	;
	v669 = v638
	goto L200
L220:
	;
	goto L193
L221:
	;
	goto L77
L222:
	;
	goto L4
L223:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_zrangebyscoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	v30 = int32(1)
	F_zrangeGenericCommand(m, v5, v30, v2, int32(2), v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_zrankCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zrankGenericCommand(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zremrangebylexCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_zremrangeGenericCommand(m, l0, int32(3))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_zrevrangebyscoreCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5+int32(24)))) = v2
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(16)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5+int32(8)))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = int32(1101)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(1102)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(1103)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1104)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
	v32 = int32(2)
	F_zrangeGenericCommand(m, v5, int32(1), v2, v32, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		m.G0 = v5 + int32(48)
		return
	}
}
func F_ztrymalloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	v2 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v47 = v2
	} else {
		if l0 != 0 {
			v9 = l0
		} else {
			v9 = int32(4)
		}
		v11 = v9 + int32(8)
		v12 = F_emscripten_builtin_malloc(m, v11)
		mBase = m.M
		if v12 == int32(0) {
			v47 = v2
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v9
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[0]))
			if v17 != int32(-1) {
				v28 = v17
			} else {
				v20 = int32(0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[0])) = v22
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[1])) = v22 + int32(1)
				v28 = v22
			}
			if v28 < int32(260) {
				v37 = v28 << (uint(int32(2)) % 32)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_ztrymalloc[2])))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_ztrymalloc[2]))) = v40 + v11
			} else {
				v31 = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc[3])) = v33 + v11
			}
			v47 = v12 + int32(8)
		}
	}
	return v47
}
func F_ztrymalloc_usable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v3 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(l0) {
		v49 = v3
		v50 = v3
	} else {
		if l0 != 0 {
			v11 = l0
		} else {
			v11 = int32(4)
		}
		v13 = v11 + int32(8)
		v14 = F_emscripten_builtin_malloc(m, v13)
		mBase = m.M
		if v14 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v11
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[0]))
			if v19 != int32(-1) {
				v30 = v19
			} else {
				v22 = int32(0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[0])) = v24
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[1])) = v24 + int32(1)
				v30 = v24
			}
			if v30 < int32(260) {
				v39 = v30 << (uint(int32(2)) % 32)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_ztrymalloc_usable[2])))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_ztrymalloc_usable[2]))) = v42 + v13
			} else {
				v33 = int32(0)
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_ztrymalloc_usable[3])) = v35 + v13
			}
			v49 = v11
			v50 = v14 + int32(8)
		} else {
			v15 = int32(0)
			v49 = v15
			v50 = v15
		}
	}
	if l1 == int32(0) {
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v49
	}
	return v50
}
func F_zuiClearIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 == int32(0) {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v5 + int32(-2) {
		case 0:
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			switch v8 + int32(-2) {
			case 0:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_hashtableReleaseIterator(m, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			default:
				F__serverPanic_1(m, int32(_a_F_zuiClearIterator_0), int32(2201), int32(_a_F_zuiClearIterator_1), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 4, 9:
				return
			}
		case 1:
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			switch v18 + int32(-7) {
			case 0, 4:
				return
			default:
				F__serverPanic_1(m, int32(_a_F_zuiClearIterator_0), int32(2210), int32(_a_F_zuiClearIterator_2), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		default:
			F__serverPanic_1(m, int32(_a_F_zuiClearIterator_0), int32(2213), int32(_a_F_zuiClearIterator_3), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
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
}
func F_zuiCompareByCardinality(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_zuiLength(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_zuiLength(m, l1)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return base.B2i32(base.Ui32(v7) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v7))
		}
	}
}
func F_zuiCompareByRevCardinality(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_zuiLength(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_zuiLength(m, l1)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return base.B2i32(base.Ui32(v3) < base.Ui32(v7)) - base.B2i32(base.Ui32(v7) < base.Ui32(v3))
		}
	}
}
func F_zuiInitIterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v6 + int32(-2) {
		case 0:
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			switch v9 + int32(-2) {
			case 0:
				v12 = F_objectGetVal(m, v3)
				mBase = m.M
				v14 = F_hashtableCreateIterator(m, v12, int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v14
					return
				}
			default:
				F__serverPanic_1(m, int32(_a_F_zuiInitIterator_0), int32(2164), int32(_a_F_zuiInitIterator_1), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 4:
				v69 = F_objectGetVal(m, v3)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
				return
			case 9:
				v17 = F_objectGetVal(m, v3)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v17
				v19 = F_lpFirst(m, v17)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v19
					return
				}
			}
		case 1:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			switch v29 + int32(-7) {
			case 0:
				v50 = F_objectGetVal(m, v3)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v53
				return
			default:
				F__serverPanic_1(m, int32(_a_F_zuiInitIterator_0), int32(2182), int32(_a_F_zuiInitIterator_2), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			case 4:
				v32 = F_objectGetVal(m, v3)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v32
				v35 = F_lpSeek(m, v32, int32(-2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v35
					if v35 == int32(0) {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v41 = F_lpNext(m, v40, v35)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v41
							if v41 != 0 {
								return
							} else {
								F__serverAssert(m, int32(_a_F_zuiInitIterator_3), int32(_a_F_zuiInitIterator_0), int32(2176))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
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
					}
				}
			}
		default:
			F__serverPanic_1(m, int32(_a_F_zuiInitIterator_0), int32(2185), int32(_a_F_zuiInitIterator_4), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
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
}
func F_zuiNext(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int64
	_ = v182
	var v187 int64
	_ = v187
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v208 float64
	_ = v208
	var v212 int64
	_ = v212
	var v214 float64
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v3 {
		v242 = v3
		m.G0 = v11 + int32(16)
		return v242
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v17&int32(1) == int32(0) {
			v27 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(56)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(48)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(40)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(32)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(24)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(16)))) = v27
			*(*int64)(unsafe.Add(mBase, uint32(l1+int32(8)))) = v27
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v57 + int32(-2) {
			case 0:
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				switch v60 + int32(-2) {
				case 0:
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v104 = F_hashtableNext(m, v103, v11)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						if v104 == int32(0) {
							v242 = v3
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v108
							v242 = int32(1)
						}
						m.G0 = v11 + int32(16)
						return v242
					}
				default:
					F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2278), int32(_a_F_zuiNext_1), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 4:
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
					if base.Ui32(v68) <= base.Ui32(v64) {
						v90 = int32(0)
					} else {
						v71 = v63 + int32(8)
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
						switch v72 + int32(-4) {
						case 0:
							v82 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(2))%32)))))
							v87 = v82
						default:
							v86 = int64(*(*int16)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(1))%32)))))
							v87 = v86
						case 4:
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(3))%32))))
							v87 = v78
						}
						*(*int64)(unsafe.Add(mBase, uint32(v11))) = v87
						v90 = int32(1)
					}
					if v90 == int32(0) {
						v242 = v3
					} else {
						v94 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = v94
						v98 = int32(1)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v98
						v242 = v98
					}
					m.G0 = v11 + int32(16)
					return v242
				case 9:
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v113 == int32(0) {
						v242 = v3
						m.G0 = v11 + int32(16)
						return v242
					} else {
						v120 = F_lpGetValue(m, v113, l1+int32(44), l1+int32(48))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v120
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v127 = F_lpNext(m, v125, v126)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v127
								v242 = int32(1)
								m.G0 = v11 + int32(16)
								return v242
							}
						}
					}
				}
			case 1:
				v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				switch v138 + int32(-7) {
				case 0:
					v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v220 == int32(0) {
						v242 = v3
					} else {
						v223 = int32(1)
						v225 = v220 + int32(16)
						v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
						v229 = v225 + v226<<(uint(int32(3))%32)
						v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v229))))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v229 + v230 + v223
						v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v236 = *(*float64)(unsafe.Add(mBase, uint32(v235)))
						*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v236
						v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v239
						v242 = v223
					}
					m.G0 = v11 + int32(16)
					return v242
				default:
					F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2298), int32(_a_F_zuiNext_2), int32(0))
					mBase = m.M
					v256 = m.ExcPending
					if v256 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 4:
					v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v141 == int32(0) {
						v242 = v3
						m.G0 = v11 + int32(16)
						return v242
					} else {
						v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v144 == int32(0) {
							v242 = v3
							m.G0 = v11 + int32(16)
							return v242
						} else {
							v151 = F_lpGetValue(m, v141, l1+int32(44), l1+int32(48))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v151
								v155 = l0 + int32(32)
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
								if v156 == int32(0) {
									F__serverAssert(m, int32(_a_F_zuiNext_3), int32(_a_F_zuiNext_0), int32(857))
									mBase = m.M
									v269 = m.ExcPending
									if v269 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v163 = F_lpGetValue(m, v156, v11+int32(12), v11)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										if v163 == int32(0) {
											v212 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
											v214 = base.F64_convert_i64_s(v212)
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v168 = int32(0)
											v172 = m.G0
											v174 = v172 - int32(32)
											m.G0 = v174
											v176 = F___errno_location(m)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v176))) = v168
											v182 = *(*int64)(unsafe.Add(mBase, _c_F_zuiNext[0]))
											*(*int64)(unsafe.Add(mBase, uint32(v174+int32(8)))) = v182
											*(*int64)(unsafe.Add(mBase, uint32(v174)+24)) = int64(0)
											v187 = *(*int64)(unsafe.Add(mBase, _c_F_zuiNext[1]))
											*(*int64)(unsafe.Add(mBase, uint32(v174))) = v187
											F_ffc_from_chars_double_options(m, v174+int32(16), v163, v163+v167, v174+int32(24), v174)
											mBase = m.M
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
											if v195 == v168 {
											} else {
												if v195 == int32(2) {
													v202 = int32(68)
												} else {
													v202 = int32(28)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v176))) = v202
											}
											v208 = *(*float64)(unsafe.Add(mBase, uint32(v174)+24))
											m.G0 = v174 + int32(32)
											v214 = v208
										}
										*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v214
										v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										F_zzlPrev(m, v216, l0+int32(28), v155)
										mBase = m.M
										v218 = m.ExcPending
										if v218 != 0 {
											return int32(0)
										} else {
											v242 = int32(1)
											m.G0 = v11 + int32(16)
											return v242
										}
									}
								}
							}
						}
					}
				}
			default:
				F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2301), int32(_a_F_zuiNext_4), int32(0))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			F_sdsfree(m, v22)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(56)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(48)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(40)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(32)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(24)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(16)))) = v27
				*(*int64)(unsafe.Add(mBase, uint32(l1+int32(8)))) = v27
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				switch v57 + int32(-2) {
				case 0:
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					switch v60 + int32(-2) {
					case 0:
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v104 = F_hashtableNext(m, v103, v11)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							if v104 == int32(0) {
								v242 = v3
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v108
								v242 = int32(1)
							}
							m.G0 = v11 + int32(16)
							return v242
						}
					default:
						F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2278), int32(_a_F_zuiNext_1), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					case 4:
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
						if base.Ui32(v68) <= base.Ui32(v64) {
							v90 = int32(0)
						} else {
							v71 = v63 + int32(8)
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
							switch v72 + int32(-4) {
							case 0:
								v82 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(2))%32)))))
								v87 = v82
							default:
								v86 = int64(*(*int16)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(1))%32)))))
								v87 = v86
							case 4:
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(3))%32))))
								v87 = v78
							}
							*(*int64)(unsafe.Add(mBase, uint32(v11))) = v87
							v90 = int32(1)
						}
						if v90 == int32(0) {
							v242 = v3
						} else {
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = v94
							v98 = int32(1)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99 + v98
							v242 = v98
						}
						m.G0 = v11 + int32(16)
						return v242
					case 9:
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v113 == int32(0) {
							v242 = v3
							m.G0 = v11 + int32(16)
							return v242
						} else {
							v120 = F_lpGetValue(m, v113, l1+int32(44), l1+int32(48))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1)+56)) = int64(4607182418800017408)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v120
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v127 = F_lpNext(m, v125, v126)
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v127
									v242 = int32(1)
									m.G0 = v11 + int32(16)
									return v242
								}
							}
						}
					}
				case 1:
					v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					switch v138 + int32(-7) {
					case 0:
						v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v220 == int32(0) {
							v242 = v3
						} else {
							v223 = int32(1)
							v225 = v220 + int32(16)
							v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
							v229 = v225 + v226<<(uint(int32(3))%32)
							v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v229))))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v229 + v230 + v223
							v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v236 = *(*float64)(unsafe.Add(mBase, uint32(v235)))
							*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v236
							v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v239
							v242 = v223
						}
						m.G0 = v11 + int32(16)
						return v242
					default:
						F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2298), int32(_a_F_zuiNext_2), int32(0))
						mBase = m.M
						v256 = m.ExcPending
						if v256 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					case 4:
						v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v141 == int32(0) {
							v242 = v3
							m.G0 = v11 + int32(16)
							return v242
						} else {
							v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							if v144 == int32(0) {
								v242 = v3
								m.G0 = v11 + int32(16)
								return v242
							} else {
								v151 = F_lpGetValue(m, v141, l1+int32(44), l1+int32(48))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v151
									v155 = l0 + int32(32)
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
									if v156 == int32(0) {
										F__serverAssert(m, int32(_a_F_zuiNext_3), int32(_a_F_zuiNext_0), int32(857))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v163 = F_lpGetValue(m, v156, v11+int32(12), v11)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												v212 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
												v214 = base.F64_convert_i64_s(v212)
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
												v168 = int32(0)
												v172 = m.G0
												v174 = v172 - int32(32)
												m.G0 = v174
												v176 = F___errno_location(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v176))) = v168
												v182 = *(*int64)(unsafe.Add(mBase, _c_F_zuiNext[0]))
												*(*int64)(unsafe.Add(mBase, uint32(v174+int32(8)))) = v182
												*(*int64)(unsafe.Add(mBase, uint32(v174)+24)) = int64(0)
												v187 = *(*int64)(unsafe.Add(mBase, _c_F_zuiNext[1]))
												*(*int64)(unsafe.Add(mBase, uint32(v174))) = v187
												F_ffc_from_chars_double_options(m, v174+int32(16), v163, v163+v167, v174+int32(24), v174)
												mBase = m.M
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
												if v195 == v168 {
												} else {
													if v195 == int32(2) {
														v202 = int32(68)
													} else {
														v202 = int32(28)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v176))) = v202
												}
												v208 = *(*float64)(unsafe.Add(mBase, uint32(v174)+24))
												m.G0 = v174 + int32(32)
												v214 = v208
											}
											*(*float64)(unsafe.Add(mBase, uint32(l1)+56)) = v214
											v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											F_zzlPrev(m, v216, l0+int32(28), v155)
											mBase = m.M
											v218 = m.ExcPending
											if v218 != 0 {
												return int32(0)
											} else {
												v242 = int32(1)
												m.G0 = v11 + int32(16)
												return v242
											}
										}
									}
								}
							}
						}
					}
				default:
					F__serverPanic_1(m, int32(_a_F_zuiNext_0), int32(2301), int32(_a_F_zuiNext_4), int32(0))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
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
		}
	}
}
func F_zunionCommand(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(0)
	F_zunionInterDiffGenericCommand(m, l0, v2, int32(1), v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_zunionInterDiffGenericCommand(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int64
	_ = v712
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 float64
	_ = v783
	var v784 float64
	_ = v784
	var v785 float64
	_ = v785
	var v791 float64
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v803 int32
	_ = v803
	var v812 float64
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 float64
	_ = v821
	var v822 float64
	_ = v822
	var v823 float64
	_ = v823
	var v826 float64
	_ = v826
	var v832 float64
	_ = v832
	var v834 float64
	_ = v834
	var v836 float64
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 float64
	_ = v844
	var v845 float64
	_ = v845
	var v846 float64
	_ = v846
	var v849 float64
	_ = v849
	var v855 float64
	_ = v855
	var v857 float64
	_ = v857
	var v859 float64
	_ = v859
	var v860 float64
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v882 float64
	_ = v882
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int64
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v946 int32
	_ = v946
	var v955 int64
	_ = v955
	var v959 int64
	_ = v959
	var v960 int64
	_ = v960
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v974 int32
	_ = v974
	var v978 int64
	_ = v978
	var v981 int64
	_ = v981
	var v983 int64
	_ = v983
	var v986 int64
	_ = v986
	var v997 int64
	_ = v997
	var v1000 int64
	_ = v1000
	var v1011 int64
	_ = v1011
	var v1014 int64
	_ = v1014
	var v1027 int64
	_ = v1027
	var v1034 int64
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1042 int64
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int64
	_ = v1058
	var v1066 int64
	_ = v1066
	var v1069 int64
	_ = v1069
	var v1081 int64
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int64
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1102 int64
	_ = v1102
	var v1110 int64
	_ = v1110
	var v1113 int64
	_ = v1113
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int64
	_ = v1127
	var v1135 int64
	_ = v1135
	var v1137 int64
	_ = v1137
	var v1142 int64
	_ = v1142
	var v1151 int32
	_ = v1151
	var v1152 int64
	_ = v1152
	var v1163 int64
	_ = v1163
	var v1168 int64
	_ = v1168
	var v1173 int64
	_ = v1173
	var v1176 int64
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1278 int32
	_ = v1278
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1320 int32
	_ = v1320
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1361 float64
	_ = v1361
	var v1362 float64
	_ = v1362
	var v1363 float64
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int64
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1387 float64
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int64
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1431 int32
	_ = v1431
	var v1440 int64
	_ = v1440
	var v1444 int64
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1453 int64
	_ = v1453
	var v1455 int64
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1463 int64
	_ = v1463
	var v1466 int64
	_ = v1466
	var v1468 int64
	_ = v1468
	var v1471 int64
	_ = v1471
	var v1482 int64
	_ = v1482
	var v1485 int64
	_ = v1485
	var v1496 int64
	_ = v1496
	var v1499 int64
	_ = v1499
	var v1512 int64
	_ = v1512
	var v1519 int64
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1527 int64
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1543 int64
	_ = v1543
	var v1551 int64
	_ = v1551
	var v1554 int64
	_ = v1554
	var v1566 int64
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1570 int64
	_ = v1570
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1587 int64
	_ = v1587
	var v1595 int64
	_ = v1595
	var v1598 int64
	_ = v1598
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int64
	_ = v1612
	var v1620 int64
	_ = v1620
	var v1622 int64
	_ = v1622
	var v1627 int64
	_ = v1627
	var v1636 int32
	_ = v1636
	var v1637 int64
	_ = v1637
	var v1648 int64
	_ = v1648
	var v1653 int64
	_ = v1653
	var v1658 int64
	_ = v1658
	var v1661 int64
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 float64
	_ = v1731
	var v1733 float64
	_ = v1733
	var v1739 float64
	_ = v1739
	var v1741 float64
	_ = v1741
	var v1743 float64
	_ = v1743
	var v1744 float64
	_ = v1744
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1777 int32
	_ = v1777
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1900 int32
	_ = v1900
	var v1906 int32
	_ = v1906
	var v1942 int64
	_ = v1942
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1983 int64
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int64
	_ = v2003
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2037 int32
	_ = v2037
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2086 float64
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(160)
	m.G0 = v22
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v6
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+l2<<(uint(int32(2))%32))))
	v40 = F_getLongFromObjectOrReply(m, l0, v36, v22+int32(156), v6)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(160)
	return
L2:
	;
	return
L3:
	;
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if int32(0) < v42 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v53 = l2 + int32(1)
	if v42 <= v51-v53 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v46
	F_addReplyErrorFormat(m, l0, int32(_a_F_zunionInterDiffGenericCommand_0), v22)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L1
L8:
	;
	v61 = v42 * int32(40)
	v62 = int32(0)
	if base.Ui32(int32(2147483646)) < base.Ui32(v61) {
		v108 = v62
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[0]))
	F_addReplyErrorObject(m, l0, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	F_valkey_free(m, v108)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L2
	} else {
		goto L440
	}
L12:
	;
	F_addReplyError(m, l0, int32(_a_F_zunionInterDiffGenericCommand_1))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L2
	} else {
		goto L439
	}
L13:
	;
	if v108 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L14:
	;
	goto L13
L15:
	;
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = v61
	goto L18
L17:
	;
	v70 = int32(4)
	goto L18
L18:
	;
	v72 = v70 + int32(8)
	v73 = F_emscripten_builtin_calloc(m, int32(1), v72)
	mBase = m.M
	if v73 == int32(0) {
		v108 = v62
		goto L14
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v70
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[1]))
	if v78 != int32(-1) {
		v89 = v78
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v89 < int32(260) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v81 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[1])) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[2])) = v83 + int32(1)
	v89 = v83
	goto L20
L22:
	;
	v108 = v73 + int32(8)
	goto L14
L23:
	;
	v98 = v89 << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_zunionInterDiffGenericCommand[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_zunionInterDiffGenericCommand[3]))) = v101 + v72
	goto L22
L24:
	;
	v92 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[4])) = v94 + v72
	goto L22
L25:
	;
	v113 = int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v114 < v113 {
		v196 = v53
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v209 <= v196 {
		v687 = v113
		v690 = int32(1)
		goto L40
	} else {
		goto L41
	}
L27:
	;
	v124 = int32(0)
	v125 = v53
	goto L28
L28:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v125<<(uint(int32(2))%32))))
	v143 = F_lookupKeyRead(m, v137, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L32
	}
L29:
	;
	v196 = v184
	goto L26
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108+v124*int32(40))+16)) = int64(4607182418800017408)
	v183 = int32(1)
	v184 = v125 + v183
	v186 = v124 + v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v186 < v187 {
		v124 = v186
		v125 = v184
		goto L28
	} else {
		goto L38
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108+v124*int32(40)))) = int32(0)
	goto L30
L32:
	;
	if v143 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v147&int32(14) == int32(2) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v160 = v108 + v124*int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v143
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v163 = int32(15)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v162 & v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = int32(base.Ui32(v166)>>(uint(int32(4))%32)) & v163
	goto L30
L35:
	;
	F_valkey_free(m, v108)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[5]))
	F_addReplyErrorObject(m, l0, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	goto L29
L39:
	;
	F_valkey_free(m, v108)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L2
	} else {
		goto L438
	}
L40:
	;
	if l3 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L41:
	;
	v211 = v209 - v196
	if v211 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v214 = int32(0)
	v218 = int32(1)
	v222 = base.B2i32(l3 == v218) | base.B2i32(l4 != v214)
	v234 = v211
	v235 = v196
	v239 = v218
	v240 = v214
	goto L44
L43:
	;
	v212 = int32(1)
	v687 = v212
	v690 = v212
	goto L40
L44:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v222|base.B2i32(v234 <= v247) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v687 = base.B2i32(v670 == int32(0))
	v690 = v669
	goto L40
L46:
	;
	if v664 != 0 {
		v234 = v664
		v235 = v665
		v239 = v669
		v240 = v670
		goto L44
	} else {
		goto L157
	}
L47:
	;
	v344 = base.B2i32(v234 < int32(2))
	if v344|(v222^v218^v218) != 0 {
		goto L69
	} else {
		goto L70
	}
L48:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v235<<(uint(int32(2))%32))))
	v255 = F_objectGetVal(m, v254)
	mBase = m.M
	v256 = int32(_a_F_zunionInterDiffGenericCommand_2)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v259 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v291-v293 != 0 {
		goto L47
	} else {
		goto L61
	}
L50:
	;
	v291 = F_tolower(m, v287)
	mBase = m.M
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v293 = F_tolower(m, v292)
	mBase = m.M
	goto L49
L51:
	;
	v261 = v255
	v262 = v256
	v263 = v259
	goto L54
L52:
	;
	v287 = int32(0)
	v288 = v256
	goto L50
L53:
	;
	v287 = v284 & int32(255)
	v288 = v283
	goto L50
L54:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v265 == int32(0) {
		v283 = v262
		v284 = v263
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v283 = v277
	v284 = int32(0)
	goto L53
L56:
	;
	v269 = v263 & int32(255)
	if v269 == v265 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v276 = int32(1)
	v277 = v262 + v276
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v278 != 0 {
		v261 = v261 + v276
		v262 = v277
		v263 = v278
		goto L54
	} else {
		goto L60
	}
L58:
	;
	v271 = F_tolower(m, v269)
	mBase = m.M
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v273 = F_tolower(m, v272)
	mBase = m.M
	if v271 == v273 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v283 = v262
	v284 = v275
	goto L53
L60:
	;
	goto L55
L61:
	;
	v296 = v234 + int32(-1)
	v298 = v235 + int32(1)
	v299 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v300 <= v299 {
		v664 = v296
		v665 = v298
		v669 = v239
		v670 = v240
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v309 = v296
	v310 = v298
	v312 = v299
	goto L63
L63:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322+v310<<(uint(int32(2))%32))))
	v333 = F_getDoubleFromObjectOrReply(m, l0, v326, v108+v312*int32(40)+int32(16), int32(_a_F_zunionInterDiffGenericCommand_3))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	if v333 != 0 {
		goto L39
	} else {
		goto L66
	}
L66:
	;
	v336 = v309 + int32(-1)
	v337 = int32(1)
	v338 = v310 + v337
	v340 = v312 + v337
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v341 <= v340 {
		v664 = v336
		v665 = v338
		v669 = v239
		v670 = v240
		goto L46
	} else {
		goto L67
	}
L67:
	;
	v309 = v336
	v310 = v338
	v312 = v340
	goto L63
L68:
	;
	v664 = v234 + int32(-2)
	v665 = v235 + int32(2)
	v669 = v653
	v670 = v240
	goto L46
L69:
	;
	if base.B2i32(v234 < int32(1))|base.B2i32(l1|l4 != v214) != 0 {
		goto L125
	} else {
		goto L126
	}
L70:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v235<<(uint(int32(2))%32))))
	v351 = F_objectGetVal(m, v350)
	mBase = m.M
	v352 = int32(_a_F_zunionInterDiffGenericCommand_4)
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if v355 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if v387-v389 != 0 {
		goto L69
	} else {
		goto L83
	}
L72:
	;
	v387 = F_tolower(m, v383)
	mBase = m.M
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v389 = F_tolower(m, v388)
	mBase = m.M
	goto L71
L73:
	;
	v357 = v351
	v358 = v352
	v359 = v355
	goto L76
L74:
	;
	v383 = int32(0)
	v384 = v352
	goto L72
L75:
	;
	v383 = v380 & int32(255)
	v384 = v379
	goto L72
L76:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v361 == int32(0) {
		v379 = v358
		v380 = v359
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v379 = v373
	v380 = int32(0)
	goto L75
L78:
	;
	v365 = v359 & int32(255)
	if v365 == v361 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v372 = int32(1)
	v373 = v358 + v372
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
	if v374 != 0 {
		v357 = v357 + v372
		v358 = v373
		v359 = v374
		goto L76
	} else {
		goto L82
	}
L80:
	;
	v367 = F_tolower(m, v365)
	mBase = m.M
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	v369 = F_tolower(m, v368)
	mBase = m.M
	if v367 == v369 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	v379 = v358
	v380 = v371
	goto L75
L82:
	;
	goto L77
L83:
	;
	v391 = int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v394 = v235 + v391
	v396 = v394 << (uint(int32(2)) % 32)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v392+v396)))
	v399 = F_objectGetVal(m, v398)
	mBase = m.M
	v400 = int32(_a_F_zunionInterDiffGenericCommand_5)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	if v403 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v435-v437 == int32(0) {
		v653 = v391
		goto L68
	} else {
		goto L96
	}
L85:
	;
	v435 = F_tolower(m, v431)
	mBase = m.M
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v437 = F_tolower(m, v436)
	mBase = m.M
	goto L84
L86:
	;
	v405 = v399
	v406 = v400
	v407 = v403
	goto L89
L87:
	;
	v431 = int32(0)
	v432 = v400
	goto L85
L88:
	;
	v431 = v428 & int32(255)
	v432 = v427
	goto L85
L89:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v409 == int32(0) {
		v427 = v406
		v428 = v407
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v427 = v421
	v428 = int32(0)
	goto L88
L91:
	;
	v413 = v407 & int32(255)
	if v413 == v409 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v420 = int32(1)
	v421 = v406 + v420
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+1)))
	if v422 != 0 {
		v405 = v405 + v420
		v406 = v421
		v407 = v422
		goto L89
	} else {
		goto L95
	}
L93:
	;
	v415 = F_tolower(m, v413)
	mBase = m.M
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v417 = F_tolower(m, v416)
	mBase = m.M
	if v415 == v417 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v427 = v406
	v428 = v419
	goto L88
L95:
	;
	goto L90
L96:
	;
	v441 = int32(2)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442+v394<<(uint(v441)%32))))
	v447 = F_objectGetVal(m, v446)
	mBase = m.M
	v448 = int32(_a_F_zunionInterDiffGenericCommand_6)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	if v451 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v483-v485 == int32(0) {
		v653 = v441
		goto L68
	} else {
		goto L109
	}
L98:
	;
	v483 = F_tolower(m, v479)
	mBase = m.M
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v485 = F_tolower(m, v484)
	mBase = m.M
	goto L97
L99:
	;
	v453 = v447
	v454 = v448
	v455 = v451
	goto L102
L100:
	;
	v479 = int32(0)
	v480 = v448
	goto L98
L101:
	;
	v479 = v476 & int32(255)
	v480 = v475
	goto L98
L102:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	if v457 == int32(0) {
		v475 = v454
		v476 = v455
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v475 = v469
	v476 = int32(0)
	goto L101
L104:
	;
	v461 = v455 & int32(255)
	if v461 == v457 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v468 = int32(1)
	v469 = v454 + v468
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	if v470 != 0 {
		v453 = v453 + v468
		v454 = v469
		v455 = v470
		goto L102
	} else {
		goto L108
	}
L106:
	;
	v463 = F_tolower(m, v461)
	mBase = m.M
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	v465 = F_tolower(m, v464)
	mBase = m.M
	if v463 == v465 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	v475 = v454
	v476 = v467
	goto L101
L108:
	;
	goto L103
L109:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v490+v396)))
	v493 = F_objectGetVal(m, v492)
	mBase = m.M
	v494 = int32(_a_F_zunionInterDiffGenericCommand_7)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v497 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if v529-v531 == int32(0) {
		v653 = int32(3)
		goto L68
	} else {
		goto L122
	}
L111:
	;
	v529 = F_tolower(m, v525)
	mBase = m.M
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	v531 = F_tolower(m, v530)
	mBase = m.M
	goto L110
L112:
	;
	v499 = v493
	v500 = v494
	v501 = v497
	goto L115
L113:
	;
	v525 = int32(0)
	v526 = v494
	goto L111
L114:
	;
	v525 = v522 & int32(255)
	v526 = v521
	goto L111
L115:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v503 == int32(0) {
		v521 = v500
		v522 = v501
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v521 = v515
	v522 = int32(0)
	goto L114
L117:
	;
	v507 = v501 & int32(255)
	if v507 == v503 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v514 = int32(1)
	v515 = v500 + v514
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	if v516 != 0 {
		v499 = v499 + v514
		v500 = v515
		v501 = v516
		goto L115
	} else {
		goto L121
	}
L119:
	;
	v509 = F_tolower(m, v507)
	mBase = m.M
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	v511 = F_tolower(m, v510)
	mBase = m.M
	if v509 == v511 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v521 = v500
	v522 = v513
	goto L114
L121:
	;
	goto L116
L122:
	;
	F_valkey_free(m, v108)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[0]))
	F_addReplyErrorObject(m, l0, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	goto L1
L125:
	;
	if l4 == int32(0) {
		goto L11
	} else {
		goto L140
	}
L126:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v544+v235<<(uint(int32(2))%32))))
	v549 = F_objectGetVal(m, v548)
	mBase = m.M
	v550 = int32(_a_F_zunionInterDiffGenericCommand_8)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	if v553 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if v585-v587 != 0 {
		goto L11
	} else {
		goto L139
	}
L128:
	;
	v585 = F_tolower(m, v581)
	mBase = m.M
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v587 = F_tolower(m, v586)
	mBase = m.M
	goto L127
L129:
	;
	v555 = v549
	v556 = v550
	v557 = v553
	goto L132
L130:
	;
	v581 = int32(0)
	v582 = v550
	goto L128
L131:
	;
	v581 = v578 & int32(255)
	v582 = v577
	goto L128
L132:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if v559 == int32(0) {
		v577 = v556
		v578 = v557
		goto L131
	} else {
		goto L134
	}
L133:
	;
	v577 = v571
	v578 = int32(0)
	goto L131
L134:
	;
	v563 = v557 & int32(255)
	if v563 == v559 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v570 = int32(1)
	v571 = v556 + v570
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+1)))
	if v572 != 0 {
		v555 = v555 + v570
		v556 = v571
		v557 = v572
		goto L132
	} else {
		goto L138
	}
L136:
	;
	v565 = F_tolower(m, v563)
	mBase = m.M
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v567 = F_tolower(m, v566)
	mBase = m.M
	if v565 == v567 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	v577 = v556
	v578 = v569
	goto L131
L138:
	;
	goto L133
L139:
	;
	v591 = int32(1)
	v664 = v234 + int32(-1)
	v665 = v235 + v591
	v669 = v239
	v670 = v591
	goto L46
L140:
	;
	if v234 < int32(2) {
		goto L11
	} else {
		goto L141
	}
L141:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v598 = v235 << (uint(int32(2)) % 32)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v596+v598)))
	v601 = F_objectGetVal(m, v600)
	mBase = m.M
	v602 = int32(_a_F_zunionInterDiffGenericCommand_9)
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	if v605 != 0 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	if v637-v639 != 0 {
		goto L11
	} else {
		goto L154
	}
L143:
	;
	v637 = F_tolower(m, v633)
	mBase = m.M
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	v639 = F_tolower(m, v638)
	mBase = m.M
	goto L142
L144:
	;
	v607 = v601
	v608 = v602
	v609 = v605
	goto L147
L145:
	;
	v633 = int32(0)
	v634 = v602
	goto L143
L146:
	;
	v633 = v630 & int32(255)
	v634 = v629
	goto L143
L147:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v611 == int32(0) {
		v629 = v608
		v630 = v609
		goto L146
	} else {
		goto L149
	}
L148:
	;
	v629 = v623
	v630 = int32(0)
	goto L146
L149:
	;
	v615 = v609 & int32(255)
	if v615 == v611 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v622 = int32(1)
	v623 = v608 + v622
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	if v624 != 0 {
		v607 = v607 + v622
		v608 = v623
		v609 = v624
		goto L147
	} else {
		goto L153
	}
L151:
	;
	v617 = F_tolower(m, v615)
	mBase = m.M
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v619 = F_tolower(m, v618)
	mBase = m.M
	if v617 == v619 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v629 = v608
	v630 = v621
	goto L146
L153:
	;
	goto L148
L154:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v641+v598+int32(4))))
	v649 = F_getPositiveLongFromObjectOrReply(m, l0, v645, v22+int32(72), int32(_a_F_zunionInterDiffGenericCommand_10))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	if v649 != 0 {
		goto L39
	} else {
		goto L156
	}
L156:
	;
	v653 = v239
	goto L68
L157:
	;
	goto L45
L158:
	;
	if l4 != 0 {
		v711 = int32(0)
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	F_qsort(m, v108, v700, int32(40), int32(1095))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v712 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(144)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(136)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(128)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(120)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(112)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(104)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(96)))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v712
	if l3 != int32(2) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v706 = F_createZsetObject(m)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L2
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v706
	v709 = F_objectGetVal(m, v706)
	mBase = m.M
	v711 = v709
	goto L161
L164:
	;
	if l1 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L165:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if l3 != 0 {
		goto L273
	} else {
		goto L274
	}
L166:
	;
	v745 = F_zuiLength(m, v108)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	if v745 == int32(0) {
		v1942 = v712
		goto L164
	} else {
		goto L168
	}
L168:
	;
	F_zuiInitIterator(m, v108)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	v753 = F_zuiNext(m, v108, v22+int32(88))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L2
	} else {
		goto L173
	}
L170:
	;
	F_zuiClearIterator(m, v108)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L2
	} else {
		goto L271
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v1260
	v1278 = v1257
	goto L170
L172:
	;
	v759 = v690 + int32(-1)
	v760 = int32(0)
	v774 = v760
	v776 = v760
	v777 = v760
	goto L175
L173:
	;
	if v753 != 0 {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v755 = int32(0)
	v1257 = v755
	v1259 = v755
	v1260 = v755
	goto L171
L175:
	;
	v783 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	v784 = *(*float64)(unsafe.Add(mBase, uint32(v22)+144))
	v785 = base.F64_mul(v783, v784)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v785)&int64(9223372036854775807)) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v1257 = v1236
	v1259 = v1238
	v1260 = v1239
	goto L171
L177:
	;
	v791 = float64(0)
	goto L179
L178:
	;
	v791 = v785
	goto L179
L179:
	;
	v792 = int32(1)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v793 < int32(2) {
		v872 = v793
		v873 = v792
		v882 = v791
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if l4 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L181:
	;
	v803 = v792
	v812 = v791
	goto L182
L182:
	;
	v817 = v108 + v803*int32(40)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v818 != v819 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v872 = v864
	v873 = v863
	v882 = v860
	goto L180
L184:
	;
	v863 = v803 + int32(1)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v863 < v864 {
		v803 = v863
		v812 = v860
		goto L182
	} else {
		goto L214
	}
L185:
	;
	v841 = F_zuiFind(m, v817, v22+int32(88), v22+int32(16))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L2
	} else {
		goto L200
	}
L186:
	;
	v821 = *(*float64)(unsafe.Add(mBase, uint32(v22)+144))
	v822 = *(*float64)(unsafe.Add(mBase, uint32(v817)+16))
	v823 = base.F64_mul(v821, v822)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+16)) = v823
	switch v759 {
	case 0:
		goto L189
	case 1:
		goto L188
	default:
		goto L187
	}
L187:
	;
	if base.F64_gt(v823, v812) != 0 {
		goto L196
	} else {
		goto L197
	}
L188:
	;
	if base.F64_lt(v823, v812) != 0 {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v826 = base.F64_add(v812, v823)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v826)&int64(9223372036854775807)) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v832 = float64(0)
	goto L192
L191:
	;
	v832 = v826
	goto L192
L192:
	;
	v860 = v832
	goto L184
L193:
	;
	v834 = v823
	goto L195
L194:
	;
	v834 = v812
	goto L195
L195:
	;
	v860 = v834
	goto L184
L196:
	;
	v836 = v823
	goto L198
L197:
	;
	v836 = v812
	goto L198
L198:
	;
	v860 = v836
	goto L184
L199:
	;
	v844 = *(*float64)(unsafe.Add(mBase, uint32(v817)+16))
	v845 = *(*float64)(unsafe.Add(mBase, uint32(v22)+16))
	v846 = base.F64_mul(v844, v845)
	*(*float64)(unsafe.Add(mBase, uint32(v22)+16)) = v846
	switch v759 {
	case 0:
		goto L204
	case 1:
		goto L203
	default:
		goto L202
	}
L200:
	;
	if v841 != 0 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	v872 = v843
	v873 = v803
	v882 = v812
	goto L180
L202:
	;
	if base.F64_gt(v846, v812) != 0 {
		goto L211
	} else {
		goto L212
	}
L203:
	;
	if base.F64_lt(v846, v812) != 0 {
		goto L208
	} else {
		goto L209
	}
L204:
	;
	v849 = base.F64_add(v812, v846)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v849)&int64(9223372036854775807)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v855 = float64(0)
	goto L207
L206:
	;
	v855 = v849
	goto L207
L207:
	;
	v860 = v855
	goto L184
L208:
	;
	v857 = v846
	goto L210
L209:
	;
	v857 = v812
	goto L210
L210:
	;
	v860 = v857
	goto L184
L211:
	;
	v859 = v846
	goto L213
L212:
	;
	v859 = v812
	goto L213
L213:
	;
	v860 = v859
	goto L184
L214:
	;
	goto L183
L215:
	;
	v1244 = F_zuiNext(m, v108, v22+int32(88))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L2
	} else {
		goto L269
	}
L216:
	;
	if v873 != v872 {
		v1236 = v774
		v1238 = v776
		v1239 = v777
		goto L215
	} else {
		goto L222
	}
L217:
	;
	if v873 != v872 {
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v892 = v774 + int32(1)
	if base.Ui32(v892) <= base.Ui32(v888+int32(-1)) {
		v1236 = v892
		v1238 = v776
		v1239 = v777
		goto L215
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v777
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+88)))
	if v896&int32(1) == int32(0) {
		v1278 = v892
		goto L170
	} else {
		goto L220
	}
L220:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	F_sdsfree(m, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L2
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = int32(0)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v906 & int32(-2)
	v1278 = v892
	goto L170
L222:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	if v912&int32(1) == int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[6]))
	if int32(311) < v946 {
		goto L235
	} else {
		goto L236
	}
L224:
	;
	if v911 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v912 & int32(-2)
	v935 = v911
	goto L223
L226:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v22)+128))
	if v926 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v924 = F_sdsdup(m, v911)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v935 = v924
	goto L223
L229:
	;
	v932 = *(*int64)(unsafe.Add(mBase, uint32(v22)+136))
	v933 = F_sdsfromlonglong(m, v932)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L2
	} else {
		goto L232
	}
L230:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	v930 = F_sdsnewlen(m, v926, v929)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L2
	} else {
		goto L231
	}
L231:
	;
	v935 = v930
	goto L223
L232:
	;
	v935 = v933
	goto L223
L233:
	;
	v1179 = int32(1)
	if v1176 == int64(0) {
		goto L249
	} else {
		goto L250
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[6])) = v1151
	v1163 = int64(base.Ui64(v1152)>>(uint(int64(29))%64))&int64(22906492245) ^ v1152
	v1168 = v1163<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v1163
	v1173 = v1168<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v1168
	v1176 = int64(base.Ui64(v1173)>>(uint(int64(43))%64)) ^ v1173
	goto L233
L235:
	;
	if v946 == int32(313) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v955 = *(*int64)(unsafe.Add(mBase, uint32(v946<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1151 = v946 + int32(1)
	v1152 = v955
	goto L234
L237:
	;
	v1039 = int32(0)
	v1042 = v1034
	goto L243
L238:
	;
	v960 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7])) = v960
	v968 = int64(1)
	v970 = v960
	goto L240
L239:
	;
	v959 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7]))
	v1034 = v959
	goto L237
L240:
	;
	v974 = int32(3)
	v978 = int64(62)
	v981 = int64(6364136223846793005)
	v983 = (int64(base.Ui64(v970)>>(uint(v978)%64))^v970)*v981 + v968
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v968)<<(uint(v974)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v983
	v986 = v968 + int64(1)
	v997 = (int64(base.Ui64(v983)>>(uint(v978)%64))^v983)*v981 + v986
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v986)<<(uint(v974)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v997
	v1000 = v968 + int64(2)
	v1011 = (int64(base.Ui64(v997)>>(uint(v978)%64))^v997)*v981 + v1000
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1000)<<(uint(v974)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1011
	v1014 = v968 + int64(3)
	if v1014 == int64(312) {
		v1034 = v960
		goto L237
	} else {
		goto L242
	}
L242:
	;
	v1027 = (int64(base.Ui64(v1011)>>(uint(int64(62))%64))^v1011)*int64(6364136223846793005) + v1014
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1014)<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1027
	v968 = v968 + int64(4)
	v970 = v1027
	goto L240
L243:
	;
	v1048 = int32(3)
	v1049 = v1039 << (uint(v1048) % 32)
	v1052 = int32(1)
	v1053 = v1039 + v1052
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v1053<<(uint(v1048)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1066 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1058)&v1052<<(uint(v1048)%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1069 = *(*int64)(unsafe.Add(mBase, uint32(v1049)+uint32(_c_F_zunionInterDiffGenericCommand[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v1049)+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1066 ^ v1069 ^ int64(base.Ui64(v1042&int64(-2147483648)|v1058&int64(2147483646))>>(uint(int64(1))%64))
	if v1053 != int32(156) {
		v1039 = v1053
		v1042 = v1058
		goto L243
	} else {
		goto L245
	}
L244:
	;
	v1081 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[9]))
	v1083 = int32(156)
	v1085 = v1081
	goto L246
L245:
	;
	goto L244
L246:
	;
	v1092 = int32(3)
	v1093 = v1083 << (uint(v1092) % 32)
	v1096 = int32(1)
	v1097 = v1083 + v1096
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v1097<<(uint(v1092)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1110 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1102)&v1096<<(uint(v1092)%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1113 = *(*int64)(unsafe.Add(mBase, uint32(v1093)+uint32(_c_F_zunionInterDiffGenericCommand[10])))
	*(*int64)(unsafe.Add(mBase, uint32(v1093)+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1110 ^ v1113 ^ int64(base.Ui64(v1085&int64(-2147483648)|v1102&int64(2147483646))>>(uint(int64(1))%64))
	if v1097 != int32(311) {
		v1083 = v1097
		v1085 = v1102
		goto L246
	} else {
		goto L248
	}
L247:
	;
	v1124 = int32(1)
	v1125 = int32(0)
	v1127 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7]))
	v1135 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1127)&v1124<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1137 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[11]))
	v1142 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[12])) = v1135 ^ v1137 ^ int64(base.Ui64(v1127&int64(2147483646)|v1142&int64(-2147483648))>>(uint(int64(1))%64))
	v1151 = v1124
	v1152 = v1127
	goto L234
L248:
	;
	goto L247
L249:
	;
	v1185 = int32(32)
	goto L251
L250:
	;
	v1185 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v1176)))>>(uint(v1179)%32)) + v1179
	goto L251
L251:
	;
	v1186 = F_zslCreateNode(m, v1185, v882, v935)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L2
	} else {
		goto L252
	}
L252:
	;
	v1188 = F_zslInsertNode(m, v936, v1186)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L2
	} else {
		goto L253
	}
L253:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v1191 = F_hashtableAdd(m, v1190, v1188)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L2
	} else {
		goto L254
	}
L254:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935+int32(-1)))))
	v1197 = v1195 & int32(7)
	switch v1197 {
	case 0:
		goto L261
	case 1:
		goto L260
	case 2:
		goto L259
	case 3:
		goto L258
	case 4:
		goto L257
	default:
		v1230 = v776
		v1231 = v777
		goto L255
	}
L255:
	;
	F_sdsfree(m, v935)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L2
	} else {
		goto L268
	}
L256:
	;
	v1213 = v777 + v1212
	if base.Ui32(v1212) <= base.Ui32(v776) {
		v1230 = v776
		v1231 = v1213
		goto L255
	} else {
		goto L262
	}
L257:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v935+int32(-17))))
	v1212 = v1211
	goto L256
L258:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v935+int32(-9))))
	v1212 = v1208
	goto L256
L259:
	;
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935+int32(-5)))))
	v1212 = v1205
	goto L256
L260:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935+int32(-3)))))
	v1212 = v1202
	goto L256
L261:
	;
	v1212 = int32(base.Ui32(v1195) >> (uint(int32(3)) % 32))
	goto L256
L262:
	;
	switch v1197 {
	default:
		goto L267
	case 1:
		goto L266
	case 2:
		goto L265
	case 3:
		goto L264
	case 4:
		goto L263
	}
L263:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v935+int32(-17))))
	v1230 = v1228
	v1231 = v1213
	goto L255
L264:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v935+int32(-9))))
	v1230 = v1225
	v1231 = v1213
	goto L255
L265:
	;
	v1222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v935+int32(-5)))))
	v1230 = v1222
	v1231 = v1213
	goto L255
L266:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935+int32(-3)))))
	v1230 = v1219
	v1231 = v1213
	goto L255
L267:
	;
	v1230 = int32(base.Ui32(v1195) >> (uint(int32(3)) % 32))
	v1231 = v1213
	goto L255
L268:
	;
	v1236 = v774
	v1238 = v1230
	v1239 = v1231
	goto L215
L269:
	;
	if v1244 != 0 {
		v774 = v1236
		v776 = v1238
		v777 = v1239
		goto L175
	} else {
		goto L270
	}
L270:
	;
	goto L176
L271:
	;
	v1942 = base.I64_extend_i32_u(v1278)
	goto L164
L272:
	;
	v1942 = int64(0)
	goto L164
L273:
	;
	F_zdiff(m, v108, v1289, v711, v22+int32(84), v22+int32(80))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L2
	} else {
		goto L379
	}
L274:
	;
	if v1289 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1821 = v22 + int32(16)
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v1823 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1821)+14)) = uint8(v1823)
	*(*int32)(unsafe.Add(mBase, uint32(v1821))) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+24)) = v1823
	*(*uint8)(unsafe.Add(mBase, uint32(v1821)+15)) = uint8(v1823)
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+8)) = int32(-1)
	if v1822 == v1823 {
		goto L367
	} else {
		goto L368
	}
L276:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v1298 = F_zuiLength(m, v108+v1289*int32(40)+int32(-40))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L2
	} else {
		goto L277
	}
L277:
	;
	v1300 = F_hashtableExpand(m, v1292, v1298)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L2
	} else {
		goto L278
	}
L278:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v1302 < int32(1) {
		goto L275
	} else {
		goto L279
	}
L279:
	;
	v1320 = int32(0)
	goto L280
L280:
	;
	v1329 = v108 + v1320*int32(40)
	v1330 = F_zuiLength(m, v1329)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L2
	} else {
		goto L283
	}
L281:
	;
	goto L275
L282:
	;
	v1798 = v1320 + int32(1)
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	if v1798 < v1799 {
		v1320 = v1798
		goto L280
	} else {
		goto L365
	}
L283:
	;
	if v1330 == int32(0) {
		goto L282
	} else {
		goto L284
	}
L284:
	;
	F_zuiInitIterator(m, v1329)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L2
	} else {
		goto L285
	}
L285:
	;
	v1338 = F_zuiNext(m, v1329, v22+int32(88))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L2
	} else {
		goto L287
	}
L286:
	;
	F_zuiClearIterator(m, v1329)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L2
	} else {
		goto L364
	}
L287:
	;
	if v1338 == int32(0) {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	goto L289
L289:
	;
	v1361 = *(*float64)(unsafe.Add(mBase, uint32(v1329)+16))
	v1362 = *(*float64)(unsafe.Add(mBase, uint32(v22)+144))
	v1363 = base.F64_mul(v1361, v1362)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	if v1365 != 0 {
		v1381 = v1365
		goto L291
	} else {
		goto L292
	}
L290:
	;
	goto L286
L291:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1363)&int64(9223372036854775807)) {
		goto L298
	} else {
		goto L299
	}
L292:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v22)+128))
	if v1366 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v1375
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v1377 | int32(1)
	v1381 = v1375
	goto L291
L294:
	;
	v1372 = *(*int64)(unsafe.Add(mBase, uint32(v22)+136))
	v1373 = F_sdsfromlonglong(m, v1372)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L2
	} else {
		goto L297
	}
L295:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	v1370 = F_sdsnewlen(m, v1366, v1369)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L2
	} else {
		goto L296
	}
L296:
	;
	v1375 = v1370
	goto L293
L297:
	;
	v1375 = v1373
	goto L293
L298:
	;
	v1387 = float64(0)
	goto L300
L299:
	;
	v1387 = v1363
	goto L300
L300:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v1393 = F_hashtableFindPositionForInsert(m, v1388, v1381, v22+int32(16), v22+int32(68))
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L2
	} else {
		goto L303
	}
L301:
	;
	v1755 = F_zuiNext(m, v1329, v22+int32(88))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L2
	} else {
		goto L362
	}
L302:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	v1731 = *(*float64)(unsafe.Add(mBase, uint32(v1730)))
	switch v690 + int32(-1) {
	case 0:
		goto L352
	case 1:
		goto L351
	default:
		goto L350
	}
L303:
	;
	if v1393 == int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	if v1398&int32(1) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[6]))
	if int32(311) < v1431 {
		goto L317
	} else {
		goto L318
	}
L306:
	;
	if v1397 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v1398 & int32(-2)
	v1421 = v1397
	goto L305
L308:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v22)+128))
	if v1412 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	v1410 = F_sdsdup(m, v1397)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L2
	} else {
		goto L310
	}
L310:
	;
	v1421 = v1410
	goto L305
L311:
	;
	v1418 = *(*int64)(unsafe.Add(mBase, uint32(v22)+136))
	v1419 = F_sdsfromlonglong(m, v1418)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L2
	} else {
		goto L314
	}
L312:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	v1416 = F_sdsnewlen(m, v1412, v1415)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L2
	} else {
		goto L313
	}
L313:
	;
	v1421 = v1416
	goto L305
L314:
	;
	v1421 = v1419
	goto L305
L315:
	;
	v1664 = int32(1)
	if v1661 == int64(0) {
		goto L331
	} else {
		goto L332
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[6])) = v1636
	v1648 = int64(base.Ui64(v1637)>>(uint(int64(29))%64))&int64(22906492245) ^ v1637
	v1653 = v1648<<(uint(int64(17))%64)&int64(8202884508482404352) ^ v1648
	v1658 = v1653<<(uint(int64(37))%64)&int64(-2270628950310912) ^ v1653
	v1661 = int64(base.Ui64(v1658)>>(uint(int64(43))%64)) ^ v1658
	goto L315
L317:
	;
	if v1431 == int32(313) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	v1440 = *(*int64)(unsafe.Add(mBase, uint32(v1431<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1636 = v1431 + int32(1)
	v1637 = v1440
	goto L316
L319:
	;
	v1524 = int32(0)
	v1527 = v1519
	goto L325
L320:
	;
	v1445 = int64(5489)
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7])) = v1445
	v1453 = int64(1)
	v1455 = v1445
	goto L322
L321:
	;
	v1444 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7]))
	v1519 = v1444
	goto L319
L322:
	;
	v1459 = int32(3)
	v1463 = int64(62)
	v1466 = int64(6364136223846793005)
	v1468 = (int64(base.Ui64(v1455)>>(uint(v1463)%64))^v1455)*v1466 + v1453
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1453)<<(uint(v1459)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1468
	v1471 = v1453 + int64(1)
	v1482 = (int64(base.Ui64(v1468)>>(uint(v1463)%64))^v1468)*v1466 + v1471
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1471)<<(uint(v1459)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1482
	v1485 = v1453 + int64(2)
	v1496 = (int64(base.Ui64(v1482)>>(uint(v1463)%64))^v1482)*v1466 + v1485
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1485)<<(uint(v1459)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1496
	v1499 = v1453 + int64(3)
	if v1499 == int64(312) {
		v1519 = v1445
		goto L319
	} else {
		goto L324
	}
L324:
	;
	v1512 = (int64(base.Ui64(v1496)>>(uint(int64(62))%64))^v1496)*int64(6364136223846793005) + v1499
	*(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1499)<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1512
	v1453 = v1453 + int64(4)
	v1455 = v1512
	goto L322
L325:
	;
	v1533 = int32(3)
	v1534 = v1524 << (uint(v1533) % 32)
	v1537 = int32(1)
	v1538 = v1524 + v1537
	v1543 = *(*int64)(unsafe.Add(mBase, uint32(v1538<<(uint(v1533)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1551 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1543)&v1537<<(uint(v1533)%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(v1534)+uint32(_c_F_zunionInterDiffGenericCommand[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v1534)+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1551 ^ v1554 ^ int64(base.Ui64(v1527&int64(-2147483648)|v1543&int64(2147483646))>>(uint(int64(1))%64))
	if v1538 != int32(156) {
		v1524 = v1538
		v1527 = v1543
		goto L325
	} else {
		goto L327
	}
L326:
	;
	v1566 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[9]))
	v1568 = int32(156)
	v1570 = v1566
	goto L328
L327:
	;
	goto L326
L328:
	;
	v1577 = int32(3)
	v1578 = v1568 << (uint(v1577) % 32)
	v1581 = int32(1)
	v1582 = v1568 + v1581
	v1587 = *(*int64)(unsafe.Add(mBase, uint32(v1582<<(uint(v1577)%32))+uint32(_c_F_zunionInterDiffGenericCommand[7])))
	v1595 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1587)&v1581<<(uint(v1577)%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1598 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+uint32(_c_F_zunionInterDiffGenericCommand[10])))
	*(*int64)(unsafe.Add(mBase, uint32(v1578)+uint32(_c_F_zunionInterDiffGenericCommand[7]))) = v1595 ^ v1598 ^ int64(base.Ui64(v1570&int64(-2147483648)|v1587&int64(2147483646))>>(uint(int64(1))%64))
	if v1582 != int32(311) {
		v1568 = v1582
		v1570 = v1587
		goto L328
	} else {
		goto L330
	}
L329:
	;
	v1609 = int32(1)
	v1610 = int32(0)
	v1612 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[7]))
	v1620 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v1612)&v1609<<(uint(int32(3))%32))+uint32(_c_F_zunionInterDiffGenericCommand[8])))
	v1622 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[11]))
	v1627 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[12])) = v1620 ^ v1622 ^ int64(base.Ui64(v1612&int64(2147483646)|v1627&int64(-2147483648))>>(uint(int64(1))%64))
	v1636 = v1609
	v1637 = v1612
	goto L316
L330:
	;
	goto L329
L331:
	;
	v1670 = int32(32)
	goto L333
L332:
	;
	v1670 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v1661)))>>(uint(v1664)%32)) + v1664
	goto L333
L333:
	;
	v1671 = F_zslCreateNode(m, v1670, v1387, v1421)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L2
	} else {
		goto L334
	}
L334:
	;
	F_sdsfree(m, v1421)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L2
	} else {
		goto L335
	}
L335:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	F_hashtableInsertAtPosition(m, v1675, v1671, v22+int32(16))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L2
	} else {
		goto L336
	}
L336:
	;
	v1681 = v1671 + int32(16)
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1681)))
	v1685 = v1681 + v1682<<(uint(int32(3))%32)
	v1686 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1685))))
	v1687 = v1685 + v1686
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687))))
	v1690 = v1688 & int32(7)
	switch v1690 {
	case 0:
		goto L342
	case 1:
		goto L341
	case 2:
		goto L340
	case 3:
		goto L339
	case 4:
		goto L338
	default:
		goto L301
	}
L337:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v1706 + v1705
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if base.Ui32(v1705) <= base.Ui32(v1709) {
		goto L301
	} else {
		goto L343
	}
L338:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1687+int32(-16))))
	v1705 = v1704
	goto L337
L339:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1687+int32(-8))))
	v1705 = v1701
	goto L337
L340:
	;
	v1698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687+int32(-4)))))
	v1705 = v1698
	goto L337
L341:
	;
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+int32(-2)))))
	v1705 = v1695
	goto L337
L342:
	;
	v1705 = int32(base.Ui32(v1688) >> (uint(int32(3)) % 32))
	goto L337
L343:
	;
	switch v1690 {
	default:
		goto L348
	case 1:
		goto L347
	case 2:
		goto L346
	case 3:
		goto L345
	case 4:
		goto L344
	}
L344:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1687+int32(-16))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1728
	goto L301
L345:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1687+int32(-8))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1724
	goto L301
L346:
	;
	v1720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687+int32(-4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1720
	goto L301
L347:
	;
	v1716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+int32(-2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1716
	goto L301
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = int32(base.Ui32(v1688) >> (uint(int32(3)) % 32))
	goto L301
L349:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1730))) = v1744
	goto L301
L350:
	;
	if base.F64_gt(v1387, v1731) != 0 {
		goto L359
	} else {
		goto L360
	}
L351:
	;
	if base.F64_lt(v1387, v1731) != 0 {
		goto L356
	} else {
		goto L357
	}
L352:
	;
	v1733 = base.F64_add(v1387, v1731)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1733)&int64(9223372036854775807)) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1739 = float64(0)
	goto L355
L354:
	;
	v1739 = v1733
	goto L355
L355:
	;
	v1744 = v1739
	goto L349
L356:
	;
	v1741 = v1387
	goto L358
L357:
	;
	v1741 = v1731
	goto L358
L358:
	;
	v1744 = v1741
	goto L349
L359:
	;
	v1743 = v1387
	goto L361
L360:
	;
	v1743 = v1731
	goto L361
L361:
	;
	v1744 = v1743
	goto L349
L362:
	;
	if v1755 != 0 {
		goto L289
	} else {
		goto L363
	}
L363:
	;
	goto L290
L364:
	;
	goto L282
L365:
	;
	goto L281
L366:
	;
	v1845 = F_hashtableNext(m, v22+int32(16), v22+int32(68))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L2
	} else {
		goto L371
	}
L367:
	;
	goto L366
L368:
	;
	goto L367
L370:
	;
	F_hashtableCleanupIterator(m, v22+int32(16))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L2
	} else {
		goto L378
	}
L371:
	;
	if v1845 == int32(0) {
		goto L370
	} else {
		goto L372
	}
L372:
	;
	goto L373
L373:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	v1870 = F_zslInsertNode(m, v1868, v1869)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L2
	} else {
		goto L375
	}
L374:
	;
	goto L370
L375:
	;
	v1876 = F_hashtableNext(m, v22+int32(16), v22+int32(68))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L2
	} else {
		goto L376
	}
L376:
	;
	if v1876 != 0 {
		goto L373
	} else {
		goto L377
	}
L377:
	;
	goto L374
L378:
	;
	goto L272
L379:
	;
	goto L272
L380:
	;
	if l4 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L381:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1948)))
	if v1949 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1988 = F_dbDelete(m, v1987, l1)
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L2
	} else {
		goto L396
	}
L383:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	F_zsetConvertToListpackIfNeeded(m, v1952, v1953, v1954)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L2
	} else {
		goto L384
	}
L384:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_setKey(m, l0, v1957, l1, v22+int32(76), int32(0))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L2
	} else {
		goto L385
	}
L385:
	;
	if l3 == int32(2) {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1968 = int32(_a_F_zunionInterDiffGenericCommand_11)
	goto L388
L387:
	;
	v1968 = int32(_a_F_zunionInterDiffGenericCommand_12)
	goto L388
L388:
	;
	if l3 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1970 = v1968
	goto L391
L390:
	;
	v1970 = int32(_a_F_zunionInterDiffGenericCommand_13)
	goto L391
L391:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+28))
	F_notifyKeyspaceEvent(m, int32(128), v1970, l1, v1972)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L2
	} else {
		goto L392
	}
L392:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v1976 = F_zsetLength(m, v1975)
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L2
	} else {
		goto L393
	}
L393:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v1976))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L2
	} else {
		goto L394
	}
L394:
	;
	v1981 = int32(_a_F_zunionInterDiffGenericCommand_14)
	v1983 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[13])) = v1983 + int64(1)
	goto L39
L395:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[14]))
	F_addReply(m, l0, v2008)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L2
	} else {
		goto L400
	}
L396:
	;
	if v1988 == int32(0) {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_signalModifiedKey(m, l0, v1992, l1)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L2
	} else {
		goto L398
	}
L398:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_zunionInterDiffGenericCommand_15), l1, v1998)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L2
	} else {
		goto L399
	}
L399:
	;
	v2001 = int32(_a_F_zunionInterDiffGenericCommand_14)
	v2003 = *(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[13])) = v2003 + int64(1)
	goto L395
L400:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	F_decrRefCount(m, v2011)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L2
	} else {
		goto L401
	}
L401:
	;
	goto L39
L402:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+12))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2018)))
	if v687 != 0 {
		v2025 = v2020
		goto L405
	} else {
		goto L406
	}
L403:
	;
	F_addReplyLongLong(m, l0, v1942)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L2
	} else {
		goto L404
	}
L404:
	;
	goto L39
L405:
	;
	F_addReplyArrayLen(m, l0, v2025)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L2
	} else {
		goto L407
	}
L406:
	;
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v2025 = v2020 << (uint(base.B2i32(v2021 == int32(2))) % 32)
	goto L405
L407:
	;
	if v2019 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v2146 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[15]))
	if v2146 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L409:
	;
	v2037 = v2019
	goto L410
L410:
	;
	if v687 != 0 {
		goto L413
	} else {
		goto L414
	}
L411:
	;
	goto L408
L412:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v2037)+12))
	if v2124 != 0 {
		v2037 = v2124
		goto L410
	} else {
		goto L433
	}
L413:
	;
	v2090 = v2037 + int32(16)
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2090)))
	v2094 = v2090 + v2091<<(uint(int32(3))%32)
	v2095 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2094))))
	v2096 = v2094 + v2095
	v2100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2096))))
	switch v2100 & int32(7) {
	case 0:
		goto L431
	case 1:
		goto L430
	case 2:
		goto L429
	case 3:
		goto L428
	case 4:
		goto L427
	default:
		v2117 = int32(0)
		goto L426
	}
L414:
	;
	v2049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if base.Ui32(v2049) < base.Ui32(int32(3)) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v2056 = v2037 + int32(16)
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v2056)))
	v2060 = v2056 + v2057<<(uint(int32(3))%32)
	v2061 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2060))))
	v2062 = v2060 + v2061
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2062))))
	switch v2066 & int32(7) {
	case 0:
		goto L423
	case 1:
		goto L422
	case 2:
		goto L421
	case 3:
		goto L420
	case 4:
		goto L419
	default:
		v2083 = int32(0)
		goto L418
	}
L416:
	;
	F_addReplyArrayLen(m, l0, int32(2))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L2
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	F_addReplyBulkCBuffer(m, l0, v2062+int32(1), v2083)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L2
	} else {
		goto L424
	}
L419:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2062+int32(-16))))
	v2083 = v2082
	goto L418
L420:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v2062+int32(-8))))
	v2083 = v2079
	goto L418
L421:
	;
	v2076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2062+int32(-4)))))
	v2083 = v2076
	goto L418
L422:
	;
	v2073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2062+int32(-2)))))
	v2083 = v2073
	goto L418
L423:
	;
	v2083 = int32(base.Ui32(v2066) >> (uint(int32(3)) % 32))
	goto L418
L424:
	;
	v2086 = *(*float64)(unsafe.Add(mBase, uint32(v2037)))
	F_addReplyDouble(m, l0, v2086)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L2
	} else {
		goto L425
	}
L425:
	;
	goto L412
L426:
	;
	F_addReplyBulkCBuffer(m, l0, v2096+int32(1), v2117)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L2
	} else {
		goto L432
	}
L427:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2096+int32(-16))))
	v2117 = v2116
	goto L426
L428:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2096+int32(-8))))
	v2117 = v2113
	goto L426
L429:
	;
	v2110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2096+int32(-4)))))
	v2117 = v2110
	goto L426
L430:
	;
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2096+int32(-2)))))
	v2117 = v2107
	goto L426
L431:
	;
	v2117 = int32(base.Ui32(v2100) >> (uint(int32(3)) % 32))
	goto L426
L432:
	;
	goto L412
L433:
	;
	goto L411
L434:
	;
	F_decrRefCount(m, v2144)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L2
	} else {
		goto L437
	}
L435:
	;
	F_freeObjAsync(m, int32(0), v2144, int32(-1))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L2
	} else {
		goto L436
	}
L436:
	;
	goto L39
L437:
	;
	goto L39
L438:
	;
	goto L1
L439:
	;
	goto L1
L440:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, _c_F_zunionInterDiffGenericCommand[0]))
	F_addReplyErrorObject(m, l0, v2183)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L2
	} else {
		goto L441
	}
L441:
	;
	goto L1
}
