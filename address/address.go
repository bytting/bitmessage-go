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

package address

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"errors"

	"bitmessage-go/base58"
	"bitmessage-go/bitecdsa"
	"bitmessage-go/bitelliptic"
	"bitmessage-go/ripemd160"
	"bitmessage-go/varint"
)

type Address struct {
	Identifier                string
	SigningKey, EncryptionKey *bitecdsa.PrivateKey
}

func New(addressVersion, stream uint64, eighteenByteRipe bool) (*Address, error) {

	var err error
	addr := new(Address)
	addr.SigningKey, err = bitecdsa.GenerateKey(bitelliptic.S256(), rand.Reader)
	if err != nil {
		return nil, errors.New("address.New: Error generating ecdsa signing keys")
	}

	var ripe []byte

	for {
		addr.EncryptionKey, err = bitecdsa.GenerateKey(bitelliptic.S256(), rand.Reader)
		if err != nil {
			return nil, errors.New("address.New: Error generating ecdsa encryption keys")
		}

		var keyMerge []byte
		keyMerge = append(keyMerge, addr.SigningKey.PublicKey.X.Bytes()...)
		keyMerge = append(keyMerge, addr.SigningKey.PublicKey.Y.Bytes()...)
		keyMerge = append(keyMerge, addr.EncryptionKey.PublicKey.X.Bytes()...)
		keyMerge = append(keyMerge, addr.EncryptionKey.PublicKey.Y.Bytes()...)

		sha := sha512.New()
		sha.Write(keyMerge)

		ripemd := ripemd160.New()
		ripemd.Write(sha.Sum(nil))
		ripe = ripemd.Sum(nil)

		if eighteenByteRipe {
			if ripe[0] == 0x00 && ripe[1] == 0x00 {
				ripe = ripe[2:]
				break
			}
		} else {
			if ripe[0] == 0x00 {
				ripe = ripe[1:]
				break
			}
		}
	}

	bmAddr := varint.Encode(addressVersion)
	bmAddr = append(bmAddr, varint.Encode(stream)...)
	bmAddr = append(bmAddr, ripe...)

	sha1, sha2 := sha512.New(), sha512.New()
	sha1.Write(bmAddr)
	sha2.Write(sha1.Sum(nil))
	checksum := sha2.Sum(nil)[:4]
	bmAddr = append(bmAddr, checksum...)

	encoded, err := base58.Encode(bmAddr)
	if err != nil {
		return nil, err
	}
	addr.Identifier = "BM-" + encoded
	return addr, nil
}

func Validate(address string) (bool, error) {

	b, err := base58.Decode(address[3:])
	if err != nil {
		return false, err
	}
	raw := b[:len(b)-4]
	cs1 := b[len(b)-4:]

	sha1, sha2 := sha512.New(), sha512.New()
	sha1.Write(raw)
	sha2.Write(sha1.Sum(nil))
	cs2 := sha2.Sum(nil)[:4]

	return bytes.Compare(cs1, cs2) == 0, nil
}

func GetStream(address string) (uint64, error) {

	valid, err := Validate(address)
	if err != nil {
		return 0, err
	}

	if !valid {
		return 0, errors.New("address.GetStream: Invalid address checksum")
	}

	b, err := base58.Decode(address[3:])
	if err != nil {
		return 0, err
	}
	_, nb := varint.Decode(b)
	s, _ := varint.Decode(b[nb:])
	return s, nil
}
