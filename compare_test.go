package fastbytes

import "testing"

func TestEqual(t *testing.T) {
	slice1 := []byte{1, 2, 3}
	slice2 := []byte{2, 3, 4}

	if !Equal(slice1, slice1) {
		t.Error(slice1, "!=", slice1)
	}

	if Equal(slice1, slice2) {
		t.Error(slice1, "==", slice2)
	}
}
