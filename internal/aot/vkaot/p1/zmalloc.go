package p1

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_zmalloc_get_private_dirty(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_zmalloc_get_rss(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v1 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v10 < int32(261) {
		if v10 < int32(1) {
			v87 = v1
		} else {
			v18 = v1
			v19 = v10
			v21 = v19 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v19) {
				v28 = int32(0)
				v30 = v18
				v31 = v28
				v35 = v28
				for {
					v38 = v31 << (uint(int32(2)) % 32)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[317])))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[318])))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[319])))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[320])))
					v54 = v41 + (v44 + (v47 + (v50 + v30)))
					v55 = int32(4)
					v56 = v31 + v55
					v58 = v35 + v55
					if v58 != v19&int32(2147483644) {
						v30 = v54
						v31 = v56
						v35 = v58
						continue
					} else {
						break
					}
					break
				}
				v60 = v54
				v61 = v56
			} else {
				v60 = v18
				v61 = int32(0)
			}
			if v21 == int32(0) {
				v87 = v60
			} else {
				v69 = v60
				v70 = v61
				v72 = int32(0)
				for {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[320])))
					v81 = v80 + v69
					v82 = int32(1)
					v85 = v72 + v82
					if v85 != v21 {
						v69 = v81
						v70 = v70 + v82
						v72 = v85
						continue
					} else {
						break
					}
					break
				}
				v87 = v81
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[316]))
		v18 = v14
		v19 = int32(260)
		v21 = v19 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v19) {
			v28 = int32(0)
			v30 = v18
			v31 = v28
			v35 = v28
			for {
				v38 = v31 << (uint(int32(2)) % 32)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[317])))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[318])))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[319])))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[320])))
				v54 = v41 + (v44 + (v47 + (v50 + v30)))
				v55 = int32(4)
				v56 = v31 + v55
				v58 = v35 + v55
				if v58 != v19&int32(2147483644) {
					v30 = v54
					v31 = v56
					v35 = v58
					continue
				} else {
					break
				}
				break
			}
			v60 = v54
			v61 = v56
		} else {
			v60 = v18
			v61 = int32(0)
		}
		if v21 == int32(0) {
			v87 = v60
		} else {
			v69 = v60
			v70 = v61
			v72 = int32(0)
			for {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[320])))
				v81 = v80 + v69
				v82 = int32(1)
				v85 = v72 + v82
				if v85 != v21 {
					v69 = v81
					v70 = v70 + v82
					v72 = v85
					continue
				} else {
					break
				}
				break
			}
			v87 = v81
		}
	}
	return v87
}
func F_zmalloc_usable_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(-8))))
	return v4 & int32(2147483647)
}
func F_zmalloc_used_memory(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v1 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	if v10 < int32(261) {
		if v10 < int32(1) {
			v87 = v1
		} else {
			v18 = v1
			v19 = v10
			v21 = v19 & int32(3)
			if base.Ui32(int32(4)) <= base.Ui32(v19) {
				v28 = int32(0)
				v30 = v18
				v31 = v28
				v35 = v28
				for {
					v38 = v31 << (uint(int32(2)) % 32)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[317])))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[318])))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[319])))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[320])))
					v54 = v41 + (v44 + (v47 + (v50 + v30)))
					v55 = int32(4)
					v56 = v31 + v55
					v58 = v35 + v55
					if v58 != v19&int32(2147483644) {
						v30 = v54
						v31 = v56
						v35 = v58
						continue
					} else {
						break
					}
					break
				}
				v60 = v54
				v61 = v56
			} else {
				v60 = v18
				v61 = int32(0)
			}
			if v21 == int32(0) {
				v87 = v60
			} else {
				v69 = v60
				v70 = v61
				v72 = int32(0)
				for {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[320])))
					v81 = v80 + v69
					v82 = int32(1)
					v85 = v72 + v82
					if v85 != v21 {
						v69 = v81
						v70 = v70 + v82
						v72 = v85
						continue
					} else {
						break
					}
					break
				}
				v87 = v81
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[316]))
		v18 = v14
		v19 = int32(260)
		v21 = v19 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(v19) {
			v28 = int32(0)
			v30 = v18
			v31 = v28
			v35 = v28
			for {
				v38 = v31 << (uint(int32(2)) % 32)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[317])))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[318])))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[319])))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[320])))
				v54 = v41 + (v44 + (v47 + (v50 + v30)))
				v55 = int32(4)
				v56 = v31 + v55
				v58 = v35 + v55
				if v58 != v19&int32(2147483644) {
					v30 = v54
					v31 = v56
					v35 = v58
					continue
				} else {
					break
				}
				break
			}
			v60 = v54
			v61 = v56
		} else {
			v60 = v18
			v61 = int32(0)
		}
		if v21 == int32(0) {
			v87 = v60
		} else {
			v69 = v60
			v70 = v61
			v72 = int32(0)
			for {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_consts[320])))
				v81 = v80 + v69
				v82 = int32(1)
				v85 = v72 + v82
				if v85 != v21 {
					v69 = v81
					v70 = v70 + v82
					v72 = v85
					continue
				} else {
					break
				}
				break
			}
			v87 = v81
		}
	}
	return v87
}
