package main

import (
	"fmt"
	"time"
)

func task(name string, channel chan int) {
	for {
		value := <-channel
		fmt.Println("Executing task", value, "by", name)
	}
}

func main() {
	channel := make(chan int)

	go task("go time 1", channel)
	go task("go time 2", channel)
	go task("go time 3", channel)
	go task("go time 4", channel)
	go task("go time 5", channel)

	for i := 0; i < 100; i++ {
		channel <- i
	}
	time.Sleep(time.Second)
}
