package main

import "fmt"

func main() {
	var a[]int;
 	a = append(a, 1)

	fmt.Println("a:", a)


	arr := [...]int{1, 2, 3}

	b := arr[0:1:1]

	fmt.Println("b:", b)

	c := b[0:1]

	fmt.Println("c:", c)
}