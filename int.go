package fastbytes

import "unsafe"

// I2B converts an int to a byte slice using unsafe. Return a byte slice of length 8!
func I2B(i int) []byte {
	bytes := make([]byte, 8)
	*(*int64)(unsafe.Pointer(&bytes[0])) = *(*int64)(unsafe.Pointer(&i))
	return bytes
}

// PutI put an int to a byte slice using unsafe. The byte slice len must be 8 or the byte slice must be nil!
func PutI(b []byte, i int) {
	*(*int64)(unsafe.Pointer(&b[0])) = *(*int64)(unsafe.Pointer(&i))
}

// I8ToB converts an int8 to a byte slice using unsafe
func I8ToB(i int8) []byte {
	bytes := make([]byte, 1)
	bytes[0] = byte(i)
	return bytes
}

// PutI8 put an int8 to a byte slice using unsafe
func PutI8(b []byte, i int8) {
	b[0] = byte(i)
}

// I16ToB converts an int16 to a byte slice using unsafe
func I16ToB(i int16) []byte {
	bytes := make([]byte, 2)
	*(*int16)(unsafe.Pointer(&bytes[0])) = i
	return bytes
}

// PutI16 put an int16 to a byte slice using unsafe
func PutI16(b []byte, i int16) {
	*(*int16)(unsafe.Pointer(&b[0])) = i
}

// I32ToB converts an int32 to a byte slice using unsafe
func I32ToB(i int32) []byte {
	bytes := make([]byte, 4)
	*(*int32)(unsafe.Pointer(&bytes[0])) = i
	return bytes
}

// PutI32 put an int32 to a byte slice using unsafe
func PutI32(b []byte, i int32) {
	*(*int32)(unsafe.Pointer(&b[0])) = i
}

// I64ToB converts an int64 to a byte slice using unsafe
func I64ToB(i int64) []byte {
	bytes := make([]byte, 8)
	*(*int64)(unsafe.Pointer(&bytes[0])) = i
	return bytes
}

// PutI64 put an int64 to a byte slice using unsafe
func PutI64(b []byte, i int64) {
	*(*int64)(unsafe.Pointer(&b[0])) = i
}

// B2I converts a byte slice to an int.
func B2I(b []byte) int {
	return int(int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24)
}

// B2I64 converts a byte slice to an int64
func B2I64(b []byte) int64 {
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}

// B2I32 converts a byte slice to an int32
func B2I32(b []byte) int32 {
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

// B2I16 converts a byte slice to an int16
func B2I16(b []byte) int16 {
	return int16(b[0]) | int16(b[1])<<8
}

// B2I8 converts a byte slice to an int8
func B2I8(b []byte) int8 {
	return int8(b[0])
}
