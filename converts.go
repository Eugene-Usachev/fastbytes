package fastbytes

import (
	"reflect"
	"unsafe"
)

// B2S converts a byte slice to a string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts a string to a byte slice
func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// B2I64 converts a byte slice to an int64
func B2I64(s []byte) (res int64, ok bool) {
	sign := len(s) > 0 && s[0] == '-'
	if sign {
		s = s[1:]
	}

	ok = true

	res = 0
	for _, c := range s {
		if v := int64(c - '0'); v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	if sign {
		res = -res
	}

	return
}
