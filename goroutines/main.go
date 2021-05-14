package main

import "fmt"
import "sync"

var tables =  map[int]map[int]int{}


func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan map[int]int)

	var consumer = func(wg *sync.WaitGroup, ch chan map[int]int) {
		fmt.Println(<-ch)
		wg.Done()
	}

	var producer = func(wg *sync.WaitGroup, ch chan map[int]int) {
		tableOf := 2
		table := map[int]int{}
		for i := 1; i < 10; i++ {
			table[i] = tableOf * i
		}
		ch<- table
		wg.Done()
	}
	wg.Add(2)
	go consumer(wg, ch)
	go producer(wg, ch)
	wg.Wait()	
			
		

}
	
