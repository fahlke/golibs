package pearson

import (
	"testing"
)

type test struct {
	name   string
	in     string
	digest uint32
}

//nolint:gochecknoglobals
var golden = []test{}

func TestSize(t *testing.T)      { t.Parallel() }
func TestBlockSize(t *testing.T) { t.Parallel() }
func TestReset(t *testing.T)     { t.Parallel() }
func TestSum(t *testing.T)       { t.Parallel() }
func TestSum32(t *testing.T)     { t.Parallel() }
func TestWrite(t *testing.T)     { t.Parallel() }
