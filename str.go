package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	if (len(os.Args) < 3) {
		fmt.Println("Usage: go run main.go <name>")
		os.Exit(-1);
	}


	fmt.Printf("Hello, %s!\n", os.Args[0])
	old, new := os.Args[1], os.Args[2]
	fmt.Printf("Replacing '%s' with '%s'\n", old, new)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), old)
		t := strings.Join(s, new)
		fmt.Println(t)
	}
	
}