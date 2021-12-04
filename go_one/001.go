package main

import "fmt"

func main() {

	fmt.Println("Welcome to my quiz game")

	// string int uint float64 bool

	name := "Tom"
	age := 21
	fmt.Printf("Hello %s ,how are you doing?  ,you are %d?\n", name, age)

	fmt.Println("Enter you name!")
	var name1 string
	fmt.Scan(&name1)
	fmt.Printf("Hello, %s,wellcome to game!", name1)

	fmt.Println("Enter your age!")
	var age1 uint
	fmt.Scan(&age1)
	fmt.Printf("Age: %d\n", age1)

	if age1 >= 18 {
		fmt.Println("Ya you can play")
	} else {
		fmt.Println("Ooops!! You can't play")
		return //跳出整個函數
	}

	score := 0
	num_quiz := 2

	fmt.Println("Continued; What is better  RTX 3080 or RTX 3090?")
	var mm1 string
	var mm2 string
	fmt.Scan(&mm1, &mm2)
	fmt.Printf("your choice is %s  %s\n", mm1, mm2) //為了跳  space

	if mm1+" "+mm2 == "RTX 3090" || mm1+" "+mm2 == "rtx 3090" {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("INCorrect")
	}
	fmt.Printf("HOw many cores does the Ryzen 9 3900x?")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct")
		score++

	} else {
		fmt.Println("Incorrect")
	}
	fmt.Printf("Your scored is %d out of %d\n", score, num_quiz)
	percent := (float64(score) / float64(num_quiz)) * 100
	fmt.Printf("your score: %f%%", percent)
}

//true && true  -> true   only
//false || false-> false   only
//!false-> true
