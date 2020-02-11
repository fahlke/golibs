package hashmap

import (
	"hash"
	"sync"
)

// maximum number of buckets possible with a Pearson-16 hash
const buckets = 1 << 16

type HashMap struct {
	mutex sync.RWMutex
	data  [buckets][]item
	hash  hash.Hash
}

type item struct {
	key   string
	value interface{}
}

func (h *HashMap) Set(key string, value interface{}) error {
	return nil
}
