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

// HashMap represents a hash table implementation
type HashMap struct {
	mutex sync.RWMutex
	data  [buckets][]item
}

type item struct {
	key   string
	value interface{}
}

// Set adds a new key/value pair to the hashmap
func (h *HashMap) Set(key string, value interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	digest := getDigest(&key)

	// If key is already in the bucket, update and return
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

// Get returns the requested key/value pair
func (h *HashMap) Get(key string) (interface{}, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	digest := getDigest(&key)

	for i := range h.data[digest] {
		if h.data[digest][i].key == key {
			return h.data[digest][i].value, nil
		}
	}

	return nil, ErrNotFound
}

// Delete removes a key/value pair from the hashmap
func (h *HashMap) Delete(key string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	digest := getDigest(&key)

	for i := range h.data[digest] {
		if h.data[digest][i].key == key {
			h.data[digest] = append(h.data[digest][:i], h.data[digest][i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func getDigest(key *string) uint16 {
	return pearson.Sum16(*key)
}
