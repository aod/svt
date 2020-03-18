package visualizer

import (
	"errors"
	"time"

	"github.com/aod/svt/pkg/sorters"
)

type state interface {
	handle(*Visualizer) error
}

type initialState struct{}

func (s *initialState) handle(v *Visualizer) error {
	select {
	case <-v.quit:
		return errors.New("quitting program")
	case comparison := <-v.update:
		v.draw(comparison)
	}
	v.state = &sortingState{}
	return nil
}

type sortingState struct{}

func (s *sortingState) handle(v *Visualizer) error {
	for {
		select {
		case <-v.quit:
			return errors.New("quitting program")
		case <-time.After(v.Config.Delay):
			if comparison, ok := <-v.update; ok {
				v.draw(comparison)
			} else {
				v.state = &highlightState{}
				return nil
			}
		}
	}
}

type highlightState struct{}

func (s *highlightState) handle(v *Visualizer) error {
	for i := 0; i < len(v.Array); i++ {
		select {
		case <-v.quit:
			return errors.New("quitting program")
		case <-time.After(time.Second / 60):
			v.visualize(sorters.Compare{
				Indexes: [2]int{i, i},
				Swapped: false,
			})
		}
	}
	v.state = &endState{}
	return nil
}

type endState struct{}

func (s *endState) handle(v *Visualizer) error {
	if v.Config.QuitWhenDone {
		<-time.After(time.Second / 10)
	} else {
		<-v.quit
	}
	return errors.New("reached end state")
}
