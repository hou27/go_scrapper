package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLenght, _ := lenAndUpper("jalapeno")
	fmt.Println(totalLenght)
	repeatMe("asgas", "dsagdgdgsa", "jalale", "houh")
}