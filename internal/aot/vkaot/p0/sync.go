package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_generateSyncSlotsEstablishCommand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F__serverAssert(m, int32(_a324), int32(_a325), int32(1417))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L83
	}
L2:
	;
	v14 = F_sdsempty(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l0 + int32(112)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l0 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v19<<(uint(int32(1))%32) + int32(8)
	v34 = F_sdscatprintf(m, v14, int32(_a326), v9+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v38 = v9 + int32(40)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v39
	goto L6
L6:
	;
	v44 = v9 + int32(40)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v46 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v9 + int32(48)
	return v227
L8:
	;
	if v46 == int32(0) {
		v227 = v34
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46+base.B2i32(v49 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	goto L9
L11:
	;
	v59 = v46
	v61 = v34
	goto L12
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v65))))
	v67 = int32(0)
	if base.Ui64(v66) < base.Ui64(int64(10)) {
		v126 = v67
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v227 = v209
	goto L7
L14:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(v65)+4)))
	v136 = int32(0)
	if base.Ui64(v135) < base.Ui64(int64(10)) {
		v195 = v136
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v133 = int32(1) + v126
	goto L14
L16:
	;
	v72 = v66
	v73 = v67
	goto L17
L17:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v72) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v126 = v120
	goto L15
L19:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v72) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v133 = int32(2) + v73
	goto L14
L21:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v72) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v133 = int32(3) + v73
	goto L14
L23:
	;
	v120 = v73 + int32(12)
	v124 = base.I64_div_u_s(v72, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v72) {
		v72 = v124
		v73 = v120
		goto L17
	} else {
		goto L45
	}
L24:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v72) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v72) {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v72) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v72) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v72) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v72) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v133 = int32(4) + v73
	goto L14
L31:
	;
	v97 = int32(6)
	goto L33
L32:
	;
	v97 = int32(5)
	goto L33
L33:
	;
	v133 = v97 + v73
	goto L14
L34:
	;
	v103 = int32(8)
	goto L36
L35:
	;
	v103 = int32(7)
	goto L36
L36:
	;
	v133 = v103 + v73
	goto L14
L37:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v72) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v72) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v111 = int32(10)
	goto L41
L40:
	;
	v111 = int32(9)
	goto L41
L41:
	;
	v133 = v111 + v73
	goto L14
L42:
	;
	v117 = int32(12)
	goto L44
L43:
	;
	v117 = int32(11)
	goto L44
L44:
	;
	v133 = v117 + v73
	goto L14
L45:
	;
	goto L18
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v133
	v209 = F_sdscatfmt(m, v61, int32(_a327), v9)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L78
	}
L47:
	;
	v202 = int32(1) + v195
	goto L46
L48:
	;
	v141 = v135
	v142 = v136
	goto L49
L49:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v141) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v195 = v189
	goto L47
L51:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v141) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v202 = int32(2) + v142
	goto L46
L53:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v141) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v202 = int32(3) + v142
	goto L46
L55:
	;
	v189 = v142 + int32(12)
	v193 = base.I64_div_u_s(v141, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v141) {
		v141 = v193
		v142 = v189
		goto L49
	} else {
		goto L77
	}
L56:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v141) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v141) {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v141) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v141) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v141) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v141) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v202 = int32(4) + v142
	goto L46
L63:
	;
	v166 = int32(6)
	goto L65
L64:
	;
	v166 = int32(5)
	goto L65
L65:
	;
	v202 = v166 + v142
	goto L46
L66:
	;
	v172 = int32(8)
	goto L68
L67:
	;
	v172 = int32(7)
	goto L68
L68:
	;
	v202 = v172 + v142
	goto L46
L69:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v141) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v141) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v180 = int32(10)
	goto L73
L72:
	;
	v180 = int32(9)
	goto L73
L73:
	;
	v202 = v180 + v142
	goto L46
L74:
	;
	v186 = int32(12)
	goto L76
L75:
	;
	v186 = int32(11)
	goto L76
L76:
	;
	v202 = v186 + v142
	goto L46
L77:
	;
	goto L50
L78:
	;
	v212 = v9 + int32(40)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v214 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v214 != 0 {
		v59 = v214
		v61 = v209
		goto L12
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v214+base.B2i32(v217 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v223
	goto L80
L82:
	;
	goto L13
L83:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_syncReadLine(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = l2 + int32(-1)
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v134
L2:
	;
	v18 = l1
	v19 = v15
	v22 = int32(0)
	goto L4
L3:
	;
	v134 = int32(0)
	goto L1
L4:
	;
	v36 = F_mstime(m)
	mBase = m.M
	goto L8
L5:
	;
	v134 = v15
	goto L1
L6:
	;
	if v95 == int32(-1) {
		v134 = int32(-1)
		goto L1
	} else {
		goto L22
	}
L7:
	;
	goto L6
L8:
	;
	v40 = v11 + int32(15)
	v41 = int32(1)
	v43 = int32(0)
	v45 = l3
	goto L11
L9:
	;
	v95 = int32(-1)
	goto L7
L10:
	;
	v82 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v80
	goto L9
L11:
	;
	v50 = F_read(m, l0, v40, v41)
	mBase = m.M
	switch v50 + int32(1) {
	case 0:
		goto L15
	case 1:
		v80 = int32(15)
		goto L10
	default:
		goto L14
	}
L12:
	;
	v80 = int32(73)
	goto L10
L13:
	;
	v66 = int64(10)
	if v66 < v45 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v57 = v50 + v43
	v58 = v41 - v50
	if v58 == int32(0) {
		v95 = v57
		goto L7
	} else {
		goto L17
	}
L15:
	;
	v53 = F___errno_location(m)
	mBase = m.M
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 == int32(6) {
		v62 = v40
		v63 = v41
		v64 = v43
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v62 = v40 + v50
	v63 = v58
	v64 = v57
	goto L13
L18:
	;
	v69 = v45
	goto L20
L19:
	;
	v69 = v66
	goto L20
L20:
	;
	v70 = F_aeWait(m, l0, int32(1), v69)
	mBase = m.M
	v71 = F_mstime(m)
	mBase = m.M
	v72 = v71 - v36
	if v72 < l3 {
		v40 = v62
		v41 = v63
		v43 = v64
		v45 = l3 - v72
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v103 != int32(10) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)) = uint8(v118)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v103)
	v121 = int32(1)
	v126 = v19 + int32(-1)
	if v126 != 0 {
		v18 = v18 + v121
		v19 = v126
		v22 = v22 + v121
		goto L4
	} else {
		goto L28
	}
L24:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v106)
	if v22 == v106 {
		v134 = v106
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v112 = v18 + int32(-1)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113 != int32(13) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v134 = v22
	goto L1
L27:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v116)
	goto L26
L28:
	;
	goto L5
}
func F_syncWithPrimaryHandleReceiveAuthReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	if v10 == v2 {
		v73 = v2
		m.G0 = v6 + int32(288)
		return v73
	} else {
		v13 = int32(-1)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[551]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
		v24 = m.T0[v23].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v18*int32(1000)))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 != int32(-1) {
				v43 = int32(_a69)
				v45 = *(*int64)(unsafe.Add(mBase, _consts[37]))
				*(*int64)(unsafe.Add(mBase, _consts[523])) = v45
				v49 = F_sdsnew(m, v6+int32(32))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v49 == int32(0) {
						v73 = v13
						m.G0 = v6 + int32(288)
						return v73
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						if v53 != int32(45) {
							F_sdsfree(m, v49)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v73 = int32(0)
								m.G0 = v6 + int32(288)
								return v73
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(3) < v57 {
								F_sdsfree(m, v49)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v73 = v13
									m.G0 = v6 + int32(288)
									return v73
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v49
								F__serverLog(m, int32(3), int32(_a1040), v6+int32(16))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v49)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v73 = v13
										m.G0 = v6 + int32(288)
										return v73
									}
								}
							}
						}
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v31 {
					v73 = v13
					m.G0 = v6 + int32(288)
					return v73
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
					v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v36
						F__serverLog(m, int32(3), int32(_a1010), v6)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v73 = v13
							m.G0 = v6 + int32(288)
							return v73
						}
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleReceiveIPReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[588]))
	if v10 == v2 {
		v71 = v2
		m.G0 = v6 + int32(288)
		return v71
	} else {
		v13 = int32(-1)
		v18 = *(*int32)(unsafe.Add(mBase, _consts[551]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
		v24 = m.T0[v23].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v18*int32(1000)))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 != int32(-1) {
				v43 = int32(_a69)
				v45 = *(*int64)(unsafe.Add(mBase, _consts[37]))
				*(*int64)(unsafe.Add(mBase, _consts[523])) = v45
				v49 = F_sdsnew(m, v6+int32(32))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v49 == int32(0) {
						v71 = v13
						m.G0 = v6 + int32(288)
						return v71
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						if v53 != int32(45) {
							F_sdsfree(m, v49)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v71 = int32(0)
								m.G0 = v6 + int32(288)
								return v71
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(2) < v57 {
								F_sdsfree(m, v49)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v71 = int32(0)
									m.G0 = v6 + int32(288)
									return v71
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v49
								F__serverLog(m, int32(2), int32(_a1041), v6+int32(16))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_sdsfree(m, v49)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v71 = int32(0)
										m.G0 = v6 + int32(288)
										return v71
									}
								}
							}
						}
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _consts[15]))
				if int32(3) < v31 {
					v71 = v13
					m.G0 = v6 + int32(288)
					return v71
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
					v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v36
						F__serverLog(m, int32(3), int32(_a1010), v6)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v71 = v13
							m.G0 = v6 + int32(288)
							return v71
						}
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleReceiveNodeIDReplyState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(288)
	m.G0 = v6
	v8 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v19 = m.T0[v18].(func(*base.Module, int32, int32, int32, int64) int32)(m, l0, v6+int32(32), int32(256), base.I64_extend_i32_s(v13*int32(1000)))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != int32(-1) {
			v38 = int32(_a69)
			v40 = *(*int64)(unsafe.Add(mBase, _consts[37]))
			*(*int64)(unsafe.Add(mBase, _consts[523])) = v40
			v44 = F_sdsnew(m, v6+int32(32))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v44 == int32(0) {
					v66 = v8
					m.G0 = v6 + int32(288)
					return v66
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
					if v48 != int32(45) {
						F_sdsfree(m, v44)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = int32(0)
							m.G0 = v6 + int32(288)
							return v66
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _consts[15]))
						if int32(2) < v52 {
							F_sdsfree(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = int32(0)
								m.G0 = v6 + int32(288)
								return v66
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
							F__serverLog(m, int32(2), int32(_a1042), v6+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_sdsfree(m, v44)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v66 = int32(0)
									m.G0 = v6 + int32(288)
									return v66
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[15]))
			if int32(3) < v26 {
				v66 = v8
				m.G0 = v6 + int32(288)
				return v66
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
				v31 = m.T0[v30].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v31
					F__serverLog(m, int32(3), int32(_a1010), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v66 = v8
						m.G0 = v6 + int32(288)
						return v66
					}
				}
			}
		}
	}
}
func F_syncWithPrimaryHandleSendHandshakeState(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int64
	_ = v274
	var v279 int64
	_ = v279
	var v285 int64
	_ = v285
	var v300 int64
	_ = v300
	var v303 int64
	_ = v303
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(192)
	return v426
L2:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(3) < v414 {
		goto L90
	} else {
		goto L91
	}
L3:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[587]))
	if v145 != 0 {
		v154 = v145
		goto L33
	} else {
		goto L34
	}
L4:
	;
	v16 = int32(0)
	v20 = v10 + int32(184)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[563]))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
	v25 = *(*int64)(unsafe.Add(mBase, _consts[564]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+176)) = v25
	v30 = v10 + int32(168)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[565]))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	v35 = *(*int64)(unsafe.Add(mBase, _consts[566]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, _consts[567]))
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v13
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-1)))))
	switch v112 & int32(7) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	case 3:
		goto L26
	case 4:
		goto L25
	default:
		v129 = v16
		goto L24
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v38
	if v38&int32(3) == int32(0) {
		v70 = v38
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v41 = int32(4)
	v106 = v10 + int32(176) | v41
	v107 = v10 + int32(160) | v41
	v108 = int32(2)
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v103
	v106 = v20
	v107 = v30
	v108 = int32(3)
	goto L5
L9:
	;
	v103 = v95 - v38
	goto L8
L10:
	;
	v74 = v70
	goto L18
L11:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v38
	goto L14
L13:
	;
	v103 = v38 - v38
	goto L8
L14:
	;
	v63 = v59 + int32(1)
	if v63&int32(3) == int32(0) {
		v70 = v63
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v68 != 0 {
		v59 = v63
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v95 = v63
	goto L9
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v83 = int32(-2139062144)
	if (int32(16843008)-v80|v80)&v83 == v83 {
		v74 = v74 + int32(4)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v89 = v74
	goto L21
L20:
	;
	goto L19
L21:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v93 != 0 {
		v89 = v89 + int32(1)
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v95 = v89
	goto L9
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v129
	v135 = F_sendCommandArgv(m, l0, v108, v10+int32(176), v10+int32(160))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-17))))
	v129 = v128
	goto L24
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(-9))))
	v129 = v125
	goto L24
L27:
	;
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+int32(-5)))))
	v129 = v122
	goto L24
L28:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(-3)))))
	v129 = v119
	goto L24
L29:
	;
	v129 = int32(base.Ui32(v112) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	return int32(0)
L31:
	;
	if v135 != 0 {
		v408 = v135
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	v157 = F_sdsfromlonglong(m, base.I64_extend_i32_s(v154))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L30
	} else {
		goto L41
	}
L34:
	;
	v146 = int32(_a69)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[105]))
	v149 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	if v147 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = v147
	goto L37
L36:
	;
	v150 = v149
	goto L37
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v153 = v150
	goto L40
L39:
	;
	v153 = v149
	goto L40
L40:
	;
	v154 = v153
	goto L33
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(_a1031)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(_a935)
	v168 = F_sendCommand(m, l0, v10+int32(48))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	F_sdsfree(m, v157)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	if v168 != 0 {
		v408 = v168
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[588]))
	if v173 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v189&int32(-2) == int32(2) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a1032)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a935)
	v185 = F_sendCommand(m, l0, v10+int32(32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	if v185 != 0 {
		v408 = v185
		goto L2
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v266 = int32(0)
	v267 = *(*int32)(unsafe.Add(mBase, _consts[589]))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(144)))) = v267
	v274 = *(*int64)(unsafe.Add(mBase, _consts[590]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(136)))) = v274
	v279 = *(*int64)(unsafe.Add(mBase, _consts[591]))
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(128)))) = v279
	v285 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(88)))) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(96)))) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = int64(17179869187)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = int64(17179869192)
	v300 = *(*int64)(unsafe.Add(mBase, _consts[592]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v300
	v303 = *(*int64)(unsafe.Add(mBase, _consts[593]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = int32(6)
	if v261 == v266 {
		goto L79
	} else {
		goto L80
	}
L50:
	;
	v228 = F_moduleAllDatatypesHandleErrors(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L30
	} else {
		goto L66
	}
L51:
	;
	v194 = int32(1)
	if v189 != v194 {
		v261 = v194
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v197 = int64(0)
	v201 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v201 < int32(1) {
		v223 = v197
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v223 != int64(0) {
		v261 = v194
		goto L49
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	v205 = v197
	v206 = int32(0)
	goto L56
L56:
	;
	v207 = F_dbHasNoKeys(m, v206)
	mBase = m.M
	if v207 != 0 {
		v217 = v205
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v223 = v217
	goto L54
L58:
	;
	v219 = v206 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v219 < v221 {
		v205 = v217
		v206 = v219
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[23]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v206<<(uint(int32(2))%32))))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v215 = F_kvstoreSize(m, v214)
	mBase = m.M
	v217 = v215 + v205
	goto L58
L60:
	;
	goto L57
L61:
	;
	goto L50
L62:
	;
	v261 = int32(1)
	goto L49
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+132))
	if v254 != 0 {
		goto L75
	} else {
		goto L76
	}
L64:
	;
	F__serverLog(m, int32(2), v247, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L30
	} else {
		goto L74
	}
L65:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[536]))
	if v237 != int32(2) {
		goto L63
	} else {
		goto L70
	}
L66:
	;
	if v228 != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v231 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v261 = int32(1)
	goto L49
L69:
	;
	v247 = int32(_a1033)
	goto L64
L70:
	;
	v240 = F_moduleAllModulesHandleReplAsyncLoad(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	if v240 != 0 {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if int32(2) < v243 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	v247 = int32(_a1034)
	goto L64
L74:
	;
	v261 = int32(1)
	goto L49
L75:
	;
	v256 = m.T0[v254].(func(*base.Module) int32)(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L30
	} else {
		goto L77
	}
L76:
	;
	v261 = int32(1)
	goto L49
L77:
	;
	v261 = base.B2i32(v256 == int32(0))
	goto L49
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[594]))
	if v337 == int32(0) {
		v360 = v333
		goto L81
	} else {
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = int32(_a1035)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = int32(_a1036)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+84)) = int64(73014444036)
	v333 = int32(7)
	v334 = v10 + int32(140)
	v335 = v10 + int32(92)
	goto L78
L80:
	;
	v333 = int32(5)
	v334 = v10 + int32(132)
	v335 = v10 + int32(84)
	goto L78
L81:
	;
	v366 = F_sendCommandArgv(m, l0, v360, v10+int32(112), v10+int32(64))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L30
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = int32(_a1035)
	v342 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v342
	v346 = int32(2)
	v349 = v333<<(uint(v346)%32) + v342
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(64)+v349))) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(112)+v349))) = int32(_a1037)
	v360 = v333 + v346
	goto L81
L83:
	;
	if v366 != 0 {
		v408 = v366
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(_a535)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a678)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a935)
	v379 = F_sendCommand(m, l0, v10+int32(16))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L30
	} else {
		goto L85
	}
L85:
	;
	if v379 != 0 {
		v408 = v379
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v382 == int32(0) {
		v426 = v368
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = int32(_a1038)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = int32(_a935)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v391 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+168)) = int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = int64(81604378632)
	v404 = F_sendCommandArgv(m, l0, int32(3), v10+int32(176), v10+int32(160))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L30
	} else {
		goto L88
	}
L88:
	;
	if v404 == int32(0) {
		v426 = v368
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v408 = v404
	goto L2
L90:
	;
	F_sdsfree(m, v408)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L30
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v408
	F__serverLog(m, int32(3), int32(_a1039), v10)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L30
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v426 = int32(-1)
	goto L1
}
