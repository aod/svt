package sorters

import (
	"math/rand"
	"sort"
	"sync"
)

func Bogo(arr sort.Interface, update chan<- Compare, mutex *sync.Mutex) {
	n := arr.Len()

	for !sort.IsSorted(arr) {
		rand.Shuffle(n, func(i, j int) {
			mutex.Lock()
			arr.Swap(i, j)
			mutex.Unlock()

			update <- Compare{
				Indexes: [2]int{i, j},
				Swapped: true,
			}
		})
	}
}
