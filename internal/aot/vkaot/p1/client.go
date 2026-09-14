package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_checkClientOutputBufferLimits(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v169 int64
	_ = v169
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v12&int32(1) != 0 {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
		v35 = v31*int32(28) + v34
		v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
		if v36 == int64(-1) {
			v46 = v35
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
			v46 = v35 + base.I32_wrap_i64(v36) + v42*int32(28)
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
		v75 = v47 + v46
	} else {
		if v12&int32(2) == int32(0) {
			if v12&int32(262144) != 0 {
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v25 == int32(0) {
				} else {
					v28 = F_isImportSlotMigrationJob(m, v25)
					mBase = m.M
				}
			}
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			v35 = v31*int32(28) + v34
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
			if v36 == int64(-1) {
				v46 = v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
				v46 = v35 + base.I32_wrap_i64(v36) + v42*int32(28)
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
			v75 = v47 + v46
		} else {
			if v12&int32(4) == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+184))
				if v50 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, _consts[475]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
					v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v55)+24)))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)+16))
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
					v64 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
					v67 = int32(44)
					v75 = base.I32_wrap_i64(v56+v57-v60) + base.I32_wrap_i64(v63-v64)*v67 + v67
				} else {
					v75 = int32(0)
				}
			} else {
				if v12&int32(262144) != 0 {
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v25 == int32(0) {
					} else {
						v28 = F_isImportSlotMigrationJob(m, v25)
						mBase = m.M
					}
				}
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				v35 = v31*int32(28) + v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+384))
				if v36 == int64(-1) {
					v46 = v35
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+376))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
					v46 = v35 + base.I32_wrap_i64(v36) + v42*int32(28)
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
				v75 = v47 + v46
			}
		}
	}
	if base.Ui32(v75) < base.Ui32(int32(1025)) {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		if v90&int32(1) != 0 {
			v124 = int32(0)
			v127 = v124
			v128 = v124
		} else {
			if v90&int32(2) == int32(0) {
				if v90&int32(262144) != 0 {
					v127 = int32(0)
					v128 = int32(2)
				} else {
					v109 = int32(0)
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v111 == v109 {
						v127 = v109
						v128 = v109
					} else {
						v114 = int32(1)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
						if base.B2i32(v116 == v114) == int32(0) {
							v127 = v114
							v128 = v114
						} else {
							v124 = int32(0)
							v127 = v124
							v128 = v124
						}
					}
				}
			} else {
				v97 = int32(1)
				if v90&int32(4) == int32(0) {
					v127 = v97
					v128 = v97
				} else {
					if v90&int32(262144) != 0 {
						v127 = int32(0)
						v128 = int32(2)
					} else {
						v109 = int32(0)
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v111 == v109 {
							v127 = v109
							v128 = v109
						} else {
							v114 = int32(1)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
							if base.B2i32(v116 == v114) == int32(0) {
								v127 = v114
								v128 = v114
							} else {
								v124 = int32(0)
								v127 = v124
								v128 = v124
							}
						}
					}
				}
			}
		}
		v131 = v128 * int32(24)
		v133 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[258])))
		v137 = *(*int64)(unsafe.Add(mBase, _consts[322]))
		v138 = base.I32_wrap_i64(v137)
		v139 = base.I32_wrap_i64(v133)
		if v133&int64(4294967295) < v137 {
			v143 = v138
		} else {
			v143 = v139
		}
		if v139 != 0 {
			v144 = v143
		} else {
			v144 = v139
		}
		if v127 != 0 {
			v145 = v144
		} else {
			v145 = v139
		}
		v147 = base.B2i32(v133 != int64(0)) & base.B2i32(base.Ui32(v145) <= base.Ui32(v75))
		v148 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[259])))
		if v148 == int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
			return v147 | int32(0)
		} else {
			v151 = base.I32_wrap_i64(v148)
			if v148&int64(4294967295) < v137 {
				v155 = v138
			} else {
				v155 = v151
			}
			if v151 != 0 {
				v156 = v155
			} else {
				v156 = v151
			}
			if v127 != 0 {
				v157 = v156
			} else {
				v157 = v151
			}
			if base.Ui32(v75) < base.Ui32(v157) {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
				return v147 | int32(0)
			} else {
				v160 = *(*int64)(unsafe.Add(mBase, _consts[109]))
				v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
				if v161 != int64(0) {
					v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
					v169 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[260])))
					return v147 | base.B2i32(v169 < v160-v165)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v160
					return v147 | int32(0)
				}
			}
		}
	} else {
		v79 = *(*int32)(unsafe.Add(mBase, _consts[22]))
		v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
		if v80&int32(6) == int32(4) {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			if v90&int32(1) != 0 {
				v124 = int32(0)
				v127 = v124
				v128 = v124
			} else {
				if v90&int32(2) == int32(0) {
					if v90&int32(262144) != 0 {
						v127 = int32(0)
						v128 = int32(2)
					} else {
						v109 = int32(0)
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
						if v111 == v109 {
							v127 = v109
							v128 = v109
						} else {
							v114 = int32(1)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
							if base.B2i32(v116 == v114) == int32(0) {
								v127 = v114
								v128 = v114
							} else {
								v124 = int32(0)
								v127 = v124
								v128 = v124
							}
						}
					}
				} else {
					v97 = int32(1)
					if v90&int32(4) == int32(0) {
						v127 = v97
						v128 = v97
					} else {
						if v90&int32(262144) != 0 {
							v127 = int32(0)
							v128 = int32(2)
						} else {
							v109 = int32(0)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v111 == v109 {
								v127 = v109
								v128 = v109
							} else {
								v114 = int32(1)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								if base.B2i32(v116 == v114) == int32(0) {
									v127 = v114
									v128 = v114
								} else {
									v124 = int32(0)
									v127 = v124
									v128 = v124
								}
							}
						}
					}
				}
			}
			v131 = v128 * int32(24)
			v133 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[258])))
			v137 = *(*int64)(unsafe.Add(mBase, _consts[322]))
			v138 = base.I32_wrap_i64(v137)
			v139 = base.I32_wrap_i64(v133)
			if v133&int64(4294967295) < v137 {
				v143 = v138
			} else {
				v143 = v139
			}
			if v139 != 0 {
				v144 = v143
			} else {
				v144 = v139
			}
			if v127 != 0 {
				v145 = v144
			} else {
				v145 = v139
			}
			v147 = base.B2i32(v133 != int64(0)) & base.B2i32(base.Ui32(v145) <= base.Ui32(v75))
			v148 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[259])))
			if v148 == int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
				return v147 | int32(0)
			} else {
				v151 = base.I32_wrap_i64(v148)
				if v148&int64(4294967295) < v137 {
					v155 = v138
				} else {
					v155 = v151
				}
				if v151 != 0 {
					v156 = v155
				} else {
					v156 = v151
				}
				if v127 != 0 {
					v157 = v156
				} else {
					v157 = v151
				}
				if base.Ui32(v75) < base.Ui32(v157) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
					return v147 | int32(0)
				} else {
					v160 = *(*int64)(unsafe.Add(mBase, _consts[109]))
					v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
					if v161 != int64(0) {
						v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
						v169 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[260])))
						return v147 | base.B2i32(v169 < v160-v165)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v160
						return v147 | int32(0)
					}
				}
			}
		} else {
			v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+206)))
			if v85&int32(384) != 0 {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v90&int32(1) != 0 {
					v124 = int32(0)
					v127 = v124
					v128 = v124
				} else {
					if v90&int32(2) == int32(0) {
						if v90&int32(262144) != 0 {
							v127 = int32(0)
							v128 = int32(2)
						} else {
							v109 = int32(0)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
							if v111 == v109 {
								v127 = v109
								v128 = v109
							} else {
								v114 = int32(1)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								if base.B2i32(v116 == v114) == int32(0) {
									v127 = v114
									v128 = v114
								} else {
									v124 = int32(0)
									v127 = v124
									v128 = v124
								}
							}
						}
					} else {
						v97 = int32(1)
						if v90&int32(4) == int32(0) {
							v127 = v97
							v128 = v97
						} else {
							if v90&int32(262144) != 0 {
								v127 = int32(0)
								v128 = int32(2)
							} else {
								v109 = int32(0)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
								if v111 == v109 {
									v127 = v109
									v128 = v109
								} else {
									v114 = int32(1)
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
									if base.B2i32(v116 == v114) == int32(0) {
										v127 = v114
										v128 = v114
									} else {
										v124 = int32(0)
										v127 = v124
										v128 = v124
									}
								}
							}
						}
					}
				}
				v131 = v128 * int32(24)
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[258])))
				v137 = *(*int64)(unsafe.Add(mBase, _consts[322]))
				v138 = base.I32_wrap_i64(v137)
				v139 = base.I32_wrap_i64(v133)
				if v133&int64(4294967295) < v137 {
					v143 = v138
				} else {
					v143 = v139
				}
				if v139 != 0 {
					v144 = v143
				} else {
					v144 = v139
				}
				if v127 != 0 {
					v145 = v144
				} else {
					v145 = v139
				}
				v147 = base.B2i32(v133 != int64(0)) & base.B2i32(base.Ui32(v145) <= base.Ui32(v75))
				v148 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[259])))
				if v148 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
					return v147 | int32(0)
				} else {
					v151 = base.I32_wrap_i64(v148)
					if v148&int64(4294967295) < v137 {
						v155 = v138
					} else {
						v155 = v151
					}
					if v151 != 0 {
						v156 = v155
					} else {
						v156 = v151
					}
					if v127 != 0 {
						v157 = v156
					} else {
						v157 = v151
					}
					if base.Ui32(v75) < base.Ui32(v157) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
						return v147 | int32(0)
					} else {
						v160 = *(*int64)(unsafe.Add(mBase, _consts[109]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
						if v161 != int64(0) {
							v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
							v169 = *(*int64)(unsafe.Add(mBase, uint32(v131)+uint32(_consts[260])))
							return v147 | base.B2i32(v169 < v160-v165)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v160
							return v147 | int32(0)
						}
					}
				}
			} else {
				return int32(1)
			}
		}
	}
}
func F_clientAcceptHandler(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 == int32(3) {
		v55 = *(*int32)(unsafe.Add(mBase, _consts[489]))
		if v55 == int32(0) {
			v105 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v107 == v105 {
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
				F_sdsfree(m, v150)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return
				} else {
					v153 = int32(_a20)
					v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
					*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
					F_moduleFireServerEvent(m, int64(4), int32(0), v9)
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			} else {
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+128))
				if v110 == int32(0) {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
					F_sdsfree(m, v150)
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return
					} else {
						v153 = int32(_a20)
						v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
						*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
						F_moduleFireServerEvent(m, int64(4), int32(0), v9)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				} else {
					v115 = m.T0[v110].(func(*base.Module, int32, int32) int32)(m, l0, v7+int32(28))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						if v115 == int32(0) {
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
							if v139 == int32(0) {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								F_sdsfree(m, v150)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(_a20)
									v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
									*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
									F_moduleFireServerEvent(m, int64(4), int32(0), v9)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							} else {
								v143 = int32(0)
								F_addACLLogEntry(m, v9, int32(6), v143, v143, v139, v143)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return
								} else {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									F_sdsfree(m, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a20)
										v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
										*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
										F_moduleFireServerEvent(m, int64(4), int32(0), v9)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+328)) = v115
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+204)) = v120 | int32(25165824)
							F_moduleNotifyUserChanged(m, v9)
							mBase = m.M
							v126 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							if int32(1) < v126 {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								F_sdsfree(m, v150)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(_a20)
									v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
									*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
									F_moduleFireServerEvent(m, int64(4), int32(0), v9)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, _consts[481]))
								if v131 != 0 {
									v133 = int32(_a1645)
								} else {
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
									v133 = v132
								}
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v133
								F__serverLog(m, int32(1), int32(_a1662), v7)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									F_sdsfree(m, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a20)
										v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
										*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
										F_moduleFireServerEvent(m, int64(4), int32(0), v9)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
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
		} else {
			v58 = int32(0)
			v59 = *(*int32)(unsafe.Add(mBase, _consts[22]))
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
			if v60&int32(4) == v58 {
				v105 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v107 == v105 {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
					F_sdsfree(m, v150)
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return
					} else {
						v153 = int32(_a20)
						v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
						*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
						F_moduleFireServerEvent(m, int64(4), int32(0), v9)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+128))
					if v110 == int32(0) {
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
						F_sdsfree(m, v150)
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							v153 = int32(_a20)
							v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
							*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
							F_moduleFireServerEvent(m, int64(4), int32(0), v9)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					} else {
						v115 = m.T0[v110].(func(*base.Module, int32, int32) int32)(m, l0, v7+int32(28))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							if v115 == int32(0) {
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								if v139 == int32(0) {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									F_sdsfree(m, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a20)
										v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
										*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
										F_moduleFireServerEvent(m, int64(4), int32(0), v9)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								} else {
									v143 = int32(0)
									F_addACLLogEntry(m, v9, int32(6), v143, v143, v139, v143)
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										F_sdsfree(m, v150)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											v153 = int32(_a20)
											v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
											*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
											F_moduleFireServerEvent(m, int64(4), int32(0), v9)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												m.G0 = v7 + int32(32)
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+328)) = v115
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+204)) = v120 | int32(25165824)
								F_moduleNotifyUserChanged(m, v9)
								mBase = m.M
								v126 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								if int32(1) < v126 {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									F_sdsfree(m, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a20)
										v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
										*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
										F_moduleFireServerEvent(m, int64(4), int32(0), v9)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, _consts[481]))
									if v131 != 0 {
										v133 = int32(_a1645)
									} else {
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
										v133 = v132
									}
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v133
									F__serverLog(m, int32(1), int32(_a1662), v7)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										F_sdsfree(m, v150)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return
										} else {
											v153 = int32(_a20)
											v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
											*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
											F_moduleFireServerEvent(m, int64(4), int32(0), v9)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
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
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
				if v66 == int32(0) {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+68))
					v78 = m.T0[v77].(func(*base.Module, int32, int32, int32) int32)(m, v73, int32(_a1663), int32(1030))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v80 = int32(_a20)
						v82 = *(*int64)(unsafe.Add(mBase, _consts[491]))
						*(*int64)(unsafe.Add(mBase, _consts[491])) = v82 + int64(1)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
						if v86&int32(1280) != 0 {
							m.G0 = v7 + int32(32)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v86 | int32(1024)
							v93 = *(*int32)(unsafe.Add(mBase, _consts[113]))
							if v93 == int32(0) {
								v101 = *(*int32)(unsafe.Add(mBase, _consts[483]))
								v102 = F_listAddNodeTail(m, v101, v9)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, _consts[483]))
								v98 = F_listSearchKey(m, v97, v9)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									if v98 != 0 {
										F__serverAssertWithInfo(m, v9, int32(0), int32(_a1641), int32(_a1630), int32(2277))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return
										} else {
											F_abort(m)
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, _consts[483]))
										v102 = F_listAddNodeTail(m, v101, v9)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v69 = m.T0[v66].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						if v69 == int32(1) {
							v105 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v107 == v105 {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
								F_sdsfree(m, v150)
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return
								} else {
									v153 = int32(_a20)
									v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
									*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
									F_moduleFireServerEvent(m, int64(4), int32(0), v9)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+128))
								if v110 == int32(0) {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									F_sdsfree(m, v150)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return
									} else {
										v153 = int32(_a20)
										v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
										*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
										F_moduleFireServerEvent(m, int64(4), int32(0), v9)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								} else {
									v115 = m.T0[v110].(func(*base.Module, int32, int32) int32)(m, l0, v7+int32(28))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										if v115 == int32(0) {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
											if v139 == int32(0) {
												v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
												F_sdsfree(m, v150)
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return
												} else {
													v153 = int32(_a20)
													v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
													*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
													F_moduleFireServerEvent(m, int64(4), int32(0), v9)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														m.G0 = v7 + int32(32)
														return
													}
												}
											} else {
												v143 = int32(0)
												F_addACLLogEntry(m, v9, int32(6), v143, v143, v139, v143)
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
													F_sdsfree(m, v150)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return
													} else {
														v153 = int32(_a20)
														v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
														*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
														F_moduleFireServerEvent(m, int64(4), int32(0), v9)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															m.G0 = v7 + int32(32)
															return
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+328)) = v115
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+204)) = v120 | int32(25165824)
											F_moduleNotifyUserChanged(m, v9)
											mBase = m.M
											v126 = *(*int32)(unsafe.Add(mBase, _consts[28]))
											if int32(1) < v126 {
												v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
												F_sdsfree(m, v150)
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return
												} else {
													v153 = int32(_a20)
													v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
													*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
													F_moduleFireServerEvent(m, int64(4), int32(0), v9)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														m.G0 = v7 + int32(32)
														return
													}
												}
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, _consts[481]))
												if v131 != 0 {
													v133 = int32(_a1645)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
													v133 = v132
												}
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v133
												F__serverLog(m, int32(1), int32(_a1662), v7)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
													F_sdsfree(m, v150)
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return
													} else {
														v153 = int32(_a20)
														v155 = *(*int64)(unsafe.Add(mBase, _consts[490]))
														*(*int64)(unsafe.Add(mBase, _consts[490])) = v155 + int64(1)
														F_moduleFireServerEvent(m, int64(4), int32(0), v9)
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
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
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+68))
							v78 = m.T0[v77].(func(*base.Module, int32, int32, int32) int32)(m, v73, int32(_a1663), int32(1030))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								v80 = int32(_a20)
								v82 = *(*int64)(unsafe.Add(mBase, _consts[491]))
								*(*int64)(unsafe.Add(mBase, _consts[491])) = v82 + int64(1)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
								if v86&int32(1280) != 0 {
									m.G0 = v7 + int32(32)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v86 | int32(1024)
									v93 = *(*int32)(unsafe.Add(mBase, _consts[113]))
									if v93 == int32(0) {
										v101 = *(*int32)(unsafe.Add(mBase, _consts[483]))
										v102 = F_listAddNodeTail(m, v101, v9)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _consts[483]))
										v98 = F_listSearchKey(m, v97, v9)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											if v98 != 0 {
												F__serverAssertWithInfo(m, v9, int32(0), int32(_a1641), int32(_a1630), int32(2277))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													F_abort(m)
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, _consts[483]))
												v102 = F_listAddNodeTail(m, v101, v9)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
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
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		if int32(3) < v14 {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
			if v36&int32(1280) != 0 {
				m.G0 = v7 + int32(32)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v36 | int32(1024)
				v43 = *(*int32)(unsafe.Add(mBase, _consts[113]))
				if v43 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[483]))
					v52 = F_listAddNodeTail(m, v51, v9)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[483]))
					v48 = F_listSearchKey(m, v47, v9)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						if v48 != 0 {
							F__serverAssertWithInfo(m, v9, int32(0), int32(_a1641), int32(_a1630), int32(2277))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_abort(m)
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, _consts[483]))
							v52 = F_listAddNodeTail(m, v51, v9)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
			v19 = m.T0[v18].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = F_getClientPeerId(m, v9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = F_getClientSockname(m, v9)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v19
						F__serverLog(m, int32(3), int32(_a1664), v7+int32(16))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+200))
							if v36&int32(1280) != 0 {
								m.G0 = v7 + int32(32)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v36 | int32(1024)
								v43 = *(*int32)(unsafe.Add(mBase, _consts[113]))
								if v43 == int32(0) {
									v51 = *(*int32)(unsafe.Add(mBase, _consts[483]))
									v52 = F_listAddNodeTail(m, v51, v9)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, _consts[483]))
									v48 = F_listSearchKey(m, v47, v9)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										if v48 != 0 {
											F__serverAssertWithInfo(m, v9, int32(0), int32(_a1641), int32(_a1630), int32(2277))
											mBase = m.M
											v173 = m.ExcPending
											if v173 != 0 {
												return
											} else {
												F_abort(m)
												mBase = m.M
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, _consts[483]))
											v52 = F_listAddNodeTail(m, v51, v9)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
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
func F_clientCapaCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 < int32(3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v8 = int32(2)
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v8<<(uint(int32(2))%32))))
	v14 = F_objectGetVal(m, v13)
	mBase = m.M
	v15 = int32(_a1736)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L1
L5:
	;
	v59 = v8 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v59 < v60 {
		v8 = v59
		goto L3
	} else {
		goto L19
	}
L6:
	;
	if v50-v52 != 0 {
		goto L5
	} else {
		goto L18
	}
L7:
	;
	v50 = F_tolower(m, v46)
	mBase = m.M
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v52 = F_tolower(m, v51)
	mBase = m.M
	goto L6
L8:
	;
	v20 = v14
	v21 = v15
	v22 = v18
	goto L11
L9:
	;
	v46 = int32(0)
	v47 = v15
	goto L7
L10:
	;
	v46 = v43 & int32(255)
	v47 = v42
	goto L7
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v24 == int32(0) {
		v42 = v21
		v43 = v22
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v42 = v36
	v43 = int32(0)
	goto L10
L13:
	;
	v28 = v22 & int32(255)
	if v28 == v24 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = int32(1)
	v36 = v21 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v37 != 0 {
		v20 = v20 + v35
		v21 = v36
		v22 = v37
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v30 = F_tolower(m, v28)
	mBase = m.M
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v32 = F_tolower(m, v31)
	mBase = m.M
	if v30 == v32 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v42 = v21
	v43 = v34
	goto L10
L17:
	;
	goto L12
L18:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
	v56 = v54 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)) = uint8(v56)
	goto L5
L19:
	;
	goto L4
L20:
	;
	return
L21:
	;
	return
}
func F_clientHasPendingReplies(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v5&int32(1) != 0 {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
		if v26 != 0 {
			v50 = int32(1)
			return v50
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
			return base.B2i32(v28 != int32(0))
		}
	} else {
		if v5&int32(2) == int32(0) {
			if v5&int32(262144) != 0 {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v18 == int32(0) {
				} else {
				}
			}
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
			if v26 != 0 {
				v50 = int32(1)
				return v50
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
				return base.B2i32(v28 != int32(0))
			}
		} else {
			if v5&int32(4) == int32(0) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				if v32 != 0 {
					F__serverAssert(m, int32(_a1631), int32(_a1630), int32(1798))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
					if v34 != 0 {
						F__serverAssert(m, int32(_a1631), int32(_a1630), int32(1798))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_abort(m)
							mBase = m.M
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v35 = int32(0)
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+184))
						if v37 == v35 {
							v50 = v35
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[475]))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
							if v42 != v37 {
								v50 = int32(1)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)+188))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
								if v44 == v46 {
									v50 = v35
								} else {
									v50 = int32(1)
								}
							}
						}
						return v50
					}
				}
			} else {
				if v5&int32(262144) != 0 {
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v18 == int32(0) {
					} else {
					}
				}
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
				if v26 != 0 {
					v50 = int32(1)
					return v50
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
					return base.B2i32(v28 != int32(0))
				}
			}
		}
	}
}
func F_clientHashtableTypeMetadataSize(m *base.Module) int32 {
	return int32(4)
}
func F_clientImportSourceCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v7 = *(*int32)(unsafe.Add(mBase, _consts[505]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = F_objectGetVal(m, v62)
	mBase = m.M
	v64 = int32(_a50)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = F_objectGetVal(m, v9)
	mBase = m.M
	v11 = int32(_a1735)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v46-v48 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	v46 = F_tolower(m, v42)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	goto L3
L5:
	;
	v16 = v10
	v17 = v11
	v18 = v14
	goto L8
L6:
	;
	v42 = int32(0)
	v43 = v11
	goto L4
L7:
	;
	v42 = v39 & int32(255)
	v43 = v38
	goto L4
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v20 == int32(0) {
		v38 = v17
		v39 = v18
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v38 = v32
	v39 = int32(0)
	goto L7
L10:
	;
	v24 = v18 & int32(255)
	if v24 == v20 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(1)
	v32 = v17 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v33 != 0 {
		v16 = v16 + v31
		v17 = v32
		v18 = v33
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v26 = F_tolower(m, v24)
	mBase = m.M
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v28 = F_tolower(m, v27)
	mBase = m.M
	if v26 == v28 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v38 = v17
	v39 = v30
	goto L7
L14:
	;
	goto L9
L15:
	;
	F_addReplyErrorLength(m, l0, int32(_a1753), int32(28))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	F_afterErrorReply(m, l0, int32(_a1753), int32(28), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	return
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v113 = F_objectGetVal(m, v112)
	mBase = m.M
	v114 = int32(_a1735)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 != 0 {
		goto L37
	} else {
		goto L38
	}
L20:
	;
	if v99-v101 != 0 {
		goto L19
	} else {
		goto L32
	}
L21:
	;
	v99 = F_tolower(m, v95)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v101 = F_tolower(m, v100)
	mBase = m.M
	goto L20
L22:
	;
	v69 = v63
	v70 = v64
	v71 = v67
	goto L25
L23:
	;
	v95 = int32(0)
	v96 = v64
	goto L21
L24:
	;
	v95 = v92 & int32(255)
	v96 = v91
	goto L21
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 == int32(0) {
		v91 = v70
		v92 = v71
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v91 = v85
	v92 = int32(0)
	goto L24
L27:
	;
	v77 = v71 & int32(255)
	if v77 == v73 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v84 = int32(1)
	v85 = v70 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v86 != 0 {
		v69 = v69 + v84
		v70 = v85
		v71 = v86
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v79 = F_tolower(m, v77)
	mBase = m.M
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v81 = F_tolower(m, v80)
	mBase = m.M
	if v79 == v81 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v91 = v70
	v92 = v83
	goto L24
L31:
	;
	goto L26
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v103 | int32(536870912)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	return
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReply(m, l0, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L16
	} else {
		goto L49
	}
L35:
	;
	if v149-v151 != 0 {
		goto L34
	} else {
		goto L47
	}
L36:
	;
	v149 = F_tolower(m, v145)
	mBase = m.M
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v151 = F_tolower(m, v150)
	mBase = m.M
	goto L35
L37:
	;
	v119 = v113
	v120 = v114
	v121 = v117
	goto L40
L38:
	;
	v145 = int32(0)
	v146 = v114
	goto L36
L39:
	;
	v145 = v142 & int32(255)
	v146 = v141
	goto L36
L40:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v123 == int32(0) {
		v141 = v120
		v142 = v121
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v141 = v135
	v142 = int32(0)
	goto L39
L42:
	;
	v127 = v121 & int32(255)
	if v127 == v123 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v134 = int32(1)
	v135 = v120 + v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v136 != 0 {
		v119 = v119 + v134
		v120 = v135
		v121 = v136
		goto L40
	} else {
		goto L46
	}
L44:
	;
	v129 = F_tolower(m, v127)
	mBase = m.M
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v131 = F_tolower(m, v130)
	mBase = m.M
	if v129 == v131 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v141 = v120
	v142 = v133
	goto L39
L46:
	;
	goto L41
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v153 & int32(-536870913)
	v158 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	return
L49:
	;
	v165 = F_objectGetVal(m, v162)
	mBase = m.M
	v167 = F_objectGetVal(m, v162)
	mBase = m.M
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+int32(-1)))))
	switch v170 & int32(7) {
	case 0:
		goto L55
	case 1:
		goto L54
	case 2:
		goto L53
	case 3:
		goto L52
	case 4:
		goto L51
	default:
		v187 = int32(0)
		goto L50
	}
L50:
	;
	F_afterErrorReply(m, l0, v165, v187+int32(-2), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L16
	} else {
		goto L56
	}
L51:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v167+int32(-17))))
	v187 = v186
	goto L50
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v167+int32(-9))))
	v187 = v183
	goto L50
L53:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167+int32(-5)))))
	v187 = v180
	goto L50
L54:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+int32(-3)))))
	v187 = v177
	goto L50
L55:
	;
	v187 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
	goto L50
L56:
	;
	return
}
func F_clientMatchesCapaFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	v33 = int32(0)
	goto L1
L1:
	;
	switch v21 & int32(7) {
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
		v42 = int32(0)
		goto L3
	}
L2:
	;
	return v43
L3:
	;
	v43 = base.B2i32(base.Ui32(v42) <= base.Ui32(v33))
	if base.Ui32(v42) <= base.Ui32(v33) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v42 = v41
	goto L3
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v42 = v40
	goto L3
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v42 = v39
	goto L3
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v42 = v38
	goto L3
L8:
	;
	v42 = int32(base.Ui32(v21) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	goto L2
L10:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v33))))
	if v45 != int32(114) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+227)))
	if v50&v48 != 0 {
		v33 = v33 + v48
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_clientMatchesFlagFilter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v29 = int32(0)
	goto L1
L1:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v33 & int32(7) {
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
		v42 = int32(0)
		goto L3
	}
L3:
	;
	v43 = base.B2i32(base.Ui32(v42) <= base.Ui32(v29))
	if base.Ui32(v42) <= base.Ui32(v29) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v42 = v41
	goto L3
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v42 = v40
	goto L3
L6:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v42 = v39
	goto L3
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v42 = v38
	goto L3
L8:
	;
	v42 = int32(base.Ui32(v33) >> (uint(int32(3)) % 32))
	goto L3
L9:
	;
	v29 = v29 + int32(1)
	goto L1
L10:
	;
	return v43
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v29))))
	switch v45 + int32(-65) {
	case 0:
		goto L21
	case 1:
		goto L25
	default:
		goto L10
	case 4:
		goto L14
	case 8:
		goto L16
	case 12:
		goto L31
	case 13:
		goto L13
	case 14:
		goto L12
	case 15:
		goto L30
	case 17:
		goto L26
	case 18:
		goto L32
	case 19:
		goto L17
	case 20:
		goto L20
	case 33:
		goto L28
	case 34:
		goto L23
	case 35:
		goto L24
	case 36:
		goto L18
	case 40:
		goto L15
	case 49:
		goto L19
	case 51:
		goto L27
	case 52:
		goto L22
	case 55:
		goto L29
	}
L12:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v124 = int32(6)
	if v123&v124 == v124 {
		goto L9
	} else {
		goto L59
	}
L13:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v113&int32(262171) != 0 {
		goto L10
	} else {
		goto L56
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v105 == int32(0) {
		goto L10
	} else {
		goto L53
	}
L15:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v99 == int32(0) {
		goto L10
	} else {
		goto L50
	}
L16:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+207)))
	if v96&int32(32) != 0 {
		goto L9
	} else {
		goto L49
	}
L17:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+206)))
	if v93&int32(1) != 0 {
		goto L9
	} else {
		goto L48
	}
L18:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v90&int32(64) != 0 {
		goto L9
	} else {
		goto L47
	}
L19:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v87&int32(2) != 0 {
		goto L9
	} else {
		goto L46
	}
L20:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v84&int32(8) != 0 {
		goto L9
	} else {
		goto L45
	}
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v81&int32(4) != 0 {
		goto L9
	} else {
		goto L44
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v78&int32(128) != 0 {
		goto L9
	} else {
		goto L43
	}
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v75&int32(64) != 0 {
		goto L9
	} else {
		goto L42
	}
L24:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v72&int32(32) != 0 {
		goto L9
	} else {
		goto L41
	}
L25:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v69&int32(16) != 0 {
		goto L9
	} else {
		goto L40
	}
L26:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v66&int32(8) != 0 {
		goto L9
	} else {
		goto L39
	}
L27:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v63&int32(4) != 0 {
		goto L9
	} else {
		goto L38
	}
L28:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v60&int32(16) != 0 {
		goto L9
	} else {
		goto L37
	}
L29:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v57&int32(8) != 0 {
		goto L9
	} else {
		goto L36
	}
L30:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+202)))
	if v54&int32(4) != 0 {
		goto L9
	} else {
		goto L35
	}
L31:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v51&int32(1) != 0 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v48&int32(2) != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	goto L10
L35:
	;
	goto L10
L36:
	;
	goto L10
L37:
	;
	goto L10
L38:
	;
	goto L10
L39:
	;
	goto L10
L40:
	;
	goto L10
L41:
	;
	goto L10
L42:
	;
	goto L10
L43:
	;
	goto L10
L44:
	;
	goto L10
L45:
	;
	goto L10
L46:
	;
	goto L10
L47:
	;
	goto L10
L48:
	;
	goto L10
L49:
	;
	goto L10
L50:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	goto L51
L51:
	;
	if v102 == int32(1) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	goto L10
L53:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	goto L54
L54:
	;
	if base.B2i32(v108 == int32(1)) == int32(0) {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	goto L10
L56:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v116&int32(536952860)|v113&int32(134368) != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v122 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	goto L9
L59:
	;
	goto L10
}
func F_clientNoEvictCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v9 = int32(_a50)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v61 = F_objectGetVal(m, v60)
	mBase = m.M
	v62 = int32(_a1735)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	if v44-v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v44 = F_tolower(m, v40)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v46 = F_tolower(m, v45)
	mBase = m.M
	goto L2
L4:
	;
	v14 = v8
	v15 = v9
	v16 = v12
	goto L7
L5:
	;
	v40 = int32(0)
	v41 = v9
	goto L3
L6:
	;
	v40 = v37 & int32(255)
	v41 = v36
	goto L3
L7:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v18 == int32(0) {
		v36 = v15
		v37 = v16
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v36 = v30
	v37 = int32(0)
	goto L6
L9:
	;
	v22 = v16 & int32(255)
	if v22 == v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = int32(1)
	v30 = v15 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v31 != 0 {
		v14 = v14 + v29
		v15 = v30
		v16 = v31
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v24 = F_tolower(m, v22)
	mBase = m.M
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v26 = F_tolower(m, v25)
	mBase = m.M
	if v24 == v26 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v36 = v15
	v37 = v28
	goto L6
L13:
	;
	goto L8
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v48 | int32(16384)
	F_removeClientFromMemUsageBucket(m, l0, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	return
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReply(m, l0, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L34
	}
L19:
	;
	if v97-v99 != 0 {
		goto L18
	} else {
		goto L31
	}
L20:
	;
	v97 = F_tolower(m, v93)
	mBase = m.M
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v99 = F_tolower(m, v98)
	mBase = m.M
	goto L19
L21:
	;
	v67 = v61
	v68 = v62
	v69 = v65
	goto L24
L22:
	;
	v93 = int32(0)
	v94 = v62
	goto L20
L23:
	;
	v93 = v90 & int32(255)
	v94 = v89
	goto L20
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v71 == int32(0) {
		v89 = v68
		v90 = v69
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v89 = v83
	v90 = int32(0)
	goto L23
L26:
	;
	v75 = v69 & int32(255)
	if v75 == v71 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v82 = int32(1)
	v83 = v68 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v84 != 0 {
		v67 = v67 + v82
		v68 = v83
		v69 = v84
		goto L24
	} else {
		goto L30
	}
L28:
	;
	v77 = F_tolower(m, v75)
	mBase = m.M
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v79 = F_tolower(m, v78)
	mBase = m.M
	if v77 == v79 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v89 = v68
	v90 = v81
	goto L23
L30:
	;
	goto L25
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v101 & int32(-16385)
	v105 = F_updateClientMemUsageAndBucket(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	return
L34:
	;
	v115 = F_objectGetVal(m, v112)
	mBase = m.M
	v117 = F_objectGetVal(m, v112)
	mBase = m.M
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(-1)))))
	switch v120 & int32(7) {
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
		v137 = int32(0)
		goto L35
	}
L35:
	;
	F_afterErrorReply(m, l0, v115, v137+int32(-2), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L41
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-17))))
	v137 = v136
	goto L35
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(-9))))
	v137 = v133
	goto L35
L38:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+int32(-5)))))
	v137 = v130
	goto L35
L39:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+int32(-3)))))
	v137 = v127
	goto L35
L40:
	;
	v137 = int32(base.Ui32(v120) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	return
}
func F_clientNoTouchCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = F_objectGetVal(m, v7)
	mBase = m.M
	v9 = int32(_a50)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	v59 = int32(_a1735)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	if v44-v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v44 = F_tolower(m, v40)
	mBase = m.M
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v46 = F_tolower(m, v45)
	mBase = m.M
	goto L2
L4:
	;
	v14 = v8
	v15 = v9
	v16 = v12
	goto L7
L5:
	;
	v40 = int32(0)
	v41 = v9
	goto L3
L6:
	;
	v40 = v37 & int32(255)
	v41 = v36
	goto L3
L7:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v18 == int32(0) {
		v36 = v15
		v37 = v16
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v36 = v30
	v37 = int32(0)
	goto L6
L9:
	;
	v22 = v16 & int32(255)
	if v22 == v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = int32(1)
	v30 = v15 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v31 != 0 {
		v14 = v14 + v29
		v15 = v30
		v16 = v31
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v24 = F_tolower(m, v22)
	mBase = m.M
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v26 = F_tolower(m, v25)
	mBase = m.M
	if v24 == v26 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v36 = v15
	v37 = v28
	goto L6
L13:
	;
	goto L8
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v48 | int32(65536)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	return
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReply(m, l0, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L32
	}
L18:
	;
	if v94-v96 != 0 {
		goto L17
	} else {
		goto L30
	}
L19:
	;
	v94 = F_tolower(m, v90)
	mBase = m.M
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v96 = F_tolower(m, v95)
	mBase = m.M
	goto L18
L20:
	;
	v64 = v58
	v65 = v59
	v66 = v62
	goto L23
L21:
	;
	v90 = int32(0)
	v91 = v59
	goto L19
L22:
	;
	v90 = v87 & int32(255)
	v91 = v86
	goto L19
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == int32(0) {
		v86 = v65
		v87 = v66
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v86 = v80
	v87 = int32(0)
	goto L22
L25:
	;
	v72 = v66 & int32(255)
	if v72 == v68 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(1)
	v80 = v65 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)))
	if v81 != 0 {
		v64 = v64 + v79
		v65 = v80
		v66 = v81
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v74 = F_tolower(m, v72)
	mBase = m.M
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v76 = F_tolower(m, v75)
	mBase = m.M
	if v74 == v76 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v86 = v65
	v87 = v78
	goto L22
L29:
	;
	goto L24
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v98 & int32(-65537)
	v103 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	return
L32:
	;
	v110 = F_objectGetVal(m, v107)
	mBase = m.M
	v112 = F_objectGetVal(m, v107)
	mBase = m.M
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-1)))))
	switch v115 & int32(7) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L34
	default:
		v132 = int32(0)
		goto L33
	}
L33:
	;
	F_afterErrorReply(m, l0, v110, v132+int32(-2), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L15
	} else {
		goto L39
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-17))))
	v132 = v131
	goto L33
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v112+int32(-9))))
	v132 = v128
	goto L33
L36:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+int32(-5)))))
	v132 = v125
	goto L33
L37:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+int32(-3)))))
	v132 = v122
	goto L33
L38:
	;
	v132 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	return
}
func F_clientSetNameCommand(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v14 = F_clientSetName(m, l0, v9, v6+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return
L2:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L40
	}
L3:
	;
	return
L4:
	;
	if v14 != int32(-1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v18&int32(3) == int32(0) {
		v40 = v18
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_addReplyErrorLength(m, l0, v18, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L22
	}
L7:
	;
	v73 = v65 - v18
	goto L6
L8:
	;
	v44 = v40
	goto L16
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = v18
	goto L12
L11:
	;
	v73 = v18 - v18
	goto L6
L12:
	;
	v33 = v29 + int32(1)
	if v33&int32(3) == int32(0) {
		v40 = v33
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v38 != 0 {
		v29 = v33
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v65 = v33
	goto L7
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v53 = int32(-2139062144)
	if (int32(16843008)-v50|v50)&v53 == v53 {
		v44 = v44 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v59 = v44
	goto L19
L18:
	;
	goto L17
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 != 0 {
		v59 = v59 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v65 = v59
	goto L7
L21:
	;
	goto L20
L22:
	;
	if v18&int32(3) == int32(0) {
		v97 = v18
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_afterErrorReply(m, l0, v18, v130, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L39
	}
L24:
	;
	v130 = v122 - v18
	goto L23
L25:
	;
	v101 = v97
	goto L33
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v83 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v86 = v18
	goto L29
L28:
	;
	v130 = v18 - v18
	goto L23
L29:
	;
	v90 = v86 + int32(1)
	if v90&int32(3) == int32(0) {
		v97 = v90
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v95 != 0 {
		v86 = v90
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v122 = v90
	goto L24
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v110 = int32(-2139062144)
	if (int32(16843008)-v107|v107)&v110 == v110 {
		v101 = v101 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v116 = v101
	goto L36
L35:
	;
	goto L34
L36:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 != 0 {
		v116 = v116 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v122 = v116
	goto L24
L38:
	;
	goto L37
L39:
	;
	goto L1
L40:
	;
	goto L1
}
func F_clientSubscriptionsCount(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	return v4 + v5 + (v9 + v10)
}
func F_clientTrackingCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v894 int64
	_ = v894
	var v896 int64
	_ = v896
	var v898 int64
	_ = v898
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(64)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(32)))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v10
	F_initClientPubSubData(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v29 < int32(4) {
		v634 = v27
		v635 = v27
		goto L9
	} else {
		goto L10
	}
L3:
	;
	m.G0 = v13 + int32(64)
	return
L4:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L234
	}
L5:
	;
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v13+int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v13+int32(16)))) = v894
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v896
	v898 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	F_enableTracking(m, l0, v898, v13+int32(8), v635, v634)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L233
	}
L6:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+8))
	v810 = F_objectGetVal(m, v809)
	mBase = m.M
	v811 = int32(_a1735)
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	if v814 != 0 {
		goto L212
	} else {
		goto L213
	}
L7:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v736&int32(4) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L8:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L171
	}
L9:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	v644 = F_objectGetVal(m, v643)
	mBase = m.M
	v645 = int32(_a50)
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	if v648 != 0 {
		goto L155
	} else {
		goto L156
	}
L10:
	;
	v33 = int32(0)
	v37 = v33
	v38 = v33
	v39 = v29
	v40 = int32(3)
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = v40 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v47)))
	v50 = F_objectGetVal(m, v49)
	mBase = m.M
	v51 = int32(_a1736)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v634 = v624
	v635 = v625
	goto L9
L13:
	;
	v91 = v40 + int32(1)
	v92 = base.B2i32(v39 == v91)
	if v39 == v91 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v86 = F_tolower(m, v82)
	mBase = m.M
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v88 = F_tolower(m, v87)
	mBase = m.M
	goto L13
L15:
	;
	v56 = v50
	v57 = v51
	v58 = v54
	goto L18
L16:
	;
	v82 = int32(0)
	v83 = v51
	goto L14
L17:
	;
	v82 = v79 & int32(255)
	v83 = v78
	goto L14
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v60 == int32(0) {
		v78 = v57
		v79 = v58
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v78 = v72
	v79 = int32(0)
	goto L17
L20:
	;
	v64 = v58 & int32(255)
	if v64 == v60 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(1)
	v72 = v57 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v73 != 0 {
		v56 = v56 + v71
		v57 = v72
		v58 = v73
		goto L18
	} else {
		goto L24
	}
L22:
	;
	v66 = F_tolower(m, v64)
	mBase = m.M
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v68 = F_tolower(m, v67)
	mBase = m.M
	if v66 == v68 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v78 = v57
	v79 = v70
	goto L17
L24:
	;
	goto L19
L25:
	;
	v629 = v626 + int32(1)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v629 < v630 {
		v37 = v624
		v38 = v625
		v39 = v630
		v40 = v629
		goto L11
	} else {
		goto L152
	}
L26:
	;
	v624 = v620
	v625 = v621
	v626 = v91
	goto L25
L27:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v374+v47)))
	v377 = F_objectGetVal(m, v376)
	mBase = m.M
	v378 = int32(_a1737)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v381 != 0 {
		goto L84
	} else {
		goto L85
	}
L28:
	;
	if v86-v88 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	if v93 == int64(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v91<<(uint(int32(2))%32))))
	v115 = F_getLongLongFromObjectOrReply(m, l0, v111, v13+int32(40), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L31:
	;
	F_addReplyErrorLength(m, l0, int32(_a1738), int32(51))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_afterErrorReply(m, l0, int32(_a1738), int32(51), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L3
L35:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	v122 = int64(56)
	v124 = int64(65280)
	v126 = int64(40)
	v129 = int64(16711680)
	v131 = int64(24)
	v133 = int64(4278190080)
	v135 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v121<<(uint(v122)%64) | v121&v124<<(uint(v126)%64) | (v121&v129<<(uint(v131)%64) | v121&v133<<(uint(v135)%64)) | (int64(base.Ui64(v121)>>(uint(v135)%64))&v133 | int64(base.Ui64(v121)>>(uint(v131)%64))&v129 | (int64(base.Ui64(v121)>>(uint(v126)%64))&v124 | int64(base.Ui64(v121)>>(uint(v122)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v163 = v13 + int32(56)
	v164 = int32(8)
	v166 = v13 + int32(52)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	goto L42
L36:
	;
	if v115 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L3
L39:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if v362 != 0 {
		v620 = v37
		v621 = v38
		goto L26
	} else {
		goto L77
	}
L40:
	;
	if v319 != v164 {
		goto L67
	} else {
		goto L68
	}
L41:
	;
	v310 = int32(0)
	v316 = v175
	v317 = v176
	v319 = v310
	v323 = v310
	goto L40
L42:
	;
	if base.Ui32(v176) < base.Ui32(int32(8)) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v187 = v175
	v188 = v176
	v190 = int32(0)
	goto L45
L44:
	;
	v316 = v300
	v317 = v301
	v319 = v303
	v323 = base.B2i32(v306 != int32(0))
	goto L40
L45:
	;
	v196 = int32(base.Ui32(v188) >> (uint(int32(3)) % 32))
	v197 = int32(4)
	v198 = v187 + v197
	if v188&v197 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v300 = v291
	v301 = v292
	v303 = v276
	v306 = v281
	goto L44
L47:
	;
	v281 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v198+v196+(v281-v196)&int32(3)+v269<<(uint(int32(2))%32))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if base.Ui32(v292) < base.Ui32(int32(8)) {
		v300 = v291
		v301 = v292
		v303 = v276
		v306 = v281
		goto L44
	} else {
		goto L65
	}
L48:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v190))))
	v247 = int32(0)
	goto L59
L49:
	;
	v203 = int32(0)
	if base.Ui32(v164) <= base.Ui32(v190) {
		v236 = v190
		v239 = v203
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v239 == v196 {
		v269 = v203
		v276 = v236
		goto L47
	} else {
		goto L57
	}
L51:
	;
	v213 = v190
	v216 = v203
	goto L52
L52:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v216))))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v213))))
	if v219 != v221 {
		v236 = v213
		v239 = v216
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v236 = v224
	v239 = v226
	goto L50
L54:
	;
	v223 = int32(1)
	v224 = v213 + v223
	v226 = v216 + v223
	if base.Ui32(v196) <= base.Ui32(v226) {
		v236 = v224
		v239 = v226
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if base.Ui32(v224) < base.Ui32(v164) {
		v213 = v224
		v216 = v226
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v300 = v187
	v301 = v188
	v303 = v236
	v306 = v239
	goto L44
L58:
	;
	if v247 != v196 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v247))))
	if v260 == v244&int32(255) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v262 = int32(1)
	v264 = v247 + v262
	if v264 != v196 {
		v247 = v264
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v316 = v187
	v317 = v188
	v319 = v190
	v323 = v262
	goto L40
L63:
	;
	v269 = v247
	v276 = v190 + int32(1)
	goto L47
L64:
	;
	v300 = v187
	v301 = v188
	v303 = v190
	v306 = v196
	goto L44
L65:
	;
	if base.Ui32(v276) < base.Ui32(v164) {
		v187 = v291
		v188 = v292
		v190 = v276
		goto L45
	} else {
		goto L66
	}
L66:
	;
	goto L46
L67:
	;
	goto L39
L68:
	;
	if v317&int32(1) == int32(0) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v331 = v317 & int32(4)
	if v323&base.B2i32(v331 != int32(0)) != 0 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	if v166 == int32(0) {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	if v317&int32(2) != 0 {
		v357 = int32(0)
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v357
	goto L67
L73:
	;
	v341 = int32(3)
	v342 = int32(base.Ui32(v317) >> (uint(v341) % 32))
	if v331 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v352 = int32(4)
	goto L76
L75:
	;
	v352 = v342 << (uint(int32(2)) % 32)
	goto L76
L76:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v316+v342+(int32(0)-v342)&v341+v352+int32(4))))
	v357 = v356
	goto L72
L77:
	;
	F_addReplyErrorLength(m, l0, int32(_a1739), int32(49))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_afterErrorReply(m, l0, int32(_a1739), int32(49), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_valkey_free(m, v38)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L3
L81:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421+v47)))
	v424 = F_objectGetVal(m, v423)
	mBase = m.M
	v425 = int32(_a1740)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v428 != 0 {
		goto L98
	} else {
		goto L99
	}
L82:
	;
	if v413-v415 != 0 {
		goto L81
	} else {
		goto L94
	}
L83:
	;
	v413 = F_tolower(m, v409)
	mBase = m.M
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v415 = F_tolower(m, v414)
	mBase = m.M
	goto L82
L84:
	;
	v383 = v377
	v384 = v378
	v385 = v381
	goto L87
L85:
	;
	v409 = int32(0)
	v410 = v378
	goto L83
L86:
	;
	v409 = v406 & int32(255)
	v410 = v405
	goto L83
L87:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v387 == int32(0) {
		v405 = v384
		v406 = v385
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v405 = v399
	v406 = int32(0)
	goto L86
L89:
	;
	v391 = v385 & int32(255)
	if v391 == v387 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v398 = int32(1)
	v399 = v384 + v398
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383)+1)))
	if v400 != 0 {
		v383 = v383 + v398
		v384 = v399
		v385 = v400
		goto L87
	} else {
		goto L93
	}
L91:
	;
	v393 = F_tolower(m, v391)
	mBase = m.M
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v395 = F_tolower(m, v394)
	mBase = m.M
	if v393 == v395 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	v405 = v384
	v406 = v397
	goto L86
L93:
	;
	goto L88
L94:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v417 | int32(16)
	v624 = v37
	v625 = v38
	v626 = v40
	goto L25
L95:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468+v47)))
	v471 = F_objectGetVal(m, v470)
	mBase = m.M
	v472 = int32(_a1741)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	if v475 != 0 {
		goto L112
	} else {
		goto L113
	}
L96:
	;
	if v460-v462 != 0 {
		goto L95
	} else {
		goto L108
	}
L97:
	;
	v460 = F_tolower(m, v456)
	mBase = m.M
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	v462 = F_tolower(m, v461)
	mBase = m.M
	goto L96
L98:
	;
	v430 = v424
	v431 = v425
	v432 = v428
	goto L101
L99:
	;
	v456 = int32(0)
	v457 = v425
	goto L97
L100:
	;
	v456 = v453 & int32(255)
	v457 = v452
	goto L97
L101:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	if v434 == int32(0) {
		v452 = v431
		v453 = v432
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v452 = v446
	v453 = int32(0)
	goto L100
L103:
	;
	v438 = v432 & int32(255)
	if v438 == v434 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v445 = int32(1)
	v446 = v431 + v445
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+1)))
	if v447 != 0 {
		v430 = v430 + v445
		v431 = v446
		v432 = v447
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v440 = F_tolower(m, v438)
	mBase = m.M
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	v442 = F_tolower(m, v441)
	mBase = m.M
	if v440 == v442 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v452 = v431
	v453 = v444
	goto L100
L107:
	;
	goto L102
L108:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v464 | int32(32)
	v624 = v37
	v625 = v38
	v626 = v40
	goto L25
L109:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v515+v47)))
	v518 = F_objectGetVal(m, v517)
	mBase = m.M
	v519 = int32(_a1742)
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v522 != 0 {
		goto L126
	} else {
		goto L127
	}
L110:
	;
	if v507-v509 != 0 {
		goto L109
	} else {
		goto L122
	}
L111:
	;
	v507 = F_tolower(m, v503)
	mBase = m.M
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	v509 = F_tolower(m, v508)
	mBase = m.M
	goto L110
L112:
	;
	v477 = v471
	v478 = v472
	v479 = v475
	goto L115
L113:
	;
	v503 = int32(0)
	v504 = v472
	goto L111
L114:
	;
	v503 = v500 & int32(255)
	v504 = v499
	goto L111
L115:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v481 == int32(0) {
		v499 = v478
		v500 = v479
		goto L114
	} else {
		goto L117
	}
L116:
	;
	v499 = v493
	v500 = int32(0)
	goto L114
L117:
	;
	v485 = v479 & int32(255)
	if v485 == v481 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v492 = int32(1)
	v493 = v478 + v492
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+1)))
	if v494 != 0 {
		v477 = v477 + v492
		v478 = v493
		v479 = v494
		goto L115
	} else {
		goto L121
	}
L119:
	;
	v487 = F_tolower(m, v485)
	mBase = m.M
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v489 = F_tolower(m, v488)
	mBase = m.M
	if v487 == v489 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v499 = v478
	v500 = v491
	goto L114
L121:
	;
	goto L116
L122:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v511 | int32(64)
	v624 = v37
	v625 = v38
	v626 = v40
	goto L25
L123:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v562+v47)))
	v565 = F_objectGetVal(m, v564)
	mBase = m.M
	v566 = int32(_a1743)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v569 != 0 {
		goto L139
	} else {
		goto L140
	}
L124:
	;
	if v554-v556 != 0 {
		goto L123
	} else {
		goto L136
	}
L125:
	;
	v554 = F_tolower(m, v550)
	mBase = m.M
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	v556 = F_tolower(m, v555)
	mBase = m.M
	goto L124
L126:
	;
	v524 = v518
	v525 = v519
	v526 = v522
	goto L129
L127:
	;
	v550 = int32(0)
	v551 = v519
	goto L125
L128:
	;
	v550 = v547 & int32(255)
	v551 = v546
	goto L125
L129:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	if v528 == int32(0) {
		v546 = v525
		v547 = v526
		goto L128
	} else {
		goto L131
	}
L130:
	;
	v546 = v540
	v547 = int32(0)
	goto L128
L131:
	;
	v532 = v526 & int32(255)
	if v532 == v528 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v539 = int32(1)
	v540 = v525 + v539
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+1)))
	if v541 != 0 {
		v524 = v524 + v539
		v525 = v540
		v526 = v541
		goto L129
	} else {
		goto L135
	}
L133:
	;
	v534 = F_tolower(m, v532)
	mBase = m.M
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v536 = F_tolower(m, v535)
	mBase = m.M
	if v534 == v536 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v546 = v525
	v547 = v538
	goto L128
L135:
	;
	goto L130
L136:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v558 | int32(256)
	v624 = v37
	v625 = v38
	v626 = v40
	goto L25
L137:
	;
	if v39 == v91 {
		goto L8
	} else {
		goto L149
	}
L138:
	;
	v601 = F_tolower(m, v597)
	mBase = m.M
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	v603 = F_tolower(m, v602)
	mBase = m.M
	goto L137
L139:
	;
	v571 = v565
	v572 = v566
	v573 = v569
	goto L142
L140:
	;
	v597 = int32(0)
	v598 = v566
	goto L138
L141:
	;
	v597 = v594 & int32(255)
	v598 = v593
	goto L138
L142:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v575 == int32(0) {
		v593 = v572
		v594 = v573
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v593 = v587
	v594 = int32(0)
	goto L141
L144:
	;
	v579 = v573 & int32(255)
	if v579 == v575 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v586 = int32(1)
	v587 = v572 + v586
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+1)))
	if v588 != 0 {
		v571 = v571 + v586
		v572 = v587
		v573 = v588
		goto L142
	} else {
		goto L148
	}
L146:
	;
	v581 = F_tolower(m, v579)
	mBase = m.M
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	v583 = F_tolower(m, v582)
	mBase = m.M
	if v581 == v583 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	v593 = v572
	v594 = v585
	goto L141
L148:
	;
	goto L143
L149:
	;
	if v601-v603 != 0 {
		goto L8
	} else {
		goto L150
	}
L150:
	;
	v605 = int32(2)
	v608 = v37 + int32(1)
	v611 = F_valkey_realloc(m, v38, v608<<(uint(v605)%32))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614+v91<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(v605)%32)+v611))) = v618
	v620 = v608
	v621 = v611
	goto L26
L152:
	;
	goto L12
L153:
	;
	if v680-v682 != 0 {
		goto L6
	} else {
		goto L165
	}
L154:
	;
	v680 = F_tolower(m, v676)
	mBase = m.M
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677))))
	v682 = F_tolower(m, v681)
	mBase = m.M
	goto L153
L155:
	;
	v650 = v644
	v651 = v645
	v652 = v648
	goto L158
L156:
	;
	v676 = int32(0)
	v677 = v645
	goto L154
L157:
	;
	v676 = v673 & int32(255)
	v677 = v672
	goto L154
L158:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	if v654 == int32(0) {
		v672 = v651
		v673 = v652
		goto L157
	} else {
		goto L160
	}
L159:
	;
	v672 = v666
	v673 = int32(0)
	goto L157
L160:
	;
	v658 = v652 & int32(255)
	if v658 == v654 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v665 = int32(1)
	v666 = v651 + v665
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+1)))
	if v667 != 0 {
		v650 = v650 + v665
		v651 = v666
		v652 = v667
		goto L158
	} else {
		goto L164
	}
L162:
	;
	v660 = F_tolower(m, v658)
	mBase = m.M
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	v662 = F_tolower(m, v661)
	mBase = m.M
	if v660 == v662 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650))))
	v672 = v651
	v673 = v664
	goto L157
L164:
	;
	goto L159
L165:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v688 = int32(base.Ui32(v684)>>(uint(int32(4))%32)) & int32(1)
	if v688 != 0 {
		goto L7
	} else {
		goto L166
	}
L166:
	;
	if v634 == int32(0) {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	F_addReplyErrorLength(m, l0, int32(_a1744), int32(47))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_afterErrorReply(m, l0, int32(_a1744), int32(47), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	goto L3
L171:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReply(m, l0, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v708 = F_objectGetVal(m, v705)
	mBase = m.M
	v710 = F_objectGetVal(m, v705)
	mBase = m.M
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710+int32(-1)))))
	switch v713 & int32(7) {
	case 0:
		goto L178
	case 1:
		goto L177
	case 2:
		goto L176
	case 3:
		goto L175
	case 4:
		goto L174
	default:
		v730 = int32(0)
		goto L173
	}
L173:
	;
	F_afterErrorReply(m, l0, v708, v730+int32(-2), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L179
	}
L174:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v710+int32(-17))))
	v730 = v729
	goto L173
L175:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v710+int32(-9))))
	v730 = v726
	goto L173
L176:
	;
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v710+int32(-5)))))
	v730 = v723
	goto L173
L177:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710+int32(-3)))))
	v730 = v720
	goto L173
L178:
	;
	v730 = int32(base.Ui32(v713) >> (uint(int32(3)) % 32))
	goto L173
L179:
	;
	goto L3
L180:
	;
	v758 = v684 & int32(96)
	v760 = v684 & int32(16)
	if v760 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L181:
	;
	if int32(base.Ui32(v736)>>(uint(int32(4))%32))&int32(1) == v688 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	F_addReplyErrorLength(m, l0, int32(_a1745), int32(124))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_afterErrorReply(m, l0, int32(_a1745), int32(124), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L3
L186:
	;
	if v758 != int32(96) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	if v758 == int32(0) {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	F_addReplyError(m, l0, int32(_a1746))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L3
L191:
	;
	if v684&int32(32) == int32(0) {
		goto L198
	} else {
		goto L199
	}
L192:
	;
	F_addReplyErrorLength(m, l0, int32(_a1747), int32(49))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_afterErrorReply(m, l0, int32(_a1747), int32(49), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	goto L3
L196:
	;
	if v760 == int32(0) {
		goto L5
	} else {
		goto L205
	}
L197:
	;
	F_addReplyError(m, l0, int32(_a1748))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L203
	}
L198:
	;
	if v684&int32(64) == int32(0) {
		goto L196
	} else {
		goto L201
	}
L199:
	;
	if v736&int32(64) != 0 {
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	if v736&int32(32) == int32(0) {
		goto L196
	} else {
		goto L202
	}
L202:
	;
	goto L197
L203:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L3
L205:
	;
	v804 = F_checkPrefixCollisionsOrReply(m, l0, v635, v634)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	if v804 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	goto L3
L209:
	;
	F_valkey_free(m, v635)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L224
	}
L210:
	;
	if v846-v848 != 0 {
		goto L209
	} else {
		goto L222
	}
L211:
	;
	v846 = F_tolower(m, v842)
	mBase = m.M
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843))))
	v848 = F_tolower(m, v847)
	mBase = m.M
	goto L210
L212:
	;
	v816 = v810
	v817 = v811
	v818 = v814
	goto L215
L213:
	;
	v842 = int32(0)
	v843 = v811
	goto L211
L214:
	;
	v842 = v839 & int32(255)
	v843 = v838
	goto L211
L215:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	if v820 == int32(0) {
		v838 = v817
		v839 = v818
		goto L214
	} else {
		goto L217
	}
L216:
	;
	v838 = v832
	v839 = int32(0)
	goto L214
L217:
	;
	v824 = v818 & int32(255)
	if v824 == v820 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v831 = int32(1)
	v832 = v817 + v831
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+1)))
	if v833 != 0 {
		v816 = v816 + v831
		v817 = v832
		v818 = v833
		goto L215
	} else {
		goto L221
	}
L219:
	;
	v826 = F_tolower(m, v824)
	mBase = m.M
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817))))
	v828 = F_tolower(m, v827)
	mBase = m.M
	if v826 == v828 {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816))))
	v838 = v817
	v839 = v830
	goto L214
L221:
	;
	goto L216
L222:
	;
	F_disableTracking(m, l0)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	goto L4
L224:
	;
	v855 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReply(m, l0, v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v858 = F_objectGetVal(m, v855)
	mBase = m.M
	v860 = F_objectGetVal(m, v855)
	mBase = m.M
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+int32(-1)))))
	switch v863 & int32(7) {
	case 0:
		goto L231
	case 1:
		goto L230
	case 2:
		goto L229
	case 3:
		goto L228
	case 4:
		goto L227
	default:
		v880 = int32(0)
		goto L226
	}
L226:
	;
	F_afterErrorReply(m, l0, v858, v880+int32(-2), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L232
	}
L227:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v860+int32(-17))))
	v880 = v879
	goto L226
L228:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v860+int32(-9))))
	v880 = v876
	goto L226
L229:
	;
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v860+int32(-5)))))
	v880 = v873
	goto L226
L230:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+int32(-3)))))
	v880 = v870
	goto L226
L231:
	;
	v880 = int32(base.Ui32(v863) >> (uint(int32(3)) % 32))
	goto L226
L232:
	;
	goto L3
L233:
	;
	goto L4
L234:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v910)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	goto L3
}
func F_clientTrackingInfoCommand(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int64
	_ = v229
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	v6 = m.G0
	v8 = v6 - int32(304)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	v11 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a62), int32(5))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L9
	}
L2:
	;
	return
L3:
	;
	if v11 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v17 = base.B2i32(v10&int32(255) == int32(2))
	if v10&int32(255) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = int32(42)
	goto L7
L6:
	;
	v22 = int32(37)
	goto L7
L7:
	;
	F__addReplyLongLongWithPrefix(m, l0, base.I64_extend_i32_u(int32(3)<<(uint(v17)%32)), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	v30 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v36 = v34 & int32(4)
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(_a50)
	goto L13
L12:
	;
	v37 = int32(_a1735)
	goto L13
L13:
	;
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(2)
	goto L16
L15:
	;
	v40 = int32(3)
	goto L16
L16:
	;
	F_addReplyBulkCBuffer(m, l0, v37, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v44&int32(16) == int32(0) {
		v55 = v44
		v56 = int32(1)
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v55&int32(32) == int32(0) {
		v77 = v55
		v78 = v56
		goto L21
	} else {
		goto L22
	}
L19:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1737), int32(5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v55 = v53
	v56 = int32(2)
	goto L18
L21:
	;
	if v77&int32(64) == int32(0) {
		v99 = v77
		v100 = v78
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1740), int32(5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v65&int32(128) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1749), int32(11))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v77 = v65
	v78 = v56 + int32(1)
	goto L21
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v77 = v76
	v78 = v56 + int32(2)
	goto L21
L27:
	;
	if v99&int32(256) == int32(0) {
		v112 = v99
		v113 = v100
		goto L33
	} else {
		goto L34
	}
L28:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1741), int32(6))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v87&int32(128) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1750), int32(10))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L32
	}
L31:
	;
	v99 = v87
	v100 = v78 + int32(1)
	goto L27
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v99 = v98
	v100 = v78 + int32(2)
	goto L27
L33:
	;
	if v112&int32(8) == int32(0) {
		v124 = v113
		goto L36
	} else {
		goto L37
	}
L34:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1742), int32(6))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v112 = v111
	v113 = v100 + int32(1)
	goto L33
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+224)))
	if v127 == int32(2) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1751), int32(15))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v124 = v113 + int32(1)
	goto L36
L39:
	;
	v130 = int32(42)
	goto L41
L40:
	;
	v130 = int32(126)
	goto L41
L41:
	;
	F_setDeferredAggregateLen(m, l0, v30, v124, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1736), int32(8))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v137&int32(4) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	F_addReplyBulkCBuffer(m, l0, int32(_a1752), int32(8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L60
	}
L45:
	;
	v146 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v142)+16))
	F_addReplyLongLong(m, l0, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	if v146 != 0 {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v148 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v148)
	v151 = v8 | int32(1)
	goto L54
L50:
	;
	v198 = int32(2573)
	*(*uint16)(unsafe.Add(mBase, uint32(v194+v8+int32(1)))) = uint16(v198)
	F__addReplyToBufferOrList(m, l0, v8, v194+int32(3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L59
	}
L51:
	;
	v194 = int32(0)
	goto L50
L53:
	;
	v175 = F_ull2string(m, v151+v166, int32(126), int64(1))
	mBase = m.M
	if v175 == int32(0) {
		goto L51
	} else {
		goto L57
	}
L54:
	;
	goto L56
L56:
	;
	v162 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v162)
	v166 = int32(1)
	goto L53
L57:
	;
	v194 = v175 + v166
	goto L50
L59:
	;
	goto L44
L60:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+204)))
	if v209&int32(4) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	m.G0 = v8 + int32(304)
	return
L62:
	;
	v269 = F_prepareClientToWrite(m, l0)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L2
	} else {
		goto L78
	}
L63:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	if v215 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v215)+8))
	goto L65
L65:
	;
	F_addReplyArrayLen(m, l0, base.I32_wrap_i64(v218))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(128)
	v229 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v8)+296)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v8 + int32(168)
	goto L67
L67:
	;
	v242 = int32(0)
	v244 = F_raxSeek(m, v8, int32(_a67), v242, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v246 = F_raxNext(m, v8)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L70
	}
L69:
	;
	F_raxStop(m, v8)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L77
	}
L70:
	;
	if v246 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	goto L72
L72:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	F_addReplyBulkCBuffer(m, l0, v255, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L74
	}
L73:
	;
	goto L69
L74:
	;
	v259 = F_raxNext(m, v8)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	if v259 != 0 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	goto L61
L78:
	;
	if v269 != 0 {
		goto L61
	} else {
		goto L79
	}
L79:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[504]))
	v273 = F_objectGetVal(m, v272)
	mBase = m.M
	F__addReplyToBufferOrList(m, l0, v273, int32(4))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	goto L61
}
func F_closeClientOnOutputBufferLimitReached(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v13&int32(268435456) != 0 {
		v236 = v3
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1641), int32(_a1630), int32(2277))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L60
	} else {
		goto L81
	}
L2:
	;
	F__serverAssert(m, int32(_a1638), int32(_a1630), int32(6286))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L60
	} else {
		goto L80
	}
L3:
	;
	F__serverAssert(m, int32(_a1637), int32(_a1630), int32(6285))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L60
	} else {
		goto L79
	}
L4:
	;
	m.G0 = v10 + int32(16)
	return v236
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	if base.Ui64(int64(4294901759)) <= base.Ui64(v19) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v19 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v22&int32(1024) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v25 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v22&int32(1) != 0 {
		v236 = v3
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v22&int32(2) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v22&int32(262144) != 0 {
		v236 = v3
		goto L4
	} else {
		goto L15
	}
L13:
	;
	if v22&int32(4) == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v38 == int32(0) {
		v236 = v3
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L17
L17:
	;
	v236 = v3
	goto L4
L18:
	;
	v60 = F_getClientOutputBufferMemoryUsage(m, l0)
	mBase = m.M
	if base.Ui32(v60) < base.Ui32(int32(1025)) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v13&int32(33554432) == int32(0) {
		v236 = v3
		goto L4
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if v168 == int32(0) {
		v236 = v3
		goto L4
	} else {
		goto L59
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v74&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v65&int32(6) == int32(4) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+206)))
	if v70&int32(384) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v168 = int32(1)
	goto L21
L26:
	;
	v113 = v110 * int32(24)
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[258])))
	v119 = *(*int64)(unsafe.Add(mBase, _consts[322]))
	v120 = base.I32_wrap_i64(v119)
	v121 = base.I32_wrap_i64(v115)
	if v115&int64(4294967295) < v119 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v106 = int32(0)
	v109 = v106
	v110 = v106
	goto L26
L28:
	;
	if v74&int32(2) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v74&int32(262144) != 0 {
		v109 = int32(0)
		v110 = int32(2)
		goto L26
	} else {
		goto L32
	}
L30:
	;
	v81 = int32(1)
	if v74&int32(4) == int32(0) {
		v109 = v81
		v110 = v81
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v93 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v95 == v93 {
		v109 = v93
		v110 = v93
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v98 = int32(1)
	v100 = F_isImportSlotMigrationJob(m, v95)
	mBase = m.M
	if v100 == int32(0) {
		v109 = v98
		v110 = v98
		goto L26
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	v125 = v120
	goto L37
L36:
	;
	v125 = v121
	goto L37
L37:
	;
	if v121 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = v125
	goto L40
L39:
	;
	v126 = v121
	goto L40
L40:
	;
	if v109 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = v126
	goto L43
L42:
	;
	v127 = v121
	goto L43
L43:
	;
	v129 = base.B2i32(v115 != int64(0)) & base.B2i32(base.Ui32(v127) <= base.Ui32(v60))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[259])))
	if v130 == int64(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v168 = v129 | int32(0)
	goto L21
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = int64(0)
	goto L44
L46:
	;
	v133 = base.I32_wrap_i64(v130)
	if v130&int64(4294967295) < v119 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v137 = v120
	goto L49
L48:
	;
	v137 = v133
	goto L49
L49:
	;
	if v133 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = v137
	goto L52
L51:
	;
	v138 = v133
	goto L52
L52:
	;
	if v109 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v139 = v138
	goto L55
L54:
	;
	v139 = v133
	goto L55
L55:
	;
	if base.Ui32(v60) < base.Ui32(v139) {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	v142 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	v143 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	if v143 != int64(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(l0)+336))
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[260])))
	v168 = v129 | base.B2i32(v151 < v142-v147)
	goto L21
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+336)) = v142
	goto L44
L59:
	;
	v171 = F_sdsempty(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	return int32(0)
L61:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[481]))
	v177 = F_catClientInfoString(m, v171, l0, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if l1|v179&int32(33554432) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	F_sdsfree(m, v177)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L60
	} else {
		goto L78
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v177
	F__serverLog(m, int32(3), v218, v10)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L60
	} else {
		goto L77
	}
L65:
	;
	v211 = F_freeClient(m, l0)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L60
	} else {
		goto L75
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v179 & int32(-33554433)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v188&int32(1280) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v207 {
		goto L63
	} else {
		goto L74
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v188 | int32(1024)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v195 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[483]))
	v204 = F_listAddNodeTail(m, v203, l0)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L60
	} else {
		goto L73
	}
L70:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[483]))
	v200 = F_listSearchKey(m, v199, l0)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	if v200 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	goto L67
L74:
	;
	v218 = int32(_a1640)
	goto L64
L75:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v214 {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	v218 = int32(_a1639)
	goto L64
L77:
	;
	goto L63
L78:
	;
	v228 = int32(_a20)
	v230 = *(*int64)(unsafe.Add(mBase, _consts[482]))
	*(*int64)(unsafe.Add(mBase, _consts[482])) = v230 + int64(1)
	v236 = int32(1)
	goto L4
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_createClient(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v64 int32
	_ = v64
	var v107 int64
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int64
	_ = v149
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v212 int64
	_ = v212
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v248 int32
	_ = v248
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_valkey_malloc(m, int32(392))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			v31 = F_zmalloc_usable(m, int32(16384), v12+int32(128))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v31
				v35 = F_selectDb(m, v12, int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = int32(_a20)
					v39 = *(*int64)(unsafe.Add(mBase, _consts[474]))
					*(*int64)(unsafe.Add(mBase, _consts[474])) = v39 + int64(1)
					v43 = int32(2)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+224)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v39
					v46 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+356)) = v46
					v48 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+348)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l0
					*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = v48
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+276)) = v53
					v56 = *(*int64)(unsafe.Add(mBase, _consts[109]))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+312)) = v56
					*(*int64)(unsafe.Add(mBase, uint32(v12+int32(76)))) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12)+68)) = v48
					v64 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v64
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+227)) = uint8(v46)
					*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(64)))) = uint16(v46)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12+int32(36)))) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12+int32(28)))) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12+int32(20)))) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = v48
					*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v46)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v46
					*(*int64)(unsafe.Add(mBase, uint32(v12)+188)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12)+284)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12+int32(208)))) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12)+200)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v12)+320)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v12)+292)) = v64
					v107 = *(*int64)(unsafe.Add(mBase, _consts[109]))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+368)) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v107
					*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v46
					v113 = *(*int32)(unsafe.Add(mBase, _consts[22]))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+328)) = v113
					*(*int32)(unsafe.Add(mBase, uint32(v12)+216)) = v46
					v120 = int32(1)
					v126 = int32(base.Ui32(v114^v64)>>(uint(v120)%32)) & int32(base.Ui32(v114)>>(uint(v43)%32)) & v120
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
					if v126 != 0 {
						v138 = v126<<(uint(int32(23))%32) | v129&int32(-25165825) | v126<<(uint(int32(24))%32)
					} else {
						v138 = v129 & int32(-8388609)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v138
					v140 = F_listCreate(m)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						v142 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+376)) = v142
						*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v12)+344)) = v142
						*(*int64)(unsafe.Add(mBase, uint32(v12)+384)) = int64(-1)
						v149 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = v149
						*(*int64)(unsafe.Add(mBase, uint32(v12)+336)) = v149
						*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = int32(952)
						*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = int32(953)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+222)) = uint8(v142)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+360)) = v149
						*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v149
						*(*int64)(unsafe.Add(mBase, uint32(v12)+100)) = v149
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(108)))) = v149
						*(*int32)(unsafe.Add(mBase, uint32(v12+int32(116)))) = v142
						*(*int64)(unsafe.Add(mBase, uint32(v12)+304)) = v149
						*(*int32)(unsafe.Add(mBase, uint32(v12)+280)) = v142
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+226)) = uint8(v142)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+223)) = uint8(v142)
						v182 = v12 + int32(168)
						*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v12
						*(*int64)(unsafe.Add(mBase, uint32(v182))) = v149
						*(*int64)(unsafe.Add(mBase, uint32(v12)+296)) = int64(0)
						if l0 == int32(0) {
							v244 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+232)) = v244
							*(*int64)(unsafe.Add(mBase, uint32(v12)+264)) = v244
							v248 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v248
							*(*int64)(unsafe.Add(mBase, uint32(v12)+256)) = v244
							*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v244
							*(*int64)(unsafe.Add(mBase, uint32(v12+int32(240)))) = v244
							*(*int64)(unsafe.Add(mBase, uint32(v12+int32(248)))) = v244
							*(*int64)(unsafe.Add(mBase, uint32(v12+int32(144)))) = v244
							*(*int32)(unsafe.Add(mBase, uint32(v12+int32(152)))) = v248
							m.G0 = v9 + int32(16)
							return v12
						} else {
							v191 = *(*int32)(unsafe.Add(mBase, _consts[26]))
							v192 = F_listAddNodeTail(m, v191, v12)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int32(0)
							} else {
								v194 = int32(_a20)
								v195 = *(*int32)(unsafe.Add(mBase, _consts[26]))
								v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+308)) = v196
								v198 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
								v199 = int64(56)
								v201 = int64(65280)
								v203 = int64(40)
								v206 = int64(16711680)
								v208 = int64(24)
								v210 = int64(4278190080)
								v212 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v198<<(uint(v199)%64) | v198&v201<<(uint(v203)%64) | (v198&v206<<(uint(v208)%64) | v198&v210<<(uint(v212)%64)) | (int64(base.Ui64(v198)>>(uint(v212)%64))&v210 | int64(base.Ui64(v198)>>(uint(v208)%64))&v206 | (int64(base.Ui64(v198)>>(uint(v203)%64))&v201 | int64(base.Ui64(v198)>>(uint(v199)%64))))
								v236 = *(*int32)(unsafe.Add(mBase, _consts[444]))
								v237 = int32(8)
								v241 = F_raxInsert(m, v236, v9+v237, v237, v12, int32(0))
								mBase = m.M
								v242 = m.ExcPending
								if v242 != 0 {
									return int32(0)
								} else {
									v244 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+232)) = v244
									*(*int64)(unsafe.Add(mBase, uint32(v12)+264)) = v244
									v248 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v248
									*(*int64)(unsafe.Add(mBase, uint32(v12)+256)) = v244
									*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v244
									*(*int64)(unsafe.Add(mBase, uint32(v12+int32(240)))) = v244
									*(*int64)(unsafe.Add(mBase, uint32(v12+int32(248)))) = v244
									*(*int64)(unsafe.Add(mBase, uint32(v12+int32(144)))) = v244
									*(*int32)(unsafe.Add(mBase, uint32(v12+int32(152)))) = v248
									m.G0 = v9 + int32(16)
									return v12
								}
							}
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
			v21 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, int32(107))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v12
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				v26 = v24 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v26)
				v31 = F_zmalloc_usable(m, int32(16384), v12+int32(128))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v31
					v35 = F_selectDb(m, v12, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = int32(_a20)
						v39 = *(*int64)(unsafe.Add(mBase, _consts[474]))
						*(*int64)(unsafe.Add(mBase, _consts[474])) = v39 + int64(1)
						v43 = int32(2)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+224)) = uint8(v43)
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v39
						v46 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+356)) = v46
						v48 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+348)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l0
						*(*int64)(unsafe.Add(mBase, uint32(v12)+180)) = v48
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+128))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+276)) = v53
						v56 = *(*int64)(unsafe.Add(mBase, _consts[109]))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+312)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(76)))) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12)+68)) = v48
						v64 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v64
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+227)) = uint8(v46)
						*(*uint16)(unsafe.Add(mBase, uint32(v12+int32(64)))) = uint16(v46)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(36)))) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(28)))) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(20)))) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = v48
						*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v46)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v46
						*(*int64)(unsafe.Add(mBase, uint32(v12)+188)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12)+284)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12+int32(208)))) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12)+200)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v12)+320)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(v12)+292)) = v64
						v107 = *(*int64)(unsafe.Add(mBase, _consts[109]))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+368)) = v107
						*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v107
						*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v46
						v113 = *(*int32)(unsafe.Add(mBase, _consts[22]))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+328)) = v113
						*(*int32)(unsafe.Add(mBase, uint32(v12)+216)) = v46
						v120 = int32(1)
						v126 = int32(base.Ui32(v114^v64)>>(uint(v120)%32)) & int32(base.Ui32(v114)>>(uint(v43)%32)) & v120
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
						if v126 != 0 {
							v138 = v126<<(uint(int32(23))%32) | v129&int32(-25165825) | v126<<(uint(int32(24))%32)
						} else {
							v138 = v129 & int32(-8388609)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = v138
						v140 = F_listCreate(m)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							v142 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+376)) = v142
							*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = v140
							*(*int32)(unsafe.Add(mBase, uint32(v12)+344)) = v142
							*(*int64)(unsafe.Add(mBase, uint32(v12)+384)) = int64(-1)
							v149 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+160)) = v149
							*(*int64)(unsafe.Add(mBase, uint32(v12)+336)) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = int32(952)
							*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = int32(953)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+222)) = uint8(v142)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+360)) = v149
							*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v149
							*(*int64)(unsafe.Add(mBase, uint32(v12)+100)) = v149
							*(*int64)(unsafe.Add(mBase, uint32(v12+int32(108)))) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v12+int32(116)))) = v142
							*(*int64)(unsafe.Add(mBase, uint32(v12)+304)) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v12)+280)) = v142
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+226)) = uint8(v142)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+223)) = uint8(v142)
							v182 = v12 + int32(168)
							*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v12
							*(*int64)(unsafe.Add(mBase, uint32(v182))) = v149
							*(*int64)(unsafe.Add(mBase, uint32(v12)+296)) = int64(0)
							if l0 == int32(0) {
								v244 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v12)+232)) = v244
								*(*int64)(unsafe.Add(mBase, uint32(v12)+264)) = v244
								v248 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v248
								*(*int64)(unsafe.Add(mBase, uint32(v12)+256)) = v244
								*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v244
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(240)))) = v244
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(248)))) = v244
								*(*int64)(unsafe.Add(mBase, uint32(v12+int32(144)))) = v244
								*(*int32)(unsafe.Add(mBase, uint32(v12+int32(152)))) = v248
								m.G0 = v9 + int32(16)
								return v12
							} else {
								v191 = *(*int32)(unsafe.Add(mBase, _consts[26]))
								v192 = F_listAddNodeTail(m, v191, v12)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int32(0)
								} else {
									v194 = int32(_a20)
									v195 = *(*int32)(unsafe.Add(mBase, _consts[26]))
									v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+308)) = v196
									v198 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
									v199 = int64(56)
									v201 = int64(65280)
									v203 = int64(40)
									v206 = int64(16711680)
									v208 = int64(24)
									v210 = int64(4278190080)
									v212 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v198<<(uint(v199)%64) | v198&v201<<(uint(v203)%64) | (v198&v206<<(uint(v208)%64) | v198&v210<<(uint(v212)%64)) | (int64(base.Ui64(v198)>>(uint(v212)%64))&v210 | int64(base.Ui64(v198)>>(uint(v208)%64))&v206 | (int64(base.Ui64(v198)>>(uint(v203)%64))&v201 | int64(base.Ui64(v198)>>(uint(v199)%64))))
									v236 = *(*int32)(unsafe.Add(mBase, _consts[444]))
									v237 = int32(8)
									v241 = F_raxInsert(m, v236, v9+v237, v237, v12, int32(0))
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int32(0)
									} else {
										v244 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+232)) = v244
										*(*int64)(unsafe.Add(mBase, uint32(v12)+264)) = v244
										v248 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v248
										*(*int64)(unsafe.Add(mBase, uint32(v12)+256)) = v244
										*(*int64)(unsafe.Add(mBase, uint32(v12)+136)) = v244
										*(*int64)(unsafe.Add(mBase, uint32(v12+int32(240)))) = v244
										*(*int64)(unsafe.Add(mBase, uint32(v12+int32(248)))) = v244
										*(*int64)(unsafe.Add(mBase, uint32(v12+int32(144)))) = v244
										*(*int32)(unsafe.Add(mBase, uint32(v12+int32(152)))) = v248
										m.G0 = v9 + int32(16)
										return v12
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
func F_freeClientArgv(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v4&int32(1) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v39
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(l0+int32(28)))) = v43
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15 < int32(1) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	if v3 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	if v3 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L1
L6:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_tryOffloadFreeArgvToIOThreads(m, l0, v9, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v11 != int32(-1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_valkey_free(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L16
	}
L11:
	;
	v20 = int32(0)
	goto L12
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v20<<(uint(int32(2))%32))))
	F_decrRefCount(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	goto L10
L14:
	;
	v29 = v20 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v29 < v30 {
		v20 = v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L1
}
func F_freeClientOriginalArgv(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v6&int32(1) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v14 = F_tryOffloadFreeArgvToIOThreads(m, l0, v13, v3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	return
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	goto L1
L6:
	;
	return
L7:
	;
	if v14 != int32(-1) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v19 <= v18 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_valkey_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L15
	}
L10:
	;
	v23 = v18
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	F_decrRefCount(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v32 = v23 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	if v32 < v33 {
		v23 = v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L5
}
func F_freeClientPubSubData(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v84 int64
	_ = v84
	var v94 int64
	_ = v94
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
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
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v198 int64
	_ = v198
	var v206 int64
	_ = v206
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
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
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(144)
	return
L2:
	;
	v15 = v9 + int32(88)
	v16 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[508]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	v20 = v9 + int32(80)
	v22 = *(*int64)(unsafe.Add(mBase, _consts[509]))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v22
	v25 = v9 + int32(72)
	v27 = *(*int64)(unsafe.Add(mBase, _consts[510]))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v27
	v30 = *(*int64)(unsafe.Add(mBase, _consts[511]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[512]))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v137
	v140 = *(*int64)(unsafe.Add(mBase, _consts[513]))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v140
	v143 = *(*int64)(unsafe.Add(mBase, _consts[514]))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v143
	v146 = *(*int64)(unsafe.Add(mBase, _consts[515]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	v149 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L23
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	goto L6
L6:
	;
	if v35+v36 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v41 = v9 + int32(96)
	v42 = m.T0[v32].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v44 = int32(1)
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)) = uint8(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(-1)
	if v42 == v45 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v66 = F_hashtableNext(m, v9+int32(96), v9+int32(92))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L14
	}
L10:
	;
	goto L9
L11:
	;
	goto L12
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = v41
	goto L10
L13:
	;
	F_hashtableCleanupIterator(m, v9+int32(96))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L21
	}
L14:
	;
	if v66 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(40)))) = v84
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(48)))) = v94
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(88))))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(56)))) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v9)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
	v112 = F_pubsubUnsubscribeChannel(m, l0, v108, int32(0), v9+int32(32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v118 = F_hashtableNext(m, v9+int32(96), v9+int32(92))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v118 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	goto L3
L22:
	;
	v245 = F_pubsubUnsubscribeAllPatterns(m, l0, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L40
	}
L23:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	goto L24
L24:
	;
	if v151+v152 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v157 = v9 + int32(96)
	v158 = m.T0[v148].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v160 = int32(1)
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+14)) = uint8(v161)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v157)+24)) = v161
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+15)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = int32(-1)
	if v158 == v161 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v182 = F_hashtableNext(m, v9+int32(96), v9+int32(92))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L32
	}
L28:
	;
	goto L27
L29:
	;
	goto L30
L30:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v158)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+24)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v158)+40)) = v157
	goto L28
L31:
	;
	F_hashtableCleanupIterator(m, v9+int32(96))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L39
	}
L32:
	;
	if v182 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	goto L34
L34:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(72))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(8)))) = v198
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(80))))
	*(*int64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v206
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(88))))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v9)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v9)+92))
	v220 = F_pubsubUnsubscribeChannel(m, l0, v218, int32(0), v9)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L31
L36:
	;
	v226 = F_hashtableNext(m, v9+int32(96), v9+int32(92))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v226 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	goto L22
L40:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v247&int32(262144) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	F_hashtableRelease(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v247 & int32(-262145)
	v255 = int32(_a20)
	v257 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v257 + int32(-1)
	goto L41
L43:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	F_hashtableRelease(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	F_hashtableRelease(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v278 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+24))
	if v280 == v278 {
		v286 = v277
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_valkey_free(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	F_disableTracking(m, l0)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v286 = v285
	goto L46
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(0)
	goto L1
}
func F_getClientPeerId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v14 = F__emscripten_memset_bulkmem(m, v7+int32(16), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	if v15 != 0 {
		v65 = v15
		m.G0 = v7 + int32(288)
		return v65
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v16 == int32(0) {
			v62 = F_sdsnew(m, v7+int32(16))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v62
				v65 = v62
				m.G0 = v7 + int32(288)
				return v65
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
			if v20 == int32(0) {
				v62 = F_sdsnew(m, v7+int32(16))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v62
					v65 = v62
					m.G0 = v7 + int32(288)
					return v65
				}
			} else {
				v29 = m.T0[v20].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v16, v7+int32(160), int32(128), v7+int32(156), int32(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v29 < int32(0) {
						v62 = F_sdsnew(m, v7+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v62
							v65 = v62
							m.G0 = v7 + int32(288)
							return v65
						}
					} else {
						v37 = int32(58)
						v38 = F___strchrnul(m, v7+int32(160), v37)
						mBase = m.M
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
						if v40 == v37 {
							v44 = v38
						} else {
							v44 = int32(0)
						}
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+156))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(160)
						if v44 != 0 {
							v55 = int32(_a243)
						} else {
							v55 = int32(_a244)
						}
						v56 = F_snprintf(m, v7+int32(16), int32(128), v55, v7)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v62 = F_sdsnew(m, v7+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v62
								v65 = v62
								m.G0 = v7 + int32(288)
								return v65
							}
						}
					}
				}
			}
		}
	}
}
func F_getClientPubSubShardChannels(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	return v3
}
func F_getClientTypeByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v3 = int32(_a529)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v44 = int32(1)
	v45 = int32(_a530)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	if v38-v40 != 0 {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v38 = F_tolower(m, v34)
	mBase = m.M
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v40 = F_tolower(m, v39)
	mBase = m.M
	goto L2
L4:
	;
	v8 = l0
	v9 = v3
	v10 = v6
	goto L7
L5:
	;
	v34 = int32(0)
	v35 = v3
	goto L3
L6:
	;
	v34 = v31 & int32(255)
	v35 = v30
	goto L3
L7:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v12 == int32(0) {
		v30 = v9
		v31 = v10
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v30 = v24
	v31 = int32(0)
	goto L6
L9:
	;
	v16 = v10 & int32(255)
	if v16 == v12 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v23 = int32(1)
	v24 = v9 + v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v25 != 0 {
		v8 = v8 + v23
		v9 = v24
		v10 = v25
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v18 = F_tolower(m, v16)
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v20 = F_tolower(m, v19)
	mBase = m.M
	if v18 == v20 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v30 = v9
	v31 = v22
	goto L6
L13:
	;
	goto L8
L14:
	;
	return int32(0)
L15:
	;
	return v251
L16:
	;
	if v80-v82 == int32(0) {
		v251 = v44
		goto L15
	} else {
		goto L28
	}
L17:
	;
	v80 = F_tolower(m, v76)
	mBase = m.M
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v82 = F_tolower(m, v81)
	mBase = m.M
	goto L16
L18:
	;
	v50 = l0
	v51 = v45
	v52 = v48
	goto L21
L19:
	;
	v76 = int32(0)
	v77 = v45
	goto L17
L20:
	;
	v76 = v73 & int32(255)
	v77 = v72
	goto L17
L21:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54 == int32(0) {
		v72 = v51
		v73 = v52
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v72 = v66
	v73 = int32(0)
	goto L20
L23:
	;
	v58 = v52 & int32(255)
	if v58 == v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v65 = int32(1)
	v66 = v51 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v67 != 0 {
		v50 = v50 + v65
		v51 = v66
		v52 = v67
		goto L21
	} else {
		goto L27
	}
L25:
	;
	v60 = F_tolower(m, v58)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	if v60 == v62 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v72 = v51
	v73 = v64
	goto L20
L27:
	;
	goto L22
L28:
	;
	v86 = int32(_a531)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v89 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v121-v123 == int32(0) {
		v251 = v44
		goto L15
	} else {
		goto L41
	}
L30:
	;
	v121 = F_tolower(m, v117)
	mBase = m.M
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v123 = F_tolower(m, v122)
	mBase = m.M
	goto L29
L31:
	;
	v91 = l0
	v92 = v86
	v93 = v89
	goto L34
L32:
	;
	v117 = int32(0)
	v118 = v86
	goto L30
L33:
	;
	v117 = v114 & int32(255)
	v118 = v113
	goto L30
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v95 == int32(0) {
		v113 = v92
		v114 = v93
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v113 = v107
	v114 = int32(0)
	goto L33
L36:
	;
	v99 = v93 & int32(255)
	if v99 == v95 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v106 = int32(1)
	v107 = v92 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v108 != 0 {
		v91 = v91 + v106
		v92 = v107
		v93 = v108
		goto L34
	} else {
		goto L40
	}
L38:
	;
	v101 = F_tolower(m, v99)
	mBase = m.M
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v103 = F_tolower(m, v102)
	mBase = m.M
	if v101 == v103 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v113 = v92
	v114 = v105
	goto L33
L40:
	;
	goto L35
L41:
	;
	v127 = int32(_a532)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v130 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v168 = int32(_a533)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v171 != 0 {
		goto L59
	} else {
		goto L60
	}
L43:
	;
	if v162-v164 != 0 {
		goto L42
	} else {
		goto L55
	}
L44:
	;
	v162 = F_tolower(m, v158)
	mBase = m.M
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v164 = F_tolower(m, v163)
	mBase = m.M
	goto L43
L45:
	;
	v132 = l0
	v133 = v127
	v134 = v130
	goto L48
L46:
	;
	v158 = int32(0)
	v159 = v127
	goto L44
L47:
	;
	v158 = v155 & int32(255)
	v159 = v154
	goto L44
L48:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v136 == int32(0) {
		v154 = v133
		v155 = v134
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v154 = v148
	v155 = int32(0)
	goto L47
L50:
	;
	v140 = v134 & int32(255)
	if v140 == v136 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v147 = int32(1)
	v148 = v133 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v149 != 0 {
		v132 = v132 + v147
		v133 = v148
		v134 = v149
		goto L48
	} else {
		goto L54
	}
L52:
	;
	v142 = F_tolower(m, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = F_tolower(m, v143)
	mBase = m.M
	if v142 == v144 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v154 = v133
	v155 = v146
	goto L47
L54:
	;
	goto L49
L55:
	;
	return int32(2)
L56:
	;
	v211 = int32(_a534)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v214 != 0 {
		goto L72
	} else {
		goto L73
	}
L57:
	;
	if v203-v205 != 0 {
		goto L56
	} else {
		goto L69
	}
L58:
	;
	v203 = F_tolower(m, v199)
	mBase = m.M
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v205 = F_tolower(m, v204)
	mBase = m.M
	goto L57
L59:
	;
	v173 = l0
	v174 = v168
	v175 = v171
	goto L62
L60:
	;
	v199 = int32(0)
	v200 = v168
	goto L58
L61:
	;
	v199 = v196 & int32(255)
	v200 = v195
	goto L58
L62:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v177 == int32(0) {
		v195 = v174
		v196 = v175
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v195 = v189
	v196 = int32(0)
	goto L61
L64:
	;
	v181 = v175 & int32(255)
	if v181 == v177 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v188 = int32(1)
	v189 = v174 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+1)))
	if v190 != 0 {
		v173 = v173 + v188
		v174 = v189
		v175 = v190
		goto L62
	} else {
		goto L68
	}
L66:
	;
	v183 = F_tolower(m, v181)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v185 = F_tolower(m, v184)
	mBase = m.M
	if v183 == v185 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v195 = v174
	v196 = v187
	goto L61
L68:
	;
	goto L63
L69:
	;
	return int32(3)
L70:
	;
	if v246-v248 != 0 {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v246 = F_tolower(m, v242)
	mBase = m.M
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v248 = F_tolower(m, v247)
	mBase = m.M
	goto L70
L72:
	;
	v216 = l0
	v217 = v211
	v218 = v214
	goto L75
L73:
	;
	v242 = int32(0)
	v243 = v211
	goto L71
L74:
	;
	v242 = v239 & int32(255)
	v243 = v238
	goto L71
L75:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v220 == int32(0) {
		v238 = v217
		v239 = v218
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v238 = v232
	v239 = int32(0)
	goto L74
L77:
	;
	v224 = v218 & int32(255)
	if v224 == v220 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v231 = int32(1)
	v232 = v217 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	if v233 != 0 {
		v216 = v216 + v231
		v217 = v232
		v218 = v233
		goto L75
	} else {
		goto L81
	}
L79:
	;
	v226 = F_tolower(m, v224)
	mBase = m.M
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v228 = F_tolower(m, v227)
	mBase = m.M
	if v226 == v228 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	v238 = v217
	v239 = v230
	goto L74
L81:
	;
	goto L76
L82:
	;
	v250 = int32(-1)
	goto L84
L83:
	;
	v250 = int32(3)
	goto L84
L84:
	;
	v251 = v250
	goto L15
}
func F_getClientType_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(1) == int32(0) {
		v11 = int32(2)
		if v4&v11 == int32(0) {
			if v4&int32(262144) != 0 {
				v32 = v11
				return v32
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v22 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v27 == int32(1) {
						v30 = int32(4)
					} else {
						v30 = int32(5)
					}
					v32 = v30
					return v32
				} else {
					return int32(0)
				}
			}
		} else {
			if v4&int32(4) != 0 {
				if v4&int32(262144) != 0 {
					v32 = v11
					return v32
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v22 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						if v27 == int32(1) {
							v30 = int32(4)
						} else {
							v30 = int32(5)
						}
						v32 = v30
						return v32
					} else {
						return int32(0)
					}
				}
			} else {
				return int32(1)
			}
		}
	} else {
		return int32(3)
	}
}
func F_getClientType_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v4&int32(1) == int32(0) {
		v11 = int32(2)
		if v4&v11 == int32(0) {
			if v4&int32(262144) != 0 {
				v32 = v11
				return v32
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
				if v22 != 0 {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					if v27 == int32(1) {
						v30 = int32(4)
					} else {
						v30 = int32(5)
					}
					v32 = v30
					return v32
				} else {
					return int32(0)
				}
			}
		} else {
			if v4&int32(4) != 0 {
				if v4&int32(262144) != 0 {
					v32 = v11
					return v32
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
					if v22 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						if v27 == int32(1) {
							v30 = int32(4)
						} else {
							v30 = int32(5)
						}
						v32 = v30
						return v32
					} else {
						return int32(0)
					}
				}
			} else {
				return int32(1)
			}
		}
	} else {
		return int32(3)
	}
}
func F_initClientBlockingState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v4 != 0 {
		return
	} else {
		v6 = F_valkey_malloc(m, int32(56))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v6
			v9 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v9
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
			v16 = F_dictCreate(m, int32(_a182))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				v19 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v19
				*(*int64)(unsafe.Add(mBase, uint32(v18)+28)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
				*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v19
				return
			}
		}
	}
}
func F_lookupClientByID(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int64(56)
	v10 = int64(65280)
	v12 = int64(40)
	v15 = int64(16711680)
	v17 = int64(24)
	v19 = int64(4278190080)
	v21 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = l0<<(uint(v8)%64) | l0&v10<<(uint(v12)%64) | (l0&v15<<(uint(v17)%64) | l0&v19<<(uint(v21)%64)) | (int64(base.Ui64(l0)>>(uint(v21)%64))&v19 | int64(base.Ui64(l0)>>(uint(v17)%64))&v15 | (int64(base.Ui64(l0)>>(uint(v12)%64))&v10 | int64(base.Ui64(l0)>>(uint(v8)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v48 = int32(8)
	v49 = v6 + v48
	v52 = v6 + int32(4)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	goto L4
L1:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	m.G0 = v6 + int32(16)
	return v248
L2:
	;
	if v205 != v48 {
		goto L29
	} else {
		goto L30
	}
L3:
	;
	v196 = int32(0)
	v202 = v61
	v203 = v62
	v205 = v196
	v209 = v196
	goto L2
L4:
	;
	if base.Ui32(v62) < base.Ui32(int32(8)) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v73 = v61
	v74 = v62
	v76 = int32(0)
	goto L7
L6:
	;
	v202 = v186
	v203 = v187
	v205 = v189
	v209 = base.B2i32(v192 != int32(0))
	goto L2
L7:
	;
	v82 = int32(base.Ui32(v74) >> (uint(int32(3)) % 32))
	v83 = int32(4)
	v84 = v73 + v83
	if v74&v83 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v186 = v177
	v187 = v178
	v189 = v162
	v192 = v167
	goto L6
L9:
	;
	v167 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v84+v82+(v167-v82)&int32(3)+v155<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if base.Ui32(v178) < base.Ui32(int32(8)) {
		v186 = v177
		v187 = v178
		v189 = v162
		v192 = v167
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v76))))
	v133 = int32(0)
	goto L21
L11:
	;
	v89 = int32(0)
	if base.Ui32(v48) <= base.Ui32(v76) {
		v122 = v76
		v125 = v89
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v125 == v82 {
		v155 = v89
		v162 = v122
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v99 = v76
	v102 = v89
	goto L14
L14:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v102))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v99))))
	if v105 != v107 {
		v122 = v99
		v125 = v102
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v122 = v110
	v125 = v112
	goto L12
L16:
	;
	v109 = int32(1)
	v110 = v99 + v109
	v112 = v102 + v109
	if base.Ui32(v82) <= base.Ui32(v112) {
		v122 = v110
		v125 = v112
		goto L12
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v110) < base.Ui32(v48) {
		v99 = v110
		v102 = v112
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v186 = v73
	v187 = v74
	v189 = v122
	v192 = v125
	goto L6
L20:
	;
	if v133 != v82 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v133))))
	if v146 == v130&int32(255) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v148 = int32(1)
	v150 = v133 + v148
	if v150 != v82 {
		v133 = v150
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v202 = v73
	v203 = v74
	v205 = v76
	v209 = v148
	goto L2
L25:
	;
	v155 = v133
	v162 = v76 + int32(1)
	goto L9
L26:
	;
	v186 = v73
	v187 = v74
	v189 = v76
	v192 = v82
	goto L6
L27:
	;
	if base.Ui32(v162) < base.Ui32(v48) {
		v73 = v177
		v74 = v178
		v76 = v162
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
	if v203&int32(1) == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v217 = v203 & int32(4)
	if v209&base.B2i32(v217 != int32(0)) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v52 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v203&int32(2) != 0 {
		v243 = int32(0)
		goto L34
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v243
	goto L29
L35:
	;
	v227 = int32(3)
	v228 = int32(base.Ui32(v203) >> (uint(v227) % 32))
	if v217 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v238 = int32(4)
	goto L38
L37:
	;
	v238 = v228 << (uint(int32(2)) % 32)
	goto L38
L38:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v202+v228+(int32(0)-v228)&v227+v238+int32(4))))
	v243 = v242
	goto L34
}
func F_parseClientFiltersOrReply(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int64
	_ = v175
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v197 int64
	_ = v197
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v218 int64
	_ = v218
	var v242 int64
	_ = v242
	var v261 int32
	_ = v261
	var v270 int64
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v433 int64
	_ = v433
	var v439 int32
	_ = v439
	var v441 int64
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v455 int64
	_ = v455
	var v460 int64
	_ = v460
	var v464 int32
	_ = v464
	var v466 int64
	_ = v466
	var v468 int32
	_ = v468
	var v476 int64
	_ = v476
	var v500 int64
	_ = v500
	var v519 int32
	_ = v519
	var v528 int64
	_ = v528
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int64
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1512 int32
	_ = v1512
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1633 int32
	_ = v1633
	var v1644 int32
	_ = v1644
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1664 int32
	_ = v1664
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2091 int32
	_ = v2091
	var v2098 int32
	_ = v2098
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2165 int32
	_ = v2165
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2220 int32
	_ = v2220
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2251 int32
	_ = v2251
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
	var v2284 int32
	_ = v2284
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2375 int32
	_ = v2375
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2408 int32
	_ = v2408
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2526 int32
	_ = v2526
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2553 int32
	_ = v2553
	var v2559 int32
	_ = v2559
	var v2573 int32
	_ = v2573
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v15 < int32(3) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(160)
	return v2577
L2:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v2573)
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L24
	} else {
		goto L728
	}
L3:
	;
	v2577 = int32(0)
	goto L1
L4:
	;
	v22 = v15
	v23 = int32(2)
	goto L5
L5:
	;
	v30 = v23 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = v23 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33)))
	v36 = F_objectGetVal(m, v35)
	mBase = m.M
	v37 = int32(_a1698)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L3
L7:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2553 < v2559 {
		v22 = v2559
		v23 = v2553
		goto L5
	} else {
		goto L727
	}
L8:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293+v33)))
	v296 = F_objectGetVal(m, v295)
	mBase = m.M
	v297 = int32(_a1699)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v300 != 0 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	if v72-v74 != 0 {
		goto L8
	} else {
		goto L21
	}
L10:
	;
	v72 = F_tolower(m, v68)
	mBase = m.M
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v74 = F_tolower(m, v73)
	mBase = m.M
	goto L9
L11:
	;
	v42 = v36
	v43 = v37
	v44 = v40
	goto L14
L12:
	;
	v68 = int32(0)
	v69 = v37
	goto L10
L13:
	;
	v68 = v65 & int32(255)
	v69 = v64
	goto L10
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v46 == int32(0) {
		v64 = v43
		v65 = v44
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v64 = v58
	v65 = int32(0)
	goto L13
L16:
	;
	v50 = v44 & int32(255)
	if v50 == v46 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v57 = int32(1)
	v58 = v43 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v59 != 0 {
		v42 = v42 + v57
		v43 = v58
		v44 = v59
		goto L14
	} else {
		goto L20
	}
L18:
	;
	v52 = F_tolower(m, v50)
	mBase = m.M
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v54 = F_tolower(m, v53)
	mBase = m.M
	if v52 == v54 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v64 = v43
	v65 = v56
	goto L13
L20:
	;
	goto L15
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v82 <= v30 {
		v2553 = v30
		goto L7
	} else {
		goto L26
	}
L23:
	;
	v77 = F_intsetNew(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v77
	goto L22
L26:
	;
	v88 = v30
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v96 = v88 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v96)))
	v99 = F_objectGetVal(m, v98)
	mBase = m.M
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101+v96)))
	v104 = F_objectGetVal(m, v103)
	mBase = m.M
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+int32(-1)))))
	switch v107 & int32(7) {
	case 0:
		goto L34
	case 1:
		goto L33
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		v124 = int32(0)
		goto L29
	}
L29:
	;
	v126 = v13 + int32(152)
	v127 = int32(0)
	if base.Ui32(v124+int32(-21)) < base.Ui32(int32(-20)) {
		v261 = v127
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(-17))))
	v124 = v123
	goto L29
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v104+int32(-9))))
	v124 = v120
	goto L29
L32:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+int32(-5)))))
	v124 = v117
	goto L29
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+int32(-3)))))
	v124 = v114
	goto L29
L34:
	;
	v124 = int32(base.Ui32(v107) >> (uint(int32(3)) % 32))
	goto L29
L35:
	;
	if v261 == int32(0) {
		v2553 = v88
		goto L7
	} else {
		goto L62
	}
L36:
	;
	goto L35
L37:
	;
	v139 = int32(1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v124 != v139 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v261 = int32(1)
	goto L36
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v242
	goto L38
L40:
	;
	if v140&int32(255) == int32(45) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v144 = v140 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v144&int32(255)) {
		v261 = v127
		goto L36
	} else {
		goto L42
	}
L42:
	;
	if v126 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v242 = base.I64_extend_i32_u(v144) & int64(255)
	goto L39
L44:
	;
	if base.Ui32(int32(8)) < base.Ui32((v163+int32(-49))&int32(255)) {
		v261 = v127
		goto L36
	} else {
		goto L47
	}
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v162 = int32(2)
	v163 = v160
	v164 = v99 + int32(1)
	goto L44
L46:
	;
	v162 = v139
	v163 = v140
	v164 = v99
	goto L44
L47:
	;
	v175 = base.I64_extend_i32_u(v163+int32(-48)) & int64(255)
	if base.Ui32(v124) <= base.Ui32(v162) {
		v218 = v175
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v140&int32(255) != int32(45) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v181 = v162
	v183 = v175
	v185 = v164
	goto L50
L50:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	if base.Ui32((v187+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v261 = v127
		goto L36
	} else {
		goto L52
	}
L51:
	;
	v218 = v208
	goto L48
L52:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v183) {
		v261 = v127
		goto L36
	} else {
		goto L53
	}
L53:
	;
	v197 = v183 * int64(10)
	v202 = base.I64_extend_i32_u(v187+int32(-48)) & int64(255)
	if base.Ui64(v202^int64(-1)) < base.Ui64(v197) {
		v261 = v127
		goto L36
	} else {
		goto L54
	}
L54:
	;
	v206 = int32(1)
	v208 = v197 + v202
	v210 = v181 + v206
	if v210 != v124 {
		v181 = v210
		v183 = v208
		v185 = v185 + v206
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	if v218 < int64(0) {
		v261 = v127
		goto L36
	} else {
		goto L60
	}
L57:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v218) {
		v261 = v127
		goto L36
	} else {
		goto L58
	}
L58:
	;
	if v126 == int32(0) {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	v242 = int64(0) - v218
	goto L39
L60:
	;
	if v126 == int32(0) {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v242 = v218
	goto L39
L62:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	if int64(0) < v270 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v286 = F_intsetAdd(m, v283, v270, v13+int32(151))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L24
	} else {
		goto L67
	}
L64:
	;
	F_addReplyErrorLength(m, l0, int32(_a1700), int32(34))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L24
	} else {
		goto L65
	}
L65:
	;
	F_afterErrorReply(m, l0, int32(_a1700), int32(34), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L24
	} else {
		goto L66
	}
L66:
	;
	v2577 = int32(-1)
	goto L1
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v286
	v290 = v88 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v290 < v291 {
		v88 = v290
		goto L27
	} else {
		goto L68
	}
L68:
	;
	v2553 = v290
	goto L7
L69:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v551+v33)))
	v554 = F_objectGetVal(m, v553)
	mBase = m.M
	v555 = int32(_a1701)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554))))
	if v558 != 0 {
		goto L132
	} else {
		goto L133
	}
L70:
	;
	if v332-v334 != 0 {
		goto L69
	} else {
		goto L82
	}
L71:
	;
	v332 = F_tolower(m, v328)
	mBase = m.M
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	v334 = F_tolower(m, v333)
	mBase = m.M
	goto L70
L72:
	;
	v302 = v296
	v303 = v297
	v304 = v300
	goto L75
L73:
	;
	v328 = int32(0)
	v329 = v297
	goto L71
L74:
	;
	v328 = v325 & int32(255)
	v329 = v324
	goto L71
L75:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v306 == int32(0) {
		v324 = v303
		v325 = v304
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v324 = v318
	v325 = int32(0)
	goto L74
L77:
	;
	v310 = v304 & int32(255)
	if v310 == v306 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v317 = int32(1)
	v318 = v303 + v317
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	if v319 != 0 {
		v302 = v302 + v317
		v303 = v318
		v304 = v319
		goto L75
	} else {
		goto L81
	}
L79:
	;
	v312 = F_tolower(m, v310)
	mBase = m.M
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	v314 = F_tolower(m, v313)
	mBase = m.M
	if v312 == v314 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v324 = v303
	v325 = v316
	goto L74
L81:
	;
	goto L76
L82:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v336 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v340 <= v30 {
		v2553 = v30
		goto L7
	} else {
		goto L86
	}
L84:
	;
	v337 = F_intsetNew(m)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L24
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v337
	goto L83
L86:
	;
	v346 = v30
	goto L87
L87:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v354 = v346 << (uint(int32(2)) % 32)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352+v354)))
	v357 = F_objectGetVal(m, v356)
	mBase = m.M
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v359+v354)))
	v362 = F_objectGetVal(m, v361)
	mBase = m.M
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+int32(-1)))))
	switch v365 & int32(7) {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	default:
		v382 = int32(0)
		goto L89
	}
L89:
	;
	v384 = v13 + int32(152)
	v385 = int32(0)
	if base.Ui32(v382+int32(-21)) < base.Ui32(int32(-20)) {
		v519 = v385
		goto L96
	} else {
		goto L97
	}
L90:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v362+int32(-17))))
	v382 = v381
	goto L89
L91:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v362+int32(-9))))
	v382 = v378
	goto L89
L92:
	;
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362+int32(-5)))))
	v382 = v375
	goto L89
L93:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+int32(-3)))))
	v382 = v372
	goto L89
L94:
	;
	v382 = int32(base.Ui32(v365) >> (uint(int32(3)) % 32))
	goto L89
L95:
	;
	if v519 == int32(0) {
		v2553 = v346
		goto L7
	} else {
		goto L122
	}
L96:
	;
	goto L95
L97:
	;
	v397 = int32(1)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	if v382 != v397 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v519 = int32(1)
	goto L96
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v384))) = v500
	goto L98
L100:
	;
	if v398&int32(255) == int32(45) {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v402 = v398 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v402&int32(255)) {
		v519 = v385
		goto L96
	} else {
		goto L102
	}
L102:
	;
	if v384 == int32(0) {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v500 = base.I64_extend_i32_u(v402) & int64(255)
	goto L99
L104:
	;
	if base.Ui32(int32(8)) < base.Ui32((v421+int32(-49))&int32(255)) {
		v519 = v385
		goto L96
	} else {
		goto L107
	}
L105:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
	v420 = int32(2)
	v421 = v418
	v422 = v357 + int32(1)
	goto L104
L106:
	;
	v420 = v397
	v421 = v398
	v422 = v357
	goto L104
L107:
	;
	v433 = base.I64_extend_i32_u(v421+int32(-48)) & int64(255)
	if base.Ui32(v382) <= base.Ui32(v420) {
		v476 = v433
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if v398&int32(255) != int32(45) {
		goto L116
	} else {
		goto L117
	}
L109:
	;
	v439 = v420
	v441 = v433
	v443 = v422
	goto L110
L110:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443)+1)))
	if base.Ui32((v445+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v519 = v385
		goto L96
	} else {
		goto L112
	}
L111:
	;
	v476 = v466
	goto L108
L112:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v441) {
		v519 = v385
		goto L96
	} else {
		goto L113
	}
L113:
	;
	v455 = v441 * int64(10)
	v460 = base.I64_extend_i32_u(v445+int32(-48)) & int64(255)
	if base.Ui64(v460^int64(-1)) < base.Ui64(v455) {
		v519 = v385
		goto L96
	} else {
		goto L114
	}
L114:
	;
	v464 = int32(1)
	v466 = v455 + v460
	v468 = v439 + v464
	if v468 != v382 {
		v439 = v468
		v441 = v466
		v443 = v443 + v464
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	if v476 < int64(0) {
		v519 = v385
		goto L96
	} else {
		goto L120
	}
L117:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v476) {
		v519 = v385
		goto L96
	} else {
		goto L118
	}
L118:
	;
	if v384 == int32(0) {
		goto L98
	} else {
		goto L119
	}
L119:
	;
	v500 = int64(0) - v476
	goto L99
L120:
	;
	if v384 == int32(0) {
		goto L98
	} else {
		goto L121
	}
L121:
	;
	v500 = v476
	goto L99
L122:
	;
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	if int64(0) < v528 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v544 = F_intsetAdd(m, v541, v528, v13+int32(151))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L24
	} else {
		goto L127
	}
L124:
	;
	F_addReplyErrorLength(m, l0, int32(_a1700), int32(34))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L24
	} else {
		goto L125
	}
L125:
	;
	F_afterErrorReply(m, l0, int32(_a1700), int32(34), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L24
	} else {
		goto L126
	}
L126:
	;
	v2577 = int32(-1)
	goto L1
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v544
	v548 = v346 + int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v548 < v549 {
		v346 = v548
		goto L87
	} else {
		goto L128
	}
L128:
	;
	v2553 = v548
	goto L7
L129:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v622+v33)))
	v625 = F_objectGetVal(m, v624)
	mBase = m.M
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v627 = int32(_a1702)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	if v630 != 0 {
		goto L154
	} else {
		goto L155
	}
L130:
	;
	if v590-v592 != 0 {
		goto L129
	} else {
		goto L142
	}
L131:
	;
	v590 = F_tolower(m, v586)
	mBase = m.M
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	v592 = F_tolower(m, v591)
	mBase = m.M
	goto L130
L132:
	;
	v560 = v554
	v561 = v555
	v562 = v558
	goto L135
L133:
	;
	v586 = int32(0)
	v587 = v555
	goto L131
L134:
	;
	v586 = v583 & int32(255)
	v587 = v582
	goto L131
L135:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	if v564 == int32(0) {
		v582 = v561
		v583 = v562
		goto L134
	} else {
		goto L137
	}
L136:
	;
	v582 = v576
	v583 = int32(0)
	goto L134
L137:
	;
	v568 = v562 & int32(255)
	if v568 == v564 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v575 = int32(1)
	v576 = v561 + v575
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+1)))
	if v577 != 0 {
		v560 = v560 + v575
		v561 = v576
		v562 = v577
		goto L135
	} else {
		goto L141
	}
L139:
	;
	v570 = F_tolower(m, v568)
	mBase = m.M
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	v572 = F_tolower(m, v571)
	mBase = m.M
	if v570 == v572 {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	v582 = v561
	v583 = v574
	goto L134
L141:
	;
	goto L136
L142:
	;
	if v22 <= v30 {
		goto L129
	} else {
		goto L143
	}
L143:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595+v30<<(uint(int32(2))%32))))
	v603 = F_getLongLongFromObjectOrReply(m, l0, v599, v13+int32(152), int32(_a1703))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L24
	} else {
		goto L146
	}
L144:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v605
	v2553 = v23 + int32(2)
	goto L7
L145:
	;
	v2577 = int32(-1)
	goto L1
L146:
	;
	if v603 != 0 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	if int64(0) < v605 {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	F_addReplyErrorLength(m, l0, int32(_a1704), int32(31))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L24
	} else {
		goto L149
	}
L149:
	;
	F_afterErrorReply(m, l0, int32(_a1704), int32(31), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L24
	} else {
		goto L150
	}
L150:
	;
	goto L145
L151:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v626+v33)))
	v717 = F_objectGetVal(m, v716)
	mBase = m.M
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v719 = int32(_a1705)
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if v722 != 0 {
		goto L185
	} else {
		goto L186
	}
L152:
	;
	if v662-v664 != 0 {
		goto L151
	} else {
		goto L164
	}
L153:
	;
	v662 = F_tolower(m, v658)
	mBase = m.M
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	v664 = F_tolower(m, v663)
	mBase = m.M
	goto L152
L154:
	;
	v632 = v625
	v633 = v627
	v634 = v630
	goto L157
L155:
	;
	v658 = int32(0)
	v659 = v627
	goto L153
L156:
	;
	v658 = v655 & int32(255)
	v659 = v654
	goto L153
L157:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	if v636 == int32(0) {
		v654 = v633
		v655 = v634
		goto L156
	} else {
		goto L159
	}
L158:
	;
	v654 = v648
	v655 = int32(0)
	goto L156
L159:
	;
	v640 = v634 & int32(255)
	if v640 == v636 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v647 = int32(1)
	v648 = v633 + v647
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+1)))
	if v649 != 0 {
		v632 = v632 + v647
		v633 = v648
		v634 = v649
		goto L157
	} else {
		goto L163
	}
L161:
	;
	v642 = F_tolower(m, v640)
	mBase = m.M
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v644 = F_tolower(m, v643)
	mBase = m.M
	if v642 == v644 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	v654 = v633
	v655 = v646
	goto L156
L163:
	;
	goto L158
L164:
	;
	if v22 <= v30 {
		goto L151
	} else {
		goto L165
	}
L165:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v626+v30<<(uint(int32(2))%32))))
	v671 = F_objectGetVal(m, v670)
	mBase = m.M
	v674 = F_strcasecmp(m, v671, int32(_a529))
	mBase = m.M
	if v674 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v698
	if v698 != int32(-1) {
		goto L179
	} else {
		goto L180
	}
L167:
	;
	v676 = int32(1)
	v678 = F_strcasecmp(m, v671, int32(_a530))
	mBase = m.M
	if v678 == int32(0) {
		v696 = v676
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v698 = int32(0)
	goto L166
L169:
	;
	v698 = v696
	goto L166
L170:
	;
	v682 = F_strcasecmp(m, v671, int32(_a531))
	mBase = m.M
	if v682 == int32(0) {
		v696 = v676
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v686 = F_strcasecmp(m, v671, int32(_a532))
	mBase = m.M
	if v686 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v689 = F_strcasecmp(m, v671, int32(_a533))
	mBase = m.M
	if v689 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v698 = int32(2)
	goto L166
L174:
	;
	v694 = F_strcasecmp(m, v671, int32(_a534))
	mBase = m.M
	if v694 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v698 = int32(3)
	goto L166
L176:
	;
	v695 = int32(-1)
	goto L178
L177:
	;
	v695 = int32(3)
	goto L178
L178:
	;
	v696 = v695
	goto L169
L179:
	;
	v2553 = v23 + int32(2)
	goto L7
L180:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v702+v30<<(uint(int32(2))%32))))
	v707 = F_objectGetVal(m, v706)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v707
	F_addReplyErrorFormat(m, l0, int32(_a1706), v13)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L24
	} else {
		goto L181
	}
L181:
	;
	v2577 = int32(-1)
	goto L1
L182:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v718+v33)))
	v811 = F_objectGetVal(m, v810)
	mBase = m.M
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v813 = int32(_a1707)
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811))))
	if v816 != 0 {
		goto L216
	} else {
		goto L217
	}
L183:
	;
	if v754-v756 != 0 {
		goto L182
	} else {
		goto L195
	}
L184:
	;
	v754 = F_tolower(m, v750)
	mBase = m.M
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	v756 = F_tolower(m, v755)
	mBase = m.M
	goto L183
L185:
	;
	v724 = v717
	v725 = v719
	v726 = v722
	goto L188
L186:
	;
	v750 = int32(0)
	v751 = v719
	goto L184
L187:
	;
	v750 = v747 & int32(255)
	v751 = v746
	goto L184
L188:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725))))
	if v728 == int32(0) {
		v746 = v725
		v747 = v726
		goto L187
	} else {
		goto L190
	}
L189:
	;
	v746 = v740
	v747 = int32(0)
	goto L187
L190:
	;
	v732 = v726 & int32(255)
	if v732 == v728 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v739 = int32(1)
	v740 = v725 + v739
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724)+1)))
	if v741 != 0 {
		v724 = v724 + v739
		v725 = v740
		v726 = v741
		goto L188
	} else {
		goto L194
	}
L192:
	;
	v734 = F_tolower(m, v732)
	mBase = m.M
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725))))
	v736 = F_tolower(m, v735)
	mBase = m.M
	if v734 == v736 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	v746 = v725
	v747 = v738
	goto L187
L194:
	;
	goto L189
L195:
	;
	if v22 <= v30 {
		goto L182
	} else {
		goto L196
	}
L196:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v718+v30<<(uint(int32(2))%32))))
	v763 = F_objectGetVal(m, v762)
	mBase = m.M
	v766 = F_strcasecmp(m, v763, int32(_a529))
	mBase = m.M
	if v766 != 0 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v790
	if v790 != int32(-1) {
		goto L210
	} else {
		goto L211
	}
L198:
	;
	v768 = int32(1)
	v770 = F_strcasecmp(m, v763, int32(_a530))
	mBase = m.M
	if v770 == int32(0) {
		v788 = v768
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v790 = int32(0)
	goto L197
L200:
	;
	v790 = v788
	goto L197
L201:
	;
	v774 = F_strcasecmp(m, v763, int32(_a531))
	mBase = m.M
	if v774 == int32(0) {
		v788 = v768
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v778 = F_strcasecmp(m, v763, int32(_a532))
	mBase = m.M
	if v778 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v781 = F_strcasecmp(m, v763, int32(_a533))
	mBase = m.M
	if v781 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v790 = int32(2)
	goto L197
L205:
	;
	v786 = F_strcasecmp(m, v763, int32(_a534))
	mBase = m.M
	if v786 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v790 = int32(3)
	goto L197
L207:
	;
	v787 = int32(-1)
	goto L209
L208:
	;
	v787 = int32(3)
	goto L209
L209:
	;
	v788 = v787
	goto L200
L210:
	;
	v2553 = v23 + int32(2)
	goto L7
L211:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v794+v30<<(uint(int32(2))%32))))
	v799 = F_objectGetVal(m, v798)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v799
	F_addReplyErrorFormat(m, l0, int32(_a1706), v13+int32(16))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L24
	} else {
		goto L212
	}
L212:
	;
	v2577 = int32(-1)
	goto L1
L213:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v812+v33)))
	v863 = F_objectGetVal(m, v862)
	mBase = m.M
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v865 = int32(_a1708)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	if v868 != 0 {
		goto L231
	} else {
		goto L232
	}
L214:
	;
	if v848-v850 != 0 {
		goto L213
	} else {
		goto L226
	}
L215:
	;
	v848 = F_tolower(m, v844)
	mBase = m.M
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	v850 = F_tolower(m, v849)
	mBase = m.M
	goto L214
L216:
	;
	v818 = v811
	v819 = v813
	v820 = v816
	goto L219
L217:
	;
	v844 = int32(0)
	v845 = v813
	goto L215
L218:
	;
	v844 = v841 & int32(255)
	v845 = v840
	goto L215
L219:
	;
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819))))
	if v822 == int32(0) {
		v840 = v819
		v841 = v820
		goto L218
	} else {
		goto L221
	}
L220:
	;
	v840 = v834
	v841 = int32(0)
	goto L218
L221:
	;
	v826 = v820 & int32(255)
	if v826 == v822 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v833 = int32(1)
	v834 = v819 + v833
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+1)))
	if v835 != 0 {
		v818 = v818 + v833
		v819 = v834
		v820 = v835
		goto L219
	} else {
		goto L225
	}
L223:
	;
	v828 = F_tolower(m, v826)
	mBase = m.M
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819))))
	v830 = F_tolower(m, v829)
	mBase = m.M
	if v828 == v830 {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818))))
	v840 = v819
	v841 = v832
	goto L218
L225:
	;
	goto L220
L226:
	;
	if v22 <= v30 {
		goto L213
	} else {
		goto L227
	}
L227:
	;
	v853 = int32(2)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v812+v30<<(uint(v853)%32))))
	v857 = F_objectGetVal(m, v856)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v857
	v2553 = v23 + v853
	goto L7
L228:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v864+v33)))
	v915 = F_objectGetVal(m, v914)
	mBase = m.M
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v917 = int32(_a1709)
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915))))
	if v920 != 0 {
		goto L246
	} else {
		goto L247
	}
L229:
	;
	if v900-v902 != 0 {
		goto L228
	} else {
		goto L241
	}
L230:
	;
	v900 = F_tolower(m, v896)
	mBase = m.M
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897))))
	v902 = F_tolower(m, v901)
	mBase = m.M
	goto L229
L231:
	;
	v870 = v863
	v871 = v865
	v872 = v868
	goto L234
L232:
	;
	v896 = int32(0)
	v897 = v865
	goto L230
L233:
	;
	v896 = v893 & int32(255)
	v897 = v892
	goto L230
L234:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	if v874 == int32(0) {
		v892 = v871
		v893 = v872
		goto L233
	} else {
		goto L236
	}
L235:
	;
	v892 = v886
	v893 = int32(0)
	goto L233
L236:
	;
	v878 = v872 & int32(255)
	if v878 == v874 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v885 = int32(1)
	v886 = v871 + v885
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+1)))
	if v887 != 0 {
		v870 = v870 + v885
		v871 = v886
		v872 = v887
		goto L234
	} else {
		goto L240
	}
L238:
	;
	v880 = F_tolower(m, v878)
	mBase = m.M
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	v882 = F_tolower(m, v881)
	mBase = m.M
	if v880 == v882 {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870))))
	v892 = v871
	v893 = v884
	goto L233
L240:
	;
	goto L235
L241:
	;
	if v22 <= v30 {
		goto L228
	} else {
		goto L242
	}
L242:
	;
	v905 = int32(2)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v864+v30<<(uint(v905)%32))))
	v909 = F_objectGetVal(m, v908)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v909
	v2553 = v23 + v905
	goto L7
L243:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v916+v33)))
	v967 = F_objectGetVal(m, v966)
	mBase = m.M
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v969 = int32(_a1710)
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	if v972 != 0 {
		goto L261
	} else {
		goto L262
	}
L244:
	;
	if v952-v954 != 0 {
		goto L243
	} else {
		goto L256
	}
L245:
	;
	v952 = F_tolower(m, v948)
	mBase = m.M
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949))))
	v954 = F_tolower(m, v953)
	mBase = m.M
	goto L244
L246:
	;
	v922 = v915
	v923 = v917
	v924 = v920
	goto L249
L247:
	;
	v948 = int32(0)
	v949 = v917
	goto L245
L248:
	;
	v948 = v945 & int32(255)
	v949 = v944
	goto L245
L249:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923))))
	if v926 == int32(0) {
		v944 = v923
		v945 = v924
		goto L248
	} else {
		goto L251
	}
L250:
	;
	v944 = v938
	v945 = int32(0)
	goto L248
L251:
	;
	v930 = v924 & int32(255)
	if v930 == v926 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v937 = int32(1)
	v938 = v923 + v937
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+1)))
	if v939 != 0 {
		v922 = v922 + v937
		v923 = v938
		v924 = v939
		goto L249
	} else {
		goto L255
	}
L253:
	;
	v932 = F_tolower(m, v930)
	mBase = m.M
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v923))))
	v934 = F_tolower(m, v933)
	mBase = m.M
	if v932 == v934 {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	v944 = v923
	v945 = v936
	goto L248
L255:
	;
	goto L250
L256:
	;
	if v22 <= v30 {
		goto L243
	} else {
		goto L257
	}
L257:
	;
	v957 = int32(2)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v916+v30<<(uint(v957)%32))))
	v961 = F_objectGetVal(m, v960)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v961
	v2553 = v23 + v957
	goto L7
L258:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v968+v33)))
	v1019 = F_objectGetVal(m, v1018)
	mBase = m.M
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1021 = int32(_a23)
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019))))
	if v1024 != 0 {
		goto L276
	} else {
		goto L277
	}
L259:
	;
	if v1004-v1006 != 0 {
		goto L258
	} else {
		goto L271
	}
L260:
	;
	v1004 = F_tolower(m, v1000)
	mBase = m.M
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001))))
	v1006 = F_tolower(m, v1005)
	mBase = m.M
	goto L259
L261:
	;
	v974 = v967
	v975 = v969
	v976 = v972
	goto L264
L262:
	;
	v1000 = int32(0)
	v1001 = v969
	goto L260
L263:
	;
	v1000 = v997 & int32(255)
	v1001 = v996
	goto L260
L264:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975))))
	if v978 == int32(0) {
		v996 = v975
		v997 = v976
		goto L263
	} else {
		goto L266
	}
L265:
	;
	v996 = v990
	v997 = int32(0)
	goto L263
L266:
	;
	v982 = v976 & int32(255)
	if v982 == v978 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v989 = int32(1)
	v990 = v975 + v989
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974)+1)))
	if v991 != 0 {
		v974 = v974 + v989
		v975 = v990
		v976 = v991
		goto L264
	} else {
		goto L270
	}
L268:
	;
	v984 = F_tolower(m, v982)
	mBase = m.M
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v975))))
	v986 = F_tolower(m, v985)
	mBase = m.M
	if v984 == v986 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974))))
	v996 = v975
	v997 = v988
	goto L263
L270:
	;
	goto L265
L271:
	;
	if v22 <= v30 {
		goto L258
	} else {
		goto L272
	}
L272:
	;
	v1009 = int32(2)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v968+v30<<(uint(v1009)%32))))
	v1013 = F_objectGetVal(m, v1012)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v1013
	v2553 = v23 + v1009
	goto L7
L273:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1020+v33)))
	v1125 = F_objectGetVal(m, v1124)
	mBase = m.M
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1127 = int32(_a1711)
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125))))
	if v1130 != 0 {
		goto L301
	} else {
		goto L302
	}
L274:
	;
	if v1056-v1058 != 0 {
		goto L273
	} else {
		goto L286
	}
L275:
	;
	v1056 = F_tolower(m, v1052)
	mBase = m.M
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1053))))
	v1058 = F_tolower(m, v1057)
	mBase = m.M
	goto L274
L276:
	;
	v1026 = v1019
	v1027 = v1021
	v1028 = v1024
	goto L279
L277:
	;
	v1052 = int32(0)
	v1053 = v1021
	goto L275
L278:
	;
	v1052 = v1049 & int32(255)
	v1053 = v1048
	goto L275
L279:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027))))
	if v1030 == int32(0) {
		v1048 = v1027
		v1049 = v1028
		goto L278
	} else {
		goto L281
	}
L280:
	;
	v1048 = v1042
	v1049 = int32(0)
	goto L278
L281:
	;
	v1034 = v1028 & int32(255)
	if v1034 == v1030 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1041 = int32(1)
	v1042 = v1027 + v1041
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026)+1)))
	if v1043 != 0 {
		v1026 = v1026 + v1041
		v1027 = v1042
		v1028 = v1043
		goto L279
	} else {
		goto L285
	}
L283:
	;
	v1036 = F_tolower(m, v1034)
	mBase = m.M
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1027))))
	v1038 = F_tolower(m, v1037)
	mBase = m.M
	if v1036 == v1038 {
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026))))
	v1048 = v1027
	v1049 = v1040
	goto L278
L285:
	;
	goto L280
L286:
	;
	if v22 <= v30 {
		goto L273
	} else {
		goto L287
	}
L287:
	;
	v1062 = v30 << (uint(int32(2)) % 32)
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1020+v1062)))
	v1065 = F_objectGetVal(m, v1064)
	mBase = m.M
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1067+v1062)))
	v1070 = F_objectGetVal(m, v1069)
	mBase = m.M
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070+int32(-1)))))
	switch v1073 & int32(7) {
	case 0:
		goto L293
	case 1:
		goto L292
	case 2:
		goto L291
	case 3:
		goto L290
	case 4:
		goto L289
	default:
		v1090 = int32(0)
		goto L288
	}
L288:
	;
	v1091 = int32(0)
	v1092 = m.G0
	v1093 = int32(16)
	v1094 = v1092 - v1093
	m.G0 = v1094
	*(*int32)(unsafe.Add(mBase, uint32(v1094)+12)) = v1091
	v1099 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1102 = F_raxFind(m, v1099, v1065, v1090, v1094+int32(12))
	mBase = m.M
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+12))
	m.G0 = v1094 + v1093
	goto L294
L289:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1070+int32(-17))))
	v1090 = v1089
	goto L288
L290:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1070+int32(-9))))
	v1090 = v1086
	goto L288
L291:
	;
	v1083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1070+int32(-5)))))
	v1090 = v1083
	goto L288
L292:
	;
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070+int32(-3)))))
	v1090 = v1080
	goto L288
L293:
	;
	v1090 = int32(base.Ui32(v1073) >> (uint(int32(3)) % 32))
	goto L288
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v1103
	if v1103 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v2553 = v23 + int32(2)
	goto L7
L296:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1108+v30<<(uint(int32(2))%32))))
	v1113 = F_objectGetVal(m, v1112)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v1113
	F_addReplyErrorFormat(m, l0, int32(_a1712), v13+int32(32))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L24
	} else {
		goto L297
	}
L297:
	;
	v2577 = int32(-1)
	goto L1
L298:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1126+v33)))
	v1235 = F_objectGetVal(m, v1234)
	mBase = m.M
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1237 = int32(_a1713)
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	if v1240 != 0 {
		goto L327
	} else {
		goto L328
	}
L299:
	;
	if v1162-v1164 != 0 {
		goto L298
	} else {
		goto L311
	}
L300:
	;
	v1162 = F_tolower(m, v1158)
	mBase = m.M
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159))))
	v1164 = F_tolower(m, v1163)
	mBase = m.M
	goto L299
L301:
	;
	v1132 = v1125
	v1133 = v1127
	v1134 = v1130
	goto L304
L302:
	;
	v1158 = int32(0)
	v1159 = v1127
	goto L300
L303:
	;
	v1158 = v1155 & int32(255)
	v1159 = v1154
	goto L300
L304:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133))))
	if v1136 == int32(0) {
		v1154 = v1133
		v1155 = v1134
		goto L303
	} else {
		goto L306
	}
L305:
	;
	v1154 = v1148
	v1155 = int32(0)
	goto L303
L306:
	;
	v1140 = v1134 & int32(255)
	if v1140 == v1136 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1147 = int32(1)
	v1148 = v1133 + v1147
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+1)))
	if v1149 != 0 {
		v1132 = v1132 + v1147
		v1133 = v1148
		v1134 = v1149
		goto L304
	} else {
		goto L310
	}
L308:
	;
	v1142 = F_tolower(m, v1140)
	mBase = m.M
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133))))
	v1144 = F_tolower(m, v1143)
	mBase = m.M
	if v1142 == v1144 {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132))))
	v1154 = v1133
	v1155 = v1146
	goto L303
L310:
	;
	goto L305
L311:
	;
	if v22 <= v30 {
		goto L298
	} else {
		goto L312
	}
L312:
	;
	v1168 = v30 << (uint(int32(2)) % 32)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1126+v1168)))
	v1171 = F_objectGetVal(m, v1170)
	mBase = m.M
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1172+v1168)))
	v1175 = F_objectGetVal(m, v1174)
	mBase = m.M
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175+int32(-1)))))
	switch v1181 & int32(7) {
	case 0:
		goto L319
	case 1:
		goto L318
	case 2:
		goto L317
	case 3:
		goto L316
	case 4:
		goto L315
	default:
		v1198 = int32(0)
		goto L314
	}
L313:
	;
	v1201 = int32(0)
	v1202 = m.G0
	v1203 = int32(16)
	v1204 = v1202 - v1203
	m.G0 = v1204
	*(*int32)(unsafe.Add(mBase, uint32(v1204)+12)) = v1201
	v1209 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1212 = F_raxFind(m, v1209, v1171, v1200, v1204+int32(12))
	mBase = m.M
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+12))
	m.G0 = v1204 + v1203
	goto L320
L314:
	;
	v1200 = v1198
	goto L313
L315:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1175+int32(-17))))
	v1198 = v1197
	goto L314
L316:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1175+int32(-9))))
	v1200 = v1194
	goto L313
L317:
	;
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175+int32(-5)))))
	v1200 = v1191
	goto L313
L318:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175+int32(-3)))))
	v1200 = v1188
	goto L313
L319:
	;
	v1200 = int32(base.Ui32(v1181) >> (uint(int32(3)) % 32))
	goto L313
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v1213
	if v1213 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v2553 = v23 + int32(2)
	goto L7
L322:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1218+v30<<(uint(int32(2))%32))))
	v1223 = F_objectGetVal(m, v1222)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v1223
	F_addReplyErrorFormat(m, l0, int32(_a1712), v13+int32(48))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L24
	} else {
		goto L323
	}
L323:
	;
	v2577 = int32(-1)
	goto L1
L324:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v33)))
	v1372 = F_objectGetVal(m, v1371)
	mBase = m.M
	v1373 = int32(_a1714)
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1372))))
	if v1376 != 0 {
		goto L370
	} else {
		goto L371
	}
L325:
	;
	if v1272-v1274 != 0 {
		goto L324
	} else {
		goto L337
	}
L326:
	;
	v1272 = F_tolower(m, v1268)
	mBase = m.M
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269))))
	v1274 = F_tolower(m, v1273)
	mBase = m.M
	goto L325
L327:
	;
	v1242 = v1235
	v1243 = v1237
	v1244 = v1240
	goto L330
L328:
	;
	v1268 = int32(0)
	v1269 = v1237
	goto L326
L329:
	;
	v1268 = v1265 & int32(255)
	v1269 = v1264
	goto L326
L330:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243))))
	if v1246 == int32(0) {
		v1264 = v1243
		v1265 = v1244
		goto L329
	} else {
		goto L332
	}
L331:
	;
	v1264 = v1258
	v1265 = int32(0)
	goto L329
L332:
	;
	v1250 = v1244 & int32(255)
	if v1250 == v1246 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1257 = int32(1)
	v1258 = v1243 + v1257
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)))
	if v1259 != 0 {
		v1242 = v1242 + v1257
		v1243 = v1258
		v1244 = v1259
		goto L330
	} else {
		goto L336
	}
L334:
	;
	v1252 = F_tolower(m, v1250)
	mBase = m.M
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243))))
	v1254 = F_tolower(m, v1253)
	mBase = m.M
	if v1252 == v1254 {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
	v1264 = v1243
	v1265 = v1256
	goto L329
L336:
	;
	goto L331
L337:
	;
	if v22 <= v30 {
		goto L324
	} else {
		goto L338
	}
L338:
	;
	v1278 = v30 << (uint(int32(2)) % 32)
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1278)))
	v1281 = F_objectGetVal(m, v1280)
	mBase = m.M
	v1282 = int32(_a510)
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1281))))
	if v1285 != 0 {
		goto L343
	} else {
		goto L344
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v1366
	v2553 = v23 + int32(2)
	goto L7
L340:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1322+v1278)))
	v1325 = F_objectGetVal(m, v1324)
	mBase = m.M
	v1326 = int32(_a512)
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1325))))
	if v1329 != 0 {
		goto L356
	} else {
		goto L357
	}
L341:
	;
	if v1317-v1319 != 0 {
		goto L340
	} else {
		goto L353
	}
L342:
	;
	v1317 = F_tolower(m, v1313)
	mBase = m.M
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314))))
	v1319 = F_tolower(m, v1318)
	mBase = m.M
	goto L341
L343:
	;
	v1287 = v1281
	v1288 = v1282
	v1289 = v1285
	goto L346
L344:
	;
	v1313 = int32(0)
	v1314 = v1282
	goto L342
L345:
	;
	v1313 = v1310 & int32(255)
	v1314 = v1309
	goto L342
L346:
	;
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288))))
	if v1291 == int32(0) {
		v1309 = v1288
		v1310 = v1289
		goto L345
	} else {
		goto L348
	}
L347:
	;
	v1309 = v1303
	v1310 = int32(0)
	goto L345
L348:
	;
	v1295 = v1289 & int32(255)
	if v1295 == v1291 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1302 = int32(1)
	v1303 = v1288 + v1302
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287)+1)))
	if v1304 != 0 {
		v1287 = v1287 + v1302
		v1288 = v1303
		v1289 = v1304
		goto L346
	} else {
		goto L352
	}
L350:
	;
	v1297 = F_tolower(m, v1295)
	mBase = m.M
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288))))
	v1299 = F_tolower(m, v1298)
	mBase = m.M
	if v1297 == v1299 {
		goto L349
	} else {
		goto L351
	}
L351:
	;
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287))))
	v1309 = v1288
	v1310 = v1301
	goto L345
L352:
	;
	goto L347
L353:
	;
	v1366 = int32(1)
	goto L339
L354:
	;
	if v1361-v1363 != 0 {
		goto L2
	} else {
		goto L366
	}
L355:
	;
	v1361 = F_tolower(m, v1357)
	mBase = m.M
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1358))))
	v1363 = F_tolower(m, v1362)
	mBase = m.M
	goto L354
L356:
	;
	v1331 = v1325
	v1332 = v1326
	v1333 = v1329
	goto L359
L357:
	;
	v1357 = int32(0)
	v1358 = v1326
	goto L355
L358:
	;
	v1357 = v1354 & int32(255)
	v1358 = v1353
	goto L355
L359:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332))))
	if v1335 == int32(0) {
		v1353 = v1332
		v1354 = v1333
		goto L358
	} else {
		goto L361
	}
L360:
	;
	v1353 = v1347
	v1354 = int32(0)
	goto L358
L361:
	;
	v1339 = v1333 & int32(255)
	if v1339 == v1335 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1346 = int32(1)
	v1347 = v1332 + v1346
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331)+1)))
	if v1348 != 0 {
		v1331 = v1331 + v1346
		v1332 = v1347
		v1333 = v1348
		goto L359
	} else {
		goto L365
	}
L363:
	;
	v1341 = F_tolower(m, v1339)
	mBase = m.M
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332))))
	v1343 = F_tolower(m, v1342)
	mBase = m.M
	if v1341 == v1343 {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331))))
	v1353 = v1332
	v1354 = v1345
	goto L358
L365:
	;
	goto L360
L366:
	;
	v1366 = int32(0)
	goto L339
L367:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1434+v33)))
	v1437 = F_objectGetVal(m, v1436)
	mBase = m.M
	v1438 = int32(_a62)
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1437))))
	if v1441 != 0 {
		goto L391
	} else {
		goto L392
	}
L368:
	;
	if v1408-v1410 != 0 {
		goto L367
	} else {
		goto L380
	}
L369:
	;
	v1408 = F_tolower(m, v1404)
	mBase = m.M
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1405))))
	v1410 = F_tolower(m, v1409)
	mBase = m.M
	goto L368
L370:
	;
	v1378 = v1372
	v1379 = v1373
	v1380 = v1376
	goto L373
L371:
	;
	v1404 = int32(0)
	v1405 = v1373
	goto L369
L372:
	;
	v1404 = v1401 & int32(255)
	v1405 = v1400
	goto L369
L373:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1379))))
	if v1382 == int32(0) {
		v1400 = v1379
		v1401 = v1380
		goto L372
	} else {
		goto L375
	}
L374:
	;
	v1400 = v1394
	v1401 = int32(0)
	goto L372
L375:
	;
	v1386 = v1380 & int32(255)
	if v1386 == v1382 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1393 = int32(1)
	v1394 = v1379 + v1393
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378)+1)))
	if v1395 != 0 {
		v1378 = v1378 + v1393
		v1379 = v1394
		v1380 = v1395
		goto L373
	} else {
		goto L379
	}
L377:
	;
	v1388 = F_tolower(m, v1386)
	mBase = m.M
	v1389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1379))))
	v1390 = F_tolower(m, v1389)
	mBase = m.M
	if v1388 == v1390 {
		goto L376
	} else {
		goto L378
	}
L378:
	;
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	v1400 = v1379
	v1401 = v1392
	goto L372
L379:
	;
	goto L374
L380:
	;
	if v22 <= v30 {
		goto L367
	} else {
		goto L381
	}
L381:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1413+v30<<(uint(int32(2))%32))))
	v1421 = F_getLongLongFromObjectOrReply(m, l0, v1417, v13+int32(152), int32(_a1715))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L24
	} else {
		goto L384
	}
L382:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = v1423
	v2553 = v23 + int32(2)
	goto L7
L383:
	;
	v2577 = int32(-1)
	goto L1
L384:
	;
	if v1421 != 0 {
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1423 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	if int64(0) < v1423 {
		goto L382
	} else {
		goto L386
	}
L386:
	;
	F_addReplyError(m, l0, int32(_a1716))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L24
	} else {
		goto L387
	}
L387:
	;
	goto L383
L388:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1555+v33)))
	v1558 = F_objectGetVal(m, v1557)
	mBase = m.M
	v1559 = int32(_a1717)
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1558))))
	if v1562 != 0 {
		goto L428
	} else {
		goto L429
	}
L389:
	;
	if v1473-v1475 != 0 {
		goto L388
	} else {
		goto L401
	}
L390:
	;
	v1473 = F_tolower(m, v1469)
	mBase = m.M
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
	v1475 = F_tolower(m, v1474)
	mBase = m.M
	goto L389
L391:
	;
	v1443 = v1437
	v1444 = v1438
	v1445 = v1441
	goto L394
L392:
	;
	v1469 = int32(0)
	v1470 = v1438
	goto L390
L393:
	;
	v1469 = v1466 & int32(255)
	v1470 = v1465
	goto L390
L394:
	;
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1444))))
	if v1447 == int32(0) {
		v1465 = v1444
		v1466 = v1445
		goto L393
	} else {
		goto L396
	}
L395:
	;
	v1465 = v1459
	v1466 = int32(0)
	goto L393
L396:
	;
	v1451 = v1445 & int32(255)
	if v1451 == v1447 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1458 = int32(1)
	v1459 = v1444 + v1458
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443)+1)))
	if v1460 != 0 {
		v1443 = v1443 + v1458
		v1444 = v1459
		v1445 = v1460
		goto L394
	} else {
		goto L400
	}
L398:
	;
	v1453 = F_tolower(m, v1451)
	mBase = m.M
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1444))))
	v1455 = F_tolower(m, v1454)
	mBase = m.M
	if v1453 == v1455 {
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443))))
	v1465 = v1444
	v1466 = v1457
	goto L393
L400:
	;
	goto L395
L401:
	;
	if v22 <= v30 {
		goto L388
	} else {
		goto L402
	}
L402:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v1478 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v30<<(uint(int32(2))%32))))
	v1490 = F_objectGetVal(m, v1489)
	mBase = m.M
	v1491 = F_sdsnew(m, v1490)
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L24
	} else {
		goto L406
	}
L404:
	;
	F_sdsfree(m, v1478)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L24
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = int32(0)
	goto L403
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v1491
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491+int32(-1)))))
	v1523 = int32(0)
	goto L409
L407:
	;
	v2553 = v23 + int32(2)
	goto L7
L408:
	;
	if v1543 != int32(-1) {
		goto L407
	} else {
		goto L423
	}
L409:
	;
	switch v1512 & int32(7) {
	case 0:
		goto L416
	case 1:
		goto L415
	case 2:
		goto L414
	case 3:
		goto L413
	case 4:
		goto L412
	default:
		v1532 = int32(0)
		goto L411
	}
L410:
	;
	if base.Ui32(v1523) < base.Ui32(v1532) {
		goto L420
	} else {
		goto L421
	}
L411:
	;
	if base.Ui32(v1532) <= base.Ui32(v1523) {
		goto L417
	} else {
		goto L418
	}
L412:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1491+int32(-17))))
	v1532 = v1531
	goto L411
L413:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1491+int32(-9))))
	v1532 = v1530
	goto L411
L414:
	;
	v1529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1491+int32(-5)))))
	v1532 = v1529
	goto L411
L415:
	;
	v1528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491+int32(-3)))))
	v1532 = v1528
	goto L411
L416:
	;
	v1532 = int32(base.Ui32(v1512) >> (uint(int32(3)) % 32))
	goto L411
L417:
	;
	goto L410
L418:
	;
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491+v1523))))
	switch v1535 + int32(-65) {
	case 0, 1, 4, 8, 12, 13, 14, 15, 17, 18, 19, 20, 33, 34, 35, 36, 40, 49, 51, 52, 55:
		goto L419
	default:
		goto L417
	}
L419:
	;
	v1523 = v1523 + int32(1)
	goto L409
L420:
	;
	v1543 = int32(-1)
	goto L422
L421:
	;
	v1543 = int32(0)
	goto L422
L422:
	;
	goto L408
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v1491
	F_addReplyErrorFormat(m, l0, int32(_a1718), v13+int32(64))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L24
	} else {
		goto L424
	}
L424:
	;
	v2577 = int32(-1)
	goto L1
L425:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1676+v33)))
	v1679 = F_objectGetVal(m, v1678)
	mBase = m.M
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1681 = int32(_a373)
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1679))))
	if v1684 != 0 {
		goto L465
	} else {
		goto L466
	}
L426:
	;
	if v1594-v1596 != 0 {
		goto L425
	} else {
		goto L438
	}
L427:
	;
	v1594 = F_tolower(m, v1590)
	mBase = m.M
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1591))))
	v1596 = F_tolower(m, v1595)
	mBase = m.M
	goto L426
L428:
	;
	v1564 = v1558
	v1565 = v1559
	v1566 = v1562
	goto L431
L429:
	;
	v1590 = int32(0)
	v1591 = v1559
	goto L427
L430:
	;
	v1590 = v1587 & int32(255)
	v1591 = v1586
	goto L427
L431:
	;
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
	if v1568 == int32(0) {
		v1586 = v1565
		v1587 = v1566
		goto L430
	} else {
		goto L433
	}
L432:
	;
	v1586 = v1580
	v1587 = int32(0)
	goto L430
L433:
	;
	v1572 = v1566 & int32(255)
	if v1572 == v1568 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1579 = int32(1)
	v1580 = v1565 + v1579
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+1)))
	if v1581 != 0 {
		v1564 = v1564 + v1579
		v1565 = v1580
		v1566 = v1581
		goto L431
	} else {
		goto L437
	}
L435:
	;
	v1574 = F_tolower(m, v1572)
	mBase = m.M
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
	v1576 = F_tolower(m, v1575)
	mBase = m.M
	if v1574 == v1576 {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564))))
	v1586 = v1565
	v1587 = v1578
	goto L430
L437:
	;
	goto L432
L438:
	;
	if v22 <= v30 {
		goto L425
	} else {
		goto L439
	}
L439:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v1599 == int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1606+v30<<(uint(int32(2))%32))))
	v1611 = F_objectGetVal(m, v1610)
	mBase = m.M
	v1612 = F_sdsnew(m, v1611)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L24
	} else {
		goto L443
	}
L441:
	;
	F_sdsfree(m, v1599)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L24
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = int32(0)
	goto L440
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v1612
	v1633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612+int32(-1)))))
	v1644 = int32(0)
	goto L446
L444:
	;
	v2553 = v23 + int32(2)
	goto L7
L445:
	;
	if v1664 != int32(-1) {
		goto L444
	} else {
		goto L460
	}
L446:
	;
	switch v1633 & int32(7) {
	case 0:
		goto L453
	case 1:
		goto L452
	case 2:
		goto L451
	case 3:
		goto L450
	case 4:
		goto L449
	default:
		v1653 = int32(0)
		goto L448
	}
L447:
	;
	if base.Ui32(v1644) < base.Ui32(v1653) {
		goto L457
	} else {
		goto L458
	}
L448:
	;
	if base.Ui32(v1653) <= base.Ui32(v1644) {
		goto L454
	} else {
		goto L455
	}
L449:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1612+int32(-17))))
	v1653 = v1652
	goto L448
L450:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1612+int32(-9))))
	v1653 = v1651
	goto L448
L451:
	;
	v1650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1612+int32(-5)))))
	v1653 = v1650
	goto L448
L452:
	;
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612+int32(-3)))))
	v1653 = v1649
	goto L448
L453:
	;
	v1653 = int32(base.Ui32(v1633) >> (uint(int32(3)) % 32))
	goto L448
L454:
	;
	goto L447
L455:
	;
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1612+v1644))))
	switch v1656 + int32(-65) {
	case 0, 1, 4, 8, 12, 13, 14, 15, 17, 18, 19, 20, 33, 34, 35, 36, 40, 49, 51, 52, 55:
		goto L456
	default:
		goto L454
	}
L456:
	;
	v1644 = v1644 + int32(1)
	goto L446
L457:
	;
	v1664 = int32(-1)
	goto L459
L458:
	;
	v1664 = int32(0)
	goto L459
L459:
	;
	goto L445
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v1612
	F_addReplyErrorFormat(m, l0, int32(_a1719), v13+int32(80))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L24
	} else {
		goto L461
	}
L461:
	;
	v2577 = int32(-1)
	goto L1
L462:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1680+v33)))
	v1731 = F_objectGetVal(m, v1730)
	mBase = m.M
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1733 = int32(_a1720)
	v1736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1731))))
	if v1736 != 0 {
		goto L480
	} else {
		goto L481
	}
L463:
	;
	if v1716-v1718 != 0 {
		goto L462
	} else {
		goto L475
	}
L464:
	;
	v1716 = F_tolower(m, v1712)
	mBase = m.M
	v1717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1713))))
	v1718 = F_tolower(m, v1717)
	mBase = m.M
	goto L463
L465:
	;
	v1686 = v1679
	v1687 = v1681
	v1688 = v1684
	goto L468
L466:
	;
	v1712 = int32(0)
	v1713 = v1681
	goto L464
L467:
	;
	v1712 = v1709 & int32(255)
	v1713 = v1708
	goto L464
L468:
	;
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687))))
	if v1690 == int32(0) {
		v1708 = v1687
		v1709 = v1688
		goto L467
	} else {
		goto L470
	}
L469:
	;
	v1708 = v1702
	v1709 = int32(0)
	goto L467
L470:
	;
	v1694 = v1688 & int32(255)
	if v1694 == v1690 {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v1701 = int32(1)
	v1702 = v1687 + v1701
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686)+1)))
	if v1703 != 0 {
		v1686 = v1686 + v1701
		v1687 = v1702
		v1688 = v1703
		goto L468
	} else {
		goto L474
	}
L472:
	;
	v1696 = F_tolower(m, v1694)
	mBase = m.M
	v1697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687))))
	v1698 = F_tolower(m, v1697)
	mBase = m.M
	if v1696 == v1698 {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686))))
	v1708 = v1687
	v1709 = v1700
	goto L467
L474:
	;
	goto L469
L475:
	;
	if v22 <= v30 {
		goto L462
	} else {
		goto L476
	}
L476:
	;
	v1721 = int32(2)
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1680+v30<<(uint(v1721)%32))))
	v1725 = F_objectGetVal(m, v1724)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v1725
	v2553 = v23 + v1721
	goto L7
L477:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1732+v33)))
	v1783 = F_objectGetVal(m, v1782)
	mBase = m.M
	v1784 = int32(_a1721)
	v1787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1783))))
	if v1787 != 0 {
		goto L495
	} else {
		goto L496
	}
L478:
	;
	if v1768-v1770 != 0 {
		goto L477
	} else {
		goto L490
	}
L479:
	;
	v1768 = F_tolower(m, v1764)
	mBase = m.M
	v1769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1765))))
	v1770 = F_tolower(m, v1769)
	mBase = m.M
	goto L478
L480:
	;
	v1738 = v1731
	v1739 = v1733
	v1740 = v1736
	goto L483
L481:
	;
	v1764 = int32(0)
	v1765 = v1733
	goto L479
L482:
	;
	v1764 = v1761 & int32(255)
	v1765 = v1760
	goto L479
L483:
	;
	v1742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1739))))
	if v1742 == int32(0) {
		v1760 = v1739
		v1761 = v1740
		goto L482
	} else {
		goto L485
	}
L484:
	;
	v1760 = v1754
	v1761 = int32(0)
	goto L482
L485:
	;
	v1746 = v1740 & int32(255)
	if v1746 == v1742 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1753 = int32(1)
	v1754 = v1739 + v1753
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1738)+1)))
	if v1755 != 0 {
		v1738 = v1738 + v1753
		v1739 = v1754
		v1740 = v1755
		goto L483
	} else {
		goto L489
	}
L487:
	;
	v1748 = F_tolower(m, v1746)
	mBase = m.M
	v1749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1739))))
	v1750 = F_tolower(m, v1749)
	mBase = m.M
	if v1748 == v1750 {
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1738))))
	v1760 = v1739
	v1761 = v1752
	goto L482
L489:
	;
	goto L484
L490:
	;
	if v22 <= v30 {
		goto L477
	} else {
		goto L491
	}
L491:
	;
	v1773 = int32(2)
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1732+v30<<(uint(v1773)%32))))
	v1777 = F_objectGetVal(m, v1776)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v1777
	v2553 = v23 + v1773
	goto L7
L492:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1841+v33)))
	v1844 = F_objectGetVal(m, v1843)
	mBase = m.M
	v1845 = int32(_a1722)
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1844))))
	if v1848 != 0 {
		goto L514
	} else {
		goto L515
	}
L493:
	;
	if v1819-v1821 != 0 {
		goto L492
	} else {
		goto L505
	}
L494:
	;
	v1819 = F_tolower(m, v1815)
	mBase = m.M
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816))))
	v1821 = F_tolower(m, v1820)
	mBase = m.M
	goto L493
L495:
	;
	v1789 = v1783
	v1790 = v1784
	v1791 = v1787
	goto L498
L496:
	;
	v1815 = int32(0)
	v1816 = v1784
	goto L494
L497:
	;
	v1815 = v1812 & int32(255)
	v1816 = v1811
	goto L494
L498:
	;
	v1793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790))))
	if v1793 == int32(0) {
		v1811 = v1790
		v1812 = v1791
		goto L497
	} else {
		goto L500
	}
L499:
	;
	v1811 = v1805
	v1812 = int32(0)
	goto L497
L500:
	;
	v1797 = v1791 & int32(255)
	if v1797 == v1793 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1804 = int32(1)
	v1805 = v1790 + v1804
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1789)+1)))
	if v1806 != 0 {
		v1789 = v1789 + v1804
		v1790 = v1805
		v1791 = v1806
		goto L498
	} else {
		goto L504
	}
L502:
	;
	v1799 = F_tolower(m, v1797)
	mBase = m.M
	v1800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790))))
	v1801 = F_tolower(m, v1800)
	mBase = m.M
	if v1799 == v1801 {
		goto L501
	} else {
		goto L503
	}
L503:
	;
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1789))))
	v1811 = v1790
	v1812 = v1803
	goto L497
L504:
	;
	goto L499
L505:
	;
	if v22 <= v30 {
		goto L492
	} else {
		goto L506
	}
L506:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v1824 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1831+v30<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v1835
	F_incrRefCount(m, v1835)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L24
	} else {
		goto L510
	}
L508:
	;
	F_decrRefCount(m, v1824)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L24
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = int32(0)
	goto L507
L510:
	;
	v2553 = v23 + int32(2)
	goto L7
L511:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1902+v33)))
	v1905 = F_objectGetVal(m, v1904)
	mBase = m.M
	v1906 = int32(_a1723)
	v1909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1905))))
	if v1909 != 0 {
		goto L533
	} else {
		goto L534
	}
L512:
	;
	if v1880-v1882 != 0 {
		goto L511
	} else {
		goto L524
	}
L513:
	;
	v1880 = F_tolower(m, v1876)
	mBase = m.M
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877))))
	v1882 = F_tolower(m, v1881)
	mBase = m.M
	goto L512
L514:
	;
	v1850 = v1844
	v1851 = v1845
	v1852 = v1848
	goto L517
L515:
	;
	v1876 = int32(0)
	v1877 = v1845
	goto L513
L516:
	;
	v1876 = v1873 & int32(255)
	v1877 = v1872
	goto L513
L517:
	;
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1851))))
	if v1854 == int32(0) {
		v1872 = v1851
		v1873 = v1852
		goto L516
	} else {
		goto L519
	}
L518:
	;
	v1872 = v1866
	v1873 = int32(0)
	goto L516
L519:
	;
	v1858 = v1852 & int32(255)
	if v1858 == v1854 {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v1865 = int32(1)
	v1866 = v1851 + v1865
	v1867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850)+1)))
	if v1867 != 0 {
		v1850 = v1850 + v1865
		v1851 = v1866
		v1852 = v1867
		goto L517
	} else {
		goto L523
	}
L521:
	;
	v1860 = F_tolower(m, v1858)
	mBase = m.M
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1851))))
	v1862 = F_tolower(m, v1861)
	mBase = m.M
	if v1860 == v1862 {
		goto L520
	} else {
		goto L522
	}
L522:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850))))
	v1872 = v1851
	v1873 = v1864
	goto L516
L523:
	;
	goto L518
L524:
	;
	if v22 <= v30 {
		goto L511
	} else {
		goto L525
	}
L525:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v1885 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1892+v30<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v1896
	F_incrRefCount(m, v1896)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L24
	} else {
		goto L529
	}
L527:
	;
	F_decrRefCount(m, v1885)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L24
	} else {
		goto L528
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = int32(0)
	goto L526
L529:
	;
	v2553 = v23 + int32(2)
	goto L7
L530:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1963+v33)))
	v1966 = F_objectGetVal(m, v1965)
	mBase = m.M
	v1967 = int32(_a1724)
	v1970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1966))))
	if v1970 != 0 {
		goto L552
	} else {
		goto L553
	}
L531:
	;
	if v1941-v1943 != 0 {
		goto L530
	} else {
		goto L543
	}
L532:
	;
	v1941 = F_tolower(m, v1937)
	mBase = m.M
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1938))))
	v1943 = F_tolower(m, v1942)
	mBase = m.M
	goto L531
L533:
	;
	v1911 = v1905
	v1912 = v1906
	v1913 = v1909
	goto L536
L534:
	;
	v1937 = int32(0)
	v1938 = v1906
	goto L532
L535:
	;
	v1937 = v1934 & int32(255)
	v1938 = v1933
	goto L532
L536:
	;
	v1915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912))))
	if v1915 == int32(0) {
		v1933 = v1912
		v1934 = v1913
		goto L535
	} else {
		goto L538
	}
L537:
	;
	v1933 = v1927
	v1934 = int32(0)
	goto L535
L538:
	;
	v1919 = v1913 & int32(255)
	if v1919 == v1915 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v1926 = int32(1)
	v1927 = v1912 + v1926
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911)+1)))
	if v1928 != 0 {
		v1911 = v1911 + v1926
		v1912 = v1927
		v1913 = v1928
		goto L536
	} else {
		goto L542
	}
L540:
	;
	v1921 = F_tolower(m, v1919)
	mBase = m.M
	v1922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912))))
	v1923 = F_tolower(m, v1922)
	mBase = m.M
	if v1921 == v1923 {
		goto L539
	} else {
		goto L541
	}
L541:
	;
	v1925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911))))
	v1933 = v1912
	v1934 = v1925
	goto L535
L542:
	;
	goto L537
L543:
	;
	if v22 <= v30 {
		goto L530
	} else {
		goto L544
	}
L544:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v1946 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1953+v30<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v1957
	F_incrRefCount(m, v1957)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L24
	} else {
		goto L548
	}
L546:
	;
	F_decrRefCount(m, v1946)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L24
	} else {
		goto L547
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = int32(0)
	goto L545
L548:
	;
	v2553 = v23 + int32(2)
	goto L7
L549:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v2024+v33)))
	v2027 = F_objectGetVal(m, v2026)
	mBase = m.M
	v2028 = int32(_a232)
	v2031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2027))))
	if v2031 != 0 {
		goto L571
	} else {
		goto L572
	}
L550:
	;
	if v2002-v2004 != 0 {
		goto L549
	} else {
		goto L562
	}
L551:
	;
	v2002 = F_tolower(m, v1998)
	mBase = m.M
	v2003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1999))))
	v2004 = F_tolower(m, v2003)
	mBase = m.M
	goto L550
L552:
	;
	v1972 = v1966
	v1973 = v1967
	v1974 = v1970
	goto L555
L553:
	;
	v1998 = int32(0)
	v1999 = v1967
	goto L551
L554:
	;
	v1998 = v1995 & int32(255)
	v1999 = v1994
	goto L551
L555:
	;
	v1976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	if v1976 == int32(0) {
		v1994 = v1973
		v1995 = v1974
		goto L554
	} else {
		goto L557
	}
L556:
	;
	v1994 = v1988
	v1995 = int32(0)
	goto L554
L557:
	;
	v1980 = v1974 & int32(255)
	if v1980 == v1976 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v1987 = int32(1)
	v1988 = v1973 + v1987
	v1989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972)+1)))
	if v1989 != 0 {
		v1972 = v1972 + v1987
		v1973 = v1988
		v1974 = v1989
		goto L555
	} else {
		goto L561
	}
L559:
	;
	v1982 = F_tolower(m, v1980)
	mBase = m.M
	v1983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1973))))
	v1984 = F_tolower(m, v1983)
	mBase = m.M
	if v1982 == v1984 {
		goto L558
	} else {
		goto L560
	}
L560:
	;
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972))))
	v1994 = v1973
	v1995 = v1986
	goto L554
L561:
	;
	goto L556
L562:
	;
	if v22 <= v30 {
		goto L549
	} else {
		goto L563
	}
L563:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v2007 == int32(0) {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v2014+v30<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v2018
	F_incrRefCount(m, v2018)
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L24
	} else {
		goto L567
	}
L565:
	;
	F_decrRefCount(m, v2007)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L24
	} else {
		goto L566
	}
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = int32(0)
	goto L564
L567:
	;
	v2553 = v23 + int32(2)
	goto L7
L568:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2098+v33)))
	v2101 = F_objectGetVal(m, v2100)
	mBase = m.M
	v2102 = int32(_a1725)
	v2105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2101))))
	if v2105 != 0 {
		goto L594
	} else {
		goto L595
	}
L569:
	;
	if v2063-v2065 != 0 {
		goto L568
	} else {
		goto L581
	}
L570:
	;
	v2063 = F_tolower(m, v2059)
	mBase = m.M
	v2064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2060))))
	v2065 = F_tolower(m, v2064)
	mBase = m.M
	goto L569
L571:
	;
	v2033 = v2027
	v2034 = v2028
	v2035 = v2031
	goto L574
L572:
	;
	v2059 = int32(0)
	v2060 = v2028
	goto L570
L573:
	;
	v2059 = v2056 & int32(255)
	v2060 = v2055
	goto L570
L574:
	;
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034))))
	if v2037 == int32(0) {
		v2055 = v2034
		v2056 = v2035
		goto L573
	} else {
		goto L576
	}
L575:
	;
	v2055 = v2049
	v2056 = int32(0)
	goto L573
L576:
	;
	v2041 = v2035 & int32(255)
	if v2041 == v2037 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2048 = int32(1)
	v2049 = v2034 + v2048
	v2050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033)+1)))
	if v2050 != 0 {
		v2033 = v2033 + v2048
		v2034 = v2049
		v2035 = v2050
		goto L574
	} else {
		goto L580
	}
L578:
	;
	v2043 = F_tolower(m, v2041)
	mBase = m.M
	v2044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034))))
	v2045 = F_tolower(m, v2044)
	mBase = m.M
	if v2043 == v2045 {
		goto L577
	} else {
		goto L579
	}
L579:
	;
	v2047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033))))
	v2055 = v2034
	v2056 = v2047
	goto L573
L580:
	;
	goto L575
L581:
	;
	if v22 <= v30 {
		goto L568
	} else {
		goto L582
	}
L582:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2068+v30<<(uint(int32(2))%32))))
	v2076 = F_getIntFromObjectOrReply(m, l0, v2072, v13+int32(152), int32(_a1726))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L24
	} else {
		goto L585
	}
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v2080
	v2553 = v23 + int32(2)
	goto L7
L584:
	;
	v2577 = int32(-1)
	goto L1
L585:
	;
	if v2076 != 0 {
		goto L584
	} else {
		goto L586
	}
L586:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	if v2080 < int32(0) {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v2079 + int32(-1)
	F_addReplyErrorFormat(m, l0, int32(_a1727), v13+int32(96))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L24
	} else {
		goto L590
	}
L588:
	;
	if v2080 < v2079 {
		goto L583
	} else {
		goto L589
	}
L589:
	;
	goto L587
L590:
	;
	goto L584
L591:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2172+v33)))
	v2175 = F_objectGetVal(m, v2174)
	mBase = m.M
	v2176 = int32(_a443)
	v2179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2175))))
	if v2179 != 0 {
		goto L617
	} else {
		goto L618
	}
L592:
	;
	if v2137-v2139 != 0 {
		goto L591
	} else {
		goto L604
	}
L593:
	;
	v2137 = F_tolower(m, v2133)
	mBase = m.M
	v2138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134))))
	v2139 = F_tolower(m, v2138)
	mBase = m.M
	goto L592
L594:
	;
	v2107 = v2101
	v2108 = v2102
	v2109 = v2105
	goto L597
L595:
	;
	v2133 = int32(0)
	v2134 = v2102
	goto L593
L596:
	;
	v2133 = v2130 & int32(255)
	v2134 = v2129
	goto L593
L597:
	;
	v2111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2108))))
	if v2111 == int32(0) {
		v2129 = v2108
		v2130 = v2109
		goto L596
	} else {
		goto L599
	}
L598:
	;
	v2129 = v2123
	v2130 = int32(0)
	goto L596
L599:
	;
	v2115 = v2109 & int32(255)
	if v2115 == v2111 {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v2122 = int32(1)
	v2123 = v2108 + v2122
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2107)+1)))
	if v2124 != 0 {
		v2107 = v2107 + v2122
		v2108 = v2123
		v2109 = v2124
		goto L597
	} else {
		goto L603
	}
L601:
	;
	v2117 = F_tolower(m, v2115)
	mBase = m.M
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2108))))
	v2119 = F_tolower(m, v2118)
	mBase = m.M
	if v2117 == v2119 {
		goto L600
	} else {
		goto L602
	}
L602:
	;
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2107))))
	v2129 = v2108
	v2130 = v2121
	goto L596
L603:
	;
	goto L598
L604:
	;
	if v22 <= v30 {
		goto L591
	} else {
		goto L605
	}
L605:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2142+v30<<(uint(int32(2))%32))))
	v2150 = F_getIntFromObjectOrReply(m, l0, v2146, v13+int32(152), int32(_a1728))
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L24
	} else {
		goto L608
	}
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v2154
	v2553 = v23 + int32(2)
	goto L7
L607:
	;
	v2577 = int32(-1)
	goto L1
L608:
	;
	if v2150 != 0 {
		goto L607
	} else {
		goto L609
	}
L609:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v13)+152))
	if v2154 < int32(0) {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v2153 + int32(-1)
	F_addReplyErrorFormat(m, l0, int32(_a1729), v13+int32(112))
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L24
	} else {
		goto L613
	}
L611:
	;
	if v2154 < v2153 {
		goto L606
	} else {
		goto L612
	}
L612:
	;
	goto L610
L613:
	;
	goto L607
L614:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v33)))
	v2299 = F_objectGetVal(m, v2298)
	mBase = m.M
	v2300 = int32(_a1730)
	v2303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2299))))
	if v2303 != 0 {
		goto L654
	} else {
		goto L655
	}
L615:
	;
	if v2211-v2213 != 0 {
		goto L614
	} else {
		goto L627
	}
L616:
	;
	v2211 = F_tolower(m, v2207)
	mBase = m.M
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2208))))
	v2213 = F_tolower(m, v2212)
	mBase = m.M
	goto L615
L617:
	;
	v2181 = v2175
	v2182 = v2176
	v2183 = v2179
	goto L620
L618:
	;
	v2207 = int32(0)
	v2208 = v2176
	goto L616
L619:
	;
	v2207 = v2204 & int32(255)
	v2208 = v2203
	goto L616
L620:
	;
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2182))))
	if v2185 == int32(0) {
		v2203 = v2182
		v2204 = v2183
		goto L619
	} else {
		goto L622
	}
L621:
	;
	v2203 = v2197
	v2204 = int32(0)
	goto L619
L622:
	;
	v2189 = v2183 & int32(255)
	if v2189 == v2185 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v2196 = int32(1)
	v2197 = v2182 + v2196
	v2198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2181)+1)))
	if v2198 != 0 {
		v2181 = v2181 + v2196
		v2182 = v2197
		v2183 = v2198
		goto L620
	} else {
		goto L626
	}
L624:
	;
	v2191 = F_tolower(m, v2189)
	mBase = m.M
	v2192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2182))))
	v2193 = F_tolower(m, v2192)
	mBase = m.M
	if v2191 == v2193 {
		goto L623
	} else {
		goto L625
	}
L625:
	;
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2181))))
	v2203 = v2182
	v2204 = v2195
	goto L619
L626:
	;
	goto L621
L627:
	;
	if v22 <= v30 {
		goto L614
	} else {
		goto L628
	}
L628:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2216 == int32(0) {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2223+v30<<(uint(int32(2))%32))))
	v2228 = F_objectGetVal(m, v2227)
	mBase = m.M
	v2229 = F_sdsnew(m, v2228)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L24
	} else {
		goto L632
	}
L630:
	;
	F_sdsfree(m, v2216)
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L24
	} else {
		goto L631
	}
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = int32(0)
	goto L629
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v2229
	v2251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229+int32(-1)))))
	v2265 = int32(0)
	goto L635
L633:
	;
	v2553 = v23 + int32(2)
	goto L7
L634:
	;
	if v2284 != int32(-1) {
		goto L633
	} else {
		goto L649
	}
L635:
	;
	switch v2251 & int32(7) {
	case 0:
		goto L642
	case 1:
		goto L641
	case 2:
		goto L640
	case 3:
		goto L639
	case 4:
		goto L638
	default:
		v2272 = int32(0)
		goto L637
	}
L636:
	;
	if base.Ui32(v2265) < base.Ui32(v2272) {
		goto L646
	} else {
		goto L647
	}
L637:
	;
	if base.Ui32(v2272) <= base.Ui32(v2265) {
		goto L643
	} else {
		goto L644
	}
L638:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2229+int32(-17))))
	v2272 = v2271
	goto L637
L639:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2229+int32(-9))))
	v2272 = v2270
	goto L637
L640:
	;
	v2269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2229+int32(-5)))))
	v2272 = v2269
	goto L637
L641:
	;
	v2268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229+int32(-3)))))
	v2272 = v2268
	goto L637
L642:
	;
	v2272 = int32(base.Ui32(v2251) >> (uint(int32(3)) % 32))
	goto L637
L643:
	;
	goto L636
L644:
	;
	v2277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229+v2265))))
	if v2277 == int32(114) {
		v2265 = v2265 + int32(1)
		goto L635
	} else {
		goto L645
	}
L645:
	;
	goto L643
L646:
	;
	v2284 = int32(-1)
	goto L648
L647:
	;
	v2284 = int32(0)
	goto L648
L648:
	;
	goto L634
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v2229
	F_addReplyErrorFormat(m, l0, int32(_a1731), v13+int32(128))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L24
	} else {
		goto L650
	}
L650:
	;
	v2577 = int32(-1)
	goto L1
L651:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2420+v33)))
	v2423 = F_objectGetVal(m, v2422)
	mBase = m.M
	v2424 = int32(_a1732)
	v2427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	if v2427 != 0 {
		goto L691
	} else {
		goto L692
	}
L652:
	;
	if v2335-v2337 != 0 {
		goto L651
	} else {
		goto L664
	}
L653:
	;
	v2335 = F_tolower(m, v2331)
	mBase = m.M
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2332))))
	v2337 = F_tolower(m, v2336)
	mBase = m.M
	goto L652
L654:
	;
	v2305 = v2299
	v2306 = v2300
	v2307 = v2303
	goto L657
L655:
	;
	v2331 = int32(0)
	v2332 = v2300
	goto L653
L656:
	;
	v2331 = v2328 & int32(255)
	v2332 = v2327
	goto L653
L657:
	;
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2306))))
	if v2309 == int32(0) {
		v2327 = v2306
		v2328 = v2307
		goto L656
	} else {
		goto L659
	}
L658:
	;
	v2327 = v2321
	v2328 = int32(0)
	goto L656
L659:
	;
	v2313 = v2307 & int32(255)
	if v2313 == v2309 {
		goto L660
	} else {
		goto L661
	}
L660:
	;
	v2320 = int32(1)
	v2321 = v2306 + v2320
	v2322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2305)+1)))
	if v2322 != 0 {
		v2305 = v2305 + v2320
		v2306 = v2321
		v2307 = v2322
		goto L657
	} else {
		goto L663
	}
L661:
	;
	v2315 = F_tolower(m, v2313)
	mBase = m.M
	v2316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2306))))
	v2317 = F_tolower(m, v2316)
	mBase = m.M
	if v2315 == v2317 {
		goto L660
	} else {
		goto L662
	}
L662:
	;
	v2319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2305))))
	v2327 = v2306
	v2328 = v2319
	goto L656
L663:
	;
	goto L658
L664:
	;
	if v22 <= v30 {
		goto L651
	} else {
		goto L665
	}
L665:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v2340 == int32(0) {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2347+v30<<(uint(int32(2))%32))))
	v2352 = F_objectGetVal(m, v2351)
	mBase = m.M
	v2353 = F_sdsnew(m, v2352)
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L24
	} else {
		goto L669
	}
L667:
	;
	F_sdsfree(m, v2340)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L24
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = int32(0)
	goto L666
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v2353
	v2375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353+int32(-1)))))
	v2389 = int32(0)
	goto L672
L670:
	;
	v2553 = v23 + int32(2)
	goto L7
L671:
	;
	if v2408 != int32(-1) {
		goto L670
	} else {
		goto L686
	}
L672:
	;
	switch v2375 & int32(7) {
	case 0:
		goto L679
	case 1:
		goto L678
	case 2:
		goto L677
	case 3:
		goto L676
	case 4:
		goto L675
	default:
		v2396 = int32(0)
		goto L674
	}
L673:
	;
	if base.Ui32(v2389) < base.Ui32(v2396) {
		goto L683
	} else {
		goto L684
	}
L674:
	;
	if base.Ui32(v2396) <= base.Ui32(v2389) {
		goto L680
	} else {
		goto L681
	}
L675:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v2353+int32(-17))))
	v2396 = v2395
	goto L674
L676:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2353+int32(-9))))
	v2396 = v2394
	goto L674
L677:
	;
	v2393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2353+int32(-5)))))
	v2396 = v2393
	goto L674
L678:
	;
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353+int32(-3)))))
	v2396 = v2392
	goto L674
L679:
	;
	v2396 = int32(base.Ui32(v2375) >> (uint(int32(3)) % 32))
	goto L674
L680:
	;
	goto L673
L681:
	;
	v2401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353+v2389))))
	if v2401 == int32(114) {
		v2389 = v2389 + int32(1)
		goto L672
	} else {
		goto L682
	}
L682:
	;
	goto L680
L683:
	;
	v2408 = int32(-1)
	goto L685
L684:
	;
	v2408 = int32(0)
	goto L685
L685:
	;
	goto L671
L686:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v2353
	F_addReplyErrorFormat(m, l0, int32(_a1733), v13+int32(144))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L24
	} else {
		goto L687
	}
L687:
	;
	v2577 = int32(-1)
	goto L1
L688:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2482+v33)))
	v2485 = F_objectGetVal(m, v2484)
	mBase = m.M
	v2486 = int32(_a1734)
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2485))))
	if v2489 != 0 {
		goto L710
	} else {
		goto L711
	}
L689:
	;
	if v2459-v2461 != 0 {
		goto L688
	} else {
		goto L701
	}
L690:
	;
	v2459 = F_tolower(m, v2455)
	mBase = m.M
	v2460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2456))))
	v2461 = F_tolower(m, v2460)
	mBase = m.M
	goto L689
L691:
	;
	v2429 = v2423
	v2430 = v2424
	v2431 = v2427
	goto L694
L692:
	;
	v2455 = int32(0)
	v2456 = v2424
	goto L690
L693:
	;
	v2455 = v2452 & int32(255)
	v2456 = v2451
	goto L690
L694:
	;
	v2433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	if v2433 == int32(0) {
		v2451 = v2430
		v2452 = v2431
		goto L693
	} else {
		goto L696
	}
L695:
	;
	v2451 = v2445
	v2452 = int32(0)
	goto L693
L696:
	;
	v2437 = v2431 & int32(255)
	if v2437 == v2433 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	v2444 = int32(1)
	v2445 = v2430 + v2444
	v2446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2429)+1)))
	if v2446 != 0 {
		v2429 = v2429 + v2444
		v2430 = v2445
		v2431 = v2446
		goto L694
	} else {
		goto L700
	}
L698:
	;
	v2439 = F_tolower(m, v2437)
	mBase = m.M
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	v2441 = F_tolower(m, v2440)
	mBase = m.M
	if v2439 == v2441 {
		goto L697
	} else {
		goto L699
	}
L699:
	;
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2429))))
	v2451 = v2430
	v2452 = v2443
	goto L693
L700:
	;
	goto L695
L701:
	;
	if v22 <= v30 {
		goto L688
	} else {
		goto L702
	}
L702:
	;
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v2464 == int32(0) {
		goto L703
	} else {
		goto L704
	}
L703:
	;
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2471+v30<<(uint(int32(2))%32))))
	v2476 = F_objectGetVal(m, v2475)
	mBase = m.M
	v2477 = F_sdsnew(m, v2476)
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L24
	} else {
		goto L706
	}
L704:
	;
	F_sdsfree(m, v2464)
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L24
	} else {
		goto L705
	}
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = int32(0)
	goto L703
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v2477
	v2553 = v23 + int32(2)
	goto L7
L707:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v2545)
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L24
	} else {
		goto L726
	}
L708:
	;
	if v2521-v2523 != 0 {
		goto L707
	} else {
		goto L720
	}
L709:
	;
	v2521 = F_tolower(m, v2517)
	mBase = m.M
	v2522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2518))))
	v2523 = F_tolower(m, v2522)
	mBase = m.M
	goto L708
L710:
	;
	v2491 = v2485
	v2492 = v2486
	v2493 = v2489
	goto L713
L711:
	;
	v2517 = int32(0)
	v2518 = v2486
	goto L709
L712:
	;
	v2517 = v2514 & int32(255)
	v2518 = v2513
	goto L709
L713:
	;
	v2495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492))))
	if v2495 == int32(0) {
		v2513 = v2492
		v2514 = v2493
		goto L712
	} else {
		goto L715
	}
L714:
	;
	v2513 = v2507
	v2514 = int32(0)
	goto L712
L715:
	;
	v2499 = v2493 & int32(255)
	if v2499 == v2495 {
		goto L716
	} else {
		goto L717
	}
L716:
	;
	v2506 = int32(1)
	v2507 = v2492 + v2506
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491)+1)))
	if v2508 != 0 {
		v2491 = v2491 + v2506
		v2492 = v2507
		v2493 = v2508
		goto L713
	} else {
		goto L719
	}
L717:
	;
	v2501 = F_tolower(m, v2499)
	mBase = m.M
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2492))))
	v2503 = F_tolower(m, v2502)
	mBase = m.M
	if v2501 == v2503 {
		goto L716
	} else {
		goto L718
	}
L718:
	;
	v2505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491))))
	v2513 = v2492
	v2514 = v2505
	goto L712
L719:
	;
	goto L714
L720:
	;
	if v22 <= v30 {
		goto L707
	} else {
		goto L721
	}
L721:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v2526 == int32(0) {
		goto L722
	} else {
		goto L723
	}
L722:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2533+v30<<(uint(int32(2))%32))))
	v2538 = F_objectGetVal(m, v2537)
	mBase = m.M
	v2539 = F_sdsnew(m, v2538)
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L24
	} else {
		goto L725
	}
L723:
	;
	F_sdsfree(m, v2526)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L24
	} else {
		goto L724
	}
L724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+116)) = int32(0)
	goto L722
L725:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+116)) = v2539
	v2553 = v23 + int32(2)
	goto L7
L726:
	;
	v2577 = int32(-1)
	goto L1
L727:
	;
	goto L6
L728:
	;
	v2577 = int32(-1)
	goto L1
}
func F_removeClientFromMemUsageBucket(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	if v4 == int32(0) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v7 - v8
		if l1 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
			F_listDelNode(m, v11, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+296)) = int64(0)
				return
			}
		}
	}
}
func F_removeClientFromTimeoutTable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v9&int32(512) == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v9 & int32(-513)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
		v19 = int64(56)
		v21 = int64(65280)
		v23 = int64(40)
		v26 = int64(16711680)
		v28 = int64(24)
		v30 = int64(4278190080)
		v32 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v18<<(uint(v19)%64) | v18&v21<<(uint(v23)%64) | (v18&v26<<(uint(v28)%64) | v18&v30<<(uint(v32)%64)) | (int64(base.Ui64(v18)>>(uint(v32)%64))&v30 | int64(base.Ui64(v18)>>(uint(v28)%64))&v26 | (int64(base.Ui64(v18)>>(uint(v23)%64))&v21 | int64(base.Ui64(v18)>>(uint(v19)%64))))
		v56 = v7 | int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v56))) = l0
		v61 = *(*int32)(unsafe.Add(mBase, _consts[878]))
		v64 = F_raxRemove(m, v61, v7, int32(16), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_replaceClientCommandVector(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	F_backupAndUpdateClientArgv(m, l0, l1, l2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v10 & int32(-1073741825)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 < int32(1) {
		v39 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = F_lookupCommandOrOriginal(m, v42, v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v18 = v14
	v19 = v7
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v19<<(uint(int32(2))%32))))
	if v25 == int32(0) {
		v34 = v18
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v39 = v34
	goto L3
L7:
	;
	v36 = v19 + int32(1)
	if v36 < v34 {
		v18 = v34
		v19 = v36
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v28 = F_getStringObjectLen(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v28 + v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = v33
	goto L7
L10:
	;
	goto L6
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v43
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1755), int32(_a1630), int32(6089))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_resetClientIOState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+223)) = uint8(v2)
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+222)) = uint8(v2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v12 & int32(-3)
	return
}
func F_resetClientMultiState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v11 < int32(1) {
		v57 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_valkey_free(m, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L17
	}
L5:
	;
	v16 = v5
	v18 = int32(0)
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = v20 + v18*int32(20)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v24 < int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v57 = v54
	goto L4
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	F_valkey_free(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L15
	}
L9:
	;
	v29 = int32(0)
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v29<<(uint(int32(2))%32))))
	F_decrRefCount(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L8
L12:
	;
	return
L13:
	;
	v39 = v29 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v39 < v40 {
		v29 = v39
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v50 = v18 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v50 < v52 {
		v16 = v51
		v18 = v50
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L7
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v61+int32(16)))) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v61+int32(8)))) = v62
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+52)) = v73
	goto L1
}
