package pow

import (
	"bitmessage/varint"
	"io/ioutil"
	"testing"
)

func TestPOW(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	payload, err := ioutil.ReadFile("pow.go")
	if err != nil {
		t.Error("Unable to read from file\n")
	}

	nonce := Nonce(payload)

	new_payload := varint.Encode(nonce)
	new_payload = append(new_payload, payload...)

	if !Validate(new_payload) {
		t.Error("Nonce %d is not valid for payload", nonce)
	}
}
