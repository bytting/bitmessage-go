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
	"errors"

	"bitmessage-go/varint"
)

type addr struct {

	// Number of address entries (max: 1000)
	//Count uint64

	// Address of other nodes on the network.
	AddrList []*netaddr
}

func NewAddr() *addr {

	return new(addr)
}

func (a *addr) Clear() {

	a.AddrList = nil
	//a.Count = 0
}

func (a *addr) Add(na *netaddr) {

	a.AddrList = append(a.AddrList, na)
	//a.Count++
}

func (a *addr) Serialize() ([]byte, error) {

	var buf bytes.Buffer

	buf.Write(varint.Encode(uint64(len(a.AddrList))))
	for i := range a.AddrList {
		b, _ := a.AddrList[i].Serialize()
		buf.Write(b)
	}

	return buf.Bytes(), nil
}

func (a *addr) Deserialize(packet []byte) error {

	if len(packet) < 8 {
		return errors.New("addr.Deserialize: packet is too short for count extraction")
	}

	cnt, nb := varint.Decode(packet[:8])

	if uint64(len(packet)) < uint64(nb)+(cnt*38) { // sizeof(netaddr) == 38
		return errors.New("addr.Deserialize: packet is too short for netaddr extraction")
	}

	for i := uint64(0); i < cnt; i++ {
		off := uint64(nb) + (i * 38)
		na := NewNetaddr()
		na.Deserialize(packet[off : off+38])
		a.Add(na)
	}

	return nil
}
