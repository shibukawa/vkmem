package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___procfdname(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v9 = int32(0)
	for {
		v12 = l0 + v9
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1008]))))
		*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v15)
		if v9 != int32(14) {
			v9 = v9 + int32(1)
			continue
		} else {
			break
		}
		break
	}
	if l1 == int32(0) {
		v56 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v56)
		v58 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)) = uint8(v58)
		return
	} else {
		v26 = int32(14)
		v28 = l1
		for {
			v30 = v26 + int32(1)
			v34 = base.I32_div_u_s(v28, int32(10))
			if base.Ui32(int32(9)) < base.Ui32(v28) {
				v26 = v30
				v28 = v34
				continue
			} else {
				break
			}
			break
		}
		v36 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+v30))) = uint8(v36)
		v39 = l1
		v40 = v30
		for {
			v44 = v40 + int32(-1)
			v46 = int32(10)
			v47 = base.I32_div_u_s(v39, v46)
			v52 = v39 - v47*v46 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l0+v44))) = uint8(v52)
			if base.Ui32(int32(9)) < base.Ui32(v39) {
				v39 = v47
				v40 = v44
				continue
			} else {
				break
			}
			break
		}
		return
	}
}
func F_parseExtendedCommandArgumentsOrReply(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	if l6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 <= l2 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(-1)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	v30 = base.B2i32(l1 != int32(3))
	v32 = base.B2i32(base.Ui32(int32(4)) < base.Ui32(l1))
	v35 = l2
	goto L5
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v55 = v35 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v55)))
	v58 = F_objectGetVal(m, v57)
	mBase = m.M
	if v35 == l3+int32(-1) {
		v65 = int32(0)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v68 = v66 + int32(-70)
	switch v68 {
	case 0, 32:
		goto L13
	case 1, 33:
		goto L11
	default:
		goto L10
	case 3, 35:
		goto L12
	case 8, 40:
		goto L15
	case 18, 50:
		goto L14
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v55+int32(4))))
	v65 = v64
	goto L7
L9:
	;
	v464 = v458 + int32(1)
	if v464 < l3 {
		v35 = v464
		goto L5
	} else {
		goto L137
	}
L10:
	;
	v214 = int32(_a2321)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v217 != 0 {
		goto L60
	} else {
		goto L61
	}
L11:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v192|int32(32) != int32(101) {
		goto L10
	} else {
		goto L53
	}
L12:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v155|int32(32) != int32(102) {
		goto L43
	} else {
		goto L44
	}
L13:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v115 = v113 | int32(32)
	if v115 != int32(110) {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v91|int32(32) != int32(120) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v69|int32(32) != int32(120) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	switch v66 + int32(-88) {
	case 0:
		goto L14
	default:
		goto L10
	case 14:
		goto L13
	case 15:
		goto L11
	case 17:
		goto L12
	}
L17:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v74 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v75&int32(514) != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(int32(4)) < base.Ui32(l1) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if int32(1)<<(uint(l1)%32)&int32(26) == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v75 | int32(1)
	v458 = v35
	goto L9
L22:
	;
	switch v66 + int32(-102) {
	case 0:
		goto L13
	case 1:
		goto L11
	default:
		goto L10
	case 3:
		goto L12
	}
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v96 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v97&int32(513) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(int32(4)) < base.Ui32(l1) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if int32(1)<<(uint(l1)%32)&int32(26) == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v97 | int32(2)
	v458 = v35
	goto L9
L28:
	;
	switch v68 {
	case 0, 32:
		goto L35
	case 1:
		goto L11
	default:
		goto L10
	case 3:
		goto L12
	}
L29:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v118|int32(32) != int32(120) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v123 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v124&int32(4096) != 0 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if l1 != int32(3) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	if v124&int32(512) != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v124 | int32(2048)
	v458 = v35
	goto L9
L35:
	;
	if v115 != int32(120) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	switch v66 + int32(-71) {
	case 0:
		goto L11
	default:
		goto L10
	case 2:
		goto L12
	}
L37:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v135|int32(32) != int32(120) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v140 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v141&int32(2048) != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	if l1 != int32(3) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	if v141&int32(512) != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v141 | int32(4096)
	v458 = v35
	goto L9
L43:
	;
	if v66 != int32(103) {
		goto L10
	} else {
		goto L52
	}
L44:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v160|int32(32) != int32(101) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v165|int32(32) != int32(113) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v170 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v65 == int32(0) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v173&int32(3) != 0 {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	if l1 != int32(1) {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	if v173&int32(512) != 0 {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v173 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v65
	v458 = v35 + int32(1)
	goto L9
L52:
	;
	goto L11
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v197|int32(32) != int32(116) {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	if l1 != int32(1) {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v204&int32(255) != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v207 | int32(32)
	v458 = v35
	goto L9
L57:
	;
	v266 = int32(_a2322)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v269 != 0 {
		goto L77
	} else {
		goto L78
	}
L58:
	;
	if v249-v251 != 0 {
		goto L57
	} else {
		goto L70
	}
L59:
	;
	v249 = F_tolower(m, v245)
	mBase = m.M
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v251 = F_tolower(m, v250)
	mBase = m.M
	goto L58
L60:
	;
	v219 = v58
	v220 = v214
	v221 = v217
	goto L63
L61:
	;
	v245 = int32(0)
	v246 = v214
	goto L59
L62:
	;
	v245 = v242 & int32(255)
	v246 = v241
	goto L59
L63:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v223 == int32(0) {
		v241 = v220
		v242 = v221
		goto L62
	} else {
		goto L65
	}
L64:
	;
	v241 = v235
	v242 = int32(0)
	goto L62
L65:
	;
	v227 = v221 & int32(255)
	if v227 == v223 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v234 = int32(1)
	v235 = v220 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	if v236 != 0 {
		v219 = v219 + v234
		v220 = v235
		v221 = v236
		goto L63
	} else {
		goto L69
	}
L67:
	;
	v229 = F_tolower(m, v227)
	mBase = m.M
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	v231 = F_tolower(m, v230)
	mBase = m.M
	if v229 == v231 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v241 = v220
	v242 = v233
	goto L62
L69:
	;
	goto L64
L70:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v253&int32(460) != 0 {
		goto L57
	} else {
		goto L71
	}
L71:
	;
	if base.Ui32(int32(4)) < base.Ui32(l1) {
		goto L57
	} else {
		goto L72
	}
L72:
	;
	if int32(1)<<(uint(l1)%32)&int32(26) == int32(0) {
		goto L57
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v253 | int32(16)
	v458 = v35
	goto L9
L74:
	;
	v314 = v66 + int32(-69)
	switch v314 {
	case 0:
		goto L94
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		goto L89
	case 11:
		goto L93
	default:
		goto L95
	}
L75:
	;
	if v301-v303|l1&int32(-3) != 0 {
		goto L74
	} else {
		goto L87
	}
L76:
	;
	v301 = F_tolower(m, v297)
	mBase = m.M
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	v303 = F_tolower(m, v302)
	mBase = m.M
	goto L75
L77:
	;
	v271 = v58
	v272 = v266
	v273 = v269
	goto L80
L78:
	;
	v297 = int32(0)
	v298 = v266
	goto L76
L79:
	;
	v297 = v294 & int32(255)
	v298 = v293
	goto L76
L80:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v275 == int32(0) {
		v293 = v272
		v294 = v273
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v293 = v287
	v294 = int32(0)
	goto L79
L82:
	;
	v279 = v273 & int32(255)
	if v279 == v275 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v286 = int32(1)
	v287 = v272 + v286
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	if v288 != 0 {
		v271 = v271 + v286
		v272 = v287
		v273 = v288
		goto L80
	} else {
		goto L86
	}
L84:
	;
	v281 = F_tolower(m, v279)
	mBase = m.M
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v283 = F_tolower(m, v282)
	mBase = m.M
	if v281 == v283 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v293 = v272
	v294 = v285
	goto L79
L86:
	;
	goto L81
L87:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v306&int32(220) != 0 {
		goto L74
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v306 | int32(256)
	v458 = v35
	goto L9
L89:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	F_addReplyErrorObject(m, l0, v451)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L135
	} else {
		goto L136
	}
L90:
	;
	if v410|int32(32) != int32(120) {
		goto L89
	} else {
		goto L126
	}
L91:
	;
	if v373|int32(32) != int32(120) {
		goto L115
	} else {
		goto L116
	}
L92:
	;
	if v66 != int32(101) {
		goto L89
	} else {
		goto L114
	}
L93:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v342|int32(32) != int32(120) {
		goto L104
	} else {
		goto L105
	}
L94:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v317|int32(32) != int32(120) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	switch v66 + int32(-101) {
	case 0:
		goto L94
	default:
		goto L89
	case 11:
		goto L93
	}
L96:
	;
	switch v314 {
	case 0:
		v373 = v317
		goto L91
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		goto L89
	case 11:
		goto L93
	default:
		goto L92
	}
L97:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v322 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v323&int32(344) != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	if v323&int32(128) != 0 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	if v65 == int32(0) {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v323 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
	if l6 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v458 = v35 + int32(1)
	goto L9
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v35
	goto L102
L104:
	;
	switch v66 + int32(-101) {
	case 0:
		v373 = v342
		goto L91
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		goto L89
	case 11:
		v410 = v342
		goto L90
	default:
		goto L112
	}
L105:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v347 != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v348&int32(340) != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if v348&int32(128) != 0 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	if v65 == int32(0) {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v348 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
	if l6 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v458 = v35 + int32(1)
	goto L9
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v35
	goto L110
L112:
	;
	if v66 == int32(80) {
		v410 = v342
		goto L90
	} else {
		goto L113
	}
L113:
	;
	goto L89
L114:
	;
	v373 = v317
	goto L91
L115:
	;
	if v66 != int32(80) {
		goto L89
	} else {
		goto L125
	}
L116:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v380|int32(32) != int32(97) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v385|int32(32) != int32(116) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v390 != 0 {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v391&int32(284) != 0 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	if v391&int32(128) != 0 {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	if v65 == int32(0) {
		goto L115
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v391 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
	if l6 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v458 = v35 + int32(1)
	goto L9
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v35
	goto L123
L125:
	;
	v410 = v373
	goto L90
L126:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	if v417|int32(32) != int32(97) {
		goto L89
	} else {
		goto L127
	}
L127:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v422|int32(32) != int32(116) {
		goto L89
	} else {
		goto L128
	}
L128:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v427 != 0 {
		goto L89
	} else {
		goto L129
	}
L129:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v428&int32(340) != 0 {
		goto L89
	} else {
		goto L130
	}
L130:
	;
	if v428&int32(8) != 0 {
		goto L89
	} else {
		goto L131
	}
L131:
	;
	if v65 == int32(0) {
		goto L89
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v428 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
	if l6 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v458 = v35 + int32(1)
	goto L9
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v35
	goto L133
L135:
	;
	return int32(0)
L136:
	;
	return int32(-1)
L137:
	;
	goto L6
}
func F_parseLoadexArguments(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 < int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v189
L2:
	;
	v120 = int32(_a1596)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v123 != 0 {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	v115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v115
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v115
	v189 = v115
	goto L1
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(0)
	goto L5
L5:
	;
	v28 = v16 + v23<<(uint(int32(2))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = F_objectGetVal(m, v29)
	mBase = m.M
	v31 = int32(_a1597)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L3
L7:
	;
	if v66-v68 != 0 {
		goto L2
	} else {
		goto L19
	}
L8:
	;
	v66 = F_tolower(m, v62)
	mBase = m.M
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = F_tolower(m, v67)
	mBase = m.M
	goto L7
L9:
	;
	v36 = v30
	v37 = v31
	v38 = v34
	goto L12
L10:
	;
	v62 = int32(0)
	v63 = v31
	goto L8
L11:
	;
	v62 = v59 & int32(255)
	v63 = v58
	goto L8
L12:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 == int32(0) {
		v58 = v37
		v59 = v38
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v58 = v52
	v59 = int32(0)
	goto L11
L14:
	;
	v44 = v38 & int32(255)
	if v44 == v40 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = int32(1)
	v52 = v37 + v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v53 != 0 {
		v36 = v36 + v51
		v37 = v52
		v38 = v53
		goto L12
	} else {
		goto L18
	}
L16:
	;
	v46 = F_tolower(m, v44)
	mBase = m.M
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v48 = F_tolower(m, v47)
	mBase = m.M
	if v46 == v48 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v58 = v37
	v59 = v50
	goto L11
L18:
	;
	goto L13
L19:
	;
	v71 = v23 + int32(2)
	if v71 < v13 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(4))))
	v88 = F_objectGetVal(m, v87)
	mBase = m.M
	v89 = F_sdsdup(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L23
	} else {
		goto L25
	}
L21:
	;
	v73 = int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v75 {
		v189 = v73
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F__serverLog(m, int32(2), int32(_a1598), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v189 = v73
	goto L1
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v16+v71<<(uint(int32(2))%32))))
	v95 = F_objectGetVal(m, v94)
	mBase = m.M
	v96 = F_sdsdup(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v100 = F_dictReplace(m, v99, v89, v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v105 = v23 + int32(3)
	if v105 < v13 {
		v23 = v105
		goto L5
	} else {
		goto L31
	}
L28:
	;
	if v100 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_sdsfree(m, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L6
L32:
	;
	v173 = int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	if int32(2) < v175 {
		v189 = v173
		goto L1
	} else {
		goto L52
	}
L33:
	;
	if v155-v157 != 0 {
		goto L32
	} else {
		goto L45
	}
L34:
	;
	v155 = F_tolower(m, v151)
	mBase = m.M
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	v157 = F_tolower(m, v156)
	mBase = m.M
	goto L33
L35:
	;
	v125 = v30
	v126 = v120
	v127 = v123
	goto L38
L36:
	;
	v151 = int32(0)
	v152 = v120
	goto L34
L37:
	;
	v151 = v148 & int32(255)
	v152 = v147
	goto L34
L38:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v129 == int32(0) {
		v147 = v126
		v148 = v127
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v147 = v141
	v148 = int32(0)
	goto L37
L40:
	;
	v133 = v127 & int32(255)
	if v133 == v129 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v140 = int32(1)
	v141 = v126 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
	if v142 != 0 {
		v125 = v125 + v140
		v126 = v141
		v127 = v142
		goto L38
	} else {
		goto L44
	}
L42:
	;
	v135 = F_tolower(m, v133)
	mBase = m.M
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v137 = F_tolower(m, v136)
	mBase = m.M
	if v135 == v137 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v147 = v126
	v148 = v139
	goto L37
L44:
	;
	goto L39
L45:
	;
	v159 = int32(0)
	v161 = v23 + int32(1)
	v166 = base.B2i32(v161 < v13)
	if v161 < v13 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v167 = v16 + v161<<(uint(int32(2))%32)
	goto L48
L47:
	;
	v167 = v159
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v167
	if v161 < v13 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v171 = v13 - v161
	goto L51
L50:
	;
	v171 = int32(0)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v171
	v189 = v159
	goto L1
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v30
	F__serverLog(m, int32(2), int32(_a1599), v11)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v189 = v173
	goto L1
}
func F_parseMultibulk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v262 int64
	_ = v262
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v284 int64
	_ = v284
	var v289 int64
	_ = v289
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v297 int32
	_ = v297
	var v305 int64
	_ = v305
	var v329 int64
	_ = v329
	var v348 int32
	_ = v348
	var v357 int64
	_ = v357
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int64
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v675 int64
	_ = v675
	var v681 int32
	_ = v681
	var v683 int64
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v697 int64
	_ = v697
	var v702 int64
	_ = v702
	var v706 int32
	_ = v706
	var v708 int64
	_ = v708
	var v710 int32
	_ = v710
	var v718 int64
	_ = v718
	var v742 int64
	_ = v742
	var v761 int32
	_ = v761
	var v770 int64
	_ = v770
	var v774 int64
	_ = v774
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
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
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v955 int32
	_ = v955
	var v958 int64
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int64
	_ = v999
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1220 int64
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1244 int32
	_ = v1244
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v26 = v24 & int32(65536)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v27 != 0 {
		v405 = v27
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v22 + int32(16)
	return v1244
L2:
	;
	v1220 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v1220 + base.I64_extend_i32_u(v1221<<(uint(int32(1))%32)+v1224)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	v1244 = int32(8192)
	goto L1
L3:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1690), int32(_a1630), int32(3786))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L94
	} else {
		goto L299
	}
L4:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1691), int32(_a1630), int32(3728))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L94
	} else {
		goto L298
	}
L5:
	;
	F__serverAssertWithInfo(m, l0, int32(0), int32(_a1692), int32(_a1630), int32(3707))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L94
	} else {
		goto L297
	}
L6:
	;
	if v405 <= int32(0) {
		goto L3
	} else {
		goto L100
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v28 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v31 = v29 + v30
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-1)))))
	v37 = v35 & int32(7)
	switch v37 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		v52 = int32(0)
		goto L9
	}
L9:
	;
	v54 = v52 - v30
	v55 = int32(0)
	v58 = base.B2i32(v54 != v55)
	if v31&int32(3) == v55 {
		v84 = v31
		v86 = v54
		v87 = v58
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
	v52 = v51
	goto L9
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v52 = v48
	goto L9
L12:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
	v52 = v45
	goto L9
L13:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v52 = v42
	goto L9
L14:
	;
	v52 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	goto L9
L15:
	;
	v180 = int32(0)
	switch v37 {
	case 0:
		goto L53
	case 1:
		goto L52
	case 2:
		goto L51
	case 3:
		goto L50
	case 4:
		goto L49
	default:
		v196 = v180
		goto L48
	}
L16:
	;
	if v157 != 0 {
		goto L15
	} else {
		goto L41
	}
L17:
	;
	v157 = int32(0)
	goto L16
L18:
	;
	v135 = v128
	v137 = v130
	goto L36
L19:
	;
	if v87 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L20:
	;
	if v54 == int32(0) {
		v84 = v31
		v86 = v54
		v87 = v58
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = v31
	v69 = v54
	goto L22
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(13) {
		v128 = v67
		v130 = v69
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v84 = v79
	v86 = v75
	v87 = v77
	goto L19
L24:
	;
	v75 = v69 + int32(-1)
	v76 = int32(0)
	v77 = base.B2i32(v75 != v76)
	v79 = v67 + int32(1)
	if v79&int32(3) == v76 {
		v84 = v79
		v86 = v75
		v87 = v77
		goto L19
	} else {
		goto L25
	}
L25:
	;
	if v75 != 0 {
		v67 = v79
		v69 = v75
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v91 == int32(13) {
		v121 = v84
		v123 = v86
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v123 == int32(0) {
		goto L17
	} else {
		goto L35
	}
L29:
	;
	if base.Ui32(v86) < base.Ui32(int32(4)) {
		v121 = v84
		v123 = v86
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v101 = v84
	v103 = v86
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v108 = v107 ^ int32(218959117)
	v111 = int32(-2139062144)
	if (int32(16843008)-v108|v108)&v111 != v111 {
		v128 = v101
		v130 = v103
		goto L18
	} else {
		goto L33
	}
L32:
	;
	v121 = v116
	v123 = v118
	goto L28
L33:
	;
	v116 = v101 + int32(4)
	v118 = v103 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v118) {
		v101 = v116
		v103 = v118
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v128 = v121
	v130 = v123
	goto L18
L36:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v140 != int32(13) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L17
L38:
	;
	v145 = v137 + int32(-1)
	if v145 != 0 {
		v135 = v135 + int32(1)
		v137 = v145
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v157 = v135
	goto L16
L40:
	;
	goto L37
L41:
	;
	switch v37 {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v173 = int32(0)
		goto L42
	}
L42:
	;
	v1244 = base.B2i32(base.Ui32(int32(65536)) < base.Ui32(v173-v30)) << (uint(int32(2)) % 32)
	goto L1
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
	v173 = v172
	goto L42
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v173 = v169
	goto L42
L45:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
	v173 = v166
	goto L42
L46:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v173 = v163
	goto L42
L47:
	;
	v173 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	goto L42
L48:
	;
	if v196-v30+int32(-2) < v157-v31 {
		v1244 = v180
		goto L1
	} else {
		goto L54
	}
L49:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-17))))
	v196 = v195
	goto L48
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(-9))))
	v196 = v192
	goto L48
L51:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(-5)))))
	v196 = v189
	goto L48
L52:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+int32(-3)))))
	v196 = v186
	goto L48
L53:
	;
	v196 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
	goto L48
L54:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v201 == int32(10) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v205 != int32(42) {
		goto L4
	} else {
		goto L57
	}
L56:
	;
	v1244 = int32(4194304)
	goto L1
L57:
	;
	v208 = int32(8)
	v210 = v31 + int32(1)
	v211 = v157 - v210
	v213 = v22 + v208
	v214 = int32(0)
	if base.Ui32(v211+int32(-21)) < base.Ui32(int32(-20)) {
		v348 = v214
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v348 == int32(0) {
		v1244 = v208
		goto L1
	} else {
		goto L85
	}
L59:
	;
	goto L58
L60:
	;
	v226 = int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v211 != v226 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v348 = int32(1)
	goto L59
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v329
	goto L61
L63:
	;
	if v227&int32(255) == int32(45) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v231 = v227 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v231&int32(255)) {
		v348 = v214
		goto L59
	} else {
		goto L65
	}
L65:
	;
	if v213 == int32(0) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v329 = base.I64_extend_i32_u(v231) & int64(255)
	goto L62
L67:
	;
	if base.Ui32(int32(8)) < base.Ui32((v250+int32(-49))&int32(255)) {
		v348 = v214
		goto L59
	} else {
		goto L70
	}
L68:
	;
	v1255 = int32(2)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	v249 = v1255
	v250 = v247
	v251 = v31 + v1255
	goto L67
L69:
	;
	v249 = v226
	v250 = v227
	v251 = v210
	goto L67
L70:
	;
	v262 = base.I64_extend_i32_u(v250+int32(-48)) & int64(255)
	if base.Ui32(v211) <= base.Ui32(v249) {
		v305 = v262
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v227&int32(255) != int32(45) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v268 = v249
	v270 = v262
	v272 = v251
	goto L73
L73:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	if base.Ui32((v274+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v348 = v214
		goto L59
	} else {
		goto L75
	}
L74:
	;
	v305 = v295
	goto L71
L75:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v270) {
		v348 = v214
		goto L59
	} else {
		goto L76
	}
L76:
	;
	v284 = v270 * int64(10)
	v289 = base.I64_extend_i32_u(v274+int32(-48)) & int64(255)
	if base.Ui64(v289^int64(-1)) < base.Ui64(v284) {
		v348 = v214
		goto L59
	} else {
		goto L77
	}
L77:
	;
	v293 = int32(1)
	v295 = v284 + v289
	v297 = v268 + v293
	if v297 != v211 {
		v268 = v297
		v270 = v295
		v272 = v272 + v293
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	if v305 < int64(0) {
		v348 = v214
		goto L59
	} else {
		goto L83
	}
L80:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v305) {
		v348 = v214
		goto L59
	} else {
		goto L81
	}
L81:
	;
	if v213 == int32(0) {
		goto L61
	} else {
		goto L82
	}
L82:
	;
	v329 = int64(0) - v305
	goto L62
L83:
	;
	if v213 == int32(0) {
		goto L61
	} else {
		goto L84
	}
L84:
	;
	v329 = v305
	goto L62
L85:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	if int64(2147483647) < v357 {
		v1244 = v208
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v357 < int64(11) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v157 - v365 + int32(2)
	if int64(1) <= v357 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	if v26 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v1244 = int32(16)
	goto L1
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(-1)
	v375 = base.I32_wrap_i64(v357)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v377 == int32(0) {
		v385 = v375
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v1244 = int32(4096)
	goto L1
L92:
	;
	v386 = int32(1024)
	if v385 < v386 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	F_valkey_free(m, v377)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	return int32(0)
L95:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v385 = v384
	goto L92
L96:
	;
	v389 = v385
	goto L98
L97:
	;
	v389 = v386
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v389
	v393 = F_valkey_malloc(m, v389<<(uint(int32(2))%32))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L94
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v393
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v398 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v398 + base.I64_extend_i32_u(v211+int32(3))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v405 = v404
	goto L6
L100:
	;
	v418 = v24 & int32(16384)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v427 = v419
	goto L101
L101:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v427 != int32(-1) {
		v1005 = v427
		v1006 = v439
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1189 == int32(0) {
		goto L2
	} else {
		goto L296
	}
L103:
	;
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006+int32(-1)))))
	switch v1018 & int32(7) {
	case 0:
		goto L259
	case 1:
		goto L258
	case 2:
		goto L257
	case 3:
		goto L256
	case 4:
		goto L255
	default:
		v1035 = int32(0)
		goto L254
	}
L104:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v443 = v439 + v442
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+int32(-1)))))
	v449 = v447 & int32(7)
	switch v449 {
	case 0:
		goto L110
	case 1:
		goto L109
	case 2:
		goto L108
	case 3:
		goto L107
	case 4:
		goto L106
	default:
		v464 = int32(0)
		goto L105
	}
L105:
	;
	v466 = v464 - v442
	v467 = int32(0)
	v470 = base.B2i32(v466 != v467)
	if v443&int32(3) == v467 {
		v496 = v443
		v498 = v466
		v499 = v470
		goto L115
	} else {
		goto L116
	}
L106:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-17))))
	v464 = v463
	goto L105
L107:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-9))))
	v464 = v460
	goto L105
L108:
	;
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+int32(-5)))))
	v464 = v457
	goto L105
L109:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+int32(-3)))))
	v464 = v454
	goto L105
L110:
	;
	v464 = int32(base.Ui32(v447) >> (uint(int32(3)) % 32))
	goto L105
L111:
	;
	v592 = int32(0)
	switch v449 {
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
		v608 = v592
		goto L144
	}
L112:
	;
	if v569 != 0 {
		goto L111
	} else {
		goto L137
	}
L113:
	;
	v569 = int32(0)
	goto L112
L114:
	;
	v547 = v540
	v549 = v542
	goto L132
L115:
	;
	if v499 == int32(0) {
		goto L113
	} else {
		goto L123
	}
L116:
	;
	if v466 == int32(0) {
		v496 = v443
		v498 = v466
		v499 = v470
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v479 = v443
	v481 = v466
	goto L118
L118:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v484 == int32(13) {
		v540 = v479
		v542 = v481
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v496 = v491
	v498 = v487
	v499 = v489
	goto L115
L120:
	;
	v487 = v481 + int32(-1)
	v488 = int32(0)
	v489 = base.B2i32(v487 != v488)
	v491 = v479 + int32(1)
	if v491&int32(3) == v488 {
		v496 = v491
		v498 = v487
		v499 = v489
		goto L115
	} else {
		goto L121
	}
L121:
	;
	if v487 != 0 {
		v479 = v491
		v481 = v487
		goto L118
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v503 == int32(13) {
		v533 = v496
		v535 = v498
		goto L124
	} else {
		goto L125
	}
L124:
	;
	if v535 == int32(0) {
		goto L113
	} else {
		goto L131
	}
L125:
	;
	if base.Ui32(v498) < base.Ui32(int32(4)) {
		v533 = v496
		v535 = v498
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v513 = v496
	v515 = v498
	goto L127
L127:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v520 = v519 ^ int32(218959117)
	v523 = int32(-2139062144)
	if (int32(16843008)-v520|v520)&v523 != v523 {
		v540 = v513
		v542 = v515
		goto L114
	} else {
		goto L129
	}
L128:
	;
	v533 = v528
	v535 = v530
	goto L124
L129:
	;
	v528 = v513 + int32(4)
	v530 = v515 + int32(-4)
	if base.Ui32(int32(3)) < base.Ui32(v530) {
		v513 = v528
		v515 = v530
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v540 = v533
	v542 = v535
	goto L114
L132:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	if v552 != int32(13) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L113
L134:
	;
	v557 = v549 + int32(-1)
	if v557 != 0 {
		v547 = v547 + int32(1)
		v549 = v557
		goto L132
	} else {
		goto L136
	}
L135:
	;
	v569 = v547
	goto L112
L136:
	;
	goto L133
L137:
	;
	switch v449 {
	case 0:
		goto L143
	case 1:
		goto L142
	case 2:
		goto L141
	case 3:
		goto L140
	case 4:
		goto L139
	default:
		v585 = int32(0)
		goto L138
	}
L138:
	;
	v1244 = base.B2i32(base.Ui32(int32(65536)) < base.Ui32(v585-v442)) << (uint(int32(6)) % 32)
	goto L1
L139:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-17))))
	v585 = v584
	goto L138
L140:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-9))))
	v585 = v581
	goto L138
L141:
	;
	v578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+int32(-5)))))
	v585 = v578
	goto L138
L142:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+int32(-3)))))
	v585 = v575
	goto L138
L143:
	;
	v585 = int32(base.Ui32(v447) >> (uint(int32(3)) % 32))
	goto L138
L144:
	;
	if v608-v442+int32(-2) < v569-v443 {
		v1244 = v592
		goto L1
	} else {
		goto L150
	}
L145:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-17))))
	v608 = v607
	goto L144
L146:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v439+int32(-9))))
	v608 = v604
	goto L144
L147:
	;
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439+int32(-5)))))
	v608 = v601
	goto L144
L148:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439+int32(-3)))))
	v608 = v598
	goto L144
L149:
	;
	v608 = int32(base.Ui32(v447) >> (uint(int32(3)) % 32))
	goto L144
L150:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v613 == int32(36) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)))
	if v617 == int32(10) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v1244 = int32(128)
	goto L1
L153:
	;
	v621 = int32(256)
	v623 = v443 + int32(1)
	v624 = v569 - v623
	v626 = v22 + int32(8)
	v627 = int32(0)
	if base.Ui32(v624+int32(-21)) < base.Ui32(int32(-20)) {
		v761 = v627
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v1244 = int32(4194304)
	goto L1
L155:
	;
	if v761 == int32(0) {
		v1244 = v621
		goto L1
	} else {
		goto L182
	}
L156:
	;
	goto L155
L157:
	;
	v639 = int32(1)
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
	if v624 != v639 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v761 = int32(1)
	goto L156
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v626))) = v742
	goto L158
L160:
	;
	if v640&int32(255) == int32(45) {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	v644 = v640 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v644&int32(255)) {
		v761 = v627
		goto L156
	} else {
		goto L162
	}
L162:
	;
	if v626 == int32(0) {
		goto L158
	} else {
		goto L163
	}
L163:
	;
	v742 = base.I64_extend_i32_u(v644) & int64(255)
	goto L159
L164:
	;
	if base.Ui32(int32(8)) < base.Ui32((v663+int32(-49))&int32(255)) {
		v761 = v627
		goto L156
	} else {
		goto L167
	}
L165:
	;
	v1256 = int32(2)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+1)))
	v662 = v1256
	v663 = v660
	v664 = v443 + v1256
	goto L164
L166:
	;
	v662 = v639
	v663 = v640
	v664 = v623
	goto L164
L167:
	;
	v675 = base.I64_extend_i32_u(v663+int32(-48)) & int64(255)
	if base.Ui32(v624) <= base.Ui32(v662) {
		v718 = v675
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if v640&int32(255) != int32(45) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v681 = v662
	v683 = v675
	v685 = v664
	goto L170
L170:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)))
	if base.Ui32((v687+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v761 = v627
		goto L156
	} else {
		goto L172
	}
L171:
	;
	v718 = v708
	goto L168
L172:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v683) {
		v761 = v627
		goto L156
	} else {
		goto L173
	}
L173:
	;
	v697 = v683 * int64(10)
	v702 = base.I64_extend_i32_u(v687+int32(-48)) & int64(255)
	if base.Ui64(v702^int64(-1)) < base.Ui64(v697) {
		v761 = v627
		goto L156
	} else {
		goto L174
	}
L174:
	;
	v706 = int32(1)
	v708 = v697 + v702
	v710 = v681 + v706
	if v710 != v624 {
		v681 = v710
		v683 = v708
		v685 = v685 + v706
		goto L170
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	if v718 < int64(0) {
		v761 = v627
		goto L156
	} else {
		goto L180
	}
L177:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v718) {
		v761 = v627
		goto L156
	} else {
		goto L178
	}
L178:
	;
	if v626 == int32(0) {
		goto L158
	} else {
		goto L179
	}
L179:
	;
	v742 = int64(0) - v718
	goto L159
L180:
	;
	if v626 == int32(0) {
		goto L158
	} else {
		goto L181
	}
L181:
	;
	v742 = v718
	goto L159
L182:
	;
	v770 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	if v770 < int64(0) {
		v1244 = v621
		goto L1
	} else {
		goto L183
	}
L183:
	;
	if v418 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if base.Ui64(v770) < base.Ui64(int64(16385)) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	v774 = *(*int64)(unsafe.Add(mBase, _consts[493]))
	if v774 < v770 {
		v1244 = v621
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v784 = v569 - v781 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v784
	if v418 != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	if v26 == int32(0) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v1244 = int32(32)
	goto L1
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v993
	v999 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(l5))) = v999 + base.I64_extend_i32_u(v624+int32(3))
	v1005 = v993
	v1006 = v994
	goto L103
L191:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781+int32(-1)))))
	switch v792 & int32(7) {
	case 0:
		goto L200
	case 1:
		goto L199
	case 2:
		goto L198
	case 3:
		goto L197
	case 4:
		goto L196
	default:
		v809 = int32(0)
		goto L195
	}
L192:
	;
	v993 = base.I32_wrap_i64(v770)
	v994 = v781
	goto L190
L193:
	;
	if base.Ui64(int64(32767)) < base.Ui64(v770) {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v811 = base.I32_wrap_i64(v770)
	if base.Ui32(v811+int32(2)) < base.Ui32(v809-v784) {
		v993 = v811
		v994 = v781
		goto L190
	} else {
		goto L201
	}
L196:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v781+int32(-17))))
	v809 = v808
	goto L195
L197:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v781+int32(-9))))
	v809 = v805
	goto L195
L198:
	;
	v802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v781+int32(-5)))))
	v809 = v802
	goto L195
L199:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781+int32(-3)))))
	v809 = v799
	goto L195
L200:
	;
	v809 = int32(base.Ui32(v792) >> (uint(int32(3)) % 32))
	goto L195
L201:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _consts[492]))
	if v781 != v816 {
		v852 = v781
		v853 = v784
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v854 = int32(-1)
	v862 = v852 + v854
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	v865 = v863 & int32(7)
	switch v865 {
	case 0:
		goto L219
	case 1:
		goto L218
	case 2:
		goto L217
	case 3:
		goto L216
	case 4:
		goto L215
	default:
		goto L213
	}
L203:
	;
	v818 = int32(0)
	v821 = F_sdsnewlen(m, v818, int32(16384))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L94
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, _consts[492])) = v821
	v826 = v821 + int32(-1)
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	switch v827 & int32(7) {
	case 0:
		goto L211
	case 1:
		goto L210
	case 2:
		goto L209
	case 3:
		goto L208
	case 4:
		goto L207
	default:
		goto L206
	}
L205:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v852 = v851
	v853 = v850
	goto L202
L206:
	;
	v848 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v821))) = uint8(v848)
	goto L205
L207:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v821+int32(-17)))) = int64(0)
	goto L206
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v821+int32(-9)))) = int32(0)
	goto L206
L209:
	;
	v838 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v821+int32(-5)))) = uint16(v838)
	goto L206
L210:
	;
	v834 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v821+int32(-3)))) = uint8(v834)
	goto L206
L211:
	;
	v830 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v830)
	goto L206
L212:
	;
	v955 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v955
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959+int32(-1)))))
	switch v962 & int32(7) {
	case 0:
		goto L251
	case 1:
		goto L250
	case 2:
		goto L249
	case 3:
		goto L248
	case 4:
		goto L247
	default:
		v979 = v955
		goto L246
	}
L213:
	;
	goto L212
L214:
	;
	if v880 == int32(0) {
		goto L213
	} else {
		goto L220
	}
L215:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v852+int32(-17))))
	v880 = v879
	goto L214
L216:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v852+int32(-9))))
	v880 = v876
	goto L214
L217:
	;
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v852+int32(-5)))))
	v880 = v873
	goto L214
L218:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852+int32(-3)))))
	v880 = v870
	goto L214
L219:
	;
	v880 = int32(base.Ui32(v863) >> (uint(int32(3)) % 32))
	goto L214
L220:
	;
	v886 = int32(-1)&v880 + v854
	v890 = v853>>(uint(int32(31))%32)&v880 + v853
	v893 = v886 - v890 + int32(1)
	switch v865 {
	default:
		goto L226
	case 1:
		goto L225
	case 2:
		goto L224
	case 3:
		goto L223
	case 4:
		goto L222
	}
L221:
	;
	v909 = int32(0)
	v911 = base.B2i32(base.Ui32(v890) < base.Ui32(v908))
	if base.Ui32(v890) < base.Ui32(v908) {
		goto L228
	} else {
		goto L229
	}
L222:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v852+int32(-17))))
	v908 = v907
	goto L221
L223:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v852+int32(-9))))
	v908 = v904
	goto L221
L224:
	;
	v901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v852+int32(-5)))))
	v908 = v901
	goto L221
L225:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852+int32(-3)))))
	v908 = v898
	goto L221
L226:
	;
	v908 = int32(base.Ui32(v863) >> (uint(int32(3)) % 32))
	goto L221
L227:
	;
	v925 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v852+v919))) = uint8(v925)
	switch v865 {
	default:
		goto L245
	case 1:
		goto L244
	case 2:
		goto L243
	case 3:
		goto L242
	case 4:
		goto L241
	}
L228:
	;
	v912 = v890
	goto L230
L229:
	;
	v912 = v909
	goto L230
L230:
	;
	v913 = v908 - v912
	if base.Ui32(v893) < base.Ui32(v913) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v915 = v893
	goto L233
L232:
	;
	v915 = v913
	goto L233
L233:
	;
	if v886 < v890 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v917 = v909
	goto L236
L235:
	;
	v917 = v915
	goto L236
L236:
	;
	if base.Ui32(v890) < base.Ui32(v908) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v919 = v917
	goto L239
L238:
	;
	v919 = int32(0)
	goto L239
L239:
	;
	if v919 == int32(0) {
		goto L227
	} else {
		goto L240
	}
L240:
	;
	v923 = F_memmove(m, v852, v852+v912, v919)
	mBase = m.M
	goto L227
L241:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v852+int32(-17)))) = base.I64_extend_i32_u(v919)
	goto L213
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v852+int32(-9)))) = v919
	goto L212
L243:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v852+int32(-5)))) = uint16(v919)
	goto L212
L244:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v852+int32(-3)))) = uint8(v919)
	goto L212
L245:
	;
	v928 = v919 << (uint(int32(3)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v862))) = uint8(v928)
	goto L212
L246:
	;
	v984 = F_sdsMakeRoomForNonGreedy(m, v959, base.I32_wrap_i64(v958)-v979+int32(2))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L94
	} else {
		goto L252
	}
L247:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v959+int32(-17))))
	v979 = v978
	goto L246
L248:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v959+int32(-9))))
	v979 = v975
	goto L246
L249:
	;
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v959+int32(-5)))))
	v979 = v972
	goto L246
L250:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959+int32(-3)))))
	v979 = v969
	goto L246
L251:
	;
	v979 = int32(base.Ui32(v962) >> (uint(int32(3)) % 32))
	goto L246
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v984
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+320))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v990 = v988 + int32(2)
	if base.Ui32(v990) <= base.Ui32(v987) {
		v993 = v988
		v994 = v984
		goto L190
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+320)) = v990
	v993 = v988
	v994 = v984
	goto L190
L254:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v1035-v1036) < base.Ui32(v1005+int32(2)) {
		goto L260
	} else {
		goto L261
	}
L255:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1006+int32(-17))))
	v1035 = v1034
	goto L254
L256:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1006+int32(-9))))
	v1035 = v1031
	goto L254
L257:
	;
	v1028 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006+int32(-5)))))
	v1035 = v1028
	goto L254
L258:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006+int32(-3)))))
	v1035 = v1025
	goto L254
L259:
	;
	v1035 = int32(base.Ui32(v1018) >> (uint(int32(3)) % 32))
	goto L254
L260:
	;
	goto L102
L261:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1041 < v1042 {
		v1063 = v1005
		v1064 = v1006
		v1065 = v1036
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1066 = int32(4194304)
	v1067 = v1064 + v1065
	v1068 = v1067 + v1063
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068))))
	if v1069 != int32(13) {
		v1244 = v1066
		goto L1
	} else {
		goto L271
	}
L263:
	;
	v1045 = v1042 << (uint(int32(1)) % 32)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1047 = v1046 + v1041
	if v1045 < v1047 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1049 = v1045
	goto L266
L265:
	;
	v1049 = v1047
	goto L266
L266:
	;
	if v1042 < int32(1073741823) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1052 = v1049
	goto L269
L268:
	;
	v1052 = v1047
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1052
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1057 = F_valkey_realloc(m, v1054, v1052<<(uint(int32(2))%32))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L94
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1057
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1063 = v1060
	v1064 = v1062
	v1065 = v1061
	goto L262
L271:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068+int32(1)))))
	if v1074 != int32(10) {
		v1244 = v1066
		goto L1
	} else {
		goto L272
	}
L272:
	;
	if v1065|v418 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1182 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1182
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1187 = v1185 + v1182
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1187
	if v1187 != 0 {
		v427 = v1182
		goto L101
	} else {
		goto L295
	}
L274:
	;
	v1158 = F_createStringObject_1(m, v1067, v1063)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L94
	} else {
		goto L294
	}
L275:
	;
	if v1063 < int32(32768) {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1064+int32(-1)))))
	switch v1082 & int32(7) {
	case 0:
		goto L282
	case 1:
		goto L281
	case 2:
		goto L280
	case 3:
		goto L279
	case 4:
		goto L278
	default:
		goto L274
	}
L277:
	;
	if v1099 != v1063+int32(2) {
		goto L274
	} else {
		goto L283
	}
L278:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1064+int32(-17))))
	v1099 = v1098
	goto L277
L279:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1064+int32(-9))))
	v1099 = v1095
	goto L277
L280:
	;
	v1092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064+int32(-5)))))
	v1099 = v1092
	goto L277
L281:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1064+int32(-3)))))
	v1099 = v1089
	goto L277
L282:
	;
	v1099 = int32(base.Ui32(v1082) >> (uint(int32(3)) % 32))
	goto L277
L283:
	;
	v1104 = F_createObject(m, int32(0), v1064)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L94
	} else {
		goto L284
	}
L284:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1107 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1106+v1107<<(uint(int32(2))%32)))) = v1104
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1115 + v1116
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_sdsIncrLen(m, v1119, int32(-2))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L94
	} else {
		goto L285
	}
L285:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1128 = F_sdsnewlen(m, v1124, v1125+int32(2))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L94
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1128
	v1133 = v1128 + int32(-1)
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133))))
	switch v1134 & int32(7) {
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
		goto L288
	}
L287:
	;
	goto L273
L288:
	;
	v1155 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1128))) = uint8(v1155)
	goto L287
L289:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1128+int32(-17)))) = int64(0)
	goto L288
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1128+int32(-9)))) = int32(0)
	goto L288
L291:
	;
	v1145 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1128+int32(-5)))) = uint16(v1145)
	goto L288
L292:
	;
	v1141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1128+int32(-3)))) = uint8(v1141)
	goto L288
L293:
	;
	v1137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1133))) = uint8(v1137)
	goto L288
L294:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v1161 + int32(1)
	v1165 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1160+v1161<<(uint(v1165)%32)))) = v1158
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1169 + v1170
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1173 + v1174 + v1165
	goto L273
L295:
	;
	goto L2
L296:
	;
	v1244 = int32(0)
	goto L1
L297:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parseMultibulkBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	v16 = F_parseMultibulk(m, l0, l0+int32(24), l0+int32(20), l0+int32(28), l0+int32(32), l0+int32(240))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v19 = v16 | v18
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v19
	if v19&int32(65536) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F__serverAssert(m, int32(_a1689), int32(_a1630), int32(3654))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L34
	}
L4:
	;
	return
L5:
	;
	v24 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v25&v24 != 0 {
		v32 = v24
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v35 != 0 {
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v35 = v32
	goto L6
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = F_isImportSlotMigrationJob(m, v28)
	mBase = m.M
	v32 = v30
	goto L7
L10:
	;
	v35 = int32(0)
	goto L6
L11:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if v36 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v16&int32(8192) == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-1)))))
	switch v49 & int32(7) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		goto L4
	}
L15:
	;
	goto L4
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v66) <= base.Ui32(v67) {
		goto L4
	} else {
		goto L22
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-17))))
	v66 = v65
	goto L16
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(-9))))
	v66 = v62
	goto L16
L19:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+int32(-5)))))
	v66 = v59
	goto L16
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+int32(-3)))))
	v66 = v56
	goto L16
L21:
	;
	v66 = int32(base.Ui32(v49) >> (uint(int32(3)) % 32))
	goto L16
L22:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v67))))
	if v70 != int32(42) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(2)
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	if v75 == v76 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v98 = v96 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)) = uint16(v98)
	v104 = v95 + v96&int32(65535)*int32(40)
	v105 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = v105
	*(*int64)(unsafe.Add(mBase, uint32(v104+int32(32)))) = v105
	v112 = v104 + int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v112))) = v105
	v116 = v104 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v105
	v120 = v104 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v120))) = v105
	v129 = F_parseMultibulk(m, l0, v104+int32(4), v120, v104+int32(12), v104+int32(20), v112)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	if v75 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v95 = v78
	v96 = v75
	goto L24
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v91 = F_valkey_realloc(m, v86, v84&int32(65535)*int32(40))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	if base.Ui32(int32(512)) < base.Ui32(v75) {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v84 = int32(16)
	goto L27
L30:
	;
	v84 = v75 << (uint(int32(1)) % 32)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v91
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	v95 = v91
	v96 = v94
	goto L24
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v129
	if v129&int32(8192) != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	goto L15
L34:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pathIsBaseName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v2 = int32(0)
	v4 = int32(47)
	v5 = F___strchrnul(m, l0, v4)
	mBase = m.M
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v7 == v4 {
		v11 = v5
	} else {
		v11 = v2
	}
	if v11 != 0 {
		v22 = v2
	} else {
		v12 = int32(92)
		v13 = F___strchrnul(m, l0, v12)
		mBase = m.M
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v15 == v12 {
			v19 = v13
		} else {
			v19 = int32(0)
		}
		v22 = base.B2i32(v19 == int32(0))
	}
	return v22
}
func F_performModuleConfigSetFromName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v6 = F_dictFind(m, v5, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a499)
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v12 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a499)
				return int32(0)
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
				if v15&int32(1) != 0 {
					v23 = F_performInterfaceSet(m, v12, l1, l2)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(_a499)
					return int32(0)
				}
			}
		}
	}
}
func F_perror(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	goto L1
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v8 = F___strerror_l(m, v7, v7)
	mBase = m.M
	goto L2
L2:
	;
	v9 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1016]))
	if v9 <= v10 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1017]))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1018]))
	if l0 == v19 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L6
L5:
	;
	v18 = int32(1)
	goto L3
L6:
	;
	v18 = int32(0)
	goto L3
L7:
	;
	if v8&int32(3) == int32(0) {
		v116 = v8
		goto L32
	} else {
		goto L33
	}
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l0&int32(3) == int32(0) {
		v49 = l0
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v85 = F_fwrite(m, l0, v82, int32(1), int32(_a2769))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	v82 = v74 - l0
	goto L10
L12:
	;
	v53 = v49
	goto L20
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = l0
	goto L16
L15:
	;
	v82 = l0 - l0
	goto L10
L16:
	;
	v42 = v38 + int32(1)
	if v42&int32(3) == int32(0) {
		v49 = v42
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 != 0 {
		v38 = v42
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v74 = v42
	goto L11
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v62 = int32(-2139062144)
	if (int32(16843008)-v59|v59)&v62 == v62 {
		v53 = v53 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v68 = v53
	goto L23
L22:
	;
	goto L21
L23:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 != 0 {
		v68 = v68 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v74 = v68
	goto L11
L25:
	;
	goto L24
L26:
	;
	return
L27:
	;
	v89 = F_fputc(m, int32(58), int32(_a2769))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v93 = F_fputc(m, int32(32), int32(_a2769))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	v152 = F_fwrite(m, v8, v149, int32(1), int32(_a2769))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L26
	} else {
		goto L46
	}
L31:
	;
	v149 = v141 - v8
	goto L30
L32:
	;
	v120 = v116
	goto L40
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v102 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = v8
	goto L36
L35:
	;
	v149 = v8 - v8
	goto L30
L36:
	;
	v109 = v105 + int32(1)
	if v109&int32(3) == int32(0) {
		v116 = v109
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v114 != 0 {
		v105 = v109
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v141 = v109
	goto L31
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 == v129 {
		v120 = v120 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v135 = v120
	goto L43
L42:
	;
	goto L41
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 != 0 {
		v135 = v135 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v141 = v135
	goto L31
L45:
	;
	goto L44
L46:
	;
	v156 = F_fputc(m, int32(10), int32(_a2769))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v158 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1018])) = v22
	*(*int32)(unsafe.Add(mBase, _consts[1017])) = v20
	if v18 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return
L49:
	;
	goto L50
L50:
	;
	goto L48
}
func F_persistCommand(m *base.Module, l0 int32) {
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	v5 = F_lookupKeyWrite(m, v2, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if v5 == int32(0) {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[71]))
			F_addReply(m, l0, v44)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				return
			}
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v12 = F_removeExpire(m, v9, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[71]))
					F_addReply(m, l0, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						return
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					F_signalModifiedKey(m, l0, v16, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
						F_notifyKeyspaceEvent(m, int32(4), int32(_a616), v24, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _consts[331]))
							F_addReply(m, l0, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								v33 = int32(_a20)
								v35 = *(*int64)(unsafe.Add(mBase, _consts[180]))
								*(*int64)(unsafe.Add(mBase, _consts[180])) = v35 + int64(1)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pexpireatCommand(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	F_expireGenericCommand(m, l0, int64(0), int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_pexpiretimeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v8 = F_lookupKeyReadWithFlags(m, v4, v6, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if v8 != 0 {
			v13 = int64(-1)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v17&int32(1) == int32(0) {
				v28 = v13
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v8+(v17&int32(4)^int32(12)))))
				v28 = v27
			}
			v29 = int64(0)
			if v29 < v28 {
				v32 = v28
			} else {
				v32 = v29
			}
			if v28 == int64(-1) {
				v35 = v13
			} else {
				v35 = v32
			}
			F_addReplyLongLong(m, l0, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				return
			}
		} else {
			F_addReplyLongLong(m, l0, int64(-2))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pfcountCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(16400)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 < int32(3) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16400)
	return
L2:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = F_lookupKeyRead(m, v58, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L19
	}
L3:
	;
	v18 = F__emscripten_memset_bulkmem(m, v10, base.I32_extend8_s(int32(0)), int32(16400))
	mBase = m.M
	goto L4
L4:
	;
	v19 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v19)
	v28 = int32(1)
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v28<<(uint(int32(2))%32))))
	v37 = F_lookupKeyRead(m, v31, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v54 = F_hllCount(m, v18, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L16
	}
L7:
	;
	v50 = v28 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v50 < v51 {
		v28 = v50
		goto L5
	} else {
		goto L15
	}
L8:
	;
	return
L9:
	;
	if v37 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v41 = F_isHLLObjectOrReply(m, l0, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = F_hllMerge(m, v18+int32(16), v37)
	mBase = m.M
	if v43 != int32(-1) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	F_addReplyError(m, l0, int32(_a666))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	goto L6
L16:
	;
	F_addReplyLongLong(m, l0, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	goto L1
L18:
	;
	v67 = F_isHLLObjectOrReply(m, l0, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L22
	}
L19:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	F_addReply(m, l0, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = F_dbUnshareStringValue(m, v69, v71, v61)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L27
	}
L24:
	;
	F_addReplyError(m, l0, int32(_a666))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L33
	}
L25:
	;
	F_addReplyLongLong(m, l0, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L32
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
	v96 = F_hllCount(m, v74, v10)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	v74 = F_objectGetVal(m, v72)
	mBase = m.M
	v75 = int64(*(*int8)(unsafe.Add(mBase, uint32(v74)+15)))
	if v75 < int64(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v80 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v74)+8)))
	v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+12)))
	v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+13)))
	v90 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+14)))
	v111 = v75<<(uint(int64(56))%64) | v80 | v82<<(uint(int64(32))%64) | v86<<(uint(int64(40))%64) | v90<<(uint(int64(48))%64)
	goto L25
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v98 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v96
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	F_signalModifiedKey(m, l0, v100, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v105 = int32(_a20)
	v107 = *(*int64)(unsafe.Add(mBase, _consts[180]))
	*(*int64)(unsafe.Add(mBase, _consts[180])) = v107 + int64(1)
	v111 = v96
	goto L25
L32:
	;
	goto L1
L33:
	;
	goto L1
}
func F_pfselftestCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v215 int64
	_ = v215
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v246 int64
	_ = v246
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 float64
	_ = v333
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v355 int64
	_ = v355
	var v359 int64
	_ = v359
	var v363 int64
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16432)
	m.G0 = v20
	v25 = F_sdsnewlen(m, v2, int32(12304))
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
	v28 = v25 + int32(16)
	v31 = v2
	goto L4
L3:
	;
	m.G0 = v20 + int32(16432)
	return
L4:
	;
	v52 = int32(0)
	goto L6
L5:
	;
	v171 = F__emscripten_memset_bulkmem(m, v28, base.I32_extend8_s(int32(0)), int32(12288))
	mBase = m.M
	goto L18
L6:
	;
	v68 = int32(0)
	v70 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v74 = v70*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v74
	goto L8
L7:
	;
	v117 = int32(0)
	goto L10
L8:
	;
	v79 = int32(63)
	v80 = base.I32_wrap_i64(int64(base.Ui64(v74)>>(uint(int64(33))%64))) & v79
	*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(48)+v52))) = uint8(v80)
	v82 = int32(6)
	v83 = v52 * v82
	v86 = v28 + int32(base.Ui32(v83)>>(uint(int32(3))%32))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v90 = v83 & v82
	v96 = v87&(v79<<(uint(v90)%32)^int32(-1)) | v80<<(uint(v90)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v96)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v101 = int32(8) - v90
	v105 = v98&(int32(-64)>>(uint(v101)%32)) | int32(base.Ui32(v80)>>(uint(v101)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)) = uint8(v105)
	v108 = v52 + int32(1)
	if v108 != int32(16384) {
		v52 = v108
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v129 = int32(6)
	v130 = v117 * v129
	v133 = v28 + int32(base.Ui32(v130)>>(uint(int32(3))%32))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = v130 & v129
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = (v134<<(uint(int32(8)-v137)%32) | int32(base.Ui32(v140)>>(uint(v137)%32))) & int32(63)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+int32(48)+v117))))
	if v144 == v148 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v165 = v31 + int32(1)
	if v165 != int32(1000) {
		v31 = v165
		goto L4
	} else {
		goto L17
	}
L12:
	;
	v161 = v117 + int32(1)
	if v161 != int32(16384) {
		v117 = v161
		goto L10
	} else {
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v117
	F_addReplyErrorFormat(m, l0, int32(_a667), v20+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_sdsfree(m, v25)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	goto L11
L17:
	;
	goto L5
L18:
	;
	v174 = F_sdsnewlen(m, int32(0), int32(18))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v176 = int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v174)+16)) = uint16(v176)
	v179 = F_createObject(m, int32(0), v174)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v181 = F_objectGetVal(m, v179)
	mBase = m.M
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+4)) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = int32(1280072008)
	v187 = int32(0)
	v189 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v193 = v189*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v193
	goto L21
L21:
	;
	v199 = int32(0)
	v201 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v205 = v201*int64(6364136223846793005) + int64(1)
	*(*int64)(unsafe.Add(mBase, _consts[343])) = v205
	goto L22
L22:
	;
	v215 = int64(1)
	v228 = v215
	v229 = v215
	goto L25
L23:
	;
	F_sdsfree(m, v25)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L52
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v228
	F_addReplyErrorFormat(m, l0, int32(_a668), v20)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L51
	}
L25:
	;
	v234 = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v205)>>(uint(int64(33))%64))))<<(uint(int64(32))%64) | base.I64_extend_i32_s(base.I32_wrap_i64(int64(base.Ui64(v193)>>(uint(int64(33))%64)))) ^ v228
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v234
	v236 = int64(-4132994306676758123)
	v237 = v234 * v236
	v238 = int64(47)
	v246 = ((int64(base.Ui64(v237)>>(uint(v238)%64))^v237)*v236 ^ int64(3829533692205168561)) * v236
	v251 = (int64(base.Ui64(v246)>>(uint(v238)%64)) ^ v246) * v236
	v254 = int64(base.Ui64(v251)>>(uint(v238)%64)) ^ v251
	v258 = int32(6)
	v259 = base.I32_wrap_i64(v254) & int32(16383) * v258
	v262 = v171 + int32(base.Ui32(v259)>>(uint(int32(3))%32))
	v264 = v262 + int32(1)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v268 = v259 & v258
	v269 = int32(8) - v268
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	v281 = base.I32_wrap_i64(base.I64_ctz(int64(base.Ui64(v254)>>(uint(int64(14))%64)) | int64(1125899906842624)))
	if base.Ui32(v281) < base.Ui32((v265<<(uint(v269)%32)|int32(base.Ui32(v271)>>(uint(v268)%32)))&int32(63)) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	F_addReply(m, l0, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L50
	}
L27:
	;
	v303 = F_hllAdd(m, v179, v20+int32(40), int32(8))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v287 = v281 + int32(1)
	v289 = v265&(int32(-64)>>(uint(v269)%32)) | int32(base.Ui32(v287)>>(uint(v269)%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v264))) = uint8(v289)
	v297 = v271&(int32(63)<<(uint(v268)%32)^int32(-1)) | v287<<(uint(v268)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v297)
	goto L27
L29:
	;
	if v229 != v228 {
		v359 = v229
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v363 = v228 + int64(1)
	if v363 != int64(10000001) {
		v228 = v363
		v229 = v359
		goto L25
	} else {
		goto L49
	}
L31:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[344]))
	if base.Ui64(base.I64_extend_i32_u(int32(base.Ui32(v307)>>(uint(int32(1))%32)))) <= base.Ui64(v228) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v320 = F_hllCount(m, v25, int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v312 = F_objectGetVal(m, v179)
	mBase = m.M
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+4)))
	if v313 == int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	F_addReplyError(m, l0, int32(_a669))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v333 = base.F64_ceil(base.F64_mul(base.F64_convert_i64_u(v228), float64(0.04875)))
	if base.F64_lt(v333, float64(1.8446744073709552e+19))&base.F64_ge(v333, float64(0)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v322 = F_objectGetVal(m, v179)
	mBase = m.M
	v324 = F_hllCount(m, v322, int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v320 == v324 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_addReplyError(m, l0, int32(_a670))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L23
L41:
	;
	v345 = F_hllCount(m, v25, int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v343 = int64(0)
	goto L41
L43:
	;
	v341 = base.I64_trunc_f64_u(v333)
	v343 = v341
	goto L41
L44:
	;
	v347 = v228 - v345
	v349 = v347 >> (uint(int64(63)) % 64)
	v351 = v347 ^ v349 - v349
	if v228 == int64(10) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v355 = int64(1)
	goto L47
L46:
	;
	v355 = v343
	goto L47
L47:
	;
	if v355 < v351 {
		goto L24
	} else {
		goto L48
	}
L48:
	;
	v359 = v228 * int64(10)
	goto L30
L49:
	;
	goto L26
L50:
	;
	goto L23
L51:
	;
	goto L23
L52:
	;
	if v179 == int32(0) {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	F_decrRefCount(m, v179)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L3
}
func F_poll(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = m.Env.X__syscall_poll(m, l0, l1, l2)
	mBase = m.M
	if base.Ui32(v4) < base.Ui32(int32(-4095)) {
		v12 = v4
	} else {
		v7 = F___errno_location(m)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0) - v4
		v12 = int32(-1)
	}
	return v12
}
func F_populateCommandTable(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	v1 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	if v6 == v1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a2197), int32(_a2157), int32(3375))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L16
	}
L2:
	;
	return
L3:
	;
	v9 = v1
	v10 = v6
	goto L4
L4:
	;
	v14 = v9 * int32(224)
	v15 = int32(_a2198) + v14
	v16 = F_sdsnew(m, v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L2
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[747]))) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_consts[748]))) = v16
	v20 = F_populateCommandStructure(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L9
	}
L8:
	;
	v40 = v9 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40*int32(224))+uint32(_consts[746])))
	if v44 != 0 {
		v9 = v40
		v10 = v44
		goto L4
	} else {
		goto L15
	}
L9:
	;
	if v20 == int32(-1) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v26 = F_hashtableAdd(m, v25, v15)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v30 = F_hashtableAdd(m, v29, v15)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v30 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	goto L5
L16:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pow(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v71 float64
	_ = v71
	var v80 float64
	_ = v80
	var v92 int32
	_ = v92
	var v103 int64
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v153 int64
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 float64
	_ = v167
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v191 float64
	_ = v191
	var v201 float64
	_ = v201
	var v204 float64
	_ = v204
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v217 float64
	_ = v217
	var v222 int32
	_ = v222
	var v229 int64
	_ = v229
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v243 int32
	_ = v243
	var v246 float64
	_ = v246
	var v250 int64
	_ = v250
	var v255 float64
	_ = v255
	var v258 float64
	_ = v258
	var v261 float64
	_ = v261
	var v264 float64
	_ = v264
	var v265 float64
	_ = v265
	var v267 float64
	_ = v267
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v278 float64
	_ = v278
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v284 float64
	_ = v284
	var v285 float64
	_ = v285
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v295 float64
	_ = v295
	var v299 float64
	_ = v299
	var v302 float64
	_ = v302
	var v307 float64
	_ = v307
	var v310 float64
	_ = v310
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v322 float64
	_ = v322
	var v323 float64
	_ = v323
	var v326 float64
	_ = v326
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 float64
	_ = v349
	var v351 float64
	_ = v351
	var v353 int32
	_ = v353
	var v359 float64
	_ = v359
	var v360 float64
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 float64
	_ = v364
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v369 float64
	_ = v369
	var v371 float64
	_ = v371
	var v374 float64
	_ = v374
	var v378 float64
	_ = v378
	var v379 float64
	_ = v379
	var v382 float64
	_ = v382
	var v385 float64
	_ = v385
	var v389 float64
	_ = v389
	var v392 float64
	_ = v392
	var v395 int64
	_ = v395
	var v400 int32
	_ = v400
	var v403 float64
	_ = v403
	var v406 float64
	_ = v406
	var v409 int64
	_ = v409
	var v414 int64
	_ = v414
	var v415 float64
	_ = v415
	var v416 float64
	_ = v416
	var v427 float64
	_ = v427
	var v432 float64
	_ = v432
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v21 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
	v25 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l1)) >> (uint(int64(52)) % 64)))
	v27 = v25 & int32(2047)
	v29 = v27 + int32(-1086)
	v30 = base.I64_reinterpret_f64(l1)
	v31 = base.I64_reinterpret_f64(l0)
	if base.Ui32(v21+int32(-2047)) < base.Ui32(int32(-2046)) {
		if base.B2i32(base.Ui64(v30<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993))) == int32(0) {
			if base.B2i32(base.Ui64(v31<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993))) == int32(0) {
				if int64(-1) < v31 {
					v178 = v21
					v179 = v31
					v180 = int32(0)
					if base.Ui32(int32(-129)) < base.Ui32(v29) {
						if v178 != 0 {
							v213 = v179
							v214 = v180
						} else {
							v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) + int64(-234187180623265792)
							v214 = v180
						}
						v217 = base.F64_reinterpret_i64(v30 & int64(-134217728))
						v222 = int32(0)
						v229 = v213 + int64(-4604531861337669632)
						v233 = base.F64_convert_i32_s(base.I32_wrap_i64(v229 >> (uint(int64(52)) % 64)))
						v235 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
						v243 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
						v246 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1020])))
						v250 = v213 - v229&int64(-4503599627370496)
						v255 = base.F64_reinterpret_i64((v250 + int64(2147483648)) & int64(-4294967296))
						v258 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1021])))
						v261 = base.F64_add(base.F64_mul(v255, v258), float64(-1))
						v264 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v250), v255), v258)
						v265 = base.F64_add(v261, v264)
						v267 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
						v271 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1023])))
						v272 = base.F64_add(base.F64_mul(v233, v267), v271)
						v273 = base.F64_add(v265, v272)
						v278 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
						v279 = base.F64_mul(v265, v278)
						v280 = base.F64_mul(v261, v278)
						v284 = base.F64_mul(v261, v280)
						v285 = base.F64_add(v273, v284)
						v289 = base.F64_mul(v265, v279)
						v292 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
						v295 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
						v299 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
						v302 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
						v307 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
						v310 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
						v314 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v233, v235), v246), base.F64_add(v265, base.F64_sub(v272, v273))), base.F64_mul(v264, base.F64_add(v279, v280))), base.F64_add(v284, base.F64_sub(v273, v285))), base.F64_mul(base.F64_mul(v265, v289), base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v265, v292), v295)), base.F64_add(base.F64_mul(v265, v299), v302))), base.F64_add(base.F64_mul(v265, v307), v310))))
						v315 = base.F64_add(v285, v314)
						*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = base.F64_add(v314, base.F64_sub(v285, v315))
						v322 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v315) & int64(-134217728))
						v323 = base.F64_mul(v217, v322)
						v326 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
						v337 = F_top12_2(m, v323)
						mBase = m.M
						v339 = v337 & int32(2047)
						v341 = F_top12_2(m, float64(5.551115123125783e-17))
						mBase = m.M
						v344 = F_top12_2(m, float64(512))
						mBase = m.M
						if base.Ui32(v339-v341) < base.Ui32(v344-v341) {
							v361 = v339
							v363 = int32(0)
							v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
							v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
							v368 = base.F64_add(base.F64_mul(v323, v364), v367)
							v369 = base.F64_sub(v368, v367)
							v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
							v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
							v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
							v379 = base.F64_mul(v378, v378)
							v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
							v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
							v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
							v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
							v395 = base.I64_reinterpret_f64(v368)
							v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
							v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
							v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
							v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
							v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
							if v361 != 0 {
								v416 = base.F64_reinterpret_i64(v414)
								v427 = base.F64_add(base.F64_mul(v416, v406), v416)
							} else {
								v415 = F_specialcase_2(m, v406, v414, v395)
								mBase = m.M
								v427 = v415
							}
						} else {
							if base.Ui32(v341) <= base.Ui32(v339) {
								v353 = F_top12_2(m, float64(1024))
								mBase = m.M
								if base.Ui32(v339) < base.Ui32(v353) {
									v361 = int32(0)
									v363 = int32(0)
									v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
									v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
									v368 = base.F64_add(base.F64_mul(v323, v364), v367)
									v369 = base.F64_sub(v368, v367)
									v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
									v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
									v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
									v379 = base.F64_mul(v378, v378)
									v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
									v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
									v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
									v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
									v395 = base.I64_reinterpret_f64(v368)
									v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
									v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
									v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
									v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
									v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
									if v361 != 0 {
										v416 = base.F64_reinterpret_i64(v414)
										v427 = base.F64_add(base.F64_mul(v416, v406), v416)
									} else {
										v415 = F_specialcase_2(m, v406, v414, v395)
										mBase = m.M
										v427 = v415
									}
								} else {
									if int64(-1) < base.I64_reinterpret_f64(v323) {
										v360 = F___math_oflow(m, v214)
										mBase = m.M
										v427 = v360
									} else {
										v359 = F___math_uflow(m, v214)
										mBase = m.M
										v427 = v359
									}
								}
							} else {
								v349 = base.F64_add(v323, float64(1))
								if v214 != 0 {
									v351 = base.F64_neg(v349)
								} else {
									v351 = v349
								}
								v427 = v351
							}
						}
						v432 = v427
					} else {
						if v179 == int64(4607182418800017408) {
							v432 = float64(1)
						} else {
							if base.Ui32(int32(957)) < base.Ui32(v27) {
								if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v25)) == base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
									v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
									mBase = m.M
									v432 = v204
								} else {
									v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
									mBase = m.M
									v432 = v201
								}
							} else {
								if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
									v191 = l1
								} else {
									v191 = base.F64_neg(l1)
								}
								v432 = base.F64_add(v191, float64(1))
							}
						}
					}
				} else {
					v142 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v142) < base.Ui32(int32(1023)) {
						v165 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v142) {
							v165 = int32(2)
						} else {
							v153 = int64(1) << (uint(base.I64_extend_i32_u(int32(1075)-v142)) % 64)
							if (v153+int64(-1))&v30 != int64(0) {
								v165 = int32(0)
							} else {
								if v153&v30 == int64(0) {
									v164 = int32(2)
								} else {
									v164 = int32(1)
								}
								v165 = v164
							}
						}
					}
					if v165 != 0 {
						v178 = v21 & int32(2047)
						v179 = base.I64_reinterpret_f64(l0) & int64(9223372036854775807)
						v180 = base.B2i32(v165 == int32(1)) << (uint(int32(18)) % 32)
						if base.Ui32(int32(-129)) < base.Ui32(v29) {
							if v178 != 0 {
								v213 = v179
								v214 = v180
							} else {
								v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) + int64(-234187180623265792)
								v214 = v180
							}
							v217 = base.F64_reinterpret_i64(v30 & int64(-134217728))
							v222 = int32(0)
							v229 = v213 + int64(-4604531861337669632)
							v233 = base.F64_convert_i32_s(base.I32_wrap_i64(v229 >> (uint(int64(52)) % 64)))
							v235 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
							v243 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
							v246 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1020])))
							v250 = v213 - v229&int64(-4503599627370496)
							v255 = base.F64_reinterpret_i64((v250 + int64(2147483648)) & int64(-4294967296))
							v258 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1021])))
							v261 = base.F64_add(base.F64_mul(v255, v258), float64(-1))
							v264 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v250), v255), v258)
							v265 = base.F64_add(v261, v264)
							v267 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
							v271 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1023])))
							v272 = base.F64_add(base.F64_mul(v233, v267), v271)
							v273 = base.F64_add(v265, v272)
							v278 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
							v279 = base.F64_mul(v265, v278)
							v280 = base.F64_mul(v261, v278)
							v284 = base.F64_mul(v261, v280)
							v285 = base.F64_add(v273, v284)
							v289 = base.F64_mul(v265, v279)
							v292 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
							v295 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
							v299 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
							v302 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
							v307 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
							v310 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
							v314 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v233, v235), v246), base.F64_add(v265, base.F64_sub(v272, v273))), base.F64_mul(v264, base.F64_add(v279, v280))), base.F64_add(v284, base.F64_sub(v273, v285))), base.F64_mul(base.F64_mul(v265, v289), base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v265, v292), v295)), base.F64_add(base.F64_mul(v265, v299), v302))), base.F64_add(base.F64_mul(v265, v307), v310))))
							v315 = base.F64_add(v285, v314)
							*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = base.F64_add(v314, base.F64_sub(v285, v315))
							v322 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v315) & int64(-134217728))
							v323 = base.F64_mul(v217, v322)
							v326 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
							v337 = F_top12_2(m, v323)
							mBase = m.M
							v339 = v337 & int32(2047)
							v341 = F_top12_2(m, float64(5.551115123125783e-17))
							mBase = m.M
							v344 = F_top12_2(m, float64(512))
							mBase = m.M
							if base.Ui32(v339-v341) < base.Ui32(v344-v341) {
								v361 = v339
								v363 = int32(0)
								v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
								v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
								v368 = base.F64_add(base.F64_mul(v323, v364), v367)
								v369 = base.F64_sub(v368, v367)
								v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
								v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
								v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
								v379 = base.F64_mul(v378, v378)
								v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
								v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
								v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
								v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
								v395 = base.I64_reinterpret_f64(v368)
								v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
								v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
								v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
								v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
								v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
								if v361 != 0 {
									v416 = base.F64_reinterpret_i64(v414)
									v427 = base.F64_add(base.F64_mul(v416, v406), v416)
								} else {
									v415 = F_specialcase_2(m, v406, v414, v395)
									mBase = m.M
									v427 = v415
								}
							} else {
								if base.Ui32(v341) <= base.Ui32(v339) {
									v353 = F_top12_2(m, float64(1024))
									mBase = m.M
									if base.Ui32(v339) < base.Ui32(v353) {
										v361 = int32(0)
										v363 = int32(0)
										v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
										v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
										v368 = base.F64_add(base.F64_mul(v323, v364), v367)
										v369 = base.F64_sub(v368, v367)
										v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
										v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
										v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
										v379 = base.F64_mul(v378, v378)
										v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
										v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
										v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
										v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
										v395 = base.I64_reinterpret_f64(v368)
										v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
										v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
										v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
										v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
										v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
										if v361 != 0 {
											v416 = base.F64_reinterpret_i64(v414)
											v427 = base.F64_add(base.F64_mul(v416, v406), v416)
										} else {
											v415 = F_specialcase_2(m, v406, v414, v395)
											mBase = m.M
											v427 = v415
										}
									} else {
										if int64(-1) < base.I64_reinterpret_f64(v323) {
											v360 = F___math_oflow(m, v214)
											mBase = m.M
											v427 = v360
										} else {
											v359 = F___math_uflow(m, v214)
											mBase = m.M
											v427 = v359
										}
									}
								} else {
									v349 = base.F64_add(v323, float64(1))
									if v214 != 0 {
										v351 = base.F64_neg(v349)
									} else {
										v351 = v349
									}
									v427 = v351
								}
							}
							v432 = v427
						} else {
							if v179 == int64(4607182418800017408) {
								v432 = float64(1)
							} else {
								if base.Ui32(int32(957)) < base.Ui32(v27) {
									if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v25)) == base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
										v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
										mBase = m.M
										v432 = v204
									} else {
										v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
										mBase = m.M
										v432 = v201
									}
								} else {
									if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
										v191 = l1
									} else {
										v191 = base.F64_neg(l1)
									}
									v432 = base.F64_add(v191, float64(1))
								}
							}
						}
					} else {
						v167 = base.F64_sub(l0, l0)
						v432 = base.F64_div(v167, v167)
					}
				}
			} else {
				v80 = base.F64_mul(l0, l0)
				if int64(-1) < v31 {
					v120 = v80
				} else {
					v92 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v92) < base.Ui32(int32(1023)) {
						v115 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v92) {
							v115 = int32(2)
						} else {
							v103 = int64(1) << (uint(base.I64_extend_i32_u(int32(1075)-v92)) % 64)
							if (v103+int64(-1))&v30 != int64(0) {
								v115 = int32(0)
							} else {
								if v103&v30 == int64(0) {
									v114 = int32(2)
								} else {
									v114 = int32(1)
								}
								v115 = v114
							}
						}
					}
					if v115 == int32(1) {
						v119 = base.F64_neg(v80)
					} else {
						v119 = v80
					}
					v120 = v119
				}
				if int64(-1) < v30 {
					v432 = v120
				} else {
					v124 = base.F64_div(float64(1), v120)
					v126 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v126-int32(16))+8)) = v124
					v432 = v124
				}
			}
		} else {
			v48 = float64(1)
			if v31 == int64(4607182418800017408) {
				v432 = v48
			} else {
				v52 = v30 << (uint(int64(1)) % 64)
				if v52 == int64(0) {
					v432 = v48
				} else {
					v56 = v31 << (uint(int64(1)) % 64)
					if base.Ui64(int64(-9007199254740992)) < base.Ui64(v56) {
						v432 = base.F64_add(l0, l1)
					} else {
						if base.Ui64(v52) < base.Ui64(int64(-9007199254740991)) {
							if v56 == int64(9214364837600034816) {
								v432 = v48
							} else {
								if base.B2i32(base.Ui64(v56) < base.Ui64(int64(9214364837600034816)))^base.B2i32(v30 < int64(0)) != 0 {
									v71 = float64(0)
								} else {
									v71 = base.F64_mul(l1, l1)
								}
								v432 = v71
							}
						} else {
							v432 = base.F64_add(l0, l1)
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(int32(-129)) < base.Ui32(v29) {
			v213 = v31
			v214 = int32(0)
			v217 = base.F64_reinterpret_i64(v30 & int64(-134217728))
			v222 = int32(0)
			v229 = v213 + int64(-4604531861337669632)
			v233 = base.F64_convert_i32_s(base.I32_wrap_i64(v229 >> (uint(int64(52)) % 64)))
			v235 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
			v243 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
			v246 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1020])))
			v250 = v213 - v229&int64(-4503599627370496)
			v255 = base.F64_reinterpret_i64((v250 + int64(2147483648)) & int64(-4294967296))
			v258 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1021])))
			v261 = base.F64_add(base.F64_mul(v255, v258), float64(-1))
			v264 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v250), v255), v258)
			v265 = base.F64_add(v261, v264)
			v267 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
			v271 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1023])))
			v272 = base.F64_add(base.F64_mul(v233, v267), v271)
			v273 = base.F64_add(v265, v272)
			v278 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
			v279 = base.F64_mul(v265, v278)
			v280 = base.F64_mul(v261, v278)
			v284 = base.F64_mul(v261, v280)
			v285 = base.F64_add(v273, v284)
			v289 = base.F64_mul(v265, v279)
			v292 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
			v295 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
			v299 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
			v302 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
			v307 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
			v310 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
			v314 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v233, v235), v246), base.F64_add(v265, base.F64_sub(v272, v273))), base.F64_mul(v264, base.F64_add(v279, v280))), base.F64_add(v284, base.F64_sub(v273, v285))), base.F64_mul(base.F64_mul(v265, v289), base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v265, v292), v295)), base.F64_add(base.F64_mul(v265, v299), v302))), base.F64_add(base.F64_mul(v265, v307), v310))))
			v315 = base.F64_add(v285, v314)
			*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = base.F64_add(v314, base.F64_sub(v285, v315))
			v322 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v315) & int64(-134217728))
			v323 = base.F64_mul(v217, v322)
			v326 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
			v337 = F_top12_2(m, v323)
			mBase = m.M
			v339 = v337 & int32(2047)
			v341 = F_top12_2(m, float64(5.551115123125783e-17))
			mBase = m.M
			v344 = F_top12_2(m, float64(512))
			mBase = m.M
			if base.Ui32(v339-v341) < base.Ui32(v344-v341) {
				v361 = v339
				v363 = int32(0)
				v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
				v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
				v368 = base.F64_add(base.F64_mul(v323, v364), v367)
				v369 = base.F64_sub(v368, v367)
				v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
				v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
				v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
				v379 = base.F64_mul(v378, v378)
				v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
				v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
				v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
				v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
				v395 = base.I64_reinterpret_f64(v368)
				v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
				v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
				v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
				v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
				v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
				if v361 != 0 {
					v416 = base.F64_reinterpret_i64(v414)
					v427 = base.F64_add(base.F64_mul(v416, v406), v416)
				} else {
					v415 = F_specialcase_2(m, v406, v414, v395)
					mBase = m.M
					v427 = v415
				}
			} else {
				if base.Ui32(v341) <= base.Ui32(v339) {
					v353 = F_top12_2(m, float64(1024))
					mBase = m.M
					if base.Ui32(v339) < base.Ui32(v353) {
						v361 = int32(0)
						v363 = int32(0)
						v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
						v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
						v368 = base.F64_add(base.F64_mul(v323, v364), v367)
						v369 = base.F64_sub(v368, v367)
						v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
						v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
						v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
						v379 = base.F64_mul(v378, v378)
						v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
						v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
						v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
						v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
						v395 = base.I64_reinterpret_f64(v368)
						v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
						v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
						v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
						v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
						v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
						if v361 != 0 {
							v416 = base.F64_reinterpret_i64(v414)
							v427 = base.F64_add(base.F64_mul(v416, v406), v416)
						} else {
							v415 = F_specialcase_2(m, v406, v414, v395)
							mBase = m.M
							v427 = v415
						}
					} else {
						if int64(-1) < base.I64_reinterpret_f64(v323) {
							v360 = F___math_oflow(m, v214)
							mBase = m.M
							v427 = v360
						} else {
							v359 = F___math_uflow(m, v214)
							mBase = m.M
							v427 = v359
						}
					}
				} else {
					v349 = base.F64_add(v323, float64(1))
					if v214 != 0 {
						v351 = base.F64_neg(v349)
					} else {
						v351 = v349
					}
					v427 = v351
				}
			}
			v432 = v427
		} else {
			if base.B2i32(base.Ui64(v30<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993))) == int32(0) {
				if base.B2i32(base.Ui64(v31<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993))) == int32(0) {
					if int64(-1) < v31 {
						v178 = v21
						v179 = v31
						v180 = int32(0)
						if base.Ui32(int32(-129)) < base.Ui32(v29) {
							if v178 != 0 {
								v213 = v179
								v214 = v180
							} else {
								v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) + int64(-234187180623265792)
								v214 = v180
							}
							v217 = base.F64_reinterpret_i64(v30 & int64(-134217728))
							v222 = int32(0)
							v229 = v213 + int64(-4604531861337669632)
							v233 = base.F64_convert_i32_s(base.I32_wrap_i64(v229 >> (uint(int64(52)) % 64)))
							v235 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
							v243 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
							v246 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1020])))
							v250 = v213 - v229&int64(-4503599627370496)
							v255 = base.F64_reinterpret_i64((v250 + int64(2147483648)) & int64(-4294967296))
							v258 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1021])))
							v261 = base.F64_add(base.F64_mul(v255, v258), float64(-1))
							v264 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v250), v255), v258)
							v265 = base.F64_add(v261, v264)
							v267 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
							v271 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1023])))
							v272 = base.F64_add(base.F64_mul(v233, v267), v271)
							v273 = base.F64_add(v265, v272)
							v278 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
							v279 = base.F64_mul(v265, v278)
							v280 = base.F64_mul(v261, v278)
							v284 = base.F64_mul(v261, v280)
							v285 = base.F64_add(v273, v284)
							v289 = base.F64_mul(v265, v279)
							v292 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
							v295 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
							v299 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
							v302 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
							v307 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
							v310 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
							v314 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v233, v235), v246), base.F64_add(v265, base.F64_sub(v272, v273))), base.F64_mul(v264, base.F64_add(v279, v280))), base.F64_add(v284, base.F64_sub(v273, v285))), base.F64_mul(base.F64_mul(v265, v289), base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v265, v292), v295)), base.F64_add(base.F64_mul(v265, v299), v302))), base.F64_add(base.F64_mul(v265, v307), v310))))
							v315 = base.F64_add(v285, v314)
							*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = base.F64_add(v314, base.F64_sub(v285, v315))
							v322 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v315) & int64(-134217728))
							v323 = base.F64_mul(v217, v322)
							v326 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
							v337 = F_top12_2(m, v323)
							mBase = m.M
							v339 = v337 & int32(2047)
							v341 = F_top12_2(m, float64(5.551115123125783e-17))
							mBase = m.M
							v344 = F_top12_2(m, float64(512))
							mBase = m.M
							if base.Ui32(v339-v341) < base.Ui32(v344-v341) {
								v361 = v339
								v363 = int32(0)
								v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
								v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
								v368 = base.F64_add(base.F64_mul(v323, v364), v367)
								v369 = base.F64_sub(v368, v367)
								v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
								v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
								v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
								v379 = base.F64_mul(v378, v378)
								v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
								v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
								v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
								v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
								v395 = base.I64_reinterpret_f64(v368)
								v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
								v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
								v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
								v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
								v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
								if v361 != 0 {
									v416 = base.F64_reinterpret_i64(v414)
									v427 = base.F64_add(base.F64_mul(v416, v406), v416)
								} else {
									v415 = F_specialcase_2(m, v406, v414, v395)
									mBase = m.M
									v427 = v415
								}
							} else {
								if base.Ui32(v341) <= base.Ui32(v339) {
									v353 = F_top12_2(m, float64(1024))
									mBase = m.M
									if base.Ui32(v339) < base.Ui32(v353) {
										v361 = int32(0)
										v363 = int32(0)
										v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
										v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
										v368 = base.F64_add(base.F64_mul(v323, v364), v367)
										v369 = base.F64_sub(v368, v367)
										v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
										v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
										v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
										v379 = base.F64_mul(v378, v378)
										v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
										v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
										v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
										v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
										v395 = base.I64_reinterpret_f64(v368)
										v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
										v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
										v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
										v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
										v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
										if v361 != 0 {
											v416 = base.F64_reinterpret_i64(v414)
											v427 = base.F64_add(base.F64_mul(v416, v406), v416)
										} else {
											v415 = F_specialcase_2(m, v406, v414, v395)
											mBase = m.M
											v427 = v415
										}
									} else {
										if int64(-1) < base.I64_reinterpret_f64(v323) {
											v360 = F___math_oflow(m, v214)
											mBase = m.M
											v427 = v360
										} else {
											v359 = F___math_uflow(m, v214)
											mBase = m.M
											v427 = v359
										}
									}
								} else {
									v349 = base.F64_add(v323, float64(1))
									if v214 != 0 {
										v351 = base.F64_neg(v349)
									} else {
										v351 = v349
									}
									v427 = v351
								}
							}
							v432 = v427
						} else {
							if v179 == int64(4607182418800017408) {
								v432 = float64(1)
							} else {
								if base.Ui32(int32(957)) < base.Ui32(v27) {
									if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v25)) == base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
										v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
										mBase = m.M
										v432 = v204
									} else {
										v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
										mBase = m.M
										v432 = v201
									}
								} else {
									if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
										v191 = l1
									} else {
										v191 = base.F64_neg(l1)
									}
									v432 = base.F64_add(v191, float64(1))
								}
							}
						}
					} else {
						v142 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(52))%64))) & int32(2047)
						if base.Ui32(v142) < base.Ui32(int32(1023)) {
							v165 = int32(0)
						} else {
							if base.Ui32(int32(1075)) < base.Ui32(v142) {
								v165 = int32(2)
							} else {
								v153 = int64(1) << (uint(base.I64_extend_i32_u(int32(1075)-v142)) % 64)
								if (v153+int64(-1))&v30 != int64(0) {
									v165 = int32(0)
								} else {
									if v153&v30 == int64(0) {
										v164 = int32(2)
									} else {
										v164 = int32(1)
									}
									v165 = v164
								}
							}
						}
						if v165 != 0 {
							v178 = v21 & int32(2047)
							v179 = base.I64_reinterpret_f64(l0) & int64(9223372036854775807)
							v180 = base.B2i32(v165 == int32(1)) << (uint(int32(18)) % 32)
							if base.Ui32(int32(-129)) < base.Ui32(v29) {
								if v178 != 0 {
									v213 = v179
									v214 = v180
								} else {
									v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) + int64(-234187180623265792)
									v214 = v180
								}
								v217 = base.F64_reinterpret_i64(v30 & int64(-134217728))
								v222 = int32(0)
								v229 = v213 + int64(-4604531861337669632)
								v233 = base.F64_convert_i32_s(base.I32_wrap_i64(v229 >> (uint(int64(52)) % 64)))
								v235 = *(*float64)(unsafe.Add(mBase, _consts[1019]))
								v243 = base.I32_wrap_i64(int64(base.Ui64(v229)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
								v246 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1020])))
								v250 = v213 - v229&int64(-4503599627370496)
								v255 = base.F64_reinterpret_i64((v250 + int64(2147483648)) & int64(-4294967296))
								v258 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1021])))
								v261 = base.F64_add(base.F64_mul(v255, v258), float64(-1))
								v264 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v250), v255), v258)
								v265 = base.F64_add(v261, v264)
								v267 = *(*float64)(unsafe.Add(mBase, _consts[1022]))
								v271 = *(*float64)(unsafe.Add(mBase, uint32(v243)+uint32(_consts[1023])))
								v272 = base.F64_add(base.F64_mul(v233, v267), v271)
								v273 = base.F64_add(v265, v272)
								v278 = *(*float64)(unsafe.Add(mBase, _consts[1024]))
								v279 = base.F64_mul(v265, v278)
								v280 = base.F64_mul(v261, v278)
								v284 = base.F64_mul(v261, v280)
								v285 = base.F64_add(v273, v284)
								v289 = base.F64_mul(v265, v279)
								v292 = *(*float64)(unsafe.Add(mBase, _consts[1025]))
								v295 = *(*float64)(unsafe.Add(mBase, _consts[1026]))
								v299 = *(*float64)(unsafe.Add(mBase, _consts[1027]))
								v302 = *(*float64)(unsafe.Add(mBase, _consts[1028]))
								v307 = *(*float64)(unsafe.Add(mBase, _consts[1029]))
								v310 = *(*float64)(unsafe.Add(mBase, _consts[1030]))
								v314 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v233, v235), v246), base.F64_add(v265, base.F64_sub(v272, v273))), base.F64_mul(v264, base.F64_add(v279, v280))), base.F64_add(v284, base.F64_sub(v273, v285))), base.F64_mul(base.F64_mul(v265, v289), base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v289, base.F64_add(base.F64_mul(v265, v292), v295)), base.F64_add(base.F64_mul(v265, v299), v302))), base.F64_add(base.F64_mul(v265, v307), v310))))
								v315 = base.F64_add(v285, v314)
								*(*float64)(unsafe.Add(mBase, uint32(v16+int32(8)))) = base.F64_add(v314, base.F64_sub(v285, v315))
								v322 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v315) & int64(-134217728))
								v323 = base.F64_mul(v217, v322)
								v326 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
								v337 = F_top12_2(m, v323)
								mBase = m.M
								v339 = v337 & int32(2047)
								v341 = F_top12_2(m, float64(5.551115123125783e-17))
								mBase = m.M
								v344 = F_top12_2(m, float64(512))
								mBase = m.M
								if base.Ui32(v339-v341) < base.Ui32(v344-v341) {
									v361 = v339
									v363 = int32(0)
									v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
									v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
									v368 = base.F64_add(base.F64_mul(v323, v364), v367)
									v369 = base.F64_sub(v368, v367)
									v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
									v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
									v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
									v379 = base.F64_mul(v378, v378)
									v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
									v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
									v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
									v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
									v395 = base.I64_reinterpret_f64(v368)
									v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
									v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
									v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
									v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
									v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
									if v361 != 0 {
										v416 = base.F64_reinterpret_i64(v414)
										v427 = base.F64_add(base.F64_mul(v416, v406), v416)
									} else {
										v415 = F_specialcase_2(m, v406, v414, v395)
										mBase = m.M
										v427 = v415
									}
								} else {
									if base.Ui32(v341) <= base.Ui32(v339) {
										v353 = F_top12_2(m, float64(1024))
										mBase = m.M
										if base.Ui32(v339) < base.Ui32(v353) {
											v361 = int32(0)
											v363 = int32(0)
											v364 = *(*float64)(unsafe.Add(mBase, _consts[1031]))
											v367 = *(*float64)(unsafe.Add(mBase, _consts[1032]))
											v368 = base.F64_add(base.F64_mul(v323, v364), v367)
											v369 = base.F64_sub(v368, v367)
											v371 = *(*float64)(unsafe.Add(mBase, _consts[1033]))
											v374 = *(*float64)(unsafe.Add(mBase, _consts[1034]))
											v378 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v217), v322), base.F64_mul(l1, base.F64_add(v326, base.F64_sub(v315, v322)))), base.F64_add(base.F64_mul(v369, v371), base.F64_add(base.F64_mul(v369, v374), v323)))
											v379 = base.F64_mul(v378, v378)
											v382 = *(*float64)(unsafe.Add(mBase, _consts[1035]))
											v385 = *(*float64)(unsafe.Add(mBase, _consts[1036]))
											v389 = *(*float64)(unsafe.Add(mBase, _consts[1037]))
											v392 = *(*float64)(unsafe.Add(mBase, _consts[1038]))
											v395 = base.I64_reinterpret_f64(v368)
											v400 = base.I32_wrap_i64(v395) << (uint(int32(4)) % 32) & int32(2032)
											v403 = *(*float64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1039])))
											v406 = base.F64_add(base.F64_mul(base.F64_mul(v379, v379), base.F64_add(base.F64_mul(v378, v382), v385)), base.F64_add(base.F64_mul(v379, base.F64_add(base.F64_mul(v378, v389), v392)), base.F64_add(v403, v378)))
											v409 = *(*int64)(unsafe.Add(mBase, uint32(v400)+uint32(_consts[1040])))
											v414 = v409 + (v395+base.I64_extend_i32_u(v214))<<(uint(int64(45))%64)
											if v361 != 0 {
												v416 = base.F64_reinterpret_i64(v414)
												v427 = base.F64_add(base.F64_mul(v416, v406), v416)
											} else {
												v415 = F_specialcase_2(m, v406, v414, v395)
												mBase = m.M
												v427 = v415
											}
										} else {
											if int64(-1) < base.I64_reinterpret_f64(v323) {
												v360 = F___math_oflow(m, v214)
												mBase = m.M
												v427 = v360
											} else {
												v359 = F___math_uflow(m, v214)
												mBase = m.M
												v427 = v359
											}
										}
									} else {
										v349 = base.F64_add(v323, float64(1))
										if v214 != 0 {
											v351 = base.F64_neg(v349)
										} else {
											v351 = v349
										}
										v427 = v351
									}
								}
								v432 = v427
							} else {
								if v179 == int64(4607182418800017408) {
									v432 = float64(1)
								} else {
									if base.Ui32(int32(957)) < base.Ui32(v27) {
										if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v25)) == base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
											v204 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
											mBase = m.M
											v432 = v204
										} else {
											v201 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
											mBase = m.M
											v432 = v201
										}
									} else {
										if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
											v191 = l1
										} else {
											v191 = base.F64_neg(l1)
										}
										v432 = base.F64_add(v191, float64(1))
									}
								}
							}
						} else {
							v167 = base.F64_sub(l0, l0)
							v432 = base.F64_div(v167, v167)
						}
					}
				} else {
					v80 = base.F64_mul(l0, l0)
					if int64(-1) < v31 {
						v120 = v80
					} else {
						v92 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(52))%64))) & int32(2047)
						if base.Ui32(v92) < base.Ui32(int32(1023)) {
							v115 = int32(0)
						} else {
							if base.Ui32(int32(1075)) < base.Ui32(v92) {
								v115 = int32(2)
							} else {
								v103 = int64(1) << (uint(base.I64_extend_i32_u(int32(1075)-v92)) % 64)
								if (v103+int64(-1))&v30 != int64(0) {
									v115 = int32(0)
								} else {
									if v103&v30 == int64(0) {
										v114 = int32(2)
									} else {
										v114 = int32(1)
									}
									v115 = v114
								}
							}
						}
						if v115 == int32(1) {
							v119 = base.F64_neg(v80)
						} else {
							v119 = v80
						}
						v120 = v119
					}
					if int64(-1) < v30 {
						v432 = v120
					} else {
						v124 = base.F64_div(float64(1), v120)
						v126 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v126-int32(16))+8)) = v124
						v432 = v124
					}
				}
			} else {
				v48 = float64(1)
				if v31 == int64(4607182418800017408) {
					v432 = v48
				} else {
					v52 = v30 << (uint(int64(1)) % 64)
					if v52 == int64(0) {
						v432 = v48
					} else {
						v56 = v31 << (uint(int64(1)) % 64)
						if base.Ui64(int64(-9007199254740992)) < base.Ui64(v56) {
							v432 = base.F64_add(l0, l1)
						} else {
							if base.Ui64(v52) < base.Ui64(int64(-9007199254740991)) {
								if v56 == int64(9214364837600034816) {
									v432 = v48
								} else {
									if base.B2i32(base.Ui64(v56) < base.Ui64(int64(9214364837600034816)))^base.B2i32(v30 < int64(0)) != 0 {
										v71 = float64(0)
									} else {
										v71 = base.F64_mul(l1, l1)
									}
									v432 = v71
								}
							} else {
								v432 = base.F64_add(l0, l1)
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v16 + int32(16)
	return v432
}
func F_prefetchCommandQueueKeys(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v215 int32
	_ = v215
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v16&int32(2097152) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F__serverAssert(m, int32(_a1693), int32(_a1630), int32(4193))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v16 | int32(2097152)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v30 = v14 - (v23*int32(40)+int32(15))&int32(-16)
	m.G0 = v30
	if v23 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L2
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v34 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	if base.Ui32(v61) <= base.Ui32(v60) {
		v125 = v59
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v45 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v47 == v45 {
		v59 = v45
		goto L6
	} else {
		goto L12
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+290)))
	if v37&int32(4) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v43 = F_addKeysToIncrFindBatch(m, l0, v34, v40, v41, v30, int32(0), v23)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v59 = v43
	goto L6
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v50&int32(8192) == int32(0) {
		v59 = v45
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v50&int32(393216) != 0 {
		v59 = v45
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v59 = v45
	goto L6
L16:
	;
	if v125 < int32(2) {
		goto L4
	} else {
		goto L34
	}
L17:
	;
	if v23 <= v59 {
		v125 = v59
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v67 = v59
	v71 = v60
	v72 = v61
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v78 = v75 + v71*int32(40)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v79 | int32(2097152)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	if v83 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v125 = v114
	goto L16
L21:
	;
	v117 = v71 + int32(1)
	if base.Ui32(v115&int32(65535)) <= base.Ui32(v117) {
		v125 = v114
		goto L16
	} else {
		goto L32
	}
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v111 = F_addKeysToIncrFindBatch(m, l0, v83, v109, v110, v30, v67, v23)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L10
	} else {
		goto L31
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v91 == int32(0) {
		v114 = v67
		v115 = v72
		goto L21
	} else {
		goto L26
	}
L24:
	;
	if v79&int32(262144) == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if v79&int32(8192) == int32(0) {
		v114 = v67
		v115 = v72
		goto L21
	} else {
		goto L27
	}
L27:
	;
	if v79&int32(393216) != 0 {
		v114 = v67
		v115 = v72
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v100 == int32(0) {
		v114 = v67
		v115 = v72
		goto L21
	} else {
		goto L29
	}
L29:
	;
	F__serverAssert(m, int32(_a1694), int32(_a1630), int32(4206))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+60)))
	v114 = v111
	v115 = v113
	goto L21
L32:
	;
	if v114 < v23 {
		v67 = v114
		v71 = v117
		v72 = v115
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L20
L34:
	;
	v135 = int32(0)
	v138 = v135
	v144 = v135
	goto L35
L35:
	;
	v151 = F_hashtableIncrementalFindStep(m, v30+v138*int32(40))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	v160 = v157
	goto L40
L37:
	;
	v153 = v144 + v151
	v155 = v138 + int32(1)
	if v155 != v125 {
		v138 = v155
		v144 = v153
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v157 = int32(0)
	if v153 != 0 {
		v138 = v157
		v144 = v157
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v175 = F_hashtableIncrementalFindGetResult(m, v30+v160*int32(40), v14+int32(12))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L4
L42:
	;
	v184 = v160 + int32(1)
	if v184 != v125 {
		v160 = v184
		goto L40
	} else {
		goto L46
	}
L43:
	;
	if v175 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v180 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v181 = F_objectGetVal(m, v179)
	mBase = m.M
	goto L42
L46:
	;
	goto L41
L47:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prefixmatchlen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(1)
	if l1 != v12 {
		if l1 < int32(1) {
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v28
			v33 = F_stringmatchlen_impl(m, l0, l1, l2, l3, l4, v10+int32(12), v28)
			mBase = m.M
			v34 = v33
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))))
			if v24 != int32(42) {
				v34 = int32(0)
			} else {
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v28
				v33 = F_stringmatchlen_impl(m, l0, l1, l2, l3, l4, v10+int32(12), v28)
				mBase = m.M
				v34 = v33
			}
		}
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v15 != int32(42) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1+int32(-1)))))
			if v24 != int32(42) {
				v34 = int32(0)
			} else {
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v28
				v33 = F_stringmatchlen_impl(m, l0, l1, l2, l3, l4, v10+int32(12), v28)
				mBase = m.M
				v34 = v33
			}
		} else {
			v34 = v12
		}
	}
	m.G0 = v10 + int32(16)
	return v34
}
func F_prepareReplicasToWrite(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[78]))
	v11 = v6 + int32(8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v12
	goto L1
L1:
	;
	v16 = int32(0)
	v18 = v6 + int32(8)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v20 == v16 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v67
L3:
	;
	if v20 == int32(0) {
		v67 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20+base.B2i32(v23 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v29
	goto L4
L6:
	;
	v34 = v16
	v35 = v20
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+205)))
	if v37&int32(32) != 0 {
		v51 = v34
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v67 = v51
	goto L2
L9:
	;
	v53 = v6 + int32(8)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v55 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == int32(6) {
		v51 = v34
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v44 = F_prepareClientToWrite(m, v36)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v51 = v34 + base.B2i32(v44 != int32(-1))
	goto L9
L14:
	;
	if v55 != 0 {
		v34 = v51
		v35 = v55
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v55+base.B2i32(v58 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v64
	goto L15
L17:
	;
	goto L8
}
func F_printHelpMessage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_sdsempty(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v19
	v23 = F_sdscatfmt(m, v13, int32(_a2032), v11)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_createObject(m, v17, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v29 = F_listAddNodeTail(m, v28, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[658]))
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_printCommandHelp(m, int32(_a2033))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[659]))
	if v42 == v41 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_printCommandHelp(m, int32(_a2034))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_scriptingEngineDebuggerFlushLogs(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L13:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	v52 = v45
	v53 = v42
	v54 = v46
	goto L14
L14:
	;
	v58 = v54 + v52*int32(40)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
	if v59 != 0 {
		v66 = v53
		v67 = v54
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L12
L16:
	;
	v69 = v52 + int32(1)
	if base.Ui32(v69) < base.Ui32(v66) {
		v52 = v69
		v53 = v66
		v54 = v67
		goto L14
	} else {
		goto L19
	}
L17:
	;
	F_printCommandHelp(m, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, _consts[659]))
	v65 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	v66 = v63
	v67 = v65
	goto L16
L19:
	;
	goto L15
L20:
	;
	m.G0 = v11 + int32(16)
	return int32(1)
}
func F_printf_core(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v514 int32
	_ = v514
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v562 int64
	_ = v562
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v604 int64
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v640 int64
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int64
	_ = v675
	var v679 int64
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int64
	_ = v706
	var v710 int64
	_ = v710
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int64
	_ = v727
	var v734 int64
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	var v743 int64
	_ = v743
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int64
	_ = v755
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int64
	_ = v792
	var v798 int32
	_ = v798
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int64
	_ = v825
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v963 int32
	_ = v963
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v988 float64
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1011 int32
	_ = v1011
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1056 int32
	_ = v1056
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1234 int32
	_ = v1234
	var v1278 int32
	_ = v1278
	v8 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(64)
	m.G0 = v29
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l1
	v35 = v29 + int32(40)
	v39 = l1
	v48 = v8
	v49 = v8
	goto L5
L1:
	;
	m.G0 = v29 + int32(64)
	return v1278
L2:
	;
	v1278 = int32(-1)
	goto L1
L3:
	;
	goto L275
L4:
	;
	v1234 = int32(61)
	goto L3
L5:
	;
	v66 = v39
	v75 = v48
	v76 = v49
	v77 = int32(0)
	goto L7
L6:
	;
	v1278 = int32(0)
	goto L1
L7:
	;
	if v76^int32(2147483647) < v77 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v94 = v77 + v76
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v95 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L8
L11:
	;
	v1125 = v1122 - v1118
	if v1125 < v1121 {
		goto L261
	} else {
		goto L262
	}
L12:
	;
	v1118 = v1111
	v1119 = v1112
	v1120 = v1113
	v1121 = v1114
	v1122 = v35
	v1123 = v1115
	goto L11
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+39)) = uint8(v1106)
	v1118 = v29 + int32(39)
	v1119 = v584
	v1120 = v580
	v1121 = int32(1)
	v1122 = v35
	v1123 = v585
	goto L11
L14:
	;
	v1234 = int32(28)
	goto L3
L15:
	;
	if l0 != 0 {
		v1278 = v94
		goto L1
	} else {
		goto L247
	}
L16:
	;
	v110 = v66
	v112 = v95
	goto L17
L17:
	;
	v125 = v112 & int32(255)
	if v125 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	v110 = v110 + int32(1)
	v112 = v993
	goto L17
L20:
	;
	v190 = v176 - v66
	v192 = v94 ^ int32(2147483647)
	if v192 < v190 {
		goto L4
	} else {
		goto L29
	}
L21:
	;
	if v125 != int32(37) {
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v165 = v110
	v176 = v110
	goto L20
L23:
	;
	v140 = v110
	v142 = v110
	goto L24
L24:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v154 == int32(37) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v165 = v161
	v176 = v158
	goto L20
L26:
	;
	v158 = v140 + int32(1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+2)))
	v161 = v142 + int32(2)
	if v159 == int32(37) {
		v140 = v158
		v142 = v161
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v165 = v142
	v176 = v140
	goto L20
L28:
	;
	goto L25
L29:
	;
	if l0 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v190 != 0 {
		v66 = v165
		v76 = v94
		v77 = v190
		goto L7
	} else {
		goto L34
	}
L31:
	;
	F_out(m, l0, v66, v190)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v165
	v202 = v165 + int32(1)
	v203 = int32(-1)
	v204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165)+1)))
	v206 = v204 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v206) {
		v215 = v75
		v216 = v202
		v217 = v203
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v216
	v220 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216))))
	v222 = v220 + int32(-32)
	if base.Ui32(v222) <= base.Ui32(int32(31)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
	if v209 != int32(36) {
		v215 = v75
		v216 = v202
		v217 = v203
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v215 = int32(1)
	v216 = v165 + int32(3)
	v217 = v206
	goto L35
L38:
	;
	if v289 != int32(42) {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	v225 = int32(0)
	v227 = int32(1) << (uint(v222) % 32)
	if v227&int32(75913) == v225 {
		v286 = v216
		v288 = v225
		v289 = v220
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v286 = v216
	v288 = int32(0)
	v289 = v220
	goto L38
L41:
	;
	v233 = v227
	v244 = v216
	v249 = v225
	goto L42
L42:
	;
	v259 = v244 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v259
	v261 = v233 | v249
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v244)+1)))
	v264 = v262 + int32(-32)
	if base.Ui32(int32(32)) <= base.Ui32(v264) {
		v286 = v259
		v288 = v261
		v289 = v262
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v286 = v259
	v288 = v261
	v289 = v262
	goto L38
L44:
	;
	v268 = int32(1) << (uint(v264) % 32)
	if v268&int32(75913) != 0 {
		v233 = v268
		v244 = v259
		v249 = v261
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v402 == int32(46) {
		goto L73
	} else {
		goto L74
	}
L47:
	;
	v344 = v29 + int32(60)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v352 = int32(*(*int8)(unsafe.Add(mBase, uint32(v351))))
	v354 = v352 + int32(-48)
	if base.Ui32(v354) <= base.Ui32(int32(9)) {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	v299 = int32(*(*int8)(unsafe.Add(mBase, uint32(v286)+1)))
	v301 = v299 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v301) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v332
	if int32(-1) < v335 {
		v395 = v332
		v396 = v333
		v398 = v288
		v399 = v335
		goto L46
	} else {
		goto L59
	}
L50:
	;
	if v215 != 0 {
		goto L14
	} else {
		goto L56
	}
L51:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+2)))
	if v304 != int32(36) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	if l0 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v332 = v286 + int32(3)
	v333 = int32(1)
	v335 = v317
	goto L49
L54:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l3+v301<<(uint(int32(3))%32))))
	v317 = v316
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v301<<(uint(int32(2))%32)))) = int32(10)
	v317 = int32(0)
	goto L53
L56:
	;
	v322 = v286 + int32(1)
	if l0 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v326 + int32(4)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	v332 = v322
	v333 = int32(0)
	v335 = v330
	goto L49
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v322
	v324 = int32(0)
	v395 = v322
	v396 = v324
	v398 = v288
	v399 = v324
	goto L46
L59:
	;
	v395 = v332
	v396 = v333
	v398 = v288 | int32(8192)
	v399 = int32(0) - v335
	goto L46
L60:
	;
	if v391 < int32(0) {
		goto L4
	} else {
		goto L71
	}
L61:
	;
	v359 = int32(0)
	v360 = v351
	v361 = v354
	goto L63
L62:
	;
	v391 = int32(0)
	goto L60
L63:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v359) {
		v376 = int32(-1)
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v391 = v376
	goto L60
L65:
	;
	v378 = v360 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v378
	v380 = int32(*(*int8)(unsafe.Add(mBase, uint32(v360)+1)))
	v382 = v380 + int32(-48)
	if base.Ui32(v382) < base.Ui32(int32(10)) {
		v359 = v376
		v360 = v378
		v361 = v382
		goto L63
	} else {
		goto L70
	}
L66:
	;
	v369 = v359 * int32(10)
	if base.Ui32(v369^int32(2147483647)) < base.Ui32(v361) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v374 = int32(-1)
	goto L69
L68:
	;
	v374 = v361 + v369
	goto L69
L69:
	;
	v376 = v374
	goto L65
L70:
	;
	goto L64
L71:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v395 = v394
	v396 = v215
	v398 = v288
	v399 = v391
	goto L46
L72:
	;
	v503 = v498
	v514 = int32(0)
	goto L98
L73:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)))
	if v406 != int32(42) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v498 = v395
	v500 = int32(-1)
	v501 = int32(0)
	goto L72
L75:
	;
	v444 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v395 + v444
	v449 = v29 + int32(60)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v457 = int32(*(*int8)(unsafe.Add(mBase, uint32(v456))))
	v459 = v457 + int32(-48)
	if base.Ui32(v459) <= base.Ui32(int32(9)) {
		goto L88
	} else {
		goto L89
	}
L76:
	;
	v409 = int32(*(*int8)(unsafe.Add(mBase, uint32(v395)+2)))
	v411 = v409 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v411) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v438
	v498 = v438
	v500 = v440
	v501 = base.B2i32(int32(-1) < v440)
	goto L72
L78:
	;
	if v396 != 0 {
		goto L14
	} else {
		goto L84
	}
L79:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+3)))
	if v414 != int32(36) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	if l0 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v438 = v395 + int32(4)
	v440 = v427
	goto L77
L82:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l3+v411<<(uint(int32(3))%32))))
	v427 = v426
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v411<<(uint(int32(2))%32)))) = int32(10)
	v427 = int32(0)
	goto L81
L84:
	;
	v431 = v395 + int32(2)
	if l0 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v433 + int32(4)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v438 = v431
	v440 = v437
	goto L77
L86:
	;
	v438 = v431
	v440 = int32(0)
	goto L77
L87:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v498 = v497
	v500 = v496
	v501 = v444
	goto L72
L88:
	;
	v464 = int32(0)
	v465 = v456
	v466 = v459
	goto L90
L89:
	;
	v496 = int32(0)
	goto L87
L90:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v464) {
		v481 = int32(-1)
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v496 = v481
	goto L87
L92:
	;
	v483 = v465 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v483
	v485 = int32(*(*int8)(unsafe.Add(mBase, uint32(v465)+1)))
	v487 = v485 + int32(-48)
	if base.Ui32(v487) < base.Ui32(int32(10)) {
		v464 = v481
		v465 = v483
		v466 = v487
		goto L90
	} else {
		goto L97
	}
L93:
	;
	v474 = v464 * int32(10)
	if base.Ui32(v474^int32(2147483647)) < base.Ui32(v466) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v479 = int32(-1)
	goto L96
L95:
	;
	v479 = v466 + v474
	goto L96
L96:
	;
	v481 = v479
	goto L92
L97:
	;
	goto L91
L98:
	;
	v528 = int32(28)
	v529 = int32(*(*int8)(unsafe.Add(mBase, uint32(v503))))
	if base.Ui32(v529+int32(-123)) < base.Ui32(int32(-58)) {
		v1234 = v528
		goto L3
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v535
	if v541 == int32(27) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v535 = v503 + int32(1)
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v514*int32(58))+uint32(_consts[1068]))))
	if base.Ui32((v541+int32(-1))&int32(255)) < base.Ui32(int32(8)) {
		v503 = v535
		v514 = v541
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v576&int32(32) != 0 {
		goto L2
	} else {
		goto L114
	}
L103:
	;
	if int32(-1) < v217 {
		v1234 = v528
		goto L3
	} else {
		goto L112
	}
L104:
	;
	if v541 == int32(0) {
		v1234 = v528
		goto L3
	} else {
		goto L105
	}
L105:
	;
	if v217 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L110
	}
L107:
	;
	if l0 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v562 = *(*int64)(unsafe.Add(mBase, uint32(l3+v217<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v562
	goto L102
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v217<<(uint(int32(2))%32)))) = v541
	v39 = v535
	v48 = v396
	v49 = v94
	goto L5
L110:
	;
	F_pop_arg(m, v29+int32(48), v541, l2, l6)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L32
	} else {
		goto L111
	}
L111:
	;
	goto L102
L112:
	;
	v572 = int32(0)
	if l0 == v572 {
		v66 = v535
		v75 = v396
		v76 = v94
		v77 = v572
		goto L7
	} else {
		goto L113
	}
L113:
	;
	goto L102
L114:
	;
	v580 = v398 & int32(-65537)
	if v398&int32(8192) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v583 = v580
	goto L117
L116:
	;
	v583 = v398
	goto L117
L117:
	;
	v584 = int32(0)
	v585 = int32(_a2780)
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	v587 = base.I32_extend8_s(v586)
	if v586&int32(15) == int32(3) {
		goto L135
	} else {
		goto L136
	}
L118:
	;
	if v501&base.B2i32(v500 < int32(0)) != 0 {
		goto L4
	} else {
		goto L244
	}
L119:
	;
	F_pad(m, l0, int32(32), v399, v963, v583^int32(8192))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L32
	} else {
		goto L240
	}
L120:
	;
	v864 = int32(0)
	v866 = v848
	goto L218
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(0)
	*(*uint32)(unsafe.Add(mBase, uint32(v29)+8)) = uint32(v825)
	v843 = v29 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v843
	v848 = v843
	v849 = int32(-1)
	goto L120
L122:
	;
	if v500 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L123:
	;
	v825 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if base.B2i32(v825 == int64(0)) == int32(0) {
		goto L121
	} else {
		goto L213
	}
L124:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v809 != 0 {
		goto L200
	} else {
		goto L201
	}
L125:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
	v1106 = v808
	goto L13
L126:
	;
	if v501&base.B2i32(v790 < int32(0)) != 0 {
		goto L4
	} else {
		goto L190
	}
L127:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v727) {
		goto L180
	} else {
		goto L181
	}
L128:
	;
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if int64(-1) < v706 {
		goto L171
	} else {
		goto L172
	}
L129:
	;
	v673 = int32(0)
	v674 = int32(_a2780)
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v675 == int64(0) {
		v695 = v35
		goto L162
	} else {
		goto L163
	}
L130:
	;
	v632 = int32(0)
	v633 = int32(_a2780)
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v634 == int64(0) {
		v659 = v35
		goto L154
	} else {
		goto L155
	}
L131:
	;
	v622 = int32(8)
	if base.Ui32(v622) < base.Ui32(v500) {
		goto L150
	} else {
		goto L151
	}
L132:
	;
	v605 = int32(0)
	switch v514 {
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
		v66 = v535
		v75 = v396
		v76 = v94
		v77 = v605
		goto L7
	case 6:
		goto L144
	case 7:
		goto L143
	}
L133:
	;
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	v725 = int32(0)
	v726 = int32(_a2780)
	v727 = v604
	goto L127
L134:
	;
	switch v595 + int32(-65) {
	case 0, 4, 5, 6:
		goto L118
	case 1, 3:
		v1118 = v66
		v1119 = v584
		v1120 = v583
		v1121 = v500
		v1122 = v35
		v1123 = v585
		goto L11
	case 2:
		goto L123
	default:
		goto L141
	}
L135:
	;
	v594 = v587 & int32(-45)
	goto L137
L136:
	;
	v594 = v587
	goto L137
L137:
	;
	if v514 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v595 = v594
	goto L140
L139:
	;
	v595 = v587
	goto L140
L140:
	;
	switch v595 + int32(-88) {
	case 0, 32:
		v629 = v595
		v630 = v583
		v631 = v500
		goto L130
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 25, 26, 28, 30, 31:
		v1118 = v66
		v1119 = v584
		v1120 = v583
		v1121 = v500
		v1122 = v35
		v1123 = v585
		goto L11
	case 9, 13, 14, 15:
		goto L118
	case 11:
		goto L125
	case 12, 17:
		goto L128
	case 22:
		goto L132
	case 23:
		goto L129
	case 24:
		goto L131
	case 27:
		goto L124
	case 29:
		goto L133
	default:
		goto L134
	}
L141:
	;
	if v595 == int32(83) {
		goto L122
	} else {
		goto L142
	}
L142:
	;
	v1111 = v66
	v1112 = v584
	v1113 = v583
	v1114 = v500
	v1115 = v585
	goto L12
L143:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v619))) = base.I64_extend_i32_s(v94)
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L144:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v617))) = v94
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L145:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v615))) = uint8(v94)
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L146:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v613))) = uint16(v94)
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L147:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v610))) = base.I64_extend_i32_s(v94)
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L148:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v608))) = v94
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L149:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = v94
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v605
	goto L7
L150:
	;
	v625 = v500
	goto L152
L151:
	;
	v625 = v622
	goto L152
L152:
	;
	v629 = int32(120)
	v630 = v583 | int32(8)
	v631 = v625
	goto L130
L153:
	;
	if v634 == int64(0) {
		v787 = v659
		v788 = v632
		v789 = v630
		v790 = v631
		v791 = v633
		v792 = v634
		goto L126
	} else {
		goto L159
	}
L154:
	;
	goto L153
L155:
	;
	v640 = v634
	v641 = v35
	goto L156
L156:
	;
	v645 = v641 + int32(-1)
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v640)&int32(15))+uint32(_consts[1069]))))
	v652 = v651 | v629&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v645))) = uint8(v652)
	if base.Ui64(int64(15)) < base.Ui64(v640) {
		v640 = int64(base.Ui64(v640) >> (uint(int64(4)) % 64))
		v641 = v645
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v659 = v645
	goto L154
L158:
	;
	goto L157
L159:
	;
	if v630&int32(8) == int32(0) {
		v787 = v659
		v788 = v632
		v789 = v630
		v790 = v631
		v791 = v633
		v792 = v634
		goto L126
	} else {
		goto L160
	}
L160:
	;
	v787 = v659
	v788 = int32(2)
	v789 = v630
	v790 = v631
	v791 = int32(base.Ui32(v629)>>(uint(int32(4))%32)) + int32(_a2780)
	v792 = v634
	goto L126
L161:
	;
	if v583&int32(8) == int32(0) {
		v787 = v695
		v788 = v673
		v789 = v583
		v790 = v500
		v791 = v674
		v792 = v675
		goto L126
	} else {
		goto L167
	}
L162:
	;
	goto L161
L163:
	;
	v679 = v675
	v680 = v35
	goto L164
L164:
	;
	v683 = v680 + int32(-1)
	v688 = base.I32_wrap_i64(v679)&int32(7) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v683))) = uint8(v688)
	if base.Ui64(int64(7)) < base.Ui64(v679) {
		v679 = int64(base.Ui64(v679) >> (uint(int64(3)) % 64))
		v680 = v683
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v695 = v683
	goto L162
L166:
	;
	goto L165
L167:
	;
	v701 = v35 - v695
	if v701 < v500 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v705 = v500
	goto L170
L169:
	;
	v705 = v701 + int32(1)
	goto L170
L170:
	;
	v787 = v695
	v788 = v673
	v789 = v583
	v790 = v705
	v791 = v674
	v792 = v675
	goto L126
L171:
	;
	if v583&int32(2048) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v710 = int64(0) - v706
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v710
	v725 = int32(1)
	v726 = int32(_a2780)
	v727 = v710
	goto L127
L173:
	;
	v723 = v583 & int32(1)
	if v723 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v725 = int32(1)
	v726 = int32(_a2781)
	v727 = v706
	goto L127
L175:
	;
	v724 = int32(_a2782)
	goto L177
L176:
	;
	v724 = int32(_a2780)
	goto L177
L177:
	;
	v725 = v723
	v726 = v724
	v727 = v706
	goto L127
L178:
	;
	v787 = v781
	v788 = v725
	v789 = v583
	v790 = v500
	v791 = v726
	v792 = v727
	goto L126
L179:
	;
	if v755 == int64(0) {
		v781 = v754
		goto L185
	} else {
		goto L186
	}
L180:
	;
	v734 = v727
	v735 = v35
	goto L182
L181:
	;
	v754 = v35
	v755 = v727
	goto L179
L182:
	;
	v741 = v735 + int32(-1)
	v742 = int64(10)
	v743 = base.I64_div_u_s(v734, v742)
	v749 = base.I32_wrap_i64(v734-v743*v742) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v749)
	if base.Ui64(int64(42949672959)) < base.Ui64(v734) {
		v734 = v743
		v735 = v741
		goto L182
	} else {
		goto L184
	}
L183:
	;
	v754 = v741
	v755 = v743
	goto L179
L184:
	;
	goto L183
L185:
	;
	goto L178
L186:
	;
	v763 = v754
	v765 = base.I32_wrap_i64(v755)
	goto L187
L187:
	;
	v769 = v763 + int32(-1)
	v770 = int32(10)
	v771 = base.I32_div_u_s(v765, v770)
	v776 = v765 - v771*v770 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v769))) = uint8(v776)
	if base.Ui32(int32(9)) < base.Ui32(v765) {
		v763 = v769
		v765 = v771
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v781 = v769
	goto L185
L189:
	;
	goto L188
L190:
	;
	if v501 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v798 = v789 & int32(-65537)
	goto L193
L192:
	;
	v798 = v789
	goto L193
L193:
	;
	if v792 != int64(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v805 = v35 - v787 + base.B2i32(v792 == int64(0))
	if v805 < v790 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	if v790 != 0 {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v1118 = v35
	v1119 = v788
	v1120 = v798
	v1121 = int32(0)
	v1122 = v35
	v1123 = v791
	goto L11
L197:
	;
	v807 = v790
	goto L199
L198:
	;
	v807 = v805
	goto L199
L199:
	;
	v1111 = v787
	v1112 = v788
	v1113 = v798
	v1114 = v807
	v1115 = v791
	goto L12
L200:
	;
	v811 = v809
	goto L202
L201:
	;
	v811 = int32(_a2668)
	goto L202
L202:
	;
	v812 = int32(2147483647)
	if base.Ui32(v500) < base.Ui32(v812) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v815 = v500
	goto L205
L204:
	;
	v815 = v812
	goto L205
L205:
	;
	v818 = F_memchr(m, v811, int32(0), v815)
	mBase = m.M
	if v818 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v821 = v811 + v820
	if v500 <= int32(-1) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	v820 = v818 - v811
	goto L209
L208:
	;
	v820 = v815
	goto L209
L209:
	;
	goto L206
L210:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	if v824 != 0 {
		goto L4
	} else {
		goto L212
	}
L211:
	;
	v1118 = v811
	v1119 = v584
	v1120 = v580
	v1121 = v820
	v1122 = v821
	v1123 = v585
	goto L11
L212:
	;
	v1118 = v811
	v1119 = v584
	v1120 = v580
	v1121 = v820
	v1122 = v821
	v1123 = v585
	goto L11
L213:
	;
	v1106 = int32(0)
	goto L13
L214:
	;
	v834 = int32(0)
	F_pad(m, l0, int32(32), v399, v834, v583)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L32
	} else {
		goto L216
	}
L215:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v848 = v833
	v849 = v500
	goto L120
L216:
	;
	v963 = v834
	goto L119
L217:
	;
	if v895 < int32(0) {
		v1234 = int32(61)
		goto L3
	} else {
		goto L227
	}
L218:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	if v878 == int32(0) {
		v895 = v864
		goto L217
	} else {
		goto L220
	}
L219:
	;
	v895 = v893
	goto L217
L220:
	;
	v882 = v29 + int32(4)
	if v882 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	if v886 < int32(0) {
		goto L2
	} else {
		goto L224
	}
L222:
	;
	v885 = F_wcrtomb(m, v882, v878, int32(0))
	mBase = m.M
	v886 = v885
	goto L221
L223:
	;
	v886 = int32(0)
	goto L221
L224:
	;
	if base.Ui32(v849-v864) < base.Ui32(v886) {
		v895 = v864
		goto L217
	} else {
		goto L225
	}
L225:
	;
	v893 = v886 + v864
	if base.Ui32(v893) < base.Ui32(v849) {
		v864 = v893
		v866 = v866 + int32(4)
		goto L218
	} else {
		goto L226
	}
L226:
	;
	goto L219
L227:
	;
	F_pad(m, l0, int32(32), v399, v895, v583)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L32
	} else {
		goto L228
	}
L228:
	;
	if v895 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v921 = v906
	v922 = int32(0)
	goto L231
L230:
	;
	v963 = int32(0)
	goto L119
L231:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	if v933 == int32(0) {
		v963 = v895
		goto L119
	} else {
		goto L233
	}
L232:
	;
	v963 = v895
	goto L119
L233:
	;
	v937 = v29 + int32(4)
	if v937 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v942 = v941 + v922
	if base.Ui32(v895) < base.Ui32(v942) {
		v963 = v895
		goto L119
	} else {
		goto L237
	}
L235:
	;
	v940 = F_wcrtomb(m, v937, v933, int32(0))
	mBase = m.M
	v941 = v940
	goto L234
L236:
	;
	v941 = int32(0)
	goto L234
L237:
	;
	F_out(m, l0, v29+int32(4), v941)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L32
	} else {
		goto L238
	}
L238:
	;
	if base.Ui32(v942) < base.Ui32(v895) {
		v921 = v921 + int32(4)
		v922 = v942
		goto L231
	} else {
		goto L239
	}
L239:
	;
	goto L232
L240:
	;
	if v963 < v399 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v983 = v399
	goto L243
L242:
	;
	v983 = v963
	goto L243
L243:
	;
	v66 = v535
	v75 = v396
	v76 = v94
	v77 = v983
	goto L7
L244:
	;
	v988 = *(*float64)(unsafe.Add(mBase, uint32(v29)+48))
	v989 = m.T0[l5].(func(*base.Module, int32, float64, int32, int32, int32, int32) int32)(m, l0, v988, v399, v500, v583, v595)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L32
	} else {
		goto L245
	}
L245:
	;
	if int32(0) <= v989 {
		v66 = v535
		v75 = v396
		v76 = v94
		v77 = v989
		goto L7
	} else {
		goto L246
	}
L246:
	;
	v1234 = int32(61)
	goto L3
L247:
	;
	if v75 == int32(0) {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	v1011 = int32(1)
	goto L250
L249:
	;
	if base.Ui32(v1011) < base.Ui32(int32(10)) {
		goto L255
	} else {
		goto L256
	}
L250:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1011<<(uint(int32(2))%32))))
	if v1028 == int32(0) {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	F_pop_arg(m, l3+v1011<<(uint(int32(3))%32), v1028, l2, l6)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L32
	} else {
		goto L253
	}
L253:
	;
	v1036 = int32(1)
	v1038 = v1011 + v1036
	if v1038 != int32(10) {
		v1011 = v1038
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v1278 = v1036
	goto L1
L255:
	;
	v1056 = v1011
	goto L257
L256:
	;
	v1278 = int32(1)
	goto L1
L257:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1056<<(uint(int32(2))%32))))
	if v1073 != 0 {
		goto L14
	} else {
		goto L259
	}
L259:
	;
	v1074 = int32(1)
	v1076 = v1056 + v1074
	if v1076 == int32(10) {
		v1278 = v1074
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v1056 = v1076
	goto L257
L261:
	;
	v1127 = v1121
	goto L263
L262:
	;
	v1127 = v1125
	goto L263
L263:
	;
	if v1119^int32(2147483647) < v1127 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v1132 = v1119 + v1127
	if v1132 < v399 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1134 = v399
	goto L267
L266:
	;
	v1134 = v1132
	goto L267
L267:
	;
	if v192 < v1134 {
		v1234 = int32(61)
		goto L3
	} else {
		goto L268
	}
L268:
	;
	F_pad(m, l0, int32(32), v1134, v1132, v1120)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L32
	} else {
		goto L269
	}
L269:
	;
	F_out(m, l0, v1123, v1119)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L32
	} else {
		goto L270
	}
L270:
	;
	F_pad(m, l0, int32(48), v1134, v1132, v1120^int32(65536))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L32
	} else {
		goto L271
	}
L271:
	;
	F_pad(m, l0, int32(48), v1127, v1125, int32(0))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L32
	} else {
		goto L272
	}
L272:
	;
	F_out(m, l0, v1118, v1125)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L32
	} else {
		goto L273
	}
L273:
	;
	F_pad(m, l0, int32(32), v1134, v1132, v1120^int32(8192))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L32
	} else {
		goto L274
	}
L274:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v66 = v1157
	v75 = v396
	v76 = v94
	v77 = v1134
	goto L7
L275:
	;
	*(*int32)(unsafe.Add(mBase, _consts[18])) = v1234
	goto L2
}
func F_processAnnotations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	v9 = m.G0
	v11 = v9 - int32(1120)
	m.G0 = v11
	v14 = F___ftello(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, _consts[903])) = v14
		v22 = F_fgets(m, v11+int32(96), int32(1024), l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
				v99 = F_iprintf(m, int32(_a2470), v11)
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					m.Env.Exit(m, int32(1))
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v26 = int32(1)
				v28 = *(*int64)(unsafe.Add(mBase, _consts[904]))
				if v28 == int64(0) {
					v89 = v26
					m.G0 = v11 + int32(1120)
					return v89
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
					if v31 != int32(978539555) {
						v89 = v26
						m.G0 = v11 + int32(1120)
						return v89
					} else {
						v34 = int32(9116376)
						*(*int32)(unsafe.Add(mBase, _consts[18])) = int32(0)
						v45 = F_strtox_2(m, v11+int32(96)|int32(4), v11+int32(92), int32(10), int64(2147483648))
						mBase = m.M
						v47 = *(*int32)(unsafe.Add(mBase, _consts[18]))
						if v47 != 0 {
							v104 = F_puts(m, int32(_a2471))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								m.Env.Exit(m, int32(1))
								mBase = m.M
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
							if v49 != int32(13) {
								v104 = F_puts(m, int32(_a2471))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									m.Env.Exit(m, int32(1))
									mBase = m.M
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v54 = *(*int64)(unsafe.Add(mBase, _consts[904]))
								if base.I64_extend_i32_s(base.I32_wrap_i64(v45)) <= v54 {
									v89 = int32(1)
									m.G0 = v11 + int32(1120)
									return v89
								} else {
									v58 = *(*int64)(unsafe.Add(mBase, _consts[903]))
									if v58 == int64(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v54
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
										v113 = F_iprintf(m, int32(_a2472), v11+int32(16))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											m.Env.Exit(m, int32(1))
											mBase = m.M
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if l2 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v11+int32(48)))) = v58
											*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v54
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
											v125 = F_iprintf(m, int32(_a2473), v11+int32(32))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												v128 = F_puts(m, int32(_a2474))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													m.Env.Exit(m, int32(1))
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
											if int32(-1) < v66 {
												v70 = F___lockfile(m, l0)
												mBase = m.M
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
												if v70 == int32(0) {
													v75 = v71
												} else {
													F___unlockfile(m, l0)
													mBase = m.M
													v75 = v71
												}
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
												v75 = v69
											}
											if int32(-1) < v75 {
												v83 = v75
											} else {
												v79 = F___errno_location(m)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(8)
												v83 = int32(-1)
											}
											v85 = *(*int64)(unsafe.Add(mBase, _consts[903]))
											v86 = F_ftruncate(m, v83, v85)
											mBase = m.M
											if v86 == int32(-1) {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l1
												v134 = *(*int64)(unsafe.Add(mBase, _consts[904]))
												*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v134
												v139 = F_iprintf(m, int32(_a2475), v11+int32(64))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													m.Env.Exit(m, int32(1))
													mBase = m.M
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v89 = int32(0)
												m.G0 = v11 + int32(1120)
												return v89
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
func F_propagateSyncSlotsFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
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
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	v8 = m.G0
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v11 != 0 {
		v12 = int32(9)
	} else {
		v12 = int32(7)
	}
	v19 = v8 - (v12<<(uint(int32(2))%32)+int32(15))&int32(112)
	m.G0 = v19
	v21 = int32(_a388)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, _consts[217]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v31
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v36 == int32(20) {
		v39 = int32(464)
	} else {
		v39 = int32(468)
	}
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v44
	v49 = F_createStringObject_1(m, l0+int32(112), int32(40))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v49
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
		if v52 != 0 {
			v55 = *(*int32)(unsafe.Add(mBase, _consts[219]))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v55
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-1)))))
			switch v60 & int32(7) {
			case 0:
				v77 = int32(base.Ui32(v60) >> (uint(int32(3)) % 32))
			case 1:
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(-3)))))
				v77 = v67
			case 2:
				v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(-5)))))
				v77 = v70
			case 3:
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-9))))
				v77 = v73
			case 4:
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(-17))))
				v77 = v76
			default:
				v77 = int32(0)
			}
			v78 = F_createStringObject_1(m, v52, v77)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
				v81 = v78
				v86 = int32(0)
				v90 = *(*int32)(unsafe.Add(mBase, _consts[177]))
				*(*int32)(unsafe.Add(mBase, _consts[177])) = v90 + int32(1)
				if v90 != 0 {
				} else {
					v98 = F_ustime(m)
					mBase = m.M
					v100 = int32(0)
					*(*int64)(unsafe.Add(mBase, _consts[178])) = v98
					v104 = base.I64_div_s(v98, int64(1000))
					*(*int64)(unsafe.Add(mBase, _consts[54])) = v104
					v108 = base.I64_div_s(v98, int64(1000000))
					*(*int64)(unsafe.Add(mBase, _consts[109])) = v108
					v111 = *(*int32)(unsafe.Add(mBase, _consts[179]))
					F_lrulfu_updateClockAndPolicy(m, v104, int32(base.Ui32(v111&int32(2))>>(uint(int32(1))%32)))
					mBase = m.M
					v119 = *(*int64)(unsafe.Add(mBase, _consts[54]))
					*(*int64)(unsafe.Add(mBase, _consts[12])) = v119
				}
				v123 = int32(-1)
				F_alsoPropagate(m, v123, v19, v12, int32(3), v123)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					v128 = int32(0)
					v130 = *(*int32)(unsafe.Add(mBase, _consts[177]))
					*(*int32)(unsafe.Add(mBase, _consts[177])) = v130 + int32(-1)
					F_postExecutionUnitOperations(m)
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						if v49 == int32(0) {
							if v81 == int32(0) {
								m.G0 = v8
								return
							} else {
								F_decrRefCount(m, v81)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return
								} else {
									m.G0 = v8
									return
								}
							}
						} else {
							F_decrRefCount(m, v49)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return
							} else {
								if v81 == int32(0) {
									m.G0 = v8
									return
								} else {
									F_decrRefCount(m, v81)
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return
									} else {
										m.G0 = v8
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v81 = int32(0)
			v86 = int32(0)
			v90 = *(*int32)(unsafe.Add(mBase, _consts[177]))
			*(*int32)(unsafe.Add(mBase, _consts[177])) = v90 + int32(1)
			if v90 != 0 {
			} else {
				v98 = F_ustime(m)
				mBase = m.M
				v100 = int32(0)
				*(*int64)(unsafe.Add(mBase, _consts[178])) = v98
				v104 = base.I64_div_s(v98, int64(1000))
				*(*int64)(unsafe.Add(mBase, _consts[54])) = v104
				v108 = base.I64_div_s(v98, int64(1000000))
				*(*int64)(unsafe.Add(mBase, _consts[109])) = v108
				v111 = *(*int32)(unsafe.Add(mBase, _consts[179]))
				F_lrulfu_updateClockAndPolicy(m, v104, int32(base.Ui32(v111&int32(2))>>(uint(int32(1))%32)))
				mBase = m.M
				v119 = *(*int64)(unsafe.Add(mBase, _consts[54]))
				*(*int64)(unsafe.Add(mBase, _consts[12])) = v119
			}
			v123 = int32(-1)
			F_alsoPropagate(m, v123, v19, v12, int32(3), v123)
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				v128 = int32(0)
				v130 = *(*int32)(unsafe.Add(mBase, _consts[177]))
				*(*int32)(unsafe.Add(mBase, _consts[177])) = v130 + int32(-1)
				F_postExecutionUnitOperations(m)
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return
				} else {
					if v49 == int32(0) {
						if v81 == int32(0) {
							m.G0 = v8
							return
						} else {
							F_decrRefCount(m, v81)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return
							} else {
								m.G0 = v8
								return
							}
						}
					} else {
						F_decrRefCount(m, v49)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							if v81 == int32(0) {
								m.G0 = v8
								return
							} else {
								F_decrRefCount(m, v81)
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return
								} else {
									m.G0 = v8
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
func F_protectClient(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v3 | int32(-2147483648)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == int32(0) {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
		v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, v7, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v16 = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
			v20 = m.T0[v19].(func(*base.Module, int32, int32, int32) int32)(m, v15, v16, v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_psubscribeCommand(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+205)))
	if v3&int32(16) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v14 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v8&int32(8) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_addReplyError(m, l0, int32(_a1783))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v33&int32(262144) != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v19 = int32(1)
	goto L8
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v19<<(uint(int32(2))%32))))
	v25 = F_pubsubSubscribePattern(m, l0, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v28 = v19 + int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v28 < v29 {
		v19 = v28
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v33 | int32(262144)
	v39 = int32(_a20)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v41 + int32(1)
	goto L12
}
func F_push_onecapture(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if l1 < v5 {
		v19 = l0 + l1<<(uint(int32(3))%32)
		v21 = v19 + int32(16)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(20))))
		switch v24 + int32(2) {
		case 0:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(3)
			*(*float64)(unsafe.Add(mBase, uint32(v34))) = base.F64_convert_i32_s(v28 - v29 + int32(1))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v39 + int32(16)
			return
		case 1:
			v43 = m.G3
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v48 = F_luaL_error(m, v44, v43+int32(_a2734), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				F_lua_pushlstring(m, v51, v52, v24)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					return
				}
			}
		default:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			F_lua_pushlstring(m, v51, v52, v24)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if l1 != 0 {
			v11 = m.G3
			v15 = F_luaL_error(m, v7, v11+int32(_a2732), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		} else {
			F_lua_pushlstring(m, v7, l2, l3-l2)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_puts(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1042]))
	if v2 <= v5 {
		v13 = int32(0)
	} else {
		v13 = int32(1)
	}
	v15 = F_fputs(m, l0, int32(_a2770))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v15 {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[1043]))
			if v23 == int32(10) {
				v41 = F___overflow(m, int32(_a2770), int32(10))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v45 = v41 >> (uint(int32(31)) % 32)
					if v13 != 0 {
					} else {
					}
					return v45
				}
			} else {
				v26 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[1044]))
				v29 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
				if v27 == v29 {
					v41 = F___overflow(m, int32(_a2770), int32(10))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v45 = v41 >> (uint(int32(31)) % 32)
						if v13 != 0 {
						} else {
						}
						return v45
					}
				} else {
					v31 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1044])) = v27 + int32(1)
					v36 = int32(10)
					*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v36)
					v45 = v31
					if v13 != 0 {
					} else {
					}
					return v45
				}
			}
		} else {
			v45 = int32(-1)
			if v13 != 0 {
			} else {
			}
			return v45
		}
	}
}
