package main

import "fmt"

//Person is some type
type Person struct {
	firstName string
	lastName  string
}

func changeName(p *Person) {
	fmt.Print(p)
	p.firstName = "Bob"
}

func main() {

	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}
	fmt.Print(&person)
	changeName(&person)

	fmt.Println(person)

	var x []string
	fmt.Print(x)
}

func test(p []string) []string {
	fmt.Println(p)
	return p
}
