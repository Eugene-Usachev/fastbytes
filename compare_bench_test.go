package fastbytes

// The same

//var (
//	little = 2
//	medium = 15
//	large  = 1000
//
//	littleSlice = func() []byte {
//		slice := make([]byte, little)
//		for i := 0; i < little; i++ {
//			slice[i] = byte(i % 255)
//		}
//		return slice
//	}()
//
//	mediumSlice = func() []byte {
//		slice := make([]byte, medium)
//		for i := 0; i < medium; i++ {
//			slice[i] = byte(i % 255)
//		}
//		return slice
//	}()
//
//	largeSlice = func() []byte {
//		slice := make([]byte, large)
//		for i := 0; i < large; i++ {
//			slice[i] = byte(i % 255)
//		}
//		return slice
//	}()
//
//	littleSlice2 = func() []byte {
//		slice := make([]byte, little)
//		for i := 0; i < little; i++ {
//			slice[i] = byte(i % 255)
//		}
//		slice[little-1] /= 2
//		return slice
//	}()
//
//	mediumSlice2 = func() []byte {
//		slice := make([]byte, medium)
//		for i := 0; i < medium; i++ {
//			slice[i] = byte(i % 255)
//		}
//		slice[medium-1] /= 2
//		return slice
//	}()
//
//	largeSlice2 = func() []byte {
//		slice := make([]byte, large)
//		for i := 0; i < large; i++ {
//			slice[i] = byte(i % 255)
//		}
//		slice[large-1] /= 2
//		return slice
//	}()
//)
//
//func BenchmarkEqual_Little(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(littleSlice, littleSlice)
//	}
//}
//
//func BenchmarkEqualBytes_Little(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(littleSlice, littleSlice)
//	}
//}
//
//func BenchmarkEqual_Medium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(mediumSlice, mediumSlice)
//	}
//}
//
//func BenchmarkEqualBytes_Medium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(mediumSlice, mediumSlice)
//	}
//}
//
//func BenchmarkEqual_Large(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(largeSlice, largeSlice)
//	}
//}
//
//func BenchmarkEqualBytes_Large(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(largeSlice, largeSlice)
//	}
//}
//
//func BenchmarkEqualDifferentLittle(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(littleSlice2, littleSlice)
//	}
//}
//
//func BenchmarkEqualBytesDifferentLittle(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(littleSlice2, littleSlice)
//	}
//}
//
//func BenchmarkEqualDifferentMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(mediumSlice2, mediumSlice)
//	}
//}
//
//func BenchmarkEqualBytesDifferentMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(mediumSlice2, mediumSlice)
//	}
//}
//
//func BenchmarkEqualDifferentLarge(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(largeSlice2, largeSlice)
//	}
//}
//
//func BenchmarkEqualBytesDifferentLarge(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(largeSlice2, largeSlice)
//	}
//}
//
//func BenchmarkEqual_LittleAndMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(littleSlice, mediumSlice)
//	}
//}
//
//func BenchmarkEqualBytes_LittleAndMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(littleSlice, mediumSlice)
//	}
//}
//
//func BenchmarkEqual_LargeAndMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Equal(largeSlice, mediumSlice)
//	}
//}
//
//func BenchmarkEqualBytes_LargeAndMedium(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		bytes.Equal(largeSlice, mediumSlice)
//	}
//}
