package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_luaH_setnum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 float64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l2 < int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v88
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v10))) = base.F64_convert_i32_s(l2)
	v79 = F_newkey(m, l0, l1, v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	v66 = m.G398
	if v64 != v66 {
		v88 = v64
		goto L1
	} else {
		goto L17
	}
L4:
	;
	v51 = v44
	goto L11
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v27 = base.I64_reinterpret_f64(v25)
	v32 = int32(-1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v39 = base.I32_rem_u_s(base.I32_wrap_i64(int64(base.Ui64(v27)>>(uint(int64(32))%64))+v27), v32<<(uint(v33)%32)^v32|int32(1))
	v43 = v25
	v44 = v26 + v39<<(uint(int32(5))%32)
	goto L4
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v64 = v19 + l2<<(uint(int32(4))%32) + int32(-16)
	goto L3
L7:
	;
	v17 = base.F64_convert_i32_s(l2)
	if l2 != 0 {
		v25 = v17
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if l2 <= v14 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v25 = base.F64_convert_i32_u(l2)
	goto L5
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v43 = v17
	v44 = v18
	goto L4
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	if v53 != int32(3) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	if v58 != 0 {
		v51 = v58
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v51)+16))
	if base.F64_eq(v56, v43) != 0 {
		v64 = v51
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L2
L17:
	;
	goto L2
L18:
	;
	return int32(0)
L19:
	;
	v88 = v79
	goto L1
}
