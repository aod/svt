// Package sorters implements steppable sorting algorithms functions
package sorters

import "sync"

// A Stepped function takes an int slice and sorts it whilst updating the
// channel for each comparison and uses the mutex when swapping elements.
type Stepped func([]int, chan<- int, *sync.Mutex)
