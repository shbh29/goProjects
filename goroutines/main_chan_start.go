package main

import "fmt"
import "time"

func main_chan_start() {
	ch1 := make(chan int)
	//	ch2 := make(chan int)

	go func(ch chan int) {
		
		fmt.Println(<-ch)
	}(ch1)

	go func(ch chan int) {
		ch <- 42
	}(ch1)

	time.Sleep(2 * time.Second)

}
