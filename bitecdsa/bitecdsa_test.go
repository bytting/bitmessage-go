// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2011 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitecdsa

import (
	"bitmessage/bitelliptic"
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"testing"
)

func testKeyGeneration(t *testing.T, c *bitelliptic.BitCurve, tag string) {
	priv, err := GenerateKey(c, rand.Reader)
	if err != nil {
		t.Errorf("%s: error: %s", tag, err)
		return
	}
	if !c.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
		t.Errorf("%s: public key invalid: %s", tag, err)
	}
}

func TestKeyGeneration(t *testing.T) {
	testKeyGeneration(t, bitelliptic.S256(), "S256")
	if testing.Short() {
		return
	}
	testKeyGeneration(t, bitelliptic.S160(), "S160")
	testKeyGeneration(t, bitelliptic.S192(), "S192")
	testKeyGeneration(t, bitelliptic.S224(), "S224")
}

func testSignAndVerify(t *testing.T, c *bitelliptic.BitCurve, tag string) {
	priv, _ := GenerateKey(c, rand.Reader)

	hashed := []byte("testing")
	r, s, err := Sign(rand.Reader, priv, hashed)
	if err != nil {
		t.Errorf("%s: error signing: %s", tag, err)
		return
	}

	if !Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify failed", tag)
	}

	hashed[0] ^= 0xff
	if Verify(&priv.PublicKey, hashed, r, s) {
		t.Errorf("%s: Verify always works!", tag)
	}
}

func TestSignAndVerify(t *testing.T) {
	testSignAndVerify(t, bitelliptic.S256(), "S256")
	if testing.Short() {
		return
	}
	testSignAndVerify(t, bitelliptic.S160(), "S160")
	testSignAndVerify(t, bitelliptic.S192(), "S192")
	testSignAndVerify(t, bitelliptic.S224(), "S224")
}

func fromHex(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("bad hex")
	}
	return r
}

// These test vectors were generated with OpenSSL using vectorgen.rb
var testVectors = []struct {
	hash   string
	Qx, Qy string
	r, s   string
	ok     bool
}{
	{
		"rWO/YB4Ur9u5yXRkOC0QrRmutzuGReCslws/wgt3uaE=",
		"75180709339c672ffb8db4fe8ca27a0603b2c85a4460371f7bd69618000935e6",
		"54e67cde188fcd0e3ade6b94422834505e17ba44c98b945f80596ed4c5d57ff1",
		"acbd37fb3876e5580556855f51b158b7462069283c15e337b47562d9246b7cf9",
		"117ff1e940b28f3add88d138f7c4b67145c508019c4fa0bf5b2fdbf622db6226",
		true,
	},
	{
		"MhMQsu+hkicpORVDX+liavMDRvH2IBidH9Z1UO77GWk=",
		"2a81f04c0ed31850f7701c72179b4f0cf5f438705c821c2c340775f31589694d",
		"9c7e176c645de82d00e80a6980eda165b6840ad4fb66f305f0994f299dee83cc",
		"3f154d5d2214958197fa777c7aeb114c29ce3c9d1978e08413b72a56476317c6",
		"b79b8da65ec1f31607de2d5697c4d8e4460c34bcf0b595979ee8886a7a22490f",
		true,
	},
	{
		"mPVLf1SunM/ffOxrFaC+NV5dHbaHDeFJy0ViI3bSddw=",
		"b3b5e83a406478b26ff5a051286b9295c7bb11e350a75806c7e21fd067dd3caa",
		"404da2f031a5e1697d7e3550db992b3f1093b1cfb26a43ac1136d1e34175520c",
		"7334611111a2df05bcc8955ea12eda1187be693d59977aef8a537e83a7cf0228",
		"227a8054c147a5b89c88cd524a4b81da8f07457add995e35fa5ff02a96b01dbb",
		true,
	},
	{
		"2MNMltlFPVVmpqOx524b7l0W+aRcHtokKZSKBA2/8z8=",
		"51bc15c20131a92f3161be567dcde6675b866d11e6d64f9094f93ec91536136d",
		"69f6d1b74743e1857d02cd720e4378822e0598485ce65422d27b0090f7276c64",
		"2225f7e1dfb852989e8ab085afcc3a4941e10e73fdd31b06b4cdf315ee982980",
		"d1bc806b74a85232b3cd843595a201d6910e6a370faf639657b45e5e7dd06d68",
		true,
	},
	{
		"05oUbaIeMArcXpe2Q4U6s5UfcmyoluHoRdZb0H0LOP8=",
		"7653996b68c20e98e32f1f062bf66c4d906f4573fba5ba317d614f886535d4fe",
		"bf9e3f48523b2d49962c5cc29f67eb4ddf6de6afebda0d777b10a1bc94b58cee",
		"8f42e09fc02c7caa7471d1fb010b99ab6d0fdc173dcb546d8014eace25c6b4e0",
		"3c951d468ba26c59d3c3008310b46c5986b1e6fc21c030e5fd5824d677e76ea3",
		true,
	},
	{
		"XQKZz3vOw9t76d102XJty/1gIeMr4jxoV0vakfB64hg=",
		"54e298ad4bbe26e1935eacb057e621b0cf496a11ca2ee485d56567c631989978",
		"a96039cee81be6fc95c4d4c5f6f0c69f48ce4be56d0156ee694f7bdaef3ee0cd",
		"6ce6c3cc8ead129d0da35d378da07fa86330993fd3f166b3fccf8c9ea9067af0",
		"c481fd7360ceef79357f5beefcedc21fa6b8aa3fb8eea6c7a13b18d8b679256c",
		true,
	},
	{
		"UZ/P1J1K+64bLnNc+K5JSxZX5ZupRSaxMLI0JedXdRs=",
		"1b994a0fdecff7cf3a4392743727a3a74be726f4d7c02224d0fad910f1714ee2",
		"8b3b37234126c6764fbebbb45adb9be0e13a08b3e30659232b54bd73d0dd9508",
		"8ca7216aaa8ff2ccc101b6a15710fe273a41a04d5a43360c7d6a889c9f71fb78",
		"44a74dd5c2b39257308fbc36f83291a8748de26e3cf42d0533e4348a58d4a899",
		true,
	},
	{
		"U16DczVy8AqDWnFvJGVPfLM8UZL6WHMfkIjeiRGCRSY=",
		"77f6e0d2a20c161b8788a18dfd77662dfb60e3cd408b74591f48acb6dac2d7d4",
		"5557af02ad7c227e71639f00c08ec34ab94d2877c9089701b619f23789334f0d",
		"3b75c6d5124b2e71c8bea53acce1bf7587b00a68705571f609c33547f696f469",
		"b6024e1ff27e2775690b37f18e74c8e3845226557bb3b40be0206a6c64fd928b",
		true,
	},
	{
		"7vXxpKXjgoeuGYrSSZsemM8pLehnB22kR8He88iyGEM=",
		"0d6b65ac0699fab1c22f9fdfbb63b540d692aa6a0e6127e495af55d11d49d60e",
		"214b168e985cd03acc0c7e0c79b8cefea0f57fd482b042bbea898ab9dadb69fe",
		"86af531e21cb21bb26a460d48f0859865e13d421018607983827d18e6dac9115",
		"11de053699e0ab61c09840e7f2f2b994bd3a642246852e6d781303fb3cde0f21",
		true,
	},
	{
		"Ys16vX2HP0xLQ/qs/B0ibIG74OklkW7HGHFBt+Jrars=",
		"1dcb5989e0ea5dafe23c20f16fcbd5304f0111d1b2d9787bc8e27e4b309df58b",
		"6c7fd68e517bc301e028b9b4b857a0c0c17dcab5091dda7cccebe41b023bab58",
		"b91bfb1e04f142359e204c547d5d9d0103b3242f22481266c0dd2d0f550afce2",
		"f27b01154367ada77ef0864159f67e5d8a8b9c147c109cc7dbb58edf4d31fb3d",
		true,
	},
	{
		"WQ6cVAnMSU1EY/TqxPlEf6NiukEqwg6I9SQawOafdfA=",
		"434709b03bc6a4f85e41b36cb363fcffacfa6004cd88319c4f198dfb55817c9d",
		"a328f6f34135070008db8627a2bcc059d2d22f94ab997f324efd8ed7b46eeca9",
		"0ec1168e5492a30c6d7416728fbff65e09d8a5604cccd300ce78a03f3bb6a819",
		"e67a8c628a62c9d8c290ad80d720c301099eac6980d70d460b83496fd3897880",
		false,
	},
	{
		"HKl7FWfgW6Ak94ul+d0h2EBoWMrmQ+UvObs02G253fI=",
		"6fe145fd43ce35776a7a100572dce5a56ed99e0d016e9a3492001f07e3b94bb4",
		"51f409be7a06f45a5b50b968a074475261d2c295a6361f9fba7938d84d98a07f",
		"6bf5928119b50383d5a0827246c54332e5195a16f67c7d0d0395785e8ec6b37a",
		"9893394f8b8d72d83e170a6fc75f6cab3014023ffb9dba78132d07fff303aac9",
		false,
	},
	{
		"HMiKpuZ77GXybslRjcXVGQeZtsXoO+43RXYOfZNqaco=",
		"df869c02298562f61daf77b14ec61d9846bee5b3a0ad889ba7579bcea654d930",
		"a90586b5f7edac6bfd36082f1caa75b9ad6c5b14608c37292f291018de2459aa",
		"3a63a928dd7ae2a893375a5ff59ceeb514af4f50cc59403ed1ce678fbb740678",
		"72dbfd3ed248ec227308d84c1c7b51ae3d45e1f56ff9da6a4469f5a75353a20e",
		false,
	},
	{
		"I9yE0WBaIWJi1b30cXLnCU6f9ZHpZYnqu/iexCSOAis=",
		"6b12278a10906d7a7bdea6f4acd27b8aab7aff741291e2c0936a7e9195a139ea",
		"2eb90f005a796cb9e4ac916eb9ada784b0a6e7f00b400a4d1112163d18ce642c",
		"78825efc06632ce1e8c18740e7c890e43874d409f4990ebbf574157724166b24",
		"8fbc10a9cffd6f51da9a1b3dca808575d7b5848ba5744fab85bfed9722617008",
		false,
	},
	{
		"zdwUrQPYBZBYiRTFHkpT7m5Y2g9dmIOEUOeg1UT8Q3E=",
		"14860c07172ee7bbfedefe36a88732a3b6b0c881b9267b06c1ddad3df0e74cc0",
		"c02612cdf33db35ed3a396c461f3fab0c0a02fda8377aed94927315fbf37f3fb",
		"51abd6dc6a5128640cfaece311444c81758317806ebcd5c7cd5e4be87939aa48",
		"ca9aeadc7ab478de83fa4ed7e7462340c7cb85292b077db0b553ebd6afa8e27d",
		false,
	},
	{
		"DRyVx9Nh8TzP0Dvf6mWg49PzVDu1lFzgvplIeOyCryI=",
		"e11380da9773efa620793e250e3aae5f35968e752f869b268b4a85a840051012",
		"8a22802d8bd396c442d82f8143c85c2a2f0f38fa9f4c7abd020b24c4e60d7592",
		"dc499e27300a12a0c64a682a921c7b493f025734f14404cb87b6b16afa73fcba",
		"334b81cae281a0d255ec32174a2c71960a45bf7dbc83aa063da476ea08e47ad0",
		false,
	},
	{
		"A4gPHKhpDW6TfM+iqTs9ERCMEH1D4JWDp/ikIXx+8DE=",
		"5d72e49d777228163a02ab7a44627ccfeed65a539db2d10b9b870244a101c56b",
		"15e4e579a75bc535c4597292b4c4e7f465681b8b5e1cdfb7e2832a3340984628",
		"a035ad65506baa9dd031fdeea7fbf8c2561cf8b078c488d1d2e5a953bc1dd9fc",
		"5c0f7f2b250d3ec79b200f2f8ee2e1d07b55a8d8f1ba7cf1f97634a6463c6676",
		false,
	},
	{
		"O9VfPlCHlELFnfyk3XyWe61sqrWFdIjhX/9YKuHcpc0=",
		"160a5c3e40c612755aa92c7bafa2ec0d30beaf3c5e8502357598ef1ff5a30d39",
		"ec0157612f8bc12fd3401a623d74cf995474652b02d9157bf6a516e28e581b2d",
		"753af6d5ef6bc9b619a65c935abc10ca9669fa36fc0f04eb0d770cfa9851bb53",
		"cde3845f37da3b48a5d73bc4da9721e0420ffdc90fcff3d4e6327741ce489d67",
		false,
	},
	{
		"mslv4d71bz33jMEVkhYCYZambPkF1AGl1XeKzAqUHdI=",
		"177fac6032f5e7943887c649f2d1d644e46fe9a4855deff3dbee7501658eebb9",
		"b936a5174434aa416190f3b934d33517560a1e9986ca2c6fdd30988425090e62",
		"e85f7bc9d1f497e70b24a950fa1247fa45b4abc125445de96499f511298d7f1b",
		"b86ad799a14a580074c3be6f04947d71fcf0fe65e325dc601a9ea4a5f05e722f",
		false,
	},
	{
		"jHOF5JG6NDz3qvXxc7BMFvQnuaDFu6w2WNOzrVS3c5s=",
		"1cab54e51f4f0a29bdb6469e5991db6808ec9feee87c390f23850f45dbb46cde",
		"4eb3a2ba5d5c3f6f5236e46f01d4c1f6e7a6bb1d28925d5a4ee4b4c39198b481",
		"efebf536c797cf0544f604590c9308785ea01b9ef9383f037dde205dcc19abed",
		"f7c8463a43b722b11bfff0043e4c4b4da3b4a7a3d532b1a0b2e7492869ecf877",
		false,
	},
}

func TestVectors(t *testing.T) {
	for i, test := range testVectors {
		pub := PublicKey{
			BitCurve: bitelliptic.S256(),
			X:        fromHex(test.Qx),
			Y:        fromHex(test.Qy),
		}
		hashed, _ := base64.StdEncoding.DecodeString(test.hash)
		r := fromHex(test.r)
		s := fromHex(test.s)
		if Verify(&pub, hashed, r, s) != test.ok {
			t.Errorf("%d: bad result", i)
		}
		if testing.Short() {
			break
		}
	}
}

func BenchmarkVerify(b *testing.B) {
	b.StopTimer()
	data := testVectors[0]
	pub := &PublicKey{
		BitCurve: bitelliptic.S256(),
		X:        fromHex(data.Qx),
		Y:        fromHex(data.Qy),
	}
	hashed, _ := base64.StdEncoding.DecodeString(data.hash)
	r := fromHex(data.r)
	s := fromHex(data.s)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Verify(pub, hashed, r, s)
	}
}

func BenchmarkSign(b *testing.B) {
	b.StopTimer()
	priv, _ := GenerateKey(bitelliptic.S256(), rand.Reader)
	hashed := []byte("testing")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Sign(rand.Reader, priv, hashed)
	}
}
