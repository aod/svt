package sorters

import (
	"math/rand"
	"sort"
	"sync"
)

func Bogo(arr sort.Interface, update chan<- int, mutex *sync.Mutex) {
	n := arr.Len()

	for !sort.IsSorted(arr) {
		rand.Shuffle(n, func(i, j int) {
			mutex.Lock()
			arr.Swap(i, j)
			mutex.Unlock()
			update <- i
			update <- j
		})
	}
}
