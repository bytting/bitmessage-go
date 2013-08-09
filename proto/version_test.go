package proto

import (
	"testing"
)

func TestVersion(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// t.Error("Command has changed after serialize/deserialize")
}
