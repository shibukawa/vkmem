package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___ofl_lock(m *base.Module) int32 {
	return int32(9116556)
}
func F_opendir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v2 = int32(0)
	v6 = F_open(m, l0, int32(589824), v2)
	mBase = m.M
	if v6 < v2 {
		v46 = v2
		return v46
	} else {
		v17 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(2072)))
		v30 = F_emscripten_builtin_malloc(m, v17)
		mBase = m.M
		if v30 == int32(0) {
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(-4)))))
			if v35&int32(3) == int32(0) {
			} else {
				v41 = F___memset(m, v30, int32(0), v17)
				mBase = m.M
			}
		}
		if v30 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v6
			v46 = v30
			return v46
		} else {
			v42 = m.Wasi_snapshot_preview1.Fd_close(m, v6)
			mBase = m.M
			return int32(0)
		}
	}
}
func F_openlog(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == int32(0) {
		v32 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[1057])) = uint8(v32)
	} else {
		v18 = int32(31)
		v21 = F_memchr(m, l0, int32(0), v18)
		mBase = m.M
		if v21 != 0 {
			v23 = v21 - l0
		} else {
			v23 = v18
		}
		if v23 == int32(0) {
		} else {
			v26 = F__emscripten_memcpy_bulkmem(m, int32(9128384), l0, v23)
			mBase = m.M
		}
		v30 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[1057]))) = uint8(v30)
	}
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1058])) = l2
	*(*int32)(unsafe.Add(mBase, _consts[1059])) = l1
	if l1&int32(8) == v36 {
	} else {
		v45 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
		if int32(-1) < v45 {
		} else {
			v48 = int32(0)
			v53 = F_socket(m, int32(1), int32(524290), v48)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[1060])) = v53
			if v53 < v48 {
			} else {
				v59 = F_connect(m, v53, int32(_a2772), int32(12))
				mBase = m.M
			}
		}
	}
	m.G0 = v8 + int32(16)
	return
}
func F_out(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4&int32(32) != 0 {
		return
	} else {
		v7 = F___fwritex(m, l1, l2, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_overMaxmemoryAfterAlloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	v2 = int32(0)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if v8 == int64(0) {
		v276 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v276
L2:
	;
	v11 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v20 < int32(261) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v105 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	if base.Ui64(base.I64_extend_i32_u(v97+l0)) <= base.Ui64(v105) {
		v276 = v2
		goto L1
	} else {
		goto L19
	}
L4:
	;
	goto L3
L5:
	;
	v31 = v29 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v29) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	if v20 < int32(1) {
		v97 = v11
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v28 = v24
	v29 = int32(260)
	goto L5
L8:
	;
	v28 = v11
	v29 = v20
	goto L5
L9:
	;
	if v31 == int32(0) {
		v97 = v70
		goto L4
	} else {
		goto L15
	}
L10:
	;
	v38 = int32(0)
	v40 = v28
	v41 = v38
	v45 = v38
	goto L12
L11:
	;
	v70 = v28
	v71 = int32(0)
	goto L9
L12:
	;
	v48 = v41 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[317])))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[318])))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[319])))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[320])))
	v64 = v51 + (v54 + (v57 + (v60 + v40)))
	v65 = int32(4)
	v66 = v41 + v65
	v68 = v45 + v65
	if v68 != v29&int32(2147483644) {
		v40 = v64
		v41 = v66
		v45 = v68
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v70 = v64
	v71 = v66
	goto L9
L14:
	;
	goto L13
L15:
	;
	v79 = v70
	v80 = v71
	v82 = int32(0)
	goto L16
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80<<(uint(int32(2))%32))+uint32(_consts[320])))
	v91 = v90 + v79
	v92 = int32(1)
	v95 = v82 + v92
	if v95 != v31 {
		v79 = v91
		v80 = v80 + v92
		v82 = v95
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v97 = v91
	goto L4
L18:
	;
	goto L17
L19:
	;
	v110 = int32(_a20)
	v111 = *(*int64)(unsafe.Add(mBase, _consts[322]))
	v113 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	if base.I64_extend_i32_u(v113) <= v111 {
		v128 = int32(0)
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[38]))
	if v130 == int32(0) {
		v176 = v128
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v118 = base.I64_div_s(v111, int64(16384))
	v125 = v113 - base.I32_wrap_i64(v111+v118*int64(44)) + int32(-44)
	if base.Ui32(v113) < base.Ui32(v125) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v127 = int32(0)
	goto L24
L23:
	;
	v127 = v125
	goto L24
L24:
	;
	v128 = v127
	goto L20
L25:
	;
	v177 = int32(0)
	v180 = m.G0
	v182 = v180 - int32(16)
	m.G0 = v182
	v186 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	if v186 == v177 {
		v218 = v177
		goto L42
	} else {
		goto L43
	}
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	v141 = v134 + int32(-1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v144 = v142 & int32(7)
	switch v144 {
	case 0:
		goto L33
	case 1:
		v150 = int32(4)
		goto L28
	case 2:
		goto L32
	case 3:
		goto L31
	case 4:
		goto L30
	default:
		goto L29
	}
L27:
	;
	v176 = v174 + v128
	goto L25
L28:
	;
	switch v144 {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	default:
		v170 = int32(0)
		goto L34
	}
L29:
	;
	v150 = int32(1)
	goto L28
L30:
	;
	v150 = int32(18)
	goto L28
L31:
	;
	v150 = int32(10)
	goto L28
L32:
	;
	v150 = int32(6)
	goto L28
L33:
	;
	v145 = F_zmalloc_usable_size(m, v141)
	mBase = m.M
	v174 = v145
	goto L27
L34:
	;
	v174 = v150 + v170
	goto L27
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-9))))
	v170 = v169
	goto L34
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(-5))))
	v174 = v150 + v165
	goto L27
L37:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134+int32(-3)))))
	v174 = v150 + v161
	goto L27
L38:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+int32(-2)))))
	v174 = v150 + v157
	goto L27
L39:
	;
	v174 = v150 + int32(base.Ui32(v142)>>(uint(int32(3))%32))
	goto L27
L40:
	;
	v268 = *(*int64)(unsafe.Add(mBase, _consts[321]))
	v270 = v97 - v266
	if base.Ui32(v97) < base.Ui32(v270) {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	if v218 == int32(0) {
		v266 = v176
		goto L40
	} else {
		goto L52
	}
L42:
	;
	m.G0 = v182 + int32(16)
	goto L41
L43:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v190 == int32(0) {
		v218 = v177
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[171])))
	v195 = v182 + int32(8)
	F_listRewind(m, v193, v195)
	mBase = m.M
	v197 = int32(0)
	v200 = F_listNext(m, v195)
	mBase = m.M
	if v200 == v197 {
		v218 = v197
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v205 = v200
	goto L46
L46:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v207 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v218 = v197
	goto L42
L48:
	;
	v216 = F_listNext(m, v182+int32(8))
	mBase = m.M
	if v216 != 0 {
		v205 = v216
		goto L46
	} else {
		goto L51
	}
L49:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)+156))
	if base.Ui32(v208+int32(-18)) <= base.Ui32(int32(2)) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v218 = int32(1)
	goto L42
L51:
	;
	goto L47
L52:
	;
	v225 = int32(0)
	v228 = m.G0
	v230 = v228 - int32(16)
	m.G0 = v230
	v233 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+uint32(_consts[171])))
	v236 = v230 + int32(8)
	F_listRewind(m, v234, v236)
	mBase = m.M
	v241 = F_listNext(m, v236)
	mBase = m.M
	if v241 == v225 {
		v260 = v225
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v266 = v260 + v176
	goto L40
L54:
	;
	m.G0 = v230 + int32(16)
	goto L53
L55:
	;
	v245 = v225
	v246 = v241
	goto L56
L56:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v248 != 0 {
		v254 = v245
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v260 = v254
	goto L54
L58:
	;
	v258 = F_listNext(m, v230+int32(8))
	mBase = m.M
	if v258 != 0 {
		v245 = v254
		v246 = v258
		goto L56
	} else {
		goto L61
	}
L59:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v247)+152))
	if v249 == int32(0) {
		v254 = v245
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v252 = F_getClientOutputBufferMemoryUsage(m, v249)
	mBase = m.M
	v254 = v252 + v245
	goto L58
L61:
	;
	goto L57
L62:
	;
	v272 = int32(0)
	goto L64
L63:
	;
	v272 = v270
	goto L64
L64:
	;
	v276 = base.B2i32(base.Ui64(v268) < base.Ui64(base.I64_extend_i32_u(v272+l0)))
	goto L1
}
