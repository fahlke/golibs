package pearson

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name   string
	in     string
	digest uint16
}

//nolint:gochecknoglobals
var golden = []test{
	{"empty string", "", 0x0},
	{"foo", "foo", 0xad5c},
	{"bar", "bar", 0x9be8},
	{"baz", "baz", 0x488b},
	{"Rainbow unicorn", "🦄🌈", 0x45ce},
}

func TestSize(t *testing.T) {
	t.Parallel()

	var d = New()

	assert.Equal(t, 2, d.Size())
}

func TestBlockSize(t *testing.T) {
	t.Parallel()

	var d = New()

	assert.Equal(t, 1, d.BlockSize())
}

func TestReset(t *testing.T) {
	t.Parallel()

	var digest = New()

	digest.Write([]byte("dummy")) //nolint:errcheck
	digest.Reset()

	assert.Equal(t, uint16(0x0), binary.BigEndian.Uint16(digest.Sum(nil))) //nolint:gomnd
}

func TestSum(t *testing.T) { t.Parallel() }

func TestWrite(t *testing.T) {
	t.Parallel()

	for _, tt := range golden {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var digest = New()

			digest.Write([]byte(tt.in)) //nolint:errcheck
			assert.Equal(t, tt.digest, binary.BigEndian.Uint16(digest.Sum(nil)))
		})
	}
}
