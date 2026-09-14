package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_lookupKeyReadOrReply(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v6 = F_lookupKey(m, v4, l1, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			return v6
		} else {
			F_addReplyOrErrorObject(m, l0, l2)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F_lookupKeyWrite(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_lookupKey(m, l0, l1, int32(8))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lookupKeyWriteWithFlags(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_lookupKey(m, l0, l1, l2|int32(8))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_setKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	if l4&int32(4) != 0 {
		F_dbSetValue(m, l1, l2, l3, int32(1), int32(0))
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if l4&int32(1) != 0 {
				if l4&int32(2) != 0 {
					return
				} else {
					F_touchWatchedKey(m, l1, l2)
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_trackingInvalidateKey(m, l0, l2, int32(1))
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v27 = F_removeExpire(m, l1, l2)
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if l4&int32(2) != 0 {
						return
					} else {
						F_touchWatchedKey(m, l1, l2)
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_trackingInvalidateKey(m, l0, l2, int32(1))
							v35 = m.ExcPending
							if v35 != 0 {
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
		if l4&int32(16) != 0 {
			F_dbAddInternal(m, l1, l2, l3, int32(1))
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if l4&int32(1) != 0 {
					if l4&int32(2) != 0 {
						return
					} else {
						F_touchWatchedKey(m, l1, l2)
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_trackingInvalidateKey(m, l0, l2, int32(1))
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v27 = F_removeExpire(m, l1, l2)
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if l4&int32(2) != 0 {
							return
						} else {
							F_touchWatchedKey(m, l1, l2)
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_trackingInvalidateKey(m, l0, l2, int32(1))
								v35 = m.ExcPending
								if v35 != 0 {
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
			if l4&int32(8) != 0 {
				F_dbAddInternal(m, l1, l2, l3, int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					if l4&int32(1) != 0 {
						if l4&int32(2) != 0 {
							return
						} else {
							F_touchWatchedKey(m, l1, l2)
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_trackingInvalidateKey(m, l0, l2, int32(1))
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v27 = F_removeExpire(m, l1, l2)
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							if l4&int32(2) != 0 {
								return
							} else {
								F_touchWatchedKey(m, l1, l2)
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_trackingInvalidateKey(m, l0, l2, int32(1))
									v35 = m.ExcPending
									if v35 != 0 {
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
				v13 = F_lookupKey(m, l1, l2, int32(8))
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					if v13 != 0 {
						F_dbSetValue(m, l1, l2, l3, int32(1), int32(0))
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l4&int32(1) != 0 {
								if l4&int32(2) != 0 {
									return
								} else {
									F_touchWatchedKey(m, l1, l2)
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										F_trackingInvalidateKey(m, l0, l2, int32(1))
										v35 = m.ExcPending
										if v35 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v27 = F_removeExpire(m, l1, l2)
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									if l4&int32(2) != 0 {
										return
									} else {
										F_touchWatchedKey(m, l1, l2)
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											F_trackingInvalidateKey(m, l0, l2, int32(1))
											v35 = m.ExcPending
											if v35 != 0 {
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
						F_dbAddInternal(m, l1, l2, l3, int32(0))
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							if l4&int32(1) != 0 {
								if l4&int32(2) != 0 {
									return
								} else {
									F_touchWatchedKey(m, l1, l2)
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										F_trackingInvalidateKey(m, l0, l2, int32(1))
										v35 = m.ExcPending
										if v35 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v27 = F_removeExpire(m, l1, l2)
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									if l4&int32(2) != 0 {
										return
									} else {
										F_touchWatchedKey(m, l1, l2)
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											F_trackingInvalidateKey(m, l0, l2, int32(1))
											v35 = m.ExcPending
											if v35 != 0 {
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
