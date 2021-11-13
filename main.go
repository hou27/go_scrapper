package main

import "fmt"

func main() {
	jalapeno := map[string/* typeof key */]string/* typeof value */{"name": "hou27", "age": "22"}
	for key, value := range jalapeno {
		fmt.Println(key, value)
	}
}