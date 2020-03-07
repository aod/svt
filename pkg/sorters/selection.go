package sorters

import (
	"sort"
	"sync"
)

func Selection(arr sort.Interface, update chan<- int, mutex *sync.Mutex) {
	n := arr.Len()

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			update <- j
			if arr.Less(j, minIdx) {
				minIdx = j
			}
		}

		mutex.Lock()
		arr.Swap(minIdx, i)
		mutex.Unlock()
		update <- i
	}
}
