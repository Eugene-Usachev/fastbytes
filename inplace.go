package fastbytes

func bytesToLowerInplace(buf []byte) {
	for i, ch := range buf {
		if ch >= 'A' && ch <= 'Z' {
			buf[i] += 'a' - 'A'
		}
	}
}

func bytesTrimLeftInplace(buf []byte) []byte {
	i, l := 0, len(buf)
	for ; i < l && buf[i] == ' '; i++ {
	}
	return buf[i:]
}
