package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	coordinateWithContext()

	withDeadline()
}

func coordinateWithContext() {
	var wg sync.WaitGroup

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with context.Context]\n", num)
	max := int32(10)
	go addNum(&num, 3, max, cancel)
	<-ctx.Done()
}

func addNum(num *int32, val int, max int32, f func()) {
	fmt.Printf("num=%d, val=%d, max=%d\n", *num, val, max)
	f()
}

func withDeadline() {
	ctx := context.Background()

	ctx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(time.Second*10))

	go func(ctx context.Context) {
		i := 0
		for {
			fmt.Println(i)
			i++
			time.Sleep(time.Second)
		}
	}(ctx)

	<-ctx.Done()

	cancelFunc()
}
