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

type msg struct {
	Nonce        uint64
	Time         int64
	StreamNumber uint64
	Encrypted    []byte
}

func NewMsg() (*msg, error) {
	return nil, nil
}

func (v *msg) Serialize() ([]byte, error) {
	return nil, nil
}

func (v *msg) Deserialize(packet []byte) error {
	return nil
}
