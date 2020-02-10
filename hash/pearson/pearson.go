package pearson

import "hash"

type digest uint32

func New() hash.Hash32 {
	panic("implement me")
}

func (d *digest) Size() int {
	panic("implement me")
}

func (d *digest) BlockSize() int {
	panic("implement me")
}

func (d *digest) Reset() {
	panic("implement me")
}

func (d *digest) Sum(b []byte) []byte {
	panic("implement me")
}

func (d *digest) Sum32() uint32 {
	panic("implement me")
}

func (d *digest) Write(p []byte) (n int, err error) {
	panic("implement me")
}
