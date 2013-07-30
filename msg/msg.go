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
	Magic    [4]byte
	Command  [12]byte
	Length   [4]byte
	Checksum [4]byte
	Payload  []byte
}

func NewMessage() *Message {

	m := new(Message)
	binary.BigEndian.PutUint32(m.Magic[:], 0xe9beb4d9)
	return m
}

func NewMessageFromCommand(cmd string, payload []byte) (*Message, error) {

	m := new(Message)

	if len(cmd) >= 12 {
		return nil, errors.New("msg.NewMessage: Command is too long")
	}

	binary.BigEndian.PutUint32(m.Magic[:], 0xe9beb4d9)

	var i int = 0
	for ; i < len(cmd); i++ {
		m.Command[i] = cmd[i]
	}

	for ; i < len(m.Command); i++ {
		m.Command[i] = 0
	}

	copy(m.Payload, payload)

	binary.BigEndian.PutUint32(m.Length[:], uint32(len(m.Payload)))

	sha := sha512.New()
	sha.Write(m.Payload)
	copy(m.Checksum[:], sha.Sum(nil)[:4])

	return m, nil
}

func (m *Message) GetCommand() string {

	var buf bytes.Buffer

	for i := 0; i < len(m.Command); i++ {
		buf.WriteByte(m.Command[i])
		if m.Command[i] == 0 {
			break
		}
	}

	return buf.String()
}

func (m *Message) GetLength() uint32 {

	return binary.BigEndian.Uint32(m.Length[:])
}

func (m *Message) Serialize() []byte {

	var buf bytes.Buffer

	buf.Write(m.Magic[:])
	buf.Write(m.Command[:])
	buf.Write(m.Length[:])
	buf.Write(m.Checksum[:])
	buf.Write(m.Payload)

	return buf.Bytes()
}

func (m *Message) Deserialize(packet []byte) {

	copy(m.Magic[:], packet[:4])
	copy(m.Command[:], packet[4:16])
	copy(m.Length[:], packet[20:24])
	length := binary.BigEndian.Uint32(m.Length[:])
	copy(m.Checksum[:], packet[24:28])
	copy(m.Payload[:], packet[28:length])
}
