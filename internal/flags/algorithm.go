package flags

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/aod/svt/pkg/sorters"
)

var (
	algorithmsTable = map[string]sorters.Stepped{
		"bubble":    sorters.Bubble,
		"bogo":      sorters.Bogo,
		"cocktail":  sorters.Cocktail,
		"comb":      sorters.Comb,
		"selection": sorters.Selection,
	}
	algorithms []string
)

func init() {
	for k := range algorithmsTable {
		algorithms = append(algorithms, k)
	}
	sort.Strings(algorithms)
}

func Algorithms() []string {
	a := make([]string, len(algorithms))
	copy(a, algorithms)
	return a
}

type algorithmValue struct {
	value      string
	algorithm  *sorters.Stepped
	algorithms map[string]sorters.Stepped
	usage      string
}

func (v *algorithmValue) String() string {
	return v.value
}

func (v *algorithmValue) Set(key string) error {
	algo, ok := algorithmsTable[key]
	if !ok {
		return fmt.Errorf("algorithm does not exist")
	}
	*v.algorithm = algo
	return nil
}

func (v *algorithmValue) Usage() string {
	return fmt.Sprintf("Sorting `algorithm`. Choose from: %s or %s",
		strings.Join(algorithms[:len(algorithms)-1], ", "),
		algorithms[len(algorithms)-1])
}

func AlgorithmVar(sorter *sorters.Stepped, name string) {
	*sorter = sorters.Bubble
	a := &algorithmValue{
		value:     "bubble",
		algorithm: sorter,
	}
	flag.Var(a, name, a.Usage())
}
