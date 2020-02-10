package pearson

import "hash"

const (
	size = 8
)

type digest uint64

func New() hash.Hash64 {
	d := digest(0)
	d.Reset()

	return &d
}

func (d *digest) Size() int {
	return size
}

func (d *digest) BlockSize() int {
	panic("implement me")
}

func (d *digest) Reset() {
	*d = digest(0)
}

func (d *digest) Sum(b []byte) []byte {
	panic("implement me")
}

func (d *digest) Sum64() uint64 {
	panic("implement me")
}

func (d *digest) Write(p []byte) (n int, err error) {
	panic("implement me")
}
