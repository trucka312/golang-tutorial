package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for {
		fmt.Println("Executing task", name)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("go time 1")
	go task("go time 2")
	go task("go time 3")
	go task("go time 4")
	go task("go time 5")

	task("main")
}
