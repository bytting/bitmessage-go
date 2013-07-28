/*
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (BM-2DAS9BAs92wLKajVy9DS1LFcDiey5dxp5c)

package wif

import (
	"bitmessage/base58"
	"bitmessage/bitecdsa"
	"bitmessage/bitelliptic"
	"bytes"
	"crypto/sha256"
	"math/big"
)

func Encode(keys *bitecdsa.PrivateKey) (wif string) {

	var extended bytes.Buffer
	extended.WriteByte(byte(0x80))
	extended.Write(keys.D.Bytes())
	sha1, sha2 := sha256.New(), sha256.New()
	sha1.Write(extended.Bytes())
	sha2.Write(sha1.Sum(nil))
	checksum := sha2.Sum(nil)[:4]
	extended.Write(checksum)
	return base58.Encode(extended.Bytes())
}

func Decode(wif string) *bitecdsa.PrivateKey {

	if len(wif) < 6 {
		panic("WIF is too short in Decode")
	}

	extended := base58.Decode(wif)
	decoded := extended[1 : len(extended)-4]
	keys := new(bitecdsa.PrivateKey)
	keys.D = new(big.Int).SetBytes(decoded)
	keys.PublicKey.BitCurve = bitelliptic.S256()
	for keys.PublicKey.X == nil {
		keys.PublicKey.X, keys.PublicKey.Y = keys.PublicKey.BitCurve.ScalarBaseMult(decoded)
	}
	return keys
}

func Validate(wif string) bool {

	if len(wif) < 5 {
		panic("WIF is too short in Validate")
	}

	extended := base58.Decode(wif)
	raw := extended[1 : len(extended)-4]
	cs1 := extended[len(extended)-4:]
	sha1, sha2 := sha256.New(), sha256.New()
	sha1.Write(raw)
	sha2.Write(sha1.Sum(nil))
	cs2 := sha2.Sum(nil)[:4]
	return bytes.Compare(cs1, cs2) == 0
}
