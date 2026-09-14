package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_appendCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
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
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v198 int32
	_ = v198
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = F_lookupKeyWrite(m, v14, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 != 0 {
			v61 = int32(0)
			v63 = F_checkType(m, l0, v17, v61)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v63 != 0 {
					m.G0 = v12 + int32(16)
					return
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
					v67 = F_stringObjectLen(m, v17)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = F_objectGetVal(m, v66)
						mBase = m.M
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-1)))))
						switch v72 & int32(7) {
						case 0:
							v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
						case 1:
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+int32(-3)))))
							v89 = v79
						case 2:
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+int32(-5)))))
							v89 = v82
						case 3:
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-9))))
							v89 = v85
						case 4:
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v69+int32(-17))))
							v89 = v88
						default:
							v89 = v61
						}
						v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						if v91 != int64(-1) {
							v95 = int32(1)
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
							if v96&v95 != 0 {
								v103 = v95
								v106 = v103
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v99 != 0 {
									v101 = F_isImportSlotMigrationJob(m, v99)
									mBase = m.M
									v103 = v101
									v106 = v103
								} else {
									v106 = int32(0)
								}
							}
						} else {
							v106 = int32(1)
						}
						if v106 != 0 {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
							v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
							v119 = F_dbUnshareStringValue(m, v116, v118, v17)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								v121 = F_objectGetVal(m, v119)
								mBase = m.M
								v122 = F_objectGetVal(m, v66)
								mBase = m.M
								v123 = int32(0)
								v125 = F_objectGetVal(m, v66)
								mBase = m.M
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-1)))))
								switch v128 & int32(7) {
								case 0:
									v145 = int32(base.Ui32(v128) >> (uint(int32(3)) % 32))
								case 1:
									v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-3)))))
									v145 = v135
								case 2:
									v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125+int32(-5)))))
									v145 = v138
								case 3:
									v141 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-9))))
									v145 = v141
								case 4:
									v144 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-17))))
									v145 = v144
								default:
									v145 = v123
								}
								v146 = F_sdscatlen(m, v121, v122, v145)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return
								} else {
									F_objectSetVal(m, v119, v146)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return
									} else {
										v150 = F_objectGetVal(m, v119)
										mBase = m.M
										v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-1)))))
										switch v153 & int32(7) {
										case 0:
											v170 = int32(base.Ui32(v153) >> (uint(int32(3)) % 32))
										case 1:
											v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-3)))))
											v170 = v160
										case 2:
											v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150+int32(-5)))))
											v170 = v163
										case 3:
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-9))))
											v170 = v166
										case 4:
											v169 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-17))))
											v170 = v169
										default:
											v170 = v123
										}
										v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
										F_signalModifiedKey(m, l0, v177, v179)
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
											v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
											F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_appendCommand_0), v185, v187)
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return
											} else {
												v190 = int32(_a_F_appendCommand_1)
												v192 = *(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0]))
												*(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0])) = v192 + int64(1)
												F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v170))
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v111 = *(*int64)(unsafe.Add(mBase, _c_F_appendCommand[1]))
							if base.I64_extend_i32_u(v89)+base.I64_extend_i32_u(v67) <= v111 {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
								v119 = F_dbUnshareStringValue(m, v116, v118, v17)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									v121 = F_objectGetVal(m, v119)
									mBase = m.M
									v122 = F_objectGetVal(m, v66)
									mBase = m.M
									v123 = int32(0)
									v125 = F_objectGetVal(m, v66)
									mBase = m.M
									v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-1)))))
									switch v128 & int32(7) {
									case 0:
										v145 = int32(base.Ui32(v128) >> (uint(int32(3)) % 32))
									case 1:
										v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+int32(-3)))))
										v145 = v135
									case 2:
										v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125+int32(-5)))))
										v145 = v138
									case 3:
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-9))))
										v145 = v141
									case 4:
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(-17))))
										v145 = v144
									default:
										v145 = v123
									}
									v146 = F_sdscatlen(m, v121, v122, v145)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										F_objectSetVal(m, v119, v146)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return
										} else {
											v150 = F_objectGetVal(m, v119)
											mBase = m.M
											v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-1)))))
											switch v153 & int32(7) {
											case 0:
												v170 = int32(base.Ui32(v153) >> (uint(int32(3)) % 32))
											case 1:
												v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+int32(-3)))))
												v170 = v160
											case 2:
												v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150+int32(-5)))))
												v170 = v163
											case 3:
												v166 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-9))))
												v170 = v166
											case 4:
												v169 = *(*int32)(unsafe.Add(mBase, uint32(v150+int32(-17))))
												v170 = v169
											default:
												v170 = v123
											}
											v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
											F_signalModifiedKey(m, l0, v177, v179)
											mBase = m.M
											v181 = m.ExcPending
											if v181 != 0 {
												return
											} else {
												v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
												v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
												v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
												F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_appendCommand_0), v185, v187)
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return
												} else {
													v190 = int32(_a_F_appendCommand_1)
													v192 = *(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0]))
													*(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0])) = v192 + int64(1)
													F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v170))
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							} else {
								F_addReplyError(m, l0, int32(_a_F_appendCommand_2))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v20
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			if v22&int32(1) == int32(0) {
				v43 = F_tryObjectEncoding(m, v20)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v43
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
					F_dbAdd(m, v46, v48, v12+int32(12))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						F_incrRefCount(m, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v57
							v59 = F_stringObjectLen(m, v57)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v170 = v59
								v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
								F_signalModifiedKey(m, l0, v177, v179)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
									v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
									F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_appendCommand_0), v185, v187)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return
									} else {
										v190 = int32(_a_F_appendCommand_1)
										v192 = *(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0])) = v192 + int64(1)
										F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v170))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_incrRefCount(m, v20)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_dbAdd(m, v29, v31, v12+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						F_rewriteClientCommandArgument(m, l0, int32(2), v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v41 = F_stringObjectLen(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v170 = v41
								v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
								F_signalModifiedKey(m, l0, v177, v179)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
									v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
									v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
									F_notifyKeyspaceEvent(m, int32(8), int32(_a_F_appendCommand_0), v185, v187)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return
									} else {
										v190 = int32(_a_F_appendCommand_1)
										v192 = *(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0]))
										*(*int64)(unsafe.Add(mBase, _c_F_appendCommand[0])) = v192 + int64(1)
										F_addReplyLongLong(m, l0, base.I64_extend_i32_u(v170))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
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
func F_commandGroupStr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_commandGroupStr[0])))
	return v6
}
func F_commandHelpCommand(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v11 = F__emscripten_memcpy_bulkmem(m, v5, int32(_a_F_commandHelpCommand_0), int32(76))
	F_addReplyHelp(m, l0, v11)
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		m.G0 = v11 + int32(80)
		return
	}
}
func F_commandProcessed(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v8&int32(16) != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
		if v14 == int32(-1) {
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_commandProcessed[0]))
			if v18 == int32(0) {
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_commandProcessed[1]))
				if v22 == int32(0) {
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_commandProcessed[2]))
					if v25&int32(16)|v29 != 0 {
					} else {
						v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
						if v33 != int32(17) {
							v39 = v31
						} else {
							v37 = v31 + int64(15)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v37
							v39 = v37
						}
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_commandProcessed[3]))
						v44 = v41 + v14*int32(24)
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_commandProcessed[4])))
						*(*int64)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_commandProcessed[4]))) = v47 + v39
					}
				}
			}
		}
		F_resetClient(m, l0)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v54 == int32(0) {
				return
			} else {
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v54)+48))
				v59 = int32(1)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
				if v60&v59 != 0 {
					v67 = v59
					v70 = v67
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v63 != 0 {
						v65 = F_isImportSlotMigrationJob(m, v63)
						mBase = m.M
						v67 = v65
						v70 = v67
					} else {
						v70 = int32(0)
					}
				}
				if v70 == int32(0) {
				} else {
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
					if v73&int32(8) != 0 {
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+int32(-1)))))
						switch v82 & int32(7) {
						case 0:
							v99 = int32(base.Ui32(v82) >> (uint(int32(3)) % 32))
						case 1:
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+int32(-3)))))
							v99 = v89
						case 2:
							v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79+int32(-5)))))
							v99 = v92
						case 3:
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(-9))))
							v99 = v95
						case 4:
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v79+int32(-17))))
							v99 = v98
						default:
							v99 = int32(0)
						}
						v102 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+16)))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v77 - base.I64_extend_i32_u(v99) + v102
					}
				}
				v111 = int32(1)
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
				if v112&v111 != 0 {
					v119 = v111
					v122 = v119
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v115 != 0 {
						v117 = F_isImportSlotMigrationJob(m, v115)
						mBase = m.M
						v119 = v117
						v122 = v119
					} else {
						v122 = int32(0)
					}
				}
				if v122 == int32(0) {
					return
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v125)+48))
					if v126 == v57 {
						return
					} else {
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+56))
						v131 = v126 - v57
						F_replicationFeedStreamFromPrimaryStream(m, v128+v129, base.I32_wrap_i64(v131))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							v136 = *(*int64)(unsafe.Add(mBase, uint32(v135)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v135)+56)) = v136 + v131
							return
						}
					}
				}
			}
		}
	}
}
func F_copyCommand(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
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
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
	v18 = int32(3)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v18 < v19 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[0]))
	F_addReplyErrorObject(m, l0, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L38
	} else {
		goto L118
	}
L3:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v255 = F_lookupKey(m, v253, v248, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L38
	} else {
		goto L74
	}
L4:
	;
	v186 = F_objectGetVal(m, v181)
	mBase = m.M
	v187 = F_objectGetVal(m, v180)
	mBase = m.M
	v188 = int32(0)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+int32(-1)))))
	switch v195 & int32(7) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	case 3:
		goto L55
	case 4:
		goto L54
	default:
		v212 = v188
		goto L53
	}
L5:
	;
	v31 = v18
	v32 = v19
	v33 = v15
	v34 = int32(0)
	goto L7
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v179 = int32(1)
	v180 = v23
	v181 = v24
	v182 = v15
	goto L4
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v39 = v31 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v39)))
	v42 = F_objectGetVal(m, v41)
	mBase = m.M
	v43 = int32(_a_F_copyCommand_0)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v171 = base.B2i32(v164 == int32(0))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v15 != v163 {
		v246 = v171
		v247 = v173
		v248 = v174
		v249 = v163
		goto L3
	} else {
		goto L51
	}
L9:
	;
	v167 = v161 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v167 < v168 {
		v31 = v167
		v32 = v168
		v33 = v163
		v34 = v164
		goto L7
	} else {
		goto L50
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v39)))
	v86 = F_objectGetVal(m, v85)
	mBase = m.M
	v87 = int32(_a_F_copyCommand_1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	if v78-v80 != 0 {
		goto L10
	} else {
		goto L23
	}
L12:
	;
	v78 = F_tolower(m, v74)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v80 = F_tolower(m, v79)
	mBase = m.M
	goto L11
L13:
	;
	v48 = v42
	v49 = v43
	v50 = v46
	goto L16
L14:
	;
	v74 = int32(0)
	v75 = v43
	goto L12
L15:
	;
	v74 = v71 & int32(255)
	v75 = v70
	goto L12
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v52 == int32(0) {
		v70 = v49
		v71 = v50
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v70 = v64
	v71 = int32(0)
	goto L15
L18:
	;
	v56 = v50 & int32(255)
	if v56 == v52 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v63 = int32(1)
	v64 = v49 + v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v65 != 0 {
		v48 = v48 + v63
		v49 = v64
		v50 = v65
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v58 = F_tolower(m, v56)
	mBase = m.M
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v60 = F_tolower(m, v59)
	mBase = m.M
	if v58 == v60 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v70 = v49
	v71 = v62
	goto L15
L22:
	;
	goto L17
L23:
	;
	v161 = v31
	v163 = v33
	v164 = int32(1)
	goto L9
L24:
	;
	if int32(-2) < v31-v32 {
		goto L2
	} else {
		goto L36
	}
L25:
	;
	v122 = F_tolower(m, v118)
	mBase = m.M
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v124 = F_tolower(m, v123)
	mBase = m.M
	goto L24
L26:
	;
	v92 = v86
	v93 = v87
	v94 = v90
	goto L29
L27:
	;
	v118 = int32(0)
	v119 = v87
	goto L25
L28:
	;
	v118 = v115 & int32(255)
	v119 = v114
	goto L25
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v96 == int32(0) {
		v114 = v93
		v115 = v94
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v114 = v108
	v115 = int32(0)
	goto L28
L31:
	;
	v100 = v94 & int32(255)
	if v100 == v96 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v107 = int32(1)
	v108 = v93 + v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v109 != 0 {
		v92 = v92 + v107
		v93 = v108
		v94 = v109
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v102 = F_tolower(m, v100)
	mBase = m.M
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v104 = F_tolower(m, v103)
	mBase = m.M
	if v102 == v104 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v114 = v93
	v115 = v106
	goto L28
L35:
	;
	goto L30
L36:
	;
	if v122-v124 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v131 = v31 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+v131<<(uint(int32(2))%32))))
	v139 = F_getIntFromObjectOrReply(m, l0, v135, v13+int32(12), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	if v139 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v141 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v150 = F_createDatabaseIfNeeded(m, v141)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L38
	} else {
		goto L46
	}
L42:
	;
	F_addReplyError(m, l0, int32(_a_F_copyCommand_2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[1]))
	if v141 < v145 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L1
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v150
	if v16 < int32(0) {
		v161 = v131
		v163 = v150
		v164 = v34
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[1]))
	if v156 <= v16 {
		v161 = v131
		v163 = v150
		v164 = v34
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v158 = F_createDatabaseIfNeeded(m, v16)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v158
	v161 = v131
	v163 = v150
	v164 = v34
	goto L9
L50:
	;
	goto L8
L51:
	;
	v179 = v171
	v180 = v173
	v181 = v174
	v182 = v163
	goto L4
L52:
	;
	if v238 != 0 {
		v246 = v179
		v247 = v180
		v248 = v181
		v249 = v182
		goto L3
	} else {
		goto L71
	}
L53:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+int32(-1)))))
	switch v215 & int32(7) {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v232 = v188
		goto L59
	}
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v186+int32(-17))))
	v212 = v211
	goto L53
L55:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v186+int32(-9))))
	v212 = v208
	goto L53
L56:
	;
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186+int32(-5)))))
	v212 = v205
	goto L53
L57:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+int32(-3)))))
	v212 = v202
	goto L53
L58:
	;
	v212 = int32(base.Ui32(v195) >> (uint(int32(3)) % 32))
	goto L53
L59:
	;
	v233 = base.B2i32(base.Ui32(v212) < base.Ui32(v232))
	if base.Ui32(v212) < base.Ui32(v232) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v187+int32(-17))))
	v232 = v231
	goto L59
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v187+int32(-9))))
	v232 = v228
	goto L59
L62:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187+int32(-5)))))
	v232 = v225
	goto L59
L63:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+int32(-3)))))
	v232 = v222
	goto L59
L64:
	;
	v232 = int32(base.Ui32(v215) >> (uint(int32(3)) % 32))
	goto L59
L65:
	;
	v234 = v212
	goto L67
L66:
	;
	v234 = v232
	goto L67
L67:
	;
	v235 = F_memcmp(m, v186, v187, v234)
	mBase = m.M
	if v235 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v238 = v235
	goto L70
L69:
	;
	v238 = base.B2i32(base.Ui32(v232) < base.Ui32(v212)) - v233
	goto L70
L70:
	;
	goto L52
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[2]))
	F_addReplyErrorObject(m, l0, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L38
	} else {
		goto L72
	}
L72:
	;
	goto L1
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v264&int32(1) == int32(0) {
		v275 = int64(-1)
		goto L78
	} else {
		goto L79
	}
L74:
	;
	if v255 != 0 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[3]))
	F_addReply(m, l0, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L38
	} else {
		goto L76
	}
L76:
	;
	goto L1
L77:
	;
	v277 = F_lookupKey(m, v249, v247, int32(8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L38
	} else {
		goto L81
	}
L78:
	;
	goto L77
L79:
	;
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v255+(v264&int32(4)^int32(12)))))
	v275 = v274
	goto L78
L80:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	switch v288 & int32(15) {
	case 0:
		goto L86
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	case 5:
		goto L88
	case 6:
		goto L89
	default:
		goto L87
	}
L81:
	;
	if base.B2i32(v277 != int32(0))&v246 != int32(1) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[3]))
	F_addReply(m, l0, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L38
	} else {
		goto L83
	}
L83:
	;
	goto L1
L84:
	;
	if v277 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v310
	goto L84
L86:
	;
	v308 = F_dupStringObject(m, v255)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L38
	} else {
		goto L102
	}
L87:
	;
	F_addReplyError(m, l0, int32(_a_F_copyCommand_3))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L38
	} else {
		goto L101
	}
L88:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v249)+28))
	v302 = F_moduleTypeDupOrReply(m, l0, v248, v247, v301, v255)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L38
	} else {
		goto L99
	}
L89:
	;
	v299 = F_streamDup(m, v255)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L38
	} else {
		goto L98
	}
L90:
	;
	v297 = F_hashTypeDup(m, v255)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L38
	} else {
		goto L97
	}
L91:
	;
	v295 = F_zsetDup(m, v255)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L38
	} else {
		goto L96
	}
L92:
	;
	v293 = F_setTypeDup(m, v255)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L38
	} else {
		goto L95
	}
L93:
	;
	v291 = F_listTypeDup(m, v255)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L38
	} else {
		goto L94
	}
L94:
	;
	v310 = v291
	goto L85
L95:
	;
	v310 = v293
	goto L85
L96:
	;
	v310 = v295
	goto L85
L97:
	;
	v310 = v297
	goto L85
L98:
	;
	v310 = v299
	goto L85
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v302
	if v302 != 0 {
		goto L84
	} else {
		goto L100
	}
L100:
	;
	goto L1
L101:
	;
	goto L1
L102:
	;
	v310 = v308
	goto L85
L103:
	;
	F_dbAddInternal(m, v249, v247, v13+int32(8), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L38
	} else {
		goto L110
	}
L104:
	;
	v315 = int32(_a_F_copyCommand_4)
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[4]))
	v317 = F_objectGetVal(m, v247)
	mBase = m.M
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[5]))
	if v319 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v325 = F_dbGenericDeleteWithDictIndex(m, v249, v247, v316, int32(1), v323)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L38
	} else {
		goto L109
	}
L106:
	;
	v321 = F_getKeySlot(m, v317)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L38
	} else {
		goto L108
	}
L107:
	;
	v323 = int32(0)
	goto L105
L108:
	;
	v323 = v321
	goto L105
L109:
	;
	goto L103
L110:
	;
	if v275 == int64(-1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	F_touchWatchedKey(m, v249, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L38
	} else {
		goto L114
	}
L112:
	;
	v336 = F_setExpire(m, l0, v249, v247, v275)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L38
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v336
	goto L111
L114:
	;
	F_trackingInvalidateKey(m, l0, v340, int32(1))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L38
	} else {
		goto L115
	}
L115:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v249)+28))
	F_notifyKeyspaceEvent(m, int32(4), int32(_a_F_copyCommand_5), v349, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L38
	} else {
		goto L116
	}
L116:
	;
	v353 = int32(_a_F_copyCommand_4)
	v355 = *(*int64)(unsafe.Add(mBase, _c_F_copyCommand[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_copyCommand[6])) = v355 + int64(1)
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_copyCommand[7]))
	F_addReply(m, l0, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L38
	} else {
		goto L117
	}
L117:
	;
	goto L1
L118:
	;
	goto L1
}
func F_evalCommandHandler(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
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
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v425 = m.G3
	v431 = m.G8
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	m.T0[v432].(func(*base.Module, int32, int32, int32))(m, v425+int32(_a_F_evalCommandHandler_0), v425+int32(_a_F_evalCommandHandler_1), int32(614))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L102
	}
L2:
	;
	v17 = m.G3
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v19 = int32(0)
	v23 = m.G13
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v25 = m.T0[v24].(func(*base.Module, int32, int32, int32) int32)(m, v19, v17+int32(_a_F_evalCommandHandler_2), v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if l1 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v83 = m.G15
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v86 = m.G7
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = m.T0[v87].(func(*base.Module, int32, int32) int32)(m, v25, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L15
	}
L6:
	;
	v36 = int32(1)
	goto L7
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0+v36<<(uint(int32(2))%32))))
	v48 = m.G7
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = m.G16
	v54 = m.T0[v49].(func(*base.Module, int32, int32) int32)(m, v47, v13+int32(44))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v58 = m.T0[v57].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v25, v54, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v36 == l1+int32(-1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = v36 + int32(1)
	if v71 != l1 {
		v36 = v71
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v62 = m.G3
	v66 = m.G16
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = m.T0[v67].(func(*base.Module, int32, int32, int32, int32) int32)(m, int32(0), v25, v62+int32(_a_F_evalCommandHandler_3), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v88
	v92 = m.G3
	v97 = m.T0[v84].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v92+int32(_a_F_evalCommandHandler_4), v13+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v102 = m.T0[v101].(func(*base.Module, int32, int32) int32)(m, v25, v13+int32(44))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v107 = m.T0[v106].(func(*base.Module, int32, int32) int32)(m, v97, v13+int32(40))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	v416 = m.G14
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	m.T0[v417].(func(*base.Module))(m)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L101
	}
L19:
	;
	v244 = m.G17
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v245].(func(*base.Module, int32, int32))(m, int32(0), v25)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L3
	} else {
		goto L56
	}
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v112 = F_luaL_loadbuffer(m, v18, v107, v109, v92+int32(_a_F_evalCommandHandler_5))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	if v112 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L25
L23:
	;
	v144 = m.G3
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v148 = F_luaL_loadbuffer(m, v18, v102, v145, v144+int32(_a_F_evalCommandHandler_5))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L31
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v134 + int32(-16)
	goto L23
L25:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	goto L24
L31:
	;
	if v148 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v154 = F_lua_tolstring(m, v18, int32(-1), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v154
	v157 = m.G3
	v158 = m.G13
	v163 = F_lm_asprintf(m, v157+int32(_a_F_evalCommandHandler_6), v13+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v166 = m.G12
	v167 = int32(0)
	if v163&int32(3) == v167 {
		v189 = v163
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v223 = m.T0[v165].(func(*base.Module, int32, int32, int32) int32)(m, v167, v163, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L51
	}
L36:
	;
	v222 = v214 - v163
	goto L35
L37:
	;
	v193 = v189
	goto L45
L38:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v175 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v178 = v163
	goto L41
L40:
	;
	v222 = v163 - v163
	goto L35
L41:
	;
	v182 = v178 + int32(1)
	if v182&int32(3) == int32(0) {
		v189 = v182
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v187 != 0 {
		v178 = v182
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v214 = v182
	goto L36
L45:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v202 = int32(-2139062144)
	if (int32(16843008)-v199|v199)&v202 == v202 {
		v193 = v193 + int32(4)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v208 = v193
	goto L48
L47:
	;
	goto L46
L48:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v212 != 0 {
		v208 = v208 + int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v214 = v208
	goto L36
L50:
	;
	goto L49
L51:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	m.T0[v226].(func(*base.Module, int32, int32))(m, v223, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v229 = m.G11
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	m.T0[v230].(func(*base.Module, int32))(m, v163)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v234 = m.G17
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	m.T0[v235].(func(*base.Module, int32, int32))(m, int32(0), v25)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	m.T0[v239].(func(*base.Module, int32, int32))(m, int32(0), v97)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L18
L56:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	m.T0[v249].(func(*base.Module, int32, int32))(m, int32(0), v97)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v252 = int32(0)
	v255 = F_lua_pcall(m, v18, v252, int32(1), v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L59
	}
L58:
	;
	v366 = m.G3
	v367 = m.G13
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v369 = m.G12
	v374 = m.T0[v368].(func(*base.Module, int32, int32, int32) int32)(m, int32(0), v366+int32(_a_F_evalCommandHandler_7), int32(9))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L3
	} else {
		goto L90
	}
L59:
	;
	if v255 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v261 = F_lua_tolstring(m, v18, int32(-1), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v261
	v264 = m.G3
	v265 = m.G13
	v268 = F_lm_asprintf(m, v264+int32(_a_F_evalCommandHandler_8), v13)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v271 = m.G12
	v272 = int32(0)
	if v268&int32(3) == v272 {
		v294 = v268
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v328 = m.T0[v270].(func(*base.Module, int32, int32, int32) int32)(m, v272, v268, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L3
	} else {
		goto L79
	}
L64:
	;
	v327 = v319 - v268
	goto L63
L65:
	;
	v298 = v294
	goto L73
L66:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v280 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v283 = v268
	goto L69
L68:
	;
	v327 = v268 - v268
	goto L63
L69:
	;
	v287 = v283 + int32(1)
	if v287&int32(3) == int32(0) {
		v294 = v287
		goto L65
	} else {
		goto L71
	}
L71:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v292 != 0 {
		v283 = v287
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v319 = v287
	goto L64
L73:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	v307 = int32(-2139062144)
	if (int32(16843008)-v304|v304)&v307 == v307 {
		v298 = v298 + int32(4)
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v313 = v298
	goto L76
L75:
	;
	goto L74
L76:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v317 != 0 {
		v313 = v313 + int32(1)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v319 = v313
	goto L64
L78:
	;
	goto L77
L79:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	m.T0[v331].(func(*base.Module, int32, int32))(m, v328, int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v334 = m.G11
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	m.T0[v335].(func(*base.Module, int32))(m, v268)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	goto L84
L82:
	;
	goto L18
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v356 + int32(-16)
	goto L82
L84:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	goto L83
L90:
	;
	v378 = F_ldbCatStackValueRec(m, v374, v18, int32(-1), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	m.T0[v381].(func(*base.Module, int32, int32))(m, v378, int32(1))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	goto L95
L93:
	;
	goto L18
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v402 + int32(-16)
	goto L93
L95:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	goto L94
L101:
	;
	m.G0 = v13 + int32(48)
	return int32(1)
L102:
	;
	m.Env.Exit(m, int32(1))
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_execCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v291 int32
	_ = v291
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v16&int32(8) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_execCommand_0), int32(_a_F_execCommand_1), int32(278))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L61
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_addReplyError(m, l0, int32(_a_F_execCommand_2))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	goto L2
L7:
	;
	if v105&int32(4128) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L8:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v105 = v102
	goto L7
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	if v25 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = v14 + int32(8)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32
	goto L11
L11:
	;
	v37 = v14 + int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v39 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39+base.B2i32(v42 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v48
	goto L13
L15:
	;
	v54 = v39
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+24)))
	if v64&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L8
L18:
	;
	v78 = v14 + int32(8)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v80 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v69 = F_keyIsExpired(m, v67, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if v69 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v75 = v73 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v75
	v105 = v75
	goto L7
L22:
	;
	if v80 != 0 {
		v54 = v80
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80+base.B2i32(v83 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v89
	goto L23
L25:
	;
	goto L17
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v142 | int32(4096)
	F_unwatchAllKeys(m, l0)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L35
	}
L27:
	;
	if v105&int32(4096) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_resetClientMultiState(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L33
	}
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127<<(uint(int32(2))%32))+uint32(_c_F_execCommand[0])))
	F_addReply(m, l0, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_execCommand[1]))
	F_addReplyErrorObject(m, l0, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	goto L28
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v136 & int32(-4137)
	F_unwatchAllKeys(m, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	goto L2
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execCommand[2])) = int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+52)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	F_addReplyArrayLen(m, l0, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v163 < int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v142&int32(4096) != 0 {
		goto L57
	} else {
		goto L58
	}
L38:
	;
	v169 = v162
	v175 = int32(0)
	goto L39
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v180 = v175 * int32(20)
	v181 = v178 + v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v188
	v194 = F_ACLCheckAllPerm(m, l0, v14+int32(8))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L47
	}
L40:
	;
	goto L37
L41:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = v223 + v180
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+8)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = v231
	F_freeClientOriginalArgv(m, l0)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L55
	}
L42:
	;
	v212 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v212 == int64(-1) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v202 = int32(0)
	F_addACLLogEntry(m, l0, v194, int32(2), v201, v202, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L48
	}
L44:
	;
	v199 = int32(_a_F_execCommand_3)
	goto L43
L45:
	;
	v199 = int32(_a_F_execCommand_4)
	goto L43
L46:
	;
	v199 = int32(_a_F_execCommand_5)
	goto L43
L47:
	;
	switch v194 {
	case 0:
		goto L42
	default:
		goto L44
	case 2:
		v199 = int32(_a_F_execCommand_6)
		goto L43
	case 3:
		goto L46
	case 5:
		goto L45
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v199
	F_addReplyErrorFormat(m, l0, int32(_a_F_execCommand_7), v14)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v215 = int32(0)
	goto L52
L51:
	;
	v215 = int32(3)
	goto L52
L52:
	;
	F_call(m, l0, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v218&int32(16) != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L41
L55:
	;
	v236 = v175 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v236 < v238 {
		v169 = v237
		v175 = v236
		goto L39
	} else {
		goto L56
	}
L56:
	;
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v151
	F_resetClientMultiState(m, l0)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v253 & int32(-4097)
	goto L57
L59:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v264 & int32(-4137)
	F_unwatchAllKeys(m, l0)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execCommand[2])) = int32(0)
	goto L2
L61:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generateSelectCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	if base.Ui32(int32(9)) < base.Ui32(l0) {
		v16 = v6 + int32(16)
		v18 = base.I64_extend_i32_s(l0)
		if v18 <= int64(-1) {
			v27 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v27)
			v36 = v6 + int32(17)
			v37 = int32(20)
			v38 = int64(0) - v18
			v39 = int32(1)
		} else {
			v36 = v16
			v37 = int32(21)
			v38 = v18
			v39 = int32(0)
		}
		v40 = F_ull2string(m, v36, v37, v38)
		mBase = m.M
		if v40 == int32(0) {
			v59 = int32(0)
		} else {
			v59 = v40 + v39
		}
		v60 = F_sdsempty(m)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v6 + int32(16)
			v70 = F_sdscatfmt(m, v60, int32(_a_F_generateSelectCommand_0), v6)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				v72 = F_createObject(m, int32(0), v70)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = v72
					m.G0 = v6 + int32(48)
					return v74
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_generateSelectCommand[0])))
		v74 = v14
		m.G0 = v6 + int32(48)
		return v74
	}
}
func F_getCommandFlags(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+56))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
	if v6 == int32(295) {
		v11 = F_fcallGetCommandFlags(m, l0, v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return v11
		}
	} else {
		if v6 != int32(297) {
			if v6 == int32(293) {
				v24 = F_evalGetCommandFlags(m, l0, v5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = v24
					return v26
				}
			} else {
				if v6 == int32(292) {
					v24 = F_evalGetCommandFlags(m, l0, v5)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = v24
						return v26
					}
				} else {
					if v6 == int32(290) {
						v24 = F_evalGetCommandFlags(m, l0, v5)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							v26 = v24
							return v26
						}
					} else {
						if v6 != int32(294) {
							v26 = v5
							return v26
						} else {
							v24 = F_evalGetCommandFlags(m, l0, v5)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int64(0)
							} else {
								v26 = v24
								return v26
							}
						}
					}
				}
			}
		} else {
			v11 = F_fcallGetCommandFlags(m, l0, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				return v11
			}
		}
	}
}
func F_getFlushCommandFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 != int32(2) {
		v103 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v103 != int32(1) {
		goto L32
	} else {
		goto L33
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = F_objectGetVal(m, v8)
	mBase = m.M
	v10 = int32(_a_F_getFlushCommandFlags_0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v53 != int32(2) {
		v103 = v53
		goto L1
	} else {
		goto L17
	}
L4:
	;
	if v45-v47 != 0 {
		goto L3
	} else {
		goto L16
	}
L5:
	;
	v45 = F_tolower(m, v41)
	mBase = m.M
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v47 = F_tolower(m, v46)
	mBase = m.M
	goto L4
L6:
	;
	v15 = v9
	v16 = v10
	v17 = v13
	goto L9
L7:
	;
	v41 = int32(0)
	v42 = v10
	goto L5
L8:
	;
	v41 = v38 & int32(255)
	v42 = v37
	goto L5
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v19 == int32(0) {
		v37 = v16
		v38 = v17
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v37 = v31
	v38 = int32(0)
	goto L8
L11:
	;
	v23 = v17 & int32(255)
	if v23 == v19 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = int32(1)
	v31 = v16 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v32 != 0 {
		v15 = v15 + v30
		v16 = v31
		v17 = v32
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v25 = F_tolower(m, v23)
	mBase = m.M
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v27 = F_tolower(m, v26)
	mBase = m.M
	if v25 == v27 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v37 = v16
	v38 = v29
	goto L8
L15:
	;
	goto L10
L16:
	;
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v49
	return v49
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a_F_getFlushCommandFlags_1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v103 = v102
	goto L1
L19:
	;
	if v94-v96 != 0 {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L19
L21:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L24
L22:
	;
	v90 = int32(0)
	v91 = v59
	goto L20
L23:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L20
L24:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v86 = v80
	v87 = int32(0)
	goto L23
L26:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L23
L30:
	;
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	return int32(0)
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_getFlushCommandFlags[0]))
	F_addReplyErrorObject(m, l0, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_getFlushCommandFlags[1]))
	v108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.B2i32(v107 != v108)
	return v108
L34:
	;
	return int32(0)
L35:
	;
	return int32(-1)
}
func F_lookupCommandByCString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_lookupCommandByCString[0]))
	v5 = F_sdsnew(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_lookupCommandBySdsLogic(m, v4, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_sdsfree(m, v5)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	}
}
func F_prepareCommandQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_prepareCommandGeneric(m, v4, v5, l0+int32(288), l0+int32(80), l0+int32(292))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v15) <= base.Ui32(v14) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v18 = v14
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = v20 + v18*int32(40)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_prepareCommandGeneric(m, v24, v25, v23, v23+int32(32), v23+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v33 = v18 + int32(1)
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v33) < base.Ui32(v34) {
		v18 = v33
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_saveCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v4 = m.G0
	v6 = v4 - int32(64)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[0]))
	if v9 != int32(1) {
		v15 = int32(_a_F_saveCommand_0)
		v17 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[1]))
		*(*int64)(unsafe.Add(mBase, _c_F_saveCommand[1])) = v17 + int64(1)
		v23 = int32(0)
		v24 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(56)))) = v24
		v29 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(48)))) = v29
		v34 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(40)))) = v34
		v39 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(32)))) = v39
		v44 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[6]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(24)))) = v44
		v49 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[7]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(16)))) = v49
		v54 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[8]))
		*(*int64)(unsafe.Add(mBase, uint32(v6+int32(8)))) = v54
		v57 = *(*int64)(unsafe.Add(mBase, _c_F_saveCommand[9]))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v57
		v60 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[10]))
		if v60 != 0 {
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[11]))
			if v72 == int32(0) {
				v78 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[12]))
				if v78 != 0 {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+96))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
					v82 = v81
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
					v84 = v6
				} else {
					v84 = int32(0)
				}
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+96))
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
				v82 = v76
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
				v84 = v6
			}
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[13]))
			if v62 == int32(0) {
				v72 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[11]))
				if v72 == int32(0) {
					v78 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[12]))
					if v78 != 0 {
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+96))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
						v82 = v81
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
						v84 = v6
					} else {
						v84 = int32(0)
					}
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)+96))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
					v82 = v76
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
					v84 = v6
				}
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[14]))
				if v67 == int32(-1) {
					v70 = int32(0)
				} else {
					v70 = v67
				}
				v82 = v70
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
				v84 = v6
			}
		}
		v85 = int32(0)
		v87 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[15]))
		v89 = F_rdbSave(m, v85, v87, v84, v85)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return
		} else {
			if v89 != 0 {
				v96 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[16]))
				F_addReplyErrorObject(m, l0, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					m.G0 = v6 + int32(64)
					return
				}
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_saveCommand[17]))
				F_addReply(m, l0, v92)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					m.G0 = v6 + int32(64)
					return
				}
			}
		}
	} else {
		F_addReplyError(m, l0, int32(_a_F_saveCommand_1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			m.G0 = v6 + int32(64)
			return
		}
	}
}
func F_selectCommand(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v13 = F_getIntFromObjectOrReply(m, l0, v9, v6+int32(12), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			m.G0 = v6 + int32(16)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v15 < int32(0) {
				F_addReplyError(m, l0, int32(_a_F_selectCommand_0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_selectCommand[0]))
				if v15 < v19 {
					v24 = F_createDatabaseIfNeeded(m, v15)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v24
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
						if v27&int32(8) == int32(0) {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_selectCommand[1]))
							F_addReply(m, l0, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if v32 == int32(0) {
								F__serverAssert(m, int32(_a_F_selectCommand_1), int32(_a_F_selectCommand_2), int32(918))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v35
								v39 = *(*int32)(unsafe.Add(mBase, _c_F_selectCommand[1]))
								F_addReply(m, l0, v39)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				} else {
					F_addReplyError(m, l0, int32(_a_F_selectCommand_0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_sendCommand(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_sdsempty(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l1 + int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v109
	v115 = F_sdscatprintf(m, v13, int32(_a_F_sendCommand_0), v11+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L5:
	;
	v27 = v23
	v30 = v17
	v32 = int32(0)
	goto L7
L6:
	;
	v107 = v17
	v109 = int32(0)
	goto L4
L7:
	;
	if v27&int32(3) == int32(0) {
		v55 = v27
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v107 = v94
	v109 = v101
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v88
	v94 = F_sdscatprintf(m, v30, int32(_a_F_sendCommand_1), v11+int32(32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v88 = v80 - v27
	goto L9
L11:
	;
	v59 = v55
	goto L19
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = v27
	goto L15
L14:
	;
	v88 = v27 - v27
	goto L9
L15:
	;
	v48 = v44 + int32(1)
	if v48&int32(3) == int32(0) {
		v55 = v48
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v53 != 0 {
		v44 = v48
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v80 = v48
	goto L10
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 == v68 {
		v59 = v59 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v74 = v59
	goto L22
L21:
	;
	goto L20
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 != 0 {
		v74 = v74 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v80 = v74
	goto L10
L24:
	;
	goto L23
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v96 + int32(4)
	v101 = v32 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v102 != 0 {
		v27 = v102
		v30 = v94
		v32 = v101
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	v117 = F_sdscatsds(m, v115, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_sdsfree(m, v107)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(-1)))))
	switch v123 & int32(7) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	case 3:
		goto L32
	case 4:
		goto L31
	default:
		v140 = int32(0)
		goto L30
	}
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_sendCommand[0]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+92))
	v149 = m.T0[v148].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v117, v140, base.I64_extend_i32_s(v143*int32(1000)))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-17))))
	v140 = v139
	goto L30
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-9))))
	v140 = v136
	goto L30
L33:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(-5)))))
	v140 = v133
	goto L30
L34:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(-3)))))
	v140 = v130
	goto L30
L35:
	;
	v140 = int32(base.Ui32(v123) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	F_sdsfree(m, v117)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	if v149 != int32(-1) {
		v163 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v153 = F_sdsempty(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+88))
	v157 = m.T0[v156].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v157
	v161 = F_sdscatprintf(m, v153, int32(_a_F_sendCommand_2), v11)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L36
L42:
	;
	m.G0 = v11 + int32(48)
	return v163
}
func F_sortCommand(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_sortCommandGeneric(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_updateCommandLatencyHistogram(m *base.Module, l0 int32, l1 int64) {
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
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 != 0 {
		v11 = v4
		v12 = int64(1)
		if v12 < l1 {
			v15 = l1
		} else {
			v15 = v12
		}
		v16 = int64(1000000000)
		if v15 < v16 {
			v19 = v15
		} else {
			v19 = v16
		}
		if int64(0) <= v19 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
			v37 = v30 + v31 + base.I32_wrap_i64(base.I64_clz(v33|v19))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			v48 = (int32(64)-v37)<<(uint(v30)%32) - v40 + base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(base.I64_extend_i32_u(v31-v37+int32(63)))%64)))
			if v48 < int32(0) {
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
				if v51 <= v48 {
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
					if v53 == int32(0) {
						v66 = v48
					} else {
						v56 = int32(0)
						v59 = v48 - v53
						if v59 < v51 {
							v61 = v56
						} else {
							v61 = v56 - v51
						}
						if v59 < int32(0) {
							v64 = v51
						} else {
							v64 = v61
						}
						v66 = v64 + v59
					}
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
					v70 = v67 + v66<<(uint(int32(3))%32)
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
					v72 = int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71 + v72
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v75 + v72
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
					if v79 < v19 {
						v81 = v19
					} else {
						v81 = v79
					}
					*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v81
					v83 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
					if v19 < v83 {
						v85 = v19
					} else {
						v85 = v83
					}
					if v19 == int64(0) {
						v88 = v83
					} else {
						v88 = v85
					}
					*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v88
				}
			}
		} else {
		}
		return
	} else {
		v8 = F_hdr_init(m, int64(1), int64(1000000000), int32(2), l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v11 = v10
			v12 = int64(1)
			if v12 < l1 {
				v15 = l1
			} else {
				v15 = v12
			}
			v16 = int64(1000000000)
			if v15 < v16 {
				v19 = v15
			} else {
				v19 = v16
			}
			if int64(0) <= v19 {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
				v37 = v30 + v31 + base.I32_wrap_i64(base.I64_clz(v33|v19))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				v48 = (int32(64)-v37)<<(uint(v30)%32) - v40 + base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(base.I64_extend_i32_u(v31-v37+int32(63)))%64)))
				if v48 < int32(0) {
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
					if v51 <= v48 {
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
						if v53 == int32(0) {
							v66 = v48
						} else {
							v56 = int32(0)
							v59 = v48 - v53
							if v59 < v51 {
								v61 = v56
							} else {
								v61 = v56 - v51
							}
							if v59 < int32(0) {
								v64 = v51
							} else {
								v64 = v61
							}
							v66 = v64 + v59
						}
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
						v70 = v67 + v66<<(uint(int32(3))%32)
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						v72 = int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v70))) = v71 + v72
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v11)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = v75 + v72
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
						if v79 < v19 {
							v81 = v19
						} else {
							v81 = v79
						}
						*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v81
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
						if v19 < v83 {
							v85 = v19
						} else {
							v85 = v83
						}
						if v19 == int64(0) {
							v88 = v83
						} else {
							v88 = v85
						}
						*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v88
					}
				}
			} else {
			}
			return
		}
	}
}
