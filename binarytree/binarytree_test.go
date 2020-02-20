package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals
var golden = []Item{
	{"Alpha", "álfa (άλφα)"},
	{"Beta", "víta (βήτα)"},
	{"Gamma", "gám(m)a (γάμ(μ)α)"},
	{"Delta", "délta (δέλτα)"},
	{"Epsilon", "épsilon (έψιλον)"},
	{"Zeta", "zíta (ζήτα)"},
}

func TestBinaryTree_Set(t *testing.T) {
	t.Parallel()

	bt := BinaryTree{}

	for _, item := range golden {
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
}
