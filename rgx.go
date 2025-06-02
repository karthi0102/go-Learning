package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main(){
	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)
	ch := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println("upper",ch)

	fmt.Println("Matches:", mm)
	fmt.Println("Indices:", id)

	for _, d := range id {
		fmt.Printf("Match: %s\n", te[d[0]:d[1]], )
	}
}