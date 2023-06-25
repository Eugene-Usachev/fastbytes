package fastbytes

// The same

//func BenchmarkU2B(b *testing.B) {
//	u := uint(4294967294)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 8)
//		slice = U2B(u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkPutU(b *testing.B) {
//	u := uint(4294967294)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 8)
//		PutU(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryU2B(b *testing.B) {
//	u := uint(4294967294)
//	u32 := uint32(u)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 8)
//		binary.LittleEndian.PutUint32(slice, u32)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkU8ToB(b *testing.B) {
//	u := uint8(255)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 1)
//		slice = U8ToB(u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkPutU8(b *testing.B) {
//	u := uint8(255)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 1)
//		PutU8(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryU8ToB(b *testing.B) {
//	u := uint8(255)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 1)
//		slice[0] = u
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkU16ToB(b *testing.B) {
//	u := uint16(65534)
//	for i := 0; i < b.N; i++ {
//		var slice []byte
//		slice = U16ToB(u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkPutU16(b *testing.B) {
//	u := uint16(65534)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 2)
//		PutU16(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryU16ToB(b *testing.B) {
//	u := uint16(65534)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 2)
//		binary.LittleEndian.PutUint16(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkU32ToB(b *testing.B) {
//	u := uint32(4294967294)
//	for i := 0; i < b.N; i++ {
//		var slice []byte
//		slice = U32ToB(u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkPutU32(b *testing.B) {
//	u := uint32(4294967294)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 4)
//		PutU32(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryU32ToB(b *testing.B) {
//	u := uint32(4294967294)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 4)
//		binary.LittleEndian.PutUint32(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkU64ToB(b *testing.B) {
//	u := uint64((2 << 63) - 2)
//	for i := 0; i < b.N; i++ {
//		var slice []byte
//		slice = U64ToB(u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkPutU64(b *testing.B) {
//	u := uint64((2 << 63) - 2)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 8)
//		PutU64(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryU64ToB(b *testing.B) {
//	u := uint64((2 << 63) - 2)
//	for i := 0; i < b.N; i++ {
//		var slice = make([]byte, 8)
//		binary.LittleEndian.PutUint64(slice, u)
//		if len(slice) < 1 {
//		}
//	}
//}

// The same

//func BenchmarkB2U(b *testing.B) {
//	slice := U2B(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := B2U(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2U(b *testing.B) {
//	slice := U2B(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := binary.LittleEndian.Uint32(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkB2U8(b *testing.B) {
//	slice := U8ToB(254)
//	for i := 0; i < b.N; i++ {
//		u := B2U8(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2U8(b *testing.B) {
//	slice := U8ToB(254)
//	for i := 0; i < b.N; i++ {
//		u := slice[0]
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkB2U16(b *testing.B) {
//	slice := U16ToB(65534)
//	for i := 0; i < b.N; i++ {
//		u := B2U16(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2U16(b *testing.B) {
//	slice := U16ToB(65534)
//	for i := 0; i < b.N; i++ {
//		u := binary.LittleEndian.Uint16(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkB2U32(b *testing.B) {
//	slice := U32ToB(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := B2U32(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2U32(b *testing.B) {
//	slice := U32ToB(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := binary.LittleEndian.Uint32(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkB2U64(b *testing.B) {
//	slice := U64ToB(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := B2U64(slice)
//		if u < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2U64(b *testing.B) {
//	slice := U64ToB(4294967294)
//	for i := 0; i < b.N; i++ {
//		u := binary.LittleEndian.Uint64(slice)
//		if u < 1 {
//		}
//	}
//}
