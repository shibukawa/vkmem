package host

import (
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"hash/crc64"
)

// Hashes on the host: Valkey's sha1.c/sha256.c are replaced (wasm/*.inc)
// by contexts that hold a handle to a Go hash.Hash, and crc64() calls
// straight into hash/crc64. Same bytes out, hardware-accelerated in.

// Valkey's CRC-64 (crc64.c, "Jones" polynomial 0xad93d23594c935a9 fed
// LSB-first and reflected on output, init 0, no final xor). That is the
// reflected CRC with the reversed polynomial, which is what hash/crc64
// computes behind an init/xorout of all ones; the raw state maps through
// ^Update(^crc).
var crc64Jones = crc64.MakeTable(0x95ac9329ac4bc9b5)

func crc64Raw(crc uint64, b []byte) uint64 {
	return ^crc64.Update(^crc, crc64Jones, b)
}

func (h *Host) newHashHandle(typ int32) int32 {
	var hh hash.Hash
	if free := h.hashFree[typ]; len(free) > 0 {
		hh = free[len(free)-1]
		h.hashFree[typ] = free[:len(free)-1]
		hh.Reset()
	} else {
		switch typ {
		case 1:
			hh = sha1.New()
		case 2:
			hh = sha256.New()
		default:
			return 0
		}
	}
	if h.hashes == nil {
		h.hashes = map[int32]hash.Hash{}
		h.hashType = map[int32]int32{}
	}
	for {
		h.nextHash++
		if h.nextHash <= 0 {
			h.nextHash = 1
		}
		if _, taken := h.hashes[h.nextHash]; !taken {
			break
		}
	}
	h.hashes[h.nextHash] = hh
	h.hashType[h.nextHash] = typ
	return h.nextHash
}

var cryptoTable = []Fn{
	{"env", "vkmem_hash_create", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return ret32(h.newHashHandle(i32(a[0])))
	}},
	{"env", "vkmem_hash_update", "iii", "", func(h *Host, m Memory, a []uint64) uint64 {
		if hh := h.hashes[i32(a[0])]; hh != nil {
			if b, ok := m.Read(u32(a[1]), u32(a[2])); ok {
				hh.Write(b)
			}
		}
		return 0
	}},
	// vkmem_hash_final writes the digest and releases the handle.
	{"env", "vkmem_hash_final", "iii", "", func(h *Host, m Memory, a []uint64) uint64 {
		id := i32(a[0])
		hh := h.hashes[id]
		if hh == nil {
			return 0
		}
		if uint32(hh.Size()) <= u32(a[2]) {
			m.Write(u32(a[1]), hh.Sum(nil))
		}
		typ := h.hashType[id]
		delete(h.hashes, id)
		delete(h.hashType, id)
		if h.hashFree == nil {
			h.hashFree = map[int32][]hash.Hash{}
		}
		if len(h.hashFree[typ]) < 8 {
			h.hashFree[typ] = append(h.hashFree[typ], hh)
		}
		return 0
	}},
	{"env", "vkmem_crc64", "jij", "j", func(h *Host, m Memory, a []uint64) uint64 {
		b, ok := m.Read(u32(a[1]), uint32(a[2]))
		if !ok {
			return a[0]
		}
		return crc64Raw(a[0], b)
	}},
}

func init() {
	table = append(table, cryptoTable...)
}
