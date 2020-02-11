package hashmap

import (
	"errors"
	"sync"

	"github.com/fahlke/golibs/hash/pearson"
)

// maximum number of buckets possible with a Pearson-16 hash
const buckets = 1 << 16

var (
	// ErrNotFound means that the requested item does not exist
	ErrNotFound = errors.New("item not found")
)

type HashMap struct {
	mutex sync.RWMutex
	data  [buckets][]item
}

type item struct {
	key   string
	value interface{}
}

func (h *HashMap) Set(key string, value interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	digest := pearson.Sum16(key)

	// if key is already in the bucket, update and exit
	for i := range h.data[digest] {
		if h.data[digest][i].key == key {
			h.data[digest][i].value = value

			return
		}
	}

	// Add a new key/value pair to the bucket
	h.data[digest] = append(h.data[digest], item{
		key:   key,
		value: value,
	})
}

func (h *HashMap) Get(key string) (interface{}, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	digest := pearson.Sum16(key)

	x := h.data[digest]
	for i := range x {
		if h.data[digest][i].key == key {
			return h.data[digest][i], nil
		}
	}

	return nil, ErrNotFound
}
