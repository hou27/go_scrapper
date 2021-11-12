package main

import (
	"fmt"

	"github.com/hou27/learngo/example"
)

func main() {
	example.PrintTwo()
	name := "hou27" // declaration + assignment ( go guess the type for you. )
	name = "jalapeno"
	fmt.Println(name) // if you want to export something, start with uppercase.
}