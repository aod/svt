package visualizer

import (
	"math/rand"
	"sync"
	"time"

	"github.com/aod/svt/pkg/sorters"
	"github.com/gdamore/tcell"
)

type Visualizer struct {
	Config        Config
	Array         []int
	Width, Height int
	Screen        tcell.Screen

	quit   chan struct{}
	update chan sorters.Compare
	mutex  *sync.Mutex
	state  state
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
		v.Config.Algorithm(v, v.update, v.mutex)
	}()

	v.refreshDimensions()
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
	for v.state != nil {
		v.state.handle(v)
	}
}

func (v *Visualizer) draw(c sorters.Compare) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.visualize(c)
}

func (v *Visualizer) visualize(c sorters.Compare) {
	compareColor, swapColor := v.colors()
	comparedStyle := v.Config.Style.Foreground(compareColor)
	swappedStyle := v.Config.Style.Foreground(swapColor)

	if !c.Swapped {
		v.Screen.Clear()
		defer v.Screen.Show()

		for i := range v.Array {
			v.drawColumnForIndex(i, v.Config.Style)
		}
		v.drawColumnForIndex(c.Indexes[0], comparedStyle)
		v.drawColumnForIndex(c.Indexes[1], comparedStyle)
	} else {
		v.Screen.Clear()

		i1 := c.Indexes[0]
		i2 := c.Indexes[1]

		// Show the array first as if it were not sorted
		for i := range v.Array {
			switch {
			case i == i1:
				v.drawColumn(i, v.Array[i2], v.Config.Style)
			case i == i2:
				v.drawColumn(i, v.Array[i1], v.Config.Style)
			default:
				v.drawColumnForIndex(i, v.Config.Style)
			}
		}
		v.drawColumn(i1, v.Array[i2], comparedStyle)
		v.drawColumn(i2, v.Array[i1], comparedStyle)
		v.Screen.Show()

		<-time.After(v.Config.Delay)
		v.Screen.Clear()
		defer v.Screen.Show()

		for i := range v.Array {
			v.drawColumnForIndex(i, v.Config.Style)
		}
		v.drawColumnForIndex(i1, swappedStyle)
		v.drawColumnForIndex(i2, swappedStyle)
	}
}

func (v *Visualizer) drawColumnForIndex(idx int, style tcell.Style) {
	x, y := v.position()

	x += idx * v.Config.ColumnThiccness
	value := v.Array[idx]

	for i := 0; i <= value; i++ {
		for j := 0; j < v.Config.ColumnThiccness; j++ {
			v.Screen.SetContent(x+j, y-i, '█', nil, style)
		}
	}
}

func (v *Visualizer) drawColumn(idx, height int, style tcell.Style) {
	x, y := v.position()

	x += idx * v.Config.ColumnThiccness

	for i := 0; i <= height; i++ {
		for j := 0; j < v.Config.ColumnThiccness; j++ {
			v.Screen.SetContent(x+j, y-i, '█', nil, style)
		}
	}
}

func (v *Visualizer) colors() (compareColor, swapColor tcell.Color) {
	return tcell.ColorMediumPurple, tcell.ColorLightGreen
}

func (v *Visualizer) position() (x, y int) {
	return v.Width/2 - v.Config.ArraySize*v.Config.ColumnThiccness/2, v.Height/2 + v.Config.ArraySize/2
}

func (v *Visualizer) refreshDimensions() {
	v.Width, v.Height = v.Screen.Size()
}

func (v *Visualizer) Len() int           { return len(v.Array) }
func (v *Visualizer) Swap(i, j int)      { v.Array[i], v.Array[j] = v.Array[j], v.Array[i] }
func (v *Visualizer) Less(i, j int) bool { return v.Array[i] < v.Array[j] }

func Make(c Config) *Visualizer {
	v := &Visualizer{}
	v.Config = c

	v.quit = make(chan struct{})
	v.update = make(chan sorters.Compare)
	v.mutex = &sync.Mutex{}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	v.Array = random.Perm(v.Config.ArraySize)

	v.state = &initialState{}

	return v
}
