package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___wake_1(m *base.Module, l0 int32) {
	return
}
func F___wake_2(m *base.Module, l0 int32) {
	return
}
func F_waitForClientIO(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
	if v2 != 0 {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
		if v6 != int32(1) {
		} else {
			for {
				v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
				if v10 == int32(1) {
					continue
				} else {
					break
				}
				break
			}
		}
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
		if v14 != int32(1) {
		} else {
			for {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
				if v18 == int32(1) {
					continue
				} else {
					break
				}
				break
			}
		}
	} else {
		v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
		if v3 == int32(0) {
		} else {
			v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
			if v6 != int32(1) {
			} else {
				for {
					v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)))
					if v10 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
			if v14 != int32(1) {
			} else {
				for {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)))
					if v18 == int32(1) {
						continue
					} else {
						break
					}
					break
				}
			}
		}
	}
	return
}
func F_watchdogScheduleSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = m.G0
	v5 = int32(32)
	v6 = v4 - v5
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(0)
	v12 = int32(1000)
	v13 = base.I32_div_s(l0, v12)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = base.I64_extend_i32_s(v13)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = (l0 - v13*v12) * v12
	v24 = F_setitimer(m, v2, v6, v2)
	mBase = m.M
	m.G0 = v6 + v5
	return
}
func F_watchedKeyGetKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	return v2
}
func F_wcslen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v5 = l0
	for {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v9 != 0 {
			v5 = v5 + int32(4)
			continue
		} else {
			break
		}
		break
	}
	return (v5 - l0) >> (uint(int32(2)) % 32)
}
func F_whileBlockedCron(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v36 int64
	_ = v36
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	v7 = *(*int64)(unsafe.Add(mBase, _consts[488]))
	if v7 == int64(0) {
		F__serverAssert(m, int32(_a1247), int32(_a1240), int32(1752))
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, _consts[32]))
		if v11 <= v7 {
			return
		} else {
			v13 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[148]))
			v19 = *(*int32)(unsafe.Add(mBase, _consts[149]))
			v20 = base.I32_div_s(int32(1000), v19)
			v26 = base.I64_div_s(v11-v7+base.I64_extend_i32_s(v20+int32(-1)), base.I64_extend_i32_s(v20))
			v27 = base.I32_wrap_i64(v26)
			*(*int32)(unsafe.Add(mBase, _consts[148])) = v15 + v27
			*(*int64)(unsafe.Add(mBase, _consts[488])) = v7 + base.I64_extend_i32_s(v20*v27)
			v36 = *(*int64)(unsafe.Add(mBase, _consts[270]))
			if base.B2i32(v36 == int64(0)) == v13 {
				v42 = F_ustime(m)
				mBase = m.M
				v43 = v42
			} else {
				v43 = int64(0)
			}
			v44 = int32(0)
			v45 = *(*int32)(unsafe.Add(mBase, _consts[116]))
			if v45 == v44 {
				v51 = *(*int64)(unsafe.Add(mBase, _consts[270]))
				if v51 == int64(0) {
					v69 = int32(0)
					v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
					if v70 == v69 {
						return
					} else {
						v73 = int32(0)
						v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
						if v74 == v73 {
							return
						} else {
							v79 = F_prepareForShutdown(m, int32(0), int32(2))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								if v79 == int32(0) {
									m.Env.Exit(m, int32(0))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if int32(3) < v84 {
										v92 = int32(0)
										*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
										*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
										return
									} else {
										F__serverLog(m, int32(3), int32(_a1248), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v92 = int32(0)
											*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
											*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
											return
										}
									}
								}
							}
						}
					}
				} else {
					v54 = F_ustime(m)
					mBase = m.M
					v56 = *(*int64)(unsafe.Add(mBase, _consts[270]))
					if v56 == int64(0) {
						v69 = int32(0)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
						if v70 == v69 {
							return
						} else {
							v73 = int32(0)
							v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
							if v74 == v73 {
								return
							} else {
								v79 = F_prepareForShutdown(m, int32(0), int32(2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									if v79 == int32(0) {
										m.Env.Exit(m, int32(0))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v84 {
											v92 = int32(0)
											*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
											*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
											return
										} else {
											F__serverLog(m, int32(3), int32(_a1248), int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v92 = int32(0)
												*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
												*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
												return
											}
										}
									}
								}
							}
						}
					} else {
						v59 = v54 - v43
						if v59 < v56*int64(1000) {
							v69 = int32(0)
							v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
							if v70 == v69 {
								return
							} else {
								v73 = int32(0)
								v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
								if v74 == v73 {
									return
								} else {
									v79 = F_prepareForShutdown(m, int32(0), int32(2))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										if v79 == int32(0) {
											m.Env.Exit(m, int32(0))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(3) < v84 {
												v92 = int32(0)
												*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
												*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
												return
											} else {
												F__serverLog(m, int32(3), int32(_a1248), int32(0))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return
												} else {
													v92 = int32(0)
													*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
													*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
													return
												}
											}
										}
									}
								}
							}
						} else {
							F_latencyAddSample(m, int32(_a1249), v59)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v69 = int32(0)
								v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
								if v70 == v69 {
									return
								} else {
									v73 = int32(0)
									v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
									if v74 == v73 {
										return
									} else {
										v79 = F_prepareForShutdown(m, int32(0), int32(2))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											if v79 == int32(0) {
												m.Env.Exit(m, int32(0))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
												if int32(3) < v84 {
													v92 = int32(0)
													*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
													*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
													return
												} else {
													F__serverLog(m, int32(3), int32(_a1248), int32(0))
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														v92 = int32(0)
														*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
														*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
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
				F_cronUpdateMemoryStats(m)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v51 = *(*int64)(unsafe.Add(mBase, _consts[270]))
					if v51 == int64(0) {
						v69 = int32(0)
						v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
						if v70 == v69 {
							return
						} else {
							v73 = int32(0)
							v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
							if v74 == v73 {
								return
							} else {
								v79 = F_prepareForShutdown(m, int32(0), int32(2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									if v79 == int32(0) {
										m.Env.Exit(m, int32(0))
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(3) < v84 {
											v92 = int32(0)
											*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
											*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
											return
										} else {
											F__serverLog(m, int32(3), int32(_a1248), int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v92 = int32(0)
												*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
												*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
												return
											}
										}
									}
								}
							}
						}
					} else {
						v54 = F_ustime(m)
						mBase = m.M
						v56 = *(*int64)(unsafe.Add(mBase, _consts[270]))
						if v56 == int64(0) {
							v69 = int32(0)
							v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
							if v70 == v69 {
								return
							} else {
								v73 = int32(0)
								v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
								if v74 == v73 {
									return
								} else {
									v79 = F_prepareForShutdown(m, int32(0), int32(2))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										if v79 == int32(0) {
											m.Env.Exit(m, int32(0))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(3) < v84 {
												v92 = int32(0)
												*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
												*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
												return
											} else {
												F__serverLog(m, int32(3), int32(_a1248), int32(0))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return
												} else {
													v92 = int32(0)
													*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
													*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
													return
												}
											}
										}
									}
								}
							}
						} else {
							v59 = v54 - v43
							if v59 < v56*int64(1000) {
								v69 = int32(0)
								v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
								if v70 == v69 {
									return
								} else {
									v73 = int32(0)
									v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
									if v74 == v73 {
										return
									} else {
										v79 = F_prepareForShutdown(m, int32(0), int32(2))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											if v79 == int32(0) {
												m.Env.Exit(m, int32(0))
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
												if int32(3) < v84 {
													v92 = int32(0)
													*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
													*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
													return
												} else {
													F__serverLog(m, int32(3), int32(_a1248), int32(0))
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														v92 = int32(0)
														*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
														*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
														return
													}
												}
											}
										}
									}
								}
							} else {
								F_latencyAddSample(m, int32(_a1249), v59)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v69 = int32(0)
									v70 = *(*int32)(unsafe.Add(mBase, _consts[686]))
									if v70 == v69 {
										return
									} else {
										v73 = int32(0)
										v74 = *(*int32)(unsafe.Add(mBase, _consts[116]))
										if v74 == v73 {
											return
										} else {
											v79 = F_prepareForShutdown(m, int32(0), int32(2))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												if v79 == int32(0) {
													m.Env.Exit(m, int32(0))
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, _consts[15]))
													if int32(3) < v84 {
														v92 = int32(0)
														*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
														*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
														return
													} else {
														F__serverLog(m, int32(3), int32(_a1248), int32(0))
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return
														} else {
															v92 = int32(0)
															*(*int32)(unsafe.Add(mBase, _consts[687])) = v92
															*(*int32)(unsafe.Add(mBase, _consts[686])) = v92
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
func F_wholeCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(1)
	v15 = m.G6
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+276))
	if v16 < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v134 = m.G14
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	m.T0[v135].(func(*base.Module))(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L22
	}
L2:
	;
	v23 = v14
	v24 = v16
	goto L3
L3:
	;
	v29 = int32(1) - v23
	v31 = v29 >> (uint(int32(31)) % 32)
	if base.Ui32(int32(1000000)) < base.Ui32(v29^v31-v31) {
		v118 = v24
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	if v23 < v118 {
		v23 = v23 + int32(1)
		v24 = v118
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v36 = m.G3
	v37 = int32(0)
	v38 = m.G6
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+272))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+v23<<(uint(int32(2))%32)+int32(-4))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+260))
	if v37 < v46 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v23
	v92 = m.G6
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+280))
	if v93 == v23 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v59 = v37
	goto L11
L9:
	;
	v86 = v36 + int32(_a1718)
	v87 = v36 + int32(_a1717)
	goto L7
L10:
	;
	v86 = v62 + int32(_a1720)
	v87 = v62 + int32(_a1719)
	goto L7
L11:
	;
	v62 = m.G3
	v63 = m.G6
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v59<<(uint(int32(2))%32))+4))
	if v67 == v23 {
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v86 = v69 + int32(_a1718)
	v87 = v69 + int32(_a1717)
	goto L7
L13:
	;
	v69 = m.G3
	v71 = v59 + int32(1)
	if v71 != v46 {
		v59 = v71
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v95 = v87
	goto L17
L16:
	;
	v95 = v86
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v95
	v97 = m.G3
	v98 = m.G15
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = m.G12
	v104 = m.T0[v99].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v97+int32(_a1721), v12)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	m.T0[v109].(func(*base.Module, int32, int32))(m, v104, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v92)+276))
	v118 = v112
	goto L5
L21:
	;
	goto L4
L22:
	;
	m.G0 = v12 + int32(16)
	return int32(1)
}
