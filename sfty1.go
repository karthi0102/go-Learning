package main

import "fmt"


func main() {
	items := [][2]byte{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	a := [][] byte{}

	fmt.Println("======= Before =======")
	fmt.Println(a)
	fmt.Println(items);

	for _,item := range items {
		a = append(a, item[:])
	}

	fmt.Println("======= After =======")
	fmt.Println(items);
	fmt.Println(a)
}