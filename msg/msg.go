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

package msg

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"errors"
)

type Message struct {
	Magic    uint32
	Command  string
	Length   uint32
	Checksum []byte
	Payload  []byte
}

func NewMessage() *Message {

	m := new(Message)
	m.Magic = 0xe9beb4d9 // FIXME check endianess
	return m
}

func NewMessageFromCommand(cmd string, payload []byte) (*Message, error) {

	m := new(Message)

	if len(cmd) >= 12 {
		return nil, errors.New("msg.NewMessage: Command is too long")
	}

	m.Magic = 0xe9beb4d9 // FIXME check endianess
	m.Command = cmd
	copy(m.Payload, payload)
	m.Length = uint32(len(m.Payload))

	sha := sha512.New()
	sha.Write(m.Payload)
	copy(m.Checksum, sha.Sum(nil)[:4])

	return m, nil
}

func (m *Message) Serialize() []byte {

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, m.Magic)
	buf.Write([]byte(m.Command))
	binary.Write(buf, binary.BigEndian, m.Length)
	buf.Write(m.Checksum)
	buf.Write(m.Payload)

	return buf.Bytes()
}

func (m *Message) Deserialize(packet []byte) {

	m.Magic = binary.BigEndian.Uint32(packet[:4])
	m.Command = string(packet[4:16])
	m.Length = binary.BigEndian.Uint32(packet[16:20])
	m.Checksum = packet[20:24]
	m.Payload = packet[24:]
}
