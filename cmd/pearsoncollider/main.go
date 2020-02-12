package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fahlke/golibs/hash/pearson"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890äöüÄÖÜ") //nolint:gochecknoglobals

// Main starts a tool to show hash collisions for an arbitrary input string against a random generated string
// in my Pearson-16 hash implementation.
func main() {
	var str1 = "foo"

	if value, ok := os.LookupEnv("TEST_STRING"); ok {
		str1 = value
	}

	fmt.Printf("Trying to find a maximum of 10 hash collisions with Pearson-16...\n\n")

	for collisions := 0; collisions < 10; {
		str2 := randSeq() //nolint:gomnd
		if str1 == str2 {
			continue
		}

		if pearson.Sum16(str1) == pearson.Sum16(str2) {
			fmt.Printf("digest `0x%x` is the same for `%s` and `%s`\n", pearson.Sum16(str2), str1, str2)
			collisions++
		}
	}
}

func randSeq() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, rand.Intn(20)) //nolint:gomnd
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
