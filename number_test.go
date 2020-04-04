package bytesutil

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"testing/quick"
)

func TestAppendNumbers(t *testing.T) {
	f := func(a uint64, b uint32, c uint16) bool {
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:8], a)
		if !assert.EqualValues(t, buf[:], AppendUint64LE(nil, a)) {
			return false
		}
		if !assert.EqualValues(t, a, Uint64LE(buf[:8])) {
			return false
		}
		binary.LittleEndian.PutUint32(buf[:4], b)
		if !assert.EqualValues(t, buf[:4], AppendUint32LE(nil, b)) {
			return false
		}
		if !assert.EqualValues(t, b, Uint32LE(buf[:4])) {
			return false
		}
		binary.LittleEndian.PutUint16(buf[:2], c)
		if !assert.EqualValues(t, buf[:2], AppendUint16LE(nil, c)) {
			return false
		}
		if !assert.EqualValues(t, c, Uint16LE(buf[:2])) {
			return false
		}

		binary.BigEndian.PutUint64(buf[:8], a)
		if !assert.EqualValues(t, buf[:], AppendUint64BE(nil, a)) {
			return false
		}
		if !assert.EqualValues(t, a, Uint64BE(buf[:8])) {
			return false
		}
		binary.BigEndian.PutUint32(buf[:4], b)
		if !assert.EqualValues(t, buf[:4], AppendUint32BE(nil, b)) {
			return false
		}
		if !assert.EqualValues(t, b, Uint32BE(buf[:4])) {
			return false
		}
		binary.BigEndian.PutUint16(buf[:2], c)
		if !assert.EqualValues(t, buf[:2], AppendUint16BE(nil, c)) {
			return false
		}
		if !assert.EqualValues(t, c, Uint16BE(buf[:2])) {
			return false
		}

		return true
	}

	require.NoError(t, quick.Check(f, nil))
}

func TestAppendVarInt(t *testing.T) {
	a := func(a uint64, b uint32, c uint16) bool {
		var buf [binary.MaxVarintLen64]byte

		size := binary.PutUvarint(buf[:], a)
		if !assert.EqualValues(t, buf[:size], AppendUvarInt(nil, a)) {
			return false
		}
		size = binary.PutUvarint(buf[:], uint64(b))
		if !assert.EqualValues(t, buf[:size], AppendUvarInt(nil, uint64(b))) {
			return false
		}
		size = binary.PutUvarint(buf[:], uint64(c))
		if !assert.EqualValues(t, buf[:size], AppendUvarInt(nil, uint64(c))) {
			return false
		}

		return true
	}

	require.NoError(t, quick.Check(a, nil))

	b := func(a int64, b int32, c int16) bool {
		var buf [binary.MaxVarintLen64]byte

		size := binary.PutVarint(buf[:], a)
		if !assert.EqualValues(t, buf[:size], AppendVarInt(nil, a)) {
			return false
		}
		size = binary.PutVarint(buf[:], int64(b))
		if !assert.EqualValues(t, buf[:size], AppendVarInt(nil, int64(b))) {
			return false
		}
		size = binary.PutVarint(buf[:], int64(c))
		if !assert.EqualValues(t, buf[:size], AppendVarInt(nil, int64(c))) {
			return false
		}

		return true
	}

	require.NoError(t, quick.Check(b, nil))
}
