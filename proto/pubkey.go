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

package proto

import ()

type pubkey3 struct {
	Nonce               uint64
	Time                uint32 // FIXME uint64
	AddressVersion      uint64
	StreamNumber        uint64
	Behavior            uint32
	PublicSigningKey    []byte
	PublicEncryptionKey []byte
	NonceTrialsPerByte  uint64
	ExtraBytes          uint64
	SignatureLength     uint64
	Signature           []byte
}

func NewPubkey3() (*pubkey3, error) {
	return nil, nil
}

func (v *pubkey3) Serialize() ([]byte, error) {
	return nil, nil
}

func (v *pubkey3) Deserialize(packet []byte) error {
	return nil
}
