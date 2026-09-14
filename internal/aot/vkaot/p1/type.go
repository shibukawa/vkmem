package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_checkType(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = int32(0)
	if l1 == v4 {
		v19 = v4
		return v19
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v8&int32(15) == l2 {
			v19 = v4
			return v19
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[507]))
			F_addReplyErrorObject(m, l0, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = int32(1)
				return v19
			}
		}
	}
}
func F_setTypeAddAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int64
	_ = v224
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v246 int64
	_ = v246
	var v251 int64
	_ = v251
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v267 int64
	_ = v267
	var v291 int64
	_ = v291
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int64
	_ = v354
	var v358 int64
	_ = v358
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int64
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v433 int32
	_ = v433
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v484 int64
	_ = v484
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int64
	_ = v498
	var v499 int64
	_ = v499
	var v500 int64
	_ = v500
	var v501 int64
	_ = v501
	var v506 int32
	_ = v506
	var v508 int64
	_ = v508
	var v511 int64
	_ = v511
	var v512 int32
	_ = v512
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int64
	_ = v570
	var v571 int32
	_ = v571
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int64
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int64
	_ = v647
	var v651 int64
	_ = v651
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int64
	_ = v691
	var v692 int64
	_ = v692
	var v693 int64
	_ = v693
	var v694 int64
	_ = v694
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v803 int32
	_ = v803
	v10 = m.G0
	v12 = v10 - int32(64)
	m.G0 = v12
	if l1 != 0 {
		v84 = l1
		v85 = l2
		v86 = l4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(64)
	return v803
L2:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v87)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		goto L24
	default:
		goto L21
	case 4:
		goto L22
	case 9:
		goto L23
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14&int32(240) != int32(96) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = v12 + int32(32)
	v39 = int32(0)
	if l3 <= int64(-1) {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v19)
	v22 = F_objectGetVal(m, l0)
	mBase = m.M
	v25 = F_intsetAdd(m, v22, l3, v12+int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	F_objectSetVal(m, l0, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
	if v31 == int32(0) {
		v803 = v19
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_maybeConvertIntset(m, l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
	v803 = v36
	goto L1
L11:
	;
	v84 = v38
	v85 = v83
	v86 = v39
	goto L2
L12:
	;
	v83 = v39
	goto L11
L14:
	;
	v64 = F_ull2string(m, v60, v61, v62)
	mBase = m.M
	if v64 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L15:
	;
	goto L17
L16:
	;
	v60 = v38
	v61 = int32(21)
	v62 = l3
	v63 = int32(0)
	goto L14
L17:
	;
	v51 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v51)
	v60 = v12 + int32(33)
	v61 = int32(20)
	v62 = int64(0) - l3
	v63 = int32(1)
	goto L14
L18:
	;
	v83 = v64 + v63
	goto L11
L20:
	;
	v803 = int32(1)
	goto L1
L21:
	;
	F__serverPanic_1(m, int32(_a2391), int32(222), int32(_a2392), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L302
	}
L22:
	;
	v175 = v12 + int32(16)
	v176 = int32(0)
	if base.Ui32(v85+int32(-21)) < base.Ui32(int32(-20)) {
		v310 = v176
		goto L66
	} else {
		goto L67
	}
L23:
	;
	v117 = F_objectGetVal(m, l0)
	mBase = m.M
	v118 = F_lpFirst(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L39
	}
L24:
	;
	if v86 != 0 {
		v96 = v84
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = F_objectGetVal(m, l0)
	mBase = m.M
	v101 = F_hashtableFindPositionForInsert(m, v97, v96, v12+int32(16), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L29
	}
L26:
	;
	v94 = F_sdsnewlen(m, v84, v85)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v96 = v94
	goto L25
L28:
	;
	if v96 == v84 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v101 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v96 != v84 {
		v108 = v96
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_hashtableInsertAtPosition(m, v97, v108, v12+int32(16))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	v106 = F_sdsdup(m, v96)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v108 = v106
	goto L31
L34:
	;
	goto L20
L35:
	;
	v803 = int32(0)
	goto L1
L36:
	;
	F_sdsfree(m, v96)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v127 = F_lpLength(m, v117)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L44
	}
L39:
	;
	if v118 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v122 = int32(0)
	v124 = F_lpFind(m, v117, v118, v84, v85, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v124 != 0 {
		v803 = v122
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v156 = F_lpLength(m, v117)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L58
	}
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[545]))
	if base.Ui32(v130) <= base.Ui32(v127) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if base.Ui32(v133) < base.Ui32(v85) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v117 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if base.B2i32(base.Ui32(v137+v85) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L43
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v137 = v136
	goto L48
L50:
	;
	v137 = int32(0)
	goto L48
L51:
	;
	if v84 != v12+int32(32) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v150 = F_lpAppend(m, v117, v84, v85)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L56
	}
L53:
	;
	v146 = F_lpAppendInteger(m, v117, l3)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_objectSetVal(m, l0, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L20
L56:
	;
	F_objectSetVal(m, l0, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	goto L20
L58:
	;
	v158 = int32(1)
	v161 = F_setTypeConvertAndExpand(m, l0, int32(2), v156+v158, v158)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v163 = F_objectGetVal(m, l0)
	mBase = m.M
	v164 = F_sdsnewlen(m, v84, v85)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v166 = F_hashtableAdd(m, v163, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v166 != 0 {
		v803 = int32(1)
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F__serverAssert(m, int32(_a2394), int32(_a2391), int32(177))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v335 = F_objectGetVal(m, l0)
	mBase = m.M
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	goto L99
L65:
	;
	if v310 == int32(0) {
		goto L64
	} else {
		goto L92
	}
L66:
	;
	goto L65
L67:
	;
	v188 = int32(1)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 != v188 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v310 = int32(1)
	goto L66
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v175))) = v291
	goto L68
L70:
	;
	if v189&int32(255) == int32(45) {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v193 = v189 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v193&int32(255)) {
		v310 = v176
		goto L66
	} else {
		goto L72
	}
L72:
	;
	if v175 == int32(0) {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v291 = base.I64_extend_i32_u(v193) & int64(255)
	goto L69
L74:
	;
	if base.Ui32(int32(8)) < base.Ui32((v212+int32(-49))&int32(255)) {
		v310 = v176
		goto L66
	} else {
		goto L77
	}
L75:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v211 = int32(2)
	v212 = v209
	v213 = v84 + int32(1)
	goto L74
L76:
	;
	v211 = v188
	v212 = v189
	v213 = v84
	goto L74
L77:
	;
	v224 = base.I64_extend_i32_u(v212+int32(-48)) & int64(255)
	if base.Ui32(v85) <= base.Ui32(v211) {
		v267 = v224
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v189&int32(255) != int32(45) {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v230 = v211
	v232 = v224
	v234 = v213
	goto L80
L80:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	if base.Ui32((v236+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v310 = v176
		goto L66
	} else {
		goto L82
	}
L81:
	;
	v267 = v257
	goto L78
L82:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v232) {
		v310 = v176
		goto L66
	} else {
		goto L83
	}
L83:
	;
	v246 = v232 * int64(10)
	v251 = base.I64_extend_i32_u(v236+int32(-48)) & int64(255)
	if base.Ui64(v251^int64(-1)) < base.Ui64(v246) {
		v310 = v176
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v255 = int32(1)
	v257 = v246 + v251
	v259 = v230 + v255
	if v259 != v85 {
		v230 = v259
		v232 = v257
		v234 = v234 + v255
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	if v267 < int64(0) {
		v310 = v176
		goto L66
	} else {
		goto L90
	}
L87:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v267) {
		v310 = v176
		goto L66
	} else {
		goto L88
	}
L88:
	;
	if v175 == int32(0) {
		goto L68
	} else {
		goto L89
	}
L89:
	;
	v291 = int64(0) - v267
	goto L69
L90:
	;
	if v175 == int32(0) {
		goto L68
	} else {
		goto L91
	}
L91:
	;
	v291 = v267
	goto L69
L92:
	;
	v319 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v319)
	v322 = F_objectGetVal(m, l0)
	mBase = m.M
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v326 = F_intsetAdd(m, v322, v323, v12+int32(15))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	F_objectSetVal(m, l0, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v330 == int32(0) {
		v803 = v319
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_maybeConvertIntset(m, l0)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	goto L20
L97:
	;
	v733 = F_objectGetVal(m, l0)
	mBase = m.M
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+4))
	goto L282
L98:
	;
	v339 = F_objectGetVal(m, l0)
	mBase = m.M
	v344 = v339 + int32(8)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v347 = v345 + int32(-1)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	switch v348 + int32(-4) {
	case 0:
		goto L103
	default:
		goto L102
	case 4:
		goto L104
	}
L99:
	;
	if v336 != 0 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v337 = int32(0)
	v729 = v337
	v732 = v337
	goto L97
L101:
	;
	if int64(-1) < v363 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v362 = int64(*(*int16)(unsafe.Add(mBase, uint32(v344+v347<<(uint(int32(1))%32)))))
	v363 = v362
	goto L101
L103:
	;
	v358 = int64(*(*int32)(unsafe.Add(mBase, uint32(v344+v347<<(uint(int32(2))%32)))))
	v363 = v358
	goto L101
L104:
	;
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v339+v345<<(uint(int32(3))%32))))
	v363 = v354
	goto L101
L105:
	;
	v494 = F_objectGetVal(m, l0)
	mBase = m.M
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	switch v495 + int32(-4) {
	case 0:
		goto L173
	default:
		goto L172
	case 4:
		goto L174
	}
L106:
	;
	v429 = int32(0)
	if base.Ui64(v363) < base.Ui64(int64(10)) {
		v486 = v429
		goto L140
	} else {
		goto L141
	}
L107:
	;
	v368 = int32(0)
	v370 = int64(0) - v363
	if base.Ui64(v370) < base.Ui64(int64(10)) {
		v420 = v368
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v493 = v424 + v425 + int32(1)
	goto L105
L109:
	;
	v424 = v420
	v425 = int32(1)
	goto L108
L110:
	;
	v373 = v370
	v374 = v368
	goto L111
L111:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v373) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v420 = v414
	goto L109
L113:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v373) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v424 = v374
	v425 = int32(2)
	goto L108
L115:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v373) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v424 = v374
	v425 = int32(3)
	goto L108
L117:
	;
	v414 = v374 + int32(12)
	v418 = base.I64_div_u_s(v373, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v373) {
		v373 = v418
		v374 = v414
		goto L111
	} else {
		goto L139
	}
L118:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v373) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v373) {
		goto L131
	} else {
		goto L132
	}
L120:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v373) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v373) {
		goto L128
	} else {
		goto L129
	}
L122:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v373) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v373) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v424 = v374
	v425 = int32(4)
	goto L108
L125:
	;
	v395 = int32(6)
	goto L127
L126:
	;
	v395 = int32(5)
	goto L127
L127:
	;
	v424 = v374
	v425 = v395
	goto L108
L128:
	;
	v400 = int32(8)
	goto L130
L129:
	;
	v400 = int32(7)
	goto L130
L130:
	;
	v424 = v374
	v425 = v400
	goto L108
L131:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v373) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v373) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v407 = int32(10)
	goto L135
L134:
	;
	v407 = int32(9)
	goto L135
L135:
	;
	v424 = v374
	v425 = v407
	goto L108
L136:
	;
	v412 = int32(12)
	goto L138
L137:
	;
	v412 = int32(11)
	goto L138
L138:
	;
	v424 = v374
	v425 = v412
	goto L108
L139:
	;
	goto L112
L140:
	;
	v493 = int32(1) + v486
	goto L105
L141:
	;
	v432 = v363
	v433 = v429
	goto L142
L142:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v432) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v486 = v480
	goto L140
L144:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v432) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v493 = int32(2) + v433
	goto L105
L146:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v432) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v493 = int32(3) + v433
	goto L105
L148:
	;
	v480 = v433 + int32(12)
	v484 = base.I64_div_u_s(v432, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v432) {
		v432 = v484
		v433 = v480
		goto L142
	} else {
		goto L170
	}
L149:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v432) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v432) {
		goto L162
	} else {
		goto L163
	}
L151:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v432) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v432) {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v432) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v432) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v493 = int32(4) + v433
	goto L105
L156:
	;
	v457 = int32(6)
	goto L158
L157:
	;
	v457 = int32(5)
	goto L158
L158:
	;
	v493 = v457 + v433
	goto L105
L159:
	;
	v463 = int32(8)
	goto L161
L160:
	;
	v463 = int32(7)
	goto L161
L161:
	;
	v493 = v463 + v433
	goto L105
L162:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v432) {
		goto L167
	} else {
		goto L168
	}
L163:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v432) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v471 = int32(10)
	goto L166
L165:
	;
	v471 = int32(9)
	goto L166
L166:
	;
	v493 = v471 + v433
	goto L105
L167:
	;
	v477 = int32(12)
	goto L169
L168:
	;
	v477 = int32(11)
	goto L169
L169:
	;
	v493 = v477 + v433
	goto L105
L170:
	;
	goto L143
L171:
	;
	if int64(-1) < v501 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v500 = int64(*(*int16)(unsafe.Add(mBase, uint32(v494)+8)))
	v501 = v500
	goto L171
L173:
	;
	v499 = int64(*(*int32)(unsafe.Add(mBase, uint32(v494)+8)))
	v501 = v499
	goto L171
L174:
	;
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v494)+8))
	v501 = v498
	goto L171
L175:
	;
	v632 = F_objectGetVal(m, l0)
	mBase = m.M
	v637 = v632 + int32(8)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	v640 = v638 + int32(-1)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	switch v641 + int32(-4) {
	case 0:
		goto L243
	default:
		goto L242
	case 4:
		goto L244
	}
L176:
	;
	v567 = int32(0)
	if base.Ui64(v501) < base.Ui64(int64(10)) {
		v624 = v567
		goto L210
	} else {
		goto L211
	}
L177:
	;
	v506 = int32(0)
	v508 = int64(0) - v501
	if base.Ui64(v508) < base.Ui64(int64(10)) {
		v558 = v506
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v631 = v562 + v563 + int32(1)
	goto L175
L179:
	;
	v562 = v558
	v563 = int32(1)
	goto L178
L180:
	;
	v511 = v508
	v512 = v506
	goto L181
L181:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v511) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v558 = v552
	goto L179
L183:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v511) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v562 = v512
	v563 = int32(2)
	goto L178
L185:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v511) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v562 = v512
	v563 = int32(3)
	goto L178
L187:
	;
	v552 = v512 + int32(12)
	v556 = base.I64_div_u_s(v511, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v511) {
		v511 = v556
		v512 = v552
		goto L181
	} else {
		goto L209
	}
L188:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v511) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v511) {
		goto L201
	} else {
		goto L202
	}
L190:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v511) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v511) {
		goto L198
	} else {
		goto L199
	}
L192:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v511) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v511) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v562 = v512
	v563 = int32(4)
	goto L178
L195:
	;
	v533 = int32(6)
	goto L197
L196:
	;
	v533 = int32(5)
	goto L197
L197:
	;
	v562 = v512
	v563 = v533
	goto L178
L198:
	;
	v538 = int32(8)
	goto L200
L199:
	;
	v538 = int32(7)
	goto L200
L200:
	;
	v562 = v512
	v563 = v538
	goto L178
L201:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v511) {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v511) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v545 = int32(10)
	goto L205
L204:
	;
	v545 = int32(9)
	goto L205
L205:
	;
	v562 = v512
	v563 = v545
	goto L178
L206:
	;
	v550 = int32(12)
	goto L208
L207:
	;
	v550 = int32(11)
	goto L208
L208:
	;
	v562 = v512
	v563 = v550
	goto L178
L209:
	;
	goto L182
L210:
	;
	v631 = int32(1) + v624
	goto L175
L211:
	;
	v570 = v501
	v571 = v567
	goto L212
L212:
	;
	if base.Ui64(int64(100)) <= base.Ui64(v570) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v624 = v618
	goto L210
L214:
	;
	if base.Ui64(int64(1000)) <= base.Ui64(v570) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v631 = int32(2) + v571
	goto L175
L216:
	;
	if base.Ui64(int64(999999999999)) < base.Ui64(v570) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v631 = int32(3) + v571
	goto L175
L218:
	;
	v618 = v571 + int32(12)
	v622 = base.I64_div_u_s(v570, int64(1000000000000))
	if base.Ui64(int64(9999999999999)) < base.Ui64(v570) {
		v570 = v622
		v571 = v618
		goto L212
	} else {
		goto L240
	}
L219:
	;
	if base.Ui64(int64(99999999)) < base.Ui64(v570) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	if base.Ui64(int64(9999999999)) < base.Ui64(v570) {
		goto L232
	} else {
		goto L233
	}
L221:
	;
	if base.Ui64(int64(999999)) < base.Ui64(v570) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if base.Ui64(int64(9999999)) < base.Ui64(v570) {
		goto L229
	} else {
		goto L230
	}
L223:
	;
	if base.Ui64(int64(10000)) <= base.Ui64(v570) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	if base.Ui64(int64(99999)) < base.Ui64(v570) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v631 = int32(4) + v571
	goto L175
L226:
	;
	v595 = int32(6)
	goto L228
L227:
	;
	v595 = int32(5)
	goto L228
L228:
	;
	v631 = v595 + v571
	goto L175
L229:
	;
	v601 = int32(8)
	goto L231
L230:
	;
	v601 = int32(7)
	goto L231
L231:
	;
	v631 = v601 + v571
	goto L175
L232:
	;
	if base.Ui64(int64(99999999999)) < base.Ui64(v570) {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	if base.Ui64(int64(999999999)) < base.Ui64(v570) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v609 = int32(10)
	goto L236
L235:
	;
	v609 = int32(9)
	goto L236
L236:
	;
	v631 = v609 + v571
	goto L175
L237:
	;
	v615 = int32(12)
	goto L239
L238:
	;
	v615 = int32(11)
	goto L239
L239:
	;
	v631 = v615 + v571
	goto L175
L240:
	;
	goto L213
L241:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v656) {
		goto L247
	} else {
		goto L248
	}
L242:
	;
	v655 = int64(*(*int16)(unsafe.Add(mBase, uint32(v637+v640<<(uint(int32(1))%32)))))
	v656 = v655
	goto L241
L243:
	;
	v651 = int64(*(*int32)(unsafe.Add(mBase, uint32(v637+v640<<(uint(int32(2))%32)))))
	v656 = v651
	goto L241
L244:
	;
	v647 = *(*int64)(unsafe.Add(mBase, uint32(v632+v638<<(uint(int32(3))%32))))
	v656 = v647
	goto L241
L245:
	;
	v687 = F_objectGetVal(m, l0)
	mBase = m.M
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	switch v688 + int32(-4) {
	case 0:
		goto L260
	default:
		goto L259
	case 4:
		goto L261
	}
L246:
	;
	v686 = v683*v336 + int32(7)
	goto L245
L247:
	;
	if base.Ui64(int64(8192)) <= base.Ui64(v656+int64(4096)) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v683 = int32(2)
	goto L246
L249:
	;
	if base.Ui64(int64(65536)) <= base.Ui64(v656+int64(32768)) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v683 = int32(3)
	goto L246
L251:
	;
	if base.Ui64(int64(16777216)) <= base.Ui64(v656+int64(8388608)) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v683 = int32(4)
	goto L246
L253:
	;
	if base.Ui64(v656+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v683 = int32(5)
	goto L246
L255:
	;
	v682 = int32(6)
	goto L257
L256:
	;
	v682 = int32(10)
	goto L257
L257:
	;
	v683 = v682
	goto L246
L258:
	;
	if base.Ui64(int64(128)) <= base.Ui64(v694) {
		goto L264
	} else {
		goto L265
	}
L259:
	;
	v693 = int64(*(*int16)(unsafe.Add(mBase, uint32(v687)+8)))
	v694 = v693
	goto L258
L260:
	;
	v692 = int64(*(*int32)(unsafe.Add(mBase, uint32(v687)+8)))
	v694 = v692
	goto L258
L261:
	;
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v687)+8))
	v694 = v691
	goto L258
L262:
	;
	if base.Ui32(v724) < base.Ui32(v686) {
		goto L275
	} else {
		goto L276
	}
L263:
	;
	v724 = v721*v336 + int32(7)
	goto L262
L264:
	;
	if base.Ui64(int64(8192)) <= base.Ui64(v694+int64(4096)) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v721 = int32(2)
	goto L263
L266:
	;
	if base.Ui64(int64(65536)) <= base.Ui64(v694+int64(32768)) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v721 = int32(3)
	goto L263
L268:
	;
	if base.Ui64(int64(16777216)) <= base.Ui64(v694+int64(8388608)) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v721 = int32(4)
	goto L263
L270:
	;
	if base.Ui64(v694+int64(2147483648)) < base.Ui64(int64(4294967296)) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v721 = int32(5)
	goto L263
L272:
	;
	v720 = int32(6)
	goto L274
L273:
	;
	v720 = int32(10)
	goto L274
L274:
	;
	v721 = v720
	goto L263
L275:
	;
	v726 = v686
	goto L277
L276:
	;
	v726 = v724
	goto L277
L277:
	;
	if base.Ui32(v631) < base.Ui32(v493) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v728 = v493
	goto L280
L279:
	;
	v728 = v631
	goto L280
L280:
	;
	v729 = v728
	v732 = v726
	goto L97
L281:
	;
	v769 = F_objectGetVal(m, l0)
	mBase = m.M
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	goto L296
L282:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _consts[545]))
	if base.Ui32(v736) <= base.Ui32(v734) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _consts[546]))
	if base.Ui32(v739) < base.Ui32(v85) {
		goto L281
	} else {
		goto L284
	}
L284:
	;
	if base.Ui32(v739) < base.Ui32(v729) {
		goto L281
	} else {
		goto L285
	}
L285:
	;
	goto L289
L286:
	;
	if base.B2i32(base.Ui32(int32(0)+(v732+v85)) < base.Ui32(int32(1073741825))) == int32(0) {
		goto L281
	} else {
		goto L290
	}
L287:
	;
	goto L286
L289:
	;
	goto L287
L290:
	;
	v753 = F_objectGetVal(m, l0)
	mBase = m.M
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	goto L291
L291:
	;
	v755 = int32(1)
	v758 = F_setTypeConvertAndExpand(m, l0, int32(11), v754+v755, v755)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	v760 = F_objectGetVal(m, l0)
	mBase = m.M
	v761 = F_lpAppend(m, v760, v84, v85)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L6
	} else {
		goto L293
	}
L293:
	;
	v763 = F_lpShrinkToFit(m, v761)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	F_objectSetVal(m, l0, v763)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L6
	} else {
		goto L295
	}
L295:
	;
	goto L20
L296:
	;
	v771 = int32(1)
	v774 = F_setTypeConvertAndExpand(m, l0, int32(2), v770+v771, v771)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L297
	}
L297:
	;
	v776 = F_objectGetVal(m, l0)
	mBase = m.M
	v777 = F_sdsnewlen(m, v84, v85)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L6
	} else {
		goto L298
	}
L298:
	;
	v779 = F_hashtableAdd(m, v776, v777)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L6
	} else {
		goto L299
	}
L299:
	;
	if v779 != 0 {
		goto L20
	} else {
		goto L300
	}
L300:
	;
	F__serverAssert(m, int32(_a2394), int32(_a2391), int32(217))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L301
	}
L301:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setTypeCreate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v3 = int32(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-1)))))
	switch v9 & int32(7) {
	case 0:
		v26 = int32(base.Ui32(v9) >> (uint(int32(3)) % 32))
	case 1:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(-3)))))
		v26 = v16
	case 2:
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(-5)))))
		v26 = v19
	case 3:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-9))))
		v26 = v22
	case 4:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-17))))
		v26 = v25
	default:
		v26 = v3
	}
	v29 = F_string2ll(m, l0, v26, v3)
	mBase = m.M
	if v29 != 0 {
		v30 = int32(0)
	} else {
		v30 = int32(-1)
	}
	if v30 != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, _consts[545]))
		if base.Ui32(v40) < base.Ui32(l1) {
			v45 = F_createSetObject(m)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = F_objectGetVal(m, v45)
				mBase = m.M
				v48 = F_hashtableExpand(m, v47, l1)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					return v45
				}
			}
		} else {
			v42 = F_createSetListpackObject(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				return v42
			}
		}
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _consts[544]))
		if base.Ui32(v32) < base.Ui32(l1) {
			v40 = *(*int32)(unsafe.Add(mBase, _consts[545]))
			if base.Ui32(v40) < base.Ui32(l1) {
				v45 = F_createSetObject(m)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = F_objectGetVal(m, v45)
					mBase = m.M
					v48 = F_hashtableExpand(m, v47, l1)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						return v45
					}
				}
			} else {
				v42 = F_createSetListpackObject(m)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					return v42
				}
			}
		} else {
			v34 = F_createIntsetObject(m)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				return v34
			}
		}
	}
}
func F_setTypeInitIterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = F_valkey_malloc(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(base.Ui32(v10)>>(uint(int32(4))%32)) & int32(15)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v14
		switch v14 + int32(-2) {
		case 0:
			v31 = F_objectGetVal(m, l0)
			mBase = m.M
			v33 = F_hashtableCreateIterator(m, v31, int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v33
				return v5
			}
		default:
			F__serverPanic_1(m, int32(_a2391), int32(331), int32(_a2392), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_abort(m)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		case 4:
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
			return v5
		case 9:
			*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(0)
			return v5
		}
	}
}
func F_setTypeIsMemberAux(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v86 int32
	_ = v86
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int64
	_ = v153
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v175 int64
	_ = v175
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v196 int64
	_ = v196
	var v220 int64
	_ = v220
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l1 != 0 {
		v87 = l1
		v88 = l2
		v89 = l4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v303
L2:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = int32(base.Ui32(v90)>>(uint(int32(4))%32)) & int32(15)
	switch v94 + int32(-6) {
	case 0:
		goto L26
	default:
		goto L25
	case 5:
		goto L27
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12&int32(240) != int32(96) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = v10 + int32(16)
	v42 = int32(0)
	if l3 <= int64(-1) {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v17 = F_objectGetVal(m, l0)
	mBase = m.M
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if base.Ui64(l3+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v303 = v39
	goto L1
L7:
	;
	goto L6
L8:
	;
	v28 = int32(4)
	goto L10
L9:
	;
	v28 = int32(2)
	goto L10
L10:
	;
	if base.Ui64(l3+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = int32(8)
	goto L13
L12:
	;
	v33 = v28
	goto L13
L13:
	;
	if base.Ui32(v20) < base.Ui32(v33) {
		v39 = int32(0)
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(0)
	v36 = F_intsetSearch(m, v17, l3, v35)
	mBase = m.M
	v39 = base.B2i32(v36 != v35)
	goto L7
L15:
	;
	v87 = v41
	v88 = v86
	v89 = v42
	goto L2
L16:
	;
	v86 = v42
	goto L15
L18:
	;
	v67 = F_ull2string(m, v63, v64, v65)
	mBase = m.M
	if v67 == int32(0) {
		goto L16
	} else {
		goto L22
	}
L19:
	;
	goto L21
L20:
	;
	v63 = v41
	v64 = int32(21)
	v65 = l3
	v66 = int32(0)
	goto L18
L21:
	;
	v54 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v54)
	v63 = v10 + int32(17)
	v64 = int32(20)
	v65 = int64(0) - l3
	v66 = int32(1)
	goto L18
L22:
	;
	v86 = v67 + v66
	goto L15
L24:
	;
	v299 = F_lpFind(m, v97, v98, v87, v88, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L28
	} else {
		goto L79
	}
L25:
	;
	if v89 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L26:
	;
	v104 = v10 + int32(8)
	v105 = int32(0)
	if base.Ui32(v88+int32(-21)) < base.Ui32(int32(-20)) {
		v239 = v105
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v97 = F_objectGetVal(m, l0)
	mBase = m.M
	v98 = F_lpFirst(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	if v98 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v303 = int32(0)
	goto L1
L31:
	;
	v247 = F_objectGetVal(m, l0)
	mBase = m.M
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if base.Ui64(v248+int64(-32768)) < base.Ui64(int64(-65536)) {
		goto L62
	} else {
		goto L63
	}
L32:
	;
	if v239 != 0 {
		goto L31
	} else {
		goto L59
	}
L33:
	;
	goto L32
L34:
	;
	v117 = int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != v117 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v239 = int32(1)
	goto L33
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = v220
	goto L35
L37:
	;
	if v118&int32(255) == int32(45) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v122 = v118 + int32(-48)
	if base.Ui32(int32(9)) < base.Ui32(v122&int32(255)) {
		v239 = v105
		goto L33
	} else {
		goto L39
	}
L39:
	;
	if v104 == int32(0) {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v220 = base.I64_extend_i32_u(v122) & int64(255)
	goto L36
L41:
	;
	if base.Ui32(int32(8)) < base.Ui32((v141+int32(-49))&int32(255)) {
		v239 = v105
		goto L33
	} else {
		goto L44
	}
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v140 = int32(2)
	v141 = v138
	v142 = v87 + int32(1)
	goto L41
L43:
	;
	v140 = v117
	v141 = v118
	v142 = v87
	goto L41
L44:
	;
	v153 = base.I64_extend_i32_u(v141+int32(-48)) & int64(255)
	if base.Ui32(v88) <= base.Ui32(v140) {
		v196 = v153
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v118&int32(255) != int32(45) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	v159 = v140
	v161 = v153
	v163 = v142
	goto L47
L47:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if base.Ui32((v165+int32(-58))&int32(255)) < base.Ui32(int32(246)) {
		v239 = v105
		goto L33
	} else {
		goto L49
	}
L48:
	;
	v196 = v186
	goto L45
L49:
	;
	if base.Ui64(int64(1844674407370955161)) < base.Ui64(v161) {
		v239 = v105
		goto L33
	} else {
		goto L50
	}
L50:
	;
	v175 = v161 * int64(10)
	v180 = base.I64_extend_i32_u(v165+int32(-48)) & int64(255)
	if base.Ui64(v180^int64(-1)) < base.Ui64(v175) {
		v239 = v105
		goto L33
	} else {
		goto L51
	}
L51:
	;
	v184 = int32(1)
	v186 = v175 + v180
	v188 = v159 + v184
	if v188 != v88 {
		v159 = v188
		v161 = v186
		v163 = v163 + v184
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	if v196 < int64(0) {
		v239 = v105
		goto L33
	} else {
		goto L57
	}
L54:
	;
	if base.Ui64(int64(-9223372036854775807-1)) < base.Ui64(v196) {
		v239 = v105
		goto L33
	} else {
		goto L55
	}
L55:
	;
	if v104 == int32(0) {
		goto L35
	} else {
		goto L56
	}
L56:
	;
	v220 = int64(0) - v196
	goto L36
L57:
	;
	if v104 == int32(0) {
		goto L35
	} else {
		goto L58
	}
L58:
	;
	v220 = v196
	goto L36
L59:
	;
	v303 = int32(0)
	goto L1
L60:
	;
	v303 = base.B2i32(v270 != int32(0))
	goto L1
L61:
	;
	goto L60
L62:
	;
	v259 = int32(4)
	goto L64
L63:
	;
	v259 = int32(2)
	goto L64
L64:
	;
	if base.Ui64(v248+int64(-2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v264 = int32(8)
	goto L67
L66:
	;
	v264 = v259
	goto L67
L67:
	;
	if base.Ui32(v251) < base.Ui32(v264) {
		v270 = int32(0)
		goto L61
	} else {
		goto L68
	}
L68:
	;
	v266 = int32(0)
	v267 = F_intsetSearch(m, v247, v248, v266)
	mBase = m.M
	v270 = base.B2i32(v267 != v266)
	goto L61
L69:
	;
	if v94 != int32(2) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v94 != int32(2) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v277 = F_objectGetVal(m, l0)
	mBase = m.M
	v279 = F_hashtableFind(m, v277, v87, int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L28
	} else {
		goto L72
	}
L72:
	;
	v303 = v279
	goto L1
L73:
	;
	F__serverPanic_1(m, int32(_a2391), int32(316), int32(_a2392), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L28
	} else {
		goto L78
	}
L74:
	;
	v283 = F_sdsnewlen(m, v87, v88)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L28
	} else {
		goto L75
	}
L75:
	;
	v285 = F_objectGetVal(m, l0)
	mBase = m.M
	v287 = F_hashtableFind(m, v285, v283, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L28
	} else {
		goto L76
	}
L76:
	;
	F_sdsfree(m, v283)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L28
	} else {
		goto L77
	}
L77:
	;
	v303 = v287
	goto L1
L78:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v303 = base.B2i32(v299 != int32(0))
	goto L1
}
func F_setTypeNext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v12 + int32(-2) {
	case 0:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = F_hashtableNext(m, v15, v10+int32(12))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				v116 = int32(-1)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-1)))))
				switch v29 & int32(7) {
				case 0:
					v46 = int32(base.Ui32(v29) >> (uint(int32(3)) % 32))
				case 1:
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(-3)))))
					v46 = v36
				case 2:
					v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(-5)))))
					v46 = v39
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-9))))
					v46 = v42
				case 4:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(-17))))
					v46 = v45
				default:
					v46 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v46
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(-123456789)
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v116 = v115
			}
			m.G0 = v10 + int32(16)
			return v116
		}
	default:
		F__serverPanic_1(m, int32(_a2391), int32(386), int32(_a2393), int32(0))
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v52 = F_objectGetVal(m, v51)
		mBase = m.M
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v53 + int32(1)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
		if base.Ui32(v60) <= base.Ui32(v53) {
			v82 = int32(0)
		} else {
			v63 = v52 + int32(8)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
			switch v64 + int32(-4) {
			case 0:
				v74 = int64(*(*int32)(unsafe.Add(mBase, uint32(v63+v53<<(uint(int32(2))%32)))))
				v79 = v74
			default:
				v78 = int64(*(*int16)(unsafe.Add(mBase, uint32(v63+v53<<(uint(int32(1))%32)))))
				v79 = v78
			case 4:
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v63+v53<<(uint(int32(3))%32))))
				v79 = v70
			}
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v79
			v82 = int32(1)
		}
		if v82 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v116 = v115
		} else {
			v116 = int32(-1)
		}
		m.G0 = v10 + int32(16)
		return v116
	case 9:
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v86 = F_objectGetVal(m, v85)
		mBase = m.M
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v87 != 0 {
			v90 = F_lpNext(m, v86, v87)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = v90
				if v92 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v92
					v97 = F_lpGetValue(m, v92, v10+int32(8), l3)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v97
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v100
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v116 = v115
						m.G0 = v10 + int32(16)
						return v116
					}
				} else {
					v116 = int32(-1)
					m.G0 = v10 + int32(16)
					return v116
				}
			}
		} else {
			v88 = F_lpFirst(m, v86)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v92 = v88
				if v92 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v92
					v97 = F_lpGetValue(m, v92, v10+int32(8), l3)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v97
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v100
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v116 = v115
						m.G0 = v10 + int32(16)
						return v116
					}
				} else {
					v116 = int32(-1)
					m.G0 = v10 + int32(16)
					return v116
				}
			}
		}
	}
}
func F_setTypeSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch int32(base.Ui32(v2)>>(uint(int32(4))%32))&int32(15) + int32(-2) {
	case 0:
		v25 = F_objectGetVal(m, l0)
		mBase = m.M
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
		return v26 + v27
	default:
		F__serverPanic_1(m, int32(_a2391), int32(481), int32(_a2392), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v9 = F_objectGetVal(m, l0)
		mBase = m.M
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		return v10
	case 9:
		v12 = F_objectGetVal(m, l0)
		mBase = m.M
		v13 = F_lpLength(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
