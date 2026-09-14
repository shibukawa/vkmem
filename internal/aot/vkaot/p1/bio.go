package p1

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_allocBioJob(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_valkey_malloc(m, l0+int32(24))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_bioCreateCloseAofJob(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = F_allocBioJob(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(3)
		v12 = int32(1)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
		v21 = l2<<(uint(v12)%32)&int32(2) | v16&int32(252) | v12
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v21)
		v23 = int32(0)
		v25 = *(*int32)(unsafe.Add(mBase, _consts[70]))
		*(*int32)(unsafe.Add(mBase, _consts[70])) = v25 + v12
		F_bioExecuteJob(m, v6)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bioCreateCloseJob(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = F_allocBioJob(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		v9 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
		v11 = int32(1)
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
		v21 = l2<<(uint(v11)%32)&int32(2) | l1&v11 | v18&int32(252)
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v21)
		v25 = *(*int32)(unsafe.Add(mBase, _consts[69]))
		*(*int32)(unsafe.Add(mBase, _consts[69])) = v25 + v11
		F_bioExecuteJob(m, v6)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			return
		}
	}
}
func F_bioCreateLazyFreeJob(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v18 = F_allocBioJob(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l2
		if l1 < int32(1) {
		} else {
			v25 = l1 & int32(3)
			v27 = v18 + int32(8)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			if base.Ui32(int32(4)) <= base.Ui32(l1) {
				v35 = int32(0)
				v38 = v35
				v39 = v29
				v46 = v35
				for {
					v48 = int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v39 + v48
					v53 = v27 + v38<<(uint(int32(2))%32)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v53))) = v54
					v56 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v39 + v56
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v53+v48))) = v61
					v63 = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v39 + v63
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v53+v56))) = v68
					v71 = v39 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v71
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v53+v63))) = v75
					v78 = v38 + v48
					v80 = v46 + v48
					if v80 != l1&int32(2147483644) {
						v38 = v78
						v39 = v71
						v46 = v80
						continue
					} else {
						break
					}
					break
				}
				v82 = v71
				v83 = v78
			} else {
				v82 = v29
				v83 = int32(0)
			}
			if v25 == int32(0) {
			} else {
				v95 = v82
				v96 = v83
				v102 = int32(0)
				for {
					v107 = v95 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v107
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v96<<(uint(int32(2))%32)))) = v112
					v114 = int32(1)
					v117 = v102 + v114
					if v117 != v25 {
						v95 = v107
						v96 = v96 + v114
						v102 = v117
						continue
					} else {
						break
					}
					break
				}
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(2)
		v132 = int32(0)
		v134 = *(*int32)(unsafe.Add(mBase, _consts[68]))
		*(*int32)(unsafe.Add(mBase, _consts[68])) = v134 + int32(1)
		F_bioExecuteJob(m, v18)
		mBase = m.M
		v139 = m.ExcPending
		if v139 != 0 {
			return
		} else {
			m.G0 = v14 + int32(16)
			return
		}
	}
}
func F_bioPendingJobsOfType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[69])))
	return v6
}
