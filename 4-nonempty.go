package main

import "fmt"

func main() {
	initial := []string{"this", "", "is", "", "test"}
	res := nonempty(initial...)
	fmt.Println(res)
	fmt.Println(initial)
	res2 := nonempty2(initial)

	fmt.Println(res2)
	fmt.Println(initial)

	initial2 := []string{"this", "", "is", "", "test"}
	initial3 := initial2[:]
	initial3[0] = "test"
	fmt.Println(initial3)
	fmt.Println(initial2)
	// fmt.Println(res2)

}

func nonempty(strings ...string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
