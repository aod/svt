package visualizer

import (
	"time"

	"github.com/aod/svt/pkg/sorters"
	"github.com/gdamore/tcell"
)

type Config struct {
	ArraySize       int
	Delay           time.Duration
	Algorithm       sorters.Stepped
	ColumnThiccness int
	Style           tcell.Style
	QuitWhenDone    bool
}
