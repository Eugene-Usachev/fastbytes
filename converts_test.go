package fastbytes

import "testing"

func TestB2S(t *testing.T) {
	b := []byte("abc")
	s := B2S(b)
	if s != "abc" {
		t.Errorf("expected %s, got %s", "abc", s)
	}
}

func TestS2B(t *testing.T) {
	s := "abc"
	b := S2B(s)
	if !Equal(b, []byte("abc")) {
		t.Errorf("expected %s, got %s", b, []byte("abc"))
	}
}
