package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var foo string = "bar"
var (
	x1 bool      = false
	x2 string    = "info"
	x3 time.Time = time.Now()
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)

	var baz string = "qux"
	fmt.Println(foo, baz)
	fmt.Println(x1, x2, x3)

	var seed1 int64 = 1234456789

	rand.Seed(seed1) //重複隨機數  一個位置

	d1 := false
	l1 := "inter"
	s1 := time.Now()
	fmt.Println(d1, l1, s1)
}
