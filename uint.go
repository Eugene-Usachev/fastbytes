package fastbytes

import "unsafe"

// U2B converts an uint to a byte slice using unsafe. Return a byte slice of length 8!
func U2B(u uint) []byte {
	bytes := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&bytes[0])) = *(*uint64)(unsafe.Pointer(&u))
	return bytes
}

// PutU put an uint to a byte slice using unsafe. The byte slice len must be 8!
func PutU(b []byte, u uint) {
	*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&u))
}

// U8ToB converts an uint8 to a byte slice using unsafe
func U8ToB(u uint8) []byte {
	bytes := make([]byte, 1)
	bytes[0] = u
	return bytes
}

// PutU8 put an uint8 to a byte slice using unsafe
func PutU8(b []byte, u uint8) {
	b[0] = u
}

// U16ToB converts an uint16 to a byte slice using unsafe
func U16ToB(u uint16) []byte {
	bytes := make([]byte, 2)
	*(*uint16)(unsafe.Pointer(&bytes[0])) = u
	return bytes
}

// PutU16 put an uint16 to a byte slice using unsafe
func PutU16(b []byte, u uint16) {
	*(*uint16)(unsafe.Pointer(&b[0])) = u
}

// U32ToB converts an uint32 to a byte slice using unsafe
func U32ToB(u uint32) []byte {
	bytes := make([]byte, 4)
	*(*uint32)(unsafe.Pointer(&bytes[0])) = u
	return bytes
}

// PutU32 put an uint32 to a byte slice using unsafe
func PutU32(b []byte, u uint32) {
	*(*uint32)(unsafe.Pointer(&b[0])) = u
}

// U64ToB converts an uint64 to a byte slice using unsafe
func U64ToB(u uint64) []byte {
	bytes := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&bytes[0])) = u
	return bytes
}

// PutU64 put an uint64 to a byte slice using unsafe
func PutU64(b []byte, u uint64) {
	*(*uint64)(unsafe.Pointer(&b[0])) = u
}

// B2U converts a byte slice to an uint
func B2U(b []byte) uint {
	return uint(b[0]) | uint(b[1])<<8 | uint(b[2])<<16 | uint(b[3])<<24
}

// B2U64 converts a byte slice to an uint64
func B2U64(b []byte) uint64 {
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// B2U32 converts a byte slice to an uint32
func B2U32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// B2U16 converts a byte slice to an uint16
func B2U16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

// B2U8 converts a byte slice to an uint8
func B2U8(b []byte) uint8 {
	return b[0]
}
