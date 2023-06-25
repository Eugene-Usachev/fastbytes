package fastbytes

import (
	"strings"
	"testing"
)

var (
	little = 10
	medium = 250
	large  = 5000

	littleSlice = func() string {
		builder := strings.Builder{}

		for i := 0; i < little; i++ {
			builder.WriteString("a")
		}

		return builder.String()
	}()

	mediumSlice = func() string {
		builder := strings.Builder{}

		for i := 0; i < medium; i++ {
			builder.WriteString("a")
		}

		return builder.String()
	}()

	largeSlice = func() string {
		builder := strings.Builder{}

		for i := 0; i < large; i++ {
			builder.WriteString("a")
		}

		return builder.String()
	}()

	littleSliceBytes = func() []byte {
		slice := make([]byte, little)

		for i := 0; i < little; i++ {
			slice = append(slice, byte('a'))
		}

		return slice
	}()

	mediumSliceBytes = func() []byte {
		slice := make([]byte, medium)

		for i := 0; i < medium; i++ {
			slice = append(slice, byte('a'))
		}

		return slice
	}()

	largeSliceBytes = func() []byte {
		slice := make([]byte, large)

		for i := 0; i < large; i++ {
			slice = append(slice, byte('a'))
		}

		return slice
	}()
)

func BenchmarkS2B_Little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := S2B(littleSlice)
		if len(slice) < little {
		}
	}
}

func BenchmarkS2BStd_Little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []byte(littleSlice)
		if len(slice) < little {
		}
	}
}

func BenchmarkS2B_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := S2B(mediumSlice)
		if len(slice) < medium {
		}
	}
}

func BenchmarkS2BStd_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []byte(mediumSlice)
		if len(slice) < medium {
		}
	}
}

func BenchmarkS2B_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := S2B(largeSlice)
		if len(slice) < large {
		}
	}
}

func BenchmarkS2BStd_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []byte(largeSlice)
		if len(slice) < large {
		}
	}
}

func BenchmarkB2S_Little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := B2S(littleSliceBytes)
		if len(str) < little {
		}
	}
}

func BenchmarkB2SStd_Little(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := string(littleSliceBytes)
		if len(str) < little {
		}
	}
}

func BenchmarkB2S_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := B2S(mediumSliceBytes)
		if len(str) < medium {
		}
	}
}

func BenchmarkB2SStd_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := string(mediumSliceBytes)
		if len(str) < medium {
		}
	}
}

func BenchmarkB2S_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := B2S(largeSliceBytes)
		if len(str) < large {
		}
	}
}

func BenchmarkB2SStd_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := string(largeSliceBytes)
		if len(str) < large {
		}
	}
}
