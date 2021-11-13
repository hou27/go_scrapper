package main

import (
	"fmt"

	"github.com/hou27/learngo/myDict"
)

func main() {
	dictionary := myDict.Diiiii{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}