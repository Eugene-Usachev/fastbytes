package fastbytes

// The same

/*func BenchmarkI2B(b *testing.B) {
	intVar := -2147483648
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 8)
		slice = I2B(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkPutI(b *testing.B) {
	intVar := -2147483648
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 8)
		PutI(slice, intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkBinaryI2B(b *testing.B) {
	intVar := -2147483648
	i32 := int32(intVar)
	for i := 0; i < b.N; i++ {
		u32 := uint32(i32)
		var slice = make([]byte, 8)
		binary.LittleEndian.PutUint32(slice, u32)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkI8ToB(b *testing.B) {
	intVar := int8(-128)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 1)
		slice = I8ToB(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkPutI8(b *testing.B) {
	intVar := int8(-128)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 1)
		PutI8(slice, intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkBinaryI8ToB(b *testing.B) {
	intVar := int8(-128)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 1)
		slice[0] = byte(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkI16ToB(b *testing.B) {
	intVar := int16(-32768)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 2)
		slice = I16ToB(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkPutI16(b *testing.B) {
	intVar := int16(-32768)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 2)
		PutI16(slice, intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkBinaryI16ToB(b *testing.B) {
	intVar := int16(-32768)
	for i := 0; i < b.N; i++ {
		i16 := uint16(intVar)
		var slice = make([]byte, 8)
		binary.LittleEndian.PutUint16(slice, i16)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkI32ToB(b *testing.B) {
	intVar := int32(-2147483648)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 4)
		slice = I32ToB(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkPutI32(b *testing.B) {
	intVar := int32(-2147483648)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 4)
		PutI32(slice, intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkBinaryI32ToB(b *testing.B) {
	intVar := int32(-2147483648)
	for i := 0; i < b.N; i++ {
		i32 := uint32(intVar)
		var slice = make([]byte, 4)
		binary.LittleEndian.PutUint32(slice, i32)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkI64ToB(b *testing.B) {
	intVar := int64(-9223372036854775808)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 8)
		slice = I64ToB(intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkPutI64(b *testing.B) {
	intVar := int64(-9223372036854775808)
	for i := 0; i < b.N; i++ {
		var slice = make([]byte, 8)
		PutI64(slice, intVar)
		if len(slice) > 1 {
		}
	}
}

func BenchmarkBinaryI64ToB(b *testing.B) {
	intVar := int64(-9223372036854775808)
	for i := 0; i < b.N; i++ {
		i64 := uint64(intVar)
		var slice = make([]byte, 8)
		binary.LittleEndian.PutUint64(slice, i64)
		if len(slice) > 1 {
		}
	}
}*/

// The same

//func BenchmarkB2I(b *testing.B) {
//	slice := I2B(-2147483648)
//	for i := 0; i < b.N; i++ {
//		intVar := B2I(slice)
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2I(b *testing.B) {
//	slice := I2B(-2147483648)
//	for i := 0; i < b.N; i++ {
//		intVar := int(binary.LittleEndian.Uint32(slice))
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkB2I8(b *testing.B) {
//	slice := I8ToB(-128)
//	for i := 0; i < b.N; i++ {
//		intVar := B2I8(slice)
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2I8(b *testing.B) {
//	slice := I8ToB(-128)
//	for i := 0; i < b.N; i++ {
//		intVar := slice[0]
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkB2I16(b *testing.B) {
//	slice := I16ToB(-32768)
//	for i := 0; i < b.N; i++ {
//		intVar := B2I16(slice)
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2I16(b *testing.B) {
//	slice := I16ToB(-32768)
//	for i := 0; i < b.N; i++ {
//		intVar := int(binary.LittleEndian.Uint16(slice))
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkB2I32(b *testing.B) {
//	slice := I32ToB(-2147483648)
//	for i := 0; i < b.N; i++ {
//		intVar := B2I32(slice)
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2I32(b *testing.B) {
//	slice := I32ToB(-2147483648)
//	for i := 0; i < b.N; i++ {
//		intVar := int(binary.LittleEndian.Uint32(slice))
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkB2I64(b *testing.B) {
//	slice := I64ToB(-9223372036854775808)
//	for i := 0; i < b.N; i++ {
//		intVar := B2I64(slice)
//		if intVar < 1 {
//		}
//	}
//}
//
//func BenchmarkBinaryB2I64(b *testing.B) {
//	slice := I64ToB(-9223372036854775808)
//	for i := 0; i < b.N; i++ {
//		intVar := int(binary.LittleEndian.Uint64(slice))
//		if intVar < 1 {
//		}
//	}
//}
