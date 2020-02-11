package hashmap

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"sync"

	"github.com/fahlke/golibs/hash/pearson"
)

// maximum number of buckets possible with a Pearson-16 hash
const buckets = 1 << 16

var (
	ErrDigestFailed = errors.New("failed to get digest")
)

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
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Create a new instance if not already initialized
	if h.hash == nil {
		h.hash = pearson.New()
	}

	// Generate hash value for the key
	// It never returns an error
	h.hash.Write([]byte(key)) //nolint:errcheck
	defer h.hash.Reset()

	var digest uint16

	buf := bytes.NewBuffer(h.hash.Sum(nil))
	if err := binary.Read(buf, binary.BigEndian, &digest); err != nil {
		return fmt.Errorf("%s (%s)", ErrDigestFailed, err)
	}

	// Add the new k/v pair to the hash table
	h.data[digest] = append(h.data[digest], item{
		key:   key,
		value: value,
	})

	return nil
}
