package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zslDeleteNode(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v10 < int32(1) {
		v73 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v80 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v14 = l1 + int32(12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v16 != l1 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = int32(1)
	if v10 != v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v18
	goto L3
L5:
	;
	v31 = v20
	goto L7
L6:
	;
	v73 = int32(1)
	goto L1
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2+v31<<(uint(int32(2))%32))))
	v40 = v31 << (uint(int32(3)) % 32)
	v41 = v38 + v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v42 != l1 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v73 = v68
	goto L1
L9:
	;
	v67 = v31 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v67 < v68 {
		v31 = v67
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v59 = v41 + int32(16)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60 + int32(-1)
	goto L9
L11:
	;
	v45 = v41 + int32(16)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v40)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v47 + v48 + int32(-1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v14+v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+int32(12)))) = v56
	goto L9
L12:
	;
	goto L8
L13:
	;
	if v86 < int32(2) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v79
	v86 = v73
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v79
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v86 = v84
	goto L13
L16:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v118 + int32(-1)
	return
L17:
	;
	v94 = v86
	goto L18
L18:
	;
	v101 = v94 + int32(-1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v101<<(uint(int32(3))%32))))
	if v105 != 0 {
		goto L16
	} else {
		goto L20
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v101
	if base.Ui32(int32(2)) < base.Ui32(v94) {
		v94 = v101
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
}
func F_zslDeleteRangeByRank(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var __phi50 int32
	_ = __phi50
	var v52 int32
	_ = v52
	var __phi52 int32
	_ = __phi52
	var v54 int32
	_ = v54
	var __phi54 int32
	_ = __phi54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v17 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v18 < v17 {
		v113 = v17
		v114 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v120 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if v121 == v120 {
		v290 = v120
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v28 = l0
	v29 = v18
	v30 = int32(0)
	goto L4
L3:
	;
	v113 = v102 + int32(1)
	v114 = v100
	goto L1
L4:
	;
	v35 = v29 + int32(-1)
	v37 = v35 << (uint(int32(3)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28+v37+int32(12))))
	if v41 == int32(0) {
		v82 = v28
		v84 = v30
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v100 = v82
	v102 = v84
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v35<<(uint(int32(2))%32)))) = v82
	if int32(1) < v29 {
		v28 = v82
		v29 = v35
		v30 = v84
		goto L4
	} else {
		goto L16
	}
L7:
	;
	__phi50 = v28
	__phi52 = v30
	__phi54 = v41
	v50 = __phi50
	v52 = __phi52
	v54 = __phi54
	goto L8
L8:
	;
	if v29 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v82 = v54
	v84 = v71
	goto L6
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37+int32(12))))
	if v75 != 0 {
		__phi50 = v54
		__phi52 = v71
		__phi54 = v75
		v50 = __phi50
		v52 = __phi52
		v54 = __phi54
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v65 = v52 + int32(1)
	if base.Ui32(v65) < base.Ui32(l1) {
		v71 = v65
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v50+v37+int32(16))))
	v62 = v61 + v52
	if base.Ui32(l1) <= base.Ui32(v62) {
		v82 = v50
		v84 = v52
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v71 = v62
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v35<<(uint(int32(2))%32)))) = v50
	v100 = v50
	v102 = v52
	goto L3
L15:
	;
	goto L9
L16:
	;
	goto L5
L17:
	;
	m.G0 = v15 + int32(128)
	return v290
L18:
	;
	if base.Ui32(l2) < base.Ui32(v113) {
		v290 = v120
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v130 = v113
	v135 = v121
	v136 = v120
	goto L20
L20:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v144 < int32(1) {
		v207 = v144
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v290 = v273
	goto L17
L22:
	;
	v257 = v135 + int32(16)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v261 = v257 + v258<<(uint(int32(3))%32)
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v261))))
	v266 = F_hashtableDelete(m, l3, v261+v262+int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L44
	} else {
		goto L45
	}
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	if v214 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	v148 = v135 + int32(12)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	if v150 != v135 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v154 = int32(1)
	if v144 != v154 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v152
	goto L25
L27:
	;
	v165 = v154
	goto L29
L28:
	;
	v207 = int32(1)
	goto L23
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v15+v165<<(uint(int32(2))%32))))
	v174 = v165 << (uint(int32(3)) % 32)
	v175 = v172 + v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if v176 != v135 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v207 = v202
	goto L23
L31:
	;
	v201 = v165 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v201 < v202 {
		v165 = v201
		goto L29
	} else {
		goto L34
	}
L32:
	;
	v193 = v175 + int32(16)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v194 + int32(-1)
	goto L31
L33:
	;
	v179 = v175 + int32(16)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v135+int32(16)+v174)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v179))) = v181 + v182 + int32(-1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v148+v174)))
	*(*int32)(unsafe.Add(mBase, uint32(v175+int32(12)))) = v190
	goto L31
L34:
	;
	goto L30
L35:
	;
	if v220 < int32(2) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v213
	v220 = v207
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = v213
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v220 = v218
	goto L35
L38:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v252 + int32(-1)
	goto L22
L39:
	;
	v228 = v220
	goto L40
L40:
	;
	v235 = v228 + int32(-1)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v235<<(uint(int32(3))%32))))
	if v239 != 0 {
		goto L38
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v235
	if base.Ui32(int32(2)) < base.Ui32(v228) {
		v228 = v235
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	return int32(0)
L45:
	;
	F_valkey_free(m, v135)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v273 = v136 + int32(1)
	if v137 == int32(0) {
		v290 = v273
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v277 = v130 + int32(1)
	if base.Ui32(v277) <= base.Ui32(l2) {
		v130 = v277
		v135 = v137
		v136 = v273
		goto L20
	} else {
		goto L48
	}
L48:
	;
	goto L21
}
func F_zslGetElementByRank(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var __phi32 int32
	_ = __phi32
	var v36 int32
	_ = v36
	var __phi36 int32
	_ = __phi36
	var v39 int32
	_ = v39
	var __phi39 int32
	_ = __phi39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v11 <= v3 {
		v73 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v73
L2:
	;
	v15 = l0
	v18 = v11
	v19 = int32(0)
	goto L4
L3:
	;
	if v36 == l1 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v25 = v18 + int32(-1)
	v27 = v25 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15+v27)+12))
	if v29 == int32(0) {
		v55 = v15
		v59 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return v55
L6:
	;
	if v59 == l1 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	__phi32 = v15
	__phi36 = v19
	__phi39 = v29
	v32 = __phi32
	v36 = __phi36
	v39 = __phi39
	goto L8
L8:
	;
	if v25 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v55 = v39
	v59 = v52
	goto L6
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39+v27)+12))
	if v54 != 0 {
		__phi32 = v39
		__phi36 = v52
		__phi39 = v54
		v32 = __phi32
		v36 = __phi36
		v39 = __phi39
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v50 = v36 + int32(1)
	if base.Ui32(l1) < base.Ui32(v50) {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32+v27+int32(16))))
	v47 = v46 + v36
	if base.Ui32(v47) <= base.Ui32(l1) {
		v52 = v47
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v55 = v32
	v59 = v36
	goto L6
L14:
	;
	v52 = v50
	goto L10
L15:
	;
	goto L9
L16:
	;
	goto L5
L17:
	;
	if v18 < int32(2) {
		v73 = v3
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v15 = v55
	v18 = v25
	v19 = v59
	goto L4
L19:
	;
	v70 = v32
	goto L21
L20:
	;
	v70 = int32(0)
	goto L21
L21:
	;
	v73 = v70
	goto L1
}
func F_zslGetHeader(m *base.Module, l0 int32) int32 {
	return l0
}
func F_zslGetLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v2
}
func F_zslGetNodeElement(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = l0 + int32(16)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v7 = v3 + v4<<(uint(int32(3))%32)
	v8 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7))))
	return v7 + v8 + int32(1)
}
func F_zslGetRank(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v3 = int32(0)
	if l1 == v3 {
		v41 = v3
	} else {
		v10 = l1
		v11 = v3
		for {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v16 = v14 + int32(-1)
			if v16 < int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v10+v16<<(uint(int32(3))%32)+int32(12))))
				v35 = v32
				v36 = base.B2i32(v32 != int32(0))
			} else {
				v22 = v16 << (uint(int32(3)) % 32)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10+v22)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(16)+v22)))
				v35 = v24
				v36 = v26
			}
			v38 = v36 + v11
			if v35 != 0 {
				v10 = v35
				v11 = v38
				continue
			} else {
				break
			}
			break
		}
		v41 = v38
	}
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return v44 - v41
}
func F_zslGetTail(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
func F_zslLexValueLteMax(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
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
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v7 == int32(0) {
		v79 = int32(1)
		if l0 == v6 {
			v143 = v79
		} else {
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_zslLexValueLteMax[0]))
			if l0 == v82 {
				v143 = v79
			} else {
				v85 = *(*int32)(unsafe.Add(mBase, _c_F_zslLexValueLteMax[1]))
				if v6 == v85 {
					v143 = v79
				} else {
					v87 = int32(0)
					if v6 == v82 {
						v143 = v87
					} else {
						if l0 == v85 {
							v143 = v87
						} else {
							v90 = int32(0)
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							switch v97 & int32(7) {
							case 0:
								v114 = int32(base.Ui32(v97) >> (uint(int32(3)) % 32))
							case 1:
								v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
								v114 = v104
							case 2:
								v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
								v114 = v107
							case 3:
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
								v114 = v110
							case 4:
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
								v114 = v113
							default:
								v114 = v90
							}
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
							switch v117 & int32(7) {
							case 0:
								v134 = int32(base.Ui32(v117) >> (uint(int32(3)) % 32))
							case 1:
								v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
								v134 = v124
							case 2:
								v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
								v134 = v127
							case 3:
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
								v134 = v130
							case 4:
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
								v134 = v133
							default:
								v134 = v90
							}
							v135 = base.B2i32(base.Ui32(v114) < base.Ui32(v134))
							if base.Ui32(v114) < base.Ui32(v134) {
								v136 = v114
							} else {
								v136 = v134
							}
							v137 = F_memcmp(m, l0, v6, v136)
							mBase = m.M
							if v137 != 0 {
								v140 = v137
							} else {
								v140 = base.B2i32(base.Ui32(v134) < base.Ui32(v114)) - v135
							}
							v143 = base.B2i32(v140 < int32(1))
						}
					}
				}
			}
		}
		return v143
	} else {
		if l0 != v6 {
			v13 = int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_zslLexValueLteMax[0]))
			if l0 == v15 {
				v143 = v13
				return v143
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_zslLexValueLteMax[1]))
				if v6 == v18 {
					v143 = v13
					return v143
				} else {
					if v6 != v15 {
						if l0 == v18 {
							v143 = int32(0)
							return v143
						} else {
							v25 = int32(0)
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
							switch v32 & int32(7) {
							case 0:
								v49 = int32(base.Ui32(v32) >> (uint(int32(3)) % 32))
							case 1:
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
								v49 = v39
							case 2:
								v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
								v49 = v42
							case 3:
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
								v49 = v45
							case 4:
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
								v49 = v48
							default:
								v49 = v25
							}
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-1)))))
							switch v52 & int32(7) {
							case 0:
								v69 = int32(base.Ui32(v52) >> (uint(int32(3)) % 32))
							case 1:
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6+int32(-3)))))
								v69 = v59
							case 2:
								v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6+int32(-5)))))
								v69 = v62
							case 3:
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-9))))
								v69 = v65
							case 4:
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(-17))))
								v69 = v68
							default:
								v69 = v25
							}
							v70 = base.B2i32(base.Ui32(v49) < base.Ui32(v69))
							if base.Ui32(v49) < base.Ui32(v69) {
								v71 = v49
							} else {
								v71 = v69
							}
							v72 = F_memcmp(m, l0, v6, v71)
							mBase = m.M
							if v72 != 0 {
								v75 = v72
							} else {
								v75 = base.B2i32(base.Ui32(v69) < base.Ui32(v49)) - v70
							}
							return int32(base.Ui32(v75) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(0)
					}
				}
			}
		} else {
			return int32(0)
		}
	}
}
func F_zslParseLexRangeItem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v6 = int32(-1)
	v7 = F_objectGetVal(m, l0)
	mBase = m.M
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	switch v8 + int32(-40) {
	case 0:
		v21 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
		switch v28 & int32(7) {
		case 0:
			v45 = int32(base.Ui32(v28) >> (uint(int32(3)) % 32))
		case 1:
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
			v45 = v35
		case 2:
			v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
			v45 = v38
		case 3:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
			v45 = v41
		case 4:
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
			v45 = v44
		default:
			v45 = int32(0)
		}
		v48 = F_sdsnewlen(m, v7+v21, v45+int32(-1))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v84 = v48
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
			v89 = int32(0)
			return v89
		}
	case 1, 2, 4:
		v89 = v6
		return v89
	case 3:
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
		if v11 != 0 {
			v89 = v6
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_zslParseLexRangeItem[0]))
			v84 = v15
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
			v89 = int32(0)
		}
		return v89
	case 5:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
		if v16 != 0 {
			v89 = v6
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_zslParseLexRangeItem[1]))
			v84 = v20
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
			v89 = int32(0)
		}
		return v89
	default:
		if v8 != int32(91) {
			v89 = v6
			return v89
		} else {
			v54 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-1)))))
			switch v61 & int32(7) {
			case 0:
				v78 = int32(base.Ui32(v61) >> (uint(int32(3)) % 32))
			case 1:
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+int32(-3)))))
				v78 = v68
			case 2:
				v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+int32(-5)))))
				v78 = v71
			case 3:
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-9))))
				v78 = v74
			case 4:
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(-17))))
				v78 = v77
			default:
				v78 = v54
			}
			v81 = F_sdsnewlen(m, v7+int32(1), v78+int32(-1))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v84 = v81
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
				v89 = int32(0)
				return v89
			}
		}
	}
}
func F_zslParseRange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int64
	_ = v66
	var v71 int64
	_ = v71
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 float64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int64
	_ = v121
	var v126 int64
	_ = v126
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 float64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
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
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int64
	_ = v215
	var v220 int64
	_ = v220
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 float64
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v270 int64
	_ = v270
	var v275 int64
	_ = v275
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 float64
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = int64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_objectGetVal(m, l0)
	mBase = m.M
	if v14&int32(240) != v9 {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-1)))))
		switch v25 & int32(7) {
		case 0:
			v42 = int32(base.Ui32(v25) >> (uint(int32(3)) % 32))
		case 1:
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(-3)))))
			v42 = v32
		case 2:
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+int32(-5)))))
			v42 = v35
		case 3:
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-9))))
			v42 = v38
		case 4:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(-17))))
			v42 = v41
		default:
			v42 = int32(0)
		}
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v43 != int32(40) {
			v107 = v10 + int32(12)
			v108 = int32(0)
			v111 = m.G0
			v113 = v111 - int32(32)
			m.G0 = v113
			v115 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v115))) = v108
			v121 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v113+int32(8)))) = v121
			*(*int64)(unsafe.Add(mBase, uint32(v113)+24)) = int64(0)
			v126 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v113))) = v126
			F_ffc_from_chars_double_options(m, v113+int32(16), v15, v15+v42, v113+int32(24), v113)
			mBase = m.M
			v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
			if v134 == v108 {
			} else {
				if v134 == int32(2) {
					v141 = int32(68)
				} else {
					v141 = int32(28)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v115))) = v141
			}
			if v107 == int32(0) {
			} else {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v107))) = v145
			}
			v147 = *(*float64)(unsafe.Add(mBase, uint32(v113)+24))
			m.G0 = v113 + int32(32)
			*(*float64)(unsafe.Add(mBase, uint32(l2))) = v147
			v152 = int32(-1)
			v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
			if v154 != 0 {
				v315 = v152
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v147)&int64(9223372036854775807)) {
					v315 = v152
				} else {
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v164 = F_objectGetVal(m, l1)
					mBase = m.M
					if v163&int32(240) != int32(16) {
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-1)))))
						switch v174 & int32(7) {
						case 0:
							v191 = int32(base.Ui32(v174) >> (uint(int32(3)) % 32))
						case 1:
							v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-3)))))
							v191 = v181
						case 2:
							v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+int32(-5)))))
							v191 = v184
						case 3:
							v187 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-9))))
							v191 = v187
						case 4:
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-17))))
							v191 = v190
						default:
							v191 = int32(0)
						}
						v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
						if v192 != int32(40) {
							v256 = v10 + int32(12)
							v257 = int32(0)
							v260 = m.G0
							v262 = v260 - int32(32)
							m.G0 = v262
							v264 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v264))) = v257
							v270 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
							*(*int64)(unsafe.Add(mBase, uint32(v262+int32(8)))) = v270
							*(*int64)(unsafe.Add(mBase, uint32(v262)+24)) = int64(0)
							v275 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v262))) = v275
							F_ffc_from_chars_double_options(m, v262+int32(16), v164, v164+v191, v262+int32(24), v262)
							mBase = m.M
							v283 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
							if v283 == v257 {
							} else {
								if v283 == int32(2) {
									v290 = int32(68)
								} else {
									v290 = int32(28)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v264))) = v290
							}
							if v256 == int32(0) {
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v256))) = v294
							}
							v296 = *(*float64)(unsafe.Add(mBase, uint32(v262)+24))
							m.G0 = v262 + int32(32)
							*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v296
							v301 = int32(-1)
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
							if v303 != 0 {
								v315 = v301
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v296)&int64(9223372036854775807)) {
									v315 = v301
								} else {
									v315 = int32(0)
								}
							}
						} else {
							v195 = int32(-1)
							v197 = v164 + int32(1)
							v201 = v10 + int32(12)
							v202 = int32(0)
							v205 = m.G0
							v207 = v205 - int32(32)
							m.G0 = v207
							v209 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v209))) = v202
							v215 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
							*(*int64)(unsafe.Add(mBase, uint32(v207+int32(8)))) = v215
							*(*int64)(unsafe.Add(mBase, uint32(v207)+24)) = int64(0)
							v220 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v207))) = v220
							F_ffc_from_chars_double_options(m, v207+int32(16), v197, v197+(v191+v195), v207+int32(24), v207)
							mBase = m.M
							v228 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
							if v228 == v202 {
							} else {
								if v228 == int32(2) {
									v235 = int32(68)
								} else {
									v235 = int32(28)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v209))) = v235
							}
							if v201 == int32(0) {
							} else {
								v239 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v201))) = v239
							}
							v241 = *(*float64)(unsafe.Add(mBase, uint32(v207)+24))
							m.G0 = v207 + int32(32)
							*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v241
							v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
							if v247 != 0 {
								v315 = v195
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v241)&int64(9223372036854775807)) {
									v315 = v195
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(1)
									v315 = int32(0)
								}
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_convert_i32_s(v164)
						v315 = int32(0)
					}
				}
			}
		} else {
			v46 = int32(-1)
			v48 = v15 + int32(1)
			v52 = v10 + int32(12)
			v53 = int32(0)
			v56 = m.G0
			v58 = v56 - int32(32)
			m.G0 = v58
			v60 = F___errno_location(m)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v60))) = v53
			v66 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v58+int32(8)))) = v66
			*(*int64)(unsafe.Add(mBase, uint32(v58)+24)) = int64(0)
			v71 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v58))) = v71
			F_ffc_from_chars_double_options(m, v58+int32(16), v48, v48+(v42+v46), v58+int32(24), v58)
			mBase = m.M
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
			if v79 == v53 {
			} else {
				if v79 == int32(2) {
					v86 = int32(68)
				} else {
					v86 = int32(28)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v60))) = v86
			}
			if v52 == int32(0) {
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v52))) = v90
			}
			v92 = *(*float64)(unsafe.Add(mBase, uint32(v58)+24))
			m.G0 = v58 + int32(32)
			*(*float64)(unsafe.Add(mBase, uint32(l2))) = v92
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
			if v98 != 0 {
				v315 = v46
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v92)&int64(9223372036854775807)) {
					v315 = v46
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(1)
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v164 = F_objectGetVal(m, l1)
					mBase = m.M
					if v163&int32(240) != int32(16) {
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-1)))))
						switch v174 & int32(7) {
						case 0:
							v191 = int32(base.Ui32(v174) >> (uint(int32(3)) % 32))
						case 1:
							v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-3)))))
							v191 = v181
						case 2:
							v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+int32(-5)))))
							v191 = v184
						case 3:
							v187 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-9))))
							v191 = v187
						case 4:
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-17))))
							v191 = v190
						default:
							v191 = int32(0)
						}
						v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
						if v192 != int32(40) {
							v256 = v10 + int32(12)
							v257 = int32(0)
							v260 = m.G0
							v262 = v260 - int32(32)
							m.G0 = v262
							v264 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v264))) = v257
							v270 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
							*(*int64)(unsafe.Add(mBase, uint32(v262+int32(8)))) = v270
							*(*int64)(unsafe.Add(mBase, uint32(v262)+24)) = int64(0)
							v275 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v262))) = v275
							F_ffc_from_chars_double_options(m, v262+int32(16), v164, v164+v191, v262+int32(24), v262)
							mBase = m.M
							v283 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
							if v283 == v257 {
							} else {
								if v283 == int32(2) {
									v290 = int32(68)
								} else {
									v290 = int32(28)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v264))) = v290
							}
							if v256 == int32(0) {
							} else {
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v256))) = v294
							}
							v296 = *(*float64)(unsafe.Add(mBase, uint32(v262)+24))
							m.G0 = v262 + int32(32)
							*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v296
							v301 = int32(-1)
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
							if v303 != 0 {
								v315 = v301
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v296)&int64(9223372036854775807)) {
									v315 = v301
								} else {
									v315 = int32(0)
								}
							}
						} else {
							v195 = int32(-1)
							v197 = v164 + int32(1)
							v201 = v10 + int32(12)
							v202 = int32(0)
							v205 = m.G0
							v207 = v205 - int32(32)
							m.G0 = v207
							v209 = F___errno_location(m)
							mBase = m.M
							*(*int32)(unsafe.Add(mBase, uint32(v209))) = v202
							v215 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
							*(*int64)(unsafe.Add(mBase, uint32(v207+int32(8)))) = v215
							*(*int64)(unsafe.Add(mBase, uint32(v207)+24)) = int64(0)
							v220 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v207))) = v220
							F_ffc_from_chars_double_options(m, v207+int32(16), v197, v197+(v191+v195), v207+int32(24), v207)
							mBase = m.M
							v228 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
							if v228 == v202 {
							} else {
								if v228 == int32(2) {
									v235 = int32(68)
								} else {
									v235 = int32(28)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v209))) = v235
							}
							if v201 == int32(0) {
							} else {
								v239 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v201))) = v239
							}
							v241 = *(*float64)(unsafe.Add(mBase, uint32(v207)+24))
							m.G0 = v207 + int32(32)
							*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v241
							v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
							if v247 != 0 {
								v315 = v195
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v241)&int64(9223372036854775807)) {
									v315 = v195
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(1)
									v315 = int32(0)
								}
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_convert_i32_s(v164)
						v315 = int32(0)
					}
				}
			}
		}
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_convert_i32_s(v15)
		v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v164 = F_objectGetVal(m, l1)
		mBase = m.M
		if v163&int32(240) != int32(16) {
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-1)))))
			switch v174 & int32(7) {
			case 0:
				v191 = int32(base.Ui32(v174) >> (uint(int32(3)) % 32))
			case 1:
				v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+int32(-3)))))
				v191 = v181
			case 2:
				v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+int32(-5)))))
				v191 = v184
			case 3:
				v187 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-9))))
				v191 = v187
			case 4:
				v190 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(-17))))
				v191 = v190
			default:
				v191 = int32(0)
			}
			v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
			if v192 != int32(40) {
				v256 = v10 + int32(12)
				v257 = int32(0)
				v260 = m.G0
				v262 = v260 - int32(32)
				m.G0 = v262
				v264 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v264))) = v257
				v270 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v262+int32(8)))) = v270
				*(*int64)(unsafe.Add(mBase, uint32(v262)+24)) = int64(0)
				v275 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v262))) = v275
				F_ffc_from_chars_double_options(m, v262+int32(16), v164, v164+v191, v262+int32(24), v262)
				mBase = m.M
				v283 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
				if v283 == v257 {
				} else {
					if v283 == int32(2) {
						v290 = int32(68)
					} else {
						v290 = int32(28)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v264))) = v290
				}
				if v256 == int32(0) {
				} else {
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v256))) = v294
				}
				v296 = *(*float64)(unsafe.Add(mBase, uint32(v262)+24))
				m.G0 = v262 + int32(32)
				*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v296
				v301 = int32(-1)
				v302 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
				if v303 != 0 {
					v315 = v301
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v296)&int64(9223372036854775807)) {
						v315 = v301
					} else {
						v315 = int32(0)
					}
				}
			} else {
				v195 = int32(-1)
				v197 = v164 + int32(1)
				v201 = v10 + int32(12)
				v202 = int32(0)
				v205 = m.G0
				v207 = v205 - int32(32)
				m.G0 = v207
				v209 = F___errno_location(m)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v209))) = v202
				v215 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v207+int32(8)))) = v215
				*(*int64)(unsafe.Add(mBase, uint32(v207)+24)) = int64(0)
				v220 = *(*int64)(unsafe.Add(mBase, _c_F_zslParseRange[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v207))) = v220
				F_ffc_from_chars_double_options(m, v207+int32(16), v197, v197+(v191+v195), v207+int32(24), v207)
				mBase = m.M
				v228 = *(*int32)(unsafe.Add(mBase, uint32(v207)+20))
				if v228 == v202 {
				} else {
					if v228 == int32(2) {
						v235 = int32(68)
					} else {
						v235 = int32(28)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v209))) = v235
				}
				if v201 == int32(0) {
				} else {
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v201))) = v239
				}
				v241 = *(*float64)(unsafe.Add(mBase, uint32(v207)+24))
				m.G0 = v207 + int32(32)
				*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = v241
				v246 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
				if v247 != 0 {
					v315 = v195
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v241)&int64(9223372036854775807)) {
						v315 = v195
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(1)
						v315 = int32(0)
					}
				}
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(l2)+8)) = base.F64_convert_i32_s(v164)
			v315 = int32(0)
		}
	}
	m.G0 = v10 + int32(16)
	return v315
}
func F_zslUpdateScore(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v174 int32
	_ = v174
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
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
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v347 int32
	_ = v347
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F__serverAssert(m, int32(_a_F_zslUpdateScore_0), int32(_a_F_zslUpdateScore_1), int32(384))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L70
	} else {
		goto L72
	}
L2:
	;
	m.G0 = v16 + int32(128)
	return v330
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v35 < int32(1) {
		v193 = l0
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	if base.F64_lt(v21, l2) == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = l2
	v330 = int32(0)
	goto L2
L8:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v25)))
	if base.F64_gt(v28, l2) == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	if v201 != l1 {
		goto L1
	} else {
		goto L47
	}
L11:
	;
	v39 = l1 + int32(16)
	v45 = l0
	v46 = v35
	goto L12
L12:
	;
	v54 = v46 + int32(-1)
	v56 = v54 << (uint(int32(3)) % 32)
	v59 = v45 + v56 + int32(12)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 == l1 {
		v174 = v45
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v193 = v174
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v54<<(uint(int32(2))%32)))) = v174
	if int32(1) < v46 {
		v45 = v174
		v46 = v54
		goto L12
	} else {
		goto L46
	}
L15:
	;
	if v60 == int32(0) {
		v174 = v45
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
	v65 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_gt(v64, v65) != 0 {
		v174 = v45
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v71 = v60
	v72 = v45
	v77 = v59
	v78 = v64
	v79 = v65
	goto L18
L18:
	;
	if base.F64_lt(v78, v79) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v174 = v156
	goto L14
L20:
	;
	v159 = v156 + v56 + int32(12)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v160 == l1 {
		v174 = v156
		goto L14
	} else {
		goto L43
	}
L21:
	;
	v84 = v71 + int32(16)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(3)
	v88 = v84 + v85<<(uint(v86)%32)
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88))))
	v90 = v88 + v89
	v91 = int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v96 = v39 + v93<<(uint(v86)%32)
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	v98 = v96 + v97
	v101 = int32(0)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v101))))
	switch v108 & int32(7) {
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
		v125 = v101
		goto L24
	}
L22:
	;
	v156 = v71
	goto L20
L23:
	;
	if int32(-1) < v151 {
		v174 = v72
		goto L14
	} else {
		goto L42
	}
L24:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(0)))))
	switch v128 & int32(7) {
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
		v145 = v101
		goto L30
	}
L25:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-16))))
	v125 = v124
	goto L24
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v90+int32(-8))))
	v125 = v121
	goto L24
L27:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90+int32(-4)))))
	v125 = v118
	goto L24
L28:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+int32(-2)))))
	v125 = v115
	goto L24
L29:
	;
	v125 = int32(base.Ui32(v108) >> (uint(int32(3)) % 32))
	goto L24
L30:
	;
	v146 = base.B2i32(base.Ui32(v125) < base.Ui32(v145))
	if base.Ui32(v125) < base.Ui32(v145) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-16))))
	v145 = v144
	goto L30
L32:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(-8))))
	v145 = v141
	goto L30
L33:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98+int32(-4)))))
	v145 = v138
	goto L30
L34:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(-2)))))
	v145 = v135
	goto L30
L35:
	;
	v145 = int32(base.Ui32(v128) >> (uint(int32(3)) % 32))
	goto L30
L36:
	;
	v147 = v125
	goto L38
L37:
	;
	v147 = v145
	goto L38
L38:
	;
	v148 = F_memcmp(m, v90+v91, v98+v91, v147)
	mBase = m.M
	if v148 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v151 = v148
	goto L41
L40:
	;
	v151 = base.B2i32(base.Ui32(v145) < base.Ui32(v125)) - v146
	goto L41
L41:
	;
	goto L23
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v156 = v154
	goto L20
L43:
	;
	if v160 == int32(0) {
		v174 = v156
		goto L14
	} else {
		goto L44
	}
L44:
	;
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v160)))
	v165 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_gt(v164, v165) == int32(0) {
		v71 = v160
		v72 = v156
		v77 = v159
		v78 = v164
		v79 = v165
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L19
L46:
	;
	goto L13
L47:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v209 < int32(1) {
		v272 = v209
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = l2
	v322 = F_zslInsertNode(m, l0, l1)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L70
	} else {
		goto L71
	}
L49:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v279 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v213 = l1 + int32(12)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	if v215 != l1 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v219 = int32(1)
	if v209 != v219 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+12)) = v217
	goto L51
L53:
	;
	v230 = v219
	goto L55
L54:
	;
	v272 = int32(1)
	goto L49
L55:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16+v230<<(uint(int32(2))%32))))
	v239 = v230 << (uint(int32(3)) % 32)
	v240 = v237 + v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	if v241 != l1 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v272 = v267
	goto L49
L57:
	;
	v266 = v230 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v266 < v267 {
		v230 = v266
		goto L55
	} else {
		goto L60
	}
L58:
	;
	v258 = v240 + int32(16)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v259 + int32(-1)
	goto L57
L59:
	;
	v244 = v240 + int32(16)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v239)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v246 + v247 + int32(-1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v213+v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v240+int32(12)))) = v255
	goto L57
L60:
	;
	goto L56
L61:
	;
	if v285 < int32(2) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v278
	v285 = v272
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+8)) = v278
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v285 = v283
	goto L61
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v317 + int32(-1)
	goto L48
L65:
	;
	v293 = v285
	goto L66
L66:
	;
	v300 = v293 + int32(-1)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(12)+v300<<(uint(int32(3))%32))))
	if v304 != 0 {
		goto L64
	} else {
		goto L68
	}
L67:
	;
	goto L64
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v300
	if base.Ui32(int32(2)) < base.Ui32(v293) {
		v293 = v300
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	return int32(0)
L71:
	;
	v330 = v322
	goto L2
L72:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
