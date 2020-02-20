package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals
var insertTests = []Item{
	{"Alpha", "álfa (άλφα)"},
	{"Beta", "víta (βήτα)"},
	{"Gamma", "gám(m)a (γάμ(μ)α)"},
	{"Delta", "délta (δέλτα)"},
	{"Epsilon", "épsilon (έψιλον)"},
	{"Zeta", "zíta (ζήτα)"},
}

//nolint:gochecknoglobals
var updateTests = []Item{
	{"Alpha", "álfa (άλφα) (updated)"},
	{"Delta", "délta (δέλτα) (updated)"},
}

func TestBinaryTree_Set(t *testing.T) {
	t.Parallel()

	bt := BinaryTree{}

	t.Run("insert", func(t *testing.T) {
		for _, item := range insertTests {
			bt.Set(item.Key, item.Value)
		}

		assert.Equal(t, "Alpha", bt.root.key)
		assert.Equal(t, "álfa (άλφα)", bt.root.value)
		assert.Equal(t, "Beta", bt.root.right.key)
		assert.Equal(t, "víta (βήτα)", bt.root.right.value)
		assert.Equal(t, "Gamma", bt.root.right.right.key)
		assert.Equal(t, "gám(m)a (γάμ(μ)α)", bt.root.right.right.value)
		assert.Equal(t, "Delta", bt.root.right.right.left.key)
		assert.Equal(t, "délta (δέλτα)", bt.root.right.right.left.value)
		assert.Equal(t, "Epsilon", bt.root.right.right.left.right.key)
		assert.Equal(t, "épsilon (έψιλον)", bt.root.right.right.left.right.value)
		assert.Equal(t, "Zeta", bt.root.right.right.right.key)
		assert.Equal(t, "zíta (ζήτα)", bt.root.right.right.right.value)
	})

	t.Run("update", func(t *testing.T) {
		for _, item := range updateTests {
			bt.Set(item.Key, item.Value)
		}

		assert.Equal(t, "Alpha", bt.root.key)
		assert.Equal(t, "álfa (άλφα) (updated)", bt.root.value)
		assert.Equal(t, "Delta", bt.root.right.right.left.key)
		assert.Equal(t, "délta (δέλτα) (updated)", bt.root.right.right.left.value)
	})
}

func TestBinaryTree_Height(t *testing.T) {
	t.Parallel()

	bt := BinaryTree{}

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, 0, bt.Height())
	})

	t.Run("one", func(t *testing.T) {
		bt.Set(insertTests[0].Key, insertTests[0].Value)

		assert.Equal(t, 1, bt.Height())
	})

	t.Run("five", func(t *testing.T) {
		for _, item := range insertTests {
			bt.Set(item.Key, item.Value)
		}

		assert.Equal(t, 5, bt.Height())
	})

	t.Run("five after update", func(t *testing.T) {
		for _, item := range updateTests {
			bt.Set(item.Key, item.Value)
		}

		assert.Equal(t, 5, bt.Height())
	})
}
