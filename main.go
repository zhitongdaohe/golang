package main

import (
	"fmt"
	"time"
)

// 1,3,5,7
// 2,4,6,8

var index int = 1

var ch1, ch2 = make(chan int), make(chan int)

func main() {
	go work1()
	go work2()

	time.Sleep(time.Second * 10)

	// arr := []int{1, 1, 2, 3, 3, 4, 5}

	// removeDuplicateNums(arr)
}

func work1() {
	fmt.Println(index)
	index++
	ch1 <- index

	for {
		select {
		case val := <-ch2:
			fmt.Println(val)
			time.Sleep(time.Second * 1)
			index++
			ch1 <- index
		}
	}
}

func work2() {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
			time.Sleep(time.Second * 1)
			index++
			ch2 <- index
		}
	}
}

/*
给你一个 **升序排列** 的数组 `nums` ，
请你删除重复出现的元素，使每个元素 **只出现一次** ，
返回删除后数组的新长度。元素的 **相对顺序** 应该保持 **一致**
*/

// 1, 1, 2, 3, 3, 4, 5
func removeDuplicateNums(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	i, j := 1, 1

	for j < len(arr) {
		if arr[j] == arr[i-1] {
			j++
		} else {
			arr[i] = arr[j]
			i++
			j++
		}
	}

	fmt.Printf("i=%d, j=%d\n", i, j)

	arr = arr[0:i]
	fmt.Println(arr)

	return i + 1
}
