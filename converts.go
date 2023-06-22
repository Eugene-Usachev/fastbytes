package fastbytes

import (
	"reflect"
	"unsafe"
)

// S2B converts a string to a byte slice
func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

// U2B converts an uint to a byte slice using unsafe
func U2B(u uint) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&u))
	h.Len = int(unsafe.Sizeof(u))
	h.Cap = h.Len
	return b
}

// U8ToB converts an uint8 to a byte slice using unsafe
func U8ToB(u uint8) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&u))
	h.Len = int(unsafe.Sizeof(u))
	h.Cap = h.Len
	return b
}

// U16ToB converts an uint16 to a byte slice using unsafe
func U16ToB(u uint16) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&u))
	h.Len = int(unsafe.Sizeof(u))
	h.Cap = h.Len
	return b
}

// U32ToB converts an uint32 to a byte slice using unsafe
func U32ToB(u uint32) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&u))
	h.Len = int(unsafe.Sizeof(u))
	h.Cap = h.Len
	return b
}

// U64ToB converts an uint64 to a byte slice using unsafe
func U64ToB(u uint64) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&u))
	h.Len = int(unsafe.Sizeof(u))
	h.Cap = h.Len
	return b
}

// I2B converts an int to a byte slice using unsafe
func I2B(i int) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&i))
	h.Len = int(unsafe.Sizeof(i))
	h.Cap = h.Len
	return b
}

// I8ToB converts an int8 to a byte slice using unsafe
func I8ToB(i int8) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&i))
	h.Len = int(unsafe.Sizeof(i))
	h.Cap = h.Len
	return b
}

// I16ToB converts an int16 to a byte slice using unsafe
func I16ToB(i int16) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&i))
	h.Len = int(unsafe.Sizeof(i))
	h.Cap = h.Len
	return b
}

// I32ToB converts an int32 to a byte slice using unsafe
func I32ToB(i int32) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&i))
	h.Len = int(unsafe.Sizeof(i))
	h.Cap = h.Len
	return b
}

// I64ToB converts an int64 to a byte slice using unsafe
func I64ToB(i int64) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = uintptr(unsafe.Pointer(&i))
	h.Len = int(unsafe.Sizeof(i))
	h.Cap = h.Len
	return b
}

// B2S converts a byte slice to a string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
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

// B2U64 converts a byte slice to an uint64
func B2U64(s []byte) (res uint64, ok bool) {
	res = 0
	for _, c := range s {
		if v := uint64(c - '0'); v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	return res, ok
}

// B2I32 converts a byte slice to an int32
func B2I32(s []byte) (res int32, ok bool) {
	sign := len(s) > 0 && s[0] == '-'
	if sign {
		s = s[1:]
	}

	ok = true

	res = 0
	for _, c := range s {
		if v := int32(c - '0'); v < 0 || v > 9 {
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

// B2U32 converts a byte slice to an uint32
func B2U32(s []byte) (res uint32, ok bool) {
	res = 0
	for _, c := range s {
		if v := uint32(c - '0'); v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	return res, ok
}

// B2I16 converts a byte slice to an int16
func B2I16(s []byte) (res int16, ok bool) {
	sign := len(s) > 0 && s[0] == '-'
	if sign {
		s = s[1:]
	}

	ok = true

	res = 0
	for _, c := range s {
		if v := int16(c - '0'); v < 0 || v > 9 {
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

// B2U16 converts a byte slice to an uint16
func B2U16(s []byte) (res uint16, ok bool) {
	res = 0
	for _, c := range s {
		if v := uint16(c - '0'); v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	return res, ok
}

// B2I8 converts a byte slice to an int8
func B2I8(s []byte) (res int8, ok bool) {
	sign := len(s) > 0 && s[0] == '-'
	if sign {
		s = s[1:]
	}

	ok = true

	res = 0
	for _, c := range s {
		if v := int8(c - '0'); v < 0 || v > 9 {
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

// B2U8 converts a byte slice to an uint8
func B2U8(s []byte) (res uint8, ok bool) {
	res = 0
	for _, c := range s {
		if v := c - '0'; v < 0 || v > 9 {
			ok = false
			break
		} else {
			res = res*10 + v
		}
	}

	return res, ok
}
