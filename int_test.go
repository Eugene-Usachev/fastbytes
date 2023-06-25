package fastbytes

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestI2B(t *testing.T) {
	i := int(-214748364)
	i2b := I2B(i)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, uint64(i))
	if bytes.Equal(i2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", i2b, slice)
	}
}

func TestPutI(t *testing.T) {
	i := int(-214748364)
	var putI = make([]byte, 8)
	PutI(putI, i)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, uint64(i))
	if bytes.Equal(putI, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putI, slice)
	}
}

func TestI8ToB(t *testing.T) {
	i := int8(-127)
	i2b := I8ToB(i)
	var slice = make([]byte, 1)
	slice[0] = byte(i)
	if bytes.Equal(i2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", i2b, slice)
	}
}

func TestPutI8(t *testing.T) {
	i := int8(-127)
	var putI = make([]byte, 1)
	PutI8(putI, i)
	var slice = make([]byte, 1)
	slice[0] = byte(i)
	if bytes.Equal(putI, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putI, slice)
	}
}

func TestI16ToB(t *testing.T) {
	i := int16(-32767)
	i2b := I16ToB(i)
	var slice = make([]byte, 2)
	binary.LittleEndian.PutUint16(slice, uint16(i))
	if bytes.Equal(i2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", i2b, slice)
	}
}

func TestPutI16(t *testing.T) {
	i := int16(-32767)
	var putI = make([]byte, 2)
	PutI16(putI, i)
	var slice = make([]byte, 2)
	binary.LittleEndian.PutUint16(slice, uint16(i))
	if bytes.Equal(putI, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putI, slice)
	}
}

func TestI32ToB(t *testing.T) {
	i := int32(-2147483647)
	i2b := I32ToB(i)
	var slice = make([]byte, 4)
	binary.LittleEndian.PutUint32(slice, uint32(i))
	if bytes.Equal(i2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", i2b, slice)
	}
}

func TestPutI32(t *testing.T) {
	i := int32(-2147483647)
	var putI = make([]byte, 4)
	PutI32(putI, i)
	var slice = make([]byte, 4)
	binary.LittleEndian.PutUint32(slice, uint32(i))
	if bytes.Equal(putI, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putI, slice)
	}
}

func TestI64ToB(t *testing.T) {
	i := int64(-9223372036854775808)
	i2b := I64ToB(i)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, uint64(i))
	if bytes.Equal(i2b, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", i2b, slice)
	}
}

func TestPutI64(t *testing.T) {
	i := int64(-9223372036854775808)
	var putI = make([]byte, 8)
	PutI64(putI, i)
	var slice = make([]byte, 8)
	binary.LittleEndian.PutUint64(slice, uint64(i))
	if bytes.Equal(putI, slice) {
		t.Log("Pass")
	} else {
		t.Error("Fail", putI, slice)
	}
}

func TestB2I(t *testing.T) {
	b := I2B(-32000)
	b2i := B2I(b)
	i := binary.LittleEndian.Uint64(b)
	if int(i) == b2i {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2i, i)
	}
}

func TestB2I8(t *testing.T) {
	b := I8ToB(-127)
	b2i := B2I8(b)
	i := b[0]
	if int8(i) == b2i {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2i, i)
	}
}

func TestB2I16(t *testing.T) {
	b := I16ToB(-32767)
	b2i := B2I16(b)
	i := binary.LittleEndian.Uint16(b)
	if int16(i) == b2i {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2i, i)
	}
}

func TestB2I32(t *testing.T) {
	b := I32ToB(-2147483647)
	b2i := B2I32(b)
	i := binary.LittleEndian.Uint32(b)
	if int32(i) == b2i {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2i, i)
	}
}

func TestB2I64(t *testing.T) {
	b := I64ToB(-9223372036854775808)
	b2i := B2I64(b)
	i := binary.LittleEndian.Uint64(b)
	if int64(i) == b2i {
		t.Log("Pass")
	} else {
		t.Error("Fail", b2i, i)
	}
}
