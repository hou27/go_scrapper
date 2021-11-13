package main

import (
	"fmt"

	"github.com/hou27/learngo/myDict"
)

func main() {
	dictionary := myDict.Diiiii{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	mean, _ := dictionary.Search(word)
	fmt.Println("found -", word, "\ndefinition:", mean)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}