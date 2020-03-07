// Package sorters implements steppable sorting algorithms functions
package sorters

import (
	"sort"
	"sync"
)

// A Stepped function takes an sort.Interface and sorts it whilst updating the
// channel for each comparison and uses the mutex when swapping elements.
type Stepped func(sort.Interface, chan<- int, *sync.Mutex)
