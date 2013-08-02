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
	"crypto/sha512"
	"encoding/binary"
	"errors"
)

type message struct {
	Magic    uint32
	Command  string
	Length   uint32
	Checksum []byte
	Payload  []byte
}

func NewMessage(cmd string, payload []byte) (*message, error) {

	m := new(message)

	if len(cmd) >= 12 {
		return nil, errors.New("msg.NewMessage: Command is too long")
	}

	m.Magic = 0xe9beb4d9 // FIXME check endianess
	m.Command = cmd
	m.Payload = payload
	m.Length = uint32(len(m.Payload))

	sha := sha512.New()
	sha.Write(m.Payload)
	m.Checksum = sha.Sum(nil)[:4]

	return m, nil
}

func (m *message) Serialize() ([]byte, error) {

	if len(m.Command) == 0 || len(m.Checksum) != 4 || len(m.Payload) == 0 {
		return nil, errors.New("msg.Serialize: Message is incomplete")
	}

	if m.Magic != 0xe9beb4d9 {
		return nil, errors.New("msg.Serialize: Magic number is invalid")
	}

	if int(m.Length) != len(m.Payload) {
		return nil, errors.New("msg.Serialize: Message length is invalid")
	}

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, m.Magic)
	for i := 0; i < 12; i++ {
		if i < len(m.Command) {
			buf.WriteByte(m.Command[i])
		} else {
			buf.WriteByte(byte(0))
		}
	}
	binary.Write(buf, binary.BigEndian, m.Length)
	buf.Write(m.Checksum)
	buf.Write(m.Payload)

	return buf.Bytes(), nil
}

func Deserialize(packet []byte) (*message, error) {

	if len(packet) < 25 {
		return nil, errors.New("msg.Deserialize: Packet length is too small")
	}

	m := new(message)

	m.Magic = binary.BigEndian.Uint32(packet[:4])
	if m.Magic != 0xe9beb4d9 {
		return nil, errors.New("msg.Deserialize: Magic number is invalid")
	}

	var cmd bytes.Buffer
	for i := 0; i < 12; i++ {
		if packet[4+i] == 0 {
			break
		}
		cmd.WriteByte(packet[4+i])
	}
	m.Command = cmd.String()

	m.Length = binary.BigEndian.Uint32(packet[16:20])
	m.Checksum = packet[20:24]
	m.Payload = packet[24:]

	if int(m.Length) != len(m.Payload) {
		return nil, errors.New("msg.Deserialize: Message length is invalid")
	}

	return m, nil
}
