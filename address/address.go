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
	"strings"

	"bitmessage-go/base58"
	"bitmessage-go/bitecdsa"
	"bitmessage-go/bitelliptic"
	"bitmessage-go/ripemd160"
	"bitmessage-go/varint"
)

type address struct {
	Identifier    string
	SigningKey    *bitecdsa.PrivateKey
	EncryptionKey *bitecdsa.PrivateKey
}

func New(addressVersion, stream uint64, eighteenByteRipe bool) (*address, error) {

	var err error
	addr := new(address)
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

		var keyMerge bytes.Buffer
		keyMerge.Write(addr.SigningKey.PublicKey.X.Bytes())
		keyMerge.Write(addr.SigningKey.PublicKey.Y.Bytes())
		keyMerge.Write(addr.EncryptionKey.PublicKey.X.Bytes())
		keyMerge.Write(addr.EncryptionKey.PublicKey.Y.Bytes())

		sha := sha512.New()
		sha.Write(keyMerge.Bytes())

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

	var bmAddr bytes.Buffer
	bmAddr.Write(varint.Encode(addressVersion))
	bmAddr.Write(varint.Encode(stream))
	bmAddr.Write(ripe)

	sha1, sha2 := sha512.New(), sha512.New()
	sha1.Write(bmAddr.Bytes())
	sha2.Write(sha1.Sum(nil))
	checksum := sha2.Sum(nil)[:4]
	bmAddr.Write(checksum)

	encoded, err := base58.Encode(bmAddr.Bytes())
	if err != nil {
		return nil, err
	}
	addr.Identifier = "BM-" + encoded

	return addr, nil
}

func (addr *address) Version() (uint64, error) {

	b, err := base58.Decode(addr.Identifier[3:])
	if err != nil {
		return 0, err
	}
	v, _ := varint.Decode(b)

	return v, nil
}

func (addr *address) Stream() (uint64, error) {

	b, err := base58.Decode(addr.Identifier[3:])
	if err != nil {
		return 0, err
	}
	_, nb := varint.Decode(b)
	s, _ := varint.Decode(b[nb:])

	return s, nil
}

func ValidateIdentifier(identifier string) bool {

	if !strings.HasPrefix(identifier, "BM-") {
		return false
	}

	if len(identifier) < 25 { // prefix + ripe + checksum
		return false
	}

	return true
}

func ValidateChecksum(address string) (bool, error) {

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
