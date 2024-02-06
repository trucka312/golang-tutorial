package main

import "fmt"

func main() {
	var x int 
	x = 10
	fmt.Printf("=)) %d\n", x)

	a := 20
	fmt.Printf("=)))) %d\n", a)
	a = 30
	fmt.Printf("-)) %d\n", a)

	// string
	name := "wilson"
	fmt.Println("...", name)

	// float64
	pi := 3.14
	fmt.Printf(".... %0.6f\n", pi)

	// bool
	verdade := true
	mentira := false

	fmt.Println(",..", verdade == mentira)
}
