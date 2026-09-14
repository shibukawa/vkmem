package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F_handleQbLimitReached(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_sdsempty(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[0]))
		v13 = F_catClientInfoString(m, v9, l0, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = F_sdsempty(m)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[0]))
				if v18 == int32(0) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v34 = F_sdscatrepr(m, v15, v32, int32(64))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[1]))
						if int32(3) < v37 {
							v46 = v34
							F_sdsfree(m, v13)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_sdsfree(m, v46)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v51&int32(1280) != 0 {
										v69 = int32(_a_F_handleQbLimitReached_0)
										v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
										*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
										m.G0 = v7 + int32(32)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v51 | int32(1024)
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[3]))
										if v58 == int32(0) {
											v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
											v67 = F_listAddNodeTail(m, v66, l0)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v69 = int32(_a_F_handleQbLimitReached_0)
												v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
												*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
												m.G0 = v7 + int32(32)
												return
											}
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
											v63 = F_listSearchKey(m, v62, l0)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												if v63 != 0 {
													F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_handleQbLimitReached_1), int32(_a_F_handleQbLimitReached_2), int32(2277))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
													v67 = F_listAddNodeTail(m, v66, l0)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v69 = int32(_a_F_handleQbLimitReached_0)
														v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
														*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
														m.G0 = v7 + int32(32)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13
							F__serverLog(m, int32(3), int32(_a_F_handleQbLimitReached_3), v7)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v46 = v34
								F_sdsfree(m, v13)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_sdsfree(m, v46)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										if v51&int32(1280) != 0 {
											v69 = int32(_a_F_handleQbLimitReached_0)
											v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
											*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
											m.G0 = v7 + int32(32)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v51 | int32(1024)
											v58 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[3]))
											if v58 == int32(0) {
												v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
												v67 = F_listAddNodeTail(m, v66, l0)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													v69 = int32(_a_F_handleQbLimitReached_0)
													v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
													*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
													m.G0 = v7 + int32(32)
													return
												}
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
												v63 = F_listSearchKey(m, v62, l0)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return
												} else {
													if v63 != 0 {
														F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_handleQbLimitReached_1), int32(_a_F_handleQbLimitReached_2), int32(2277))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
														v67 = F_listAddNodeTail(m, v66, l0)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v69 = int32(_a_F_handleQbLimitReached_0)
															v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
															*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
															m.G0 = v7 + int32(32)
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
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[1]))
					if int32(3) < v22 {
						v46 = v15
						F_sdsfree(m, v13)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_sdsfree(m, v46)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								if v51&int32(1280) != 0 {
									v69 = int32(_a_F_handleQbLimitReached_0)
									v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
									*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
									m.G0 = v7 + int32(32)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v51 | int32(1024)
									v58 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[3]))
									if v58 == int32(0) {
										v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
										v67 = F_listAddNodeTail(m, v66, l0)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = int32(_a_F_handleQbLimitReached_0)
											v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
											*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
											m.G0 = v7 + int32(32)
											return
										}
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
										v63 = F_listSearchKey(m, v62, l0)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											if v63 != 0 {
												F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_handleQbLimitReached_1), int32(_a_F_handleQbLimitReached_2), int32(2277))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
												v67 = F_listAddNodeTail(m, v66, l0)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													v69 = int32(_a_F_handleQbLimitReached_0)
													v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
													*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
													m.G0 = v7 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v13
						F__serverLog(m, int32(3), int32(_a_F_handleQbLimitReached_4), v7+int32(16))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v46 = v15
							F_sdsfree(m, v13)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_sdsfree(m, v46)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									if v51&int32(1280) != 0 {
										v69 = int32(_a_F_handleQbLimitReached_0)
										v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
										*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
										m.G0 = v7 + int32(32)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v51 | int32(1024)
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[3]))
										if v58 == int32(0) {
											v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
											v67 = F_listAddNodeTail(m, v66, l0)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v69 = int32(_a_F_handleQbLimitReached_0)
												v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
												*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
												m.G0 = v7 + int32(32)
												return
											}
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
											v63 = F_listSearchKey(m, v62, l0)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return
											} else {
												if v63 != 0 {
													F__serverAssertWithInfo(m, l0, int32(0), int32(_a_F_handleQbLimitReached_1), int32(_a_F_handleQbLimitReached_2), int32(2277))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, _c_F_handleQbLimitReached[4]))
													v67 = F_listAddNodeTail(m, v66, l0)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v69 = int32(_a_F_handleQbLimitReached_0)
														v71 = *(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2]))
														*(*int64)(unsafe.Add(mBase, _c_F_handleQbLimitReached[2])) = v71 + int64(1)
														m.G0 = v7 + int32(32)
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
func F_hdr_init(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v98 int32
	_ = v98
	var v106 int64
	_ = v106
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v186 int32
	_ = v186
	v19 = int32(28)
	if l0 < int64(1) {
		v186 = v19
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v186
L2:
	;
	if base.Ui32(l2+int32(-6)) < base.Ui32(int32(-5)) {
		v186 = v19
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 < l0<<(uint(int64(1))%64) {
		v186 = v19
		goto L1
	} else {
		goto L4
	}
L4:
	;
	switch l2 + int32(-1) {
	case 0:
		v36 = float64(20)
		goto L5
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		goto L9
	}
L5:
	;
	v37 = F_log(m, v36)
	mBase = m.M
	v40 = base.F64_ceil(base.F64_div(v37, float64(0.6931471805599453)))
	if base.F64_lt(base.F64_abs(v40), float64(2.147483648e+09)) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v36 = float64(20000)
	goto L5
L7:
	;
	v36 = float64(2000)
	goto L5
L8:
	;
	v36 = float64(200)
	goto L5
L9:
	;
	v36 = float64(200000)
	goto L5
L10:
	;
	v50 = F_log(m, base.F64_convert_i64_u(l0))
	mBase = m.M
	v52 = base.F64_div(v50, float64(0.6931471805599453))
	if base.F64_gt(v52, float64(2.147483647e+09)) != 0 {
		v186 = v19
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v48 = int32(-2147483648)
	goto L10
L12:
	;
	v46 = base.I32_trunc_f64_s(v40)
	v48 = v46
	goto L10
L13:
	;
	v56 = int32(1)
	if v56 < v48 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v70 = v59 + int32(-1)
	v72 = base.I32_div_s(v68, int32(2))
	if base.F64_lt(base.F64_abs(v52), float64(2.147483648e+09)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v68 = int32(-2147483648)
	goto L14
L16:
	;
	v59 = v48
	goto L18
L17:
	;
	v59 = v56
	goto L18
L18:
	;
	v60 = F_scalbn(m, float64(1), v59)
	mBase = m.M
	goto L19
L19:
	;
	if base.F64_lt(base.F64_abs(v60), float64(2.147483648e+09)) == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v66 = base.I32_trunc_f64_s(v60)
	v68 = v66
	goto L14
L21:
	;
	v82 = base.I64_extend_i32_s(v80)
	if int64(61) < base.I64_extend_i32_u(v70)+v82 {
		v186 = v19
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v80 = int32(-2147483648)
	goto L21
L23:
	;
	v78 = base.I32_trunc_f64_s(v52)
	v80 = v78
	goto L21
L24:
	;
	v87 = base.I64_extend_i32_s(v68)
	v89 = v87 << (uint(base.I64_extend_i32_u(v80)) % 64)
	if l1 < v89 {
		v125 = int32(1)
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v140 = (v125 + int32(1)) * v72
	v142 = F_zcalloc_num(m, v140, int32(8))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	v98 = int32(1)
	v106 = v89
	goto L27
L27:
	;
	if v106 < int64(4611686018427387904) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v125 = v115
	goto L25
L29:
	;
	v115 = v98 + int32(1)
	v117 = v106 << (uint(int64(1)) % 64)
	if v117 <= l1 {
		v98 = v115
		v106 = v117
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v125 = v98 + int32(1)
	goto L25
L31:
	;
	goto L28
L32:
	;
	return int32(0)
L33:
	;
	if v142 == int32(0) {
		v186 = int32(48)
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v150 = F_zcalloc_num(m, int32(1), int32(104))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L32
	} else {
		goto L36
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v150)+72)) = int64(4607182418800017408)
	v158 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+64)) = v158
	v161 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v150)+56)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v150)+48)) = int64(9223372036854775807)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+40)) = v68
	*(*int64)(unsafe.Add(mBase, uint32(v150)+32)) = (v87 + int64(-1)) << (uint(v82) % 64)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+28)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v150)+24)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v150)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v150))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v150)+96)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v150)+88)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v150)+80)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v150)+44)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v150
	v186 = v158
	goto L1
L36:
	;
	if v150 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	F_valkey_free(m, v142)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	return int32(48)
}
func F_hdr_iter_log_init(m *base.Module, l0 int32, l1 int32, l2 int64, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v10 int64
	_ = v10
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v46 int64
	_ = v46
	v5 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(32)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(40)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(72)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = l2
	*(*float64)(unsafe.Add(mBase, uint32(l0)+80)) = l3
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(1159)
	v46 = base.I64_extend_i32_u(int32(63) - (v36 + base.I32_wrap_i64(base.I64_clz(v37|l2))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = base.I64_extend32_s(l2>>(uint(v46)%64)) << (uint(v46) % 64)
	return
}
func F_hdr_record_value(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v75 int32
	_ = v75
	if int64(0) <= l1 {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		v21 = v14 + v15 + base.I32_wrap_i64(base.I64_clz(v17|l1))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v32 = (int32(64)-v21)<<(uint(v14)%32) - v24 + base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(base.I64_extend_i32_u(v15-v21+int32(63)))%64)))
		if v32 < v12 {
			v75 = v12
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v35 <= v32 {
				v75 = v12
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v37 == int32(0) {
					v50 = v32
				} else {
					v40 = int32(0)
					v43 = v32 - v37
					if v43 < v35 {
						v45 = v40
					} else {
						v45 = v40 - v35
					}
					if v43 < int32(0) {
						v48 = v35
					} else {
						v48 = v45
					}
					v50 = v48 + v43
				}
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v54 = v51 + v50<<(uint(int32(3))%32)
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
				v56 = int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(v54))) = v55 + v56
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v59 + v56
				v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				if v63 < l1 {
					v65 = l1
				} else {
					v65 = v63
				}
				*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v65
				v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < v67 {
					v69 = l1
				} else {
					v69 = v67
				}
				if l1 == int64(0) {
					v72 = v67
				} else {
					v72 = v69
				}
				*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v72
				v75 = int32(1)
			}
		}
		return v75
	} else {
		return int32(0)
	}
}
func F_hdr_value_at_percentile(m *base.Module, l0 int32, l1 float64) int64 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 float64
	_ = v13
	var v16 int64
	_ = v16
	var v20 float64
	_ = v20
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v90 int64
	_ = v90
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v120 int64
	_ = v120
	v10 = float64(100)
	if base.F64_lt(l1, v10) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v29 < int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v28 = int64(-9223372036854775807 - 1)
	goto L1
L3:
	;
	v13 = l1
	goto L5
L4:
	;
	v13 = v10
	goto L5
L5:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v20 = base.F64_add(base.F64_mul(base.F64_div(v13, float64(100)), base.F64_convert_i64_s(v16)), float64(0.5))
	if base.F64_lt(base.F64_abs(v20), float64(9.223372036854776e+18)) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = base.I64_trunc_f64_s(v20)
	v28 = v26
	goto L1
L7:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v103 = int32(63) - (v97 + base.I32_wrap_i64(base.I64_clz(v98|v90)))
	v104 = base.I64_extend_i32_u(v103)
	v105 = v90 >> (uint(v104) % 64)
	v107 = base.I64_extend32_s(v105) << (uint(v104) % 64)
	if base.F64_eq(l1, float64(0)) != 0 {
		v120 = v107
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v90 = int64(0)
	goto L7
L9:
	;
	v32 = int64(1)
	if v32 < v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v28
	goto L12
L11:
	;
	v35 = v32
	goto L12
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v42 = int64(0)
	v46 = int32(0)
	goto L13
L13:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v36+v46<<(uint(int32(3))%32))))
	v52 = v51 + v42
	if v52 < v35 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L8
L15:
	;
	v75 = v46 + int32(1)
	if v75 != v29 {
		v42 = v52
		v46 = v75
		goto L13
	} else {
		goto L23
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v60 = int32(base.Ui32(v46) >> (uint(v59) % 32))
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v54
	goto L19
L18:
	;
	v61 = int32(0)
	goto L19
L19:
	;
	v64 = int32(1)
	if v64 < v60 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v67 = v60
	goto L22
L21:
	;
	v67 = v64
	goto L22
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v90 = base.I64_extend_i32_s((v54+int32(2147483647))&v46+v61) << (uint(base.I64_extend_i32_u(v67+v68+int32(-1))) % 64)
	goto L7
L23:
	;
	goto L14
L24:
	;
	return v120
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v120 = v107 + int64(1)<<(uint(base.I64_extend_i32_u(v103+base.B2i32(v111 <= base.I32_wrap_i64(v105))))%64) + int64(-1)
	goto L24
}
func F_helloCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v339 int64
	_ = v339
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v19 < int32(2) {
		v36 = int32(1)
		v37 = v19
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if v339 == int64(0) {
		goto L92
	} else {
		goto L93
	}
L3:
	;
	F_addReplyErrorLength(m, l0, int32(_a_F_helloCommand_0), int32(214))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L8
	} else {
		goto L90
	}
L4:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v305&int32(128) != 0 {
		goto L2
	} else {
		goto L89
	}
L5:
	;
	F_addReplyErrorLength(m, l0, int32(_a_F_helloCommand_1), int32(37))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L8
	} else {
		goto L87
	}
L6:
	;
	if v37 <= v36 {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = F_getLongLongFromObjectOrReply(m, l0, v23, v14+int32(8), int32(_a_F_helloCommand_2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui64(v29+int64(-4)) < base.Ui64(int64(-2)) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v36 = int32(2)
	v37 = v34
	goto L6
L12:
	;
	v39 = int32(0)
	v44 = v36
	v45 = v37
	v46 = v39
	v47 = v39
	v48 = v39
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v55 = v44 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v55)))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a_F_helloCommand_3)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v251 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L15:
	;
	v100 = v45 + (v44 ^ int32(-1))
	if v100 < int32(2) {
		goto L29
	} else {
		goto L30
	}
L16:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L15
L17:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L20
L18:
	;
	v90 = int32(0)
	v91 = v59
	goto L16
L19:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L16
L20:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v86 = v80
	v87 = int32(0)
	goto L19
L22:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L20
	} else {
		goto L26
	}
L24:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L19
L26:
	;
	goto L21
L27:
	;
	v259 = v256 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v259 < v260 {
		v44 = v259
		v45 = v260
		v46 = v251
		v47 = v252
		v48 = v253
		goto L13
	} else {
		goto L73
	}
L28:
	;
	F__serverAssert(m, int32(_a_F_helloCommand_4), int32(_a_F_helloCommand_5), int32(6052))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L72
	}
L29:
	;
	v130 = int32(_a_F_helloCommand_6)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	if v94-v96 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v44 <= int32(-1) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32(v44) < base.Ui32(int32(31)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v110 = int32(2) << (uint(v44) % 32)
	goto L35
L34:
	;
	v110 = int32(1)
	goto L35
L35:
	;
	if base.Ui32(v44) < base.Ui32(int32(30)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = int32(4) << (uint(v44) % 32)
	goto L38
L37:
	;
	v116 = int32(1)
	goto L38
L38:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v110 | v116 | v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v122 = int32(2)
	v123 = v44 + v122
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+v123<<(uint(v122)%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121+v55)+4))
	v251 = v129
	v252 = v47
	v253 = v127
	v256 = v123
	goto L27
L39:
	;
	if v100 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v165 = F_tolower(m, v161)
	mBase = m.M
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v167 = F_tolower(m, v166)
	mBase = m.M
	goto L39
L41:
	;
	v135 = v58
	v136 = v130
	v137 = v133
	goto L44
L42:
	;
	v161 = int32(0)
	v162 = v130
	goto L40
L43:
	;
	v161 = v158 & int32(255)
	v162 = v157
	goto L40
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v139 == int32(0) {
		v157 = v136
		v158 = v137
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v157 = v151
	v158 = int32(0)
	goto L43
L46:
	;
	v143 = v137 & int32(255)
	if v143 == v139 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v150 = int32(1)
	v151 = v136 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v152 != 0 {
		v135 = v135 + v150
		v136 = v151
		v137 = v152
		goto L44
	} else {
		goto L50
	}
L48:
	;
	v145 = F_tolower(m, v143)
	mBase = m.M
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v147 = F_tolower(m, v146)
	mBase = m.M
	if v145 == v147 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v157 = v136
	v158 = v149
	goto L43
L50:
	;
	goto L45
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v58
	F_addReplyErrorFormat(m, l0, int32(_a_F_helloCommand_7), v14)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L8
	} else {
		goto L71
	}
L52:
	;
	if v165-v167 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v173 = v44 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173<<(uint(int32(2))%32))))
	if v177 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v179 = F_objectGetVal(m, v177)
	mBase = m.M
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+int32(-1)))))
	switch v182 & int32(7) {
	case 0:
		goto L61
	case 1:
		goto L60
	case 2:
		goto L59
	case 3:
		goto L58
	case 4:
		goto L57
	default:
		v251 = v46
		v252 = v177
		v253 = v48
		v256 = v173
		goto L27
	}
L55:
	;
	v251 = v46
	v252 = int32(0)
	v253 = v48
	v256 = v173
	goto L27
L56:
	;
	if v199 == int32(0) {
		v251 = v46
		v252 = v177
		v253 = v48
		v256 = v173
		goto L27
	} else {
		goto L62
	}
L57:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v179+int32(-17))))
	v199 = v198
	goto L56
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v179+int32(-9))))
	v199 = v195
	goto L56
L59:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179+int32(-5)))))
	v199 = v192
	goto L56
L60:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+int32(-3)))))
	v199 = v189
	goto L56
L61:
	;
	v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
	goto L56
L62:
	;
	v202 = F_objectGetVal(m, v177)
	mBase = m.M
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v203 == int32(0) {
		v251 = v46
		v252 = v177
		v253 = v48
		v256 = v173
		goto L27
	} else {
		goto L63
	}
L63:
	;
	v208 = v203
	v209 = v202
	goto L65
L64:
	;
	F_addReplyErrorLength(m, l0, int32(_a_F_helloCommand_8), int32(67))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L69
	}
L65:
	;
	if base.Ui32((v208+int32(-127))&int32(255)) < base.Ui32(int32(162)) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v223 == int32(0) {
		v251 = v46
		v252 = v177
		v253 = v48
		v256 = v173
		goto L27
	} else {
		goto L68
	}
L68:
	;
	v208 = v223
	v209 = v209 + int32(1)
	goto L65
L69:
	;
	F_afterErrorReply(m, l0, int32(_a_F_helloCommand_8), int32(67), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	goto L1
L71:
	;
	goto L1
L72:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	goto L14
L74:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v286&int32(128) == int32(0) {
		goto L3
	} else {
		goto L84
	}
L75:
	;
	if v253 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(0)
	v270 = F_ACLAuthenticateUser(m, l0, v251, v253, v14+int32(4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L78
	}
L77:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v277 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	if v270 != int32(1) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_addAuthErrReply(m, l0, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	switch v270 + int32(-1) {
	case 0, 2:
		goto L1
	default:
		goto L74
	}
L82:
	;
	F_decrRefCount(m, v277)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if v252 == int32(0) {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v294 = F_clientSetName(m, l0, v252, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	goto L2
L87:
	;
	F_afterErrorReply(m, l0, int32(_a_F_helloCommand_1), int32(37), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	goto L3
L90:
	;
	F_afterErrorReply(m, l0, int32(_a_F_helloCommand_0), int32(214), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	goto L1
L92:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[0]))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+int32(-1)))))
	switch v348 & int32(7) {
	case 0:
		goto L99
	case 1:
		goto L98
	case 2:
		goto L97
	case 3:
		goto L96
	case 4:
		goto L95
	default:
		v365 = int32(0)
		goto L94
	}
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)) = uint8(v339)
	goto L92
L94:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[1]))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v369 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L101
	}
L95:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v345+int32(-17))))
	v365 = v364
	goto L94
L96:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v345+int32(-9))))
	v365 = v361
	goto L94
L97:
	;
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345+int32(-5)))))
	v365 = v358
	goto L94
L98:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+int32(-3)))))
	v365 = v355
	goto L94
L99:
	;
	v365 = int32(base.Ui32(v348) >> (uint(int32(3)) % 32))
	goto L94
L100:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_9), int32(6))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L8
	} else {
		goto L110
	}
L101:
	;
	if v369 != 0 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	if v367 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v373 = int32(6)
	goto L105
L104:
	;
	v373 = int32(7)
	goto L105
L105:
	;
	v380 = base.B2i32(v368&int32(255) == int32(2))
	if v368&int32(255) == int32(2) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v385 = int32(42)
	goto L108
L107:
	;
	v385 = int32(37)
	goto L108
L108:
	;
	F__addReplyLongLongWithPrefix(m, l0, base.I64_extend_i32_u((v373+base.B2i32(v365 != int32(0)))<<(uint(v380)%32)), v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L100
L110:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[2]))
	if v396 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v397 = int32(_a_F_helloCommand_10)
	goto L113
L112:
	;
	v397 = int32(_a_F_helloCommand_11)
	goto L113
L113:
	;
	if v396 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v400 = int32(5)
	goto L116
L115:
	;
	v400 = int32(6)
	goto L116
L116:
	;
	F_addReplyBulkCBuffer(m, l0, v397, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_12), int32(7))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[2]))
	if v410 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v411 = int32(_a_F_helloCommand_13)
	goto L121
L120:
	;
	v411 = int32(_a_F_helloCommand_14)
	goto L121
L121:
	;
	F_addReplyBulkCBuffer(m, l0, v411, int32(5))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_15), int32(5))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	v419 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	F_addReplyLongLong(m, l0, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_16), int32(2))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_addReplyLongLong(m, l0, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_17), int32(4))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[1]))
	if v434 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[1]))
	if v454 != 0 {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[3]))
	if v442 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_18), int32(8))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_19), int32(10))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L8
	} else {
		goto L135
	}
L133:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_20), int32(7))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	goto L128
L135:
	;
	goto L128
L136:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_21), int32(7))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L146
	}
L137:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_22), int32(4))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L8
	} else {
		goto L138
	}
L138:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[4]))
	if v462 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v463 = int32(_a_F_helloCommand_23)
	goto L141
L140:
	;
	v463 = int32(_a_F_helloCommand_24)
	goto L141
L141:
	;
	if v462 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v466 = int32(7)
	goto L144
L143:
	;
	v466 = int32(6)
	goto L144
L144:
	;
	F_addReplyBulkCBuffer(m, l0, v463, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	goto L136
L146:
	;
	F_addReplyLoadedModules(m, l0)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L147
	}
L147:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[0]))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+int32(-1)))))
	switch v480 & int32(7) {
	case 0:
		goto L153
	case 1:
		goto L152
	case 2:
		goto L151
	case 3:
		goto L150
	case 4:
		goto L149
	default:
		goto L1
	}
L148:
	;
	if v497 == int32(0) {
		goto L1
	} else {
		goto L154
	}
L149:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v477+int32(-17))))
	v497 = v496
	goto L148
L150:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v477+int32(-9))))
	v497 = v493
	goto L148
L151:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v477+int32(-5)))))
	v497 = v490
	goto L148
L152:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+int32(-3)))))
	v497 = v487
	goto L148
L153:
	;
	v497 = int32(base.Ui32(v480) >> (uint(int32(3)) % 32))
	goto L148
L154:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a_F_helloCommand_25), int32(17))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_helloCommand[0]))
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506+int32(-1)))))
	switch v509 & int32(7) {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	case 3:
		goto L158
	case 4:
		goto L157
	default:
		v526 = int32(0)
		goto L156
	}
L156:
	;
	F_addReplyBulkCBuffer(m, l0, v506, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L8
	} else {
		goto L162
	}
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v506+int32(-17))))
	v526 = v525
	goto L156
L158:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v506+int32(-9))))
	v526 = v522
	goto L156
L159:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506+int32(-5)))))
	v526 = v519
	goto L156
L160:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506+int32(-3)))))
	v526 = v516
	goto L156
L161:
	;
	v526 = int32(base.Ui32(v509) >> (uint(int32(3)) % 32))
	goto L156
L162:
	;
	goto L1
}
func F_hex_digit_to_int(m *base.Module, l0 int32) int32 {
	var v35 int32
	_ = v35
	switch l0 + int32(-49) {
	case 0:
		v35 = int32(1)
		return v35
	case 1:
		return int32(2)
	case 2:
		return int32(3)
	case 3:
		return int32(4)
	case 4:
		return int32(5)
	case 5:
		return int32(6)
	case 6:
		return int32(7)
	case 7:
		return int32(8)
	case 8:
		return int32(9)
	default:
		v35 = int32(0)
		return v35
	case 16, 48:
		return int32(10)
	case 17, 49:
		return int32(11)
	case 18, 50:
		return int32(12)
	case 19, 51:
		return int32(13)
	case 20, 52:
		return int32(14)
	case 21, 53:
		return int32(15)
	}
}
func F_hexfloat(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v98 int64
	_ = v98
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int64
	_ = v227
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v258 int64
	_ = v258
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v276 int64
	_ = v276
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v291 int64
	_ = v291
	var v296 int64
	_ = v296
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v338 int32
	_ = v338
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int64
	_ = v388
	var v390 int64
	_ = v390
	var v394 int64
	_ = v394
	var v414 int64
	_ = v414
	var v428 int32
	_ = v428
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v473 int64
	_ = v473
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v486 int64
	_ = v486
	var v501 int32
	_ = v501
	var v514 int64
	_ = v514
	var v515 int32
	_ = v515
	var v520 int64
	_ = v520
	var v523 int64
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v544 int64
	_ = v544
	var v545 int64
	_ = v545
	var v549 int32
	_ = v549
	var v555 int64
	_ = v555
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int64
	_ = v571
	var v573 int64
	_ = v573
	var v577 int64
	_ = v577
	var v597 int64
	_ = v597
	var v611 int32
	_ = v611
	var v620 int64
	_ = v620
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v625 int64
	_ = v625
	var v626 int64
	_ = v626
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v648 int64
	_ = v648
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int64
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v680 int64
	_ = v680
	var v695 int64
	_ = v695
	var v696 int64
	_ = v696
	var v697 int64
	_ = v697
	var v706 int64
	_ = v706
	var v711 int64
	_ = v711
	var v712 int64
	_ = v712
	var v713 int64
	_ = v713
	var v717 int64
	_ = v717
	var v722 int64
	_ = v722
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v750 int32
	_ = v750
	var v751 int64
	_ = v751
	var v752 int64
	_ = v752
	var v755 int64
	_ = v755
	var v761 int64
	_ = v761
	var v765 int64
	_ = v765
	var v769 int32
	_ = v769
	var v773 int64
	_ = v773
	var v774 int64
	_ = v774
	var v778 int32
	_ = v778
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int64
	_ = v824
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v832 int64
	_ = v832
	var v833 int64
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int64
	_ = v839
	var v844 int64
	_ = v844
	var v845 int64
	_ = v845
	var v860 int32
	_ = v860
	var v861 int64
	_ = v861
	var v862 int64
	_ = v862
	var v865 int64
	_ = v865
	var v872 int64
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v893 int64
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v907 int64
	_ = v907
	var v922 int64
	_ = v922
	var v923 int64
	_ = v923
	var v924 int64
	_ = v924
	var v933 int64
	_ = v933
	var v934 int64
	_ = v934
	var v935 int64
	_ = v935
	var v938 int32
	_ = v938
	var v939 float64
	_ = v939
	var v941 int32
	_ = v941
	var v945 float64
	_ = v945
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v961 float64
	_ = v961
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 float64
	_ = v974
	var v975 int32
	_ = v975
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int64
	_ = v993
	var v995 int64
	_ = v995
	var v999 int64
	_ = v999
	var v1019 int64
	_ = v1019
	var v1033 int32
	_ = v1033
	var v1042 int64
	_ = v1042
	var v1045 int64
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1047 int64
	_ = v1047
	var v1048 int64
	_ = v1048
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int64
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1085 int64
	_ = v1085
	var v1100 int64
	_ = v1100
	var v1101 int64
	_ = v1101
	var v1102 int64
	_ = v1102
	var v1110 int32
	_ = v1110
	var v1111 int64
	_ = v1111
	var v1116 int64
	_ = v1116
	var v1117 int64
	_ = v1117
	var v1122 int64
	_ = v1122
	var v1124 int64
	_ = v1124
	var v1146 int64
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1148 int64
	_ = v1148
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1151 int64
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int64
	_ = v1160
	var v1169 int64
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1198 int32
	_ = v1198
	var v1210 int32
	_ = v1210
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int64
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1248 int64
	_ = v1248
	var v1257 int64
	_ = v1257
	var v1259 int64
	_ = v1259
	var v1260 int64
	_ = v1260
	var v1268 int64
	_ = v1268
	var v1273 int64
	_ = v1273
	var v1277 int64
	_ = v1277
	var v1282 int64
	_ = v1282
	var v1287 int64
	_ = v1287
	var v1289 int64
	_ = v1289
	var v1293 int64
	_ = v1293
	var v1298 int64
	_ = v1298
	var v1299 int64
	_ = v1299
	var v1304 int64
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int64
	_ = v1308
	var v1313 int64
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1322 int64
	_ = v1322
	var v1325 int64
	_ = v1325
	var v1331 int64
	_ = v1331
	var v1336 int64
	_ = v1336
	var v1337 int64
	_ = v1337
	var v1346 int64
	_ = v1346
	var v1347 int64
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1375 int32
	_ = v1375
	var v1387 int32
	_ = v1387
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1423 int64
	_ = v1423
	var v1424 int64
	_ = v1424
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1444 int64
	_ = v1444
	var v1445 int64
	_ = v1445
	var v1457 int64
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1478 int64
	_ = v1478
	var v1479 int64
	_ = v1479
	var v1480 int64
	_ = v1480
	var v1481 int64
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1492 int64
	_ = v1492
	var v1494 int64
	_ = v1494
	var v1503 int64
	_ = v1503
	var v1504 int64
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1518 int64
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1532 int64
	_ = v1532
	var v1547 int64
	_ = v1547
	var v1548 int64
	_ = v1548
	var v1549 int64
	_ = v1549
	var v1558 int64
	_ = v1558
	var v1563 int64
	_ = v1563
	var v1564 int64
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1569 int64
	_ = v1569
	var v1574 int64
	_ = v1574
	var v1582 int64
	_ = v1582
	var v1583 int64
	_ = v1583
	var v1597 int64
	_ = v1597
	var v1601 int64
	_ = v1601
	v22 = m.G0
	v24 = v22 - int32(432)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v26 == v27 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = int32(0)
	v46 = v35
	v49 = v36
	goto L9
L2:
	;
	v33 = F___shgetc(m, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v26 + int32(1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v35 = v32
	goto L1
L4:
	;
	return
L5:
	;
	v35 = v33
	goto L1
L6:
	;
	v148 = int32(0)
	v149 = int64(0)
	v161 = v133
	v162 = v134
	v163 = v135
	v164 = v136
	v165 = int64(4611404543450677248)
	v166 = v148
	v167 = v149
	v168 = v149
	v169 = v149
	v170 = v148
	v171 = v149
	goto L29
L7:
	;
	v85 = int64(0)
	if v84 == int32(48) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v82 = F___shgetc(m, l1)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L18
	}
L9:
	;
	if v46 == int32(48) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v71 == v72 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v46 != int32(46) {
		v133 = v46
		v134 = v36
		v135 = int64(0)
		v136 = v49
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v64 == v65 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v64 + int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v84 = v70
	goto L7
L15:
	;
	v80 = F___shgetc(m, l1)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v74 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v71 + v74
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v46 = v78
	v49 = v74
	goto L9
L17:
	;
	v46 = v80
	v49 = int32(1)
	goto L9
L18:
	;
	v84 = v82
	goto L7
L19:
	;
	v98 = v85
	goto L21
L20:
	;
	v133 = v84
	v134 = int32(1)
	v135 = v85
	v136 = v49
	goto L6
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v110 == v111 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v124 = int32(1)
	v133 = v119
	v134 = v124
	v135 = v121
	v136 = v124
	goto L6
L23:
	;
	v121 = v98 + int64(-1)
	if v119 == int32(48) {
		v98 = v121
		goto L21
	} else {
		goto L27
	}
L24:
	;
	v117 = F___shgetc(m, l1)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v110 + int32(1)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v119 = v116
	goto L23
L26:
	;
	v119 = v117
	goto L23
L27:
	;
	goto L22
L28:
	;
	if v164 != 0 {
		goto L57
	} else {
		goto L58
	}
L29:
	;
	v176 = v161 + int32(-48)
	if base.Ui32(v176) < base.Ui32(int32(10)) {
		v190 = v161
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v326 == v327 {
		goto L53
	} else {
		goto L54
	}
L32:
	;
	if int32(57) < v161 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v180 = v161 | int32(32)
	if v161 == int32(46) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v161 != int32(46) {
		v190 = v180
		goto L32
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(5)) < base.Ui32(v180+int32(-97)) {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v162 != 0 {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v315 = int32(1)
	v316 = v171
	v317 = v164
	v318 = v165
	v319 = v166
	v320 = v167
	v321 = v168
	v322 = v169
	v323 = v170
	v324 = v171
	goto L31
L39:
	;
	v195 = v190 + int32(-87)
	goto L41
L40:
	;
	v195 = v176
	goto L41
L41:
	;
	if int64(7) < v171 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v315 = v162
	v316 = v163
	v317 = int32(1)
	v318 = v305
	v319 = v306
	v320 = v307
	v321 = v308
	v322 = v309
	v323 = v310
	v324 = v171 + int64(1)
	goto L31
L43:
	;
	if base.Ui64(int64(28)) < base.Ui64(v171) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v305 = v165
	v306 = v195 + v166<<(uint(int32(4))%32)
	v307 = v167
	v308 = v168
	v309 = v169
	v310 = v170
	goto L42
L45:
	;
	if v195 == int32(0) {
		v305 = v165
		v306 = v166
		v307 = v167
		v308 = v168
		v309 = v169
		v310 = v170
		goto L42
	} else {
		goto L51
	}
L46:
	;
	v204 = v24 + int32(48)
	v209 = m.G0
	v211 = v209 - int32(16)
	m.G0 = v211
	if v195 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F___multf3(m, v24+int32(32), v169, v165, int64(0), int64(4610278643543834624))
	mBase = m.M
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(56))))
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v24)+32))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(40))))
	F___multf3(m, v24+int32(16), v258, v263, v264, v269)
	mBase = m.M
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(24))))
	F___addtf3(m, v24, v271, v276, v167, v168)
	mBase = m.M
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(8))))
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v305 = v269
	v306 = v166
	v307 = v281
	v308 = v280
	v309 = v264
	v310 = v170
	goto L42
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v204)+8)) = v244
	m.G0 = v211 + int32(16)
	goto L47
L49:
	;
	v216 = v195 >> (uint(int32(31)) % 32)
	v218 = v195 ^ v216 - v216
	v221 = base.I32_clz(v218)
	F___ashlti3(m, v211, base.I64_extend_i32_u(v218), int64(0), v221+int32(81))
	mBase = m.M
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v211+int32(8))))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v211)))
	v243 = v242
	v244 = v227 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v221)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v195&int32(-2147483648))<<(uint(int64(32))%64)
	goto L48
L50:
	;
	v213 = int64(0)
	v243 = v213
	v244 = v213
	goto L48
L51:
	;
	if v170 != 0 {
		v305 = v165
		v306 = v166
		v307 = v167
		v308 = v168
		v309 = v169
		v310 = v170
		goto L42
	} else {
		goto L52
	}
L52:
	;
	F___multf3(m, v24+int32(80), v169, v165, int64(0), int64(4611123068473966592))
	mBase = m.M
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v24)+80))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(88))))
	F___addtf3(m, v24+int32(64), v291, v296, v167, v168)
	mBase = m.M
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(72))))
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v24)+64))
	v305 = v165
	v306 = v166
	v307 = v304
	v308 = v302
	v309 = v169
	v310 = int32(1)
	goto L42
L53:
	;
	v333 = F___shgetc(m, l1)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v326 + int32(1)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	v161 = v332
	v162 = v315
	v163 = v316
	v164 = v317
	v165 = v318
	v166 = v319
	v167 = v320
	v168 = v321
	v169 = v322
	v170 = v323
	v171 = v324
	goto L29
L55:
	;
	v161 = v333
	v162 = v315
	v163 = v316
	v164 = v317
	v165 = v318
	v166 = v319
	v167 = v320
	v168 = v321
	v169 = v322
	v170 = v323
	v171 = v324
	goto L29
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1597
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v1601
	m.G0 = v24 + int32(432)
	return
L57:
	;
	if int64(7) < v171 {
		v501 = v166
		goto L81
	} else {
		goto L82
	}
L58:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v335 < int64(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v374 = v24 + int32(96)
	v384 = m.G0
	v386 = v384 - int32(16)
	m.G0 = v386
	v388 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(l4)))
	v390 = v388 & int64(4503599627370495)
	v394 = int64(base.Ui64(v388)>>(uint(int64(52))%64)) & int64(2047)
	if v394 == int64(0) {
		goto L72
	} else {
		goto L73
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = int64(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v357 - v358)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L67
L61:
	;
	if l5 != 0 {
		goto L59
	} else {
		goto L65
	}
L62:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v338 + int32(-1)
	if l5 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v338 + int32(-2)
	if v162 == int32(0) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v338 + int32(-3)
	goto L59
L65:
	;
	goto L60
L66:
	;
	goto L59
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v362
	goto L66
L70:
	;
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(104))))
	v459 = *(*int64)(unsafe.Add(mBase, uint32(v24)+96))
	v1597 = v459
	v1601 = v458
	goto L56
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v374))) = v441
	*(*int64)(unsafe.Add(mBase, uint32(v374)+8)) = v442<<(uint(int64(48))%64) | v388&int64(-9223372036854775807-1) | v443
	m.G0 = v386 + int32(16)
	goto L70
L72:
	;
	if base.B2i32(v390 == int64(0)) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v394 == int64(2047) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v441 = v390 << (uint(int64(60)) % 64)
	v442 = int64(32767)
	v443 = int64(base.Ui64(v390) >> (uint(int64(4)) % 64))
	goto L71
L75:
	;
	v441 = v390 << (uint(int64(60)) % 64)
	v442 = v394 + int64(15360)
	v443 = int64(base.Ui64(v390) >> (uint(int64(4)) % 64))
	goto L71
L76:
	;
	if base.Ui64(v390) < base.Ui64(int64(4294967296)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v414 = int64(0)
	v441 = v414
	v442 = v414
	v443 = v414
	goto L71
L78:
	;
	v428 = base.I32_clz(base.I32_wrap_i64(v388)) | int32(32)
	goto L80
L79:
	;
	v428 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v390) >> (uint(int64(32)) % 64))))
	goto L80
L80:
	;
	F___ashlti3(m, v386, v390, int64(0), v428+int32(49))
	mBase = m.M
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v386+int32(8))))
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v386)))
	v441 = v440
	v442 = base.I64_extend_i32_u(int32(15372) - v428)
	v443 = v437 ^ int64(281474976710656)
	goto L71
L81:
	;
	if v161&int32(-33) != int32(80) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v473 = v171
	v474 = v166
	goto L83
L83:
	;
	v484 = v474 << (uint(int32(4)) % 32)
	v486 = v473 + int64(1)
	if v486 != int64(8) {
		v473 = v486
		v474 = v484
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v501 = v484
	goto L81
L85:
	;
	goto L84
L86:
	;
	if v501 != 0 {
		goto L101
	} else {
		goto L102
	}
L87:
	;
	v555 = int64(0)
	goto L86
L88:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v549 + int32(-1)
	goto L87
L89:
	;
	v544 = int64(0)
	v545 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v545 < v544 {
		v555 = v544
		goto L86
	} else {
		goto L100
	}
L90:
	;
	v514 = F_scanexp(m, l1, l5)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v514 != int64(-9223372036854775807-1) {
		v555 = v514
		goto L86
	} else {
		goto L92
	}
L92:
	;
	if l5 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v523 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v523
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v528 - v529)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L97
L94:
	;
	v520 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(-1) < v520 {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	goto L87
L96:
	;
	v1597 = v523
	v1601 = int64(0)
	goto L56
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v533
	goto L96
L100:
	;
	goto L88
L101:
	;
	if v162 != 0 {
		goto L115
	} else {
		goto L116
	}
L102:
	;
	v557 = v24 + int32(112)
	v567 = m.G0
	v569 = v567 - int32(16)
	m.G0 = v569
	v571 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(l4)))
	v573 = v571 & int64(4503599627370495)
	v577 = int64(base.Ui64(v571)>>(uint(int64(52))%64)) & int64(2047)
	if v577 == int64(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(120))))
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v24)+112))
	v1597 = v642
	v1601 = v641
	goto L56
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v557))) = v624
	*(*int64)(unsafe.Add(mBase, uint32(v557)+8)) = v625<<(uint(int64(48))%64) | v571&int64(-9223372036854775807-1) | v626
	m.G0 = v569 + int32(16)
	goto L103
L105:
	;
	if base.B2i32(v573 == int64(0)) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	if v577 == int64(2047) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v624 = v573 << (uint(int64(60)) % 64)
	v625 = int64(32767)
	v626 = int64(base.Ui64(v573) >> (uint(int64(4)) % 64))
	goto L104
L108:
	;
	v624 = v573 << (uint(int64(60)) % 64)
	v625 = v577 + int64(15360)
	v626 = int64(base.Ui64(v573) >> (uint(int64(4)) % 64))
	goto L104
L109:
	;
	if base.Ui64(v573) < base.Ui64(int64(4294967296)) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v597 = int64(0)
	v624 = v597
	v625 = v597
	v626 = v597
	goto L104
L111:
	;
	v611 = base.I32_clz(base.I32_wrap_i64(v571)) | int32(32)
	goto L113
L112:
	;
	v611 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v573) >> (uint(int64(32)) % 64))))
	goto L113
L113:
	;
	F___ashlti3(m, v569, v573, int64(0), v611+int32(49))
	mBase = m.M
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v569+int32(8))))
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v569)))
	v624 = v623
	v625 = base.I64_extend_i32_u(int32(15372) - v611)
	v626 = v620 ^ int64(281474976710656)
	goto L104
L114:
	;
	if v648 < base.I64_extend_i32_s(l3+int32(-226)) {
		goto L124
	} else {
		goto L125
	}
L115:
	;
	v643 = v163
	goto L117
L116:
	;
	v643 = v171
	goto L117
L117:
	;
	v648 = v643<<(uint(int64(2))%64) + v555 + int64(-32)
	if v648 <= base.I64_extend_i32_u(int32(0)-l3) {
		goto L114
	} else {
		goto L118
	}
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hexfloat[0])) = int32(68)
	v657 = v24 + int32(160)
	v662 = m.G0
	v664 = v662 - int32(16)
	m.G0 = v664
	if l4 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v24)+160))
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(168))))
	v712 = int64(-1)
	v713 = int64(9223090561878065151)
	F___multf3(m, v24+int32(144), v706, v711, v712, v713)
	mBase = m.M
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v24)+144))
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(152))))
	F___multf3(m, v24+int32(128), v717, v722, v712, v713)
	mBase = m.M
	v730 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(136))))
	v731 = *(*int64)(unsafe.Add(mBase, uint32(v24)+128))
	v1597 = v731
	v1601 = v730
	goto L56
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v657))) = v696
	*(*int64)(unsafe.Add(mBase, uint32(v657)+8)) = v697
	m.G0 = v664 + int32(16)
	goto L120
L122:
	;
	v669 = l4 >> (uint(int32(31)) % 32)
	v671 = l4 ^ v669 - v669
	v674 = base.I32_clz(v671)
	F___ashlti3(m, v664, base.I64_extend_i32_u(v671), int64(0), v674+int32(81))
	mBase = m.M
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v664+int32(8))))
	v695 = *(*int64)(unsafe.Add(mBase, uint32(v664)))
	v696 = v695
	v697 = v680 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v674)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l4&int32(-2147483648))<<(uint(int64(32))%64)
	goto L121
L123:
	;
	v666 = int64(0)
	v696 = v666
	v697 = v666
	goto L121
L124:
	;
	goto L280
L125:
	;
	if v501 <= int32(-1) {
		v860 = v501
		v861 = v167
		v862 = v168
		v865 = v648
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v872 = v865 + base.I64_extend_i32_u(int32(32)-l3)
	v873 = base.I32_wrap_i64(v872)
	v874 = int32(0)
	if v874 < v873 {
		goto L161
	} else {
		goto L162
	}
L127:
	;
	v750 = v501
	v751 = v167
	v752 = v168
	v755 = v648
	goto L128
L128:
	;
	v761 = int64(0)
	F___addtf3(m, v24+int32(416), v751, v752, v761, int64(-4611967493404098560))
	mBase = m.M
	v765 = int64(4611123068473966592)
	v769 = int32(-1)
	v773 = v752 & int64(9223372036854775807)
	v774 = int64(9223090561878065152)
	if v773 == v774 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v860 = v837
	v861 = v845
	v862 = v844
	v865 = v839
	goto L126
L130:
	;
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v24)+416))
	v826 = base.B2i32(int32(-1) < v821)
	if int32(-1) < v821 {
		goto L152
	} else {
		goto L153
	}
L131:
	;
	v821 = v817
	goto L130
L132:
	;
	v778 = base.B2i32(v751 != v761)
	goto L134
L133:
	;
	v778 = base.B2i32(base.Ui64(v774) < base.Ui64(v773))
	goto L134
L134:
	;
	if v778 != 0 {
		v817 = v769
		goto L131
	} else {
		goto L135
	}
L135:
	;
	goto L137
L137:
	;
	goto L138
L138:
	;
	goto L139
L139:
	;
	if base.B2i32(v761|v751|(int64(4611123068473966592)|v773) == int64(0)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v765&v752 < int64(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v821 = int32(0)
	goto L130
L142:
	;
	if v752 == v765 {
		goto L148
	} else {
		goto L149
	}
L143:
	;
	if v752 == v765 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v802 = base.B2i32(base.Ui64(v751) < base.Ui64(v761))
	goto L146
L145:
	;
	v802 = base.B2i32(v752 < v765)
	goto L146
L146:
	;
	if v802 != 0 {
		v817 = v769
		goto L131
	} else {
		goto L147
	}
L147:
	;
	v821 = base.B2i32(v751^v761|(v752^v765) != int64(0))
	goto L130
L148:
	;
	v811 = base.B2i32(base.Ui64(v761) < base.Ui64(v751))
	goto L150
L149:
	;
	v811 = base.B2i32(v765 < v752)
	goto L150
L150:
	;
	if v811 != 0 {
		v817 = v769
		goto L131
	} else {
		goto L151
	}
L151:
	;
	v817 = base.B2i32(v751^v761|(v752^v765) != int64(0))
	goto L131
L152:
	;
	v827 = v824
	goto L154
L153:
	;
	v827 = v751
	goto L154
L154:
	;
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(424))))
	if int32(-1) < v821 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v833 = v832
	goto L157
L156:
	;
	v833 = v752
	goto L157
L157:
	;
	F___addtf3(m, v24+int32(400), v751, v752, v827, v833)
	mBase = m.M
	v836 = v750 << (uint(int32(1)) % 32)
	v837 = v836 | v826
	v839 = v755 + int64(-1)
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(408))))
	v845 = *(*int64)(unsafe.Add(mBase, uint32(v24)+400))
	if int32(-1) < v836 {
		v750 = v837
		v751 = v845
		v752 = v844
		v755 = v839
		goto L128
	} else {
		goto L158
	}
L158:
	;
	goto L129
L159:
	;
	v1153 = v24 + int32(320)
	v1154 = int32(1)
	v1160 = int64(0)
	v1169 = v862 & int64(9223372036854775807)
	v1170 = int64(9223090561878065152)
	if v1169 == v1170 {
		goto L205
	} else {
		goto L206
	}
L160:
	;
	v938 = v24 + int32(352)
	v939 = float64(1)
	v941 = int32(144) - v880
	if v941 < int32(1024) {
		goto L174
	} else {
		goto L175
	}
L161:
	;
	v877 = v873
	goto L163
L162:
	;
	v877 = v874
	goto L163
L163:
	;
	if v872 < base.I64_extend_i32_u(l2) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v880 = v877
	goto L166
L165:
	;
	v880 = l2
	goto L166
L166:
	;
	if base.Ui32(v880) < base.Ui32(int32(113)) {
		goto L160
	} else {
		goto L167
	}
L167:
	;
	v884 = v24 + int32(384)
	v889 = m.G0
	v891 = v889 - int32(16)
	m.G0 = v891
	if l4 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v933 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(392))))
	v934 = int64(0)
	v935 = *(*int64)(unsafe.Add(mBase, uint32(v24)+384))
	v1148 = v933
	v1149 = v934
	v1150 = v935
	v1151 = v934
	goto L159
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v884))) = v923
	*(*int64)(unsafe.Add(mBase, uint32(v884)+8)) = v924
	m.G0 = v891 + int32(16)
	goto L168
L170:
	;
	v896 = l4 >> (uint(int32(31)) % 32)
	v898 = l4 ^ v896 - v896
	v901 = base.I32_clz(v898)
	F___ashlti3(m, v891, base.I64_extend_i32_u(v898), int64(0), v901+int32(81))
	mBase = m.M
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v891+int32(8))))
	v922 = *(*int64)(unsafe.Add(mBase, uint32(v891)))
	v923 = v922
	v924 = v907 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v901)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l4&int32(-2147483648))<<(uint(int64(32))%64)
	goto L169
L171:
	;
	v893 = int64(0)
	v923 = v893
	v924 = v893
	goto L169
L172:
	;
	v989 = m.G0
	v991 = v989 - int32(16)
	m.G0 = v991
	v993 = base.I64_reinterpret_f64(base.F64_mul(v974, base.F64_reinterpret_i64(base.I64_extend_i32_u(v975+int32(1023))<<(uint(int64(52))%64))))
	v995 = v993 & int64(4503599627370495)
	v999 = int64(base.Ui64(v993)>>(uint(int64(52))%64)) & int64(2047)
	if v999 == int64(0) {
		goto L189
	} else {
		goto L190
	}
L173:
	;
	goto L172
L174:
	;
	if int32(-1023) < v941 {
		v974 = v939
		v975 = v941
		goto L173
	} else {
		goto L181
	}
L175:
	;
	v945 = base.F64_mul(v939, float64(8.98846567431158e+307))
	if base.Ui32(int32(2047)) <= base.Ui32(v941) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v952 = int32(3069)
	if base.Ui32(v941) < base.Ui32(v952) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v974 = v945
	v975 = v941 + int32(-1023)
	goto L173
L178:
	;
	v955 = v941
	goto L180
L179:
	;
	v955 = v952
	goto L180
L180:
	;
	v974 = base.F64_mul(v945, float64(8.98846567431158e+307))
	v975 = v955 + int32(-2046)
	goto L173
L181:
	;
	v961 = base.F64_mul(v939, float64(2.004168360008973e-292))
	if base.Ui32(v941) <= base.Ui32(int32(-1992)) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v968 = int32(-2960)
	if base.Ui32(v968) < base.Ui32(v941) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v974 = v961
	v975 = v941 + int32(969)
	goto L173
L184:
	;
	v971 = v941
	goto L186
L185:
	;
	v971 = v968
	goto L186
L186:
	;
	v974 = base.F64_mul(v961, float64(2.004168360008973e-292))
	v975 = v971 + int32(1938)
	goto L173
L187:
	;
	v1062 = v24 + int32(336)
	v1067 = m.G0
	v1069 = v1067 - int32(16)
	m.G0 = v1069
	if l4 != 0 {
		goto L200
	} else {
		goto L201
	}
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v938))) = v1046
	*(*int64)(unsafe.Add(mBase, uint32(v938)+8)) = v1047<<(uint(int64(48))%64) | v993&int64(-9223372036854775807-1) | v1048
	m.G0 = v991 + int32(16)
	goto L187
L189:
	;
	if base.B2i32(v995 == int64(0)) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	if v999 == int64(2047) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1046 = v995 << (uint(int64(60)) % 64)
	v1047 = int64(32767)
	v1048 = int64(base.Ui64(v995) >> (uint(int64(4)) % 64))
	goto L188
L192:
	;
	v1046 = v995 << (uint(int64(60)) % 64)
	v1047 = v999 + int64(15360)
	v1048 = int64(base.Ui64(v995) >> (uint(int64(4)) % 64))
	goto L188
L193:
	;
	if base.Ui64(v995) < base.Ui64(int64(4294967296)) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1019 = int64(0)
	v1046 = v1019
	v1047 = v1019
	v1048 = v1019
	goto L188
L195:
	;
	v1033 = base.I32_clz(base.I32_wrap_i64(v993)) | int32(32)
	goto L197
L196:
	;
	v1033 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v995) >> (uint(int64(32)) % 64))))
	goto L197
L197:
	;
	F___ashlti3(m, v991, v995, int64(0), v1033+int32(49))
	mBase = m.M
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(v991+int32(8))))
	v1045 = *(*int64)(unsafe.Add(mBase, uint32(v991)))
	v1046 = v1045
	v1047 = base.I64_extend_i32_u(int32(15372) - v1033)
	v1048 = v1042 ^ int64(281474976710656)
	goto L188
L198:
	;
	v1110 = v24 + int32(368)
	v1111 = *(*int64)(unsafe.Add(mBase, uint32(v24)+352))
	v1116 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(360))))
	v1117 = *(*int64)(unsafe.Add(mBase, uint32(v24)+336))
	v1122 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(344))))
	*(*int64)(unsafe.Add(mBase, uint32(v1110))) = v1111
	v1124 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v1110)+8)) = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v1122)>>(uint(v1124)%64)))&int32(32768)|base.I32_wrap_i64(int64(base.Ui64(v1116&int64(9223090561878065152))>>(uint(v1124)%64))))<<(uint(v1124)%64) | v1116&int64(281474976710655)
	goto L202
L199:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1062))) = v1101
	*(*int64)(unsafe.Add(mBase, uint32(v1062)+8)) = v1102
	m.G0 = v1069 + int32(16)
	goto L198
L200:
	;
	v1074 = l4 >> (uint(int32(31)) % 32)
	v1076 = l4 ^ v1074 - v1074
	v1079 = base.I32_clz(v1076)
	F___ashlti3(m, v1069, base.I64_extend_i32_u(v1076), int64(0), v1079+int32(81))
	mBase = m.M
	v1085 = *(*int64)(unsafe.Add(mBase, uint32(v1069+int32(8))))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1069)))
	v1101 = v1100
	v1102 = v1085 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1079)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l4&int32(-2147483648))<<(uint(int64(32))%64)
	goto L199
L201:
	;
	v1071 = int64(0)
	v1101 = v1071
	v1102 = v1071
	goto L199
L202:
	;
	v1146 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(376))))
	v1147 = *(*int64)(unsafe.Add(mBase, uint32(v24)+368))
	v1148 = v1122
	v1149 = v1147
	v1150 = v1117
	v1151 = v1146
	goto L159
L203:
	;
	v1227 = base.B2i32(v860&v1154 == int32(0)) & (base.B2i32(base.Ui32(v880) < base.Ui32(int32(32))) & base.B2i32(v1223 != int32(0)))
	v1228 = v860 | v1227
	v1232 = m.G0
	v1234 = v1232 - int32(16)
	m.G0 = v1234
	if v1228 != 0 {
		goto L229
	} else {
		goto L230
	}
L204:
	;
	v1223 = v1219
	goto L203
L205:
	;
	v1174 = base.B2i32(v861 != v1160)
	goto L207
L206:
	;
	v1174 = base.B2i32(base.Ui64(v1170) < base.Ui64(v1169))
	goto L207
L207:
	;
	if v1174 != 0 {
		v1219 = v1154
		goto L204
	} else {
		goto L208
	}
L208:
	;
	goto L210
L210:
	;
	goto L211
L211:
	;
	goto L212
L212:
	;
	if base.B2i32(v1160|v861|(int64(0)|v1169) == int64(0)) == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if v1160&v862 < int64(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v1223 = int32(0)
	goto L203
L215:
	;
	if v862 == v1160 {
		goto L223
	} else {
		goto L224
	}
L216:
	;
	if v862 == v1160 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1223 = base.B2i32(v861^v1160|(v862^v1160) != int64(0))
	goto L203
L218:
	;
	v1198 = base.B2i32(base.Ui64(v861) < base.Ui64(v1160))
	goto L220
L219:
	;
	v1198 = base.B2i32(v862 < v1160)
	goto L220
L220:
	;
	if v1198 == int32(0) {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	v1223 = int32(-1)
	goto L203
L222:
	;
	v1219 = base.B2i32(v861^v1160|(v862^v1160) != int64(0))
	goto L204
L223:
	;
	v1210 = base.B2i32(base.Ui64(v1160) < base.Ui64(v861))
	goto L225
L224:
	;
	v1210 = base.B2i32(v1160 < v862)
	goto L225
L225:
	;
	if v1210 == int32(0) {
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v1223 = int32(-1)
	goto L203
L227:
	;
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v24)+320))
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(328))))
	F___multf3(m, v24+int32(304), v1150, v1148, v1268, v1273)
	mBase = m.M
	v1277 = *(*int64)(unsafe.Add(mBase, uint32(v24)+304))
	v1282 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(312))))
	F___addtf3(m, v24+int32(272), v1277, v1282, v1149, v1151)
	mBase = m.M
	if v1227 != 0 {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1153))) = v1259
	*(*int64)(unsafe.Add(mBase, uint32(v1153)+8)) = v1260
	m.G0 = v1234 + int32(16)
	goto L227
L229:
	;
	v1241 = base.I32_clz(v1228)
	F___ashlti3(m, v1234, base.I64_extend_i32_u(v1228), int64(0), int32(112)-(v1241^int32(31)))
	mBase = m.M
	v1248 = *(*int64)(unsafe.Add(mBase, uint32(v1234+int32(8))))
	v1257 = *(*int64)(unsafe.Add(mBase, uint32(v1234)))
	v1259 = v1257
	v1260 = v1248 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1241)<<(uint(int64(48))%64)
	goto L228
L230:
	;
	v1236 = int64(0)
	v1259 = v1236
	v1260 = v1236
	goto L228
L231:
	;
	v1287 = int64(0)
	goto L233
L232:
	;
	v1287 = v861
	goto L233
L233:
	;
	if v1227 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1289 = int64(0)
	goto L236
L235:
	;
	v1289 = v862
	goto L236
L236:
	;
	F___multf3(m, v24+int32(288), v1150, v1148, v1287, v1289)
	mBase = m.M
	v1293 = *(*int64)(unsafe.Add(mBase, uint32(v24)+288))
	v1298 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(296))))
	v1299 = *(*int64)(unsafe.Add(mBase, uint32(v24)+272))
	v1304 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(280))))
	F___addtf3(m, v24+int32(256), v1293, v1298, v1299, v1304)
	mBase = m.M
	v1307 = v24 + int32(240)
	v1308 = *(*int64)(unsafe.Add(mBase, uint32(v24)+256))
	v1313 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(264))))
	v1315 = m.G0
	v1316 = int32(16)
	v1317 = v1315 - v1316
	m.G0 = v1317
	F___addtf3(m, v1317, v1308, v1313, v1149, v1151^int64(-9223372036854775807-1))
	mBase = m.M
	v1322 = *(*int64)(unsafe.Add(mBase, uint32(v1317)))
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v1317+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v1307)+8)) = v1325
	*(*int64)(unsafe.Add(mBase, uint32(v1307))) = v1322
	m.G0 = v1317 + v1316
	goto L237
L237:
	;
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(v24)+240))
	v1336 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(248))))
	v1337 = int64(0)
	v1346 = v1336 & int64(9223372036854775807)
	v1347 = int64(9223090561878065152)
	if v1346 == v1347 {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	v1405 = v24 + int32(224)
	v1406 = base.I32_wrap_i64(v865)
	v1408 = m.G0
	v1410 = v1408 - int32(80)
	m.G0 = v1410
	if v1406 < int32(16384) {
		goto L267
	} else {
		goto L268
	}
L239:
	;
	if v1400 != 0 {
		goto L238
	} else {
		goto L263
	}
L240:
	;
	v1400 = v1396
	goto L239
L241:
	;
	v1351 = base.B2i32(v1331 != v1337)
	goto L243
L242:
	;
	v1351 = base.B2i32(base.Ui64(v1347) < base.Ui64(v1346))
	goto L243
L243:
	;
	if v1351 != 0 {
		v1396 = int32(1)
		goto L240
	} else {
		goto L244
	}
L244:
	;
	goto L246
L246:
	;
	goto L247
L247:
	;
	goto L248
L248:
	;
	if base.B2i32(v1337|v1331|(int64(0)|v1346) == int64(0)) == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	if v1337&v1336 < int64(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1400 = int32(0)
	goto L239
L251:
	;
	if v1336 == v1337 {
		goto L259
	} else {
		goto L260
	}
L252:
	;
	if v1336 == v1337 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1400 = base.B2i32(v1331^v1337|(v1336^v1337) != int64(0))
	goto L239
L254:
	;
	v1375 = base.B2i32(base.Ui64(v1331) < base.Ui64(v1337))
	goto L256
L255:
	;
	v1375 = base.B2i32(v1336 < v1337)
	goto L256
L256:
	;
	if v1375 == int32(0) {
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v1400 = int32(-1)
	goto L239
L258:
	;
	v1396 = base.B2i32(v1331^v1337|(v1336^v1337) != int64(0))
	goto L240
L259:
	;
	v1387 = base.B2i32(base.Ui64(v1337) < base.Ui64(v1331))
	goto L261
L260:
	;
	v1387 = base.B2i32(v1337 < v1336)
	goto L261
L261:
	;
	if v1387 == int32(0) {
		goto L258
	} else {
		goto L262
	}
L262:
	;
	v1400 = int32(-1)
	goto L239
L263:
	;
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hexfloat[0])) = int32(68)
	goto L238
L265:
	;
	v1503 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(232))))
	v1504 = *(*int64)(unsafe.Add(mBase, uint32(v24)+224))
	v1597 = v1504
	v1601 = v1503
	goto L56
L266:
	;
	F___multf3(m, v1410, v1480, v1481, int64(0), base.I64_extend_i32_u(v1482+int32(16383))<<(uint(int64(48))%64))
	mBase = m.M
	v1492 = *(*int64)(unsafe.Add(mBase, uint32(v1410+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+8)) = v1492
	v1494 = *(*int64)(unsafe.Add(mBase, uint32(v1410)))
	*(*int64)(unsafe.Add(mBase, uint32(v1405))) = v1494
	m.G0 = v1410 + int32(80)
	goto L265
L267:
	;
	if int32(-16383) < v1406 {
		v1480 = v1331
		v1481 = v1336
		v1482 = v1406
		goto L266
	} else {
		goto L274
	}
L268:
	;
	F___multf3(m, v1410+int32(32), v1331, v1336, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1423 = *(*int64)(unsafe.Add(mBase, uint32(v1410+int32(40))))
	v1424 = *(*int64)(unsafe.Add(mBase, uint32(v1410)+32))
	if base.Ui32(int32(32767)) <= base.Ui32(v1406) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	F___multf3(m, v1410+int32(16), v1424, v1423, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1434 = int32(49149)
	if base.Ui32(v1406) < base.Ui32(v1434) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1480 = v1424
	v1481 = v1423
	v1482 = v1406 + int32(-16383)
	goto L266
L271:
	;
	v1437 = v1406
	goto L273
L272:
	;
	v1437 = v1434
	goto L273
L273:
	;
	v1444 = *(*int64)(unsafe.Add(mBase, uint32(v1410+int32(24))))
	v1445 = *(*int64)(unsafe.Add(mBase, uint32(v1410)+16))
	v1480 = v1445
	v1481 = v1444
	v1482 = v1437 + int32(-32766)
	goto L266
L274:
	;
	F___multf3(m, v1410+int32(64), v1331, v1336, int64(0), int64(32088147345014784))
	mBase = m.M
	v1457 = *(*int64)(unsafe.Add(mBase, uint32(v1410+int32(72))))
	v1458 = *(*int64)(unsafe.Add(mBase, uint32(v1410)+64))
	if base.Ui32(v1406) <= base.Ui32(int32(-32652)) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	F___multf3(m, v1410+int32(48), v1458, v1457, int64(0), int64(32088147345014784))
	mBase = m.M
	v1468 = int32(-48920)
	if base.Ui32(v1468) < base.Ui32(v1406) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1480 = v1458
	v1481 = v1457
	v1482 = v1406 + int32(16269)
	goto L266
L277:
	;
	v1471 = v1406
	goto L279
L278:
	;
	v1471 = v1468
	goto L279
L279:
	;
	v1478 = *(*int64)(unsafe.Add(mBase, uint32(v1410+int32(56))))
	v1479 = *(*int64)(unsafe.Add(mBase, uint32(v1410)+48))
	v1480 = v1479
	v1481 = v1478
	v1482 = v1471 + int32(32538)
	goto L266
L280:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hexfloat[0])) = int32(68)
	v1509 = v24 + int32(208)
	v1514 = m.G0
	v1516 = v1514 - int32(16)
	m.G0 = v1516
	if l4 != 0 {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	v1558 = *(*int64)(unsafe.Add(mBase, uint32(v24)+208))
	v1563 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(216))))
	v1564 = int64(0)
	v1565 = int64(281474976710656)
	F___multf3(m, v24+int32(192), v1558, v1563, v1564, v1565)
	mBase = m.M
	v1569 = *(*int64)(unsafe.Add(mBase, uint32(v24)+192))
	v1574 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(200))))
	F___multf3(m, v24+int32(176), v1569, v1574, v1564, v1565)
	mBase = m.M
	v1582 = *(*int64)(unsafe.Add(mBase, uint32(v24+int32(184))))
	v1583 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	v1597 = v1583
	v1601 = v1582
	goto L56
L282:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1509))) = v1548
	*(*int64)(unsafe.Add(mBase, uint32(v1509)+8)) = v1549
	m.G0 = v1516 + int32(16)
	goto L281
L283:
	;
	v1521 = l4 >> (uint(int32(31)) % 32)
	v1523 = l4 ^ v1521 - v1521
	v1526 = base.I32_clz(v1523)
	F___ashlti3(m, v1516, base.I64_extend_i32_u(v1523), int64(0), v1526+int32(81))
	mBase = m.M
	v1532 = *(*int64)(unsafe.Add(mBase, uint32(v1516+int32(8))))
	v1547 = *(*int64)(unsafe.Add(mBase, uint32(v1516)))
	v1548 = v1547
	v1549 = v1532 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1526)<<(uint(int64(48))%64) | base.I64_extend_i32_u(l4&int32(-2147483648))<<(uint(int64(32))%64)
	goto L282
L284:
	;
	v1518 = int64(0)
	v1548 = v1518
	v1549 = v1518
	goto L282
}
func F_hexpireCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_hexpireCommand[0]))
	F_hexpireGenericCommand(m, l0, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_hgetCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_hgetCommand[0])))
	v17 = F_lookupKeyReadOrReply(m, l0, v10, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 == int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			v22 = F_checkType(m, l0, v17, int32(4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					m.G0 = v7 + int32(16)
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					v26 = F_objectGetVal(m, v25)
					mBase = m.M
					v27 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(9223372036854775807)
					v38 = F_hashTypeGetValue(m, v17, v26, v7+int32(12), v7+int32(8), v7, v27)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						if v38 != 0 {
							F_addReplyNull(m, l0)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							if v40 == int32(0) {
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
								F_addReplyBulkLongLong(m, l0, v46)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
								F_addReplyBulkCBuffer(m, l0, v40, v43)
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
				}
			}
		}
	}
}
func F_hgetdelCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v231 int32
	_ = v231
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v18 = F_objectGetVal(m, v17)
	mBase = m.M
	v19 = int32(_a_F_hgetdelCommand_0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v69 = F_getLongLongFromObjectOrReply(m, l0, v65, v12+int32(8), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L18
	}
L3:
	;
	if v54-v56 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v54 = F_tolower(m, v50)
	mBase = m.M
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v56 = F_tolower(m, v55)
	mBase = m.M
	goto L3
L5:
	;
	v24 = v18
	v25 = v19
	v26 = v22
	goto L8
L6:
	;
	v50 = int32(0)
	v51 = v19
	goto L4
L7:
	;
	v50 = v47 & int32(255)
	v51 = v46
	goto L4
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v28 == int32(0) {
		v46 = v25
		v47 = v26
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v46 = v40
	v47 = int32(0)
	goto L7
L10:
	;
	v32 = v26 & int32(255)
	if v32 == v28 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(1)
	v40 = v25 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v41 != 0 {
		v24 = v24 + v39
		v25 = v40
		v26 = v41
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v34 = F_tolower(m, v32)
	mBase = m.M
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = F_tolower(m, v35)
	mBase = m.M
	if v34 == v36 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v46 = v25
	v47 = v38
	goto L7
L14:
	;
	goto L9
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_hgetdelCommand[0]))
	F_addReplyErrorObject(m, l0, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L1
L18:
	;
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	if v71 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = F_lookupKeyWrite(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	F_addReplyError(m, l0, int32(_a_F_hgetdelCommand_1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v71 == base.I64_extend_i32_s(v74+int32(-4)) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L1
L25:
	;
	v88 = F_checkType(m, l0, v85, int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v90 = F_hashTypeHasVolatileFields(m, v85)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	if v85 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_initDeferredReplyBuffer(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L33
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v94&int32(240) != int32(32) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = F_objectGetVal(m, v85)
	mBase = m.M
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+28)))
	v102 = v100 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+28)) = uint16(v102)
	goto L32
L32:
	;
	goto L29
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_addReplyArrayLen(m, l0, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(5) <= v110 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v173&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v119 = v85
	v121 = v109
	v122 = int32(0)
	v123 = int32(4)
	goto L38
L37:
	;
	v171 = v85
	v173 = v109
	v174 = int32(0)
	goto L35
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = v123 << (uint(int32(2)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v127)))
	v130 = F_objectGetVal(m, v129)
	mBase = m.M
	F_addHashFieldToReply(m, l0, v119, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L40
	}
L39:
	;
	v171 = v161
	v173 = v162
	v174 = v163
	goto L35
L40:
	;
	if v119 == int32(0) {
		v158 = v121
		v159 = v122
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v165 = v123 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v165 < v166 {
		v119 = v161
		v121 = v162
		v122 = v163
		v123 = v165
		goto L38
	} else {
		goto L52
	}
L42:
	;
	v161 = int32(0)
	v162 = v158
	v163 = v159
	goto L41
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135+v127)))
	v138 = F_objectGetVal(m, v137)
	mBase = m.M
	v139 = F_hashTypeDelete(m, v119, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	if v139 == int32(0) {
		v161 = v119
		v162 = v121
		v163 = v122
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v144 = v122 + int32(1)
	v145 = F_hashTypeLength(m, v119)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	if v145 != 0 {
		v161 = v119
		v162 = v121
		v163 = v144
		goto L41
	} else {
		goto L47
	}
L47:
	;
	if v90 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v155 = F_dbDelete(m, v152, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L51
	}
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUntrackKeyWithVolatileItems(m, v149, v119)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v158 = int32(1)
	v159 = v144
	goto L42
L52:
	;
	goto L39
L53:
	;
	if v174 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	if v171 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v181&int32(240) != int32(32) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v186 = F_objectGetVal(m, v171)
	mBase = m.M
	F_hashtableResumeAutoShrink(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	F_commitDeferredReplyBuffer(m, l0, int32(1))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L70
	}
L59:
	;
	v192 = v173 & int32(1)
	if v192 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	F_signalModifiedKey(m, l0, v199, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L65
	}
L61:
	;
	v193 = F_hashTypeHasVolatileFields(m, v171)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	if v90 == v193 {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v196, v171)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	goto L60
L65:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a_F_hgetdelCommand_2), v207, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	if v192 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v222 = int32(_a_F_hgetdelCommand_3)
	v224 = *(*int64)(unsafe.Add(mBase, _c_F_hgetdelCommand[1]))
	*(*int64)(unsafe.Add(mBase, _c_F_hgetdelCommand[1])) = v224 + base.I64_extend_i32_s(v174)
	goto L58
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_hgetdelCommand_4), v217, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L1
}
func F_hincrbyCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int64
	_ = v109
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v152 int64
	_ = v152
	var v176 int64
	_ = v176
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int64
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int64
	_ = v282
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v291 int64
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int64
	_ = v303
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v311 int32
	_ = v311
	var v317 int64
	_ = v317
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v339 int32
	_ = v339
	var v343 int64
	_ = v343
	var v346 int64
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	v9 = m.G0
	v11 = v9 - int32(5184)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[0]))) = int64(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v20 = F_getLongLongFromObjectOrReply(m, l0, v16, v11+int32(5168), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(5184)
	return
L2:
	;
	return
L3:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = F_lookupKeyWrite(m, v22, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v28 = F_checkType(m, l0, v25, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v25 != 0 {
		v41 = v25
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = F_objectGetVal(m, v43)
	mBase = m.M
	v53 = F_hashTypeGetValue(m, v41, v44, v11+int32(5164), v11+int32(5160), v11+int32(5176), v11+int32(5152))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L15
	}
L9:
	;
	v30 = F_createHashObject(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbAdd(m, v33, v24, v11+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v38 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v41 = v38
	goto L8
L13:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1])))
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[2])))
	if int64(-1) < v209 {
		goto L49
	} else {
		goto L50
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1]))) = int64(0)
	goto L13
L15:
	;
	if v53 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[3])))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[4])))
	v60 = v11 + int32(5176)
	v61 = int32(0)
	if base.Ui32(v58+int32(-21)) < base.Ui32(int32(-20)) {
		v195 = v61
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v195 != 0 {
		goto L13
	} else {
		goto L45
	}
L19:
	;
	goto L18
L20:
	;
	v73 = int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v58 != v73 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v195 = int32(1)
	goto L19
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1]))) = v176
	goto L21
L23:
	;
	if v74&int32(255) == int32(45) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v78 = v74 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v78&int32(255)) {
		v195 = v61
		goto L19
	} else {
		goto L25
	}
L25:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v176 = base.I64_extend_i32_u(v78) & int64(255)
	goto L22
L27:
	;
	if base.Ui32(int32(8)) < base.Ui32((v97+int32(-49))&int32(255)) {
		v195 = v61
		goto L19
	} else {
		goto L30
	}
L28:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v96 = int32(2)
	v97 = v94
	v98 = v55 + int32(1)
	goto L27
L29:
	;
	v96 = v73
	v97 = v74
	v98 = v55
	goto L27
L30:
	;
	v109 = base.I64_extend_i32_u(v97+int32(-48)) & int64(255)
	if base.Ui32(v58) <= base.Ui32(v96) {
		v152 = v109
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v74&int32(255) != int32(45) {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v115 = v96
	v117 = v109
	v119 = v98
	goto L33
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if base.Ui32((v121+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v195 = v61
		goto L19
	} else {
		goto L35
	}
L34:
	;
	v152 = v142
	goto L31
L35:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v117) {
		v195 = v61
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v131 = v117 * int64(10)
	v136 = base.I64_extend_i32_u(v121+int32(-48)) & int64(255)
	if base.Ui64(v136^int64(-1)) < base.Ui64(v131) {
		v195 = v61
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v140 = int32(1)
	v142 = v131 + v136
	v144 = v115 + v140
	if v144 != v58 {
		v115 = v144
		v117 = v142
		v119 = v119 + v140
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	if v152 < int64(0) {
		v195 = v61
		goto L19
	} else {
		goto L43
	}
L40:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v152) {
		v195 = v61
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	v176 = int64(0) - v152
	goto L22
L43:
	;
	if v60 == int32(0) {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v176 = v152
	goto L22
L45:
	;
	F_addReplyError(m, l0, int32(_a_F_hincrbyCommand_0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L1
L47:
	;
	v227 = v209 + v208
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1]))) = v227
	v229 = F_sdsfromlonglong(m, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L57
	}
L48:
	;
	F_addReplyError(m, l0, int32(_a_F_hincrbyCommand_1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L2
	} else {
		goto L56
	}
L49:
	;
	if v209 < int64(1) {
		goto L47
	} else {
		goto L53
	}
L50:
	;
	if int64(-1) < v208 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if v209 < int64(-9223372036854775807-1)-v208 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	if v208 < int64(1) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	if v209 <= v208^int64(9223372036854775807) {
		goto L47
	} else {
		goto L55
	}
L55:
	;
	goto L48
L56:
	;
	goto L1
L57:
	;
	v231 = F_hashTypeHasVolatileFields(m, v41)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[5]))) = uint8(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	v237 = F_objectGetVal(m, v236)
	mBase = m.M
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[0])))
	v242 = F_hashTypeSet(m, v41, v237, v229, v238, int32(2), v11+int32(5151))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v244 = F_hashTypeHasVolatileFields(m, v41)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L61
	}
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	F_signalModifiedKey(m, l0, v250, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L64
	}
L61:
	;
	if v231 == v244 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v247, v41)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[5]))))
	if v255 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a_F_hincrbyCommand_2), v275, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L68
	}
L66:
	;
	v258 = int32(_a_F_hincrbyCommand_3)
	v260 = *(*int64)(unsafe.Add(mBase, _c_F_hincrbyCommand[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_hincrbyCommand[6])) = v260 + int64(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a_F_hincrbyCommand_4), v267, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v280 = int32(_a_F_hincrbyCommand_3)
	v282 = *(*int64)(unsafe.Add(mBase, _c_F_hincrbyCommand[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_hincrbyCommand[7])) = v282 + int64(1)
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1])))
	F_addReplyLongLong(m, l0, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	if v231 == int32(0) {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[1])))
	v295 = m.G0
	v297 = v295 - int32(16)
	m.G0 = v297
	if base.B2i32(v291 == int64(0)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v339 = v11 + int32(16)
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(8))))
	v348 = F_ld2string(m, v339, int32(5120), v343, v346, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L75
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v330
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v331
	m.G0 = v297 + int32(16)
	goto L71
L73:
	;
	v306 = v291 >> (uint(int64(63)) % 64)
	v308 = v291 ^ v306 - v306
	v311 = base.I32_wrap_i64(base.I64_clz(v308))
	F___ashlti3(m, v297, v308, int64(0), v311+int32(49))
	mBase = m.M
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v297+int32(8))))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	v330 = v329
	v331 = v317 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16446)-v311)<<(uint(int64(48))%64) | v291&int64(-9223372036854775807-1)
	goto L72
L74:
	;
	v303 = int64(0)
	v330 = v303
	v331 = v303
	goto L72
L75:
	;
	v350 = F_createRawStringObject(m, v339, v348)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[0])))
	if v352 != int64(-1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v366 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L82
	}
L78:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyCommand[8]))
	F_rewriteClientCommandArgument(m, l0, int32(0), v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	F_rewriteClientCommandArgument(m, l0, int32(3), v350)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	F_decrRefCount(m, v350)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyCommand[0])))
	v369 = F_createStringObjectFromLongLong(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyCommand[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v366))) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = v375
	F_incrRefCount(m, v375)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v379 = int32(_a_F_hincrbyCommand_5)
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyCommand[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v380
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyCommand[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+16)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyCommand[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = v390
	F_incrRefCount(m, v390)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+28)) = v350
	F_replaceClientCommandVector(m, l0, int32(8), v366)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	goto L1
}
func F_hincrbyfloatCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v165 int64
	_ = v165
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v242 int64
	_ = v242
	var v247 int64
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int64
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int64
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	v9 = m.G0
	v11 = v9 - int32(5232)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[0]))) = int64(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v20 = F_getLongDoubleFromObjectOrReply(m, l0, v16, v11+int32(5200), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(5232)
	return
L2:
	;
	return
L3:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[1])))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[2])))
	v30 = v25 & int64(281474976710655)
	v34 = int32(32767)
	v35 = base.I32_wrap_i64(int64(base.Ui64(v25)>>(uint(int64(48))%64))) & v34
	if v35 == v34 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = F_lookupKeyWrite(m, v85, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L26
	}
L6:
	;
	F_addReplyError(m, l0, int32(_a_F_hincrbyfloatCommand_0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L25
	}
L7:
	;
	if v50 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L8:
	;
	v50 = v48
	goto L7
L9:
	;
	v48 = base.B2i32(v30|v22 == int64(0))
	goto L8
L10:
	;
	if v35 != 0 {
		v48 = int32(4)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v30|v22 == int64(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = int32(2)
	goto L14
L13:
	;
	v44 = int32(3)
	goto L14
L14:
	;
	v50 = v44
	goto L7
L15:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[1])))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[2])))
	v59 = v54 & int64(281474976710655)
	v63 = int32(32767)
	v64 = base.I32_wrap_i64(int64(base.Ui64(v54)>>(uint(int64(48))%64))) & v63
	if v64 == v63 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v79 != int32(1) {
		goto L5
	} else {
		goto L24
	}
L17:
	;
	v79 = v77
	goto L16
L18:
	;
	v77 = base.B2i32(v59|v53 == int64(0))
	goto L17
L19:
	;
	if v64 != 0 {
		v77 = int32(4)
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if v59|v53 == int64(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v73 = int32(2)
	goto L23
L22:
	;
	v73 = int32(3)
	goto L23
L23:
	;
	v79 = v73
	goto L16
L24:
	;
	goto L6
L25:
	;
	goto L1
L26:
	;
	v91 = F_checkType(m, l0, v88, int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v88 != 0 {
		v104 = v88
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v105 = int64(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v109 = F_objectGetVal(m, v108)
	mBase = m.M
	v118 = F_hashTypeGetValue(m, v104, v109, v11+int32(5188), v11+int32(5184), v11+int32(5192), v11+int32(5176))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L35
	}
L30:
	;
	v93 = F_createHashObject(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbAdd(m, v96, v87, v11+int32(48))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	if v101 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v104 = v101
	goto L29
L34:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[1])))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[2])))
	F___addtf3(m, v11+int32(8), v195, v200, v191, v192)
	mBase = m.M
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[3]))) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[4]))) = v208
	v214 = v206 & int64(281474976710655)
	v218 = int32(32767)
	v219 = base.I32_wrap_i64(int64(base.Ui64(v206)>>(uint(int64(48))%64))) & v218
	if v219 == v218 {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	if v118 != 0 {
		v191 = v105
		v192 = v105
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[12])))
	if v120 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v138 = v11 + int32(24)
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[13])))
	v143 = m.G0
	v145 = v143 - int32(16)
	m.G0 = v145
	if base.B2i32(v139 == int64(0)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[14])))
	v126 = F_string2ld(m, v120, v123, v11+int32(5216))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	F_addReplyError(m, l0, int32(_a_F_hincrbyfloatCommand_6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	if v126 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[3])))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[4])))
	v191 = v133
	v192 = v132
	goto L34
L42:
	;
	goto L1
L43:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(32))))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	v191 = v189
	v192 = v188
	goto L34
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v138)+8)) = v179
	m.G0 = v145 + int32(16)
	goto L43
L45:
	;
	v154 = v139 >> (uint(int64(63)) % 64)
	v156 = v139 ^ v154 - v154
	v159 = base.I32_wrap_i64(base.I64_clz(v156))
	F___ashlti3(m, v145, v156, int64(0), v159+int32(49))
	mBase = m.M
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v145+int32(8))))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	v178 = v177
	v179 = v165 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16446)-v159)<<(uint(int64(48))%64) | v139&int64(-9223372036854775807-1)
	goto L44
L46:
	;
	v151 = int64(0)
	v178 = v151
	v179 = v151
	goto L44
L47:
	;
	v274 = v11 + int32(48)
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[4])))
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[3])))
	v283 = F_ld2string(m, v274, int32(5120), v278, v281, int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L68
	}
L48:
	;
	F_addReplyError(m, l0, int32(_a_F_hincrbyfloatCommand_1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L67
	}
L49:
	;
	if v234 == int32(0) {
		goto L48
	} else {
		goto L57
	}
L50:
	;
	v234 = v232
	goto L49
L51:
	;
	v232 = base.B2i32(v214|v208 == int64(0))
	goto L50
L52:
	;
	if v219 != 0 {
		v232 = int32(4)
		goto L50
	} else {
		goto L53
	}
L53:
	;
	if v214|v208 == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v228 = int32(2)
	goto L56
L55:
	;
	v228 = int32(3)
	goto L56
L56:
	;
	v234 = v228
	goto L49
L57:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[4])))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[3])))
	v247 = v242 & int64(281474976710655)
	v251 = int32(32767)
	v252 = base.I32_wrap_i64(int64(base.Ui64(v242)>>(uint(int64(48))%64))) & v251
	if v252 == v251 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if v267 != int32(1) {
		goto L47
	} else {
		goto L66
	}
L59:
	;
	v267 = v265
	goto L58
L60:
	;
	v265 = base.B2i32(v247|v237 == int64(0))
	goto L59
L61:
	;
	if v252 != 0 {
		v265 = int32(4)
		goto L59
	} else {
		goto L62
	}
L62:
	;
	if v247|v237 == int64(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v261 = int32(2)
	goto L65
L64:
	;
	v261 = int32(3)
	goto L65
L65:
	;
	v267 = v261
	goto L58
L66:
	;
	goto L48
L67:
	;
	goto L1
L68:
	;
	v285 = F_sdsnewlen(m, v274, v283)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v287 = F_hashTypeHasVolatileFields(m, v104)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)) = uint8(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v293 = F_objectGetVal(m, v292)
	mBase = m.M
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[0])))
	v298 = F_hashTypeSet(m, v104, v293, v285, v294, int32(2), v11+int32(47))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v300 = F_hashTypeHasVolatileFields(m, v104)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L73
	}
L72:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	F_signalModifiedKey(m, l0, v306, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L2
	} else {
		goto L76
	}
L73:
	;
	if v287 == v300 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_dbUpdateObjectWithVolatileItemsTracking(m, v303, v104)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
	if v311 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a_F_hincrbyfloatCommand_2), v331, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L80
	}
L78:
	;
	v314 = int32(_a_F_hincrbyfloatCommand_3)
	v316 = *(*int64)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[11])) = v316 + int64(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+28))
	F_notifyKeyspaceEvent(m, int32(64), int32(_a_F_hincrbyfloatCommand_5), v323, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v336 = int32(_a_F_hincrbyfloatCommand_3)
	v338 = *(*int64)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[5])) = v338 + int64(1)
	F_addReplyBulkCBuffer(m, l0, v11+int32(48), v283)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v348 = F_createRawStringObject(m, v11+int32(48), v283)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[0])))
	if v350 != int64(-1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v364 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L88
	}
L84:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[10]))
	F_rewriteClientCommandArgument(m, l0, int32(0), v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	F_rewriteClientCommandArgument(m, l0, int32(3), v348)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	F_decrRefCount(m, v348)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L1
L88:
	;
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_hincrbyfloatCommand[0])))
	v367 = F_createStringObjectFromLongLong(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v373
	F_incrRefCount(m, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v377 = int32(_a_F_hincrbyfloatCommand_4)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+12)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = v378
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v382
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_hincrbyfloatCommand[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+20)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+24)) = v388
	F_incrRefCount(m, v388)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+28)) = v348
	F_replaceClientCommandVector(m, l0, int32(8), v364)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	goto L1
}
func F_hkeysCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_genericHgetallCommand(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_hllAdd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int64
	_ = v139
	var v144 int64
	_ = v144
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v252 int64
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int64
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int64
	_ = v367
	var v372 int64
	_ = v372
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v396 int64
	_ = v396
	var v397 int64
	_ = v397
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v411 int64
	_ = v411
	var v414 int64
	_ = v414
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	v9 = F_objectGetVal(m, l0)
	mBase = m.M
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
	switch v10 {
	case 0:
		v24 = base.I64_extend_i32_s(l2)*int64(-4132994306676758123) ^ base.I64_extend_i32_u(int32(-1379386599))
		v26 = l2 & int32(-8)
		if v26 == int32(0) {
			v135 = l1
			v139 = v24
		} else {
			v30 = l2 + int32(-8)
			v31 = int32(24)
			if v30&v31 != v31 {
				v35 = int32(3)
				v43 = l1
				v44 = int32(0)
				v46 = v24
				for {
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
					v52 = int64(-4132994306676758123)
					v53 = v51 * v52
					v61 = ((int64(base.Ui64(v53)>>(uint(int64(47))%64))^v53)*v52 ^ v46) * v52
					v63 = v43 + int32(8)
					v65 = v44 + int32(1)
					if v65 != (int32(base.Ui32(v30)>>(uint(v35)%32))+int32(1))&v35 {
						v43 = v63
						v44 = v65
						v46 = v61
						continue
					} else {
						break
					}
					break
				}
				v68 = v63
				v71 = v61
			} else {
				v68 = l1
				v71 = v24
			}
			v76 = l1 + v26
			if base.Ui32(v30) < base.Ui32(int32(24)) {
				v135 = v76
				v139 = v71
			} else {
				v80 = v68
				v83 = v71
				for {
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v80)+24))
					v89 = int64(-4132994306676758123)
					v90 = v88 * v89
					v91 = int64(47)
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v80)+16))
					v98 = v96 * v89
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v80)+8))
					v106 = v104 * v89
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
					v114 = v112 * v89
					v131 = ((int64(base.Ui64(v90)>>(uint(v91)%64))^v90)*v89 ^ ((int64(base.Ui64(v98)>>(uint(v91)%64))^v98)*v89^((int64(base.Ui64(v106)>>(uint(v91)%64))^v106)*v89^((int64(base.Ui64(v114)>>(uint(v91)%64))^v114)*v89^v83)*v89)*v89)*v89) * v89
					v133 = v80 + int32(32)
					if v133 != v76 {
						v80 = v133
						v83 = v131
						continue
					} else {
						break
					}
					break
				}
				v135 = v76
				v139 = v131
			}
		}
		switch l2 & int32(7) {
		default:
			v178 = v139
		case 1:
			v173 = v139
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 2:
			v168 = v139
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 3:
			v163 = v139
			v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
			v168 = v164<<(uint(int64(16))%64) ^ v163
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 4:
			v158 = v139
			v159 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
			v163 = v159<<(uint(int64(24))%64) ^ v158
			v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
			v168 = v164<<(uint(int64(16))%64) ^ v163
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 5:
			v153 = v139
			v154 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
			v158 = v154<<(uint(int64(32))%64) ^ v153
			v159 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
			v163 = v159<<(uint(int64(24))%64) ^ v158
			v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
			v168 = v164<<(uint(int64(16))%64) ^ v163
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 6:
			v148 = v139
			v149 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+5)))
			v153 = v149<<(uint(int64(40))%64) ^ v148
			v154 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
			v158 = v154<<(uint(int64(32))%64) ^ v153
			v159 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
			v163 = v159<<(uint(int64(24))%64) ^ v158
			v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
			v168 = v164<<(uint(int64(16))%64) ^ v163
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		case 7:
			v144 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+6)))
			v148 = v144<<(uint(int64(48))%64) ^ v139
			v149 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+5)))
			v153 = v149<<(uint(int64(40))%64) ^ v148
			v154 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
			v158 = v154<<(uint(int64(32))%64) ^ v153
			v159 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+3)))
			v163 = v159<<(uint(int64(24))%64) ^ v158
			v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
			v168 = v164<<(uint(int64(16))%64) ^ v163
			v169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
			v173 = v169<<(uint(int64(8))%64) ^ v168
			v174 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
			v178 = (v173 ^ v174) * int64(-4132994306676758123)
		}
		v179 = int64(47)
		v183 = (int64(base.Ui64(v178)>>(uint(v179)%64)) ^ v178) * int64(-4132994306676758123)
		v186 = int64(base.Ui64(v183)>>(uint(v179)%64)) ^ v183
		v190 = int32(6)
		v191 = base.I32_wrap_i64(v186) & int32(16383) * v190
		v194 = v9 + int32(base.Ui32(v191)>>(uint(int32(3))%32))
		v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+int32(17)))))
		v200 = v191 & v190
		v201 = int32(8) - v200
		v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+16)))
		v213 = base.I32_wrap_i64(base.I64_ctz(int64(base.Ui64(v186)>>(uint(int64(14))%64)) | int64(1125899906842624)))
		if base.Ui32((v197<<(uint(v201)%32)|int32(base.Ui32(v203)>>(uint(v200)%32)))&int32(63)) <= base.Ui32(v213) {
			v219 = int32(1)
			v225 = v213 + v219
			v227 = v197&(int32(-64)>>(uint(v201)%32)) | int32(base.Ui32(v225)>>(uint(v201)%32))
			*(*uint8)(unsafe.Add(mBase, uint32(v194+int32(17)))) = uint8(v227)
			v235 = v203&(int32(63)<<(uint(v200)%32)^int32(-1)) | v225<<(uint(v200)%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v194+int32(16)))) = uint8(v235)
			return v219
		} else {
			return int32(0)
		}
	case 1:
		v252 = base.I64_extend_i32_s(l2)*int64(-4132994306676758123) ^ base.I64_extend_i32_u(int32(-1379386599))
		v254 = l2 & int32(-8)
		if v254 == int32(0) {
			v363 = l1
			v367 = v252
		} else {
			v258 = l2 + int32(-8)
			v259 = int32(24)
			if v258&v259 != v259 {
				v263 = int32(3)
				v271 = l1
				v272 = int32(0)
				v274 = v252
				for {
					v279 = *(*int64)(unsafe.Add(mBase, uint32(v271)))
					v280 = int64(-4132994306676758123)
					v281 = v279 * v280
					v289 = ((int64(base.Ui64(v281)>>(uint(int64(47))%64))^v281)*v280 ^ v274) * v280
					v291 = v271 + int32(8)
					v293 = v272 + int32(1)
					if v293 != (int32(base.Ui32(v258)>>(uint(v263)%32))+int32(1))&v263 {
						v271 = v291
						v272 = v293
						v274 = v289
						continue
					} else {
						break
					}
					break
				}
				v296 = v291
				v299 = v289
			} else {
				v296 = l1
				v299 = v252
			}
			v304 = l1 + v254
			if base.Ui32(v258) < base.Ui32(int32(24)) {
				v363 = v304
				v367 = v299
			} else {
				v308 = v296
				v311 = v299
				for {
					v316 = *(*int64)(unsafe.Add(mBase, uint32(v308)+24))
					v317 = int64(-4132994306676758123)
					v318 = v316 * v317
					v319 = int64(47)
					v324 = *(*int64)(unsafe.Add(mBase, uint32(v308)+16))
					v326 = v324 * v317
					v332 = *(*int64)(unsafe.Add(mBase, uint32(v308)+8))
					v334 = v332 * v317
					v340 = *(*int64)(unsafe.Add(mBase, uint32(v308)))
					v342 = v340 * v317
					v359 = ((int64(base.Ui64(v318)>>(uint(v319)%64))^v318)*v317 ^ ((int64(base.Ui64(v326)>>(uint(v319)%64))^v326)*v317^((int64(base.Ui64(v334)>>(uint(v319)%64))^v334)*v317^((int64(base.Ui64(v342)>>(uint(v319)%64))^v342)*v317^v311)*v317)*v317)*v317) * v317
					v361 = v308 + int32(32)
					if v361 != v304 {
						v308 = v361
						v311 = v359
						continue
					} else {
						break
					}
					break
				}
				v363 = v304
				v367 = v359
			}
		}
		switch l2 & int32(7) {
		default:
			v406 = v367
		case 1:
			v401 = v367
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 2:
			v396 = v367
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 3:
			v391 = v367
			v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+2)))
			v396 = v392<<(uint(int64(16))%64) ^ v391
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 4:
			v386 = v367
			v387 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+3)))
			v391 = v387<<(uint(int64(24))%64) ^ v386
			v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+2)))
			v396 = v392<<(uint(int64(16))%64) ^ v391
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 5:
			v381 = v367
			v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+4)))
			v386 = v382<<(uint(int64(32))%64) ^ v381
			v387 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+3)))
			v391 = v387<<(uint(int64(24))%64) ^ v386
			v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+2)))
			v396 = v392<<(uint(int64(16))%64) ^ v391
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 6:
			v376 = v367
			v377 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+5)))
			v381 = v377<<(uint(int64(40))%64) ^ v376
			v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+4)))
			v386 = v382<<(uint(int64(32))%64) ^ v381
			v387 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+3)))
			v391 = v387<<(uint(int64(24))%64) ^ v386
			v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+2)))
			v396 = v392<<(uint(int64(16))%64) ^ v391
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		case 7:
			v372 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+6)))
			v376 = v372<<(uint(int64(48))%64) ^ v367
			v377 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+5)))
			v381 = v377<<(uint(int64(40))%64) ^ v376
			v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+4)))
			v386 = v382<<(uint(int64(32))%64) ^ v381
			v387 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+3)))
			v391 = v387<<(uint(int64(24))%64) ^ v386
			v392 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+2)))
			v396 = v392<<(uint(int64(16))%64) ^ v391
			v397 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
			v401 = v397<<(uint(int64(8))%64) ^ v396
			v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
			v406 = (v401 ^ v402) * int64(-4132994306676758123)
		}
		v407 = int64(47)
		v411 = (int64(base.Ui64(v406)>>(uint(v407)%64)) ^ v406) * int64(-4132994306676758123)
		v414 = int64(base.Ui64(v411)>>(uint(v407)%64)) ^ v411
		v426 = F_hllSparseSet(m, l0, base.I32_wrap_i64(v414)&int32(16383), base.I32_wrap_i64(base.I64_ctz(int64(base.Ui64(v414)>>(uint(int64(14))%64))|int64(1125899906842624)))+int32(1))
		mBase = m.M
		v429 = m.ExcPending
		if v429 != 0 {
			return int32(0)
		} else {
			v430 = v426
			return v430
		}
	default:
		v430 = int32(-1)
		return v430
	}
}
func F_hllCount(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v426 int64
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v557 float64
	_ = v557
	var v559 int32
	_ = v559
	var v563 float64
	_ = v563
	var v568 float64
	_ = v568
	var v583 float64
	_ = v583
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v589 float64
	_ = v589
	var v590 float64
	_ = v590
	var v593 float64
	_ = v593
	var v595 float64
	_ = v595
	var v613 float64
	_ = v613
	var v622 int32
	_ = v622
	var v631 float64
	_ = v631
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 float64
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v657 float64
	_ = v657
	var v662 float64
	_ = v662
	var v664 int32
	_ = v664
	var v667 float64
	_ = v667
	var v683 float64
	_ = v683
	var v684 float64
	_ = v684
	var v685 float64
	_ = v685
	var v687 float64
	_ = v687
	var v689 float64
	_ = v689
	var v708 float64
	_ = v708
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int64
	_ = v724
	var v726 int64
	_ = v726
	var v730 int64
	_ = v730
	var v750 int64
	_ = v750
	var v764 int32
	_ = v764
	var v773 int64
	_ = v773
	var v776 int64
	_ = v776
	var v777 int64
	_ = v777
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v792 int64
	_ = v792
	var v793 int32
	_ = v793
	var v795 int64
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int64
	_ = v802
	var v805 int64
	_ = v805
	var v806 int64
	_ = v806
	var v810 float64
	_ = v810
	var v818 int64
	_ = v818
	var v820 int64
	_ = v820
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	v27 = F__emscripten_memset_bulkmem(m, v20+int32(16), base.I32_extend8_s(int32(0)), int32(256))
	mBase = m.M
	goto L1
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	switch v28 {
	case 0:
		goto L4
	case 1:
		goto L6
	default:
		goto L5
	}
L2:
	;
	v557 = float64(0)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v20)+220))
	v563 = base.F64_mul(base.F64_sub(float64(16384), base.F64_convert_i32_s(v559)), float64(6.103515625e-05))
	if base.F64_eq(v563, v557) != 0 {
		v613 = v557
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v412 = l0 + int32(16)
	v413 = int32(0)
	goto L36
L4:
	;
	v177 = int32(16)
	v180 = v20 + v177
	v196 = l0 + v177
	v198 = int32(0)
	goto L33
L5:
	;
	if v28 == int32(255) {
		goto L3
	} else {
		goto L29
	}
L6:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v32 & int32(7) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L11
	case 3:
		goto L10
	case 4:
		goto L9
	default:
		v146 = int32(1)
		goto L7
	}
L7:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L27
	}
L8:
	;
	v50 = int32(1)
	if v49+int32(-16) < v50 {
		v146 = v50
		goto L7
	} else {
		goto L14
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v49 = v48
	goto L8
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v49 = v45
	goto L8
L11:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v49 = v42
	goto L8
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v49 = v39
	goto L8
L13:
	;
	v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
	goto L8
L14:
	;
	v55 = int32(16)
	v62 = int32(0)
	v67 = l0 + v55
	v68 = v62
	v71 = v62
	goto L16
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v135
	v146 = base.B2i32(v134 != int32(16384)) | v136
	goto L7
L16:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v83 = v81 & int32(192)
	if v83 == int32(64) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v134 = v124
	v135 = v125
	v136 = int32(0)
	goto L15
L18:
	;
	v129 = v67 + v123
	if base.Ui32(v129) < base.Ui32(l0+v49) {
		v67 = v129
		v68 = v124
		v71 = v125
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v108 = int32(1)
	v112 = v81&int32(3) + v108
	v113 = v112 + v68
	if int32(16384) < v113 {
		v134 = v68
		v135 = v71
		v136 = v108
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v94 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v102 = v81<<(uint(int32(8))%32)&int32(16128) | v99 + v94
	v103 = v102 + v68
	if int32(16384) < v103 {
		v134 = v68
		v135 = v71
		v136 = v94
		goto L15
	} else {
		goto L24
	}
L21:
	;
	if v83 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v86 = int32(1)
	v88 = v81 + v86
	v89 = v88 + v68
	if int32(16384) < v89 {
		v134 = v68
		v135 = v71
		v136 = v86
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v123 = int32(1)
	v124 = v89
	v125 = v71 + v88
	goto L18
L24:
	;
	v123 = int32(2)
	v124 = v103
	v125 = v71 + v102
	goto L18
L25:
	;
	v118 = v20 + v55 | int32(4) + v81&int32(124)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v119 + v112
	v123 = int32(1)
	v124 = v113
	v125 = v71
	goto L18
L26:
	;
	goto L17
L27:
	;
	if v146 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	goto L2
L29:
	;
	F__serverPanic_1(m, int32(_a_F_hllCount_0), int32(1101), int32(_a_F_hllCount_1), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int64(0)
L31:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L2
L33:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+11)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+10)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+9)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+8)))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+7)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+6)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+5)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+3)))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+2)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v224 = int32(63)
	v226 = int32(2)
	v228 = v180 + v223&v224<<(uint(v226)%32)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v230 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v229 + v230
	v235 = int32(60)
	v237 = int32(6)
	v242 = v180 + (v222<<(uint(v226)%32)&v235|int32(base.Ui32(v223)>>(uint(v237)%32)))<<(uint(v226)%32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v243 + v230
	v247 = int32(4)
	v249 = int32(48)
	v256 = v180 + (v221<<(uint(v247)%32)&v249|int32(base.Ui32(v222)>>(uint(v247)%32)))<<(uint(v226)%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v257 + v230
	v261 = int32(252)
	v263 = v180 + v221&v261
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v264 + v230
	v272 = v180 + v220&v224<<(uint(v226)%32)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v273 + v230
	v286 = v180 + (v219<<(uint(v226)%32)&v235|int32(base.Ui32(v220)>>(uint(v237)%32)))<<(uint(v226)%32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = v287 + v230
	v300 = v180 + (v218<<(uint(v247)%32)&v249|int32(base.Ui32(v219)>>(uint(v247)%32)))<<(uint(v226)%32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = v301 + v230
	v307 = v180 + v218&v261
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = v308 + v230
	v316 = v180 + v217&v224<<(uint(v226)%32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v317 + v230
	v330 = v180 + (v216<<(uint(v226)%32)&v235|int32(base.Ui32(v217)>>(uint(v237)%32)))<<(uint(v226)%32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v331 + v230
	v344 = v180 + (v215<<(uint(v247)%32)&v249|int32(base.Ui32(v216)>>(uint(v247)%32)))<<(uint(v226)%32)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v345 + v230
	v351 = v180 + v215&v261
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v352 + v230
	v360 = v180 + v214&v224<<(uint(v226)%32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v361 + v230
	v374 = v180 + (v213<<(uint(v226)%32)&v235|int32(base.Ui32(v214)>>(uint(v237)%32)))<<(uint(v226)%32)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	*(*int32)(unsafe.Add(mBase, uint32(v374))) = v375 + v230
	v388 = v180 + (v212<<(uint(v247)%32)&v249|int32(base.Ui32(v213)>>(uint(v247)%32)))<<(uint(v226)%32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v389 + v230
	v395 = v180 + v212&v261
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v396 + v230
	v403 = v198 + v230
	if v403 != int32(1024) {
		v196 = v196 + int32(12)
		v198 = v403
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	goto L34
L36:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v412)))
	if v426 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L2
L38:
	;
	v537 = v413 + int32(1)
	if v537 != int32(2048) {
		v412 = v412 + int32(8)
		v413 = v537
		goto L36
	} else {
		goto L41
	}
L39:
	;
	v434 = v20 + int32(16)
	v435 = base.I32_wrap_i64(v426)
	v436 = int32(255)
	v438 = int32(2)
	v440 = v434 + v435&v436<<(uint(v438)%32)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v442 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v441 + v442
	v449 = int32(1020)
	v451 = v434 + int32(base.Ui32(v435)>>(uint(int32(6))%32))&v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v452 + v442
	v462 = v434 + int32(base.Ui32(v435)>>(uint(int32(14))%32))&v449
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v463 + v442
	v473 = v434 + int32(base.Ui32(v435)>>(uint(int32(22))%32))&v449
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v474 + v442
	v487 = v434 + base.I32_wrap_i64(int64(base.Ui64(v426)>>(uint(int64(32))%64)))&v436<<(uint(v438)%32)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v488 + v442
	v501 = v434 + base.I32_wrap_i64(int64(base.Ui64(v426)>>(uint(int64(40))%64)))&v436<<(uint(v438)%32)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v502 + v442
	v515 = v434 + base.I32_wrap_i64(int64(base.Ui64(v426)>>(uint(int64(48))%64)))&v436<<(uint(v438)%32)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v516 + v442
	v527 = v434 + base.I32_wrap_i64(int64(base.Ui64(v426)>>(uint(int64(56))%64)))<<(uint(v438)%32)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v528 + v442
	goto L38
L40:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v429 + int32(8)
	goto L38
L41:
	;
	goto L37
L42:
	;
	v622 = int32(50)
	v631 = v613
	goto L48
L43:
	;
	if base.F64_eq(v563, float64(1)) != 0 {
		v613 = v557
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v568 = float64(1)
	v583 = base.F64_sub(v568, v563)
	v584 = v563
	v585 = v568
	goto L45
L45:
	;
	v589 = base.F64_sqrt(v584)
	v590 = base.F64_sub(float64(1), v589)
	v593 = base.F64_mul(v585, float64(0.5))
	v595 = base.F64_sub(v583, base.F64_mul(base.F64_mul(v590, v590), v593))
	if base.F64_ne(v583, v595) != 0 {
		v583 = v595
		v584 = v589
		v585 = v593
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v613 = base.F64_mul(base.F64_div(v595, float64(3)), float64(16384))
	goto L42
L47:
	;
	goto L46
L48:
	;
	v637 = v20 + int32(16)
	v638 = int32(2)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v637+v622<<(uint(v638)%32))))
	v644 = float64(0.5)
	v649 = v622 + int32(-1)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v637+v649<<(uint(v638)%32))))
	v657 = base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(v631, base.F64_convert_i32_s(v641)), v644), base.F64_convert_i32_s(v653)), v644)
	if v649 != int32(1) {
		v622 = v622 + int32(-2)
		v631 = v657
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v662 = float64(1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v667 = base.F64_mul(base.F64_convert_i32_s(v664), float64(6.103515625e-05))
	if base.F64_eq(v667, v662) != 0 {
		v708 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v720 = m.G0
	v722 = v720 - int32(16)
	m.G0 = v722
	v724 = base.I64_reinterpret_f64(base.F64_div(float64(1.9363525058498377e+08), base.F64_add(base.F64_mul(v708, float64(16384)), v657)))
	v726 = v724 & int64(4503599627370495)
	v730 = int64(base.Ui64(v724)>>(uint(int64(52))%64)) & int64(2047)
	if v730 == int64(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v683 = v662
	v684 = v667
	v685 = v667
	goto L53
L53:
	;
	v687 = base.F64_mul(v684, v684)
	v689 = base.F64_add(base.F64_mul(v687, v683), v685)
	if base.F64_ne(v685, v689) != 0 {
		v683 = base.F64_add(v683, v683)
		v684 = v687
		v685 = v689
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v708 = v689
	goto L51
L55:
	;
	goto L54
L56:
	;
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v793 = int32(8)
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v20+v793)))
	v797 = m.G0
	v798 = int32(16)
	v799 = v797 - v798
	m.G0 = v799
	F_roundl(m, v799, v792, v795)
	mBase = m.M
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v799)))
	v805 = *(*int64)(unsafe.Add(mBase, uint32(v799+v793)))
	v806 = F___fixtfdi(m, v802, v805)
	mBase = m.M
	m.G0 = v799 + v798
	goto L69
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v777
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v778<<(uint(int64(48))%64) | v724&int64(-9223372036854775807-1) | v779
	m.G0 = v722 + int32(16)
	goto L56
L58:
	;
	if base.B2i32(v726 == int64(0)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	if v730 == int64(2047) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v777 = v726 << (uint(int64(60)) % 64)
	v778 = int64(32767)
	v779 = int64(base.Ui64(v726) >> (uint(int64(4)) % 64))
	goto L57
L61:
	;
	v777 = v726 << (uint(int64(60)) % 64)
	v778 = v730 + int64(15360)
	v779 = int64(base.Ui64(v726) >> (uint(int64(4)) % 64))
	goto L57
L62:
	;
	if base.Ui64(v726) < base.Ui64(int64(4294967296)) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v750 = int64(0)
	v777 = v750
	v778 = v750
	v779 = v750
	goto L57
L64:
	;
	v764 = base.I32_clz(base.I32_wrap_i64(v724)) | int32(32)
	goto L66
L65:
	;
	v764 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v726) >> (uint(int64(32)) % 64))))
	goto L66
L66:
	;
	F___ashlti3(m, v722, v726, int64(0), v764+int32(49))
	mBase = m.M
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v722+int32(8))))
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v722)))
	v777 = v776
	v778 = base.I64_extend_i32_u(int32(15372) - v764)
	v779 = v773 ^ int64(281474976710656)
	goto L57
L67:
	;
	m.G0 = v20 + int32(272)
	return v820
L68:
	;
	v820 = int64(0)
	goto L67
L69:
	;
	v810 = base.F64_convert_i64_s(v806)
	if base.F64_lt(v810, float64(1.8446744073709552e+19))&base.F64_ge(v810, float64(0)) == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v818 = base.I64_trunc_f64_u(v810)
	v820 = v818
	goto L67
}
func F_hllSparseSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var __phi166 int32
	_ = __phi166
	var v168 int32
	_ = v168
	var __phi168 int32
	_ = __phi168
	var v170 int32
	_ = v170
	var __phi170 int32
	_ = __phi170
	var v176 int32
	_ = v176
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
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
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v910 int32
	_ = v910
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v942 int32
	_ = v942
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if base.Ui32(int32(32)) < base.Ui32(l2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_hllSparseSet_0), int32(_a_F_hllSparseSet_1), int32(949))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L32
	} else {
		goto L230
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v924
L3:
	;
	v877 = F_hllSparseToDense(m, l0)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L32
	} else {
		goto L227
	}
L4:
	;
	v22 = F_objectGetVal(m, l0)
	mBase = m.M
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-1)))))
	switch v25 & int32(7) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v42 = int32(0)
		goto L5
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_hllSparseSet[0]))
	if base.Ui32(v44) <= base.Ui32(v42) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-9))))
	v42 = v41
	goto L5
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(-5))))
	v42 = v38
	goto L5
L8:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(-3)))))
	v42 = v35
	goto L5
L9:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(-2)))))
	v42 = v32
	goto L5
L10:
	;
	v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
	goto L5
L11:
	;
	v131 = int32(-1)
	v132 = F_objectGetVal(m, l0)
	mBase = m.M
	v133 = F_objectGetVal(m, l0)
	mBase = m.M
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v131))))
	switch v136 & int32(7) {
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
		v924 = v131
		goto L2
	}
L12:
	;
	v46 = F_objectGetVal(m, l0)
	mBase = m.M
	v47 = int32(-1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))))
	switch v49&int32(7) + v47 {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	default:
		goto L13
	}
L13:
	;
	v88 = F_objectGetVal(m, l0)
	mBase = m.M
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(-1)))))
	switch v91 & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	default:
		v108 = int32(0)
		goto L20
	}
L14:
	;
	if base.Ui32(int32(2)) < base.Ui32(v83) {
		goto L11
	} else {
		goto L19
	}
L15:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v46+int32(-9))))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v46+int32(-17))))
	v83 = base.I32_wrap_i64(v77 - v80)
	goto L14
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-5))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-9))))
	v83 = v70 - v73
	goto L14
L17:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+int32(-3)))))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+int32(-5)))))
	v83 = v63 - v66
	goto L14
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-2)))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-3)))))
	v83 = v56 - v59
	goto L14
L19:
	;
	goto L13
L20:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_hllSparseSet[0]))
	v111 = F_objectGetVal(m, l0)
	mBase = m.M
	v113 = v108 + int32(3)
	v114 = int32(300)
	if base.Ui32(v113) < base.Ui32(v114) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(-17))))
	v108 = v107
	goto L20
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(-9))))
	v108 = v104
	goto L20
L23:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+int32(-5)))))
	v108 = v101
	goto L20
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+int32(-3)))))
	v108 = v98
	goto L20
L25:
	;
	v108 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
	goto L20
L26:
	;
	v117 = v113
	goto L28
L27:
	;
	v117 = v114
	goto L28
L28:
	;
	v118 = v117 + v113
	if base.Ui32(v118) < base.Ui32(v110) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v120 = v118
	goto L31
L30:
	;
	v120 = v110
	goto L31
L31:
	;
	v122 = F_sdsResize(m, v111, v120, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	F_objectSetVal(m, l0, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	goto L11
L35:
	;
	if v153 < int32(17) {
		v924 = v131
		goto L2
	} else {
		goto L41
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(-17))))
	v153 = v152
	goto L35
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(-9))))
	v153 = v149
	goto L35
L38:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133+int32(-5)))))
	v153 = v146
	goto L35
L39:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+int32(-3)))))
	v153 = v143
	goto L35
L40:
	;
	v153 = int32(base.Ui32(v136) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	v156 = v132 + v153
	v157 = int32(0)
	v159 = v132 + int32(16)
	__phi166 = v159
	__phi168 = v157
	__phi170 = v157
	v166 = __phi166
	v168 = __phi168
	v170 = __phi170
	goto L42
L42:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if base.Ui32(v176) < base.Ui32(int32(64)) {
		v192 = v176
		v193 = int32(1)
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if base.Ui32(v200) < base.Ui32(v156) {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v194 = v192 + v168
	if v194 < l1 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if int32(-1) < base.I32_extend8_s(v176) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v192 = v176<<(uint(int32(8))%32)&int32(16128) | v189
	v193 = int32(2)
	goto L44
L47:
	;
	v192 = v176 & int32(3)
	v193 = int32(1)
	goto L44
L48:
	;
	goto L43
L49:
	;
	v197 = v194 + int32(1)
	v198 = v166 + v193
	if base.Ui32(v198) < base.Ui32(v156) {
		__phi166 = v198
		__phi168 = v197
		__phi170 = v166
		v166 = __phi166
		v168 = __phi168
		v170 = __phi170
		goto L42
	} else {
		goto L51
	}
L50:
	;
	v200 = v166
	v201 = v168
	v202 = v170
	goto L48
L51:
	;
	v200 = v198
	v201 = v197
	v202 = v166
	goto L48
L52:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v207 = v205 & int32(192)
	v209 = base.B2i32(v207 == int32(64))
	if v207 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	v924 = int32(-1)
	goto L2
L54:
	;
	F__serverAssert(m, int32(_a_F_hllSparseSet_2), int32(_a_F_hllSparseSet_1), int32(888))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L32
	} else {
		goto L226
	}
L55:
	;
	if v202 != 0 {
		goto L160
	} else {
		goto L161
	}
L56:
	;
	v350 = v349 + v348
	if v350 < int32(1) {
		goto L83
	} else {
		goto L84
	}
L57:
	;
	v348 = v343
	v349 = int32(-2)
	goto L56
L58:
	;
	v332 = int32(-1)
	v334 = v267 + v332
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)) = uint8(v334)
	v340 = v258 - (v17 + int32(11)) + int32(2)
	if v207 == int32(0) {
		v348 = v340
		v349 = v332
		goto L56
	} else {
		goto L82
	}
L59:
	;
	v330 = v258 + int32(1) - (v17 + int32(11))
	if v207 != 0 {
		v343 = v330
		goto L57
	} else {
		goto L81
	}
L60:
	;
	v284 = v201 + v192
	if l1 == v201 {
		v299 = v17 + int32(11)
		goto L76
	} else {
		goto L77
	}
L61:
	;
	v236 = v201 + v192
	if l1 == v201 {
		v258 = v17 + int32(11)
		goto L69
	} else {
		goto L70
	}
L62:
	;
	if v205 != 0 {
		goto L61
	} else {
		goto L68
	}
L63:
	;
	if v207 == int32(64) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v215 = int32(base.Ui32(v205)>>(uint(int32(2))%32)) & int32(31)
	if base.Ui32(v215+int32(1)) < base.Ui32(l2) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v205&int32(3) != 0 {
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v924 = int32(0)
	goto L2
L67:
	;
	v227 = l2<<(uint(int32(2))%32) + int32(124) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v227)
	v604 = v156
	goto L55
L68:
	;
	v234 = l2<<(uint(int32(2))%32) + int32(124) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v234)
	v604 = v156
	goto L55
L69:
	;
	v264 = l2<<(uint(int32(2))%32) + int32(124) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v264)
	if v236 == l1 {
		goto L59
	} else {
		goto L73
	}
L70:
	;
	v240 = l1 - v201
	if v240 < int32(65) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v254 = v240 + int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)) = uint8(v254)
	v258 = v17 + int32(12)
	goto L69
L72:
	;
	v244 = v240 + int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v244)
	v249 = int32(base.Ui32(v244)>>(uint(int32(8))%32)) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)) = uint8(v249)
	v258 = v17 + int32(13)
	goto L69
L73:
	;
	v267 = v236 - l1
	if v267 < int32(65) {
		goto L58
	} else {
		goto L74
	}
L74:
	;
	v270 = int32(-1)
	v272 = v267 + v270
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+2)) = uint8(v272)
	v277 = int32(base.Ui32(v272)>>(uint(int32(8))%32)) | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)) = uint8(v277)
	v283 = v258 - (v17 + int32(11)) + int32(3)
	if v207 != 0 {
		v343 = v283
		goto L57
	} else {
		goto L75
	}
L75:
	;
	v348 = v283
	v349 = v270
	goto L56
L76:
	;
	v305 = l2<<(uint(int32(2))%32) + int32(124) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v299))) = uint8(v305)
	if l1 != v284 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v295 = v215<<(uint(int32(2))%32) | (l1 + (v201 ^ int32(-1))) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+11)) = uint8(v295)
	v299 = v17 + int32(12)
	goto L76
L78:
	;
	v348 = v321 - (v17 + int32(11))
	v349 = int32(-1)
	goto L56
L79:
	;
	v310 = int32(2)
	v317 = v215<<(uint(v310)%32) | (v284 + (l1 ^ int32(-1))) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+1)) = uint8(v317)
	v321 = v299 + v310
	goto L78
L80:
	;
	v321 = v299 + int32(1)
	goto L78
L81:
	;
	v348 = v330
	v349 = int32(-1)
	goto L56
L82:
	;
	v343 = v340
	goto L57
L83:
	;
	v382 = int32(0)
	v384 = F_objectGetVal(m, l0)
	mBase = m.M
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+int32(-1)))))
	switch v387 & int32(7) {
	case 0:
		goto L97
	case 1:
		goto L96
	case 2:
		goto L95
	case 3:
		goto L94
	case 4:
		goto L93
	default:
		v404 = v382
		goto L92
	}
L84:
	;
	v354 = F_objectGetVal(m, l0)
	mBase = m.M
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+int32(-1)))))
	switch v357 & int32(7) {
	case 0:
		goto L90
	case 1:
		goto L89
	case 2:
		goto L88
	case 3:
		goto L87
	case 4:
		goto L86
	default:
		v374 = int32(0)
		goto L85
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_hllSparseSet[0]))
	if base.Ui32(v377) < base.Ui32(v374+v350) {
		goto L3
	} else {
		goto L91
	}
L86:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v354+int32(-17))))
	v374 = v373
	goto L85
L87:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v354+int32(-9))))
	v374 = v370
	goto L85
L88:
	;
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354+int32(-5)))))
	v374 = v367
	goto L85
L89:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354+int32(-3)))))
	v374 = v364
	goto L85
L90:
	;
	v374 = int32(base.Ui32(v357) >> (uint(int32(3)) % 32))
	goto L85
L91:
	;
	goto L83
L92:
	;
	v406 = F_objectGetVal(m, l0)
	mBase = m.M
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+int32(-1)))))
	switch v409 & int32(7) {
	case 0:
		goto L103
	case 1:
		goto L102
	case 2:
		goto L101
	case 3:
		goto L100
	case 4:
		goto L99
	default:
		v426 = v382
		goto L98
	}
L93:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v384+int32(-17))))
	v404 = v403
	goto L92
L94:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v384+int32(-9))))
	v404 = v400
	goto L92
L95:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384+int32(-5)))))
	v404 = v397
	goto L92
L96:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+int32(-3)))))
	v404 = v394
	goto L92
L97:
	;
	v404 = int32(base.Ui32(v387) >> (uint(int32(3)) % 32))
	goto L92
L98:
	;
	if base.Ui32(v426) < base.Ui32(v404+v350) {
		goto L54
	} else {
		goto L104
	}
L99:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v406+int32(-9))))
	v426 = v425
	goto L98
L100:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v406+int32(-5))))
	v426 = v422
	goto L98
L101:
	;
	v419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406+int32(-3)))))
	v426 = v419
	goto L98
L102:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+int32(-2)))))
	v426 = v416
	goto L98
L103:
	;
	v426 = int32(base.Ui32(v409) >> (uint(int32(3)) % 32))
	goto L98
L104:
	;
	if v350 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v589 = F_objectGetVal(m, l0)
	mBase = m.M
	F_sdsIncrLen(m, v589, v350)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L32
	} else {
		goto L155
	}
L106:
	;
	if v207 == int32(64) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v432 = int32(2)
	goto L109
L108:
	;
	v432 = int32(1)
	goto L109
L109:
	;
	v433 = v200 + v432
	if base.Ui32(v156) <= base.Ui32(v433) {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	if base.Ui32(v433) < base.Ui32(v156) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v437 = v433
	goto L113
L112:
	;
	v437 = int32(0)
	goto L113
L113:
	;
	v438 = v437 + v350
	v439 = v156 - v437
	if v438 == v437 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L105
L115:
	;
	goto L114
L116:
	;
	v443 = v439 + v438
	if base.Ui32(int32(0)-v439<<(uint(int32(1))%32)) < base.Ui32(v437-v443) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v453 = (v437 ^ v438) & int32(3)
	if base.Ui32(v437) <= base.Ui32(v438) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v450 = F___memcpy(m, v438, v437, v439)
	mBase = m.M
	goto L114
L119:
	;
	if v559 == int32(0) {
		goto L115
	} else {
		goto L151
	}
L120:
	;
	if base.Ui32(v537) <= base.Ui32(int32(3)) {
		v558 = v536
		v559 = v537
		v560 = v538
		goto L119
	} else {
		goto L147
	}
L121:
	;
	if v453 != 0 {
		v519 = v439
		goto L131
	} else {
		goto L132
	}
L122:
	;
	if v453 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if v438&int32(3) != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v558 = v437
	v559 = v439
	v560 = v438
	goto L119
L125:
	;
	v460 = v437
	v461 = v439
	v462 = v438
	goto L127
L126:
	;
	v536 = v437
	v537 = v439
	v538 = v438
	goto L120
L127:
	;
	if v461 == int32(0) {
		goto L115
	} else {
		goto L129
	}
L129:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v466)
	v468 = int32(1)
	v469 = v460 + v468
	v471 = v461 + int32(-1)
	v473 = v462 + v468
	if v473&int32(3) == int32(0) {
		v536 = v469
		v537 = v471
		v538 = v473
		goto L120
	} else {
		goto L130
	}
L130:
	;
	v460 = v469
	v461 = v471
	v462 = v473
	goto L127
L131:
	;
	if v519 == int32(0) {
		goto L115
	} else {
		goto L143
	}
L132:
	;
	if v443&int32(3) == int32(0) {
		v499 = v439
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if base.Ui32(v499) <= base.Ui32(int32(3)) {
		v519 = v499
		goto L131
	} else {
		goto L139
	}
L134:
	;
	v484 = v439
	goto L135
L135:
	;
	if v484 == int32(0) {
		goto L115
	} else {
		goto L137
	}
L136:
	;
	v499 = v490
	goto L133
L137:
	;
	v490 = v484 + int32(-1)
	v491 = v438 + v490
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437+v490))))
	*(*uint8)(unsafe.Add(mBase, uint32(v491))) = uint8(v493)
	if v491&int32(3) != 0 {
		v484 = v490
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v506 = v499
	goto L140
L140:
	;
	v510 = v506 + int32(-4)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v437+v510)))
	*(*int32)(unsafe.Add(mBase, uint32(v438+v510))) = v513
	if base.Ui32(int32(3)) < base.Ui32(v510) {
		v506 = v510
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v519 = v510
	goto L131
L142:
	;
	goto L141
L143:
	;
	v526 = v519
	goto L144
L144:
	;
	v530 = v526 + int32(-1)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437+v530))))
	*(*uint8)(unsafe.Add(mBase, uint32(v438+v530))) = uint8(v533)
	if v530 != 0 {
		v526 = v530
		goto L144
	} else {
		goto L146
	}
L146:
	;
	goto L115
L147:
	;
	v543 = v536
	v544 = v537
	v545 = v538
	goto L148
L148:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v547
	v549 = int32(4)
	v550 = v543 + v549
	v552 = v545 + v549
	v554 = v544 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v554) {
		v543 = v550
		v544 = v554
		v545 = v552
		goto L148
	} else {
		goto L150
	}
L149:
	;
	v558 = v550
	v559 = v554
	v560 = v552
	goto L119
L150:
	;
	goto L149
L151:
	;
	v565 = v558
	v566 = v559
	v567 = v560
	goto L152
L152:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	*(*uint8)(unsafe.Add(mBase, uint32(v567))) = uint8(v569)
	v571 = int32(1)
	v576 = v566 + int32(-1)
	if v576 != 0 {
		v565 = v565 + v571
		v566 = v576
		v567 = v567 + v571
		goto L152
	} else {
		goto L154
	}
L153:
	;
	goto L115
L154:
	;
	goto L153
L155:
	;
	if v348 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v604 = v156 + v350
	goto L55
L157:
	;
	goto L156
L158:
	;
	v596 = F__emscripten_memcpy_bulkmem(m, v200, v17+int32(11), v348)
	mBase = m.M
	goto L157
L159:
	;
	v850 = F_objectGetVal(m, l0)
	mBase = m.M
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+15)))
	v853 = v851 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v850)+15)) = uint8(v853)
	v924 = int32(1)
	goto L2
L160:
	;
	v607 = v202
	goto L162
L161:
	;
	v607 = v159
	goto L162
L162:
	;
	if base.Ui32(v604) <= base.Ui32(v607) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v614 = v607
	v616 = int32(5)
	v618 = v604
	goto L164
L164:
	;
	v625 = v614 + int32(1)
	v632 = v616
	v634 = v618
	goto L166
L165:
	;
	goto L159
L166:
	;
	if v632 == int32(0) {
		goto L159
	} else {
		goto L168
	}
L167:
	;
	if base.Ui32(v833) < base.Ui32(v634) {
		v614 = v833
		v616 = v643
		v618 = v634
		goto L164
	} else {
		goto L225
	}
L168:
	;
	v643 = v632 + int32(-1)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v646 = v644 & int32(192)
	if v646 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L167
L170:
	;
	if v646 != int32(64) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v833 = v625
	goto L169
L172:
	;
	if base.Ui32(v625) < base.Ui32(v634) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v833 = v614 + int32(2)
	goto L169
L174:
	;
	v652 = int32(*(*int8)(unsafe.Add(mBase, uint32(v625))))
	if v652 <= int32(-1) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v833 = v625
	goto L169
L176:
	;
	v656 = v652 & int32(255)
	if (v656^v644)&int32(124) == int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v833 = v625
	goto L169
L178:
	;
	v662 = int32(3)
	v666 = v656&v662 + v644&v662
	if base.Ui32(v666) <= base.Ui32(int32(2)) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v833 = v625
	goto L169
L180:
	;
	v675 = v666 + int32(1) | v644&int32(124) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v614)+1)) = uint8(v675)
	v677 = v634 - v614
	if v614 == v625 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v833 = v625
	goto L169
L182:
	;
	v826 = F_objectGetVal(m, l0)
	mBase = m.M
	F_sdsIncrLen(m, v826, int32(-1))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L32
	} else {
		goto L223
	}
L183:
	;
	v825 = v614
	goto L182
L184:
	;
	v681 = v677 + v614
	if base.Ui32(int32(0)-v677<<(uint(int32(1))%32)) < base.Ui32(v625-v681) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v691 = (v625 ^ v614) & int32(3)
	if base.Ui32(v625) <= base.Ui32(v614) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v688 = F___memcpy(m, v614, v625, v677)
	mBase = m.M
	v825 = v688
	goto L182
L187:
	;
	if v797 == int32(0) {
		goto L183
	} else {
		goto L219
	}
L188:
	;
	if base.Ui32(v775) <= base.Ui32(int32(3)) {
		v796 = v774
		v797 = v775
		v798 = v776
		goto L187
	} else {
		goto L215
	}
L189:
	;
	if v691 != 0 {
		v757 = v677
		goto L199
	} else {
		goto L200
	}
L190:
	;
	if v691 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if v614&int32(3) != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v796 = v625
	v797 = v677
	v798 = v614
	goto L187
L193:
	;
	v698 = v625
	v699 = v677
	v700 = v614
	goto L195
L194:
	;
	v774 = v625
	v775 = v677
	v776 = v614
	goto L188
L195:
	;
	if v699 == int32(0) {
		goto L183
	} else {
		goto L197
	}
L197:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	*(*uint8)(unsafe.Add(mBase, uint32(v700))) = uint8(v704)
	v706 = int32(1)
	v707 = v698 + v706
	v709 = v699 + int32(-1)
	v711 = v700 + v706
	if v711&int32(3) == int32(0) {
		v774 = v707
		v775 = v709
		v776 = v711
		goto L188
	} else {
		goto L198
	}
L198:
	;
	v698 = v707
	v699 = v709
	v700 = v711
	goto L195
L199:
	;
	if v757 == int32(0) {
		goto L183
	} else {
		goto L211
	}
L200:
	;
	if v681&int32(3) == int32(0) {
		v737 = v677
		goto L201
	} else {
		goto L202
	}
L201:
	;
	if base.Ui32(v737) <= base.Ui32(int32(3)) {
		v757 = v737
		goto L199
	} else {
		goto L207
	}
L202:
	;
	v722 = v677
	goto L203
L203:
	;
	if v722 == int32(0) {
		goto L183
	} else {
		goto L205
	}
L204:
	;
	v737 = v728
	goto L201
L205:
	;
	v728 = v722 + int32(-1)
	v729 = v614 + v728
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v728))))
	*(*uint8)(unsafe.Add(mBase, uint32(v729))) = uint8(v731)
	if v729&int32(3) != 0 {
		v722 = v728
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	v744 = v737
	goto L208
L208:
	;
	v748 = v744 + int32(-4)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v625+v748)))
	*(*int32)(unsafe.Add(mBase, uint32(v614+v748))) = v751
	if base.Ui32(int32(3)) < base.Ui32(v748) {
		v744 = v748
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v757 = v748
	goto L199
L210:
	;
	goto L209
L211:
	;
	v764 = v757
	goto L212
L212:
	;
	v768 = v764 + int32(-1)
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v768))))
	*(*uint8)(unsafe.Add(mBase, uint32(v614+v768))) = uint8(v771)
	if v768 != 0 {
		v764 = v768
		goto L212
	} else {
		goto L214
	}
L214:
	;
	goto L183
L215:
	;
	v781 = v774
	v782 = v775
	v783 = v776
	goto L216
L216:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = v785
	v787 = int32(4)
	v788 = v781 + v787
	v790 = v783 + v787
	v792 = v782 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v792) {
		v781 = v788
		v782 = v792
		v783 = v790
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v796 = v788
	v797 = v792
	v798 = v790
	goto L187
L218:
	;
	goto L217
L219:
	;
	v803 = v796
	v804 = v797
	v805 = v798
	goto L220
L220:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	*(*uint8)(unsafe.Add(mBase, uint32(v805))) = uint8(v807)
	v809 = int32(1)
	v814 = v804 + int32(-1)
	if v814 != 0 {
		v803 = v803 + v809
		v804 = v814
		v805 = v805 + v809
		goto L220
	} else {
		goto L222
	}
L221:
	;
	goto L183
L222:
	;
	goto L221
L223:
	;
	v831 = v634 + int32(-1)
	if base.Ui32(v825) < base.Ui32(v831) {
		v632 = v643
		v634 = v831
		goto L166
	} else {
		goto L224
	}
L224:
	;
	goto L159
L225:
	;
	goto L165
L226:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	if v877 == int32(-1) {
		v924 = int32(-1)
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v881 = F_objectGetVal(m, l0)
	mBase = m.M
	v882 = int32(6)
	v883 = l1 * v882
	v884 = int32(8)
	v885 = base.I32_div_s(v883, v884)
	v886 = v881 + v885
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886+int32(17)))))
	v892 = v883 & v882
	v893 = v884 - v892
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+16)))
	if base.Ui32(l2) <= base.Ui32((v889<<(uint(v893)%32)|int32(base.Ui32(v895)>>(uint(v892)%32)))&int32(63)) {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v910 = v889&(int32(-64)>>(uint(v893)%32)) | int32(base.Ui32(l2)>>(uint(v893)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v886+int32(17)))) = uint8(v910)
	v918 = v895&(int32(63)<<(uint(v892)%32)^int32(-1)) | l2<<(uint(v892)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v886+int32(16)))) = uint8(v918)
	v924 = int32(1)
	goto L2
L230:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hllSparseToDense(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
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
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	v2 = int32(0)
	v14 = F_objectGetVal(m, l0)
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-1)))))
	switch v17 & int32(7) {
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
		v34 = v2
		goto L1
	}
L1:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	if v35 == int32(0) {
		v265 = v2
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-17))))
	v34 = v33
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(-9))))
	v34 = v30
	goto L1
L4:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+int32(-5)))))
	v34 = v27
	goto L1
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(-3)))))
	v34 = v24
	goto L1
L6:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return v265
L8:
	;
	v38 = int32(0)
	v41 = F_sdsnewlen(m, v38, int32(12304))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v45
	v47 = int32(8)
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v14+v47)))
	*(*int64)(unsafe.Add(mBase, uint32(v41+v47))) = v51
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)) = uint8(v53)
	if base.Ui32(v34) < base.Ui32(int32(17)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v258 = F_objectGetVal(m, l0)
	mBase = m.M
	F_sdsfree(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L33
	}
L12:
	;
	F_sdsfree(m, v41)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L9
	} else {
		goto L32
	}
L13:
	;
	v58 = int32(16)
	v59 = v41 + v58
	v63 = v38
	v65 = v14 + v58
	goto L14
L14:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v75 = v73 & int32(192)
	if v75 == int32(64) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if v233 == int32(16384) {
		goto L11
	} else {
		goto L31
	}
L16:
	;
	v239 = v65 + v235
	if base.Ui32(v239) < base.Ui32(v14+v34) {
		v63 = v233
		v65 = v239
		goto L14
	} else {
		goto L30
	}
L17:
	;
	v97 = v73 & int32(3)
	if int32(16384) < v97+v63+int32(1) {
		goto L12
	} else {
		goto L23
	}
L18:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	v93 = v63 + v73<<(uint(int32(8))%32)&int32(16128) + v90 + int32(1)
	if v93 <= int32(16384) {
		v233 = v93
		v235 = int32(2)
		goto L16
	} else {
		goto L22
	}
L19:
	;
	if v75 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v78 = int32(1)
	v81 = v63 + v73 + v78
	if v81 <= int32(16384) {
		v233 = v81
		v235 = v78
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	goto L12
L23:
	;
	v103 = int32(6)
	v104 = v63 * v103
	v105 = int32(8)
	v106 = base.I32_div_s(v104, v105)
	v107 = v59 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v111 = v104 & v103
	v120 = int32(1)
	v121 = int32(base.Ui32(v73)>>(uint(int32(2))%32))&int32(31) + v120
	v123 = v108&(int32(63)<<(uint(v111)%32)^int32(-1)) | v121<<(uint(v111)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v123)
	v127 = v107 + v120
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v131 = v105 - v111
	v135 = v128&(int32(-64)>>(uint(v131)%32)) | int32(base.Ui32(v121)>>(uint(v131)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v135)
	v138 = v63 + v120
	if v97 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v139 = int32(6)
	v140 = v138 * v139
	v141 = int32(8)
	v142 = base.I32_div_s(v140, v141)
	v143 = v59 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v147 = v140 & v139
	v153 = v144&(int32(63)<<(uint(v147)%32)^int32(-1)) | v121<<(uint(v147)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v153)
	v155 = int32(1)
	v157 = v143 + v155
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v161 = v141 - v147
	v165 = v158&(int32(-64)>>(uint(v161)%32)) | int32(base.Ui32(v121)>>(uint(v161)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v165)
	v168 = v63 + int32(2)
	if v97 != v155 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v233 = v138
	v235 = v120
	goto L16
L26:
	;
	v171 = int32(6)
	v172 = v168 * v171
	v173 = int32(8)
	v174 = base.I32_div_s(v172, v173)
	v175 = v59 + v174
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v179 = v172 & v171
	v185 = v176&(int32(63)<<(uint(v179)%32)^int32(-1)) | v121<<(uint(v179)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v185)
	v187 = int32(1)
	v189 = v175 + v187
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v193 = v173 - v179
	v197 = v190&(int32(-64)>>(uint(v193)%32)) | int32(base.Ui32(v121)>>(uint(v193)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v197)
	v200 = v63 + int32(3)
	if v97 != int32(2) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v233 = v168
	v235 = v155
	goto L16
L28:
	;
	v203 = int32(6)
	v204 = v200 * v203
	v205 = int32(8)
	v206 = base.I32_div_s(v204, v205)
	v207 = v59 + v206
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	v211 = v204 & v203
	v217 = v208&(int32(63)<<(uint(v211)%32)^int32(-1)) | v121<<(uint(v211)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v217)
	v219 = int32(1)
	v221 = v207 + v219
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	v225 = v205 - v211
	v229 = v222&(int32(-64)>>(uint(v225)%32)) | int32(base.Ui32(v121)>>(uint(v225)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v221))) = uint8(v229)
	v233 = v63 + int32(4)
	v235 = v219
	goto L16
L29:
	;
	v233 = v200
	v235 = v187
	goto L16
L30:
	;
	goto L15
L31:
	;
	goto L12
L32:
	;
	return int32(-1)
L33:
	;
	F_objectSetVal(m, l0, v41)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v265 = int32(0)
	goto L7
}
func F_hpexpireatCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_hexpireGenericCommand(m, l0, int64(0), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_hstrlenCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_hstrlenCommand[0]))
	v14 = F_lookupKeyReadOrReply(m, l0, v11, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	return
L3:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = F_checkType(m, l0, v14, int32(4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = F_objectGetVal(m, v22)
	mBase = m.M
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
	v36 = F_hashTypeGetValue(m, v14, v23, v8+int32(12), v8+int32(8), v8, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v173))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L78
	}
L8:
	;
	if v36 != 0 {
		v173 = v24
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if int64(-1) < v42 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v173 = v41
	goto L7
L12:
	;
	v173 = v172
	goto L7
L13:
	;
	v108 = int32(0)
	if base.Ui64(v42) < base.Ui64(int64(10)) {
		v165 = v108
		goto L47
	} else {
		goto L48
	}
L14:
	;
	v47 = int32(0)
	v49 = int64(0) - v42
	if base.Ui64(v49) < base.Ui64(int64(10)) {
		v99 = v47
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v172 = v103 + v104 + int32(1)
	goto L12
L16:
	;
	v103 = v99
	v104 = int32(1)
	goto L15
L17:
	;
	v52 = v49
	v53 = v47
	goto L18
L18:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v52) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v99 = v93
	goto L16
L20:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v52) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v103 = v53
	v104 = int32(2)
	goto L15
L22:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v52) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v103 = v53
	v104 = int32(3)
	goto L15
L24:
	;
	v93 = v53 + int32(12)
	v97 = base.I64_div_u_s(v52, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v52) {
		v52 = v97
		v53 = v93
		goto L18
	} else {
		goto L46
	}
L25:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v52) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v52) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v52) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v52) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v52) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v52) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v103 = v53
	v104 = int32(4)
	goto L15
L32:
	;
	v74 = int32(6)
	goto L34
L33:
	;
	v74 = int32(5)
	goto L34
L34:
	;
	v103 = v53
	v104 = v74
	goto L15
L35:
	;
	v79 = int32(8)
	goto L37
L36:
	;
	v79 = int32(7)
	goto L37
L37:
	;
	v103 = v53
	v104 = v79
	goto L15
L38:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v52) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v52) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v86 = int32(10)
	goto L42
L41:
	;
	v86 = int32(9)
	goto L42
L42:
	;
	v103 = v53
	v104 = v86
	goto L15
L43:
	;
	v91 = int32(12)
	goto L45
L44:
	;
	v91 = int32(11)
	goto L45
L45:
	;
	v103 = v53
	v104 = v91
	goto L15
L46:
	;
	goto L19
L47:
	;
	v172 = int32(1) + v165
	goto L12
L48:
	;
	v111 = v42
	v112 = v108
	goto L49
L49:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v111) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v165 = v159
	goto L47
L51:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v111) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v172 = int32(2) + v112
	goto L12
L53:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v111) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v172 = int32(3) + v112
	goto L12
L55:
	;
	v159 = v112 + int32(12)
	v163 = base.I64_div_u_s(v111, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v111) {
		v111 = v163
		v112 = v159
		goto L49
	} else {
		goto L77
	}
L56:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v111) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v111) {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v111) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v111) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v111) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v111) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v172 = int32(4) + v112
	goto L12
L63:
	;
	v136 = int32(6)
	goto L65
L64:
	;
	v136 = int32(5)
	goto L65
L65:
	;
	v172 = v136 + v112
	goto L12
L66:
	;
	v142 = int32(8)
	goto L68
L67:
	;
	v142 = int32(7)
	goto L68
L68:
	;
	v172 = v142 + v112
	goto L12
L69:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v111) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v111) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v150 = int32(10)
	goto L73
L72:
	;
	v150 = int32(9)
	goto L73
L73:
	;
	v172 = v150 + v112
	goto L12
L74:
	;
	v156 = int32(12)
	goto L76
L75:
	;
	v156 = int32(11)
	goto L76
L76:
	;
	v172 = v156 + v112
	goto L12
L77:
	;
	goto L50
L78:
	;
	goto L1
}
