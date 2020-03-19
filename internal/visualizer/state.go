package visualizer

import (
	"time"

	"github.com/aod/svt/pkg/sorters"
)

type state interface {
	handle(*Visualizer)
}

type initialState struct{}

func (s *initialState) handle(v *Visualizer) {
	select {
	case <-v.quit:
		v.state = nil
		return
	case comparison := <-v.update:
		v.draw(comparison)
	}
	v.state = &sortingState{}
}

type sortingState struct{}

func (s *sortingState) handle(v *Visualizer) {
	for {
		select {
		case <-v.quit:
			v.state = nil
			return
		case <-time.After(v.Config.Delay):
			if comparison, ok := <-v.update; ok {
				v.draw(comparison)
			} else {
				v.state = &highlightState{}
				return
			}
		}
	}
}

type highlightState struct{}

func (s *highlightState) handle(v *Visualizer) {
	for i := 0; i < len(v.Array); i++ {
		select {
		case <-v.quit:
			v.state = nil
			return
		case <-time.After(time.Second / 60):
			v.visualize(sorters.Compare{
				Indexes: [2]int{i, i},
				Swapped: false,
			})
		}
	}
	v.state = &doneState{}
}

type doneState struct{}

func (s *doneState) handle(v *Visualizer) {
	if v.Config.QuitWhenDone {
		<-time.After(time.Second / 10)
	} else {
		<-v.quit
	}
	v.state = nil
}
