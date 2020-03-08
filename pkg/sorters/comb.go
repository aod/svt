package sorters

import (
	"sort"
	"sync"
)

func Comb(arr sort.Interface, update chan<- Compare, mutex *sync.Mutex) {
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
			c := Compare{
				Indexes: [2]int{i, i + gap},
				Swapped: false,
			}

			if arr.Less(i+gap, i) {
				mutex.Lock()
				arr.Swap(i+gap, i)
				mutex.Unlock()

				c.Swapped = true
				sorted = false
			}

			update <- c
		}
	}
}
