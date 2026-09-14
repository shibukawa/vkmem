package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"math"
	"unsafe"
)

func F___expo2(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v102 float64
	_ = v102
	var v113 float64
	_ = v113
	v6 = base.F64_add(l0, float64(-1416.0996898839683))
	v13 = F_top12_1(m, v6)
	mBase = m.M
	v15 = v13 & int32(2047)
	v17 = F_top12_1(m, float64(5.551115123125783e-17))
	mBase = m.M
	v20 = F_top12_1(m, float64(512))
	mBase = m.M
	if base.Ui32(v20-v17) <= base.Ui32(v15-v17) {
		if base.Ui32(v17) <= base.Ui32(v15) {
			v28 = F_top12_1(m, float64(1024))
			mBase = m.M
			if base.Ui32(v15) < base.Ui32(v28) {
				v45 = int32(0)
				v46 = int32(0)
				v47 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
				v50 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
				v51 = base.F64_add(base.F64_mul(v6, v47), v50)
				v52 = base.F64_sub(v51, v50)
				v54 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
				v57 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
				v60 = base.F64_add(base.F64_mul(v52, v54), base.F64_add(base.F64_mul(v52, v57), v6))
				v61 = base.F64_mul(v60, v60)
				v64 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
				v67 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
				v71 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
				v74 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
				v77 = base.I64_reinterpret_f64(v51)
				v82 = base.I32_wrap_i64(v77) << (uint(int32(4)) % 32) & int32(2032)
				v85 = *(*float64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1033])))
				v88 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), base.F64_add(base.F64_mul(v60, v64), v67)), base.F64_add(base.F64_mul(v61, base.F64_add(base.F64_mul(v60, v71), v74)), base.F64_add(v85, v60)))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1034])))
				v94 = v91 + v77<<(uint(int64(45))%64)
				if v45 != 0 {
					v96 = base.F64_reinterpret_i64(v94)
					v102 = base.F64_add(base.F64_mul(v96, v88), v96)
					v113 = v102
				} else {
					v95 = F_specialcase_1(m, v88, v94, v77)
					mBase = m.M
					v113 = v95
				}
			} else {
				v31 = base.I64_reinterpret_f64(v6)
				if v31 == int64(-4503599627370496) {
					v102 = float64(0)
					v113 = v102
				} else {
					v35 = F_top12_1(m, math.Float64frombits(uint64(0x7ff0000000000000)))
					mBase = m.M
					if base.Ui32(v15) < base.Ui32(v35) {
						if int64(-1) < v31 {
							v44 = F___math_oflow(m, int32(0))
							mBase = m.M
							v113 = v44
						} else {
							v42 = F___math_uflow(m, int32(0))
							mBase = m.M
							v113 = v42
						}
					} else {
						v113 = base.F64_add(v6, float64(1))
					}
				}
			}
		} else {
			v113 = base.F64_add(v6, float64(1))
		}
	} else {
		v45 = v15
		v46 = int32(0)
		v47 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
		v50 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
		v51 = base.F64_add(base.F64_mul(v6, v47), v50)
		v52 = base.F64_sub(v51, v50)
		v54 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
		v57 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
		v60 = base.F64_add(base.F64_mul(v52, v54), base.F64_add(base.F64_mul(v52, v57), v6))
		v61 = base.F64_mul(v60, v60)
		v64 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
		v67 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
		v71 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
		v74 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
		v77 = base.I64_reinterpret_f64(v51)
		v82 = base.I32_wrap_i64(v77) << (uint(int32(4)) % 32) & int32(2032)
		v85 = *(*float64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1033])))
		v88 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), base.F64_add(base.F64_mul(v60, v64), v67)), base.F64_add(base.F64_mul(v61, base.F64_add(base.F64_mul(v60, v71), v74)), base.F64_add(v85, v60)))
		v91 = *(*int64)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1034])))
		v94 = v91 + v77<<(uint(int64(45))%64)
		if v45 != 0 {
			v96 = base.F64_reinterpret_i64(v94)
			v102 = base.F64_add(base.F64_mul(v96, v88), v96)
			v113 = v102
		} else {
			v95 = F_specialcase_1(m, v88, v94, v77)
			mBase = m.M
			v113 = v95
		}
	}
	return base.F64_mul(base.F64_mul(base.F64_mul(l1, float64(2.247116418577895e+307)), v113), float64(2.247116418577895e+307))
}
func F_emptyDbAsync(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v18 != 0 {
		v19 = int32(14)
	} else {
		v19 = int32(0)
	}
	if v18 != 0 {
		v22 = int32(3)
	} else {
		v22 = int32(1)
	}
	v23 = F_kvstoreCreate(m, int32(_a886), v19, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23
		v27 = F_kvstoreCreate(m, int32(_a887), v19, v22)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27
			v31 = F_kvstoreCreate(m, int32(_a887), v19, v22)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				if v34 == int32(1) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					if v39 != 0 {
						v41 = F_hashtableSize(m, v39)
						mBase = m.M
						v44 = base.I64_extend_i32_u(v41)
					} else {
						v44 = int64(0)
					}
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
					v44 = v37
				}
				v45 = int32(0)
				v47 = *(*int32)(unsafe.Add(mBase, _consts[503]))
				*(*int32)(unsafe.Add(mBase, _consts[503])) = v47 + base.I32_wrap_i64(v44)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
				*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = v12
				F_bioCreateLazyFreeJob(m, int32(552), int32(3), v10)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	}
}
func F_enableBcastTrackingForPrefix(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
	v11 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	v13 = v7 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if l2 == v4 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v229 = int32(0)
	v231 = F_raxTryInsert(m, v225, v7+int32(12), int32(4), v229, v229)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L42
	} else {
		goto L48
	}
L2:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v224 = v223
	goto L1
L3:
	;
	if v207 != 0 {
		goto L2
	} else {
		goto L41
	}
L4:
	;
	if v166 != l2 {
		v207 = v4
		goto L31
	} else {
		goto L32
	}
L5:
	;
	v157 = int32(0)
	v163 = v22
	v164 = v23
	v166 = v157
	v170 = v157
	goto L4
L6:
	;
	if base.Ui32(v23) < base.Ui32(int32(8)) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v34 = v22
	v35 = v23
	v37 = int32(0)
	goto L9
L8:
	;
	v163 = v147
	v164 = v148
	v166 = v150
	v170 = base.B2i32(v153 != int32(0))
	goto L4
L9:
	;
	v43 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	v44 = int32(4)
	v45 = v34 + v44
	if v35&v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v147 = v138
	v148 = v139
	v150 = v123
	v153 = v128
	goto L8
L11:
	;
	v128 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v45+v43+(v128-v43)&int32(3)+v116<<(uint(int32(2))%32))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if base.Ui32(v139) < base.Ui32(int32(8)) {
		v147 = v138
		v148 = v139
		v150 = v123
		v153 = v128
		goto L8
	} else {
		goto L29
	}
L12:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v37))))
	v94 = int32(0)
	goto L23
L13:
	;
	v50 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v37) {
		v83 = v37
		v86 = v50
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v86 == v43 {
		v116 = v50
		v123 = v83
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v60 = v37
	v63 = v50
	goto L16
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v60))))
	if v66 != v68 {
		v83 = v60
		v86 = v63
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v83 = v71
	v86 = v73
	goto L14
L18:
	;
	v70 = int32(1)
	v71 = v60 + v70
	v73 = v63 + v70
	if base.Ui32(v43) <= base.Ui32(v73) {
		v83 = v71
		v86 = v73
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v71) < base.Ui32(l2) {
		v60 = v71
		v63 = v73
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v147 = v34
	v148 = v35
	v150 = v83
	v153 = v86
	goto L8
L22:
	;
	if v94 != v43 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v94))))
	if v107 == v91&int32(255) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v109 = int32(1)
	v111 = v94 + v109
	if v111 != v43 {
		v94 = v111
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v163 = v34
	v164 = v35
	v166 = v37
	v170 = v109
	goto L4
L27:
	;
	v116 = v94
	v123 = v37 + int32(1)
	goto L11
L28:
	;
	v147 = v34
	v148 = v35
	v150 = v37
	v153 = v43
	goto L8
L29:
	;
	if base.Ui32(v123) < base.Ui32(l2) {
		v34 = v138
		v35 = v139
		v37 = v123
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L10
L31:
	;
	goto L3
L32:
	;
	v172 = int32(0)
	if v164&int32(1) == v172 {
		v207 = v172
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v178 = v164 & int32(4)
	if v170&base.B2i32(v178 != int32(0)) != 0 {
		v207 = v172
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v182 = int32(1)
	if v13 == int32(0) {
		v207 = v182
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v164&int32(2) != 0 {
		v204 = int32(0)
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v204
	v207 = v182
	goto L31
L37:
	;
	v188 = int32(3)
	v189 = int32(base.Ui32(v164) >> (uint(v188) % 32))
	if v178 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v199 = int32(4)
	goto L40
L39:
	;
	v199 = v189 << (uint(int32(2)) % 32)
	goto L40
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v163+v189+(int32(0)-v189)&v188+v199+int32(4))))
	v204 = v203
	goto L36
L41:
	;
	v210 = F_valkey_malloc(m, int32(8))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
	;
	v212 = F_raxNew(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v212
	v215 = F_raxNew(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v215
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	v221 = F_raxInsert(m, v219, l1, l2, v210, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v224 = v210
	goto L1
L47:
	;
	m.G0 = v7 + int32(16)
	return
L48:
	;
	if v231 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+100))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+24))
	if v237 != 0 {
		v243 = v237
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v244 = int32(0)
	v246 = F_raxInsert(m, v243, l1, l2, v244, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L42
	} else {
		goto L53
	}
L51:
	;
	v238 = F_raxNew(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L42
	} else {
		goto L52
	}
L52:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+24)) = v238
	v243 = v238
	goto L50
L53:
	;
	goto L47
}
func F_enumConfigGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4&int32(256) == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v17 = v4
		v18 = v16
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v24 = F_configEnumGetName(m, v19, v18, int32(base.Ui32(v17)>>(uint(int32(3))%32))&int32(1))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		v10 = F_getModuleEnumConfig(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v17 = v14
			v18 = v10
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v24 = F_configEnumGetName(m, v19, v18, int32(base.Ui32(v17)>>(uint(int32(3))%32))&int32(1))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_escapeJsonString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_sdscatlen(m, l0, int32(_a1761), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 == int32(0) {
		v68 = v12
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v75 = F_sdscatlen(m, v68, int32(_a1761), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L28
	}
L4:
	;
	v18 = v12
	v19 = l1
	v20 = l2
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	switch v23 + int32(-8) {
	case 0:
		goto L9
	case 1:
		goto L10
	case 2:
		goto L13
	case 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto L8
	case 4:
		goto L12
	case 5:
		goto L11
	case 26:
		goto L14
	default:
		goto L15
	}
L6:
	;
	v68 = v63
	goto L3
L7:
	;
	v67 = v20 + int32(-1)
	if v67 != 0 {
		v18 = v63
		v19 = v19 + int32(1)
		v20 = v67
		goto L5
	} else {
		goto L27
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = base.I32_extend8_s(v23)
	if base.Ui32(v23) < base.Ui32(int32(32)) {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v52 = F_sdscatlen(m, v18, int32(_a1762), int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v48 = F_sdscatlen(m, v18, int32(_a1763), int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v44 = F_sdscatlen(m, v18, int32(_a1764), int32(2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v40 = F_sdscatlen(m, v18, int32(_a1765), int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v36 = F_sdscatlen(m, v18, int32(_a1766), int32(2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v23
	v32 = F_sdscatprintf(m, v18, int32(_a1767), v8+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	if v23 != int32(92) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v63 = v32
	goto L7
L18:
	;
	v63 = v36
	goto L7
L19:
	;
	v63 = v40
	goto L7
L20:
	;
	v63 = v44
	goto L7
L21:
	;
	v63 = v48
	goto L7
L22:
	;
	v63 = v52
	goto L7
L23:
	;
	v60 = int32(_a1768)
	goto L25
L24:
	;
	v60 = int32(_a1769)
	goto L25
L25:
	;
	v61 = F_sdscatprintf(m, v18, v60, v8)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v63 = v61
	goto L7
L27:
	;
	goto L6
L28:
	;
	m.G0 = v8 + int32(32)
	return v75
}
func F_evalRoCommand(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_evalCommand(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_evalScriptsDict(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	return v2
}
func F_evictionPoolPopulate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v650 int32
	_ = v650
	v15 = m.G0
	v17 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	v24 = v15 - (v17<<(uint(int32(2))%32)+int32(15))&int32(-16)
	m.G0 = v24
	v26 = F_kvstoreGetFairRandomHashtableIndex(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[416]))
	v37 = F_kvstoreHashtableSampleEntries(m, l1, v26, v24, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L6
	}
L2:
	;
	return int32(0)
L3:
	;
	if v26 != int32(-1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	m.G0 = v15
	return int32(0)
L5:
	;
	m.G0 = v15
	return v37
L6:
	;
	if v37 < int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v42 = l2 + int32(24)
	v47 = int32(0)
	goto L8
L8:
	;
	v58 = int32(2)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24+v47<<(uint(v58)%32))))
	v62 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65&v58 == v62 {
		v85 = v62
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L5
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	if v87&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	v79 = v61 + (v65&int32(4) ^ int32(12)) + v65<<(uint(int32(3))%32)&int32(8)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v85 = v79 + v80 + int32(1)
	goto L11
L13:
	;
	v650 = v47 + int32(1)
	if v650 != v37 {
		v47 = v650
		goto L8
	} else {
		goto L168
	}
L14:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-1)))))
	switch v572 & int32(7) {
	case 0:
		goto L151
	case 1:
		goto L156
	case 2:
		goto L155
	case 3:
		goto L154
	case 4:
		goto L153
	default:
		v598 = int32(0)
		goto L150
	}
L15:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l2)+372))
	v411 = int32(24)
	v412 = v236 + v411
	v416 = (v232 ^ int32(15)) * v411
	if v412 == v236 {
		goto L109
	} else {
		goto L110
	}
L16:
	;
	F__serverPanic_1(m, int32(_a780), int32(135), int32(_a781), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L2
	} else {
		goto L107
	}
L17:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v134 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	if v87 != int32(512) {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	v94 = m.G0
	v95 = int32(16)
	v96 = v94 - v95
	m.G0 = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v99 = int32(8)
	v103 = F_lrulfu_getIdleness(m, int32(base.Ui32(v98)>>(uint(v99)%32)), v96+int32(12))
	mBase = m.M
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v104 | v103<<(uint(v99)%32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	m.G0 = v96 + v95
	goto L20
L20:
	;
	v133 = base.I64_extend_i32_u(v109)
	goto L17
L21:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v119&int32(1) == int32(0) {
		v130 = int64(-1)
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v133 = v130 ^ int64(-1)
	goto L17
L23:
	;
	goto L22
L24:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v61+(v119&int32(4)^int32(12)))))
	v130 = v129
	goto L23
L25:
	;
	v246 = v244 + int32(-1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v134 == v247 {
		goto L63
	} else {
		goto L64
	}
L26:
	;
	v236 = l2 + v232*int32(24)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+8))
	if v237 == int32(0) {
		v566 = v232
		goto L14
	} else {
		goto L61
	}
L27:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l2)+368))
	if v230 != 0 {
		goto L13
	} else {
		goto L60
	}
L28:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui64(v133) <= base.Ui64(v137) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v139 = int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v140 == int32(0) {
		v232 = v139
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	if base.Ui64(v133) <= base.Ui64(v143) {
		v232 = v139
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v145 = int32(2)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v146 == int32(0) {
		v232 = v145
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
	if base.Ui64(v133) <= base.Ui64(v149) {
		v232 = v145
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v151 = int32(3)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	if v152 == int32(0) {
		v232 = v151
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l2)+72))
	if base.Ui64(v133) <= base.Ui64(v155) {
		v232 = v151
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v157 = int32(4)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+104))
	if v158 == int32(0) {
		v232 = v157
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(l2)+96))
	if base.Ui64(v133) <= base.Ui64(v161) {
		v232 = v157
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v163 = int32(5)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v164 == int32(0) {
		v232 = v163
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v167 = *(*int64)(unsafe.Add(mBase, uint32(l2)+120))
	if base.Ui64(v133) <= base.Ui64(v167) {
		v232 = v163
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v169 = int32(6)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v170 == int32(0) {
		v232 = v169
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l2)+144))
	if base.Ui64(v133) <= base.Ui64(v173) {
		v232 = v169
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v175 = int32(7)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+176))
	if v176 == int32(0) {
		v232 = v175
		goto L26
	} else {
		goto L42
	}
L42:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(l2)+168))
	if base.Ui64(v133) <= base.Ui64(v179) {
		v232 = v175
		goto L26
	} else {
		goto L43
	}
L43:
	;
	v181 = int32(8)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l2)+200))
	if v182 == int32(0) {
		v232 = v181
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l2)+192))
	if base.Ui64(v133) <= base.Ui64(v185) {
		v232 = v181
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v187 = int32(9)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	if v188 == int32(0) {
		v232 = v187
		goto L26
	} else {
		goto L46
	}
L46:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(l2)+216))
	if base.Ui64(v133) <= base.Ui64(v191) {
		v232 = v187
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v193 = int32(10)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	if v194 == int32(0) {
		v232 = v193
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l2)+240))
	if base.Ui64(v133) <= base.Ui64(v197) {
		v232 = v193
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v199 = int32(11)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l2)+272))
	if v200 == int32(0) {
		v232 = v199
		goto L26
	} else {
		goto L50
	}
L50:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l2)+264))
	if base.Ui64(v133) <= base.Ui64(v203) {
		v232 = v199
		goto L26
	} else {
		goto L51
	}
L51:
	;
	v205 = int32(12)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l2)+296))
	if v206 == int32(0) {
		v232 = v205
		goto L26
	} else {
		goto L52
	}
L52:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(l2)+288))
	if base.Ui64(v133) <= base.Ui64(v209) {
		v232 = v205
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v211 = int32(13)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)+320))
	if v212 == int32(0) {
		v232 = v211
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(l2)+312))
	if base.Ui64(v133) <= base.Ui64(v215) {
		v232 = v211
		goto L26
	} else {
		goto L55
	}
L55:
	;
	v217 = int32(14)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+344))
	if v218 == int32(0) {
		v232 = v217
		goto L26
	} else {
		goto L56
	}
L56:
	;
	v221 = *(*int64)(unsafe.Add(mBase, uint32(l2)+336))
	if base.Ui64(v133) <= base.Ui64(v221) {
		v232 = v217
		goto L26
	} else {
		goto L57
	}
L57:
	;
	v223 = int32(15)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l2)+368))
	if v224 == int32(0) {
		v232 = v223
		goto L26
	} else {
		goto L58
	}
L58:
	;
	v228 = *(*int64)(unsafe.Add(mBase, uint32(l2)+360))
	if base.Ui64(v133) <= base.Ui64(v228) {
		v232 = v223
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v244 = int32(16)
	goto L25
L60:
	;
	v232 = int32(0)
	goto L26
L61:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l2)+368))
	if v240 == int32(0) {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	v244 = v232
	goto L25
L63:
	;
	v252 = v246 * int32(24)
	if l2 == v42 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	F_sdsfree(m, v134)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400+v252)+12)) = v247
	v566 = v246
	goto L14
L67:
	;
	v400 = l2
	goto L66
L68:
	;
	v256 = v252 + l2
	if base.Ui32(int32(0)-v252<<(uint(int32(1))%32)) < base.Ui32(v42-v256) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v266 = (v42 ^ l2) & int32(3)
	if base.Ui32(v42) <= base.Ui32(l2) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v263 = F___memcpy(m, l2, v42, v252)
	mBase = m.M
	v400 = v263
	goto L66
L71:
	;
	if v372 == int32(0) {
		goto L67
	} else {
		goto L103
	}
L72:
	;
	if base.Ui32(v350) <= base.Ui32(int32(3)) {
		v371 = v349
		v372 = v350
		v373 = v351
		goto L71
	} else {
		goto L99
	}
L73:
	;
	if v266 != 0 {
		v332 = v252
		goto L83
	} else {
		goto L84
	}
L74:
	;
	if v266 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if l2&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v371 = v42
	v372 = v252
	v373 = l2
	goto L71
L77:
	;
	v273 = v42
	v274 = v252
	v275 = l2
	goto L79
L78:
	;
	v349 = v42
	v350 = v252
	v351 = l2
	goto L72
L79:
	;
	if v274 == int32(0) {
		goto L67
	} else {
		goto L81
	}
L81:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	*(*uint8)(unsafe.Add(mBase, uint32(v275))) = uint8(v279)
	v281 = int32(1)
	v282 = v273 + v281
	v284 = v274 + int32(-1)
	v286 = v275 + v281
	if v286&int32(3) == int32(0) {
		v349 = v282
		v350 = v284
		v351 = v286
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v273 = v282
	v274 = v284
	v275 = v286
	goto L79
L83:
	;
	if v332 == int32(0) {
		goto L67
	} else {
		goto L95
	}
L84:
	;
	if v256&int32(3) == int32(0) {
		v312 = v252
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if base.Ui32(v312) <= base.Ui32(int32(3)) {
		v332 = v312
		goto L83
	} else {
		goto L91
	}
L86:
	;
	v297 = v252
	goto L87
L87:
	;
	if v297 == int32(0) {
		goto L67
	} else {
		goto L89
	}
L88:
	;
	v312 = v303
	goto L85
L89:
	;
	v303 = v297 + int32(-1)
	v304 = l2 + v303
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v303))))
	*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v306)
	if v304&int32(3) != 0 {
		v297 = v303
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v319 = v312
	goto L92
L92:
	;
	v323 = v319 + int32(-4)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v42+v323)))
	*(*int32)(unsafe.Add(mBase, uint32(l2+v323))) = v326
	if base.Ui32(int32(3)) < base.Ui32(v323) {
		v319 = v323
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v332 = v323
	goto L83
L94:
	;
	goto L93
L95:
	;
	v339 = v332
	goto L96
L96:
	;
	v343 = v339 + int32(-1)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v343))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v343))) = uint8(v346)
	if v343 != 0 {
		v339 = v343
		goto L96
	} else {
		goto L98
	}
L98:
	;
	goto L67
L99:
	;
	v356 = v349
	v357 = v350
	v358 = v351
	goto L100
L100:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v360
	v362 = int32(4)
	v363 = v356 + v362
	v365 = v358 + v362
	v367 = v357 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v367) {
		v356 = v363
		v357 = v367
		v358 = v365
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v371 = v363
	v372 = v367
	v373 = v365
	goto L71
L102:
	;
	goto L101
L103:
	;
	v378 = v371
	v379 = v372
	v380 = v373
	goto L104
L104:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	*(*uint8)(unsafe.Add(mBase, uint32(v380))) = uint8(v382)
	v384 = int32(1)
	v389 = v379 + int32(-1)
	if v389 != 0 {
		v378 = v378 + v384
		v379 = v389
		v380 = v380 + v384
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L67
L106:
	;
	goto L105
L107:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v410
	v566 = v232
	goto L14
L109:
	;
	goto L108
L110:
	;
	v420 = v416 + v412
	if base.Ui32(int32(0)-v416<<(uint(int32(1))%32)) < base.Ui32(v236-v420) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v430 = (v236 ^ v412) & int32(3)
	if base.Ui32(v236) <= base.Ui32(v412) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v427 = F___memcpy(m, v412, v236, v416)
	mBase = m.M
	goto L108
L113:
	;
	if v536 == int32(0) {
		goto L109
	} else {
		goto L145
	}
L114:
	;
	if base.Ui32(v514) <= base.Ui32(int32(3)) {
		v535 = v513
		v536 = v514
		v537 = v515
		goto L113
	} else {
		goto L141
	}
L115:
	;
	if v430 != 0 {
		v496 = v416
		goto L125
	} else {
		goto L126
	}
L116:
	;
	if v430 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v412&int32(3) != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v535 = v236
	v536 = v416
	v537 = v412
	goto L113
L119:
	;
	v437 = v236
	v438 = v416
	v439 = v412
	goto L121
L120:
	;
	v513 = v236
	v514 = v416
	v515 = v412
	goto L114
L121:
	;
	if v438 == int32(0) {
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	*(*uint8)(unsafe.Add(mBase, uint32(v439))) = uint8(v443)
	v445 = int32(1)
	v446 = v437 + v445
	v448 = v438 + int32(-1)
	v450 = v439 + v445
	if v450&int32(3) == int32(0) {
		v513 = v446
		v514 = v448
		v515 = v450
		goto L114
	} else {
		goto L124
	}
L124:
	;
	v437 = v446
	v438 = v448
	v439 = v450
	goto L121
L125:
	;
	if v496 == int32(0) {
		goto L109
	} else {
		goto L137
	}
L126:
	;
	if v420&int32(3) == int32(0) {
		v476 = v416
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if base.Ui32(v476) <= base.Ui32(int32(3)) {
		v496 = v476
		goto L125
	} else {
		goto L133
	}
L128:
	;
	v461 = v416
	goto L129
L129:
	;
	if v461 == int32(0) {
		goto L109
	} else {
		goto L131
	}
L130:
	;
	v476 = v467
	goto L127
L131:
	;
	v467 = v461 + int32(-1)
	v468 = v412 + v467
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v467))))
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v470)
	if v468&int32(3) != 0 {
		v461 = v467
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v483 = v476
	goto L134
L134:
	;
	v487 = v483 + int32(-4)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v236+v487)))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v487))) = v490
	if base.Ui32(int32(3)) < base.Ui32(v487) {
		v483 = v487
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v496 = v487
	goto L125
L136:
	;
	goto L135
L137:
	;
	v503 = v496
	goto L138
L138:
	;
	v507 = v503 + int32(-1)
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v507))))
	*(*uint8)(unsafe.Add(mBase, uint32(v412+v507))) = uint8(v510)
	if v507 != 0 {
		v503 = v507
		goto L138
	} else {
		goto L140
	}
L140:
	;
	goto L109
L141:
	;
	v520 = v513
	v521 = v514
	v522 = v515
	goto L142
L142:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v524
	v526 = int32(4)
	v527 = v520 + v526
	v529 = v522 + v526
	v531 = v521 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v531) {
		v520 = v527
		v521 = v531
		v522 = v529
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v535 = v527
	v536 = v531
	v537 = v529
	goto L113
L144:
	;
	goto L143
L145:
	;
	v542 = v535
	v543 = v536
	v544 = v537
	goto L146
L146:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v546)
	v548 = int32(1)
	v553 = v543 + int32(-1)
	if v553 != 0 {
		v542 = v542 + v548
		v543 = v553
		v544 = v544 + v548
		goto L146
	} else {
		goto L148
	}
L147:
	;
	goto L109
L148:
	;
	goto L147
L149:
	;
	v639 = l2 + v566*int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v639))) = v133
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+20)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v639)+16)) = v641
	goto L13
L150:
	;
	v601 = l2 + v566*int32(24)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v604 = v598 + int32(1)
	if v604 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L151:
	;
	v598 = int32(base.Ui32(v572) >> (uint(int32(3)) % 32))
	goto L150
L152:
	;
	if v587 < int32(256) {
		v598 = v587
		goto L150
	} else {
		goto L157
	}
L153:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-17))))
	v587 = v586
	goto L152
L154:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(-9))))
	v587 = v583
	goto L152
L155:
	;
	v580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85+int32(-5)))))
	v587 = v580
	goto L152
L156:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+int32(-3)))))
	v598 = v577
	goto L150
L157:
	;
	v593 = F_sdsdup(m, v85)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v566*int32(24))+8)) = v593
	goto L149
L159:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v611 = v609 + int32(-1)
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	switch v612 & int32(7) {
	case 0:
		goto L167
	case 1:
		goto L166
	case 2:
		goto L165
	case 3:
		goto L164
	case 4:
		goto L163
	default:
		goto L162
	}
L160:
	;
	goto L159
L161:
	;
	v607 = F__emscripten_memcpy_bulkmem(m, v602, v85, v604)
	mBase = m.M
	goto L160
L162:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+8)) = v631
	goto L149
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v609+int32(-17)))) = base.I64_extend_i32_u(v598)
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609+int32(-9)))) = v598
	goto L162
L165:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v609+int32(-5)))) = uint16(v598)
	goto L162
L166:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v609+int32(-3)))) = uint8(v598)
	goto L162
L167:
	;
	v616 = v598 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v616)
	goto L162
L168:
	;
	goto L9
}
func F_execve(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[5])) = int32(45)
	return int32(-1)
}
func F_exp_inline(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v46 float64
	_ = v46
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v115 float64
	_ = v115
	var v121 int64
	_ = v121
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	v15 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(52))%64))) & int32(2047)
	v20 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
	if base.Ui32(v15-v20) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v20) {
		v51 = v15
		v53 = int32(0)
		v54 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
		v57 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
		v58 = base.F64_add(base.F64_mul(l0, v54), v57)
		v59 = base.F64_sub(v58, v57)
		v61 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
		v64 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
		v68 = base.F64_add(l1, base.F64_add(base.F64_mul(v59, v61), base.F64_add(base.F64_mul(v59, v64), l0)))
		v69 = base.F64_mul(v68, v68)
		v72 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
		v75 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
		v79 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
		v82 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
		v85 = base.I64_reinterpret_f64(v58)
		v90 = base.I32_wrap_i64(v85) << (uint(int32(4)) % 32) & int32(2032)
		v93 = *(*float64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1033])))
		v96 = base.F64_add(base.F64_mul(base.F64_mul(v69, v69), base.F64_add(base.F64_mul(v68, v72), v75)), base.F64_add(base.F64_mul(v69, base.F64_add(base.F64_mul(v68, v79), v82)), base.F64_add(v93, v68)))
		v99 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1034])))
		v104 = v99 + (v85+base.I64_extend_i32_u(l2))<<(uint(int64(45))%64)
		if v51 != 0 {
			v167 = base.F64_reinterpret_i64(v104)
			return base.F64_add(base.F64_mul(v167, v96), v167)
		} else {
			if v85&int64(2147483648) != int64(0) {
				v121 = v104 + int64(4602678819172646912)
				v122 = base.F64_reinterpret_i64(v121)
				v123 = base.F64_mul(v122, v96)
				v124 = base.F64_add(v123, v122)
				v125 = F_fabs(m, v124)
				mBase = m.M
				if base.F64_lt(v125, float64(1)) == int32(0) {
					v154 = v124
				} else {
					v130 = float64(2.2250738585072014e-308)
					v131 = F_fp_barrier_4(m, v130)
					mBase = m.M
					F_fp_force_eval_2(m, base.F64_mul(v131, v130))
					mBase = m.M
					if base.F64_lt(v124, float64(0)) != 0 {
						v142 = float64(-1)
					} else {
						v142 = float64(1)
					}
					v143 = base.F64_add(v124, v142)
					v150 = base.F64_sub(base.F64_add(v143, base.F64_add(base.F64_add(v123, base.F64_sub(v122, v124)), base.F64_add(v124, base.F64_sub(v142, v143)))), v142)
					if base.F64_eq(v150, float64(0)) != 0 {
						v153 = base.F64_reinterpret_i64(v121 & int64(-9223372036854775807-1))
					} else {
						v153 = v150
					}
					v154 = v153
				}
				v165 = base.F64_mul(v154, float64(2.2250738585072014e-308))
			} else {
				v115 = base.F64_reinterpret_i64(v104 + int64(-4544132024016830464))
				v165 = base.F64_mul(base.F64_add(base.F64_mul(v115, v96), v115), float64(5.486124068793689e+303))
			}
			return v165
		}
	} else {
		if base.Ui32(v20) <= base.Ui32(v15) {
			if base.Ui32(v15) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
				v51 = int32(0)
				v53 = int32(0)
				v54 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
				v57 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
				v58 = base.F64_add(base.F64_mul(l0, v54), v57)
				v59 = base.F64_sub(v58, v57)
				v61 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
				v64 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
				v68 = base.F64_add(l1, base.F64_add(base.F64_mul(v59, v61), base.F64_add(base.F64_mul(v59, v64), l0)))
				v69 = base.F64_mul(v68, v68)
				v72 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
				v75 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
				v79 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
				v82 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
				v85 = base.I64_reinterpret_f64(v58)
				v90 = base.I32_wrap_i64(v85) << (uint(int32(4)) % 32) & int32(2032)
				v93 = *(*float64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1033])))
				v96 = base.F64_add(base.F64_mul(base.F64_mul(v69, v69), base.F64_add(base.F64_mul(v68, v72), v75)), base.F64_add(base.F64_mul(v69, base.F64_add(base.F64_mul(v68, v79), v82)), base.F64_add(v93, v68)))
				v99 = *(*int64)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1034])))
				v104 = v99 + (v85+base.I64_extend_i32_u(l2))<<(uint(int64(45))%64)
				if v51 != 0 {
					v167 = base.F64_reinterpret_i64(v104)
					return base.F64_add(base.F64_mul(v167, v96), v167)
				} else {
					if v85&int64(2147483648) != int64(0) {
						v121 = v104 + int64(4602678819172646912)
						v122 = base.F64_reinterpret_i64(v121)
						v123 = base.F64_mul(v122, v96)
						v124 = base.F64_add(v123, v122)
						v125 = F_fabs(m, v124)
						mBase = m.M
						if base.F64_lt(v125, float64(1)) == int32(0) {
							v154 = v124
						} else {
							v130 = float64(2.2250738585072014e-308)
							v131 = F_fp_barrier_4(m, v130)
							mBase = m.M
							F_fp_force_eval_2(m, base.F64_mul(v131, v130))
							mBase = m.M
							if base.F64_lt(v124, float64(0)) != 0 {
								v142 = float64(-1)
							} else {
								v142 = float64(1)
							}
							v143 = base.F64_add(v124, v142)
							v150 = base.F64_sub(base.F64_add(v143, base.F64_add(base.F64_add(v123, base.F64_sub(v122, v124)), base.F64_add(v124, base.F64_sub(v142, v143)))), v142)
							if base.F64_eq(v150, float64(0)) != 0 {
								v153 = base.F64_reinterpret_i64(v121 & int64(-9223372036854775807-1))
							} else {
								v153 = v150
							}
							v154 = v153
						}
						v165 = base.F64_mul(v154, float64(2.2250738585072014e-308))
					} else {
						v115 = base.F64_reinterpret_i64(v104 + int64(-4544132024016830464))
						v165 = base.F64_mul(base.F64_add(base.F64_mul(v115, v96), v115), float64(5.486124068793689e+303))
					}
					return v165
				}
			} else {
				if int64(-1) < base.I64_reinterpret_f64(l0) {
					v49 = F___math_xflow(m, l2, float64(3.105036184601418e+231))
					mBase = m.M
					return v49
				} else {
					v46 = F___math_xflow(m, l2, float64(1.2882297539194267e-231))
					mBase = m.M
					return v46
				}
			}
		} else {
			v31 = base.F64_add(l0, float64(1))
			if l2 != 0 {
				v33 = base.F64_neg(v31)
			} else {
				v33 = v31
			}
			return v33
		}
	}
}
func F_expand_tilde(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v12 + int32(1)
	goto L7
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v125 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v116 == int32(0) {
		v123 = v105
		goto L1
	} else {
		goto L24
	}
L3:
	;
	goto L2
L4:
	;
	v96 = v91
	goto L20
L5:
	;
	v91 = v82
	goto L4
L7:
	;
	if v14&int32(3) == int32(0) {
		v42 = v14
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v51 = int32(-2139062144)
	if (int32(16843008)-v48|v48)&v51 != v51 {
		v82 = v42
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v29 = v14
	goto L10
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v34 == int32(0) {
		v105 = v29
		goto L3
	} else {
		goto L12
	}
L11:
	;
	v42 = v39
	goto L8
L12:
	;
	if v34 == int32(47) {
		v105 = v29
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v39 = v29 + int32(1)
	if v39&int32(3) != 0 {
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
	v57 = v42
	v60 = v48
	goto L16
L16:
	;
	v63 = v60 ^ int32(791621423)
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 != v66 {
		v82 = v57
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v72 = v57 + int32(4)
	v76 = int32(-2139062144)
	if (v70|(int32(16843008)-v70))&v76 == v76 {
		v57 = v72
		v60 = v70
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v91 = v72
	goto L4
L20:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 == int32(0) {
		v105 = v96
		goto L3
	} else {
		goto L22
	}
L21:
	;
	v105 = v96
	goto L3
L22:
	;
	if v97 != int32(47) {
		v96 = v96 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v119)
	v123 = v105 + int32(1)
	goto L1
L25:
	;
	m.G0 = v10 + int32(32)
	return v238
L26:
	;
	v198 = v195
	v203 = int32(0)
	goto L57
L27:
	;
	if v188 != int32(48) {
		goto L50
	} else {
		goto L51
	}
L28:
	;
	v182 = F___syscall_getuid32(m)
	mBase = m.M
	goto L48
L29:
	;
	goto L47
L30:
	;
	v126 = int32(_a2172)
	v132 = F___strchrnul(m, v126, int32(61))
	mBase = m.M
	if v132 != v126 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v173 != 0 {
		v195 = v173
		goto L26
	} else {
		goto L45
	}
L32:
	;
	v135 = int32(0)
	v136 = v132 - v126
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[1049]))))
	if v138 != 0 {
		v165 = v135
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v173 = int32(0)
	goto L31
L34:
	;
	v173 = v165
	goto L31
L35:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[1021]))
	if v140 == v139 {
		v165 = v135
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v143 == int32(0) {
		v165 = v135
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v147 = v140
	v150 = v143
	goto L39
L38:
	;
	v165 = v153 + int32(1)
	goto L34
L39:
	;
	v151 = F_strncmp(m, v126, v150, v136)
	mBase = m.M
	if v151 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v158 != 0 {
		v147 = v147 + int32(4)
		v150 = v158
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v153 = v152 + v136
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v154 == int32(61) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v165 = v135
	goto L34
L45:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v174 == int32(0) {
		goto L28
	} else {
		goto L46
	}
L46:
	;
	goto L29
L47:
	;
	v188 = int32(44)
	goto L27
L48:
	;
	goto L49
L49:
	;
	v188 = int32(44)
	goto L27
L50:
	;
	if v188 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v238 = int32(1)
	goto L25
L52:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v195 = v194
	goto L26
L53:
	;
	v238 = int32(3)
	goto L25
L54:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v192 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v220))) = uint8(v116)
	if v116 == int32(0) {
		v230 = v220
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v205 == int32(0) {
		v220 = v203
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v217 != 0 {
		v238 = int32(3)
		goto L25
	} else {
		goto L61
	}
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v203))) = uint8(v205)
	v210 = int32(1)
	v211 = v198 + v210
	v213 = v203 + v210
	if v213 != int32(4094) {
		v198 = v211
		v203 = v213
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v220 = int32(4094)
	goto L56
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v230
	v238 = int32(0)
	goto L25
L63:
	;
	v226 = v220 + int32(1)
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v226))) = uint8(v228)
	v230 = v226
	goto L62
}
func F_expm1(m *base.Module, l0 float64) float64 {
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v74 float64
	_ = v74
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v105 float64
	_ = v105
	var v113 float64
	_ = v113
	var v132 float64
	_ = v132
	var v142 float64
	_ = v142
	var v147 float64
	_ = v147
	var v154 float64
	_ = v154
	var v163 float64
	_ = v163
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v13) < base.Ui32(int32(1078159482)) {
		if base.Ui32(v13) < base.Ui32(int32(1071001155)) {
			if base.Ui32(v13) < base.Ui32(int32(1016070144)) {
				v176 = l0
				return v176
			} else {
				v74 = l0
				v75 = int32(0)
				v77 = float64(0)
				v79 = base.F64_mul(v74, float64(0.5))
				v80 = base.F64_mul(v74, v79)
				v96 = base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
				v99 = base.F64_sub(float64(3), base.F64_mul(v96, v79))
				v105 = base.F64_mul(v80, base.F64_div(base.F64_sub(v96, v99), base.F64_sub(float64(6), base.F64_mul(v74, v99))))
				if v75 != 0 {
					v113 = base.F64_sub(base.F64_sub(base.F64_mul(v74, base.F64_sub(v105, v77)), v77), v80)
					switch v75 + int32(1) {
					case 0:
						return base.F64_add(base.F64_mul(base.F64_sub(v74, v113), float64(0.5)), float64(-0.5))
					default:
						v142 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v75+int32(1023)) << (uint(int64(52)) % 64))
						if base.Ui32(v75) < base.Ui32(int32(57)) {
							v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v75) << (uint(int64(52)) % 64))
							if base.Ui32(int32(19)) < base.Ui32(v75) {
								v174 = base.F64_add(base.F64_sub(v74, base.F64_add(v113, v163)), float64(1))
							} else {
								v174 = base.F64_add(base.F64_sub(float64(1), v163), base.F64_sub(v74, v113))
							}
							v176 = base.F64_mul(v174, v142)
							return v176
						} else {
							v147 = base.F64_add(base.F64_sub(v74, v113), float64(1))
							if v75 == int32(1024) {
								v154 = base.F64_mul(base.F64_add(v147, v147), float64(8.98846567431158e+307))
							} else {
								v154 = base.F64_mul(v147, v142)
							}
							return base.F64_add(v154, float64(-1))
						}
					case 2:
						if base.F64_lt(v74, float64(-0.25)) == int32(0) {
							v132 = base.F64_sub(v74, v113)
							return base.F64_add(base.F64_add(v132, v132), float64(1))
						} else {
							return base.F64_mul(base.F64_sub(v113, base.F64_add(v74, float64(0.5))), float64(-2))
						}
					}
				} else {
					return base.F64_sub(v74, base.F64_sub(base.F64_mul(v74, v105), v80))
				}
			}
		} else {
			if base.Ui32(int32(1072734897)) < base.Ui32(v13) {
				v50 = base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0))
				if base.F64_lt(base.F64_abs(v50), float64(2.147483648e+09)) == int32(0) {
					v58 = int32(-2147483648)
				} else {
					v56 = base.I32_trunc_f64_s(v50)
					v58 = v56
				}
				v59 = base.F64_convert_i32_s(v58)
				v65 = v58
				v66 = base.F64_add(l0, base.F64_mul(v59, float64(-0.6931471803691238)))
				v67 = base.F64_mul(v59, float64(1.9082149292705877e-10))
			} else {
				if v8 < int64(0) {
					v65 = int32(-1)
					v66 = base.F64_add(l0, float64(0.6931471803691238))
					v67 = float64(-1.9082149292705877e-10)
				} else {
					v65 = int32(1)
					v66 = base.F64_add(l0, float64(-0.6931471803691238))
					v67 = float64(1.9082149292705877e-10)
				}
			}
			v68 = base.F64_sub(v66, v67)
			v74 = v68
			v75 = v65
			v77 = base.F64_sub(base.F64_sub(v66, v68), v67)
			v79 = base.F64_mul(v74, float64(0.5))
			v80 = base.F64_mul(v74, v79)
			v96 = base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
			v99 = base.F64_sub(float64(3), base.F64_mul(v96, v79))
			v105 = base.F64_mul(v80, base.F64_div(base.F64_sub(v96, v99), base.F64_sub(float64(6), base.F64_mul(v74, v99))))
			if v75 != 0 {
				v113 = base.F64_sub(base.F64_sub(base.F64_mul(v74, base.F64_sub(v105, v77)), v77), v80)
				switch v75 + int32(1) {
				case 0:
					return base.F64_add(base.F64_mul(base.F64_sub(v74, v113), float64(0.5)), float64(-0.5))
				default:
					v142 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v75+int32(1023)) << (uint(int64(52)) % 64))
					if base.Ui32(v75) < base.Ui32(int32(57)) {
						v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v75) << (uint(int64(52)) % 64))
						if base.Ui32(int32(19)) < base.Ui32(v75) {
							v174 = base.F64_add(base.F64_sub(v74, base.F64_add(v113, v163)), float64(1))
						} else {
							v174 = base.F64_add(base.F64_sub(float64(1), v163), base.F64_sub(v74, v113))
						}
						v176 = base.F64_mul(v174, v142)
						return v176
					} else {
						v147 = base.F64_add(base.F64_sub(v74, v113), float64(1))
						if v75 == int32(1024) {
							v154 = base.F64_mul(base.F64_add(v147, v147), float64(8.98846567431158e+307))
						} else {
							v154 = base.F64_mul(v147, v142)
						}
						return base.F64_add(v154, float64(-1))
					}
				case 2:
					if base.F64_lt(v74, float64(-0.25)) == int32(0) {
						v132 = base.F64_sub(v74, v113)
						return base.F64_add(base.F64_add(v132, v132), float64(1))
					} else {
						return base.F64_mul(base.F64_sub(v113, base.F64_add(v74, float64(0.5))), float64(-2))
					}
				}
			} else {
				return base.F64_sub(v74, base.F64_sub(base.F64_mul(v74, v105), v80))
			}
		}
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v176 = l0
			return v176
		} else {
			if int64(0) <= v8 {
				if base.F64_gt(l0, float64(709.782712893384)) == int32(0) {
					v50 = base.F64_add(base.F64_mul(l0, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), l0))
					if base.F64_lt(base.F64_abs(v50), float64(2.147483648e+09)) == int32(0) {
						v58 = int32(-2147483648)
					} else {
						v56 = base.I32_trunc_f64_s(v50)
						v58 = v56
					}
					v59 = base.F64_convert_i32_s(v58)
					v65 = v58
					v66 = base.F64_add(l0, base.F64_mul(v59, float64(-0.6931471803691238)))
					v67 = base.F64_mul(v59, float64(1.9082149292705877e-10))
					v68 = base.F64_sub(v66, v67)
					v74 = v68
					v75 = v65
					v77 = base.F64_sub(base.F64_sub(v66, v68), v67)
					v79 = base.F64_mul(v74, float64(0.5))
					v80 = base.F64_mul(v74, v79)
					v96 = base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, base.F64_add(base.F64_mul(v80, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v99 = base.F64_sub(float64(3), base.F64_mul(v96, v79))
					v105 = base.F64_mul(v80, base.F64_div(base.F64_sub(v96, v99), base.F64_sub(float64(6), base.F64_mul(v74, v99))))
					if v75 != 0 {
						v113 = base.F64_sub(base.F64_sub(base.F64_mul(v74, base.F64_sub(v105, v77)), v77), v80)
						switch v75 + int32(1) {
						case 0:
							return base.F64_add(base.F64_mul(base.F64_sub(v74, v113), float64(0.5)), float64(-0.5))
						default:
							v142 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v75+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(v75) < base.Ui32(int32(57)) {
								v163 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v75) << (uint(int64(52)) % 64))
								if base.Ui32(int32(19)) < base.Ui32(v75) {
									v174 = base.F64_add(base.F64_sub(v74, base.F64_add(v113, v163)), float64(1))
								} else {
									v174 = base.F64_add(base.F64_sub(float64(1), v163), base.F64_sub(v74, v113))
								}
								v176 = base.F64_mul(v174, v142)
								return v176
							} else {
								v147 = base.F64_add(base.F64_sub(v74, v113), float64(1))
								if v75 == int32(1024) {
									v154 = base.F64_mul(base.F64_add(v147, v147), float64(8.98846567431158e+307))
								} else {
									v154 = base.F64_mul(v147, v142)
								}
								return base.F64_add(v154, float64(-1))
							}
						case 2:
							if base.F64_lt(v74, float64(-0.25)) == int32(0) {
								v132 = base.F64_sub(v74, v113)
								return base.F64_add(base.F64_add(v132, v132), float64(1))
							} else {
								return base.F64_mul(base.F64_sub(v113, base.F64_add(v74, float64(0.5))), float64(-2))
							}
						}
					} else {
						return base.F64_sub(v74, base.F64_sub(base.F64_mul(v74, v105), v80))
					}
				} else {
					return base.F64_mul(l0, float64(8.98846567431158e+307))
				}
			} else {
				return float64(-1)
			}
		}
	}
}
func F_extractBoxOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = F_getDoubleFromObjectOrReply(m, l0, v15, v12, int32(_a823))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v54 = v14
			m.G0 = v12 + int32(16)
			return v54
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v25 = F_getDoubleFromObjectOrReply(m, l0, v21, v12+int32(8), int32(_a824))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					v54 = v14
					m.G0 = v12 + int32(16)
					return v54
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
					if base.F64_lt(v27, float64(0)) != 0 {
						F_addReplyError(m, l0, int32(_a825))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v54 = v14
							m.G0 = v12 + int32(16)
							return v54
						}
					} else {
						v30 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
						if base.F64_lt(v30, float64(0)) == int32(0) {
							if l4 == int32(0) {
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l4))) = v27
							}
							if l3 == int32(0) {
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l3))) = v30
							}
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v46 = F_extractUnitOrReply(m, l0, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								if base.F64_lt(v46, float64(0)) != 0 {
									v54 = v14
								} else {
									v50 = int32(0)
									if l2 == v50 {
										v54 = v50
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(l2))) = v46
										v54 = v50
									}
								}
								m.G0 = v12 + int32(16)
								return v54
							}
						} else {
							F_addReplyError(m, l0, int32(_a825))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v54 = v14
								m.G0 = v12 + int32(16)
								return v54
							}
						}
					}
				}
			}
		}
	}
}
