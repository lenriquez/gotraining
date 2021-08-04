package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i) + ") " + sep + arg + "\n"
	}
	fmt.Println(s)
}
