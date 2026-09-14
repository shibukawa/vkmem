package p0

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_pubsubMemOverhead(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		v7 = F_hashtableMemUsage(m, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v13 = F_hashtableMemUsage(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				v18 = F_hashtableMemUsage(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v7 + v13 + v18
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_pubsubUnsubscribeAllPatterns(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	v6 = m.G0
	v8 = v6 - int32(64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v10 != 0 {
		v36 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	goto L10
L2:
	;
	v12 = F_valkey_malloc(m, int32(32))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v12
	v18 = F_hashtableCreate(m, int32(_a862))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v18
	v23 = F_hashtableCreate(m, int32(_a862))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v23
	v28 = F_hashtableCreate(m, int32(_a862))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v28
	v36 = v30
	goto L1
L8:
	;
	if l1 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v44 = v8 + int32(16)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = int32(1)
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+14)) = uint8(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v48
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+15)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = int32(-1)
	if v46 == v48 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v39+v40 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v101 = int32(0)
	goto L8
L12:
	;
	v65 = int32(0)
	v70 = F_hashtableNext(m, v8+int32(16), v8+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L17
	}
L13:
	;
	goto L12
L14:
	;
	goto L15
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v44
	goto L13
L16:
	;
	F_hashtableCleanupIterator(m, v8+int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L24
	}
L17:
	;
	if v70 == int32(0) {
		v92 = v65
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v77 = v65
	goto L19
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v80 = F_pubsubUnsubscribePattern(m, l0, v79, l1)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L21
	}
L20:
	;
	v92 = v82
	goto L16
L21:
	;
	v82 = v80 + v77
	v87 = F_hashtableNext(m, v8+int32(16), v8+int32(12))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v87 != 0 {
		v77 = v82
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v101 = v92
	goto L8
L25:
	;
	m.G0 = v8 + int32(64)
	return v101
L26:
	;
	if v101 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_addReplyPubsubPatUnsubscribed(m, l0, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L25
}
