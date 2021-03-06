package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap_Set(t *testing.T) {
	t.Parallel()

	t.Run("check hashmap size", func(t *testing.T) {
		t.Parallel()

		m := HashMap{} //nolint:staticcheck

		assert.EqualValues(t, 1<<16, len(m.data)) //nolint:gomnd
	})

	t.Run("add one", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		m.Set("foo", "bar")
		assert.Equal(t, 1, len(m.data[0xad5c]))
		assert.Equal(t, "foo", m.data[0xad5c][0].key)
		assert.Equal(t, "bar", m.data[0xad5c][0].value)
	})

	t.Run("add one and update it", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		m.Set("foo", "bar")
		assert.Equal(t, 1, len(m.data[0xad5c]))
		assert.Equal(t, "foo", m.data[0xad5c][0].key)
		assert.Equal(t, "bar", m.data[0xad5c][0].value)

		m.Set("foo", "overwritten")
		assert.Equal(t, 1, len(m.data[0xad5c]))
		assert.Equal(t, "foo", m.data[0xad5c][0].key)
		assert.Equal(t, "overwritten", m.data[0xad5c][0].value)
	})

	t.Run("collision", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		m.Set("🦄🌈", "rainbow unicorn")
		m.Set("QBFZUHJwURIf", "collision")
		assert.Equal(t, 2, len(m.data[0x45ce]))
		assert.Equal(t, "🦄🌈", m.data[0x45ce][0].key)
		assert.Equal(t, "rainbow unicorn", m.data[0x45ce][0].value)
		assert.Equal(t, "QBFZUHJwURIf", m.data[0x45ce][1].key)
		assert.Equal(t, "collision", m.data[0x45ce][1].value)
	})
}

func TestHashMap_Get(t *testing.T) {
	t.Parallel()

	t.Run("existing item", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		m.Set("foo", "bar")

		value, err := m.Get("foo")
		assert.NoError(t, err)
		assert.Equal(t, "bar", value)
	})

	t.Run("missing item", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		m.Set("foo", "bar")

		_, err := m.Get("non-existing")
		assert.EqualError(t, err, "item not found")
	})
}

func TestHashMap_Delete(t *testing.T) {
	t.Parallel()

	t.Run("existing item", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}
		m.Set("foo", "bar")

		err := m.Delete("foo")
		assert.Equal(t, 0, len(m.data[0xad5c]))
		assert.NoError(t, err)
	})

	t.Run("empty map", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}

		err := m.Delete("non existing key")
		assert.Equal(t, 0, len(m.data[0x32]))
		assert.EqualError(t, err, "item not found")
	})

	t.Run("missing item", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}
		m.Set("foo", "bar")

		err := m.Delete("non existing key")
		assert.Equal(t, 0, len(m.data[0x32]))
		assert.EqualError(t, err, "item not found")
	})

	t.Run("colliding item", func(t *testing.T) {
		t.Parallel()

		m := HashMap{}
		m.Set("foo", "bar")
		m.Set("x8ÜutM", "collision")

		err := m.Delete("x8ÜutM")
		assert.Equal(t, 1, len(m.data[0xad5c]))
		assert.NoError(t, err)
	})
}
