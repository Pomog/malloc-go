package dublicates

import (
	"log"
	"sync"
)

func HasDuplicatesParallel(nums []int, numWorkers int) bool {
	if len(nums) == 0 || numWorkers == 0 {
		log.Fatal("0 error")
	}

	chunkSize := (int)(len(nums) / numWorkers)

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	foundDuplicate := false

	checkChunk := func(start, end int) {
		defer wg.Done()
		localSeen := make(map[int]struct{})
		for i := start; i < end; i++ {
			num := nums[i]

			if _, exists := localSeen[num]; exists {
				mu.Lock()
				foundDuplicate = true
				mu.Unlock()
				//TODO  I need to brake all go routines here

				return
			}
			localSeen[num] = struct{}{}
		}

	}

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go checkChunk(start, end)
	}

	wg.Wait()

	return foundDuplicate
}
