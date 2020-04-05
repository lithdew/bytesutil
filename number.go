package bytesutil

import (
	"encoding/binary"
	"math"
)

func AppendBitcoinUvarInt(dst []byte, v uint64) []byte {
	if v < 0xfd {
		return append(dst, byte(v))
	}

	if v <= math.MaxUint16 {
		dst = append(dst, 0xfd)
		dst = AppendUint16LE(dst, uint16(v))
		return dst
	}

	if v <= math.MaxUint32 {
		dst = append(dst, 0xfe)
		dst = AppendUint32LE(dst, uint32(v))
		return dst
	}

	dst = append(dst, 0xff)
	return AppendUint64LE(dst, v)
}

func BitcoinUvarInt(b []byte) (uint64, int) {
	if len(b) == 0 {
		return 0, 0
	}

	switch b[0] {
	case 0xff:
		if len(b) < 9 {
			return 0, -len(b)
		}
		return Uint64LE(b[1 : 1+8]), 9
	case 0xfe:
		if len(b) < 5 {
			return 0, -len(b)
		}
		return uint64(Uint32LE(b[1 : 1+4])), 5
	case 0xfd:
		if len(b) < 3 {
			return 0, -len(b)
		}
		return uint64(Uint16LE(b[1 : 1+2])), 3
	default:
		return uint64(b[0]), 1
	}
}

func AppendUvarInt(dst []byte, n uint64) []byte {
	for n >= 0x80 {
		dst = append(dst, byte(n)|0x80)
		n >>= 7
	}
	return append(dst, byte(n))
}

func AppendVarInt(dst []byte, n int64) []byte {
	un := uint64(n) << 1
	if n < 0 {
		un = ^un
	}
	return AppendUvarInt(dst, un)
}

func UvarInt(b []byte) (uint64, int) {
	return binary.Uvarint(b)
}

func VarInt(b []byte) (int64, int) {
	return binary.Varint(b)
}

func AppendUint64LE(dst []byte, n uint64) []byte {
	return append(dst, byte(n), byte(n>>8), byte(n>>16), byte(n>>24), byte(n>>32), byte(n>>40), byte(n>>48), byte(n>>56))
}

func AppendUint32LE(dst []byte, n uint32) []byte {
	return append(dst, byte(n), byte(n>>8), byte(n>>16), byte(n>>24))
}

func AppendUint16LE(dst []byte, n uint16) []byte {
	return append(dst, byte(n), byte(n>>8))
}

func AppendUint64BE(dst []byte, n uint64) []byte {
	return append(dst, byte(n>>56), byte(n>>48), byte(n>>40), byte(n>>32), byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func AppendUint32BE(dst []byte, n uint32) []byte {
	return append(dst, byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func AppendUint16BE(dst []byte, n uint16) []byte {
	return append(dst, byte(n>>8), byte(n))
}

func Uint64LE(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Uint32LE(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func Uint16LE(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func Uint64BE(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func Uint32BE(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func Uint16BE(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}
