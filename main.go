package main

import "fmt"

func main() {
	names := []string{"jalapeno", "hou27", "cow"} // Slice : Array that without the length.
	names = append(names, "magic") // append doesn't modify the slice.
	fmt.Println(names)
}