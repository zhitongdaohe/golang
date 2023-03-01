package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestClosure(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("not paramter, i:", i)
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("iniput paramter, i:", i)
		}(i)
	}
	wg.Wait()
}
