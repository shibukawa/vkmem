package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F__serverAssert(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v121 int32
	_ = v121
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	goto L1
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	goto L11
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[297])) = uint8(v31)
	goto L2
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[296]))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v22 = int32(_a582)
	goto L8
L7:
	;
	v22 = int32(_a583)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v22
	F__serverLog(m, int32(1027), int32(_a584), v8+int32(32))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v36 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v13|base.B2i32(v61 == int32(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v13 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = int32(_a589)
	goto L16
L15:
	;
	v41 = int32(_a320)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v41
	F__serverLog(m, int32(3), int32(_a590), v8+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v50 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	F__serverLog(m, int32(3), int32(_a591), v8)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(40)))) = int64(0)
	goto L34
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[299])) = int32(1)
	F_logServerInfo(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	F_logCurrentClient(m, v71, int32(_a585))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[300]))
	F_logCurrentClient(m, v76, int32(_a586))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_serverLogRaw(m, int32(1027), int32(_a587))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v85 = F_sdsempty(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v87 = int32(0)
	v90 = F_modulesCollectInfo(m, v85, v87, int32(1), v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_serverLogRaw(m, int32(1027), v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_sdsfree(m, v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v96 = F_getConfigDebugInfo(m)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_serverLogRaw(m, int32(1027), int32(_a588))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	F_serverLogRaw(m, int32(1027), v96)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	F_sdsfree(m, v96)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L20
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = int32(-1073741824)
	v121 = v8 + int32(36)
	goto L36
L35:
	;
	v149 = v8 + int32(36)
	goto L43
L36:
	;
	goto L38
L38:
	;
	if v121 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	goto L35
L41:
	;
	v144 = F___memcpy(m, int32(9118724), v121, int32(140))
	mBase = m.M
	goto L40
L42:
	;
	v177 = v8 + int32(36)
	goto L50
L43:
	;
	goto L45
L45:
	;
	if v149 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	goto L42
L48:
	;
	v172 = F___memcpy(m, int32(9118164), v149, int32(140))
	mBase = m.M
	goto L47
L49:
	;
	v205 = v8 + int32(36)
	goto L57
L50:
	;
	goto L52
L52:
	;
	if v177 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	goto L49
L55:
	;
	v200 = F___memcpy(m, int32(9118304), v177, int32(140))
	mBase = m.M
	goto L54
L56:
	;
	v233 = v8 + int32(36)
	goto L64
L57:
	;
	goto L59
L59:
	;
	if v205 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	goto L56
L62:
	;
	v228 = F___memcpy(m, int32(9117744), v205, int32(140))
	mBase = m.M
	goto L61
L63:
	;
	v259 = int32(0)
	F_bugReportEnd(m, v259, v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L70
	}
L64:
	;
	goto L66
L66:
	;
	if v233 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	goto L63
L69:
	;
	v256 = F___memcpy(m, int32(9118024), v233, int32(140))
	mBase = m.M
	goto L68
L70:
	;
	m.G0 = v8 + int32(176)
	return
}
func F__serverAssertPrintObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	v3 = m.G0
	v5 = v3 - int32(64)
	m.G0 = v5
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v10 != 0 {
		v33 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(3) < v33 {
			m.G0 = v5 + int32(64)
			return
		} else {
			F__serverLog(m, int32(3), int32(_a592), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(3) < v42 {
					m.G0 = v5 + int32(64)
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v45 & int32(15)
					F__serverLog(m, int32(3), int32(_a593), v5+int32(32))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v56 {
							m.G0 = v5 + int32(64)
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(base.Ui32(v59)>>(uint(int32(4))%32)) & int32(15)
							F__serverLog(m, int32(3), int32(_a594), v5+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(3) < v72 {
									m.G0 = v5 + int32(64)
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v76 = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(base.Ui32(v75) >> (uint(v76) % 32))
									F__serverLog(m, v76, int32(_a595), v5)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										m.G0 = v5 + int32(64)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(3) < v12 {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[297])) = uint8(v28)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			if int32(3) < v33 {
				m.G0 = v5 + int32(64)
				return
			} else {
				F__serverLog(m, int32(3), int32(_a592), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(3) < v42 {
						m.G0 = v5 + int32(64)
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v45 & int32(15)
						F__serverLog(m, int32(3), int32(_a593), v5+int32(32))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(3) < v56 {
								m.G0 = v5 + int32(64)
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(base.Ui32(v59)>>(uint(int32(4))%32)) & int32(15)
								F__serverLog(m, int32(3), int32(_a594), v5+int32(16))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									v72 = *(*int32)(unsafe.Add(mBase, _consts[28]))
									if int32(3) < v72 {
										m.G0 = v5 + int32(64)
										return
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v76 = int32(3)
										*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(base.Ui32(v75) >> (uint(v76) % 32))
										F__serverLog(m, v76, int32(_a595), v5)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											m.G0 = v5 + int32(64)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[296]))
			if v18 != 0 {
				v19 = int32(_a582)
			} else {
				v19 = int32(_a583)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v19
			F__serverLog(m, int32(1027), int32(_a584), v5+int32(48))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _consts[297])) = uint8(v28)
				v33 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(3) < v33 {
					m.G0 = v5 + int32(64)
					return
				} else {
					F__serverLog(m, int32(3), int32(_a592), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						if int32(3) < v42 {
							m.G0 = v5 + int32(64)
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v45 & int32(15)
							F__serverLog(m, int32(3), int32(_a593), v5+int32(32))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(3) < v56 {
									m.G0 = v5 + int32(64)
									return
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(base.Ui32(v59)>>(uint(int32(4))%32)) & int32(15)
									F__serverLog(m, int32(3), int32(_a594), v5+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, _consts[28]))
										if int32(3) < v72 {
											m.G0 = v5 + int32(64)
											return
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v76 = int32(3)
											*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(base.Ui32(v75) >> (uint(v76) % 32))
											F__serverLog(m, v76, int32(_a595), v5)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												m.G0 = v5 + int32(64)
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
func F__serverAssertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		if l1 == int32(0) {
			F__serverAssert(m, l2, l3, l4)
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		} else {
			F__serverAssertPrintObject(m, l1)
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F__serverAssert(m, l2, l3, l4)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F__serverAssertPrintClientInfo(m, l0)
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			if l1 == int32(0) {
				F__serverAssert(m, l2, l3, l4)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			} else {
				F__serverAssertPrintObject(m, l1)
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					F__serverAssert(m, l2, l3, l4)
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F__serverPanic_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	v6 = m.G0
	v8 = v6 - int32(432)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+288)) = l3
	v14 = F_vsnprintf(m, v8+int32(32), int32(256), l2, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[297])))
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L12
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v21 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[297])) = uint8(v37)
	goto L4
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[296]))
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = int32(_a582)
	goto L10
L9:
	;
	v28 = int32(_a583)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v28
	F__serverLog(m, int32(1027), int32(_a584), v8+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v42 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v19|base.B2i32(v73 == int32(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	F__serverLog(m, int32(3), int32(_a579), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v51 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	F__serverLog(m, int32(3), int32(_a580), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v60 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
	F__serverLog(m, int32(3), int32(_a581), v8)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(296)))) = int64(0)
	goto L23
L21:
	;
	F_printCrashReport(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+292)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+424)) = int32(-1073741824)
	v92 = v8 + int32(292)
	goto L25
L24:
	;
	v120 = v8 + int32(292)
	goto L32
L25:
	;
	goto L27
L27:
	;
	if v92 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	goto L24
L30:
	;
	v115 = F___memcpy(m, int32(9118724), v92, int32(140))
	mBase = m.M
	goto L29
L31:
	;
	v148 = v8 + int32(292)
	goto L39
L32:
	;
	goto L34
L34:
	;
	if v120 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	goto L31
L37:
	;
	v143 = F___memcpy(m, int32(9118164), v120, int32(140))
	mBase = m.M
	goto L36
L38:
	;
	v176 = v8 + int32(292)
	goto L46
L39:
	;
	goto L41
L41:
	;
	if v148 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	goto L38
L44:
	;
	v171 = F___memcpy(m, int32(9118304), v148, int32(140))
	mBase = m.M
	goto L43
L45:
	;
	v204 = v8 + int32(292)
	goto L53
L46:
	;
	goto L48
L48:
	;
	if v176 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	goto L45
L51:
	;
	v199 = F___memcpy(m, int32(9117744), v176, int32(140))
	mBase = m.M
	goto L50
L52:
	;
	v230 = int32(0)
	F_bugReportEnd(m, v230, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L53:
	;
	goto L55
L55:
	;
	if v204 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	goto L52
L58:
	;
	v227 = F___memcpy(m, int32(9118024), v204, int32(140))
	mBase = m.M
	goto L57
L59:
	;
	m.G0 = v8 + int32(432)
	return
}
func F_freeServerClientMemUsageBuckets(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v1 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	if v3 == v1 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_listRelease(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[745]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			F_listRelease(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[745]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				F_listRelease(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _consts[745]))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
					F_listRelease(m, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _consts[745]))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
						F_listRelease(m, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _consts[745]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
							F_listRelease(m, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, _consts[745]))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
								F_listRelease(m, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, _consts[745]))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
									F_listRelease(m, v41)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, _consts[745]))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+64))
										F_listRelease(m, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, _consts[745]))
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
											F_listRelease(m, v51)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, _consts[745]))
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+80))
												F_listRelease(m, v56)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													v60 = *(*int32)(unsafe.Add(mBase, _consts[745]))
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
													F_listRelease(m, v61)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														v65 = *(*int32)(unsafe.Add(mBase, _consts[745]))
														v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+96))
														F_listRelease(m, v66)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v70 = *(*int32)(unsafe.Add(mBase, _consts[745]))
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+104))
															F_listRelease(m, v71)
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+112))
																F_listRelease(m, v76)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return
																} else {
																	v80 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+120))
																	F_listRelease(m, v81)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return
																	} else {
																		v85 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																		v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+128))
																		F_listRelease(m, v86)
																		mBase = m.M
																		v88 = m.ExcPending
																		if v88 != 0 {
																			return
																		} else {
																			v90 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																			v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+136))
																			F_listRelease(m, v91)
																			mBase = m.M
																			v93 = m.ExcPending
																			if v93 != 0 {
																				return
																			} else {
																				v95 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+144))
																				F_listRelease(m, v96)
																				mBase = m.M
																				v98 = m.ExcPending
																				if v98 != 0 {
																					return
																				} else {
																					v100 = *(*int32)(unsafe.Add(mBase, _consts[745]))
																					F_valkey_free(m, v100)
																					mBase = m.M
																					v102 = m.ExcPending
																					if v102 != 0 {
																						return
																					} else {
																						v103 = int32(0)
																						*(*int32)(unsafe.Add(mBase, _consts[745])) = v103
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
				}
			}
		}
	}
}
func F_loadServerConfigFromString(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
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
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v980 int32
	_ = v980
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(1152)
	m.G0 = v12
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[244])) = uint8(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v2
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v23 & int32(7) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v40 = v2
		goto L1
	}
L1:
	;
	v45 = F_sdssplitlen(m, l0, v40, int32(_a26), int32(1), v12+int32(120))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v40 = v39
	goto L1
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v40 = v36
	goto L1
L4:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v40 = v33
	goto L1
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v40 = v30
	goto L1
L6:
	;
	v40 = int32(base.Ui32(v23) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	if v47 < int32(1) {
		v1005 = v2
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(_a469)
	v1079 = *(*int32)(unsafe.Add(mBase, _consts[245]))
	v1083 = F_fiprintf(m, v1079, int32(_a470), v12+int32(48))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L7
	} else {
		goto L352
	}
L10:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013))))
	if v1014 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L11:
	;
	v52 = v2
	goto L12
L12:
	;
	v61 = v45 + v52<<(uint(int32(2))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(_a27)
	v69 = v62 + int32(-1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v70 & int32(7) {
	case 0:
		goto L20
	case 1:
		goto L19
	case 2:
		goto L18
	case 3:
		goto L17
	case 4:
		goto L16
	default:
		v87 = int32(0)
		goto L15
	}
L13:
	;
	v1005 = v156
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v62
	v156 = v52 + int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v157 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L15:
	;
	v90 = v62 + v87 + int32(-1)
	if base.Ui32(v90) < base.Ui32(v62) {
		v108 = v62
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(-17))))
	v87 = v86
	goto L15
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(-9))))
	v87 = v83
	goto L15
L18:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+int32(-5)))))
	v87 = v80
	goto L15
L19:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(-3)))))
	v87 = v77
	goto L15
L20:
	;
	v87 = int32(base.Ui32(v70) >> (uint(int32(3)) % 32))
	goto L15
L21:
	;
	if base.Ui32(v90) <= base.Ui32(v108) {
		v124 = v90
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v96 = v62
	goto L23
L23:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	v98 = F_strchr(m, v63, v97)
	mBase = m.M
	if v98 == int32(0) {
		v108 = v96
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v108 = v102
	goto L21
L25:
	;
	v102 = v96 + int32(1)
	if base.Ui32(v102) <= base.Ui32(v90) {
		v96 = v102
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v129 = v124 - v108 + int32(1)
	if v62 == v108 {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	v112 = v90
	goto L29
L29:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	v116 = F_strchr(m, v63, v115)
	mBase = m.M
	if v116 == int32(0) {
		v124 = v112
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v124 = v108
	goto L27
L31:
	;
	v120 = v112 + int32(-1)
	if base.Ui32(v108) < base.Ui32(v120) {
		v112 = v120
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v129))) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	switch v135 & int32(7) {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	default:
		goto L35
	}
L34:
	;
	v131 = F_memmove(m, v62, v108, v129)
	mBase = m.M
	goto L33
L35:
	;
	goto L14
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62+int32(-17)))) = base.I64_extend_i32_u(v129)
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62+int32(-9)))) = v129
	goto L14
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v62+int32(-5)))) = uint16(v129)
	goto L14
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(-3)))) = uint8(v129)
	goto L14
L40:
	;
	v139 = v129 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v139)
	goto L14
L41:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	if v156 < v1001 {
		v52 = v156
		goto L12
	} else {
		goto L333
	}
L42:
	;
	if v157 == int32(35) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v164 = F_sdssplitargs(m, v62, v12+int32(116))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L45
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	if v168 != 0 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v164 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a471)
	v1069 = v156
	v1071 = v52
	goto L9
L47:
	;
	F_sdsfreesplitres(m, v164, v987)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L7
	} else {
		goto L332
	}
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+int32(-1)))))
	switch v176 & int32(7) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	default:
		goto L51
	}
L49:
	;
	v987 = int32(0)
	goto L47
L50:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v215 = F_dictFind(m, v213, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L70
	}
L51:
	;
	goto L50
L52:
	;
	if v193 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-17))))
	v193 = v192
	goto L52
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v170+int32(-9))))
	v193 = v189
	goto L52
L55:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170+int32(-5)))))
	v193 = v186
	goto L52
L56:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+int32(-3)))))
	v193 = v183
	goto L52
L57:
	;
	v193 = int32(base.Ui32(v176) >> (uint(int32(3)) % 32))
	goto L52
L58:
	;
	v198 = int32(0)
	goto L59
L59:
	;
	v201 = v170 + v198
	v202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v201))))
	v203 = F_tolower(m, v202)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v203)
	v206 = v198 + int32(1)
	if v206 != v193 {
		v198 = v206
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L51
L61:
	;
	goto L60
L62:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v987 = v980
	goto L47
L63:
	;
	v966 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v967 = F_dictReplace(m, v966, v857, v960)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L7
	} else {
		goto L329
	}
L64:
	;
	v933 = int32(2)
	v937 = v861
	goto L325
L65:
	;
	F__serverAssert(m, int32(_a472), int32(_a473), int32(556))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L7
	} else {
		goto L324
	}
L66:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	F_sdsfreesplitres(m, v164, v924)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L7
	} else {
		goto L323
	}
L67:
	;
	v620 = int32(_a474)
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v623 != 0 {
		goto L219
	} else {
		goto L220
	}
L68:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v559 = v557 & int32(8)
	if v559 != 0 {
		goto L194
	} else {
		goto L195
	}
L69:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v223 = int32(_a475)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v226 != 0 {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	if v215 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	goto L72
L72:
	;
	if v219 != 0 {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v264 = int32(_a476)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v267 != 0 {
		goto L92
	} else {
		goto L93
	}
L75:
	;
	if v258-v260 != 0 {
		goto L74
	} else {
		goto L87
	}
L76:
	;
	v258 = F_tolower(m, v254)
	mBase = m.M
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	v260 = F_tolower(m, v259)
	mBase = m.M
	goto L75
L77:
	;
	v228 = v222
	v229 = v223
	v230 = v226
	goto L80
L78:
	;
	v254 = int32(0)
	v255 = v223
	goto L76
L79:
	;
	v254 = v251 & int32(255)
	v255 = v250
	goto L76
L80:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v232 == int32(0) {
		v250 = v229
		v251 = v230
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v250 = v244
	v251 = int32(0)
	goto L79
L82:
	;
	v236 = v230 & int32(255)
	if v236 == v232 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v243 = int32(1)
	v244 = v229 + v243
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v245 != 0 {
		v228 = v228 + v243
		v229 = v244
		v230 = v245
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v238 = F_tolower(m, v236)
	mBase = m.M
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v240 = F_tolower(m, v239)
	mBase = m.M
	if v238 == v240 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v250 = v229
	v251 = v242
	goto L79
L86:
	;
	goto L81
L87:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L88
	}
L88:
	;
	goto L74
L89:
	;
	v305 = int32(_a477)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v308 != 0 {
		goto L107
	} else {
		goto L108
	}
L90:
	;
	if v299-v301 != 0 {
		goto L89
	} else {
		goto L102
	}
L91:
	;
	v299 = F_tolower(m, v295)
	mBase = m.M
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v301 = F_tolower(m, v300)
	mBase = m.M
	goto L90
L92:
	;
	v269 = v222
	v270 = v264
	v271 = v267
	goto L95
L93:
	;
	v295 = int32(0)
	v296 = v264
	goto L91
L94:
	;
	v295 = v292 & int32(255)
	v296 = v291
	goto L91
L95:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v273 == int32(0) {
		v291 = v270
		v292 = v271
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v291 = v285
	v292 = int32(0)
	goto L94
L97:
	;
	v277 = v271 & int32(255)
	if v277 == v273 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v284 = int32(1)
	v285 = v270 + v284
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v286 != 0 {
		v269 = v269 + v284
		v270 = v285
		v271 = v286
		goto L95
	} else {
		goto L101
	}
L99:
	;
	v279 = F_tolower(m, v277)
	mBase = m.M
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v281 = F_tolower(m, v280)
	mBase = m.M
	if v279 == v281 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v291 = v270
	v292 = v283
	goto L94
L101:
	;
	goto L96
L102:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L103
	}
L103:
	;
	goto L89
L104:
	;
	v346 = int32(_a478)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v349 != 0 {
		goto L122
	} else {
		goto L123
	}
L105:
	;
	if v340-v342 != 0 {
		goto L104
	} else {
		goto L117
	}
L106:
	;
	v340 = F_tolower(m, v336)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	v342 = F_tolower(m, v341)
	mBase = m.M
	goto L105
L107:
	;
	v310 = v222
	v311 = v305
	v312 = v308
	goto L110
L108:
	;
	v336 = int32(0)
	v337 = v305
	goto L106
L109:
	;
	v336 = v333 & int32(255)
	v337 = v332
	goto L106
L110:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v314 == int32(0) {
		v332 = v311
		v333 = v312
		goto L109
	} else {
		goto L112
	}
L111:
	;
	v332 = v326
	v333 = int32(0)
	goto L109
L112:
	;
	v318 = v312 & int32(255)
	if v318 == v314 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v325 = int32(1)
	v326 = v311 + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	if v327 != 0 {
		v310 = v310 + v325
		v311 = v326
		v312 = v327
		goto L110
	} else {
		goto L116
	}
L114:
	;
	v320 = F_tolower(m, v318)
	mBase = m.M
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	v322 = F_tolower(m, v321)
	mBase = m.M
	if v320 == v322 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	v332 = v311
	v333 = v324
	goto L109
L116:
	;
	goto L111
L117:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L118
	}
L118:
	;
	goto L104
L119:
	;
	v387 = int32(_a479)
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v390 != 0 {
		goto L137
	} else {
		goto L138
	}
L120:
	;
	if v381-v383 != 0 {
		goto L119
	} else {
		goto L132
	}
L121:
	;
	v381 = F_tolower(m, v377)
	mBase = m.M
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v383 = F_tolower(m, v382)
	mBase = m.M
	goto L120
L122:
	;
	v351 = v222
	v352 = v346
	v353 = v349
	goto L125
L123:
	;
	v377 = int32(0)
	v378 = v346
	goto L121
L124:
	;
	v377 = v374 & int32(255)
	v378 = v373
	goto L121
L125:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v355 == int32(0) {
		v373 = v352
		v374 = v353
		goto L124
	} else {
		goto L127
	}
L126:
	;
	v373 = v367
	v374 = int32(0)
	goto L124
L127:
	;
	v359 = v353 & int32(255)
	if v359 == v355 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v366 = int32(1)
	v367 = v352 + v366
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+1)))
	if v368 != 0 {
		v351 = v351 + v366
		v352 = v367
		v353 = v368
		goto L125
	} else {
		goto L131
	}
L129:
	;
	v361 = F_tolower(m, v359)
	mBase = m.M
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v363 = F_tolower(m, v362)
	mBase = m.M
	if v361 == v363 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	v373 = v352
	v374 = v365
	goto L124
L131:
	;
	goto L126
L132:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L133
	}
L133:
	;
	goto L119
L134:
	;
	v428 = int32(_a480)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v431 != 0 {
		goto L152
	} else {
		goto L153
	}
L135:
	;
	if v422-v424 != 0 {
		goto L134
	} else {
		goto L147
	}
L136:
	;
	v422 = F_tolower(m, v418)
	mBase = m.M
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	v424 = F_tolower(m, v423)
	mBase = m.M
	goto L135
L137:
	;
	v392 = v222
	v393 = v387
	v394 = v390
	goto L140
L138:
	;
	v418 = int32(0)
	v419 = v387
	goto L136
L139:
	;
	v418 = v415 & int32(255)
	v419 = v414
	goto L136
L140:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v396 == int32(0) {
		v414 = v393
		v415 = v394
		goto L139
	} else {
		goto L142
	}
L141:
	;
	v414 = v408
	v415 = int32(0)
	goto L139
L142:
	;
	v400 = v394 & int32(255)
	if v400 == v396 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v407 = int32(1)
	v408 = v393 + v407
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	if v409 != 0 {
		v392 = v392 + v407
		v393 = v408
		v394 = v409
		goto L140
	} else {
		goto L146
	}
L144:
	;
	v402 = F_tolower(m, v400)
	mBase = m.M
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	v404 = F_tolower(m, v403)
	mBase = m.M
	if v402 == v404 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	v414 = v393
	v415 = v406
	goto L139
L146:
	;
	goto L141
L147:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L148
	}
L148:
	;
	goto L134
L149:
	;
	v469 = int32(_a481)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v472 != 0 {
		goto L167
	} else {
		goto L168
	}
L150:
	;
	if v463-v465 != 0 {
		goto L149
	} else {
		goto L162
	}
L151:
	;
	v463 = F_tolower(m, v459)
	mBase = m.M
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v465 = F_tolower(m, v464)
	mBase = m.M
	goto L150
L152:
	;
	v433 = v222
	v434 = v428
	v435 = v431
	goto L155
L153:
	;
	v459 = int32(0)
	v460 = v428
	goto L151
L154:
	;
	v459 = v456 & int32(255)
	v460 = v455
	goto L151
L155:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v437 == int32(0) {
		v455 = v434
		v456 = v435
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v455 = v449
	v456 = int32(0)
	goto L154
L157:
	;
	v441 = v435 & int32(255)
	if v441 == v437 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v448 = int32(1)
	v449 = v434 + v448
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	if v450 != 0 {
		v433 = v433 + v448
		v434 = v449
		v435 = v450
		goto L155
	} else {
		goto L161
	}
L159:
	;
	v443 = F_tolower(m, v441)
	mBase = m.M
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v445 = F_tolower(m, v444)
	mBase = m.M
	if v443 == v445 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v455 = v434
	v456 = v447
	goto L154
L161:
	;
	goto L156
L162:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L163
	}
L163:
	;
	goto L149
L164:
	;
	v510 = int32(_a482)
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v513 != 0 {
		goto L181
	} else {
		goto L182
	}
L165:
	;
	if v504-v506 != 0 {
		goto L164
	} else {
		goto L177
	}
L166:
	;
	v504 = F_tolower(m, v500)
	mBase = m.M
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	v506 = F_tolower(m, v505)
	mBase = m.M
	goto L165
L167:
	;
	v474 = v222
	v475 = v469
	v476 = v472
	goto L170
L168:
	;
	v500 = int32(0)
	v501 = v469
	goto L166
L169:
	;
	v500 = v497 & int32(255)
	v501 = v496
	goto L166
L170:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	if v478 == int32(0) {
		v496 = v475
		v497 = v476
		goto L169
	} else {
		goto L172
	}
L171:
	;
	v496 = v490
	v497 = int32(0)
	goto L169
L172:
	;
	v482 = v476 & int32(255)
	if v482 == v478 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v489 = int32(1)
	v490 = v475 + v489
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
	if v491 != 0 {
		v474 = v474 + v489
		v475 = v490
		v476 = v491
		goto L170
	} else {
		goto L176
	}
L174:
	;
	v484 = F_tolower(m, v482)
	mBase = m.M
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v486 = F_tolower(m, v485)
	mBase = m.M
	if v484 == v486 {
		goto L173
	} else {
		goto L175
	}
L175:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v496 = v475
	v497 = v488
	goto L169
L176:
	;
	goto L171
L177:
	;
	if v221 == int32(2) {
		v987 = v221
		goto L47
	} else {
		goto L178
	}
L178:
	;
	goto L164
L179:
	;
	if v545-v547 != 0 {
		goto L67
	} else {
		goto L191
	}
L180:
	;
	v545 = F_tolower(m, v541)
	mBase = m.M
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	v547 = F_tolower(m, v546)
	mBase = m.M
	goto L179
L181:
	;
	v515 = v222
	v516 = v510
	v517 = v513
	goto L184
L182:
	;
	v541 = int32(0)
	v542 = v510
	goto L180
L183:
	;
	v541 = v538 & int32(255)
	v542 = v537
	goto L180
L184:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	if v519 == int32(0) {
		v537 = v516
		v538 = v517
		goto L183
	} else {
		goto L186
	}
L185:
	;
	v537 = v531
	v538 = int32(0)
	goto L183
L186:
	;
	v523 = v517 & int32(255)
	if v523 == v519 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v530 = int32(1)
	v531 = v516 + v530
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v532 != 0 {
		v515 = v515 + v530
		v516 = v531
		v517 = v532
		goto L184
	} else {
		goto L190
	}
L188:
	;
	v525 = F_tolower(m, v523)
	mBase = m.M
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	v527 = F_tolower(m, v526)
	mBase = m.M
	if v525 == v527 {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	v537 = v516
	v538 = v529
	goto L183
L190:
	;
	goto L185
L191:
	;
	if v221 != int32(2) {
		goto L67
	} else {
		goto L192
	}
L192:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v552 = int32(0)
	F_loadServerConfig(m, v551, v552, v552)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L7
	} else {
		goto L193
	}
L193:
	;
	goto L62
L194:
	;
	if v559 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	if v556 == int32(2) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a483)
	goto L66
L197:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v618 = m.T0[v617].(func(*base.Module, int32, int32, int32, int32) int32)(m, v219, v164+int32(4), v556+int32(-1), v12+int32(124))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L7
	} else {
		goto L214
	}
L198:
	;
	if v556 != int32(2) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+int32(-1)))))
	switch v571 & int32(7) {
	case 0:
		goto L205
	case 1:
		goto L204
	case 2:
		goto L203
	case 3:
		goto L202
	case 4:
		goto L201
	default:
		goto L197
	}
L200:
	;
	if v588 == int32(0) {
		goto L197
	} else {
		goto L206
	}
L201:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v568+int32(-17))))
	v588 = v587
	goto L200
L202:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v568+int32(-9))))
	v588 = v584
	goto L200
L203:
	;
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v568+int32(-5)))))
	v588 = v581
	goto L200
L204:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+int32(-3)))))
	v588 = v578
	goto L200
L205:
	;
	v588 = int32(base.Ui32(v571) >> (uint(int32(3)) % 32))
	goto L200
L206:
	;
	v593 = F_sdssplitargs(m, v568, v12+int32(128))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L208
	}
L207:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_sdsfreesplitres(m, v593, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L7
	} else {
		goto L213
	}
L208:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v599 = m.T0[v598].(func(*base.Module, int32, int32, int32, int32) int32)(m, v219, v593, v595, v12+int32(124))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	if v599 != 0 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	if v593 == int32(0) {
		goto L66
	} else {
		goto L211
	}
L211:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
	F_sdsfreesplitres(m, v593, v603)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
	} else {
		goto L212
	}
L212:
	;
	goto L66
L213:
	;
	goto L62
L214:
	;
	if v618 != 0 {
		goto L62
	} else {
		goto L215
	}
L215:
	;
	goto L66
L216:
	;
	v712 = int32(_a23)
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v715 != 0 {
		goto L252
	} else {
		goto L253
	}
L217:
	;
	if v655-v657 != 0 {
		goto L216
	} else {
		goto L229
	}
L218:
	;
	v655 = F_tolower(m, v651)
	mBase = m.M
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	v657 = F_tolower(m, v656)
	mBase = m.M
	goto L217
L219:
	;
	v625 = v222
	v626 = v620
	v627 = v623
	goto L222
L220:
	;
	v651 = int32(0)
	v652 = v620
	goto L218
L221:
	;
	v651 = v648 & int32(255)
	v652 = v647
	goto L218
L222:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if v629 == int32(0) {
		v647 = v626
		v648 = v627
		goto L221
	} else {
		goto L224
	}
L223:
	;
	v647 = v641
	v648 = int32(0)
	goto L221
L224:
	;
	v633 = v627 & int32(255)
	if v633 == v629 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v640 = int32(1)
	v641 = v626 + v640
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+1)))
	if v642 != 0 {
		v625 = v625 + v640
		v626 = v641
		v627 = v642
		goto L222
	} else {
		goto L228
	}
L226:
	;
	v635 = F_tolower(m, v633)
	mBase = m.M
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	v637 = F_tolower(m, v636)
	mBase = m.M
	if v635 == v637 {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	v647 = v626
	v648 = v639
	goto L221
L228:
	;
	goto L223
L229:
	;
	if v221 != int32(3) {
		goto L216
	} else {
		goto L230
	}
L230:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v662 = F_lookupCommandBySds(m, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L7
	} else {
		goto L232
	}
L231:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v669 = F_hashtableDelete(m, v667, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L7
	} else {
		goto L234
	}
L232:
	;
	if v662 != 0 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a484)
	goto L66
L234:
	;
	if v669 == int32(0) {
		goto L65
	} else {
		goto L235
	}
L235:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+int32(-1)))))
	switch v676 & int32(7) {
	case 0:
		goto L241
	case 1:
		goto L240
	case 2:
		goto L239
	case 3:
		goto L238
	case 4:
		goto L237
	default:
		goto L62
	}
L236:
	;
	if v693 == int32(0) {
		goto L62
	} else {
		goto L242
	}
L237:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v673+int32(-17))))
	v693 = v692
	goto L236
L238:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v673+int32(-9))))
	v693 = v689
	goto L236
L239:
	;
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v673+int32(-5)))))
	v693 = v686
	goto L236
L240:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673+int32(-3)))))
	v693 = v683
	goto L236
L241:
	;
	v693 = int32(base.Ui32(v676) >> (uint(int32(3)) % 32))
	goto L236
L242:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v662)+144))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v662)+140))
	if v696 == v697 {
		v702 = v673
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v703 = F_sdsdup(m, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L7
	} else {
		goto L246
	}
L244:
	;
	F_sdsfree(m, v696)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L7
	} else {
		goto L245
	}
L245:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v702 = v701
	goto L243
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v662)+144)) = v703
	v707 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v708 = F_hashtableAdd(m, v707, v662)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L7
	} else {
		goto L247
	}
L247:
	;
	if v708 != 0 {
		goto L62
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a485)
	goto L66
L249:
	;
	v795 = int32(_a486)
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v798 != 0 {
		goto L281
	} else {
		goto L282
	}
L250:
	;
	if v747-v749 != 0 {
		goto L249
	} else {
		goto L262
	}
L251:
	;
	v747 = F_tolower(m, v743)
	mBase = m.M
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	v749 = F_tolower(m, v748)
	mBase = m.M
	goto L250
L252:
	;
	v717 = v222
	v718 = v712
	v719 = v715
	goto L255
L253:
	;
	v743 = int32(0)
	v744 = v712
	goto L251
L254:
	;
	v743 = v740 & int32(255)
	v744 = v739
	goto L251
L255:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	if v721 == int32(0) {
		v739 = v718
		v740 = v719
		goto L254
	} else {
		goto L257
	}
L256:
	;
	v739 = v733
	v740 = int32(0)
	goto L254
L257:
	;
	v725 = v719 & int32(255)
	if v725 == v721 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v732 = int32(1)
	v733 = v718 + v732
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717)+1)))
	if v734 != 0 {
		v717 = v717 + v732
		v718 = v733
		v719 = v734
		goto L255
	} else {
		goto L261
	}
L259:
	;
	v727 = F_tolower(m, v725)
	mBase = m.M
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	v729 = F_tolower(m, v728)
	mBase = m.M
	if v727 == v729 {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	v739 = v718
	v740 = v731
	goto L254
L261:
	;
	goto L256
L262:
	;
	if v221 < int32(2) {
		goto L249
	} else {
		goto L263
	}
L263:
	;
	v755 = F_ACLAppendUserForLoading(m, v164, v221, v12+int32(112))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L7
	} else {
		goto L264
	}
L264:
	;
	if v755 != int32(-1) {
		goto L62
	} else {
		goto L265
	}
L265:
	;
	v761 = F___errno_location(m)
	mBase = m.M
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	switch v762 + int32(-7) {
	case 0:
		goto L271
	default:
		goto L268
	case 2:
		goto L272
	case 5:
		goto L270
	case 13:
		goto L275
	case 21:
		goto L276
	case 24:
		goto L274
	case 36:
		goto L273
	case 37:
		v774 = int32(_a34)
		goto L267
	case 61:
		goto L269
	}
L266:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v164+v777<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v781
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v776
	v790 = F_snprintf(m, v12+int32(128), int32(1024), int32(_a487), v12+int32(80))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L7
	} else {
		goto L277
	}
L267:
	;
	v776 = v774
	goto L266
L268:
	;
	v774 = int32(_a35)
	goto L267
L269:
	;
	v776 = int32(_a36)
	goto L266
L270:
	;
	v776 = int32(_a37)
	goto L266
L271:
	;
	v776 = int32(_a38)
	goto L266
L272:
	;
	v776 = int32(_a39)
	goto L266
L273:
	;
	v776 = int32(_a40)
	goto L266
L274:
	;
	v776 = int32(_a41)
	goto L266
L275:
	;
	v776 = int32(_a42)
	goto L266
L276:
	;
	v776 = int32(_a43)
	goto L266
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v12 + int32(128)
	goto L66
L278:
	;
	v843 = int32(46)
	v844 = F___strchrnul(m, v222, v843)
	mBase = m.M
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844))))
	if v846 == v843 {
		goto L296
	} else {
		goto L297
	}
L279:
	;
	if v830-v832 != 0 {
		goto L278
	} else {
		goto L291
	}
L280:
	;
	v830 = F_tolower(m, v826)
	mBase = m.M
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827))))
	v832 = F_tolower(m, v831)
	mBase = m.M
	goto L279
L281:
	;
	v800 = v222
	v801 = v795
	v802 = v798
	goto L284
L282:
	;
	v826 = int32(0)
	v827 = v795
	goto L280
L283:
	;
	v826 = v823 & int32(255)
	v827 = v822
	goto L280
L284:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	if v804 == int32(0) {
		v822 = v801
		v823 = v802
		goto L283
	} else {
		goto L286
	}
L285:
	;
	v822 = v816
	v823 = int32(0)
	goto L283
L286:
	;
	v808 = v802 & int32(255)
	if v808 == v804 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v815 = int32(1)
	v816 = v801 + v815
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+1)))
	if v817 != 0 {
		v800 = v800 + v815
		v801 = v816
		v802 = v817
		goto L284
	} else {
		goto L290
	}
L288:
	;
	v810 = F_tolower(m, v808)
	mBase = m.M
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801))))
	v812 = F_tolower(m, v811)
	mBase = m.M
	if v810 == v812 {
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800))))
	v822 = v801
	v823 = v814
	goto L283
L290:
	;
	goto L285
L291:
	;
	if v221 < int32(2) {
		goto L278
	} else {
		goto L292
	}
L292:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	F_moduleEnqueueLoadModule(m, v836, v164+int32(8), v221+int32(-2))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L7
	} else {
		goto L293
	}
L293:
	;
	goto L62
L294:
	;
	v866 = int32(_a488)
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v869 != 0 {
		goto L308
	} else {
		goto L309
	}
L295:
	;
	if v850 == int32(0) {
		goto L294
	} else {
		goto L299
	}
L296:
	;
	v850 = v844
	goto L298
L297:
	;
	v850 = int32(0)
	goto L298
L298:
	;
	goto L295
L299:
	;
	if int32(1) < v221 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v857 = F_sdsdup(m, v222)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L7
	} else {
		goto L302
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a489)
	goto L66
L302:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v861 = F_sdsdup(m, v860)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	if int32(2) < v863 {
		goto L64
	} else {
		goto L304
	}
L304:
	;
	v960 = v861
	goto L63
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a490)
	goto L66
L306:
	;
	if v901-v903 != 0 {
		goto L305
	} else {
		goto L318
	}
L307:
	;
	v901 = F_tolower(m, v897)
	mBase = m.M
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	v903 = F_tolower(m, v902)
	mBase = m.M
	goto L306
L308:
	;
	v871 = v222
	v872 = v866
	v873 = v869
	goto L311
L309:
	;
	v897 = int32(0)
	v898 = v866
	goto L307
L310:
	;
	v897 = v894 & int32(255)
	v898 = v893
	goto L307
L311:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872))))
	if v875 == int32(0) {
		v893 = v872
		v894 = v873
		goto L310
	} else {
		goto L313
	}
L312:
	;
	v893 = v887
	v894 = int32(0)
	goto L310
L313:
	;
	v879 = v873 & int32(255)
	if v879 == v875 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v886 = int32(1)
	v887 = v872 + v886
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871)+1)))
	if v888 != 0 {
		v871 = v871 + v886
		v872 = v887
		v873 = v888
		goto L311
	} else {
		goto L317
	}
L315:
	;
	v881 = F_tolower(m, v879)
	mBase = m.M
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872))))
	v883 = F_tolower(m, v882)
	mBase = m.M
	if v881 == v883 {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	v893 = v872
	v894 = v885
	goto L310
L317:
	;
	goto L312
L318:
	;
	if v221 == int32(1) {
		goto L62
	} else {
		goto L319
	}
L319:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _consts[250]))
	if v908 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	F_queueSentinelConfig(m, v164+int32(4), v221+int32(-1), v156, v915)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L7
	} else {
		goto L322
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = int32(_a491)
	goto L66
L322:
	;
	goto L62
L323:
	;
	v1069 = v156
	v1071 = v52
	goto L9
L324:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v164+v933<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v945
	v950 = F_sdscatfmt(m, v937, int32(_a492), v12+int32(96))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L7
	} else {
		goto L327
	}
L326:
	;
	v960 = v950
	goto L63
L327:
	;
	v953 = v933 + int32(1)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	if v953 < v954 {
		v933 = v953
		v937 = v950
		goto L325
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	if v967 != 0 {
		goto L62
	} else {
		goto L330
	}
L330:
	;
	F_sdsfree(m, v857)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L7
	} else {
		goto L331
	}
L331:
	;
	goto L62
L332:
	;
	goto L41
L333:
	;
	goto L13
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v1065
	v1069 = v1005
	v1071 = v1005
	goto L9
L335:
	;
	v1054 = F_sdsempty(m)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L7
	} else {
		goto L348
	}
L336:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v1025 == int32(0) {
		goto L340
	} else {
		goto L341
	}
L337:
	;
	v1018 = F_fopen(m, v1013, int32(_a493))
	mBase = m.M
	if v1018 == int32(0) {
		goto L335
	} else {
		goto L338
	}
L338:
	;
	v1021 = F_fclose(m, v1018)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L7
	} else {
		goto L339
	}
L339:
	;
	goto L336
L340:
	;
	v1033 = int32(1)
	v1035 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	if v1035 < v1033 {
		v1041 = v1033
		goto L344
	} else {
		goto L345
	}
L341:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v1029 == int32(0) {
		goto L340
	} else {
		goto L342
	}
L342:
	;
	v1065 = int32(_a494)
	goto L334
L343:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	F_sdsfreesplitres(m, v45, v1045)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L7
	} else {
		goto L347
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, _consts[251])) = v1041
	goto L343
L345:
	;
	v1038 = int32(500)
	if base.Ui32(v1035) <= base.Ui32(v1038) {
		goto L343
	} else {
		goto L346
	}
L346:
	;
	v1041 = v1038
	goto L344
L347:
	;
	v1048 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[244])) = uint8(v1048)
	m.G0 = v12 + int32(1152)
	return
L348:
	;
	goto L349
L349:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v1058 = F___strerror_l(m, v1057, v1057)
	mBase = m.M
	goto L350
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v1058
	v1063 = F_sdscatprintf(m, v1054, int32(_a495), v12+int32(64))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L7
	} else {
		goto L351
	}
L351:
	;
	v1065 = v1063
	goto L334
L352:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	if v1085 <= v1071 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v1103
	v1106 = F_fiprintf(m, v1079, int32(_a496), v12)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L7
	} else {
		goto L357
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v1069
	v1091 = F_fiprintf(m, v1079, int32(_a497), v12+int32(32))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L7
	} else {
		goto L355
	}
L355:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v45+v1071<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v1096
	v1101 = F_fiprintf(m, v1079, int32(_a498), v12+int32(16))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L7
	} else {
		goto L356
	}
L356:
	;
	goto L353
L357:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_resetServerSaveParams(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[243]))
	F_valkey_free(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[243])) = int64(0)
		return
	}
}
func F_resetServerStats(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int64
	_ = v2
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	v1 = int32(0)
	v2 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[602])) = v2
	*(*int64)(unsafe.Add(mBase, _consts[754])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[755])) = v1
	v14 = F__emscripten_memset_bulkmem(m, int32(_a2187), base.I32_extend8_s(v1), int32(96))
	mBase = m.M
	v15 = int32(0)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[349])) = v16
	*(*int64)(unsafe.Add(mBase, _consts[735])) = v16
	*(*uint8)(unsafe.Add(mBase, _consts[720])) = uint8(v15)
	v28 = F__emscripten_memset_bulkmem(m, int32(_a2204), base.I32_extend8_s(v15), int32(176))
	mBase = m.M
	v33 = F__emscripten_memset_bulkmem(m, int32(_a2205), base.I32_extend8_s(int32(0)), int32(80))
	mBase = m.M
	v34 = int32(0)
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[756])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[757])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[596])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[601])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[758])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[759])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[760])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[421])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[541])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[761])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[762])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[763])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[364])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[764])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[765])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[766])) = v35
	*(*int64)(unsafe.Add(mBase, _consts[482])) = v35
	v89 = F__emscripten_memset_bulkmem(m, int32(_a2206), base.I32_extend8_s(v34), int32(148))
	mBase = m.M
	v94 = F__emscripten_memset_bulkmem(m, int32(_a2207), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v99 = F__emscripten_memset_bulkmem(m, int32(_a2208), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v104 = F__emscripten_memset_bulkmem(m, int32(_a2209), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v109 = F__emscripten_memset_bulkmem(m, int32(_a2210), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v114 = F__emscripten_memset_bulkmem(m, int32(_a2211), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v119 = F__emscripten_memset_bulkmem(m, int32(_a2212), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v124 = F__emscripten_memset_bulkmem(m, int32(_a2213), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v129 = F__emscripten_memset_bulkmem(m, int32(_a673), base.I32_extend8_s(int32(0)), int32(148))
	mBase = m.M
	v130 = int32(0)
	v131 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[767])) = v131
	*(*int64)(unsafe.Add(mBase, _consts[768])) = v131
	*(*int32)(unsafe.Add(mBase, _consts[374])) = v130
	return
}
func F_serverAsciiArt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_valkey_malloc(m, int32(16384))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[250]))
		if v19 != 0 {
			v20 = int32(_a488)
		} else {
			v20 = int32(_a2285)
		}
		v22 = *(*int32)(unsafe.Add(mBase, _consts[110]))
		if v22 != 0 {
			v23 = int32(_a2286)
		} else {
			v23 = v20
		}
		v25 = *(*int32)(unsafe.Add(mBase, _consts[792]))
		if v25 != 0 {
			v58 = *(*int32)(unsafe.Add(mBase, _consts[793]))
			if v58 != 0 {
				v81 = F_strtox_2(m, int32(_a1908), int32(0), int32(10), int64(2147483648))
				mBase = m.M
				v83 = int32(0)
				v84 = *(*int32)(unsafe.Add(mBase, _consts[139]))
				v86 = *(*int32)(unsafe.Add(mBase, _consts[137]))
				v89 = F___syscall_getpid(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v89
				if v86 != 0 {
					v93 = v86
				} else {
					v93 = v84
				}
				*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v93
				*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(_a2287)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = base.B2i32(int32(0) < base.I32_wrap_i64(v81))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a2283)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a469)
				v108 = F_snprintf(m, v13, int32(16384), int32(_a2288), v10)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return
				} else {
					F_serverLogRaw(m, int32(1026), v13)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						F_valkey_free(m, v13)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(2) < v60 {
					F_valkey_free(m, v13)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v23
					v64 = int32(0)
					v65 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					v67 = *(*int32)(unsafe.Add(mBase, _consts[139]))
					if v65 != 0 {
						v68 = v65
					} else {
						v68 = v67
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v68
					F__serverLog(m, int32(2), int32(_a2289), v10+int32(32))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						F_valkey_free(m, v13)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, _consts[246]))
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
			if v28 != 0 {
				v58 = *(*int32)(unsafe.Add(mBase, _consts[793]))
				if v58 != 0 {
					v81 = F_strtox_2(m, int32(_a1908), int32(0), int32(10), int64(2147483648))
					mBase = m.M
					v83 = int32(0)
					v84 = *(*int32)(unsafe.Add(mBase, _consts[139]))
					v86 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					v89 = F___syscall_getpid(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v89
					if v86 != 0 {
						v93 = v86
					} else {
						v93 = v84
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v93
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(_a2287)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = base.B2i32(int32(0) < base.I32_wrap_i64(v81))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a2283)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a469)
					v108 = F_snprintf(m, v13, int32(16384), int32(_a2288), v10)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						F_serverLogRaw(m, int32(1026), v13)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							F_valkey_free(m, v13)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(2) < v60 {
						F_valkey_free(m, v13)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v23
						v64 = int32(0)
						v65 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						v67 = *(*int32)(unsafe.Add(mBase, _consts[139]))
						if v65 != 0 {
							v68 = v65
						} else {
							v68 = v67
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v68
						F__serverLog(m, int32(2), int32(_a2289), v10+int32(32))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							F_valkey_free(m, v13)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[794]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
				if int32(-1) < v33 {
					v37 = F___lockfile(m, v30)
					mBase = m.M
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
					if v37 == int32(0) {
						v42 = v38
					} else {
						F___unlockfile(m, v30)
						mBase = m.M
						v42 = v38
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
					v42 = v36
				}
				if int32(-1) < v42 {
					v50 = v42
				} else {
					v46 = F___errno_location(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(8)
					v50 = int32(-1)
				}
				v51 = F_isatty(m, v50)
				mBase = m.M
				v52 = int32(0)
				v53 = *(*int32)(unsafe.Add(mBase, _consts[793]))
				if v51|v53 == v52 {
					v60 = *(*int32)(unsafe.Add(mBase, _consts[28]))
					if int32(2) < v60 {
						F_valkey_free(m, v13)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v23
						v64 = int32(0)
						v65 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						v67 = *(*int32)(unsafe.Add(mBase, _consts[139]))
						if v65 != 0 {
							v68 = v65
						} else {
							v68 = v67
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v68
						F__serverLog(m, int32(2), int32(_a2289), v10+int32(32))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							F_valkey_free(m, v13)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				} else {
					v81 = F_strtox_2(m, int32(_a1908), int32(0), int32(10), int64(2147483648))
					mBase = m.M
					v83 = int32(0)
					v84 = *(*int32)(unsafe.Add(mBase, _consts[139]))
					v86 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					v89 = F___syscall_getpid(m)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(24)))) = v89
					if v86 != 0 {
						v93 = v86
					} else {
						v93 = v84
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v93
					*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(_a2287)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = base.B2i32(int32(0) < base.I32_wrap_i64(v81))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a2283)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a469)
					v108 = F_snprintf(m, v13, int32(16384), int32(_a2288), v10)
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return
					} else {
						F_serverLogRaw(m, int32(1026), v13)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							F_valkey_free(m, v13)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_serverBitpos(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int64
	_ = v159
	var v174 int32
	_ = v174
	var v177 int64
	_ = v177
	var v208 int64
	_ = v208
	var v218 int32
	_ = v218
	var v235 int64
	_ = v235
	v4 = int64(0)
	if l0&int32(3) == int32(0) {
		v69 = l0
		v70 = l1
		v71 = v4
		goto L5
	} else {
		goto L6
	}
L1:
	;
	if l2 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L2:
	;
	v157 = v150 << (uint(int32(8)) % 32)
	v159 = v149
	goto L1
L3:
	;
	v149 = v139
	v150 = v140 << (uint(int32(8)) % 32)
	goto L2
L4:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v117 = v115 << (uint(int32(8)) % 32)
	if v108 == int32(1) {
		v139 = v110
		v140 = v117
		goto L3
	} else {
		goto L38
	}
L5:
	;
	if base.Ui32(v70) < base.Ui32(int32(4)) {
		v98 = v69
		v99 = v70
		v101 = v71
		goto L28
	} else {
		goto L29
	}
L6:
	;
	if l1 == int32(0) {
		v69 = l0
		v70 = l1
		v71 = v4
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v23 = l1 + int32(-1)
	v26 = l0 + int32(1)
	if v26&int32(3) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v18 = int32(0)
	goto L11
L10:
	;
	v18 = int32(255)
	goto L11
L11:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == v19 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v107 = l0
	v108 = l1
	v110 = int64(0)
	goto L4
L13:
	;
	v69 = v26
	v70 = v23
	v71 = int64(8)
	goto L5
L14:
	;
	if v23 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v18 == v33 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v37 = l1 + int32(-2)
	v40 = l0 + int32(2)
	if v40&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v107 = v26
	v108 = v23
	v110 = int64(8)
	goto L4
L18:
	;
	v69 = v40
	v70 = v37
	v71 = int64(16)
	goto L5
L19:
	;
	if v37 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v18 == v47 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v51 = l1 + int32(-3)
	v53 = int32(3)
	v54 = l0 + v53
	if v54&v53 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v107 = v40
	v108 = v37
	v110 = int64(16)
	goto L4
L23:
	;
	v69 = v54
	v70 = v51
	v71 = int64(24)
	goto L5
L24:
	;
	if v51 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v18 == v61 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v69 = l0 + int32(4)
	v70 = l1 + int32(-4)
	v71 = int64(32)
	goto L5
L27:
	;
	v107 = v54
	v108 = v51
	v110 = int64(24)
	goto L4
L28:
	;
	if v99 != 0 {
		v107 = v98
		v108 = v99
		v110 = v101
		goto L4
	} else {
		goto L37
	}
L29:
	;
	if l2 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = int32(0)
	goto L32
L31:
	;
	v79 = int32(-1)
	goto L32
L32:
	;
	v80 = v69
	v81 = v70
	v83 = v71
	goto L33
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v88 != v79 {
		v107 = v80
		v108 = v81
		v110 = v83
		goto L4
	} else {
		goto L35
	}
L34:
	;
	v98 = v93
	v99 = v95
	v101 = v91
	goto L28
L35:
	;
	v91 = v83 + int64(32)
	v93 = v80 + int32(4)
	v95 = v81 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v95) {
		v80 = v93
		v81 = v95
		v83 = v91
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v139 = v101
	v140 = int32(0)
	goto L3
L38:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v123 = (v117 | v120) << (uint(int32(8)) % 32)
	v125 = v108 + int32(-2)
	if v125 == int32(0) {
		v149 = v110
		v150 = v123
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
	v131 = (v123 | v128) << (uint(int32(8)) % 32)
	if v125 == int32(1) {
		v157 = v131
		v159 = v110
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
	v157 = v131 | v134
	v159 = v110
	goto L1
L41:
	;
	return v235
L42:
	;
	v174 = int32(-2147483648)
	v177 = v159
	goto L48
L43:
	;
	if v157 == int32(0) {
		v235 = int64(-1)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v235 = v177 + int64(1)
	goto L41
L46:
	;
	return v177 + int64(2)
L47:
	;
	return v177 + int64(3)
L48:
	;
	if l2 != base.B2i32(v174&v157 != int32(0)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	F__serverPanic_1(m, int32(_a177), int32(348), int32(_a178), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	if l2 == base.B2i32(int32(base.Ui32(v174)>>(uint(int32(1))%32))&v157 != int32(0)) {
		goto L45
	} else {
		goto L52
	}
L51:
	;
	return v177
L52:
	;
	if l2 == base.B2i32(int32(base.Ui32(v174)>>(uint(int32(2))%32))&v157 != int32(0)) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	if l2 == base.B2i32(int32(base.Ui32(v174)>>(uint(int32(3))%32))&v157 != int32(0)) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v208 = v177 + int64(4)
	if base.I32_wrap_i64(v159)+int32(32) != base.I32_wrap_i64(v208) {
		v174 = int32(base.Ui32(v174) >> (uint(int32(4)) % 32))
		v177 = v208
		goto L48
	} else {
		goto L55
	}
L55:
	;
	goto L49
L56:
	;
	return int64(0)
L57:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_serverBuildIdString(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[568])))
	if v7 != 0 {
		m.G0 = v4 + int32(16)
		return int32(_a1909)
	} else {
		v11 = F_crc64(m, int64(0), int32(_a1910), int64(73))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v11
		v16 = F_snprintf(m, int32(_a1909), int32(32), int32(_a1911), v4)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[568])) = uint8(v21)
			m.G0 = v4 + int32(16)
			return int32(_a1909)
		}
	}
}
func F_serverGitDirty(m *base.Module) int32 {
	return int32(_a1908)
}
func F_serverLogRawFromHandler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if l0&int32(255) < v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v17&int32(255) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v17&int32(255) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	if l0&int32(1024) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(420)
	v28 = F_open(m, v16, int32(1089), v8)
	mBase = m.M
	if v28 == int32(-1) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v31 = int32(1)
	goto L6
L9:
	;
	v31 = v28
	goto L6
L10:
	;
	if v17&int32(255) == int32(0) {
		goto L1
	} else {
		goto L101
	}
L11:
	;
	v93 = v8 + int32(16)
	v95 = F___syscall_getpid(m)
	mBase = m.M
	goto L29
L12:
	;
	if l1&int32(3) == int32(0) {
		v57 = l1
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v91 = F_write(m, v31, l1, v90)
	mBase = m.M
	goto L10
L14:
	;
	v90 = v82 - l1
	goto L13
L15:
	;
	v61 = v57
	goto L23
L16:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = l1
	goto L19
L18:
	;
	v90 = l1 - l1
	goto L13
L19:
	;
	v50 = v46 + int32(1)
	if v50&int32(3) == int32(0) {
		v57 = v50
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v55 != 0 {
		v46 = v50
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v82 = v50
	goto L14
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = int32(-2139062144)
	if (int32(16843008)-v67|v67)&v70 == v70 {
		v61 = v61 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v76 = v61
	goto L26
L25:
	;
	goto L24
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		v76 = v76 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v82 = v76
	goto L14
L28:
	;
	goto L27
L29:
	;
	v96 = base.I64_extend_i32_s(v95)
	if v96 <= int64(-1) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v139 = v8 + int32(16)
	if v139&int32(3) == int32(0) {
		v163 = v139
		goto L41
	} else {
		goto L42
	}
L31:
	;
	goto L30
L33:
	;
	v118 = F_ull2string(m, v114, v115, v116)
	mBase = m.M
	if v118 == int32(0) {
		goto L31
	} else {
		goto L37
	}
L34:
	;
	goto L36
L35:
	;
	v114 = v93
	v115 = int32(64)
	v116 = v96
	goto L33
L36:
	;
	v105 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v105)
	v114 = v8 + int32(17)
	v115 = int32(63)
	v116 = int64(0) - v96
	goto L33
L37:
	;
	goto L30
L39:
	;
	v197 = F_write(m, v31, v139, v196)
	mBase = m.M
	if v197 == int32(-1) {
		goto L10
	} else {
		goto L55
	}
L40:
	;
	v196 = v188 - v139
	goto L39
L41:
	;
	v167 = v163
	goto L49
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v149 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = v139
	goto L45
L44:
	;
	v196 = v139 - v139
	goto L39
L45:
	;
	v156 = v152 + int32(1)
	if v156&int32(3) == int32(0) {
		v163 = v156
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 != 0 {
		v152 = v156
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v188 = v156
	goto L40
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(-2139062144)
	if (int32(16843008)-v173|v173)&v176 == v176 {
		v167 = v167 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v182 = v167
	goto L52
L51:
	;
	goto L50
L52:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 != 0 {
		v182 = v182 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v188 = v182
	goto L40
L54:
	;
	goto L53
L55:
	;
	v202 = F_write(m, v31, int32(_a2155), int32(17))
	mBase = m.M
	if v202 == int32(-1) {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v206 = v8 + int32(16)
	v209 = F___time(m, int32(0))
	mBase = m.M
	if v209 <= int64(-1) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v252 = v8 + int32(16)
	if v252&int32(3) == int32(0) {
		v276 = v252
		goto L68
	} else {
		goto L69
	}
L58:
	;
	goto L57
L60:
	;
	v231 = F_ull2string(m, v227, v228, v229)
	mBase = m.M
	if v231 == int32(0) {
		goto L58
	} else {
		goto L64
	}
L61:
	;
	goto L63
L62:
	;
	v227 = v206
	v228 = int32(64)
	v229 = v209
	goto L60
L63:
	;
	v218 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v218)
	v227 = v8 + int32(17)
	v228 = int32(63)
	v229 = int64(0) - v209
	goto L60
L64:
	;
	goto L57
L66:
	;
	v310 = F_write(m, v31, v252, v309)
	mBase = m.M
	if v310 == int32(-1) {
		goto L10
	} else {
		goto L82
	}
L67:
	;
	v309 = v301 - v252
	goto L66
L68:
	;
	v280 = v276
	goto L76
L69:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v262 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v265 = v252
	goto L72
L71:
	;
	v309 = v252 - v252
	goto L66
L72:
	;
	v269 = v265 + int32(1)
	if v269&int32(3) == int32(0) {
		v276 = v269
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v274 != 0 {
		v265 = v269
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v301 = v269
	goto L67
L76:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v289 = int32(-2139062144)
	if (int32(16843008)-v286|v286)&v289 == v289 {
		v280 = v280 + int32(4)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v295 = v280
	goto L79
L78:
	;
	goto L77
L79:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v299 != 0 {
		v295 = v295 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v301 = v295
	goto L67
L81:
	;
	goto L80
L82:
	;
	v315 = F_write(m, v31, int32(_a2156), int32(2))
	mBase = m.M
	if v315 == int32(-1) {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	if l1&int32(3) == int32(0) {
		v339 = l1
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v373 = F_write(m, v31, l1, v372)
	mBase = m.M
	if v373 == int32(-1) {
		goto L10
	} else {
		goto L100
	}
L85:
	;
	v372 = v364 - l1
	goto L84
L86:
	;
	v343 = v339
	goto L94
L87:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v325 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v328 = l1
	goto L90
L89:
	;
	v372 = l1 - l1
	goto L84
L90:
	;
	v332 = v328 + int32(1)
	if v332&int32(3) == int32(0) {
		v339 = v332
		goto L86
	} else {
		goto L92
	}
L92:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v337 != 0 {
		v328 = v332
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v364 = v332
	goto L85
L94:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v352 = int32(-2139062144)
	if (int32(16843008)-v349|v349)&v352 == v352 {
		v343 = v343 + int32(4)
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v358 = v343
	goto L97
L96:
	;
	goto L95
L97:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v362 != 0 {
		v358 = v358 + int32(1)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v364 = v358
	goto L85
L99:
	;
	goto L98
L100:
	;
	v378 = F_write(m, v31, int32(_a26), int32(1))
	mBase = m.M
	goto L10
L101:
	;
	v383 = F_close(m, v31)
	mBase = m.M
	goto L1
}
func F_serverPubsubShardSubscriptionCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	if v3 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v8 != 0 {
			v10 = F_hashtableSize(m, v8)
			mBase = m.M
			v13 = base.I64_extend_i32_u(v10)
		} else {
			v13 = int64(0)
		}
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+40))
		v13 = v6
	}
	return base.I32_wrap_i64(v13)
}
func F_server_math_random(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	v3 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[985]))
	v16 = int32(58989)
	v17 = v15 * v16
	v20 = int32(65535)
	*(*int32)(unsafe.Add(mBase, _consts[985])) = (v17 + int32(11)) & v20
	v24 = int32(16)
	v30 = int32(base.Ui32(v17)>>(uint(v24)%32)) + base.B2i32(base.Ui32(int32(65524)) < base.Ui32(v17&v20))
	v34 = *(*int32)(unsafe.Add(mBase, _consts[986]))
	v36 = v34 * v16
	v39 = v30&v20 + v36&v20
	v42 = int32(57068)
	v43 = v15 * v42
	v46 = v39&v20 + v43&int32(65532)
	v48 = v46 & v20
	*(*int32)(unsafe.Add(mBase, _consts[986])) = v48
	v63 = *(*int32)(unsafe.Add(mBase, _consts[987]))
	v77 = (int32(base.Ui32(v43)>>(uint(v24)%32)) + v15*int32(5) + v34*v42 + int32(base.Ui32(v36)>>(uint(v24)%32)) + v63*v16 + base.B2i32(base.Ui32(v20) < base.Ui32(v30)) + base.B2i32(base.Ui32(v20) < base.Ui32(v39)) + base.B2i32(base.Ui32(v20) < base.Ui32(v46))) & v20
	*(*int32)(unsafe.Add(mBase, _consts[987])) = v77
	v85 = base.I32_rem_s(v77<<(uint(int32(15))%32)|int32(base.Ui32(v48)>>(uint(int32(1))%32)), int32(2147483647))
	v88 = base.F64_div(base.F64_convert_i32_s(v85), float64(2.147483647e+09))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch (v89 - v90) >> (uint(int32(4)) % 32) {
	case 0:
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = int32(3)
		*(*float64)(unsafe.Add(mBase, uint32(v95))) = v88
		v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v99 + int32(16)
		return int32(1)
	case 1:
		v106 = F_luaL_checkinteger(m, l0, int32(1))
		mBase = m.M
		v109 = m.ExcPending
		if v109 != 0 {
			return int32(0)
		} else {
			if int32(0) < v106 {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = int32(3)
				*(*float64)(unsafe.Add(mBase, uint32(v124))) = base.F64_add(base.F64_floor(base.F64_mul(v88, base.F64_convert_i32_s(v106))), float64(1))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v128 + int32(16)
				return int32(1)
			} else {
				v113 = m.G3
				v116 = F_luaL_argerror(m, l0, int32(1), v113+int32(_a2637))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v124))) = base.F64_add(base.F64_floor(base.F64_mul(v88, base.F64_convert_i32_s(v106))), float64(1))
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v128 + int32(16)
					return int32(1)
				}
			}
		}
	case 2:
		v135 = F_luaL_checkinteger(m, l0, int32(1))
		mBase = m.M
		v136 = m.ExcPending
		if v136 != 0 {
			return int32(0)
		} else {
			v138 = F_luaL_checkinteger(m, l0, int32(2))
			mBase = m.M
			v139 = m.ExcPending
			if v139 != 0 {
				return int32(0)
			} else {
				if v135 <= v138 {
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = int32(3)
					*(*float64)(unsafe.Add(mBase, uint32(v156))) = base.F64_add(base.F64_floor(base.F64_mul(v88, base.F64_convert_i32_s(v138-v135+int32(1)))), base.F64_convert_i32_s(v135))
					v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v160 + int32(16)
					return int32(1)
				} else {
					v142 = m.G3
					v145 = F_luaL_argerror(m, l0, int32(2), v142+int32(_a2637))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v156)+8)) = int32(3)
						*(*float64)(unsafe.Add(mBase, uint32(v156))) = base.F64_add(base.F64_floor(base.F64_mul(v88, base.F64_convert_i32_s(v138-v135+int32(1)))), base.F64_convert_i32_s(v135))
						v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v160 + int32(16)
						return int32(1)
					}
				}
			}
		}
	default:
		v166 = m.G3
		v170 = F_luaL_error(m, l0, v166+int32(_a483), int32(0))
		mBase = m.M
		v171 = m.ExcPending
		if v171 != 0 {
			return int32(0)
		} else {
			return v170
		}
	}
}
