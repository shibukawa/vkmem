package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_kvstoreBuckets(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v8 != 0 {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
			if v14 == int32(255) {
				v18 = int32(0)
			} else {
				v18 = int32(1) << (uint(v14) % 32)
			}
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
			if v21 == int32(255) {
				v25 = int32(0)
			} else {
				v25 = int32(1) << (uint(v21) % 32)
			}
			return v18 + v25
		} else {
			return int32(0)
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		return v5
	}
}
func F_kvstoreGetHashtable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3+l1<<(uint(int32(2))%32))))
	return v7
}
func F_kvstoreHashtableDelete(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+l1<<(uint(int32(2))%32))))
	if v12 == v4 {
		v68 = v4
		return v68
	} else {
		v15 = F_hashtableDelete(m, v12, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v68 = v4
				return v68
			} else {
				F_cumulativeKeyCountAdd(m, l0, l1, int32(-1))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(1)
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v25&int32(2) == int32(0) {
						v68 = v24
						return v68
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32))))
						if v34 == int32(0) {
							v68 = v24
							return v68
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
							if v37+v38 != 0 {
								v68 = v24
								return v68
							} else {
								v40 = int32(0)
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v43 = l1 << (uint(int32(2)) % 32)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v43)))
								if v45 == v40 {
									v56 = v40
									F_hashtableRelease(m, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v59+v43))) = int32(0)
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 + int32(-1)
										v68 = v24
										return v68
									}
								} else {
									v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+26)))
									if int32(0) < v48 {
										v68 = v24
										return v68
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32))))
										v56 = v55
										F_hashtableRelease(m, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v59+v43))) = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 + int32(-1)
											v68 = v24
											return v68
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
func F_kvstoreHashtableExpand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	if l2 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = l1 << (uint(int32(2)) % 32)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9)))
		if v11 != 0 {
			v38 = v11
			v39 = F_hashtableExpand(m, v38, l2)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = F_hashtableCreate(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17+v9))) = v13
				*(*int32)(unsafe.Add(mBase, uint32(v13+int32(44))+4)) = l0
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23+v9)))
				v26 = F_hashtableMemUsage(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v26 + v28
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v31 + int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35+v9)))
					v38 = v37
					v39 = F_hashtableExpand(m, v38, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v39
					}
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtableFairRandomEntry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4+l1<<(uint(int32(2))%32))))
	if v8 != 0 {
		v11 = F_hashtableFairRandomEntry(m, v8, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v11
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreHashtablePop(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+l1<<(uint(int32(2))%32))))
	if v12 == v5 {
		v69 = v5
		return v69
	} else {
		v15 = F_hashtablePop(m, v12, l2, l3)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v69 = v5
				return v69
			} else {
				F_cumulativeKeyCountAdd(m, l0, l1, int32(-1))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(1)
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v25&int32(2) == int32(0) {
						v69 = v24
						return v69
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32))))
						if v34 == int32(0) {
							v69 = v24
							return v69
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
							if v37+v38 != 0 {
								v69 = v24
								return v69
							} else {
								v40 = int32(0)
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v43 = l1 << (uint(int32(2)) % 32)
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v43)))
								if v45 == v40 {
									v56 = v40
									F_hashtableRelease(m, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v59+v43))) = int32(0)
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 + int32(-1)
										v69 = v24
										return v69
									}
								} else {
									v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+26)))
									if int32(0) < v48 {
										v69 = v24
										return v69
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+l1<<(uint(int32(2))%32))))
										v56 = v55
										F_hashtableRelease(m, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v59+v43))) = int32(0)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63 + int32(-1)
											v69 = v24
											return v69
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
func F_kvstoreHashtableTwoPhasePopDelete(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = l1 << (uint(int32(2)) % 32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8)))
	F_hashtableTwoPhasePopDelete(m, v10, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_cumulativeKeyCountAdd(m, l0, l1, int32(-1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v16&int32(2) == int32(0) {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21+v8)))
				if v23 == int32(0) {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					if v26+v27 != 0 {
						return
					} else {
						v29 = int32(0)
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v32 = l1 << (uint(int32(2)) % 32)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v32)))
						if v34 == v29 {
							v45 = v29
							F_hashtableRelease(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v48+v32))) = int32(0)
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52 + int32(-1)
								return
							}
						} else {
							v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+26)))
							if int32(0) < v37 {
								return
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+l1<<(uint(int32(2))%32))))
								v45 = v44
								F_hashtableRelease(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v48+v32))) = int32(0)
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v52 + int32(-1)
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
func F_kvstoreHashtableTwoPhasePopFindRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5+l1<<(uint(int32(2))%32))))
	if v9 != 0 {
		v12 = F_hashtableTwoPhasePopFindRef(m, v9, l2, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		return int32(0)
	}
}
func F_kvstoreImportingSize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	return v2
}
func F_kvstoreIteratorRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	F_hashtableCleanupIterator(m, l0+int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v12 == int64(-1) {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			if v64 == int32(0) {
				F_valkey_free(m, l0)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					return
				}
			} else {
				F_hashtableReleaseIterator(m, v64)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_valkey_free(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v16&int32(2) == int32(0) {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				if v64 == int32(0) {
					F_valkey_free(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						return
					}
				} else {
					F_hashtableReleaseIterator(m, v64)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						F_valkey_free(m, l0)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v22 = base.I32_wrap_i64(v12)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32))))
				if v26 == int32(0) {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					if v64 == int32(0) {
						F_valkey_free(m, l0)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							return
						}
					} else {
						F_hashtableReleaseIterator(m, v64)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_valkey_free(m, l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					if v29+v30 != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
						if v64 == int32(0) {
							F_valkey_free(m, l0)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								return
							}
						} else {
							F_hashtableReleaseIterator(m, v64)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_valkey_free(m, l0)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v32 = int32(0)
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						v35 = v22 << (uint(int32(2)) % 32)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v35)))
						if v37 == v32 {
							v48 = v32
							F_hashtableRelease(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v51+v35))) = int32(0)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v55 + int32(-1)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								if v64 == int32(0) {
									F_valkey_free(m, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										return
									}
								} else {
									F_hashtableReleaseIterator(m, v64)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_valkey_free(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+26)))
							if int32(0) < v40 {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
								if v64 == int32(0) {
									F_valkey_free(m, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										return
									}
								} else {
									F_hashtableReleaseIterator(m, v64)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_valkey_free(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v22<<(uint(int32(2))%32))))
								v48 = v47
								F_hashtableRelease(m, v48)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v51+v35))) = int32(0)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v55 + int32(-1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
									if v64 == int32(0) {
										F_valkey_free(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											return
										}
									} else {
										F_hashtableReleaseIterator(m, v64)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_valkey_free(m, l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
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
func F_kvstoreMemUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v10 = v3 + v5*int32(12) + int32(80)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v11 == int32(0) {
		v20 = v10
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v20 = v10 + v14<<(uint(int32(3))%32) + int32(8)
	}
	return v20
}
func F_kvstoreNumNonEmptyHashtables(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	return v2
}
