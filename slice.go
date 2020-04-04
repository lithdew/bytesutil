package bytesutil

import (
	"math/rand"
	"reflect"
	"unsafe"
)

const (
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetIdxBits = 6
	charsetIdxMask = 1<<charsetIdxBits - 1
	charsetIdxMax  = 63 / charsetIdxBits
)

func RandomSlice(dst []byte) []byte {
	n := len(dst)

	for i, cache, remain := n-1, rand.Int63(), charsetIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), charsetIdxMax
		}

		if idx := int(cache & charsetIdxMask); idx < len(charset) {
			dst[i] = charset[idx]
			i--
		}

		cache >>= charsetIdxBits
		remain--
	}

	return dst
}
func ExtendSlice(dst []byte, size int) []byte {
	n := size - cap(dst)
	if n > 0 {
		dst = append(dst[:cap(dst)], make([]byte, n)...)
	}
	return dst[:size]
}

func String(b []byte) string {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: header.Data, Len: header.Len}))
}

func Slice(s string) []byte {
	header := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: header.Data, Len: header.Len, Cap: header.Len}))
}
