package hashmap_test

import (
	"fmt"

	"github.com/fahlke/golibs/hashmap"
)

func ExampleHashMap_Set() {
	hm := hashmap.HashMap{}
	hm.Set("foo", "bar")
}

func ExampleHashMap_Get() {
	hm := hashmap.HashMap{}
	hm.Set("foo", "bar")

	value, err := hm.Get("foo")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
	// output:
	// bar
}

func ExampleHashMap_Delete() {
	hm := hashmap.HashMap{}

	err := hm.Delete("foo")
	if err != nil {
		fmt.Println(err)
	}
	// output:
	// item not found
}
