package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	n := 0

	for i := 0; i < len(c1); i++ {
		for b := c1[i] ^ c2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	fmt.Println(n)
}
