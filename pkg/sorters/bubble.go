package sorters

import (
	"sort"
	"sync"
)

func Bubble(arr sort.Interface, update chan<- int, mutex *sync.Mutex) {
	n := arr.Len()

	for n > 1 {
		newN := 0
		for i := 1; i <= n-1; i++ {
			update <- i - 1

			if arr.Less(i, i-1) {
				mutex.Lock()
				arr.Swap(i, i-1)
				mutex.Unlock()

				newN = i
			}
		}

		n = newN
	}
}
