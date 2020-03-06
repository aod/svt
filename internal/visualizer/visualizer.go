package visualizer

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gdamore/tcell"
)

type Visualizer struct {
	Config        Config
	Array         []int
	Width, Height int
	Screen        tcell.Screen

	quit   chan struct{}
	update chan int
	mutex  *sync.Mutex
}

func (v *Visualizer) Init() error {
	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err = screen.Init(); err != nil {
		return err
	}
	v.Screen = screen
	return nil
}

func (v *Visualizer) close() {
	v.Screen.Fini()
}

func (v *Visualizer) Visualize() {
	defer v.close()
	go v.handleEvents()

	go func() {
		defer close(v.update)
		v.Config.Algorithm(v.Array, v.update, v.mutex)
	}()

	v.mainLoop()
}

func (v *Visualizer) handleEvents() {
	defer close(v.quit)

	for {
		polledEvent := v.Screen.PollEvent()

		switch event := polledEvent.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyRune:
				switch event.Rune() {
				case 'q', 'Q':
					return
				}
			}
		case *tcell.EventResize:
			v.refreshDimensions()
			v.Screen.Sync()
		}
	}
}

func (v *Visualizer) mainLoop() {
	v.refreshDimensions()

	// Initial draw to screen
	select {
	case <-v.quit:
		return
	case updateIdx := <-v.update:
		v.draw(updateIdx)
	}

	// Main sorting & drawing loop
mainLoop:
	for {
		select {
		case <-v.quit:
			return
		case <-time.After(v.Config.Delay):
			if updateIdx, ok := <-v.update; ok {
				v.draw(updateIdx)
			} else {
				break mainLoop
			}
		}
	}

	// Final visualization. Show sorted array, highlight each element after another
	for i := 0; i < len(v.Array); i++ {
		select {
		case <-v.quit:
			return
		case <-time.After(time.Second / 60):
			v.drawVisualization(i)
		}
	}

	<-v.quit
}

func (v *Visualizer) draw(updateIdx int) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.drawVisualization(updateIdx)
}

func (v *Visualizer) drawVisualization(updateIdx int) {
	v.Screen.Clear()

	x := v.Width/2 - v.Config.ArraySize*v.Config.ColumnThiccness/2
	y := v.Height/2 + v.Config.ArraySize/2

	for idx, value := range v.Array {
		style := tcell.StyleDefault
		if updateIdx == idx {
			style = style.Foreground(tcell.ColorMediumPurple)
		}

		for i := 0; i <= value; i++ {
			for j := 0; j < v.Config.ColumnThiccness; j++ {
				v.Screen.SetContent(x+j, y-i, '█', nil, style)
			}
		}
		x += v.Config.ColumnThiccness
	}

	v.Screen.Show()
}

func (v *Visualizer) refreshDimensions() {
	v.Width, v.Height = v.Screen.Size()
}

func Make(c Config) *Visualizer {
	v := &Visualizer{}
	v.Config = c

	v.quit = make(chan struct{})
	v.update = make(chan int)
	v.mutex = &sync.Mutex{}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	v.Array = random.Perm(v.Config.ArraySize)

	return v
}
