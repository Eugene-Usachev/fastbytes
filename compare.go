package fastbytes

// Equal returns true if the two bytes slice are equal. If you are sure that two bytes slice have same length, use EqualValue instead.
func Equal(a, b []byte) bool {
	shA := B2S(a)
	shB := B2S(b)
	return len(shA) == len(shB) && shA == shB
}

// EqualValue returns true if the two bytes slice are equal. If you are not sure that two bytes slice have same length, use Equal instead.
func EqualValue(a, b []byte) bool {
	return B2S(a) == B2S(b)
}
