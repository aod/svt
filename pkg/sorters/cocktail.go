package sorters

import (
	"sort"
	"sync"
)

func Cocktail(arr sort.Interface, update chan<- Compare, mutex *sync.Mutex) {
	n := arr.Len()
	beginIdx := 0
	endIdx := n - 1

	for beginIdx <= endIdx {
		newBeginIdx := endIdx
		newEndIdx := beginIdx

		for i := beginIdx; i < endIdx; i++ {
			c := Compare{
				Indexes: [2]int{i + 1, i},
				Swapped: false,
			}

			if arr.Less(i+1, i) {
				mutex.Lock()
				arr.Swap(i+1, i)
				mutex.Unlock()

				c.Swapped = true
				newEndIdx = i
			}

			update <- c
		}

		endIdx = newEndIdx

		for i := endIdx - 1; i >= beginIdx; i-- {
			c := Compare{
				Indexes: [2]int{i + 1, i},
				Swapped: false,
			}

			if arr.Less(i+1, i) {
				mutex.Lock()
				arr.Swap(i+1, i)
				mutex.Unlock()

				c.Swapped = true
				newBeginIdx = i
			}

			update <- c
		}

		beginIdx = newBeginIdx
	}
}
