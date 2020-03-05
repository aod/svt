package sorters

import "sync"

func Cocktail(arr []int, update chan<- int, mutex *sync.Mutex) {
	n := len(arr)
	beginIdx := 0
	endIdx := n - 1

	for beginIdx <= endIdx {
		newBeginIdx := endIdx
		newEndIdx := beginIdx

		for i := beginIdx; i < endIdx; i++ {
			update <- i

			if arr[i] > arr[i+1] {
				mutex.Lock()
				arr[i], arr[i+1] = arr[i+1], arr[i]
				mutex.Unlock()

				newEndIdx = i
			}
		}

		endIdx = newEndIdx

		for i := endIdx - 1; i >= beginIdx; i-- {
			update <- i + 1

			if arr[i] > arr[i+1] {
				mutex.Lock()
				arr[i], arr[i+1] = arr[i+1], arr[i]
				mutex.Unlock()

				newBeginIdx = i
			}
		}

		beginIdx = newBeginIdx
	}
}
