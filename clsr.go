package main

func do() (int, func() int){
	a,b := 1, 2
	return a, func() int {
		a,b = b, a+b
		return b
	}
}

func main(){
	a, b := do()
	println(a,b(),b(),b()) 
}