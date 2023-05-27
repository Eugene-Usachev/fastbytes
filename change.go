package fastbytes

// BytesToLowerCase converts all bytes in buf to lowercase.
func BytesToLowerCase(buf []byte) {
	for i, ch := range buf {
		if ch >= 'A' && ch <= 'Z' {
			buf[i] += 'a' - 'A'
		}
	}
}

// BytesToUpperCase converts all bytes in buf to uppercase.
func BytesToUpperCase(buf []byte) {
	for i, ch := range buf {
		if ch >= 'a' && ch <= 'z' {
			buf[i] -= 'a' - 'A'
		}
	}
}

// BytesTrimLeft trims all empty bytes (' ') from the left and returns a new slice.
func BytesTrimLeft(buf []byte) []byte {
	i, l := 0, len(buf)
	for ; i < l && buf[i] == ' '; i++ {
	}
	return buf[i:]
}
