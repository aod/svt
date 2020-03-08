package sorters

import (
	"sort"
	"sync"
)

func Comb(arr sort.Interface, update chan<- int, mutex *sync.Mutex) {
	n := arr.Len()
	gap := n
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			sorted = true
			gap = 1
		}

		for i := 0; i+gap < n; i++ {
			update <- i
			update <- i + gap

			if arr.Less(i+gap, i) {
				mutex.Lock()
				arr.Swap(i+gap, i)
				mutex.Unlock()

				sorted = false
			}
		}
	}
}
