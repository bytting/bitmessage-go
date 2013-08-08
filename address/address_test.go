package address

import (
	"strings"
	"testing"
)

func TestAddress(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	addr, err := New(3, 1, false)
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.HasPrefix(addr.Identifier, "BM-2D") {
		t.Error("Address does not start with correct prefix. Want BM-2D, got %s\n", addr.Identifier[:5])
	}

	valid, err := ValidateChecksum(addr.Identifier)
	if err != nil {
		t.Error(err.Error())
	}

	if !valid {
		t.Error("Address checksum incorrect\n")
	}

	stream, err := GetStream(addr.Identifier)
	if err != nil {
		t.Error(err.Error())
	}
	if stream != 1 {
		t.Error("Address stream number incorrect. Want 1, got %d\n", stream)
	}
}
