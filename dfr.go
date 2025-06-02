package main

import "fmt"


func do() int{
	var a int
	defer func(){
		fmt.Print("defer")
		a=11
	}()
	fmt.Print("do")
	a=9
	return a
}

func main(){
	a := 1
	defer fmt.Printf("%d", a)

	a = 2
	fmt.Printf("%d",a)

	a=do()
   fmt.Printf("%d",a)

}