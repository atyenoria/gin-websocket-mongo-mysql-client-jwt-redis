package main

import "fmt"

type Android struct {
	Person Person
	Model  string
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Print(p.Name)
}

func main() {
	a := new(Android)
	a.Person.Name="sample"
	a.Person.Talk()
}
