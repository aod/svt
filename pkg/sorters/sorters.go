// Package sorters implements steppable sorting algorithms functions
package sorters

import (
	"sort"
	"sync"
)

// Compare is the comparison between 2 indexes and states if they have been
// swapped or not.
type Compare struct {
	Indexes [2]int
	Swapped bool
}

// A Stepped function takes an sort.Interface and sorts it whilst updating the
// channel for each comparison and uses the mutex when swapping elements.
type Stepped func(sort.Interface, chan<- Compare, *sync.Mutex)
