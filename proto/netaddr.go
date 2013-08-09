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
	"bytes"
	"encoding/binary"
	"errors"
	"time"
)

type netaddr struct {

	// the Time. Protocol version 1 clients use 4 byte time while protocol version 2 clients use 8 byte time.
	Time int64

	// Stream number for this node
	Stream uint32

	// same service(s) listed in version
	Services uint64

	// IPv6 address. The original client only supports IPv4 and only reads the last 4 bytes to get the IPv4 address. However, the IPv4 address is written into the message as a 16 byte IPv4-mapped IPv6 address
	// (12 bytes 00 00 00 00 00 00 00 00 00 00 FF FF, followed by the 4 bytes of the IPv4 address).
	IP [16]byte

	// port number
	Port uint16
}

func NewNetaddr() *netaddr {

	return new(netaddr)
}

func NewNetaddrFrom(stream uint32, services uint64, ip []byte, port uint16) *netaddr {

	na := new(netaddr)

	na.Time = time.Now().Unix()
	na.Stream = stream
	na.Services = services
	copy(na.IP[:], ip)
	na.Port = port

	return na
}

func (na *netaddr) Serialize() ([]byte, error) {

	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, na.Time)
	binary.Write(&buf, binary.BigEndian, na.Stream)
	binary.Write(&buf, binary.BigEndian, na.Services)
	buf.Write(na.IP[:])
	binary.Write(&buf, binary.BigEndian, na.Port)

	return buf.Bytes(), nil
}

func (na *netaddr) Deserialize(packet []byte) error {

	if len(packet) < 38 { // time + stream + services + ip + port
		return errors.New("netaddr.Deserialize: packet is too short")
	}

	na.Time = int64(binary.BigEndian.Uint64(packet[:8]))
	na.Stream = binary.BigEndian.Uint32(packet[8:12])
	na.Services = binary.BigEndian.Uint64(packet[12:20])
	copy(na.IP[:], packet[20:36])
	na.Port = binary.BigEndian.Uint16(packet[36:38])

	return nil
}
