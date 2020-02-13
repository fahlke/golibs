package util_test

import (
	"fmt"

	"github.com/fahlke/golibs/util"
)

func ExampleMax() {
	max := util.Max(10, 31, 4)
	fmt.Println(max)
	// Output:
	// 31
}

func ExampleMin() {
	min := util.Min(10, 31, 4)
	fmt.Println(min)
	// Output:
	// 4
}
