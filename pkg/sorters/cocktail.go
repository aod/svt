package sorters

import (
	"sort"
	"sync"
)

func Cocktail(arr sort.Interface, update chan<- int, mutex *sync.Mutex) {
	n := arr.Len()
	beginIdx := 0
	endIdx := n - 1

	for beginIdx <= endIdx {
		newBeginIdx := endIdx
		newEndIdx := beginIdx

		for i := beginIdx; i < endIdx; i++ {
			update <- i

			if arr.Less(i+1, i) {
				mutex.Lock()
				arr.Swap(i, i+1)
				mutex.Unlock()

				newEndIdx = i
			}
		}

		endIdx = newEndIdx

		for i := endIdx - 1; i >= beginIdx; i-- {
			update <- i + 1

			if arr.Less(i+1, i) {
				mutex.Lock()
				arr.Swap(i, i+1)
				mutex.Unlock()

				newBeginIdx = i
			}
		}

		beginIdx = newBeginIdx
	}
}
