package base58

import (
	"bytes"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var buf bytes.Buffer
	buf.WriteString("This is a string")

	encoded := Encode(buf.Bytes())
	decoded := Decode(encoded)

	if string(decoded) != buf.String() {
		t.Error("Decoded base58 does not match. Expected %s, got %s", buf.String(), string(decoded))
	}

}
