package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 5
	var d = 5
	var e = a + b - c*d/5

	fmt.Println(e)

	// Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	// Unary Operator
	var j = 1
	j++ // 1 + j = 2
	j++ // 2 + j = 3
	j-- // 2 - j = 2
	fmt.Println(j)

}
