package address

import (
	"strings"
	"testing"
)

func TestAddress(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	addr := New(3, 1, false)

	if !strings.HasPrefix(addr.Identifier, "BM-2D") {
		t.Error("Address does not start with correct prefix. Want BM-2D, got %s\n", addr.Identifier[:5])
	}

	if !Validate(addr.Identifier) {
		t.Error("Address checksum incorrect\n")
	}

	if GetStream(addr.Identifier) != 1 {
		t.Error("Address stream number incorrect. Want 1, got %d\n", GetStream(addr.Identifier))
	}
}
