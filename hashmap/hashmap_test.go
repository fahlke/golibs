package hashmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fahlke/golibs/hashmap"
)

func TestSet(t *testing.T) {
	t.Parallel()

	m := hashmap.HashMap{}

	err := m.Set("Hello", "World!")
	assert.NoError(t, err)
}