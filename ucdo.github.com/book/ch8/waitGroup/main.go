package main

import "sync"

func testWaitGroup() {
	str := `.//'.\`
	ch := make(chan int, len(str))
	var wg sync.WaitGroup
	for v := range str {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			// do something
			ch <- v
		}(v)
	}

	// wait all goroutine done
	go func() {
		wg.Wait()
		close(ch)
	}()
}
