package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_entryConstruct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	v11 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if base.B2i32(l2 != v11)&base.B2i32(l3 == v11) != 0 {
		v32 = F_zmalloc_usable(m, l0, v17+int32(12))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			if l7 == int32(0) {
				v43 = v32
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = l4
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v39 - l7
				v43 = v32 + l7
			}
			if l5 != 0 {
				if l2 == int32(0) {
					v91 = v43
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
					switch v100 & int32(7) {
					case 0:
						v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
					case 1:
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v117 = v107
					case 2:
						v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v117 = v110
					case 3:
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v117 = v113
					case 4:
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v117 = v116
					default:
						v117 = int32(0)
					}
					v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v122 = v119 + int32(-1)
						v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
						if v123&int32(7) != 0 {
							if l5 != 0 {
								v132 = int32(0)
							} else {
								v132 = int32(16)
							}
							v133 = v123&int32(239) | v132
							*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
							v136 = base.B2i32(l7 != int32(0))
							if v123&int32(7) != 0 {
								v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
								v147 = v136
								v148 = v145
							} else {
								v147 = v136
								v148 = v133
							}
						} else {
							v147 = base.B2i32(l7 != int32(0))
							v148 = v123
						}
						v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
						if v151 == int32(0) {
							m.G0 = v17 + int32(16)
							return v119
						} else {
							v155 = v148 & int32(7)
							if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
								F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if v155 != 0 {
									v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
								} else {
									v169 = int32(0)
								}
								if v169 != v147 {
									F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									m.G0 = v17 + int32(16)
									return v119
								}
							}
						}
					}
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
					switch v59 & int32(7) {
					case 0:
						v76 = int32(base.Ui32(v59) >> (uint(int32(3)) % 32))
					case 1:
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
						v76 = v66
					case 2:
						v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
						v76 = v69
					case 3:
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
						v76 = v72
					case 4:
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
						v76 = v75
					default:
						v76 = int32(0)
					}
					v78 = F_sdswrite(m, v43+l9, v53-l9, int32(1), l2, v76)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_sdsfree(m, l2)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
							v84 = v43
							v85 = v82 - l8
							*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
							v91 = v84
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
							switch v100 & int32(7) {
							case 0:
								v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
							case 1:
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
								v117 = v107
							case 2:
								v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
								v117 = v110
							case 3:
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
								v117 = v113
							case 4:
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
								v117 = v116
							default:
								v117 = int32(0)
							}
							v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								v122 = v119 + int32(-1)
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
								if v123&int32(7) != 0 {
									if l5 != 0 {
										v132 = int32(0)
									} else {
										v132 = int32(16)
									}
									v133 = v123&int32(239) | v132
									*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
									v136 = base.B2i32(l7 != int32(0))
									if v123&int32(7) != 0 {
										v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
										v147 = v136
										v148 = v145
									} else {
										v147 = v136
										v148 = v133
									}
								} else {
									v147 = base.B2i32(l7 != int32(0))
									v148 = v123
								}
								v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
								if v151 == int32(0) {
									m.G0 = v17 + int32(16)
									return v119
								} else {
									v155 = v148 & int32(7)
									if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
										F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
										mBase = m.M
										v186 = m.ExcPending
										if v186 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if v155 != 0 {
											v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
										} else {
											v169 = int32(0)
										}
										if v169 != v147 {
											F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
											mBase = m.M
											v192 = m.ExcPending
											if v192 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											m.G0 = v17 + int32(16)
											return v119
										}
									}
								}
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v44 = l2
				} else {
					v44 = l3
				}
				*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
				v84 = v43 + int32(4)
				v85 = v48 + int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
				v91 = v84
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
				switch v100 & int32(7) {
				case 0:
					v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
				case 1:
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
					v117 = v107
				case 2:
					v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
					v117 = v110
				case 3:
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
					v117 = v113
				case 4:
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
					v117 = v116
				default:
					v117 = int32(0)
				}
				v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					v122 = v119 + int32(-1)
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
					if v123&int32(7) != 0 {
						if l5 != 0 {
							v132 = int32(0)
						} else {
							v132 = int32(16)
						}
						v133 = v123&int32(239) | v132
						*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
						v136 = base.B2i32(l7 != int32(0))
						if v123&int32(7) != 0 {
							v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
							v147 = v136
							v148 = v145
						} else {
							v147 = v136
							v148 = v133
						}
					} else {
						v147 = base.B2i32(l7 != int32(0))
						v148 = v123
					}
					v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
					if v151 == int32(0) {
						m.G0 = v17 + int32(16)
						return v119
					} else {
						v155 = v148 & int32(7)
						if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
							F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if v155 != 0 {
								v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
							} else {
								v169 = int32(0)
							}
							if v169 != v147 {
								F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								m.G0 = v17 + int32(16)
								return v119
							}
						}
					}
				}
			}
		}
	} else {
		if base.B2i32(l2|l3 == int32(0))&l5 != 0 {
			v32 = F_zmalloc_usable(m, l0, v17+int32(12))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if l7 == int32(0) {
					v43 = v32
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v32))) = l4
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v39 - l7
					v43 = v32 + l7
				}
				if l5 != 0 {
					if l2 == int32(0) {
						v91 = v43
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
						switch v100 & int32(7) {
						case 0:
							v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
						case 1:
							v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
							v117 = v107
						case 2:
							v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
							v117 = v110
						case 3:
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
							v117 = v113
						case 4:
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
							v117 = v116
						default:
							v117 = int32(0)
						}
						v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = v119 + int32(-1)
							v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
							if v123&int32(7) != 0 {
								if l5 != 0 {
									v132 = int32(0)
								} else {
									v132 = int32(16)
								}
								v133 = v123&int32(239) | v132
								*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
								v136 = base.B2i32(l7 != int32(0))
								if v123&int32(7) != 0 {
									v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
									v147 = v136
									v148 = v145
								} else {
									v147 = v136
									v148 = v133
								}
							} else {
								v147 = base.B2i32(l7 != int32(0))
								v148 = v123
							}
							v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
							if v151 == int32(0) {
								m.G0 = v17 + int32(16)
								return v119
							} else {
								v155 = v148 & int32(7)
								if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
									F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if v155 != 0 {
										v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
									} else {
										v169 = int32(0)
									}
									if v169 != v147 {
										F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
										mBase = m.M
										v192 = m.ExcPending
										if v192 != 0 {
											return int32(0)
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										m.G0 = v17 + int32(16)
										return v119
									}
								}
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
						switch v59 & int32(7) {
						case 0:
							v76 = int32(base.Ui32(v59) >> (uint(int32(3)) % 32))
						case 1:
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
							v76 = v66
						case 2:
							v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
							v76 = v69
						case 3:
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
							v76 = v72
						case 4:
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
							v76 = v75
						default:
							v76 = int32(0)
						}
						v78 = F_sdswrite(m, v43+l9, v53-l9, int32(1), l2, v76)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_sdsfree(m, l2)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
								v84 = v43
								v85 = v82 - l8
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
								v91 = v84
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
								switch v100 & int32(7) {
								case 0:
									v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
								case 1:
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
									v117 = v107
								case 2:
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
									v117 = v110
								case 3:
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
									v117 = v113
								case 4:
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
									v117 = v116
								default:
									v117 = int32(0)
								}
								v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = v119 + int32(-1)
									v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
									if v123&int32(7) != 0 {
										if l5 != 0 {
											v132 = int32(0)
										} else {
											v132 = int32(16)
										}
										v133 = v123&int32(239) | v132
										*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
										v136 = base.B2i32(l7 != int32(0))
										if v123&int32(7) != 0 {
											v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
											v147 = v136
											v148 = v145
										} else {
											v147 = v136
											v148 = v133
										}
									} else {
										v147 = base.B2i32(l7 != int32(0))
										v148 = v123
									}
									v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
									if v151 == int32(0) {
										m.G0 = v17 + int32(16)
										return v119
									} else {
										v155 = v148 & int32(7)
										if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
											F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if v155 != 0 {
												v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
											} else {
												v169 = int32(0)
											}
											if v169 != v147 {
												F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												m.G0 = v17 + int32(16)
												return v119
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						v44 = l2
					} else {
						v44 = l3
					}
					*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v84 = v43 + int32(4)
					v85 = v48 + int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
					v91 = v84
					v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
					switch v100 & int32(7) {
					case 0:
						v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
					case 1:
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
						v117 = v107
					case 2:
						v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
						v117 = v110
					case 3:
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
						v117 = v113
					case 4:
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
						v117 = v116
					default:
						v117 = int32(0)
					}
					v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						v122 = v119 + int32(-1)
						v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
						if v123&int32(7) != 0 {
							if l5 != 0 {
								v132 = int32(0)
							} else {
								v132 = int32(16)
							}
							v133 = v123&int32(239) | v132
							*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
							v136 = base.B2i32(l7 != int32(0))
							if v123&int32(7) != 0 {
								v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
								*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
								v147 = v136
								v148 = v145
							} else {
								v147 = v136
								v148 = v133
							}
						} else {
							v147 = base.B2i32(l7 != int32(0))
							v148 = v123
						}
						v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
						if v151 == int32(0) {
							m.G0 = v17 + int32(16)
							return v119
						} else {
							v155 = v148 & int32(7)
							if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
								F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									F_abort(m)
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if v155 != 0 {
									v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
								} else {
									v169 = int32(0)
								}
								if v169 != v147 {
									F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										F_abort(m)
										mBase = m.M
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									m.G0 = v17 + int32(16)
									return v119
								}
							}
						}
					}
				}
			}
		} else {
			if l2 != 0 {
				F__serverAssert(m, int32(_a_F_entryConstruct_3), int32(_a_F_entryConstruct_1), int32(312))
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					F_abort(m)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l3 == int32(0) {
					F__serverAssert(m, int32(_a_F_entryConstruct_3), int32(_a_F_entryConstruct_1), int32(312))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if l5 != 0 {
						F__serverAssert(m, int32(_a_F_entryConstruct_3), int32(_a_F_entryConstruct_1), int32(312))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v32 = F_zmalloc_usable(m, l0, v17+int32(12))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if l7 == int32(0) {
								v43 = v32
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v32))) = l4
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v39 - l7
								v43 = v32 + l7
							}
							if l5 != 0 {
								if l2 == int32(0) {
									v91 = v43
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
									switch v100 & int32(7) {
									case 0:
										v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
									case 1:
										v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
										v117 = v107
									case 2:
										v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
										v117 = v110
									case 3:
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
										v117 = v113
									case 4:
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
										v117 = v116
									default:
										v117 = int32(0)
									}
									v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v122 = v119 + int32(-1)
										v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
										if v123&int32(7) != 0 {
											if l5 != 0 {
												v132 = int32(0)
											} else {
												v132 = int32(16)
											}
											v133 = v123&int32(239) | v132
											*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
											v136 = base.B2i32(l7 != int32(0))
											if v123&int32(7) != 0 {
												v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
												*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
												v147 = v136
												v148 = v145
											} else {
												v147 = v136
												v148 = v133
											}
										} else {
											v147 = base.B2i32(l7 != int32(0))
											v148 = v123
										}
										v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
										if v151 == int32(0) {
											m.G0 = v17 + int32(16)
											return v119
										} else {
											v155 = v148 & int32(7)
											if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
												F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if v155 != 0 {
													v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
												} else {
													v169 = int32(0)
												}
												if v169 != v147 {
													F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return int32(0)
													} else {
														F_abort(m)
														mBase = m.M
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													m.G0 = v17 + int32(16)
													return v119
												}
											}
										}
									}
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-1)))))
									switch v59 & int32(7) {
									case 0:
										v76 = int32(base.Ui32(v59) >> (uint(int32(3)) % 32))
									case 1:
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(-3)))))
										v76 = v66
									case 2:
										v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2+int32(-5)))))
										v76 = v69
									case 3:
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-9))))
										v76 = v72
									case 4:
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(-17))))
										v76 = v75
									default:
										v76 = int32(0)
									}
									v78 = F_sdswrite(m, v43+l9, v53-l9, int32(1), l2, v76)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_sdsfree(m, l2)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
											v84 = v43
											v85 = v82 - l8
											*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
											v91 = v84
											v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
											switch v100 & int32(7) {
											case 0:
												v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
											case 1:
												v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
												v117 = v107
											case 2:
												v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
												v117 = v110
											case 3:
												v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
												v117 = v113
											case 4:
												v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
												v117 = v116
											default:
												v117 = int32(0)
											}
											v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v122 = v119 + int32(-1)
												v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
												if v123&int32(7) != 0 {
													if l5 != 0 {
														v132 = int32(0)
													} else {
														v132 = int32(16)
													}
													v133 = v123&int32(239) | v132
													*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
													v136 = base.B2i32(l7 != int32(0))
													if v123&int32(7) != 0 {
														v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
														*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
														v147 = v136
														v148 = v145
													} else {
														v147 = v136
														v148 = v133
													}
												} else {
													v147 = base.B2i32(l7 != int32(0))
													v148 = v123
												}
												v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
												if v151 == int32(0) {
													m.G0 = v17 + int32(16)
													return v119
												} else {
													v155 = v148 & int32(7)
													if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
														F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															F_abort(m)
															mBase = m.M
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														if v155 != 0 {
															v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
														} else {
															v169 = int32(0)
														}
														if v169 != v147 {
															F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
															mBase = m.M
															v192 = m.ExcPending
															if v192 != 0 {
																return int32(0)
															} else {
																F_abort(m)
																mBase = m.M
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															m.G0 = v17 + int32(16)
															return v119
														}
													}
												}
											}
										}
									}
								}
							} else {
								if l2 != 0 {
									v44 = l2
								} else {
									v44 = l3
								}
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
								v84 = v43 + int32(4)
								v85 = v48 + int32(-4)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v85
								v91 = v84
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
								switch v100 & int32(7) {
								case 0:
									v117 = int32(base.Ui32(v100) >> (uint(int32(3)) % 32))
								case 1:
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
									v117 = v107
								case 2:
									v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
									v117 = v110
								case 3:
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
									v117 = v113
								case 4:
									v116 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
									v117 = v116
								default:
									v117 = int32(0)
								}
								v119 = F_sdswrite(m, v91, l9, base.I32_extend8_s(l6), l1, v117)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = v119 + int32(-1)
									v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
									if v123&int32(7) != 0 {
										if l5 != 0 {
											v132 = int32(0)
										} else {
											v132 = int32(16)
										}
										v133 = v123&int32(239) | v132
										*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v133)
										v136 = base.B2i32(l7 != int32(0))
										if v123&int32(7) != 0 {
											v145 = v133&int32(247) | base.B2i32(l7 != int32(0))<<(uint(int32(3))%32)
											*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v145)
											v147 = v136
											v148 = v145
										} else {
											v147 = v136
											v148 = v133
										}
									} else {
										v147 = base.B2i32(l7 != int32(0))
										v148 = v123
									}
									v151 = *(*int32)(unsafe.Add(mBase, _c_F_entryConstruct[0]))
									if v151 == int32(0) {
										m.G0 = v17 + int32(16)
										return v119
									} else {
										v155 = v148 & int32(7)
										if l5 == base.B2i32(v155 != int32(0))&int32(base.Ui32(v148&int32(16))>>(uint(int32(4))%32)) {
											F__serverAssert(m, int32(_a_F_entryConstruct_0), int32(_a_F_entryConstruct_1), int32(344))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return int32(0)
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if v155 != 0 {
												v169 = int32(base.Ui32(v148)>>(uint(int32(3))%32)) & int32(1)
											} else {
												v169 = int32(0)
											}
											if v169 != v147 {
												F__serverAssert(m, int32(_a_F_entryConstruct_2), int32(_a_F_entryConstruct_1), int32(345))
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return int32(0)
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												m.G0 = v17 + int32(16)
												return v119
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
func F_entryCreate(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
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
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	if l1 == int32(0) {
		v35 = int32(-1)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
		switch v18 & int32(7) {
		case 0:
			v35 = int32(base.Ui32(v18) >> (uint(int32(3)) % 32))
		case 1:
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
			v35 = v25
		case 2:
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
			v35 = v28
		case 3:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
			v35 = v31
		case 4:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
			v35 = v34
		default:
			v35 = int32(0)
		}
	}
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v40 & int32(7) {
	case 0:
		v57 = int32(base.Ui32(v40) >> (uint(int32(3)) % 32))
	case 1:
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v57 = v47
	case 2:
		v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v57 = v50
	case 3:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v57 = v53
	case 4:
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v57 = v56
	default:
		v57 = int32(0)
	}
	if base.Ui32(int32(32)) <= base.Ui32(v57) {
		if base.Ui32(int32(253)) <= base.Ui32(v57) {
			if base.Ui32(v57) < base.Ui32(int32(65531)) {
				v69 = int32(2)
			} else {
				v69 = int32(3)
			}
			v70 = v69
		} else {
			v70 = int32(1)
		}
	} else {
		v70 = int32(0)
	}
	if v70 != 0 {
		v72 = v70
	} else {
		v72 = int32(1)
	}
	v74 = base.B2i32(l2 != int64(-1))
	if l2 != int64(-1) {
		v75 = v72
	} else {
		v75 = v70
	}
	v79 = v75 & int32(7)
	if base.Ui32(int32(4)) < base.Ui32(v79) {
		v87 = int32(0)
	} else {
		v86 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_c_F_entryCreate[0])))
		v87 = v86
	}
	v89 = v57 + int32(1)
	v90 = v87 + v89
	v92 = v74 << (uint(int32(3)) % 32)
	if v35 != int32(-1) {
		v109 = *(*int32)(unsafe.Add(mBase, _c_F_entryCreate[1]))
		v113 = v35 + v109 + int32(1)
		v114 = v90 + v92
		v115 = v113 + v114
		if base.Ui32(v115) < base.Ui32(int32(129)) {
			v141 = v115
			v142 = v90
			v143 = int32(1)
			v144 = v113
			v145 = v75
		} else {
			v118 = int32(0)
			if v75 == v118 {
				v135 = *(*int32)(unsafe.Add(mBase, _c_F_entryCreate[1]))
				v137 = v135 + v89
				v141 = v137 + v92 + int32(4)
				v142 = v137
				v143 = v118
				v144 = v113
				v145 = int32(1)
			} else {
				v141 = v114 + int32(4)
				v142 = v90
				v143 = v118
				v144 = v113
				v145 = v75
			}
		}
	} else {
		v96 = int32(0)
		v141 = v90 + v92
		v142 = v90
		v143 = v96
		v144 = v96
		v145 = v75
	}
	v148 = F_entryConstruct(m, v141, l0, l1, int32(0), l2, v143, v145, v92, v144, v142)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		return int32(0)
	} else {
		return v148
	}
}
func F_entryFreeValuePtr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	v5 = l0 + int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6&int32(7) == int32(0) {
		F__serverAssert(m, int32(_a_F_entryFreeValuePtr_0), int32(_a_F_entryFreeValuePtr_1), int32(175))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		if v6&int32(16) == int32(0) {
			F__serverAssert(m, int32(_a_F_entryFreeValuePtr_0), int32(_a_F_entryFreeValuePtr_1), int32(175))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v22 = v20 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v22) {
				v30 = int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_entryFreeValuePtr[0])))
				v30 = v29
			}
			v33 = l0 + v30 + int32(-4)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
			if v35&int32(7) == int32(0) {
				F_sdsfree(m, v34)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
					return
				}
			} else {
				v40 = int32(48)
				if v35&v40 != v40 {
					F_sdsfree(m, v34)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
						return
					}
				} else {
					F_valkey_free(m, v34)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(0)
						return
					}
				}
			}
		}
	}
}
func F_entryGetExpiry(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	v5 = int64(-1)
	v7 = l0 + int32(-1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8&int32(7) == int32(0) {
		v59 = v5
	} else {
		if v8&int32(8) == int32(0) {
			v59 = v5
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v24 = v22 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v24) {
				v32 = int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_entryGetExpiry[0])))
				v32 = v31
			}
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if int32(base.Ui32(v36&int32(16))>>(uint(int32(4))%32)) != 0 {
				v41 = int32(-4)
			} else {
				v41 = int32(0)
			}
			v44 = v36 & int32(7)
			if v44 != 0 {
				v45 = v41
			} else {
				v45 = int32(0)
			}
			if int32(base.Ui32(v36&int32(8))>>(uint(int32(3))%32)) != 0 {
				v53 = int32(-8)
			} else {
				v53 = int32(0)
			}
			if v44 != 0 {
				v55 = v53
			} else {
				v55 = int32(0)
			}
			v57 = *(*int64)(unsafe.Add(mBase, uint32(l0+v32+v45+v55)))
			v59 = v57
		}
	}
	return v59
}
func F_entryGetValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v10 = v8 & int32(7)
	if v10 == int32(0) {
		switch v10 {
		case 0:
			v30 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
		case 1:
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
			v30 = v20
		case 2:
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
			v30 = v23
		case 3:
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
			v30 = v26
		case 4:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
			v30 = v29
		default:
			v30 = int32(0)
		}
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetValue[0]))
		v45 = l0 + v30 + v43
		v47 = v45 + int32(1)
		if l1 == int32(0) {
			v138 = v47
			return v138
		} else {
			v50 = int32(0)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v50))))
			switch v53 & int32(7) {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v53) >> (uint(int32(3)) % 32))
				return v47
			case 1:
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-2)))))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v62
				return v47
			case 2:
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-4)))))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
				return v47
			case 3:
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-8))))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v72
				return v47
			case 4:
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-16))))
				v78 = v77
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
				return v47
			default:
				v78 = v50
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
				return v47
			}
		}
	} else {
		if v8&int32(16) != 0 {
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
			v88 = v86 & int32(7)
			if base.Ui32(int32(4)) < base.Ui32(v88) {
				v96 = int32(0)
			} else {
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(int32(2))%32))+uint32(_c_F_entryGetValue[1])))
				v96 = v95
			}
			v99 = l0 + v96 + int32(-4)
			if v8&int32(32) == int32(0) {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
				if l1 == int32(0) {
					v138 = v113
				} else {
					v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+int32(-1)))))
					switch v119 & int32(7) {
					case 0:
						v136 = int32(base.Ui32(v119) >> (uint(int32(3)) % 32))
					case 1:
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+int32(-3)))))
						v136 = v126
					case 2:
						v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113+int32(-5)))))
						v136 = v129
					case 3:
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v113+int32(-9))))
						v136 = v132
					case 4:
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v113+int32(-17))))
						v136 = v135
					default:
						v136 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v136
					v138 = v113
				}
				return v138
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
				if v104 != 0 {
					if l1 == int32(0) {
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v109
					}
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
					return v111
				} else {
					return int32(0)
				}
			}
		} else {
			switch v10 {
			case 0:
				v30 = int32(base.Ui32(v8) >> (uint(int32(3)) % 32))
			case 1:
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
				v30 = v20
			case 2:
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
				v30 = v23
			case 3:
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
				v30 = v26
			case 4:
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
				v30 = v29
			default:
				v30 = int32(0)
			}
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_entryGetValue[0]))
			v45 = l0 + v30 + v43
			v47 = v45 + int32(1)
			if l1 == int32(0) {
				v138 = v47
				return v138
			} else {
				v50 = int32(0)
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v50))))
				switch v53 & int32(7) {
				case 0:
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(v53) >> (uint(int32(3)) % 32))
					return v47
				case 1:
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+int32(-2)))))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v62
					return v47
				case 2:
					v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45+int32(-4)))))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
					return v47
				case 3:
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-8))))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v72
					return v47
				case 4:
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(-16))))
					v78 = v77
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
					return v47
				default:
					v78 = v50
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
					return v47
				}
			}
		}
	}
}
func F_entryHasStringRef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	v2 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	if v6&int32(7) == v2 {
		v19 = v2
	} else {
		if v6&int32(16) == int32(0) {
			v19 = v2
		} else {
			v19 = int32(base.Ui32(v6&int32(32)) >> (uint(int32(5)) % 32))
		}
	}
	return v19
}
