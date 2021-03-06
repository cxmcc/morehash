// Package bsdsum implements the BSD checksum algorithm
// BSD checksum is also known as UNIX sum
package bsdsum

import "github.com/cxmcc/unixsums"

// The size of a BSD checksum value
const Size = 2

type digest struct {
	sum uint16
}

// New returns a new hash.Hash computing the BSD checksum value
func New() unixsums.Hash16 {
	return &digest{0}
}

func (d *digest) Reset() {
	d.sum = 0
}

func (d *digest) Size() int {
	return Size
}

func (d *digest) BlockSize() int {
	return 1
}

func (d *digest) Write(data []byte) (int, error) {
	for _, b := range data {
		d.sum = (d.sum >> 1) + (d.sum << 15) + uint16(b)
	}
	return len(data), nil
}

func (d *digest) Sum16() uint16 {
	return d.sum
}

func (d *digest) Sum(in []byte) []byte {
	s := d.Sum16()
	return append(in, byte(s>>8), byte(s))
}

// Bsdsum returns the BSD checksum of the given byte array
func Bsdsum(data []byte) uint16 {
	d := New()
	d.Write(data)
	return d.Sum16()
}
