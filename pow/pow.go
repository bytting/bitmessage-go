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

package pow

import (
	"bitmessage/varint"
	"crypto/sha512"
	"encoding/binary"
	"runtime"
)

func scan(offset_start, offset_end, target uint64, payload_hash []byte, out chan<- uint64, done chan<- bool, shutdown *bool) {

	var trials uint64 = 18446744073709551615
	var nonce uint64 = offset_start
	h1, h2 := sha512.New(), sha512.New()

	for trials > target {
		if *shutdown {
			done <- true
			return
		}
		nonce++
		if nonce > offset_end {
			return
		}
		b := varint.Encode(nonce)
		b = append(b, payload_hash...)
		h1.Write(b)
		h2.Write(h1.Sum(nil))
		trials = binary.BigEndian.Uint64(h2.Sum(nil)[:8])
		h1.Reset()
		h2.Reset()
	}
	out <- nonce
	done <- true
}

func Nonce(payload []byte) uint64 {

	ncpu := runtime.NumCPU() - 1
	if ncpu < 1 {
		ncpu = 1
	}
	runtime.GOMAXPROCS(ncpu)

	sha := sha512.New()
	sha.Write(payload)
	payload_hash := sha.Sum(nil)
	var target uint64 = 18446744073709551615 / uint64((len(payload)+14000+8)*320)

	var nprocs int = 1000
	var i, slice uint64 = 0, 18446744073709551615 / uint64(nprocs)

	recv := make(chan uint64, nprocs)
	done := make(chan bool, nprocs)
	shutdown := false

	for ; i < uint64(nprocs); i++ {
		go scan(i*slice, i*slice+slice, target, payload_hash, recv, done, &shutdown)
	}

	nonce := <-recv

	shutdown = true
	for i = 0; i < uint64(nprocs); i++ {
		<-done
	}

	return nonce
}

func Validate(payload []byte) bool {

	if len(payload) < 2 {
		return false
	}

	var offset int
	var nonce, trials_test uint64
	var hash_test, initial_payload []byte

	nonce, offset = varint.Decode(payload)
	initial_payload = payload[offset:]

	sha := sha512.New()
	sha.Write(initial_payload)
	initial_hash := sha.Sum(nil)

	var target uint64 = 18446744073709551615 / uint64((len(payload)+14000+8)*320)

	hash_test = varint.Encode(nonce)
	hash_test = append(hash_test, initial_hash...)
	sha1, sha2 := sha512.New(), sha512.New()
	sha1.Write(hash_test)
	sha2.Write(sha1.Sum(nil))

	trials_test = binary.BigEndian.Uint64(sha2.Sum(nil)[:8])

	return trials_test <= target
}
