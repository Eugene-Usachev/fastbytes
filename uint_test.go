package fastbytes

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestU2B(t *testing.T) {
	u := uint(4294967294)
	u2b := U2B(u)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint32(slice, uint32(u))
	if bytes.Equal(u2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", u2b, slice)
	}
}

func TestPutU(t *testing.T) {
	u := uint(4294967294)
	var putU = make([]byte, 8)
	PutU(putU, u)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint32(slice, uint32(u))
	if bytes.Equal(putU, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putU, slice)
	}
}

func TestU8ToB(t *testing.T) {
	u := uint8(254)
	u2b := U8ToB(u)
	var slice = make([]byte, 1)
	slice[0] = u
	if bytes.Equal(u2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", u2b, slice)
	}
}

func TestPutU8(t *testing.T) {
	u := uint8(254)

	var putU = make([]byte, 1)
	PutU8(putU, u)
	var slice = make([]byte, 1)
	slice[0] = u
	if bytes.Equal(putU, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putU, slice)
	}
}

func TestU16ToB(t *testing.T) {
	u := uint16(65534)
	u2b := U16ToB(u)
	var slice = make([]byte, 2)
	binary.LittleEndian.PutUint16(slice, u)
	if bytes.Equal(u2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", u2b, slice)
	}
}

func TestPutU16(t *testing.T) {
	u := uint16(65534)
	var putU = make([]byte, 2)
	PutU16(putU, u)
	var slice = make([]byte, 2)
	binary.LittleEndian.PutUint16(slice, u)
	if bytes.Equal(putU, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putU, slice)
	}
}

func TestU32ToB(t *testing.T) {
	u := uint32(4294967294)
	u2b := U32ToB(u)
	var slice = make([]byte, 4)
	binary.LittleEndian.PutUint32(slice, u)
	if bytes.Equal(u2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", u2b, slice)
	}
}

func TestPutU32(t *testing.T) {
	u := uint32(4294967294)
	putU := make([]byte, 4)
	PutU32(putU, u)
	var slice = make([]byte, 4)
	binary.LittleEndian.PutUint32(slice, u)
	if bytes.Equal(putU, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putU, slice)
	}
}

func TestU64ToB(t *testing.T) {
	u := uint64((2 << 63) - 2)
	u2b := U64ToB(u)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, u)
	if bytes.Equal(u2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", u2b, slice)
	}
}

func TestPutU64(t *testing.T) {
	u := uint64((2 << 63) - 2)
	putU := make([]byte, 8)
	PutU64(putU, u)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, u)
	if bytes.Equal(putU, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putU, slice)
	}
}

func TestB2U(t *testing.T) {
	b := U2B(4294967294)
	b2u := B2U(b)
	u := binary.LittleEndian.Uint32(b)
	if b2u == uint(u) {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2u, u)
	}
}

func TestB2U8(t *testing.T) {
	b := U8ToB(254)
	b2u := B2U8(b)
	u := b[0]
	if b2u == u {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2u, u)
	}
}

func TestB2U16(t *testing.T) {
	b := U16ToB(65534)
	b2u := B2U16(b)
	u := binary.LittleEndian.Uint16(b)
	if b2u == u {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2u, u)
	}
}

func TestB2U32(t *testing.T) {
	b := U32ToB(4294967294)
	b2u := B2U32(b)
	u := binary.LittleEndian.Uint32(b)
	if b2u == u {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2u, u)
	}
}

func TestB2U64(t *testing.T) {
	b := U64ToB((2 << 63) - 2)
	b2u := B2U64(b)
	u := binary.LittleEndian.Uint64(b)
	if b2u == u {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2u, u)
	}
}
