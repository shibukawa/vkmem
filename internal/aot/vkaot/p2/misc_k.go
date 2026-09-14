package p2

import (
	base "github.com/shibukawa/valkeymem/internal/aot/vkaot/base"
	"unsafe"
)

func F_keyspaceEventsStringToFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = l0
	v6 = int32(0)
	goto L3
L1:
	;
	return v29
L2:
	;
	v29 = int32(-1)
	goto L1
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	switch v9 {
	case 0:
		v29 = v6
		goto L1
	default:
		goto L2
	case 36:
		goto L18
	case 65:
		v24 = int32(10236)
		goto L5
	case 69:
		goto L10
	case 75:
		goto L11
	case 100:
		goto L7
	case 101:
		goto L12
	case 103:
		goto L19
	case 104:
		goto L15
	case 108:
		goto L17
	case 109:
		goto L8
	case 110:
		goto L6
	case 115:
		goto L16
	case 116:
		goto L9
	case 120:
		goto L13
	case 122:
		goto L14
	}
L5:
	;
	v5 = v5 + int32(1)
	v6 = v6 | v24
	goto L3
L6:
	;
	v24 = int32(16384)
	goto L5
L7:
	;
	v24 = int32(8192)
	goto L5
L8:
	;
	v24 = int32(2048)
	goto L5
L9:
	;
	v24 = int32(1024)
	goto L5
L10:
	;
	v24 = int32(2)
	goto L5
L11:
	;
	v24 = int32(1)
	goto L5
L12:
	;
	v24 = int32(512)
	goto L5
L13:
	;
	v24 = int32(256)
	goto L5
L14:
	;
	v24 = int32(128)
	goto L5
L15:
	;
	v24 = int32(64)
	goto L5
L16:
	;
	v24 = int32(32)
	goto L5
L17:
	;
	v24 = int32(16)
	goto L5
L18:
	;
	v24 = int32(8)
	goto L5
L19:
	;
	v24 = int32(4)
	goto L5
}
