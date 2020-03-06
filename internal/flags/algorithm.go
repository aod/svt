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

func (v *AlgorithmValue) Usage() string {
	var algorithms []string
	for name := range v.Algorithms {
		algorithms = append(algorithms, "- "+name)
	}
	return fmt.Sprintf("Sorting algorithm. Possible values are:\n\t%s\n", strings.Join(algorithms, "\n\t"))
}

var algorithms = map[string]sorters.Stepped{
	"bubble":    sorters.Bubble,
	"cocktail":  sorters.Cocktail,
	"selection": sorters.Selection,
}

func AlgorithmVar(sorter *sorters.Stepped, name, value string) {
	*sorter = algorithms[value]

	a := &AlgorithmValue{
		Value:      value,
		Algorithm:  sorter,
		Algorithms: algorithms,
	}
	flag.Var(a, name, a.Usage())
}
