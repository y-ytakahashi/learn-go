package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) fullName() string {
	return p.firstName + p.lastName
}

func NewPerson(fistName, lastName string) *Person {
	return &Person{
		firstName: fistName,
		lastName:  lastName,
	}

}

func main() {
	p := []*Person{
		NewPerson("hoge", "huga"),
		NewPerson("opai", "mohumohu"),
	}

	for index, v := range p {
		fmt.Println(index, ":", v.fullName())
	}
}
