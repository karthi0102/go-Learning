package main

import "fmt"

type user struct {
	name  string
	age   int
}

func main(){
	users:= []user{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	alice := &users[0]

	fmt.Printf("Users: %#p\n", users)

	fmt.Println("Users:", users)
	fmt.Println("Alice:", alice.name, "Age:", alice.age)

	amy := user{"amy", 10}
	users = append(users, amy)
	fmt.Printf("Users: %#p\n", users)

	alice.age ++;

	fmt.Printf("Users: %v\n", users)

}