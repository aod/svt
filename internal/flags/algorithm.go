package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aod/svt/pkg/sorters"
)

type AlgorithmValue struct {
	Value      string
	Algorithm  *sorters.Stepped
	Algorithms map[string]sorters.Stepped
}

func (v *AlgorithmValue) String() string {
	return v.Value
}

func (v *AlgorithmValue) Set(key string) error {
	algo, ok := v.Algorithms[key]
	if !ok {
		return fmt.Errorf("Algorithm does not exist")
	}
	v.Value = key
	*v.Algorithm = algo
	return nil
}

var algorithms = []string{
	"bubble",
	"cocktail",
	"selection",
}

func (v *AlgorithmValue) Usage() string {
	return fmt.Sprintf("Sorting algorithm, choose from:\n%s or %s",
		strings.Join(algorithms[:len(algorithms)-1], ", "),
		algorithms[len(algorithms)-1])
}

var algorithmsTable = map[string]sorters.Stepped{
	"bubble":    sorters.Bubble,
	"cocktail":  sorters.Cocktail,
	"selection": sorters.Selection,
}

func AlgorithmVar(sorter *sorters.Stepped, name, value string) {
	*sorter = algorithmsTable[value]

	a := &AlgorithmValue{
		Value:      value,
		Algorithm:  sorter,
		Algorithms: algorithmsTable,
	}
	flag.Var(a, name, a.Usage())
}
