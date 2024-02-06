package main

import "fmt"

func main() {
	people := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	people["junior"] = 10

	for name, age := range people {
		fmt.Printf(" %s hẻeeee: %d\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("loggg: %d\n", people["wilson"])
}
