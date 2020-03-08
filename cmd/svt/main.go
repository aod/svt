package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aod/svt/internal/flags"
	"github.com/aod/svt/internal/visualizer"
	"github.com/gdamore/tcell"
)

var config visualizer.Config

var printAlgorithms = false

func init() {
	flag.IntVar(&config.ArraySize, "a", 12, "Array size")
	flag.DurationVar(&config.Delay, "d", time.Millisecond*16, "Delay between sorts")
	flag.IntVar(&config.ColumnThiccness, "t", 4, "Column thiccness")
	flags.AlgorithmVar(&config.Algorithm, "s", "bubble")
	flag.BoolVar(&config.QuitWhenDone, "q", false, "Automatically quit after it's done sorting")
	config.Style = tcell.StyleDefault

	flag.BoolVar(&printAlgorithms, "algorithms", false, "Print out all available sorting algorithms")

	flag.Parse()
}

func main() {
	if printAlgorithms {
		for _, v := range flags.Algorithms {
			fmt.Fprintln(os.Stdout, v)
		}
		return
	}

	v := visualizer.Make(config)

	if err := v.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	v.Visualize()
}
