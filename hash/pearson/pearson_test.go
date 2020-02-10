package pearson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name   string
	in     string
	digest uint64
}

//nolint:gochecknoglobals
var golden = []test{}

func TestSize(t *testing.T) {
	t.Parallel()

	var d = New()

	assert.Equal(t, 8, d.Size())
}

func TestBlockSize(t *testing.T) {
	t.Parallel()

	var d = New()

	assert.Equal(t, 1, d.BlockSize())
}

func TestReset(t *testing.T) { t.Parallel() }
func TestSum(t *testing.T)   { t.Parallel() }
func TestSum32(t *testing.T) { t.Parallel() }
func TestWrite(t *testing.T) { t.Parallel() }
