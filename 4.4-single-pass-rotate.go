package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	rotate := 3
	for i := 0; i < rotate; i++ {
		tmp := a[i]
		tmp2 := 0
		for j := i; j < len(a); j = j + rotate {

			if j == 0 {
				tmp = a[(j+rotate)%len(a)]
				a[(j+rotate)%len(a)] = a[j]
			} else {
				tmp2 = a[(j+rotate)%len(a)]
				a[(j+rotate)%len(a)] = tmp
				tmp = tmp2
			}
		}
	}
	fmt.Println(a)
}
