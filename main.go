package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 19 { // variable expression
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}