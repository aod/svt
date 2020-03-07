package sorters

import "sync"
import "math/rand"

func Bogo(arr []int, update chan<- int, mutex *sync.Mutex) {
	n := len(arr)

	for {
		rand.Shuffle(n, func(i, j int) {
			mutex.Lock()
			arr[i], arr[j] = arr[j], arr[i]
			mutex.Unlock()
			update <- i
			update <- j
		})

		sorted := true
		for i := 0; i < n - 1; i++ {
			if arr[i] > arr[i + 1] {
				sorted = false
			}
		}

		if sorted == true {
			break
		}
	}
}