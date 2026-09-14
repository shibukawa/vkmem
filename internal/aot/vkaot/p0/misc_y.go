package p0

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F___year_to_secs(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	if base.Ui64(int64(136)) < base.Ui64(l0+int64(-2)) {
		v43 = l0 + int64(-100)
		v44 = int64(400)
		v45 = base.I64_div_s(v43, v44)
		v48 = v43 - v45*v44
		v54 = base.I32_wrap_i64(v48)
		if v48 < int64(0) {
			v59 = v54 + int32(400)
		} else {
			v59 = v54
		}
		if v59 != 0 {
			if v59 < int32(200) {
				v75 = base.B2i32(int32(99) < v59)
				if int32(99) < v59 {
					v76 = v59 + int32(-100)
				} else {
					v76 = v59
				}
				v77 = v76
				v78 = v75
			} else {
				if base.Ui32(v59) < base.Ui32(int32(300)) {
					v77 = v59 + int32(-200)
					v78 = int32(2)
				} else {
					v77 = v59 + int32(-300)
					v78 = int32(3)
				}
			}
			if v77 != 0 {
				v84 = int32(base.Ui32(v77) >> (uint(int32(2)) % 32))
				v87 = int32(0)
				v88 = base.B2i32(v77&int32(3) == v87)
				if l1 == v87 {
					v95 = v88
					v96 = v78
					v97 = v84
				} else {
					v91 = v88
					v92 = v78
					v93 = v84
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
					v95 = v91
					v96 = v92
					v97 = v93
				}
			} else {
				v80 = int32(0)
				v81 = v78
				v82 = int32(0)
				if l1 != 0 {
					v91 = v80
					v92 = v81
					v93 = v82
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
					v95 = v91
					v96 = v92
					v97 = v93
				} else {
					v95 = v80
					v96 = v81
					v97 = v82
				}
			}
		} else {
			v80 = int32(1)
			v81 = int32(0)
			v82 = int32(0)
			if l1 != 0 {
				v91 = v80
				v92 = v81
				v93 = v82
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
				v95 = v91
				v96 = v92
				v97 = v93
			} else {
				v95 = v80
				v96 = v81
				v97 = v82
			}
		}
		return v43*int64(31536000) + base.I64_extend_i32_s(v97+(v96*int32(24)+(base.I32_wrap_i64(v48>>(uint(int64(63))%64))+base.I32_wrap_i64(v45))*int32(97))-v95)*int64(86400) + int64(946771200)
	} else {
		v13 = base.I32_wrap_i64(l0)
		v17 = (v13 + int32(-68)) >> (uint(int32(2)) % 32)
		if v13&int32(3) != 0 {
			if l1 == int32(0) {
				v31 = v17
			} else {
				v28 = v17
				v29 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
				v31 = v28
			}
		} else {
			v21 = v17 + int32(-1)
			if l1 == int32(0) {
				v31 = v21
			} else {
				v28 = v21
				v29 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
				v31 = v28
			}
		}
		return base.I64_extend_i32_s(v13*int32(31536000) + v31*int32(86400) + int32(2087447296))
	}
}
