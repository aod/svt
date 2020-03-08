package sorters

import (
	"sort"
	"sync"
)

func Selection(arr sort.Interface, update chan<- Compare, mutex *sync.Mutex) {
	n := arr.Len()

	for i := 0; i < n-1; i++ {
		c := Compare{}
		c.Indexes[0] = i
		minIdx := i

		for j := i + 1; j < n; j++ {
			c.Indexes[1] = j
			update <- c
			if arr.Less(j, minIdx) {
				minIdx = j
			}
		}

		mutex.Lock()
		arr.Swap(minIdx, i)
		mutex.Unlock()

		update <- Compare{
			Indexes: [2]int{minIdx, i},
			Swapped: true,
		}
	}
}
