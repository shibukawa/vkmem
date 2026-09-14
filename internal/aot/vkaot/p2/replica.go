package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_findReplica(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	v8 = m.G0
	v10 = v8 - int32(64)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v15 = v10 + int32(56)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16
	goto L1
L1:
	;
	v21 = v10 + int32(56)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	m.G0 = v10 + int32(64)
	return v135
L3:
	;
	v135 = int32(0)
	goto L2
L4:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23+base.B2i32(v26 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
	goto L5
L7:
	;
	v39 = v23
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+152))
	if v45 != 0 {
		v62 = v45
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L3
L10:
	;
	v110 = v10 + int32(56)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v112 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v46 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	if v50 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v56 = m.T0[v50].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v46, v10, int32(46), int32(0), int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v56 == int32(-1) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v62 = v10
	goto L11
L18:
	;
	if v99-v101 != 0 {
		goto L10
	} else {
		goto L30
	}
L19:
	;
	v99 = F_tolower(m, v95)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v101 = F_tolower(m, v100)
	mBase = m.M
	goto L18
L20:
	;
	v69 = l0
	v70 = v62
	v71 = v67
	goto L23
L21:
	;
	v95 = int32(0)
	v96 = v62
	goto L19
L22:
	;
	v95 = v92 & int32(255)
	v96 = v91
	goto L19
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 == int32(0) {
		v91 = v70
		v92 = v71
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v91 = v85
	v92 = int32(0)
	goto L22
L25:
	;
	v77 = v71 & int32(255)
	if v77 == v73 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = int32(1)
	v85 = v70 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v86 != 0 {
		v69 = v69 + v84
		v70 = v85
		v71 = v86
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v79 = F_tolower(m, v77)
	mBase = m.M
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v81 = F_tolower(m, v80)
	mBase = m.M
	if v79 == v81 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v91 = v70
	v92 = v83
	goto L22
L29:
	;
	goto L24
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+148))
	if l1 == v104 {
		v135 = v43
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	if v112 != 0 {
		v39 = v112
		goto L8
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v112+base.B2i32(v115 == int32(0))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v121
	goto L33
L35:
	;
	goto L9
}
func F_flushReplicaKeysWithExpireList(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[422]))
	if v4 == v2 {
		return
	} else {
		if l0 == int32(0) {
			F_dictRelease(m, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[422])) = v13
				return
			}
		} else {
			F_freeReplicaKeysWithExpireAsync(m, v4)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v13 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[422])) = v13
				return
			}
		}
	}
}
