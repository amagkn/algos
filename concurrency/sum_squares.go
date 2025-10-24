package concurrency

import (
	"sync"
	"sync/atomic"
)

// SumSquares считает сумму квадратов слайса многопоточно (сколько элементов, столько горутин)
func SumSquares(nums []int) int {
	var counter int64 = 0

	var wg sync.WaitGroup
	for _, n := range nums {
		wg.Add(1)

		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, int64(n*n))
		}()
	}

	wg.Wait()
	return int(counter)
}
