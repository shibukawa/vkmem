package p2

import (
	base "github.com/shibukawa/vkmem/internal/aot/vkaot/base"
	"unsafe"
)

func F_receiveChildInfo(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 float64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[0]))
	if v7 == int32(-1) {
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1]))
		if v11 != int32(32) {
			v18 = v11
		} else {
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1])) = v14
			v18 = v14
		}
		v19 = int32(_a_F_receiveChildInfo_0)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[0]))
		v25 = F_read(m, v20, v18+int32(_a_F_receiveChildInfo_1), int32(32)-v18)
		mBase = m.M
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1]))
		if v25 < int32(1) {
			v33 = v27
		} else {
			v31 = v27 + v25
			*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1])) = v31
			v33 = v31
		}
		if v33 != int32(32) {
		} else {
			v36 = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[2]))
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[3]))
			v41 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[4]))
			if base.Ui32(v39) <= base.Ui32(v41) {
				v45 = v41
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[4])) = v39
				v45 = v39
			}
			switch v37 {
			case 0:
				v61 = int32(0)
				v62 = *(*float64)(unsafe.Add(mBase, _c_F_receiveChildInfo[5]))
				v64 = *(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[6]))
				v65 = int32(_a_F_receiveChildInfo_0)
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[7]))
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[8])) = v67
				*(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[9])) = v64
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[10])) = v39
				if base.F64_eq(v62, float64(-1)) != 0 {
				} else {
					*(*float64)(unsafe.Add(mBase, _c_F_receiveChildInfo[11])) = v62
				}
			case 1:
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[12])) = v45
			case 2:
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[13])) = v45
			case 3:
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[14])) = v45
			case 4:
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[15])) = v45
			case 5:
				v46 = int32(_a_F_receiveChildInfo_0)
				v48 = *(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[16]))
				v50 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_receiveChildInfo[17])))
				*(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[16])) = v48 + v50
			default:
			}
			for {
				v84 = int32(_a_F_receiveChildInfo_0)
				*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1])) = int32(0)
				v88 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[0]))
				v91 = F_read(m, v88, int32(_a_F_receiveChildInfo_1), int32(32))
				mBase = m.M
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1]))
				if v91 < int32(1) {
					v99 = v93
				} else {
					v97 = v93 + v91
					*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[1])) = v97
					v99 = v97
				}
				if v99 != int32(32) {
					break
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[3]))
					v105 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[4]))
					if base.Ui32(v103) <= base.Ui32(v105) {
						v109 = v105
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[4])) = v103
						v109 = v103
					}
					v111 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[2]))
					switch v111 {
					case 0:
						v112 = int32(0)
						v113 = *(*float64)(unsafe.Add(mBase, _c_F_receiveChildInfo[5]))
						v115 = *(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[6]))
						v116 = int32(_a_F_receiveChildInfo_0)
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[7]))
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[8])) = v118
						*(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[9])) = v115
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[10])) = v103
						if base.F64_eq(v113, float64(-1)) != 0 {
							continue
						} else {
							*(*float64)(unsafe.Add(mBase, _c_F_receiveChildInfo[11])) = v113
							continue
						}
						continue
					case 1:
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[12])) = v109
						continue
					case 2:
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[13])) = v109
						continue
					case 3:
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[14])) = v109
						continue
					case 4:
						*(*int32)(unsafe.Add(mBase, _c_F_receiveChildInfo[15])) = v109
						continue
					case 5:
						v136 = int32(_a_F_receiveChildInfo_0)
						v138 = *(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[16]))
						v140 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_receiveChildInfo[17])))
						*(*int64)(unsafe.Add(mBase, _c_F_receiveChildInfo[16])) = v138 + v140
						continue
					default:
						continue
					}
					continue
				}
				break
			}
		}
	}
	return
}
func F_resetChildState(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int64
	_ = v3
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	v1 = int32(0)
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[0])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[1])) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[2])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[3])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[4])) = v3
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[5]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[6]))
	if v25 != 0 {
		v26 = int32(2)
	} else {
		v26 = base.B2i32(v19 == v1) << (uint(int32(1)) % 32)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_resetChildState[7])) = v26
	*(*int32)(unsafe.Add(mBase, _c_F_resetChildState[8])) = v26
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[9]))
	if v33 != int32(-1) {
		v40 = F_close(m, v33)
		mBase = m.M
		v41 = int32(_a_F_resetChildState_0)
		v42 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[10]))
		v43 = F_close(m, v42)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_resetChildState[11])) = int32(0)
		*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[9])) = int64(-1)
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[10]))
		if v37 == int32(-1) {
		} else {
			v40 = F_close(m, v33)
			mBase = m.M
			v41 = int32(_a_F_resetChildState_0)
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_resetChildState[10]))
			v43 = F_close(m, v42)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _c_F_resetChildState[11])) = int32(0)
			*(*int64)(unsafe.Add(mBase, _c_F_resetChildState[9])) = int64(-1)
		}
	}
	F_moduleFireServerEvent(m, int64(13), int32(1), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		return
	} else {
		return
	}
}
func F_sendChildCowInfo(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = int32(0)
	F_sendChildInfoGeneric(m, l0, v3, v3, float64(-1), l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
