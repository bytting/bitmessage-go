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

package main

import (
	"bitmessage-go/address"
	"bitmessage-go/base58"
	"bitmessage-go/varint"
	"bytes"
	"fmt"
)

func main() {

	var buf bytes.Buffer
	buf.WriteString("This is a string")

	encoded := base58.Encode(buf.Bytes())
	decoded := base58.Decode(encoded)

	var val uint64 = 12345
	v := varint.Encode(val)
	v2, nb := varint.Decode(v)

	fmt.Printf("%s, %d %d\n", decoded, nb, v2)

	addr := address.New(3, 1, false)
	fmt.Printf("%s\n", addr.Identifier)

	if !address.Validate(addr.Identifier) {
		panic("Panic!")
	}
}
