package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aod/svt/internal/flags"
	"github.com/aod/svt/pkg/visualizer"
	"github.com/gdamore/tcell"
)

var config visualizer.Config

func init() {
	flag.IntVar(&config.ArraySize, "a", 12, "Array size")
	flag.DurationVar(&config.Delay, "d", time.Millisecond*16, "Delay between sorts")
	flag.IntVar(&config.ColumnThiccness, "t", 4, "Column thiccness")
	flags.AlgorithmVar(&config.Algorithm, "s", "bubble")
	config.Style = tcell.StyleDefault.Foreground(tcell.ColorGhostWhite)

	flag.Parse()
}

func main() {
	v := visualizer.Make(config)

	if err := v.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	v.Visualize()
}
