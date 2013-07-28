package varint

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	val1 := uint64(1)
	buf := Encode(val1)
	val2, nb := Decode(buf)

	if nb != 1 {
		t.Error("Reported varint size is not correct. Expected 1, got ", nb)
	}

	if val1 != val2 {
		t.Error("Decoded varint does not match. Expected 1, got ", val2)
	}

	val1 = uint64(12345)
	buf = Encode(val1)
	val2, nb = Decode(buf)

	if nb != 3 {
		t.Error("Reported varint size is not correct. Expected 3, got ", nb)
	}

	if val1 != val2 {
		t.Error("Decoded varint does not match. Expected 12345, got ", val2)
	}

	val1 = uint64(1234567890)
	buf = Encode(val1)
	val2, nb = Decode(buf)

	if nb != 5 {
		t.Error("Reported varint size is not correct. Expected 5, got ", nb)
	}

	if val1 != val2 {
		t.Error("Decoded varint does not match. Expected 1234567890, got ", val2)
	}

	val1 = uint64(1234567890123456)
	buf = Encode(val1)
	val2, nb = Decode(buf)

	if nb != 9 {
		t.Error("Reported varint size is not correct. Expected 9, got ", nb)
	}

	if val1 != val2 {
		t.Error("Decoded varint does not match. Expected 1234567890123456, got ", val2)
	}
}
