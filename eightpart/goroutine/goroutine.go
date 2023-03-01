package goroutine

import "fmt"

func Closure() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:", i)
		}()
	}
}
