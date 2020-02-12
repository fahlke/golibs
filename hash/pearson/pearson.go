package pearson

import (
	"errors"
	"hash"
)

const (
	size      = 2
	blockSize = 1
)

var (
	ErrDigestFailed = errors.New("failed to get digest")
)

type digest uint16

// 256-byte lookup table containing a permutation of the values 0 through 255.
// https://www.random.org/sequences/?min=0&max=255&col=16&format=plain&rnd=new
//nolint:gochecknoglobals
var lookupTable = []uint8{165, 199, 36, 25, 48, 88, 201, 164, 173, 226, 211, 225, 103, 66, 208, 62,
	252, 169, 251, 130, 162, 116, 133, 212, 12, 239, 218, 58, 113, 135, 9, 17,
	151, 215, 221, 27, 45, 21, 46, 138, 235, 118, 111, 245, 214, 178, 224, 213,
	246, 193, 69, 126, 167, 216, 244, 158, 41, 247, 198, 210, 137, 160, 86, 4,
	189, 7, 197, 176, 35, 84, 129, 49, 231, 128, 78, 217, 23, 51, 29, 153,
	180, 24, 19, 11, 161, 229, 237, 186, 33, 97, 219, 71, 77, 108, 52, 102,
	233, 174, 70, 53, 14, 125, 99, 144, 123, 166, 195, 209, 223, 110, 106, 112,
	172, 85, 6, 94, 131, 100, 15, 124, 243, 18, 156, 38, 73, 32, 175, 20,
	114, 127, 91, 204, 242, 119, 68, 141, 206, 157, 168, 96, 57, 43, 76, 250,
	150, 182, 98, 122, 238, 56, 145, 241, 81, 60, 143, 171, 115, 163, 44, 196,
	75, 253, 187, 40, 61, 177, 8, 109, 28, 202, 31, 117, 134, 92, 188, 90,
	54, 227, 191, 222, 26, 203, 146, 120, 184, 64, 30, 101, 170, 2, 154, 107,
	59, 65, 10, 22, 1, 236, 230, 39, 87, 228, 248, 205, 95, 79, 132, 254,
	0, 136, 207, 232, 152, 220, 147, 183, 47, 149, 67, 139, 192, 55, 121, 255,
	104, 200, 249, 159, 240, 234, 148, 181, 82, 105, 142, 190, 3, 89, 42, 74,
	72, 5, 13, 16, 37, 185, 34, 93, 155, 140, 179, 83, 80, 50, 63, 194}

// New returns a new 16-bit Pearson hash.Hash. Its Sum method will lay the value out in big-endian byte order.
func New() hash.Hash {
	d := digest(0)
	d.Reset()

	return &d
}

func (d *digest) Size() int {
	return size
}

func (d *digest) BlockSize() int {
	return blockSize
}

func (d *digest) Reset() {
	*d = digest(0)
}

func (d *digest) Sum(b []byte) []byte {
	s := uint16(*d)

	return append(b, byte(s>>8), byte(s)) //nolint:gomnd
}

func (d *digest) Write(p []byte) (n int, err error) {
	var h, hh uint8

	for i := range p {
		switch i {
		case 0:
			h = lookupTable[p[0]]
			hh = lookupTable[(p[0]+1)%255]
		default:
			h = lookupTable[h^p[i]]
			hh = lookupTable[hh^p[i]]
		}
	}

	*d = digest((uint16(h) << 8) | uint16(hh))

	return len(p), nil
}

func Sum16(in string) uint16 {
	d := digest(0)

	d.Write([]byte(in)) //nolint:errcheck

	return uint16(d)
}
