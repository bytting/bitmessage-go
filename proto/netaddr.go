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

import (
	"encoding/binary"
	"time"
)

type netaddr struct {

	// the Time. Protocol version 1 clients use 4 byte time while protocol version 2 clients use 8 byte time.
	time [8]byte

	// Stream number for this node
	stream [4]byte

	// same service(s) listed in version
	services [8]byte

	// IPv6 address. The original client only supports IPv4 and only reads the last 4 bytes to get the IPv4 address. However, the IPv4 address is written into the message as a 16 byte IPv4-mapped IPv6 address
	// (12 bytes 00 00 00 00 00 00 00 00 00 00 FF FF, followed by the 4 bytes of the IPv4 address).
	ip [16]byte

	// port number
	port [2]byte
}

func NewNetaddr(stream uint32, services uint64, ip []byte, port uint16) *netaddr {

	a := new(netaddr)
	binary.BigEndian.PutUint64(a.time[:], uint64(time.Now().Unix()))
	binary.BigEndian.PutUint32(a.stream[:], stream)
	binary.BigEndian.PutUint64(a.services[:], services)
	copy(a.ip[:], ip)
	binary.BigEndian.PutUint16(a.port[:], port)
	return a
}

func (na *netaddr) Time() int64 {

	return int64(binary.BigEndian.Uint64(na.time[:]))
}

func (na *netaddr) Stream() uint32 {

	return binary.BigEndian.Uint32(na.stream[:])
}

func (na *netaddr) Services() uint64 {

	return binary.BigEndian.Uint64(na.services[:])
}

func (na *netaddr) IP() []byte {

	return na.ip[:]
}

func (na *netaddr) Port() uint16 {

	return binary.BigEndian.Uint16(na.port[:])
}
