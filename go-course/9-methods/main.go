package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	if p.Age < 18 {
		return fmt.Sprintf("logg jovem %s, đâyyyy", p.Name)
	} else if p.Age > 50 {
		return fmt.Sprintf("logg senhor %s, đâyyyy", p.Name)
	}

	return fmt.Sprintf("logg %s, đâyyyy", p.Name)
}

func (p *Person) privateMethod() int {
	return p.Age * 4
}

func main() {
	person := Person{Name: "junior", Age: 17}
	fmt.Println(person.String())

	person = Person{Name: "Wilson", Age: 27}
	fmt.Println(person.String())

	person = Person{Name: "mr thanh kiu", Age: 60}
	fmt.Println(person.String())

}
