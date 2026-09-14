package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_sdsCatPatternString(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v3 + int32(-1) {
	case 0:
		v8 = F_sdscatlen(m, l0, int32(_a13), int32(3))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v27 = v8
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v29 = F_sdscatsds(m, v27, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	case 1:
		v14 = F_sdscatlen(m, l0, int32(_a14), int32(3))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v27 = v14
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v29 = F_sdscatsds(m, v27, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	case 2:
		v25 = F_sdscatlen(m, l0, int32(_a15), int32(1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v29 = F_sdscatsds(m, v27, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	default:
		F__serverPanic_1(m, int32(_a6), int32(347), int32(_a16), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
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
func F_sdsConfigRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	if v5&int32(1) == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = v14
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		F_rewriteConfigSdsOption(m, l2, l1, v15, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v15 == int32(0) {
				return
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
				if v21&int32(1) == int32(0) {
					return
				} else {
					F_sdsfree(m, v15)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v11 = F_getModuleStringConfig(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v15 = v11
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_rewriteConfigSdsOption(m, l2, l1, v15, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v15 == int32(0) {
					return
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
					if v21&int32(1) == int32(0) {
						return
					} else {
						F_sdsfree(m, v15)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
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
func F_sdsConfigSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
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
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v8 == int32(0) {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		if v18&int32(1) == int32(0) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v28 = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v30 == int32(0) {
				v58 = v29
			} else {
				v33 = int32(0)
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
				switch v36 & int32(7) {
				case 0:
					v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
					if v53 == int32(0) {
						v58 = v33
					} else {
						v58 = v29
					}
				case 1:
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
					v53 = v43
					if v53 == int32(0) {
						v58 = v33
					} else {
						v58 = v29
					}
				case 2:
					v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
					v53 = v46
					if v53 == int32(0) {
						v58 = v33
					} else {
						v58 = v29
					}
				case 3:
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
					v53 = v49
					if v53 == int32(0) {
						v58 = v33
					} else {
						v58 = v29
					}
				case 4:
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
					v53 = v52
					if v53 == int32(0) {
						v58 = v33
					} else {
						v58 = v29
					}
				default:
					v58 = v33
				}
			}
			if v58 == v28 {
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v137&int32(256) == int32(0) {
					v147 = v137
					if v147&int32(512) != 0 {
						v152 = int32(1)
					} else {
						v152 = int32(2)
					}
					return v152
				} else {
					if v28 == int32(0) {
						v147 = v137
						if v147&int32(512) != 0 {
							v152 = int32(1)
						} else {
							v152 = int32(2)
						}
						return v152
					} else {
						F_sdsfree(m, v28)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v147 = v146
							if v147&int32(512) != 0 {
								v152 = int32(1)
							} else {
								v152 = int32(2)
							}
							return v152
						}
					}
				}
			} else {
				if v58 == int32(0) {
					F_sdsfree(m, v28)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
						if v120&int32(1) == int32(0) {
							if v58 != 0 {
								v130 = F_sdsdup(m, v58)
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int32(0)
								} else {
									v132 = v130
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
									return int32(1)
								}
							} else {
								v132 = int32(0)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
								return int32(1)
							}
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							v126 = F_setModuleStringConfig(m, v125, v58, l3)
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								return v126
							}
						}
					}
				} else {
					if v28 == int32(0) {
						F_sdsfree(m, v28)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
							if v120&int32(1) == int32(0) {
								if v58 != 0 {
									v130 = F_sdsdup(m, v58)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										v132 = v130
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
										return int32(1)
									}
								} else {
									v132 = int32(0)
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
									return int32(1)
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								v126 = F_setModuleStringConfig(m, v125, v58, l3)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									return v126
								}
							}
						}
					} else {
						v65 = int32(0)
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
						switch v72 & int32(7) {
						case 0:
							v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
						case 1:
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
							v89 = v79
						case 2:
							v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
							v89 = v82
						case 3:
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
							v89 = v85
						case 4:
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
							v89 = v88
						default:
							v89 = v65
						}
						v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-1)))))
						switch v92 & int32(7) {
						case 0:
							v109 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
						case 1:
							v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-3)))))
							v109 = v99
						case 2:
							v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+int32(-5)))))
							v109 = v102
						case 3:
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-9))))
							v109 = v105
						case 4:
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-17))))
							v109 = v108
						default:
							v109 = v65
						}
						v110 = base.B2i32(base.Ui32(v89) < base.Ui32(v109))
						if base.Ui32(v89) < base.Ui32(v109) {
							v111 = v89
						} else {
							v111 = v109
						}
						v112 = F_memcmp(m, v28, v58, v111)
						mBase = m.M
						if v112 != 0 {
							v115 = v112
						} else {
							v115 = base.B2i32(base.Ui32(v109) < base.Ui32(v89)) - v110
						}
						if v115 == int32(0) {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v137&int32(256) == int32(0) {
								v147 = v137
								if v147&int32(512) != 0 {
									v152 = int32(1)
								} else {
									v152 = int32(2)
								}
								return v152
							} else {
								if v28 == int32(0) {
									v147 = v137
									if v147&int32(512) != 0 {
										v152 = int32(1)
									} else {
										v152 = int32(2)
									}
									return v152
								} else {
									F_sdsfree(m, v28)
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int32(0)
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v147 = v146
										if v147&int32(512) != 0 {
											v152 = int32(1)
										} else {
											v152 = int32(2)
										}
										return v152
									}
								}
							}
						} else {
							F_sdsfree(m, v28)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
								if v120&int32(1) == int32(0) {
									if v58 != 0 {
										v130 = F_sdsdup(m, v58)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = v130
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v132 = int32(0)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
										return int32(1)
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									v126 = F_setModuleStringConfig(m, v125, v58, l3)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										return v126
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			v24 = F_getModuleStringConfig(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = v24
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v30 == int32(0) {
					v58 = v29
				} else {
					v33 = int32(0)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
					switch v36 & int32(7) {
					case 0:
						v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
						if v53 == int32(0) {
							v58 = v33
						} else {
							v58 = v29
						}
					case 1:
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
						v53 = v43
						if v53 == int32(0) {
							v58 = v33
						} else {
							v58 = v29
						}
					case 2:
						v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
						v53 = v46
						if v53 == int32(0) {
							v58 = v33
						} else {
							v58 = v29
						}
					case 3:
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
						v53 = v49
						if v53 == int32(0) {
							v58 = v33
						} else {
							v58 = v29
						}
					case 4:
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
						v53 = v52
						if v53 == int32(0) {
							v58 = v33
						} else {
							v58 = v29
						}
					default:
						v58 = v33
					}
				}
				if v58 == v28 {
					v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v137&int32(256) == int32(0) {
						v147 = v137
						if v147&int32(512) != 0 {
							v152 = int32(1)
						} else {
							v152 = int32(2)
						}
						return v152
					} else {
						if v28 == int32(0) {
							v147 = v137
							if v147&int32(512) != 0 {
								v152 = int32(1)
							} else {
								v152 = int32(2)
							}
							return v152
						} else {
							F_sdsfree(m, v28)
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v147 = v146
								if v147&int32(512) != 0 {
									v152 = int32(1)
								} else {
									v152 = int32(2)
								}
								return v152
							}
						}
					}
				} else {
					if v58 == int32(0) {
						F_sdsfree(m, v28)
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
							if v120&int32(1) == int32(0) {
								if v58 != 0 {
									v130 = F_sdsdup(m, v58)
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										v132 = v130
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
										return int32(1)
									}
								} else {
									v132 = int32(0)
									v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
									return int32(1)
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								v126 = F_setModuleStringConfig(m, v125, v58, l3)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									return v126
								}
							}
						}
					} else {
						if v28 == int32(0) {
							F_sdsfree(m, v28)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
								if v120&int32(1) == int32(0) {
									if v58 != 0 {
										v130 = F_sdsdup(m, v58)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = v130
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v132 = int32(0)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
										return int32(1)
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									v126 = F_setModuleStringConfig(m, v125, v58, l3)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										return v126
									}
								}
							}
						} else {
							v65 = int32(0)
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
							switch v72 & int32(7) {
							case 0:
								v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
							case 1:
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
								v89 = v79
							case 2:
								v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
								v89 = v82
							case 3:
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
								v89 = v85
							case 4:
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
								v89 = v88
							default:
								v89 = v65
							}
							v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-1)))))
							switch v92 & int32(7) {
							case 0:
								v109 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
							case 1:
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-3)))))
								v109 = v99
							case 2:
								v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+int32(-5)))))
								v109 = v102
							case 3:
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-9))))
								v109 = v105
							case 4:
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-17))))
								v109 = v108
							default:
								v109 = v65
							}
							v110 = base.B2i32(base.Ui32(v89) < base.Ui32(v109))
							if base.Ui32(v89) < base.Ui32(v109) {
								v111 = v89
							} else {
								v111 = v109
							}
							v112 = F_memcmp(m, v28, v58, v111)
							mBase = m.M
							if v112 != 0 {
								v115 = v112
							} else {
								v115 = base.B2i32(base.Ui32(v109) < base.Ui32(v89)) - v110
							}
							if v115 == int32(0) {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v137&int32(256) == int32(0) {
									v147 = v137
									if v147&int32(512) != 0 {
										v152 = int32(1)
									} else {
										v152 = int32(2)
									}
									return v152
								} else {
									if v28 == int32(0) {
										v147 = v137
										if v147&int32(512) != 0 {
											v152 = int32(1)
										} else {
											v152 = int32(2)
										}
										return v152
									} else {
										F_sdsfree(m, v28)
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
											return int32(0)
										} else {
											v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v147 = v146
											if v147&int32(512) != 0 {
												v152 = int32(1)
											} else {
												v152 = int32(2)
											}
											return v152
										}
									}
								}
							} else {
								F_sdsfree(m, v28)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
									if v120&int32(1) == int32(0) {
										if v58 != 0 {
											v130 = F_sdsdup(m, v58)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v132 = v130
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
												return int32(1)
											}
										} else {
											v132 = int32(0)
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										v126 = F_setModuleStringConfig(m, v125, v58, l3)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											return v126
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
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, v11, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
				if v18&int32(1) == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v28 = v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v30 == int32(0) {
						v58 = v29
					} else {
						v33 = int32(0)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
						switch v36 & int32(7) {
						case 0:
							v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
							if v53 == int32(0) {
								v58 = v33
							} else {
								v58 = v29
							}
						case 1:
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
							v53 = v43
							if v53 == int32(0) {
								v58 = v33
							} else {
								v58 = v29
							}
						case 2:
							v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
							v53 = v46
							if v53 == int32(0) {
								v58 = v33
							} else {
								v58 = v29
							}
						case 3:
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
							v53 = v49
							if v53 == int32(0) {
								v58 = v33
							} else {
								v58 = v29
							}
						case 4:
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
							v53 = v52
							if v53 == int32(0) {
								v58 = v33
							} else {
								v58 = v29
							}
						default:
							v58 = v33
						}
					}
					if v58 == v28 {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v137&int32(256) == int32(0) {
							v147 = v137
							if v147&int32(512) != 0 {
								v152 = int32(1)
							} else {
								v152 = int32(2)
							}
							return v152
						} else {
							if v28 == int32(0) {
								v147 = v137
								if v147&int32(512) != 0 {
									v152 = int32(1)
								} else {
									v152 = int32(2)
								}
								return v152
							} else {
								F_sdsfree(m, v28)
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return int32(0)
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v147 = v146
									if v147&int32(512) != 0 {
										v152 = int32(1)
									} else {
										v152 = int32(2)
									}
									return v152
								}
							}
						}
					} else {
						if v58 == int32(0) {
							F_sdsfree(m, v28)
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
								if v120&int32(1) == int32(0) {
									if v58 != 0 {
										v130 = F_sdsdup(m, v58)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = v130
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v132 = int32(0)
										v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
										return int32(1)
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									v126 = F_setModuleStringConfig(m, v125, v58, l3)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										return v126
									}
								}
							}
						} else {
							if v28 == int32(0) {
								F_sdsfree(m, v28)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
									if v120&int32(1) == int32(0) {
										if v58 != 0 {
											v130 = F_sdsdup(m, v58)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v132 = v130
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
												return int32(1)
											}
										} else {
											v132 = int32(0)
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										v126 = F_setModuleStringConfig(m, v125, v58, l3)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											return v126
										}
									}
								}
							} else {
								v65 = int32(0)
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
								switch v72 & int32(7) {
								case 0:
									v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
								case 1:
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
									v89 = v79
								case 2:
									v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
									v89 = v82
								case 3:
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
									v89 = v85
								case 4:
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
									v89 = v88
								default:
									v89 = v65
								}
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-1)))))
								switch v92 & int32(7) {
								case 0:
									v109 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
								case 1:
									v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-3)))))
									v109 = v99
								case 2:
									v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+int32(-5)))))
									v109 = v102
								case 3:
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-9))))
									v109 = v105
								case 4:
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-17))))
									v109 = v108
								default:
									v109 = v65
								}
								v110 = base.B2i32(base.Ui32(v89) < base.Ui32(v109))
								if base.Ui32(v89) < base.Ui32(v109) {
									v111 = v89
								} else {
									v111 = v109
								}
								v112 = F_memcmp(m, v28, v58, v111)
								mBase = m.M
								if v112 != 0 {
									v115 = v112
								} else {
									v115 = base.B2i32(base.Ui32(v109) < base.Ui32(v89)) - v110
								}
								if v115 == int32(0) {
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v137&int32(256) == int32(0) {
										v147 = v137
										if v147&int32(512) != 0 {
											v152 = int32(1)
										} else {
											v152 = int32(2)
										}
										return v152
									} else {
										if v28 == int32(0) {
											v147 = v137
											if v147&int32(512) != 0 {
												v152 = int32(1)
											} else {
												v152 = int32(2)
											}
											return v152
										} else {
											F_sdsfree(m, v28)
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
											} else {
												v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v147 = v146
												if v147&int32(512) != 0 {
													v152 = int32(1)
												} else {
													v152 = int32(2)
												}
												return v152
											}
										}
									}
								} else {
									F_sdsfree(m, v28)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
										if v120&int32(1) == int32(0) {
											if v58 != 0 {
												v130 = F_sdsdup(m, v58)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = v130
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
													return int32(1)
												}
											} else {
												v132 = int32(0)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
												return int32(1)
											}
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											v126 = F_setModuleStringConfig(m, v125, v58, l3)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												return v126
											}
										}
									}
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					v24 = F_getModuleStringConfig(m, v23)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v28 = v24
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v30 == int32(0) {
							v58 = v29
						} else {
							v33 = int32(0)
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
							switch v36 & int32(7) {
							case 0:
								v53 = int32(base.Ui32(v36) >> (uint(int32(3)) % 32))
								if v53 == int32(0) {
									v58 = v33
								} else {
									v58 = v29
								}
							case 1:
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
								v53 = v43
								if v53 == int32(0) {
									v58 = v33
								} else {
									v58 = v29
								}
							case 2:
								v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
								v53 = v46
								if v53 == int32(0) {
									v58 = v33
								} else {
									v58 = v29
								}
							case 3:
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
								v53 = v49
								if v53 == int32(0) {
									v58 = v33
								} else {
									v58 = v29
								}
							case 4:
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
								v53 = v52
								if v53 == int32(0) {
									v58 = v33
								} else {
									v58 = v29
								}
							default:
								v58 = v33
							}
						}
						if v58 == v28 {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v137&int32(256) == int32(0) {
								v147 = v137
								if v147&int32(512) != 0 {
									v152 = int32(1)
								} else {
									v152 = int32(2)
								}
								return v152
							} else {
								if v28 == int32(0) {
									v147 = v137
									if v147&int32(512) != 0 {
										v152 = int32(1)
									} else {
										v152 = int32(2)
									}
									return v152
								} else {
									F_sdsfree(m, v28)
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int32(0)
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v147 = v146
										if v147&int32(512) != 0 {
											v152 = int32(1)
										} else {
											v152 = int32(2)
										}
										return v152
									}
								}
							}
						} else {
							if v58 == int32(0) {
								F_sdsfree(m, v28)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
									if v120&int32(1) == int32(0) {
										if v58 != 0 {
											v130 = F_sdsdup(m, v58)
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int32(0)
											} else {
												v132 = v130
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
												return int32(1)
											}
										} else {
											v132 = int32(0)
											v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
											return int32(1)
										}
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										v126 = F_setModuleStringConfig(m, v125, v58, l3)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											return v126
										}
									}
								}
							} else {
								if v28 == int32(0) {
									F_sdsfree(m, v28)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
										if v120&int32(1) == int32(0) {
											if v58 != 0 {
												v130 = F_sdsdup(m, v58)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = v130
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
													return int32(1)
												}
											} else {
												v132 = int32(0)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
												return int32(1)
											}
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											v126 = F_setModuleStringConfig(m, v125, v58, l3)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												return v126
											}
										}
									}
								} else {
									v65 = int32(0)
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-1)))))
									switch v72 & int32(7) {
									case 0:
										v89 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
									case 1:
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+int32(-3)))))
										v89 = v79
									case 2:
										v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(-5)))))
										v89 = v82
									case 3:
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-9))))
										v89 = v85
									case 4:
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(-17))))
										v89 = v88
									default:
										v89 = v65
									}
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-1)))))
									switch v92 & int32(7) {
									case 0:
										v109 = int32(base.Ui32(v92) >> (uint(int32(3)) % 32))
									case 1:
										v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+int32(-3)))))
										v109 = v99
									case 2:
										v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+int32(-5)))))
										v109 = v102
									case 3:
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-9))))
										v109 = v105
									case 4:
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v58+int32(-17))))
										v109 = v108
									default:
										v109 = v65
									}
									v110 = base.B2i32(base.Ui32(v89) < base.Ui32(v109))
									if base.Ui32(v89) < base.Ui32(v109) {
										v111 = v89
									} else {
										v111 = v109
									}
									v112 = F_memcmp(m, v28, v58, v111)
									mBase = m.M
									if v112 != 0 {
										v115 = v112
									} else {
										v115 = base.B2i32(base.Ui32(v109) < base.Ui32(v89)) - v110
									}
									if v115 == int32(0) {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v137&int32(256) == int32(0) {
											v147 = v137
											if v147&int32(512) != 0 {
												v152 = int32(1)
											} else {
												v152 = int32(2)
											}
											return v152
										} else {
											if v28 == int32(0) {
												v147 = v137
												if v147&int32(512) != 0 {
													v152 = int32(1)
												} else {
													v152 = int32(2)
												}
												return v152
											} else {
												F_sdsfree(m, v28)
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v147 = v146
													if v147&int32(512) != 0 {
														v152 = int32(1)
													} else {
														v152 = int32(2)
													}
													return v152
												}
											}
										}
									} else {
										F_sdsfree(m, v28)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
											if v120&int32(1) == int32(0) {
												if v58 != 0 {
													v130 = F_sdsdup(m, v58)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v132 = v130
														v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
														*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
														return int32(1)
													}
												} else {
													v132 = int32(0)
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													*(*int32)(unsafe.Add(mBase, uint32(v133))) = v132
													return int32(1)
												}
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												v126 = F_setModuleStringConfig(m, v125, v58, l3)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													return v126
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
				return int32(0)
			}
		}
	}
}
func F_sdsMakeRoomForNonGreedy(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F__sdsMakeRoomFor(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_sdsResize(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
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
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
	var v441 int32
	_ = v441
	var v456 int32
	_ = v456
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	v23 = v21 & int32(7)
	switch v23 {
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
		v43 = v4
		v44 = v4
		goto L1
	}
L1:
	;
	switch v23 {
	case 0:
		goto L12
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	default:
		v59 = v4
		goto L7
	}
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
	v43 = int32(17)
	v44 = v41
	goto L1
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v43 = int32(9)
	v44 = v37
	goto L1
L4:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
	v43 = int32(5)
	v44 = v33
	goto L1
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v43 = int32(3)
	v44 = v29
	goto L1
L6:
	;
	v43 = int32(1)
	v44 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	if v59 != l1 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
	v59 = v58
	goto L7
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-5))))
	v59 = v55
	goto L7
L10:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
	v59 = v52
	goto L7
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-2)))))
	v59 = v49
	goto L7
L12:
	;
	v59 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	F__serverAssert(m, int32(_a1126), int32(_a1127), int32(399))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L40
	} else {
		goto L141
	}
L14:
	;
	m.G0 = v14 + int32(16)
	return v441
L15:
	;
	if base.Ui32(int32(32)) <= base.Ui32(l1) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v441 = l0
	goto L14
L17:
	;
	v74 = l0 - v43
	if base.Ui32(l1) < base.Ui32(v44) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if base.Ui32(int32(253)) <= base.Ui32(l1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v73 = int32(0)
	goto L17
L20:
	;
	if base.Ui32(l1) < base.Ui32(int32(65531)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v73 = int32(1)
	goto L17
L22:
	;
	v72 = int32(2)
	goto L24
L23:
	;
	v72 = int32(3)
	goto L24
L24:
	;
	v73 = v72
	goto L17
L25:
	;
	v75 = l1
	goto L27
L26:
	;
	v75 = v44
	goto L27
L27:
	;
	if v73 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v75))) = uint8(v397)
	v400 = v392 + int32(-1)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	switch v401 & int32(7) {
	case 0:
		goto L136
	case 1:
		goto L135
	case 2:
		goto L134
	case 3:
		goto L133
	case 4:
		goto L132
	default:
		v420 = v401
		goto L131
	}
L29:
	;
	v307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v307
	v315 = F_zmalloc_usable(m, l1+v84+int32(1), v14+int32(12))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L40
	} else {
		goto L105
	}
L30:
	;
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v89
	v97 = F_zrealloc_usable(m, v74, l1+v43+int32(1), v14+int32(12))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v77 = v73
	goto L33
L32:
	;
	v77 = int32(1)
	goto L33
L33:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v78 = v77
	goto L36
L35:
	;
	v78 = v73
	goto L36
L36:
	;
	if v23 == v78 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[638])))
	if base.Ui32(v23) <= base.Ui32(v78) {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v78) < base.Ui32(int32(2)) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	goto L30
L40:
	;
	return int32(0)
L41:
	;
	if v97 == int32(0) {
		v441 = v89
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v103 = v97 + v43
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v107 = v104 + (v43 ^ int32(-1))
	switch v23 {
	case 0:
		v389 = v107
		v392 = v103
		goto L28
	case 1:
		v110 = int32(255)
		goto L44
	case 2:
		goto L45
	default:
		v286 = v107
		v288 = v103
		v290 = v23
		goto L43
	}
L43:
	;
	if v290 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L44:
	;
	if base.Ui32(v107) <= base.Ui32(v110) {
		v286 = v107
		v288 = v103
		v290 = v23
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v110 = int32(65535)
	goto L44
L46:
	;
	v115 = base.B2i32(base.Ui32(v107) < base.Ui32(int32(65531)))
	if base.Ui32(v107) < base.Ui32(int32(65531)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v116 = int32(5)
	goto L49
L48:
	;
	v116 = int32(9)
	goto L49
L49:
	;
	v117 = v97 + v116
	v119 = v75 + int32(1)
	if v117 == v103 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if base.Ui32(v107) < base.Ui32(int32(65531)) {
		goto L91
	} else {
		goto L92
	}
L51:
	;
	v267 = v117
	goto L50
L52:
	;
	v123 = v119 + v117
	if base.Ui32(int32(0)-v119<<(uint(int32(1))%32)) < base.Ui32(v103-v123) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v133 = (v103 ^ v117) & int32(3)
	if base.Ui32(v103) <= base.Ui32(v117) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v130 = F___memcpy(m, v117, v103, v119)
	mBase = m.M
	v267 = v130
	goto L50
L55:
	;
	if v239 == int32(0) {
		goto L51
	} else {
		goto L87
	}
L56:
	;
	if base.Ui32(v217) <= base.Ui32(int32(3)) {
		v238 = v216
		v239 = v217
		v240 = v218
		goto L55
	} else {
		goto L83
	}
L57:
	;
	if v133 != 0 {
		v199 = v119
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v133 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v117&int32(3) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v238 = v103
	v239 = v119
	v240 = v117
	goto L55
L61:
	;
	v140 = v103
	v141 = v119
	v142 = v117
	goto L63
L62:
	;
	v216 = v103
	v217 = v119
	v218 = v117
	goto L56
L63:
	;
	if v141 == int32(0) {
		goto L51
	} else {
		goto L65
	}
L65:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v146)
	v148 = int32(1)
	v149 = v140 + v148
	v151 = v141 + int32(-1)
	v153 = v142 + v148
	if v153&int32(3) == int32(0) {
		v216 = v149
		v217 = v151
		v218 = v153
		goto L56
	} else {
		goto L66
	}
L66:
	;
	v140 = v149
	v141 = v151
	v142 = v153
	goto L63
L67:
	;
	if v199 == int32(0) {
		goto L51
	} else {
		goto L79
	}
L68:
	;
	if v123&int32(3) == int32(0) {
		v179 = v119
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if base.Ui32(v179) <= base.Ui32(int32(3)) {
		v199 = v179
		goto L67
	} else {
		goto L75
	}
L70:
	;
	v164 = v119
	goto L71
L71:
	;
	if v164 == int32(0) {
		goto L51
	} else {
		goto L73
	}
L72:
	;
	v179 = v170
	goto L69
L73:
	;
	v170 = v164 + int32(-1)
	v171 = v117 + v170
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v173)
	if v171&int32(3) != 0 {
		v164 = v170
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v186 = v179
	goto L76
L76:
	;
	v190 = v186 + int32(-4)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v103+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v190))) = v193
	if base.Ui32(int32(3)) < base.Ui32(v190) {
		v186 = v190
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v199 = v190
	goto L67
L78:
	;
	goto L77
L79:
	;
	v206 = v199
	goto L80
L80:
	;
	v210 = v206 + int32(-1)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117+v210))) = uint8(v213)
	if v210 != 0 {
		v206 = v210
		goto L80
	} else {
		goto L82
	}
L82:
	;
	goto L51
L83:
	;
	v223 = v216
	v224 = v217
	v225 = v218
	goto L84
L84:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v227
	v229 = int32(4)
	v230 = v223 + v229
	v232 = v225 + v229
	v234 = v224 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v234) {
		v223 = v230
		v224 = v234
		v225 = v232
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v238 = v230
	v239 = v234
	v240 = v232
	goto L55
L86:
	;
	goto L85
L87:
	;
	v245 = v238
	v246 = v239
	v247 = v240
	goto L88
L88:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v249)
	v251 = int32(1)
	v256 = v246 + int32(-1)
	if v256 != 0 {
		v245 = v245 + v251
		v246 = v256
		v247 = v247 + v251
		goto L88
	} else {
		goto L90
	}
L89:
	;
	goto L51
L90:
	;
	goto L89
L91:
	;
	v272 = int32(2)
	goto L93
L92:
	;
	v272 = int32(3)
	goto L93
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v267+int32(-1)))) = uint8(v272)
	if base.Ui32(int32(65530)) < base.Ui32(v107) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v286 = v282 + (v116 ^ int32(-1))
	v288 = v267
	v290 = v272
	goto L43
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267+int32(-9)))) = v75
	goto L94
L96:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v267+int32(-5)))) = uint16(v75)
	goto L94
L97:
	;
	v296 = int32(255)
	goto L99
L98:
	;
	v296 = int32(-1)
	goto L99
L99:
	;
	if v290 == int32(2) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v299 = int32(65535)
	goto L102
L101:
	;
	v299 = v296
	goto L102
L102:
	;
	if base.Ui32(v286) <= base.Ui32(v299) {
		v389 = v286
		v392 = v288
		goto L28
	} else {
		goto L103
	}
L103:
	;
	F__serverAssert(m, int32(_a1128), int32(_a1127), int32(389))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L40
	} else {
		goto L104
	}
L104:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	if v315 == int32(0) {
		v441 = v307
		goto L14
	} else {
		goto L106
	}
L106:
	;
	v319 = int32(-1)
	v320 = v84 ^ v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v324 = v78 + v319
	switch v324 {
	case 0:
		v326 = int32(255)
		goto L111
	case 1:
		goto L112
	default:
		goto L110
	}
L107:
	;
	if base.Ui32(v384) < base.Ui32(v383) {
		goto L13
	} else {
		goto L130
	}
L108:
	;
	v383 = v377
	v384 = int32(65535)
	v385 = v379
	goto L107
L109:
	;
	v366 = v315 + int32(5)
	if v335 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L110:
	;
	v350 = v315 + v84
	v352 = v75 + int32(1)
	if v352 == int32(0) {
		v356 = v350
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v327 = v321 + v320
	if base.Ui32(v327) <= base.Ui32(v326) {
		goto L110
	} else {
		goto L113
	}
L112:
	;
	v326 = int32(65535)
	goto L111
L113:
	;
	v332 = base.B2i32(base.Ui32(v327) < base.Ui32(int32(65531)))
	if base.Ui32(v327) < base.Ui32(int32(65531)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v333 = int32(-6)
	goto L116
L115:
	;
	v333 = int32(-10)
	goto L116
L116:
	;
	v335 = v75 + int32(1)
	if base.Ui32(v327) < base.Ui32(int32(65531)) {
		goto L109
	} else {
		goto L117
	}
L117:
	;
	v337 = v315 + int32(9)
	if v335 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	F_valkey_free(m, v74)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L40
	} else {
		goto L121
	}
L119:
	;
	goto L118
L120:
	;
	v340 = F__emscripten_memcpy_bulkmem(m, v337, l0, v335)
	mBase = m.M
	goto L119
L121:
	;
	v344 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+8)) = uint8(v344)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v389 = v346 + v333
	v392 = v337
	goto L28
L122:
	;
	F_valkey_free(m, v74)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L40
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	v355 = F__emscripten_memcpy_bulkmem(m, v350, l0, v352)
	mBase = m.M
	v356 = v355
	goto L123
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v356+int32(-1)))) = uint8(v78)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v363 = v362 + v320
	switch v324 {
	case 0:
		v383 = v363
		v384 = int32(255)
		v385 = v350
		goto L107
	case 1:
		v377 = v363
		v379 = v350
		goto L108
	default:
		v389 = v363
		v392 = v350
		goto L28
	}
L126:
	;
	F_valkey_free(m, v74)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L40
	} else {
		goto L129
	}
L127:
	;
	goto L126
L128:
	;
	v369 = F__emscripten_memcpy_bulkmem(m, v366, l0, v335)
	mBase = m.M
	goto L127
L129:
	;
	v373 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+4)) = uint8(v373)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v377 = v375 + v333
	v379 = v366
	goto L108
L130:
	;
	v389 = v383
	v392 = v385
	goto L28
L131:
	;
	switch v420&int32(7) + int32(-1) {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		goto L138
	case 3:
		goto L137
	default:
		v441 = v392
		goto L14
	}
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v392+int32(-17)))) = base.I64_extend_i32_u(v75)
	v420 = v401
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392+int32(-9)))) = v75
	v420 = v401
	goto L131
L134:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v392+int32(-5)))) = uint16(v75)
	v420 = v401
	goto L131
L135:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v392+int32(-3)))) = uint8(v75)
	v420 = v401
	goto L131
L136:
	;
	v405 = v75 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v400))) = uint8(v405)
	v420 = v405
	goto L131
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v392+int32(-9)))) = base.I64_extend_i32_u(v389)
	v441 = v392
	goto L14
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392+int32(-5)))) = v389
	v441 = v392
	goto L14
L139:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v392+int32(-3)))) = uint16(v389)
	v441 = v392
	goto L14
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v392+int32(-2)))) = uint8(v389)
	v441 = v392
	goto L14
L141:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
