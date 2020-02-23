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

	bt := New()

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

func TestBinaryTree_Get(t *testing.T) {
	t.Parallel()

	bt := New()

	t.Run("empty", func(t *testing.T) {
		_, err := bt.Get("empty")
		assert.EqualError(t, err, "the binary tree is empty")
	})

	for _, item := range insertTests {
		bt.Set(item.Key, item.Value)
	}

	t.Run("root", func(t *testing.T) {
		value, err := bt.Get("Alpha")
		assert.NoError(t, err)
		assert.Equal(t, "álfa (άλφα)", value)
	})

	t.Run("left", func(t *testing.T) {
		value, err := bt.Get("Delta")
		assert.NoError(t, err)
		assert.Equal(t, "délta (δέλτα)", value)
	})

	t.Run("right", func(t *testing.T) {
		value, err := bt.Get("Zeta")
		assert.NoError(t, err)
		assert.Equal(t, "zíta (ζήτα)", value)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := bt.Get("non existing key")
		assert.EqualError(t, err, "not found")
	})
}

func TestBinaryTree_Height(t *testing.T) {
	t.Parallel()

	bt := New()

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

func TestBinaryTree_Iterate(t *testing.T) {
	t.Parallel()

	bt := New()

	t.Run("insert", func(t *testing.T) {
		for _, item := range insertTests {
			bt.Set(item.Key, item.Value)
		}

		var n int

		for item := range bt.Iterate() {
			//nolint:gomnd
			switch n {
			case 0:
				assert.Equal(t, "Alpha", item.Key)
				assert.Equal(t, "álfa (άλφα)", item.Value)
			case 1:
				assert.Equal(t, "Beta", item.Key)
				assert.Equal(t, "víta (βήτα)", item.Value)
			case 2:
				assert.Equal(t, "Delta", item.Key)
				assert.Equal(t, "délta (δέλτα)", item.Value)
			case 3:
				assert.Equal(t, "Epsilon", item.Key)
				assert.Equal(t, "épsilon (έψιλον)", item.Value)
			case 5:
				assert.Equal(t, "Zeta", item.Key)
				assert.Equal(t, "zíta (ζήτα)", item.Value)
			case 6:
				assert.Equal(t, "Gamma", item.Key)
				assert.Equal(t, "gám(m)a (γάμ(μ)α)", item.Value)
			}

			n++
		}
	})
}

func TestBinaryTree_String(t *testing.T) {
	t.Parallel()

	bt := New()

	for _, item := range insertTests {
		bt.Set(item.Key, item.Value)
	}

	expected := "" +
		"root: Alpha\n" +
		"  right: Beta\n" +
		"    right: Gamma\n" +
		"      left: Delta\n" +
		"        right: Epsilon\n" +
		"      right: Zeta\n"

	assert.Equal(t, expected, bt.String())
}
