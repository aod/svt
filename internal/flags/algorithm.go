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

var Algorithms = []string{
	"bubble",
	"bogo",
	"cocktail",
	"comb",
	"selection",
}

func (v *AlgorithmValue) Usage() string {
	return fmt.Sprintf("Sorting algorithm, choose from:\n%s or %s",
		strings.Join(Algorithms[:len(Algorithms)-1], ", "),
		Algorithms[len(Algorithms)-1])
}

var algorithmsTable = map[string]sorters.Stepped{
	"bubble":    sorters.Bubble,
	"bogo":      sorters.Bogo,
	"cocktail":  sorters.Cocktail,
	"comb":      sorters.Comb,
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
