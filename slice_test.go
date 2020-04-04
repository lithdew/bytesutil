package bytesutil

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"testing/quick"
)

func TestRandomSlice(t *testing.T) {
	f := func(a uint16) bool {
		slice := make([]byte, a)
		return assert.EqualValues(t, slice, RandomSlice(slice))
	}

	require.NoError(t, quick.Check(f, nil))
}

func TestExtendSlice(t *testing.T) {
	a := make([]byte, 0, 30)
	require.Equal(t, 30, cap(a))

	a = ExtendSlice(a, 35)
	require.Equal(t, 64, cap(a))
}

func TestString(t *testing.T) {
	f := func(a []byte) bool {
		return assert.EqualValues(t, string(a), String(a))
	}

	require.NoError(t, quick.Check(f, nil))
}

func TestSlice(t *testing.T) {
	f := func(a string) bool {
		return assert.EqualValues(t, []byte(a), Slice(a))
	}

	require.NoError(t, quick.Check(f, nil))
}

func TestStringSlice(t *testing.T) {
	f := func(a []byte) bool {
		return assert.EqualValues(t, []byte(string(a)), Slice(String(a)))
	}

	require.NoError(t, quick.Check(f, nil))
}

var (
	sinkSlice  []byte
	sinkString string
)

func BenchmarkStdStringSlice(b *testing.B) {
	a := "hello"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkSlice = []byte(a)
		sinkString = string(sinkSlice)
	}

	_, _ = sinkSlice, sinkString
}

func BenchmarkStringSlice(b *testing.B) {
	a := "hello"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sinkSlice = Slice(a)
		sinkString = String(sinkSlice)
	}
}
