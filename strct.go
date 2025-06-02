package main

import "fmt"
var s1 struct {
	a int
	b string
}

var s2 struct {
	a int
	b string
}


type p1 struct {
	x int `json:"x_cord"`
	y int `json:"y_cord"`
}

type p2 struct {
	x int `json:"x_cord"`
	y int `json:"y_cord"`
}

type p3 struct {
	x int `json:"x_cord"`
	y int `json:"y_cordinate"`
}

type p4 struct {
	y int `json:"y_cord"`
	x int `json:"x_cord"`
}

type p5 struct {
	x int `json:"x_cord"`
	y int
}

type p6 struct {
	x int `json:"x_cord"`
	y int `json:"y_cord"`
	z int `json:"z_cord"`
}


func main(){
	s1 = struct{a int; b string}{a: 1, b: "hello"}
	s2 = s1
	fmt.Printf("s1: %v, s2: %v %v\n", s1, s2, s1 == s2)

	point1 := p1{x: 10, y: 20}

	fmt.Println(point1)

	point2 := p2(point1)
	fmt.Println(point2)

	point3 := p3(point1)
	fmt.Println(point3)

}


 
