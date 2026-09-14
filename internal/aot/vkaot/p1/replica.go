package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_copyReplicaOutputBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v4 != 0 {
		F__serverAssert(m, int32(_a1661), int32(_a1630), int32(1784))
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
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		if v6 != 0 {
			F__serverAssert(m, int32(_a1661), int32(_a1630), int32(1784))
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
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+184))
			if v8 == int32(0) {
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+184)) = v8
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+188)) = v13
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16 + int32(1)
			}
			return
		}
	}
}
func F_freeReplicaKeysWithExpireAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = v8 + v9
	if base.Ui32(v10) < base.Ui32(int32(65)) {
		F_dictRelease(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v13 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[373]))
		*(*int32)(unsafe.Add(mBase, _consts[373])) = v15 + v10
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		F_bioCreateLazyFreeJob(m, int32(558), int32(1), v6)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_freeReplicaReferencedReplBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if v10&int32(4) == int32(0) {
		v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+184))
		if v78 == int32(0) {
			v92 = v77
			*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
			m.G0 = v8 + int32(32)
			return
		} else {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
			v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
			if v82 <= int32(0) {
				F__serverAssert(m, int32(_a1920), int32(_a1913), int32(435))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 + int32(-1)
				F_incrementalTrimReplicationBacklog(m, int32(64))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v92 = v91
					*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
					m.G0 = v8 + int32(32)
					return
				}
			}
		}
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v16 = int64(56)
		v18 = int64(65280)
		v20 = int64(40)
		v23 = int64(16711680)
		v25 = int64(24)
		v27 = int64(4278190080)
		v29 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v15<<(uint(v16)%64) | v15&v18<<(uint(v20)%64) | (v15&v23<<(uint(v25)%64) | v15&v27<<(uint(v29)%64)) | (int64(base.Ui64(v15)>>(uint(v29)%64))&v27 | int64(base.Ui64(v15)>>(uint(v25)%64))&v23 | (int64(base.Ui64(v15)>>(uint(v20)%64))&v18 | int64(base.Ui64(v15)>>(uint(v16)%64))))
		v53 = *(*int32)(unsafe.Add(mBase, _consts[569]))
		v58 = F_raxRemove(m, v53, v8+int32(24), int32(8), int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			if v58 == int32(0) {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+184))
				if v78 == int32(0) {
					v92 = v77
					*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
					m.G0 = v8 + int32(32)
					return
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
					if v82 <= int32(0) {
						F__serverAssert(m, int32(_a1920), int32(_a1913), int32(435))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 + int32(-1)
						F_incrementalTrimReplicationBacklog(m, int32(64))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v92 = v91
							*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				if int32(0) < v63 {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+184))
					if v78 == int32(0) {
						v92 = v77
						*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
						m.G0 = v8 + int32(32)
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
						if v82 <= int32(0) {
							F__serverAssert(m, int32(_a1920), int32(_a1913), int32(435))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 + int32(-1)
							F_incrementalTrimReplicationBacklog(m, int32(64))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
								v92 = v91
								*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					v66 = F_replicationGetReplicaName(m, l0)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v66
						F__serverLog(m, int32(0), int32(_a1923), v8)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+184))
							if v78 == int32(0) {
								v92 = v77
								*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
								m.G0 = v8 + int32(32)
								return
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
								if v82 <= int32(0) {
									F__serverAssert(m, int32(_a1920), int32(_a1913), int32(435))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v81))) = v82 + int32(-1)
									F_incrementalTrimReplicationBacklog(m, int32(64))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
										v92 = v91
										*(*int64)(unsafe.Add(mBase, uint32(v92)+184)) = int64(0)
										m.G0 = v8 + int32(32)
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
func F_getReplicaKeyWithExpireCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[330]))
	if v5 == v1 {
		v11 = v1
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		v11 = v8 + v9
	}
	return v11
}
func F_removeReplicaFromPsyncWait(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v294 int32
	_ = v294
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+168))
	v14 = int64(56)
	v16 = int64(65280)
	v18 = int64(40)
	v21 = int64(16711680)
	v23 = int64(24)
	v25 = int64(4278190080)
	v27 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v13<<(uint(v14)%64) | v13&v16<<(uint(v18)%64) | (v13&v21<<(uint(v23)%64) | v13&v25<<(uint(v27)%64)) | (int64(base.Ui64(v13)>>(uint(v27)%64))&v25 | int64(base.Ui64(v13)>>(uint(v23)%64))&v21 | (int64(base.Ui64(v13)>>(uint(v18)%64))&v16 | int64(base.Ui64(v13)>>(uint(v14)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[569]))
	v55 = v10 + int32(40)
	v56 = int32(8)
	v58 = v10 + int32(36)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	goto L4
L1:
	;
	v254 = int32(_a1919)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+104))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+184))
	if v257 == int32(0) {
		v270 = v254
		goto L40
	} else {
		goto L41
	}
L2:
	;
	if v211 != v56 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v202 = int32(0)
	v208 = v67
	v209 = v68
	v211 = v202
	v215 = v202
	goto L2
L4:
	;
	if base.Ui32(v68) < base.Ui32(int32(8)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v79 = v67
	v80 = v68
	v82 = int32(0)
	goto L7
L6:
	;
	v208 = v192
	v209 = v193
	v211 = v195
	v215 = base.B2i32(v198 != int32(0))
	goto L2
L7:
	;
	v88 = int32(base.Ui32(v80) >> (uint(int32(3)) % 32))
	v89 = int32(4)
	v90 = v79 + v89
	if v80&v89 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v192 = v183
	v193 = v184
	v195 = v168
	v198 = v173
	goto L6
L9:
	;
	v173 = int32(0)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v90+v88+(v173-v88)&int32(3)+v161<<(uint(int32(2))%32))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if base.Ui32(v184) < base.Ui32(int32(8)) {
		v192 = v183
		v193 = v184
		v195 = v168
		v198 = v173
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v82))))
	v139 = int32(0)
	goto L21
L11:
	;
	v95 = int32(0)
	if base.Ui32(v56) <= base.Ui32(v82) {
		v128 = v82
		v131 = v95
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v131 == v88 {
		v161 = v95
		v168 = v128
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v105 = v82
	v108 = v95
	goto L14
L14:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v108))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v105))))
	if v111 != v113 {
		v128 = v105
		v131 = v108
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v128 = v116
	v131 = v118
	goto L12
L16:
	;
	v115 = int32(1)
	v116 = v105 + v115
	v118 = v108 + v115
	if base.Ui32(v88) <= base.Ui32(v118) {
		v128 = v116
		v131 = v118
		goto L12
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v116) < base.Ui32(v56) {
		v105 = v116
		v108 = v118
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v192 = v79
	v193 = v80
	v195 = v128
	v198 = v131
	goto L6
L20:
	;
	if v139 != v88 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v139))))
	if v152 == v136&int32(255) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v154 = int32(1)
	v156 = v139 + v154
	if v156 != v88 {
		v139 = v156
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v208 = v79
	v209 = v80
	v211 = v82
	v215 = v154
	goto L2
L25:
	;
	v161 = v139
	v168 = v82 + int32(1)
	goto L9
L26:
	;
	v192 = v79
	v193 = v80
	v195 = v82
	v198 = v88
	goto L6
L27:
	;
	if base.Ui32(v168) < base.Ui32(v56) {
		v79 = v183
		v80 = v184
		v82 = v168
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L8
L29:
	;
	goto L1
L30:
	;
	if v209&int32(1) == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v223 = v209 & int32(4)
	if v215&base.B2i32(v223 != int32(0)) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v58 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v209&int32(2) != 0 {
		v249 = int32(0)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v249
	goto L29
L35:
	;
	v233 = int32(3)
	v234 = int32(base.Ui32(v209) >> (uint(v233) % 32))
	if v223 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v244 = int32(4)
	goto L38
L37:
	;
	v244 = v234 << (uint(int32(2)) % 32)
	goto L38
L38:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v208+v234+(int32(0)-v234)&v233+v244+int32(4))))
	v249 = v248
	goto L34
L39:
	;
	F__serverAssert(m, int32(_a1920), int32(_a1913), int32(303))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L46
	} else {
		goto L50
	}
L40:
	;
	v272 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v256)+184)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+204)) = v274 & int32(-33554433)
	v279 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v272 < v279 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	if v260 == int32(0) {
		v270 = v254
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v263 <= int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v263 + int32(-1)
	v270 = int32(_a1921)
	goto L40
L44:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
	v298 = int64(56)
	v300 = int64(65280)
	v302 = int64(40)
	v305 = int64(16711680)
	v307 = int64(24)
	v309 = int64(4278190080)
	v311 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v297<<(uint(v298)%64) | v297&v300<<(uint(v302)%64) | (v297&v305<<(uint(v307)%64) | v297&v309<<(uint(v311)%64)) | (int64(base.Ui64(v297)>>(uint(v311)%64))&v309 | int64(base.Ui64(v297)>>(uint(v307)%64))&v305 | (int64(base.Ui64(v297)>>(uint(v302)%64))&v300 | int64(base.Ui64(v297)>>(uint(v298)%64))))
	v335 = *(*int32)(unsafe.Add(mBase, _consts[569]))
	v340 = F_raxRemove(m, v335, v10+int32(40), int32(8), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L46
	} else {
		goto L49
	}
L45:
	;
	v282 = F_replicationGetReplicaName(m, l0)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return
L47:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v284)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)))) = v270
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v282
	F__serverLog(m, int32(0), int32(_a1922), v10)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	m.G0 = v10 + int32(48)
	return
L50:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replicaStartCommandStream(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v7&int32(32) != 0 {
		F__serverAssert(m, int32(_a1943), int32(_a1913), int32(1617))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
		v14 = *(*int64)(unsafe.Add(mBase, _consts[47]))
		if v14 != int64(0) {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v28&int32(4194304) != 0 {
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v31 == int32(0) {
					v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
					mBase = m.M
					if v38 == int32(0) {
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
						F_listLinkNodeHead(m, v46, l0+int32(168))
						mBase = m.M
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					switch v34 {
					case 0:
						v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
						mBase = m.M
						if v38 == int32(0) {
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
							v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
							F_listLinkNodeHead(m, v46, l0+int32(168))
							mBase = m.M
						}
					default:
					case 9, 11:
						if v28&int32(1024) != 0 {
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v37 != 0 {
							} else {
								v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
								mBase = m.M
								if v38 == int32(0) {
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
									v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
									F_listLinkNodeHead(m, v46, l0+int32(168))
									mBase = m.M
								}
							}
						}
					}
				}
			}
			m.G0 = v5 + int32(16)
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[594]))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v18
			F_replicationFeedReplicas(m, int32(-1), v5+int32(12), int32(1))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v28&int32(4194304) != 0 {
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v31 == int32(0) {
						v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
						mBase = m.M
						if v38 == int32(0) {
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
							v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
							F_listLinkNodeHead(m, v46, l0+int32(168))
							mBase = m.M
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						switch v34 {
						case 0:
							v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
							mBase = m.M
							if v38 == int32(0) {
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
								v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
								F_listLinkNodeHead(m, v46, l0+int32(168))
								mBase = m.M
							}
						default:
						case 9, 11:
							if v28&int32(1024) != 0 {
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
								if v37 != 0 {
								} else {
									v38 = F_clusterSlotMigrationShouldInstallWriteHandler(m, l0)
									mBase = m.M
									if v38 == int32(0) {
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v41 | int32(4194304)
										v46 = *(*int32)(unsafe.Add(mBase, _consts[478]))
										F_listLinkNodeHead(m, v46, l0+int32(168))
										mBase = m.M
									}
								}
							}
						}
					}
				}
				m.G0 = v5 + int32(16)
				return
			}
		}
	}
}
