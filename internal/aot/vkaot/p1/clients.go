package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_clientsCronHandleTimeout(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
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
	v5 = base.I64_div_s(l1, int64(1000))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[0]))
	if v7 == int32(0) {
		v53 = int32(0)
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v54&int32(16) == v53 {
			v71 = v53
			return v71
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[1]))
			if v60 == int32(0) {
				v71 = v53
				return v71
			} else {
				v63 = F_clusterRedirectBlockedClientIfNeeded(m, l0)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					if v63 == int32(0) {
						v71 = v53
						return v71
					} else {
						v67 = int32(0)
						F_unblockClientOnError(m, l0, v67)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = v67
							return v71
						}
					}
				}
			}
		}
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
		if v10&int32(2) != 0 {
			v53 = int32(0)
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
			if v54&int32(16) == v53 {
				v71 = v53
				return v71
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[1]))
				if v60 == int32(0) {
					v71 = v53
					return v71
				} else {
					v63 = F_clusterRedirectBlockedClientIfNeeded(m, l0)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 == int32(0) {
							v71 = v53
							return v71
						} else {
							v67 = int32(0)
							F_unblockClientOnError(m, l0, v67)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = v67
								return v71
							}
						}
					}
				}
			}
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			if v14 != int64(-1) {
				v18 = int32(1)
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
				if v19&v18 != 0 {
					v26 = v18
					v29 = v26
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v22 != 0 {
						v24 = F_isImportSlotMigrationJob(m, v22)
						mBase = m.M
						v26 = v24
						v29 = v26
					} else {
						v29 = int32(0)
					}
				}
			} else {
				v29 = int32(1)
			}
			if v29 != 0 {
				v53 = int32(0)
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
				if v54&int32(16) == v53 {
					v71 = v53
					return v71
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[1]))
					if v60 == int32(0) {
						v71 = v53
						return v71
					} else {
						v63 = F_clusterRedirectBlockedClientIfNeeded(m, l0)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 == int32(0) {
								v71 = v53
								return v71
							} else {
								v67 = int32(0)
								F_unblockClientOnError(m, l0, v67)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = v67
									return v71
								}
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v30&int32(262160) != 0 {
					v53 = int32(0)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
					if v54&int32(16) == v53 {
						v71 = v53
						return v71
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[1]))
						if v60 == int32(0) {
							v71 = v53
							return v71
						} else {
							v63 = F_clusterRedirectBlockedClientIfNeeded(m, l0)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v63 == int32(0) {
									v71 = v53
									return v71
								} else {
									v67 = int32(0)
									F_unblockClientOnError(m, l0, v67)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = v67
										return v71
									}
								}
							}
						}
					}
				} else {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
					v36 = int64(*(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[0])))
					if v5-v33 <= v36 {
						v53 = int32(0)
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
						if v54&int32(16) == v53 {
							v71 = v53
							return v71
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[1]))
							if v60 == int32(0) {
								v71 = v53
								return v71
							} else {
								v63 = F_clusterRedirectBlockedClientIfNeeded(m, l0)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(0) {
										v71 = v53
										return v71
									} else {
										v67 = int32(0)
										F_unblockClientOnError(m, l0, v67)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = v67
											return v71
										}
									}
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_clientsCronHandleTimeout[2]))
						if int32(1) < v39 {
							v49 = F_freeClient(m, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return int32(1)
							}
						} else {
							F__serverLog(m, int32(1), int32(_a_F_clientsCronHandleTimeout_0), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = F_freeClient(m, l0)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									return int32(1)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_clientsCronResizeQueryBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-1)))))
	switch v16 & int32(7) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		v33 = int32(0)
		goto L3
	}
L3:
	;
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_clientsCronResizeQueryBuffer[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = int32(-1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38))))
	v42 = v40 & int32(7)
	switch v42 + v38 {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	default:
		v171 = v37
		v173 = v40
		goto L10
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-9))))
	v33 = v32
	goto L3
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-5))))
	v33 = v29
	goto L3
L6:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+int32(-3)))))
	v33 = v26
	goto L3
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(-2)))))
	v33 = v23
	goto L3
L8:
	;
	v33 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+320)) = v197
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v203 == int32(-1) {
		goto L1
	} else {
		goto L60
	}
L10:
	;
	switch v173 & int32(7) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		goto L56
	case 4:
		goto L55
	default:
		v197 = int32(0)
		goto L9
	}
L11:
	;
	if base.Ui32(v79) < base.Ui32(int32(4097)) {
		v164 = v37
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(-9))))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v37+int32(-17))))
	v79 = base.I32_wrap_i64(v70 - v73)
	v80 = base.I32_wrap_i64(int64(base.Ui64(v70) >> (uint(int64(48)) % 64)))
	goto L11
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-5))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
	v79 = v61 - v64
	v80 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
	goto L11
L14:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
	v79 = v54 - v57
	v80 = v54
	goto L11
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-2)))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
	v79 = v47 - v50
	v80 = v50
	goto L11
L16:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-1)))))
	v171 = v164
	v173 = v170
	goto L10
L17:
	;
	if v36-v34 < int64(3) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v158
	if v158 != 0 {
		v164 = v158
		goto L16
	} else {
		goto L54
	}
L19:
	;
	if base.Ui32(v33) < base.Ui32(int32(32769)) {
		v164 = v37
		goto L16
	} else {
		goto L37
	}
L20:
	;
	switch v42 + int32(-1) {
	default:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	}
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v101&int32(1) != 0 {
		v119 = v37
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
	v100 = v99
	goto L21
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
	v100 = v96
	goto L21
L24:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
	v100 = v93
	goto L21
L25:
	;
	v100 = v80 & int32(255)
	goto L21
L26:
	;
	v123 = F_sdsRemoveFreeSpace(m, v119, int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L34
	} else {
		goto L36
	}
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v105 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v112 != 0 {
		v119 = v111
		goto L26
	} else {
		goto L32
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	goto L31
L30:
	;
	v111 = v37
	v112 = int32(0)
	goto L28
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v111 = v110
	v112 = base.B2i32(v107 == int32(1))
	goto L28
L32:
	;
	if v100 != v104 {
		v119 = v111
		goto L26
	} else {
		goto L33
	}
L33:
	;
	F_sdsfree(m, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	v158 = int32(0)
	goto L18
L36:
	;
	v158 = v123
	goto L18
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+320))
	if base.Ui32(int32(base.Ui32(v33)>>(uint(int32(1))%32))) <= base.Ui32(v129) {
		v164 = v37
		goto L16
	} else {
		goto L38
	}
L38:
	;
	switch v42 + int32(-1) {
	default:
		goto L43
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	}
L39:
	;
	if base.Ui32(v129) < base.Ui32(v144) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
	v144 = v143
	goto L39
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
	v144 = v140
	goto L39
L42:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
	v144 = v137
	goto L39
L43:
	;
	v144 = v80 & int32(255)
	goto L39
L44:
	;
	v146 = v144
	goto L46
L45:
	;
	v146 = v129
	goto L46
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v149 = v147 + int32(2)
	if base.Ui32(v149) < base.Ui32(v146) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v151 = v146
	goto L49
L48:
	;
	v151 = v149
	goto L49
L49:
	;
	if v147 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v154 = v146
	goto L52
L51:
	;
	v154 = v151
	goto L52
L52:
	;
	v156 = F_sdsResize(m, v37, v154, int32(1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v158 = v156
	goto L18
L54:
	;
	v197 = int32(0)
	goto L9
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v171+int32(-17))))
	v197 = v195
	goto L9
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v171+int32(-9))))
	v197 = v192
	goto L9
L57:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171+int32(-5)))))
	v197 = v189
	goto L9
L58:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+int32(-3)))))
	v197 = v186
	goto L9
L59:
	;
	v197 = int32(base.Ui32(v173&int32(248)) >> (uint(int32(3)) % 32))
	goto L9
L60:
	;
	v207 = v203 + int32(2)
	if base.Ui32(v207) <= base.Ui32(v197) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+320)) = v207
	goto L1
}
