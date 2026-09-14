package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_geohashCommand(m *base.Module, l0 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v144 float64
	_ = v144
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v157 float64
	_ = v157
	var v170 float64
	_ = v170
	var v173 float64
	_ = v173
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 float64
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = F_lookupKeyRead(m, v15, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(112)
	return
L2:
	;
	return
L3:
	;
	v21 = F_checkType(m, l0, v18, int32(3))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_addReplyArrayLen(m, l0, v23+int32(-2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v28 < int32(3) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = v13 + int32(64)
	v44 = int32(2)
	goto L8
L8:
	;
	if v18 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L1
L10:
	;
	v313 = v44 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v313 < v314 {
		v44 = v313
		goto L8
	} else {
		goto L51
	}
L11:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v13)+104))
	v68 = v13 + int32(56)
	v69 = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v69)
	if base.F64_lt(v64, float64(1.8446744073709552e+19))&base.F64_ge(v64, float64(0)) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L16
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v44<<(uint(int32(2))%32))))
	v55 = F_objectGetVal(m, v54)
	mBase = m.M
	v58 = F_zsetScore(m, v18, v55, v13+int32(104))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if v58 != int32(-1) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v80
	v82 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(60)))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(57)))) = v82
	v86 = int32(8)
	v87 = v13 + v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v90
	v95 = m.G0
	v96 = int32(16)
	v97 = v95 - v96
	m.G0 = v97
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v97+v86))) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v105
	v107 = F_geohashDecodeToLongLatType(m, v97, v13+int32(80))
	mBase = m.M
	m.G0 = v97 + v96
	goto L21
L18:
	;
	v80 = int64(0)
	goto L17
L19:
	;
	v78 = base.I64_trunc_f64_u(v64)
	v80 = v78
	goto L17
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = int64(4636033603912859648)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = int64(-4587338432941916160)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = int64(4640537203540230144)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = int64(-4582834833314545664)
	v122 = v13 + int32(48)
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v13)+80))
	v124 = *(*float64)(unsafe.Add(mBase, uint32(v13)+88))
	v125 = int32(26)
	v127 = v13 + int32(32)
	if v32 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	if v107 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L10
L24:
	;
	v218 = int32(48)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+30)) = uint16(v218)
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	v221 = base.I32_wrap_i64(v220)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(27))%32)))+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v226)
	v230 = int32(31)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(2))%32))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+29)) = uint8(v234)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(7))%32))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)) = uint8(v242)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(12))%32))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)) = uint8(v250)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(17))%32))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v258)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v221)>>(uint(int32(22))%32))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)) = uint8(v266)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(32))%64)))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)) = uint8(v275)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(37))%64)))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)) = uint8(v284)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(42))%64)))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+21)) = uint8(v293)
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(47))%64)))&v230)+uint32(_consts[338]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v302)
	F_addReplyBulkCBuffer(m, l0, v13+int32(20), int32(11))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L50
	}
L25:
	;
	goto L24
L26:
	;
	if v127 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	goto L28
L28:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	if base.F64_ne(v144, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v122)+8))
	if base.F64_ne(v154, float64(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	if v122 == int32(0) {
		goto L25
	} else {
		goto L34
	}
L31:
	;
	if v122 == int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	if base.F64_ne(v149, float64(0)) != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	goto L29
L35:
	;
	if base.F64_gt(base.F64_abs(v123), float64(180)) != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v122)))
	if base.F64_eq(v157, float64(0)) != 0 {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	if base.F64_gt(base.F64_abs(v124), float64(85.05112878)) != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)) = uint8(v125)
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = int64(0)
	if base.F64_gt(v124, v144) != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	if base.F64_lt(v124, v170) != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	if base.F64_gt(v123, v154) != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v122)))
	if base.F64_lt(v123, v173) != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	v181 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v125)) % 64))
	v182 = base.F64_mul(base.F64_div(base.F64_sub(v123, v173), base.F64_sub(v154, v173)), v181)
	if base.F64_lt(v182, float64(4.294967296e+09))&base.F64_ge(v182, float64(0)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v196 = base.F64_mul(base.F64_div(base.F64_sub(v124, v170), base.F64_sub(v144, v170)), v181)
	if base.F64_lt(v196, float64(4.294967296e+09))&base.F64_ge(v196, float64(0)) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v192 = int32(0)
	goto L44
L46:
	;
	v190 = base.I32_trunc_f64_u(v182)
	v192 = v190
	goto L44
L47:
	;
	v207 = F_interleave64(m, v206, v192)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v207
	goto L25
L48:
	;
	v206 = int32(0)
	goto L47
L49:
	;
	v204 = base.I32_trunc_f64_u(v196)
	v206 = v204
	goto L47
L50:
	;
	goto L10
L51:
	;
	goto L9
}
func F_geohashDecodeToLongLatType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 int64
	_ = v23
	var v29 int64
	_ = v29
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v75 float64
	_ = v75
	var v78 float64
	_ = v78
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v86 int64
	_ = v86
	var v91 int64
	_ = v91
	var v96 int64
	_ = v96
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v112 int32
	_ = v112
	var v115 float64
	_ = v115
	var v129 float64
	_ = v129
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	v3 = int32(0)
	if l1 == v3 {
		v138 = v3
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int64(0) {
			v20 = int32(1)
			v21 = float64(-180)
			v22 = float64(180)
			v23 = int64(1)
			v29 = int64(base.Ui64(v13)>>(uint(v23)%64))&int64(4919131752989213764) | v13&int64(2459565876494606882)
			v36 = (int64(base.Ui64(v29)>>(uint(v23)%64)) | int64(base.Ui64(v29)>>(uint(int64(3))%64))) & int64(1085102592571150095)
			v41 = (int64(base.Ui64(v36)>>(uint(int64(4))%64)) | v36) & int64(71777214294589695)
			v44 = int64(base.Ui64(v41)>>(uint(int64(8))%64)) | v41
			v52 = base.I32_wrap_i64(int64(base.Ui64(v44)>>(uint(int64(16))%64))&int64(4294901760) | v44&int64(65535))
			v59 = base.F64_convert_i64_u(v23 << (uint(base.I64_extend_i32_u(v12)&int64(255)) % 64))
			v61 = float64(360)
			v75 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v52), v59), v61), v21), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v52+v20), v59), v61), v21)), float64(0.5))
			if base.F64_gt(v75, v22) != 0 {
				v78 = v22
			} else {
				v78 = v75
			}
			if base.F64_lt(v78, float64(-180)) != 0 {
				v81 = v21
			} else {
				v81 = v78
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = v81
			v83 = float64(-85.05112878)
			v84 = float64(85.05112878)
			v86 = v13 & int64(6148914691236517205)
			v91 = (int64(base.Ui64(v86)>>(uint(int64(1))%64)) | v86) & int64(3689348814741910323)
			v96 = (int64(base.Ui64(v91)>>(uint(int64(2))%64)) | v91) & int64(1085102592571150095)
			v101 = (int64(base.Ui64(v96)>>(uint(int64(4))%64)) | v96) & int64(71777214294589695)
			v104 = int64(base.Ui64(v101)>>(uint(int64(8))%64)) | v101
			v112 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(int64(16))%64))&int64(4294901760) | v104&int64(65535))
			v115 = float64(170.10225756)
			v129 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v112), v59), v115), v83), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v112+int32(1)), v59), v115), v83)), float64(0.5))
			if base.F64_gt(v129, v84) != 0 {
				v132 = v84
			} else {
				v132 = v129
			}
			if base.F64_lt(v132, float64(-85.05112878)) != 0 {
				v135 = v83
			} else {
				v135 = v132
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v135
			v138 = v20
		} else {
			if v12&int32(255) == int32(0) {
				v138 = v3
			} else {
				v20 = int32(1)
				v21 = float64(-180)
				v22 = float64(180)
				v23 = int64(1)
				v29 = int64(base.Ui64(v13)>>(uint(v23)%64))&int64(4919131752989213764) | v13&int64(2459565876494606882)
				v36 = (int64(base.Ui64(v29)>>(uint(v23)%64)) | int64(base.Ui64(v29)>>(uint(int64(3))%64))) & int64(1085102592571150095)
				v41 = (int64(base.Ui64(v36)>>(uint(int64(4))%64)) | v36) & int64(71777214294589695)
				v44 = int64(base.Ui64(v41)>>(uint(int64(8))%64)) | v41
				v52 = base.I32_wrap_i64(int64(base.Ui64(v44)>>(uint(int64(16))%64))&int64(4294901760) | v44&int64(65535))
				v59 = base.F64_convert_i64_u(v23 << (uint(base.I64_extend_i32_u(v12)&int64(255)) % 64))
				v61 = float64(360)
				v75 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v52), v59), v61), v21), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v52+v20), v59), v61), v21)), float64(0.5))
				if base.F64_gt(v75, v22) != 0 {
					v78 = v22
				} else {
					v78 = v75
				}
				if base.F64_lt(v78, float64(-180)) != 0 {
					v81 = v21
				} else {
					v81 = v78
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v81
				v83 = float64(-85.05112878)
				v84 = float64(85.05112878)
				v86 = v13 & int64(6148914691236517205)
				v91 = (int64(base.Ui64(v86)>>(uint(int64(1))%64)) | v86) & int64(3689348814741910323)
				v96 = (int64(base.Ui64(v91)>>(uint(int64(2))%64)) | v91) & int64(1085102592571150095)
				v101 = (int64(base.Ui64(v96)>>(uint(int64(4))%64)) | v96) & int64(71777214294589695)
				v104 = int64(base.Ui64(v101)>>(uint(int64(8))%64)) | v101
				v112 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(int64(16))%64))&int64(4294901760) | v104&int64(65535))
				v115 = float64(170.10225756)
				v129 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v112), v59), v115), v83), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v112+int32(1)), v59), v115), v83)), float64(0.5))
				if base.F64_gt(v129, v84) != 0 {
					v132 = v84
				} else {
					v132 = v129
				}
				if base.F64_lt(v132, float64(-85.05112878)) != 0 {
					v135 = v83
				} else {
					v135 = v132
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v135
				v138 = v20
			}
		}
	}
	return v138
}
func F_geohashDecodeToLongLatWGS84(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 int64
	_ = v36
	var v42 int64
	_ = v42
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v65 int32
	_ = v65
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v88 float64
	_ = v88
	var v91 float64
	_ = v91
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v109 int64
	_ = v109
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v125 int32
	_ = v125
	var v128 float64
	_ = v128
	var v142 float64
	_ = v142
	var v145 float64
	_ = v145
	var v148 float64
	_ = v148
	var v151 int32
	_ = v151
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0+v8)))
	*(*int64)(unsafe.Add(mBase, uint32(v6+v8))) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v14
	if l1 == v3 {
		v151 = v3
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		if v26 != int64(0) {
			v33 = int32(1)
			v34 = float64(-180)
			v35 = float64(180)
			v36 = int64(1)
			v42 = int64(base.Ui64(v26)>>(uint(v36)%64))&int64(4919131752989213764) | v26&int64(2459565876494606882)
			v49 = (int64(base.Ui64(v42)>>(uint(v36)%64)) | int64(base.Ui64(v42)>>(uint(int64(3))%64))) & int64(1085102592571150095)
			v54 = (int64(base.Ui64(v49)>>(uint(int64(4))%64)) | v49) & int64(71777214294589695)
			v57 = int64(base.Ui64(v54)>>(uint(int64(8))%64)) | v54
			v65 = base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(16))%64))&int64(4294901760) | v57&int64(65535))
			v72 = base.F64_convert_i64_u(v36 << (uint(base.I64_extend_i32_u(v25)&int64(255)) % 64))
			v74 = float64(360)
			v88 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v65), v72), v74), v34), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v65+v33), v72), v74), v34)), float64(0.5))
			if base.F64_gt(v88, v35) != 0 {
				v91 = v35
			} else {
				v91 = v88
			}
			if base.F64_lt(v91, float64(-180)) != 0 {
				v94 = v34
			} else {
				v94 = v91
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = v94
			v96 = float64(-85.05112878)
			v97 = float64(85.05112878)
			v99 = v26 & int64(6148914691236517205)
			v104 = (int64(base.Ui64(v99)>>(uint(int64(1))%64)) | v99) & int64(3689348814741910323)
			v109 = (int64(base.Ui64(v104)>>(uint(int64(2))%64)) | v104) & int64(1085102592571150095)
			v114 = (int64(base.Ui64(v109)>>(uint(int64(4))%64)) | v109) & int64(71777214294589695)
			v117 = int64(base.Ui64(v114)>>(uint(int64(8))%64)) | v114
			v125 = base.I32_wrap_i64(int64(base.Ui64(v117)>>(uint(int64(16))%64))&int64(4294901760) | v117&int64(65535))
			v128 = float64(170.10225756)
			v142 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v125), v72), v128), v96), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v125+int32(1)), v72), v128), v96)), float64(0.5))
			if base.F64_gt(v142, v97) != 0 {
				v145 = v97
			} else {
				v145 = v142
			}
			if base.F64_lt(v145, float64(-85.05112878)) != 0 {
				v148 = v96
			} else {
				v148 = v145
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v148
			v151 = v33
		} else {
			if v25&int32(255) == int32(0) {
				v151 = v3
			} else {
				v33 = int32(1)
				v34 = float64(-180)
				v35 = float64(180)
				v36 = int64(1)
				v42 = int64(base.Ui64(v26)>>(uint(v36)%64))&int64(4919131752989213764) | v26&int64(2459565876494606882)
				v49 = (int64(base.Ui64(v42)>>(uint(v36)%64)) | int64(base.Ui64(v42)>>(uint(int64(3))%64))) & int64(1085102592571150095)
				v54 = (int64(base.Ui64(v49)>>(uint(int64(4))%64)) | v49) & int64(71777214294589695)
				v57 = int64(base.Ui64(v54)>>(uint(int64(8))%64)) | v54
				v65 = base.I32_wrap_i64(int64(base.Ui64(v57)>>(uint(int64(16))%64))&int64(4294901760) | v57&int64(65535))
				v72 = base.F64_convert_i64_u(v36 << (uint(base.I64_extend_i32_u(v25)&int64(255)) % 64))
				v74 = float64(360)
				v88 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v65), v72), v74), v34), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v65+v33), v72), v74), v34)), float64(0.5))
				if base.F64_gt(v88, v35) != 0 {
					v91 = v35
				} else {
					v91 = v88
				}
				if base.F64_lt(v91, float64(-180)) != 0 {
					v94 = v34
				} else {
					v94 = v91
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1))) = v94
				v96 = float64(-85.05112878)
				v97 = float64(85.05112878)
				v99 = v26 & int64(6148914691236517205)
				v104 = (int64(base.Ui64(v99)>>(uint(int64(1))%64)) | v99) & int64(3689348814741910323)
				v109 = (int64(base.Ui64(v104)>>(uint(int64(2))%64)) | v104) & int64(1085102592571150095)
				v114 = (int64(base.Ui64(v109)>>(uint(int64(4))%64)) | v109) & int64(71777214294589695)
				v117 = int64(base.Ui64(v114)>>(uint(int64(8))%64)) | v114
				v125 = base.I32_wrap_i64(int64(base.Ui64(v117)>>(uint(int64(16))%64))&int64(4294901760) | v117&int64(65535))
				v128 = float64(170.10225756)
				v142 = base.F64_mul(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v125), v72), v128), v96), base.F64_add(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v125+int32(1)), v72), v128), v96)), float64(0.5))
				if base.F64_gt(v142, v97) != 0 {
					v145 = v97
				} else {
					v145 = v142
				}
				if base.F64_lt(v145, float64(-85.05112878)) != 0 {
					v148 = v96
				} else {
					v148 = v145
				}
				*(*float64)(unsafe.Add(mBase, uint32(l1)+8)) = v148
				v151 = v33
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v151
}
func F_geohashEncodeType(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v74 float64
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v119 int32
	_ = v119
	v3 = l2
	v5 = int32(0)
	if l3 == v5 {
		v119 = v5
	} else {
		if base.Ui32((v3+int32(-33))&int32(255)) < base.Ui32(int32(224)) {
			v119 = v5
		} else {
			if base.F64_gt(base.F64_abs(l0), float64(180)) != 0 {
				v119 = v5
			} else {
				if base.F64_gt(base.F64_abs(l1), float64(85.05112878)) != 0 {
					v119 = v5
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v3)
					v32 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v3)) % 64))
					v33 = base.F64_mul(base.F64_div(base.F64_add(l1, float64(85.05112878)), float64(170.10225756)), v32)
					if base.F64_lt(v33, float64(4.294967296e+09))&base.F64_ge(v33, float64(0)) == int32(0) {
						v43 = int32(0)
					} else {
						v41 = base.I32_trunc_f64_u(v33)
						v43 = v41
					}
					v44 = base.I64_extend_i32_u(v43)
					v49 = (v44<<(uint(int64(16))%64) | v44) & int64(281470681808895)
					v54 = (v49<<(uint(int64(8))%64) | v49) & int64(71777214294589695)
					v59 = (v54<<(uint(int64(4))%64) | v54) & int64(1085102592571150095)
					v64 = (v59<<(uint(int64(2))%64) | v59) & int64(3689348814741910323)
					v74 = base.F64_mul(base.F64_div(base.F64_add(l0, float64(180)), float64(360)), v32)
					if base.F64_lt(v74, float64(4.294967296e+09))&base.F64_ge(v74, float64(0)) == int32(0) {
						v84 = int32(0)
					} else {
						v82 = base.I32_trunc_f64_u(v74)
						v84 = v82
					}
					v85 = base.I64_extend_i32_u(v84)
					v90 = (v85<<(uint(int64(16))%64) | v85) & int64(281470681808895)
					v95 = (v90<<(uint(int64(8))%64) | v90) & int64(71777214294589695)
					v100 = (v95<<(uint(int64(4))%64) | v95) & int64(1085102592571150095)
					v101 = int64(2)
					v105 = (v100<<(uint(v101)%64) | v100) & int64(3689348814741910323)
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = (v105<<(uint(v101)%64)|v105<<(uint(int64(1))%64))&int64(-6148914691236517206) | (v64<<(uint(int64(1))%64)|v64)&int64(6148914691236517205)
					v119 = int32(1)
				}
			}
		}
	}
	return v119
}
func F_geohashEncodeWGS84(m *base.Module, l0 float64, l1 float64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v74 float64
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v119 int32
	_ = v119
	v3 = l2
	v5 = int32(0)
	if l3 == v5 {
		v119 = v5
	} else {
		if base.Ui32((v3+int32(-33))&int32(255)) < base.Ui32(int32(224)) {
			v119 = v5
		} else {
			if base.F64_gt(base.F64_abs(l0), float64(180)) != 0 {
				v119 = v5
			} else {
				if base.F64_gt(base.F64_abs(l1), float64(85.05112878)) != 0 {
					v119 = v5
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v3)
					v32 = base.F64_convert_i64_u(int64(1) << (uint(base.I64_extend_i32_u(v3)) % 64))
					v33 = base.F64_mul(base.F64_div(base.F64_add(l1, float64(85.05112878)), float64(170.10225756)), v32)
					if base.F64_lt(v33, float64(4.294967296e+09))&base.F64_ge(v33, float64(0)) == int32(0) {
						v43 = int32(0)
					} else {
						v41 = base.I32_trunc_f64_u(v33)
						v43 = v41
					}
					v44 = base.I64_extend_i32_u(v43)
					v49 = (v44<<(uint(int64(16))%64) | v44) & int64(281470681808895)
					v54 = (v49<<(uint(int64(8))%64) | v49) & int64(71777214294589695)
					v59 = (v54<<(uint(int64(4))%64) | v54) & int64(1085102592571150095)
					v64 = (v59<<(uint(int64(2))%64) | v59) & int64(3689348814741910323)
					v74 = base.F64_mul(base.F64_div(base.F64_add(l0, float64(180)), float64(360)), v32)
					if base.F64_lt(v74, float64(4.294967296e+09))&base.F64_ge(v74, float64(0)) == int32(0) {
						v84 = int32(0)
					} else {
						v82 = base.I32_trunc_f64_u(v74)
						v84 = v82
					}
					v85 = base.I64_extend_i32_u(v84)
					v90 = (v85<<(uint(int64(16))%64) | v85) & int64(281470681808895)
					v95 = (v90<<(uint(int64(8))%64) | v90) & int64(71777214294589695)
					v100 = (v95<<(uint(int64(4))%64) | v95) & int64(1085102592571150095)
					v101 = int64(2)
					v105 = (v100<<(uint(v101)%64) | v100) & int64(3689348814741910323)
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = (v105<<(uint(v101)%64)|v105<<(uint(int64(1))%64))&int64(-6148914691236517206) | (v64<<(uint(int64(1))%64)|v64)&int64(6148914691236517205)
					v119 = int32(1)
				}
			}
		}
	}
	return v119
}
func F_geohashGetDistanceIfInRectangle(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 float64
	_ = v100
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v144 float64
	_ = v144
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v172 float64
	_ = v172
	var v178 int64
	_ = v178
	var v183 int32
	_ = v183
	var v204 float64
	_ = v204
	var v208 float64
	_ = v208
	var v211 float64
	_ = v211
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v223 float64
	_ = v223
	var v227 float64
	_ = v227
	var v236 float64
	_ = v236
	var v243 float64
	_ = v243
	var v248 float64
	_ = v248
	var v249 float64
	_ = v249
	var v257 float64
	_ = v257
	var v260 float64
	_ = v260
	var v265 float64
	_ = v265
	var v268 float64
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v288 float64
	_ = v288
	var v292 int32
	_ = v292
	var v293 float64
	_ = v293
	var v294 float64
	_ = v294
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v305 float64
	_ = v305
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v338 float64
	_ = v338
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v344 float64
	_ = v344
	var v347 float64
	_ = v347
	var v349 float64
	_ = v349
	var v351 float64
	_ = v351
	var v354 float64
	_ = v354
	var v357 float64
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v380 float64
	_ = v380
	var v384 int32
	_ = v384
	var v385 float64
	_ = v385
	var v386 float64
	_ = v386
	var v389 float64
	_ = v389
	var v391 float64
	_ = v391
	var v393 float64
	_ = v393
	var v396 float64
	_ = v396
	var v399 float64
	_ = v399
	var v404 float64
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v424 float64
	_ = v424
	var v428 int32
	_ = v428
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v434 float64
	_ = v434
	var v435 float64
	_ = v435
	var v437 float64
	_ = v437
	var v439 float64
	_ = v439
	var v441 float64
	_ = v441
	var v452 float64
	_ = v452
	var v458 int64
	_ = v458
	var v463 int32
	_ = v463
	var v484 float64
	_ = v484
	var v488 float64
	_ = v488
	var v491 float64
	_ = v491
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v498 float64
	_ = v498
	var v503 float64
	_ = v503
	var v507 float64
	_ = v507
	var v516 float64
	_ = v516
	var v523 float64
	_ = v523
	var v528 float64
	_ = v528
	var v529 float64
	_ = v529
	var v537 float64
	_ = v537
	var v542 float64
	_ = v542
	var v550 int32
	_ = v550
	v8 = int32(0)
	v12 = float64(0.017453292519943295)
	v13 = base.F64_mul(l3, v12)
	v15 = base.F64_mul(l5, v12)
	if base.F64_gt(base.F64_mul(base.F64_abs(base.F64_sub(v13, v15)), float64(6.372797560856e+06)), base.F64_mul(l1, float64(0.5))) != 0 {
		v550 = v8
	} else {
		v23 = base.F64_sub(v15, v15)
		v24 = float64(0.017453292519943295)
		v25 = base.F64_mul(l2, v24)
		v27 = base.F64_mul(l4, v24)
		v30 = base.F64_mul(base.F64_sub(v25, v27), float64(0.5))
		v34 = m.G0
		v36 = v34 - int32(16)
		m.G0 = v36
		v43 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v30))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1072243195)) < base.Ui32(v43) {
			if base.Ui32(v43) < base.Ui32(int32(2146435072)) {
				v54 = F___rem_pio2(m, v30, v36)
				mBase = m.M
				v55 = *(*float64)(unsafe.Add(mBase, uint32(v36)+8))
				v56 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
				switch v54 & int32(3) {
				default:
					v60 = F___sin(m, v56, v55, int32(1))
					mBase = m.M
					v67 = v60
				case 1:
					v61 = F___cos(m, v56, v55)
					mBase = m.M
					v67 = v61
				case 2:
					v63 = F___sin(m, v56, v55, int32(1))
					mBase = m.M
					v67 = base.F64_neg(v63)
				case 3:
					v65 = F___cos(m, v56, v55)
					mBase = m.M
					v67 = base.F64_neg(v65)
				}
			} else {
				v67 = base.F64_sub(v30, v30)
			}
		} else {
			if base.Ui32(v43) < base.Ui32(int32(1045430272)) {
				v67 = v30
			} else {
				v50 = F___sin(m, v30, float64(0), int32(0))
				mBase = m.M
				v67 = v50
			}
		}
		m.G0 = v36 + int32(16)
		if base.F64_le(base.F64_abs(v67), float64(1e-15)) == int32(0) {
			v84 = m.G0
			v86 = v84 - int32(16)
			m.G0 = v86
			v93 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v15))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v93) {
				if base.Ui32(v93) < base.Ui32(int32(2146435072)) {
					v104 = F___rem_pio2(m, v15, v86)
					mBase = m.M
					v105 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
					v106 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
					switch v104 & int32(3) {
					default:
						v109 = F___cos(m, v106, v105)
						mBase = m.M
						v119 = v109
					case 1:
						v111 = F___sin(m, v106, v105, int32(1))
						mBase = m.M
						v119 = base.F64_neg(v111)
					case 2:
						v113 = F___cos(m, v106, v105)
						mBase = m.M
						v119 = base.F64_neg(v113)
					case 3:
						v116 = F___sin(m, v106, v105, int32(1))
						mBase = m.M
						v119 = v116
					}
				} else {
					v119 = base.F64_sub(v15, v15)
				}
			} else {
				if base.Ui32(v93) < base.Ui32(int32(1044816030)) {
					v119 = float64(1)
				} else {
					v100 = F___cos(m, v15, float64(0))
					mBase = m.M
					v119 = v100
				}
			}
			m.G0 = v86 + int32(16)
			v124 = base.F64_mul(v23, float64(0.5))
			v128 = m.G0
			v130 = v128 - int32(16)
			m.G0 = v130
			v137 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v124))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v137) {
				if base.Ui32(v137) < base.Ui32(int32(2146435072)) {
					v148 = F___rem_pio2(m, v124, v130)
					mBase = m.M
					v149 = *(*float64)(unsafe.Add(mBase, uint32(v130)+8))
					v150 = *(*float64)(unsafe.Add(mBase, uint32(v130)))
					switch v148 & int32(3) {
					default:
						v154 = F___sin(m, v150, v149, int32(1))
						mBase = m.M
						v161 = v154
					case 1:
						v155 = F___cos(m, v150, v149)
						mBase = m.M
						v161 = v155
					case 2:
						v157 = F___sin(m, v150, v149, int32(1))
						mBase = m.M
						v161 = base.F64_neg(v157)
					case 3:
						v159 = F___cos(m, v150, v149)
						mBase = m.M
						v161 = base.F64_neg(v159)
					}
				} else {
					v161 = base.F64_sub(v124, v124)
				}
			} else {
				if base.Ui32(v137) < base.Ui32(int32(1045430272)) {
					v161 = v124
				} else {
					v144 = F___sin(m, v124, float64(0), int32(0))
					mBase = m.M
					v161 = v144
				}
			}
			m.G0 = v130 + int32(16)
			v172 = base.F64_sqrt(base.F64_add(base.F64_mul(v161, v161), base.F64_mul(v67, base.F64_mul(base.F64_mul(v119, v119), v67))))
			v178 = base.I64_reinterpret_f64(v172)
			v183 = base.I32_wrap_i64(int64(base.Ui64(v178)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v183) < base.Ui32(int32(1072693248)) {
				if base.Ui32(int32(1071644671)) < base.Ui32(v183) {
					v208 = F_fabs(m, v172)
					mBase = m.M
					v211 = base.F64_mul(base.F64_sub(float64(1), v208), float64(0.5))
					v212 = F_sqrt(m, v211)
					mBase = m.M
					v213 = F_R_2(m, v211)
					mBase = m.M
					if base.Ui32(v183) < base.Ui32(int32(1072640819)) {
						v223 = float64(0.7853981633974483)
						v227 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v212) & int64(-4294967296))
						v236 = base.F64_div(base.F64_sub(v211, base.F64_mul(v227, v227)), base.F64_add(v212, v227))
						v243 = base.F64_add(base.F64_sub(base.F64_sub(v223, base.F64_add(v227, v227)), base.F64_sub(base.F64_mul(base.F64_add(v212, v212), v213), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v236, v236)))), v223)
					} else {
						v218 = base.F64_add(base.F64_mul(v212, v213), v212)
						v243 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v218, v218), float64(-6.123233995736766e-17)))
					}
					if v178 < int64(0) {
						v248 = base.F64_neg(v243)
					} else {
						v248 = v243
					}
					v249 = v248
					v257 = v249
				} else {
					if base.Ui32(v183+int32(-1048576)) < base.Ui32(int32(1044381696)) {
						v249 = v172
						v257 = v249
					} else {
						v204 = F_R_2(m, base.F64_mul(v172, v172))
						mBase = m.M
						v257 = base.F64_add(base.F64_mul(v172, v204), v172)
					}
				}
			} else {
				if v183+int32(-1072693248)|base.I32_wrap_i64(v178) != 0 {
					v257 = base.F64_div(float64(0), base.F64_sub(v172, v172))
				} else {
					v257 = base.F64_add(base.F64_mul(v172, float64(1.5707963267948966)), float64(7.52316384526264e-37))
				}
			}
			v260 = base.F64_mul(v257, float64(1.2745595121712e+07))
		} else {
			v260 = base.F64_mul(base.F64_abs(v23), float64(6.372797560856e+06))
		}
		if base.F64_gt(v260, base.F64_mul(l0, float64(0.5))) != 0 {
			v550 = v8
		} else {
			v265 = base.F64_sub(v15, v13)
			v268 = base.F64_mul(base.F64_sub(v27, v25), float64(0.5))
			v272 = m.G0
			v274 = v272 - int32(16)
			m.G0 = v274
			v281 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v268))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1072243195)) < base.Ui32(v281) {
				if base.Ui32(v281) < base.Ui32(int32(2146435072)) {
					v292 = F___rem_pio2(m, v268, v274)
					mBase = m.M
					v293 = *(*float64)(unsafe.Add(mBase, uint32(v274)+8))
					v294 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
					switch v292 & int32(3) {
					default:
						v298 = F___sin(m, v294, v293, int32(1))
						mBase = m.M
						v305 = v298
					case 1:
						v299 = F___cos(m, v294, v293)
						mBase = m.M
						v305 = v299
					case 2:
						v301 = F___sin(m, v294, v293, int32(1))
						mBase = m.M
						v305 = base.F64_neg(v301)
					case 3:
						v303 = F___cos(m, v294, v293)
						mBase = m.M
						v305 = base.F64_neg(v303)
					}
				} else {
					v305 = base.F64_sub(v268, v268)
				}
			} else {
				if base.Ui32(v281) < base.Ui32(int32(1045430272)) {
					v305 = v268
				} else {
					v288 = F___sin(m, v268, float64(0), int32(0))
					mBase = m.M
					v305 = v288
				}
			}
			m.G0 = v274 + int32(16)
			if base.F64_le(base.F64_abs(v305), float64(1e-15)) == int32(0) {
				v322 = m.G0
				v324 = v322 - int32(16)
				m.G0 = v324
				v331 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v13))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v331) {
					if base.Ui32(v331) < base.Ui32(int32(2146435072)) {
						v342 = F___rem_pio2(m, v13, v324)
						mBase = m.M
						v343 = *(*float64)(unsafe.Add(mBase, uint32(v324)+8))
						v344 = *(*float64)(unsafe.Add(mBase, uint32(v324)))
						switch v342 & int32(3) {
						default:
							v347 = F___cos(m, v344, v343)
							mBase = m.M
							v357 = v347
						case 1:
							v349 = F___sin(m, v344, v343, int32(1))
							mBase = m.M
							v357 = base.F64_neg(v349)
						case 2:
							v351 = F___cos(m, v344, v343)
							mBase = m.M
							v357 = base.F64_neg(v351)
						case 3:
							v354 = F___sin(m, v344, v343, int32(1))
							mBase = m.M
							v357 = v354
						}
					} else {
						v357 = base.F64_sub(v13, v13)
					}
				} else {
					if base.Ui32(v331) < base.Ui32(int32(1044816030)) {
						v357 = float64(1)
					} else {
						v338 = F___cos(m, v13, float64(0))
						mBase = m.M
						v357 = v338
					}
				}
				m.G0 = v324 + int32(16)
				v364 = m.G0
				v366 = v364 - int32(16)
				m.G0 = v366
				v373 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v15))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v373) {
					if base.Ui32(v373) < base.Ui32(int32(2146435072)) {
						v384 = F___rem_pio2(m, v15, v366)
						mBase = m.M
						v385 = *(*float64)(unsafe.Add(mBase, uint32(v366)+8))
						v386 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
						switch v384 & int32(3) {
						default:
							v389 = F___cos(m, v386, v385)
							mBase = m.M
							v399 = v389
						case 1:
							v391 = F___sin(m, v386, v385, int32(1))
							mBase = m.M
							v399 = base.F64_neg(v391)
						case 2:
							v393 = F___cos(m, v386, v385)
							mBase = m.M
							v399 = base.F64_neg(v393)
						case 3:
							v396 = F___sin(m, v386, v385, int32(1))
							mBase = m.M
							v399 = v396
						}
					} else {
						v399 = base.F64_sub(v15, v15)
					}
				} else {
					if base.Ui32(v373) < base.Ui32(int32(1044816030)) {
						v399 = float64(1)
					} else {
						v380 = F___cos(m, v15, float64(0))
						mBase = m.M
						v399 = v380
					}
				}
				m.G0 = v366 + int32(16)
				v404 = base.F64_mul(v265, float64(0.5))
				v408 = m.G0
				v410 = v408 - int32(16)
				m.G0 = v410
				v417 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v404))>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072243195)) < base.Ui32(v417) {
					if base.Ui32(v417) < base.Ui32(int32(2146435072)) {
						v428 = F___rem_pio2(m, v404, v410)
						mBase = m.M
						v429 = *(*float64)(unsafe.Add(mBase, uint32(v410)+8))
						v430 = *(*float64)(unsafe.Add(mBase, uint32(v410)))
						switch v428 & int32(3) {
						default:
							v434 = F___sin(m, v430, v429, int32(1))
							mBase = m.M
							v441 = v434
						case 1:
							v435 = F___cos(m, v430, v429)
							mBase = m.M
							v441 = v435
						case 2:
							v437 = F___sin(m, v430, v429, int32(1))
							mBase = m.M
							v441 = base.F64_neg(v437)
						case 3:
							v439 = F___cos(m, v430, v429)
							mBase = m.M
							v441 = base.F64_neg(v439)
						}
					} else {
						v441 = base.F64_sub(v404, v404)
					}
				} else {
					if base.Ui32(v417) < base.Ui32(int32(1045430272)) {
						v441 = v404
					} else {
						v424 = F___sin(m, v404, float64(0), int32(0))
						mBase = m.M
						v441 = v424
					}
				}
				m.G0 = v410 + int32(16)
				v452 = base.F64_sqrt(base.F64_add(base.F64_mul(v441, v441), base.F64_mul(v305, base.F64_mul(base.F64_mul(v357, v399), v305))))
				v458 = base.I64_reinterpret_f64(v452)
				v463 = base.I32_wrap_i64(int64(base.Ui64(v458)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(v463) < base.Ui32(int32(1072693248)) {
					if base.Ui32(int32(1071644671)) < base.Ui32(v463) {
						v488 = F_fabs(m, v452)
						mBase = m.M
						v491 = base.F64_mul(base.F64_sub(float64(1), v488), float64(0.5))
						v492 = F_sqrt(m, v491)
						mBase = m.M
						v493 = F_R_2(m, v491)
						mBase = m.M
						if base.Ui32(v463) < base.Ui32(int32(1072640819)) {
							v503 = float64(0.7853981633974483)
							v507 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v492) & int64(-4294967296))
							v516 = base.F64_div(base.F64_sub(v491, base.F64_mul(v507, v507)), base.F64_add(v492, v507))
							v523 = base.F64_add(base.F64_sub(base.F64_sub(v503, base.F64_add(v507, v507)), base.F64_sub(base.F64_mul(base.F64_add(v492, v492), v493), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v516, v516)))), v503)
						} else {
							v498 = base.F64_add(base.F64_mul(v492, v493), v492)
							v523 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v498, v498), float64(-6.123233995736766e-17)))
						}
						if v458 < int64(0) {
							v528 = base.F64_neg(v523)
						} else {
							v528 = v523
						}
						v529 = v528
						v537 = v529
					} else {
						if base.Ui32(v463+int32(-1048576)) < base.Ui32(int32(1044381696)) {
							v529 = v452
							v537 = v529
						} else {
							v484 = F_R_2(m, base.F64_mul(v452, v452))
							mBase = m.M
							v537 = base.F64_add(base.F64_mul(v452, v484), v452)
						}
					}
				} else {
					if v463+int32(-1072693248)|base.I32_wrap_i64(v458) != 0 {
						v537 = base.F64_div(float64(0), base.F64_sub(v452, v452))
					} else {
						v537 = base.F64_add(base.F64_mul(v452, float64(1.5707963267948966)), float64(7.52316384526264e-37))
					}
				}
				v542 = base.F64_mul(v537, float64(1.2745595121712e+07))
			} else {
				v542 = base.F64_mul(base.F64_abs(v265), float64(6.372797560856e+06))
			}
			*(*float64)(unsafe.Add(mBase, uint32(l6))) = v542
			v550 = int32(1)
		}
	}
	return v550
}
