package pearson_test

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/fahlke/golibs/hash/pearson"
)

// This example shows the basic usage of the package: Create a Pearson hash.Hash,
// write a value and receive the hash as hex value.
func ExampleNew_hexValue() {
	digest := pearson.New()
	digest.Write([]byte("Hello World!"))

	fmt.Printf("%x\n", digest.Sum(nil))
	// Output: e0a2
}

// This example shows the usage of the package: Create a Pearson hash.Hash,
// write a value and receive the hash as an unsigned 16bit integer.
func ExampleNew_intValue() {
	digest := pearson.New()
	digest.Write([]byte("Hello World!"))

	var ret uint16
	buf := bytes.NewBuffer(digest.Sum(nil))
	binary.Read(buf, binary.BigEndian, &ret)

	fmt.Printf("%d\n", ret)
	// Output: 57506
}
