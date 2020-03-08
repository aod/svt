package sorters

import (
	"sort"
	"sync"
)

func Bubble(arr sort.Interface, update chan<- Compare, mutex *sync.Mutex) {
	n := arr.Len()

	for n > 1 {
		newN := 0
		for i := 1; i <= n-1; i++ {
			c := Compare{
				Indexes: [2]int{i, i - 1},
				Swapped: false,
			}

			if arr.Less(i, i-1) {
				mutex.Lock()
				arr.Swap(i, i-1)
				mutex.Unlock()

				c.Swapped = true
				newN = i
			}

			update <- c
		}

		n = newN
	}
}
