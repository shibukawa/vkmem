package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_vsetInitIterator(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+376)) = v3
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+368)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+356)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l1)+352)) = v3
	return
}
func F_vsetIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != 0 {
		return base.B2i32(v2 == int32(-1))
	} else {
		F__serverAssert(m, int32(_a1878), int32(_a1861), int32(2363))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_abort(m)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_vsetIsValid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		v15 = int32(0)
	} else {
		v5 = int32(1)
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v6 + v5 {
		case 0:
			v15 = v5
		case 1:
			v15 = int32(0)
		default:
			if v6&int32(7) != 0 {
				v15 = v5
			} else {
				v15 = int32(0)
			}
		}
	}
	return v15
}
func F_vsetMemUsage(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(304)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 + int32(1) {
	case 0:
		v145 = v2
		goto L1
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(304)
	return v145
L2:
	;
	if v9&int32(1) != 0 {
		v145 = v2
		goto L1
	} else {
		goto L6
	}
L3:
	;
	F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L6:
	;
	switch v9 & int32(6) {
	default:
		goto L10
	case 1, 3, 5:
		goto L11
	case 2:
		goto L14
	case 4:
		goto L13
	case 6:
		goto L12
	}
L7:
	;
	F__serverAssert(m, int32(_a1874), int32(_a1861), int32(1696))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L43
	}
L8:
	;
	F__serverAssert(m, int32(_a1868), int32(_a1861), int32(1710))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L42
	}
L9:
	;
	F__serverAssert(m, int32(_a1874), int32(_a1861), int32(1696))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L41
	}
L10:
	;
	F__serverPanic_1(m, int32(_a1861), int32(2264), int32(_a1875), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L40
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v40 = v9 & int32(-8)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	goto L17
L13:
	;
	F__serverPanic_1(m, int32(_a1861), int32(2260), int32(_a1876), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v25 = v9 & int32(-8)
	if v25 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v145 = base.I32_wrap_i64(int64(base.Ui64(v28) >> (uint(int64(30)) % 64)))
	goto L1
L16:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(128)
	v47 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v6)+296)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v6)+160)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v6 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+156)) = v6 + int32(168)
	goto L18
L18:
	;
	v60 = int32(0)
	v62 = F_raxSeek(m, v6, int32(_a263), v60, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v62 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v66 = F_raxNext(m, v6)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	F_raxStop(m, v6)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L39
	}
L22:
	;
	if v66 == int32(0) {
		v115 = v41
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = v41
	goto L24
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	switch v73 + int32(1) {
	case 0:
		v110 = v72
		goto L26
	case 1:
		goto L28
	default:
		goto L27
	}
L25:
	;
	v115 = v110
	goto L21
L26:
	;
	v111 = F_raxNext(m, v6)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L37
	}
L27:
	;
	if v73&int32(1) != 0 {
		v110 = v72
		goto L26
	} else {
		goto L30
	}
L28:
	;
	F__serverAssert(m, int32(_a1864), int32(_a1861), int32(781))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	switch v73&int32(6) + int32(-2) {
	case 0:
		goto L33
	default:
		goto L32
	case 2:
		goto L31
	}
L31:
	;
	v106 = F_hashtableMemUsage(m, v73&int32(-8))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L36
	}
L32:
	;
	F__serverPanic_1(m, int32(_a1861), int32(1726), int32(_a1877), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	v89 = v73 & int32(-8)
	if v89 == int32(0) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v110 = v72 + base.I32_wrap_i64(int64(base.Ui64(v92)>>(uint(int64(30))%64)))
	goto L26
L35:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v110 = v106 + v72
	goto L26
L37:
	;
	if v111 != 0 {
		v72 = v110
		goto L24
	} else {
		goto L38
	}
L38:
	;
	goto L25
L39:
	;
	v145 = v115
	goto L1
L40:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vsetRelease(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 == int32(-1) {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		return
	} else {
		F_freeVsetBucket(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return
		}
	}
}
