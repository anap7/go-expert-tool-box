package main

func main() {
	a := 5
	b := 2
	c := 3

	switch a {
	case 1: 
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default: 
		println("d")
	}
	
	if a > b || c > a {
		println(a)
	} else {
		println(b)
	}
}