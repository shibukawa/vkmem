package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_createStreamObject(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_streamNew(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v14 = int32(12)
		v17 = F_zmalloc_usable(m, v14, v6+v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v8
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(34359738374)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22&int32(-241) | int32(160)
			m.G0 = v6 + int32(16)
			return v17
		}
	}
}
func F_freeStream(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_raxFreeWithCallback(m, v3, int32(1092))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(0) {
			F_valkey_free(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		} else {
			F_raxFreeWithCallback(m, v7, int32(1093))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_valkey_free(m, l0)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_streamCreateCG(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int64
	_ = v221
	var v223 int32
	_ = v223
	var v227 int64
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != 0 {
		v14 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if l2 == v15 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v9 = F_raxNew(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v9
	v14 = v9
	goto L1
L5:
	;
	return v234
L6:
	;
	if v210 != 0 {
		v234 = v15
		goto L5
	} else {
		goto L44
	}
L7:
	;
	if v169 != l2 {
		v210 = v15
		goto L34
	} else {
		goto L35
	}
L8:
	;
	v160 = int32(0)
	v167 = v26
	v169 = v160
	v173 = v160
	goto L7
L9:
	;
	if base.Ui32(v26) < base.Ui32(int32(8)) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v37 = v25
	v38 = v26
	v40 = int32(0)
	goto L12
L11:
	;
	v167 = v151
	v169 = v153
	v173 = base.B2i32(v156 != int32(0))
	goto L7
L12:
	;
	v46 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	v47 = int32(4)
	v48 = v37 + v47
	if v38&v47 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v151 = v142
	v153 = v126
	v156 = v131
	goto L11
L14:
	;
	v131 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v48+v46+(v131-v46)&int32(3)+v119<<(uint(int32(2))%32))))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui32(v142) < base.Ui32(int32(8)) {
		v151 = v142
		v153 = v126
		v156 = v131
		goto L11
	} else {
		goto L32
	}
L15:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v40))))
	v97 = int32(0)
	goto L26
L16:
	;
	v53 = int32(0)
	if base.Ui32(l2) <= base.Ui32(v40) {
		v86 = v40
		v89 = v53
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v89 == v46 {
		v119 = v53
		v126 = v86
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v63 = v40
	v66 = v53
	goto L19
L19:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v66))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v63))))
	if v69 != v71 {
		v86 = v63
		v89 = v66
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v86 = v74
	v89 = v76
	goto L17
L21:
	;
	v73 = int32(1)
	v74 = v63 + v73
	v76 = v66 + v73
	if base.Ui32(v46) <= base.Ui32(v76) {
		v86 = v74
		v89 = v76
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(v74) < base.Ui32(l2) {
		v63 = v74
		v66 = v76
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v151 = v38
	v153 = v86
	v156 = v89
	goto L11
L25:
	;
	if v97 != v46 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v97))))
	if v110 == v94&int32(255) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v112 = int32(1)
	v114 = v97 + v112
	if v114 != v46 {
		v97 = v114
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v167 = v38
	v169 = v40
	v173 = v112
	goto L7
L30:
	;
	v119 = v97
	v126 = v40 + int32(1)
	goto L14
L31:
	;
	v151 = v38
	v153 = v40
	v156 = v46
	goto L11
L32:
	;
	if base.Ui32(v126) < base.Ui32(l2) {
		v37 = v141
		v38 = v142
		v40 = v126
		goto L12
	} else {
		goto L33
	}
L33:
	;
	goto L13
L34:
	;
	goto L6
L35:
	;
	v175 = int32(0)
	if v167&int32(1) == v175 {
		v210 = v175
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v173&base.B2i32(v167&int32(4) != int32(0)) != 0 {
		v210 = v175
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v210 = int32(1)
	goto L34
L44:
	;
	v213 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v215 = F_raxNew(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+24)) = v215
	v218 = F_raxNew(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+28)) = v218
	v221 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v221
	v223 = int32(8)
	v227 = *(*int64)(unsafe.Add(mBase, uint32(l3+v223)))
	*(*int64)(unsafe.Add(mBase, uint32(v213+v223))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(v213)+16)) = l4
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v232 = F_raxInsert(m, v230, l1, l2, v213, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v234 = v213
	goto L5
}
func F_streamCreateNACK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	v5 = F_valkey_malloc(m, int32(24))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, _consts[98]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(1)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = v10
		return v5
	}
}
func F_streamDecrID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v17 int64
	_ = v17
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != int64(0) {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v3 + int64(-1)
		return int32(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if v6 != int64(0) {
			v17 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v17
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v6 + v17
			return int32(0)
		} else {
			v9 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
			*(*int64)(unsafe.Add(mBase, uint32(l0+int32(8)))) = v9
			return int32(-1)
		}
	}
}
func F_streamEstimateDistanceFromFirstEverEntry(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v124 int64
	_ = v124
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if base.B2i32(v10 == int64(0)) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	return int64(0)
L3:
	;
	return v124
L4:
	;
	if base.Ui64(v56) < base.Ui64(v57) {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v49) <= base.Ui64(v52) {
		v56 = v49
		v57 = v52
		goto L4
	} else {
		goto L20
	}
L6:
	;
	v44 = int64(-1)
	if base.Ui64(v42) < base.Ui64(v43) {
		v124 = v44
		goto L3
	} else {
		goto L18
	}
L7:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui64(v40) < base.Ui64(v17) {
		v49 = v17
		goto L5
	} else {
		goto L17
	}
L8:
	;
	if v17 != int64(0) {
		goto L7
	} else {
		goto L14
	}
L9:
	;
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v21) < base.Ui64(v17) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if base.Ui64(v21) <= base.Ui64(v17) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v26) < base.Ui64(v25) {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	return v10
L13:
	;
	return v10
L14:
	;
	v32 = int64(0)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if v33 != v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = v32
	v43 = v38
	goto L6
L16:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = int64(0)
	v57 = v36
	goto L4
L17:
	;
	v42 = v17
	v43 = v40
	goto L6
L18:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(v46) < base.Ui64(v47) {
		v124 = v44
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v49 = v42
	goto L5
L20:
	;
	return int64(-1)
L21:
	;
	v69 = int32(1)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui64(v71) < base.Ui64(v56) {
		v85 = v69
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v60) <= base.Ui64(v61) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if base.Ui64(v60) < base.Ui64(v61) {
		goto L21
	} else {
		goto L25
	}
L24:
	;
	return int64(-1)
L25:
	;
	return v10
L26:
	;
	v86 = int32(0)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui64(v71) < base.Ui64(v87) {
		v107 = v86
		v110 = v69
		goto L33
	} else {
		goto L34
	}
L27:
	;
	if base.Ui64(v56) < base.Ui64(v71) {
		v85 = int32(-1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(v77) < base.Ui64(v76) {
		v85 = int32(1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	if base.Ui64(v76) < base.Ui64(v77) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v82 = int32(-1)
	goto L32
L31:
	;
	v82 = int32(0)
	goto L32
L32:
	;
	v85 = v82
	goto L26
L33:
	;
	v111 = int64(-1)
	if v107 != 0 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if base.Ui64(v71) <= base.Ui64(v87) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v87 != int64(0) {
		v107 = v86
		v110 = v101
		goto L33
	} else {
		goto L43
	}
L36:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(v91) <= base.Ui64(v92) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v101 = int32(-1)
	goto L35
L38:
	;
	if base.Ui64(v91) < base.Ui64(v92) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v101 = int32(1)
	goto L35
L40:
	;
	v98 = int32(-1)
	goto L42
L41:
	;
	v98 = int32(0)
	goto L42
L42:
	;
	v101 = v98
	goto L35
L43:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v107 = base.B2i32(v104 == int64(0))
	v110 = v101
	goto L33
L44:
	;
	if int32(-1) < v85 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if int32(-1) < v110 {
		v124 = v111
		goto L3
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if v85 != 0 {
		v124 = v111
		goto L3
	} else {
		goto L49
	}
L48:
	;
	return v10 - v18
L49:
	;
	v124 = v10 - v18 + int64(1)
	goto L3
}
func F_streamFreeNACK(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_valkey_free(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_streamGenericParseIDOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v40 int32
	_ = v40
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v106 int64
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int64
	_ = v237
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v259 int64
	_ = v259
	var v264 int64
	_ = v264
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v280 int64
	_ = v280
	var v291 int64
	_ = v291
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int64
	_ = v337
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v476 int64
	_ = v476
	var v482 int32
	_ = v482
	var v484 int64
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v498 int64
	_ = v498
	var v503 int64
	_ = v503
	var v507 int32
	_ = v507
	var v509 int64
	_ = v509
	var v511 int32
	_ = v511
	var v519 int64
	_ = v519
	var v530 int64
	_ = v530
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v576 int64
	_ = v576
	var v583 int32
	_ = v583
	var v597 int64
	_ = v597
	var v599 int64
	_ = v599
	var v601 int64
	_ = v601
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v628 int32
	_ = v628
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v16 = F_objectGetVal(m, l1)
	mBase = m.M
	v17 = int32(-1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v17))))
	switch v19&int32(7) + v17 {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	default:
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(144)
	return v617
L2:
	;
	v609 = int32(-1)
	if l0 == int32(0) {
		v617 = v609
		goto L1
	} else {
		goto L131
	}
L3:
	;
	v40 = F_objectGetVal(m, l1)
	mBase = m.M
	v42 = F_objectGetVal(m, l1)
	mBase = m.M
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
L4:
	;
	if base.Ui32(int32(127)) < base.Ui32(v36) {
		goto L2
	} else {
		goto L9
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-17))))
	v36 = v35
	goto L4
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(-9))))
	v36 = v32
	goto L4
L7:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+int32(-5)))))
	v36 = v29
	goto L4
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(-3)))))
	v36 = v26
	goto L4
L9:
	;
	goto L3
L10:
	;
	v66 = v62 + int32(1)
	if v66 == int32(0) {
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
	if l4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v69 = F__emscripten_memcpy_bulkmem(m, v14+int32(16), v40, v66)
	mBase = m.M
	goto L17
L19:
	;
	if l5 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
	if base.B2i32(v73 != int32(45))&base.B2i32(v73 != int32(43)) != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+17)))
	if v79&int32(255) == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+17)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
	if v90 != int32(45) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
	goto L23
L25:
	;
	if v90 != int32(43) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v89&int32(255) != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v95 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v95
	*(*int64)(unsafe.Add(mBase, uint32(l2+int32(8)))) = v95
	v617 = int32(0)
	goto L1
L28:
	;
	v115 = int32(45)
	v116 = F___strchrnul(m, v14+int32(16), v115)
	mBase = m.M
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v118 == v115 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	if v89&int32(255) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v106 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v106
	*(*int64)(unsafe.Add(mBase, uint32(l2+int32(8)))) = v106
	v617 = int32(0)
	goto L1
L31:
	;
	v128 = v14 + int32(16)
	if v128&int32(3) == int32(0) {
		v152 = v128
		goto L39
	} else {
		goto L40
	}
L32:
	;
	if v122 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L33:
	;
	v122 = v116
	goto L35
L34:
	;
	v122 = int32(0)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v125)
	goto L31
L37:
	;
	v187 = v14 + int32(8)
	v195 = m.G0
	v197 = v195 - int32(16)
	m.G0 = v197
	if base.Ui32(v185+int32(-21)) < base.Ui32(int32(-20)) {
		goto L56
	} else {
		goto L57
	}
L38:
	;
	v185 = v177 - v128
	goto L37
L39:
	;
	v156 = v152
	goto L47
L40:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v138 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = v128
	goto L43
L42:
	;
	v185 = v128 - v128
	goto L37
L43:
	;
	v145 = v141 + int32(1)
	if v145&int32(3) == int32(0) {
		v152 = v145
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v150 != 0 {
		v141 = v145
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v177 = v145
	goto L38
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v165 = int32(-2139062144)
	if (int32(16843008)-v162|v162)&v165 == v165 {
		v156 = v156 + int32(4)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v171 = v156
	goto L50
L49:
	;
	goto L48
L50:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v175 != 0 {
		v171 = v171 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v177 = v171
	goto L38
L52:
	;
	goto L51
L53:
	;
	if v344 == int32(0) {
		goto L2
	} else {
		goto L80
	}
L54:
	;
	m.G0 = v197 + int32(16)
	goto L53
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = v337
	v344 = int32(1)
	goto L54
L56:
	;
	v308 = int32(0)
	v309 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v308
	v317 = F_strtoull(m, v128, v197+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if v319 == int32(28) {
		v344 = v308
		goto L54
	} else {
		goto L77
	}
L57:
	;
	v203 = int32(1)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v185 != v203 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v204&int32(255) != int32(45) {
		v224 = v203
		v225 = v204
		v226 = v128
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v208 = v204 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v208&int32(255)) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v337 = base.I64_extend_i32_u(v208) & int64(255)
	goto L55
L61:
	;
	if base.Ui32(int32(8)) < base.Ui32((v225+int32(-49))&int32(255)) {
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v224 = int32(2)
	v225 = v222
	v226 = v14 + int32(17)
	goto L61
L63:
	;
	v237 = base.I64_extend_i32_u(v225+int32(-48)) & int64(255)
	if base.Ui32(v185) <= base.Ui32(v224) {
		v280 = v237
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v204&int32(255) != int32(45) {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	v243 = v224
	v245 = v237
	v247 = v226
	goto L66
L66:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	if base.Ui32((v249+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L56
	} else {
		goto L68
	}
L67:
	;
	v280 = v270
	goto L64
L68:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v245) {
		goto L56
	} else {
		goto L69
	}
L69:
	;
	v259 = v245 * int64(10)
	v264 = base.I64_extend_i32_u(v249+int32(-48)) & int64(255)
	if base.Ui64(v264^int64(-1)) < base.Ui64(v259) {
		goto L56
	} else {
		goto L70
	}
L70:
	;
	v268 = int32(1)
	v270 = v259 + v264
	v272 = v243 + v268
	if v272 != v185 {
		v243 = v272
		v245 = v270
		v247 = v247 + v268
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L67
L72:
	;
	if int64(0) <= v280 {
		v337 = v280
		goto L55
	} else {
		goto L76
	}
L73:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v280) {
		goto L56
	} else {
		goto L74
	}
L74:
	;
	v291 = int64(-1)
	if v291 < v280+v291 {
		v344 = int32(0)
		goto L54
	} else {
		goto L75
	}
L75:
	;
	v337 = int64(0)
	goto L55
L76:
	;
	goto L56
L77:
	;
	if v319 == int32(68) {
		v344 = v308
		goto L54
	} else {
		goto L78
	}
L78:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v324 == int32(0) {
		v344 = v308
		goto L54
	} else {
		goto L79
	}
L79:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v344 = base.B2i32(v328 == int32(0))
	goto L54
L80:
	;
	if v122 == int32(0) {
		v599 = l3
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v599
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v601
	v617 = int32(0)
	goto L1
L82:
	;
	v361 = v122 + int32(1)
	if v361&int32(3) == int32(0) {
		v383 = v361
		goto L85
	} else {
		goto L86
	}
L83:
	;
	if l5 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	v416 = v408 - v361
	goto L83
L85:
	;
	v387 = v383
	goto L93
L86:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v369 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v372 = v361
	goto L89
L88:
	;
	v416 = v361 - v361
	goto L83
L89:
	;
	v376 = v372 + int32(1)
	if v376&int32(3) == int32(0) {
		v383 = v376
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v381 != 0 {
		v372 = v376
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v408 = v376
	goto L84
L93:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	v396 = int32(-2139062144)
	if (int32(16843008)-v393|v393)&v396 == v396 {
		v387 = v387 + int32(4)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v402 = v387
	goto L96
L95:
	;
	goto L94
L96:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	if v406 != 0 {
		v402 = v402 + int32(1)
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v408 = v402
	goto L84
L98:
	;
	goto L97
L99:
	;
	v434 = m.G0
	v436 = v434 - int32(16)
	m.G0 = v436
	if base.Ui32(v416+int32(-21)) < base.Ui32(int32(-20)) {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	if v416 != int32(1) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v421 != int32(42) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	v599 = int64(0)
	goto L81
L103:
	;
	if v583 == int32(0) {
		goto L2
	} else {
		goto L130
	}
L104:
	;
	m.G0 = v436 + int32(16)
	goto L103
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v576
	v583 = int32(1)
	goto L104
L106:
	;
	v547 = int32(0)
	v548 = F___errno_location(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = v547
	v556 = F_strtoull(m, v361, v436+int32(12), int32(10))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	if v558 == int32(28) {
		v583 = v547
		goto L104
	} else {
		goto L127
	}
L107:
	;
	v442 = int32(1)
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v416 != v442 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if v443&int32(255) != int32(45) {
		v463 = v442
		v464 = v443
		v465 = v361
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v447 = v443 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v447&int32(255)) {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v576 = base.I64_extend_i32_u(v447) & int64(255)
	goto L105
L111:
	;
	if base.Ui32(int32(8)) < base.Ui32((v464+int32(-49))&int32(255)) {
		goto L106
	} else {
		goto L113
	}
L112:
	;
	v628 = int32(2)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	v463 = v628
	v464 = v461
	v465 = v122 + v628
	goto L111
L113:
	;
	v476 = base.I64_extend_i32_u(v464+int32(-48)) & int64(255)
	if base.Ui32(v416) <= base.Ui32(v463) {
		v519 = v476
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if v443&int32(255) != int32(45) {
		goto L122
	} else {
		goto L123
	}
L115:
	;
	v482 = v463
	v484 = v476
	v486 = v465
	goto L116
L116:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)))
	if base.Ui32((v488+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		goto L106
	} else {
		goto L118
	}
L117:
	;
	v519 = v509
	goto L114
L118:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v484) {
		goto L106
	} else {
		goto L119
	}
L119:
	;
	v498 = v484 * int64(10)
	v503 = base.I64_extend_i32_u(v488+int32(-48)) & int64(255)
	if base.Ui64(v503^int64(-1)) < base.Ui64(v498) {
		goto L106
	} else {
		goto L120
	}
L120:
	;
	v507 = int32(1)
	v509 = v498 + v503
	v511 = v482 + v507
	if v511 != v416 {
		v482 = v511
		v484 = v509
		v486 = v486 + v507
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	if int64(0) <= v519 {
		v576 = v519
		goto L105
	} else {
		goto L126
	}
L123:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v519) {
		goto L106
	} else {
		goto L124
	}
L124:
	;
	v530 = int64(-1)
	if v530 < v519+v530 {
		v583 = int32(0)
		goto L104
	} else {
		goto L125
	}
L125:
	;
	v576 = int64(0)
	goto L105
L126:
	;
	goto L106
L127:
	;
	if v558 == int32(68) {
		v583 = v547
		goto L104
	} else {
		goto L128
	}
L128:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v563 == int32(0) {
		v583 = v547
		goto L104
	} else {
		goto L129
	}
L129:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v436)+12))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	v583 = base.B2i32(v567 == int32(0))
	goto L104
L130:
	;
	v597 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v599 = v597
	goto L81
L131:
	;
	F_addReplyError(m, l0, int32(_a1697))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	return int32(0)
L133:
	;
	v617 = v609
	goto L1
}
func F_streamIteratorGetID(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int64
	_ = v200
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v222 int64
	_ = v222
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v243 int64
	_ = v243
	var v267 int64
	_ = v267
	var v286 int32
	_ = v286
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int64
	_ = v348
	var v353 int32
	_ = v353
	var v359 int64
	_ = v359
	var v364 int32
	_ = v364
	var v370 int64
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v382 int64
	_ = v382
	var v387 int32
	_ = v387
	var v399 int64
	_ = v399
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v485 int64
	_ = v485
	var v491 int32
	_ = v491
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v507 int64
	_ = v507
	var v512 int64
	_ = v512
	var v516 int32
	_ = v516
	var v518 int64
	_ = v518
	var v520 int32
	_ = v520
	var v528 int64
	_ = v528
	var v552 int64
	_ = v552
	var v571 int32
	_ = v571
	var v580 int64
	_ = v580
	var v581 int64
	_ = v581
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
	var v592 int64
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int64
	_ = v603
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v642 int32
	_ = v642
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int64
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v724 int64
	_ = v724
	var v730 int32
	_ = v730
	var v732 int64
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v746 int64
	_ = v746
	var v751 int64
	_ = v751
	var v755 int32
	_ = v755
	var v757 int64
	_ = v757
	var v759 int32
	_ = v759
	var v767 int64
	_ = v767
	var v791 int64
	_ = v791
	var v810 int32
	_ = v810
	var v819 int64
	_ = v819
	var v820 int64
	_ = v820
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v831 int64
	_ = v831
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int64
	_ = v842
	var v853 int32
	_ = v853
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int64
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v916 int64
	_ = v916
	var v922 int32
	_ = v922
	var v924 int64
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v938 int64
	_ = v938
	var v943 int64
	_ = v943
	var v947 int32
	_ = v947
	var v949 int64
	_ = v949
	var v951 int32
	_ = v951
	var v959 int64
	_ = v959
	var v983 int64
	_ = v983
	var v1002 int32
	_ = v1002
	var v1011 int64
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1022 int64
	_ = v1022
	var v1024 int64
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int64
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1082 int64
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1090 int64
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1104 int64
	_ = v1104
	var v1109 int64
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1115 int64
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1125 int64
	_ = v1125
	var v1149 int64
	_ = v1149
	var v1168 int32
	_ = v1168
	var v1177 int64
	_ = v1177
	var v1178 int64
	_ = v1178
	var v1179 int64
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int64
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1242 int64
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1250 int64
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1264 int64
	_ = v1264
	var v1269 int64
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1275 int64
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1285 int64
	_ = v1285
	var v1309 int64
	_ = v1309
	var v1328 int32
	_ = v1328
	var v1337 int64
	_ = v1337
	var v1338 int64
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int64
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int64
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int64
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1408 int64
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1416 int64
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1430 int64
	_ = v1430
	var v1435 int64
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1441 int64
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1451 int64
	_ = v1451
	var v1475 int64
	_ = v1475
	var v1494 int32
	_ = v1494
	var v1503 int64
	_ = v1503
	var v1504 int64
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int64
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int64
	_ = v1513
	var v1516 int64
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int64
	_ = v1518
	var v1521 int64
	_ = v1521
	var v1522 int64
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1531 int64
	_ = v1531
	var v1534 int64
	_ = v1534
	var v1535 int64
	_ = v1535
	var v1537 int64
	_ = v1537
	var v1540 int64
	_ = v1540
	var v1541 int64
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1549 int64
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int64
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1565 int64
	_ = v1565
	var v1577 int32
	_ = v1577
	var v1578 int64
	_ = v1578
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int64
	_ = v1589
	var v1593 int64
	_ = v1593
	var v1600 int32
	_ = v1600
	var v1601 int64
	_ = v1601
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int64
	_ = v1612
	var v1621 int32
	_ = v1621
	var v1627 int32
	_ = v1627
	var v1633 int32
	_ = v1633
	var v1639 int32
	_ = v1639
	var v1647 int32
	_ = v1647
	var v1653 int32
	_ = v1653
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1680 int32
	_ = v1680
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v21 = l0 + int32(88)
	goto L8
L1:
	;
	m.G0 = v16 + int32(16)
	return v1680
L2:
	;
	v1680 = int32(0)
	goto L1
L3:
	;
	F__serverAssert(m, int32(_a1691), int32(_a1692), int32(1181))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L24
	} else {
		goto L361
	}
L4:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L24
	} else {
		goto L360
	}
L5:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L24
	} else {
		goto L359
	}
L6:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L24
	} else {
		goto L358
	}
L7:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L24
	} else {
		goto L357
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	if v35 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+392)) = int64(0)
	goto L8
L11:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L24
	} else {
		goto L356
	}
L12:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L24
	} else {
		goto L355
	}
L13:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1692), int32(282))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L24
	} else {
		goto L354
	}
L14:
	;
	F__serverAssert(m, int32(_a1694), int32(_a1692), int32(1103))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L24
	} else {
		goto L353
	}
L15:
	;
	v655 = v642
	goto L134
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v628
	v642 = v628
	goto L15
L17:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v427 == int32(0) {
		v642 = v38
		goto L15
	} else {
		goto L94
	}
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v40 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v38 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v54 != int32(16) {
		goto L14
	} else {
		goto L30
	}
L22:
	;
	v50 = F_raxPrev(m, v21)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L24
	} else {
		goto L28
	}
L23:
	;
	v41 = F_raxNext(m, v21)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	if v41 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v47 == int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	if v50 == int32(0) {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+392)) = v60
	v62 = int64(56)
	v64 = int64(65280)
	v66 = int64(40)
	v69 = int64(16711680)
	v71 = int64(24)
	v73 = int64(4278190080)
	v75 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v59<<(uint(v62)%64) | v59&v64<<(uint(v66)%64) | (v59&v69<<(uint(v71)%64) | v59&v73<<(uint(v75)%64)) | (int64(base.Ui64(v59)>>(uint(v75)%64))&v73 | int64(base.Ui64(v59)>>(uint(v71)%64))&v69 | (int64(base.Ui64(v59)>>(uint(v66)%64))&v64 | int64(base.Ui64(v59)>>(uint(v62)%64))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v58<<(uint(v62)%64) | v58&v64<<(uint(v66)%64) | (v58&v69<<(uint(v71)%64) | v58&v73<<(uint(v75)%64)) | (int64(base.Ui64(v58)>>(uint(v75)%64))&v73 | int64(base.Ui64(v58)>>(uint(v71)%64))&v69 | (int64(base.Ui64(v58)>>(uint(v66)%64))&v64 | int64(base.Ui64(v58)>>(uint(v62)%64))))
	v134 = F_lpFirst(m, v60)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v138 = F_lpNext(m, v137, v134)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v142 = F_lpNext(m, v141, v138)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v142
	v148 = F_lpGet(m, v142, v16+int32(8), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L24
	} else {
		goto L36
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v300 = F_lpNext(m, v298, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L24
	} else {
		goto L66
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v152 = int32(0)
	if base.Ui32(v151+int32(-21)) < base.Ui32(int32(-20)) {
		v286 = v152
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v148 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v296 = v150
	goto L34
L38:
	;
	if v286 == int32(0) {
		goto L13
	} else {
		goto L65
	}
L39:
	;
	goto L38
L40:
	;
	v164 = int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v151 != v164 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v286 = int32(1)
	goto L39
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v267
	goto L41
L43:
	;
	if v165&int32(255) == int32(45) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v169 = v165 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v169&int32(255)) {
		v286 = v152
		goto L39
	} else {
		goto L45
	}
L45:
	;
	if v16 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v267 = base.I64_extend_i32_u(v169) & int64(255)
	goto L42
L47:
	;
	if base.Ui32(int32(8)) < base.Ui32((v188+int32(-49))&int32(255)) {
		v286 = v152
		goto L39
	} else {
		goto L50
	}
L48:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v187 = int32(2)
	v188 = v185
	v189 = v148 + int32(1)
	goto L47
L49:
	;
	v187 = v164
	v188 = v165
	v189 = v148
	goto L47
L50:
	;
	v200 = base.I64_extend_i32_u(v188+int32(-48)) & int64(255)
	if base.Ui32(v151) <= base.Ui32(v187) {
		v243 = v200
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v165&int32(255) != int32(45) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	v206 = v187
	v208 = v200
	v210 = v189
	goto L53
L53:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	if base.Ui32((v212+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v286 = v152
		goto L39
	} else {
		goto L55
	}
L54:
	;
	v243 = v233
	goto L51
L55:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v208) {
		v286 = v152
		goto L39
	} else {
		goto L56
	}
L56:
	;
	v222 = v208 * int64(10)
	v227 = base.I64_extend_i32_u(v212+int32(-48)) & int64(255)
	if base.Ui64(v227^int64(-1)) < base.Ui64(v222) {
		v286 = v152
		goto L39
	} else {
		goto L57
	}
L57:
	;
	v231 = int32(1)
	v233 = v222 + v227
	v235 = v206 + v231
	if v235 != v151 {
		v206 = v235
		v208 = v233
		v210 = v210 + v231
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	if v243 < int64(0) {
		v286 = v152
		goto L39
	} else {
		goto L63
	}
L60:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v243) {
		v286 = v152
		goto L39
	} else {
		goto L61
	}
L61:
	;
	if v16 == int32(0) {
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v267 = int64(0) - v243
	goto L42
L63:
	;
	if v16 == int32(0) {
		goto L41
	} else {
		goto L64
	}
L64:
	;
	v267 = v243
	goto L42
L65:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v296 = v295
	goto L34
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v300
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v304 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v336 == int32(7) {
		v423 = int32(0)
		goto L75
	} else {
		goto L76
	}
L68:
	;
	v305 = int64(0)
	v306 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v306 == v305 {
		v642 = v300
		goto L15
	} else {
		goto L69
	}
L69:
	;
	v315 = v300
	v316 = v305
	goto L70
L70:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v323 = F_lpNext(m, v322, v315)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L24
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v323
	v327 = v316 + int64(1)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v327) < base.Ui64(v328) {
		v315 = v323
		v316 = v327
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v642 = v323
	goto L15
L74:
	;
	v628 = v423
	goto L16
L75:
	;
	goto L74
L76:
	;
	v339 = int32(-1)
	v340 = v330 + v336
	v345 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340+int32(-2)))))
	v348 = base.I64_extend_i32_u(v345 & int32(127))
	if v339 < v345 {
		v416 = v339
		v418 = v348
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v423 = v340 + v339 + (v416 - base.I32_wrap_i64(v418))
	goto L75
L78:
	;
	v353 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340+int32(-3)))))
	v359 = base.I64_extend_i32_u(v353&int32(127))<<(uint(int64(7))%64) | v348
	if int32(-1) < v353 {
		v399 = v359
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v416 = v373
	v418 = int64(-1)
	goto L77
L80:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v399) {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v364 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340+int32(-4)))))
	v370 = base.I64_extend_i32_u(v364&int32(127))<<(uint(int64(14))%64) | v359
	if int32(-1) < v364 {
		v399 = v370
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v373 = int32(-5)
	v376 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340+v373))))
	v382 = base.I64_extend_i32_u(v376&int32(127))<<(uint(int64(21))%64) | v370
	if int32(-1) < v376 {
		v399 = v382
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v387 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340+int32(-6)))))
	if v387 < int32(0) {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v399 = base.I64_extend_i32_u(v387&int32(127))<<(uint(int64(28))%64) | v382
	goto L80
L85:
	;
	if base.Ui64(int64(16384)) <= base.Ui64(v399) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v416 = int32(-1)
	v418 = v399
	goto L77
L87:
	;
	if base.Ui64(int64(2097152)) <= base.Ui64(v399) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v416 = int32(-2)
	v418 = v399
	goto L77
L89:
	;
	if base.Ui64(v399) < base.Ui64(int64(268435456)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v416 = int32(-3)
	v418 = v399
	goto L77
L91:
	;
	v413 = int32(-4)
	goto L93
L92:
	;
	v413 = int32(-5)
	goto L93
L93:
	;
	v416 = v413
	v418 = v399
	goto L77
L94:
	;
	v433 = F_lpGet(m, v38, v16+int32(8), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L24
	} else {
		goto L97
	}
L95:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	if v581 == int64(0) {
		v612 = v582
		goto L127
	} else {
		goto L128
	}
L96:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v437 = int32(0)
	if base.Ui32(v436+int32(-21)) < base.Ui32(int32(-20)) {
		v571 = v437
		goto L100
	} else {
		goto L101
	}
L97:
	;
	if v433 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v581 = v435
	goto L95
L99:
	;
	if v571 == int32(0) {
		goto L12
	} else {
		goto L126
	}
L100:
	;
	goto L99
L101:
	;
	v449 = int32(1)
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v436 != v449 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v571 = int32(1)
	goto L100
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v552
	goto L102
L104:
	;
	if v450&int32(255) == int32(45) {
		goto L109
	} else {
		goto L110
	}
L105:
	;
	v454 = v450 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v454&int32(255)) {
		v571 = v437
		goto L100
	} else {
		goto L106
	}
L106:
	;
	if v16 == int32(0) {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	v552 = base.I64_extend_i32_u(v454) & int64(255)
	goto L103
L108:
	;
	if base.Ui32(int32(8)) < base.Ui32((v473+int32(-49))&int32(255)) {
		v571 = v437
		goto L100
	} else {
		goto L111
	}
L109:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	v472 = int32(2)
	v473 = v470
	v474 = v433 + int32(1)
	goto L108
L110:
	;
	v472 = v449
	v473 = v450
	v474 = v433
	goto L108
L111:
	;
	v485 = base.I64_extend_i32_u(v473+int32(-48)) & int64(255)
	if base.Ui32(v436) <= base.Ui32(v472) {
		v528 = v485
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if v450&int32(255) != int32(45) {
		goto L120
	} else {
		goto L121
	}
L113:
	;
	v491 = v472
	v493 = v485
	v495 = v474
	goto L114
L114:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+1)))
	if base.Ui32((v497+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v571 = v437
		goto L100
	} else {
		goto L116
	}
L115:
	;
	v528 = v518
	goto L112
L116:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v493) {
		v571 = v437
		goto L100
	} else {
		goto L117
	}
L117:
	;
	v507 = v493 * int64(10)
	v512 = base.I64_extend_i32_u(v497+int32(-48)) & int64(255)
	if base.Ui64(v512^int64(-1)) < base.Ui64(v507) {
		v571 = v437
		goto L100
	} else {
		goto L118
	}
L118:
	;
	v516 = int32(1)
	v518 = v507 + v512
	v520 = v491 + v516
	if v520 != v436 {
		v491 = v520
		v493 = v518
		v495 = v495 + v516
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L115
L120:
	;
	if v528 < int64(0) {
		v571 = v437
		goto L100
	} else {
		goto L124
	}
L121:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v528) {
		v571 = v437
		goto L100
	} else {
		goto L122
	}
L122:
	;
	if v16 == int32(0) {
		goto L102
	} else {
		goto L123
	}
L123:
	;
	v552 = int64(0) - v528
	goto L103
L124:
	;
	if v16 == int32(0) {
		goto L102
	} else {
		goto L125
	}
L125:
	;
	v552 = v528
	goto L103
L126:
	;
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v581 = v580
	goto L95
L127:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v620 = F_lpPrev(m, v619, v612)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L24
	} else {
		goto L133
	}
L128:
	;
	v591 = v582
	v592 = v581
	goto L129
L129:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v599 = F_lpPrev(m, v598, v591)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L24
	} else {
		goto L131
	}
L130:
	;
	v612 = v599
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v599
	v603 = v592 + int64(-1)
	if v603 != int64(0) {
		v591 = v599
		v592 = v603
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v628 = v620
	goto L16
L134:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v662 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+400)) = v853
	v864 = F_lpGet(m, v853, v16+int32(8), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L24
	} else {
		goto L180
	}
L137:
	;
	v672 = F_lpGet(m, v655, v16+int32(8), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L24
	} else {
		goto L143
	}
L138:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v664 = F_lpNext(m, v663, v655)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L24
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v664
	if v664 == int32(0) {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	v853 = v664
	goto L136
L141:
	;
	if v820 == int64(0) {
		goto L10
	} else {
		goto L173
	}
L142:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v676 = int32(0)
	if base.Ui32(v675+int32(-21)) < base.Ui32(int32(-20)) {
		v810 = v676
		goto L146
	} else {
		goto L147
	}
L143:
	;
	if v672 != 0 {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v820 = v674
	goto L141
L145:
	;
	if v810 == int32(0) {
		goto L11
	} else {
		goto L172
	}
L146:
	;
	goto L145
L147:
	;
	v688 = int32(1)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
	if v675 != v688 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v810 = int32(1)
	goto L146
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v791
	goto L148
L150:
	;
	if v689&int32(255) == int32(45) {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v693 = v689 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v693&int32(255)) {
		v810 = v676
		goto L146
	} else {
		goto L152
	}
L152:
	;
	if v16 == int32(0) {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v791 = base.I64_extend_i32_u(v693) & int64(255)
	goto L149
L154:
	;
	if base.Ui32(int32(8)) < base.Ui32((v712+int32(-49))&int32(255)) {
		v810 = v676
		goto L146
	} else {
		goto L157
	}
L155:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)))
	v711 = int32(2)
	v712 = v709
	v713 = v672 + int32(1)
	goto L154
L156:
	;
	v711 = v688
	v712 = v689
	v713 = v672
	goto L154
L157:
	;
	v724 = base.I64_extend_i32_u(v712+int32(-48)) & int64(255)
	if base.Ui32(v675) <= base.Ui32(v711) {
		v767 = v724
		goto L158
	} else {
		goto L159
	}
L158:
	;
	if v689&int32(255) != int32(45) {
		goto L166
	} else {
		goto L167
	}
L159:
	;
	v730 = v711
	v732 = v724
	v734 = v713
	goto L160
L160:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+1)))
	if base.Ui32((v736+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v810 = v676
		goto L146
	} else {
		goto L162
	}
L161:
	;
	v767 = v757
	goto L158
L162:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v732) {
		v810 = v676
		goto L146
	} else {
		goto L163
	}
L163:
	;
	v746 = v732 * int64(10)
	v751 = base.I64_extend_i32_u(v736+int32(-48)) & int64(255)
	if base.Ui64(v751^int64(-1)) < base.Ui64(v746) {
		v810 = v676
		goto L146
	} else {
		goto L164
	}
L164:
	;
	v755 = int32(1)
	v757 = v746 + v751
	v759 = v730 + v755
	if v759 != v675 {
		v730 = v759
		v732 = v757
		v734 = v734 + v755
		goto L160
	} else {
		goto L165
	}
L165:
	;
	goto L161
L166:
	;
	if v767 < int64(0) {
		v810 = v676
		goto L146
	} else {
		goto L170
	}
L167:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v767) {
		v810 = v676
		goto L146
	} else {
		goto L168
	}
L168:
	;
	if v16 == int32(0) {
		goto L148
	} else {
		goto L169
	}
L169:
	;
	v791 = int64(0) - v767
	goto L149
L170:
	;
	if v16 == int32(0) {
		goto L148
	} else {
		goto L171
	}
L171:
	;
	v791 = v767
	goto L149
L172:
	;
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v820 = v819
	goto L141
L173:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v830 = v823
	v831 = v820
	goto L174
L174:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v838 = F_lpPrev(m, v837, v830)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L24
	} else {
		goto L176
	}
L175:
	;
	v853 = v838
	goto L136
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v838
	v842 = v831 + int64(-1)
	if base.B2i32(v842 == int64(0)) == int32(0) {
		v830 = v838
		v831 = v842
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v1015 = F_lpNext(m, v1013, v1014)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L24
	} else {
		goto L210
	}
L179:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v868 = int32(0)
	if base.Ui32(v867+int32(-21)) < base.Ui32(int32(-20)) {
		v1002 = v868
		goto L183
	} else {
		goto L184
	}
L180:
	;
	if v864 != 0 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v866 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v1012 = v866
	goto L178
L182:
	;
	if v1002 == int32(0) {
		goto L7
	} else {
		goto L209
	}
L183:
	;
	goto L182
L184:
	;
	v880 = int32(1)
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	if v867 != v880 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	v1002 = int32(1)
	goto L183
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v983
	goto L185
L187:
	;
	if v881&int32(255) == int32(45) {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v885 = v881 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v885&int32(255)) {
		v1002 = v868
		goto L183
	} else {
		goto L189
	}
L189:
	;
	if v16 == int32(0) {
		goto L185
	} else {
		goto L190
	}
L190:
	;
	v983 = base.I64_extend_i32_u(v885) & int64(255)
	goto L186
L191:
	;
	if base.Ui32(int32(8)) < base.Ui32((v904+int32(-49))&int32(255)) {
		v1002 = v868
		goto L183
	} else {
		goto L194
	}
L192:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+1)))
	v903 = int32(2)
	v904 = v901
	v905 = v864 + int32(1)
	goto L191
L193:
	;
	v903 = v880
	v904 = v881
	v905 = v864
	goto L191
L194:
	;
	v916 = base.I64_extend_i32_u(v904+int32(-48)) & int64(255)
	if base.Ui32(v867) <= base.Ui32(v903) {
		v959 = v916
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if v881&int32(255) != int32(45) {
		goto L203
	} else {
		goto L204
	}
L196:
	;
	v922 = v903
	v924 = v916
	v926 = v905
	goto L197
L197:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926)+1)))
	if base.Ui32((v928+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1002 = v868
		goto L183
	} else {
		goto L199
	}
L198:
	;
	v959 = v949
	goto L195
L199:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v924) {
		v1002 = v868
		goto L183
	} else {
		goto L200
	}
L200:
	;
	v938 = v924 * int64(10)
	v943 = base.I64_extend_i32_u(v928+int32(-48)) & int64(255)
	if base.Ui64(v943^int64(-1)) < base.Ui64(v938) {
		v1002 = v868
		goto L183
	} else {
		goto L201
	}
L201:
	;
	v947 = int32(1)
	v949 = v938 + v943
	v951 = v922 + v947
	if v951 != v867 {
		v922 = v951
		v924 = v949
		v926 = v926 + v947
		goto L197
	} else {
		goto L202
	}
L202:
	;
	goto L198
L203:
	;
	if v959 < int64(0) {
		v1002 = v868
		goto L183
	} else {
		goto L207
	}
L204:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v959) {
		v1002 = v868
		goto L183
	} else {
		goto L205
	}
L205:
	;
	if v16 == int32(0) {
		goto L185
	} else {
		goto L206
	}
L206:
	;
	v983 = int64(0) - v959
	goto L186
L207:
	;
	if v16 == int32(0) {
		goto L185
	} else {
		goto L208
	}
L208:
	;
	v983 = v959
	goto L186
L209:
	;
	v1011 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v1012 = v1011
	goto L178
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1015
	v1018 = int32(8)
	v1022 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(l1+v1018))) = v1022
	v1024 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(8))))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v1024
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v1030 = F_lpGet(m, v1026, v16+v1018, int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L24
	} else {
		goto L213
	}
L211:
	;
	v1179 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v1179 + v1178
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v1184 = F_lpNext(m, v1182, v1183)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L24
	} else {
		goto L243
	}
L212:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v1034 = int32(0)
	if base.Ui32(v1033+int32(-21)) < base.Ui32(int32(-20)) {
		v1168 = v1034
		goto L216
	} else {
		goto L217
	}
L213:
	;
	if v1030 != 0 {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1032 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v1178 = v1032
	goto L211
L215:
	;
	if v1168 == int32(0) {
		goto L6
	} else {
		goto L242
	}
L216:
	;
	goto L215
L217:
	;
	v1046 = int32(1)
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030))))
	if v1033 != v1046 {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	v1168 = int32(1)
	goto L216
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v1149
	goto L218
L220:
	;
	if v1047&int32(255) == int32(45) {
		goto L225
	} else {
		goto L226
	}
L221:
	;
	v1051 = v1047 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1051&int32(255)) {
		v1168 = v1034
		goto L216
	} else {
		goto L222
	}
L222:
	;
	if v16 == int32(0) {
		goto L218
	} else {
		goto L223
	}
L223:
	;
	v1149 = base.I64_extend_i32_u(v1051) & int64(255)
	goto L219
L224:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1070+int32(-49))&int32(255)) {
		v1168 = v1034
		goto L216
	} else {
		goto L227
	}
L225:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030)+1)))
	v1069 = int32(2)
	v1070 = v1067
	v1071 = v1030 + int32(1)
	goto L224
L226:
	;
	v1069 = v1046
	v1070 = v1047
	v1071 = v1030
	goto L224
L227:
	;
	v1082 = base.I64_extend_i32_u(v1070+int32(-48)) & int64(255)
	if base.Ui32(v1033) <= base.Ui32(v1069) {
		v1125 = v1082
		goto L228
	} else {
		goto L229
	}
L228:
	;
	if v1047&int32(255) != int32(45) {
		goto L236
	} else {
		goto L237
	}
L229:
	;
	v1088 = v1069
	v1090 = v1082
	v1092 = v1071
	goto L230
L230:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+1)))
	if base.Ui32((v1094+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1168 = v1034
		goto L216
	} else {
		goto L232
	}
L231:
	;
	v1125 = v1115
	goto L228
L232:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1090) {
		v1168 = v1034
		goto L216
	} else {
		goto L233
	}
L233:
	;
	v1104 = v1090 * int64(10)
	v1109 = base.I64_extend_i32_u(v1094+int32(-48)) & int64(255)
	if base.Ui64(v1109^int64(-1)) < base.Ui64(v1104) {
		v1168 = v1034
		goto L216
	} else {
		goto L234
	}
L234:
	;
	v1113 = int32(1)
	v1115 = v1104 + v1109
	v1117 = v1088 + v1113
	if v1117 != v1033 {
		v1088 = v1117
		v1090 = v1115
		v1092 = v1092 + v1113
		goto L230
	} else {
		goto L235
	}
L235:
	;
	goto L231
L236:
	;
	if v1125 < int64(0) {
		v1168 = v1034
		goto L216
	} else {
		goto L240
	}
L237:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1125) {
		v1168 = v1034
		goto L216
	} else {
		goto L238
	}
L238:
	;
	if v16 == int32(0) {
		goto L218
	} else {
		goto L239
	}
L239:
	;
	v1149 = int64(0) - v1125
	goto L219
L240:
	;
	if v16 == int32(0) {
		goto L218
	} else {
		goto L241
	}
L241:
	;
	v1149 = v1125
	goto L219
L242:
	;
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v1178 = v1177
	goto L211
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1184
	v1190 = F_lpGet(m, v1184, v16+int32(8), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L24
	} else {
		goto L246
	}
L244:
	;
	v1339 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v1339 + v1338
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v1344 = F_lpNext(m, v1342, v1343)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L24
	} else {
		goto L276
	}
L245:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v1194 = int32(0)
	if base.Ui32(v1193+int32(-21)) < base.Ui32(int32(-20)) {
		v1328 = v1194
		goto L249
	} else {
		goto L250
	}
L246:
	;
	if v1190 != 0 {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1192 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v1338 = v1192
	goto L244
L248:
	;
	if v1328 == int32(0) {
		goto L5
	} else {
		goto L275
	}
L249:
	;
	goto L248
L250:
	;
	v1206 = int32(1)
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190))))
	if v1193 != v1206 {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	v1328 = int32(1)
	goto L249
L252:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v1309
	goto L251
L253:
	;
	if v1207&int32(255) == int32(45) {
		goto L258
	} else {
		goto L259
	}
L254:
	;
	v1211 = v1207 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1211&int32(255)) {
		v1328 = v1194
		goto L249
	} else {
		goto L255
	}
L255:
	;
	if v16 == int32(0) {
		goto L251
	} else {
		goto L256
	}
L256:
	;
	v1309 = base.I64_extend_i32_u(v1211) & int64(255)
	goto L252
L257:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1230+int32(-49))&int32(255)) {
		v1328 = v1194
		goto L249
	} else {
		goto L260
	}
L258:
	;
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1190)+1)))
	v1229 = int32(2)
	v1230 = v1227
	v1231 = v1190 + int32(1)
	goto L257
L259:
	;
	v1229 = v1206
	v1230 = v1207
	v1231 = v1190
	goto L257
L260:
	;
	v1242 = base.I64_extend_i32_u(v1230+int32(-48)) & int64(255)
	if base.Ui32(v1193) <= base.Ui32(v1229) {
		v1285 = v1242
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if v1207&int32(255) != int32(45) {
		goto L269
	} else {
		goto L270
	}
L262:
	;
	v1248 = v1229
	v1250 = v1242
	v1252 = v1231
	goto L263
L263:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252)+1)))
	if base.Ui32((v1254+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1328 = v1194
		goto L249
	} else {
		goto L265
	}
L264:
	;
	v1285 = v1275
	goto L261
L265:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1250) {
		v1328 = v1194
		goto L249
	} else {
		goto L266
	}
L266:
	;
	v1264 = v1250 * int64(10)
	v1269 = base.I64_extend_i32_u(v1254+int32(-48)) & int64(255)
	if base.Ui64(v1269^int64(-1)) < base.Ui64(v1264) {
		v1328 = v1194
		goto L249
	} else {
		goto L267
	}
L267:
	;
	v1273 = int32(1)
	v1275 = v1264 + v1269
	v1277 = v1248 + v1273
	if v1277 != v1193 {
		v1248 = v1277
		v1250 = v1275
		v1252 = v1252 + v1273
		goto L263
	} else {
		goto L268
	}
L268:
	;
	goto L264
L269:
	;
	if v1285 < int64(0) {
		v1328 = v1194
		goto L249
	} else {
		goto L273
	}
L270:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1285) {
		v1328 = v1194
		goto L249
	} else {
		goto L271
	}
L271:
	;
	if v16 == int32(0) {
		goto L251
	} else {
		goto L272
	}
L272:
	;
	v1309 = int64(0) - v1285
	goto L252
L273:
	;
	if v16 == int32(0) {
		goto L251
	} else {
		goto L274
	}
L274:
	;
	v1309 = v1285
	goto L252
L275:
	;
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v1338 = v1337
	goto L244
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1344
	v1348 = v1012 & int64(2)
	v1350 = base.B2i32(v1348 == int64(0))
	if v1348 == int64(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	if v1513 <= int64(-1) {
		goto L3
	} else {
		goto L313
	}
L278:
	;
	v1356 = F_lpGet(m, v1344, v16+int32(8), int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L24
	} else {
		goto L282
	}
L279:
	;
	v1351 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v1351
	v1512 = v1344
	v1513 = v1351
	goto L277
L280:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v1504
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+396))
	v1508 = F_lpNext(m, v1506, v1507)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L24
	} else {
		goto L312
	}
L281:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v1360 = int32(0)
	if base.Ui32(v1359+int32(-21)) < base.Ui32(int32(-20)) {
		v1494 = v1360
		goto L285
	} else {
		goto L286
	}
L282:
	;
	if v1356 != 0 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1358 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v1504 = v1358
	goto L280
L284:
	;
	if v1494 == int32(0) {
		goto L4
	} else {
		goto L311
	}
L285:
	;
	goto L284
L286:
	;
	v1372 = int32(1)
	v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1356))))
	if v1359 != v1372 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1494 = int32(1)
	goto L285
L288:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v1475
	goto L287
L289:
	;
	if v1373&int32(255) == int32(45) {
		goto L294
	} else {
		goto L295
	}
L290:
	;
	v1377 = v1373 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1377&int32(255)) {
		v1494 = v1360
		goto L285
	} else {
		goto L291
	}
L291:
	;
	if v16 == int32(0) {
		goto L287
	} else {
		goto L292
	}
L292:
	;
	v1475 = base.I64_extend_i32_u(v1377) & int64(255)
	goto L288
L293:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1396+int32(-49))&int32(255)) {
		v1494 = v1360
		goto L285
	} else {
		goto L296
	}
L294:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1356)+1)))
	v1395 = int32(2)
	v1396 = v1393
	v1397 = v1356 + int32(1)
	goto L293
L295:
	;
	v1395 = v1372
	v1396 = v1373
	v1397 = v1356
	goto L293
L296:
	;
	v1408 = base.I64_extend_i32_u(v1396+int32(-48)) & int64(255)
	if base.Ui32(v1359) <= base.Ui32(v1395) {
		v1451 = v1408
		goto L297
	} else {
		goto L298
	}
L297:
	;
	if v1373&int32(255) != int32(45) {
		goto L305
	} else {
		goto L306
	}
L298:
	;
	v1414 = v1395
	v1416 = v1408
	v1418 = v1397
	goto L299
L299:
	;
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418)+1)))
	if base.Ui32((v1420+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1494 = v1360
		goto L285
	} else {
		goto L301
	}
L300:
	;
	v1451 = v1441
	goto L297
L301:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1416) {
		v1494 = v1360
		goto L285
	} else {
		goto L302
	}
L302:
	;
	v1430 = v1416 * int64(10)
	v1435 = base.I64_extend_i32_u(v1420+int32(-48)) & int64(255)
	if base.Ui64(v1435^int64(-1)) < base.Ui64(v1430) {
		v1494 = v1360
		goto L285
	} else {
		goto L303
	}
L303:
	;
	v1439 = int32(1)
	v1441 = v1430 + v1435
	v1443 = v1414 + v1439
	if v1443 != v1359 {
		v1414 = v1443
		v1416 = v1441
		v1418 = v1418 + v1439
		goto L299
	} else {
		goto L304
	}
L304:
	;
	goto L300
L305:
	;
	if v1451 < int64(0) {
		v1494 = v1360
		goto L285
	} else {
		goto L309
	}
L306:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1451) {
		v1494 = v1360
		goto L285
	} else {
		goto L307
	}
L307:
	;
	if v16 == int32(0) {
		goto L287
	} else {
		goto L308
	}
L308:
	;
	v1475 = int64(0) - v1451
	goto L288
L309:
	;
	if v16 == int32(0) {
		goto L287
	} else {
		goto L310
	}
L310:
	;
	v1475 = v1451
	goto L288
L311:
	;
	v1503 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v1504 = v1503
	goto L280
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1508
	v1511 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v1512 = v1508
	v1513 = v1511
	goto L277
L313:
	;
	v1516 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1517 != 0 {
		goto L317
	} else {
		goto L318
	}
L314:
	;
	if v1348 == int64(0) {
		goto L346
	} else {
		goto L347
	}
L315:
	;
	if v1513 == int64(0) {
		v655 = v1512
		goto L134
	} else {
		goto L341
	}
L316:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)) = uint32(v1012)
	if v1348 == int64(0) {
		goto L339
	} else {
		goto L340
	}
L317:
	;
	v1537 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if base.Ui64(v1537) < base.Ui64(v1516) {
		goto L314
	} else {
		goto L329
	}
L318:
	;
	v1518 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(v1518) < base.Ui64(v1516) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1524 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	if base.Ui64(v1516) < base.Ui64(v1518) {
		goto L315
	} else {
		goto L321
	}
L321:
	;
	v1521 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1522 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui64(v1521) < base.Ui64(v1522) {
		goto L315
	} else {
		goto L322
	}
L322:
	;
	goto L319
L323:
	;
	v1530 = int32(0)
	v1531 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if base.Ui64(v1531) < base.Ui64(v1516) {
		v1680 = v1530
		goto L1
	} else {
		goto L326
	}
L324:
	;
	if base.I32_wrap_i64(v1012&int64(1)) != 0 {
		goto L315
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	if base.Ui64(v1516) < base.Ui64(v1531) {
		goto L316
	} else {
		goto L327
	}
L327:
	;
	v1534 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1535 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	if base.Ui64(v1535) < base.Ui64(v1534) {
		v1680 = v1530
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L316
L329:
	;
	if base.Ui64(v1516) < base.Ui64(v1537) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1543 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	v1540 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1541 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	if base.Ui64(v1541) < base.Ui64(v1540) {
		goto L314
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	v1549 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui64(v1549) < base.Ui64(v1516) {
		goto L316
	} else {
		goto L336
	}
L334:
	;
	if base.I32_wrap_i64(v1012&int64(1)) != 0 {
		goto L314
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	v1551 = int32(0)
	if base.Ui64(v1516) < base.Ui64(v1549) {
		v1680 = v1551
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v1553 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if base.Ui64(v1553) < base.Ui64(v1554) {
		v1680 = v1551
		goto L1
	} else {
		goto L338
	}
L338:
	;
	goto L316
L339:
	;
	v1680 = int32(1)
	goto L1
L340:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1560
	goto L339
L341:
	;
	v1565 = int64(1)
	v1577 = v1512
	v1578 = int64(0)
	goto L342
L342:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1585 = F_lpNext(m, v1584, v1577)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L24
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1585
	v1589 = v1578 + int64(1)
	if v1589 != v1513<<(uint(int64(base.Ui64(v1348)>>(uint(v1565)%64))^v1565)%64) {
		v1577 = v1585
		v1578 = v1589
		goto L342
	} else {
		goto L345
	}
L345:
	;
	v655 = v1585
	goto L134
L346:
	;
	v1593 = int64(5)
	goto L348
L347:
	;
	v1593 = int64(4)
	goto L348
L348:
	;
	v1600 = v1512
	v1601 = v1593
	goto L349
L349:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+392))
	v1608 = F_lpPrev(m, v1607, v1600)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L24
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+396)) = v1608
	v1612 = v1601 + int64(-1)
	if base.B2i32(v1612 == int64(0)) == int32(0) {
		v1600 = v1608
		v1601 = v1612
		goto L349
	} else {
		goto L352
	}
L352:
	;
	v655 = v1608
	goto L134
L353:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L360:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_streamLastValidID(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(464)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(80)))) = v5
	v16 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(96)))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8)+88)) = v16
	v23 = v8 + int32(104)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(2)
	v28 = int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v23)+296)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v23)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v8 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v8 + int32(272)
	v43 = int32(0)
	v45 = F_raxSeek(m, v23, int32(_a1695), v43, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+408)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v8)+60)) = int64(4294967297)
		v56 = F_streamIteratorGetID(m, v8+int32(16), l1, v8+int32(8))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			if v56 != 0 {
				F_raxStop(m, v23)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					m.G0 = v8 + int32(464)
					return
				}
			} else {
				v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				if v58 == int64(0) {
					F_raxStop(m, v23)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						m.G0 = v8 + int32(464)
						return
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v58
					F__serverPanic_1(m, int32(_a1692), int32(1350), int32(_a1696), v8)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_abort(m)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_streamLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_objectGetVal(m, l0)
	mBase = m.M
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	return v3
}
func F_streamLookupCG(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v234
L2:
	;
	v12 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v12
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v17 & int32(7) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		v34 = v12
		goto L4
	}
L3:
	;
	v234 = int32(0)
	goto L1
L4:
	;
	v36 = v8 + int32(12)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v34 = v33
	goto L4
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v34 = v30
	goto L4
L7:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v34 = v27
	goto L4
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v34 = v24
	goto L4
L9:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L4
L10:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v234 = v232
	goto L1
L11:
	;
	if v189 != v34 {
		goto L38
	} else {
		goto L39
	}
L12:
	;
	v180 = int32(0)
	v186 = v45
	v187 = v46
	v189 = v180
	v193 = v180
	goto L11
L13:
	;
	if base.Ui32(v46) < base.Ui32(int32(8)) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v57 = v45
	v58 = v46
	v60 = int32(0)
	goto L16
L15:
	;
	v186 = v170
	v187 = v171
	v189 = v173
	v193 = base.B2i32(v176 != int32(0))
	goto L11
L16:
	;
	v66 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
	v67 = int32(4)
	v68 = v57 + v67
	if v58&v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v170 = v161
	v171 = v162
	v173 = v146
	v176 = v151
	goto L15
L18:
	;
	v151 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v68+v66+(v151-v66)&int32(3)+v139<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if base.Ui32(v162) < base.Ui32(int32(8)) {
		v170 = v161
		v171 = v162
		v173 = v146
		v176 = v151
		goto L15
	} else {
		goto L36
	}
L19:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v60))))
	v117 = int32(0)
	goto L30
L20:
	;
	v73 = int32(0)
	if base.Ui32(v34) <= base.Ui32(v60) {
		v106 = v60
		v109 = v73
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v109 == v66 {
		v139 = v73
		v146 = v106
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v83 = v60
	v86 = v73
	goto L23
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v86))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v83))))
	if v89 != v91 {
		v106 = v83
		v109 = v86
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v106 = v94
	v109 = v96
	goto L21
L25:
	;
	v93 = int32(1)
	v94 = v83 + v93
	v96 = v86 + v93
	if base.Ui32(v66) <= base.Ui32(v96) {
		v106 = v94
		v109 = v96
		goto L21
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v94) < base.Ui32(v34) {
		v83 = v94
		v86 = v96
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v170 = v57
	v171 = v58
	v173 = v106
	v176 = v109
	goto L15
L29:
	;
	if v117 != v66 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v117))))
	if v130 == v114&int32(255) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(1)
	v134 = v117 + v132
	if v134 != v66 {
		v117 = v134
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v186 = v57
	v187 = v58
	v189 = v60
	v193 = v132
	goto L11
L34:
	;
	v139 = v117
	v146 = v60 + int32(1)
	goto L18
L35:
	;
	v170 = v57
	v171 = v58
	v173 = v60
	v176 = v66
	goto L15
L36:
	;
	if base.Ui32(v146) < base.Ui32(v34) {
		v57 = v161
		v58 = v162
		v60 = v146
		goto L16
	} else {
		goto L37
	}
L37:
	;
	goto L17
L38:
	;
	goto L10
L39:
	;
	if v187&int32(1) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v201 = v187 & int32(4)
	if v193&base.B2i32(v201 != int32(0)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	if v36 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	if v187&int32(2) != 0 {
		v227 = int32(0)
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v227
	goto L38
L44:
	;
	v211 = int32(3)
	v212 = int32(base.Ui32(v187) >> (uint(v211) % 32))
	if v201 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v222 = int32(4)
	goto L47
L46:
	;
	v222 = v212 << (uint(int32(2)) % 32)
	goto L47
L47:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v186+v212+(int32(0)-v212)&v211+v222+int32(4))))
	v227 = v226
	goto L43
}
func F_streamLookupConsumer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v234
L2:
	;
	v11 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-1)))))
	switch v17 & int32(7) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		v34 = v11
		goto L4
	}
L3:
	;
	v234 = int32(0)
	goto L1
L4:
	;
	v36 = v8 + int32(12)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-17))))
	v34 = v33
	goto L4
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(-9))))
	v34 = v30
	goto L4
L7:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(-5)))))
	v34 = v27
	goto L4
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(-3)))))
	v34 = v24
	goto L4
L9:
	;
	v34 = int32(base.Ui32(v17) >> (uint(int32(3)) % 32))
	goto L4
L10:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v234 = v232
	goto L1
L11:
	;
	if v189 != v34 {
		goto L38
	} else {
		goto L39
	}
L12:
	;
	v180 = int32(0)
	v186 = v45
	v187 = v46
	v189 = v180
	v193 = v180
	goto L11
L13:
	;
	if base.Ui32(v46) < base.Ui32(int32(8)) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v57 = v45
	v58 = v46
	v60 = int32(0)
	goto L16
L15:
	;
	v186 = v170
	v187 = v171
	v189 = v173
	v193 = base.B2i32(v176 != int32(0))
	goto L11
L16:
	;
	v66 = int32(base.Ui32(v58) >> (uint(int32(3)) % 32))
	v67 = int32(4)
	v68 = v57 + v67
	if v58&v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v170 = v161
	v171 = v162
	v173 = v146
	v176 = v151
	goto L15
L18:
	;
	v151 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v68+v66+(v151-v66)&int32(3)+v139<<(uint(int32(2))%32))))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if base.Ui32(v162) < base.Ui32(int32(8)) {
		v170 = v161
		v171 = v162
		v173 = v146
		v176 = v151
		goto L15
	} else {
		goto L36
	}
L19:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v60))))
	v117 = int32(0)
	goto L30
L20:
	;
	v73 = int32(0)
	if base.Ui32(v34) <= base.Ui32(v60) {
		v106 = v60
		v109 = v73
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v109 == v66 {
		v139 = v73
		v146 = v106
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v83 = v60
	v86 = v73
	goto L23
L23:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v86))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v83))))
	if v89 != v91 {
		v106 = v83
		v109 = v86
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v106 = v94
	v109 = v96
	goto L21
L25:
	;
	v93 = int32(1)
	v94 = v83 + v93
	v96 = v86 + v93
	if base.Ui32(v66) <= base.Ui32(v96) {
		v106 = v94
		v109 = v96
		goto L21
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v94) < base.Ui32(v34) {
		v83 = v94
		v86 = v96
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v170 = v57
	v171 = v58
	v173 = v106
	v176 = v109
	goto L15
L29:
	;
	if v117 != v66 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v117))))
	if v130 == v114&int32(255) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(1)
	v134 = v117 + v132
	if v134 != v66 {
		v117 = v134
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v186 = v57
	v187 = v58
	v189 = v60
	v193 = v132
	goto L11
L34:
	;
	v139 = v117
	v146 = v60 + int32(1)
	goto L18
L35:
	;
	v170 = v57
	v171 = v58
	v173 = v60
	v176 = v66
	goto L15
L36:
	;
	if base.Ui32(v146) < base.Ui32(v34) {
		v57 = v161
		v58 = v162
		v60 = v146
		goto L16
	} else {
		goto L37
	}
L37:
	;
	goto L17
L38:
	;
	goto L10
L39:
	;
	if v187&int32(1) == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v201 = v187 & int32(4)
	if v193&base.B2i32(v201 != int32(0)) != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	if v36 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	if v187&int32(2) != 0 {
		v227 = int32(0)
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v227
	goto L38
L44:
	;
	v211 = int32(3)
	v212 = int32(base.Ui32(v187) >> (uint(v211) % 32))
	if v201 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v222 = int32(4)
	goto L47
L46:
	;
	v222 = v212 << (uint(int32(2)) % 32)
	goto L47
L47:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v186+v212+(int32(0)-v212)&v211+v222+int32(4))))
	v227 = v226
	goto L43
}
func F_streamNew(m *base.Module) int32 {
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
	var v10 int64
	_ = v10
	v4 = F_valkey_malloc(m, int32(72))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_raxNew(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v8
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(16)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(24)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(32)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(40)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(48)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(56)))) = v10
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(64)))) = v10
			return v4
		}
	}
}
func F_streamPropagateGroupID(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int64
	_ = v169
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v12 = int32(_a578)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[933]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[934]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v16
	v19 = v8 + int32(32)
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	v28 = int32(1)
	if base.Ui64(v23) < base.Ui64(int64(10)) {
		v85 = v28
		v86 = v5
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v160 = v19 + v159
	v161 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v161)
	v163 = int32(1)
	v164 = v160 + v163
	v165 = int32(0)
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui64(v169) < base.Ui64(int64(10)) {
		v231 = v163
		v232 = v165
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v89 = v85 + v86
	if base.Ui32(int32(21)) <= base.Ui32(v89) {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	v36 = v5
	v37 = v23
	goto L4
L4:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v37) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v85 = v28
	v86 = v77
	goto L2
L6:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v37) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v85 = int32(2)
	v86 = v36
	goto L2
L8:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v37) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v85 = int32(3)
	v86 = v36
	goto L2
L10:
	;
	v77 = v36 + int32(12)
	v81 = base.I64_div_u_s(v37, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v37) {
		v36 = v77
		v37 = v81
		goto L4
	} else {
		goto L32
	}
L11:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v37) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v37) {
		goto L24
	} else {
		goto L25
	}
L13:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v37) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v37) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v37) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v37) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v85 = int32(4)
	v86 = v36
	goto L2
L18:
	;
	v58 = int32(6)
	goto L20
L19:
	;
	v58 = int32(5)
	goto L20
L20:
	;
	v85 = v58
	v86 = v36
	goto L2
L21:
	;
	v63 = int32(8)
	goto L23
L22:
	;
	v63 = int32(7)
	goto L23
L23:
	;
	v85 = v63
	v86 = v36
	goto L2
L24:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v37) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v37) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v70 = int32(10)
	goto L28
L27:
	;
	v70 = int32(9)
	goto L28
L28:
	;
	v85 = v70
	v86 = v36
	goto L2
L29:
	;
	v75 = int32(12)
	goto L31
L30:
	;
	v75 = int32(11)
	goto L31
L31:
	;
	v85 = v75
	v86 = v36
	goto L2
L32:
	;
	goto L5
L33:
	;
	goto L44
L34:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19+v89))) = uint8(v92)
	v95 = v89 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v23) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v131 = v19 + v128
	if base.Ui64(int64(9)) < base.Ui64(v129) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v102 = v23
	v104 = v95
	goto L38
L37:
	;
	v128 = v95
	v129 = v23
	goto L35
L38:
	;
	v108 = int64(100)
	v109 = base.I64_div_u_s(v102, v108)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v102-v109*v108)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v8+int32(31)+v104))) = uint16(v118)
	v121 = v104 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v102) {
		v102 = v109
		v104 = v121
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v128 = v121
	v129 = v109
	goto L35
L40:
	;
	goto L39
L41:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v129)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v131+int32(-1)))) = uint16(v145)
	v159 = v89
	goto L1
L42:
	;
	v136 = base.I32_wrap_i64(v129) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v136)
	v159 = v89
	goto L1
L43:
	;
	v159 = int32(0)
	goto L1
L44:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v149)
	goto L43
L45:
	;
	v310 = F_sdsnewlen(m, v8+int32(32), v164+v305-(v8+int32(32)))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v235 = v231 + v232
	if base.Ui32(int32(21)) <= base.Ui32(v235) {
		goto L77
	} else {
		goto L78
	}
L47:
	;
	v182 = v165
	v183 = v169
	goto L48
L48:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v183) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v231 = v163
	v232 = v223
	goto L46
L50:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v183) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v231 = int32(2)
	v232 = v182
	goto L46
L52:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v183) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v231 = int32(3)
	v232 = v182
	goto L46
L54:
	;
	v223 = v182 + int32(12)
	v227 = base.I64_div_u_s(v183, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v183) {
		v182 = v223
		v183 = v227
		goto L48
	} else {
		goto L76
	}
L55:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v183) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v183) {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v183) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v183) {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v183) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v183) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v231 = int32(4)
	v232 = v182
	goto L46
L62:
	;
	v204 = int32(6)
	goto L64
L63:
	;
	v204 = int32(5)
	goto L64
L64:
	;
	v231 = v204
	v232 = v182
	goto L46
L65:
	;
	v209 = int32(8)
	goto L67
L66:
	;
	v209 = int32(7)
	goto L67
L67:
	;
	v231 = v209
	v232 = v182
	goto L46
L68:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v183) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v183) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v216 = int32(10)
	goto L72
L71:
	;
	v216 = int32(9)
	goto L72
L72:
	;
	v231 = v216
	v232 = v182
	goto L46
L73:
	;
	v221 = int32(12)
	goto L75
L74:
	;
	v221 = int32(11)
	goto L75
L75:
	;
	v231 = v221
	v232 = v182
	goto L46
L76:
	;
	goto L49
L77:
	;
	goto L88
L78:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164+v235))) = uint8(v238)
	v241 = v235 + int32(-1)
	if base.Ui64(int64(100)) <= base.Ui64(v169) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v277 = v164 + v274
	if base.Ui64(int64(9)) < base.Ui64(v275) {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	v248 = v169
	v250 = v241
	goto L82
L81:
	;
	v274 = v241
	v275 = v169
	goto L79
L82:
	;
	v254 = int64(100)
	v255 = base.I64_div_u_s(v248, v254)
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v248-v255*v254)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v160+int32(0)+v250))) = uint16(v264)
	v267 = v250 + int32(-2)
	if base.Ui64(int64(9999)) < base.Ui64(v248) {
		v248 = v255
		v250 = v267
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v274 = v267
	v275 = v255
	goto L79
L84:
	;
	goto L83
L85:
	;
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v275)<<(uint(int32(1))%32))+uint32(_consts[723]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v277+int32(-1)))) = uint16(v291)
	v305 = v235
	goto L45
L86:
	;
	v282 = base.I32_wrap_i64(v275) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v282)
	v305 = v235
	goto L45
L87:
	;
	v305 = int32(0)
	goto L45
L88:
	;
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v295)
	goto L87
L89:
	;
	return
L90:
	;
	v312 = F_createObject(m, v165, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v312
	v316 = *(*int32)(unsafe.Add(mBase, _consts[935]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	v319 = F_createStringObjectFromLongLong(m, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v319
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+28))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	F_alsoPropagate(m, v323, v8, int32(7), int32(3), v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	F_decrRefCount(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	F_decrRefCount(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L89
	} else {
		goto L95
	}
L95:
	;
	m.G0 = v8 + int32(80)
	return
}
func F_streamTrimByID(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(32)))) = int64(8589934592)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v5
	v31 = *(*int64)(unsafe.Add(mBase, _consts[932]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = l2
	if l2 != 0 {
		v36 = v31 * int64(100)
	} else {
		v36 = v5
	}
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v36
	v40 = int32(8)
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l1+v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v8+int32(72)))) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v44
	v48 = F_streamTrim(m, l0, v8+v40)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		return int64(0)
	} else {
		m.G0 = v8 + int32(80)
		return v48
	}
}
func F_streamTrimByLength(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(32)))) = int64(4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(72)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+64)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = l1
	v33 = *(*int64)(unsafe.Add(mBase, _consts[932]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = l2
	if l2 != 0 {
		v38 = v33 * int64(100)
	} else {
		v38 = v11
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v38
	v42 = F_streamTrim(m, l0, v7+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(80)
		return v42
	}
}
func F_streamValidateListpackIntegrity(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v172 int32
	_ = v172
	var v178 int64
	_ = v178
	var v183 int32
	_ = v183
	var v189 int64
	_ = v189
	var v194 int32
	_ = v194
	var v200 int64
	_ = v200
	var v205 int32
	_ = v205
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v314 int64
	_ = v314
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v336 int64
	_ = v336
	var v341 int64
	_ = v341
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v349 int32
	_ = v349
	var v357 int64
	_ = v357
	var v381 int64
	_ = v381
	var v400 int32
	_ = v400
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int64
	_ = v544
	var v549 int32
	_ = v549
	var v555 int64
	_ = v555
	var v560 int32
	_ = v560
	var v566 int64
	_ = v566
	var v571 int32
	_ = v571
	var v577 int64
	_ = v577
	var v582 int32
	_ = v582
	var v592 int64
	_ = v592
	var v594 int64
	_ = v594
	var v599 int32
	_ = v599
	var v610 int32
	_ = v610
	var v628 int32
	_ = v628
	var v633 int64
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int64
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v694 int64
	_ = v694
	var v700 int32
	_ = v700
	var v702 int64
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v716 int64
	_ = v716
	var v721 int64
	_ = v721
	var v725 int32
	_ = v725
	var v727 int64
	_ = v727
	var v729 int32
	_ = v729
	var v737 int64
	_ = v737
	var v761 int64
	_ = v761
	var v780 int32
	_ = v780
	var v789 int64
	_ = v789
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v924 int64
	_ = v924
	var v929 int32
	_ = v929
	var v935 int64
	_ = v935
	var v940 int32
	_ = v940
	var v946 int64
	_ = v946
	var v951 int32
	_ = v951
	var v957 int64
	_ = v957
	var v962 int32
	_ = v962
	var v972 int64
	_ = v972
	var v974 int64
	_ = v974
	var v979 int32
	_ = v979
	var v990 int32
	_ = v990
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int64
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1069 int64
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1077 int64
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1091 int64
	_ = v1091
	var v1096 int64
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1102 int64
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1112 int64
	_ = v1112
	var v1136 int64
	_ = v1136
	var v1155 int32
	_ = v1155
	var v1162 int64
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1282 int32
	_ = v1282
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1300 int64
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1311 int64
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1322 int64
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1333 int64
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1348 int64
	_ = v1348
	var v1350 int64
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1366 int32
	_ = v1366
	var v1384 int32
	_ = v1384
	var v1400 int64
	_ = v1400
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1536 int64
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1547 int64
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1558 int64
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1569 int64
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1584 int64
	_ = v1584
	var v1586 int64
	_ = v1586
	var v1591 int32
	_ = v1591
	var v1602 int32
	_ = v1602
	var v1620 int32
	_ = v1620
	var v1624 int64
	_ = v1624
	var v1631 int32
	_ = v1631
	var v1645 int64
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
	var v1729 int32
	_ = v1729
	var v1736 int32
	_ = v1736
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1763 int32
	_ = v1763
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1781 int64
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1792 int64
	_ = v1792
	var v1797 int32
	_ = v1797
	var v1803 int64
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1814 int64
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1829 int64
	_ = v1829
	var v1831 int64
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1847 int32
	_ = v1847
	var v1865 int32
	_ = v1865
	var v1869 int64
	_ = v1869
	var v1876 int32
	_ = v1876
	var v1882 int64
	_ = v1882
	var v1883 int64
	_ = v1883
	var v1884 int64
	_ = v1884
	var v1894 int64
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1976 int32
	_ = v1976
	var v1983 int32
	_ = v1983
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2010 int32
	_ = v2010
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2028 int64
	_ = v2028
	var v2033 int32
	_ = v2033
	var v2039 int64
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2050 int64
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2061 int64
	_ = v2061
	var v2066 int32
	_ = v2066
	var v2076 int64
	_ = v2076
	var v2078 int64
	_ = v2078
	var v2083 int32
	_ = v2083
	var v2094 int32
	_ = v2094
	var v2112 int32
	_ = v2112
	var v2117 int64
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2188 int32
	_ = v2188
	var v2199 int32
	_ = v2199
	var v2206 int32
	_ = v2206
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2251 int64
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2262 int64
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2273 int64
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2284 int64
	_ = v2284
	var v2289 int32
	_ = v2289
	var v2299 int64
	_ = v2299
	var v2301 int64
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2317 int32
	_ = v2317
	var v2335 int32
	_ = v2335
	var v2340 int64
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2357 int32
	_ = v2357
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2411 int32
	_ = v2411
	var v2422 int32
	_ = v2422
	var v2429 int32
	_ = v2429
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2456 int32
	_ = v2456
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2474 int64
	_ = v2474
	var v2479 int32
	_ = v2479
	var v2485 int64
	_ = v2485
	var v2490 int32
	_ = v2490
	var v2496 int64
	_ = v2496
	var v2501 int32
	_ = v2501
	var v2507 int64
	_ = v2507
	var v2512 int32
	_ = v2512
	var v2522 int64
	_ = v2522
	var v2524 int64
	_ = v2524
	var v2529 int32
	_ = v2529
	var v2540 int32
	_ = v2540
	var v2558 int32
	_ = v2558
	var v2568 int64
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2610 int32
	_ = v2610
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2641 int32
	_ = v2641
	var v2652 int32
	_ = v2652
	var v2659 int32
	_ = v2659
	var v2670 int32
	_ = v2670
	var v2673 int32
	_ = v2673
	var v2686 int32
	_ = v2686
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2699 int32
	_ = v2699
	var v2701 int32
	_ = v2701
	var v2704 int64
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2715 int64
	_ = v2715
	var v2720 int32
	_ = v2720
	var v2726 int64
	_ = v2726
	var v2731 int32
	_ = v2731
	var v2737 int64
	_ = v2737
	var v2742 int32
	_ = v2742
	var v2752 int64
	_ = v2752
	var v2754 int64
	_ = v2754
	var v2759 int32
	_ = v2759
	var v2770 int32
	_ = v2770
	var v2788 int32
	_ = v2788
	var v2791 int64
	_ = v2791
	var v2804 int64
	_ = v2804
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2829 int32
	_ = v2829
	var v2831 int32
	_ = v2831
	var v2833 int32
	_ = v2833
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2871 int32
	_ = v2871
	var v2877 int32
	_ = v2877
	var v2888 int32
	_ = v2888
	var v2895 int32
	_ = v2895
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2922 int32
	_ = v2922
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2932 int32
	_ = v2932
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2940 int64
	_ = v2940
	var v2945 int32
	_ = v2945
	var v2951 int64
	_ = v2951
	var v2956 int32
	_ = v2956
	var v2962 int64
	_ = v2962
	var v2967 int32
	_ = v2967
	var v2973 int64
	_ = v2973
	var v2978 int32
	_ = v2978
	var v2988 int64
	_ = v2988
	var v2990 int64
	_ = v2990
	var v2995 int32
	_ = v2995
	var v3006 int32
	_ = v3006
	var v3024 int32
	_ = v3024
	var v3028 int64
	_ = v3028
	var v3030 int32
	_ = v3030
	var v3039 int32
	_ = v3039
	var v3058 int32
	_ = v3058
	var v3068 int64
	_ = v3068
	var v3069 int64
	_ = v3069
	var v3070 int64
	_ = v3070
	var v3083 int64
	_ = v3083
	var v3090 int32
	_ = v3090
	var v3092 int32
	_ = v3092
	var v3093 int32
	_ = v3093
	var v3102 int32
	_ = v3102
	var v3106 int32
	_ = v3106
	var v3108 int32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3120 int32
	_ = v3120
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3125 int32
	_ = v3125
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3150 int32
	_ = v3150
	var v3156 int32
	_ = v3156
	var v3167 int32
	_ = v3167
	var v3174 int32
	_ = v3174
	var v3185 int32
	_ = v3185
	var v3188 int32
	_ = v3188
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3219 int64
	_ = v3219
	var v3224 int32
	_ = v3224
	var v3230 int64
	_ = v3230
	var v3235 int32
	_ = v3235
	var v3241 int64
	_ = v3241
	var v3246 int32
	_ = v3246
	var v3252 int64
	_ = v3252
	var v3257 int32
	_ = v3257
	var v3267 int64
	_ = v3267
	var v3269 int64
	_ = v3269
	var v3274 int32
	_ = v3274
	var v3285 int32
	_ = v3285
	var v3303 int32
	_ = v3303
	var v3307 int64
	_ = v3307
	var v3314 int32
	_ = v3314
	var v3328 int64
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3337 int64
	_ = v3337
	var v3338 int64
	_ = v3338
	var v3343 int32
	_ = v3343
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3355 int32
	_ = v3355
	var v3359 int32
	_ = v3359
	var v3361 int32
	_ = v3361
	var v3363 int32
	_ = v3363
	var v3365 int32
	_ = v3365
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3403 int32
	_ = v3403
	var v3409 int32
	_ = v3409
	var v3420 int32
	_ = v3420
	var v3427 int32
	_ = v3427
	var v3438 int32
	_ = v3438
	var v3441 int32
	_ = v3441
	var v3454 int32
	_ = v3454
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3472 int64
	_ = v3472
	var v3477 int32
	_ = v3477
	var v3483 int64
	_ = v3483
	var v3488 int32
	_ = v3488
	var v3494 int64
	_ = v3494
	var v3499 int32
	_ = v3499
	var v3505 int64
	_ = v3505
	var v3510 int32
	_ = v3510
	var v3520 int64
	_ = v3520
	var v3522 int64
	_ = v3522
	var v3527 int32
	_ = v3527
	var v3538 int32
	_ = v3538
	var v3556 int32
	_ = v3556
	var v3561 int32
	_ = v3561
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v25 = F_lpValidateIntegrity(m, l0, l1, v4, v4)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(32)
	return v3561
L2:
	;
	return int32(0)
L3:
	;
	if v25 == int32(0) {
		v3561 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v34 == int32(255) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v37
	v40 = v20 + int32(12)
	v41 = int32(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v50 == v41 {
		v233 = v41
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v37 = int32(0)
	goto L8
L7:
	;
	v37 = l0 + int32(6)
	goto L8
L8:
	;
	goto L5
L9:
	;
	if v251 == int32(0) {
		v3561 = v4
		goto L1
	} else {
		goto L61
	}
L10:
	;
	v251 = v233
	goto L9
L11:
	;
	v54 = l0 + int32(6)
	if base.Ui32(v50) < base.Ui32(v54) {
		v233 = v41
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v56 = l0 + l1
	v58 = v56 + int32(-1)
	if base.Ui32(v58) < base.Ui32(v50) {
		v233 = v41
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v60 != int32(255) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v222
	v233 = int32(1)
	goto L10
L15:
	;
	v68 = base.I32_extend8_s(v60)
	v69 = int32(192)
	v70 = v60 & v69
	v71 = int32(1)
	v73 = v60 & int32(224)
	if v73 == v69 {
		v94 = v71
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v50+int32(1) == v56 {
		v222 = int32(0)
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v251 = int32(0)
	goto L9
L18:
	;
	v95 = v50 + v94
	if base.Ui32(v95) < base.Ui32(v54) {
		v233 = v41
		goto L10
	} else {
		goto L26
	}
L19:
	;
	if int32(-1) < v68 {
		v94 = v71
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v68+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v94 = v71
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v70 == int32(128) {
		v94 = v71
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if v60&int32(240) != int32(224) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v68 != int32(-16) {
		v233 = v41
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v94 = int32(2)
	goto L18
L25:
	;
	v94 = int32(5)
	goto L18
L26:
	;
	if base.Ui32(v58) < base.Ui32(v95) {
		v233 = v41
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v98 = int32(1)
	if v68 <= int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v158 = v157 + v155
	v159 = v50 + v158
	if base.Ui32(v159) < base.Ui32(v54) {
		v233 = v41
		goto L10
	} else {
		goto L50
	}
L29:
	;
	if v70 != int32(128) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v155 = int32(1)
	v157 = v98
	goto L28
L31:
	;
	if v73 != int32(192) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v104 = int32(1)
	v155 = v60&int32(63) + v104
	v157 = v104
	goto L28
L33:
	;
	v115 = (v68 + int32(15)) & int32(255)
	if base.Ui32(v115) < base.Ui32(int32(4)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v155 = int32(2)
	v157 = v98
	goto L28
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v115<<(uint(int32(2))%32))+uint32(_consts[505])))
	v155 = v154
	v157 = v98
	goto L28
L36:
	;
	if v60&int32(240) != int32(224) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(v136) < base.Ui32(int32(128)) {
		v155 = v136
		v157 = v98
		goto L28
	} else {
		goto L42
	}
L38:
	;
	if v68 == int32(-16) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v136 = v122 | v60<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L37
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v50)+1))
	v136 = v133 + int32(5)
	goto L37
L41:
	;
	v155 = int32(0)
	v157 = v98
	goto L28
L42:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v136) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v136) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v155 = v136
	v157 = int32(2)
	goto L28
L45:
	;
	if base.Ui32(v136) < base.Ui32(int32(268435456)) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v155 = v136
	v157 = int32(3)
	goto L28
L47:
	;
	v149 = int32(4)
	goto L49
L48:
	;
	v149 = int32(5)
	goto L49
L49:
	;
	v155 = v136
	v157 = v149
	goto L28
L50:
	;
	if base.Ui32(v58) < base.Ui32(v159) {
		v233 = v41
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v162 = int32(-1)
	v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+v162))))
	v167 = base.I64_extend_i32_u(v164 & int32(127))
	if v162 < v164 {
		v217 = v167
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if v217+base.I64_extend_i32_u(v157) != base.I64_extend_i32_u(v158) {
		v233 = v41
		goto L10
	} else {
		goto L60
	}
L53:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+int32(-2)))))
	v178 = base.I64_extend_i32_u(v172&int32(127))<<(uint(int64(7))%64) | v167
	if int32(-1) < v172 {
		v217 = v178
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v183 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+int32(-3)))))
	v189 = base.I64_extend_i32_u(v183&int32(127))<<(uint(int64(14))%64) | v178
	if int32(-1) < v183 {
		v217 = v189
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v194 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+int32(-4)))))
	v200 = base.I64_extend_i32_u(v194&int32(127))<<(uint(int64(21))%64) | v189
	if int32(-1) < v194 {
		v217 = v200
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v205 = int32(*(*int8)(unsafe.Add(mBase, uint32(v159+int32(-5)))))
	if int32(-1) < v205 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v215 = base.I64_extend_i32_u(v205&int32(127))<<(uint(int64(28))%64) | v200
	goto L59
L58:
	;
	v215 = int64(-1)
	goto L59
L59:
	;
	v217 = v215
	goto L52
L60:
	;
	v222 = v159
	goto L14
L61:
	;
	if v37 == int32(0) {
		v3561 = v4
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v256 = int32(0)
	v260 = F_lpGet(m, v37, v20+int32(24), v256)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L65
	}
L63:
	;
	v412 = int32(0)
	if v410 != 0 {
		v3561 = v412
		goto L1
	} else {
		goto L94
	}
L64:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v265 = v20 + int32(16)
	v266 = int32(0)
	if base.Ui32(v263+int32(-21)) < base.Ui32(int32(-20)) {
		v400 = v266
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v260 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	v410 = v256
	v411 = v262
	goto L63
L67:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	v410 = base.B2i32(v400 == int32(0))
	v411 = v409
	goto L63
L68:
	;
	goto L67
L69:
	;
	v278 = int32(1)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v263 != v278 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v400 = int32(1)
	goto L68
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v265))) = v381
	goto L70
L72:
	;
	if v279&int32(255) == int32(45) {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	v283 = v279 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v283&int32(255)) {
		v400 = v266
		goto L68
	} else {
		goto L74
	}
L74:
	;
	if v265 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v381 = base.I64_extend_i32_u(v283) & int64(255)
	goto L71
L76:
	;
	if base.Ui32(int32(8)) < base.Ui32((v302+int32(-49))&int32(255)) {
		v400 = v266
		goto L68
	} else {
		goto L79
	}
L77:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v301 = int32(2)
	v302 = v299
	v303 = v260 + int32(1)
	goto L76
L78:
	;
	v301 = v278
	v302 = v279
	v303 = v260
	goto L76
L79:
	;
	v314 = base.I64_extend_i32_u(v302+int32(-48)) & int64(255)
	if base.Ui32(v263) <= base.Ui32(v301) {
		v357 = v314
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v279&int32(255) != int32(45) {
		goto L88
	} else {
		goto L89
	}
L81:
	;
	v320 = v301
	v322 = v314
	v324 = v303
	goto L82
L82:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+1)))
	if base.Ui32((v326+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v400 = v266
		goto L68
	} else {
		goto L84
	}
L83:
	;
	v357 = v347
	goto L80
L84:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v322) {
		v400 = v266
		goto L68
	} else {
		goto L85
	}
L85:
	;
	v336 = v322 * int64(10)
	v341 = base.I64_extend_i32_u(v326+int32(-48)) & int64(255)
	if base.Ui64(v341^int64(-1)) < base.Ui64(v336) {
		v400 = v266
		goto L68
	} else {
		goto L86
	}
L86:
	;
	v345 = int32(1)
	v347 = v336 + v341
	v349 = v320 + v345
	if v349 != v263 {
		v320 = v349
		v322 = v347
		v324 = v324 + v345
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	if v357 < int64(0) {
		v400 = v266
		goto L68
	} else {
		goto L92
	}
L89:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v357) {
		v400 = v266
		goto L68
	} else {
		goto L90
	}
L90:
	;
	if v265 == int32(0) {
		goto L70
	} else {
		goto L91
	}
L91:
	;
	v381 = int64(0) - v357
	goto L71
L92:
	;
	if v265 == int32(0) {
		goto L70
	} else {
		goto L93
	}
L93:
	;
	v381 = v357
	goto L71
L94:
	;
	if v411 < int64(0) {
		v3561 = v412
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v417 = v20 + int32(12)
	v418 = int32(0)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	if v427 == v418 {
		v610 = v418
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v628 == int32(0) {
		v3561 = v412
		goto L1
	} else {
		goto L148
	}
L97:
	;
	v628 = v610
	goto L96
L98:
	;
	v431 = l0 + int32(6)
	if base.Ui32(v427) < base.Ui32(v431) {
		v610 = v418
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v433 = l0 + l1
	v435 = v433 + int32(-1)
	if base.Ui32(v435) < base.Ui32(v427) {
		v610 = v418
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v437 != int32(255) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v599
	v610 = int32(1)
	goto L97
L102:
	;
	v445 = base.I32_extend8_s(v437)
	v446 = int32(192)
	v447 = v437 & v446
	v448 = int32(1)
	v450 = v437 & int32(224)
	if v450 == v446 {
		v471 = v448
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v427+int32(1) == v433 {
		v599 = int32(0)
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v628 = int32(0)
	goto L96
L105:
	;
	v472 = v427 + v471
	if base.Ui32(v472) < base.Ui32(v431) {
		v610 = v418
		goto L97
	} else {
		goto L113
	}
L106:
	;
	if int32(-1) < v445 {
		v471 = v448
		goto L105
	} else {
		goto L107
	}
L107:
	;
	if base.Ui32((v445+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v471 = v448
		goto L105
	} else {
		goto L108
	}
L108:
	;
	if v447 == int32(128) {
		v471 = v448
		goto L105
	} else {
		goto L109
	}
L109:
	;
	if v437&int32(240) != int32(224) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if v445 != int32(-16) {
		v610 = v418
		goto L97
	} else {
		goto L112
	}
L111:
	;
	v471 = int32(2)
	goto L105
L112:
	;
	v471 = int32(5)
	goto L105
L113:
	;
	if base.Ui32(v435) < base.Ui32(v472) {
		v610 = v418
		goto L97
	} else {
		goto L114
	}
L114:
	;
	v475 = int32(1)
	if v445 <= int32(-1) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v535 = v534 + v532
	v536 = v427 + v535
	if base.Ui32(v536) < base.Ui32(v431) {
		v610 = v418
		goto L97
	} else {
		goto L137
	}
L116:
	;
	if v447 != int32(128) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v532 = int32(1)
	v534 = v475
	goto L115
L118:
	;
	if v450 != int32(192) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v481 = int32(1)
	v532 = v437&int32(63) + v481
	v534 = v481
	goto L115
L120:
	;
	v492 = (v445 + int32(15)) & int32(255)
	if base.Ui32(v492) < base.Ui32(int32(4)) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v532 = int32(2)
	v534 = v475
	goto L115
L122:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v492<<(uint(int32(2))%32))+uint32(_consts[505])))
	v532 = v531
	v534 = v475
	goto L115
L123:
	;
	if v437&int32(240) != int32(224) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if base.Ui32(v513) < base.Ui32(int32(128)) {
		v532 = v513
		v534 = v475
		goto L115
	} else {
		goto L129
	}
L125:
	;
	if v445 == int32(-16) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	v513 = v499 | v437<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L124
L127:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v427)+1))
	v513 = v510 + int32(5)
	goto L124
L128:
	;
	v532 = int32(0)
	v534 = v475
	goto L115
L129:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v513) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v513) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v532 = v513
	v534 = int32(2)
	goto L115
L132:
	;
	if base.Ui32(v513) < base.Ui32(int32(268435456)) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v532 = v513
	v534 = int32(3)
	goto L115
L134:
	;
	v526 = int32(4)
	goto L136
L135:
	;
	v526 = int32(5)
	goto L136
L136:
	;
	v532 = v513
	v534 = v526
	goto L115
L137:
	;
	if base.Ui32(v435) < base.Ui32(v536) {
		v610 = v418
		goto L97
	} else {
		goto L138
	}
L138:
	;
	v539 = int32(-1)
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536+v539))))
	v544 = base.I64_extend_i32_u(v541 & int32(127))
	if v539 < v541 {
		v594 = v544
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v594+base.I64_extend_i32_u(v534) != base.I64_extend_i32_u(v535) {
		v610 = v418
		goto L97
	} else {
		goto L147
	}
L140:
	;
	v549 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536+int32(-2)))))
	v555 = base.I64_extend_i32_u(v549&int32(127))<<(uint(int64(7))%64) | v544
	if int32(-1) < v549 {
		v594 = v555
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v560 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536+int32(-3)))))
	v566 = base.I64_extend_i32_u(v560&int32(127))<<(uint(int64(14))%64) | v555
	if int32(-1) < v560 {
		v594 = v566
		goto L139
	} else {
		goto L142
	}
L142:
	;
	v571 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536+int32(-4)))))
	v577 = base.I64_extend_i32_u(v571&int32(127))<<(uint(int64(21))%64) | v566
	if int32(-1) < v571 {
		v594 = v577
		goto L139
	} else {
		goto L143
	}
L143:
	;
	v582 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536+int32(-5)))))
	if int32(-1) < v582 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v592 = base.I64_extend_i32_u(v582&int32(127))<<(uint(int64(28))%64) | v577
	goto L146
L145:
	;
	v592 = int64(-1)
	goto L146
L146:
	;
	v594 = v592
	goto L139
L147:
	;
	v599 = v536
	goto L101
L148:
	;
	if l2 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v636 = int32(0)
	v640 = F_lpGet(m, v415, v20+int32(24), v636)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L2
	} else {
		goto L153
	}
L150:
	;
	v633 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v633 + v411
	goto L149
L151:
	;
	v792 = int32(0)
	if v790 != 0 {
		v3561 = v792
		goto L1
	} else {
		goto L182
	}
L152:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v645 = v20 + int32(16)
	v646 = int32(0)
	if base.Ui32(v643+int32(-21)) < base.Ui32(int32(-20)) {
		v780 = v646
		goto L156
	} else {
		goto L157
	}
L153:
	;
	if v640 != 0 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	v790 = v636
	v791 = v642
	goto L151
L155:
	;
	v789 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	v790 = base.B2i32(v780 == int32(0))
	v791 = v789
	goto L151
L156:
	;
	goto L155
L157:
	;
	v658 = int32(1)
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v643 != v658 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v780 = int32(1)
	goto L156
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v645))) = v761
	goto L158
L160:
	;
	if v659&int32(255) == int32(45) {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	v663 = v659 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v663&int32(255)) {
		v780 = v646
		goto L156
	} else {
		goto L162
	}
L162:
	;
	if v645 == int32(0) {
		goto L158
	} else {
		goto L163
	}
L163:
	;
	v761 = base.I64_extend_i32_u(v663) & int64(255)
	goto L159
L164:
	;
	if base.Ui32(int32(8)) < base.Ui32((v682+int32(-49))&int32(255)) {
		v780 = v646
		goto L156
	} else {
		goto L167
	}
L165:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+1)))
	v681 = int32(2)
	v682 = v679
	v683 = v640 + int32(1)
	goto L164
L166:
	;
	v681 = v658
	v682 = v659
	v683 = v640
	goto L164
L167:
	;
	v694 = base.I64_extend_i32_u(v682+int32(-48)) & int64(255)
	if base.Ui32(v643) <= base.Ui32(v681) {
		v737 = v694
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v659&int32(255) != int32(45) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v700 = v681
	v702 = v694
	v704 = v683
	goto L170
L170:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+1)))
	if base.Ui32((v706+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v780 = v646
		goto L156
	} else {
		goto L172
	}
L171:
	;
	v737 = v727
	goto L168
L172:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v702) {
		v780 = v646
		goto L156
	} else {
		goto L173
	}
L173:
	;
	v716 = v702 * int64(10)
	v721 = base.I64_extend_i32_u(v706+int32(-48)) & int64(255)
	if base.Ui64(v721^int64(-1)) < base.Ui64(v716) {
		v780 = v646
		goto L156
	} else {
		goto L174
	}
L174:
	;
	v725 = int32(1)
	v727 = v716 + v721
	v729 = v700 + v725
	if v729 != v643 {
		v700 = v729
		v702 = v727
		v704 = v704 + v725
		goto L170
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	if v737 < int64(0) {
		v780 = v646
		goto L156
	} else {
		goto L180
	}
L177:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v737) {
		v780 = v646
		goto L156
	} else {
		goto L178
	}
L178:
	;
	if v645 == int32(0) {
		goto L158
	} else {
		goto L179
	}
L179:
	;
	v761 = int64(0) - v737
	goto L159
L180:
	;
	if v645 == int32(0) {
		goto L158
	} else {
		goto L181
	}
L181:
	;
	v761 = v737
	goto L159
L182:
	;
	if v791 < int64(0) {
		v3561 = v792
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v797 = v20 + int32(12)
	v798 = int32(0)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	if v807 == v798 {
		v990 = v798
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v1008 == int32(0) {
		v3561 = v792
		goto L1
	} else {
		goto L236
	}
L185:
	;
	v1008 = v990
	goto L184
L186:
	;
	v811 = l0 + int32(6)
	if base.Ui32(v807) < base.Ui32(v811) {
		v990 = v798
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v813 = l0 + l1
	v815 = v813 + int32(-1)
	if base.Ui32(v815) < base.Ui32(v807) {
		v990 = v798
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v817 != int32(255) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v797))) = v979
	v990 = int32(1)
	goto L185
L190:
	;
	v825 = base.I32_extend8_s(v817)
	v826 = int32(192)
	v827 = v817 & v826
	v828 = int32(1)
	v830 = v817 & int32(224)
	if v830 == v826 {
		v851 = v828
		goto L193
	} else {
		goto L194
	}
L191:
	;
	if v807+int32(1) == v813 {
		v979 = int32(0)
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v1008 = int32(0)
	goto L184
L193:
	;
	v852 = v807 + v851
	if base.Ui32(v852) < base.Ui32(v811) {
		v990 = v798
		goto L185
	} else {
		goto L201
	}
L194:
	;
	if int32(-1) < v825 {
		v851 = v828
		goto L193
	} else {
		goto L195
	}
L195:
	;
	if base.Ui32((v825+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v851 = v828
		goto L193
	} else {
		goto L196
	}
L196:
	;
	if v827 == int32(128) {
		v851 = v828
		goto L193
	} else {
		goto L197
	}
L197:
	;
	if v817&int32(240) != int32(224) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	if v825 != int32(-16) {
		v990 = v798
		goto L185
	} else {
		goto L200
	}
L199:
	;
	v851 = int32(2)
	goto L193
L200:
	;
	v851 = int32(5)
	goto L193
L201:
	;
	if base.Ui32(v815) < base.Ui32(v852) {
		v990 = v798
		goto L185
	} else {
		goto L202
	}
L202:
	;
	v855 = int32(1)
	if v825 <= int32(-1) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v915 = v914 + v912
	v916 = v807 + v915
	if base.Ui32(v916) < base.Ui32(v811) {
		v990 = v798
		goto L185
	} else {
		goto L225
	}
L204:
	;
	if v827 != int32(128) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v912 = int32(1)
	v914 = v855
	goto L203
L206:
	;
	if v830 != int32(192) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v861 = int32(1)
	v912 = v817&int32(63) + v861
	v914 = v861
	goto L203
L208:
	;
	v872 = (v825 + int32(15)) & int32(255)
	if base.Ui32(v872) < base.Ui32(int32(4)) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v912 = int32(2)
	v914 = v855
	goto L203
L210:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v872<<(uint(int32(2))%32))+uint32(_consts[505])))
	v912 = v911
	v914 = v855
	goto L203
L211:
	;
	if v817&int32(240) != int32(224) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	if base.Ui32(v893) < base.Ui32(int32(128)) {
		v912 = v893
		v914 = v855
		goto L203
	} else {
		goto L217
	}
L213:
	;
	if v825 == int32(-16) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+1)))
	v893 = v879 | v817<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L212
L215:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v807)+1))
	v893 = v890 + int32(5)
	goto L212
L216:
	;
	v912 = int32(0)
	v914 = v855
	goto L203
L217:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v893) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v893) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v912 = v893
	v914 = int32(2)
	goto L203
L220:
	;
	if base.Ui32(v893) < base.Ui32(int32(268435456)) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v912 = v893
	v914 = int32(3)
	goto L203
L222:
	;
	v906 = int32(4)
	goto L224
L223:
	;
	v906 = int32(5)
	goto L224
L224:
	;
	v912 = v893
	v914 = v906
	goto L203
L225:
	;
	if base.Ui32(v815) < base.Ui32(v916) {
		v990 = v798
		goto L185
	} else {
		goto L226
	}
L226:
	;
	v919 = int32(-1)
	v921 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+v919))))
	v924 = base.I64_extend_i32_u(v921 & int32(127))
	if v919 < v921 {
		v974 = v924
		goto L227
	} else {
		goto L228
	}
L227:
	;
	if v974+base.I64_extend_i32_u(v914) != base.I64_extend_i32_u(v915) {
		v990 = v798
		goto L185
	} else {
		goto L235
	}
L228:
	;
	v929 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+int32(-2)))))
	v935 = base.I64_extend_i32_u(v929&int32(127))<<(uint(int64(7))%64) | v924
	if int32(-1) < v929 {
		v974 = v935
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v940 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+int32(-3)))))
	v946 = base.I64_extend_i32_u(v940&int32(127))<<(uint(int64(14))%64) | v935
	if int32(-1) < v940 {
		v974 = v946
		goto L227
	} else {
		goto L230
	}
L230:
	;
	v951 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+int32(-4)))))
	v957 = base.I64_extend_i32_u(v951&int32(127))<<(uint(int64(21))%64) | v946
	if int32(-1) < v951 {
		v974 = v957
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v962 = int32(*(*int8)(unsafe.Add(mBase, uint32(v916+int32(-5)))))
	if int32(-1) < v962 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v972 = base.I64_extend_i32_u(v962&int32(127))<<(uint(int64(28))%64) | v957
	goto L234
L233:
	;
	v972 = int64(-1)
	goto L234
L234:
	;
	v974 = v972
	goto L227
L235:
	;
	v979 = v916
	goto L189
L236:
	;
	v1014 = F_lpGet(m, v795, v20+int32(24), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L2
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v1163
	v1166 = int32(0)
	if v1163 == v1166 {
		v3561 = v1166
		goto L1
	} else {
		goto L268
	}
L238:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v1020 = v20 + int32(16)
	v1021 = int32(0)
	if base.Ui32(v1018+int32(-21)) < base.Ui32(int32(-20)) {
		v1155 = v1021
		goto L242
	} else {
		goto L243
	}
L239:
	;
	if v1014 != 0 {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v1017 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	v1163 = int32(1)
	v1164 = v1017
	goto L237
L241:
	;
	v1162 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	v1163 = v1155
	v1164 = v1162
	goto L237
L242:
	;
	goto L241
L243:
	;
	v1033 = int32(1)
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014))))
	if v1018 != v1033 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	v1155 = int32(1)
	goto L242
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1020))) = v1136
	goto L244
L246:
	;
	if v1034&int32(255) == int32(45) {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v1038 = v1034 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v1038&int32(255)) {
		v1155 = v1021
		goto L242
	} else {
		goto L248
	}
L248:
	;
	if v1020 == int32(0) {
		goto L244
	} else {
		goto L249
	}
L249:
	;
	v1136 = base.I64_extend_i32_u(v1038) & int64(255)
	goto L245
L250:
	;
	if base.Ui32(int32(8)) < base.Ui32((v1057+int32(-49))&int32(255)) {
		v1155 = v1021
		goto L242
	} else {
		goto L253
	}
L251:
	;
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+1)))
	v1056 = int32(2)
	v1057 = v1054
	v1058 = v1014 + int32(1)
	goto L250
L252:
	;
	v1056 = v1033
	v1057 = v1034
	v1058 = v1014
	goto L250
L253:
	;
	v1069 = base.I64_extend_i32_u(v1057+int32(-48)) & int64(255)
	if base.Ui32(v1018) <= base.Ui32(v1056) {
		v1112 = v1069
		goto L254
	} else {
		goto L255
	}
L254:
	;
	if v1034&int32(255) != int32(45) {
		goto L262
	} else {
		goto L263
	}
L255:
	;
	v1075 = v1056
	v1077 = v1069
	v1079 = v1058
	goto L256
L256:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+1)))
	if base.Ui32((v1081+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v1155 = v1021
		goto L242
	} else {
		goto L258
	}
L257:
	;
	v1112 = v1102
	goto L254
L258:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v1077) {
		v1155 = v1021
		goto L242
	} else {
		goto L259
	}
L259:
	;
	v1091 = v1077 * int64(10)
	v1096 = base.I64_extend_i32_u(v1081+int32(-48)) & int64(255)
	if base.Ui64(v1096^int64(-1)) < base.Ui64(v1091) {
		v1155 = v1021
		goto L242
	} else {
		goto L260
	}
L260:
	;
	v1100 = int32(1)
	v1102 = v1091 + v1096
	v1104 = v1075 + v1100
	if v1104 != v1018 {
		v1075 = v1104
		v1077 = v1102
		v1079 = v1079 + v1100
		goto L256
	} else {
		goto L261
	}
L261:
	;
	goto L257
L262:
	;
	if v1112 < int64(0) {
		v1155 = v1021
		goto L242
	} else {
		goto L266
	}
L263:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v1112) {
		v1155 = v1021
		goto L242
	} else {
		goto L264
	}
L264:
	;
	if v1020 == int32(0) {
		goto L244
	} else {
		goto L265
	}
L265:
	;
	v1136 = int64(0) - v1112
	goto L245
L266:
	;
	if v1020 == int32(0) {
		goto L244
	} else {
		goto L267
	}
L267:
	;
	v1136 = v1112
	goto L245
L268:
	;
	if v1164 < int64(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L269
	}
L269:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v1173 = v20 + int32(12)
	v1174 = int32(0)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1173)))
	if v1183 == v1174 {
		v1366 = v1174
		goto L271
	} else {
		goto L272
	}
L270:
	;
	if v1384 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L322
	}
L271:
	;
	v1384 = v1366
	goto L270
L272:
	;
	v1187 = l0 + int32(6)
	if base.Ui32(v1183) < base.Ui32(v1187) {
		v1366 = v1174
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1189 = l0 + l1
	v1191 = v1189 + int32(-1)
	if base.Ui32(v1191) < base.Ui32(v1183) {
		v1366 = v1174
		goto L271
	} else {
		goto L274
	}
L274:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	if v1193 != int32(255) {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1173))) = v1355
	v1366 = int32(1)
	goto L271
L276:
	;
	v1201 = base.I32_extend8_s(v1193)
	v1202 = int32(192)
	v1203 = v1193 & v1202
	v1204 = int32(1)
	v1206 = v1193 & int32(224)
	if v1206 == v1202 {
		v1227 = v1204
		goto L279
	} else {
		goto L280
	}
L277:
	;
	if v1183+int32(1) == v1189 {
		v1355 = int32(0)
		goto L275
	} else {
		goto L278
	}
L278:
	;
	v1384 = int32(0)
	goto L270
L279:
	;
	v1228 = v1183 + v1227
	if base.Ui32(v1228) < base.Ui32(v1187) {
		v1366 = v1174
		goto L271
	} else {
		goto L287
	}
L280:
	;
	if int32(-1) < v1201 {
		v1227 = v1204
		goto L279
	} else {
		goto L281
	}
L281:
	;
	if base.Ui32((v1201+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v1227 = v1204
		goto L279
	} else {
		goto L282
	}
L282:
	;
	if v1203 == int32(128) {
		v1227 = v1204
		goto L279
	} else {
		goto L283
	}
L283:
	;
	if v1193&int32(240) != int32(224) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	if v1201 != int32(-16) {
		v1366 = v1174
		goto L271
	} else {
		goto L286
	}
L285:
	;
	v1227 = int32(2)
	goto L279
L286:
	;
	v1227 = int32(5)
	goto L279
L287:
	;
	if base.Ui32(v1191) < base.Ui32(v1228) {
		v1366 = v1174
		goto L271
	} else {
		goto L288
	}
L288:
	;
	v1231 = int32(1)
	if v1201 <= int32(-1) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1291 = v1290 + v1288
	v1292 = v1183 + v1291
	if base.Ui32(v1292) < base.Ui32(v1187) {
		v1366 = v1174
		goto L271
	} else {
		goto L311
	}
L290:
	;
	if v1203 != int32(128) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1288 = int32(1)
	v1290 = v1231
	goto L289
L292:
	;
	if v1206 != int32(192) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1237 = int32(1)
	v1288 = v1193&int32(63) + v1237
	v1290 = v1237
	goto L289
L294:
	;
	v1248 = (v1201 + int32(15)) & int32(255)
	if base.Ui32(v1248) < base.Ui32(int32(4)) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	v1288 = int32(2)
	v1290 = v1231
	goto L289
L296:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1248<<(uint(int32(2))%32))+uint32(_consts[505])))
	v1288 = v1287
	v1290 = v1231
	goto L289
L297:
	;
	if v1193&int32(240) != int32(224) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if base.Ui32(v1269) < base.Ui32(int32(128)) {
		v1288 = v1269
		v1290 = v1231
		goto L289
	} else {
		goto L303
	}
L299:
	;
	if v1201 == int32(-16) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183)+1)))
	v1269 = v1255 | v1193<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L298
L301:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+1))
	v1269 = v1266 + int32(5)
	goto L298
L302:
	;
	v1288 = int32(0)
	v1290 = v1231
	goto L289
L303:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v1269) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v1269) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v1288 = v1269
	v1290 = int32(2)
	goto L289
L306:
	;
	if base.Ui32(v1269) < base.Ui32(int32(268435456)) {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1288 = v1269
	v1290 = int32(3)
	goto L289
L308:
	;
	v1282 = int32(4)
	goto L310
L309:
	;
	v1282 = int32(5)
	goto L310
L310:
	;
	v1288 = v1269
	v1290 = v1282
	goto L289
L311:
	;
	if base.Ui32(v1191) < base.Ui32(v1292) {
		v1366 = v1174
		goto L271
	} else {
		goto L312
	}
L312:
	;
	v1295 = int32(-1)
	v1297 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1292+v1295))))
	v1300 = base.I64_extend_i32_u(v1297 & int32(127))
	if v1295 < v1297 {
		v1350 = v1300
		goto L313
	} else {
		goto L314
	}
L313:
	;
	if v1350+base.I64_extend_i32_u(v1290) != base.I64_extend_i32_u(v1291) {
		v1366 = v1174
		goto L271
	} else {
		goto L321
	}
L314:
	;
	v1305 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1292+int32(-2)))))
	v1311 = base.I64_extend_i32_u(v1305&int32(127))<<(uint(int64(7))%64) | v1300
	if int32(-1) < v1305 {
		v1350 = v1311
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1316 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1292+int32(-3)))))
	v1322 = base.I64_extend_i32_u(v1316&int32(127))<<(uint(int64(14))%64) | v1311
	if int32(-1) < v1316 {
		v1350 = v1322
		goto L313
	} else {
		goto L316
	}
L316:
	;
	v1327 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1292+int32(-4)))))
	v1333 = base.I64_extend_i32_u(v1327&int32(127))<<(uint(int64(21))%64) | v1322
	if int32(-1) < v1327 {
		v1350 = v1333
		goto L313
	} else {
		goto L317
	}
L317:
	;
	v1338 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1292+int32(-5)))))
	if int32(-1) < v1338 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1348 = base.I64_extend_i32_u(v1338&int32(127))<<(uint(int64(28))%64) | v1333
	goto L320
L319:
	;
	v1348 = int64(-1)
	goto L320
L320:
	;
	v1350 = v1348
	goto L313
L321:
	;
	v1355 = v1292
	goto L275
L322:
	;
	if v1164 == int64(0) {
		v1631 = v1171
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1645 = F_lpGetIntegerIfValid(m, v1631, v20+int32(16))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L2
	} else {
		goto L381
	}
L324:
	;
	v1400 = int64(0)
	goto L325
L325:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v1409 = v20 + int32(12)
	v1410 = int32(0)
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	if v1419 == v1410 {
		v1602 = v1410
		goto L328
	} else {
		goto L329
	}
L326:
	;
	v1631 = v1407
	goto L323
L327:
	;
	if v1620 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L379
	}
L328:
	;
	v1620 = v1602
	goto L327
L329:
	;
	v1423 = l0 + int32(6)
	if base.Ui32(v1419) < base.Ui32(v1423) {
		v1602 = v1410
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1425 = l0 + l1
	v1427 = v1425 + int32(-1)
	if base.Ui32(v1427) < base.Ui32(v1419) {
		v1602 = v1410
		goto L328
	} else {
		goto L331
	}
L331:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419))))
	if v1429 != int32(255) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1409))) = v1591
	v1602 = int32(1)
	goto L328
L333:
	;
	v1437 = base.I32_extend8_s(v1429)
	v1438 = int32(192)
	v1439 = v1429 & v1438
	v1440 = int32(1)
	v1442 = v1429 & int32(224)
	if v1442 == v1438 {
		v1463 = v1440
		goto L336
	} else {
		goto L337
	}
L334:
	;
	if v1419+int32(1) == v1425 {
		v1591 = int32(0)
		goto L332
	} else {
		goto L335
	}
L335:
	;
	v1620 = int32(0)
	goto L327
L336:
	;
	v1464 = v1419 + v1463
	if base.Ui32(v1464) < base.Ui32(v1423) {
		v1602 = v1410
		goto L328
	} else {
		goto L344
	}
L337:
	;
	if int32(-1) < v1437 {
		v1463 = v1440
		goto L336
	} else {
		goto L338
	}
L338:
	;
	if base.Ui32((v1437+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v1463 = v1440
		goto L336
	} else {
		goto L339
	}
L339:
	;
	if v1439 == int32(128) {
		v1463 = v1440
		goto L336
	} else {
		goto L340
	}
L340:
	;
	if v1429&int32(240) != int32(224) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	if v1437 != int32(-16) {
		v1602 = v1410
		goto L328
	} else {
		goto L343
	}
L342:
	;
	v1463 = int32(2)
	goto L336
L343:
	;
	v1463 = int32(5)
	goto L336
L344:
	;
	if base.Ui32(v1427) < base.Ui32(v1464) {
		v1602 = v1410
		goto L328
	} else {
		goto L345
	}
L345:
	;
	v1467 = int32(1)
	if v1437 <= int32(-1) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1527 = v1526 + v1524
	v1528 = v1419 + v1527
	if base.Ui32(v1528) < base.Ui32(v1423) {
		v1602 = v1410
		goto L328
	} else {
		goto L368
	}
L347:
	;
	if v1439 != int32(128) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1524 = int32(1)
	v1526 = v1467
	goto L346
L349:
	;
	if v1442 != int32(192) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v1473 = int32(1)
	v1524 = v1429&int32(63) + v1473
	v1526 = v1473
	goto L346
L351:
	;
	v1484 = (v1437 + int32(15)) & int32(255)
	if base.Ui32(v1484) < base.Ui32(int32(4)) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1524 = int32(2)
	v1526 = v1467
	goto L346
L353:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1484<<(uint(int32(2))%32))+uint32(_consts[505])))
	v1524 = v1523
	v1526 = v1467
	goto L346
L354:
	;
	if v1429&int32(240) != int32(224) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	if base.Ui32(v1505) < base.Ui32(int32(128)) {
		v1524 = v1505
		v1526 = v1467
		goto L346
	} else {
		goto L360
	}
L356:
	;
	if v1437 == int32(-16) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419)+1)))
	v1505 = v1491 | v1429<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L355
L358:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+1))
	v1505 = v1502 + int32(5)
	goto L355
L359:
	;
	v1524 = int32(0)
	v1526 = v1467
	goto L346
L360:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v1505) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v1505) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1524 = v1505
	v1526 = int32(2)
	goto L346
L363:
	;
	if base.Ui32(v1505) < base.Ui32(int32(268435456)) {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	v1524 = v1505
	v1526 = int32(3)
	goto L346
L365:
	;
	v1518 = int32(4)
	goto L367
L366:
	;
	v1518 = int32(5)
	goto L367
L367:
	;
	v1524 = v1505
	v1526 = v1518
	goto L346
L368:
	;
	if base.Ui32(v1427) < base.Ui32(v1528) {
		v1602 = v1410
		goto L328
	} else {
		goto L369
	}
L369:
	;
	v1531 = int32(-1)
	v1533 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1528+v1531))))
	v1536 = base.I64_extend_i32_u(v1533 & int32(127))
	if v1531 < v1533 {
		v1586 = v1536
		goto L370
	} else {
		goto L371
	}
L370:
	;
	if v1586+base.I64_extend_i32_u(v1526) != base.I64_extend_i32_u(v1527) {
		v1602 = v1410
		goto L328
	} else {
		goto L378
	}
L371:
	;
	v1541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1528+int32(-2)))))
	v1547 = base.I64_extend_i32_u(v1541&int32(127))<<(uint(int64(7))%64) | v1536
	if int32(-1) < v1541 {
		v1586 = v1547
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1552 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1528+int32(-3)))))
	v1558 = base.I64_extend_i32_u(v1552&int32(127))<<(uint(int64(14))%64) | v1547
	if int32(-1) < v1552 {
		v1586 = v1558
		goto L370
	} else {
		goto L373
	}
L373:
	;
	v1563 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1528+int32(-4)))))
	v1569 = base.I64_extend_i32_u(v1563&int32(127))<<(uint(int64(21))%64) | v1558
	if int32(-1) < v1563 {
		v1586 = v1569
		goto L370
	} else {
		goto L374
	}
L374:
	;
	v1574 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1528+int32(-5)))))
	if int32(-1) < v1574 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1584 = base.I64_extend_i32_u(v1574&int32(127))<<(uint(int64(28))%64) | v1569
	goto L377
L376:
	;
	v1584 = int64(-1)
	goto L377
L377:
	;
	v1586 = v1584
	goto L370
L378:
	;
	v1591 = v1528
	goto L332
L379:
	;
	v1624 = v1400 + int64(1)
	if v1624 != v1164 {
		v1400 = v1624
		goto L325
	} else {
		goto L380
	}
L380:
	;
	goto L326
L381:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v1647 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L382
	}
L382:
	;
	if v1645 != int64(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v1654 = v20 + int32(12)
	v1655 = int32(0)
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1654)))
	if v1664 == v1655 {
		v1847 = v1655
		goto L385
	} else {
		goto L386
	}
L384:
	;
	if v1865 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L436
	}
L385:
	;
	v1865 = v1847
	goto L384
L386:
	;
	v1668 = l0 + int32(6)
	if base.Ui32(v1664) < base.Ui32(v1668) {
		v1847 = v1655
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1670 = l0 + l1
	v1672 = v1670 + int32(-1)
	if base.Ui32(v1672) < base.Ui32(v1664) {
		v1847 = v1655
		goto L385
	} else {
		goto L388
	}
L388:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664))))
	if v1674 != int32(255) {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1654))) = v1836
	v1847 = int32(1)
	goto L385
L390:
	;
	v1682 = base.I32_extend8_s(v1674)
	v1683 = int32(192)
	v1684 = v1674 & v1683
	v1685 = int32(1)
	v1687 = v1674 & int32(224)
	if v1687 == v1683 {
		v1708 = v1685
		goto L393
	} else {
		goto L394
	}
L391:
	;
	if v1664+int32(1) == v1670 {
		v1836 = int32(0)
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v1865 = int32(0)
	goto L384
L393:
	;
	v1709 = v1664 + v1708
	if base.Ui32(v1709) < base.Ui32(v1668) {
		v1847 = v1655
		goto L385
	} else {
		goto L401
	}
L394:
	;
	if int32(-1) < v1682 {
		v1708 = v1685
		goto L393
	} else {
		goto L395
	}
L395:
	;
	if base.Ui32((v1682+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v1708 = v1685
		goto L393
	} else {
		goto L396
	}
L396:
	;
	if v1684 == int32(128) {
		v1708 = v1685
		goto L393
	} else {
		goto L397
	}
L397:
	;
	if v1674&int32(240) != int32(224) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	if v1682 != int32(-16) {
		v1847 = v1655
		goto L385
	} else {
		goto L400
	}
L399:
	;
	v1708 = int32(2)
	goto L393
L400:
	;
	v1708 = int32(5)
	goto L393
L401:
	;
	if base.Ui32(v1672) < base.Ui32(v1709) {
		v1847 = v1655
		goto L385
	} else {
		goto L402
	}
L402:
	;
	v1712 = int32(1)
	if v1682 <= int32(-1) {
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v1772 = v1771 + v1769
	v1773 = v1664 + v1772
	if base.Ui32(v1773) < base.Ui32(v1668) {
		v1847 = v1655
		goto L385
	} else {
		goto L425
	}
L404:
	;
	if v1684 != int32(128) {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1769 = int32(1)
	v1771 = v1712
	goto L403
L406:
	;
	if v1687 != int32(192) {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v1718 = int32(1)
	v1769 = v1674&int32(63) + v1718
	v1771 = v1718
	goto L403
L408:
	;
	v1729 = (v1682 + int32(15)) & int32(255)
	if base.Ui32(v1729) < base.Ui32(int32(4)) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1769 = int32(2)
	v1771 = v1712
	goto L403
L410:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1729<<(uint(int32(2))%32))+uint32(_consts[505])))
	v1769 = v1768
	v1771 = v1712
	goto L403
L411:
	;
	if v1674&int32(240) != int32(224) {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	if base.Ui32(v1750) < base.Ui32(int32(128)) {
		v1769 = v1750
		v1771 = v1712
		goto L403
	} else {
		goto L417
	}
L413:
	;
	if v1682 == int32(-16) {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	v1736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664)+1)))
	v1750 = v1736 | v1674<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L412
L415:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1664)+1))
	v1750 = v1747 + int32(5)
	goto L412
L416:
	;
	v1769 = int32(0)
	v1771 = v1712
	goto L403
L417:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v1750) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v1750) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1769 = v1750
	v1771 = int32(2)
	goto L403
L420:
	;
	if base.Ui32(v1750) < base.Ui32(int32(268435456)) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v1769 = v1750
	v1771 = int32(3)
	goto L403
L422:
	;
	v1763 = int32(4)
	goto L424
L423:
	;
	v1763 = int32(5)
	goto L424
L424:
	;
	v1769 = v1750
	v1771 = v1763
	goto L403
L425:
	;
	if base.Ui32(v1672) < base.Ui32(v1773) {
		v1847 = v1655
		goto L385
	} else {
		goto L426
	}
L426:
	;
	v1776 = int32(-1)
	v1778 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1773+v1776))))
	v1781 = base.I64_extend_i32_u(v1778 & int32(127))
	if v1776 < v1778 {
		v1831 = v1781
		goto L427
	} else {
		goto L428
	}
L427:
	;
	if v1831+base.I64_extend_i32_u(v1771) != base.I64_extend_i32_u(v1772) {
		v1847 = v1655
		goto L385
	} else {
		goto L435
	}
L428:
	;
	v1786 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1773+int32(-2)))))
	v1792 = base.I64_extend_i32_u(v1786&int32(127))<<(uint(int64(7))%64) | v1781
	if int32(-1) < v1786 {
		v1831 = v1792
		goto L427
	} else {
		goto L429
	}
L429:
	;
	v1797 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1773+int32(-3)))))
	v1803 = base.I64_extend_i32_u(v1797&int32(127))<<(uint(int64(14))%64) | v1792
	if int32(-1) < v1797 {
		v1831 = v1803
		goto L427
	} else {
		goto L430
	}
L430:
	;
	v1808 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1773+int32(-4)))))
	v1814 = base.I64_extend_i32_u(v1808&int32(127))<<(uint(int64(21))%64) | v1803
	if int32(-1) < v1808 {
		v1831 = v1814
		goto L427
	} else {
		goto L431
	}
L431:
	;
	v1819 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1773+int32(-5)))))
	if int32(-1) < v1819 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v1829 = base.I64_extend_i32_u(v1819&int32(127))<<(uint(int64(28))%64) | v1814
	goto L434
L433:
	;
	v1829 = int64(-1)
	goto L434
L434:
	;
	v1831 = v1829
	goto L427
L435:
	;
	v1836 = v1773
	goto L389
L436:
	;
	v1869 = int64(0)
	v1876 = v1652
	v1882 = v791 + v411
	v1883 = v1869
	v1884 = v1869
	goto L437
L437:
	;
	if v1882 == int64(0) {
		goto L441
	} else {
		goto L442
	}
L438:
	;
	v3561 = v1166
	goto L1
L439:
	;
	v3070 = int64(0)
	if v3069 <= v3070 {
		v3314 = v3058
		goto L725
	} else {
		goto L726
	}
L440:
	;
	v3058 = v3039
	v3068 = v2568 + int64(4)
	v3069 = v2568
	goto L439
L441:
	;
	v3030 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v3030 != 0 {
		v3561 = v1166
		goto L1
	} else {
		goto L724
	}
L442:
	;
	if v1876 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1894 = F_lpGetIntegerIfValid(m, v1876, v20+int32(16))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L2
	} else {
		goto L444
	}
L444:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v1896 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L445
	}
L445:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v1901 = v20 + int32(12)
	v1902 = int32(0)
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1901)))
	if v1911 == v1902 {
		v2094 = v1902
		goto L447
	} else {
		goto L448
	}
L446:
	;
	if v2112 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L498
	}
L447:
	;
	v2112 = v2094
	goto L446
L448:
	;
	v1915 = l0 + int32(6)
	if base.Ui32(v1911) < base.Ui32(v1915) {
		v2094 = v1902
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1917 = l0 + l1
	v1919 = v1917 + int32(-1)
	if base.Ui32(v1919) < base.Ui32(v1911) {
		v2094 = v1902
		goto L447
	} else {
		goto L450
	}
L450:
	;
	v1921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911))))
	if v1921 != int32(255) {
		goto L452
	} else {
		goto L453
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1901))) = v2083
	v2094 = int32(1)
	goto L447
L452:
	;
	v1929 = base.I32_extend8_s(v1921)
	v1930 = int32(192)
	v1931 = v1921 & v1930
	v1932 = int32(1)
	v1934 = v1921 & int32(224)
	if v1934 == v1930 {
		v1955 = v1932
		goto L455
	} else {
		goto L456
	}
L453:
	;
	if v1911+int32(1) == v1917 {
		v2083 = int32(0)
		goto L451
	} else {
		goto L454
	}
L454:
	;
	v2112 = int32(0)
	goto L446
L455:
	;
	v1956 = v1911 + v1955
	if base.Ui32(v1956) < base.Ui32(v1915) {
		v2094 = v1902
		goto L447
	} else {
		goto L463
	}
L456:
	;
	if int32(-1) < v1929 {
		v1955 = v1932
		goto L455
	} else {
		goto L457
	}
L457:
	;
	if base.Ui32((v1929+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v1955 = v1932
		goto L455
	} else {
		goto L458
	}
L458:
	;
	if v1931 == int32(128) {
		v1955 = v1932
		goto L455
	} else {
		goto L459
	}
L459:
	;
	if v1921&int32(240) != int32(224) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	if v1929 != int32(-16) {
		v2094 = v1902
		goto L447
	} else {
		goto L462
	}
L461:
	;
	v1955 = int32(2)
	goto L455
L462:
	;
	v1955 = int32(5)
	goto L455
L463:
	;
	if base.Ui32(v1919) < base.Ui32(v1956) {
		v2094 = v1902
		goto L447
	} else {
		goto L464
	}
L464:
	;
	v1959 = int32(1)
	if v1929 <= int32(-1) {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v2019 = v2018 + v2016
	v2020 = v1911 + v2019
	if base.Ui32(v2020) < base.Ui32(v1915) {
		v2094 = v1902
		goto L447
	} else {
		goto L487
	}
L466:
	;
	if v1931 != int32(128) {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	v2016 = int32(1)
	v2018 = v1959
	goto L465
L468:
	;
	if v1934 != int32(192) {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	v1965 = int32(1)
	v2016 = v1921&int32(63) + v1965
	v2018 = v1965
	goto L465
L470:
	;
	v1976 = (v1929 + int32(15)) & int32(255)
	if base.Ui32(v1976) < base.Ui32(int32(4)) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v2016 = int32(2)
	v2018 = v1959
	goto L465
L472:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1976<<(uint(int32(2))%32))+uint32(_consts[505])))
	v2016 = v2015
	v2018 = v1959
	goto L465
L473:
	;
	if v1921&int32(240) != int32(224) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	if base.Ui32(v1997) < base.Ui32(int32(128)) {
		v2016 = v1997
		v2018 = v1959
		goto L465
	} else {
		goto L479
	}
L475:
	;
	if v1929 == int32(-16) {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v1983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911)+1)))
	v1997 = v1983 | v1921<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L474
L477:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1911)+1))
	v1997 = v1994 + int32(5)
	goto L474
L478:
	;
	v2016 = int32(0)
	v2018 = v1959
	goto L465
L479:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v1997) {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v1997) {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	v2016 = v1997
	v2018 = int32(2)
	goto L465
L482:
	;
	if base.Ui32(v1997) < base.Ui32(int32(268435456)) {
		goto L484
	} else {
		goto L485
	}
L483:
	;
	v2016 = v1997
	v2018 = int32(3)
	goto L465
L484:
	;
	v2010 = int32(4)
	goto L486
L485:
	;
	v2010 = int32(5)
	goto L486
L486:
	;
	v2016 = v1997
	v2018 = v2010
	goto L465
L487:
	;
	if base.Ui32(v1919) < base.Ui32(v2020) {
		v2094 = v1902
		goto L447
	} else {
		goto L488
	}
L488:
	;
	v2023 = int32(-1)
	v2025 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2020+v2023))))
	v2028 = base.I64_extend_i32_u(v2025 & int32(127))
	if v2023 < v2025 {
		v2078 = v2028
		goto L489
	} else {
		goto L490
	}
L489:
	;
	if v2078+base.I64_extend_i32_u(v2018) != base.I64_extend_i32_u(v2019) {
		v2094 = v1902
		goto L447
	} else {
		goto L497
	}
L490:
	;
	v2033 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2020+int32(-2)))))
	v2039 = base.I64_extend_i32_u(v2033&int32(127))<<(uint(int64(7))%64) | v2028
	if int32(-1) < v2033 {
		v2078 = v2039
		goto L489
	} else {
		goto L491
	}
L491:
	;
	v2044 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2020+int32(-3)))))
	v2050 = base.I64_extend_i32_u(v2044&int32(127))<<(uint(int64(14))%64) | v2039
	if int32(-1) < v2044 {
		v2078 = v2050
		goto L489
	} else {
		goto L492
	}
L492:
	;
	v2055 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2020+int32(-4)))))
	v2061 = base.I64_extend_i32_u(v2055&int32(127))<<(uint(int64(21))%64) | v2050
	if int32(-1) < v2055 {
		v2078 = v2061
		goto L489
	} else {
		goto L493
	}
L493:
	;
	v2066 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2020+int32(-5)))))
	if int32(-1) < v2066 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2076 = base.I64_extend_i32_u(v2066&int32(127))<<(uint(int64(28))%64) | v2061
	goto L496
L495:
	;
	v2076 = int64(-1)
	goto L496
L496:
	;
	v2078 = v2076
	goto L489
L497:
	;
	v2083 = v2020
	goto L451
L498:
	;
	v2117 = F_lpGetIntegerIfValid(m, v1899, v20+int32(16))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L2
	} else {
		goto L499
	}
L499:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v2119 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L500
	}
L500:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v2124 = v20 + int32(12)
	v2125 = int32(0)
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2124)))
	if v2134 == v2125 {
		v2317 = v2125
		goto L502
	} else {
		goto L503
	}
L501:
	;
	if v2335 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L553
	}
L502:
	;
	v2335 = v2317
	goto L501
L503:
	;
	v2138 = l0 + int32(6)
	if base.Ui32(v2134) < base.Ui32(v2138) {
		v2317 = v2125
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v2140 = l0 + l1
	v2142 = v2140 + int32(-1)
	if base.Ui32(v2142) < base.Ui32(v2134) {
		v2317 = v2125
		goto L502
	} else {
		goto L505
	}
L505:
	;
	v2144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134))))
	if v2144 != int32(255) {
		goto L507
	} else {
		goto L508
	}
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2124))) = v2306
	v2317 = int32(1)
	goto L502
L507:
	;
	v2152 = base.I32_extend8_s(v2144)
	v2153 = int32(192)
	v2154 = v2144 & v2153
	v2155 = int32(1)
	v2157 = v2144 & int32(224)
	if v2157 == v2153 {
		v2178 = v2155
		goto L510
	} else {
		goto L511
	}
L508:
	;
	if v2134+int32(1) == v2140 {
		v2306 = int32(0)
		goto L506
	} else {
		goto L509
	}
L509:
	;
	v2335 = int32(0)
	goto L501
L510:
	;
	v2179 = v2134 + v2178
	if base.Ui32(v2179) < base.Ui32(v2138) {
		v2317 = v2125
		goto L502
	} else {
		goto L518
	}
L511:
	;
	if int32(-1) < v2152 {
		v2178 = v2155
		goto L510
	} else {
		goto L512
	}
L512:
	;
	if base.Ui32((v2152+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v2178 = v2155
		goto L510
	} else {
		goto L513
	}
L513:
	;
	if v2154 == int32(128) {
		v2178 = v2155
		goto L510
	} else {
		goto L514
	}
L514:
	;
	if v2144&int32(240) != int32(224) {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	if v2152 != int32(-16) {
		v2317 = v2125
		goto L502
	} else {
		goto L517
	}
L516:
	;
	v2178 = int32(2)
	goto L510
L517:
	;
	v2178 = int32(5)
	goto L510
L518:
	;
	if base.Ui32(v2142) < base.Ui32(v2179) {
		v2317 = v2125
		goto L502
	} else {
		goto L519
	}
L519:
	;
	v2182 = int32(1)
	if v2152 <= int32(-1) {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	v2242 = v2241 + v2239
	v2243 = v2134 + v2242
	if base.Ui32(v2243) < base.Ui32(v2138) {
		v2317 = v2125
		goto L502
	} else {
		goto L542
	}
L521:
	;
	if v2154 != int32(128) {
		goto L523
	} else {
		goto L524
	}
L522:
	;
	v2239 = int32(1)
	v2241 = v2182
	goto L520
L523:
	;
	if v2157 != int32(192) {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	v2188 = int32(1)
	v2239 = v2144&int32(63) + v2188
	v2241 = v2188
	goto L520
L525:
	;
	v2199 = (v2152 + int32(15)) & int32(255)
	if base.Ui32(v2199) < base.Ui32(int32(4)) {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	v2239 = int32(2)
	v2241 = v2182
	goto L520
L527:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v2199<<(uint(int32(2))%32))+uint32(_consts[505])))
	v2239 = v2238
	v2241 = v2182
	goto L520
L528:
	;
	if v2144&int32(240) != int32(224) {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	if base.Ui32(v2220) < base.Ui32(int32(128)) {
		v2239 = v2220
		v2241 = v2182
		goto L520
	} else {
		goto L534
	}
L530:
	;
	if v2152 == int32(-16) {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	v2206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2134)+1)))
	v2220 = v2206 | v2144<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L529
L532:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2134)+1))
	v2220 = v2217 + int32(5)
	goto L529
L533:
	;
	v2239 = int32(0)
	v2241 = v2182
	goto L520
L534:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v2220) {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v2220) {
		goto L537
	} else {
		goto L538
	}
L536:
	;
	v2239 = v2220
	v2241 = int32(2)
	goto L520
L537:
	;
	if base.Ui32(v2220) < base.Ui32(int32(268435456)) {
		goto L539
	} else {
		goto L540
	}
L538:
	;
	v2239 = v2220
	v2241 = int32(3)
	goto L520
L539:
	;
	v2233 = int32(4)
	goto L541
L540:
	;
	v2233 = int32(5)
	goto L541
L541:
	;
	v2239 = v2220
	v2241 = v2233
	goto L520
L542:
	;
	if base.Ui32(v2142) < base.Ui32(v2243) {
		v2317 = v2125
		goto L502
	} else {
		goto L543
	}
L543:
	;
	v2246 = int32(-1)
	v2248 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2243+v2246))))
	v2251 = base.I64_extend_i32_u(v2248 & int32(127))
	if v2246 < v2248 {
		v2301 = v2251
		goto L544
	} else {
		goto L545
	}
L544:
	;
	if v2301+base.I64_extend_i32_u(v2241) != base.I64_extend_i32_u(v2242) {
		v2317 = v2125
		goto L502
	} else {
		goto L552
	}
L545:
	;
	v2256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2243+int32(-2)))))
	v2262 = base.I64_extend_i32_u(v2256&int32(127))<<(uint(int64(7))%64) | v2251
	if int32(-1) < v2256 {
		v2301 = v2262
		goto L544
	} else {
		goto L546
	}
L546:
	;
	v2267 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2243+int32(-3)))))
	v2273 = base.I64_extend_i32_u(v2267&int32(127))<<(uint(int64(14))%64) | v2262
	if int32(-1) < v2267 {
		v2301 = v2273
		goto L544
	} else {
		goto L547
	}
L547:
	;
	v2278 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2243+int32(-4)))))
	v2284 = base.I64_extend_i32_u(v2278&int32(127))<<(uint(int64(21))%64) | v2273
	if int32(-1) < v2278 {
		v2301 = v2284
		goto L544
	} else {
		goto L548
	}
L548:
	;
	v2289 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2243+int32(-5)))))
	if int32(-1) < v2289 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v2299 = base.I64_extend_i32_u(v2289&int32(127))<<(uint(int64(28))%64) | v2284
	goto L551
L550:
	;
	v2299 = int64(-1)
	goto L551
L551:
	;
	v2301 = v2299
	goto L544
L552:
	;
	v2306 = v2243
	goto L506
L553:
	;
	v2340 = F_lpGetIntegerIfValid(m, v2122, v20+int32(16))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L2
	} else {
		goto L554
	}
L554:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v2342 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L555
	}
L555:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v2347 = v20 + int32(12)
	v2348 = int32(0)
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2347)))
	if v2357 == v2348 {
		v2540 = v2348
		goto L557
	} else {
		goto L558
	}
L556:
	;
	if v2558 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L608
	}
L557:
	;
	v2558 = v2540
	goto L556
L558:
	;
	v2361 = l0 + int32(6)
	if base.Ui32(v2357) < base.Ui32(v2361) {
		v2540 = v2348
		goto L557
	} else {
		goto L559
	}
L559:
	;
	v2363 = l0 + l1
	v2365 = v2363 + int32(-1)
	if base.Ui32(v2365) < base.Ui32(v2357) {
		v2540 = v2348
		goto L557
	} else {
		goto L560
	}
L560:
	;
	v2367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357))))
	if v2367 != int32(255) {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2347))) = v2529
	v2540 = int32(1)
	goto L557
L562:
	;
	v2375 = base.I32_extend8_s(v2367)
	v2376 = int32(192)
	v2377 = v2367 & v2376
	v2378 = int32(1)
	v2380 = v2367 & int32(224)
	if v2380 == v2376 {
		v2401 = v2378
		goto L565
	} else {
		goto L566
	}
L563:
	;
	if v2357+int32(1) == v2363 {
		v2529 = int32(0)
		goto L561
	} else {
		goto L564
	}
L564:
	;
	v2558 = int32(0)
	goto L556
L565:
	;
	v2402 = v2357 + v2401
	if base.Ui32(v2402) < base.Ui32(v2361) {
		v2540 = v2348
		goto L557
	} else {
		goto L573
	}
L566:
	;
	if int32(-1) < v2375 {
		v2401 = v2378
		goto L565
	} else {
		goto L567
	}
L567:
	;
	if base.Ui32((v2375+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v2401 = v2378
		goto L565
	} else {
		goto L568
	}
L568:
	;
	if v2377 == int32(128) {
		v2401 = v2378
		goto L565
	} else {
		goto L569
	}
L569:
	;
	if v2367&int32(240) != int32(224) {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	if v2375 != int32(-16) {
		v2540 = v2348
		goto L557
	} else {
		goto L572
	}
L571:
	;
	v2401 = int32(2)
	goto L565
L572:
	;
	v2401 = int32(5)
	goto L565
L573:
	;
	if base.Ui32(v2365) < base.Ui32(v2402) {
		v2540 = v2348
		goto L557
	} else {
		goto L574
	}
L574:
	;
	v2405 = int32(1)
	if v2375 <= int32(-1) {
		goto L576
	} else {
		goto L577
	}
L575:
	;
	v2465 = v2464 + v2462
	v2466 = v2357 + v2465
	if base.Ui32(v2466) < base.Ui32(v2361) {
		v2540 = v2348
		goto L557
	} else {
		goto L597
	}
L576:
	;
	if v2377 != int32(128) {
		goto L578
	} else {
		goto L579
	}
L577:
	;
	v2462 = int32(1)
	v2464 = v2405
	goto L575
L578:
	;
	if v2380 != int32(192) {
		goto L580
	} else {
		goto L581
	}
L579:
	;
	v2411 = int32(1)
	v2462 = v2367&int32(63) + v2411
	v2464 = v2411
	goto L575
L580:
	;
	v2422 = (v2375 + int32(15)) & int32(255)
	if base.Ui32(v2422) < base.Ui32(int32(4)) {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	v2462 = int32(2)
	v2464 = v2405
	goto L575
L582:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2422<<(uint(int32(2))%32))+uint32(_consts[505])))
	v2462 = v2461
	v2464 = v2405
	goto L575
L583:
	;
	if v2367&int32(240) != int32(224) {
		goto L585
	} else {
		goto L586
	}
L584:
	;
	if base.Ui32(v2443) < base.Ui32(int32(128)) {
		v2462 = v2443
		v2464 = v2405
		goto L575
	} else {
		goto L589
	}
L585:
	;
	if v2375 == int32(-16) {
		goto L587
	} else {
		goto L588
	}
L586:
	;
	v2429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2357)+1)))
	v2443 = v2429 | v2367<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L584
L587:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2357)+1))
	v2443 = v2440 + int32(5)
	goto L584
L588:
	;
	v2462 = int32(0)
	v2464 = v2405
	goto L575
L589:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v2443) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v2443) {
		goto L592
	} else {
		goto L593
	}
L591:
	;
	v2462 = v2443
	v2464 = int32(2)
	goto L575
L592:
	;
	if base.Ui32(v2443) < base.Ui32(int32(268435456)) {
		goto L594
	} else {
		goto L595
	}
L593:
	;
	v2462 = v2443
	v2464 = int32(3)
	goto L575
L594:
	;
	v2456 = int32(4)
	goto L596
L595:
	;
	v2456 = int32(5)
	goto L596
L596:
	;
	v2462 = v2443
	v2464 = v2456
	goto L575
L597:
	;
	if base.Ui32(v2365) < base.Ui32(v2466) {
		v2540 = v2348
		goto L557
	} else {
		goto L598
	}
L598:
	;
	v2469 = int32(-1)
	v2471 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2466+v2469))))
	v2474 = base.I64_extend_i32_u(v2471 & int32(127))
	if v2469 < v2471 {
		v2524 = v2474
		goto L599
	} else {
		goto L600
	}
L599:
	;
	if v2524+base.I64_extend_i32_u(v2464) != base.I64_extend_i32_u(v2465) {
		v2540 = v2348
		goto L557
	} else {
		goto L607
	}
L600:
	;
	v2479 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2466+int32(-2)))))
	v2485 = base.I64_extend_i32_u(v2479&int32(127))<<(uint(int64(7))%64) | v2474
	if int32(-1) < v2479 {
		v2524 = v2485
		goto L599
	} else {
		goto L601
	}
L601:
	;
	v2490 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2466+int32(-3)))))
	v2496 = base.I64_extend_i32_u(v2490&int32(127))<<(uint(int64(14))%64) | v2485
	if int32(-1) < v2490 {
		v2524 = v2496
		goto L599
	} else {
		goto L602
	}
L602:
	;
	v2501 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2466+int32(-4)))))
	v2507 = base.I64_extend_i32_u(v2501&int32(127))<<(uint(int64(21))%64) | v2496
	if int32(-1) < v2501 {
		v2524 = v2507
		goto L599
	} else {
		goto L603
	}
L603:
	;
	v2512 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2466+int32(-5)))))
	if int32(-1) < v2512 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v2522 = base.I64_extend_i32_u(v2512&int32(127))<<(uint(int64(28))%64) | v2507
	goto L606
L605:
	;
	v2522 = int64(-1)
	goto L606
L606:
	;
	v2524 = v2522
	goto L599
L607:
	;
	v2529 = v2466
	goto L561
L608:
	;
	if v1894&int64(2) == int64(0) {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v2568 = F_lpGetIntegerIfValid(m, v2345, v20+int32(16))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L2
	} else {
		goto L611
	}
L610:
	;
	v3058 = v2345
	v3068 = int64(3)
	v3069 = v1164
	goto L439
L611:
	;
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v2570 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L612
	}
L612:
	;
	if v2568 < int64(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v2577 = v20 + int32(12)
	v2578 = int32(0)
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2577)))
	if v2587 == v2578 {
		v2770 = v2578
		goto L615
	} else {
		goto L616
	}
L614:
	;
	if v2788 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L666
	}
L615:
	;
	v2788 = v2770
	goto L614
L616:
	;
	v2591 = l0 + int32(6)
	if base.Ui32(v2587) < base.Ui32(v2591) {
		v2770 = v2578
		goto L615
	} else {
		goto L617
	}
L617:
	;
	v2593 = l0 + l1
	v2595 = v2593 + int32(-1)
	if base.Ui32(v2595) < base.Ui32(v2587) {
		v2770 = v2578
		goto L615
	} else {
		goto L618
	}
L618:
	;
	v2597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2587))))
	if v2597 != int32(255) {
		goto L620
	} else {
		goto L621
	}
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2577))) = v2759
	v2770 = int32(1)
	goto L615
L620:
	;
	v2605 = base.I32_extend8_s(v2597)
	v2606 = int32(192)
	v2607 = v2597 & v2606
	v2608 = int32(1)
	v2610 = v2597 & int32(224)
	if v2610 == v2606 {
		v2631 = v2608
		goto L623
	} else {
		goto L624
	}
L621:
	;
	if v2587+int32(1) == v2593 {
		v2759 = int32(0)
		goto L619
	} else {
		goto L622
	}
L622:
	;
	v2788 = int32(0)
	goto L614
L623:
	;
	v2632 = v2587 + v2631
	if base.Ui32(v2632) < base.Ui32(v2591) {
		v2770 = v2578
		goto L615
	} else {
		goto L631
	}
L624:
	;
	if int32(-1) < v2605 {
		v2631 = v2608
		goto L623
	} else {
		goto L625
	}
L625:
	;
	if base.Ui32((v2605+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v2631 = v2608
		goto L623
	} else {
		goto L626
	}
L626:
	;
	if v2607 == int32(128) {
		v2631 = v2608
		goto L623
	} else {
		goto L627
	}
L627:
	;
	if v2597&int32(240) != int32(224) {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	if v2605 != int32(-16) {
		v2770 = v2578
		goto L615
	} else {
		goto L630
	}
L629:
	;
	v2631 = int32(2)
	goto L623
L630:
	;
	v2631 = int32(5)
	goto L623
L631:
	;
	if base.Ui32(v2595) < base.Ui32(v2632) {
		v2770 = v2578
		goto L615
	} else {
		goto L632
	}
L632:
	;
	v2635 = int32(1)
	if v2605 <= int32(-1) {
		goto L634
	} else {
		goto L635
	}
L633:
	;
	v2695 = v2694 + v2692
	v2696 = v2587 + v2695
	if base.Ui32(v2696) < base.Ui32(v2591) {
		v2770 = v2578
		goto L615
	} else {
		goto L655
	}
L634:
	;
	if v2607 != int32(128) {
		goto L636
	} else {
		goto L637
	}
L635:
	;
	v2692 = int32(1)
	v2694 = v2635
	goto L633
L636:
	;
	if v2610 != int32(192) {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v2641 = int32(1)
	v2692 = v2597&int32(63) + v2641
	v2694 = v2641
	goto L633
L638:
	;
	v2652 = (v2605 + int32(15)) & int32(255)
	if base.Ui32(v2652) < base.Ui32(int32(4)) {
		goto L640
	} else {
		goto L641
	}
L639:
	;
	v2692 = int32(2)
	v2694 = v2635
	goto L633
L640:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2652<<(uint(int32(2))%32))+uint32(_consts[505])))
	v2692 = v2691
	v2694 = v2635
	goto L633
L641:
	;
	if v2597&int32(240) != int32(224) {
		goto L643
	} else {
		goto L644
	}
L642:
	;
	if base.Ui32(v2673) < base.Ui32(int32(128)) {
		v2692 = v2673
		v2694 = v2635
		goto L633
	} else {
		goto L647
	}
L643:
	;
	if v2605 == int32(-16) {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	v2659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2587)+1)))
	v2673 = v2659 | v2597<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L642
L645:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2587)+1))
	v2673 = v2670 + int32(5)
	goto L642
L646:
	;
	v2692 = int32(0)
	v2694 = v2635
	goto L633
L647:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v2673) {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v2673) {
		goto L650
	} else {
		goto L651
	}
L649:
	;
	v2692 = v2673
	v2694 = int32(2)
	goto L633
L650:
	;
	if base.Ui32(v2673) < base.Ui32(int32(268435456)) {
		goto L652
	} else {
		goto L653
	}
L651:
	;
	v2692 = v2673
	v2694 = int32(3)
	goto L633
L652:
	;
	v2686 = int32(4)
	goto L654
L653:
	;
	v2686 = int32(5)
	goto L654
L654:
	;
	v2692 = v2673
	v2694 = v2686
	goto L633
L655:
	;
	if base.Ui32(v2595) < base.Ui32(v2696) {
		v2770 = v2578
		goto L615
	} else {
		goto L656
	}
L656:
	;
	v2699 = int32(-1)
	v2701 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2696+v2699))))
	v2704 = base.I64_extend_i32_u(v2701 & int32(127))
	if v2699 < v2701 {
		v2754 = v2704
		goto L657
	} else {
		goto L658
	}
L657:
	;
	if v2754+base.I64_extend_i32_u(v2694) != base.I64_extend_i32_u(v2695) {
		v2770 = v2578
		goto L615
	} else {
		goto L665
	}
L658:
	;
	v2709 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2696+int32(-2)))))
	v2715 = base.I64_extend_i32_u(v2709&int32(127))<<(uint(int64(7))%64) | v2704
	if int32(-1) < v2709 {
		v2754 = v2715
		goto L657
	} else {
		goto L659
	}
L659:
	;
	v2720 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2696+int32(-3)))))
	v2726 = base.I64_extend_i32_u(v2720&int32(127))<<(uint(int64(14))%64) | v2715
	if int32(-1) < v2720 {
		v2754 = v2726
		goto L657
	} else {
		goto L660
	}
L660:
	;
	v2731 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2696+int32(-4)))))
	v2737 = base.I64_extend_i32_u(v2731&int32(127))<<(uint(int64(21))%64) | v2726
	if int32(-1) < v2731 {
		v2754 = v2737
		goto L657
	} else {
		goto L661
	}
L661:
	;
	v2742 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2696+int32(-5)))))
	if int32(-1) < v2742 {
		goto L662
	} else {
		goto L663
	}
L662:
	;
	v2752 = base.I64_extend_i32_u(v2742&int32(127))<<(uint(int64(28))%64) | v2737
	goto L664
L663:
	;
	v2752 = int64(-1)
	goto L664
L664:
	;
	v2754 = v2752
	goto L657
L665:
	;
	v2759 = v2696
	goto L619
L666:
	;
	v2791 = int64(0)
	if v2568 == v2791 {
		v3039 = v2575
		goto L440
	} else {
		goto L667
	}
L667:
	;
	v2804 = v2791
	goto L668
L668:
	;
	v2811 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v2813 = v20 + int32(12)
	v2814 = int32(0)
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v2813)))
	if v2823 == v2814 {
		v3006 = v2814
		goto L671
	} else {
		goto L672
	}
L670:
	;
	if v3024 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L722
	}
L671:
	;
	v3024 = v3006
	goto L670
L672:
	;
	v2827 = l0 + int32(6)
	if base.Ui32(v2823) < base.Ui32(v2827) {
		v3006 = v2814
		goto L671
	} else {
		goto L673
	}
L673:
	;
	v2829 = l0 + l1
	v2831 = v2829 + int32(-1)
	if base.Ui32(v2831) < base.Ui32(v2823) {
		v3006 = v2814
		goto L671
	} else {
		goto L674
	}
L674:
	;
	v2833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2823))))
	if v2833 != int32(255) {
		goto L676
	} else {
		goto L677
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2813))) = v2995
	v3006 = int32(1)
	goto L671
L676:
	;
	v2841 = base.I32_extend8_s(v2833)
	v2842 = int32(192)
	v2843 = v2833 & v2842
	v2844 = int32(1)
	v2846 = v2833 & int32(224)
	if v2846 == v2842 {
		v2867 = v2844
		goto L679
	} else {
		goto L680
	}
L677:
	;
	if v2823+int32(1) == v2829 {
		v2995 = int32(0)
		goto L675
	} else {
		goto L678
	}
L678:
	;
	v3024 = int32(0)
	goto L670
L679:
	;
	v2868 = v2823 + v2867
	if base.Ui32(v2868) < base.Ui32(v2827) {
		v3006 = v2814
		goto L671
	} else {
		goto L687
	}
L680:
	;
	if int32(-1) < v2841 {
		v2867 = v2844
		goto L679
	} else {
		goto L681
	}
L681:
	;
	if base.Ui32((v2841+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v2867 = v2844
		goto L679
	} else {
		goto L682
	}
L682:
	;
	if v2843 == int32(128) {
		v2867 = v2844
		goto L679
	} else {
		goto L683
	}
L683:
	;
	if v2833&int32(240) != int32(224) {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	if v2841 != int32(-16) {
		v3006 = v2814
		goto L671
	} else {
		goto L686
	}
L685:
	;
	v2867 = int32(2)
	goto L679
L686:
	;
	v2867 = int32(5)
	goto L679
L687:
	;
	if base.Ui32(v2831) < base.Ui32(v2868) {
		v3006 = v2814
		goto L671
	} else {
		goto L688
	}
L688:
	;
	v2871 = int32(1)
	if v2841 <= int32(-1) {
		goto L690
	} else {
		goto L691
	}
L689:
	;
	v2931 = v2930 + v2928
	v2932 = v2823 + v2931
	if base.Ui32(v2932) < base.Ui32(v2827) {
		v3006 = v2814
		goto L671
	} else {
		goto L711
	}
L690:
	;
	if v2843 != int32(128) {
		goto L692
	} else {
		goto L693
	}
L691:
	;
	v2928 = int32(1)
	v2930 = v2871
	goto L689
L692:
	;
	if v2846 != int32(192) {
		goto L694
	} else {
		goto L695
	}
L693:
	;
	v2877 = int32(1)
	v2928 = v2833&int32(63) + v2877
	v2930 = v2877
	goto L689
L694:
	;
	v2888 = (v2841 + int32(15)) & int32(255)
	if base.Ui32(v2888) < base.Ui32(int32(4)) {
		goto L696
	} else {
		goto L697
	}
L695:
	;
	v2928 = int32(2)
	v2930 = v2871
	goto L689
L696:
	;
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v2888<<(uint(int32(2))%32))+uint32(_consts[505])))
	v2928 = v2927
	v2930 = v2871
	goto L689
L697:
	;
	if v2833&int32(240) != int32(224) {
		goto L699
	} else {
		goto L700
	}
L698:
	;
	if base.Ui32(v2909) < base.Ui32(int32(128)) {
		v2928 = v2909
		v2930 = v2871
		goto L689
	} else {
		goto L703
	}
L699:
	;
	if v2841 == int32(-16) {
		goto L701
	} else {
		goto L702
	}
L700:
	;
	v2895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2823)+1)))
	v2909 = v2895 | v2833<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L698
L701:
	;
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2823)+1))
	v2909 = v2906 + int32(5)
	goto L698
L702:
	;
	v2928 = int32(0)
	v2930 = v2871
	goto L689
L703:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v2909) {
		goto L704
	} else {
		goto L705
	}
L704:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v2909) {
		goto L706
	} else {
		goto L707
	}
L705:
	;
	v2928 = v2909
	v2930 = int32(2)
	goto L689
L706:
	;
	if base.Ui32(v2909) < base.Ui32(int32(268435456)) {
		goto L708
	} else {
		goto L709
	}
L707:
	;
	v2928 = v2909
	v2930 = int32(3)
	goto L689
L708:
	;
	v2922 = int32(4)
	goto L710
L709:
	;
	v2922 = int32(5)
	goto L710
L710:
	;
	v2928 = v2909
	v2930 = v2922
	goto L689
L711:
	;
	if base.Ui32(v2831) < base.Ui32(v2932) {
		v3006 = v2814
		goto L671
	} else {
		goto L712
	}
L712:
	;
	v2935 = int32(-1)
	v2937 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2932+v2935))))
	v2940 = base.I64_extend_i32_u(v2937 & int32(127))
	if v2935 < v2937 {
		v2990 = v2940
		goto L713
	} else {
		goto L714
	}
L713:
	;
	if v2990+base.I64_extend_i32_u(v2930) != base.I64_extend_i32_u(v2931) {
		v3006 = v2814
		goto L671
	} else {
		goto L721
	}
L714:
	;
	v2945 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2932+int32(-2)))))
	v2951 = base.I64_extend_i32_u(v2945&int32(127))<<(uint(int64(7))%64) | v2940
	if int32(-1) < v2945 {
		v2990 = v2951
		goto L713
	} else {
		goto L715
	}
L715:
	;
	v2956 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2932+int32(-3)))))
	v2962 = base.I64_extend_i32_u(v2956&int32(127))<<(uint(int64(14))%64) | v2951
	if int32(-1) < v2956 {
		v2990 = v2962
		goto L713
	} else {
		goto L716
	}
L716:
	;
	v2967 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2932+int32(-4)))))
	v2973 = base.I64_extend_i32_u(v2967&int32(127))<<(uint(int64(21))%64) | v2962
	if int32(-1) < v2967 {
		v2990 = v2973
		goto L713
	} else {
		goto L717
	}
L717:
	;
	v2978 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2932+int32(-5)))))
	if int32(-1) < v2978 {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	v2988 = base.I64_extend_i32_u(v2978&int32(127))<<(uint(int64(28))%64) | v2973
	goto L720
L719:
	;
	v2988 = int64(-1)
	goto L720
L720:
	;
	v2990 = v2988
	goto L713
L721:
	;
	v2995 = v2932
	goto L675
L722:
	;
	v3028 = v2804 + int64(1)
	if v3028 == v2568 {
		v3039 = v2811
		goto L440
	} else {
		goto L723
	}
L723:
	;
	v2804 = v3028
	goto L668
L724:
	;
	v3561 = base.B2i32(v1883 == v411) & base.B2i32(v1884 == v791)
	goto L1
L725:
	;
	v3328 = F_lpGetIntegerIfValid(m, v3314, v20+int32(16))
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L2
	} else {
		goto L783
	}
L726:
	;
	v3083 = v3070
	goto L727
L727:
	;
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v3092 = v20 + int32(12)
	v3093 = int32(0)
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(v3092)))
	if v3102 == v3093 {
		v3285 = v3093
		goto L730
	} else {
		goto L731
	}
L728:
	;
	v3314 = v3090
	goto L725
L729:
	;
	if v3303 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L781
	}
L730:
	;
	v3303 = v3285
	goto L729
L731:
	;
	v3106 = l0 + int32(6)
	if base.Ui32(v3102) < base.Ui32(v3106) {
		v3285 = v3093
		goto L730
	} else {
		goto L732
	}
L732:
	;
	v3108 = l0 + l1
	v3110 = v3108 + int32(-1)
	if base.Ui32(v3110) < base.Ui32(v3102) {
		v3285 = v3093
		goto L730
	} else {
		goto L733
	}
L733:
	;
	v3112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3102))))
	if v3112 != int32(255) {
		goto L735
	} else {
		goto L736
	}
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3092))) = v3274
	v3285 = int32(1)
	goto L730
L735:
	;
	v3120 = base.I32_extend8_s(v3112)
	v3121 = int32(192)
	v3122 = v3112 & v3121
	v3123 = int32(1)
	v3125 = v3112 & int32(224)
	if v3125 == v3121 {
		v3146 = v3123
		goto L738
	} else {
		goto L739
	}
L736:
	;
	if v3102+int32(1) == v3108 {
		v3274 = int32(0)
		goto L734
	} else {
		goto L737
	}
L737:
	;
	v3303 = int32(0)
	goto L729
L738:
	;
	v3147 = v3102 + v3146
	if base.Ui32(v3147) < base.Ui32(v3106) {
		v3285 = v3093
		goto L730
	} else {
		goto L746
	}
L739:
	;
	if int32(-1) < v3120 {
		v3146 = v3123
		goto L738
	} else {
		goto L740
	}
L740:
	;
	if base.Ui32((v3120+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v3146 = v3123
		goto L738
	} else {
		goto L741
	}
L741:
	;
	if v3122 == int32(128) {
		v3146 = v3123
		goto L738
	} else {
		goto L742
	}
L742:
	;
	if v3112&int32(240) != int32(224) {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	if v3120 != int32(-16) {
		v3285 = v3093
		goto L730
	} else {
		goto L745
	}
L744:
	;
	v3146 = int32(2)
	goto L738
L745:
	;
	v3146 = int32(5)
	goto L738
L746:
	;
	if base.Ui32(v3110) < base.Ui32(v3147) {
		v3285 = v3093
		goto L730
	} else {
		goto L747
	}
L747:
	;
	v3150 = int32(1)
	if v3120 <= int32(-1) {
		goto L749
	} else {
		goto L750
	}
L748:
	;
	v3210 = v3209 + v3207
	v3211 = v3102 + v3210
	if base.Ui32(v3211) < base.Ui32(v3106) {
		v3285 = v3093
		goto L730
	} else {
		goto L770
	}
L749:
	;
	if v3122 != int32(128) {
		goto L751
	} else {
		goto L752
	}
L750:
	;
	v3207 = int32(1)
	v3209 = v3150
	goto L748
L751:
	;
	if v3125 != int32(192) {
		goto L753
	} else {
		goto L754
	}
L752:
	;
	v3156 = int32(1)
	v3207 = v3112&int32(63) + v3156
	v3209 = v3156
	goto L748
L753:
	;
	v3167 = (v3120 + int32(15)) & int32(255)
	if base.Ui32(v3167) < base.Ui32(int32(4)) {
		goto L755
	} else {
		goto L756
	}
L754:
	;
	v3207 = int32(2)
	v3209 = v3150
	goto L748
L755:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v3167<<(uint(int32(2))%32))+uint32(_consts[505])))
	v3207 = v3206
	v3209 = v3150
	goto L748
L756:
	;
	if v3112&int32(240) != int32(224) {
		goto L758
	} else {
		goto L759
	}
L757:
	;
	if base.Ui32(v3188) < base.Ui32(int32(128)) {
		v3207 = v3188
		v3209 = v3150
		goto L748
	} else {
		goto L762
	}
L758:
	;
	if v3120 == int32(-16) {
		goto L760
	} else {
		goto L761
	}
L759:
	;
	v3174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3102)+1)))
	v3188 = v3174 | v3112<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L757
L760:
	;
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3102)+1))
	v3188 = v3185 + int32(5)
	goto L757
L761:
	;
	v3207 = int32(0)
	v3209 = v3150
	goto L748
L762:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v3188) {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v3188) {
		goto L765
	} else {
		goto L766
	}
L764:
	;
	v3207 = v3188
	v3209 = int32(2)
	goto L748
L765:
	;
	if base.Ui32(v3188) < base.Ui32(int32(268435456)) {
		goto L767
	} else {
		goto L768
	}
L766:
	;
	v3207 = v3188
	v3209 = int32(3)
	goto L748
L767:
	;
	v3201 = int32(4)
	goto L769
L768:
	;
	v3201 = int32(5)
	goto L769
L769:
	;
	v3207 = v3188
	v3209 = v3201
	goto L748
L770:
	;
	if base.Ui32(v3110) < base.Ui32(v3211) {
		v3285 = v3093
		goto L730
	} else {
		goto L771
	}
L771:
	;
	v3214 = int32(-1)
	v3216 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3211+v3214))))
	v3219 = base.I64_extend_i32_u(v3216 & int32(127))
	if v3214 < v3216 {
		v3269 = v3219
		goto L772
	} else {
		goto L773
	}
L772:
	;
	if v3269+base.I64_extend_i32_u(v3209) != base.I64_extend_i32_u(v3210) {
		v3285 = v3093
		goto L730
	} else {
		goto L780
	}
L773:
	;
	v3224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3211+int32(-2)))))
	v3230 = base.I64_extend_i32_u(v3224&int32(127))<<(uint(int64(7))%64) | v3219
	if int32(-1) < v3224 {
		v3269 = v3230
		goto L772
	} else {
		goto L774
	}
L774:
	;
	v3235 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3211+int32(-3)))))
	v3241 = base.I64_extend_i32_u(v3235&int32(127))<<(uint(int64(14))%64) | v3230
	if int32(-1) < v3235 {
		v3269 = v3241
		goto L772
	} else {
		goto L775
	}
L775:
	;
	v3246 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3211+int32(-4)))))
	v3252 = base.I64_extend_i32_u(v3246&int32(127))<<(uint(int64(21))%64) | v3241
	if int32(-1) < v3246 {
		v3269 = v3252
		goto L772
	} else {
		goto L776
	}
L776:
	;
	v3257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3211+int32(-5)))))
	if int32(-1) < v3257 {
		goto L777
	} else {
		goto L778
	}
L777:
	;
	v3267 = base.I64_extend_i32_u(v3257&int32(127))<<(uint(int64(28))%64) | v3252
	goto L779
L778:
	;
	v3267 = int64(-1)
	goto L779
L779:
	;
	v3269 = v3267
	goto L772
L780:
	;
	v3274 = v3211
	goto L734
L781:
	;
	v3307 = v3083 + int64(1)
	if v3307 != v3069 {
		v3083 = v3307
		goto L727
	} else {
		goto L782
	}
L782:
	;
	goto L728
L783:
	;
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v3330 == int32(0) {
		v3561 = v1166
		goto L1
	} else {
		goto L784
	}
L784:
	;
	if v3328 != v3068+v3069 {
		v3561 = v1166
		goto L1
	} else {
		goto L785
	}
L785:
	;
	v3337 = int64(1)
	v3338 = v1894 & v3337
	v3343 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v3345 = v20 + int32(12)
	v3346 = int32(0)
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3345)))
	if v3355 == v3346 {
		v3538 = v3346
		goto L787
	} else {
		goto L788
	}
L786:
	;
	if v3556 != 0 {
		v1876 = v3343
		v1882 = v1882 + int64(-1)
		v1883 = v3338 ^ v3337 + v1883
		v1884 = v3338 + v1884
		goto L437
	} else {
		goto L838
	}
L787:
	;
	v3556 = v3538
	goto L786
L788:
	;
	v3359 = l0 + int32(6)
	if base.Ui32(v3355) < base.Ui32(v3359) {
		v3538 = v3346
		goto L787
	} else {
		goto L789
	}
L789:
	;
	v3361 = l0 + l1
	v3363 = v3361 + int32(-1)
	if base.Ui32(v3363) < base.Ui32(v3355) {
		v3538 = v3346
		goto L787
	} else {
		goto L790
	}
L790:
	;
	v3365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3355))))
	if v3365 != int32(255) {
		goto L792
	} else {
		goto L793
	}
L791:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3345))) = v3527
	v3538 = int32(1)
	goto L787
L792:
	;
	v3373 = base.I32_extend8_s(v3365)
	v3374 = int32(192)
	v3375 = v3365 & v3374
	v3376 = int32(1)
	v3378 = v3365 & int32(224)
	if v3378 == v3374 {
		v3399 = v3376
		goto L795
	} else {
		goto L796
	}
L793:
	;
	if v3355+int32(1) == v3361 {
		v3527 = int32(0)
		goto L791
	} else {
		goto L794
	}
L794:
	;
	v3556 = int32(0)
	goto L786
L795:
	;
	v3400 = v3355 + v3399
	if base.Ui32(v3400) < base.Ui32(v3359) {
		v3538 = v3346
		goto L787
	} else {
		goto L803
	}
L796:
	;
	if int32(-1) < v3373 {
		v3399 = v3376
		goto L795
	} else {
		goto L797
	}
L797:
	;
	if base.Ui32((v3373+int32(15))&int32(255)) < base.Ui32(int32(4)) {
		v3399 = v3376
		goto L795
	} else {
		goto L798
	}
L798:
	;
	if v3375 == int32(128) {
		v3399 = v3376
		goto L795
	} else {
		goto L799
	}
L799:
	;
	if v3365&int32(240) != int32(224) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	if v3373 != int32(-16) {
		v3538 = v3346
		goto L787
	} else {
		goto L802
	}
L801:
	;
	v3399 = int32(2)
	goto L795
L802:
	;
	v3399 = int32(5)
	goto L795
L803:
	;
	if base.Ui32(v3363) < base.Ui32(v3400) {
		v3538 = v3346
		goto L787
	} else {
		goto L804
	}
L804:
	;
	v3403 = int32(1)
	if v3373 <= int32(-1) {
		goto L806
	} else {
		goto L807
	}
L805:
	;
	v3463 = v3462 + v3460
	v3464 = v3355 + v3463
	if base.Ui32(v3464) < base.Ui32(v3359) {
		v3538 = v3346
		goto L787
	} else {
		goto L827
	}
L806:
	;
	if v3375 != int32(128) {
		goto L808
	} else {
		goto L809
	}
L807:
	;
	v3460 = int32(1)
	v3462 = v3403
	goto L805
L808:
	;
	if v3378 != int32(192) {
		goto L810
	} else {
		goto L811
	}
L809:
	;
	v3409 = int32(1)
	v3460 = v3365&int32(63) + v3409
	v3462 = v3409
	goto L805
L810:
	;
	v3420 = (v3373 + int32(15)) & int32(255)
	if base.Ui32(v3420) < base.Ui32(int32(4)) {
		goto L812
	} else {
		goto L813
	}
L811:
	;
	v3460 = int32(2)
	v3462 = v3403
	goto L805
L812:
	;
	v3459 = *(*int32)(unsafe.Add(mBase, uint32(v3420<<(uint(int32(2))%32))+uint32(_consts[505])))
	v3460 = v3459
	v3462 = v3403
	goto L805
L813:
	;
	if v3365&int32(240) != int32(224) {
		goto L815
	} else {
		goto L816
	}
L814:
	;
	if base.Ui32(v3441) < base.Ui32(int32(128)) {
		v3460 = v3441
		v3462 = v3403
		goto L805
	} else {
		goto L819
	}
L815:
	;
	if v3373 == int32(-16) {
		goto L817
	} else {
		goto L818
	}
L816:
	;
	v3427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3355)+1)))
	v3441 = v3427 | v3365<<(uint(int32(8))%32)&int32(3840) + int32(2)
	goto L814
L817:
	;
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(v3355)+1))
	v3441 = v3438 + int32(5)
	goto L814
L818:
	;
	v3460 = int32(0)
	v3462 = v3403
	goto L805
L819:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v3441) {
		goto L820
	} else {
		goto L821
	}
L820:
	;
	if base.Ui32(int32(2097152)) <= base.Ui32(v3441) {
		goto L822
	} else {
		goto L823
	}
L821:
	;
	v3460 = v3441
	v3462 = int32(2)
	goto L805
L822:
	;
	if base.Ui32(v3441) < base.Ui32(int32(268435456)) {
		goto L824
	} else {
		goto L825
	}
L823:
	;
	v3460 = v3441
	v3462 = int32(3)
	goto L805
L824:
	;
	v3454 = int32(4)
	goto L826
L825:
	;
	v3454 = int32(5)
	goto L826
L826:
	;
	v3460 = v3441
	v3462 = v3454
	goto L805
L827:
	;
	if base.Ui32(v3363) < base.Ui32(v3464) {
		v3538 = v3346
		goto L787
	} else {
		goto L828
	}
L828:
	;
	v3467 = int32(-1)
	v3469 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3464+v3467))))
	v3472 = base.I64_extend_i32_u(v3469 & int32(127))
	if v3467 < v3469 {
		v3522 = v3472
		goto L829
	} else {
		goto L830
	}
L829:
	;
	if v3522+base.I64_extend_i32_u(v3462) != base.I64_extend_i32_u(v3463) {
		v3538 = v3346
		goto L787
	} else {
		goto L837
	}
L830:
	;
	v3477 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3464+int32(-2)))))
	v3483 = base.I64_extend_i32_u(v3477&int32(127))<<(uint(int64(7))%64) | v3472
	if int32(-1) < v3477 {
		v3522 = v3483
		goto L829
	} else {
		goto L831
	}
L831:
	;
	v3488 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3464+int32(-3)))))
	v3494 = base.I64_extend_i32_u(v3488&int32(127))<<(uint(int64(14))%64) | v3483
	if int32(-1) < v3488 {
		v3522 = v3494
		goto L829
	} else {
		goto L832
	}
L832:
	;
	v3499 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3464+int32(-4)))))
	v3505 = base.I64_extend_i32_u(v3499&int32(127))<<(uint(int64(21))%64) | v3494
	if int32(-1) < v3499 {
		v3522 = v3505
		goto L829
	} else {
		goto L833
	}
L833:
	;
	v3510 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3464+int32(-5)))))
	if int32(-1) < v3510 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	v3520 = base.I64_extend_i32_u(v3510&int32(127))<<(uint(int64(28))%64) | v3505
	goto L836
L835:
	;
	v3520 = int64(-1)
	goto L836
L836:
	;
	v3522 = v3520
	goto L829
L837:
	;
	v3527 = v3464
	goto L791
L838:
	;
	goto L438
}
