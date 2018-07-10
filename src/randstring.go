package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		// bytes[i] = byte(randomInt(30, 300))
		bytes[i] = byte(rand.Intn(1000))
	}
	return string(bytes)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan(randomString(10)) // print 10 chars
}
