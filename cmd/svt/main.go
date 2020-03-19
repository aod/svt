package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aod/svt/internal/flags"
	"github.com/aod/svt/internal/visualizer"
	"github.com/gdamore/tcell"
)

func main() {
	config := visualizer.Config{}
	flag.IntVar(&config.ArraySize, "a", 12, "Array size")
	flag.DurationVar(&config.Delay, "d", time.Millisecond*16, "Delay between sorts")
	flag.IntVar(&config.ColumnThiccness, "t", 4, "Column thiccness")
	flag.BoolVar(&config.QuitWhenDone, "q", false, "Automatically quit after it's done sorting")

	flags.AlgorithmVar(&config.Algorithm, "s")
	config.Style = tcell.StyleDefault

	printAlgorithms := flag.Bool("algorithms", false, "Print out all available sorting algorithms")

	flag.Parse()

	if *printAlgorithms {
		for _, v := range flags.Algorithms() {
			fmt.Println(v)
		}
		return
	}

	v := visualizer.Make(config)
	if err := v.Init(); err != nil {
		log.Fatalln(err)
	}
	v.Visualize()
}
