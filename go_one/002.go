package main

import "fmt"

var t float32 = 42

func main() {
	var i int = 40
	var j float64 = 42
	k := 99.
	fmt.Println("hello ,you")
	fmt.Println(i)
	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)
	fmt.Printf("%v , %T\n", t, t)

}
