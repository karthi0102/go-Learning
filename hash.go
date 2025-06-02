package main

import "fmt"


// func do(m1 map[int]int){
// 	m1[1]=3
// 	m1 = make(map[int]int)
// 	m1[1]=4
// 	fmt.Printf("%#v",m1)
// }


// func main(){
// 	m := map[int]int{	1: 2};
// 	fmt.Printf("%#v", m)
// 	do(m)
// 	fmt.Printf("%#v", m)

// }

func do(m1 *map[int]int) {
	(*m1)[1] = 3
	*m1 = make(map[int]int)
	(*m1)[1] = 11
	fmt.Printf("%#v", *m1)
}

func main() {
	m := map[int]int{4: 2}
	do(&m);
	fmt.Printf("%#v", m)
}