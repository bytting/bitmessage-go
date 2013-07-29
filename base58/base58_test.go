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

	encoded, err := Encode(buf.Bytes())
	if err != nil {
		t.Error(err.Error())
	}

	decoded, err := Decode(encoded)
	if err != nil {
		t.Error(err.Error())
	}

	if string(decoded) != buf.String() {
		t.Error("Decoded base58 does not match. Expected %s, got %s", buf.String(), string(decoded))
	}

}
