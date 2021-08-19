package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var t = flag.Int("n", 256, "omit training newline")

func main() {
	flag.Parse()

	var r []byte
	fmt.Println(*t)

	if *t == 256 {
		// The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation;
		//or a field selector of an addressable struct operand; or an array indexing operation of an addressable array.
		// This is why this does not work
		// r := sha256.Sum256([]byte("x"))[:]
		bytes := sha256.Sum256([]byte("x"))
		r = bytes[:]
	} else if *t == 384 {
		bytes := sha512.Sum384([]byte("x"))
		r = bytes[:]
	} else if *t == 512 {
		bytes := sha512.Sum512([]byte("x"))
		r = bytes[:]
	}
	fmt.Println(r)
}

// https://stackoverflow.com/questions/48067303/function-should-return-sha256-sha384-sha512-result-as-byte-slice
// go run diffsha.go -n 384
