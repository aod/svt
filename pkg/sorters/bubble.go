package sorters

import "sync"

func Bubble(arr []int, update chan<- int, mutex *sync.Mutex) {
	n := len(arr)

	for n > 1 {
		newN := 0
		for i := 1; i <= n-1; i++ {
			update <- i - 1

			if arr[i-1] > arr[i] {
				mutex.Lock()
				arr[i-1], arr[i] = arr[i], arr[i-1]
				mutex.Unlock()

				newN = i
			}
		}

		n = newN
	}
}
