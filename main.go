package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("func done") // execute code after the function finish.
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, up := lenAndUpper("jalapeno")
	fmt.Println(totalLenght, up)
}