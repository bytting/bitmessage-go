package msg

import (
	"testing"
)

func TestMessage(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// t.Error("Reported varint size is not correct. Expected 1, got ", nb)
}
