package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"jelly", "sushi"}
	jalapeno := person{name: "jalapeno", age: 22, favFood: favFood} // == person{"jalapeno", 22, favFood}
	fmt.Println(jalapeno)
}