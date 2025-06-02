package main

import "fmt"

type album struct {
	title  string
	copies int
}

func addCopies(a album){
	a.copies += 1
	fmt.Printf("Album: %s, Copies: %d\n", a.title, a.copies)
}

func addCopiesRef(a *album){
	a.copies += 1
	fmt.Printf("Album: %s, Copies: %d\n", a.title, a.copies)
}

func main(){
	a := album{title: "The Dark Side of the Moon", copies: 5}
	addCopies(a)
	fmt.Printf("After addCopies, Copies: %d\n", a.copies)
	addCopiesRef(&a)
	fmt.Printf("After addCopiesRef, Copies: %d\n", a.copies)
}