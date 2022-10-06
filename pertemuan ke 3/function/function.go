package main

import "fmt"

func main() {
	greet("Airell", 23)
	greet1("Airell", "Jalan Sudriman")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d year ol\n", name, age)
}

func greet1(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in ", address)
}
