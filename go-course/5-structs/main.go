package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	p := Person{
		Name:   "Wilson",
		Age:    24,
		Active: true,
	}
	fmt.Println("name", p.Name)
	fmt.Println("age", p.Age)
	fmt.Println("Active", p.Active)
}
