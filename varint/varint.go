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

package varint

import (
	"encoding/binary"
)

func Encode(value uint64) []byte {

	buf := make([]byte, 16)

	if value < 253 {
		buf[0] = byte(value)
		return buf[:1]
	} else if value >= 253 && value < 65536 {
		buf[0] = 253
		binary.BigEndian.PutUint16(buf[1:], uint16(value))
		return buf[:3]
	} else if value >= 65536 && value < 4294967296 {
		buf[0] = 254
		binary.BigEndian.PutUint32(buf[1:], uint32(value))
		return buf[:5]
	} else {
		buf[0] = 255
		binary.BigEndian.PutUint64(buf[1:], uint64(value))
		return buf[:9]
	}
}

func Decode(buffer []byte) (uint64, int) {

	firstByte := buffer[0]

	if firstByte < 253 {
		return uint64(firstByte), 1
	} else if firstByte == 253 {
		return uint64(binary.BigEndian.Uint16(buffer[1:])), 3
	} else if firstByte == 254 {
		return uint64(binary.BigEndian.Uint32(buffer[1:])), 5
	} else {
		return uint64(binary.BigEndian.Uint64(buffer[1:])), 9
	}
}
