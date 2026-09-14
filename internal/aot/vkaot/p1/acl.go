package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_ACLAddCommandCategory(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if base.Ui32(int32(63)) < base.Ui32(v7) {
		v35 = v3
		return v35
	} else {
		v10 = F_zstrdup(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[2]))
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1]))
			v20 = v15 + v17<<(uint(int32(4))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v10
			if l1 == int64(0) {
				v27 = int64(1) << (uint(base.I64_extend_i32_u(v17)) % 64)
			} else {
				v27 = l1
			}
			*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v27
			v29 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[1])) = v17 + v29
			v35 = v29
			return v35
		}
	}
}
func F_ACLAppendUserForLoading(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 < int32(2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v294
L2:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v67 = F_listSearchKey(m, v65, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v59 = int32(-1)
	if l2 == int32(0) {
		v294 = v59
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(_a23)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v53-v55 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L6:
	;
	v53 = F_tolower(m, v49)
	mBase = m.M
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v55 = F_tolower(m, v54)
	mBase = m.M
	goto L5
L7:
	;
	v23 = v17
	v24 = v18
	v25 = v21
	goto L10
L8:
	;
	v49 = int32(0)
	v50 = v18
	goto L6
L9:
	;
	v49 = v46 & int32(255)
	v50 = v45
	goto L6
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == int32(0) {
		v45 = v24
		v46 = v25
		goto L9
	} else {
		goto L12
	}
L11:
	;
	v45 = v39
	v46 = int32(0)
	goto L9
L12:
	;
	v31 = v25 & int32(255)
	if v31 == v27 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(1)
	v39 = v24 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v40 != 0 {
		v23 = v23 + v38
		v24 = v39
		v25 = v40
		goto L10
	} else {
		goto L16
	}
L14:
	;
	v33 = F_tolower(m, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v35 = F_tolower(m, v34)
	mBase = m.M
	if v33 == v35 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v45 = v24
	v46 = v37
	goto L9
L16:
	;
	goto L11
L17:
	;
	goto L3
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v294 = v59
	goto L1
L19:
	;
	v87 = F_ACLMergeSelectorArguments(m, l0+int32(8), l1+int32(-2), v13+int32(12), l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L27
	}
L20:
	;
	return int32(0)
L21:
	;
	if v67 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	goto L25
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(7)
	v294 = int32(-1)
	goto L1
L26:
	;
	v90 = F_ACLCreateUnlinkedUser(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	if v87 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v294 = int32(-1)
	goto L1
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v92 < int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_valkey_free(m, v87)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L20
	} else {
		goto L78
	}
L31:
	;
	v183 = v92 << (uint(int32(2)) % 32)
	v186 = F_valkey_malloc(m, v183+int32(8))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L20
	} else {
		goto L57
	}
L32:
	;
	v100 = int32(0)
	goto L33
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87+v100<<(uint(int32(2))%32))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-1)))))
	switch v113 & int32(7) {
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
		v130 = int32(0)
		goto L35
	}
L34:
	;
	goto L31
L35:
	;
	v131 = F_ACLSetUser(m, v90, v110, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L42
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-17))))
	v130 = v129
	goto L35
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v110+int32(-9))))
	v130 = v126
	goto L35
L38:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+int32(-5)))))
	v130 = v123
	goto L35
L39:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+int32(-3)))))
	v130 = v120
	goto L35
L40:
	;
	v130 = int32(base.Ui32(v113) >> (uint(int32(3)) % 32))
	goto L35
L41:
	;
	v170 = v100 + int32(1)
	if v170 != v92 {
		v100 = v170
		goto L33
	} else {
		goto L56
	}
L42:
	;
	if v131 != int32(-1) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	goto L44
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v136 == int32(44) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	F_ACLFreeUser(m, v90)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	if l2 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v144 = int32(1)
	if v144 < v92 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v100
	goto L47
L49:
	;
	v147 = v92
	goto L51
L50:
	;
	v147 = v144
	goto L51
L51:
	;
	v151 = int32(0)
	goto L52
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v87+v151<<(uint(int32(2))%32))))
	F_sdsfree(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L54
	}
L53:
	;
	v282 = int32(-1)
	goto L30
L54:
	;
	v166 = v151 + int32(1)
	if v166 != v147 {
		v151 = v166
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L34
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v189 = F_sdsdup(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v189
	v192 = int32(0)
	if v92 <= v192 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v228 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v186+v183+int32(4)))) = v228
	v235 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v236 = F_listAddNodeTail(m, v235, v186)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L20
	} else {
		goto L65
	}
L60:
	;
	v197 = v192
	goto L61
L61:
	;
	v206 = v197 + int32(1)
	v207 = int32(2)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v87+v197<<(uint(v207)%32))))
	v214 = F_sdsdup(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L20
	} else {
		goto L63
	}
L62:
	;
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186+v206<<(uint(v207)%32)))) = v214
	if v206 != v92 {
		v197 = v206
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	F_sdsfree(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L20
	} else {
		goto L66
	}
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	if v241 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	F_listRelease(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L20
	} else {
		goto L70
	}
L68:
	;
	F_decrRefCount(m, v241)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = int32(0)
	goto L67
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	F_listRelease(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_valkey_free(m, v90)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	if v92 <= int32(0) {
		v282 = v228
		goto L30
	} else {
		goto L73
	}
L73:
	;
	v261 = int32(0)
	goto L74
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v87+v261<<(uint(int32(2))%32))))
	F_sdsfree(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L76
	}
L75:
	;
	v282 = v228
	goto L30
L76:
	;
	v276 = v261 + int32(1)
	if v276 != v92 {
		v261 = v276
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v294 = v282
	goto L1
}
func F_ACLCheckChannelAgainstList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = v10 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	goto L1
L1:
	;
	v19 = v10 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return v142
L3:
	;
	v142 = int32(5)
	goto L2
L4:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+base.B2i32(v24 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v30
	goto L5
L7:
	;
	v34 = v21
	goto L8
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-1)))))
	switch v45 & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	default:
		v62 = int32(0)
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	if l3 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-17))))
	v62 = v61
	goto L10
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42+int32(-9))))
	v62 = v58
	goto L10
L13:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+int32(-5)))))
	v62 = v55
	goto L10
L14:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+int32(-3)))))
	v62 = v52
	goto L10
L15:
	;
	v62 = int32(base.Ui32(v45) >> (uint(int32(3)) % 32))
	goto L10
L16:
	;
	v115 = v10 + int32(8)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v117 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v97 = int32(0)
	v100 = m.G0
	v101 = int32(16)
	v102 = v100 - v101
	m.G0 = v102
	*(*int32)(unsafe.Add(mBase, uint32(v102)+12)) = v97
	v109 = F_stringmatchlen_impl(m, v42, v62, l1, l2, v97, v102+int32(12), v97)
	mBase = m.M
	m.G0 = v102 + v101
	goto L28
L18:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v68 == int32(0) {
		v91 = v67
		v92 = v68
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v92-v91&int32(255) != 0 {
		goto L16
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	if v68 != v67&int32(255) {
		v91 = v67
		v92 = v68
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v74 = v42
	v75 = l1
	goto L23
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v79 == int32(0) {
		v91 = v78
		v92 = v79
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v91 = v78
	v92 = v79
	goto L20
L25:
	;
	v82 = int32(1)
	if v79 == v78&int32(255) {
		v74 = v74 + v82
		v75 = v75 + v82
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v142 = int32(0)
	goto L2
L28:
	;
	if v109 != 0 {
		v142 = v97
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	if v117 != 0 {
		v34 = v117
		goto L8
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117+base.B2i32(v120 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v126
	goto L31
L33:
	;
	goto L9
}
func F_ACLCleanupCategoriesOnFailure(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v7 = v6 - l0
	if base.Ui32(v6) <= base.Ui32(v7) {
		v36 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1])) = v36
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v12 = v10
	v13 = v7
	goto L3
L3:
	;
	v16 = v13 << (uint(int32(4)) % 32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+v16)))
	F_valkey_free(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v36 = v31 - l0
	goto L1
L5:
	;
	return
L6:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v23 = v22 + v16
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v21
	v29 = v13 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if base.Ui32(v29) < base.Ui32(v31) {
		v12 = v22
		v13 = v29
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_ACLCreateUnlinkedUser(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v10 = int32(0)
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
	v17 = F_snprintf(m, v6+int32(16), int32(64), int32(_a2), v6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v89 = v6 + int32(16)
	if v89&int32(3) == v86 {
		v113 = v89
		goto L26
	} else {
		goto L27
	}
L3:
	;
	return int32(0)
L4:
	;
	v24 = v6 + int32(16)
	if v24&int32(3) == int32(0) {
		v48 = v24
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v82 = F_ACLCreateUser(m, v24, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L21
	}
L6:
	;
	v81 = v73 - v24
	goto L5
L7:
	;
	v52 = v48
	goto L15
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = v24
	goto L11
L10:
	;
	v81 = v24 - v24
	goto L5
L11:
	;
	v41 = v37 + int32(1)
	if v41&int32(3) == int32(0) {
		v48 = v41
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v46 != 0 {
		v37 = v41
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v73 = v41
	goto L6
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 == v61 {
		v52 = v52 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v67 = v52
	goto L18
L17:
	;
	goto L16
L18:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 != 0 {
		v67 = v67 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v73 = v67
	goto L6
L20:
	;
	goto L19
L21:
	;
	if v82 == int32(0) {
		v10 = v10 + int32(1)
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L2
L23:
	;
	m.G0 = v6 + int32(80)
	return v82
L24:
	;
	v148 = F_raxRemove(m, v87, v89, v146, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L40
	}
L25:
	;
	v146 = v138 - v89
	goto L24
L26:
	;
	v117 = v113
	goto L34
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v99 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v102 = v89
	goto L30
L29:
	;
	v146 = v89 - v89
	goto L24
L30:
	;
	v106 = v102 + int32(1)
	if v106&int32(3) == int32(0) {
		v113 = v106
		goto L26
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v111 != 0 {
		v102 = v106
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v138 = v106
	goto L25
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v126 = int32(-2139062144)
	if (int32(16843008)-v123|v123)&v126 == v126 {
		v117 = v117 + int32(4)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v132 = v117
	goto L37
L36:
	;
	goto L35
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v136 != 0 {
		v132 = v132 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v138 = v132
	goto L25
L39:
	;
	goto L38
L40:
	;
	if v148 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	F__serverAssert(m, int32(_a3), int32(_a4), int32(466))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ACLDescribeUser(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v359 int32
	_ = v359
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int64
	_ = v398
	var v402 int64
	_ = v402
	var v406 int64
	_ = v406
	var v407 int64
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	v9 = m.G0
	v11 = v9 - int32(64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.G0 = v11 + int32(64)
	return v502
L2:
	;
	v21 = F_sdsempty(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L3:
	;
	F_incrRefCount(m, v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L1
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, _consts[4]))
	if v24 == int64(0) {
		v64 = v21
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = v11 + int32(40)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72
	goto L16
L8:
	;
	v30 = int32(0)
	v31 = v21
	v32 = v24
	v33 = int32(_a5)
	goto L9
L9:
	;
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	if v32&v36 == int64(0) {
		v47 = v31
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v64 = v47
	goto L7
L11:
	;
	v49 = v30 + int32(1)
	v51 = v49 << (uint(int32(4)) % 32)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v51)+uint32(_consts[4])))
	if base.B2i32(v56 == int64(0)) == int32(0) {
		v30 = v49
		v31 = v47
		v32 = v56
		v33 = v51 + int32(_a5)
		goto L9
	} else {
		goto L15
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v41 = F_sdscat(m, v31, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v45 = F_sdscatlen(m, v41, int32(_a6), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v47 = v45
	goto L11
L15:
	;
	goto L10
L16:
	;
	v77 = v11 + int32(40)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v79 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v135 = v11 + int32(40)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v136
	goto L31
L18:
	;
	if v79 == int32(0) {
		v128 = v64
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79+base.B2i32(v82 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v88
	goto L19
L21:
	;
	v94 = v79
	v95 = v64
	goto L22
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v103 = F_sdscatlen(m, v95, int32(_a7), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v128 = v109
	goto L17
L24:
	;
	v105 = F_sdscatsds(m, v103, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v109 = F_sdscatlen(m, v105, int32(_a6), int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v112 = v11 + int32(40)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v114 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v114 != 0 {
		v94 = v114
		v95 = v109
		goto L22
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114+base.B2i32(v117 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v123
	goto L28
L30:
	;
	goto L23
L31:
	;
	v141 = v11 + int32(40)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v143 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v489 = F_createObject(m, int32(0), v483)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L124
	}
L33:
	;
	if v143 == int32(0) {
		v483 = v128
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v143+base.B2i32(v146 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v152
	goto L34
L36:
	;
	v158 = v143
	v159 = v128
	goto L37
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v165 = F_sdsempty(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	v483 = v462
	goto L32
L39:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v167&int32(2) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v236&int32(8) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v164)+140))
	v178 = v11 + int32(56)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v179
	goto L44
L42:
	;
	v174 = F_sdscatlen(m, v165, int32(_a8), int32(3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v230 = v174
	goto L40
L44:
	;
	v184 = v11 + int32(56)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v186 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v186 == int32(0) {
		v230 = v165
		goto L40
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186+base.B2i32(v189 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v195
	goto L46
L48:
	;
	v201 = v165
	v205 = v186
	goto L49
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	v208 = F_sdsCatPatternString(m, v201, v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	v230 = v212
	goto L40
L51:
	;
	v212 = F_sdscatlen(m, v208, int32(_a6), int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v215 = v11 + int32(56)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v217 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v217 != 0 {
		v201 = v212
		v205 = v217
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217+base.B2i32(v220 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v226
	goto L54
L56:
	;
	goto L50
L57:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v313&int32(16) != 0 {
		v443 = v307
		goto L76
	} else {
		goto L77
	}
L58:
	;
	v247 = F_sdscatlen(m, v230, int32(_a9), int32(14))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v243 = F_sdscatlen(m, v230, int32(_a10), int32(3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v307 = v243
	goto L57
L61:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v164)+144))
	v251 = v11 + int32(56)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v252
	goto L62
L62:
	;
	v257 = v11 + int32(56)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v259 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v259 == int32(0) {
		v307 = v247
		goto L57
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v259+base.B2i32(v262 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v268
	goto L64
L66:
	;
	v274 = v247
	v278 = v259
	goto L67
L67:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	v283 = F_sdscatlen(m, v274, int32(_a11), int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v307 = v289
	goto L57
L69:
	;
	v285 = F_sdscatsds(m, v283, v280)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v289 = F_sdscatlen(m, v285, int32(_a6), int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v292 = v11 + int32(56)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v294 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v294 != 0 {
		v274 = v289
		v278 = v294
		goto L67
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v294+base.B2i32(v297 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v303
	goto L73
L75:
	;
	goto L68
L76:
	;
	v449 = F_ACLDescribeSelectorCommandRules(m, v164)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L112
	}
L77:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v164)+152))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	goto L79
L78:
	;
	v324 = F_sdscatlen(m, v307, int32(_a12), int32(3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L82
	}
L79:
	;
	if v317 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v320 = F_sdscatlen(m, v307, int32(_a13), int32(9))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v443 = v320
	goto L76
L82:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v164)+152))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	goto L84
L83:
	;
	v439 = F_sdscatlen(m, v435, int32(_a6), int32(1))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L111
	}
L84:
	;
	if v327 == int32(0) {
		v435 = v324
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v164)+152))
	v331 = int32(0)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if base.Ui32(v337) <= base.Ui32(v331) {
		v359 = v331
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v371 = int32(1)
	if v327 == v371 {
		v435 = v370
		goto L83
	} else {
		goto L96
	}
L87:
	;
	if v359 == int32(0) {
		v370 = v324
		goto L86
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	switch v341 + int32(-4) {
	case 0:
		goto L92
	default:
		goto L91
	case 4:
		goto L93
	}
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v356
	v359 = int32(1)
	goto L88
L91:
	;
	v355 = int64(*(*int16)(unsafe.Add(mBase, uint32(v330+int32(8)))))
	v356 = v355
	goto L90
L92:
	;
	v351 = int64(*(*int32)(unsafe.Add(mBase, uint32(v330+int32(8)))))
	v356 = v351
	goto L90
L93:
	;
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v330+int32(8))))
	v356 = v347
	goto L90
L94:
	;
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v363
	v368 = F_sdscatfmt(m, v324, int32(_a14), v11+int32(32))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v370 = v368
	goto L86
L96:
	;
	v376 = v371
	v380 = v370
	goto L97
L97:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v164)+152))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if base.Ui32(v388) <= base.Ui32(v376) {
		v410 = int32(0)
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v435 = v425
	goto L83
L99:
	;
	v427 = v376 + int32(1)
	if v427 != v327 {
		v376 = v427
		v380 = v425
		goto L97
	} else {
		goto L110
	}
L100:
	;
	if v410 == int32(0) {
		v425 = v380
		goto L99
	} else {
		goto L107
	}
L101:
	;
	goto L100
L102:
	;
	v391 = v382 + int32(8)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	switch v392 + int32(-4) {
	case 0:
		goto L105
	default:
		goto L104
	case 4:
		goto L106
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v407
	v410 = int32(1)
	goto L101
L104:
	;
	v406 = int64(*(*int16)(unsafe.Add(mBase, uint32(v391+v376<<(uint(int32(1))%32)))))
	v407 = v406
	goto L103
L105:
	;
	v402 = int64(*(*int32)(unsafe.Add(mBase, uint32(v391+v376<<(uint(int32(2))%32)))))
	v407 = v402
	goto L103
L106:
	;
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v391+v376<<(uint(int32(3))%32))))
	v407 = v398
	goto L103
L107:
	;
	v416 = F_sdscatlen(m, v380, int32(_a15), int32(1))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v418
	v423 = F_sdscatfmt(m, v416, int32(_a14), v11+int32(16))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v425 = v423
	goto L99
L110:
	;
	goto L98
L111:
	;
	v443 = v439
	goto L76
L112:
	;
	v451 = F_sdscatsds(m, v443, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	F_sdsfree(m, v449)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v451
	if v455&int32(1) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v461 = int32(_a16)
	goto L117
L116:
	;
	v461 = int32(_a17)
	goto L117
L117:
	;
	v462 = F_sdscatfmt(m, v159, v461, v11)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_sdsfree(m, v451)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v467 = v11 + int32(40)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	if v469 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v469 != 0 {
		v158 = v469
		v159 = v462
		goto L37
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v469+base.B2i32(v472 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v467))) = v478
	goto L121
L123:
	;
	goto L38
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v489
	F_incrRefCount(m, v489)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	goto L1
}
func F_ACLListDuplicateSelector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v9 = F_valkey_malloc(m, int32(160))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v16 = F_listDup(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v20 = F_listDup(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v24 = F_intsetDup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v28 = F_sdsdup(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v28
	v31 = int32(8)
	goto L9
L7:
	;
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v43 == v40 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L7
L9:
	;
	v38 = F__emscripten_memcpy_bulkmem(m, v9+v31, l0+v31, int32(128))
	mBase = m.M
	goto L8
L10:
	;
	return v9
L11:
	;
	v48 = v40
	v49 = v43
	goto L12
L12:
	;
	v54 = v48 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49+v54)))
	if v56 == int32(0) {
		v86 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L10
L14:
	;
	v91 = v48 + int32(1)
	if v91 != int32(1024) {
		v48 = v91
		v49 = v86
		goto L12
	} else {
		goto L21
	}
L15:
	;
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v49+v54)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 == v59 {
		v86 = v49
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v70 = v59
	v71 = v62
	goto L17
L17:
	;
	F_ACLAddAllowedFirstArg(m, v9, v48, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v86 = v74
	goto L14
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74+v54)))
	v78 = v70 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+v78<<(uint(int32(2))%32))))
	if v82 != 0 {
		v70 = v78
		v71 = v82
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L13
}
func F_ACLLoadFromFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int64
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1290 int32
	_ = v1290
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
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
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	v15 = m.G0
	v17 = v15 - int32(1184)
	m.G0 = v17
	v20 = F_fopen(m, l0, int32(_a24))
	mBase = m.M
	v21 = F_sdsempty(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(1184)
	return v1682
L4:
	;
	v36 = F_fgets(m, v17+int32(144), int32(1024), v20)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v27 = F___strerror_l(m, v26, v26)
	mBase = m.M
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	v31 = F_sdscatprintf(m, v21, int32(_a25), v17)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v1682 = v31
	goto L3
L9:
	;
	v77 = F_fclose(m, v20)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L10:
	;
	if v36 == int32(0) {
		v66 = v21
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v43 = v21
	goto L12
L12:
	;
	v56 = F_sdscat(m, v43, v17+int32(144))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v66 = v56
	goto L9
L14:
	;
	v61 = F_fgets(m, v17+int32(144), int32(1024), v20)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v61 != 0 {
		v43 = v56
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v79 = F_sdsempty(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(0)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-1)))))
	switch v85 & int32(7) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		v102 = v81
		goto L19
	}
L19:
	;
	v107 = F_sdssplitlen(m, v66, v102, int32(_a26), int32(1), v17+int32(140))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L25
	}
L20:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(-17))))
	v102 = v101
	goto L19
L21:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(-9))))
	v102 = v98
	goto L19
L22:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66+int32(-5)))))
	v102 = v95
	goto L19
L23:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+int32(-3)))))
	v102 = v92
	goto L19
L24:
	;
	v102 = int32(base.Ui32(v85) >> (uint(int32(3)) % 32))
	goto L19
L25:
	;
	F_sdsfree(m, v66)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v114 = F_raxNew(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v114
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)+140))
	if v117 <= int32(0) {
		v724 = v117
		v725 = v79
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_sdsfreesplitres(m, v107, v724)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L186
	}
L29:
	;
	v122 = v81
	v124 = v79
	goto L30
L30:
	;
	v136 = v107 + v122<<(uint(int32(2))%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v138 = int32(_a27)
	v144 = v137 + int32(-1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	switch v145 & int32(7) {
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
		v162 = int32(0)
		goto L33
	}
L31:
	;
	v724 = v719
	v725 = v709
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v137
	v231 = v122 + int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v232 == int32(0) {
		v709 = v124
		goto L59
	} else {
		goto L60
	}
L33:
	;
	v165 = v137 + v162 + int32(-1)
	if base.Ui32(v165) < base.Ui32(v137) {
		v183 = v137
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-17))))
	v162 = v161
	goto L33
L35:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-9))))
	v162 = v158
	goto L33
L36:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+int32(-5)))))
	v162 = v155
	goto L33
L37:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-3)))))
	v162 = v152
	goto L33
L38:
	;
	v162 = int32(base.Ui32(v145) >> (uint(int32(3)) % 32))
	goto L33
L39:
	;
	if base.Ui32(v165) <= base.Ui32(v183) {
		v199 = v165
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v171 = v137
	goto L41
L41:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v171))))
	v173 = F_strchr(m, v138, v172)
	mBase = m.M
	if v173 == int32(0) {
		v183 = v171
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v183 = v177
	goto L39
L43:
	;
	v177 = v171 + int32(1)
	if base.Ui32(v177) <= base.Ui32(v165) {
		v171 = v177
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v204 = v199 - v183 + int32(1)
	if v137 == v183 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v187 = v165
	goto L47
L47:
	;
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v187))))
	v191 = F_strchr(m, v138, v190)
	mBase = m.M
	if v191 == int32(0) {
		v199 = v187
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v199 = v183
	goto L45
L49:
	;
	v195 = v187 + int32(-1)
	if base.Ui32(v183) < base.Ui32(v195) {
		v187 = v195
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v204))) = uint8(v208)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	switch v210 & int32(7) {
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
		goto L53
	}
L52:
	;
	v206 = F_memmove(m, v137, v183, v204)
	mBase = m.M
	goto L51
L53:
	;
	goto L32
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v137+int32(-17)))) = base.I64_extend_i32_u(v204)
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137+int32(-9)))) = v204
	goto L32
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v137+int32(-5)))) = uint16(v204)
	goto L32
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-3)))) = uint8(v204)
	goto L32
L58:
	;
	v214 = v204 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v214)
	goto L32
L59:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v17)+140))
	if v231 < v719 {
		v122 = v231
		v124 = v709
		goto L30
	} else {
		goto L185
	}
L60:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-1)))))
	switch v238 & int32(7) {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	case 4:
		goto L62
	default:
		v255 = int32(0)
		goto L61
	}
L61:
	;
	v260 = F_sdssplitlen(m, v137, v255, int32(_a6), int32(1), v17+int32(132))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L68
	}
L62:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-17))))
	v255 = v254
	goto L61
L63:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v137+int32(-9))))
	v255 = v251
	goto L61
L64:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137+int32(-5)))))
	v255 = v248
	goto L61
L65:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(-3)))))
	v255 = v245
	goto L61
L66:
	;
	v255 = int32(base.Ui32(v238) >> (uint(int32(3)) % 32))
	goto L61
L67:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	if v271 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	if v260 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v231
	v264 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v264
	v269 = F_sdscatprintf(m, v124, int32(_a28), v17+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v709 = v269
	goto L59
L71:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v276 = int32(_a23)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, _consts[20])))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v280 == int32(0) {
		v303 = v279
		v304 = v280
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_sdsfreesplitres(m, v260, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v709 = v124
	goto L59
L74:
	;
	if v271 < int32(2) {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	goto L74
L76:
	;
	if v280 != v279&int32(255) {
		v303 = v279
		v304 = v280
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v286 = v275
	v287 = v276
	goto L78
L78:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	if v291 == int32(0) {
		v303 = v290
		v304 = v291
		goto L75
	} else {
		goto L80
	}
L79:
	;
	v303 = v290
	v304 = v291
	goto L75
L80:
	;
	v294 = int32(1)
	if v291 == v290&int32(255) {
		v286 = v286 + v294
		v287 = v287 + v294
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+int32(-1)))))
	v330 = v328 & int32(7)
	switch v330 {
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
		v422 = int32(0)
		goto L88
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v231
	v314 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v314
	v319 = F_sdscatprintf(m, v124, int32(_a29), v17+int32(32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	if v304-v303&int32(255) == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	F_sdsfreesplitres(m, v260, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v709 = v319
	goto L59
L88:
	;
	v433 = F_ACLCreateUser(m, v325, v422)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L111
	}
L89:
	;
	v346 = int32(0)
	if v345 == v346 {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v325+int32(-17))))
	v345 = v344
	goto L89
L91:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v325+int32(-9))))
	v345 = v341
	goto L89
L92:
	;
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325+int32(-5)))))
	v345 = v338
	goto L89
L93:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+int32(-3)))))
	v345 = v335
	goto L89
L94:
	;
	v345 = int32(base.Ui32(v328) >> (uint(int32(3)) % 32))
	goto L89
L95:
	;
	switch v330 {
	case 0:
		goto L109
	case 1:
		goto L108
	case 2:
		goto L107
	case 3:
		goto L106
	case 4:
		goto L105
	default:
		v422 = int32(0)
		goto L88
	}
L96:
	;
	v352 = v346
	goto L98
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = v325
	v380 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v380
	v385 = F_sdscatprintf(m, v124, int32(_a30), v17+int32(112))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L103
	}
L98:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+v352))))
	if v364&int32(223) == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	if base.Ui32(base.I32_extend8_s(v364)+int32(-9)) <= base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v375 = v352 + int32(1)
	if v375 == v345 {
		goto L95
	} else {
		goto L102
	}
L102:
	;
	v352 = v375
	goto L98
L103:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	F_sdsfreesplitres(m, v260, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v709 = v385
	goto L59
L105:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v325+int32(-17))))
	v422 = v418
	goto L88
L106:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v325+int32(-9))))
	v422 = v415
	goto L88
L107:
	;
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v325+int32(-5)))))
	v422 = v412
	goto L88
L108:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+int32(-3)))))
	v422 = v409
	goto L88
L109:
	;
	v422 = int32(base.Ui32(v328) >> (uint(int32(3)) % 32))
	goto L88
L110:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	v454 = F_ACLMergeSelectorArguments(m, v260+int32(8), v448+int32(-2), v17+int32(1180), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L116
	}
L111:
	;
	if v433 != 0 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v435
	v441 = F_sdscatprintf(m, v124, int32(_a31), v17+int32(48))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	F_sdsfreesplitres(m, v260, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v709 = v441
	goto L59
L115:
	;
	v466 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1180))
	if v468 <= v466 {
		v690 = v465
		goto L119
	} else {
		goto L120
	}
L116:
	;
	if v454 != 0 {
		v465 = v124
		goto L115
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v231
	v458 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v458
	v463 = F_sdscatprintf(m, v124, int32(_a32), v17+int32(96))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v465 = v463
	goto L115
L119:
	;
	F_valkey_free(m, v454)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L183
	}
L120:
	;
	v475 = v465
	v476 = v466
	v484 = v466
	goto L121
L121:
	;
	v487 = v454 + v476<<(uint(int32(2))%32)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	v489 = int32(_a33)
	v495 = v488 + int32(-1)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	switch v496 & int32(7) {
	case 0:
		goto L129
	case 1:
		goto L128
	case 2:
		goto L127
	case 3:
		goto L126
	case 4:
		goto L125
	default:
		v513 = int32(0)
		goto L124
	}
L122:
	;
	v660 = int32(0)
	if v658 <= v660 {
		v690 = v654
		goto L119
	} else {
		goto L178
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v488
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+int32(-1)))))
	switch v584 & int32(7) {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	case 3:
		goto L152
	case 4:
		goto L151
	default:
		v601 = int32(0)
		goto L150
	}
L124:
	;
	v516 = v488 + v513 + int32(-1)
	if base.Ui32(v516) < base.Ui32(v488) {
		v534 = v488
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(-17))))
	v513 = v512
	goto L124
L126:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(-9))))
	v513 = v509
	goto L124
L127:
	;
	v506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488+int32(-5)))))
	v513 = v506
	goto L124
L128:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+int32(-3)))))
	v513 = v503
	goto L124
L129:
	;
	v513 = int32(base.Ui32(v496) >> (uint(int32(3)) % 32))
	goto L124
L130:
	;
	if base.Ui32(v516) <= base.Ui32(v534) {
		v550 = v516
		goto L136
	} else {
		goto L137
	}
L131:
	;
	v522 = v488
	goto L132
L132:
	;
	v523 = int32(*(*int8)(unsafe.Add(mBase, uint32(v522))))
	v524 = F_strchr(m, v489, v523)
	mBase = m.M
	if v524 == int32(0) {
		v534 = v522
		goto L130
	} else {
		goto L134
	}
L133:
	;
	v534 = v528
	goto L130
L134:
	;
	v528 = v522 + int32(1)
	if base.Ui32(v528) <= base.Ui32(v516) {
		v522 = v528
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v555 = v550 - v534 + int32(1)
	if v488 == v534 {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v538 = v516
	goto L138
L138:
	;
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v538))))
	v542 = F_strchr(m, v489, v541)
	mBase = m.M
	if v542 == int32(0) {
		v550 = v538
		goto L136
	} else {
		goto L140
	}
L139:
	;
	v550 = v534
	goto L136
L140:
	;
	v546 = v538 + int32(-1)
	if base.Ui32(v534) < base.Ui32(v546) {
		v538 = v546
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v559 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v488+v555))) = uint8(v559)
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	switch v561 & int32(7) {
	case 0:
		goto L149
	case 1:
		goto L148
	case 2:
		goto L147
	case 3:
		goto L146
	case 4:
		goto L145
	default:
		goto L144
	}
L143:
	;
	v557 = F_memmove(m, v488, v534, v555)
	mBase = m.M
	goto L142
L144:
	;
	goto L123
L145:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v488+int32(-17)))) = base.I64_extend_i32_u(v555)
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488+int32(-9)))) = v555
	goto L123
L147:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v488+int32(-5)))) = uint16(v555)
	goto L123
L148:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v488+int32(-3)))) = uint8(v555)
	goto L123
L149:
	;
	v565 = v555 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v495))) = uint8(v565)
	goto L123
L150:
	;
	v602 = F_ACLSetUser(m, v433, v488, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L157
	}
L151:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(-17))))
	v601 = v600
	goto L150
L152:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v488+int32(-9))))
	v601 = v597
	goto L150
L153:
	;
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488+int32(-5)))))
	v601 = v594
	goto L150
L154:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+int32(-3)))))
	v601 = v591
	goto L150
L155:
	;
	v601 = int32(base.Ui32(v584) >> (uint(int32(3)) % 32))
	goto L150
L156:
	;
	v657 = v476 + int32(1)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1180))
	if v657 < v658 {
		v475 = v654
		v476 = v657
		v484 = v655
		goto L121
	} else {
		goto L177
	}
L157:
	;
	if v602 == int32(0) {
		v654 = v475
		v655 = v484
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v608 = F___errno_location(m)
	mBase = m.M
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	switch v609 + int32(-7) {
	case 0:
		goto L164
	default:
		goto L161
	case 2:
		goto L165
	case 5:
		goto L163
	case 13:
		goto L168
	case 21:
		goto L169
	case 24:
		goto L167
	case 36:
		goto L166
	case 37:
		v621 = int32(_a34)
		goto L160
	case 61:
		goto L162
	}
L159:
	;
	goto L171
L160:
	;
	v623 = v621
	goto L159
L161:
	;
	v621 = int32(_a35)
	goto L160
L162:
	;
	v623 = int32(_a36)
	goto L159
L163:
	;
	v623 = int32(_a37)
	goto L159
L164:
	;
	v623 = int32(_a38)
	goto L159
L165:
	;
	v623 = int32(_a39)
	goto L159
L166:
	;
	v623 = int32(_a40)
	goto L159
L167:
	;
	v623 = int32(_a41)
	goto L159
L168:
	;
	v623 = int32(_a42)
	goto L159
L169:
	;
	v623 = int32(_a43)
	goto L159
L170:
	;
	if v484 != 0 {
		v650 = v475
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v625 != int32(44) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	v630 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v623
	v638 = F_sdscatprintf(m, v475, int32(_a44), v17+int32(64))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v654 = v638
	v655 = v484
	goto L156
L174:
	;
	v654 = v650
	v655 = int32(1)
	goto L156
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v623
	v643 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v643
	v648 = F_sdscatprintf(m, v475, int32(_a45), v17+int32(80))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v650 = v648
	goto L174
L177:
	;
	goto L122
L178:
	;
	v666 = v660
	goto L179
L179:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v454+v666<<(uint(int32(2))%32))))
	F_sdsfree(m, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	v690 = v654
	goto L119
L181:
	;
	v684 = v666 + int32(1)
	if v684 != v658 {
		v666 = v684
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	F_sdsfreesplitres(m, v260, v702)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v709 = v690
	goto L59
L185:
	;
	goto L31
L186:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+int32(-1)))))
	switch v739 & int32(7) {
	case 0:
		goto L192
	case 1:
		goto L193
	case 2:
		goto L189
	case 3:
		goto L190
	case 4:
		goto L191
	default:
		goto L188
	}
L187:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	F_raxFreeWithCallback(m, v1671, int32(13))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L1
	} else {
		goto L395
	}
L188:
	;
	v762 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = v762
	v765 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v767 = int32(7)
	v769 = v17 + int32(132)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	goto L202
L189:
	;
	v761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v725+int32(-5)))))
	if v761 != 0 {
		goto L187
	} else {
		goto L198
	}
L190:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v725+int32(-9))))
	if v756 == int32(0) {
		goto L188
	} else {
		goto L197
	}
L191:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v725+int32(-17))))
	if v751 == int32(0) {
		goto L188
	} else {
		goto L196
	}
L192:
	;
	if int32(base.Ui32(v739)>>(uint(int32(3))%32)) == int32(0) {
		goto L188
	} else {
		goto L195
	}
L193:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+int32(-3)))))
	if v744 != 0 {
		goto L187
	} else {
		goto L194
	}
L194:
	;
	goto L188
L195:
	;
	goto L187
L196:
	;
	goto L187
L197:
	;
	goto L187
L198:
	;
	goto L188
L199:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v17)+132))
	if v965 != 0 {
		v994 = v965
		goto L237
	} else {
		goto L238
	}
L200:
	;
	if v922 != v767 {
		goto L227
	} else {
		goto L228
	}
L201:
	;
	v913 = int32(0)
	v919 = v778
	v920 = v779
	v922 = v913
	v926 = v913
	goto L200
L202:
	;
	if base.Ui32(v779) < base.Ui32(int32(8)) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v790 = v778
	v791 = v779
	v793 = int32(0)
	goto L205
L204:
	;
	v919 = v903
	v920 = v904
	v922 = v906
	v926 = base.B2i32(v909 != int32(0))
	goto L200
L205:
	;
	v799 = int32(base.Ui32(v791) >> (uint(int32(3)) % 32))
	v800 = int32(4)
	v801 = v790 + v800
	if v791&v800 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v903 = v894
	v904 = v895
	v906 = v879
	v909 = v884
	goto L204
L207:
	;
	v884 = int32(0)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v801+v799+(v884-v799)&int32(3)+v872<<(uint(int32(2))%32))))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)))
	if base.Ui32(v895) < base.Ui32(int32(8)) {
		v903 = v894
		v904 = v895
		v906 = v879
		v909 = v884
		goto L204
	} else {
		goto L225
	}
L208:
	;
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793)+uint32(_consts[21]))))
	v850 = int32(0)
	goto L219
L209:
	;
	v806 = int32(0)
	if base.Ui32(v767) <= base.Ui32(v793) {
		v839 = v793
		v842 = v806
		goto L210
	} else {
		goto L211
	}
L210:
	;
	if v842 == v799 {
		v872 = v806
		v879 = v839
		goto L207
	} else {
		goto L217
	}
L211:
	;
	v816 = v793
	v819 = v806
	goto L212
L212:
	;
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801+v819))))
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816)+uint32(_consts[21]))))
	if v822 != v824 {
		v839 = v816
		v842 = v819
		goto L210
	} else {
		goto L214
	}
L213:
	;
	v839 = v827
	v842 = v829
	goto L210
L214:
	;
	v826 = int32(1)
	v827 = v816 + v826
	v829 = v819 + v826
	if base.Ui32(v799) <= base.Ui32(v829) {
		v839 = v827
		v842 = v829
		goto L210
	} else {
		goto L215
	}
L215:
	;
	if base.Ui32(v827) < base.Ui32(v767) {
		v816 = v827
		v819 = v829
		goto L212
	} else {
		goto L216
	}
L216:
	;
	goto L213
L217:
	;
	v903 = v790
	v904 = v791
	v906 = v839
	v909 = v842
	goto L204
L218:
	;
	if v850 != v799 {
		goto L223
	} else {
		goto L224
	}
L219:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801+v850))))
	if v863 == v847&int32(255) {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	v865 = int32(1)
	v867 = v850 + v865
	if v867 != v799 {
		v850 = v867
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v919 = v790
	v920 = v791
	v922 = v793
	v926 = v865
	goto L200
L223:
	;
	v872 = v850
	v879 = v793 + int32(1)
	goto L207
L224:
	;
	v903 = v790
	v904 = v791
	v906 = v793
	v909 = v799
	goto L204
L225:
	;
	if base.Ui32(v879) < base.Ui32(v767) {
		v790 = v894
		v791 = v895
		v793 = v879
		goto L205
	} else {
		goto L226
	}
L226:
	;
	goto L206
L227:
	;
	goto L199
L228:
	;
	if v920&int32(1) == int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v934 = v920 & int32(4)
	if v926&base.B2i32(v934 != int32(0)) != 0 {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	if v769 == int32(0) {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	if v920&int32(2) != 0 {
		v960 = int32(0)
		goto L232
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v769))) = v960
	goto L227
L233:
	;
	v944 = int32(3)
	v945 = int32(base.Ui32(v920) >> (uint(v944) % 32))
	if v934 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v955 = int32(4)
	goto L236
L235:
	;
	v955 = v945 << (uint(int32(2)) % 32)
	goto L236
L236:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v919+v945+(int32(0)-v945)&v944+v955+int32(4))))
	v960 = v959
	goto L232
L237:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	F_ACLCopyUser(m, v996, v994)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L246
	}
L238:
	;
	v968 = F_ACLCreateUser(m, int32(_a46), int32(7))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v972 = F_ACLSetUser(m, v968, int32(_a47), int32(-1))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v976 = F_ACLSetUser(m, v968, int32(_a48), int32(-1))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v980 = F_ACLSetUser(m, v968, int32(_a49), int32(-1))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v984 = F_ACLSetUser(m, v968, int32(_a50), int32(-1))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v988 = F_ACLSetUser(m, v968, int32(_a51), int32(-1))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v992 = F_ACLSetUser(m, v968, int32(_a52), int32(-1))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v994 = v968
	goto L237
L246:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v994)))
	F_sdsfree(m, v999)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v994)+16))
	if v1002 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v994)+8))
	F_listRelease(m, v1009)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	F_decrRefCount(m, v1002)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v994)+16)) = int32(0)
	goto L248
L251:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v994)+12))
	F_listRelease(m, v1012)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_valkey_free(m, v994)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1017 = int32(0)
	v1019 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1023 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v1025 = F_raxInsert(m, v1019, int32(_a46), int32(7), v1023, v1017)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v1030 = F_raxRemove(m, v112, int32(_a46), int32(7), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1034 = int32(_a20)
	v1035 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+12))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+16))
	v1039 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v1040 = F_kvstoreSize(m, v1039)
	mBase = m.M
	v1045 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v1046 = F_kvstoreSize(m, v1045)
	mBase = m.M
	goto L257
L256:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v1057 = v17 + int32(132)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1055)))
	*(*int32)(unsafe.Add(mBase, uint32(v1057)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1057))) = v1058
	goto L260
L257:
	;
	if base.I32_wrap_i64(v1040+base.I64_extend_i32_u(v1037+v1036)+v1046) < int32(1) {
		v1053 = v1017
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1051 = F_raxNew(m)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v1053 = v1051
	goto L256
L260:
	;
	v1063 = v17 + int32(132)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	if v1065 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	if v1053 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L262:
	;
	if v1065 == int32(0) {
		goto L261
	} else {
		goto L265
	}
L263:
	;
	goto L262
L264:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1065+base.B2i32(v1068 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1074
	goto L263
L265:
	;
	v1080 = v1065
	goto L266
L266:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+8))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)+328))
	if v1093 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L261
L268:
	;
	v1632 = v17 + int32(132)
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1632)))
	if v1634 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L269:
	;
	v1096 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v1096
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+int32(-1)))))
	switch v1102 & int32(7) {
	case 0:
		goto L275
	case 1:
		goto L274
	case 2:
		goto L273
	case 3:
		goto L272
	case 4:
		goto L271
	default:
		v1119 = v1096
		goto L270
	}
L270:
	;
	v1120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+1180)) = v1120
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1125 = v17 + int32(1180)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)))
	if v1119 == v1120 {
		goto L278
	} else {
		goto L279
	}
L271:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1099+int32(-17))))
	v1119 = v1118
	goto L270
L272:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1099+int32(-9))))
	v1119 = v1115
	goto L270
L273:
	;
	v1112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099+int32(-5)))))
	v1119 = v1112
	goto L270
L274:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+int32(-3)))))
	v1119 = v1109
	goto L270
L275:
	;
	v1119 = int32(base.Ui32(v1102) >> (uint(int32(3)) % 32))
	goto L270
L276:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1180))
	if v1321 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L277:
	;
	if v1278 != v1119 {
		goto L304
	} else {
		goto L305
	}
L278:
	;
	v1269 = int32(0)
	v1275 = v1134
	v1276 = v1135
	v1278 = v1269
	v1282 = v1269
	goto L277
L279:
	;
	if base.Ui32(v1135) < base.Ui32(int32(8)) {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1146 = v1134
	v1147 = v1135
	v1149 = int32(0)
	goto L282
L281:
	;
	v1275 = v1259
	v1276 = v1260
	v1278 = v1262
	v1282 = base.B2i32(v1265 != int32(0))
	goto L277
L282:
	;
	v1155 = int32(base.Ui32(v1147) >> (uint(int32(3)) % 32))
	v1156 = int32(4)
	v1157 = v1146 + v1156
	if v1147&v1156 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	v1259 = v1250
	v1260 = v1251
	v1262 = v1235
	v1265 = v1240
	goto L281
L284:
	;
	v1240 = int32(0)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1157+v1155+(v1240-v1155)&int32(3)+v1228<<(uint(int32(2))%32))))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1250)))
	if base.Ui32(v1251) < base.Ui32(int32(8)) {
		v1259 = v1250
		v1260 = v1251
		v1262 = v1235
		v1265 = v1240
		goto L281
	} else {
		goto L302
	}
L285:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+v1149))))
	v1206 = int32(0)
	goto L296
L286:
	;
	v1162 = int32(0)
	if base.Ui32(v1119) <= base.Ui32(v1149) {
		v1195 = v1149
		v1198 = v1162
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v1198 == v1155 {
		v1228 = v1162
		v1235 = v1195
		goto L284
	} else {
		goto L294
	}
L288:
	;
	v1172 = v1149
	v1175 = v1162
	goto L289
L289:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157+v1175))))
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099+v1172))))
	if v1178 != v1180 {
		v1195 = v1172
		v1198 = v1175
		goto L287
	} else {
		goto L291
	}
L290:
	;
	v1195 = v1183
	v1198 = v1185
	goto L287
L291:
	;
	v1182 = int32(1)
	v1183 = v1172 + v1182
	v1185 = v1175 + v1182
	if base.Ui32(v1155) <= base.Ui32(v1185) {
		v1195 = v1183
		v1198 = v1185
		goto L287
	} else {
		goto L292
	}
L292:
	;
	if base.Ui32(v1183) < base.Ui32(v1119) {
		v1172 = v1183
		v1175 = v1185
		goto L289
	} else {
		goto L293
	}
L293:
	;
	goto L290
L294:
	;
	v1259 = v1146
	v1260 = v1147
	v1262 = v1195
	v1265 = v1198
	goto L281
L295:
	;
	if v1206 != v1155 {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157+v1206))))
	if v1219 == v1203&int32(255) {
		goto L295
	} else {
		goto L298
	}
L298:
	;
	v1221 = int32(1)
	v1223 = v1206 + v1221
	if v1223 != v1155 {
		v1206 = v1223
		goto L296
	} else {
		goto L299
	}
L299:
	;
	v1275 = v1146
	v1276 = v1147
	v1278 = v1149
	v1282 = v1221
	goto L277
L300:
	;
	v1228 = v1206
	v1235 = v1149 + int32(1)
	goto L284
L301:
	;
	v1259 = v1146
	v1260 = v1147
	v1262 = v1149
	v1265 = v1155
	goto L281
L302:
	;
	if base.Ui32(v1235) < base.Ui32(v1119) {
		v1146 = v1250
		v1147 = v1251
		v1149 = v1235
		goto L282
	} else {
		goto L303
	}
L303:
	;
	goto L283
L304:
	;
	goto L276
L305:
	;
	if v1276&int32(1) == int32(0) {
		goto L304
	} else {
		goto L306
	}
L306:
	;
	v1290 = v1276 & int32(4)
	if v1282&base.B2i32(v1290 != int32(0)) != 0 {
		goto L304
	} else {
		goto L307
	}
L307:
	;
	if v1125 == int32(0) {
		goto L304
	} else {
		goto L308
	}
L308:
	;
	if v1276&int32(2) != 0 {
		v1316 = int32(0)
		goto L309
	} else {
		goto L310
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1125))) = v1316
	goto L304
L310:
	;
	v1300 = int32(3)
	v1301 = int32(base.Ui32(v1276) >> (uint(v1300) % 32))
	if v1290 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1311 = int32(4)
	goto L313
L312:
	;
	v1311 = v1301 << (uint(int32(2)) % 32)
	goto L313
L313:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1275+v1301+(int32(0)-v1301)&v1300+v1311+int32(4))))
	v1316 = v1315
	goto L309
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+328)) = v1321
	goto L268
L315:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+328)) = v1603
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1092)+204))
	goto L383
L316:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v17)+128))
	if v1591 == int32(0) {
		goto L314
	} else {
		goto L378
	}
L317:
	;
	if v1321 == int32(0) {
		goto L315
	} else {
		goto L377
	}
L318:
	;
	if v1053 == int32(0) {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327+int32(-1)))))
	switch v1330 & int32(7) {
	case 0:
		goto L325
	case 1:
		goto L324
	case 2:
		goto L323
	case 3:
		goto L322
	case 4:
		goto L321
	default:
		v1347 = int32(0)
		goto L320
	}
L320:
	;
	v1349 = v17 + int32(128)
	v1350 = int32(0)
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1053)))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1358)))
	if v1347 == v1350 {
		goto L328
	} else {
		goto L329
	}
L321:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1327+int32(-17))))
	v1347 = v1346
	goto L320
L322:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1327+int32(-9))))
	v1347 = v1343
	goto L320
L323:
	;
	v1340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1327+int32(-5)))))
	v1347 = v1340
	goto L320
L324:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327+int32(-3)))))
	v1347 = v1337
	goto L320
L325:
	;
	v1347 = int32(base.Ui32(v1330) >> (uint(int32(3)) % 32))
	goto L320
L326:
	;
	if v1543 != 0 {
		goto L316
	} else {
		goto L364
	}
L327:
	;
	if v1502 != v1347 {
		v1543 = v1350
		goto L354
	} else {
		goto L355
	}
L328:
	;
	v1493 = int32(0)
	v1499 = v1358
	v1500 = v1359
	v1502 = v1493
	v1506 = v1493
	goto L327
L329:
	;
	if base.Ui32(v1359) < base.Ui32(int32(8)) {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1370 = v1358
	v1371 = v1359
	v1373 = int32(0)
	goto L332
L331:
	;
	v1499 = v1483
	v1500 = v1484
	v1502 = v1486
	v1506 = base.B2i32(v1489 != int32(0))
	goto L327
L332:
	;
	v1379 = int32(base.Ui32(v1371) >> (uint(int32(3)) % 32))
	v1380 = int32(4)
	v1381 = v1370 + v1380
	if v1371&v1380 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	v1483 = v1474
	v1484 = v1475
	v1486 = v1459
	v1489 = v1464
	goto L331
L334:
	;
	v1464 = int32(0)
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1381+v1379+(v1464-v1379)&int32(3)+v1452<<(uint(int32(2))%32))))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1474)))
	if base.Ui32(v1475) < base.Ui32(int32(8)) {
		v1483 = v1474
		v1484 = v1475
		v1486 = v1459
		v1489 = v1464
		goto L331
	} else {
		goto L352
	}
L335:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327+v1373))))
	v1430 = int32(0)
	goto L346
L336:
	;
	v1386 = int32(0)
	if base.Ui32(v1347) <= base.Ui32(v1373) {
		v1419 = v1373
		v1422 = v1386
		goto L337
	} else {
		goto L338
	}
L337:
	;
	if v1422 == v1379 {
		v1452 = v1386
		v1459 = v1419
		goto L334
	} else {
		goto L344
	}
L338:
	;
	v1396 = v1373
	v1399 = v1386
	goto L339
L339:
	;
	v1402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381+v1399))))
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327+v1396))))
	if v1402 != v1404 {
		v1419 = v1396
		v1422 = v1399
		goto L337
	} else {
		goto L341
	}
L340:
	;
	v1419 = v1407
	v1422 = v1409
	goto L337
L341:
	;
	v1406 = int32(1)
	v1407 = v1396 + v1406
	v1409 = v1399 + v1406
	if base.Ui32(v1379) <= base.Ui32(v1409) {
		v1419 = v1407
		v1422 = v1409
		goto L337
	} else {
		goto L342
	}
L342:
	;
	if base.Ui32(v1407) < base.Ui32(v1347) {
		v1396 = v1407
		v1399 = v1409
		goto L339
	} else {
		goto L343
	}
L343:
	;
	goto L340
L344:
	;
	v1483 = v1370
	v1484 = v1371
	v1486 = v1419
	v1489 = v1422
	goto L331
L345:
	;
	if v1430 != v1379 {
		goto L350
	} else {
		goto L351
	}
L346:
	;
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381+v1430))))
	if v1443 == v1427&int32(255) {
		goto L345
	} else {
		goto L348
	}
L348:
	;
	v1445 = int32(1)
	v1447 = v1430 + v1445
	if v1447 != v1379 {
		v1430 = v1447
		goto L346
	} else {
		goto L349
	}
L349:
	;
	v1499 = v1370
	v1500 = v1371
	v1502 = v1373
	v1506 = v1445
	goto L327
L350:
	;
	v1452 = v1430
	v1459 = v1373 + int32(1)
	goto L334
L351:
	;
	v1483 = v1370
	v1484 = v1371
	v1486 = v1373
	v1489 = v1379
	goto L331
L352:
	;
	if base.Ui32(v1459) < base.Ui32(v1347) {
		v1370 = v1474
		v1371 = v1475
		v1373 = v1459
		goto L332
	} else {
		goto L353
	}
L353:
	;
	goto L333
L354:
	;
	goto L326
L355:
	;
	v1508 = int32(0)
	if v1500&int32(1) == v1508 {
		v1543 = v1508
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1514 = v1500 & int32(4)
	if v1506&base.B2i32(v1514 != int32(0)) != 0 {
		v1543 = v1508
		goto L354
	} else {
		goto L357
	}
L357:
	;
	v1518 = int32(1)
	if v1349 == int32(0) {
		v1543 = v1518
		goto L354
	} else {
		goto L358
	}
L358:
	;
	if v1500&int32(2) != 0 {
		v1540 = int32(0)
		goto L359
	} else {
		goto L360
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349))) = v1540
	v1543 = v1518
	goto L354
L360:
	;
	v1524 = int32(3)
	v1525 = int32(base.Ui32(v1500) >> (uint(v1524) % 32))
	if v1514 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1535 = int32(4)
	goto L363
L362:
	;
	v1535 = v1525 << (uint(int32(2)) % 32)
	goto L363
L363:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1525+(int32(0)-v1525)&v1524+v1535+int32(4))))
	v1540 = v1539
	goto L359
L364:
	;
	v1545 = F_getUpcomingChannelList(m, v1321, v1093)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v1545
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1321)))
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549+int32(-1)))))
	switch v1552 & int32(7) {
	case 0:
		goto L371
	case 1:
		goto L370
	case 2:
		goto L369
	case 3:
		goto L368
	case 4:
		goto L367
	default:
		v1581 = int32(0)
		goto L366
	}
L366:
	;
	v1583 = F_raxInsert(m, v1053, v1549, v1581, v1545, int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L376
	}
L367:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1549+int32(-17))))
	v1581 = v1580
	goto L366
L368:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1549+int32(-9))))
	v1576 = F_raxInsert(m, v1053, v1549, v1574, v1545, int32(0))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L1
	} else {
		goto L375
	}
L369:
	;
	v1568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1549+int32(-5)))))
	v1570 = F_raxInsert(m, v1053, v1549, v1568, v1545, int32(0))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L1
	} else {
		goto L374
	}
L370:
	;
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549+int32(-3)))))
	v1564 = F_raxInsert(m, v1053, v1549, v1562, v1545, int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	v1558 = F_raxInsert(m, v1053, v1549, int32(base.Ui32(v1552)>>(uint(int32(3))%32)), v1545, int32(0))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	goto L316
L373:
	;
	goto L316
L374:
	;
	goto L316
L375:
	;
	goto L316
L376:
	;
	goto L316
L377:
	;
	goto L316
L378:
	;
	v1594 = F_ACLShouldKillPubsubClient(m, v1092, v1591)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	if v1594 == int32(0) {
		goto L314
	} else {
		goto L380
	}
L380:
	;
	goto L315
L381:
	;
	F_freeClientOrCloseLater(m, v1092, int32(0))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+204)) = v1607 & int32(-8388609)
	goto L381
L385:
	;
	goto L268
L386:
	;
	if v1634 != 0 {
		v1080 = v1634
		goto L266
	} else {
		goto L389
	}
L387:
	;
	goto L386
L388:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1632)+4))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1634+base.B2i32(v1637 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1632))) = v1643
	goto L387
L389:
	;
	goto L267
L390:
	;
	F_raxFreeWithCallback(m, v112, int32(13))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L1
	} else {
		goto L393
	}
L391:
	;
	F_raxFreeWithCallback(m, v1053, int32(12))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	goto L390
L393:
	;
	F_sdsfree(m, v725)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v1682 = int32(0)
	goto L3
L395:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v112
	v1678 = F_sdscat(m, v725, int32(_a53))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1682 = v1678
	goto L3
}
func F_ACLResetFirstArgs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = int32(0)
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v19 = v12 << (uint(int32(2)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19)))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_valkey_free(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L16
	}
L5:
	;
	v60 = v12 + int32(1)
	if v60 != int32(1024) {
		v12 = v60
		goto L3
	} else {
		goto L15
	}
L6:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v25 == v24 {
		v48 = v21
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_valkey_free(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L11
	} else {
		goto L14
	}
L8:
	;
	v32 = v24
	v33 = v25
	goto L9
L9:
	;
	F_sdsfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v48 = v38
	goto L7
L11:
	;
	return
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36+v19)))
	v40 = v32 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+v40<<(uint(int32(2))%32))))
	if v44 != 0 {
		v32 = v40
		v33 = v44
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	goto L4
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L1
}
func F_ACLSetSelectorCommandBitsForCategory(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v11 = v8 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	if l0 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = F_hashtableNext(m, v8+int32(16), v8+int32(12))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L1
L3:
	;
	goto L2
L5:
	;
	F_hashtableCleanupIterator(m, v8+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L19
	}
L6:
	;
	return
L7:
	;
	if v34 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+64))
	if v44&l2 == int64(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+200))
	if v50 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_ACLChangeSelectorPerm(m, l1, v43, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v59 = F_hashtableNext(m, v8+int32(16), v8+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	F_ACLSetSelectorCommandBitsForCategory(m, v50, l1, l2, l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v59 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	m.G0 = v8 + int32(64)
	return
}
func F_ACLUpdateDefaultUserPassword(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v10 = F_ACLSetUser(m, v7, int32(_a103), int32(-1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if l0 == int32(0) {
			v67 = *(*int32)(unsafe.Add(mBase, _consts[22]))
			v70 = F_ACLSetUser(m, v67, int32(_a51), int32(-1))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				return
			}
		} else {
			v15 = F_sdsnew(m, int32(_a104))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
				switch v19 & int32(7) {
				case 0:
					v36 = int32(base.Ui32(v19) >> (uint(int32(3)) % 32))
				case 1:
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
					v36 = v26
				case 2:
					v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
					v36 = v29
				case 3:
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
					v36 = v32
				case 4:
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
					v36 = v35
				default:
					v36 = v2
				}
				v37 = F_sdscatlen(m, v15, l0, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = int32(0)
					v41 = *(*int32)(unsafe.Add(mBase, _consts[22]))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-1)))))
					switch v44 & int32(7) {
					case 0:
						v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
					case 1:
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+int32(-3)))))
						v61 = v51
					case 2:
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37+int32(-5)))))
						v61 = v54
					case 3:
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-9))))
						v61 = v57
					case 4:
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(-17))))
						v61 = v60
					default:
						v61 = v39
					}
					v62 = F_ACLSetUser(m, v41, v37, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_sdsfree(m, v37)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
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
func F_ACLUserCheckChannelPerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v159
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = v15 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v21
	goto L4
L3:
	;
	v159 = int32(0)
	goto L1
L4:
	;
	goto L5
L5:
	;
	v48 = v15 + int32(8)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v50 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v159 = v70
	goto L1
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v63&int32(8) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v50 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50+base.B2i32(v53 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
	goto L9
L11:
	;
	v159 = int32(5)
	goto L1
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+144))
	v70 = int32(0)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v72 & int32(7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v81 = v70
		goto L14
	}
L13:
	;
	v159 = int32(0)
	goto L1
L14:
	;
	v85 = m.G0
	v87 = v85 - int32(16)
	m.G0 = v87
	v90 = v87 + int32(8)
	F_listRewind(m, v69, v90)
	mBase = m.M
	v94 = F_listNext(m, v90)
	mBase = m.M
	if v94 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v81 = v80
	goto L14
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v81 = v79
	goto L14
L17:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v81 = v78
	goto L14
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v81 = v77
	goto L14
L19:
	;
	v81 = int32(base.Ui32(v72) >> (uint(int32(3)) % 32))
	goto L14
L20:
	;
	if v151 != 0 {
		goto L5
	} else {
		goto L38
	}
L21:
	;
	m.G0 = v87 + int32(16)
	goto L20
L22:
	;
	v151 = int32(5)
	goto L21
L23:
	;
	v97 = v94
	goto L24
L24:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+int32(-1)))))
	switch v108 & int32(7) {
	case 0:
		goto L31
	case 1:
		goto L30
	case 2:
		goto L29
	case 3:
		goto L28
	case 4:
		goto L27
	default:
		v125 = int32(0)
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	if l2 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v105+int32(-17))))
	v125 = v124
	goto L26
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v105+int32(-9))))
	v125 = v121
	goto L26
L29:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105+int32(-5)))))
	v125 = v118
	goto L26
L30:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+int32(-3)))))
	v125 = v115
	goto L26
L31:
	;
	v125 = int32(base.Ui32(v108) >> (uint(int32(3)) % 32))
	goto L26
L32:
	;
	v136 = F_listNext(m, v87+int32(8))
	mBase = m.M
	if v136 != 0 {
		v97 = v136
		goto L24
	} else {
		goto L37
	}
L33:
	;
	v130 = int32(0)
	v132 = F_stringmatchlen(m, v105, v125, l1, v81, v130)
	mBase = m.M
	if v132 != 0 {
		v151 = v130
		goto L21
	} else {
		goto L36
	}
L34:
	;
	v128 = F_strcmp(m, v105, l1)
	mBase = m.M
	if v128 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v151 = int32(0)
	goto L21
L36:
	;
	goto L32
L37:
	;
	goto L25
L38:
	;
	goto L6
}
func F_ACLUserCheckCmdWithUnrestrictedKeyAccess(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	v9 = m.G0
	v11 = v9 - int32(2096)
	m.G0 = v11
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(2096)
	return v187
L2:
	;
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v11 + int32(2080)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20
	goto L4
L3:
	;
	v187 = int32(1)
	goto L1
L4:
	;
	v25 = v11 + int32(2080)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v27 == int32(0) {
		v187 = v14
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27+base.B2i32(v30 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	goto L6
L8:
	;
	v40 = int32(2)
	if l5&int32(160) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = v40
	goto L11
L10:
	;
	v47 = int32(base.Ui32(l5)>>(uint(int32(5))%32)) & v40
	goto L11
L11:
	;
	v52 = v47 | int32(base.Ui32(l5)>>(uint(int32(4))%32))&int32(1)
	v53 = v27
	goto L14
L12:
	;
	F_getKeysFreeResult(m, v11+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L17
	} else {
		goto L42
	}
L13:
	;
	v164 = int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v165 == int32(0) {
		v187 = v164
		goto L1
	} else {
		goto L41
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v66 = F_ACLSelectorCheckCmd(m, v61, l1, l2, l3, v11+int32(2076), v11+int32(12), l4)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v152 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v153 == v152 {
		v187 = v152
		goto L1
	} else {
		goto L40
	}
L16:
	;
	v139 = v11 + int32(2080)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v141 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L17:
	;
	return int32(0)
L18:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v70&int32(2) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+140))
	v75 = v11 + int32(2088)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76
	goto L21
L21:
	;
	v81 = v11 + int32(2088)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v83 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v83 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83+base.B2i32(v86 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v92
	goto L23
L25:
	;
	v96 = v83
	goto L26
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105&v52 != v52 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L16
L28:
	;
	v117 = v11 + int32(2088)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v119 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 != int32(42) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v112 == int32(0) {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if v119 != 0 {
		v96 = v119
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v119+base.B2i32(v122 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v128
	goto L33
L35:
	;
	goto L27
L36:
	;
	if v141 != 0 {
		v53 = v141
		goto L14
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v141+base.B2i32(v144 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v150
	goto L37
L39:
	;
	goto L15
L40:
	;
	v175 = v152
	goto L12
L41:
	;
	v175 = v164
	goto L12
L42:
	;
	v187 = v175
	goto L1
}
func F_ACLUserGetRootSelector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v3 == int32(0) {
		F__serverAssert(m, int32(_a18), int32(_a4), int32(421))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v8&int32(1) == int32(0) {
			F__serverAssert(m, int32(_a19), int32(_a4), int32(423))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return v7
		}
	}
}
func F_aclCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int64
	_ = v540
	var v546 int32
	_ = v546
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int64
	_ = v959
	var v962 int32
	_ = v962
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int64
	_ = v973
	var v982 int64
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1000 int64
	_ = v1000
	var v1012 int32
	_ = v1012
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int64
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1383 int64
	_ = v1383
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1775 int64
	_ = v1775
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1872 int64
	_ = v1872
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1917 int32
	_ = v1917
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2129 int64
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2159 int64
	_ = v2159
	var v2165 int32
	_ = v2165
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2187 int64
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2234 int32
	_ = v2234
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2253 int64
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2303 int64
	_ = v2303
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2369 int32
	_ = v2369
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
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
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2545 int64
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2575 int32
	_ = v2575
	var v2581 int32
	_ = v2581
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2591 int32
	_ = v2591
	var v2592 int64
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2602 int32
	_ = v2602
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2642 int32
	_ = v2642
	var v2645 int32
	_ = v2645
	var v2648 int32
	_ = v2648
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2668 int32
	_ = v2668
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2684 int64
	_ = v2684
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2705 int32
	_ = v2705
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2721 int64
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2726 int32
	_ = v2726
	var v2727 int64
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2733 int64
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2779 int32
	_ = v2779
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2794 int32
	_ = v2794
	var v2801 int32
	_ = v2801
	var v2804 int32
	_ = v2804
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2819 int32
	_ = v2819
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2843 int32
	_ = v2843
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2856 int32
	_ = v2856
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2889 int32
	_ = v2889
	var v2892 int32
	_ = v2892
	var v2897 int32
	_ = v2897
	var v2900 int32
	_ = v2900
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2922 int32
	_ = v2922
	var v2929 int32
	_ = v2929
	var v2934 int32
	_ = v2934
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2963 int32
	_ = v2963
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2976 int32
	_ = v2976
	var v2984 int32
	_ = v2984
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v3005 int32
	_ = v3005
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3024 int32
	_ = v3024
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3046 int32
	_ = v3046
	var v3052 int32
	_ = v3052
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3146 int32
	_ = v3146
	var v3151 int32
	_ = v3151
	var v3153 int32
	_ = v3153
	var v3158 int32
	_ = v3158
	var v3164 int32
	_ = v3164
	var v3170 int32
	_ = v3170
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3220 int32
	_ = v3220
	var v3222 int32
	_ = v3222
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3237 int64
	_ = v3237
	var v3247 int32
	_ = v3247
	v18 = m.G0
	v20 = v18 - int32(1200)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = F_objectGetVal(m, v23)
	mBase = m.M
	v25 = int32(_a54)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v20 + int32(1200)
	return
L2:
	;
	v473 = int32(_a55)
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v476 != 0 {
		goto L98
	} else {
		goto L99
	}
L3:
	;
	if v60-v62 != 0 {
		goto L2
	} else {
		goto L15
	}
L4:
	;
	v60 = F_tolower(m, v56)
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = F_tolower(m, v61)
	mBase = m.M
	goto L3
L5:
	;
	v30 = v24
	v31 = v25
	v32 = v28
	goto L8
L6:
	;
	v56 = int32(0)
	v57 = v25
	goto L4
L7:
	;
	v56 = v53 & int32(255)
	v57 = v52
	goto L4
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(0) {
		v52 = v31
		v53 = v32
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v52 = v46
	v53 = int32(0)
	goto L7
L10:
	;
	v38 = v32 & int32(255)
	if v38 == v34 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = int32(1)
	v46 = v31 + v45
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v47 != 0 {
		v30 = v30 + v45
		v31 = v46
		v32 = v47
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v40 = F_tolower(m, v38)
	mBase = m.M
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v42 = F_tolower(m, v41)
	mBase = m.M
	if v40 == v42 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v52 = v31
	v53 = v44
	goto L7
L14:
	;
	goto L9
L15:
	;
	v64 = int32(2)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v65 <= v64 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v71 = v64
	goto L17
L17:
	;
	F_redactClientCommandArgument(m, l0, v71)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v94 = F_objectGetVal(m, v93)
	mBase = m.M
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-1)))))
	v99 = v97 & int32(7)
	switch v99 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	default:
		v187 = int32(0)
		goto L22
	}
L19:
	;
	return
L20:
	;
	v88 = v71 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v88 < v89 {
		v71 = v88
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v201 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v201
	v204 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v206 = v20 + int32(176)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	if v187 == v201 {
		goto L45
	} else {
		goto L46
	}
L23:
	;
	if v114 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-17))))
	v114 = v113
	goto L23
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-9))))
	v114 = v110
	goto L23
L26:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-5)))))
	v114 = v107
	goto L23
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-3)))))
	v114 = v104
	goto L23
L28:
	;
	v114 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
	goto L23
L29:
	;
	switch v99 {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	case 3:
		goto L39
	case 4:
		goto L38
	default:
		v187 = int32(0)
		goto L22
	}
L30:
	;
	v121 = int32(0)
	goto L32
L31:
	;
	F_addReplyError(m, l0, int32(_a56))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L19
	} else {
		goto L37
	}
L32:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v121))))
	if v136&int32(223) == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(base.I32_extend8_s(v136)+int32(-9)) < base.Ui32(int32(5)) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v147 = v121 + int32(1)
	if v147 == v114 {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v121 = v147
	goto L32
L37:
	;
	goto L1
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-17))))
	v187 = v183
	goto L22
L39:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v94+int32(-9))))
	v187 = v180
	goto L22
L40:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(-5)))))
	v187 = v177
	goto L22
L41:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+int32(-3)))))
	v187 = v174
	goto L22
L42:
	;
	v187 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
	goto L22
L43:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v406 = F_valkey_malloc(m, v403<<(uint(int32(2))%32))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L19
	} else {
		goto L81
	}
L44:
	;
	if v359 != v187 {
		goto L71
	} else {
		goto L72
	}
L45:
	;
	v350 = int32(0)
	v356 = v215
	v357 = v216
	v359 = v350
	v363 = v350
	goto L44
L46:
	;
	if base.Ui32(v216) < base.Ui32(int32(8)) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v227 = v215
	v228 = v216
	v230 = int32(0)
	goto L49
L48:
	;
	v356 = v340
	v357 = v341
	v359 = v343
	v363 = base.B2i32(v346 != int32(0))
	goto L44
L49:
	;
	v236 = int32(base.Ui32(v228) >> (uint(int32(3)) % 32))
	v237 = int32(4)
	v238 = v227 + v237
	if v228&v237 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v340 = v331
	v341 = v332
	v343 = v316
	v346 = v321
	goto L48
L51:
	;
	v321 = int32(0)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v238+v236+(v321-v236)&int32(3)+v309<<(uint(int32(2))%32))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if base.Ui32(v332) < base.Ui32(int32(8)) {
		v340 = v331
		v341 = v332
		v343 = v316
		v346 = v321
		goto L48
	} else {
		goto L69
	}
L52:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v230))))
	v287 = int32(0)
	goto L63
L53:
	;
	v243 = int32(0)
	if base.Ui32(v187) <= base.Ui32(v230) {
		v276 = v230
		v279 = v243
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v279 == v236 {
		v309 = v243
		v316 = v276
		goto L51
	} else {
		goto L61
	}
L55:
	;
	v253 = v230
	v256 = v243
	goto L56
L56:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v256))))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v253))))
	if v259 != v261 {
		v276 = v253
		v279 = v256
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v276 = v264
	v279 = v266
	goto L54
L58:
	;
	v263 = int32(1)
	v264 = v253 + v263
	v266 = v256 + v263
	if base.Ui32(v236) <= base.Ui32(v266) {
		v276 = v264
		v279 = v266
		goto L54
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v264) < base.Ui32(v187) {
		v253 = v264
		v256 = v266
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v340 = v227
	v341 = v228
	v343 = v276
	v346 = v279
	goto L48
L62:
	;
	if v287 != v236 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v287))))
	if v300 == v284&int32(255) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v302 = int32(1)
	v304 = v287 + v302
	if v304 != v236 {
		v287 = v304
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v356 = v227
	v357 = v228
	v359 = v230
	v363 = v302
	goto L44
L67:
	;
	v309 = v287
	v316 = v230 + int32(1)
	goto L51
L68:
	;
	v340 = v227
	v341 = v228
	v343 = v230
	v346 = v236
	goto L48
L69:
	;
	if base.Ui32(v316) < base.Ui32(v187) {
		v227 = v331
		v228 = v332
		v230 = v316
		goto L49
	} else {
		goto L70
	}
L70:
	;
	goto L50
L71:
	;
	goto L43
L72:
	;
	if v357&int32(1) == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v371 = v357 & int32(4)
	if v363&base.B2i32(v371 != int32(0)) != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	if v206 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	if v357&int32(2) != 0 {
		v397 = int32(0)
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v397
	goto L71
L77:
	;
	v381 = int32(3)
	v382 = int32(base.Ui32(v357) >> (uint(v381) % 32))
	if v371 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v392 = int32(4)
	goto L80
L79:
	;
	v392 = v382 << (uint(int32(2)) % 32)
	goto L80
L80:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v356+v382+(int32(0)-v382)&v381+v392+int32(4))))
	v397 = v396
	goto L76
L81:
	;
	v408 = int32(3)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v409 <= v408 {
		v445 = v409
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v462 = F_ACLStringSetUser(m, v402, v94, v406, v445+int32(-3))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L19
	} else {
		goto L87
	}
L83:
	;
	v417 = v408
	goto L84
L84:
	;
	v432 = v417 << (uint(int32(2)) % 32)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v434+v432)))
	v437 = F_objectGetVal(m, v436)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v406+int32(-12)+v432))) = v437
	v440 = v417 + int32(1)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v440 < v441 {
		v417 = v440
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v445 = v441
	goto L82
L86:
	;
	goto L85
L87:
	;
	F_valkey_free(m, v406)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L19
	} else {
		goto L88
	}
L88:
	;
	if v462 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_addReplyErrorSdsSafe(m, l0, v462)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L19
	} else {
		goto L92
	}
L90:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
	} else {
		goto L91
	}
L91:
	;
	goto L1
L92:
	;
	goto L1
L93:
	;
	F_addReplyLongLong(m, l0, v3237)
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L19
	} else {
		goto L815
	}
L94:
	;
	v3237 = base.I64_extend_i32_s(v668)
	goto L93
L95:
	;
	v674 = int32(_a57)
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v677 != 0 {
		goto L153
	} else {
		goto L154
	}
L96:
	;
	if v508-v510 != 0 {
		goto L95
	} else {
		goto L108
	}
L97:
	;
	v508 = F_tolower(m, v504)
	mBase = m.M
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v510 = F_tolower(m, v509)
	mBase = m.M
	goto L96
L98:
	;
	v478 = v24
	v479 = v473
	v480 = v476
	goto L101
L99:
	;
	v504 = int32(0)
	v505 = v473
	goto L97
L100:
	;
	v504 = v501 & int32(255)
	v505 = v500
	goto L97
L101:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v482 == int32(0) {
		v500 = v479
		v501 = v480
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v500 = v494
	v501 = int32(0)
	goto L100
L103:
	;
	v486 = v480 & int32(255)
	if v486 == v482 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v493 = int32(1)
	v494 = v479 + v493
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+1)))
	if v495 != 0 {
		v478 = v478 + v493
		v479 = v494
		v480 = v495
		goto L101
	} else {
		goto L107
	}
L105:
	;
	v488 = F_tolower(m, v486)
	mBase = m.M
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	v490 = F_tolower(m, v489)
	mBase = m.M
	if v488 == v490 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v500 = v479
	v501 = v492
	goto L100
L107:
	;
	goto L102
L108:
	;
	v512 = int32(2)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v513 <= v512 {
		goto L95
	} else {
		goto L109
	}
L109:
	;
	v519 = v512
	goto L110
L110:
	;
	F_redactClientCommandArgument(m, l0, v519)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L19
	} else {
		goto L112
	}
L111:
	;
	v539 = int32(2)
	v540 = int64(0)
	if v537 <= v539 {
		v3237 = v540
		goto L93
	} else {
		goto L114
	}
L112:
	;
	v536 = v519 + int32(1)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v536 < v537 {
		v519 = v536
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v546 = v539
	goto L117
L115:
	;
	v613 = v604
	v617 = int32(0)
	goto L131
L116:
	;
	F_addReplyError(m, l0, int32(_a58))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L19
	} else {
		goto L130
	}
L117:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560+v546<<(uint(int32(2))%32))))
	v565 = F_objectGetVal(m, v564)
	mBase = m.M
	v566 = int32(_a46)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _consts[21])))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v570 == int32(0) {
		v593 = v569
		v594 = v570
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v604 = int32(2)
	if v602 <= v604 {
		v3237 = v540
		goto L93
	} else {
		goto L129
	}
L119:
	;
	if v594-v593&int32(255) == int32(0) {
		goto L116
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	if v570 != v569&int32(255) {
		v593 = v569
		v594 = v570
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v576 = v565
	v577 = v566
	goto L123
L123:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v581 == int32(0) {
		v593 = v580
		v594 = v581
		goto L120
	} else {
		goto L125
	}
L124:
	;
	v593 = v580
	v594 = v581
	goto L120
L125:
	;
	v584 = int32(1)
	if v581 == v580&int32(255) {
		v576 = v576 + v584
		v577 = v577 + v584
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v601 = v546 + int32(1)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v601 < v602 {
		v546 = v601
		goto L117
	} else {
		goto L128
	}
L128:
	;
	goto L118
L129:
	;
	goto L115
L130:
	;
	goto L1
L131:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628+v613<<(uint(int32(2))%32))))
	v633 = F_objectGetVal(m, v632)
	mBase = m.M
	v634 = int32(0)
	v636 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+int32(-1)))))
	switch v639 & int32(7) {
	case 0:
		goto L138
	case 1:
		goto L137
	case 2:
		goto L136
	case 3:
		goto L135
	case 4:
		goto L134
	default:
		v656 = v634
		goto L133
	}
L133:
	;
	v659 = F_raxRemove(m, v636, v633, v656, v20+int32(176))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L19
	} else {
		goto L140
	}
L134:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v633+int32(-17))))
	v656 = v655
	goto L133
L135:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v633+int32(-9))))
	v656 = v652
	goto L133
L136:
	;
	v649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633+int32(-5)))))
	v656 = v649
	goto L133
L137:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+int32(-3)))))
	v656 = v646
	goto L133
L138:
	;
	v656 = int32(base.Ui32(v639) >> (uint(int32(3)) % 32))
	goto L133
L139:
	;
	v670 = v613 + int32(1)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v671 <= v670 {
		goto L94
	} else {
		goto L143
	}
L140:
	;
	if v659 == int32(0) {
		v668 = v617
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	F_ACLFreeUserAndKillClients(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L19
	} else {
		goto L142
	}
L142:
	;
	v668 = v617 + int32(1)
	goto L139
L143:
	;
	v613 = v670
	v617 = v668
	goto L131
L144:
	;
	if v1878 == int32(0) {
		goto L810
	} else {
		goto L811
	}
L145:
	;
	v3200 = F_close(m, v1885)
	mBase = m.M
	goto L144
L146:
	;
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3190 = *(*int32)(unsafe.Add(mBase, uint32(v3189)+8))
	v3191 = F_objectGetVal(m, v3190)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v3191
	F_addReplyErrorFormat(m, l0, int32(_a59), v20+int32(160))
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L19
	} else {
		goto L809
	}
L147:
	;
	F__serverAssert(m, int32(_a60), int32(_a4), int32(3204))
	mBase = m.M
	v3170 = m.ExcPending
	if v3170 != 0 {
		goto L19
	} else {
		goto L808
	}
L148:
	;
	F__serverAssert(m, int32(_a19), int32(_a4), int32(423))
	mBase = m.M
	v3164 = m.ExcPending
	if v3164 != 0 {
		goto L19
	} else {
		goto L807
	}
L149:
	;
	F__serverAssert(m, int32(_a18), int32(_a4), int32(421))
	mBase = m.M
	v3158 = m.ExcPending
	if v3158 != 0 {
		goto L19
	} else {
		goto L806
	}
L150:
	;
	v1246 = int32(_a61)
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1249 != 0 {
		goto L275
	} else {
		goto L276
	}
L151:
	;
	if v709-v711 != 0 {
		goto L150
	} else {
		goto L163
	}
L152:
	;
	v709 = F_tolower(m, v705)
	mBase = m.M
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	v711 = F_tolower(m, v710)
	mBase = m.M
	goto L151
L153:
	;
	v679 = v24
	v680 = v674
	v681 = v677
	goto L156
L154:
	;
	v705 = int32(0)
	v706 = v674
	goto L152
L155:
	;
	v705 = v702 & int32(255)
	v706 = v701
	goto L152
L156:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	if v683 == int32(0) {
		v701 = v680
		v702 = v681
		goto L155
	} else {
		goto L158
	}
L157:
	;
	v701 = v695
	v702 = int32(0)
	goto L155
L158:
	;
	v687 = v681 & int32(255)
	if v687 == v683 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v694 = int32(1)
	v695 = v680 + v694
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+1)))
	if v696 != 0 {
		v679 = v679 + v694
		v680 = v695
		v681 = v696
		goto L156
	} else {
		goto L162
	}
L160:
	;
	v689 = F_tolower(m, v687)
	mBase = m.M
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	v691 = F_tolower(m, v690)
	mBase = m.M
	if v689 == v691 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	v701 = v680
	v702 = v693
	goto L155
L162:
	;
	goto L157
L163:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v713 != int32(3) {
		goto L150
	} else {
		goto L164
	}
L164:
	;
	F_redactClientCommandArgument(m, l0, int32(2))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+8))
	v721 = F_objectGetVal(m, v720)
	mBase = m.M
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	v725 = F_objectGetVal(m, v724)
	mBase = m.M
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+int32(-1)))))
	switch v728 & int32(7) {
	case 0:
		goto L171
	case 1:
		goto L170
	case 2:
		goto L169
	case 3:
		goto L168
	case 4:
		goto L167
	default:
		v745 = int32(0)
		goto L166
	}
L166:
	;
	v746 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v746
	v749 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v751 = v20 + int32(176)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	if v745 == v746 {
		goto L174
	} else {
		goto L175
	}
L167:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v725+int32(-17))))
	v745 = v744
	goto L166
L168:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v725+int32(-9))))
	v745 = v741
	goto L166
L169:
	;
	v738 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v725+int32(-5)))))
	v745 = v738
	goto L166
L170:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+int32(-3)))))
	v745 = v735
	goto L166
L171:
	;
	v745 = int32(base.Ui32(v728) >> (uint(int32(3)) % 32))
	goto L166
L172:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	if v947 != 0 {
		goto L210
	} else {
		goto L211
	}
L173:
	;
	if v904 != v745 {
		goto L200
	} else {
		goto L201
	}
L174:
	;
	v895 = int32(0)
	v901 = v760
	v902 = v761
	v904 = v895
	v908 = v895
	goto L173
L175:
	;
	if base.Ui32(v761) < base.Ui32(int32(8)) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v772 = v760
	v773 = v761
	v775 = int32(0)
	goto L178
L177:
	;
	v901 = v885
	v902 = v886
	v904 = v888
	v908 = base.B2i32(v891 != int32(0))
	goto L173
L178:
	;
	v781 = int32(base.Ui32(v773) >> (uint(int32(3)) % 32))
	v782 = int32(4)
	v783 = v772 + v782
	if v773&v782 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v885 = v876
	v886 = v877
	v888 = v861
	v891 = v866
	goto L177
L180:
	;
	v866 = int32(0)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v783+v781+(v866-v781)&int32(3)+v854<<(uint(int32(2))%32))))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	if base.Ui32(v877) < base.Ui32(int32(8)) {
		v885 = v876
		v886 = v877
		v888 = v861
		v891 = v866
		goto L177
	} else {
		goto L198
	}
L181:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721+v775))))
	v832 = int32(0)
	goto L192
L182:
	;
	v788 = int32(0)
	if base.Ui32(v745) <= base.Ui32(v775) {
		v821 = v775
		v824 = v788
		goto L183
	} else {
		goto L184
	}
L183:
	;
	if v824 == v781 {
		v854 = v788
		v861 = v821
		goto L180
	} else {
		goto L190
	}
L184:
	;
	v798 = v775
	v801 = v788
	goto L185
L185:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v801))))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721+v798))))
	if v804 != v806 {
		v821 = v798
		v824 = v801
		goto L183
	} else {
		goto L187
	}
L186:
	;
	v821 = v809
	v824 = v811
	goto L183
L187:
	;
	v808 = int32(1)
	v809 = v798 + v808
	v811 = v801 + v808
	if base.Ui32(v781) <= base.Ui32(v811) {
		v821 = v809
		v824 = v811
		goto L183
	} else {
		goto L188
	}
L188:
	;
	if base.Ui32(v809) < base.Ui32(v745) {
		v798 = v809
		v801 = v811
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	v885 = v772
	v886 = v773
	v888 = v821
	v891 = v824
	goto L177
L191:
	;
	if v832 != v781 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v832))))
	if v845 == v829&int32(255) {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	v847 = int32(1)
	v849 = v832 + v847
	if v849 != v781 {
		v832 = v849
		goto L192
	} else {
		goto L195
	}
L195:
	;
	v901 = v772
	v902 = v773
	v904 = v775
	v908 = v847
	goto L173
L196:
	;
	v854 = v832
	v861 = v775 + int32(1)
	goto L180
L197:
	;
	v885 = v772
	v886 = v773
	v888 = v775
	v891 = v781
	goto L177
L198:
	;
	if base.Ui32(v861) < base.Ui32(v745) {
		v772 = v876
		v773 = v877
		v775 = v861
		goto L178
	} else {
		goto L199
	}
L199:
	;
	goto L179
L200:
	;
	goto L172
L201:
	;
	if v902&int32(1) == int32(0) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v916 = v902 & int32(4)
	if v908&base.B2i32(v916 != int32(0)) != 0 {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	if v751 == int32(0) {
		goto L200
	} else {
		goto L204
	}
L204:
	;
	if v902&int32(2) != 0 {
		v942 = int32(0)
		goto L205
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751))) = v942
	goto L200
L206:
	;
	v926 = int32(3)
	v927 = int32(base.Ui32(v902) >> (uint(v926) % 32))
	if v916 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v937 = int32(4)
	goto L209
L208:
	;
	v937 = v927 << (uint(int32(2)) % 32)
	goto L209
L209:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v901+v927+(int32(0)-v927)&v926+v937+int32(4))))
	v942 = v941
	goto L205
L210:
	;
	v950 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L19
	} else {
		goto L213
	}
L211:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L19
	} else {
		goto L212
	}
L212:
	;
	goto L1
L213:
	;
	F_addReplyBulkCString(m, l0, int32(_a62))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L19
	} else {
		goto L214
	}
L214:
	;
	v956 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v959 = *(*int64)(unsafe.Add(mBase, _consts[4]))
	if v959 == int64(0) {
		v1012 = int32(0)
		goto L216
	} else {
		goto L217
	}
L216:
	;
	F_setDeferredSetLen(m, l0, v956, v1012)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L19
	} else {
		goto L224
	}
L217:
	;
	v962 = int32(0)
	v968 = v962
	v969 = int32(_a5)
	v972 = v962
	v973 = v959
	goto L218
L218:
	;
	v982 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v947)+4)))
	if v973&v982 == int64(0) {
		v991 = v972
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v1012 = v991
	goto L216
L220:
	;
	v993 = v968 + int32(1)
	v995 = v993 << (uint(int32(4)) % 32)
	v1000 = *(*int64)(unsafe.Add(mBase, uint32(v995)+uint32(_consts[4])))
	if base.B2i32(v1000 == int64(0)) == int32(0) {
		v968 = v993
		v969 = v995 + int32(_a5)
		v972 = v991
		v973 = v1000
		goto L218
	} else {
		goto L223
	}
L221:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	F_addReplyBulkCString(m, l0, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L19
	} else {
		goto L222
	}
L222:
	;
	v991 = v972 + int32(1)
	goto L220
L223:
	;
	goto L219
L224:
	;
	F_addReplyBulkCString(m, l0, int32(_a63))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L19
	} else {
		goto L225
	}
L225:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v947)+8))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+20))
	F_addReplyArrayLen(m, l0, v1028)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L19
	} else {
		goto L226
	}
L226:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v947)+8))
	v1033 = v20 + int32(176)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	*(*int32)(unsafe.Add(mBase, uint32(v1033)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1033))) = v1034
	goto L227
L227:
	;
	v1039 = v20 + int32(176)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	if v1041 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v947)+12))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+20))
	if v1127 == int32(0) {
		goto L149
	} else {
		goto L246
	}
L229:
	;
	if v1041 == int32(0) {
		goto L228
	} else {
		goto L232
	}
L230:
	;
	goto L229
L231:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1041+base.B2i32(v1044 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1039))) = v1050
	goto L230
L232:
	;
	v1056 = v1041
	goto L233
L233:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+8))
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072+int32(-1)))))
	switch v1075 & int32(7) {
	case 0:
		goto L240
	case 1:
		goto L239
	case 2:
		goto L238
	case 3:
		goto L237
	case 4:
		goto L236
	default:
		v1092 = int32(0)
		goto L235
	}
L234:
	;
	goto L228
L235:
	;
	F_addReplyBulkCBuffer(m, l0, v1072, v1092)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L19
	} else {
		goto L241
	}
L236:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1072+int32(-17))))
	v1092 = v1091
	goto L235
L237:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1072+int32(-9))))
	v1092 = v1088
	goto L235
L238:
	;
	v1085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1072+int32(-5)))))
	v1092 = v1085
	goto L235
L239:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072+int32(-3)))))
	v1092 = v1082
	goto L235
L240:
	;
	v1092 = int32(base.Ui32(v1075) >> (uint(int32(3)) % 32))
	goto L235
L241:
	;
	v1096 = v20 + int32(176)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	if v1098 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v1098 != 0 {
		v1056 = v1098
		goto L233
	} else {
		goto L245
	}
L243:
	;
	goto L242
L244:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1098+base.B2i32(v1101 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1107
	goto L243
L245:
	;
	goto L234
L246:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+8))
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1131))))
	if v1132&int32(1) == int32(0) {
		goto L148
	} else {
		goto L247
	}
L247:
	;
	F_aclAddReplySelectorDescription(m, l0, v1131)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L19
	} else {
		goto L248
	}
L248:
	;
	F_addReplyBulkCString(m, l0, int32(_a64))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L19
	} else {
		goto L249
	}
L249:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v947)+12))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+20))
	F_addReplyArrayLen(m, l0, v1143+int32(-1))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L19
	} else {
		goto L250
	}
L250:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v947)+12))
	v1150 = v20 + int32(176)
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1148)))
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1150))) = v1151
	goto L251
L251:
	;
	v1156 = v20 + int32(176)
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1156)))
	if v1158 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	if v1158 == int32(0) {
		goto L147
	} else {
		goto L255
	}
L253:
	;
	goto L252
L254:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+4))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1158+base.B2i32(v1161 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1156))) = v1167
	goto L253
L255:
	;
	v1172 = v20 + int32(176)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	if v1174 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	F_setDeferredMapLen(m, l0, v950, int32(7))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L19
	} else {
		goto L270
	}
L257:
	;
	if v1174 == int32(0) {
		goto L256
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+4))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1174+base.B2i32(v1177 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1172))) = v1183
	goto L258
L260:
	;
	v1190 = v1174
	goto L261
L261:
	;
	v1204 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L19
	} else {
		goto L263
	}
L262:
	;
	goto L256
L263:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+8))
	F_aclAddReplySelectorDescription(m, l0, v1206)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L19
	} else {
		goto L264
	}
L264:
	;
	F_setDeferredMapLen(m, l0, v1204, int32(4))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L19
	} else {
		goto L265
	}
L265:
	;
	v1213 = v20 + int32(176)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1213)))
	if v1215 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	if v1215 != 0 {
		v1190 = v1215
		goto L261
	} else {
		goto L269
	}
L267:
	;
	goto L266
L268:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+4))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1215+base.B2i32(v1218 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1213))) = v1224
	goto L267
L269:
	;
	goto L262
L270:
	;
	goto L1
L271:
	;
	v1505 = int32(_a65)
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1508 != 0 {
		goto L348
	} else {
		goto L349
	}
L272:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1326 != int32(2) {
		goto L271
	} else {
		goto L299
	}
L273:
	;
	if v1281-v1283 == int32(0) {
		goto L272
	} else {
		goto L285
	}
L274:
	;
	v1281 = F_tolower(m, v1277)
	mBase = m.M
	v1282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1278))))
	v1283 = F_tolower(m, v1282)
	mBase = m.M
	goto L273
L275:
	;
	v1251 = v24
	v1252 = v1246
	v1253 = v1249
	goto L278
L276:
	;
	v1277 = int32(0)
	v1278 = v1246
	goto L274
L277:
	;
	v1277 = v1274 & int32(255)
	v1278 = v1273
	goto L274
L278:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252))))
	if v1255 == int32(0) {
		v1273 = v1252
		v1274 = v1253
		goto L277
	} else {
		goto L280
	}
L279:
	;
	v1273 = v1267
	v1274 = int32(0)
	goto L277
L280:
	;
	v1259 = v1253 & int32(255)
	if v1259 == v1255 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1266 = int32(1)
	v1267 = v1252 + v1266
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251)+1)))
	if v1268 != 0 {
		v1251 = v1251 + v1266
		v1252 = v1267
		v1253 = v1268
		goto L278
	} else {
		goto L284
	}
L282:
	;
	v1261 = F_tolower(m, v1259)
	mBase = m.M
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252))))
	v1263 = F_tolower(m, v1262)
	mBase = m.M
	if v1261 == v1263 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251))))
	v1273 = v1252
	v1274 = v1265
	goto L277
L284:
	;
	goto L279
L285:
	;
	v1287 = int32(_a66)
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1290 != 0 {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	if v1322-v1324 != 0 {
		goto L271
	} else {
		goto L298
	}
L287:
	;
	v1322 = F_tolower(m, v1318)
	mBase = m.M
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319))))
	v1324 = F_tolower(m, v1323)
	mBase = m.M
	goto L286
L288:
	;
	v1292 = v24
	v1293 = v1287
	v1294 = v1290
	goto L291
L289:
	;
	v1318 = int32(0)
	v1319 = v1287
	goto L287
L290:
	;
	v1318 = v1315 & int32(255)
	v1319 = v1314
	goto L287
L291:
	;
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293))))
	if v1296 == int32(0) {
		v1314 = v1293
		v1315 = v1294
		goto L290
	} else {
		goto L293
	}
L292:
	;
	v1314 = v1308
	v1315 = int32(0)
	goto L290
L293:
	;
	v1300 = v1294 & int32(255)
	if v1300 == v1296 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1307 = int32(1)
	v1308 = v1293 + v1307
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+1)))
	if v1309 != 0 {
		v1292 = v1292 + v1307
		v1293 = v1308
		v1294 = v1309
		goto L291
	} else {
		goto L297
	}
L295:
	;
	v1302 = F_tolower(m, v1300)
	mBase = m.M
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293))))
	v1304 = F_tolower(m, v1303)
	mBase = m.M
	if v1302 == v1304 {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292))))
	v1314 = v1293
	v1315 = v1306
	goto L290
L297:
	;
	goto L292
L298:
	;
	goto L272
L299:
	;
	v1329 = int32(_a66)
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1332 != 0 {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1370 = *(*int64)(unsafe.Add(mBase, uint32(v1369)+8))
	goto L312
L301:
	;
	v1364 = F_tolower(m, v1360)
	mBase = m.M
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1361))))
	v1366 = F_tolower(m, v1365)
	mBase = m.M
	goto L300
L302:
	;
	v1334 = v24
	v1335 = v1329
	v1336 = v1332
	goto L305
L303:
	;
	v1360 = int32(0)
	v1361 = v1329
	goto L301
L304:
	;
	v1360 = v1357 & int32(255)
	v1361 = v1356
	goto L301
L305:
	;
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335))))
	if v1338 == int32(0) {
		v1356 = v1335
		v1357 = v1336
		goto L304
	} else {
		goto L307
	}
L306:
	;
	v1356 = v1350
	v1357 = int32(0)
	goto L304
L307:
	;
	v1342 = v1336 & int32(255)
	if v1342 == v1338 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1349 = int32(1)
	v1350 = v1335 + v1349
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334)+1)))
	if v1351 != 0 {
		v1334 = v1334 + v1349
		v1335 = v1350
		v1336 = v1351
		goto L305
	} else {
		goto L311
	}
L309:
	;
	v1344 = F_tolower(m, v1342)
	mBase = m.M
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335))))
	v1346 = F_tolower(m, v1345)
	mBase = m.M
	if v1344 == v1346 {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334))))
	v1356 = v1335
	v1357 = v1348
	goto L304
L311:
	;
	goto L306
L312:
	;
	F_addReplyArrayLen(m, l0, base.I32_wrap_i64(v1370))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L19
	} else {
		goto L313
	}
L313:
	;
	v1375 = v20 + int32(176)
	v1377 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+4)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v1375))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+20)) = int32(128)
	v1383 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1375)+12)) = v1383
	*(*int64)(unsafe.Add(mBase, uint32(v1375)+296)) = v1383
	*(*int64)(unsafe.Add(mBase, uint32(v1375)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+8)) = v20 + int32(200)
	*(*int32)(unsafe.Add(mBase, uint32(v1375)+156)) = v20 + int32(344)
	goto L314
L314:
	;
	v1398 = int32(0)
	v1400 = F_raxSeek(m, v20+int32(176), int32(_a67), v1398, v1398)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L19
	} else {
		goto L315
	}
L315:
	;
	v1404 = F_raxNext(m, v20+int32(176))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L19
	} else {
		goto L317
	}
L316:
	;
	F_raxStop(m, v20+int32(176))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L19
	} else {
		goto L344
	}
L317:
	;
	if v1404 == int32(0) {
		goto L316
	} else {
		goto L318
	}
L318:
	;
	goto L319
L319:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v20)+188))
	if v1364-v1366 != 0 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	goto L316
L321:
	;
	v1482 = F_raxNext(m, v20+int32(176))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L19
	} else {
		goto L342
	}
L322:
	;
	v1459 = F_sdsnew(m, int32(_a68))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L19
	} else {
		goto L335
	}
L323:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1425)))
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+int32(-1)))))
	switch v1430 & int32(7) {
	case 0:
		goto L329
	case 1:
		goto L328
	case 2:
		goto L327
	case 3:
		goto L326
	case 4:
		goto L325
	default:
		v1455 = int32(0)
		goto L324
	}
L324:
	;
	F_addReplyBulkCBuffer(m, l0, v1427, v1455)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L19
	} else {
		goto L334
	}
L325:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1427+int32(-17))))
	v1455 = v1454
	goto L324
L326:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1427+int32(-9))))
	F_addReplyBulkCBuffer(m, l0, v1427, v1449)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L19
	} else {
		goto L333
	}
L327:
	;
	v1444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1427+int32(-5)))))
	F_addReplyBulkCBuffer(m, l0, v1427, v1444)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L19
	} else {
		goto L332
	}
L328:
	;
	v1439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1427+int32(-3)))))
	F_addReplyBulkCBuffer(m, l0, v1427, v1439)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L19
	} else {
		goto L331
	}
L329:
	;
	F_addReplyBulkCBuffer(m, l0, v1427, int32(base.Ui32(v1430)>>(uint(int32(3))%32)))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L19
	} else {
		goto L330
	}
L330:
	;
	goto L321
L331:
	;
	goto L321
L332:
	;
	goto L321
L333:
	;
	goto L321
L334:
	;
	goto L321
L335:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1425)))
	v1462 = F_sdscatsds(m, v1459, v1461)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L19
	} else {
		goto L336
	}
L336:
	;
	v1466 = F_sdscatlen(m, v1462, int32(_a6), int32(1))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L19
	} else {
		goto L337
	}
L337:
	;
	v1468 = F_ACLDescribeUser(m, v1425)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L19
	} else {
		goto L338
	}
L338:
	;
	v1470 = F_objectGetVal(m, v1468)
	mBase = m.M
	v1471 = F_sdscatsds(m, v1466, v1470)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L19
	} else {
		goto L339
	}
L339:
	;
	F_decrRefCount(m, v1468)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L19
	} else {
		goto L340
	}
L340:
	;
	F_addReplyBulkSds(m, l0, v1471)
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L19
	} else {
		goto L341
	}
L341:
	;
	goto L321
L342:
	;
	if v1482 != 0 {
		goto L319
	} else {
		goto L343
	}
L343:
	;
	goto L320
L344:
	;
	goto L1
L345:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _consts[19]))
	v1586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585))))
	if v1586 != 0 {
		goto L374
	} else {
		goto L375
	}
L346:
	;
	if v1540-v1542 != 0 {
		goto L345
	} else {
		goto L358
	}
L347:
	;
	v1540 = F_tolower(m, v1536)
	mBase = m.M
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1537))))
	v1542 = F_tolower(m, v1541)
	mBase = m.M
	goto L346
L348:
	;
	v1510 = v24
	v1511 = v1505
	v1512 = v1508
	goto L351
L349:
	;
	v1536 = int32(0)
	v1537 = v1505
	goto L347
L350:
	;
	v1536 = v1533 & int32(255)
	v1537 = v1532
	goto L347
L351:
	;
	v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511))))
	if v1514 == int32(0) {
		v1532 = v1511
		v1533 = v1512
		goto L350
	} else {
		goto L353
	}
L352:
	;
	v1532 = v1526
	v1533 = int32(0)
	goto L350
L353:
	;
	v1518 = v1512 & int32(255)
	if v1518 == v1514 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1525 = int32(1)
	v1526 = v1511 + v1525
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+1)))
	if v1527 != 0 {
		v1510 = v1510 + v1525
		v1511 = v1526
		v1512 = v1527
		goto L351
	} else {
		goto L357
	}
L355:
	;
	v1520 = F_tolower(m, v1518)
	mBase = m.M
	v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1511))))
	v1522 = F_tolower(m, v1521)
	mBase = m.M
	if v1520 == v1522 {
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510))))
	v1532 = v1511
	v1533 = v1524
	goto L350
L357:
	;
	goto L352
L358:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1544 != int32(2) {
		goto L345
	} else {
		goto L359
	}
L359:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v1547 == int32(0) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	F_addReplyNull(m, l0)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L19
	} else {
		goto L373
	}
L361:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1547)))
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1551+int32(-1)))))
	switch v1554 & int32(7) {
	case 0:
		goto L367
	case 1:
		goto L366
	case 2:
		goto L365
	case 3:
		goto L364
	case 4:
		goto L363
	default:
		v1579 = int32(0)
		goto L362
	}
L362:
	;
	F_addReplyBulkCBuffer(m, l0, v1551, v1579)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L19
	} else {
		goto L372
	}
L363:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1551+int32(-17))))
	v1579 = v1578
	goto L362
L364:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1551+int32(-9))))
	F_addReplyBulkCBuffer(m, l0, v1551, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L19
	} else {
		goto L371
	}
L365:
	;
	v1568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1551+int32(-5)))))
	F_addReplyBulkCBuffer(m, l0, v1551, v1568)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L19
	} else {
		goto L370
	}
L366:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1551+int32(-3)))))
	F_addReplyBulkCBuffer(m, l0, v1551, v1563)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L19
	} else {
		goto L369
	}
L367:
	;
	F_addReplyBulkCBuffer(m, l0, v1551, int32(base.Ui32(v1554)>>(uint(int32(3))%32)))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L19
	} else {
		goto L368
	}
L368:
	;
	goto L1
L369:
	;
	goto L1
L370:
	;
	goto L1
L371:
	;
	goto L1
L372:
	;
	goto L1
L373:
	;
	goto L1
L374:
	;
	v1670 = int32(_a69)
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1673 != 0 {
		goto L407
	} else {
		goto L408
	}
L375:
	;
	v1587 = int32(_a69)
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1590 != 0 {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	F_addReplyError(m, l0, int32(_a70))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L19
	} else {
		goto L403
	}
L377:
	;
	if v1622-v1624 == int32(0) {
		goto L376
	} else {
		goto L389
	}
L378:
	;
	v1622 = F_tolower(m, v1618)
	mBase = m.M
	v1623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619))))
	v1624 = F_tolower(m, v1623)
	mBase = m.M
	goto L377
L379:
	;
	v1592 = v24
	v1593 = v1587
	v1594 = v1590
	goto L382
L380:
	;
	v1618 = int32(0)
	v1619 = v1587
	goto L378
L381:
	;
	v1618 = v1615 & int32(255)
	v1619 = v1614
	goto L378
L382:
	;
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593))))
	if v1596 == int32(0) {
		v1614 = v1593
		v1615 = v1594
		goto L381
	} else {
		goto L384
	}
L383:
	;
	v1614 = v1608
	v1615 = int32(0)
	goto L381
L384:
	;
	v1600 = v1594 & int32(255)
	if v1600 == v1596 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1607 = int32(1)
	v1608 = v1593 + v1607
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592)+1)))
	if v1609 != 0 {
		v1592 = v1592 + v1607
		v1593 = v1608
		v1594 = v1609
		goto L382
	} else {
		goto L388
	}
L386:
	;
	v1602 = F_tolower(m, v1600)
	mBase = m.M
	v1603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593))))
	v1604 = F_tolower(m, v1603)
	mBase = m.M
	if v1602 == v1604 {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592))))
	v1614 = v1593
	v1615 = v1606
	goto L381
L388:
	;
	goto L383
L389:
	;
	v1628 = int32(_a71)
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1631 != 0 {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	if v1663-v1665 != 0 {
		goto L374
	} else {
		goto L402
	}
L391:
	;
	v1663 = F_tolower(m, v1659)
	mBase = m.M
	v1664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1660))))
	v1665 = F_tolower(m, v1664)
	mBase = m.M
	goto L390
L392:
	;
	v1633 = v24
	v1634 = v1628
	v1635 = v1631
	goto L395
L393:
	;
	v1659 = int32(0)
	v1660 = v1628
	goto L391
L394:
	;
	v1659 = v1656 & int32(255)
	v1660 = v1655
	goto L391
L395:
	;
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634))))
	if v1637 == int32(0) {
		v1655 = v1634
		v1656 = v1635
		goto L394
	} else {
		goto L397
	}
L396:
	;
	v1655 = v1649
	v1656 = int32(0)
	goto L394
L397:
	;
	v1641 = v1635 & int32(255)
	if v1641 == v1637 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1648 = int32(1)
	v1649 = v1634 + v1648
	v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633)+1)))
	if v1650 != 0 {
		v1633 = v1633 + v1648
		v1634 = v1649
		v1635 = v1650
		goto L395
	} else {
		goto L401
	}
L399:
	;
	v1643 = F_tolower(m, v1641)
	mBase = m.M
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634))))
	v1645 = F_tolower(m, v1644)
	mBase = m.M
	if v1643 == v1645 {
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633))))
	v1655 = v1634
	v1656 = v1647
	goto L394
L401:
	;
	goto L396
L402:
	;
	goto L376
L403:
	;
	goto L1
L404:
	;
	v1722 = int32(_a71)
	v1725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v1725 != 0 {
		goto L428
	} else {
		goto L429
	}
L405:
	;
	if v1705-v1707 != 0 {
		goto L404
	} else {
		goto L417
	}
L406:
	;
	v1705 = F_tolower(m, v1701)
	mBase = m.M
	v1706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1702))))
	v1707 = F_tolower(m, v1706)
	mBase = m.M
	goto L405
L407:
	;
	v1675 = v24
	v1676 = v1670
	v1677 = v1673
	goto L410
L408:
	;
	v1701 = int32(0)
	v1702 = v1670
	goto L406
L409:
	;
	v1701 = v1698 & int32(255)
	v1702 = v1697
	goto L406
L410:
	;
	v1679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1676))))
	if v1679 == int32(0) {
		v1697 = v1676
		v1698 = v1677
		goto L409
	} else {
		goto L412
	}
L411:
	;
	v1697 = v1691
	v1698 = int32(0)
	goto L409
L412:
	;
	v1683 = v1677 & int32(255)
	if v1683 == v1679 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1690 = int32(1)
	v1691 = v1676 + v1690
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675)+1)))
	if v1692 != 0 {
		v1675 = v1675 + v1690
		v1676 = v1691
		v1677 = v1692
		goto L410
	} else {
		goto L416
	}
L414:
	;
	v1685 = F_tolower(m, v1683)
	mBase = m.M
	v1686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1676))))
	v1687 = F_tolower(m, v1686)
	mBase = m.M
	if v1685 == v1687 {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675))))
	v1697 = v1676
	v1698 = v1689
	goto L409
L416:
	;
	goto L411
L417:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1709 != int32(2) {
		goto L404
	} else {
		goto L418
	}
L418:
	;
	v1712 = F_ACLLoadFromFile(m, v1585)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L19
	} else {
		goto L420
	}
L419:
	;
	F_addReplyError(m, l0, v1712)
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L19
	} else {
		goto L423
	}
L420:
	;
	if v1712 != 0 {
		goto L419
	} else {
		goto L421
	}
L421:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v1715)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L19
	} else {
		goto L422
	}
L422:
	;
	goto L1
L423:
	;
	F_sdsfree(m, v1712)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L19
	} else {
		goto L424
	}
L424:
	;
	goto L1
L425:
	;
	v2082 = int32(_a72)
	v2085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v2085 != 0 {
		goto L531
	} else {
		goto L532
	}
L426:
	;
	if v1757-v1759 != 0 {
		goto L425
	} else {
		goto L438
	}
L427:
	;
	v1757 = F_tolower(m, v1753)
	mBase = m.M
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754))))
	v1759 = F_tolower(m, v1758)
	mBase = m.M
	goto L426
L428:
	;
	v1727 = v24
	v1728 = v1722
	v1729 = v1725
	goto L431
L429:
	;
	v1753 = int32(0)
	v1754 = v1722
	goto L427
L430:
	;
	v1753 = v1750 & int32(255)
	v1754 = v1749
	goto L427
L431:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
	if v1731 == int32(0) {
		v1749 = v1728
		v1750 = v1729
		goto L430
	} else {
		goto L433
	}
L432:
	;
	v1749 = v1743
	v1750 = int32(0)
	goto L430
L433:
	;
	v1735 = v1729 & int32(255)
	if v1735 == v1731 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1742 = int32(1)
	v1743 = v1728 + v1742
	v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1727)+1)))
	if v1744 != 0 {
		v1727 = v1727 + v1742
		v1728 = v1743
		v1729 = v1744
		goto L431
	} else {
		goto L437
	}
L435:
	;
	v1737 = F_tolower(m, v1735)
	mBase = m.M
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728))))
	v1739 = F_tolower(m, v1738)
	mBase = m.M
	if v1737 == v1739 {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1727))))
	v1749 = v1728
	v1750 = v1741
	goto L430
L437:
	;
	goto L432
L438:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1761 != int32(2) {
		goto L425
	} else {
		goto L439
	}
L439:
	;
	v1764 = F_sdsempty(m)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L19
	} else {
		goto L440
	}
L440:
	;
	v1767 = v20 + int32(176)
	v1769 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1767)+4)) = v1769
	*(*int32)(unsafe.Add(mBase, uint32(v1767))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1767)+20)) = int32(128)
	v1775 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1767)+12)) = v1775
	*(*int64)(unsafe.Add(mBase, uint32(v1767)+296)) = v1775
	*(*int64)(unsafe.Add(mBase, uint32(v1767)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v1767)+8)) = v20 + int32(200)
	*(*int32)(unsafe.Add(mBase, uint32(v1767)+156)) = v20 + int32(344)
	goto L441
L441:
	;
	v1790 = int32(0)
	v1792 = F_raxSeek(m, v20+int32(176), int32(_a67), v1790, v1790)
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L19
	} else {
		goto L442
	}
L442:
	;
	v1796 = F_raxNext(m, v20+int32(176))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L19
	} else {
		goto L444
	}
L443:
	;
	F_raxStop(m, v20+int32(176))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L19
	} else {
		goto L459
	}
L444:
	;
	if v1796 == int32(0) {
		v1849 = v1764
		goto L443
	} else {
		goto L445
	}
L445:
	;
	v1802 = v1764
	goto L446
L446:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v20)+188))
	v1819 = F_sdsnew(m, int32(_a68))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L19
	} else {
		goto L448
	}
L447:
	;
	v1849 = v1839
	goto L443
L448:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1817)))
	v1822 = F_sdscatsds(m, v1819, v1821)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L19
	} else {
		goto L449
	}
L449:
	;
	v1826 = F_sdscatlen(m, v1822, int32(_a6), int32(1))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L19
	} else {
		goto L450
	}
L450:
	;
	v1828 = F_ACLDescribeUser(m, v1817)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L19
	} else {
		goto L451
	}
L451:
	;
	v1830 = F_objectGetVal(m, v1828)
	mBase = m.M
	v1831 = F_sdscatsds(m, v1826, v1830)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L19
	} else {
		goto L452
	}
L452:
	;
	F_decrRefCount(m, v1828)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L19
	} else {
		goto L453
	}
L453:
	;
	v1835 = F_sdscatsds(m, v1802, v1831)
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L19
	} else {
		goto L454
	}
L454:
	;
	v1839 = F_sdscatlen(m, v1835, int32(_a26), int32(1))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L19
	} else {
		goto L455
	}
L455:
	;
	F_sdsfree(m, v1831)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L19
	} else {
		goto L456
	}
L456:
	;
	v1845 = F_raxNext(m, v20+int32(176))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L19
	} else {
		goto L457
	}
L457:
	;
	if v1845 != 0 {
		v1802 = v1839
		goto L446
	} else {
		goto L458
	}
L458:
	;
	goto L447
L459:
	;
	v1868 = F_sdsnew(m, v1585)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L19
	} else {
		goto L460
	}
L460:
	;
	v1870 = F___syscall_getpid(m)
	mBase = m.M
	goto L461
L461:
	;
	v1872 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L462
L462:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v1870
	v1878 = F_sdscatfmt(m, v1868, int32(_a73), v20+int32(80))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L19
	} else {
		goto L463
	}
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(420)
	v1885 = F_open(m, v1878, int32(65), v20+int32(64))
	mBase = m.M
	if v1885 == int32(-1) {
		goto L465
	} else {
		goto L466
	}
L464:
	;
	v1917 = v1898
	v1929 = v1849 + v1898
	goto L473
L465:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v1902 {
		goto L144
	} else {
		goto L467
	}
L466:
	;
	v1889 = v1849 + int32(-3)
	v1891 = v1849 + int32(-5)
	v1893 = v1849 + int32(-9)
	v1895 = v1849 + int32(-17)
	v1898 = int32(0)
	goto L464
L467:
	;
	goto L468
L468:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v1907 = F___strerror_l(m, v1906, v1906)
	mBase = m.M
	goto L469
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v1907
	F__serverLog(m, int32(3), int32(_a74), v20)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L19
	} else {
		goto L470
	}
L470:
	;
	goto L144
L471:
	;
	v1985 = F_close(m, v1885)
	mBase = m.M
	v1986 = F_rename(m, v1878, v1585)
	mBase = m.M
	if v1986 != int32(-1) {
		goto L499
	} else {
		goto L500
	}
L472:
	;
	v1977 = F___strerror_l(m, v1975, v1975)
	mBase = m.M
	goto L497
L473:
	;
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1849+int32(-1)))))
	v1933 = v1931 & int32(7)
	switch v1933 {
	case 0:
		goto L480
	case 1:
		goto L479
	case 2:
		goto L478
	case 3:
		goto L477
	case 4:
		goto L476
	default:
		v1940 = int32(0)
		goto L475
	}
L474:
	;
	v1965 = F_fsync(m, v1885)
	mBase = m.M
	if v1965 != int32(-1) {
		goto L471
	} else {
		goto L494
	}
L475:
	;
	if base.Ui32(v1940) <= base.Ui32(v1917) {
		goto L481
	} else {
		goto L482
	}
L476:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1895)))
	v1940 = v1939
	goto L475
L477:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1893)))
	v1940 = v1938
	goto L475
L478:
	;
	v1937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1891))))
	v1940 = v1937
	goto L475
L479:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1889))))
	v1940 = v1936
	goto L475
L480:
	;
	v1940 = int32(base.Ui32(v1931) >> (uint(int32(3)) % 32))
	goto L475
L481:
	;
	goto L474
L482:
	;
	switch v1933 {
	case 0:
		goto L488
	case 1:
		goto L487
	case 2:
		goto L486
	case 3:
		goto L485
	case 4:
		goto L484
	default:
		v1949 = int32(0)
		goto L483
	}
L483:
	;
	v1951 = F_write(m, v1885, v1929, v1949-v1917)
	mBase = m.M
	if int32(0) < v1951 {
		goto L489
	} else {
		goto L490
	}
L484:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1895)))
	v1949 = v1948
	goto L483
L485:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1893)))
	v1949 = v1947
	goto L483
L486:
	;
	v1946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1891))))
	v1949 = v1946
	goto L483
L487:
	;
	v1945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1889))))
	v1949 = v1945
	goto L483
L488:
	;
	v1949 = int32(base.Ui32(v1931) >> (uint(int32(3)) % 32))
	goto L483
L489:
	;
	v1963 = v1951 + v1917
	v1917 = v1963
	v1929 = v1849 + v1963
	goto L473
L490:
	;
	goto L491
L491:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	if v1955 == int32(27) {
		goto L473
	} else {
		goto L492
	}
L492:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if v1960 <= int32(3) {
		v1975 = v1955
		v1976 = int32(_a75)
		goto L472
	} else {
		goto L493
	}
L493:
	;
	goto L145
L494:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v1969 {
		goto L145
	} else {
		goto L495
	}
L495:
	;
	goto L496
L496:
	;
	v1973 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v1975 = v1973
	v1976 = int32(_a76)
	goto L472
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v1977
	F__serverLog(m, int32(3), v1976, v20+int32(16))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L19
	} else {
		goto L498
	}
L498:
	;
	v1984 = F_close(m, v1885)
	mBase = m.M
	goto L144
L499:
	;
	v2006 = m.G0
	v2008 = v2006 - int32(4112)
	m.G0 = v2008
	v2010 = F_strlen(m, v1585)
	mBase = m.M
	if base.Ui32(v2010) < base.Ui32(int32(4097)) {
		goto L508
	} else {
		goto L509
	}
L500:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v1990 {
		goto L144
	} else {
		goto L501
	}
L501:
	;
	goto L502
L502:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v1995 = F___strerror_l(m, v1994, v1994)
	mBase = m.M
	goto L503
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v1995
	F__serverLog(m, int32(3), int32(_a77), v20+int32(32))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L19
	} else {
		goto L504
	}
L504:
	;
	goto L144
L505:
	;
	F_sdsfree(m, v1878)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L19
	} else {
		goto L524
	}
L506:
	;
	if v2049 != int32(-1) {
		goto L505
	} else {
		goto L519
	}
L507:
	;
	m.G0 = v2008 + int32(4112)
	goto L506
L508:
	;
	v2019 = F___memcpy(m, v2008, v1585, v2010+int32(1))
	mBase = m.M
	v2020 = F_dirname(m, v2019)
	mBase = m.M
	v2021 = int32(0)
	v2023 = F_open(m, v2020, v2021, v2021)
	mBase = m.M
	if v2023 != int32(-1) {
		goto L510
	} else {
		goto L511
	}
L509:
	;
	v2013 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2013))) = int32(37)
	v2049 = int32(-1)
	goto L507
L510:
	;
	v2033 = F_fsync(m, v2023)
	mBase = m.M
	if v2033 != int32(-1) {
		goto L515
	} else {
		goto L516
	}
L511:
	;
	v2028 = F___errno_location(m)
	mBase = m.M
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	if v2029 != int32(31) {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v2032 = int32(-1)
	goto L514
L513:
	;
	v2032 = int32(0)
	goto L514
L514:
	;
	v2049 = v2032
	goto L507
L515:
	;
	v2047 = F_close(m, v2023)
	mBase = m.M
	v2049 = int32(0)
	goto L507
L516:
	;
	v2036 = F___errno_location(m)
	mBase = m.M
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v2036)))
	if v2037 == int32(8) {
		goto L515
	} else {
		goto L517
	}
L517:
	;
	if v2037 == int32(28) {
		goto L515
	} else {
		goto L518
	}
L518:
	;
	v2042 = F_close(m, v2023)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v2036))) = v2037
	v2049 = int32(-1)
	goto L507
L519:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(3) < v2058 {
		goto L144
	} else {
		goto L520
	}
L520:
	;
	goto L521
L521:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v2063 = F___strerror_l(m, v2062, v2062)
	mBase = m.M
	goto L522
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v2063
	F__serverLog(m, int32(3), int32(_a78), v20+int32(48))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L19
	} else {
		goto L523
	}
L523:
	;
	goto L144
L524:
	;
	F_sdsfree(m, int32(0))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L19
	} else {
		goto L525
	}
L525:
	;
	F_sdsfree(m, v1849)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L19
	} else {
		goto L526
	}
L526:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v2079)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L19
	} else {
		goto L527
	}
L527:
	;
	goto L1
L528:
	;
	v2325 = int32(_a79)
	v2328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v2328 != 0 {
		goto L589
	} else {
		goto L590
	}
L529:
	;
	if v2117-v2119 != 0 {
		goto L528
	} else {
		goto L541
	}
L530:
	;
	v2117 = F_tolower(m, v2113)
	mBase = m.M
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	v2119 = F_tolower(m, v2118)
	mBase = m.M
	goto L529
L531:
	;
	v2087 = v24
	v2088 = v2082
	v2089 = v2085
	goto L534
L532:
	;
	v2113 = int32(0)
	v2114 = v2082
	goto L530
L533:
	;
	v2113 = v2110 & int32(255)
	v2114 = v2109
	goto L530
L534:
	;
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2088))))
	if v2091 == int32(0) {
		v2109 = v2088
		v2110 = v2089
		goto L533
	} else {
		goto L536
	}
L535:
	;
	v2109 = v2103
	v2110 = int32(0)
	goto L533
L536:
	;
	v2095 = v2089 & int32(255)
	if v2095 == v2091 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2102 = int32(1)
	v2103 = v2088 + v2102
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2087)+1)))
	if v2104 != 0 {
		v2087 = v2087 + v2102
		v2088 = v2103
		v2089 = v2104
		goto L534
	} else {
		goto L540
	}
L538:
	;
	v2097 = F_tolower(m, v2095)
	mBase = m.M
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2088))))
	v2099 = F_tolower(m, v2098)
	mBase = m.M
	if v2097 == v2099 {
		goto L537
	} else {
		goto L539
	}
L539:
	;
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2087))))
	v2109 = v2088
	v2110 = v2101
	goto L533
L540:
	;
	goto L535
L541:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v2121 + int32(-2) {
	case 0:
		goto L543
	case 1:
		goto L542
	default:
		goto L528
	}
L542:
	;
	v2181 = int32(0)
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2182)+8))
	v2184 = F_objectGetVal(m, v2183)
	mBase = m.M
	v2186 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v2187 = *(*int64)(unsafe.Add(mBase, uint32(v2186)+8))
	if v2187 == int64(0) {
		goto L146
	} else {
		goto L552
	}
L543:
	;
	v2124 = int32(0)
	v2125 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L19
	} else {
		goto L544
	}
L544:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v2129 = *(*int64)(unsafe.Add(mBase, uint32(v2128)+8))
	if v2129 == int64(0) {
		v2165 = v2124
		goto L545
	} else {
		goto L546
	}
L545:
	;
	F_setDeferredArrayLen(m, l0, v2125, v2165)
	mBase = m.M
	v2180 = m.ExcPending
	if v2180 != 0 {
		goto L19
	} else {
		goto L551
	}
L546:
	;
	v2134 = v2128
	v2135 = v2124
	goto L547
L547:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2134)))
	F_addReplyBulkCString(m, l0, v2149)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L19
	} else {
		goto L549
	}
L548:
	;
	v2165 = v2155
	goto L545
L549:
	;
	v2153 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v2155 = v2135 + int32(1)
	v2158 = v2153 + v2155<<(uint(int32(4))%32)
	v2159 = *(*int64)(unsafe.Add(mBase, uint32(v2158)+8))
	if v2159 != int64(0) {
		v2134 = v2158
		v2135 = v2155
		goto L547
	} else {
		goto L550
	}
L550:
	;
	goto L548
L551:
	;
	goto L1
L552:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v2186)))
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184))))
	if v2193 != 0 {
		goto L556
	} else {
		goto L557
	}
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = int32(0)
	v2314 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L19
	} else {
		goto L583
	}
L554:
	;
	if v2225-v2227 == int32(0) {
		v2303 = v2187
		goto L553
	} else {
		goto L566
	}
L555:
	;
	v2225 = F_tolower(m, v2221)
	mBase = m.M
	v2226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2222))))
	v2227 = F_tolower(m, v2226)
	mBase = m.M
	goto L554
L556:
	;
	v2195 = v2184
	v2196 = v2190
	v2197 = v2193
	goto L559
L557:
	;
	v2221 = int32(0)
	v2222 = v2190
	goto L555
L558:
	;
	v2221 = v2218 & int32(255)
	v2222 = v2217
	goto L555
L559:
	;
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	if v2199 == int32(0) {
		v2217 = v2196
		v2218 = v2197
		goto L558
	} else {
		goto L561
	}
L560:
	;
	v2217 = v2211
	v2218 = int32(0)
	goto L558
L561:
	;
	v2203 = v2197 & int32(255)
	if v2203 == v2199 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v2210 = int32(1)
	v2211 = v2196 + v2210
	v2212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2195)+1)))
	if v2212 != 0 {
		v2195 = v2195 + v2210
		v2196 = v2211
		v2197 = v2212
		goto L559
	} else {
		goto L565
	}
L563:
	;
	v2205 = F_tolower(m, v2203)
	mBase = m.M
	v2206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	v2207 = F_tolower(m, v2206)
	mBase = m.M
	if v2205 == v2207 {
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2195))))
	v2217 = v2196
	v2218 = v2209
	goto L558
L565:
	;
	goto L560
L566:
	;
	v2234 = v2181
	goto L567
L567:
	;
	v2249 = v2234 + int32(1)
	v2252 = v2186 + v2249<<(uint(int32(4))%32)
	v2253 = *(*int64)(unsafe.Add(mBase, uint32(v2252)+8))
	if v2253 == int64(0) {
		goto L146
	} else {
		goto L569
	}
L568:
	;
	v2303 = v2253
	goto L553
L569:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2252)))
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184))))
	if v2259 != 0 {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	if v2291-v2293 != 0 {
		v2234 = v2249
		goto L567
	} else {
		goto L582
	}
L571:
	;
	v2291 = F_tolower(m, v2287)
	mBase = m.M
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2288))))
	v2293 = F_tolower(m, v2292)
	mBase = m.M
	goto L570
L572:
	;
	v2261 = v2184
	v2262 = v2256
	v2263 = v2259
	goto L575
L573:
	;
	v2287 = int32(0)
	v2288 = v2256
	goto L571
L574:
	;
	v2287 = v2284 & int32(255)
	v2288 = v2283
	goto L571
L575:
	;
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2262))))
	if v2265 == int32(0) {
		v2283 = v2262
		v2284 = v2263
		goto L574
	} else {
		goto L577
	}
L576:
	;
	v2283 = v2277
	v2284 = int32(0)
	goto L574
L577:
	;
	v2269 = v2263 & int32(255)
	if v2269 == v2265 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v2276 = int32(1)
	v2277 = v2262 + v2276
	v2278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2261)+1)))
	if v2278 != 0 {
		v2261 = v2261 + v2276
		v2262 = v2277
		v2263 = v2278
		goto L575
	} else {
		goto L581
	}
L579:
	;
	v2271 = F_tolower(m, v2269)
	mBase = m.M
	v2272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2262))))
	v2273 = F_tolower(m, v2272)
	mBase = m.M
	if v2271 == v2273 {
		goto L578
	} else {
		goto L580
	}
L580:
	;
	v2275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2261))))
	v2283 = v2262
	v2284 = v2275
	goto L574
L581:
	;
	goto L576
L582:
	;
	goto L568
L583:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_aclCatWithFlags(m, l0, v2317, v2303, v20+int32(176))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L19
	} else {
		goto L584
	}
L584:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	F_setDeferredArrayLen(m, l0, v2314, v2322)
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L19
	} else {
		goto L585
	}
L585:
	;
	goto L1
L586:
	;
	v2407 = int32(_a80)
	v2410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v2410 != 0 {
		goto L612
	} else {
		goto L613
	}
L587:
	;
	if v2360-v2362 != 0 {
		goto L586
	} else {
		goto L599
	}
L588:
	;
	v2360 = F_tolower(m, v2356)
	mBase = m.M
	v2361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357))))
	v2362 = F_tolower(m, v2361)
	mBase = m.M
	goto L587
L589:
	;
	v2330 = v24
	v2331 = v2325
	v2332 = v2328
	goto L592
L590:
	;
	v2356 = int32(0)
	v2357 = v2325
	goto L588
L591:
	;
	v2356 = v2353 & int32(255)
	v2357 = v2352
	goto L588
L592:
	;
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2331))))
	if v2334 == int32(0) {
		v2352 = v2331
		v2353 = v2332
		goto L591
	} else {
		goto L594
	}
L593:
	;
	v2352 = v2346
	v2353 = int32(0)
	goto L591
L594:
	;
	v2338 = v2332 & int32(255)
	if v2338 == v2334 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v2345 = int32(1)
	v2346 = v2331 + v2345
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2330)+1)))
	if v2347 != 0 {
		v2330 = v2330 + v2345
		v2331 = v2346
		v2332 = v2347
		goto L592
	} else {
		goto L598
	}
L596:
	;
	v2340 = F_tolower(m, v2338)
	mBase = m.M
	v2341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2331))))
	v2342 = F_tolower(m, v2341)
	mBase = m.M
	if v2340 == v2342 {
		goto L595
	} else {
		goto L597
	}
L597:
	;
	v2344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2330))))
	v2352 = v2331
	v2353 = v2344
	goto L591
L598:
	;
	goto L593
L599:
	;
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2364&int32(-2) != int32(2) {
		goto L586
	} else {
		goto L600
	}
L600:
	;
	v2369 = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v2369
	if v2364 != int32(3) {
		v2393 = v2369
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v2399 = int32(base.Ui32(v2393+int32(3)) >> (uint(int32(2)) % 32))
	F_getRandomHexChars(m, v20+int32(176), v2399)
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L19
	} else {
		goto L607
	}
L602:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2374)+8))
	v2379 = F_getLongFromObjectOrReply(m, l0, v2375, v20+int32(172), int32(0))
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L19
	} else {
		goto L603
	}
L603:
	;
	if v2379 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v20)+172))
	v2382 = int32(-4097)
	if base.Ui32(v2382) < base.Ui32(v2381+v2382) {
		v2393 = v2381
		goto L601
	} else {
		goto L605
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(4096)
	F_addReplyErrorFormat(m, l0, int32(_a81), v20+int32(96))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L19
	} else {
		goto L606
	}
L606:
	;
	goto L1
L607:
	;
	F_addReplyBulkCBuffer(m, l0, v20+int32(176), v2399)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L19
	} else {
		goto L608
	}
L608:
	;
	goto L1
L609:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2742 = int32(_a82)
	v2745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v2745 != 0 {
		goto L705
	} else {
		goto L706
	}
L610:
	;
	if v2442-v2444 != 0 {
		goto L609
	} else {
		goto L622
	}
L611:
	;
	v2442 = F_tolower(m, v2438)
	mBase = m.M
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439))))
	v2444 = F_tolower(m, v2443)
	mBase = m.M
	goto L610
L612:
	;
	v2412 = v24
	v2413 = v2407
	v2414 = v2410
	goto L615
L613:
	;
	v2438 = int32(0)
	v2439 = v2407
	goto L611
L614:
	;
	v2438 = v2435 & int32(255)
	v2439 = v2434
	goto L611
L615:
	;
	v2416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2413))))
	if v2416 == int32(0) {
		v2434 = v2413
		v2435 = v2414
		goto L614
	} else {
		goto L617
	}
L616:
	;
	v2434 = v2428
	v2435 = int32(0)
	goto L614
L617:
	;
	v2420 = v2414 & int32(255)
	if v2420 == v2416 {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v2427 = int32(1)
	v2428 = v2413 + v2427
	v2429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2412)+1)))
	if v2429 != 0 {
		v2412 = v2412 + v2427
		v2413 = v2428
		v2414 = v2429
		goto L615
	} else {
		goto L621
	}
L619:
	;
	v2422 = F_tolower(m, v2420)
	mBase = m.M
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2413))))
	v2424 = F_tolower(m, v2423)
	mBase = m.M
	if v2422 == v2424 {
		goto L618
	} else {
		goto L620
	}
L620:
	;
	v2426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2412))))
	v2434 = v2413
	v2435 = v2426
	goto L614
L621:
	;
	goto L616
L622:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2446&int32(-2) != int32(2) {
		goto L609
	} else {
		goto L623
	}
L623:
	;
	v2451 = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v2451
	if v2446 != int32(3) {
		v2523 = v2451
		goto L626
	} else {
		goto L627
	}
L624:
	;
	F_addReplyArrayLen(m, l0, v2532)
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L19
	} else {
		goto L648
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v2530
	v2532 = v2530
	goto L624
L626:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v2526)+20))
	if base.Ui32(v2523) <= base.Ui32(v2527) {
		v2532 = v2523
		goto L624
	} else {
		goto L647
	}
L627:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2456)+8))
	v2458 = F_objectGetVal(m, v2457)
	mBase = m.M
	v2459 = int32(_a83)
	v2462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2458))))
	if v2462 != 0 {
		goto L631
	} else {
		goto L632
	}
L628:
	;
	v2512 = int32(0)
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2513)+8))
	v2518 = F_getLongFromObjectOrReply(m, l0, v2514, v20+int32(172), v2512)
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L19
	} else {
		goto L644
	}
L629:
	;
	if v2494-v2496 != 0 {
		goto L628
	} else {
		goto L641
	}
L630:
	;
	v2494 = F_tolower(m, v2490)
	mBase = m.M
	v2495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2491))))
	v2496 = F_tolower(m, v2495)
	mBase = m.M
	goto L629
L631:
	;
	v2464 = v2458
	v2465 = v2459
	v2466 = v2462
	goto L634
L632:
	;
	v2490 = int32(0)
	v2491 = v2459
	goto L630
L633:
	;
	v2490 = v2487 & int32(255)
	v2491 = v2486
	goto L630
L634:
	;
	v2468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2465))))
	if v2468 == int32(0) {
		v2486 = v2465
		v2487 = v2466
		goto L633
	} else {
		goto L636
	}
L635:
	;
	v2486 = v2480
	v2487 = int32(0)
	goto L633
L636:
	;
	v2472 = v2466 & int32(255)
	if v2472 == v2468 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2479 = int32(1)
	v2480 = v2465 + v2479
	v2481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2464)+1)))
	if v2481 != 0 {
		v2464 = v2464 + v2479
		v2465 = v2480
		v2466 = v2481
		goto L634
	} else {
		goto L640
	}
L638:
	;
	v2474 = F_tolower(m, v2472)
	mBase = m.M
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2465))))
	v2476 = F_tolower(m, v2475)
	mBase = m.M
	if v2474 == v2476 {
		goto L637
	} else {
		goto L639
	}
L639:
	;
	v2478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2464))))
	v2486 = v2465
	v2487 = v2478
	goto L633
L640:
	;
	goto L635
L641:
	;
	v2499 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v2499)+12)) = int32(14)
	F_listEmpty(m, v2499)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L19
	} else {
		goto L642
	}
L642:
	;
	v2504 = int32(0)
	v2505 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v2505)+12)) = v2504
	v2509 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v2509)
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L19
	} else {
		goto L643
	}
L643:
	;
	goto L1
L644:
	;
	if v2518 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v20)+172))
	if v2520 < int32(0) {
		v2530 = v2512
		goto L625
	} else {
		goto L646
	}
L646:
	;
	v2523 = v2520
	goto L626
L647:
	;
	v2530 = v2527
	goto L625
L648:
	;
	v2536 = int32(0)
	v2537 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2539 = v20 + int32(176)
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2537)))
	*(*int32)(unsafe.Add(mBase, uint32(v2539)+4)) = v2536
	*(*int32)(unsafe.Add(mBase, uint32(v2539))) = v2540
	goto L649
L649:
	;
	v2545 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L650
L650:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v20)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v2546 + int32(-1)
	if v2546 == int32(0) {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	goto L652
L652:
	;
	v2570 = v20 + int32(176)
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(v2570)))
	if v2572 == int32(0) {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	if v2572 == int32(0) {
		goto L1
	} else {
		goto L657
	}
L655:
	;
	goto L654
L656:
	;
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2570)+4))
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(v2572+base.B2i32(v2575 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2570))) = v2581
	goto L655
L657:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2572)+8))
	F_addReplyMapLen(m, l0, int32(10))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L19
	} else {
		goto L658
	}
L658:
	;
	F_addReplyBulkCString(m, l0, int32(_a84))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L19
	} else {
		goto L659
	}
L659:
	;
	v2592 = *(*int64)(unsafe.Add(mBase, uint32(v2585)))
	F_addReplyLongLong(m, l0, v2592)
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L19
	} else {
		goto L660
	}
L660:
	;
	F_addReplyBulkCString(m, l0, int32(_a85))
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L19
	} else {
		goto L661
	}
L661:
	;
	v2598 = int32(_a86)
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	v2602 = v2600 + int32(-1)
	if base.Ui32(int32(5)) < base.Ui32(v2602) {
		v2610 = v2598
		goto L662
	} else {
		goto L663
	}
L662:
	;
	F_addReplyBulkCString(m, l0, v2610)
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L19
	} else {
		goto L664
	}
L663:
	;
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2602<<(uint(int32(2))%32))+uint32(_consts[30])))
	v2610 = v2609
	goto L662
L664:
	;
	F_addReplyBulkCString(m, l0, int32(_a87))
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L19
	} else {
		goto L665
	}
L665:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+12))
	if base.Ui32(int32(4)) < base.Ui32(v2616) {
		v2624 = v2598
		goto L666
	} else {
		goto L667
	}
L666:
	;
	F_addReplyBulkCString(m, l0, v2624)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L19
	} else {
		goto L668
	}
L667:
	;
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v2616<<(uint(int32(2))%32))+uint32(_consts[31])))
	v2624 = v2623
	goto L666
L668:
	;
	F_addReplyBulkCString(m, l0, int32(_a88))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L19
	} else {
		goto L669
	}
L669:
	;
	v2630 = int32(0)
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+16))
	v2635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2632+int32(-1)))))
	switch v2635 & int32(7) {
	case 0:
		goto L675
	case 1:
		goto L674
	case 2:
		goto L673
	case 3:
		goto L672
	case 4:
		goto L671
	default:
		v2652 = v2630
		goto L670
	}
L670:
	;
	F_addReplyBulkCBuffer(m, l0, v2632, v2652)
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L19
	} else {
		goto L676
	}
L671:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2632+int32(-17))))
	v2652 = v2651
	goto L670
L672:
	;
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v2632+int32(-9))))
	v2652 = v2648
	goto L670
L673:
	;
	v2645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2632+int32(-5)))))
	v2652 = v2645
	goto L670
L674:
	;
	v2642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2632+int32(-3)))))
	v2652 = v2642
	goto L670
L675:
	;
	v2652 = int32(base.Ui32(v2635) >> (uint(int32(3)) % 32))
	goto L670
L676:
	;
	F_addReplyBulkCString(m, l0, int32(_a89))
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L19
	} else {
		goto L677
	}
L677:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+20))
	v2661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2658+int32(-1)))))
	switch v2661 & int32(7) {
	case 0:
		goto L683
	case 1:
		goto L682
	case 2:
		goto L681
	case 3:
		goto L680
	case 4:
		goto L679
	default:
		v2678 = v2630
		goto L678
	}
L678:
	;
	F_addReplyBulkCBuffer(m, l0, v2658, v2678)
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L19
	} else {
		goto L684
	}
L679:
	;
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v2658+int32(-17))))
	v2678 = v2677
	goto L678
L680:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2658+int32(-9))))
	v2678 = v2674
	goto L678
L681:
	;
	v2671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2658+int32(-5)))))
	v2678 = v2671
	goto L678
L682:
	;
	v2668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2658+int32(-3)))))
	v2678 = v2668
	goto L678
L683:
	;
	v2678 = int32(base.Ui32(v2661) >> (uint(int32(3)) % 32))
	goto L678
L684:
	;
	F_addReplyBulkCString(m, l0, int32(_a90))
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L19
	} else {
		goto L685
	}
L685:
	;
	v2684 = *(*int64)(unsafe.Add(mBase, uint32(v2585)+24))
	F_addReplyDouble(m, l0, base.F64_div(base.F64_convert_i64_s(v2545-v2684), float64(1000)))
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L19
	} else {
		goto L686
	}
L686:
	;
	F_addReplyBulkCString(m, l0, int32(_a91))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L19
	} else {
		goto L687
	}
L687:
	;
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+32))
	v2698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2695+int32(-1)))))
	switch v2698 & int32(7) {
	case 0:
		goto L693
	case 1:
		goto L692
	case 2:
		goto L691
	case 3:
		goto L690
	case 4:
		goto L689
	default:
		v2715 = int32(0)
		goto L688
	}
L688:
	;
	F_addReplyBulkCBuffer(m, l0, v2695, v2715)
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L19
	} else {
		goto L694
	}
L689:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v2695+int32(-17))))
	v2715 = v2714
	goto L688
L690:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2695+int32(-9))))
	v2715 = v2711
	goto L688
L691:
	;
	v2708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2695+int32(-5)))))
	v2715 = v2708
	goto L688
L692:
	;
	v2705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2695+int32(-3)))))
	v2715 = v2705
	goto L688
L693:
	;
	v2715 = int32(base.Ui32(v2698) >> (uint(int32(3)) % 32))
	goto L688
L694:
	;
	F_addReplyBulkCString(m, l0, int32(_a92))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L19
	} else {
		goto L695
	}
L695:
	;
	v2721 = *(*int64)(unsafe.Add(mBase, uint32(v2585)+40))
	F_addReplyLongLong(m, l0, v2721)
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		goto L19
	} else {
		goto L696
	}
L696:
	;
	F_addReplyBulkCString(m, l0, int32(_a93))
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L19
	} else {
		goto L697
	}
L697:
	;
	v2727 = *(*int64)(unsafe.Add(mBase, uint32(v2585)+48))
	F_addReplyLongLong(m, l0, v2727)
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L19
	} else {
		goto L698
	}
L698:
	;
	F_addReplyBulkCString(m, l0, int32(_a94))
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		goto L19
	} else {
		goto L699
	}
L699:
	;
	v2733 = *(*int64)(unsafe.Add(mBase, uint32(v2585)+24))
	F_addReplyLongLong(m, l0, v2733)
	mBase = m.M
	v2735 = m.ExcPending
	if v2735 != 0 {
		goto L19
	} else {
		goto L700
	}
L700:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v20)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+172)) = v2736 + int32(-1)
	if v2736 != 0 {
		goto L652
	} else {
		goto L701
	}
L701:
	;
	goto L1
L702:
	;
	if v2741 != int32(2) {
		goto L786
	} else {
		goto L787
	}
L703:
	;
	if v2777-v2779 != 0 {
		goto L702
	} else {
		goto L715
	}
L704:
	;
	v2777 = F_tolower(m, v2773)
	mBase = m.M
	v2778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2774))))
	v2779 = F_tolower(m, v2778)
	mBase = m.M
	goto L703
L705:
	;
	v2747 = v24
	v2748 = v2742
	v2749 = v2745
	goto L708
L706:
	;
	v2773 = int32(0)
	v2774 = v2742
	goto L704
L707:
	;
	v2773 = v2770 & int32(255)
	v2774 = v2769
	goto L704
L708:
	;
	v2751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2748))))
	if v2751 == int32(0) {
		v2769 = v2748
		v2770 = v2749
		goto L707
	} else {
		goto L710
	}
L709:
	;
	v2769 = v2763
	v2770 = int32(0)
	goto L707
L710:
	;
	v2755 = v2749 & int32(255)
	if v2755 == v2751 {
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v2762 = int32(1)
	v2763 = v2748 + v2762
	v2764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747)+1)))
	if v2764 != 0 {
		v2747 = v2747 + v2762
		v2748 = v2763
		v2749 = v2764
		goto L708
	} else {
		goto L714
	}
L712:
	;
	v2757 = F_tolower(m, v2755)
	mBase = m.M
	v2758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2748))))
	v2759 = F_tolower(m, v2758)
	mBase = m.M
	if v2757 == v2759 {
		goto L711
	} else {
		goto L713
	}
L713:
	;
	v2761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747))))
	v2769 = v2748
	v2770 = v2761
	goto L707
L714:
	;
	goto L709
L715:
	;
	if v2741 < int32(4) {
		goto L702
	} else {
		goto L716
	}
L716:
	;
	v2783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(v2783)+8))
	v2785 = F_objectGetVal(m, v2784)
	mBase = m.M
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v2786)+8))
	v2788 = F_objectGetVal(m, v2787)
	mBase = m.M
	v2794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2788+int32(-1)))))
	switch v2794 & int32(7) {
	case 0:
		goto L723
	case 1:
		goto L722
	case 2:
		goto L721
	case 3:
		goto L720
	case 4:
		goto L719
	default:
		v2811 = int32(0)
		goto L718
	}
L717:
	;
	v2814 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v2814
	v2817 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v2819 = v20 + int32(176)
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v2817)))
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2828)))
	if v2813 == v2814 {
		goto L726
	} else {
		goto L727
	}
L718:
	;
	v2813 = v2811
	goto L717
L719:
	;
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2788+int32(-17))))
	v2811 = v2810
	goto L718
L720:
	;
	v2807 = *(*int32)(unsafe.Add(mBase, uint32(v2788+int32(-9))))
	v2813 = v2807
	goto L717
L721:
	;
	v2804 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2788+int32(-5)))))
	v2813 = v2804
	goto L717
L722:
	;
	v2801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2788+int32(-3)))))
	v2813 = v2801
	goto L717
L723:
	;
	v2813 = int32(base.Ui32(v2794) >> (uint(int32(3)) % 32))
	goto L717
L724:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	if v3016 != 0 {
		goto L762
	} else {
		goto L763
	}
L725:
	;
	if v2972 != v2813 {
		goto L752
	} else {
		goto L753
	}
L726:
	;
	v2963 = int32(0)
	v2969 = v2828
	v2970 = v2829
	v2972 = v2963
	v2976 = v2963
	goto L725
L727:
	;
	if base.Ui32(v2829) < base.Ui32(int32(8)) {
		goto L726
	} else {
		goto L728
	}
L728:
	;
	v2840 = v2828
	v2841 = v2829
	v2843 = int32(0)
	goto L730
L729:
	;
	v2969 = v2953
	v2970 = v2954
	v2972 = v2956
	v2976 = base.B2i32(v2959 != int32(0))
	goto L725
L730:
	;
	v2849 = int32(base.Ui32(v2841) >> (uint(int32(3)) % 32))
	v2850 = int32(4)
	v2851 = v2840 + v2850
	if v2841&v2850 == int32(0) {
		goto L733
	} else {
		goto L734
	}
L731:
	;
	v2953 = v2944
	v2954 = v2945
	v2956 = v2929
	v2959 = v2934
	goto L729
L732:
	;
	v2934 = int32(0)
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(v2851+v2849+(v2934-v2849)&int32(3)+v2922<<(uint(int32(2))%32))))
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v2944)))
	if base.Ui32(v2945) < base.Ui32(int32(8)) {
		v2953 = v2944
		v2954 = v2945
		v2956 = v2929
		v2959 = v2934
		goto L729
	} else {
		goto L750
	}
L733:
	;
	v2897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2785+v2843))))
	v2900 = int32(0)
	goto L744
L734:
	;
	v2856 = int32(0)
	if base.Ui32(v2813) <= base.Ui32(v2843) {
		v2889 = v2843
		v2892 = v2856
		goto L735
	} else {
		goto L736
	}
L735:
	;
	if v2892 == v2849 {
		v2922 = v2856
		v2929 = v2889
		goto L732
	} else {
		goto L742
	}
L736:
	;
	v2866 = v2843
	v2869 = v2856
	goto L737
L737:
	;
	v2872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2851+v2869))))
	v2874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2785+v2866))))
	if v2872 != v2874 {
		v2889 = v2866
		v2892 = v2869
		goto L735
	} else {
		goto L739
	}
L738:
	;
	v2889 = v2877
	v2892 = v2879
	goto L735
L739:
	;
	v2876 = int32(1)
	v2877 = v2866 + v2876
	v2879 = v2869 + v2876
	if base.Ui32(v2849) <= base.Ui32(v2879) {
		v2889 = v2877
		v2892 = v2879
		goto L735
	} else {
		goto L740
	}
L740:
	;
	if base.Ui32(v2877) < base.Ui32(v2813) {
		v2866 = v2877
		v2869 = v2879
		goto L737
	} else {
		goto L741
	}
L741:
	;
	goto L738
L742:
	;
	v2953 = v2840
	v2954 = v2841
	v2956 = v2889
	v2959 = v2892
	goto L729
L743:
	;
	if v2900 != v2849 {
		goto L748
	} else {
		goto L749
	}
L744:
	;
	v2913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2851+v2900))))
	if v2913 == v2897&int32(255) {
		goto L743
	} else {
		goto L746
	}
L746:
	;
	v2915 = int32(1)
	v2917 = v2900 + v2915
	if v2917 != v2849 {
		v2900 = v2917
		goto L744
	} else {
		goto L747
	}
L747:
	;
	v2969 = v2840
	v2970 = v2841
	v2972 = v2843
	v2976 = v2915
	goto L725
L748:
	;
	v2922 = v2900
	v2929 = v2843 + int32(1)
	goto L732
L749:
	;
	v2953 = v2840
	v2954 = v2841
	v2956 = v2843
	v2959 = v2849
	goto L729
L750:
	;
	if base.Ui32(v2929) < base.Ui32(v2813) {
		v2840 = v2944
		v2841 = v2945
		v2843 = v2929
		goto L730
	} else {
		goto L751
	}
L751:
	;
	goto L731
L752:
	;
	goto L724
L753:
	;
	if v2970&int32(1) == int32(0) {
		goto L752
	} else {
		goto L754
	}
L754:
	;
	v2984 = v2970 & int32(4)
	if v2976&base.B2i32(v2984 != int32(0)) != 0 {
		goto L752
	} else {
		goto L755
	}
L755:
	;
	if v2819 == int32(0) {
		goto L752
	} else {
		goto L756
	}
L756:
	;
	if v2970&int32(2) != 0 {
		v3010 = int32(0)
		goto L757
	} else {
		goto L758
	}
L757:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2819))) = v3010
	goto L752
L758:
	;
	v2994 = int32(3)
	v2995 = int32(base.Ui32(v2970) >> (uint(v2994) % 32))
	if v2984 != 0 {
		goto L759
	} else {
		goto L760
	}
L759:
	;
	v3005 = int32(4)
	goto L761
L760:
	;
	v3005 = v2995 << (uint(int32(2)) % 32)
	goto L761
L761:
	;
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v2969+v2995+(int32(0)-v2995)&v2994+v3005+int32(4))))
	v3010 = v3009
	goto L757
L762:
	;
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3030 = F_lookupCommand(m, v3015+int32(12), v3027+int32(-3))
	mBase = m.M
	v3031 = m.ExcPending
	if v3031 != 0 {
		goto L19
	} else {
		goto L766
	}
L763:
	;
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v3015)+8))
	v3018 = F_objectGetVal(m, v3017)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v3018
	F_addReplyErrorFormat(m, l0, int32(_a95), v20+int32(112))
	mBase = m.M
	v3024 = m.ExcPending
	if v3024 != 0 {
		goto L19
	} else {
		goto L764
	}
L764:
	;
	goto L1
L765:
	;
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(v3030)+52))
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3046 = v3044 + int32(-3)
	if base.B2i32(int32(0) < v3041)&base.B2i32(v3041 != v3046) != 0 {
		goto L770
	} else {
		goto L771
	}
L766:
	;
	if v3030 != 0 {
		goto L765
	} else {
		goto L767
	}
L767:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3032)+12))
	v3034 = F_objectGetVal(m, v3033)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v3034
	F_addReplyErrorFormat(m, l0, int32(_a96), v20+int32(128))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L19
	} else {
		goto L768
	}
L768:
	;
	goto L1
L769:
	;
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v3066 = v3064 & int32(8)
	if v3066 != 0 {
		goto L775
	} else {
		goto L776
	}
L770:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(v3030)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v3052
	F_addReplyErrorFormat(m, l0, int32(_a97), v20+int32(144))
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L19
	} else {
		goto L773
	}
L771:
	;
	if int32(0)-v3041 <= v3046 {
		goto L769
	} else {
		goto L772
	}
L772:
	;
	goto L770
L773:
	;
	goto L1
L774:
	;
	v3096 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v3096)
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		goto L19
	} else {
		goto L785
	}
L775:
	;
	v3067 = int32(112)
	goto L777
L776:
	;
	v3067 = int32(96)
	goto L777
L777:
	;
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(l0+v3067)))
	if v3066 != 0 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v3072 = int32(52)
	goto L780
L779:
	;
	v3072 = int32(28)
	goto L780
L780:
	;
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3069+v3072)))
	v3077 = F_ACLCheckAllUserCommandPerm(m, v3016, v3030, v3059+int32(12), v3046, v3074, v20+int32(176))
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L19
	} else {
		goto L781
	}
L781:
	;
	if v3077 == int32(0) {
		goto L774
	} else {
		goto L782
	}
L782:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v3081+v3082<<(uint(int32(2))%32)+int32(12))))
	v3089 = F_objectGetVal(m, v3088)
	mBase = m.M
	v3091 = F_getAclErrorMessage(m, v3077, v3016, v3030, v3089, int32(1))
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L19
	} else {
		goto L783
	}
L783:
	;
	F_addReplyBulkSds(m, l0, v3091)
	mBase = m.M
	v3094 = m.ExcPending
	if v3094 != 0 {
		goto L19
	} else {
		goto L784
	}
L784:
	;
	goto L1
L785:
	;
	goto L1
L786:
	;
	F_addReplySubcommandSyntaxError(m, l0)
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L19
	} else {
		goto L805
	}
L787:
	;
	v3101 = int32(_a98)
	v3104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v3104 != 0 {
		goto L790
	} else {
		goto L791
	}
L788:
	;
	if v3136-v3138 != 0 {
		goto L786
	} else {
		goto L800
	}
L789:
	;
	v3136 = F_tolower(m, v3132)
	mBase = m.M
	v3137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3133))))
	v3138 = F_tolower(m, v3137)
	mBase = m.M
	goto L788
L790:
	;
	v3106 = v24
	v3107 = v3101
	v3108 = v3104
	goto L793
L791:
	;
	v3132 = int32(0)
	v3133 = v3101
	goto L789
L792:
	;
	v3132 = v3129 & int32(255)
	v3133 = v3128
	goto L789
L793:
	;
	v3110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3107))))
	if v3110 == int32(0) {
		v3128 = v3107
		v3129 = v3108
		goto L792
	} else {
		goto L795
	}
L794:
	;
	v3128 = v3122
	v3129 = int32(0)
	goto L792
L795:
	;
	v3114 = v3108 & int32(255)
	if v3114 == v3110 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v3121 = int32(1)
	v3122 = v3107 + v3121
	v3123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3106)+1)))
	if v3123 != 0 {
		v3106 = v3106 + v3121
		v3107 = v3122
		v3108 = v3123
		goto L793
	} else {
		goto L799
	}
L797:
	;
	v3116 = F_tolower(m, v3114)
	mBase = m.M
	v3117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3107))))
	v3118 = F_tolower(m, v3117)
	mBase = m.M
	if v3116 == v3118 {
		goto L796
	} else {
		goto L798
	}
L798:
	;
	v3120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3106))))
	v3128 = v3107
	v3129 = v3120
	goto L792
L799:
	;
	goto L794
L800:
	;
	goto L803
L801:
	;
	F_addReplyHelp(m, l0, v20+int32(176))
	mBase = m.M
	v3151 = m.ExcPending
	if v3151 != 0 {
		goto L19
	} else {
		goto L804
	}
L802:
	;
	goto L801
L803:
	;
	v3146 = F__emscripten_memcpy_bulkmem(m, v20+int32(176), int32(_a99), int32(108))
	mBase = m.M
	goto L802
L804:
	;
	goto L1
L805:
	;
	goto L1
L806:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L807:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L808:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L809:
	;
	goto L1
L810:
	;
	F_sdsfree(m, v1878)
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L19
	} else {
		goto L812
	}
L811:
	;
	v3220 = F_unlink(m, v1878)
	mBase = m.M
	goto L810
L812:
	;
	F_sdsfree(m, v1849)
	mBase = m.M
	v3224 = m.ExcPending
	if v3224 != 0 {
		goto L19
	} else {
		goto L813
	}
L813:
	;
	F_addReplyError(m, l0, int32(_a100))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L19
	} else {
		goto L814
	}
L814:
	;
	goto L1
L815:
	;
	goto L1
}
func F_addACLLogEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
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
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	switch l1 + int32(-1) {
	case 0:
		goto L4
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L2
	case 4:
		goto L6
	case 5:
		goto L5
	default:
		goto L3
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v61 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v54 = int32(_a20)
	v56 = *(*int64)(unsafe.Add(mBase, _consts[6]))
	*(*int64)(unsafe.Add(mBase, _consts[6])) = v56 + int64(1)
	goto L1
L3:
	;
	F__serverPanic_1(m, int32(_a4), int32(2854), int32(_a21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = int32(_a20)
	v43 = *(*int64)(unsafe.Add(mBase, _consts[7]))
	*(*int64)(unsafe.Add(mBase, _consts[7])) = v43 + int64(1)
	goto L1
L5:
	;
	v35 = int32(_a20)
	v37 = *(*int64)(unsafe.Add(mBase, _consts[8]))
	*(*int64)(unsafe.Add(mBase, _consts[8])) = v37 + int64(1)
	goto L1
L6:
	;
	v29 = int32(_a20)
	v31 = *(*int64)(unsafe.Add(mBase, _consts[9]))
	*(*int64)(unsafe.Add(mBase, _consts[9])) = v31 + int64(1)
	goto L1
L7:
	;
	v23 = int32(_a20)
	v25 = *(*int64)(unsafe.Add(mBase, _consts[10]))
	*(*int64)(unsafe.Add(mBase, _consts[10])) = v25 + int64(1)
	goto L1
L8:
	;
	v17 = int32(_a20)
	v19 = *(*int64)(unsafe.Add(mBase, _consts[11]))
	*(*int64)(unsafe.Add(mBase, _consts[11])) = v19 + int64(1)
	goto L1
L9:
	;
	return
L10:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	m.G0 = v13 + int32(16)
	return
L12:
	;
	F_trimACLLogEntriesToMaxLen(m)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L116
	}
L13:
	;
	v65 = F_valkey_malloc(m, int32(56))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = int64(1)
	if l4 != 0 {
		v72 = l4
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v73 = F_sdsdup(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = v71
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = v73
	v77 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v77
	v80 = *(*int64)(unsafe.Add(mBase, _consts[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+48)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = v80
	if l5 != 0 {
		v126 = l5
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v130 = F_sdsempty(m)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L36
	}
L20:
	;
	switch l1 + int32(-1) {
	case 0:
		goto L23
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L22
	case 4:
		goto L24
	default:
		goto L21
	}
L21:
	;
	v124 = F_sdsempty(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L35
	}
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v121 = F_objectGetVal(m, v120)
	mBase = m.M
	v122 = F_sdsdup(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L34
	}
L23:
	;
	if l3 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+l3<<(uint(int32(2))%32))))
	v102 = F_objectGetVal(m, v101)
	mBase = m.M
	v103 = F_sdsdup(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L29
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+l3<<(uint(int32(2))%32))))
	v94 = F_objectGetVal(m, v93)
	mBase = m.M
	v95 = F_sdsdup(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+140))
	v87 = F_sdsdup(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v126 = v87
	goto L19
L28:
	;
	v126 = v95
	goto L19
L29:
	;
	v126 = v103
	goto L19
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+140))
	v117 = F_sdsdup(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+l3<<(uint(int32(2))%32))))
	v112 = F_objectGetVal(m, v111)
	mBase = m.M
	v113 = F_sdsdup(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v126 = v113
	goto L19
L33:
	;
	v126 = v117
	goto L19
L34:
	;
	v126 = v122
	goto L19
L35:
	;
	v126 = v124
	goto L19
L36:
	;
	if v129 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v132 = v129
	goto L39
L38:
	;
	v132 = l0
	goto L39
L39:
	;
	v134 = F_catClientInfoString(m, v130, v132, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+32)) = v134
	if l2 != int32(4) {
		v176 = l2
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v176
	v178 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v181 = v13 + int32(8)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v182
	goto L55
L42:
	;
	v141 = F_scriptGetRunningEngineName(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v143 = int32(_a22)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[16])))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v147 == int32(0) {
		v170 = v146
		v171 = v147
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v171-v170&int32(255) != 0 {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	if v147 != v146&int32(255) {
		v170 = v146
		v171 = v147
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v153 = v141
	v154 = v143
	goto L48
L48:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v158 == int32(0) {
		v170 = v157
		v171 = v158
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v170 = v157
	v171 = v158
	goto L45
L50:
	;
	v161 = int32(1)
	if v158 == v157&int32(255) {
		v153 = v153 + v161
		v154 = v154 + v161
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v175 = int32(4)
	goto L54
L53:
	;
	v175 = int32(1)
	goto L54
L54:
	;
	v176 = v175
	goto L41
L55:
	;
	v189 = int32(9)
	goto L57
L56:
	;
	v373 = int32(0)
	v375 = *(*int64)(unsafe.Add(mBase, _consts[13]))
	*(*int64)(unsafe.Add(mBase, _consts[13])) = v375 + int64(1)
	v380 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v381 = F_listAddNodeHead(m, v380, v65)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L9
	} else {
		goto L115
	}
L57:
	;
	v198 = v13 + int32(8)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if v200 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L56
L59:
	;
	if v200 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v200+base.B2i32(v203 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v209
	goto L60
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v214 != v215 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v189 != 0 {
		v189 = v189 + int32(-1)
		goto L57
	} else {
		goto L114
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if v217 != v218 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v213)+24))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	v222 = v220 - v221
	v224 = v222 >> (uint(int64(63)) % 64)
	if base.Ui64(int64(60000)) < base.Ui64(v222^v224-v224) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v231 = int32(0)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-1)))))
	switch v238 & int32(7) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	case 4:
		goto L69
	default:
		v255 = v231
		goto L68
	}
L67:
	;
	if v281 != 0 {
		goto L63
	} else {
		goto L86
	}
L68:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+int32(-1)))))
	switch v258 & int32(7) {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2:
		goto L77
	case 3:
		goto L76
	case 4:
		goto L75
	default:
		v275 = v231
		goto L74
	}
L69:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-17))))
	v255 = v254
	goto L68
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(-9))))
	v255 = v251
	goto L68
L71:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229+int32(-5)))))
	v255 = v248
	goto L68
L72:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+int32(-3)))))
	v255 = v245
	goto L68
L73:
	;
	v255 = int32(base.Ui32(v238) >> (uint(int32(3)) % 32))
	goto L68
L74:
	;
	v276 = base.B2i32(base.Ui32(v255) < base.Ui32(v275))
	if base.Ui32(v255) < base.Ui32(v275) {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v230+int32(-17))))
	v275 = v274
	goto L74
L76:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v230+int32(-9))))
	v275 = v271
	goto L74
L77:
	;
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+int32(-5)))))
	v275 = v268
	goto L74
L78:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+int32(-3)))))
	v275 = v265
	goto L74
L79:
	;
	v275 = int32(base.Ui32(v258) >> (uint(int32(3)) % 32))
	goto L74
L80:
	;
	v277 = v255
	goto L82
L81:
	;
	v277 = v275
	goto L82
L82:
	;
	v278 = F_memcmp(m, v229, v230, v277)
	mBase = m.M
	if v278 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v281 = v278
	goto L85
L84:
	;
	v281 = base.B2i32(base.Ui32(v275) < base.Ui32(v255)) - v276
	goto L85
L85:
	;
	goto L67
L86:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v213)+20))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v284 = int32(0)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+int32(-1)))))
	switch v291 & int32(7) {
	case 0:
		goto L93
	case 1:
		goto L92
	case 2:
		goto L91
	case 3:
		goto L90
	case 4:
		goto L89
	default:
		v308 = v284
		goto L88
	}
L87:
	;
	if v334 != 0 {
		goto L63
	} else {
		goto L106
	}
L88:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+int32(-1)))))
	switch v311 & int32(7) {
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
		v328 = v284
		goto L94
	}
L89:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v282+int32(-17))))
	v308 = v307
	goto L88
L90:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v282+int32(-9))))
	v308 = v304
	goto L88
L91:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282+int32(-5)))))
	v308 = v301
	goto L88
L92:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+int32(-3)))))
	v308 = v298
	goto L88
L93:
	;
	v308 = int32(base.Ui32(v291) >> (uint(int32(3)) % 32))
	goto L88
L94:
	;
	v329 = base.B2i32(base.Ui32(v308) < base.Ui32(v328))
	if base.Ui32(v308) < base.Ui32(v328) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v283+int32(-17))))
	v328 = v327
	goto L94
L96:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v283+int32(-9))))
	v328 = v324
	goto L94
L97:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283+int32(-5)))))
	v328 = v321
	goto L94
L98:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+int32(-3)))))
	v328 = v318
	goto L94
L99:
	;
	v328 = int32(base.Ui32(v311) >> (uint(int32(3)) % 32))
	goto L94
L100:
	;
	v330 = v308
	goto L102
L101:
	;
	v330 = v328
	goto L102
L102:
	;
	v331 = F_memcmp(m, v282, v283, v330)
	mBase = m.M
	if v331 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v334 = v331
	goto L105
L104:
	;
	v334 = base.B2i32(base.Ui32(v328) < base.Ui32(v308)) - v329
	goto L105
L105:
	;
	goto L87
L106:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	F_listDelNode(m, v336, v200)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v341 = F_listAddNodeHead(m, v340, v213)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v213)+32))
	F_sdsfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+32)) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+24)) = v348
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v213)))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v350 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+32)) = int32(0)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	F_sdsfree(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	F_sdsfree(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v65)+32))
	F_sdsfree(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	F_valkey_free(m, v65)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	goto L11
L114:
	;
	goto L58
L115:
	;
	goto L12
L116:
	;
	goto L11
}
func F_trimACLLogEntriesToMaxLen(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if base.Ui32(v5) <= base.Ui32(v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = v4
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_sdsfree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_sdsfree(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_sdsfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_valkey_free(m, v12)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	F_listDelNode(m, v25, v11)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if base.Ui32(v32) < base.Ui32(v30) {
		v9 = v29
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L4
}
