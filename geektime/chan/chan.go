package main

import (
	"fmt"
	"sync"
)

func main() {
	coordinateWithChan()
}

func coordinateWithChan() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)
	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})
	<-sign
	<-sign
}

func addNum(num *int32, val int, max int32, f func()) {
	f()

	fmt.Printf("num=%d, val=%d, max=%d\n", *num, val, max)
}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)
	go addNum(&num, 3, max, wg.Done)
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
}
