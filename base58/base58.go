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

package base58

import (
	"bytes"
	"errors"
	"math/big"
	"strings"
)

const (
	alphabet58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func Encode(b []byte) (string, error) {

	if len(b) < 1 {
		return "", errors.New("base58.Encode: Byte slice is too short")
	}
	zero := big.NewInt(0)
	val := big.NewInt(0)
	val.SetBytes(b)

	var buffer bytes.Buffer

	if val.Cmp(zero) == 0 {
		buffer.WriteByte(alphabet58[0])
		return buffer.String(), nil
	}

	n := val
	r := big.NewInt(0)
	base := big.NewInt(58)

	for n.Cmp(zero) > 0 {
		r.Mod(n, base)
		n.Div(n, base) // FIXME: Use DivMod
		buffer.WriteByte(alphabet58[r.Uint64()])
	}

	length := len(buffer.Bytes())
	for i := 0; i < length/2; i++ {
		buffer.Bytes()[i], buffer.Bytes()[length-1-i] = buffer.Bytes()[length-1-i], buffer.Bytes()[i]
	}

	return buffer.String(), nil
}

func Decode(encoded string) ([]byte, error) {

	bn := big.NewInt(0)
	base := big.NewInt(58)
	tmp := big.NewInt(0)

	for i := 0; i < len(encoded); i++ {
		pos := strings.IndexRune(alphabet58, rune(encoded[i]))
		if pos == -1 {
			return nil, errors.New("base58.Decode: Character not present in base58")
		}
		tmp.SetUint64(uint64(pos))

		bn.Mul(bn, base)
		bn.Add(bn, tmp)
	}

	return bn.Bytes(), nil
}
