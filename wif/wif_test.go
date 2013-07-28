package wif

import (
	"bitmessage/bitecdsa"
	"bitmessage/bitelliptic"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func TestKeys(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	keys1, err := bitecdsa.GenerateKey(bitelliptic.S256(), rand.Reader)
	if err != nil {
		t.Error("Failed to generate private key\n")
	}

	wif := Encode(keys1)

	keys2 := Decode(wif)

	if bytes.Compare(keys1.D.Bytes(), keys2.D.Bytes()) != 0 {
		t.Error("Private keys are different. Expected %x, got %x\n", keys1.D.Bytes(), keys2.D.Bytes())
	}

	if bytes.Compare(keys1.PublicKey.X.Bytes(), keys2.PublicKey.X.Bytes()) != 0 {
		t.Error("Public point X are different. Expected %x, got %x\n", keys1.PublicKey.X.Bytes(), keys2.PublicKey.X.Bytes())
	}

	if bytes.Compare(keys1.PublicKey.Y.Bytes(), keys2.PublicKey.Y.Bytes()) != 0 {
		t.Error("Public point Y are different. Expected %x, got %x\n", keys1.PublicKey.Y.Bytes(), keys2.PublicKey.Y.Bytes())
	}

	if !keys2.PublicKey.BitCurve.IsOnCurve(keys2.PublicKey.X, keys2.PublicKey.Y) {
		t.Error("Public point is not on curve\n")
	}

	// 0437a3191fe90d9b483324c28ecd019479e708cfcff96800131c113ec30a0646ee95c31b4c5656b1e7122f071ae4471a97511f372179147277ea2a2087147f9486

	keys3 := Decode("5HtKNfWZH4QQZPUGRadud7wfyPGEKLhQJfnYPGvpiivgwfrHfpX")
	privHex := hex.EncodeToString(keys3.D.Bytes())
	if privHex != "092715c60df8c561c832ab3c804be0a0f90b108072133df7d1e348e2570be801" {
		t.Error("Private key (keys3) is wrong. Expected 092715c60df8c561c832ab3c804be0a0f90b108072133df7d1e348e2570be801, got %s\n", privHex)
	}

	//pub := elliptic.Marshal(keys3.PublicKey.Curve, keys3.PublicKey.X, keys3.PublicKey.Y)
	var pub bytes.Buffer
	pub.WriteByte(byte(0x04))
	pub.Write(keys3.PublicKey.X.Bytes())
	pub.Write(keys3.PublicKey.Y.Bytes())
	pubHex := hex.EncodeToString(pub.Bytes())
	if pubHex != "0437a3191fe90d9b483324c28ecd019479e708cfcff96800131c113ec30a0646ee95c31b4c5656b1e7122f071ae4471a97511f372179147277ea2a2087147f9486" {
		t.Error("Public key (keys3) is wrong. Expected 0437a3191fe90d9b483324c28ecd019479e708cfcff96800131c113ec30a0646ee95c31b4c5656b1e7122f071ae4471a97511f372179147277ea2a2087147f9486, got ", pubHex)
	}
}
