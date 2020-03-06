package sorters

import "sync"

func Selection(arr []int, update chan<- int, mutex *sync.Mutex) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			update <- j
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		mutex.Lock()
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
		mutex.Unlock()
	}
}
