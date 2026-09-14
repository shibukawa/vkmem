package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_addKeysToIncrFindBatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	v8 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(2064)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(1099511627776)
	v23 = F_getKeysFromCommand(m, l1, l2, l3, v14+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_getKeysFreeResult(m, v14+int32(4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L43
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(0) {
		v183 = l5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v30 == int32(0) {
		v130 = v8
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v130<<(uint(int32(2))%32))))
	goto L34
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2+v34<<(uint(int32(2))%32))))
	v39 = F_objectGetVal(m, v38)
	mBase = m.M
	v41 = F_objectGetVal(m, v38)
	mBase = m.M
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-1)))))
	switch v44 & int32(7) {
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
		v61 = int32(0)
		goto L7
	}
L7:
	;
	v62 = int32(0)
	if v61 < int32(1) {
		v82 = v62
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-17))))
	v61 = v60
	goto L7
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41+int32(-9))))
	v61 = v57
	goto L7
L10:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41+int32(-5)))))
	v61 = v54
	goto L7
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+int32(-3)))))
	v61 = v51
	goto L7
L12:
	;
	v61 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	goto L7
L13:
	;
	v130 = v124 & int32(16383)
	goto L5
L14:
	;
	goto L13
L15:
	;
	v93 = v82 + int32(1)
	if v61 <= v93 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	v91 = F_crc16(m, v39, v61)
	mBase = m.M
	v124 = v91
	goto L14
L17:
	;
	if v82 != v61 {
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v70 = v62
	goto L19
L19:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v70))))
	if v74 == int32(123) {
		v82 = v70
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v78 = v70 + int32(1)
	if v78 != v61 {
		v70 = v78
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	goto L16
L24:
	;
	v121 = F_crc16(m, v39+v82+int32(1), v99+(v82^int32(-1)))
	mBase = m.M
	v124 = v121
	goto L14
L25:
	;
	v114 = F_crc16(m, v39, v61)
	mBase = m.M
	v124 = v114
	goto L14
L26:
	;
	v99 = v93
	goto L28
L27:
	;
	if v99 == v61 {
		goto L25
	} else {
		goto L32
	}
L28:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v99))))
	if v101 == int32(125) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v105 = v99 + int32(1)
	if v105 != v61 {
		v99 = v105
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	if v99 != v93 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	if v139 == int32(0) {
		v183 = l5
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v23 < int32(1) {
		v183 = l5
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if l6 <= l5 {
		v183 = l5
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v151 = l5
	v154 = int32(0)
	goto L38
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v154<<(uint(int32(3))%32))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2+v164<<(uint(int32(2))%32))))
	v169 = F_objectGetVal(m, v168)
	mBase = m.M
	F_hashtableIncrementalFindInit(m, l4+v151*int32(40), v139, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v183 = v173
	goto L1
L40:
	;
	v172 = int32(1)
	v173 = v151 + v172
	v175 = v154 + v172
	if v23 <= v175 {
		v183 = v173
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v173 < l6 {
		v151 = v173
		v154 = v175
		goto L38
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	m.G0 = v14 + int32(2064)
	return v183
}
func F_countKeysInSlot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	if int32(1) <= v9 {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		v15 = int32(0)
		v18 = v9
		v19 = v15
		v20 = v14
		v21 = v15
		for {
			v24 = int32(0)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21<<(uint(int32(2))%32))))
			if v28 == v24 {
				v44 = v18
				v45 = v20
				v46 = v24
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+l0<<(uint(int32(2))%32))))
				if v36 != 0 {
					v38 = F_hashtableSize(m, v36)
					mBase = m.M
					v39 = v38
				} else {
					v39 = int32(0)
				}
				v40 = int32(_a44)
				v41 = *(*int32)(unsafe.Add(mBase, _consts[65]))
				v43 = *(*int32)(unsafe.Add(mBase, _consts[137]))
				v44 = v41
				v45 = v43
				v46 = v39
			}
			v47 = v46 + v19
			v49 = v21 + int32(1)
			if v49 < v44 {
				v18 = v44
				v19 = v47
				v20 = v45
				v21 = v49
				continue
			} else {
				break
			}
			break
		}
		v53 = v47
	} else {
		v53 = int32(0)
	}
	return v53
}
func F_evalGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = F_genericGetKeys(m, int32(0), int32(2), int32(3), int32(1), l1, l2, l3)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_keysCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = F_objectGetVal(m, v20)
	mBase = m.M
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-1)))))
	switch v24 & int32(7) {
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
		v41 = int32(0)
		goto L1
	}
L1:
	;
	v42 = F_addReplyDeferredLen(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-17))))
	v41 = v40
	goto L1
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-9))))
	v41 = v37
	goto L1
L4:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(-5)))))
	v41 = v34
	goto L1
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(-3)))))
	v41 = v31
	goto L1
L6:
	;
	v41 = int32(base.Ui32(v24) >> (uint(int32(3)) % 32))
	goto L1
L7:
	;
	return
L8:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v49 = base.B2i32(v44 == int32(42)) & base.B2i32(v41 == int32(1))
	v51 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	m.G0 = v16 + int32(16)
	return
L10:
	;
	v234 = int32(0)
	goto L64
L11:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v219 = F_kvstoreIteratorInit(m, v217, int32(1))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L61
	}
L12:
	;
	if v49 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v41 < int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v138 == int32(-1) {
		goto L11
	} else {
		goto L39
	}
L15:
	;
	v127 = F_crc16(m, v21, v41)
	mBase = m.M
	v138 = v127 & int32(16383)
	goto L14
L16:
	;
	v62 = int32(-1)
	v67 = v62
	v68 = int32(0)
	goto L17
L17:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v68))))
	v76 = v74 + int32(-63)
	if base.Ui32(int32(29)) < base.Ui32(v76) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L15
L19:
	;
	v117 = v68 + int32(1)
	if v117 != v41 {
		v67 = v114
		v68 = v117
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v138 = v112
	goto L14
L21:
	;
	if v74 == int32(42) {
		v112 = v62
		goto L20
	} else {
		goto L24
	}
L22:
	;
	if int32(1)<<(uint(v76)%32)&int32(805306369) != 0 {
		v112 = v62
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v67 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if int32(0) <= v67 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v74 == int32(123) {
		v114 = v68
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v95 = base.B2i32(v68 == v67+int32(1))
	if v68 == v67+int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v114 = v67
	goto L19
L30:
	;
	v96 = int32(-2)
	goto L32
L31:
	;
	v96 = v67
	goto L32
L32:
	;
	if v74 == int32(125) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = v96
	goto L35
L34:
	;
	v99 = v67
	goto L35
L35:
	;
	if v74 != int32(125) {
		v114 = v99
		goto L19
	} else {
		goto L36
	}
L36:
	;
	if v68 == v67+int32(1) {
		v114 = v99
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v108 = F_crc16(m, v21+v67+int32(1), v68+(v67^int32(-1)))
	mBase = m.M
	v112 = v108 & int32(16383)
	goto L20
L38:
	;
	goto L18
L39:
	;
	v141 = int32(0)
	v144 = m.G0
	v146 = v144 - int32(16)
	m.G0 = v146
	v149 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+uint32(_consts[212])))
	F_listRewind(m, v150, v146)
	mBase = m.M
	v153 = F_listNext(m, v146)
	mBase = m.M
	if v153 == v141 {
		v198 = v141
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v212 = F_kvstoreGetHashtableIterator(m, v210, v138, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L60
	}
L41:
	;
	if v198 == int32(0) {
		goto L40
	} else {
		goto L58
	}
L42:
	;
	m.G0 = v146 + int32(16)
	goto L41
L43:
	;
	v159 = v153
	goto L45
L44:
	;
	v198 = int32(1)
	goto L42
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v161 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v194 = F_listNext(m, v146)
	mBase = m.M
	if v194 != 0 {
		v159 = v194
		goto L45
	} else {
		goto L57
	}
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+156))
	if base.Ui32(v164+int32(-18)) < base.Ui32(int32(3)) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)+164))
	v171 = v146 + int32(8)
	F_listRewind(m, v169, v171)
	mBase = m.M
	v175 = F_listNext(m, v171)
	mBase = m.M
	if v175 == int32(0) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v181 = v175
	goto L51
L51:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v138 < v183 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L47
L53:
	;
	v189 = F_listNext(m, v146+int32(8))
	mBase = m.M
	if v189 != 0 {
		v181 = v189
		goto L51
	} else {
		goto L56
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v138 <= v185 {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L52
L57:
	;
	v198 = v141
	goto L42
L58:
	;
	F_setDeferredArrayLen(m, l0, v42, int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L9
L60:
	;
	v222 = int32(0)
	v223 = v212
	goto L10
L61:
	;
	v222 = v219
	v223 = int32(0)
	goto L10
L62:
	;
	if v222 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L63:
	;
	F_kvstoreReleaseHashtableIterator(m, v223)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L7
	} else {
		goto L109
	}
L64:
	;
	if v223 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v223 == int32(0) {
		v400 = v381
		goto L62
	} else {
		goto L108
	}
L66:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v251 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v254&int32(2) == v251 {
		v274 = v251
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v246 = F_kvstoreIteratorNext(m, v222, v16+int32(12))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L71
	}
L68:
	;
	v242 = F_kvstoreHashtableIteratorNext(m, v223, v16+int32(12))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	if v242 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v393 = v234
	goto L63
L71:
	;
	if v246 == int32(0) {
		v400 = v234
		goto L62
	} else {
		goto L72
	}
L72:
	;
	goto L66
L73:
	;
	if v49 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L73
L75:
	;
	v268 = v250 + (v254&int32(4) ^ int32(12)) + v254<<(uint(int32(3))%32)&int32(8)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v274 = v268 + v269 + int32(1)
	goto L74
L76:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+201)))
	if v385&int32(4) == int32(0) {
		v234 = v381
		goto L64
	} else {
		goto L107
	}
L77:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v316 != 0 {
		goto L87
	} else {
		goto L88
	}
L78:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-1)))))
	switch v278 & int32(7) {
	case 0:
		goto L84
	case 1:
		goto L83
	case 2:
		goto L82
	case 3:
		goto L81
	case 4:
		goto L80
	default:
		v295 = int32(0)
		goto L79
	}
L79:
	;
	v296 = int32(0)
	v298 = m.G0
	v299 = int32(16)
	v300 = v298 - v299
	m.G0 = v300
	*(*int32)(unsafe.Add(mBase, uint32(v300)+12)) = v296
	v307 = F_stringmatchlen_impl(m, v21, v41, v274, v295, v296, v300+int32(12), v296)
	mBase = m.M
	m.G0 = v300 + v299
	goto L85
L80:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-17))))
	v295 = v294
	goto L79
L81:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-9))))
	v295 = v291
	goto L79
L82:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+int32(-5)))))
	v295 = v288
	goto L79
L83:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-3)))))
	v295 = v285
	goto L79
L84:
	;
	v295 = int32(base.Ui32(v278) >> (uint(int32(3)) % 32))
	goto L79
L85:
	;
	if v307 == int32(0) {
		v381 = v234
		goto L76
	} else {
		goto L86
	}
L86:
	;
	goto L77
L87:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-1)))))
	switch v359 & int32(7) {
	case 0:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L101
	default:
		v376 = int32(0)
		goto L100
	}
L88:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v320&int32(1) == int32(0) {
		v331 = int64(-1)
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if int64(0) <= v331 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L89
L91:
	;
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v250+(v320&int32(4)^int32(12)))))
	v331 = v330
	goto L90
L92:
	;
	if v337 == int32(0) {
		goto L87
	} else {
		goto L95
	}
L93:
	;
	v335 = F_commandTimeSnapshot(m)
	mBase = m.M
	v337 = base.B2i32(v331 < v335)
	goto L92
L94:
	;
	v337 = int32(0)
	goto L92
L95:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v341 != 0 {
		v381 = v234
		goto L76
	} else {
		goto L96
	}
L96:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if v343 == int32(0) {
		v381 = v234
		goto L76
	} else {
		goto L97
	}
L97:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v347 == int32(0) {
		v381 = v234
		goto L76
	} else {
		goto L98
	}
L98:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+207)))
	if v350&int32(32) == int32(0) {
		v381 = v234
		goto L76
	} else {
		goto L99
	}
L99:
	;
	goto L87
L100:
	;
	F_addReplyBulkCBuffer(m, l0, v274, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L106
	}
L101:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-17))))
	v376 = v375
	goto L100
L102:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(-9))))
	v376 = v372
	goto L100
L103:
	;
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274+int32(-5)))))
	v376 = v369
	goto L100
L104:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+int32(-3)))))
	v376 = v366
	goto L100
L105:
	;
	v376 = int32(base.Ui32(v359) >> (uint(int32(3)) % 32))
	goto L100
L106:
	;
	v381 = v234 + int32(1)
	goto L76
L107:
	;
	goto L65
L108:
	;
	v393 = v381
	goto L63
L109:
	;
	v400 = v393
	goto L62
L110:
	;
	F_setDeferredArrayLen(m, l0, v42, v400)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L7
	} else {
		goto L113
	}
L111:
	;
	F_kvstoreIteratorRelease(m, v222)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	goto L9
}
func F_keysScanCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v16 + int32(1)
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v20 == int64(9223372036854775807) {
		v40 = int32(0)
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v43&int32(2) == v40 {
			v63 = v40
		} else {
			v57 = l1 + (v43&int32(4) ^ int32(12)) + v43<<(uint(int32(3))%32)&int32(8)
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
			v63 = v57 + v58 + int32(1)
		}
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v64 == int32(0) {
			v130 = *(*int32)(unsafe.Add(mBase, _consts[131]))
			if v130 != 0 {
				v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
				switch v182 & int32(7) {
				case 0:
					v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
				case 1:
					v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
					v199 = v189
				case 2:
					v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
					v199 = v192
				case 3:
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
					v199 = v195
				case 4:
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
					v199 = v198
				default:
					v199 = int32(0)
				}
				v200 = F_vectorPush(m, v178)
				mBase = m.M
				v201 = m.ExcPending
				if v201 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
					*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
					m.G0 = v14 + int32(16)
					return
				}
			} else {
				v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v134&int32(1) == int32(0) {
					v145 = int64(-1)
				} else {
					v144 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v134&int32(4)^int32(12)))))
					v145 = v144
				}
				if int64(0) <= v145 {
					v149 = F_commandTimeSnapshot(m)
					mBase = m.M
					v151 = base.B2i32(v145 < v149)
				} else {
					v151 = int32(0)
				}
				if v151 == int32(0) {
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
					switch v182 & int32(7) {
					case 0:
						v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
					case 1:
						v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
						v199 = v189
					case 2:
						v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
						v199 = v192
					case 3:
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
						v199 = v195
					case 4:
						v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
						v199 = v198
					default:
						v199 = int32(0)
					}
					v200 = F_vectorPush(m, v178)
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
						*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
						m.G0 = v14 + int32(16)
						return
					}
				} else {
					v155 = *(*int32)(unsafe.Add(mBase, _consts[130]))
					if v155 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
						*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
						v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return
						} else {
							if v175 != 0 {
								m.G0 = v14 + int32(16)
								return
							} else {
								v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
								switch v182 & int32(7) {
								case 0:
									v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
								case 1:
									v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
									v199 = v189
								case 2:
									v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
									v199 = v192
								case 3:
									v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
									v199 = v195
								case 4:
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
									v199 = v198
								default:
									v199 = int32(0)
								}
								v200 = F_vectorPush(m, v178)
								mBase = m.M
								v201 = m.ExcPending
								if v201 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
									*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, _consts[132]))
						if v157 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
							*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return
							} else {
								if v175 != 0 {
									m.G0 = v14 + int32(16)
									return
								} else {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
									switch v182 & int32(7) {
									case 0:
										v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
									case 1:
										v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
										v199 = v189
									case 2:
										v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
										v199 = v192
									case 3:
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
										v199 = v195
									case 4:
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
										v199 = v198
									default:
										v199 = int32(0)
									}
									v200 = F_vectorPush(m, v178)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
										*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
										m.G0 = v14 + int32(16)
										return
									}
								}
							}
						} else {
							v161 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							if v161 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
								*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									if v175 != 0 {
										m.G0 = v14 + int32(16)
										return
									} else {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							} else {
								v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+207)))
								if v164&int32(32) != 0 {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
									switch v182 & int32(7) {
									case 0:
										v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
									case 1:
										v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
										v199 = v189
									case 2:
										v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
										v199 = v192
									case 3:
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
										v199 = v195
									case 4:
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
										v199 = v198
									default:
										v199 = int32(0)
									}
									v200 = F_vectorPush(m, v178)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
										*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
										m.G0 = v14 + int32(16)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
									*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										if v175 != 0 {
											m.G0 = v14 + int32(16)
											return
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
											switch v182 & int32(7) {
											case 0:
												v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
											case 1:
												v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
												v199 = v189
											case 2:
												v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
												v199 = v192
											case 3:
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
												v199 = v195
											case 4:
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
												v199 = v198
											default:
												v199 = int32(0)
											}
											v200 = F_vectorPush(m, v178)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
												*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
												m.G0 = v14 + int32(16)
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
		} else {
			v67 = int32(0)
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-1)))))
			switch v71 & int32(7) {
			case 0:
				v88 = int32(base.Ui32(v71) >> (uint(int32(3)) % 32))
			case 1:
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-3)))))
				v88 = v78
			case 2:
				v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64+int32(-5)))))
				v88 = v81
			case 3:
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-9))))
				v88 = v84
			case 4:
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-17))))
				v88 = v87
			default:
				v88 = v67
			}
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
			switch v91 & int32(7) {
			case 0:
				v108 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
			case 1:
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
				v108 = v98
			case 2:
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
				v108 = v101
			case 3:
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
				v108 = v104
			case 4:
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
				v108 = v107
			default:
				v108 = v67
			}
			v109 = int32(0)
			v111 = m.G0
			v112 = int32(16)
			v113 = v111 - v112
			m.G0 = v113
			*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v109
			v120 = F_stringmatchlen_impl(m, v64, v88, v63, v108, v109, v113+int32(12), v109)
			mBase = m.M
			m.G0 = v113 + v112
			if v120 == int32(0) {
				m.G0 = v14 + int32(16)
				return
			} else {
				v130 = *(*int32)(unsafe.Add(mBase, _consts[131]))
				if v130 != 0 {
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
					switch v182 & int32(7) {
					case 0:
						v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
					case 1:
						v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
						v199 = v189
					case 2:
						v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
						v199 = v192
					case 3:
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
						v199 = v195
					case 4:
						v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
						v199 = v198
					default:
						v199 = int32(0)
					}
					v200 = F_vectorPush(m, v178)
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
						*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
						m.G0 = v14 + int32(16)
						return
					}
				} else {
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v134&int32(1) == int32(0) {
						v145 = int64(-1)
					} else {
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v134&int32(4)^int32(12)))))
						v145 = v144
					}
					if int64(0) <= v145 {
						v149 = F_commandTimeSnapshot(m)
						mBase = m.M
						v151 = base.B2i32(v145 < v149)
					} else {
						v151 = int32(0)
					}
					if v151 == int32(0) {
						v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
						switch v182 & int32(7) {
						case 0:
							v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
						case 1:
							v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
							v199 = v189
						case 2:
							v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
							v199 = v192
						case 3:
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
							v199 = v195
						case 4:
							v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
							v199 = v198
						default:
							v199 = int32(0)
						}
						v200 = F_vectorPush(m, v178)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
							*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
							m.G0 = v14 + int32(16)
							return
						}
					} else {
						v155 = *(*int32)(unsafe.Add(mBase, _consts[130]))
						if v155 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
							*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return
							} else {
								if v175 != 0 {
									m.G0 = v14 + int32(16)
									return
								} else {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
									switch v182 & int32(7) {
									case 0:
										v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
									case 1:
										v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
										v199 = v189
									case 2:
										v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
										v199 = v192
									case 3:
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
										v199 = v195
									case 4:
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
										v199 = v198
									default:
										v199 = int32(0)
									}
									v200 = F_vectorPush(m, v178)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
										*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
										m.G0 = v14 + int32(16)
										return
									}
								}
							}
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, _consts[132]))
							if v157 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
								*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									if v175 != 0 {
										m.G0 = v14 + int32(16)
										return
									} else {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, _consts[67]))
								if v161 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
									*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										if v175 != 0 {
											m.G0 = v14 + int32(16)
											return
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
											switch v182 & int32(7) {
											case 0:
												v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
											case 1:
												v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
												v199 = v189
											case 2:
												v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
												v199 = v192
											case 3:
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
												v199 = v195
											case 4:
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
												v199 = v198
											default:
												v199 = int32(0)
											}
											v200 = F_vectorPush(m, v178)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
												*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								} else {
									v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+207)))
									if v164&int32(32) != 0 {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
										*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return
										} else {
											if v175 != 0 {
												m.G0 = v14 + int32(16)
												return
											} else {
												v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
												switch v182 & int32(7) {
												case 0:
													v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
												case 1:
													v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
													v199 = v189
												case 2:
													v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
													v199 = v192
												case 3:
													v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
													v199 = v195
												case 4:
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
													v199 = v198
												default:
													v199 = int32(0)
												}
												v200 = F_vectorPush(m, v178)
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
													*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
													m.G0 = v14 + int32(16)
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
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v25 = v23 & int32(15)
		if v25 == int32(5) {
			v30 = F_objectGetVal(m, l1)
			mBase = m.M
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
			v36 = int64(0) - int64(base.Ui64(v32)>>(uint(int64(10))%64))
		} else {
			v36 = base.I64_extend_i32_u(v25)
		}
		if v20 != v36 {
			m.G0 = v14 + int32(16)
			return
		} else {
			v40 = int32(0)
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v43&int32(2) == v40 {
				v63 = v40
			} else {
				v57 = l1 + (v43&int32(4) ^ int32(12)) + v43<<(uint(int32(3))%32)&int32(8)
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
				v63 = v57 + v58 + int32(1)
			}
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v64 == int32(0) {
				v130 = *(*int32)(unsafe.Add(mBase, _consts[131]))
				if v130 != 0 {
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
					switch v182 & int32(7) {
					case 0:
						v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
					case 1:
						v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
						v199 = v189
					case 2:
						v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
						v199 = v192
					case 3:
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
						v199 = v195
					case 4:
						v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
						v199 = v198
					default:
						v199 = int32(0)
					}
					v200 = F_vectorPush(m, v178)
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
						*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
						m.G0 = v14 + int32(16)
						return
					}
				} else {
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v134&int32(1) == int32(0) {
						v145 = int64(-1)
					} else {
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v134&int32(4)^int32(12)))))
						v145 = v144
					}
					if int64(0) <= v145 {
						v149 = F_commandTimeSnapshot(m)
						mBase = m.M
						v151 = base.B2i32(v145 < v149)
					} else {
						v151 = int32(0)
					}
					if v151 == int32(0) {
						v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
						switch v182 & int32(7) {
						case 0:
							v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
						case 1:
							v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
							v199 = v189
						case 2:
							v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
							v199 = v192
						case 3:
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
							v199 = v195
						case 4:
							v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
							v199 = v198
						default:
							v199 = int32(0)
						}
						v200 = F_vectorPush(m, v178)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
							*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
							m.G0 = v14 + int32(16)
							return
						}
					} else {
						v155 = *(*int32)(unsafe.Add(mBase, _consts[130]))
						if v155 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
							*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return
							} else {
								if v175 != 0 {
									m.G0 = v14 + int32(16)
									return
								} else {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
									switch v182 & int32(7) {
									case 0:
										v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
									case 1:
										v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
										v199 = v189
									case 2:
										v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
										v199 = v192
									case 3:
										v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
										v199 = v195
									case 4:
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
										v199 = v198
									default:
										v199 = int32(0)
									}
									v200 = F_vectorPush(m, v178)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
										*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
										m.G0 = v14 + int32(16)
										return
									}
								}
							}
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, _consts[132]))
							if v157 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
								*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									if v175 != 0 {
										m.G0 = v14 + int32(16)
										return
									} else {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, _consts[67]))
								if v161 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
									*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										if v175 != 0 {
											m.G0 = v14 + int32(16)
											return
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
											switch v182 & int32(7) {
											case 0:
												v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
											case 1:
												v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
												v199 = v189
											case 2:
												v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
												v199 = v192
											case 3:
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
												v199 = v195
											case 4:
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
												v199 = v198
											default:
												v199 = int32(0)
											}
											v200 = F_vectorPush(m, v178)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
												*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								} else {
									v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+207)))
									if v164&int32(32) != 0 {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
										*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return
										} else {
											if v175 != 0 {
												m.G0 = v14 + int32(16)
												return
											} else {
												v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
												switch v182 & int32(7) {
												case 0:
													v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
												case 1:
													v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
													v199 = v189
												case 2:
													v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
													v199 = v192
												case 3:
													v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
													v199 = v195
												case 4:
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
													v199 = v198
												default:
													v199 = int32(0)
												}
												v200 = F_vectorPush(m, v178)
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
													*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
													m.G0 = v14 + int32(16)
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
			} else {
				v67 = int32(0)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-1)))))
				switch v71 & int32(7) {
				case 0:
					v88 = int32(base.Ui32(v71) >> (uint(int32(3)) % 32))
				case 1:
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+int32(-3)))))
					v88 = v78
				case 2:
					v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64+int32(-5)))))
					v88 = v81
				case 3:
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-9))))
					v88 = v84
				case 4:
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(-17))))
					v88 = v87
				default:
					v88 = v67
				}
				v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
				switch v91 & int32(7) {
				case 0:
					v108 = int32(base.Ui32(v91) >> (uint(int32(3)) % 32))
				case 1:
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
					v108 = v98
				case 2:
					v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
					v108 = v101
				case 3:
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
					v108 = v104
				case 4:
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
					v108 = v107
				default:
					v108 = v67
				}
				v109 = int32(0)
				v111 = m.G0
				v112 = int32(16)
				v113 = v111 - v112
				m.G0 = v113
				*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v109
				v120 = F_stringmatchlen_impl(m, v64, v88, v63, v108, v109, v113+int32(12), v109)
				mBase = m.M
				m.G0 = v113 + v112
				if v120 == int32(0) {
					m.G0 = v14 + int32(16)
					return
				} else {
					v130 = *(*int32)(unsafe.Add(mBase, _consts[131]))
					if v130 != 0 {
						v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
						switch v182 & int32(7) {
						case 0:
							v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
						case 1:
							v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
							v199 = v189
						case 2:
							v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
							v199 = v192
						case 3:
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
							v199 = v195
						case 4:
							v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
							v199 = v198
						default:
							v199 = int32(0)
						}
						v200 = F_vectorPush(m, v178)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
							*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
							m.G0 = v14 + int32(16)
							return
						}
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v134&int32(1) == int32(0) {
							v145 = int64(-1)
						} else {
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l1+(v134&int32(4)^int32(12)))))
							v145 = v144
						}
						if int64(0) <= v145 {
							v149 = F_commandTimeSnapshot(m)
							mBase = m.M
							v151 = base.B2i32(v145 < v149)
						} else {
							v151 = int32(0)
						}
						if v151 == int32(0) {
							v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
							switch v182 & int32(7) {
							case 0:
								v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
							case 1:
								v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
								v199 = v189
							case 2:
								v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
								v199 = v192
							case 3:
								v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
								v199 = v195
							case 4:
								v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
								v199 = v198
							default:
								v199 = int32(0)
							}
							v200 = F_vectorPush(m, v178)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
								*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
								m.G0 = v14 + int32(16)
								return
							}
						} else {
							v155 = *(*int32)(unsafe.Add(mBase, _consts[130]))
							if v155 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
								*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									if v175 != 0 {
										m.G0 = v14 + int32(16)
										return
									} else {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
										switch v182 & int32(7) {
										case 0:
											v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
										case 1:
											v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
											v199 = v189
										case 2:
											v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
											v199 = v192
										case 3:
											v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
											v199 = v195
										case 4:
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
											v199 = v198
										default:
											v199 = int32(0)
										}
										v200 = F_vectorPush(m, v178)
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
											*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							} else {
								v157 = *(*int32)(unsafe.Add(mBase, _consts[132]))
								if v157 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
									*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										if v175 != 0 {
											m.G0 = v14 + int32(16)
											return
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
											switch v182 & int32(7) {
											case 0:
												v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
											case 1:
												v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
												v199 = v189
											case 2:
												v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
												v199 = v192
											case 3:
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
												v199 = v195
											case 4:
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
												v199 = v198
											default:
												v199 = int32(0)
											}
											v200 = F_vectorPush(m, v178)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
												*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								} else {
									v161 = *(*int32)(unsafe.Add(mBase, _consts[67]))
									if v161 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
										*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
										v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return
										} else {
											if v175 != 0 {
												m.G0 = v14 + int32(16)
												return
											} else {
												v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
												switch v182 & int32(7) {
												case 0:
													v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
												case 1:
													v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
													v199 = v189
												case 2:
													v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
													v199 = v192
												case 3:
													v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
													v199 = v195
												case 4:
													v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
													v199 = v198
												default:
													v199 = int32(0)
												}
												v200 = F_vectorPush(m, v178)
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
													*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									} else {
										v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+207)))
										if v164&int32(32) != 0 {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
											switch v182 & int32(7) {
											case 0:
												v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
											case 1:
												v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
												v199 = v189
											case 2:
												v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
												v199 = v192
											case 3:
												v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
												v199 = v195
											case 4:
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
												v199 = v198
											default:
												v199 = int32(0)
											}
											v200 = F_vectorPush(m, v178)
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
												*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
												m.G0 = v14 + int32(16)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v63
											*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-68719476736)
											v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v175 = F_expireIfNeededWithDictIndex(m, v171, v14+int32(4), l1, int32(0), l2)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												if v175 != 0 {
													m.G0 = v14 + int32(16)
													return
												} else {
													v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-1)))))
													switch v182 & int32(7) {
													case 0:
														v199 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
													case 1:
														v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+int32(-3)))))
														v199 = v189
													case 2:
														v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63+int32(-5)))))
														v199 = v192
													case 3:
														v195 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-9))))
														v199 = v195
													case 4:
														v198 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(-17))))
														v199 = v198
													default:
														v199 = int32(0)
													}
													v200 = F_vectorPush(m, v178)
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v200)+4)) = v199
														*(*int32)(unsafe.Add(mBase, uint32(v200))) = v63
														m.G0 = v14 + int32(16)
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
	}
}
func F_sortGetKeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
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
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v10 != 0 {
		v15 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a585), int32(_a550), int32(2292))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L10
	} else {
		goto L112
	}
L2:
	;
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v16 < v17 {
		v46 = v15
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v11 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = l3 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v13
	v15 = v13
	goto L2
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(73014444033)
	if l2 < int32(3) {
		v344 = v16
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v21 = l3 + int32(12)
	if v15 == v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(2)
	v46 = v42
	goto L5
L8:
	;
	v30 = F_valkey_malloc(m, int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L12
	}
L9:
	;
	v24 = F_valkey_realloc(m, v15, int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v24
	v42 = v24
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v33 == int32(0) {
		v42 = v30
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v37 = v33 << (uint(int32(3)) % 32)
	if v37 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v42 = v30
	goto L7
L15:
	;
	goto L14
L16:
	;
	v40 = F__emscripten_memcpy_bulkmem(m, v30, v21, v37)
	mBase = m.M
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v344
	return v344
L18:
	;
	v60 = int32(2)
	v63 = int32(0)
	goto L19
L19:
	;
	v66 = l1 + v60<<(uint(int32(2))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = F_objectGetVal(m, v67)
	mBase = m.M
	v69 = int32(_a586)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v344 = v335 + int32(1)
	goto L17
L21:
	;
	if v332 < l2 {
		v60 = v332
		v63 = v335
		goto L19
	} else {
		goto L111
	}
L22:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v332 = v328 + v60 + int32(1)
	v335 = v63
	goto L21
L23:
	;
	v110 = v60 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v112 = F_objectGetVal(m, v111)
	mBase = m.M
	v113 = int32(_a587)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 != 0 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	if v104-v106 != 0 {
		goto L23
	} else {
		goto L36
	}
L25:
	;
	v104 = F_tolower(m, v100)
	mBase = m.M
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v106 = F_tolower(m, v105)
	mBase = m.M
	goto L24
L26:
	;
	v74 = v68
	v75 = v69
	v76 = v72
	goto L29
L27:
	;
	v100 = int32(0)
	v101 = v69
	goto L25
L28:
	;
	v100 = v97 & int32(255)
	v101 = v96
	goto L25
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v78 == int32(0) {
		v96 = v75
		v97 = v76
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v96 = v90
	v97 = int32(0)
	goto L28
L31:
	;
	v82 = v76 & int32(255)
	if v82 == v78 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = int32(1)
	v90 = v75 + v89
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v91 != 0 {
		v74 = v74 + v89
		v75 = v90
		v76 = v91
		goto L29
	} else {
		goto L35
	}
L33:
	;
	v84 = F_tolower(m, v82)
	mBase = m.M
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v86 = F_tolower(m, v85)
	mBase = m.M
	if v84 == v86 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v96 = v75
	v97 = v88
	goto L28
L35:
	;
	goto L30
L36:
	;
	v326 = int32(_a588)
	goto L22
L37:
	;
	v332 = v110
	v335 = v325
	goto L21
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v110
	v325 = int32(1)
	goto L37
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v154 = F_objectGetVal(m, v153)
	mBase = m.M
	v155 = int32(_a186)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v158 != 0 {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	if v148-v150 != 0 {
		goto L39
	} else {
		goto L52
	}
L41:
	;
	v148 = F_tolower(m, v144)
	mBase = m.M
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v150 = F_tolower(m, v149)
	mBase = m.M
	goto L40
L42:
	;
	v118 = v112
	v119 = v113
	v120 = v116
	goto L45
L43:
	;
	v144 = int32(0)
	v145 = v113
	goto L41
L44:
	;
	v144 = v141 & int32(255)
	v145 = v140
	goto L41
L45:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v122 == int32(0) {
		v140 = v119
		v141 = v120
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v140 = v134
	v141 = int32(0)
	goto L44
L47:
	;
	v126 = v120 & int32(255)
	if v126 == v122 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v133 = int32(1)
	v134 = v119 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v135 != 0 {
		v118 = v118 + v133
		v119 = v134
		v120 = v135
		goto L45
	} else {
		goto L51
	}
L49:
	;
	v128 = F_tolower(m, v126)
	mBase = m.M
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v130 = F_tolower(m, v129)
	mBase = m.M
	if v128 == v130 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v140 = v119
	v141 = v132
	goto L44
L51:
	;
	goto L46
L52:
	;
	if v110 < l2 {
		goto L38
	} else {
		goto L53
	}
L53:
	;
	goto L39
L54:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v196 = F_objectGetVal(m, v195)
	mBase = m.M
	v197 = int32(_a587)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v200 != 0 {
		goto L71
	} else {
		goto L72
	}
L55:
	;
	if v190-v192 != 0 {
		goto L54
	} else {
		goto L67
	}
L56:
	;
	v190 = F_tolower(m, v186)
	mBase = m.M
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v192 = F_tolower(m, v191)
	mBase = m.M
	goto L55
L57:
	;
	v160 = v154
	v161 = v155
	v162 = v158
	goto L60
L58:
	;
	v186 = int32(0)
	v187 = v155
	goto L56
L59:
	;
	v186 = v183 & int32(255)
	v187 = v182
	goto L56
L60:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == int32(0) {
		v182 = v161
		v183 = v162
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v182 = v176
	v183 = int32(0)
	goto L59
L62:
	;
	v168 = v162 & int32(255)
	if v168 == v164 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v175 = int32(1)
	v176 = v161 + v175
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)))
	if v177 != 0 {
		v160 = v160 + v175
		v161 = v176
		v162 = v177
		goto L60
	} else {
		goto L66
	}
L64:
	;
	v170 = F_tolower(m, v168)
	mBase = m.M
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v172 = F_tolower(m, v171)
	mBase = m.M
	if v170 == v172 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v182 = v161
	v183 = v174
	goto L59
L66:
	;
	goto L61
L67:
	;
	v326 = int32(_a589)
	goto L22
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v238 = F_objectGetVal(m, v237)
	mBase = m.M
	v239 = int32(_a590)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v242 != 0 {
		goto L86
	} else {
		goto L87
	}
L69:
	;
	if v232-v234 != 0 {
		goto L68
	} else {
		goto L81
	}
L70:
	;
	v232 = F_tolower(m, v228)
	mBase = m.M
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v234 = F_tolower(m, v233)
	mBase = m.M
	goto L69
L71:
	;
	v202 = v196
	v203 = v197
	v204 = v200
	goto L74
L72:
	;
	v228 = int32(0)
	v229 = v197
	goto L70
L73:
	;
	v228 = v225 & int32(255)
	v229 = v224
	goto L70
L74:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v206 == int32(0) {
		v224 = v203
		v225 = v204
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v224 = v218
	v225 = int32(0)
	goto L73
L76:
	;
	v210 = v204 & int32(255)
	if v210 == v206 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v217 = int32(1)
	v218 = v203 + v217
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v219 != 0 {
		v202 = v202 + v217
		v203 = v218
		v204 = v219
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v212 = F_tolower(m, v210)
	mBase = m.M
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v214 = F_tolower(m, v213)
	mBase = m.M
	if v212 == v214 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v224 = v203
	v225 = v216
	goto L73
L80:
	;
	goto L75
L81:
	;
	if v110 < l2 {
		goto L38
	} else {
		goto L82
	}
L82:
	;
	goto L68
L83:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v280 = F_objectGetVal(m, v279)
	mBase = m.M
	v281 = int32(_a587)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v284 != 0 {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	if v274-v276 != 0 {
		goto L83
	} else {
		goto L96
	}
L85:
	;
	v274 = F_tolower(m, v270)
	mBase = m.M
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v276 = F_tolower(m, v275)
	mBase = m.M
	goto L84
L86:
	;
	v244 = v238
	v245 = v239
	v246 = v242
	goto L89
L87:
	;
	v270 = int32(0)
	v271 = v239
	goto L85
L88:
	;
	v270 = v267 & int32(255)
	v271 = v266
	goto L85
L89:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v248 == int32(0) {
		v266 = v245
		v267 = v246
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v266 = v260
	v267 = int32(0)
	goto L88
L91:
	;
	v252 = v246 & int32(255)
	if v252 == v248 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v259 = int32(1)
	v260 = v245 + v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v261 != 0 {
		v244 = v244 + v259
		v245 = v260
		v246 = v261
		goto L89
	} else {
		goto L95
	}
L93:
	;
	v254 = F_tolower(m, v252)
	mBase = m.M
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v256 = F_tolower(m, v255)
	mBase = m.M
	if v254 == v256 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v266 = v245
	v267 = v258
	goto L88
L95:
	;
	goto L90
L96:
	;
	v326 = int32(_a591)
	goto L22
L97:
	;
	if v316-v318 != 0 {
		v325 = v63
		goto L37
	} else {
		goto L109
	}
L98:
	;
	v316 = F_tolower(m, v312)
	mBase = m.M
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	v318 = F_tolower(m, v317)
	mBase = m.M
	goto L97
L99:
	;
	v286 = v280
	v287 = v281
	v288 = v284
	goto L102
L100:
	;
	v312 = int32(0)
	v313 = v281
	goto L98
L101:
	;
	v312 = v309 & int32(255)
	v313 = v308
	goto L98
L102:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v290 == int32(0) {
		v308 = v287
		v309 = v288
		goto L101
	} else {
		goto L104
	}
L103:
	;
	v308 = v302
	v309 = int32(0)
	goto L101
L104:
	;
	v294 = v288 & int32(255)
	if v294 == v290 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v301 = int32(1)
	v302 = v287 + v301
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	if v303 != 0 {
		v286 = v286 + v301
		v287 = v302
		v288 = v303
		goto L102
	} else {
		goto L108
	}
L106:
	;
	v296 = F_tolower(m, v294)
	mBase = m.M
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v298 = F_tolower(m, v297)
	mBase = m.M
	if v296 == v298 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	v308 = v287
	v309 = v300
	goto L101
L108:
	;
	goto L103
L109:
	;
	if l2 <= v110 {
		v325 = v63
		goto L37
	} else {
		goto L110
	}
L110:
	;
	goto L38
L111:
	;
	goto L20
L112:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
