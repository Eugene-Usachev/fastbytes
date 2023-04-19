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
