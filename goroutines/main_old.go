package main

import "fmt"
import "time"
import "sync"

func main_old() {
	// write 2 time table
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			time.Sleep(150 * time.Millisecond)
			fmt.Printf("2 x %v = %v\n", i, i*2)
			wg.Done()

		}(i, wg)
	}
	wg.Wait()
	//time.Sleep(2 * time.Second)
}
