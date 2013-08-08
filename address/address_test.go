package address

import (
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

	if !ValidateIdentifier(addr.Identifier) {
		t.Error("Invalid address identifier %s\n", addr.Identifier)
	}

	valid, err := ValidateChecksum(addr.Identifier)
	if err != nil {
		t.Error(err.Error())
	}

	if !valid {
		t.Error("Address checksum incorrect\n")
	}

	ver, err := addr.Version()
	if err != nil {
		t.Error(err.Error())
	}

	if ver != 3 {
		t.Error("Address version incorrect. Want 3, got %d\n", ver)
	}

	stream, err := addr.Stream()
	if err != nil {
		t.Error(err.Error())
	}

	if stream != 1 {
		t.Error("Address stream number incorrect. Want 1, got %d\n", stream)
	}
}
