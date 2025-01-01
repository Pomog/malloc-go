package main

import (
	"Pomog/malloc-go/dublicates"
	"fmt"
	"time"
)

func generateHugeList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i + 1
	}
	return list
}

func main() {
	nums := generateHugeList(10_000_000)
	nums[9_999_999] = 1
	//nums[4_599_999] = 1

	start := time.Now()
	fmt.Println("Using single thread:", dublicates.HasDuplicates(nums), "Time:", time.Since(start)) // 2.168 sec on my comp

	start = time.Now()
	fmt.Println("Using gorutines:", dublicates.HasDuplicatesParallel(nums, 4), "Time:", time.Since(start)) // 0.816 sec on my comp
}
