package main

import "fmt"

func main() {
	var a []int
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 2)
	b := append(a, 3)
	c := append(a,
		4)
	// что окажется в a,b, с?
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
