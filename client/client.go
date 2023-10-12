package client

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			Http2()
		}()

	}
	wg.Wait()
}
