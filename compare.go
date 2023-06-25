package fastbytes

// Equal returns true if the two bytes slice are equal. If you are sure that two bytes slice have same length, use EqualValue instead.
func Equal(a, b []byte) bool {
	return B2S(a) == B2S(b)
}
