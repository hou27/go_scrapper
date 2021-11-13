package main

import "fmt"

func canIDrink(age int) bool {
	// 1

	// switch koreanAge := age + 2; koreanAge {
	// case 10:
	// 	return false
	// case 19:
	// 	return true
	// }
	// return false

	// 2

	koreanAge := age + 2;
	switch true {
	case koreanAge < 19:
		return false
	case koreanAge >= 19:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}