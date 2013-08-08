package proto

import (
	"bytes"
	"testing"
)

func TestProto(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	var buf bytes.Buffer
	buf.WriteString("A few bytes")

	m, err := NewMessageFromCommand("test", buf.Bytes())
	if err != nil {
		t.Error(err.Error())
	}

	s, err := m.Serialize()
	if err != nil {
		t.Error(err.Error())
	}

	n, _ := NewMessage()

	err = n.Deserialize(s)
	if err != nil {
		t.Error(err.Error())
	}

	if m.Command != n.Command {
		t.Error("Command has changed after serialize/deserialize")
	}

	if m.Length != n.Length {
		t.Error("Length has changed after serialize/deserialize")
	}

	if m.Magic != n.Magic {
		t.Error("Magic has changed after serialize/deserialize")
	}

	if !bytes.Equal(m.Payload, n.Payload) {
		t.Error("Payload has changed after serialize/deserialize")
	}

	if !bytes.Equal(m.Checksum, n.Checksum) {
		t.Error("Checksum has changed after serialize/deserialize")
	}

	// verack

	//addr, _ := NewAddr()

	//broadcast, _ := NewBroadcast()

	//getdata, _ := NewGetdata()

	//getpubkey, _ := NewGetpubkey()

	//inv, _ := NewInv()

	//msg, _ := NewMsg()

	//pubkey, _ := NewPubkey()

	//version, _ := NewVersion()
}
